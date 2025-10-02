import hre from "hardhat";
import path from "path";
import fs from "fs";
import { JsonRpcProvider } from "ethers";

// ERC1967 implementation slot: keccak256("eip1967.proxy.implementation") - 1
const ERC1967_IMPLEMENTATION_SLOT = "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc";
const DEFAULT_ADMIN_ROLE = "0x0000000000000000000000000000000000000000000000000000000000000000";

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
    console.error(`❌ Error loading addresses from JSON: ${error}`);
    process.exit(1);
  }
}

async function checkProxy(
  proxyAddress: string,
  localImplementation: string,
  provider: JsonRpcProvider
) {
  try {
    // Read the implementation slot form the proxy
    const implementationSlot = await provider.getStorage(proxyAddress, ERC1967_IMPLEMENTATION_SLOT);
    // Convert to address
    const remoteImplementation = implementationSlot !== '0x0000000000000000000000000000000000000000000000000000000000000000'
      ? hre.ethers.getAddress('0x' + implementationSlot.slice(-40))
      : '0x0000000000000000000000000000000000000000';

    const matches = remoteImplementation.toLowerCase() === localImplementation.toLowerCase();
    console.log(`    Local Implementation:  ${localImplementation}`);
    console.log(`    Remote Implementation: ${remoteImplementation}`);
    console.log(`    Synced: ${matches ? '✅' : '❌'}`);
    if (!matches) {
      console.log(`      Local: ${localImplementation}`);
      console.log(`      Remote:   ${remoteImplementation}`);
    }

    return matches;
  } catch (error) {
    console.error(`    Proxy Implementation: ❌ (Error)`);
    return false;
  }
}

async function compareBytecode(
  contractName: string,
  contractAddress: string,
  provider: JsonRpcProvider
) {
  try {
    const remoteBytecode = await provider.getCode(contractAddress);
    const artifact = await hre.artifacts.readArtifact(contractName);
    const localBytecode = artifact.deployedBytecode;
    const matches = remoteBytecode === localBytecode;

    console.log(`    Bytecode Match: ${matches ? '✅' : '❌'}`);
    if (!matches) {
      const lengthMatch = remoteBytecode.length === localBytecode.length;
      console.log(`      Length: ${lengthMatch ? 'Same' : 'Different'} (Remote: ${remoteBytecode.length}, Local: ${localBytecode.length})`);
    }

    return matches;
  } catch (error) {
    console.error(`    Bytecode Match: ❌ (Error)`);
    return false;
  }
}

async function checkGatewayEVMState(
  contractName: string,
  contractAddress: string,
  allContracts: Record<string, string>,
  provider: any
) {
  const gatewayEVM = await hre.ethers.getContractAt(contractName, contractAddress, provider);

  console.log(`    📊 State Variables:`);

  // GatewayEVM specific state
  const custody = await gatewayEVM.custody();
  const custodyMatches = allContracts.ERC20Custody
    ? custody.toLowerCase() === allContracts.ERC20Custody.toLowerCase()
    : false;
  console.log(`      custody: ${custodyMatches ? '✅' : '❌'} ${custody}`);
  if (!custodyMatches && allContracts.ERC20Custody) {
    console.log(`        Expected: ${allContracts.ERC20Custody}`);
  }

  const tssAddress = await gatewayEVM.tssAddress();
  const tssMatches = allContracts.TSS
    ? tssAddress.toLowerCase() === allContracts.TSS.toLowerCase()
    : false;
  console.log(`      tssAddress: ${tssMatches ? '✅' : '❌'} ${tssAddress}`);
  if (!tssMatches && allContracts.TSS) {
    console.log(`        Expected: ${allContracts.TSS}`);
  }

  const zetaConnector = await gatewayEVM.zetaConnector();
  const zetaConnectorMatches = allContracts.Connector
    ? zetaConnector.toLowerCase() === allContracts.Connector.toLowerCase()
    : false;
  console.log(`      zetaConnector: ${zetaConnectorMatches ? '✅' : '❌'} ${zetaConnector}`);
  if (!zetaConnectorMatches && allContracts.Connector) {
    console.log(`        Expected: ${allContracts.Connector}`);
  }

  const zetaToken = await gatewayEVM.zetaToken();
  const zetaTokenMatches = allContracts.ZetaToken
    ? zetaToken.toLowerCase() === allContracts.ZetaToken.toLowerCase()
    : false;

  console.log(`      zetaToken: ${zetaTokenMatches ? '✅' : '❌'} ${zetaToken}`);
  if (!zetaTokenMatches && allContracts.ZetaToken) {
    console.log(`        Expected: ${allContracts.ZetaToken}`);
  }

  // Access Control roles
  const TSS_ROLE = await gatewayEVM.TSS_ROLE();
  const ASSET_HANDLER_ROLE = await gatewayEVM.ASSET_HANDLER_ROLE();
  const PAUSER_ROLE = await gatewayEVM.PAUSER_ROLE();

  console.log(`      📋 Role Assignments:`);
  const adminHasRole = await gatewayEVM.hasRole(DEFAULT_ADMIN_ROLE, allContracts.Admin);
  console.log(`        Admin has DEFAULT_ADMIN_ROLE: ${adminHasRole ? '✅' : '❌'}`);
  const adminHasPauserRole = await gatewayEVM.hasRole(PAUSER_ROLE, allContracts.Admin);
  console.log(`        Admin has PAUSER_ROLE: ${adminHasPauserRole ? '✅' : '❌'}`);
  const tssHasRole = await gatewayEVM.hasRole(TSS_ROLE, tssAddress);
  console.log(`        TSS has TSS_ROLE: ${tssHasRole ? '✅' : '❌'}`);
  const custodyHasRole = await gatewayEVM.hasRole(ASSET_HANDLER_ROLE, custody);
  console.log(`        Custody has ASSET_HANDLER_ROLE: ${custodyHasRole ? '✅' : '❌'}`);
  const connectorHasRole = await gatewayEVM.hasRole(ASSET_HANDLER_ROLE, zetaConnector);
  console.log(`        Connector has ASSET_HANDLER_ROLE: ${connectorHasRole ? '✅' : '❌'}`);
}

