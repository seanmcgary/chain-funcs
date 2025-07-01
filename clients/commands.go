package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/seanmcgary/fass-avs/clients/pkg/bindings/FaaS"
	"github.com/urfave/cli/v2"
)

func registerFunction(c *cli.Context) error {
	if c.NArg() != 2 {
		return fmt.Errorf("expected exactly two arguments: function name and function directory path")
	}

	functionName := c.Args().Get(0)
	functionDir := c.Args().Get(1)
	rpcURL := c.String("rpc-url")
	privateKeyHex := c.String("private-key")
	faasAddress := c.String("faas-address")
	skipDeps := c.Bool("skip-deps")

	if functionName == "" {
		return fmt.Errorf("function name cannot be empty")
	}

	// Connect to Ethereum client
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}
	defer client.Close()

	// Get chain ID first
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Load private key
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return fmt.Errorf("failed to load private key: %w", err)
	}

	// Create transactor with chain ID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %w", err)
	}
	auth.GasLimit = uint64(2000000) // Set gas limit

	// Get suggested gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get suggested gas price: %w", err)
	}
	auth.GasPrice = gasPrice

	// Create FaaS contract instance
	faas, err := FaaS.NewFaaS(common.HexToAddress(faasAddress), client)
	if err != nil {
		return fmt.Errorf("failed to create FaaS contract instance: %w", err)
	}

	// Check if this is a Python function and handle dependencies
	isPythonFunction, err := checkIsPythonFunction(functionDir)
	if err != nil {
		return fmt.Errorf("failed to check function type: %w", err)
	}

	var finalFunctionDir string
	if isPythonFunction && !skipDeps {
		// Create a temporary directory with dependencies installed
		finalFunctionDir, err = preparePythonFunction(functionDir)
		if err != nil {
			return fmt.Errorf("failed to prepare Python function: %w", err)
		}
		defer os.RemoveAll(finalFunctionDir)
	} else {
		finalFunctionDir = functionDir
		if isPythonFunction && skipDeps {
			fmt.Printf("Skipping Python dependency installation (--skip-deps flag set)\n")
		}
	}

	// Create tarball from function directory (with dependencies if Python)
	tarballData, err := createTarball(finalFunctionDir)
	if err != nil {
		return fmt.Errorf("failed to create tarball: %w", err)
	}

	// Calculate function ID using current contract method (content hash only)
	functionID := crypto.Keccak256Hash(tarballData)
	fmt.Printf("Registering function '%s' from directory: %s\n", functionName, functionDir)
	fmt.Printf("Tarball size: %d bytes\n", len(tarballData))
	fmt.Printf("Function ID: %s\n", functionID.Hex())
	fmt.Printf("Using chain ID: %s\n", chainID.String())

	// Register function with embedded content (on-chain storage) - using old interface for now
	tx, err := faas.RegisterFunction(auth, tarballData)
	if err != nil {
		return fmt.Errorf("failed to register function: %w", err)
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	fmt.Printf("Waiting for confirmation...\n")

	// Wait for transaction receipt
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction: %w", err)
	}

	if receipt.Status != 1 {
		return fmt.Errorf("transaction failed")
	}

	fmt.Printf("Function registered successfully!\n")
	fmt.Printf("Function ID: %s\n", functionID.Hex())
	fmt.Printf("Gas used: %d\n", receipt.GasUsed)
	if receipt.EffectiveGasPrice != nil {
		totalCost := new(big.Int).Mul(big.NewInt(int64(receipt.GasUsed)), receipt.EffectiveGasPrice)
		fmt.Printf("Gas price: %s gwei\n", formatGasPrice(receipt.EffectiveGasPrice))
		fmt.Printf("Transaction cost: %s ETH\n", formatWei(totalCost))
	}

	return nil
}

