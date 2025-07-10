import { getDefaultConfig } from '@rainbow-me/rainbowkit';
import { localhost, sepolia, mainnet } from 'wagmi/chains';

export const config = getDefaultConfig({
  appName: 'ChainFuncs',
  projectId: import.meta.env.VITE_WALLETCONNECT_PROJECT_ID || 'YOUR_PROJECT_ID',
  chains: [localhost, sepolia, mainnet],
  ssr: false,
});