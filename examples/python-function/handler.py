def handler(input_data):
    """
    Python function handler - equivalent to exports.handler in JavaScript
    
    Args:
        input_data: The input data from the function call
        
    Returns:
        The result of the function execution
    """
    print(f"Python function called with input: {input_data}")
    
    # Handle different input types
    if isinstance(input_data, dict):
        if 'name' in input_data:
            return {
                "message": f"Hello, {input_data['name']}! This is from Python.",
                "language": "python",
                "input_type": "dict",
                "input": input_data
            }
        elif 'raw' in input_data:
            # Binary input
            print(f"Received binary data, size: {len(input_data['raw'])}")
            return {
                "message": "Processed binary data with Python",
                "language": "python", 
                "input_type": "binary",
                "size": len(input_data['raw']),
                "text_preview": input_data.get('text', '')[:100]
            }
    elif isinstance(input_data, list):
        return {
            "message": "Processed array input with Python",
            "language": "python",
            "input_type": "array",
            "length": len(input_data),
            "items": input_data
        }
    else:
        # Simple string or other input
        return {
            "message": f"Echo from Python: {input_data}",
            "language": "python",
            "input_type": type(input_data).__name__,
            "input": input_data
        }