import { CONTRACTS } from '@/lib/contracts'

const ContractInfo = () => {
  return (
    <div className="bg-card border border-border rounded-lg p-4 mt-4">
      <h3 className="text-sm font-medium text-muted-foreground mb-2">
        Contract Configuration
      </h3>
      <div className="space-y-2">
        <div>
          <span className="text-xs text-muted-foreground">FaaS Contract:</span>
          <p className="text-sm font-mono break-all">
            {CONTRACTS.FaaS.address}
          </p>
        </div>
        <div>
          <span className="text-xs text-muted-foreground">TaskMailbox Contract:</span>
          <p className="text-sm font-mono break-all">
            {CONTRACTS.TaskMailbox.address}
          </p>
        </div>
        <div>
          <span className="text-xs text-muted-foreground">Status:</span>
          <p className="text-sm text-green-600">✓ Configuration Loaded</p>
        </div>
      </div>
    </div>
  )
}

export default ContractInfo