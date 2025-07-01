import { useAccount, useBalance, useChainId } from 'wagmi'

const WalletStatus = () => {
  const { address, isConnected } = useAccount()
  const chainId = useChainId()
  const { data: balance } = useBalance({ address })

  if (!isConnected) {
    return (
      <div className="bg-card border border-border rounded-lg p-4">
        <h3 className="text-sm font-medium text-muted-foreground mb-2">
          Wallet Status
        </h3>
        <p className="text-sm text-muted-foreground">
          No wallet connected. Please connect your wallet to use ChainFuncs.
        </p>
      </div>
    )
  }

  return (
    <div className="bg-card border border-border rounded-lg p-4">
      <h3 className="text-sm font-medium text-muted-foreground mb-2">
        Wallet Status
      </h3>
      <div className="space-y-2">
        <div>
          <span className="text-xs text-muted-foreground">Address:</span>
          <p className="text-sm font-mono">
            {address?.slice(0, 6)}...{address?.slice(-4)}
          </p>
        </div>
        <div>
          <span className="text-xs text-muted-foreground">Chain ID:</span>
          <p className="text-sm">{chainId}</p>
        </div>
        {balance && (
          <div>
            <span className="text-xs text-muted-foreground">Balance:</span>
            <p className="text-sm">
              {parseFloat(balance.formatted).toFixed(4)} {balance.symbol}
            </p>
          </div>
        )}
      </div>
    </div>
  )
}

export default WalletStatus