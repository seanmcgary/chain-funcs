#!/usr/bin/env node

const fs = require('fs');
const path = require('path');

async function runFunction(functionDir, inputArgs) {
    let payload = null;
    let error = null;
    
    try {
        const indexPath = path.join(functionDir, 'index.js');
        
        if (!fs.existsSync(indexPath)) {
            throw new Error(`index.js not found in ${functionDir}`);
        }
        
        delete require.cache[require.resolve(indexPath)];
        
        const userFunction = require(indexPath);
        
        if (!userFunction.handler || typeof userFunction.handler !== 'function') {
            throw new Error('Function must export a handler function');
        }
        
        payload = await userFunction.handler(inputArgs);
        
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
        error: 'Usage: node js-runner.js <functionDir> <inputJSON>'
    }));
    process.exit(1);
}

const functionDir = process.argv[2];
const inputJSON = process.argv[3];

let inputArgs;
try {
    inputArgs = JSON.parse(inputJSON);
} catch (error) {
    console.log(JSON.stringify({
        payload: null,
        error: 'Invalid JSON input: ' + error.message
    }));
    process.exit(1);
}

runFunction(functionDir, inputArgs);