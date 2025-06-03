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
	"os/exec"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/Layr-Labs/hourglass-monorepo/ponos/pkg/performer/server"
	performerV1 "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/hourglass/v1/performer"
	"go.uber.org/zap"
)

// This offchain binary is run by Operators running the Hourglass Executor. It contains
// the business logic of the AVS and performs worked based on the tasked sent to it.
// The Hourglass Aggregator ingests tasks from the TaskMailbox and distributes work
// to Executors configured to run the AVS Performer. Performers execute the work and
// return the result to the Executor where the result is signed and return to the
// Aggregator to place in the outbox once the signing threshold is met.

type TaskWorker struct {
	logger              *zap.Logger
	functionCallArguments abi.Arguments
}

type FunctionCall struct {
	Fn    []byte   `json:"fn"`
	FnId  [32]byte `json:"fnId"`
	Input []byte   `json:"input"`
}

type ExecutionResult struct {
	Payload interface{} `json:"payload"`
	Error   *string     `json:"error"`
}

func NewTaskWorker(logger *zap.Logger) *TaskWorker {
	// Create Arguments for FunctionCall struct
	functionCallArgs := abi.Arguments{
		{Type: abi.Type{T: abi.BytesTy}, Name: "fn"},
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}, Name: "fnId"},
		{Type: abi.Type{T: abi.BytesTy}, Name: "input"},
	}
	
	return &TaskWorker{
		logger:              logger,
		functionCallArguments: functionCallArgs,
	}
}

func (tw *TaskWorker) ValidateTask(t *performerV1.TaskRequest) error {
	tw.logger.Sugar().Infow("Validating task",
		zap.Any("task", t),
	)

	call, err := tw.decodeFunctionCall(t.Payload)
	if err != nil {
		return fmt.Errorf("invalid function call payload: %w", err)
	}

	if len(call.Fn) == 0 {
		return fmt.Errorf("function content is empty")
	}

	if len(call.Input) == 0 {
		return fmt.Errorf("function input is empty")
	}

	return nil
}

func (tw *TaskWorker) HandleTask(t *performerV1.TaskRequest) (*performerV1.TaskResponse, error) {
	tw.logger.Sugar().Infow("Handling task",
		zap.Any("task", t),
	)

	call, err := tw.decodeFunctionCall(t.Payload)
	if err != nil {
		return nil, fmt.Errorf("invalid function call payload: %w", err)
	}

	result, err := tw.executeJavaScriptFunction(call)
	if err != nil {
		return nil, fmt.Errorf("failed to execute JavaScript function: %w", err)
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result: %w", err)
	}

	return &performerV1.TaskResponse{
		TaskId: t.TaskId,
		Result: resultBytes,
	}, nil
}

// decodeFunctionCall decodes ABI-encoded FunctionCall struct from payload
func (tw *TaskWorker) decodeFunctionCall(payload []byte) (*FunctionCall, error) {
	// Unpack the ABI-encoded data
	values, err := tw.functionCallArguments.Unpack(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack ABI data: %w", err)
	}

	if len(values) != 3 {
		return nil, fmt.Errorf("expected 3 values, got %d", len(values))
	}

	// Extract the values
	fn, ok := values[0].([]byte)
	if !ok {
		return nil, fmt.Errorf("fn field is not bytes")
	}

	fnId, ok := values[1].([32]byte)
	if !ok {
		return nil, fmt.Errorf("fnId field is not bytes32")
	}

	input, ok := values[2].([]byte)
	if !ok {
		return nil, fmt.Errorf("input field is not bytes")
	}

	return &FunctionCall{
		Fn:    fn,
		FnId:  fnId,
		Input: input,
	}, nil
}

func (tw *TaskWorker) executeJavaScriptFunction(call *FunctionCall) (*ExecutionResult, error) {
	// Convert function ID to hex string for directory name
	functionId := fmt.Sprintf("%x", call.FnId)
	functionDir := filepath.Join("/tmp/functions", functionId)
	
	if err := os.MkdirAll(functionDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create function directory: %w", err)
	}
	defer os.RemoveAll(functionDir)

	if err := tw.extractTarball(call.Fn, functionDir); err != nil {
		return nil, fmt.Errorf("failed to extract function tarball: %w", err)
	}

	var inputArgs interface{}
	if err := json.Unmarshal(call.Input, &inputArgs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal input arguments: %w", err)
	}

	inputJSON, err := json.Marshal(inputArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal input arguments: %w", err)
	}

	cmd := exec.Command("node", "/app/scripts/js-runner.js", functionDir, string(inputJSON))
	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			tw.logger.Error("JavaScript execution failed", zap.String("stderr", string(exitError.Stderr)))
		}
		return nil, fmt.Errorf("failed to execute JavaScript function: %w", err)
	}

	var result ExecutionResult
	if err := json.Unmarshal(output, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal execution result: %w", err)
	}

	return &result, nil
}

func (tw *TaskWorker) extractTarball(tarballData []byte, destDir string) error {
	gzipReader, err := gzip.NewReader(bytes.NewReader(tarballData))
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read tar header: %w", err)
		}

		destPath := filepath.Join(destDir, header.Name)
		
		if !filepath.HasPrefix(destPath, destDir) {
			return fmt.Errorf("invalid file path: %s", header.Name)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(destPath, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", destPath, err)
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
				return fmt.Errorf("failed to create parent directory for %s: %w", destPath, err)
			}

			file, err := os.OpenFile(destPath, os.O_CREATE|os.O_WRONLY, os.FileMode(header.Mode))
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", destPath, err)
			}

			if _, err := io.Copy(file, tarReader); err != nil {
				file.Close()
				return fmt.Errorf("failed to write file %s: %w", destPath, err)
			}
			file.Close()
		}
	}

	return nil
}

func main() {
	ctx := context.Background()
	l, _ := zap.NewProduction()

	w := NewTaskWorker(l)

	pp, err := server.NewPonosPerformerWithRpcServer(&server.PonosPerformerConfig{
		Port:    8080,
		Timeout: 5 * time.Second,
	}, w, l)
	if err != nil {
		panic(fmt.Errorf("failed to create performer: %w", err))
	}

	if err := pp.Start(ctx); err != nil {
		panic(err)
	}
}
