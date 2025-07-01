import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Eye, Download, Play } from "lucide-react"

const FunctionBrowser = () => {
  // Mock data for demonstration
  const functions = [
    { id: "0x1a2b3c", name: "hello-world", runtime: "javascript", size: "2.1 KB", deployed: "2024-07-01" },
    { id: "0x4d5e6f", name: "data-processor", runtime: "python", size: "15.3 KB", deployed: "2024-06-30" },
    { id: "0x7g8h9i", name: "api-proxy", runtime: "javascript", size: "8.7 KB", deployed: "2024-06-29" },
  ]

  return (
    <div className="p-6 space-y-6">
      <Card>
        <CardHeader>
          <CardTitle>Function Browser</CardTitle>
          <CardDescription>
            Browse and discover deployed functions on the ChainFuncs network
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="flex gap-4">
            <Input 
              placeholder="Search functions..." 
              className="max-w-sm"
            />
            <Button variant="outline">
              Filter
            </Button>
          </div>
          
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Function ID</TableHead>
                <TableHead>Name</TableHead>
                <TableHead>Runtime</TableHead>
                <TableHead>Size</TableHead>
                <TableHead>Deployed</TableHead>
                <TableHead>Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {functions.map((func) => (
                <TableRow key={func.id}>
                  <TableCell className="font-mono text-sm">{func.id}</TableCell>
                  <TableCell className="font-medium">{func.name}</TableCell>
                  <TableCell>{func.runtime}</TableCell>
                  <TableCell>{func.size}</TableCell>
                  <TableCell>{func.deployed}</TableCell>
                  <TableCell>
                    <div className="flex gap-2">
                      <Button size="sm" variant="outline" className="h-8 w-8 p-0">
                        <Eye className="h-4 w-4" />
                      </Button>
                      <Button size="sm" variant="outline" className="h-8 w-8 p-0">
                        <Download className="h-4 w-4" />
                      </Button>
                      <Button size="sm" variant="outline" className="h-8 w-8 p-0">
                        <Play className="h-4 w-4" />
                      </Button>
                    </div>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </CardContent>
      </Card>
    </div>
  )
}

export default FunctionBrowser