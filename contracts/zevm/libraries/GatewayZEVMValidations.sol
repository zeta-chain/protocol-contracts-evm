// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "../../../contracts/Revert.sol";
import "../interfaces/IGatewayZEVM.sol";

/// @title GatewayZEVMValidations
/// @notice Library containing validation functions for GatewayZEVM contract
/// @dev This library provides common validation logic used across GatewayZEVM contract
library GatewayZEVMValidations {
    /// @notice Maximum message size constant
    uint256 internal constant MAX_MESSAGE_SIZE = 2880;
    /// @notice Minimum gas limit constant
    uint256 internal constant MIN_GAS_LIMIT = 100_000;
    /// @notice Maximum gas limit constant
    uint256 internal constant MAX_GAS_LIMIT = 2_500_000;

    /// @dev Validates that an address is not zero
    /// @param addr The address to validate
    function validateNonZeroAddress(address addr) internal pure {
        if (addr == address(0)) revert IGatewayZEVMErrors.EmptyAddress();
    }

    /// @dev Validates that receiver bytes are not empty
    /// @param receiver The receiver bytes to validate
    function validateReceiver(bytes memory receiver) internal pure {
        if (receiver.length == 0) revert IGatewayZEVMErrors.EmptyAddress();
    }

    /// @dev Validates that amount is not zero
    /// @param amount The amount to validate
    function validateAmount(uint256 amount) internal pure {
        if (amount == 0) revert IGatewayZEVMErrors.InsufficientAmount();
    }

    /// @dev Validates that gas limit meets minimum requirement
    /// @param gasLimit The gas limit to validate
    function validateGasLimit(uint256 gasLimit) internal pure {
        if (gasLimit < MIN_GAS_LIMIT || gasLimit > MAX_GAS_LIMIT) revert IGatewayZEVMErrors.InvalidGasLimit();
    }

    /// @dev Validates that target address is not restricted
    /// @param target The target address to validate
    /// @param protocolAddress The protocol address to check against
    /// @param contractAddress The contract address to check against
    function validateTarget(address target, address protocolAddress, address contractAddress) private pure {
        if (target == protocolAddress || target == contractAddress) revert IGatewayZEVMErrors.InvalidTarget();
    }

    /// @dev Validates message size constraints
    /// @param messageLength The length of the main message
    /// @param revertMessageLength The length of the revert message
    function validateMessageSize(uint256 messageLength, uint256 revertMessageLength) internal pure {
        uint256 totalSize = messageLength + revertMessageLength;
        if (totalSize > MAX_MESSAGE_SIZE) {
            revert IGatewayZEVMErrors.MessageSizeExceeded(totalSize, MAX_MESSAGE_SIZE);
        }
    }

    /// @dev Validates revert options
    /// @param revertOptions The revert options to validate
    function validateRevertOptions(RevertOptions calldata revertOptions) internal pure {
        if (revertOptions.revertMessage.length > MAX_MESSAGE_SIZE) {
            revert IGatewayZEVMErrors.MessageSizeExceeded(revertOptions.revertMessage.length, MAX_MESSAGE_SIZE);
        }

        if (revertOptions.onRevertGasLimit > MAX_REVERT_GAS_LIMIT) {
            revert RevertGasLimitExceeded(revertOptions.onRevertGasLimit, MAX_REVERT_GAS_LIMIT);
        }
    }

    /// @dev Validates call options and revert options together
    /// @param callOptions The call options to validate
    /// @param revertOptions The revert options to validate
    /// @param messageLength The length of the message
    function validateCallAndRevertOptions(
        CallOptions calldata callOptions,
        RevertOptions calldata revertOptions,
        uint256 messageLength
    )
        internal
        pure
    {
        validateGasLimit(callOptions.gasLimit);
        validateMessageSize(messageLength, revertOptions.revertMessage.length);
        if (revertOptions.onRevertGasLimit > MAX_REVERT_GAS_LIMIT) {
            revert RevertGasLimitExceeded(revertOptions.onRevertGasLimit, MAX_REVERT_GAS_LIMIT);
        }
    }

    /// @dev Validates standard withdrawal parameters
    /// @param receiver The receiver address
    /// @param amount The amount to withdraw
    /// @param revertOptions The revert options
    function validateWithdrawalParams(
        bytes memory receiver,
        uint256 amount,
        RevertOptions calldata revertOptions
    )
        internal
        pure
    {
        validateReceiver(receiver);
        validateAmount(amount);
        validateRevertOptions(revertOptions);
    }

    /// @dev Validates withdrawal and call parameters
    /// @param receiver The receiver address
    /// @param amount The amount to withdraw
    /// @param message The message to send
    /// @param callOptions The call options
    /// @param revertOptions The revert options
    function validateWithdrawalAndCallParams(
        bytes memory receiver,
        uint256 amount,
        bytes calldata message,
        CallOptions calldata callOptions,
        RevertOptions calldata revertOptions
    )
        internal
        pure
    {
        validateReceiver(receiver);
        validateAmount(amount);
        validateCallAndRevertOptions(callOptions, revertOptions, message.length);
    }

    /// @dev Validates deposit parameters
    /// @param zrc20 The ZRC20 token address
    /// @param amount The amount to deposit
    /// @param target The target address
    /// @param protocolAddress The protocol address
    /// @param contractAddress The contract address
    function validateDepositParams(
        address zrc20,
        uint256 amount,
        address target,
        address protocolAddress,
        address contractAddress
    )
        internal
        pure
    {
        validateNonZeroAddress(zrc20);
        validateNonZeroAddress(target);
        validateAmount(amount);
        validateTarget(target, protocolAddress, contractAddress);
    }

    /// @dev Validates execute parameters
    /// @param zrc20 The ZRC20 token address
    /// @param target The target address
    function validateExecuteParams(address zrc20, address target) internal pure {
        validateNonZeroAddress(zrc20);
        validateNonZeroAddress(target);
    }

    /// @dev Validates ZETA deposit and call parameters
    /// @param amount The amount to deposit
    /// @param target The target address
    /// @param protocolAddress The protocol address
    /// @param contractAddress The contract address
    function validateZetaDepositParams(
        uint256 amount,
        address target,
        address protocolAddress,
        address contractAddress
    )
        internal
        pure
    {
        validateNonZeroAddress(target);
        validateAmount(amount);
        validateTarget(target, protocolAddress, contractAddress);
    }
}