func deployFunction(c *cli.Context) error {
	if c.NArg() != 2 {
		return fmt.Errorf("expected exactly two arguments: function name and function directory path")
	}

	functionName := c.Args().Get(0)
	functionDir := c.Args().Get(1)
	rpcURL := c.String("rpc-url")
	privateKeyHex := c.String("private-key")
	faasAddress := c.String("faas-address")
	s3BaseURL := c.String("s3-base-url")
	skipDeps := c.Bool("skip-deps")

	if functionName == "" {
		return fmt.Errorf("function name cannot be empty")
	}

	if s3BaseURL == "" {
		return fmt.Errorf("s3-base-url is required for deploy-function command. Use register-function for on-chain storage.")
	}

	// Connect to Ethereum client
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}
	defer client.Close()

	// Get chain ID first
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Load private key
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return fmt.Errorf("failed to load private key: %w", err)
	}

	// Create transactor with chain ID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %w", err)
	}
	auth.GasLimit = uint64(2000000) // Set gas limit

	// Get suggested gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get suggested gas price: %w", err)
	}
	auth.GasPrice = gasPrice

	// Create FaaS contract instance
	faas, err := FaaS.NewFaaS(common.HexToAddress(faasAddress), client)
	if err != nil {
		return fmt.Errorf("failed to create FaaS contract instance: %w", err)
	}

	// Check if this is a Python function and handle dependencies
	isPythonFunction, err := checkIsPythonFunction(functionDir)
	if err != nil {
		return fmt.Errorf("failed to check function type: %w", err)
	}

	var finalFunctionDir string
	if isPythonFunction && !skipDeps {
		// Create a temporary directory with dependencies installed
		finalFunctionDir, err = preparePythonFunction(functionDir)
		if err != nil {
			return fmt.Errorf("failed to prepare Python function: %w", err)
		}
		defer os.RemoveAll(finalFunctionDir)
	} else {
		finalFunctionDir = functionDir
		if isPythonFunction && skipDeps {
			fmt.Printf("Skipping Python dependency installation (--skip-deps flag set)\n")
		}
	}

	// Create tarball from function directory (with dependencies if Python)
	tarballData, err := createTarball(finalFunctionDir)
	if err != nil {
		return fmt.Errorf("failed to create tarball: %w", err)
	}

	// Calculate function ID using current contract method (content hash only)
	functionID := crypto.Keccak256Hash(tarballData)
	
	fmt.Printf("Deploying function '%s' from directory: %s\n", functionName, functionDir)
	fmt.Printf("Tarball size: %d bytes\n", len(tarballData))
	fmt.Printf("Function ID: %s\n", functionID.Hex())
	fmt.Printf("Using chain ID: %s\n", chainID.String())

	// Upload to S3
	s3URL, err := uploadToS3(tarballData, functionID, s3BaseURL)
	if err != nil {
		return fmt.Errorf("failed to upload to S3: %w", err)
	}

	fmt.Printf("Uploaded to S3: %s\n", s3URL)

	// Call deployFunction with URL - using old interface for now
	tx, err := faas.DeployFunction(auth, s3URL, functionID)
	if err != nil {
		return fmt.Errorf("failed to deploy function: %w", err)
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	fmt.Printf("Waiting for confirmation...\n")

	// Wait for transaction receipt
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction: %w", err)
	}

	if receipt.Status != 1 {
		return fmt.Errorf("transaction failed")
	}

	fmt.Printf("Function deployed successfully!\n")
	fmt.Printf("Function ID: %s\n", functionID.Hex())
	fmt.Printf("Gas used: %d\n", receipt.GasUsed)
	if receipt.EffectiveGasPrice != nil {
		totalCost := new(big.Int).Mul(big.NewInt(int64(receipt.GasUsed)), receipt.EffectiveGasPrice)
		fmt.Printf("Gas price: %s gwei\n", formatGasPrice(receipt.EffectiveGasPrice))
		fmt.Printf("Transaction cost: %s ETH\n", formatWei(totalCost))
	}

	return nil
}

