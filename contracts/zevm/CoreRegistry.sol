// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "../helpers/BaseRegistry.sol";
import "./interfaces/IGatewayZEVM.sol";
import "./interfaces/IZRC20.sol";

/// @title CoreRegistry
/// @notice Central registry for ZetaChain, managing chain info, ZRC20 data, and contract addresses across all chains.
/// @dev The contract doesn't hold any funds and should never have active allowances.
contract CoreRegistry is BaseRegistry {
    /// @notice Cross-chain message gas limit
    uint256 public constant CROSS_CHAIN_GAS_LIMIT = 500_000;
    /// @notice Instance of the GatewayZEVM contract for cross-chain communication
    IGatewayZEVM public gatewayZEVM;

    /// @notice Initialize the CoreRegistry contract.
    /// @param admin_ Address with DEFAULT_ADMIN_ROLE, authorized for upgrades and pausing actions.
    /// @param registryManager_ Address with REGISTRY_MANAGER_ROLE, authorized for all registry write actions.
    /// @param gatewayZEVM_ Address of the GatewayZEVM contract for cross-chain messaging
    function initialize(address admin_, address registryManager_, address gatewayZEVM_) public initializer {
        if (admin_ == address(0) || registryManager_ == address(0) || gatewayZEVM_ == address(0)) {
            revert ZeroAddress();
        }

        __UUPSUpgradeable_init();
        __AccessControl_init_unchained();
        __Pausable_init_unchained();

        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(REGISTRY_MANAGER_ROLE, registryManager_);
        _grantRole(PAUSER_ROLE, registryManager_);
        _grantRole(PAUSER_ROLE, admin_);

        admin = admin_;
        registryManager = registryManager_;
        gatewayZEVM = IGatewayZEVM(gatewayZEVM_);

        // Add ZetaChain to the list of supported networks
        _chains[block.chainid].active = true;
        _chains[block.chainid].registry = abi.encodePacked(address(this));
        _activeChains.push(block.chainid);
        _allChains.push(block.chainid);
    }

    /// @notice Changes status of the chain to activated/deactivated.
    /// @param chainId The ID of the chain to activate.
    /// @param gasZRC20 The address of the ZRC20 token that represents gas token for the chain.
    /// @param registry Address of the Registry contract on the connected chain.
    /// @param activation Whether activate or deactivate the chain
    function changeChainStatus(
        uint256 chainId,
        address gasZRC20,
        bytes calldata registry,
        bool activation
    )
        external
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        // Change state on ZetaChain
        _changeChainStatus(chainId, gasZRC20, registry, activation);
        // Broadcast update to satellite registries
        _broadcastChainActivation(chainId, gasZRC20, registry, activation);
        emit ChainStatusChanged(chainId, activation);
    }

    /// @notice Updates chain metadata, only for the active chains.
    /// @param chainId The ID of the chain.
    /// @param key The metadata key to update.
    /// @param value The new value for the metadata.
    function updateChainMetadata(
        uint256 chainId,
        string calldata key,
        bytes calldata value
    )
        external
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        // Change state on ZetaChain
        _updateChainMetadata(chainId, key, value);
        // Broadcast update to satellite registries
        _broadcastChainMetadataUpdate(chainId, key, value);
        emit ChainMetadataUpdated(chainId, key, value);
    }

    /// @notice Registers a new contract address for a specific chain.
    /// @param chainId The ID of the chain where the contract is deployed.
    /// @param contractType The type of the contract (e.g., "connector", "gateway").
    /// @param addressBytes The bytes representation of the non-EVM address.
    function registerContract(
        uint256 chainId,
        string calldata contractType,
        bytes calldata addressBytes
    )
        external
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        // Change state on ZetaChain
        _registerContract(chainId, contractType, addressBytes);
        // Broadcast update to satellite registries
        _broadcastContractRegistration(chainId, contractType, addressBytes);
        emit ContractRegistered(chainId, contractType, addressBytes);
    }

    /// @notice Updates contract configuration.
    /// @param chainId The ID of the chain where the contract is deployed.
    /// @param contractType The type of the contract.
    /// @param key The configuration key to update.
    /// @param value The new value for the configuration.
    function updateContractConfiguration(
        uint256 chainId,
        string calldata contractType,
        string calldata key,
        bytes calldata value
    )
        external
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        // Change state on ZetaChain
        _updateContractConfiguration(chainId, contractType, key, value);
        // Broadcast update to satellite registries
        _broadcastContractConfigUpdate(chainId, contractType, key, value);
        emit ContractConfigurationUpdated(chainId, contractType, key, value);
    }

    /// @notice Sets a contract's active status
    /// @param chainId The ID of the chain where the contract is deployed.
    /// @param contractType The type of the contract.
    /// @param active Whether the contract should be active.
    function setContractActive(
        uint256 chainId,
        string calldata contractType,
        bool active
    )
        external
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        // Change state on ZetaChain
        _setContractActive(chainId, contractType, active);
        // Broadcast update to satellite registries
        _broadcastContractStatusUpdate(chainId, contractType, active);
        emit ContractStatusChanged(_contracts[chainId][contractType].addressBytes);
    }

    /// @notice Registers a new ZRC20 token in the registry.
    /// @param address_ The address of the ZRC20 token on ZetaChain.
    /// @param symbol The symbol of the token.
    /// @param originChainId The ID of the foreign chain where the original asset exists.
    /// @param originAddress The address or identifier of the asset on its native chain.
    /// @param coinType The type of the original coin.
    /// @param decimals The number of decimals the token uses.
    function registerZRC20Token(
        address address_,
        string calldata symbol,
        uint256 originChainId,
        bytes calldata originAddress,
        string calldata coinType,
        uint8 decimals
    )
        external
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        // Change state on ZetaChain
        _registerZRC20Token(address_, symbol, originChainId, originAddress, coinType, decimals);
        // Broadcast update to satellite registries
        _broadcastZRC20Registration(address_, symbol, originChainId, originAddress, coinType, decimals);
        emit ZRC20TokenRegistered(originAddress, address_, decimals, originChainId, symbol);
    }

    /// @notice Updates ZRC20 token active status.
    function setZRC20TokenActive(address address_, bool active) external onlyRole(REGISTRY_MANAGER_ROLE) whenNotPaused {
        // Change state on ZetaChain
        _setZRC20TokenActive(address_, active);
        // Broadcast update to satellite registries
        _broadcastZRC20Update(address_, active);
        emit ZRC20TokenUpdated(address_, active);
    }

    /// @notice Broadcast chain activation update to all satellite registries.
    /// @param chainId The ID of the chain being activated/deactivated.
    /// @param gasZRC20 The address of the ZRC20 token that represents gas token for the chain.
    /// @param registry Address of the Registry contract on the connected chain.
    /// @param activation Whether activate or deactivate the chain
    function _broadcastChainActivation(
        uint256 chainId,
        address gasZRC20,
        bytes calldata registry,
        bool activation
    )
        internal
    {
        // Encode the function call for the Registry contract on the target chain
        bytes memory message = abi.encodeWithSignature(
            "changeChainStatus(uint256,address,bytes,bool)", chainId, gasZRC20, registry, activation
        );
        _broadcastToAllChains(message);
    }

    /// @notice Broadcast chain metadata to all satellite registries
    /// @param chainId The ID of the chain whose metadata is being updated
    /// @param key The metadata key being updated
    /// @param value The new value for the metadata
    function _broadcastChainMetadataUpdate(uint256 chainId, string calldata key, bytes calldata value) private {
        // Encode the function call for the Registry contract on the target chain
        bytes memory message = abi.encodeWithSignature("updateChainMetadata(uint256,string,bytes)", chainId, key, value);
        _broadcastToAllChains(message);
    }

    /// @notice Broadcast contract registration to all satellite registries
    /// @param chainId The ID of the chain where the contract is deployed
    /// @notice contractType The type of the contract
    /// @notice addressBytes The bytes representation of the non-EVM address
    function _broadcastContractRegistration(
        uint256 chainId,
        string calldata contractType,
        bytes calldata addressBytes
    )
        private
    {
        bytes memory message =
            abi.encodeWithSignature("registerContract(uint256,string,bytes)", chainId, contractType, addressBytes);
        _broadcastToAllChains(message);
    }

    /// @notice Broadcast contract configuration update to all satellite registries
    /// @param chainId The ID of the chain where the contract is deployed
    /// @notice contractType The type of the contract
    /// @notice key The configuration key being updated
    /// @notice value The new value for the configuration
    function _broadcastContractConfigUpdate(
        uint256 chainId,
        string calldata contractType,
        string calldata key,
        bytes calldata value
    )
        private
    {
        // Encode the function call for the Registry contract on the target chain
        bytes memory message = abi.encodeWithSignature(
            "updateContractConfiguration(uint256,string,string,bytes)", chainId, contractType, key, value
        );
        _broadcastToAllChains(message);
    }

    /// @notice Broadcast contract status update to all satellite registries
    /// @param chainId The ID of the chain where the contract is deployed
    /// @notice contractType The type of the contract
    /// @notice active Whether the contract should be active
    function _broadcastContractStatusUpdate(uint256 chainId, string calldata contractType, bool active) private {
        // Encode the function call for the Registry contract on the target chain
        bytes memory message =
            abi.encodeWithSignature("setContractActive(uint256,string,bool)", chainId, contractType, active);
        _broadcastToAllChains(message);
    }

    /// @notice Broadcast ZRC20 token registration to all satellite registries
    /// @param address_ The address of the ZRC20 token on ZetaChain
    /// @param symbol The symbol of the token
    /// @param originChainId The ID of the foreign chain where the original asset exists
    /// @param originAddress The address or identifier of the asset on its native chain
    /// @param coinType The type of the original coin
    /// @param decimals The number of decimals the token uses
    function _broadcastZRC20Registration(
        address address_,
        string calldata symbol,
        uint256 originChainId,
        bytes calldata originAddress,
        string calldata coinType,
        uint8 decimals
    )
        private
    {
        // Encode the function call for the Registry contract on the target chain
        bytes memory message = abi.encodeWithSignature(
            "registerZRC20Token(address,string,uint256,bytes,string,uint8)",
            address_,
            symbol,
            originChainId,
            originAddress,
            coinType,
            decimals
        );
        _broadcastToAllChains(message);
    }

    /// @notice Broadcast ZRC20 token update to all satellite registries
    /// @param address_ The address of the ZRC20 token
    /// @param active Whether the token should be active
    function _broadcastZRC20Update(address address_, bool active) private {
        // Encode the function call for the Registry contract on the target chain
        bytes memory message = abi.encodeWithSignature("setZRC20TokenActive(address,bool)", address_, active);
        _broadcastToAllChains(message);
    }

    /// @notice Generic function to broadcast encoded messages to all satellite registries
    /// @param encodedMessage The fully encoded function call to broadcast
    function _broadcastToAllChains(bytes memory encodedMessage) private {
        for (uint256 i = 0; i < _activeChains.length; i++) {
            if (_activeChains[i] == block.chainid || _chains[_activeChains[i]].registry.length == 0) {
                continue;
            }
            _sendCrossChainMessage(_activeChains[i], encodedMessage);
        }
    }

    /// @notice Sends a cross-chain message to the Registry contract on a target chain.
    /// @param targetChainId The ID of the chain to send the message to.
    /// @param message The encoded function call to execute on the target chain.
    function _sendCrossChainMessage(uint256 targetChainId, bytes memory message) private {
        // Prepare call options
        CallOptions memory callOptions = CallOptions({ gasLimit: CROSS_CHAIN_GAS_LIMIT, isArbitraryCall: false });

        // Prepare revert options
        RevertOptions memory revertOptions;

        // Approve gas token amount for cctx
        address gasZRC20 = _chains[targetChainId].gasZRC20;
        (, uint256 gasFee) = IZRC20(gasZRC20).withdrawGasFeeWithGasLimit(CROSS_CHAIN_GAS_LIMIT);
        if (!IZRC20(gasZRC20).transferFrom(msg.sender, address(this), gasFee)) {
            revert TransferFailed();
        }
        IZRC20(gasZRC20).approve(address(gatewayZEVM), gasFee);

        gatewayZEVM.call(_chains[targetChainId].registry, gasZRC20, message, callOptions, revertOptions);
    }
}
