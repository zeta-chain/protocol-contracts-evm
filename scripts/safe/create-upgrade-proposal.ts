import hre from 'hardhat';
import * as dotenv from 'dotenv';
import * as fs from 'fs';
import * as path from 'path';
import { ethers } from 'ethers';
import SafeApiKit from '@safe-global/api-kit';
import Safe from '@safe-global/protocol-kit';
import {
  MetaTransactionData,
  OperationType
} from '@safe-global/types-kit';

dotenv.config();

const SUPPORTED_CHAINS_API = {
  mainnet: "https://zetachain.blockpi.network/lcd/v1/public/zeta-chain/observer/supportedChains",
  testnet: "https://zetachain-athens.blockpi.network/lcd/v1/public/zeta-chain/observer/supportedChains"
};

interface ChainInfo {
  chain_id: string;
  name: string;
}

interface SupportedChainsResponse {
  chains: ChainInfo[];
}

async function fetchSupportedChains(): Promise<Record<string, string>> {
  try {
    const networkType = process.env.NETWORK_TYPE || 'testnet';
    const apiUrl = SUPPORTED_CHAINS_API[networkType as keyof typeof SUPPORTED_CHAINS_API];
    const response = await fetch(apiUrl);

    if (!response.ok) {
      throw new Error(`HTTP error, status: ${response.status}`);
    }

    const data: SupportedChainsResponse = await response.json();
    const chainMapping: Record<string, string> = {};

    for (const chain of data.chains) {
      chainMapping[chain.chain_id] = chain.name;
    }

    return chainMapping;
  } catch (error) {
    console.error("❌ Failed to fetch supported chains from API:", error);
    process.exit(1);
  }
}

function loadAddresses() {
  try {
    const networkType = process.env.NETWORK_TYPE || 'testnet';
    const fileName = `${networkType}.json`;

    const addressesPath = path.join(process.cwd(), `/data/checksum/${fileName}`);
    const rawData = fs.readFileSync(addressesPath, 'utf8');
    return JSON.parse(rawData);
  } catch (error) {
    console.error('❌ Error loading addresses from data/checkers/testnet.json:', error);
    process.exit(1);
  }
}

function createUpgradeCalldata(newImplementationAddress: string, upgradeData: string): string {
  const iface = new ethers.Interface([
    'function upgradeToAndCall(address newImplementation, bytes memory data) external payable',
  ]);
  return iface.encodeFunctionData('upgradeToAndCall', [
    newImplementationAddress,
    upgradeData,
  ]);
}

async function createUpgradeProposal(
  chainId: number,
  networkName: string,
  addresses: Record<string, string>,
  upgradeData: string
) {
  const safeAddress = addresses.Admin;
  const gatewayProxy = addresses.GatewayEVM;
  const newImplementationAddress = addresses.GatewayEVMImplementation;
  const networkConfig = hre.config.networks[networkName] as any;

  console.log(`\n🔄 Creating proposal for GatewayEVM upgrade on ${networkName} (Chain ID: ${chainId})`);
  console.log(`  Safe: ${safeAddress}`);
  console.log(`  Gateway: ${gatewayProxy}`);
  console.log(`  New Implementation: ${newImplementationAddress}`);

  const apiKit = new SafeApiKit({
    chainId: BigInt(chainId),
    apiKey: process.env.SAFE_API_KEY,
  });

  const protocolKitOwner1 = await Safe.init({
    provider: networkConfig.url,
    signer: process.env.PROPOSER_PRIVATE_KEY,
    safeAddress: safeAddress
  });

  const upgradeCalldata = createUpgradeCalldata(newImplementationAddress, upgradeData);

  const safeTransactionData: MetaTransactionData = {
    to: gatewayProxy,
    value: '0',
    data: upgradeCalldata,
    operation: OperationType.Call
  }

  const safeTransaction = await protocolKitOwner1.createTransaction({
    transactions: [safeTransactionData]
  })

  const safeTxHash = await protocolKitOwner1.getTransactionHash(safeTransaction)
  const signature = await protocolKitOwner1.signHash(safeTxHash)

  await apiKit.proposeTransaction({
    safeAddress: safeAddress,
    safeTransactionData: safeTransaction.data,
    safeTxHash,
    senderAddress: process.env.PROPOSER_ADDRESS!,
    senderSignature: signature.data
  })

  console.log('  ✅ Proposal submitted to Safe service');
  console.log(`  Safe Transaction Hash: ${safeTxHash}`);
}

export async function main() {
  const addresses = loadAddresses();
  const CHAIN_ID_TO_NETWORK = await fetchSupportedChains();
  const upgradeData = process.env.UPGRADE_DATA || '0x';

  console.log('🚀 Starting upgrade proposals for all networks');
  console.log('='.repeat(60));

  for (const [chainId, contracts] of Object.entries(addresses)) {
    const chainIdNum = parseInt(chainId);
    const networkName = CHAIN_ID_TO_NETWORK[chainId];
    const contractsTyped = contracts as Record<string, string>;
    if (!contractsTyped.GatewayEVM) {
      console.log(`\n⚠️ No GatewayEVM found for ${networkName} (Chain ID: ${chainId})`);
      continue;
    }

    try {
      const newImplementation = contractsTyped.GatewayEVMImplementation;
      if (!newImplementation) {
        console.log(`\n⚠️ No GatewayEVMImplementation found for ${networkName} (Chain ID: ${chainId})`);
        continue;
      }

      await createUpgradeProposal(chainIdNum, networkName, contractsTyped, upgradeData);
    } catch (error) {
      console.error(`\n❌ Error creating proposal for ${networkName} (Chain ID: ${chainId}):`, error);
    }
  }
}
