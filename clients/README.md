# FaaS CLI

A command-line interface for interacting with the FaaS (Function-as-a-Service) AVS contract.

## Installation

```bash
cd clients
make build
```

The binary will be created at `./bin/faas-cli`.

## Configuration

Copy the example environment file and configure your settings:

```bash
cp .env.example .env
# Edit .env with your actual values
```

Or set environment variables directly:

```bash
export RPC_URL="http://localhost:8545"
export PRIVATE_KEY="your_private_key_here"
export FAAS_ADDRESS="0x1234567890123456789012345678901234567890"
export TASKMAILBOX_ADDRESS="0x0987654321098765432109876543210987654321"
```

## Usage

### Deploy a Function

Deploy a JavaScript function from a directory:

```bash
./bin/faas-cli deploy-function /path/to/function/directory
```

The function directory should contain:
- `index.js` - Main function file with `exports.handler = async function(args) { ... }`
- Any additional files or node_modules needed

Example:
```bash
./bin/faas-cli deploy-function ../examples/hello-function
```

### Call a Function

Execute a deployed function:

```bash
./bin/faas-cli call-function --function-id=0x1234... "Hello World"
```

This will:
1. Call the function with the input string as the first argument
2. Wait for the transaction to be mined
3. Poll the TaskMailbox for the result
4. Display the execution result

## Command Reference

### deploy-function

```bash
./bin/faas-cli deploy-function [options] <function-directory>
```

**Options:**
- `--rpc-url` - Ethereum RPC URL (default: http://localhost:8545)
- `--private-key` - Private key for signing transactions
- `--faas-address` - FaaS contract address

**Returns:**
- Function ID (32-byte hex string) for use with call-function

### call-function

```bash
./bin/faas-cli call-function [options] --function-id=<id> <input-string>
```

**Options:**
- `--rpc-url` - Ethereum RPC URL (default: http://localhost:8545)
- `--private-key` - Private key for signing transactions
- `--faas-address` - FaaS contract address
- `--taskMailbox-address` - TaskMailbox contract address
- `--function-id` - Function ID returned from deploy-function

**Arguments:**
- `<input-string>` - String input to pass to the function (becomes args[0])

## Example Workflow

1. **Deploy a function:**
   ```bash
   ./bin/faas-cli deploy-function ../examples/hello-function
   # Output: Function ID: 0xabc123...
   ```

2. **Call the function:**
   ```bash
   ./bin/faas-cli call-function --function-id=0xabc123... "Alice"
   # Output: Execution result
   ```

3. **Check results:**
   The CLI will automatically poll for results and display them when ready.
   
   **Success Example:**
   ```json
   Function Result:
   {
     "message": "Hello, Alice!",
     "timestamp": "2024-01-01T12:00:00.000Z"
   }
   ```
   
   **Error Example:**
   ```
   Function Error: Invalid input provided
   ```

## Function Structure

JavaScript functions must follow this structure:

```javascript
exports.handler = async function(args = []) {
    // args[0] contains the input string from CLI
    const input = args[0] || "default";
    
    // Your function logic here
    return {
        message: `Hello, ${input}!`,
        timestamp: new Date().toISOString()
    };
};
```

## Error Handling

The CLI provides detailed error messages for common issues:
- Network connection problems
- Invalid private keys
- Contract interaction failures
- Function execution errors
- Transaction failures

## Environment Variables

All flags can be set via environment variables:
- `RPC_URL`
- `PRIVATE_KEY`
- `FAAS_ADDRESS`
- `TASKMAILBOX_ADDRESS`