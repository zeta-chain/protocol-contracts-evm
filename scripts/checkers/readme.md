# Protocol Checksum Verification

This script checks core ZetaChain protocol contracts deployed on ZetaChain and all connected EVM chains. It should be run after deployment/upgrade to make sure contracts are updated properly.

## What it does
- **Proxy check:** Makes sure proxy contracts point to the right implementation addresses
- **Bytecode check:** Compares deployed contract code with local compiled code
- **State check:** Verifies contracts state is correct
- **Role check:** Makes sure the right addresses have the right permissions

## Usage
```
# Generate contract artifacts first
npx hardhat compile

# Run checksum for testnet (default)
npx hardhat protocolChecksum
npx hardhat protocolChecksum --testnet

# Run checksum for mainnet
npx hardhat protocolChecksum --mainnet