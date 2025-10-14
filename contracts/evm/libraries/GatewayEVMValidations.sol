// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "../../../contracts/Errors.sol";
import "../../../contracts/Revert.sol";
import "../interfaces/IGatewayEVM.sol";

/// @title GatewayEVMValidations
/// @notice Library containing validation functions for GatewayEVM contract.
/// @dev This library provides common validation logic used across GatewayEVM contract.
library GatewayEVMValidations {
    /// @notice Maximum payload size constant.
    uint256 internal constant MAX_PAYLOAD_SIZE = 2880;

    /// @dev Validates that an address is not zero.
    /// @param addr The address to validate.
    function validateNonZeroAddress(address addr) internal pure {
        if (addr == address(0)) revert IGatewayEVMErrors.ZeroAddress();
    }

    /// @dev Validates that amount is not zero.
    /// @param amount The amount to validate.
    function validateAmount(uint256 amount) internal pure {
        if (amount == 0) revert IGatewayEVMErrors.InsufficientEVMAmount();
    }

    /// @dev Validates payload size constraints.
    /// @param payloadLength The length of the main payload.
    /// @param revertMessageLength The length of the revert message.
    function validatePayloadSize(uint256 payloadLength, uint256 revertMessageLength) internal pure {
        uint256 totalSize = payloadLength + revertMessageLength;
        if (payloadLength + revertMessageLength > MAX_PAYLOAD_SIZE) {
            revert IGatewayEVMErrors.PayloadSizeExceeded(totalSize, MAX_PAYLOAD_SIZE);
        }
    }

    /// @dev Validates on revert gas limit.
    /// @param revertOptions The revert options to validate.
    function validateGasLimitRevertOptions(RevertOptions calldata revertOptions) internal pure {
        if (revertOptions.onRevertGasLimit > MAX_REVERT_GAS_LIMIT) {
            revert RevertGasLimitExceeded(revertOptions.onRevertGasLimit, MAX_REVERT_GAS_LIMIT);
        }
    }

    /// @dev Validates on revert message length.
    /// @param revertOptions The revert options to validate.
    function validateRevertMessageLength(RevertOptions calldata revertOptions) internal pure {
        if (revertOptions.revertMessage.length > MAX_PAYLOAD_SIZE) {
            revert IGatewayEVMErrors.PayloadSizeExceeded(revertOptions.revertMessage.length, MAX_PAYLOAD_SIZE);
        }
    }

    /// @dev Validates revert options for call operations (includes callOnRevert check).
    /// @param revertOptions The revert options to validate.
    function validateCallRevertOptions(RevertOptions calldata revertOptions) internal pure {
        if (revertOptions.callOnRevert) revert INotSupportedMethods.CallOnRevertNotSupported();
        validateGasLimitRevertOptions(revertOptions);
    }

    /// @dev Validates standard deposit parameters.
    /// @param receiver The receiver address.
    /// @param amount The amount to deposit.
    /// @param revertOptions The revert options.
    function validateDepositParams(
        address receiver,
        uint256 amount,
        RevertOptions calldata revertOptions
    )
        internal
        pure
    {
        validateNonZeroAddress(receiver);
        validateAmount(amount);
        validateRevertMessageLength(revertOptions);
        validateGasLimitRevertOptions(revertOptions);
    }

    /// @dev Validates deposit and call parameters.
    /// @param receiver The receiver address.
    /// @param amount The amount to deposit.
    /// @param payload The payload to send.
    /// @param revertOptions The revert options.
    function validateDepositAndCallParams(
        address receiver,
        uint256 amount,
        bytes calldata payload,
        RevertOptions calldata revertOptions
    )
        internal
        pure
    {
        validateNonZeroAddress(receiver);
        validateAmount(amount);
        validatePayloadSize(payload.length, revertOptions.revertMessage.length);
        validateGasLimitRevertOptions(revertOptions);
    }

    /// @dev Validates call parameters.
    /// @param receiver The receiver address.
    /// @param payload The payload to send.
    /// @param revertOptions The revert options.
    function validateCallParams(
        address receiver,
        bytes calldata payload,
        RevertOptions calldata revertOptions
    )
        internal
        pure
    {
        validateNonZeroAddress(receiver);
        validateCallRevertOptions(revertOptions);
        validatePayloadSize(payload.length, revertOptions.revertMessage.length);
    }
}
