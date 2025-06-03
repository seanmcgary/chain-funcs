package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	performerV1 "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/hourglass/v1/performer"
	"go.uber.org/zap"
)

// Helper function to encode a FunctionCall struct as Solidity would
func encodeFunctionCallStructLocal(fn []byte, fnId [32]byte, input []byte) ([]byte, error) {
	functionCallType, err := abi.NewType("tuple", "FunctionCall", []abi.ArgumentMarshaling{
		{Name: "fn", Type: "bytes"},
		{Name: "fnId", Type: "bytes32"},
		{Name: "input", Type: "bytes"},
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
	}{
		Fn:    fn,
		FnId:  fnId,
		Input: input,
	}

	return functionCallArgs.Pack(structValue)
}

func TestTaskWorker_ValidateTask(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	tw := NewTaskWorker(logger)

	fn := []byte("dummy content")
	fnId := [32]byte{1, 2, 3}
	input := []byte(`["arg1", "arg2"]`)

	// Encode using ABI tuple structure
	payload, err := encodeFunctionCallStructLocal(fn, fnId, input)
	if err != nil {
		t.Fatal(err)
	}

	taskRequest := &performerV1.TaskRequest{
		TaskId:  []byte("test-task"),
		Payload: payload,
	}

	err = tw.ValidateTask(taskRequest)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestTaskWorker_ValidateTask_InvalidPayload(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	tw := NewTaskWorker(logger)

	taskRequest := &performerV1.TaskRequest{
		TaskId:  []byte("test-task"),
		Payload: []byte("invalid json"),
	}

	err := tw.ValidateTask(taskRequest)
	if err == nil {
		t.Error("Expected error for invalid payload")
	}
}

func TestTaskWorker_ValidateTask_EmptyFunction(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	tw := NewTaskWorker(logger)

	fn := []byte{}
	fnId := [32]byte{1, 2, 3}
	input := []byte(`["arg1", "arg2"]`)

	// Encode using ABI tuple structure
	payload, err := encodeFunctionCallStructLocal(fn, fnId, input)
	if err != nil {
		t.Fatal(err)
	}

	taskRequest := &performerV1.TaskRequest{
		TaskId:  []byte("test-task"),
		Payload: payload,
	}

	err = tw.ValidateTask(taskRequest)
	if err == nil {
		t.Error("Expected error for empty function content")
	}
}

func TestTaskWorker_ExtractTarball(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	tw := NewTaskWorker(logger)

	tarballData := createTestTarball(t)

	tmpDir := t.TempDir()
	
	err := tw.extractTarball(tarballData, tmpDir)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func createTestTarball(t *testing.T) []byte {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)
	tarWriter := tar.NewWriter(gzipWriter)

	content := `exports.handler = async function(args = []) {
    return { message: "Hello from test function", args: args };
};`

	header := &tar.Header{
		Name: "index.js",
		Mode: 0644,
		Size: int64(len(content)),
	}

	if err := tarWriter.WriteHeader(header); err != nil {
		t.Fatal(err)
	}

	if _, err := tarWriter.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}

	tarWriter.Close()
	gzipWriter.Close()

	return buf.Bytes()
}