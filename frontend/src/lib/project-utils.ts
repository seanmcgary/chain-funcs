import { type FileNode } from "@/components/ui/file-tree"

export interface FunctionProject {
  name: string
  files: FileNode[]
  version: string
  created: string
  lastModified: string
}

/**
 * Export function project as JSON file
 */
export function exportProject(functionName: string, files: FileNode[]): void {
  const project: FunctionProject = {
    name: functionName,
    files: files,
    version: "1.0.0",
    created: new Date().toISOString(),
    lastModified: new Date().toISOString()
  }

  const dataStr = JSON.stringify(project, null, 2)
  const dataBlob = new Blob([dataStr], { type: 'application/json' })
  const url = URL.createObjectURL(dataBlob)
  
  const link = document.createElement('a')
  link.href = url
  link.download = `${functionName}-project.json`
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
  
  URL.revokeObjectURL(url)
}

/**
 * Export function project as tarball-ready ZIP
 */
export async function exportProjectAsZip(functionName: string, files: FileNode[]): Promise<void> {
  // Import JSZip dynamically
  const JSZip = (await import('jszip')).default
  
  const zip = new JSZip()
  
  // Add all files to the zip
  files.forEach(file => {
    if (file.type === 'file' && file.content) {
      zip.file(file.name, file.content)
    }
  })
  
  const zipBlob = await zip.generateAsync({ type: 'blob' })
  const url = URL.createObjectURL(zipBlob)
  
  const link = document.createElement('a')
  link.href = url
  link.download = `${functionName}.zip`
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
  
  URL.revokeObjectURL(url)
}

/**
 * Import function project from JSON file
 */
export function importProject(): Promise<FunctionProject> {
  return new Promise((resolve, reject) => {
    const input = document.createElement('input')
    input.type = 'file'
    input.accept = '.json'
    
    input.onchange = (event) => {
      const file = (event.target as HTMLInputElement).files?.[0]
      if (!file) {
        reject(new Error('No file selected'))
        return
      }
      
      const reader = new FileReader()
      reader.onload = (e) => {
        try {
          const projectData = JSON.parse(e.target?.result as string)
          
          // Validate project structure
          if (!projectData.name || !projectData.files || !Array.isArray(projectData.files)) {
            reject(new Error('Invalid project file format'))
            return
          }
          
          resolve(projectData as FunctionProject)
        } catch (error) {
          reject(new Error('Failed to parse project file'))
        }
      }
      
      reader.onerror = () => reject(new Error('Failed to read file'))
      reader.readAsText(file)
    }
    
    input.click()
  })
}

/**
 * Import files from ZIP archive
 */
export async function importProjectFromZip(): Promise<{ name: string; files: FileNode[] }> {
  return new Promise((resolve, reject) => {
    const input = document.createElement('input')
    input.type = 'file'
    input.accept = '.zip'
    
    input.onchange = async (event) => {
      const file = (event.target as HTMLInputElement).files?.[0]
      if (!file) {
        reject(new Error('No file selected'))
        return
      }
      
      try {
        // Import JSZip dynamically
        const JSZip = (await import('jszip')).default
        
        const zip = new JSZip()
        const zipData = await zip.loadAsync(file)
        
        const files: FileNode[] = []
        const projectName = file.name.replace('.zip', '')
        
        // Extract all files from ZIP
        for (const [fileName, zipFile] of Object.entries(zipData.files)) {
          if (!zipFile.dir) {
            const content = await zipFile.async('string')
            files.push({
              id: fileName,
              name: fileName,
              type: 'file',
              content: content
            })
          }
        }
        
        resolve({ name: projectName, files })
      } catch (error) {
        reject(new Error('Failed to read ZIP file'))
      }
    }
    
    input.click()
  })
}

/**
 * Save project to localStorage
 */
export function saveProjectToLocalStorage(functionName: string, files: FileNode[]): void {
  const project: FunctionProject = {
    name: functionName,
    files: files,
    version: "1.0.0",
    created: new Date().toISOString(),
    lastModified: new Date().toISOString()
  }
  
  localStorage.setItem(`chainfuncs-project-${functionName}`, JSON.stringify(project))
}

/**
 * Load project from localStorage
 */
export function loadProjectFromLocalStorage(functionName: string): FunctionProject | null {
  const projectData = localStorage.getItem(`chainfuncs-project-${functionName}`)
  if (!projectData) return null
  
  try {
    return JSON.parse(projectData) as FunctionProject
  } catch {
    return null
  }
}

/**
 * List all saved projects in localStorage
 */
export function listSavedProjects(): string[] {
  const projects: string[] = []
  
  for (let i = 0; i < localStorage.length; i++) {
    const key = localStorage.key(i)
    if (key?.startsWith('chainfuncs-project-')) {
      const projectName = key.replace('chainfuncs-project-', '')
      projects.push(projectName)
    }
  }
  
  return projects.sort()
}

/**
 * Delete project from localStorage
 */
export function deleteProjectFromLocalStorage(functionName: string): void {
  localStorage.removeItem(`chainfuncs-project-${functionName}`)
}