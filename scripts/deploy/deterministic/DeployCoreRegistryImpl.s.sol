// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "contracts/zevm/CoreRegistry.sol";

contract DeployCoreRegistryImpl is Script {
    function run() external {
        address expectedImplAddress;
        bytes32 implSalt = keccak256("DeployCoreRegistryV1");

        vm.startBroadcast();

        // Compute expected implementation address
        expectedImplAddress = vm.computeCreate2Address(
            implSalt,
            hashInitCode(type(CoreRegistry).creationCode)
        );

        // Deploy the implementation contract using CREATE2
        CoreRegistry registryImpl = new CoreRegistry{salt: implSalt}();
        require(address(registryImpl) != address(0), "registryImpl deployment failed");
        require(expectedImplAddress == address(registryImpl), "impl address doesn't match expected address");

        vm.stopBroadcast();

        console.log("CoreRegistry Implementation deployed at: %s", address(registryImpl));
    }
}