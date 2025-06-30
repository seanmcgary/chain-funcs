package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed contracts/TaskMailbox.json
var taskMailboxJSON []byte

//go:embed contracts/FaaS.json
var faasJSON []byte

type ContractInfo struct {
	Name    string      `json:"name"`
	Address string      `json:"address"`
	ABI     interface{} `json:"abi"`
}

type EmbeddedConfig struct {
	TaskMailbox ContractInfo
	FaaS        ContractInfo
}

func LoadEmbeddedConfig() (*EmbeddedConfig, error) {
	var taskMailbox ContractInfo
	if err := json.Unmarshal(taskMailboxJSON, &taskMailbox); err != nil {
		return nil, fmt.Errorf("failed to parse embedded TaskMailbox config: %w", err)
	}

	var faas ContractInfo
	if err := json.Unmarshal(faasJSON, &faas); err != nil {
		return nil, fmt.Errorf("failed to parse embedded FaaS config: %w", err)
	}

	return &EmbeddedConfig{
		TaskMailbox: taskMailbox,
		FaaS:        faas,
	}, nil
}

func (c *EmbeddedConfig) GetDefaultRPCURL() string {
	// Default RPC URL from devnet config
	return "http://localhost:7545"
}

func (c *EmbeddedConfig) GetDefaultWSRPCURL() string {
	// Convert HTTP RPC URL to WebSocket URL
	baseURL := c.GetDefaultRPCURL()
	return convertHTTPToWebSocket(baseURL)
}

func convertHTTPToWebSocket(httpURL string) string {
	// Convert http:// to ws:// and https:// to wss://
	if strings.HasPrefix(httpURL, "http://") {
		return strings.Replace(httpURL, "http://", "ws://", 1)
	} else if strings.HasPrefix(httpURL, "https://") {
		return strings.Replace(httpURL, "https://", "wss://", 1)
	}
	// If no http/https prefix, assume ws:// should be used
	return "ws://" + httpURL
}

func (c *EmbeddedConfig) GetTaskMailboxAddress() string {
	return c.TaskMailbox.Address
}

func (c *EmbeddedConfig) GetFaaSAddress() string {
	return c.FaaS.Address
}