func callFunction(c *cli.Context) error {
	if c.NArg() != 1 {
		return fmt.Errorf("expected exactly one argument: input string")
	}

	inputString := c.Args().Get(0)
	rpcURL := c.String("rpc-url")
	wsRPCURL := c.String("ws-rpc-url")
	privateKeyHex := c.String("private-key")
	faasAddress := c.String("faas-address")
	taskMailboxAddress := c.String("taskMailbox-address")
	functionIDHex := c.String("function-id")

	// Connect to Ethereum client
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}
	defer client.Close()

	// Get chain ID first
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Load private key
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return fmt.Errorf("failed to load private key: %w", err)
	}

	// Create transactor with chain ID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %w", err)
	}

	// Set gas parameters
	auth.GasLimit = uint64(2000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get suggested gas price: %w", err)
	}
	auth.GasPrice = gasPrice

	// Parse function ID
	functionID := common.HexToHash(functionIDHex)

	// Create FaaS contract instance
	faas, err := FaaS.NewFaaS(common.HexToAddress(faasAddress), client)
	if err != nil {
		return fmt.Errorf("failed to create FaaS contract instance: %w", err)
	}

	// Prepare arguments as JSON array containing the string
	args := []string{inputString}
	argsJSON, err := json.Marshal(args)
	if err != nil {
		return fmt.Errorf("failed to marshal arguments: %w", err)
	}

	fmt.Printf("Calling function: %s\n", functionID.Hex())
	fmt.Printf("Arguments: %s\n", string(argsJSON))

	// Call the function
	tx, err := faas.CallFunction(auth, functionID, argsJSON)
	if err != nil {
		return fmt.Errorf("failed to call function: %w", err)
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
	fmt.Printf("Waiting for confirmation...\n")

	// Wait for transaction receipt
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction: %w", err)
	}

	if receipt.Status != 1 {
		return fmt.Errorf("transaction failed")
	}

	// Extract task ID from logs
	var taskID [32]byte
	faasFilterer, err := FaaS.NewFaaSFilterer(common.HexToAddress(faasAddress), client)
	if err != nil {
		return fmt.Errorf("failed to create FaaS filterer: %w", err)
	}

	// Parse logs to find the FunctionCalled event
	fmt.Printf("Logs: %+v\n", receipt)
	for _, log := range receipt.Logs {
		if len(log.Topics) > 0 && log.Address == common.HexToAddress(faasAddress) {
			event, err := faasFilterer.ParseFunctionCalled(*log)
			fmt.Printf("Event: %+v\n", event)
			if err == nil && event.FunctionId == functionID {
				taskID = event.TaskId
				break
			}
		}
	}

	if taskID == ([32]byte{}) {
		return fmt.Errorf("could not find task ID in transaction logs")
	}

	fmt.Printf("Task ID: %s\n", common.BytesToHash(taskID[:]).Hex())
	fmt.Printf("Gas used: %d\n", receipt.GasUsed)
	if receipt.EffectiveGasPrice != nil {
		totalCost := new(big.Int).Mul(big.NewInt(int64(receipt.GasUsed)), receipt.EffectiveGasPrice)
		fmt.Printf("Gas price: %s gwei\n", formatGasPrice(receipt.EffectiveGasPrice))
		fmt.Printf("Transaction cost: %s ETH\n", formatWei(totalCost))
	}
	fmt.Printf("Listening for TaskVerified event...\n")

	// Subscribe to TaskVerified events using WebSocket client
	return subscribeToTaskVerified(wsRPCURL, taskMailboxAddress, taskID)
}

func listFunctions(c *cli.Context) error {
	return fmt.Errorf("list-functions command requires updated contract with pagination support. This feature will be available after contract upgrade.")
}

func getFunctionInfo(c *cli.Context) error {
	if c.NArg() != 1 {
		return fmt.Errorf("expected exactly one argument: function ID")
	}

	functionIDHex := c.Args().Get(0)
	rpcURL := c.String("rpc-url")
	faasAddress := c.String("faas-address")

	// Connect to Ethereum client
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}
	defer client.Close()

	// Create FaaS contract instance
	faas, err := FaaS.NewFaaS(common.HexToAddress(faasAddress), client)
	if err != nil {
		return fmt.Errorf("failed to create FaaS contract instance: %w", err)
	}

	functionID := common.HexToHash(functionIDHex)

	// Get function metadata using existing methods
	metadata, err := faas.GetFunctionMetadata(nil, functionID)
	if err != nil {
		return fmt.Errorf("failed to get function metadata: %w", err)
	}

	if !metadata.HasContent && !metadata.HasUrl {
		return fmt.Errorf("function not found")
	}

	// Display function information (limited by current contract interface)
	fmt.Printf("Function Information:\n")
	fmt.Printf("  Function ID: %s\n", functionID.Hex())
	fmt.Printf("  Name: Not available (requires contract upgrade)\n")
	fmt.Printf("  Registrar: Not available (requires contract upgrade)\n")
	fmt.Printf("  Deployed: Not available (requires contract upgrade)\n")

	if metadata.HasContent {
		fmt.Printf("  Storage: On-chain content\n")
		fmt.Printf("  Size: %d bytes\n", metadata.ContentLength)
	}

	if metadata.HasUrl {
		url, err := faas.GetFunctionUrl(nil, functionID)
		if err != nil {
			return fmt.Errorf("failed to get function URL: %w", err)
		}
		fmt.Printf("  Storage: Remote URL\n")
		fmt.Printf("  URL: %s\n", url)
	}

	return nil
}

