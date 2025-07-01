// Import contract ABIs and addresses from the CLI clients
import FaaSContract from '../../../clients/contracts/FaaS.json';
import TaskMailboxContract from '../../../clients/contracts/TaskMailbox.json';

export const CONTRACTS = {
  FaaS: {
    address: FaaSContract.address as `0x${string}`,
    abi: FaaSContract.abi,
    name: FaaSContract.name,
  },
  TaskMailbox: {
    address: TaskMailboxContract.address as `0x${string}`,
    abi: TaskMailboxContract.abi,
    name: TaskMailboxContract.name,
  },
} as const;

// Chain configurations
export const CHAIN_CONFIG = {
  localhost: {
    id: 1337,
    name: 'Local Development',
    rpcUrl: 'http://localhost:7545',
    wsRpcUrl: 'ws://localhost:7545',
  },
  sepolia: {
    id: 11155111,
    name: 'Sepolia Testnet',
    rpcUrl: 'https://sepolia.infura.io/v3/YOUR_INFURA_KEY',
    wsRpcUrl: 'wss://sepolia.infura.io/ws/v3/YOUR_INFURA_KEY',
  },
} as const;

export type SupportedChainId = keyof typeof CHAIN_CONFIG;