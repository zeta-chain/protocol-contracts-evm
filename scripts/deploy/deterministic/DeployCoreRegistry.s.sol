// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "contracts/zevm/CoreRegistry.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployCoreRegistry is Script {
    function run() external {
        address admin = vm.envAddress("REGISTRY_ADMIN_ADDRESS");
        address registryManager = vm.envAddress("REGISTRY_MANAGER_ADDRESS");
        address gatewayZEVM = vm.envAddress("GATEWAY_ZEVM_ADDRESS");

        address expectedImplAddress;
        address expectedProxyAddress;

        bytes32 implSalt = keccak256("DeployCoreRegistry");
        bytes32 proxySalt = keccak256("DeployCoreRegistryProxy");

        // Add this specific check to ensure contract is not deployed at deterministic address with wrong admin
        require(admin != address(0), "Environment variable REGISTRY_ADMIN_ADDRESS is not set");
        require(registryManager != address(0), "Environment variable REGISTRY_MANAGER_ADDRESS is not set");
        require(gatewayZEVM != address(0), "Environment variable GATEWAY_ZEVM_ADDRESS is not set");

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

        // Compute expected proxy address
        expectedProxyAddress = vm.computeCreate2Address(
            proxySalt,
            hashInitCode(
                type(ERC1967Proxy).creationCode,
                abi.encode(
                    address(registryImpl),
                    abi.encodeWithSelector(CoreRegistry.initialize.selector, admin, registryManager, gatewayZEVM)
                )
            )
        );

        // Deploy the proxy contract using CREATE2
        ERC1967Proxy registryProxy = new ERC1967Proxy{salt: proxySalt}(
            address(registryImpl),
            abi.encodeWithSelector(CoreRegistry.initialize.selector, admin, registryManager, gatewayZEVM)
        );
        require(address(registryProxy) != address(0), "registryProxy deployment failed");
        require(expectedProxyAddress == address(registryProxy), "proxy address doesn't match expected address");

        CoreRegistry registry = CoreRegistry(address(registryProxy));

        // Verify initial configuration
        require(address(registry.gatewayZEVM()) == gatewayZEVM, "gatewayZEVM not set correctly");
        require(registry.hasRole(registry.DEFAULT_ADMIN_ROLE(), admin), "admin role not set correctly");
        require(registry.hasRole(registry.REGISTRY_MANAGER_ROLE(), registryManager), "registry manager role not set correctly");

        vm.stopBroadcast();

        console.log("CoreRegistry Implementation deployed at: %s", address(registryImpl));
        console.log("CoreRegistry Proxy deployed at: %s", address(registryProxy));
    }
}