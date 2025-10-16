// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { ERC20Custody } from "../../contracts/evm/ERC20Custody.sol";
import "./BaseForkTest.sol";

contract ERC20CustodyForkTest is BaseForkTest {
    // ERC20Custody proxy addresses.
    address constant ETHEREUM_ERC20CUSTODY_PROXY = 0x0Bad40D9e9C369f2223c835E108f43a45fd223B5;
    address constant BSC_ERC20CUSTODY_PROXY = 0x0Bad40D9e9C369f2223c835E108f43a45fd223B5;
    address constant POLYGON_ERC20CUSTODY_PROXY = 0x0Bad40D9e9C369f2223c835E108f43a45fd223B5;
    address constant BASE_ERC20CUSTODY_PROXY = 0x0Bad40D9e9C369f2223c835E108f43a45fd223B5;
    address constant ARBITRUM_ERC20CUSTODY_PROXY = 0xECe33274237E6422f2668eD7dEE5901b16336aA0;
    address constant AVALANCHE_ERC20CUSTODY_PROXY = 0xECe33274237E6422f2668eD7dEE5901b16336aA0;

    function _setupChains() internal override {
        chains.push(
            ChainConfig(ethereumForkId, ETHEREUM_ERC20CUSTODY_PROXY, ETHEREUM_ADMIN, ETHEREUM_RPC_URL, "Ethereum")
        );
        chains.push(ChainConfig(bscForkId, BSC_ERC20CUSTODY_PROXY, BSC_ADMIN, BSC_RPC_URL, "BSC"));
        chains.push(ChainConfig(polygonForkId, POLYGON_ERC20CUSTODY_PROXY, POLYGON_ADMIN, POLYGON_RPC_URL, "Polygon"));
        chains.push(ChainConfig(baseForkId, BASE_ERC20CUSTODY_PROXY, BASE_ADMIN, BASE_RPC_URL, "Base"));
        chains.push(
            ChainConfig(arbitrumForkId, ARBITRUM_ERC20CUSTODY_PROXY, ARBITRUM_ADMIN, ARBITRUM_RPC_URL, "Arbitrum")
        );
        chains.push(
            ChainConfig(avalancheForkId, AVALANCHE_ERC20CUSTODY_PROXY, AVALANCHE_ADMIN, AVALANCHE_RPC_URL, "Avalanche")
        );
    }

    function _testUpgradeContract(ChainConfig memory config) internal override {
        vm.selectFork(config.forkId);

        // Get the current proxy contract.
        ERC20Custody custody = ERC20Custody(config.contractAddress);

        // Storage state before the upgrade.
        address gatewayBefore = address(custody.gateway());
        address tssAddressBefore = custody.tssAddress();
        bool supportsLegacyBefore = custody.supportsLegacy();

        // Deploy new implementation.
        vm.startPrank(config.admin);
        ERC20Custody localNewImplementation = new ERC20Custody();

        // Upgrade the proxy.
        custody.upgradeToAndCall(address(localNewImplementation), "");

        // Verify upgrade.
        address onChainNewImplementation = _getImplementation(address(custody));
        assertEq(onChainNewImplementation, address(localNewImplementation), "Upgrade failed.");
        console.log("Successfully upgraded to:", onChainNewImplementation);

        // Verify state preservation.
        assertEq(address(custody.gateway()), gatewayBefore, "Gateway address changed.");
        assertEq(custody.tssAddress(), tssAddressBefore, "TSS address changed.");
        assertEq(custody.supportsLegacy(), supportsLegacyBefore, "Supports legacy changed.");
        console.log("State preserved after upgrade");

        vm.stopPrank();
    }
}
