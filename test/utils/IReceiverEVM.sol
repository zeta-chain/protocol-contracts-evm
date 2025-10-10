// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { RevertContext } from "../../contracts/Revert.sol";

/// @title IReceiverEVMEvents
/// @notice Interface for the events emitted by the ReceiverEVM contract.
interface IReceiverEVMEvents {
    /// @notice Emitted when a payable function is called.
    /// @param sender The address of the sender.
    /// @param value The amount of ETH sent with the call.
    /// @param str A string parameter.
    /// @param num A numeric parameter.
    /// @param flag A boolean parameter.
    event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag);

    /// @notice Emitted when a non-payable function is called.
    /// @param sender The address of the sender.
    /// @param strs An array of string parameters.
    /// @param nums An array of numeric parameters.
    /// @param flag A boolean parameter.
    event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag);

    /// @notice Emitted when ERC20 tokens are received.
    /// @param sender The address of the sender.
    /// @param amount The amount of ERC20 tokens received.
    /// @param token The address of the ERC20 token.
    /// @param destination The address to which the tokens are sent.
    event ReceivedERC20(address sender, uint256 amount, address token, address destination);

    /// @notice Emitted when a function without parameters is called.
    /// @param sender The address of the sender.
    event ReceivedNoParams(address sender);

    /// @notice Emitted when a revert callback function is called.
    /// @param sender The address of the sender.
    /// @param revertContext Revert Context.
    event ReceivedRevert(address sender, RevertContext revertContext);

    /// @notice Emitted when onCall function is called with authenticated call.
    /// @param sender Message context sender.
    /// @param message Message received.
    event ReceivedOnCall(address sender, bytes message);

    /// @notice Emitted when onCall function is called with new MessageContext call.
    /// @param sender Message context sender.
    /// @param message Message received.
    event ReceivedOnCallV2(address sender, address asset, uint256 amount, bytes message);
}
