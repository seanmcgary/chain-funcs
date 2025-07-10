import { useState } from 'react'
import Navigation from './components/Navigation'
import FunctionEditor from './components/FunctionEditor/FunctionEditor'
import FunctionBrowser from './components/FunctionBrowser/FunctionBrowser'
import DeploymentHub from './components/DeploymentHub/DeploymentHub'
import FunctionCaller from './components/FunctionCaller/FunctionCaller'

function App() {
  const [activeTab, setActiveTab] = useState('editor')

  const renderActiveComponent = () => {
    switch (activeTab) {
      case 'editor':
        return <FunctionEditor />
      case 'browser':
        return <FunctionBrowser />
      case 'deploy':
        return <DeploymentHub />
      case 'execute':
        return <FunctionCaller />
      default:
        return <FunctionEditor />
    }
  }

  return (
    <div className="h-screen bg-background flex flex-col">
      <Navigation activeTab={activeTab} setActiveTab={setActiveTab} />
      <main className="flex-1 overflow-hidden">
        {renderActiveComponent()}
      </main>
    </div>
  )
}

export default App