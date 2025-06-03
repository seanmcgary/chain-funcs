#!/usr/bin/env python3

import sys
import os
import json
import base64
import traceback
from io import StringIO
import importlib.util


def run_function(function_dir, input_data):
    payload = None
    error = None
    captured_logs = []
    
    try:
        # Add function directory to Python path so pip-installed packages can be found
        if function_dir not in sys.path:
            sys.path.insert(0, function_dir)
        
        handler_path = os.path.join(function_dir, 'handler.py')
        
        if not os.path.exists(handler_path):
            raise Exception(f"handler.py not found in {function_dir}")
        
        # Load the user's handler module
        spec = importlib.util.spec_from_file_location("user_handler", handler_path)
        user_handler = importlib.util.module_from_spec(spec)
        
        # Capture stdout during user function execution
        original_stdout = sys.stdout
        captured_output = StringIO()
        
        try:
            sys.stdout = captured_output
            spec.loader.exec_module(user_handler)
            
            if not hasattr(user_handler, 'handler') or not callable(getattr(user_handler, 'handler')):
                raise Exception('Function must export a handler function')
            
            payload = user_handler.handler(input_data)
            
        finally:
            # Restore stdout
            sys.stdout = original_stdout
        
        # Get captured output and split into lines
        output = captured_output.getvalue()
        if output.strip():
            captured_logs = output.strip().split('\n')
        
        # Include captured logs in the result if any
        if captured_logs:
            payload = {
                "result": payload,
                "logs": captured_logs
            }
            
    except Exception as err:
        error = str(err)
        # Also capture traceback for debugging
        if "--debug" in sys.argv:
            error = traceback.format_exc()
    
    result = {
        "payload": payload,
        "error": error
    }
    
    print(json.dumps(result))


def main():
    if len(sys.argv) < 3:
        print(json.dumps({
            "payload": None,
            "error": "Usage: python3 py-runner.py <functionDir> <inputBase64>"
        }))
        sys.exit(1)
    
    function_dir = sys.argv[1]
    input_base64 = sys.argv[2]
    
    try:
        # Decode base64 to get raw bytes
        input_buffer = base64.b64decode(input_base64)
        
        # Try to parse as JSON first, but fall back to raw data
        try:
            input_data = json.loads(input_buffer.decode('utf-8'))
        except (json.JSONDecodeError, UnicodeDecodeError):
            # If it's not JSON, provide the raw buffer and string representation
            input_data = {
                "raw": list(input_buffer),  # Convert bytes to list for JSON serialization
                "text": input_buffer.decode('utf-8', errors='replace'),
                "base64": input_base64
            }
    except Exception as error:
        print(json.dumps({
            "payload": None,
            "error": f"Invalid base64 input: {str(error)}"
        }))
        sys.exit(1)
    
    run_function(function_dir, input_data)


if __name__ == "__main__":
    main()