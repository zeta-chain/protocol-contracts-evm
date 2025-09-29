// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "contracts/evm/Registry.sol";

contract DeployRegistryImpl is Script {
    function run() external {
        address expectedImplAddress;
        bytes32 implSalt = keccak256("DeployRegistryV1");

        vm.startBroadcast();

        // Compute expected implementation address
        expectedImplAddress = vm.computeCreate2Address(
            implSalt,
            hashInitCode(type(Registry).creationCode)
        );

        // Deploy the implementation contract using CREATE2
        Registry registryImpl = new Registry{salt: implSalt}();
        require(address(registryImpl) != address(0), "registryImpl deployment failed");
        require(expectedImplAddress == address(registryImpl), "impl address doesn't match expected address");

        vm.stopBroadcast();

        console.log("Registry Implementation deployed at: %s", address(registryImpl));
    }
}