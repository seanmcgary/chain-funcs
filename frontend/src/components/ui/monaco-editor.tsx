import { useRef } from 'react'
import Editor, { type Monaco } from '@monaco-editor/react'
import { cn } from "@/lib/utils"

interface MonacoEditorProps {
  value: string
  onChange: (value: string) => void
  language?: string
  theme?: string
  className?: string
  placeholder?: string
  readOnly?: boolean
  options?: object
  style?: React.CSSProperties
}

export function MonacoEditor({
  value,
  onChange,
  language = 'javascript',
  theme = 'vs-dark',
  className,
  placeholder = '// Enter your code here...',
  readOnly = false,
  options = {},
  style
}: MonacoEditorProps) {
  const editorRef = useRef(null)
  const monacoRef = useRef<Monaco | null>(null)

  const handleEditorDidMount = (editor: any, monaco: Monaco) => {
    editorRef.current = editor
    monacoRef.current = monaco

    // Configure editor
    editor.updateOptions({
      fontSize: 14,
      fontFamily: 'ui-monospace, SFMono-Regular, "SF Mono", Menlo, Consolas, "Liberation Mono", monospace',
      minimap: { enabled: false },
      lineNumbers: 'on',
      roundedSelection: false,
      scrollBeyondLastLine: false,
      automaticLayout: true,
      tabSize: 2,
      insertSpaces: true,
      detectIndentation: false,
      wordWrap: 'on',
      wrappingIndent: 'indent',
      ...options
    })

    // Add custom theme for better integration with shadcn/ui
    monaco.editor.defineTheme('custom-dark', {
      base: 'vs-dark',
      inherit: true,
      rules: [
        { token: 'comment', foreground: '6A9955' },
        { token: 'keyword', foreground: '569CD6' },
        { token: 'string', foreground: 'CE9178' },
        { token: 'number', foreground: 'B5CEA8' },
        { token: 'function', foreground: 'DCDCAA' }
      ],
      colors: {
        'editor.background': '#0F0F0F',
        'editor.foreground': '#CCCCCC',
        'editorLineNumber.foreground': '#464647',
        'editor.selectionBackground': '#264F78',
        'editor.lineHighlightBackground': '#2D2D30'
      }
    })

    monaco.editor.defineTheme('custom-light', {
      base: 'vs',
      inherit: true,
      rules: [
        { token: 'comment', foreground: '008000' },
        { token: 'keyword', foreground: '0000FF' },
        { token: 'string', foreground: 'A31515' },
        { token: 'number', foreground: '098658' },
        { token: 'function', foreground: '795E26' }
      ],
      colors: {
        'editor.background': '#FFFFFF',
        'editor.foreground': '#000000',
        'editorLineNumber.foreground': '#237893',
        'editor.selectionBackground': '#ADD6FF',
        'editor.lineHighlightBackground': '#F7F7F7'
      }
    })

    // Set custom theme based on system theme
    const isDark = document.documentElement.classList.contains('dark')
    monaco.editor.setTheme(isDark ? 'custom-dark' : 'custom-light')

    // Watch for theme changes
    const observer = new MutationObserver((mutations) => {
      mutations.forEach((mutation) => {
        if (mutation.attributeName === 'class') {
          const isDark = document.documentElement.classList.contains('dark')
          monaco.editor.setTheme(isDark ? 'custom-dark' : 'custom-light')
        }
      })
    })

    observer.observe(document.documentElement, {
      attributes: true,
      attributeFilter: ['class']
    })

    // Add placeholder support
    if (!value && placeholder) {
      const placeholderWidget = {
        getId: () => 'placeholder-widget',
        getDomNode: () => {
          const node = document.createElement('div')
          node.style.color = '#6B7280'
          node.style.fontStyle = 'italic'
          node.style.pointerEvents = 'none'
          node.style.position = 'absolute'
          node.style.top = '0px'
          node.style.left = '0px'
          node.style.padding = '0px 0px 0px 60px'
          node.style.lineHeight = '19px'
          node.innerHTML = placeholder
          return node
        },
        getPosition: () => null
      }

      if (!value) {
        editor.addOverlayWidget(placeholderWidget)
      }

      editor.onDidChangeModelContent(() => {
        const currentValue = editor.getValue()
        if (currentValue && !value) {
          editor.removeOverlayWidget(placeholderWidget)
        } else if (!currentValue && value) {
          editor.addOverlayWidget(placeholderWidget)
        }
      })
    }

    // Setup additional language features
    if (language === 'javascript') {
      // Add common JavaScript completions
      monaco.languages.registerCompletionItemProvider('javascript', {
        provideCompletionItems: (model, position) => {
          const word = model.getWordUntilPosition(position)
          const range = {
            startLineNumber: position.lineNumber,
            endLineNumber: position.lineNumber,
            startColumn: word.startColumn,
            endColumn: word.endColumn
          }
          
          const suggestions = [
            {
              label: 'handler',
              kind: monaco.languages.CompletionItemKind.Function,
              insertText: 'exports.handler = async function(args = []) {\n\t$0\n\treturn {};\n};',
              insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
              documentation: 'ChainFuncs handler function template',
              range
            },
            {
              label: 'console.log',
              kind: monaco.languages.CompletionItemKind.Function,
              insertText: 'console.log($0)',
              insertTextRules: monaco.languages.CompletionItemInsertTextRule.InsertAsSnippet,
              documentation: 'Log to console',
              range
            }
          ]
          return { suggestions }
        }
      })
    }
  }

  const handleChange = (value: string | undefined) => {
    onChange(value || '')
  }

  return (
    <div className={cn(
      "border border-input rounded-md overflow-hidden bg-background",
      className
    )} style={style}>
      <Editor
        height={style?.height || "400px"}
        language={language}
        value={value}
        onChange={handleChange}
        onMount={handleEditorDidMount}
        loading={
          <div className="flex items-center justify-center h-[400px] text-muted-foreground">
            Loading editor...
          </div>
        }
        options={{
          readOnly,
          ...options
        }}
      />
    </div>
  )
}