func createTarball(sourceDir string) ([]byte, error) {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	tarWriter := tar.NewWriter(gzipWriter)

	// Walk through the directory
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate relative path
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		// Skip the root directory itself
		if relPath == "." {
			return nil
		}

		// Create tar header
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		header.Name = relPath

		// Write header
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// If it's a file, write its content
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := io.Copy(tarWriter, file); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Close writers
	if err := tarWriter.Close(); err != nil {
		return nil, err
	}
	if err := gzipWriter.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func subscribeToTaskVerified(wsRPCURL, taskMailboxAddress string, taskID [32]byte) error {
	// Create WebSocket client for event subscriptions
	wsClient, err := ethclient.Dial(wsRPCURL)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket RPC: %w", err)
	}
	defer wsClient.Close()

	fmt.Printf("Connected to WebSocket RPC: %s\n", wsRPCURL)
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Set up filter query for TaskVerified events
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(taskMailboxAddress)},
		Topics: [][]common.Hash{
			{crypto.Keccak256Hash([]byte("TaskVerified(address,bytes32,address,uint32,bytes)"))}, // Event signature
			{},                              // aggregator (any)
			{common.BytesToHash(taskID[:])}, // taskHash (our specific task)
			{},                              // avs (any)
		},
	}

	// Subscribe to logs using WebSocket client
	logs := make(chan types.Log)
	sub, err := wsClient.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("failed to subscribe to logs: %w", err)
	}
	defer sub.Unsubscribe()

	fmt.Printf("Subscribed to TaskVerified events for task %s...\n", common.BytesToHash(taskID[:]).Hex())

	// Wait for the event or timeout
	for {
		select {
		case err := <-sub.Err():
			return fmt.Errorf("subscription error: %w", err)
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for TaskVerified event")
		case vLog := <-logs:
			// Parse the TaskVerified event
			return handleTaskVerifiedEvent(vLog)
		}
	}
}

func handleTaskVerifiedEvent(vLog types.Log) error {
	fmt.Printf("\nTask completed successfully!\n")
	fmt.Printf("Event found in block %d, transaction %s\n", vLog.BlockNumber, vLog.TxHash.Hex())

	// Extract result from event data
	// The event structure is: TaskVerified(address aggregator, bytes32 taskHash, address avs, uint32 executorOperatorSetId, bytes result)
	// The result is in the last field of the event data
	if len(vLog.Data) == 0 {
		return fmt.Errorf("no data in TaskVerified event")
	}

	// Parse the event data - the result bytes are at the end
	// Skip the first parts (uint32 executorOperatorSetId + dynamic bytes offset)
	// and extract the bytes result
	if len(vLog.Data) < 64 {
		return fmt.Errorf("insufficient data in TaskVerified event")
	}

	// The event data contains:
	// - uint32 executorOperatorSetId (32 bytes, padded)
	// - offset to bytes result (32 bytes)
	// - length of bytes result (32 bytes)
	// - actual bytes result

	// Extract the bytes result starting from offset 64 (skip executorOperatorSetId and offset)
	if len(vLog.Data) < 96 {
		return fmt.Errorf("insufficient data for result length")
	}

	// Get the length of the result (bytes 64-96)
	resultLength := new(big.Int).SetBytes(vLog.Data[64:96]).Uint64()
	if len(vLog.Data) < int(96+resultLength) {
		return fmt.Errorf("insufficient data for result content")
	}

	// Extract the actual result bytes
	resultBytes := vLog.Data[96 : 96+resultLength]

	// Parse the structured result
	var executionResult struct {
		Payload interface{} `json:"payload"`
		Error   *string     `json:"error"`
	}

	if err := json.Unmarshal(resultBytes, &executionResult); err != nil {
		fmt.Printf("Raw Result: %s\n", string(resultBytes))
	} else {
		if executionResult.Error != nil {
			fmt.Printf("Function Error: %s\n", *executionResult.Error)
		} else {
			fmt.Printf("Function Result:\n")
			payloadBytes, _ := json.MarshalIndent(executionResult.Payload, "", "  ")
			fmt.Printf("%s\n", string(payloadBytes))
		}
	}

	return nil
}

