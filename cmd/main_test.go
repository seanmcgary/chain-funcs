package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	performerV1 "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/hourglass/v1/performer"
	"go.uber.org/zap"
)

// Helper function to encode a FunctionCall struct as Solidity would
func encodeFunctionCallStruct(fn []byte, fnId [32]byte, input []byte, url string) ([]byte, error) {
	functionCallType, err := abi.NewType("tuple", "FunctionCall", []abi.ArgumentMarshaling{
		{Name: "fn", Type: "bytes"},
		{Name: "fnId", Type: "bytes32"},
		{Name: "input", Type: "bytes"},
		{Name: "url", Type: "string"},
	})
	if err != nil {
		return nil, err
	}

	functionCallArgs := abi.Arguments{
		{Type: functionCallType, Name: "call"},
	}

	structValue := struct {
		Fn    []byte   `abi:"fn"`
		FnId  [32]byte `abi:"fnId"`
		Input []byte   `abi:"input"`
		Url   string   `abi:"url"`
	}{
		Fn:    fn,
		FnId:  fnId,
		Input: input,
		Url:   url,
	}

	return functionCallArgs.Pack(structValue)
}

func TestFunctionCallABIDecoding(t *testing.T) {
	// Create a logger for testing
	logger, _ := zap.NewDevelopment()
	
	// Create task worker
	worker := NewTaskWorker(logger)

	// Create test data
	testFunctionData := []byte("test function content - this would be a gzipped tarball")
	testFunctionId := crypto.Keccak256Hash(testFunctionData)
	testInput := []byte(`["hello", "world"]`) // JSON array of arguments

	// Pack the data as the Solidity contract would
	packedData, err := encodeFunctionCallStruct(testFunctionData, testFunctionId, testInput, "")
	if err != nil {
		t.Fatalf("Failed to pack test data: %v", err)
	}

	// Create a task request with the packed payload
	taskRequest := &performerV1.TaskRequest{
		TaskId:  []byte("test-task-123"),
		Payload: packedData,
	}

	// Test validation
	err = worker.ValidateTask(taskRequest)
	if err != nil {
		t.Fatalf("ValidateTask failed: %v", err)
	}

	// Test decoding
	decodedCall, err := worker.decodeFunctionCall(taskRequest.Payload)
	if err != nil {
		t.Fatalf("decodeFunctionCall failed: %v", err)
	}

	// Verify the decoded data matches what we encoded
	if string(decodedCall.Fn) != string(testFunctionData) {
		t.Errorf("Function data mismatch. Expected: %s, Got: %s", string(testFunctionData), string(decodedCall.Fn))
	}

	if decodedCall.FnId != testFunctionId {
		t.Errorf("Function ID mismatch. Expected: %x, Got: %x", testFunctionId, decodedCall.FnId)
	}

	if string(decodedCall.Input) != string(testInput) {
		t.Errorf("Input data mismatch. Expected: %s, Got: %s", string(testInput), string(decodedCall.Input))
	}

	// Test that the input can be parsed as JSON
	var inputArgs interface{}
	err = json.Unmarshal(decodedCall.Input, &inputArgs)
	if err != nil {
		t.Fatalf("Failed to unmarshal input as JSON: %v", err)
	}

	// Verify the JSON structure
	argsArray, ok := inputArgs.([]interface{})
	if !ok {
		t.Fatalf("Input is not a JSON array")
	}

	if len(argsArray) != 2 {
		t.Fatalf("Expected 2 arguments, got %d", len(argsArray))
	}

	if argsArray[0].(string) != "hello" || argsArray[1].(string) != "world" {
		t.Errorf("Unexpected argument values: %v", argsArray)
	}

	t.Logf("✅ All ABI encoding/decoding tests passed!")
	t.Logf("Function ID: %x", decodedCall.FnId)
	t.Logf("Function data length: %d bytes", len(decodedCall.Fn))
	t.Logf("Input arguments: %s", string(decodedCall.Input))
}

