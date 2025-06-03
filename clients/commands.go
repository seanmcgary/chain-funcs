package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

	// Create TaskMailbox contract instance
	taskMailbox, err := bindings.NewTaskMailbox(common.HexToAddress(taskMailboxAddress), client)
	if err != nil {
		return fmt.Errorf("failed to create TaskMailbox contract instance: %w", err)
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
	fmt.Printf("Polling for result...\n")

	// Poll for result
	return pollForResult(taskMailbox, taskID)
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

func pollForResult(taskMailbox *bindings.TaskMailbox, taskID [32]byte) error {
	timeout := time.After(5 * time.Minute)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			return fmt.Errorf("timeout waiting for task result")
		case <-ticker.C:
			// Try to get the task result - if it exists, the task is complete
			result, err := taskMailbox.GetTaskResult(nil, taskID)
			if err != nil {
				fmt.Printf("Task still pending... (Error: %v)\n", err)
				continue
			}
			
			// If we got a result without error, the task is complete
			if len(result) > 0 {
				fmt.Printf("Task completed successfully!\n")
				
				// Parse the structured result
				var executionResult struct {
					Payload interface{} `json:"payload"`
					Error   *string     `json:"error"`
				}
				
				if err := json.Unmarshal(result, &executionResult); err != nil {
					fmt.Printf("Raw Result: %s\n", string(result))
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
			
			fmt.Printf("Task still pending...\n")
		}
	}
}