// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "contracts/evm/Registry.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

contract DeployRegistry is Script {
    function run() external {
        address admin = vm.envAddress("REGISTRY_ADMIN_ADDRESS");
        address registryManager = vm.envAddress("REGISTRY_MANAGER_ADDRESS");
        address gatewayEVM = vm.envAddress("GATEWAY_EVM_ADDRESS");
        address coreRegistry = vm.envAddress("CORE_REGISTRY_ADDRESS");

        require(admin != address(0), "Environment variable REGISTRY_ADMIN_ADDRESS is not set");
        require(registryManager != address(0), "Environment variable REGISTRY_MANAGER_ADDRESS is not set");
        require(gatewayEVM != address(0), "Environment variable GATEWAY_EVM_ADDRESS is not set");
        require(coreRegistry != address(0), "Environment variable CORE_REGISTRY_ADDRESS is not set");

        bytes32 implSalt = keccak256("DeployRegistry");
        bytes32 proxySalt = keccak256("DeployRegistryProxy");

        vm.startBroadcast();

        // Deploy implementation and compute its address
        Registry registryImpl = new Registry{salt: implSalt}();

        // Prepare initialization data with all required parameters
        bytes memory initData = abi.encodeWithSelector(
            Registry.initialize.selector,
            admin,
            registryManager,
            gatewayEVM,
            coreRegistry
        );

        // Deploy proxy
        ERC1967Proxy registryProxy = new ERC1967Proxy{salt: proxySalt}(
            address(registryImpl),
            initData
        );

        // Create registry instance
        Registry registry = Registry(address(registryProxy));

        // Split verification into separate helper function to reduce stack pressure
        verifyDeployment(registry, admin, gatewayEVM, coreRegistry);

        vm.stopBroadcast();

        console.log("Registry Implementation deployed at: %s", address(registryImpl));
        console.log("Registry Proxy deployed at: %s", address(registryProxy));
    }

    function verifyDeployment(
        Registry registry,
        address admin,
        address gatewayEVM,
        address coreRegistry
    ) internal view {
        // Verify initial configuration
        require(address(registry.gatewayEVM()) == gatewayEVM, "gatewayEVM not set correctly");
        require(registry.coreRegistry() == coreRegistry, "coreRegistry not set correctly");
        require(registry.hasRole(registry.DEFAULT_ADMIN_ROLE(), admin), "admin role not set correctly");
        require(registry.hasRole(registry.GATEWAY_ROLE(), gatewayEVM), "gateway role not set correctly");
    }
}