func TestFunctionCallValidation(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	worker := NewTaskWorker(logger)

	tests := []struct {
		name        string
		fn          []byte
		fnId        [32]byte
		input       []byte
		expectError bool
		errorMsg    string
	}{
		{
			name:        "Valid function call",
			fn:          []byte("valid function content"),
			fnId:        [32]byte{1, 2, 3}, // Non-zero function ID
			input:       []byte(`["test"]`),
			expectError: false,
		},
		{
			name:        "Empty function content",
			fn:          []byte{},
			fnId:        [32]byte{1, 2, 3},
			input:       []byte(`["test"]`),
			expectError: true,
			errorMsg:    "function content is empty",
		},
		{
			name:        "Empty input",
			fn:          []byte("valid function content"),
			fnId:        [32]byte{1, 2, 3},
			input:       []byte{},
			expectError: true,
			errorMsg:    "function input is empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Encode the test data
			packedData, err := encodeFunctionCallStruct(tt.fn, tt.fnId, tt.input, "")
			if err != nil {
				t.Fatalf("Failed to pack test data: %v", err)
			}

			taskRequest := &performerV1.TaskRequest{
				TaskId:  []byte("test-task"),
				Payload: packedData,
			}

			err = worker.ValidateTask(taskRequest)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("Expected error containing '%s', got: %v", tt.errorMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestInvalidABIData(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	worker := NewTaskWorker(logger)

	// Test with invalid ABI data
	invalidPayload := []byte("this is not valid ABI encoded data")
	
	taskRequest := &performerV1.TaskRequest{
		TaskId:  []byte("test-task"),
		Payload: invalidPayload,
	}

	err := worker.ValidateTask(taskRequest)
	if err == nil {
		t.Error("Expected error when validating invalid ABI data, but got none")
	}

	if !strings.Contains(err.Error(), "invalid function call payload") {
		t.Errorf("Expected error about invalid payload, got: %v", err)
	}

	// Test direct decoding of invalid data
	_, err = worker.decodeFunctionCall(invalidPayload)
	if err == nil {
		t.Error("Expected error when decoding invalid ABI data, but got none")
	}

	if !strings.Contains(err.Error(), "failed to unpack ABI data") {
		t.Errorf("Expected error about unpacking ABI data, got: %v", err)
	}
}

func BenchmarkABIDecoding(b *testing.B) {
	logger, _ := zap.NewDevelopment()
	worker := NewTaskWorker(logger)

	// Prepare test data
	testFunctionData := make([]byte, 1024*10) // 10KB of test data
	for i := range testFunctionData {
		testFunctionData[i] = byte(i % 256)
	}
	
	testFunctionId := crypto.Keccak256Hash(testFunctionData)
	testInput := []byte(`["benchmark", "test", "with", "multiple", "arguments"]`)

	packedData, _ := encodeFunctionCallStruct(testFunctionData, testFunctionId, testInput, "")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := worker.decodeFunctionCall(packedData)
		if err != nil {
			b.Fatalf("Decoding failed: %v", err)
		}
	}
}

// Legacy test for basic functionality
func Test_TaskRequestPayload(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Errorf("Failed to create logger: %v", err)
	}

	taskWorker := NewTaskWorker(logger)

	// Create valid ABI-encoded test data
	testFunctionData := []byte("test function content")
	testFunctionId := crypto.Keccak256Hash(testFunctionData)
	testInput := []byte(`["test"]`)

	packedData, _ := encodeFunctionCallStruct(testFunctionData, testFunctionId, testInput, "")

	taskRequest := &performerV1.TaskRequest{
		TaskId:   []byte("test-task-id"),
		Payload:  packedData,
		Metadata: []byte("test-metadata"),
	}

	err = taskWorker.ValidateTask(taskRequest)
	if err != nil {
		t.Errorf("ValidateTask failed: %v", err)
	}

	// Note: HandleTask would fail without actual JavaScript execution environment
	// but we can test the validation and decoding parts
	t.Logf("Task validation passed for ABI-encoded payload")
}
