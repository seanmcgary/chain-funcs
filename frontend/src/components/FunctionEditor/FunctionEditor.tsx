import { useState, useEffect } from 'react'
import WalletStatus from '../WalletStatus'
import ContractInfo from '../ContractInfo'
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Button } from "@/components/ui/button"
import { MonacoEditor } from "@/components/ui/monaco-editor"
import { FileTree, type FileNode } from "@/components/ui/file-tree"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Badge } from "@/components/ui/badge"
import { validateManifest, validateFunctionStructure, generateManifestTemplate, generateFunctionTemplate } from "@/lib/validation"
import { AlertTriangle, CheckCircle, GripHorizontal } from "lucide-react"

const FunctionEditor = () => {
  const [functionName, setFunctionName] = useState('hello-world')
  const [selectedFileId, setSelectedFileId] = useState('index.js')
  const [editorHeight, setEditorHeight] = useState(400)
  const [isResizing, setIsResizing] = useState(false)
  
  // Initialize with a basic JavaScript function structure
  const [files, setFiles] = useState<FileNode[]>([
    {
      id: 'manifest.json',
      name: 'manifest.json',
      type: 'file',
      content: `{
  "runtime": "javascript"
}`
    },
    {
      id: 'index.js',
      name: 'index.js',
      type: 'file',
      content: `exports.handler = async function(args = []) {
    const name = args[0] || "World";
    const count = args[1] || 1;
    
    const messages = [];
    for (let i = 0; i < count; i++) {
        messages.push(\`Hello, \${name}! (\${i + 1})\`);
    }
    
    return {
        message: "Function executed successfully",
        results: messages,
        timestamp: new Date().toISOString(),
        inputArgs: args
    };
};`
    }
  ])

  const selectedFile = files.find(f => f.id === selectedFileId)
  
  // Validation state
  const [validationResult, setValidationResult] = useState<{
    manifest: ReturnType<typeof validateManifest>
    structure: ReturnType<typeof validateFunctionStructure>
  } | null>(null)

  // Get current runtime from manifest
  const getCurrentRuntime = (): 'javascript' | 'python' => {
    const manifestFile = files.find(f => f.name === 'manifest.json')
    if (manifestFile?.content) {
      try {
        const manifest = JSON.parse(manifestFile.content)
        return manifest.runtime || 'javascript'
      } catch {
        return 'javascript'
      }
    }
    return 'javascript'
  }

  // Validate project when files change
  useEffect(() => {
    const manifestFile = files.find(f => f.name === 'manifest.json')
    const runtime = getCurrentRuntime()
    
    if (manifestFile?.content) {
      const manifestValidation = validateManifest(manifestFile.content)
      const structureValidation = validateFunctionStructure(files, runtime)
      
      setValidationResult({
        manifest: manifestValidation,
        structure: structureValidation
      })
    }
  }, [files])

  const handleFileSelect = (fileId: string) => {
    setSelectedFileId(fileId)
  }

  const handleFileCreate = (name: string, type: 'file' | 'folder', _parentId?: string) => {
    const newFile: FileNode = {
      id: `${Date.now()}-${name}`,
      name,
      type,
      content: type === 'file' ? '' : undefined,
      children: type === 'folder' ? [] : undefined
    }
    setFiles([...files, newFile])
  }

  const handleFileDelete = (fileId: string) => {
    setFiles(files.filter(f => f.id !== fileId))
    if (selectedFileId === fileId) {
      setSelectedFileId(files[0]?.id || '')
    }
  }

  const handleFileRename = (fileId: string, newName: string) => {
    setFiles(files.map(f => 
      f.id === fileId ? { ...f, name: newName } : f
    ))
  }

  const handleContentChange = (content: string) => {
    if (selectedFile) {
      setFiles(files.map(f => 
        f.id === selectedFileId ? { ...f, content } : f
      ))
    }
  }

  const getLanguageFromFilename = (filename: string): string => {
    const ext = filename.split('.').pop()?.toLowerCase()
    switch (ext) {
      case 'js':
      case 'mjs':
        return 'javascript'
      case 'py':
        return 'python'
      case 'json':
        return 'json'
      case 'md':
        return 'markdown'
      case 'txt':
        return 'plaintext'
      default:
        return 'plaintext'
    }
  }

  const createRuntimeTemplate = (runtime: 'javascript' | 'python') => {
    const manifestContent = generateManifestTemplate(runtime)
    const handlerContent = generateFunctionTemplate(runtime)
    
    const newFiles: FileNode[] = [
      {
        id: 'manifest.json',
        name: 'manifest.json',
        type: 'file',
        content: manifestContent
      }
    ]

    if (runtime === 'javascript') {
      newFiles.push({
        id: 'index.js',
        name: 'index.js',
        type: 'file',
        content: handlerContent
      })
    } else {
      newFiles.push(
        {
          id: 'handler.py',
          name: 'handler.py',
          type: 'file',
          content: handlerContent
        },
        {
          id: 'requirements.txt',
          name: 'requirements.txt',
          type: 'file',
          content: '# Add your Python dependencies here\n# Example:\n# requests==2.28.1\n'
        }
      )
    }

    setFiles(newFiles)
    setSelectedFileId(runtime === 'javascript' ? 'index.js' : 'handler.py')
  }

  const isProjectValid = validationResult ? 
    validationResult.manifest.isValid && validationResult.structure.isValid : false

  const allErrors = validationResult ? 
    [...validationResult.manifest.errors, ...validationResult.structure.errors] : []
  
  const allWarnings = validationResult ? 
    [...validationResult.manifest.warnings, ...validationResult.structure.warnings] : []

  // Handle resize functionality
  const handleMouseDown = (e: React.MouseEvent) => {
    setIsResizing(true)
    e.preventDefault()
  }

  const handleMouseMove = (e: MouseEvent) => {
    if (!isResizing) return
    
    const editorContainer = document.querySelector('.editor-container')
    if (!editorContainer) return
    
    const rect = editorContainer.getBoundingClientRect()
    const newHeight = Math.max(200, Math.min(800, e.clientY - rect.top - 50))
    setEditorHeight(newHeight)
  }

  const handleMouseUp = () => {
    setIsResizing(false)
  }

  useEffect(() => {
    if (isResizing) {
      document.addEventListener('mousemove', handleMouseMove)
      document.addEventListener('mouseup', handleMouseUp)
      document.body.style.cursor = 'ns-resize'
      document.body.style.userSelect = 'none'
    }

    return () => {
      document.removeEventListener('mousemove', handleMouseMove)
      document.removeEventListener('mouseup', handleMouseUp)
      document.body.style.cursor = ''
      document.body.style.userSelect = ''
    }
  }, [isResizing])

  // Keyboard shortcuts for resizing
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.ctrlKey || e.metaKey) {
        if (e.key === '+' || e.key === '=') {
          e.preventDefault()
          setEditorHeight(prev => Math.min(800, prev + 50))
        } else if (e.key === '-') {
          e.preventDefault()
          setEditorHeight(prev => Math.max(200, prev - 50))
        }
      }
    }

    document.addEventListener('keydown', handleKeyDown)
    return () => document.removeEventListener('keydown', handleKeyDown)
  }, [])

  return (
    <div className="h-full bg-background flex flex-col">
      {/* Main Content Area */}
      <div className="flex-1 flex flex-col overflow-hidden">
        {/* Header */}
        <div className="border-b bg-card p-4 flex-shrink-0">
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-2xl font-bold">Function Editor</h1>
              <p className="text-muted-foreground">
                Create and edit JavaScript/Python functions for deployment on the ChainFuncs network
              </p>
            </div>
          </div>
        </div>
        
        {/* Function Controls */}
        <div className="border-b bg-card p-4 flex-shrink-0">
          <div className="flex items-center justify-between">
            <div className="space-y-2">
              <Label htmlFor="function-name">Function Name</Label>
              <Input 
                id="function-name" 
                value={functionName}
                onChange={(e) => setFunctionName(e.target.value)}
                placeholder="my-awesome-function" 
                className="max-w-sm" 
              />
            </div>
            
            <div className="flex items-center space-x-2">
              <Badge variant={isProjectValid ? "default" : "destructive"}>
                {isProjectValid ? (
                  <>
                    <CheckCircle className="h-3 w-3 mr-1" />
                    Valid
                  </>
                ) : (
                  <>
                    <AlertTriangle className="h-3 w-3 mr-1" />
                    {allErrors.length} Error{allErrors.length !== 1 ? 's' : ''}
                  </>
                )}
              </Badge>
              
              <div className="flex space-x-1">
                <Button
                  size="sm"
                  variant="outline"
                  onClick={() => createRuntimeTemplate('javascript')}
                >
                  JS Template
                </Button>
                <Button
                  size="sm"
                  variant="outline"
                  onClick={() => createRuntimeTemplate('python')}
                >
                  Python Template
                </Button>
              </div>
            </div>
          </div>
        </div>

        {/* Scrollable Content Area */}
        <div className="flex-1 flex overflow-hidden">
          {/* File Tree Sidebar */}
          <div className="w-80 border-r bg-card flex flex-col">
            <div className="p-4 border-b flex-shrink-0">
              <Label className="text-sm font-medium">Project Files</Label>
            </div>
            <div className="flex-1 overflow-hidden">
              <FileTree
                files={files}
                selectedFile={selectedFileId}
                onFileSelect={handleFileSelect}
                onFileCreate={handleFileCreate}
                onFileDelete={handleFileDelete}
                onFileRename={handleFileRename}
                className="h-full border-0 rounded-none bg-transparent"
              />
            </div>
          </div>
          
          {/* Code Editor */}
          <div className="flex-1 flex flex-col overflow-hidden">
            {/* Validation Messages - Scrollable */}
            <div className="flex-shrink-0 overflow-y-auto max-h-32">
              {allErrors.length > 0 && (
                <div className="p-4 border-b">
                  <Alert variant="destructive">
                    <AlertTriangle className="h-4 w-4" />
                    <AlertDescription>
                      <div className="space-y-1">
                        {allErrors.map((error, index) => (
                          <div key={index}>
                            <strong>{error.field}:</strong> {error.message}
                          </div>
                        ))}
                      </div>
                    </AlertDescription>
                  </Alert>
                </div>
              )}

              {allWarnings.length > 0 && (
                <div className="p-4 border-b">
                  <Alert>
                    <AlertTriangle className="h-4 w-4" />
                    <AlertDescription>
                      <div className="space-y-1">
                        {allWarnings.map((warning, index) => (
                          <div key={index}>
                            <strong>{warning.field}:</strong> {warning.message}
                          </div>
                        ))}
                      </div>
                    </AlertDescription>
                  </Alert>
                </div>
              )}
            </div>

            <div className="p-4 border-b bg-card flex-shrink-0">
              <div className="flex items-center justify-between">
                <Label htmlFor="function-code">
                  {selectedFile ? `Editing: ${selectedFile.name}` : 'Select a file to edit'}
                </Label>
                <div className="text-xs text-muted-foreground">
                  Height: {editorHeight}px (drag to resize, or Ctrl/Cmd +/-)
                </div>
              </div>
            </div>
            
            <div className="flex-1 relative overflow-hidden">
              {selectedFile ? (
                <div className="h-full relative">
                  <MonacoEditor
                    value={selectedFile.content || ''}
                    onChange={handleContentChange}
                    language={getLanguageFromFilename(selectedFile.name)}
                    className="border-0 rounded-none"
                    style={{ height: `${editorHeight}px` }}
                    placeholder={`// Enter your ${getLanguageFromFilename(selectedFile.name)} code here...`}
                  />
                  <div 
                    className="absolute bottom-0 left-0 right-0 h-3 bg-border/50 cursor-ns-resize flex items-center justify-center hover:bg-border transition-colors"
                    onMouseDown={handleMouseDown}
                  >
                    <GripHorizontal className="h-3 w-3 text-muted-foreground" />
                  </div>
                </div>
              ) : (
                <div 
                  className="h-full flex items-center justify-center text-muted-foreground bg-muted/20"
                >
                  Select a file to start editing
                </div>
              )}
            </div>
          </div>
          
          {/* Right Sidebar */}
          <div className="w-80 border-l bg-card flex flex-col">
            <div className="flex-1 overflow-y-auto p-4 space-y-6">
              <WalletStatus />
              <ContractInfo />
            </div>
          </div>
        </div>
      </div>

      {/* Sticky Action Bar */}
      <div className="border-t bg-card p-4 flex-shrink-0 shadow-lg">
        <div className="flex gap-4">
          <Button>Save Function</Button>
          <Button variant="outline">Test Locally</Button>
          <Button variant="outline">Download Project</Button>
        </div>
      </div>
    </div>
  )
}

export default FunctionEditor