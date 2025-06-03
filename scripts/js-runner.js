#!/usr/bin/env node

const fs = require('fs');
const path = require('path');

async function runFunction(functionDir, inputData) {
    let payload = null;
    let error = null;
    
    try {
        const indexPath = path.join(functionDir, 'index.js');
        
        if (!fs.existsSync(indexPath)) {
            throw new Error(`index.js not found in ${functionDir}`);
        }
        
        delete require.cache[require.resolve(indexPath)];
        
        // Capture console output during user function execution
        const originalConsole = console.log;
        const capturedLogs = [];
        console.log = (...args) => {
            capturedLogs.push(args.map(arg => String(arg)).join(' '));
        };
        
        try {
            const userFunction = require(indexPath);
            
            if (!userFunction.handler || typeof userFunction.handler !== 'function') {
                throw new Error('Function must export a handler function');
            }
            
            payload = await userFunction.handler(inputData);
            
        } finally {
            // Restore console.log
            console.log = originalConsole;
        }
        
        // Include captured logs in the result
        if (capturedLogs.length > 0) {
            payload = {
                result: payload,
                logs: capturedLogs
            };
        }
        
    } catch (err) {
        error = err.message;
    }
    
    const result = {
        payload: payload,
        error: error
    };
    
    console.log(JSON.stringify(result));
}

if (process.argv.length < 4) {
    console.log(JSON.stringify({
        payload: null,
        error: 'Usage: node js-runner.js <functionDir> <inputBase64>'
    }));
    process.exit(1);
}

const functionDir = process.argv[2];
const inputBase64 = process.argv[3];

let inputData;
try {
    // Decode base64 to get raw bytes
    const inputBuffer = Buffer.from(inputBase64, 'base64');
    
    // Try to parse as JSON first, but fall back to raw data
    try {
        inputData = JSON.parse(inputBuffer.toString('utf8'));
    } catch (jsonError) {
        // If it's not JSON, provide the raw buffer and string representation
        inputData = {
            raw: inputBuffer,
            text: inputBuffer.toString('utf8'),
            base64: inputBase64
        };
    }
} catch (error) {
    console.log(JSON.stringify({
        payload: null,
        error: 'Invalid base64 input: ' + error.message
    }));
    process.exit(1);
}

runFunction(functionDir, inputData);