// formatGasPrice formats gas price from wei to gwei with proper decimal places
func formatGasPrice(gasPrice *big.Int) string {
	gwei := new(big.Float).Quo(new(big.Float).SetInt(gasPrice), big.NewFloat(1e9))
	return gwei.Text('f', 2)
}

// formatWei formats wei amount to ETH with proper decimal places
func formatWei(wei *big.Int) string {
	eth := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
	return eth.Text('f', 6)
}

// uploadToS3 uploads the tarball to S3 and returns the public URL
func uploadToS3(tarballData []byte, functionID common.Hash, s3BaseURL string) (string, error) {
	// Parse the S3 base URL
	parsedURL, err := url.Parse(s3BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid S3 base URL: %w", err)
	}

	if parsedURL.Scheme != "s3" {
		return "", fmt.Errorf("S3 base URL must start with s3://")
	}

	bucket := parsedURL.Host
	prefix := strings.TrimPrefix(parsedURL.Path, "/")

	// Use function ID as filename - ensure proper path separator
	var key string
	if prefix == "" {
		key = fmt.Sprintf("%s.tar.gz", functionID.Hex())
	} else {
		key = fmt.Sprintf("%s/%s.tar.gz", prefix, functionID.Hex())
	}

	// Create AWS session with automatic credential resolution
	sessOptions := session.Options{
		SharedConfigState: session.SharedConfigEnable, // Enable ~/.aws/config parsing
	}

	if profile := os.Getenv("AWS_PROFILE"); profile != "" {
		fmt.Printf("Using AWS profile: %s\n", profile)
		// AWS_PROFILE env var will be automatically used by the session
	} else {
		fmt.Println("No AWS_PROFILE set, using default credentials")
	}

	sess, err := session.NewSessionWithOptions(sessOptions)
	if err != nil {
		return "", fmt.Errorf("failed to create AWS session: %w", err)
	}

	// Create S3 service client
	s3Client := s3.New(sess)

	// Upload the tarball
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(tarballData),
		ContentType: aws.String("application/gzip"),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload to S3: %w", err)
	}

	// Return the public HTTPS URL
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key), nil
}

// checkIsPythonFunction determines if this is a Python function using manifest.json
func checkIsPythonFunction(functionDir string) (bool, error) {
	manifestPath := filepath.Join(functionDir, "manifest.json")

	// Read manifest.json
	manifestData, err := os.ReadFile(manifestPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("manifest.json not found in function directory")
		}
		return false, err
	}

	// Parse manifest.json
	var manifest struct {
		Runtime string `json:"runtime"`
	}

	if err := json.Unmarshal(manifestData, &manifest); err != nil {
		return false, fmt.Errorf("failed to parse manifest.json: %w", err)
	}

	// Validate runtime
	switch manifest.Runtime {
	case "python":
		return true, nil
	case "javascript":
		return false, nil
	default:
		return false, fmt.Errorf("unsupported runtime in manifest.json: %s (supported: javascript, python)", manifest.Runtime)
	}
}

// preparePythonFunction creates a temporary directory with Python dependencies installed
func preparePythonFunction(sourceDir string) (string, error) {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "python-function-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temp directory: %w", err)
	}

	// Copy all source files to temp directory
	err = copyDir(sourceDir, tempDir)
	if err != nil {
		os.RemoveAll(tempDir)
		return "", fmt.Errorf("failed to copy source files: %w", err)
	}

	// Check if requirements.txt exists
	requirementsPath := filepath.Join(tempDir, "requirements.txt")
	if _, err := os.Stat(requirementsPath); os.IsNotExist(err) {
		// No requirements.txt, return temp dir as-is
		fmt.Printf("No requirements.txt found, packaging function without dependencies\n")
		return tempDir, nil
	}

	fmt.Printf("Installing Python dependencies from requirements.txt...\n")

	// Install dependencies using pip with --target flag
	cmd := exec.Command("pip3", "install", "-r", "requirements.txt", "--target", ".")
	cmd.Dir = tempDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		os.RemoveAll(tempDir)
		return "", fmt.Errorf("failed to install Python dependencies: %w", err)
	}

	fmt.Printf("Python dependencies installed successfully\n")
	return tempDir, nil
}

// copyDir recursively copies a directory
func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate destination path
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		// Copy file
		return copyFile(path, dstPath)
	})
}

// copyFile copies a single file
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
