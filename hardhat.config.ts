import "@nomicfoundation/hardhat-toolbox";
import "./tasks/addresses";
import { getHardhatConfigNetworks } from "@zetachain/networks";
import { HardhatUserConfig } from "hardhat/config";
import { task } from "hardhat/config";


const config: HardhatUserConfig = {
  networks: {
    ...getHardhatConfigNetworks(),
  },
  solidity: {
    version: "0.8.26",
    settings: {
      evmVersion: "cancun",
      optimizer: {
        enabled: true,
        runs: 10000
      }
    }
  }
};

task("upgradeProposal", "Create Safe upgrade proposals for GatewayEVM")
  .addFlag("testnet", "Create proposals for testnet networks")
  .addFlag("mainnet", "Create proposals for mainnet networks")
  .addOptionalParam("upgradeData", "Upgrade data to include in the proposal", "0x")
  .setAction(async (taskArgs, hre) => {
    try {
      const isMainnet = taskArgs.mainnet;
      const networkType = isMainnet ? 'mainnet' : 'testnet';
      console.log(`🚀 Creating upgrade proposals for ${networkType.toUpperCase()} networks`);

      process.env.NETWORK_TYPE = networkType;
      if (taskArgs.upgradeData) {
        process.env.UPGRADE_DATA = taskArgs.upgradeData;
      }

      const upgradeScript = await import("./scripts/safe/create-upgrade-proposal");
      await upgradeScript.main();
    } catch (error) {
      console.error("❌ Error running upgrade proposal script:", error);
      process.exit(1);
    }
  });

task("protocolChecksum", "Run EVM protocol checkers verification")
  .addFlag("testnet", "Run checkers for testnet networks")
  .addFlag("mainnet", "Run checkers for mainnet networks")
  .setAction(async (taskArgs, hre) => {
    try {
      const isMainnet = taskArgs.mainnet;
      const networkType = isMainnet ? 'mainnet' : 'testnet';
      console.log(`🔍 Running checksum for ${networkType.toUpperCase()} networks`);

      process.env.NETWORK_TYPE = networkType;

      const checksumScript = await import("./scripts/checkers/protocolChecksum");
      await checksumScript.main();
    } catch (error) {
      console.error("❌ Error running protocol checkers script:", error);
      process.exit(1);
    }
  });

task("contractDiff", "Fetch and flatten smart contracts for diff analysis")
    .addParam("oldAddress", "Address of the old implementation contract")
    .addParam("newAddress", "Address of the new implementation contract")
    .setAction(async (taskArgs, hre) => {
      try {
        console.log("🔍 Contract Diff Tool");
        console.log("=".repeat(60));
        console.log(`Network: ${hre.network.name}`);
        console.log(`Old Implementation: ${taskArgs.oldAddress}`);
        console.log(`New Implementation: ${taskArgs.newAddress}`);
        console.log("=".repeat(60));

        const contractDiffTool = await import("./scripts/contractDiff/contractDiffTool");
        await contractDiffTool.fetchAndFlattenContract(
            taskArgs.oldAddress,
            taskArgs.newAddress,
            hre.network.name
        )
      } catch (error) {
        console.error("❌ Error running contract diff tool:", error);
        process.exit(1);
      }
    });

export default config;
