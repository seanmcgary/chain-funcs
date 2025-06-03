# Function-as-a-Service (FaaS) AVS

A serverless multi-runtime execution platform built on EigenLayer's Hourglass framework, providing AWS Lambda-like functionality on a decentralized infrastructure.

## Overview

This AVS (Autonomous Verifiable Service) enables users to deploy and execute functions in multiple runtimes on-demand through a decentralized network of operators. It provides:

- **Multi-runtime serverless execution** with JavaScript (Node.js) and Python support
- **Arbitrary input support** - JSON, binary data, text, or any format
- **Automatic scaling** through EigenLayer's operator network
- **Verifiable computation** with cryptographic proofs
- **AWS Lambda-like experience** with familiar handler patterns

## Architecture

### On-Chain Components
- **FaaS.sol**: Core contract for function registration and execution requests
- **TaskMailbox**: Hourglass framework contract managing task lifecycle
- **BN254CertificateVerifier**: Validates operator signatures and stake requirements

### Off-Chain Components
- **Performer**: Go-based multi-runtime execution engine supporting JavaScript (Node.js) and Python
- **Aggregator**: Distributes tasks to operators and aggregates results
- **Executor**: Manages task execution and signs results with BLS keys

### Execution Flow
1. User deploys function (JavaScript or Python) as gzipped tarball with manifest.json via CLI
2. Function is stored on-chain with unique ID
3. Function calls create tasks in the TaskMailbox
4. Operators execute functions in isolated runtime environments (Node.js or Python)
5. Results are aggregated and verified before returning to caller

## Quick Start

