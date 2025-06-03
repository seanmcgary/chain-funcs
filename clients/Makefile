.PHONY: build clean test install

# Build the CLI binary
build:
	mkdir -p bin
	go build -o bin/faas-cli .

# Clean build artifacts
clean:
	rm -rf bin

# Run tests
test:
	go test ./...

# Install dependencies
deps:
	go mod tidy
	go mod download

# Install the CLI to $GOPATH/bin
install:
	go install .

# Format code
fmt:
	go fmt ./...

# Build for multiple platforms
build-all:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o bin/faas-cli-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -o bin/faas-cli-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o bin/faas-cli-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build -o bin/faas-cli-windows-amd64.exe .

# Default target
all: deps fmt build