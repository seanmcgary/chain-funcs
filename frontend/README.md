# FaaS AVS Frontend

A React-based web interface for the FaaS (Function-as-a-Service) Autonomous Verifiable Service built on EigenLayer.

## 🚀 Quick Start

### Prerequisites
- Node.js 18+ and npm
- MetaMask or other Web3 wallet

### Development

```bash
# Install dependencies
npm install

# Start development server
npm run dev
```

The app will be available at `http://localhost:5173`

### Production Build

```bash
# Build for production
npm run build

# Preview production build locally
npm run preview
```

### Other Commands

```bash
# Run linter
npm run lint

# TypeScript type check
npm run build  # includes tsc -b
```

## 🏗️ Project Structure

```
src/
├── components/          # React components
│   ├── ui/             # shadcn/ui base components
│   ├── WalletStatus.tsx
│   ├── ContractInfo.tsx
│   ├── Navigation.tsx
│   ├── FunctionEditor/
│   ├── FunctionBrowser/
│   ├── DeploymentHub/
│   └── FunctionCaller/
├── hooks/              # Custom React hooks
│   └── useContracts.ts
├── lib/                # Utilities and config
│   ├── utils.ts
│   ├── wagmi.ts        # Web3 configuration
│   └── contracts.ts    # Contract ABIs and addresses
└── App.tsx
```

## 🔧 Features

- **Wallet Connection**: MetaMask/WalletConnect integration via RainbowKit
- **Function Editor**: Code editor for JavaScript/Python functions (Milestone 2)
- **Function Browser**: Discover deployed functions
- **Deployment Hub**: Deploy functions to the FaaS AVS network  
- **Function Caller**: Execute functions and view real-time results
- **Contract Integration**: Direct interaction with FaaS and TaskMailbox contracts

## 🌐 Web3 Integration

- **Wagmi v2** - React hooks for Ethereum
- **Viem** - TypeScript Ethereum library
- **RainbowKit** - Beautiful wallet connection UX
- **Contract ABIs** - Imported from `../clients/contracts/`

## 🎨 UI/UX

- **Tailwind CSS** - Utility-first styling
- **shadcn/ui** - Beautiful, accessible components  
- **Responsive Design** - Works on desktop and mobile
- **Dark/Light Theme** - System preference support

## 📦 Tech Stack

- **Vite** - Fast build tool and dev server
- **React 19** - UI framework with latest features
- **TypeScript** - Type safety
- **ESLint** - Code linting
- **PostCSS** - CSS processing

## 🔗 Contract Configuration

The frontend automatically imports contract addresses and ABIs from:
- `../clients/contracts/FaaS.json`
- `../clients/contracts/TaskMailbox.json`

These are embedded at build time from the CLI's devnet deployment configuration.

## 📝 Environment Variables

Create a `.env.local` file for local overrides:

```bash
# Optional: WalletConnect Project ID for better wallet support
VITE_WALLETCONNECT_PROJECT_ID=your_project_id_here

# Optional: Custom RPC URLs
VITE_LOCALHOST_RPC_URL=http://localhost:7545
VITE_SEPOLIA_RPC_URL=https://sepolia.infura.io/v3/YOUR_KEY
```

## 🚦 Development Status

**Milestone 1 Complete ✅**
- [x] Project setup with Vite + React + TypeScript
- [x] Tailwind CSS and shadcn/ui integration
- [x] Web3 setup with Wagmi/Viem/RainbowKit
- [x] Basic navigation and app shell
- [x] Wallet connection with status display
- [x] Contract configuration from CLI
- [x] Build and development workflows

**Next: Milestone 2** - Function Editor with Monaco Editor

## 🤝 Contributing

This frontend is part of the larger FaaS AVS project. See the main project README for contribution guidelines.