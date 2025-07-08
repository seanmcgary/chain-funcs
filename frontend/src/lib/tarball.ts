import { type FileNode } from "@/components/ui/file-tree"
import pako from 'pako'
import { keccak256 } from 'viem'

/**
 * TAR format utilities for creating CLI-compatible tarballs
 * Based on the CLI's createTarball function in clients/commands.go
 */

interface TarHeader {
  name: string
  mode: number
  uid: number
  gid: number
  size: number
  mtime: number
  checksum: number
  typeflag: string
  linkname: string
  magic: string
  version: string
  uname: string
  gname: string
  devmajor: number
  devminor: number
  prefix: string
}

/**
 * Create a TAR header for a file
 */
function createTarHeader(fileName: string, fileSize: number, isDirectory: boolean = false): Uint8Array {
  const header = new Uint8Array(512)
  const encoder = new TextEncoder()
  
  // Helper function to write string at specific position
  const writeString = (str: string, offset: number, maxLength: number) => {
    const bytes = encoder.encode(str)
    const length = Math.min(bytes.length, maxLength)
    for (let i = 0; i < length; i++) {
      header[offset + i] = bytes[i]
    }
  }
  
  // File name (100 bytes)
  writeString(fileName, 0, 100)
  
  // File mode (8 bytes) - 0644 for files, 0755 for directories
  const mode = isDirectory ? '0000755' : '0000644'
  writeString(mode, 100, 8)
  
  // UID (8 bytes)
  writeString('0000000', 108, 8)
  
  // GID (8 bytes)
  writeString('0000000', 116, 8)
  
  // File size (12 bytes) - in octal
  const sizeOctal = fileSize.toString(8).padStart(11, '0')
  writeString(sizeOctal, 124, 12)
  
  // Modification time (12 bytes) - current time in octal
  const mtime = Math.floor(Date.now() / 1000)
  const mtimeOctal = mtime.toString(8).padStart(11, '0')
  writeString(mtimeOctal, 136, 12)
  
  // Checksum placeholder (8 bytes) - fill with spaces initially
  writeString('        ', 148, 8)
  
  // Type flag (1 byte) - '0' for regular file, '5' for directory
  writeString(isDirectory ? '5' : '0', 156, 1)
  
  // Link name (100 bytes) - empty
  
  // Magic (6 bytes) - "ustar" with null terminator
  writeString('ustar\0', 257, 6)
  
  // Version (2 bytes)
  writeString('00', 263, 2)
  
  // User name (32 bytes)
  writeString('user', 265, 32)
  
  // Group name (32 bytes)
  writeString('group', 297, 32)
  
  // Calculate checksum
  let checksum = 0
  for (let i = 0; i < 512; i++) {
    checksum += header[i]
  }
  
  // Write checksum (subtract the spaces we added as placeholder)
  checksum -= 8 * 32 // 8 spaces * ASCII value of space (32)
  const checksumOctal = checksum.toString(8).padStart(6, '0')
  writeString(checksumOctal + '\0 ', 148, 8)
  
  return header
}

/**
 * Create a tarball from FileNode array (matching CLI format exactly)
 */
export async function createTarball(files: FileNode[]): Promise<Uint8Array> {
  const chunks: Uint8Array[] = []
  
  // Process each file
  for (const file of files) {
    if (file.type === 'file' && file.content !== undefined) {
      // Convert string content to bytes
      const encoder = new TextEncoder()
      const contentBytes = encoder.encode(file.content)
      
      // Create header for this file
      const header = createTarHeader(file.name, contentBytes.length)
      chunks.push(new Uint8Array(header))
      
      // Add file content
      chunks.push(contentBytes)
      
      // Add padding to make it a multiple of 512 bytes
      const padding = 512 - (contentBytes.length % 512)
      if (padding < 512) {
        chunks.push(new Uint8Array(padding))
      }
    }
  }
  
  // Add two 512-byte blocks of zeros to mark end of archive
  chunks.push(new Uint8Array(1024))
  
  // Combine all chunks into a single buffer
  const totalLength = chunks.reduce((sum, chunk) => sum + chunk.length, 0)
  const tarBuffer = new Uint8Array(totalLength)
  let offset = 0
  
  for (const chunk of chunks) {
    tarBuffer.set(chunk, offset)
    offset += chunk.length
  }
  
  // Gzip compress the tar archive
  const compressed = pako.gzip(tarBuffer)
  
  return compressed
}

