# Safe Upgrade Proposal Script

A Hardhat script that creates upgrade proposals for GatewayEVM contracts across all supported EVM networks using Safe multi-sig wallets.

## Overview
This script automates the process of creating upgrade proposals for GatewayEVM proxy contracts deployed across different EVMs. It uses Safe's transaction service to propose upgrades that must be approved and executed by the Safe signers.

## Required Environment Variables
Create a ```.env``` file in your project root with the following variables:
```
# Safe API Configuration
SAFE_API_KEY=your_safe_api_key_here

# Proposer Wallet Configuration
PROPOSER_PRIVATE_KEY=0x1234567890abcdef...  # Private key of the proposer wallet
PROPOSER_ADDRESS=0x1234567890123456789012345678901234567890  # Address of the proposer wallet

# Optional: Default upgrade data
UPGRADE_DATA=0x  # Default: empty bytes (0x)
```

## Usage
### Basic Usage
```
# Create upgrade proposals for testnet networks
npx hardhat upgradeProposal --testnet

# Create upgrade proposals for mainnet networks
npx hardhat upgradeProposal --mainnet
```
### With Custom Upgrade Data
```
# Pass upgrade data directly
npx hardhat upgradeProposal --testnet --upgrade-data "0x1234567890abcdef"

# Load upgrade data from file
npx hardhat upgradeProposal --testnet --upgrade-data-file upgrade_data.txt

# Using environment variable
UPGRADE_DATA="0x1234567890abcdef" npx hardhat upgradeProposal --testnet
```