package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"time"

	"github.com/Layr-Labs/hourglass-monorepo/ponos/pkg/performer/server"
	performerV1 "github.com/Layr-Labs/protocol-apis/gen/protos/eigenlayer/hourglass/v1/performer"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"go.uber.org/zap"
)

// This offchain binary is run by Operators running the Hourglass Executor. It contains
// the business logic of the AVS and performs worked based on the tasked sent to it.
// The Hourglass Aggregator ingests tasks from the TaskMailbox and distributes work
// to Executors configured to run the AVS Performer. Performers execute the work and
// return the result to the Executor where the result is signed and return to the
// Aggregator to place in the outbox once the signing threshold is met.

type TaskWorker struct {
	logger                *zap.Logger
	functionCallArguments abi.Arguments
}

type FunctionCall struct {
	Fn    []byte   `json:"fn"`
	FnId  [32]byte `json:"fnId"`
	Input []byte   `json:"input"`
	Url   string   `json:"url"`
}

type ExecutionResult struct {
	Payload interface{} `json:"payload"`
	Error   *string     `json:"error"`
}

func NewTaskWorker(logger *zap.Logger) *TaskWorker {
	// Create Arguments for FunctionCall struct as a single tuple
	// This matches how Solidity's abi.encode(struct) works
	functionCallType, err := abi.NewType("tuple", "FunctionCall", []abi.ArgumentMarshaling{
		{Name: "fn", Type: "bytes"},
		{Name: "fnId", Type: "bytes32"},
		{Name: "input", Type: "bytes"},
		{Name: "url", Type: "string"},
	})
	if err != nil {
		logger.Fatal("Failed to create FunctionCall tuple type", zap.Error(err))
	}

	functionCallArgs := abi.Arguments{
		{Type: functionCallType, Name: "call"},
	}

	return &TaskWorker{
		logger:                logger,
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

	if len(call.Fn) == 0 && len(call.Url) == 0 {
		return fmt.Errorf("task must specify either a function or a URL")
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

	result, err := tw.executeFunction(call)
	if err != nil {
		return nil, fmt.Errorf("failed to execute function: %w", err)
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

	if len(values) != 1 {
		return nil, fmt.Errorf("expected 1 tuple value, got %d", len(values))
	}

	// The ABI library automatically creates a struct for tuple types
	// Use reflection to extract the fields
	tupleStruct := values[0]

	// Use reflection to access the fields
	v := reflect.ValueOf(tupleStruct)
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %v", v.Kind())
	}

	// Extract fields by name
	fnField := v.FieldByName("Fn")
	if !fnField.IsValid() {
		return nil, fmt.Errorf("fn field not found")
	}
	fn := fnField.Bytes()

	fnIdField := v.FieldByName("FnId")
	if !fnIdField.IsValid() {
		return nil, fmt.Errorf("fnId field not found")
	}

	// Extract [32]byte array directly
	var fnId [32]byte
	if fnIdField.Kind() == reflect.Array && fnIdField.Len() == 32 {
		for i := 0; i < 32; i++ {
			fnId[i] = byte(fnIdField.Index(i).Uint())
		}
	} else {
		return nil, fmt.Errorf("fnId field is not a 32-byte array")
	}

	inputField := v.FieldByName("Input")
	if !inputField.IsValid() {
		return nil, fmt.Errorf("input field not found")
	}
	input := inputField.Bytes()

	urlField := v.FieldByName("Url")
	if !urlField.IsValid() {
		return nil, fmt.Errorf("url field not found")
	}
	url := urlField.String()

	return &FunctionCall{
		Fn:    fn,
		FnId:  fnId,
		Input: input,
		Url:   url,
	}, nil
}

func (tw *TaskWorker) executeFunction(call *FunctionCall) (*ExecutionResult, error) {
	tw.logger.Sugar().Infow("Executing function",
		zap.String("functionId", fmt.Sprintf("%x", call.FnId)),
		zap.Int("inputLength", len(call.Input)),
		zap.Int("fnLength", len(call.Fn)),
		zap.String("url", call.Url),
	)
	// Convert function ID to hex string for directory name
	functionId := fmt.Sprintf("%x", call.FnId)
	functionDir := filepath.Join("/tmp/functions", functionId)

	if err := os.MkdirAll(functionDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create function directory: %w", err)
	}
	defer os.RemoveAll(functionDir)

	// Handle remote URL or embedded content
	if call.Url != "" {
		// Use remote URL
		if err := tw.downloadAndExtractFunction(call.Url, functionDir); err != nil {
			return nil, fmt.Errorf("failed to download and extract function: %w", err)
		}
	} else {
		// Use embedded content
		if err := tw.extractTarball(call.Fn, functionDir); err != nil {
			return nil, fmt.Errorf("failed to extract function tarball: %w", err)
		}
	}

	// Debug: list extracted files
	files, err := filepath.Glob(filepath.Join(functionDir, "*"))
	if err == nil {
		tw.logger.Sugar().Infof("Extracted files: %v", files)
	}

	// Pass raw input data as base64 to preserve binary data
	inputBase64 := base64.StdEncoding.EncodeToString(call.Input)

	// Detect runtime based on files present
	runtime, err := tw.detectRuntime(functionDir)
	if err != nil {
		return nil, fmt.Errorf("failed to detect runtime: %w", err)
	}

	var cmd *exec.Cmd
	switch runtime {
	case "javascript":
		scriptPath, err := tw.findScript("js-runner.js")
		if err != nil {
			return nil, err
		}
		tw.logger.Sugar().Infof("Executing JavaScript: node %s %s <base64-encoded-input>", scriptPath, functionDir)
		cmd = exec.Command("node", scriptPath, functionDir, inputBase64)
	case "python":
		scriptPath, err := tw.findScript("py-runner.py")
		if err != nil {
			return nil, err
		}
		tw.logger.Sugar().Infof("Executing Python: python3 %s %s <base64-encoded-input>", scriptPath, functionDir)
		cmd = exec.Command("python3", scriptPath, functionDir, inputBase64)
	default:
		return nil, fmt.Errorf("unsupported runtime: %s", runtime)
	}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		tw.logger.Error("Function execution failed",
			zap.String("runtime", runtime),
			zap.String("stderr", stderr.String()),
			zap.String("stdout", stdout.String()),
			zap.Error(err))
		return nil, fmt.Errorf("failed to execute %s function: %w", runtime, err)
	}

	output := stdout.Bytes()

	// Debug: log the raw output to see what's being returned
	tw.logger.Sugar().Infof("Node.js raw output: %q", string(output))

	var result ExecutionResult
	if err := json.Unmarshal(output, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal execution result (output: %q): %w", string(output), err)
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

// downloadAndExtractFunction downloads a tarball from URL and extracts it
func (tw *TaskWorker) downloadAndExtractFunction(url, destDir string) error {
	tw.logger.Sugar().Infow("Downloading and extracting function",
		zap.String("url", url),
		zap.String("destDir", destDir),
	)
	// Create a persistent cache directory for downloaded functions
	cacheDir := filepath.Join("/tmp/function-cache")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	// Use URL hash as cache key
	urlHash := fmt.Sprintf("%x", sha256.Sum256([]byte(url)))
	cachedPath := filepath.Join(cacheDir, urlHash+".tar.gz")

	// Check if already cached
	if _, err := os.Stat(cachedPath); err == nil {
		tw.logger.Sugar().Infof("Using cached function from %s", cachedPath)
		return tw.extractTarballFromFile(cachedPath, destDir)
	}

	// Download the tarball
	tw.logger.Sugar().Infof("Downloading function from %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download function: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download function: HTTP %d", resp.StatusCode)
	}

	// Read the content
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Cache the downloaded content
	if err := os.WriteFile(cachedPath, content, 0644); err != nil {
		tw.logger.Sugar().Warnf("Failed to cache downloaded function: %v", err)
		// Continue execution even if caching fails
	}

	// Extract directly from memory
	return tw.extractTarball(content, destDir)
}

// extractTarballFromFile extracts a tarball from a file path
func (tw *TaskWorker) extractTarballFromFile(filePath, destDir string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read cached file: %w", err)
	}
	return tw.extractTarball(content, destDir)
}

// detectRuntime determines the function runtime based on manifest.json
func (tw *TaskWorker) detectRuntime(functionDir string) (string, error) {
	manifestPath := filepath.Join(functionDir, "manifest.json")

	// Check if manifest.json exists
	manifestData, err := os.ReadFile(manifestPath)
	if err != nil {
		return "", fmt.Errorf("manifest.json not found in function directory: %w", err)
	}

	// Parse manifest.json
	var manifest struct {
		Runtime string `json:"runtime"`
	}

	if err := json.Unmarshal(manifestData, &manifest); err != nil {
		return "", fmt.Errorf("failed to parse manifest.json: %w", err)
	}

	// Validate runtime
	switch manifest.Runtime {
	case "javascript", "python":
		return manifest.Runtime, nil
	default:
		return "", fmt.Errorf("unsupported runtime in manifest.json: %s (supported: javascript, python)", manifest.Runtime)
	}
}

// findScript locates the runner script, trying both app and relative paths
func (tw *TaskWorker) findScript(scriptName string) (string, error) {
	// Try app path first (for containerized deployment)
	appPath := fmt.Sprintf("/app/scripts/%s", scriptName)
	if _, err := os.Stat(appPath); err == nil {
		return appPath, nil
	}

	// Try relative path (for local development)
	relativePath := fmt.Sprintf("./scripts/%s", scriptName)
	if _, err := os.Stat(relativePath); err == nil {
		return relativePath, nil
	}

	return "", fmt.Errorf("%s not found at /app/scripts/%s or ./scripts/%s", scriptName, scriptName, scriptName)
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
