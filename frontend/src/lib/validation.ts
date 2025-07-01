// Validation utilities for ChainFuncs function projects

export interface ManifestSchema {
  runtime: 'javascript' | 'python'
  description?: string
  version?: string
  author?: string
  dependencies?: Record<string, string>
}

export interface ValidationError {
  field: string
  message: string
}

export interface ValidationResult {
  isValid: boolean
  errors: ValidationError[]
  warnings: ValidationError[]
}

export function validateManifest(content: string): ValidationResult {
  const errors: ValidationError[] = []
  const warnings: ValidationError[] = []

  try {
    const manifest = JSON.parse(content) as any

    // Required fields validation
    if (!manifest.runtime) {
      errors.push({
        field: 'runtime',
        message: 'Runtime field is required'
      })
    } else if (!['javascript', 'python'].includes(manifest.runtime)) {
      errors.push({
        field: 'runtime',
        message: 'Runtime must be either "javascript" or "python"'
      })
    }

    // Optional field validation
    if (manifest.description && typeof manifest.description !== 'string') {
      errors.push({
        field: 'description',
        message: 'Description must be a string'
      })
    }

    if (manifest.version && typeof manifest.version !== 'string') {
      errors.push({
        field: 'version',
        message: 'Version must be a string'
      })
    }

    if (manifest.author && typeof manifest.author !== 'string') {
      errors.push({
        field: 'author',
        message: 'Author must be a string'
      })
    }

    if (manifest.dependencies && typeof manifest.dependencies !== 'object') {
      errors.push({
        field: 'dependencies',
        message: 'Dependencies must be an object'
      })
    }

    // Warnings for missing recommended fields
    if (!manifest.description) {
      warnings.push({
        field: 'description',
        message: 'Consider adding a description for your function'
      })
    }

    if (!manifest.version) {
      warnings.push({
        field: 'version',
        message: 'Consider adding a version number (e.g., "1.0.0")'
      })
    }

  } catch (e) {
    errors.push({
      field: 'json',
      message: 'Invalid JSON format'
    })
  }

  return {
    isValid: errors.length === 0,
    errors,
    warnings
  }
}

export function validateFunctionStructure(files: Array<{name: string, content?: string}>, runtime: string): ValidationResult {
  const errors: ValidationError[] = []
  const warnings: ValidationError[] = []

  // Check for required files based on runtime
  const hasManifest = files.some(f => f.name === 'manifest.json')
  if (!hasManifest) {
    errors.push({
      field: 'manifest.json',
      message: 'manifest.json file is required'
    })
  }

  if (runtime === 'javascript') {
    const hasIndexJs = files.some(f => f.name === 'index.js')
    if (!hasIndexJs) {
      errors.push({
        field: 'index.js',
        message: 'index.js file is required for JavaScript functions'
      })
    } else {
      // Validate index.js content
      const indexFile = files.find(f => f.name === 'index.js')
      if (indexFile?.content) {
        if (!indexFile.content.includes('exports.handler')) {
          errors.push({
            field: 'index.js',
            message: 'index.js must export a handler function (exports.handler = ...)'
          })
        }
      }
    }
  } else if (runtime === 'python') {
    const hasHandlerPy = files.some(f => f.name === 'handler.py')
    if (!hasHandlerPy) {
      errors.push({
        field: 'handler.py',
        message: 'handler.py file is required for Python functions'
      })
    } else {
      // Validate handler.py content
      const handlerFile = files.find(f => f.name === 'handler.py')
      if (handlerFile?.content) {
        if (!handlerFile.content.includes('def handler(')) {
          errors.push({
            field: 'handler.py',
            message: 'handler.py must define a handler function (def handler(args):)'
          })
        }
      }
    }

    // Check for requirements.txt
    const hasRequirements = files.some(f => f.name === 'requirements.txt')
    if (!hasRequirements) {
      warnings.push({
        field: 'requirements.txt',
        message: 'Consider adding requirements.txt for Python dependencies'
      })
    }
  }

  return {
    isValid: errors.length === 0,
    errors,
    warnings
  }
}

export function generateManifestTemplate(runtime: 'javascript' | 'python'): string {
  const template: ManifestSchema = {
    runtime,
    description: 'My awesome ChainFuncs function',
    version: '1.0.0',
    author: 'Your Name'
  }

  if (runtime === 'python') {
    template.dependencies = {}
  }

  return JSON.stringify(template, null, 2)
}

export function generateFunctionTemplate(runtime: 'javascript' | 'python'): string {
  if (runtime === 'javascript') {
    return `exports.handler = async function(args = []) {
    // Your function logic here
    const input = args[0] || "Hello";
    
    return {
        message: "Function executed successfully",
        result: \`Processed: \${input}\`,
        timestamp: new Date().toISOString(),
        inputArgs: args
    };
};`
  } else {
    return `def handler(args):
    """
    ChainFuncs Python handler function
    
    Args:
        args (list): Input arguments from the function call
        
    Returns:
        dict: Result object that will be returned to the caller
    """
    input_value = args[0] if args else "Hello"
    
    return {
        "message": "Function executed successfully",
        "result": f"Processed: {input_value}",
        "timestamp": __import__('datetime').datetime.now().isoformat(),
        "inputArgs": args
    }`
  }
}