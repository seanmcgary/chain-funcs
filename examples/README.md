# FaaS JavaScript Functions Examples

This directory contains example JavaScript functions that can be executed by the FaaS AVS.

## Function Structure

All JavaScript functions must:
1. Be packaged as a gzipped tarball
2. Contain an `index.js` file at the root
3. Export a `handler` function that accepts an array of arguments

```javascript
exports.handler = async function(args = []) {
    // Your function logic here
    return result;
};
```

## Creating and Deploying a Function

### 1. Create a tarball of your function

```bash
# Navigate to your function directory (e.g., hello-function)
cd examples/hello-function

# Create a gzipped tarball
tar -czf function.tar.gz index.js

# Convert to hex for contract interaction
xxd -p function.tar.gz | tr -d '\n' > function.hex
```

### 2. Register the function on-chain

```solidity
// Using the FaaS contract
bytes memory functionContent = hex"<content_from_function.hex>";
bytes32 functionId = faas.registerFunction(functionContent);
```

### 3. Call the function

```solidity
// Prepare arguments as JSON array
bytes memory arguments = abi.encode(["Alice", 3]);
bytes32 taskId = faas.callFunction(functionId, arguments);
```

### 4. Retrieve results

Results can be retrieved from the TaskMailbox using the returned `taskId`.

## Example Functions

### hello-function
- **Purpose**: Simple greeting function
- **Arguments**: `[name, count]`
- **Returns**: Array of greeting messages

### math-function  
- **Purpose**: Performs mathematical operations
- **Arguments**: `[operation, ...numbers]`
- **Operations**: `add`, `multiply`, `max`, `min`, `average`
- **Returns**: Calculation result with metadata

## Function Guidelines

1. **Error Handling**: Use try/catch and throw meaningful errors
2. **Return Values**: Return JSON-serializable objects (becomes `payload` in result)
3. **Dependencies**: Include all required `node_modules` in your tarball
4. **Timeouts**: Functions should complete within reasonable time limits
5. **Security**: Avoid accessing external resources or file system

## Result Structure

The AVS returns results in this format:
```json
{
  "payload": <your_function_return_value>,
  "error": null | "error_message"
}
```

- **Success**: `payload` contains your function's return value, `error` is `null`
- **Error**: `payload` is `null`, `error` contains the error message