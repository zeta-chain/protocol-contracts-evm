// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "../helpers/BaseRegistry.sol";
import "./interfaces/IGatewayEVM.sol";
import "./interfaces/IRegistry.sol";

/// @title Registry
/// @notice Satellite registry contract for connected chains, receiving updates from CoreRegistry.
/// @dev This contract is deployed on every connected chain and maintains a synchronized view of the registry.
contract Registry is BaseRegistry, IRegistry {
    /// @notice Identifier for the gateway role
    bytes32 public constant GATEWAY_ROLE = keccak256("GATEWAY_ROLE");
    /// @notice GatewayEVM contract that will call this contract with messages from CoreRegistry
    IGatewayEVM public gatewayEVM;
    /// @notice Represents the address of the CoreRegistry contract on the ZetaChain
    address public coreRegistry;

    /// @dev Only registry address allowed modifier.
    /// @notice Restricts function calls to only be made by this contract itself
    /// @dev This is used to ensure functions receiving cross-chain messages can only be called through
    /// the onCall function using a self-call pattern, preventing direct external calls to these functions
    modifier onlyRegistry() {
        if (msg.sender != address(this)) {
            revert InvalidSender();
        }
        _;
    }

    /// @notice Initialize the Registry contract
    /// @param admin_ Address with DEFAULT_ADMIN_ROLE, authorized for upgrades and pausing actions
    /// @param registryManager_ Address with REGISTRY_MANAGER_ROLE, authorized for all registry write actions.
    /// @param gatewayEVM_ Address of the GatewayEVM contract for cross-chain messaging
    /// @param coreRegistry_ Address of the CoreRegistry contract deployed on ZetaChain
    function initialize(address admin_, address registryManager_, address gatewayEVM_, address coreRegistry_)
        public
        initializer
    {
        if (
            admin_ == address(0) || gatewayEVM_ == address(0) || coreRegistry_ == address(0)
                || registryManager_ == address(0)
        ) {
            revert ZeroAddress();
        }

        __UUPSUpgradeable_init();
        __AccessControl_init_unchained();
        __Pausable_init_unchained();

        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(REGISTRY_MANAGER_ROLE, registryManager_);
        _grantRole(PAUSER_ROLE, registryManager_);
        _grantRole(PAUSER_ROLE, admin_);
        _grantRole(GATEWAY_ROLE, gatewayEVM_);

        admin = admin_;
        registryManager = registryManager_;
        gatewayEVM = IGatewayEVM(gatewayEVM_);
        coreRegistry = coreRegistry_;
    }

    /// @notice onCall is called by the GatewayEVM when a cross-chain message is received
    /// @param context Information about the cross-chain message
    /// @param data The encoded function call to execute
    function onCall(MessageContext calldata context, bytes calldata data)
        external
        onlyRole(GATEWAY_ROLE)
        whenNotPaused
        returns (bytes memory)
    {
        if (context.sender != coreRegistry) {
            revert InvalidSender();
        }

        // Execute a self-call with the encoded function data
        // This allows us to reuse existing function signatures and validation logic
        // By using a self-call pattern combined with onlyRegistry modifier, these functions
        // can only be called through this onCall entry point
        (bool success, bytes memory result) = address(this).call(data);
        if (!success) {
            // Extract the error message from the revert
            assembly {
                revert(add(result, 32), mload(result))
            }
        }

        return result;
    }

    /// @notice Changes status of the chain to activated/deactivated
    /// @dev Only callable through onCall from CoreRegistry
    /// @param chainId The ID of the chain being activated/deactivated.
    /// @param gasZRC20 The address of the ZRC20 token that represents gas token for the chain.
    /// @param registry Address of the Registry contract on the connected chain.
    /// @param activation Whether activate or deactivate the chain
    function changeChainStatus(uint256 chainId, address gasZRC20, bytes calldata registry, bool activation)
        external
        onlyRegistry
        whenNotPaused
    {
        _changeChainStatus(chainId, gasZRC20, registry, activation);
        emit ChainStatusChanged(chainId, activation);
    }

    /// @notice Updates chain metadata, only for the active chains
    /// @dev Only callable through onCall from CoreRegistry
    /// @param chainId The ID of the chain
    /// @param key The metadata key to update
    /// @param value The new value for the metadata
    function updateChainMetadata(uint256 chainId, string calldata key, bytes calldata value)
        external
        onlyRegistry
        whenNotPaused
    {
        _updateChainMetadata(chainId, key, value);
        emit ChainMetadataUpdated(chainId, key, value);
    }

    /// @notice Registers a new contract address for a specific chain
    /// @dev Only callable through onCall from CoreRegistry
    /// @param chainId The ID of the chain where the contract is deployed
    /// @param contractType The type of the contract (e.g., "connector", "gateway")
    /// @param addressBytes The address of the contract
    function registerContract(uint256 chainId, string calldata contractType, bytes calldata addressBytes)
        external
        onlyRegistry
        whenNotPaused
    {
        _registerContract(chainId, contractType, addressBytes);
        emit ContractRegistered(chainId, contractType, addressBytes);
    }

    /// @notice Updates contract configuration
    /// @dev Only callable through onCall from CoreRegistry
    /// @param chainId The ID of the chain where the contract is deployed
    /// @param contractType The type of the contract
    /// @param key The configuration key to update
    /// @param value The new value for the configuration
    function updateContractConfiguration(
        uint256 chainId,
        string calldata contractType,
        string calldata key,
        bytes calldata value
    )
        external
        onlyRegistry
        whenNotPaused
    {
        _updateContractConfiguration(chainId, contractType, key, value);
        emit ContractConfigurationUpdated(chainId, contractType, key, value);
    }

    /// @notice Sets a contract's active status
    /// @dev Only callable through onCall from CoreRegistry
    /// @param chainId The ID of the chain where the contract is deployed
    /// @param contractType The type of the contract
    /// @param active Whether the contract should be active
    function setContractActive(uint256 chainId, string calldata contractType, bool active) external onlyRegistry {
        _setContractActive(chainId, contractType, active);
        emit ContractStatusChanged(_contracts[chainId][contractType].addressBytes);
    }

    /// @notice Registers a new ZRC20 token in the registry
    /// @dev Only callable through onCall from CoreRegistry
    /// @param address_ The address of the ZRC20 token on ZetaChain
    /// @param symbol The symbol of the token
    /// @param originChainId The ID of the foreign chain where the original asset exists
    /// @param originAddress The address or identifier of the asset on its native chain
    /// @param coinType The type of the original coin
    /// @param decimals The number of decimals the token uses
    function registerZRC20Token(
        address address_,
        string calldata symbol,
        uint256 originChainId,
        bytes calldata originAddress,
        string calldata coinType,
        uint8 decimals
    )
        external
        onlyRegistry
        whenNotPaused
    {
        _registerZRC20Token(address_, symbol, originChainId, originAddress, coinType, decimals);
        emit ZRC20TokenRegistered(originAddress, address_, decimals, originChainId, symbol);
    }

    /// @notice Updates ZRC20 token active status
    /// @dev Only callable through onCall from CoreRegistry
    /// @param address_ The address of the ZRC20 token
    /// @param active Whether the token should be active
    function setZRC20TokenActive(address address_, bool active) external onlyRegistry whenNotPaused {
        _setZRC20TokenActive(address_, active);
        emit ZRC20TokenUpdated(address_, active);
    }

    /// @notice Bootstrap the registry with chain data
    /// @dev This function can only be called by an address with the REGISTRY_MANAGER_ROLE.
    /// @param chains Array of chain data structures to bootstrap
    /// @param metadataEntries Array of chain metadata entries
    function bootstrapChains(ChainInfoDTO[] calldata chains, ChainMetadataEntry[] calldata metadataEntries)
        external
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        // Process chain data
        for (uint256 i = 0; i < chains.length; i++) {
            ChainInfoDTO calldata chainData = chains[i];
            _changeChainStatus(chainData.chainId, chainData.gasZRC20, chainData.registry, chainData.active);
        }

        // Process metadata entries
        for (uint256 i = 0; i < metadataEntries.length; i++) {
            ChainMetadataEntry calldata metadata = metadataEntries[i];
            _updateChainMetadata(metadata.chainId, metadata.key, metadata.value);
        }
    }

    /// @notice Bootstrap the registry with contract data
    /// @dev This function can only be called by an address with the REGISTRY_MANAGER_ROLE.
    /// @param contracts Array of contract data structures to bootstrap
    /// @param configEntries Array of contract configuration entries
    function bootstrapContracts(ContractInfoDTO[] calldata contracts, ContractConfigEntry[] calldata configEntries)
        external
        onlyRole(REGISTRY_MANAGER_ROLE)
        whenNotPaused
    {
        // Process contract data
        for (uint256 i = 0; i < contracts.length; i++) {
            ContractInfoDTO calldata contractData = contracts[i];
            _registerContract(contractData.chainId, contractData.contractType, contractData.addressBytes);
        }

        // Process configuration entries
        for (uint256 i = 0; i < configEntries.length; i++) {
            ContractConfigEntry calldata configuration = configEntries[i];
            _updateContractConfiguration(
                configuration.chainId, configuration.contractType, configuration.key, configuration.value
            );
        }
    }

    /// @notice Bootstrap the registry with ZRC20 token data
    /// @dev This function can only be called by an address with the REGISTRY_MANAGER_ROLE.
    /// @param tokens Array of ZRC20 token data structures to bootstrap
    function bootstrapZRC20Tokens(ZRC20Info[] calldata tokens) external onlyRole(REGISTRY_MANAGER_ROLE) whenNotPaused {
        // Process ZRC20 token data
        for (uint256 i = 0; i < tokens.length; i++) {
            ZRC20Info calldata tokenData = tokens[i];
            _registerZRC20Token(
                tokenData.address_,
                tokenData.symbol,
                tokenData.originChainId,
                tokenData.originAddress,
                tokenData.coinType,
                tokenData.decimals
            );
        }
    }
}