async function checkERC20CustodyState(
  contractName: string,
  contractAddress: string,
  allContracts: Record<string, string>,
  provider: any
) {
  const erc20Custody = await hre.ethers.getContractAt(contractName, contractAddress, provider);

  console.log(`    📊 State Variables:`);

  // ERC20Custody specific state
  const gateway = await erc20Custody.gateway();
  const gatewayMatches = allContracts.GatewayEVM
    ? gateway.toLowerCase() === allContracts.GatewayEVM.toLowerCase()
    : false;
  console.log(`      gateway: ${gatewayMatches ? '✅' : '❌'} ${gateway}`);
  if (!gatewayMatches && allContracts.GatewayEVM) {
    console.log(`        Expected: ${allContracts.GatewayEVM}`);
  }

  const tssAddress = await erc20Custody.tssAddress();
  const tssMatches = allContracts.TSS
    ? tssAddress.toLowerCase() === allContracts.TSS.toLowerCase()
    : false;
  console.log(`      tssAddress: ${tssMatches ? '✅' : '❌'} ${tssAddress}`);
  if (!tssMatches && allContracts.TSS) {
    console.log(`        Expected: ${allContracts.TSS}`);
  }

  // Access Control roles
  const WITHDRAWER_ROLE = await erc20Custody.WITHDRAWER_ROLE();
  const WHITELISTER_ROLE = await erc20Custody.WHITELISTER_ROLE();
  const PAUSER_ROLE = await erc20Custody.PAUSER_ROLE();

  console.log(`      📋 Role Assignments:`);
  const adminHasDefaultRole = await erc20Custody.hasRole(DEFAULT_ADMIN_ROLE, allContracts.Admin);
  console.log(`        Admin has DEFAULT_ADMIN_ROLE: ${adminHasDefaultRole ? '✅' : '❌'}`);
  const adminHasPauserRole = await erc20Custody.hasRole(PAUSER_ROLE, allContracts.Admin);
  console.log(`        Admin has PAUSER_ROLE: ${adminHasPauserRole ? '✅' : '❌'}`);
  const adminHasWhitelisterRole = await erc20Custody.hasRole(WHITELISTER_ROLE, allContracts.Admin);
  console.log(`        Admin has WHITELISTER_ROLE: ${adminHasWhitelisterRole ? '✅' : '❌'}`);
  const tssHasWithdrawerRole = await erc20Custody.hasRole(WITHDRAWER_ROLE, tssAddress);
  console.log(`        TSS has WITHDRAWER_ROLE: ${tssHasWithdrawerRole ? '✅' : '❌'}`);
  const tssHasWhitelisterRole = await erc20Custody.hasRole(WHITELISTER_ROLE, tssAddress);
  console.log(`        TSS has WHITELISTER_ROLE: ${tssHasWhitelisterRole ? '✅' : '❌'}`);
}

