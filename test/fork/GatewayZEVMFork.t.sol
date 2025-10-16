// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "./BaseForkTest.sol";
import {GatewayZEVM} from "../../contracts/zevm/GatewayZEVM.sol";

contract GatewayZEVMForkTest is BaseForkTest {
    // GatewayZEVM proxy address.
    address constant ZETACHAIN_GATEWAY_PROXY = 0xfEDD7A6e3Ef1cC470fbfbF955a22D793dDC0F44E;

    function _setupChains() internal override {
        chains.push(ChainConfig(zetachainForkId, ZETACHAIN_GATEWAY_PROXY, ZETACHAIN_ADMIN, ZETACHAIN_RPC_URL, "ZetaChain"));
    }

    function _testUpgradeContract(ChainConfig memory config) internal override {
        vm.selectFork(config.forkId);

        // Get the current proxy contract.
        GatewayZEVM gateway = GatewayZEVM(payable(config.contractAddress));

        // Storage state before the upgrade.
        address zetaTokenBefore = gateway.zetaToken();

        // Deploy new implementation.
        vm.startPrank(config.admin);
        GatewayZEVM localNewImplementation = new GatewayZEVM();

        // Upgrade the proxy.
        gateway.upgradeToAndCall(address(localNewImplementation), "");

        // Verify upgrade.
        address onChainNewImplementation = _getImplementation(address(gateway));
        assertEq(onChainNewImplementation, address(localNewImplementation), "Upgrade failed.");
        console.log("Successfully upgraded to:", onChainNewImplementation);

        // Verify state preservation.
        assertEq(gateway.zetaToken(), zetaTokenBefore, "Zeta token address changed.");
        console.log("State preserved after upgrade");

        vm.stopPrank();
    }
}