### Prerequisites
- Node.js and npm
- Go 1.21+
- Docker
- Git
- [Hourglass DevKit](https://github.com/Layr-Labs/hourglass-monorepo) (install with `npm install -g @layr-labs/hourglass-devkit`)

### 1. Clone and Build

```bash
git clone <repository-url>
cd fass-avs

# Install Hourglass DevKit if not already installed
npm install -g @layr-labs/hourglass-devkit

# Build the project
devkit avs build
```

This builds:
- `./bin/performer` - Multi-runtime execution engine (JavaScript/Python)
- `./clients/bin/faas-cli` - Command-line interface

### 2. Deploy Contracts and Start Devnet

Deploy the FaaS system contracts and start a local development network:

```bash
devkit avs devnet start --skip-avs-run --port 7545
```

This command:
- Starts a local Ethereum network on port 7545
- Deploys all required contracts (TaskMailbox, FaaS, etc.)
- Sets up operator infrastructure
- Contract addresses are automatically embedded in the CLI

### 3. Start the Performer

The performer handles multi-runtime function execution for operators:

```bash
devkit avs run
```

## Using the CLI

The `faas-cli` provides an AWS Lambda-like experience for deploying and calling functions with support for both JavaScript and Python runtimes.

### Deploy a Function

#### JavaScript Functions

Create a directory with your JavaScript function:

```bash
mkdir my-function
cd my-function
```

Create `index.js` with a handler function:

```javascript
// Simple function example
exports.handler = async function(input) {
    return {
        message: "Hello from FaaS!",
        input: input,
        timestamp: new Date().toISOString()
    };
};
```

```javascript
// Advanced function with console logging
exports.handler = async function(input) {
    console.log("Function called with:", input);
    
    // Handle different input types
    if (input && typeof input === 'object' && input.text) {
        // JSON input
        return { result: `Processed: ${input.text}` };
    } else if (input && input.raw) {
        // Binary input
        console.log("Received binary data, size:", input.raw.length);
        return { size: input.raw.length, type: "binary" };
    } else {
        // Raw string or other input
        return { echo: input };
    }
};
```

#### Python Functions

Create a directory with your Python function:

```bash
mkdir my-python-function
cd my-python-function
```

Create `handler.py` with a handler function:

```python
def handler(input_data):
    """
    Python function handler
    
    Args:
        input_data: The input data from the function call
        
    Returns:
        The result of the function execution
    """
    print(f"Python function called with: {input_data}")
    
    # Handle different input types
    if isinstance(input_data, dict) and 'name' in input_data:
        return {
            "message": f"Hello, {input_data['name']}! This is from Python.",
            "language": "python",
            "input": input_data
        }
    elif isinstance(input_data, dict) and 'raw' in input_data:
        # Binary input
        return {
            "message": "Processed binary data with Python",
            "size": len(input_data['raw']),
            "type": "binary"
        }
    else:
        # Simple string or other input
        return {
            "message": f"Echo from Python: {input_data}",
            "language": "python"
        }
```

#### Required manifest.json

**All functions must include a `manifest.json` file specifying the runtime:**

For JavaScript functions:
```json
{
  "runtime": "javascript"
}
```

For Python functions:
```json
{
  "runtime": "python"
}
```

#### Local Deployment

Deploy the function:

```bash
faas-cli deploy-function \
  --private-key <YOUR_PRIVATE_KEY> \
  ./my-function
```

#### Remote Deployment (S3)

For larger functions or production deployments, use S3 storage:

```bash
# Deploy JavaScript function to S3
faas-cli deploy-function \
  --private-key <YOUR_PRIVATE_KEY> \
  --s3-base-url s3://my-bucket/functions/ \
  ./my-function

# Deploy Python function to S3
faas-cli deploy-function \
  --private-key <YOUR_PRIVATE_KEY> \
  --s3-base-url s3://my-bucket/functions/ \
  ./my-python-function
```

**Prerequisites for S3 deployment:**
- AWS CLI configured with credentials
- S3 bucket with appropriate permissions
- Functions are downloaded by operators at runtime

Output (S3 deployment):
```
Deploying function from directory: ./my-function
Tarball size: 324 bytes
Function ID: 0x1234567890abcdef...
Using chain ID: 31337
Uploaded to S3: https://my-bucket.s3.amazonaws.com/functions/0x1234567890abcdef....tar.gz
Transaction sent: 0xabcdef1234567890...
Waiting for confirmation...
Function registered successfully!
Function ID: 0x1234567890abcdef...
Gas used: 68432
Gas price: 20.00 gwei
Transaction cost: 0.001369 ETH
```

Output (Local deployment):
```
Deploying function from directory: ./my-function
Tarball size: 324 bytes
Function ID: 0x1234567890abcdef...
Using chain ID: 31337
Transaction sent: 0xabcdef1234567890...
Waiting for confirmation...
Function registered successfully!
Function ID: 0x1234567890abcdef...
Gas used: 125847
Gas price: 20.00 gwei
Transaction cost: 0.002517 ETH
```

### Call a Function

Execute your deployed function:

```bash
faas-cli call-function \
  --private-key <YOUR_PRIVATE_KEY> \
  --function-id 0x1234567890abcdef... \
  "Hello, World!"
```

Output:
```
Calling function: 0x1234567890abcdef...
Arguments: ["Hello, World!"]
Transaction sent: 0xfedcba0987654321...
Waiting for confirmation...
Task ID: 0xabcdef1234567890...
Gas used: 68432
Gas price: 20.00 gwei
Transaction cost: 0.001369 ETH
Connected to WebSocket RPC: ws://localhost:7545
Listening for TaskVerified event...
Subscribed to TaskVerified events for task 0xabcdef1234567890...

Task completed successfully!
Event found in block 12345, transaction 0x9876543210abcdef...
Function Result:
{
  "message": "Hello from FaaS!",
  "input": "Hello, World!",
  "timestamp": "2025-01-01T12:00:00.000Z"
}
```

### View Configuration

Check embedded contract addresses:

```bash
faas-cli config
```

Output:
```
Embedded Contract Configuration:

RPC URL: http://localhost:7545
WebSocket RPC URL: ws://localhost:7545
TaskMailbox Address: 0x4B7099FD879435a087C364aD2f9E7B3f94d20bBe
FaaS Address: 0x240A60DC5e0B9013Cb8CF39aa6f9dDd8f25E40D2
```

## Function Development Guide

### Handler Function

#### JavaScript Functions

All JavaScript functions must export a `handler` function:

```javascript
exports.handler = async function(input) {
    // Your logic here
    return result;
};
```

#### Python Functions

All Python functions must define a `handler` function:

```python
def handler(input_data):
    """
    Function handler
    
    Args:
        input_data: The input data from the function call
        
    Returns:
        The result of the function execution
    """
    # Your logic here
    return result
```

### Input Handling

The system supports any input type:

- **JSON data**: Parsed automatically if valid JSON
- **Text data**: Provided as string
- **Binary data**: Available as Buffer with metadata

```javascript
exports.handler = async function(input) {
    if (typeof input === 'string') {
        // Simple string input
        return { type: 'string', value: input };
    } else if (input && input.raw) {
        // Binary input with metadata
        return { 
            type: 'binary', 
            size: input.raw.length,
            text: input.text 
        };
    } else {
        // JSON object input
        return { type: 'json', data: input };
    }
};
```

### Logging

Console output is captured and included in results:

```javascript
exports.handler = async function(input) {
    console.log("Processing request:", input);
    console.log("Current time:", new Date());
    
    return { status: "completed" };
};
```

Logs appear in the result payload:
```json
{
  "result": { "status": "completed" },
  "logs": [
    "Processing request: Hello",
    "Current time: 2025-01-01T12:00:00.000Z"
  ]
}
```

### Dependencies

#### JavaScript Dependencies

Include `package.json` for npm dependencies:

```json
{
  "name": "my-function",
  "dependencies": {
    "lodash": "^4.17.21"
  }
}
```

```javascript
const _ = require('lodash');

exports.handler = async function(input) {
    return _.capitalize(input);
};
```

#### Python Dependencies

For Python functions with external dependencies, include `requirements.txt`:

```txt
requests==2.31.0
numpy==1.24.0
beautifulsoup4==4.12.0
```

**How it works:**
1. The CLI automatically detects Python functions (`handler.py` present)
2. If `requirements.txt` exists, dependencies are installed using `pip3 install --target`
3. Dependencies are bundled into the function tarball
4. The Python runner adds the function directory to `sys.path` to find packages

```python
import requests
import numpy as np
from bs4 import BeautifulSoup

def handler(input_data):
    # Use external libraries - they're automatically bundled
    response = requests.get('https://httpbin.org/json')
    data = np.array([1, 2, 3, 4, 5])
    
    return {
        "message": "Using bundled dependencies",
        "mean": float(np.mean(data)),
        "api_status": response.status_code,
        "packages": ["requests", "numpy", "beautifulsoup4"],
        "input": input_data
    }
```

**Deployment with dependencies:**
```bash
# Dependencies are automatically installed and bundled
faas-cli deploy-function \
  --private-key <YOUR_PRIVATE_KEY> \
  ./my-python-function

# Skip dependency installation (for functions without requirements.txt)
faas-cli deploy-function \
  --private-key <YOUR_PRIVATE_KEY> \
  --skip-deps \
  ./my-simple-python-function
```

**Notes:**
- Requires `pip3` to be available in your environment
- Dependencies are installed to a temporary directory and bundled into the tarball
- Works with both local and S3 deployment
- Use `--skip-deps` to skip dependency installation for functions without external packages
- Dependencies are isolated per function and don't affect your local Python environment

## CLI Reference

### Global Flags

- `--rpc-url`: Ethereum RPC endpoint (defaults to embedded)
- `--ws-rpc-url`: WebSocket RPC endpoint for event subscriptions (auto-converts from rpc-url)
- `--private-key`: Private key for transactions (required)
- `--faas-address`: FaaS contract address (defaults to embedded)
- `--taskmailbox-address`: TaskMailbox address (defaults to embedded)
- `--s3-base-url`: S3 base URL for remote deployment (deploy-function only)
- `--skip-deps`: Skip Python dependency installation (deploy-function only)

### Commands

#### `deploy-function <directory>`

Deploy a JavaScript or Python function from a directory.

**Local deployment:**
```bash
faas-cli deploy-function \
  --private-key 0x123... \
  --rpc-url https://mainnet.infura.io/v3/... \
  ./my-function
```

**S3 deployment:**
```bash
faas-cli deploy-function \
  --private-key 0x123... \
  --rpc-url https://mainnet.infura.io/v3/... \
  --s3-base-url s3://my-bucket/functions/ \
  ./my-function
```

**Python function with dependencies:**
```bash
faas-cli deploy-function \
  --private-key 0x123... \
  ./my-python-function-with-deps
```

**Skip dependency installation:**
```bash
faas-cli deploy-function \
  --private-key 0x123... \
  --skip-deps \
  ./my-simple-python-function
```

#### `call-function <input>`

Execute a deployed function.

**Example:**
```bash
faas-cli call-function \
  --function-id 0xabc123... \
  --private-key 0x123... \
  '{"name": "Alice", "age": 30}'
```

**Example with custom WebSocket URL:**
```bash
faas-cli call-function \
  --function-id 0xabc123... \
  --private-key 0x123... \
  --rpc-url https://mainnet.infura.io/v3/your-project-id \
  --ws-rpc-url wss://mainnet.infura.io/ws/v3/your-project-id \
  '{"name": "Alice", "age": 30}'
```

#### `config`

Display embedded contract configuration.

## Development Setup

### Local Testing

1. Build the project: `devkit avs build`
2. Start local devnet: `devkit avs devnet start --skip-avs-run --port 7545`
3. Start performer: `devkit avs run`
4. Deploy and test functions using the CLI

### Docker Development

Build and run in containers:

```bash
# Build container
docker build -t fass-avs .

# Run performer
docker run -p 8080:8080 fass-avs
```

## Error Handling

The system provides detailed error messages:

```javascript
exports.handler = async function(input) {
    if (!input || !input.name) {
        throw new Error("Name is required");
    }
    return { greeting: `Hello, ${input.name}!` };
};
```

Errors are returned in the response:
```json
{
  "payload": null,
  "error": "Name is required"
}
```

## Security Considerations

- Functions run in isolated Node.js processes
- No file system persistence between executions
- Network access depends on operator configuration
- Input validation is developer responsibility
- Private keys should use environment variables in production

## Real-Time Event Subscription

The CLI uses WebSocket subscriptions to listen for `TaskVerified` events instead of polling, providing real-time updates when tasks complete:

### WebSocket Connection
- **Automatic URL conversion**: HTTP URLs are automatically converted to WebSocket URLs (http → ws, https → wss)
- **Separate connection**: Uses dedicated WebSocket connection for event subscriptions while maintaining HTTP for transactions
- **Custom WebSocket URLs**: Override with `--ws-rpc-url` for providers with different WebSocket endpoints

### Features
- **Real-time notifications**: Immediate response when tasks complete
- **Efficient**: No polling overhead or unnecessary network requests  
- **Reliable**: Direct blockchain event subscription with timeout handling
- **Detailed feedback**: Shows block number and transaction hash of completion

### How It Works
1. After submitting a function call, the CLI subscribes to `TaskVerified` events
2. Filters events to only listen for the specific task ID
3. Automatically parses the result when the matching event is received
4. Provides immediate feedback with block and transaction details

### Event Structure
The `TaskVerified` event contains:
- `aggregator`: Address of the aggregator that verified the task
- `taskHash`: Unique identifier of the completed task
- `avs`: Address of the AVS that processed the task
- `executorOperatorSetId`: ID of the operator set that executed the task
- `result`: The actual function execution result (JSON encoded)

## Performance

- **Cold starts**: ~100-500ms for simple functions
- **Execution time**: Limited by operator configuration
- **Memory**: Isolated per function execution
- **Concurrency**: Scales with operator network size
- **Event delivery**: Real-time WebSocket notifications (~100-200ms after block confirmation)

## Troubleshooting

### Common Issues

**"Function not found"**
- Verify function ID is correct
- Ensure function was successfully deployed

**"Invalid JSON input"**
- Input parsing failed - check JSON syntax
- Consider using raw string input instead

**"Task timeout"**
- Increase operator timeout configuration
- Optimize function execution time
- Check operator network connectivity

### Debug Logging

Enable verbose logging in the performer:
```bash
LOG_LEVEL=debug ./bin/performer
```

## Contributing

1. Fork the repository
2. Create feature branch
3. Add tests for new functionality
4. Submit pull request

## License

MIT License - see LICENSE file for details.
