import requests
import json
from datetime import datetime

def handler(input_data):
    """
    Python function handler with external dependencies
    
    Args:
        input_data: The input data from the function call
        
    Returns:
        The result of the function execution with API data
    """
    print(f"Python function with dependencies called with: {input_data}")
    
    try:
        # Make an HTTP request using the requests library
        response = requests.get('https://httpbin.org/json', timeout=5)
        api_data = response.json()
        
        # Process the input and combine with API data
        result = {
            "message": "Python function with external dependencies",
            "timestamp": datetime.now().isoformat(),
            "input": input_data,
            "api_response": {
                "status_code": response.status_code,
                "data": api_data
            },
            "dependencies_loaded": {
                "requests": True,
                "datetime": True
            }
        }
        
        return result
        
    except requests.RequestException as e:
        return {
            "error": f"API request failed: {str(e)}",
            "input": input_data,
            "dependencies_loaded": True
        }
    except Exception as e:
        return {
            "error": f"Function execution failed: {str(e)}",
            "input": input_data
        }