/**
 * Calculate function ID from tarball (Keccak256 hash)
 */
export function calculateFunctionId(tarballData: Uint8Array): string {
  return keccak256(tarballData)
}

/**
 * Create a preview of tarball contents
 */
export function previewTarball(tarballData: Uint8Array): { files: string[], totalSize: number } {
  try {
    // Decompress
    const decompressed = pako.ungzip(tarballData)
    
    const files: string[] = []
    let offset = 0
    let totalSize = 0
    
    while (offset < decompressed.length - 1024) {
      // Read header
      const header = decompressed.slice(offset, offset + 512)
      
      // Check if we've reached the end (empty header)
      if (header[0] === 0) break
      
      // Extract filename (null-terminated string)
      let nameEnd = 0
      for (let i = 0; i < 100; i++) {
        if (header[i] === 0) {
          nameEnd = i
          break
        }
      }
      const fileName = new TextDecoder().decode(header.slice(0, nameEnd))
      
      // Extract file size (octal string)
      const sizeStr = new TextDecoder().decode(header.slice(124, 135)).trim()
      const fileSize = parseInt(sizeStr, 8) || 0
      
      if (fileName) {
        files.push(fileName)
        totalSize += fileSize
      }
      
      // Move to next file (header + file content + padding)
      offset += 512 + fileSize
      const padding = 512 - (fileSize % 512)
      if (padding < 512) {
        offset += padding
      }
    }
    
    return { files, totalSize }
  } catch (error) {
    console.error('Error previewing tarball:', error)
    return { files: [], totalSize: 0 }
  }
}

/**
 * Validate tarball structure
 */
export function validateTarball(files: FileNode[]): { valid: boolean; errors: string[] } {
  const errors: string[] = []
  
  // Check for manifest.json
  const manifest = files.find(f => f.name === 'manifest.json')
  if (!manifest) {
    errors.push('Missing required manifest.json file')
  } else if (manifest.content) {
    try {
      const manifestData = JSON.parse(manifest.content)
      if (!manifestData.runtime) {
        errors.push('manifest.json must specify runtime (javascript or python)')
      } else if (!['javascript', 'python'].includes(manifestData.runtime)) {
        errors.push('Invalid runtime. Must be "javascript" or "python"')
      } else {
        // Check for appropriate handler file
        if (manifestData.runtime === 'javascript') {
          const hasIndex = files.some(f => f.name === 'index.js')
          if (!hasIndex) {
            errors.push('JavaScript functions must have index.js file')
          }
        } else if (manifestData.runtime === 'python') {
          const hasHandler = files.some(f => f.name === 'handler.py')
          if (!hasHandler) {
            errors.push('Python functions must have handler.py file')
          }
        }
      }
    } catch (e) {
      errors.push('Invalid manifest.json format')
    }
  }
  
  return {
    valid: errors.length === 0,
    errors
  }
}

/**
 * Export function project as deployment-ready tarball
 */
export async function exportProjectAsTarball(functionName: string, files: FileNode[]): Promise<void> {
  // Validate before creating tarball
  const validation = validateTarball(files)
  if (!validation.valid) {
    throw new Error(`Invalid function structure: ${validation.errors.join(', ')}`)
  }
  
  // Create tarball
  const tarballData = await createTarball(files)
  
  // Calculate function ID
  const functionId = calculateFunctionId(tarballData)
  
  // Create download
  const blob = new Blob([tarballData], { type: 'application/gzip' })
  const url = URL.createObjectURL(blob)
  
  const link = document.createElement('a')
  link.href = url
  link.download = `${functionName}-${functionId.slice(0, 8)}.tar.gz`
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
  
  URL.revokeObjectURL(url)
  
  console.log(`Function ID: ${functionId}`)
  alert(`Function exported successfully!\nFunction ID: ${functionId}`)
}