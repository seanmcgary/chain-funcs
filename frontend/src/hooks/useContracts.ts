import { useReadContract, useWriteContract, useWaitForTransactionReceipt } from 'wagmi'
import { CONTRACTS } from '@/lib/contracts'

// Hook for reading from FaaS contract
export const useFaaSContract = () => {
  const { writeContract, data: hash, isPending } = useWriteContract()
  
  const { isLoading: isConfirming, isSuccess: isConfirmed } = 
    useWaitForTransactionReceipt({ hash })

  const registerFunction = async (content: Uint8Array) => {
    return writeContract({
      address: CONTRACTS.FaaS.address,
      abi: CONTRACTS.FaaS.abi,
      functionName: 'registerFunction',
      args: [content],
    })
  }

  const deployFunction = async (url: string, digest: `0x${string}`) => {
    return writeContract({
      address: CONTRACTS.FaaS.address,
      abi: CONTRACTS.FaaS.abi,
      functionName: 'deployFunction',
      args: [url, digest],
    })
  }

  const callFunction = async (functionId: `0x${string}`, args: Uint8Array) => {
    return writeContract({
      address: CONTRACTS.FaaS.address,
      abi: CONTRACTS.FaaS.abi,
      functionName: 'callFunction',
      args: [functionId, args],
    })
  }

  return {
    registerFunction,
    deployFunction,
    callFunction,
    hash,
    isPending,
    isConfirming,
    isConfirmed,
  }
}

// Hook for reading function metadata
export const useFunctionMetadata = (functionId?: `0x${string}`) => {
  return useReadContract({
    address: CONTRACTS.FaaS.address,
    abi: CONTRACTS.FaaS.abi,
    functionName: 'getFunctionMetadata',
    args: functionId ? [functionId] : undefined,
    query: { enabled: !!functionId },
  })
}

// Hook for reading function content
export const useFunctionContent = (functionId?: `0x${string}`) => {
  return useReadContract({
    address: CONTRACTS.FaaS.address,
    abi: CONTRACTS.FaaS.abi,
    functionName: 'getFunctionContent',
    args: functionId ? [functionId] : undefined,
    query: { enabled: !!functionId },
  })
}