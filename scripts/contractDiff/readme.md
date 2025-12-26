# Contract code difference check

This script checks the source code differences between two deployed smart contracts on supported networks.  
It should be run after deployment/upgrade to verify that the new implementation matches expectations.

## What it does
- **Fetch & verify:** Downloads verified contract source code from Etherscan/Blockscout
- **Flatten:** Normalizes multi-file and single-file contracts into comparable format
- **Reorder:** Ensures consistent ordering of contract declarations across files
- **Save files:** Stores both old and new contract versions locally under `./contract-diffs`

## Usage
```bash
# set your etherscan api key in .env file
export ETHERSCAN_API_KEY=your_api_key_here

# Compile contracts first to generate artifacts
npx hardhat compile

# Run difference check between old and new contract addresses
npx hardhat contractDiff --network base_mainnet \
  --old-address 0x8bb9EC4ae5Ae0B87B314Ae38B6DF4B30f8013AE0 \
  --new-address 0x9ef630aB803FE9a7F46214242731298EE7025794
```

## Output
- Flattened contracts will be saved in the `./contract-diffs` directory with filenames like:
    - `2025-10-02_OLD_ContractName_0x8bb9EC4a.sol`
    - `2025-10-02_NEW_ContractName_0x9ef630aB.sol`

## Comparing code

Flattened contract files can be compared using any diff tool:
  - [Diffchecker](https://www.diffchecker.com/text-compare/)
  - [Text Compare](https://text-compare.com/)

Or locally using the vimdiff utility:
```bash
vimdiff contract-diffs/OLD_FLATTENED_CONTRACT.sol contract-diffs/NEW_FLATTENED_CONTRACT.sol
```