package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed contracts/taskMailbox.json
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

func (c *EmbeddedConfig) GetTaskMailboxAddress() string {
	return c.TaskMailbox.Address
}

func (c *EmbeddedConfig) GetFaaSAddress() string {
	return c.FaaS.Address
}