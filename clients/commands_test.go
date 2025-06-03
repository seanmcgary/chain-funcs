package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateTarball(t *testing.T) {
	// Create a temporary directory with test files
	tempDir := t.TempDir()
	
	// Create test files
	indexJS := `exports.handler = async function(args = []) {
    return { message: "test", args: args };
};`
	
	err := os.WriteFile(filepath.Join(tempDir, "index.js"), []byte(indexJS), 0644)
	if err != nil {
		t.Fatal(err)
	}
	
	packageJSON := `{
    "name": "test-function",
    "version": "1.0.0"
}`
	
	err = os.WriteFile(filepath.Join(tempDir, "package.json"), []byte(packageJSON), 0644)
	if err != nil {
		t.Fatal(err)
	}
	
	// Create tarball
	tarballData, err := createTarball(tempDir)
	if err != nil {
		t.Fatalf("Failed to create tarball: %v", err)
	}
	
	// Verify tarball is not empty
	if len(tarballData) == 0 {
		t.Fatal("Tarball is empty")
	}
	
	t.Logf("Tarball size: %d bytes", len(tarballData))
}