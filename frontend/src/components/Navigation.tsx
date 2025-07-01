import { ConnectButton } from '@rainbow-me/rainbowkit'

interface NavigationProps {
  activeTab: string
  setActiveTab: (tab: string) => void
}

const Navigation = ({ activeTab, setActiveTab }: NavigationProps) => {

  const tabs = [
    { id: 'editor', label: 'Function Editor', icon: '📝' },
    { id: 'browser', label: 'Function Browser', icon: '🔍' },
    { id: 'deploy', label: 'Deploy', icon: '🚀' },
    { id: 'execute', label: 'Execute', icon: '⚡' },
  ]

  return (
    <nav className="bg-background border-b border-border">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex items-center">
            {/* Logo */}
            <div className="flex-shrink-0 flex items-center">
              <h1 className="text-xl font-bold text-foreground">
                ChainFuncs
              </h1>
            </div>
            
            {/* Navigation tabs */}
            <div className="hidden md:ml-6 md:flex md:space-x-8">
              {tabs.map((tab) => (
                <button
                  key={tab.id}
                  onClick={() => setActiveTab(tab.id)}
                  className={`inline-flex items-center px-1 pt-1 text-sm font-medium border-b-2 transition-colors ${
                    activeTab === tab.id
                      ? 'border-primary text-primary'
                      : 'border-transparent text-muted-foreground hover:text-foreground hover:border-border'
                  }`}
                >
                  <span className="mr-2">{tab.icon}</span>
                  {tab.label}
                </button>
              ))}
            </div>
          </div>

          {/* Right side - Wallet connection */}
          <div className="flex items-center">
            <ConnectButton />
          </div>
        </div>
      </div>

      {/* Mobile menu */}
      <div className="md:hidden">
        <div className="px-2 pt-2 pb-3 space-y-1 bg-background border-t border-border">
          {tabs.map((tab) => (
            <button
              key={tab.id}
              onClick={() => setActiveTab(tab.id)}
              className={`block px-3 py-2 rounded-md text-base font-medium w-full text-left transition-colors ${
                activeTab === tab.id
                  ? 'bg-primary/10 text-primary'
                  : 'text-muted-foreground hover:text-foreground hover:bg-accent'
              }`}
            >
              <span className="mr-2">{tab.icon}</span>
              {tab.label}
            </button>
          ))}
        </div>
      </div>
    </nav>
  )
}

export default Navigation