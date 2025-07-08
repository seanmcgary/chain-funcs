import { useState } from 'react'
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Archive, File, AlertCircle, CheckCircle2 } from "lucide-react"
import { type FileNode } from "@/components/ui/file-tree"
import { createTarball, previewTarball, validateTarball, calculateFunctionId } from "@/lib/tarball"

interface TarballPreviewProps {
  files: FileNode[]
  functionName: string
}

export function TarballPreview({ files, functionName }: TarballPreviewProps) {
  const [isOpen, setIsOpen] = useState(false)
  const [preview, setPreview] = useState<{ files: string[], totalSize: number } | null>(null)
  const [functionId, setFunctionId] = useState<string>('')
  const [loading, setLoading] = useState(false)

  const validation = validateTarball(files)

  const generatePreview = async () => {
    setLoading(true)
    try {
      const tarballData = await createTarball(files)
      const previewData = previewTarball(tarballData)
      setPreview(previewData)
      
      const id = calculateFunctionId(tarballData)
      setFunctionId(id)
    } catch (error) {
      console.error('Failed to generate preview:', error)
    } finally {
      setLoading(false)
    }
  }

  const formatFileSize = (bytes: number): string => {
    if (bytes < 1024) return `${bytes} B`
    if (bytes < 1024 * 1024) return `${(bytes / 1024).toFixed(2)} KB`
    return `${(bytes / (1024 * 1024)).toFixed(2)} MB`
  }

  return (
    <Dialog open={isOpen} onOpenChange={setIsOpen}>
      <DialogTrigger asChild>
        <Button 
          variant="outline" 
          size="sm"
          onClick={() => {
            setIsOpen(true)
            generatePreview()
          }}
        >
          <Archive className="h-4 w-4 mr-1" />
          Preview Tarball
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-2xl">
        <DialogHeader>
          <DialogTitle>Tarball Preview</DialogTitle>
          <DialogDescription>
            Preview the deployment tarball structure for {functionName}
          </DialogDescription>
        </DialogHeader>
        
        <div className="space-y-4">
          {/* Validation Status */}
          <div>
            <h4 className="text-sm font-medium mb-2">Validation Status</h4>
            {validation.valid ? (
              <Alert>
                <CheckCircle2 className="h-4 w-4" />
                <AlertDescription>
                  Function structure is valid and ready for deployment
                </AlertDescription>
              </Alert>
            ) : (
              <Alert variant="destructive">
                <AlertCircle className="h-4 w-4" />
                <AlertDescription>
                  <div className="space-y-1">
                    {validation.errors.map((error, index) => (
                      <div key={index}>{error}</div>
                    ))}
                  </div>
                </AlertDescription>
              </Alert>
            )}
          </div>

          {/* Function ID */}
          {functionId && (
            <div>
              <h4 className="text-sm font-medium mb-2">Function ID</h4>
              <div className="flex items-center gap-2">
                <code className="text-xs bg-muted px-2 py-1 rounded font-mono">
                  {functionId}
                </code>
                <Badge variant="outline" className="text-xs">
                  {functionId.slice(0, 8)}...
                </Badge>
              </div>
            </div>
          )}

          {/* Tarball Contents */}
          {loading ? (
            <div className="text-center py-8 text-muted-foreground">
              Generating preview...
            </div>
          ) : preview ? (
            <div>
              <div className="flex items-center justify-between mb-2">
                <h4 className="text-sm font-medium">Tarball Contents</h4>
                <span className="text-xs text-muted-foreground">
                  Total size: {formatFileSize(preview.totalSize)}
                </span>
              </div>
              <div className="border rounded-md p-4 bg-muted/20 max-h-64 overflow-y-auto">
                {preview.files.length === 0 ? (
                  <p className="text-sm text-muted-foreground">No files in tarball</p>
                ) : (
                  <ul className="space-y-1">
                    {preview.files.map((file, index) => (
                      <li key={index} className="flex items-center gap-2 text-sm">
                        <File className="h-3 w-3 text-muted-foreground" />
                        <span className="font-mono">{file}</span>
                      </li>
                    ))}
                  </ul>
                )}
              </div>
            </div>
          ) : null}

          {/* Instructions */}
          <div className="border-t pt-4">
            <h4 className="text-sm font-medium mb-2">Deployment Instructions</h4>
            <ol className="text-sm text-muted-foreground space-y-1 list-decimal list-inside">
              <li>Click "Deploy Tarball" to download the .tar.gz file</li>
              <li>Use the ChainFuncs CLI to deploy: <code className="text-xs bg-muted px-1 rounded">faas-cli deploy-function tarball.tar.gz</code></li>
              <li>Or upload to S3/IPFS and deploy with the function ID</li>
            </ol>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  )
}