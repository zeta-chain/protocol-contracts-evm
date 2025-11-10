// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { GatewayEVM } from "../../contracts/evm/GatewayEVM.sol";
import "./BaseForkTest.sol";

contract GatewayEVMForkTest is BaseForkTest {
    // GatewayEVM proxy addresses.
    address constant ETHEREUM_GATEWAY_PROXY = 0x48B9AACC350b20147001f88821d31731Ba4C30ed;
    address constant BSC_GATEWAY_PROXY = 0x48B9AACC350b20147001f88821d31731Ba4C30ed;
    address constant POLYGON_GATEWAY_PROXY = 0x48B9AACC350b20147001f88821d31731Ba4C30ed;
    address constant BASE_GATEWAY_PROXY = 0x48B9AACC350b20147001f88821d31731Ba4C30ed;
    address constant ARBITRUM_GATEWAY_PROXY = 0x1C53e188Bc2E471f9D4A4762CFf843d32C2C8549;
    address constant AVALANCHE_GATEWAY_PROXY = 0x1C53e188Bc2E471f9D4A4762CFf843d32C2C8549;

    function _setupChains() internal override {
        chains.push(ChainConfig(ethereumForkId, ETHEREUM_GATEWAY_PROXY, ETHEREUM_ADMIN, ETHEREUM_RPC_URL, "Ethereum"));
        chains.push(ChainConfig(bscForkId, BSC_GATEWAY_PROXY, BSC_ADMIN, BSC_RPC_URL, "BSC"));
        chains.push(ChainConfig(polygonForkId, POLYGON_GATEWAY_PROXY, POLYGON_ADMIN, POLYGON_RPC_URL, "Polygon"));
        chains.push(ChainConfig(baseForkId, BASE_GATEWAY_PROXY, BASE_ADMIN, BASE_RPC_URL, "Base"));
        chains.push(ChainConfig(arbitrumForkId, ARBITRUM_GATEWAY_PROXY, ARBITRUM_ADMIN, ARBITRUM_RPC_URL, "Arbitrum"));
        chains.push(
            ChainConfig(avalancheForkId, AVALANCHE_GATEWAY_PROXY, AVALANCHE_ADMIN, AVALANCHE_RPC_URL, "Avalanche")
        );
    }

    function _testUpgradeContract(ChainConfig memory config) internal override {
        vm.selectFork(config.forkId);

        // Get the current proxy contract.
        GatewayEVM gateway = GatewayEVM(config.contractAddress);

        // Storage state before the upgrade.
        address custodyBefore = gateway.custody();
        address tssAddressBefore = gateway.tssAddress();
        address zetaConnectorBefore = gateway.zetaConnector();
        address zetaTokenBefore = gateway.zetaToken();

        // Deploy new implementation.
        vm.startPrank(config.admin);
        GatewayEVM localNewImplementation = new GatewayEVM();

        // Upgrade the proxy.
        gateway.upgradeToAndCall(address(localNewImplementation), "");

        // Verify upgrade.
        address onChainNewImplementation = _getImplementation(address(gateway));
        assertEq(onChainNewImplementation, address(localNewImplementation), "Upgrade failed.");
        console.log("Successfully upgraded to:", onChainNewImplementation);

        // Verify state preservation.
        assertEq(gateway.custody(), custodyBefore, "Custody address changed.");
        assertEq(gateway.tssAddress(), tssAddressBefore, "TSS address changed.");
        assertEq(gateway.zetaConnector(), zetaConnectorBefore, "Connector address changed.");
        assertEq(gateway.zetaToken(), zetaTokenBefore, "Zeta token address changed.");
        console.log("State preserved after upgrade");

        // Check admin still has DEFAULT_ADMIN_ROLE after upgrade.
        bytes32 adminRole = gateway.DEFAULT_ADMIN_ROLE();
        bool adminHasRoleAfter = gateway.hasRole(adminRole, config.admin);
        assertTrue(adminHasRoleAfter, "Admin lost DEFAULT_ADMIN_ROLE after upgrade.");

        vm.stopPrank();
    }
}
