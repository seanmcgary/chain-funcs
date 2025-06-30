# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is **fass-avs**, a Function-as-a-Service (FaaS) Autonomous Verifiable Service (AVS) built using the Hourglass framework for EigenLayer. It enables users to register and execute arbitrary JavaScript functions in a decentralized manner, similar to AWS Lambda but running on EigenLayer's restaking infrastructure.

## Essential Commands

### Build and Development
- `make build` - Builds the Go binary to `./bin/performer`
- `make deps` - Installs and tidies Go dependencies  
- `make test` - Runs the test suite
- `make build/container` - Builds Docker container

### DevKit CLI (Primary Development Tool)
- `devkit avs build` - Compile contracts and binaries
- `devkit avs devnet start` - Launch local development network with operators
- `devkit avs call` - Simulate task execution
- `devkit avs config` - Manage project configuration

### Testing
- `go test ./...` - Run Go unit tests
- `forge test` - Run smart contract tests (in contracts/ directory)

## Architecture

### Core Components
- **Performer** (`/cmd/main.go`) - Implements business logic with `ValidateTask()` and `HandleTask()` methods
- **Aggregator** (framework-provided) - Coordinates task execution across operators
- **Executor** (framework-provided) - Manages performer containers and BLS signing

### Smart Contracts
- **L1**: `TaskAVSRegistrar` - Operator registration and BLS key management
- **L2**: `AVSTaskHook` - Task lifecycle validation and fee markets
- **L2**: `BN254CertificateVerifier` - Cryptographic verification of operator certificates
- **FaaS**: `FaaS.sol` - JavaScript function registration and execution contract

### Key Directories
- `/cmd/` - Main application entry point and JavaScript execution logic
- `/config/` - Project configuration and environment settings
- `/contracts/src/l1-contracts/` - Layer 1 smart contracts
- `/contracts/src/l2-contracts/` - Layer 2 smart contracts
- `/contracts/src/FaaS.sol` - FaaS contract for function registration and execution
- `/scripts/js-runner.js` - Node.js wrapper for executing JavaScript functions
- `/examples/` - Sample JavaScript functions and usage documentation
- `/clients/` - Go-based CLI for interacting with FaaS contracts
- `/keystores/` - BLS keystores for local development operators

## Development Workflow

### For AVS Development:
1. The Go Performer automatically handles JavaScript execution - no code changes needed
2. Add custom smart contracts alongside existing framework contracts
3. Update configuration in `/config/contexts/devnet.yaml` as needed
4. Use `devkit avs build` to compile everything
5. Use `devkit avs devnet start` to test in local environment

### For JavaScript Function Development:
1. Create JavaScript functions following the structure in `/examples/`
2. Use the CLI tool in `/clients/` to deploy and call functions:
   - `cd clients && make build`
   - `./bin/faas-cli deploy-function ../examples/hello-function`
   - `./bin/faas-cli call-function --function-id=0x... "input"`
3. Or manually package functions as gzipped tarballs and interact via smart contracts
4. CLI automatically handles tarball creation, contract interaction, and result polling

## Technology Stack

- **Go 1.23.6** - Main application language with tarball extraction and process execution
- **Node.js** - JavaScript runtime for function execution
- **Solidity ^0.8.27** - Smart contract development  
- **Foundry** - Contract compilation and testing
- **gRPC** - Inter-service communication
- **BN254 cryptography** - Operator signatures and verification
- **EigenLayer/Hourglass framework** - Core AVS infrastructure

## FaaS Function Format

All functions must:
- Be packaged as gzipped tarballs
- **Contain `manifest.json` at the root specifying the runtime**
- Include all dependencies in the tarball
- Return JSON-serializable results

### Required Files:

#### manifest.json (Required)
```json
{
  "runtime": "javascript"
}
```
or
```json
{
  "runtime": "python"
}
```

#### JavaScript Functions:
- Contain `index.js` at the root with `exports.handler = async function(args = []) { ... }`
- Include node_modules if dependencies are needed

#### Python Functions:
- Contain `handler.py` at the root with `def handler(args): ...`
- Include requirements.txt and any dependencies needed

## Development Resources

### Documentation Lookup
When working on frontend development or looking up library documentation, use the **Context7 MCP server** which provides access to up-to-date documentation for:
- React, Vite, and other frontend frameworks
- Web3 libraries (Wagmi, Viem)
- UI component libraries (Tailwind CSS, shadcn/ui)
- JavaScript/TypeScript APIs and best practices

Use Context7 to quickly resolve API questions and implementation patterns during development.