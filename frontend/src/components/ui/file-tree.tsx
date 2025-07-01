import { useState } from 'react'
import { ChevronRight, ChevronDown, File, Folder, Plus, Trash2, Edit } from 'lucide-react'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { cn } from "@/lib/utils"

export interface FileNode {
  id: string
  name: string
  type: 'file' | 'folder'
  content?: string
  children?: FileNode[]
  parent?: string
}

interface FileTreeProps {
  files: FileNode[]
  selectedFile?: string
  onFileSelect: (fileId: string) => void
  onFileCreate: (name: string, type: 'file' | 'folder', parentId?: string) => void
  onFileDelete: (fileId: string) => void
  onFileRename: (fileId: string, newName: string) => void
  className?: string
  style?: React.CSSProperties
}

interface FileTreeNodeProps {
  node: FileNode
  level: number
  selectedFile?: string
  onFileSelect: (fileId: string) => void
  onFileCreate: (name: string, type: 'file' | 'folder', parentId?: string) => void
  onFileDelete: (fileId: string) => void
  onFileRename: (fileId: string, newName: string) => void
  expandedNodes: Set<string>
  onToggleExpand: (nodeId: string) => void
}

function FileTreeNode({
  node,
  level,
  selectedFile,
  onFileSelect,
  onFileCreate,
  onFileDelete,
  onFileRename,
  expandedNodes,
  onToggleExpand
}: FileTreeNodeProps) {
  const [isRenaming, setIsRenaming] = useState(false)
  const [renameValue, setRenameValue] = useState(node.name)
  const [showActions, setShowActions] = useState(false)

  const isExpanded = expandedNodes.has(node.id)
  const isSelected = selectedFile === node.id
  const hasChildren = node.children && node.children.length > 0

  const handleToggleExpand = () => {
    if (node.type === 'folder') {
      onToggleExpand(node.id)
    }
  }

  const handleRename = () => {
    if (renameValue.trim() && renameValue !== node.name) {
      onFileRename(node.id, renameValue.trim())
    }
    setIsRenaming(false)
    setRenameValue(node.name)
  }

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter') {
      handleRename()
    } else if (e.key === 'Escape') {
      setIsRenaming(false)
      setRenameValue(node.name)
    }
  }

  const paddingLeft = level * 12 + 8

  return (
    <div>
      <div
        className={cn(
          "flex items-center h-8 hover:bg-accent/50 cursor-pointer group relative",
          isSelected && "bg-accent"
        )}
        style={{ paddingLeft: `${paddingLeft}px` }}
        onClick={() => {
          if (node.type === 'file') {
            onFileSelect(node.id)
          } else {
            handleToggleExpand()
          }
        }}
        onMouseEnter={() => setShowActions(true)}
        onMouseLeave={() => setShowActions(false)}
      >
        {node.type === 'folder' && (
          <button
            onClick={(e) => {
              e.stopPropagation()
              handleToggleExpand()
            }}
            className="mr-1 p-0.5 hover:bg-accent rounded"
          >
            {isExpanded ? (
              <ChevronDown className="h-3 w-3" />
            ) : (
              <ChevronRight className="h-3 w-3" />
            )}
          </button>
        )}
        
        {node.type === 'file' ? (
          <File className="h-4 w-4 mr-2 text-blue-500" />
        ) : (
          <Folder className="h-4 w-4 mr-2 text-yellow-500" />
        )}

        {isRenaming ? (
          <Input
            value={renameValue}
            onChange={(e) => setRenameValue(e.target.value)}
            onBlur={handleRename}
            onKeyDown={handleKeyDown}
            className="h-6 text-xs border-0 p-1 bg-transparent"
            autoFocus
            onClick={(e) => e.stopPropagation()}
          />
        ) : (
          <span className="text-sm truncate flex-1">{node.name}</span>
        )}

        {showActions && !isRenaming && (
          <div className="flex items-center space-x-1 opacity-0 group-hover:opacity-100 transition-opacity">
            {node.type === 'folder' && (
              <>
                <Button
                  size="sm"
                  variant="ghost"
                  className="h-6 w-6 p-0"
                  onClick={(e) => {
                    e.stopPropagation()
                    const fileName = prompt('Enter file name:')
                    if (fileName) {
                      onFileCreate(fileName, 'file', node.id)
                    }
                  }}
                >
                  <Plus className="h-3 w-3" />
                </Button>
              </>
            )}
            <Button
              size="sm"
              variant="ghost"
              className="h-6 w-6 p-0"
              onClick={(e) => {
                e.stopPropagation()
                setIsRenaming(true)
              }}
            >
              <Edit className="h-3 w-3" />
            </Button>
            {node.name !== 'manifest.json' && node.name !== 'index.js' && node.name !== 'handler.py' && (
              <Button
                size="sm"
                variant="ghost"
                className="h-6 w-6 p-0 text-destructive hover:text-destructive"
                onClick={(e) => {
                  e.stopPropagation()
                  if (confirm(`Delete ${node.name}?`)) {
                    onFileDelete(node.id)
                  }
                }}
              >
                <Trash2 className="h-3 w-3" />
              </Button>
            )}
          </div>
        )}
      </div>

      {node.type === 'folder' && isExpanded && hasChildren && (
        <div>
          {node.children!.map((child) => (
            <FileTreeNode
              key={child.id}
              node={child}
              level={level + 1}
              selectedFile={selectedFile}
              onFileSelect={onFileSelect}
              onFileCreate={onFileCreate}
              onFileDelete={onFileDelete}
              onFileRename={onFileRename}
              expandedNodes={expandedNodes}
              onToggleExpand={onToggleExpand}
            />
          ))}
        </div>
      )}
    </div>
  )
}

export function FileTree({
  files,
  selectedFile,
  onFileSelect,
  onFileCreate,
  onFileDelete,
  onFileRename,
  className,
  style
}: FileTreeProps) {
  const [expandedNodes, setExpandedNodes] = useState<Set<string>>(new Set(['root']))

  const handleToggleExpand = (nodeId: string) => {
    const newExpanded = new Set(expandedNodes)
    if (newExpanded.has(nodeId)) {
      newExpanded.delete(nodeId)
    } else {
      newExpanded.add(nodeId)
    }
    setExpandedNodes(newExpanded)
  }

  return (
    <div className={cn("border border-border rounded-md bg-background flex flex-col", className)} style={style}>
      <div className="p-2 border-b border-border bg-muted/50">
        <div className="flex items-center justify-between">
          <span className="text-sm font-medium">Project Files</span>
          <div className="flex space-x-1">
            <Button
              size="sm"
              variant="ghost"
              className="h-6 w-6 p-0"
              onClick={() => {
                const fileName = prompt('Enter file name:')
                if (fileName) {
                  onFileCreate(fileName, 'file')
                }
              }}
            >
              <Plus className="h-3 w-3" />
            </Button>
          </div>
        </div>
      </div>
      <div className="flex-1 overflow-y-auto" style={{ maxHeight: 'calc(100% - 60px)' }}>
        {files.map((node) => (
          <FileTreeNode
            key={node.id}
            node={node}
            level={0}
            selectedFile={selectedFile}
            onFileSelect={onFileSelect}
            onFileCreate={onFileCreate}
            onFileDelete={onFileDelete}
            onFileRename={onFileRename}
            expandedNodes={expandedNodes}
            onToggleExpand={handleToggleExpand}
          />
        ))}
      </div>
    </div>
  )
}