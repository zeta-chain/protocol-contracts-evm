// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/console.sol";
import "forge-std/Script.sol";
import "contracts/evm/GatewayEVM.sol";

contract UpgradeSimulation is Script {
    function run() external {
        // Forked environment
        address proxyAddress = vm.envAddress("PROXY_ADDRESS");
        address adminAddress = vm.envAddress("ADMIN_ADDRESS");

        // Deploy the new implementation contract
        bytes32 implSalt = keccak256("GatewayEVM");
        address gatewayImpl = address(new GatewayEVM{salt: implSalt}());

        GatewayEVM proxy = GatewayEVM(proxyAddress);

        // Get the current state
        address custody = proxy.custody();
        address tssAddress = proxy.tssAddress();
        address zetaConnector = proxy.zetaConnector();
        address zetaToken = proxy.zetaToken();

        // Simulate the upgrade
        vm.prank(adminAddress);
        proxy.upgradeToAndCall(gatewayImpl, "");

        // After upgrade, verify that the state is intact
        require(custody == proxy.custody(), "custody address mismatch");
        require(tssAddress == proxy.tssAddress(), "tss address mismatch");
        require(zetaConnector == proxy.zetaConnector(), "zetaConnector address mismatch");
        require(zetaToken == proxy.zetaToken(), "zetaToken address mismatch");

        console.log("Upgraded contract state:");
        console.log("custody address:", proxy.custody());
        console.log("tss address:", proxy.tssAddress());
        console.log("zetaConnector address:", proxy.zetaConnector());
        console.log("zetaToken address:", proxy.zetaToken());
    }
}