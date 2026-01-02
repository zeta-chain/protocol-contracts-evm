// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/console.sol";
import "forge-std/Script.sol";
import "contracts/zevm/GatewayZEVM.sol";

contract UpgradeSimulation is Script {
    function run() external {
        // Forked environment
        address payable proxyAddress = payable(vm.envAddress("PROXY_ADDRESS"));
        address adminAddress = vm.envAddress("ADMIN_ADDRESS");

        // Deploy the new implementation contract
        bytes32 implSalt = keccak256("GatewayZEVM");
        address gatewayImpl = address(new GatewayZEVM{salt: implSalt}());

        GatewayZEVM proxy = GatewayZEVM(proxyAddress);

        // Get the current state
        address zetaToken = proxy.zetaToken();

        // Simulate the upgrade
        vm.prank(adminAddress);
        proxy.upgradeToAndCall(gatewayImpl, "");

        // After upgrade, verify that the state is intact
        require(zetaToken == proxy.zetaToken(), "zetaToken address mismatch");

        console.log("Upgraded contract state:");
        console.log("zetaToken address:", proxy.zetaToken());
    }
}