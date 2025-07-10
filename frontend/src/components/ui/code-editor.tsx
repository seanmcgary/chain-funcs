import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter'
import { oneDark, oneLight } from 'react-syntax-highlighter/dist/esm/styles/prism'
import { cn } from "@/lib/utils"

interface CodeEditorProps {
  value: string
  onChange: (value: string) => void
  language?: string
  className?: string
  placeholder?: string
}

export function CodeEditor({
  value,
  onChange,
  language = 'javascript',
  className,
  placeholder = '// Enter your code here...'
}: CodeEditorProps) {

  return (
    <div className={cn(
      "relative rounded-md border border-input bg-background text-sm ring-offset-background",
      "focus-within:ring-2 focus-within:ring-ring focus-within:ring-offset-2",
      className
    )}>
      <div className="relative">
        {/* Hidden textarea for actual input */}
        <textarea
          value={value}
          onChange={(e) => onChange(e.target.value)}
          placeholder={placeholder}
          className={cn(
            "absolute inset-0 w-full h-full resize-none bg-transparent text-transparent caret-white p-4",
            "outline-none font-mono text-sm leading-6",
            "placeholder:text-muted-foreground"
          )}
          spellCheck={false}
          autoComplete="off"
          autoCorrect="off"
          autoCapitalize="off"
          style={{
            // Ensure textarea and syntax highlighter have the same font metrics
            fontFamily: 'ui-monospace, SFMono-Regular, "SF Mono", Menlo, Consolas, "Liberation Mono", monospace',
            fontSize: '14px',
            lineHeight: '1.5',
          }}
        />
        
        {/* Syntax highlighted code display */}
        <div className="pointer-events-none">
          <SyntaxHighlighter
            language={language}
            style={oneDark}
            customStyle={{
              margin: 0,
              padding: '16px',
              background: 'transparent',
              fontSize: '14px',
              lineHeight: '1.5',
              fontFamily: 'ui-monospace, SFMono-Regular, "SF Mono", Menlo, Consolas, "Liberation Mono", monospace',
            }}
            codeTagProps={{
              style: {
                fontFamily: 'ui-monospace, SFMono-Regular, "SF Mono", Menlo, Consolas, "Liberation Mono", monospace',
              }
            }}
          >
            {value || placeholder}
          </SyntaxHighlighter>
        </div>
      </div>
    </div>
  )
}