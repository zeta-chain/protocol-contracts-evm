// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

/// @title IRegistry
interface IRegistry {
    /// @notice Structure for metadata entries used during bootstrapping
    struct ChainMetadataEntry {
        /// @notice The unique identifier of the chain.
        uint256 chainId;
        /// @param key The metadata key to update.
        string key;
        /// @param value The new value for the metadata.
        bytes value;
    }

    /// @notice Structure for contract configuration entries used during bootstrapping
    struct ContractConfigEntry {
        /// @notice Represents id of the chain where contract is deployed.
        uint256 chainId;
        /// @notice The type of the contract (e.g. "connector", "gateway", "tss").
        string contractType;
        /// @param key The configuration key to update.
        string key;
        /// @param value The new value for the configuration.
        bytes value;
    }
}