async function checkGatewayZEVMState(
  contractName: string,
  contractAddress: string,
  allContracts: Record<string, string>,
  provider: any
) {
  const gatewayZEVM = await hre.ethers.getContractAt(contractName, contractAddress, provider);

  console.log(`    📊 State Variables:`);

  // GatewayZEVM specific state
  const zetaToken = await gatewayZEVM.zetaToken();
  const zetaTokenMatches = allContracts.ZetaToken
    ? zetaToken.toLowerCase() === allContracts.ZetaToken.toLowerCase()
    : false;
  console.log(`      zetaToken: ${zetaTokenMatches ? '✅' : '❌'} ${zetaToken}`);
  if (!zetaTokenMatches && allContracts.ZetaToken) {
    console.log(`        Expected: ${allContracts.ZetaToken}`);
  }

  // Access Control roles
  const PAUSER_ROLE = await gatewayZEVM.PAUSER_ROLE();

  console.log(`      📋 Role Assignments:`);
  const adminHasDefaultRole = await gatewayZEVM.hasRole(DEFAULT_ADMIN_ROLE, allContracts.Admin);
  console.log(`        Admin has DEFAULT_ADMIN_ROLE: ${adminHasDefaultRole ? '✅' : '❌'}`);
  const adminHasPauserRole = await gatewayZEVM.hasRole(PAUSER_ROLE, allContracts.Admin);
  console.log(`        Admin has PAUSER_ROLE: ${adminHasPauserRole ? '✅' : '❌'}`);
}

async function checkGatewayEVM(
  contractName: string,
  contractAddress: string,
  allContracts: Record<string, string>,
  provider: JsonRpcProvider
) {
  console.log(`\n  🏗️  GatewayEVM (Proxy):`);

  if (allContracts.GatewayEVMImplementation) {
    await checkProxy(contractAddress, allContracts.GatewayEVMImplementation, provider);
    await compareBytecode(contractName, allContracts.GatewayEVMImplementation, provider);
    await checkGatewayEVMState(contractName, contractAddress, allContracts, provider);
  } else {
    console.log(`    ⚠️  No implementation found`);
  }
}

async function checkERC20Custody(
  contractName: string,
  contractAddress: string,
  allContracts: Record<string, string>,
  provider: JsonRpcProvider
) {
  console.log(`\n  🏗️  ERC20Custody (Proxy):`);

  if (allContracts.ERC20CustodyImplementation) {
    await checkProxy(contractAddress, allContracts.ERC20CustodyImplementation, provider);
    await compareBytecode(contractName, allContracts.ERC20CustodyImplementation, provider);
    await checkERC20CustodyState(contractName, contractAddress, allContracts, provider);
  } else {
    console.log(`    ⚠️  No implementation found`);
  }
}

async function checkGatewayZEVM(
  contractName: string,
  contractAddress: string,
  allContracts: Record<string, string>,
  provider: JsonRpcProvider
) {
  console.log(`\n  🏗️  GatewayZEVM (Proxy):`);

  if (allContracts.GatewayZEVMImplementation) {
    await checkProxy(contractAddress, allContracts.GatewayZEVMImplementation, provider);
    await compareBytecode(contractName, allContracts.GatewayZEVMImplementation, provider);
    await checkGatewayZEVMState(contractName, contractAddress, allContracts, provider);
  } else {
    console.log(`    ⚠️  No implementation found`);
  }
}

async function checksumNetworks() {
  const addresses = loadAddresses();
  const CHAIN_ID_TO_NETWORK = await fetchSupportedChains();

  console.log("🔍 Contract checksum verification:");
  console.log("=".repeat(60));

  for (const [chainId, contracts] of Object.entries(addresses)) {
    const contractsTyped = contracts as Record<string, string>;
    const networkName = CHAIN_ID_TO_NETWORK[chainId];

    console.log(`\nChain ID: ${chainId} (${networkName})`);
    console.log("-".repeat(50));

    if (!networkName) {
      console.log(`  ⚠️  No network configuration found for chain ${chainId}`);
      continue;
    }

    try {
      const networkConfig = hre.config.networks[networkName] as any;
      const provider = new hre.ethers.JsonRpcProvider(networkConfig.url);

      for (const [contractName, contractAddress] of Object.entries(contractsTyped)) {
        switch (contractName) {
          case "GatewayEVM":
            await checkGatewayEVM(contractName, contractAddress, contractsTyped, provider);
            break;
          case "ERC20Custody":
            await checkERC20Custody(contractName, contractAddress, contractsTyped, provider);
            break;
          case "GatewayZEVM":
            await checkGatewayZEVM(contractName, contractAddress, contractsTyped, provider);
            break;
        }
      }
    } catch (error) {
      console.error(`  ❌ Error connecting to ${networkName}: ${error}`);
    }
  }
}

export async function main() {
  await checksumNetworks();
}
