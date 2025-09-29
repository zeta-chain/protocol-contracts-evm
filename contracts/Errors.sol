// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

/// @title INotSupportedMethods
/// @notice Interface for contracts that with non supported methods.
interface INotSupportedMethods {
    error ZETANotSupported();
    error CallOnRevertNotSupported();
}
