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
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/seanmcgary/fass-avs/clients/bindings"
	"github.com/urfave/cli/v2"
)

func deployFunction(c *cli.Context) error {
	if c.NArg() != 1 {
		return fmt.Errorf("expected exactly one argument: function directory path")
	}

	functionDir := c.Args().Get(0)
	rpcURL := c.String("rpc-url")
	privateKeyHex := c.String("private-key")
	faasAddress := c.String("faas-address")

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
	faas, err := bindings.NewFaaS(common.HexToAddress(faasAddress), client)
	if err != nil {
		return fmt.Errorf("failed to create FaaS contract instance: %w", err)
	}

	// Create tarball from function directory
	tarballData, err := createTarball(functionDir)
	if err != nil {
		return fmt.Errorf("failed to create tarball: %w", err)
	}

	fmt.Printf("Deploying function from directory: %s\n", functionDir)
	fmt.Printf("Tarball size: %d bytes\n", len(tarballData))
	fmt.Printf("Using chain ID: %s\n", chainID.String())

	// Call registerFunction
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

	// Calculate function ID (keccak256 of content)
	functionID := crypto.Keccak256Hash(tarballData)

	fmt.Printf("Function registered successfully!\n")
	fmt.Printf("Function ID: %s\n", functionID.Hex())
	fmt.Printf("Gas used: %d\n", receipt.GasUsed)

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
	faas, err := bindings.NewFaaS(common.HexToAddress(faasAddress), client)
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
	faasFilterer, err := bindings.NewFaaSFilterer(common.HexToAddress(faasAddress), client)
	if err != nil {
		return fmt.Errorf("failed to create FaaS filterer: %w", err)
	}

	// Parse logs to find the FunctionCalled event
	for _, log := range receipt.Logs {
		if len(log.Topics) > 0 && log.Address == common.HexToAddress(faasAddress) {
			event, err := faasFilterer.ParseFunctionCalled(*log)
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
	fmt.Printf("Listening for TaskVerified event...\n")

	// Subscribe to TaskVerified events using WebSocket client
	return subscribeToTaskVerified(wsRPCURL, taskMailboxAddress, taskID)
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
