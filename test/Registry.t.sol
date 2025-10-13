// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

import { Registry } from "../contracts/evm/Registry.sol";
import "../contracts/evm/interfaces/IGatewayEVM.sol";
import "../contracts/helpers/interfaces/IBaseRegistry.sol";

// Mock GatewayEVM
contract MockGatewayEVM {
    event CallEmitted(address receiver, bytes message, LegacyMessageContext messageContext);

    function onCall(LegacyMessageContext calldata messageContext, bytes calldata data) external returns (bytes memory) {
        emit CallEmitted(msg.sender, data, messageContext);
        return bytes("");
    }
}

contract RegistryTest is Test, IBaseRegistryErrors, IBaseRegistryEvents {
    address payable proxy;
    Registry registry;
    MockGatewayEVM mockGateway;

    address admin;
    address registryManager;
    address coreRegistry;
    address user;

    // Helper variables for tests
    uint256 chainId = 101;
    address gasZRC20 = address(0xcbaa);
    bytes registryAddress = abi.encodePacked(address(0x9876));
    address contractAddress = address(0xAA55);
    string contractType = "connector";
    bytes addressBytes = abi.encodePacked(contractAddress);

    function setUp() public {
        admin = address(0xABCD);
        coreRegistry = address(0x1234);
        user = address(0x5678);
        registryManager = address(0x3333);

        mockGateway = new MockGatewayEVM();

        proxy = payable(
            Upgrades.deployUUPSProxy(
                "Registry.sol",
                abi.encodeCall(Registry.initialize, (admin, registryManager, address(mockGateway), coreRegistry))
            )
        );

        registry = Registry(proxy);
    }

    function testInitialize() public view {
        assertTrue(registry.hasRole(registry.DEFAULT_ADMIN_ROLE(), admin));
        assertTrue(registry.hasRole(registry.PAUSER_ROLE(), admin));
        assertTrue(registry.hasRole(registry.PAUSER_ROLE(), registryManager));
        assertTrue(registry.hasRole(registry.REGISTRY_MANAGER_ROLE(), registryManager));
        assertTrue(registry.hasRole(registry.GATEWAY_ROLE(), address(mockGateway)));
        assertEq(address(registry.gatewayEVM()), address(mockGateway));
        assertEq(registry.coreRegistry(), coreRegistry);
    }

    function testInitializeWithZeroAddress() public {
        address payable uninitializedProxy = payable(Upgrades.deployUUPSProxy("Registry.sol", ""));

        Registry uninitializedRegistry = Registry(uninitializedProxy);

        vm.expectRevert(ZeroAddress.selector);
        uninitializedRegistry.initialize(address(0), registryManager, address(mockGateway), coreRegistry);

        vm.expectRevert(ZeroAddress.selector);
        uninitializedRegistry.initialize(admin, address(0), address(mockGateway), coreRegistry);

        vm.expectRevert(ZeroAddress.selector);
        uninitializedRegistry.initialize(admin, registryManager, address(0), coreRegistry);

        vm.expectRevert(ZeroAddress.selector);
        uninitializedRegistry.initialize(admin, registryManager, address(mockGateway), address(0));
    }

    function testPause() public {
        assertFalse(registry.paused());

        vm.prank(admin);
        registry.pause();
        assertTrue(registry.paused());

        vm.prank(admin);
        registry.unpause();
        assertFalse(registry.paused());
    }

    function testPauseUnauthorized() public {
        vm.prank(user);
        vm.expectRevert();
        registry.pause();

        vm.prank(user);
        vm.expectRevert();
        registry.unpause();
    }

    function testOnCallInvalidSender() public {
        LegacyMessageContext memory context = LegacyMessageContext({ sender: address(0x7777) });

        vm.prank(address(mockGateway));
        vm.expectRevert(InvalidSender.selector);
        registry.onCall(context, bytes(""));
    }

    function testOnCallPaused() public {
        vm.prank(admin);
        registry.pause();

        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });

        vm.prank(address(mockGateway));
        vm.expectRevert();
        registry.onCall(context, bytes(""));
    }

    function testOnCallUnauthorized() public {
        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });

        vm.prank(user);
        vm.expectRevert();
        registry.onCall(context, bytes(""));
    }

    function testChangeChainStatus() public {
        bytes memory activateData = abi.encodeWithSignature(
            "changeChainStatus(uint256,address,bytes,bool)", chainId, gasZRC20, registryAddress, true
        );
        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });

        vm.prank(address(mockGateway));
        vm.expectEmit(true, true, true, true);
        emit ChainStatusChanged(chainId, true);
        registry.onCall(context, activateData);

        uint256[] memory chains = registry.getActiveChains();
        assertEq(chains.length, 1);
        assertEq(chains[0], chainId);

        bytes memory deactivateData = abi.encodeWithSignature(
            "changeChainStatus(uint256,address,bytes,bool)", chainId, gasZRC20, registryAddress, false
        );

        vm.prank(address(mockGateway));
        vm.expectEmit(true, true, true, true);
        emit ChainStatusChanged(chainId, false);
        registry.onCall(context, deactivateData);

        chains = registry.getActiveChains();
        assertEq(chains.length, 0);
    }

    function testChangeStatusDirectCallFails() public {
        vm.prank(user);
        vm.expectRevert();
        registry.changeChainStatus(chainId, address(1), registryAddress, true);

        vm.prank(admin);
        vm.expectRevert();
        registry.changeChainStatus(chainId, address(1), registryAddress, true);
    }

    function testUpdateChainMetadata() public {
        bytes memory activateData = abi.encodeWithSignature(
            "changeChainStatus(uint256,address,bytes,bool)", chainId, gasZRC20, registryAddress, true
        );
        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        registry.onCall(context, activateData);

        string memory key = "blockTime";
        bytes memory value = abi.encode(5);
        bytes memory updateData =
            abi.encodeWithSignature("updateChainMetadata(uint256,string,bytes)", chainId, key, value);

        vm.prank(address(mockGateway));
        vm.expectEmit(true, true, true, true);
        emit ChainMetadataUpdated(chainId, key, value);
        registry.onCall(context, updateData);

        bytes memory storedValue = registry.getChainMetadata(chainId, key);
        assertEq(keccak256(storedValue), keccak256(value));
    }

    function testRegisterContract() public {
        bytes memory activateData = abi.encodeWithSignature(
            "changeChainStatus(uint256,address,bytes,bool)", chainId, gasZRC20, registryAddress, true
        );
        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        registry.onCall(context, activateData);

        bytes memory registerData =
            abi.encodeWithSignature("registerContract(uint256,string,bytes)", chainId, contractType, addressBytes);

        vm.prank(address(mockGateway));
        vm.expectEmit(true, true, true, true);
        emit ContractRegistered(chainId, contractType, addressBytes);
        registry.onCall(context, registerData);

        (bool active, bytes memory storedAddress) = registry.getContractInfo(chainId, contractType);
        assertTrue(active);
        assertEq(storedAddress, addressBytes);
    }

    function testRegisterContractWithEmptyType() public {
        bytes memory activateData = abi.encodeWithSignature(
            "changeChainStatus(uint256,address,bytes,bool)", chainId, gasZRC20, registryAddress, true
        );
        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        registry.onCall(context, activateData);

        string memory emptyType = "";
        bytes memory registerData =
            abi.encodeWithSignature("registerContract(uint256,string,bytes)", chainId, emptyType, addressBytes);

        vm.prank(address(mockGateway));
        vm.expectRevert(abi.encodeWithSelector(InvalidContractType.selector, emptyType));
        registry.onCall(context, registerData);
    }

    function testRegisterContractWithEmptyAddress() public {
        bytes memory activateData = abi.encodeWithSignature(
            "changeChainStatus(uint256,address,bytes,bool)", chainId, gasZRC20, registryAddress, true
        );
        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        registry.onCall(context, activateData);

        bytes memory emptyAddressBytes = bytes("");
        bytes memory registerData =
            abi.encodeWithSignature("registerContract(uint256,string,bytes)", chainId, contractType, emptyAddressBytes);

        vm.prank(address(mockGateway));
        vm.expectRevert(ZeroAddress.selector);
        registry.onCall(context, registerData);
    }

    function testUpdateContractConfiguration() public {
        bytes memory activateData = abi.encodeWithSignature(
            "changeChainStatus(uint256,address,bytes,bool)", chainId, gasZRC20, registryAddress, true
        );
        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        registry.onCall(context, activateData);

        bytes memory registerData =
            abi.encodeWithSignature("registerContract(uint256,string,bytes)", chainId, contractType, addressBytes);
        vm.prank(address(mockGateway));
        registry.onCall(context, registerData);

        string memory key = "gasLimit";
        bytes memory value = abi.encode(300_000);
        bytes memory updateData = abi.encodeWithSignature(
            "updateContractConfiguration(uint256,string,string,bytes)", chainId, contractType, key, value
        );

        vm.prank(address(mockGateway));
        vm.expectEmit(true, true, true, true);
        emit ContractConfigurationUpdated(chainId, contractType, key, value);
        registry.onCall(context, updateData);

        bytes memory storedValue = registry.getContractConfiguration(chainId, contractType, key);
        assertEq(keccak256(storedValue), keccak256(value));
    }

    function testUpdateContractConfigurationWithInvalidType() public {
        bytes memory activateData = abi.encodeWithSignature(
            "changeChainStatus(uint256,address,bytes,bool)", chainId, gasZRC20, registryAddress, true
        );
        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        registry.onCall(context, activateData);

        string memory key = "gasLimit";
        bytes memory value = abi.encode(300_000);
        string memory invalidType = "";
        bytes memory updateData = abi.encodeWithSignature(
            "updateContractConfiguration(uint256,string,string,bytes)", chainId, invalidType, key, value
        );

        vm.prank(address(mockGateway));
        vm.expectRevert(abi.encodeWithSelector(InvalidContractType.selector, invalidType));
        registry.onCall(context, updateData);
    }

    function testSetContractActive() public {
        bytes memory activateData = abi.encodeWithSignature(
            "changeChainStatus(uint256,address,bytes,bool)", chainId, gasZRC20, registryAddress, true
        );
        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        registry.onCall(context, activateData);

        bytes memory registerData =
            abi.encodeWithSignature("registerContract(uint256,string,bytes)", chainId, contractType, addressBytes);
        vm.prank(address(mockGateway));
        registry.onCall(context, registerData);

        (bool active,) = registry.getContractInfo(chainId, contractType);
        assertTrue(active);

        bytes memory deactivateData =
            abi.encodeWithSignature("setContractActive(uint256,string,bool)", chainId, contractType, false);

        vm.prank(address(mockGateway));
        vm.expectEmit(true, true, true, true);
        emit ContractStatusChanged(addressBytes);
        registry.onCall(context, deactivateData);

        (active,) = registry.getContractInfo(chainId, contractType);
        assertFalse(active);

        bytes memory activateContractData =
            abi.encodeWithSignature("setContractActive(uint256,string,bool)", chainId, contractType, true);

        vm.prank(address(mockGateway));
        vm.expectEmit(true, true, true, true);
        emit ContractStatusChanged(addressBytes);
        registry.onCall(context, activateContractData);

        (active,) = registry.getContractInfo(chainId, contractType);
        assertTrue(active);
    }

    function testSetContractActiveWithInvalidType() public {
        bytes memory activateData = abi.encodeWithSignature(
            "changeChainStatus(uint256,address,bytes,bool)", chainId, gasZRC20, registryAddress, true
        );
        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        registry.onCall(context, activateData);

        string memory invalidType = "";
        bytes memory updateData =
            abi.encodeWithSignature("setContractActive(uint256,string,bool)", chainId, invalidType, false);

        vm.prank(address(mockGateway));
        vm.expectRevert(abi.encodeWithSelector(InvalidContractType.selector, invalidType));
        registry.onCall(context, updateData);
    }

    function testRegisterZRC20Token() public {
        address zrc20Address = address(0xDDDD);
        string memory symbol = "TKN";
        uint256 originChainId = 101;
        bytes memory originAddress = abi.encodePacked(address(0x1357));
        string memory coinType = "ERC20";
        uint8 decimals = 18;

        bytes memory registerData = abi.encodeWithSignature(
            "registerZRC20Token(address,string,uint256,bytes,string,uint8)",
            zrc20Address,
            symbol,
            originChainId,
            originAddress,
            coinType,
            decimals
        );

        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        vm.expectEmit(true, true, true, true);
        emit ZRC20TokenRegistered(originAddress, zrc20Address, decimals, originChainId, symbol);
        registry.onCall(context, registerData);

        (
            bool active,
            string memory storedSymbol,
            uint256 storedOriginChainId,
            bytes memory storedOriginAddress,
            string memory storedCoinType,
            uint8 storedDecimals
        ) = registry.getZRC20TokenInfo(zrc20Address);

        assertTrue(active);
        assertEq(storedSymbol, symbol);
        assertEq(storedOriginChainId, originChainId);
        assertEq(keccak256(storedOriginAddress), keccak256(originAddress));
        assertEq(storedCoinType, coinType);
        assertEq(storedDecimals, decimals);
    }

    function testRegisterZRC20TokenWithZeroAddress() public {
        address zrc20Address = address(0);
        string memory symbol = "TKN";
        uint256 originChainId = 101;
        bytes memory originAddress = abi.encodePacked(address(0x1357));
        string memory coinType = "ERC20";
        uint8 decimals = 18;

        bytes memory registerData = abi.encodeWithSignature(
            "registerZRC20Token(address,string,uint256,bytes,string,uint8)",
            zrc20Address,
            symbol,
            originChainId,
            originAddress,
            coinType,
            decimals
        );

        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        vm.expectRevert(ZeroAddress.selector);
        registry.onCall(context, registerData);
    }

    function testRegisterZRC20TokenWithEmptySymbol() public {
        address zrc20Address = address(0xDDDD);
        string memory symbol = "";
        uint256 originChainId = 101;
        bytes memory originAddress = abi.encodePacked(address(0x1357));
        string memory coinType = "ERC20";
        uint8 decimals = 18;

        bytes memory registerData = abi.encodeWithSignature(
            "registerZRC20Token(address,string,uint256,bytes,string,uint8)",
            zrc20Address,
            symbol,
            originChainId,
            originAddress,
            coinType,
            decimals
        );

        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        vm.expectRevert(abi.encodeWithSelector(InvalidContractType.selector, "Symbol cannot be empty"));
        registry.onCall(context, registerData);
    }

    function testRegisterZRC20TokenWithEmptyOriginAddress() public {
        address zrc20Address = address(0xDDDD);
        string memory symbol = "TKN";
        uint256 originChainId = 101;
        bytes memory originAddress = bytes("");
        string memory coinType = "ERC20";
        uint8 decimals = 18;

        bytes memory registerData = abi.encodeWithSignature(
            "registerZRC20Token(address,string,uint256,bytes,string,uint8)",
            zrc20Address,
            symbol,
            originChainId,
            originAddress,
            coinType,
            decimals
        );

        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        vm.expectEmit(true, true, true, true);
        emit ZRC20TokenRegistered(originAddress, zrc20Address, decimals, originChainId, symbol);
        registry.onCall(context, registerData);
    }

    function testSetZRC20TokenActive() public {
        address zrc20Address = address(0xDDDD);
        string memory symbol = "TKN";
        uint256 originChainId = 101;
        bytes memory originAddress = abi.encodePacked(address(0x1357));
        string memory coinType = "ERC20";
        uint8 decimals = 18;

        bytes memory registerData = abi.encodeWithSignature(
            "registerZRC20Token(address,string,uint256,bytes,string,uint8)",
            zrc20Address,
            symbol,
            originChainId,
            originAddress,
            coinType,
            decimals
        );

        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        registry.onCall(context, registerData);

        (bool active,,,,,) = registry.getZRC20TokenInfo(zrc20Address);
        assertTrue(active);

        bytes memory deactivateData = abi.encodeWithSignature("setZRC20TokenActive(address,bool)", zrc20Address, false);

        vm.prank(address(mockGateway));
        vm.expectEmit(true, true, true, true);
        emit ZRC20TokenUpdated(zrc20Address, false);
        registry.onCall(context, deactivateData);

        (active,,,,,) = registry.getZRC20TokenInfo(zrc20Address);
        assertFalse(active);

        bytes memory activateData = abi.encodeWithSignature("setZRC20TokenActive(address,bool)", zrc20Address, true);

        vm.prank(address(mockGateway));
        vm.expectEmit(true, true, true, true);
        emit ZRC20TokenUpdated(zrc20Address, true);
        registry.onCall(context, activateData);

        (active,,,,,) = registry.getZRC20TokenInfo(zrc20Address);
        assertTrue(active);
    }

    function testSetZRC20TokenActiveWithZeroAddress() public {
        address zrc20Address = address(0);

        bytes memory updateData = abi.encodeWithSignature("setZRC20TokenActive(address,bool)", zrc20Address, false);

        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        vm.expectRevert(ZeroAddress.selector);
        registry.onCall(context, updateData);
    }

    function testMultipleActiveChains() public {
        uint256 chainId1 = 101;
        bytes memory activateData1 = abi.encodeWithSignature(
            "changeChainStatus(uint256,address,bytes,bool)", chainId1, gasZRC20, registryAddress, true
        );

        LegacyMessageContext memory context = LegacyMessageContext({ sender: coreRegistry });
        vm.prank(address(mockGateway));
        registry.onCall(context, activateData1);

        uint256 chainId2 = 102;
        bytes memory activateData2 = abi.encodeWithSignature(
            "changeChainStatus(uint256,address,bytes,bool)", chainId2, gasZRC20, registryAddress, true
        );

        vm.prank(address(mockGateway));
        registry.onCall(context, activateData2);

        uint256[] memory chains = registry.getActiveChains();
        assertEq(chains.length, 2);

        bytes memory deactivateData = abi.encodeWithSignature(
            "changeChainStatus(uint256,address,bytes,bool)", chainId1, gasZRC20, registryAddress, false
        );

        vm.prank(address(mockGateway));
        registry.onCall(context, deactivateData);

        chains = registry.getActiveChains();
        assertEq(chains.length, 1);
        assertEq(chains[0], chainId2);
    }
}
