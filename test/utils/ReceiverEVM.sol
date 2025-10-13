// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { RevertContext } from "../../contracts/Revert.sol";
import {LegacyMessageContext, MessageContext} from "../../contracts/evm/interfaces/IGatewayEVM.sol";
import "./IReceiverEVM.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/// @title ReceiverEVM
/// @notice This contract is used just for testing purposes.
/// @dev Implements various functions to receive ETH and ERC20 tokens, and emit corresponding events.
contract ReceiverEVM is IReceiverEVMEvents, ReentrancyGuard {
    using SafeERC20 for IERC20;

    /// @notice Error indicating that the amount is zero.
    error ZeroAmount();

    /// @notice Receives ETH with additional data.
    /// @param str A string parameter.
    /// @param num A numeric parameter.
    /// @param flag A boolean parameter.
    /// @dev Emits the ReceivedPayable event.
    function receivePayable(string memory str, uint256 num, bool flag) external payable {
        emit ReceivedPayable(msg.sender, msg.value, str, num, flag);
    }

    /// @notice Receives data without transferring ETH.
    /// @param strs An array of string parameters.
    /// @param nums An array of numeric parameters.
    /// @param flag A boolean parameter.
    /// @dev Emits the ReceivedNonPayable event.
    function receiveNonPayable(string[] memory strs, uint256[] memory nums, bool flag) external {
        emit ReceivedNonPayable(msg.sender, strs, nums, flag);
    }

    /// @notice Receives ERC20 tokens and transfers them to a destination address.
    /// @param amount The amount of tokens to transfer.
    /// @param token The address of the ERC20 token.
    /// @param destination The address to send the tokens to.
    /// @dev Emits the ReceivedERC20 event.
    function receiveERC20(uint256 amount, address token, address destination) external nonReentrant {
        IERC20(token).safeTransferFrom(msg.sender, destination, amount);

        emit ReceivedERC20(msg.sender, amount, token, destination);
    }

    /// @notice Receives ERC20 tokens and transfers half of them to a destination address.
    /// @param amount The amount of tokens to transfer.
    /// @param token The address of the ERC20 token.
    /// @param destination The address to send the tokens to.
    /// @dev Emits the ReceivedERC20 event. Reverts if the amount to send is zero.
    function receiveERC20Partial(uint256 amount, address token, address destination) external nonReentrant {
        uint256 amountToSend = amount / 2;
        if (amountToSend == 0) revert ZeroAmount();

        IERC20(token).safeTransferFrom(msg.sender, destination, amountToSend);

        emit ReceivedERC20(msg.sender, amountToSend, token, destination);
    }

    /// @notice Receives a call without any parameters.
    /// @dev Emits the ReceivedNoParams event.
    function receiveNoParams() external {
        emit ReceivedNoParams(msg.sender);
    }

    /// @notice Called when a revert occurs.
    /// @param revertContext Revert context.
    /// @dev Emits the ReceivedRevert event.
    function onRevert(RevertContext calldata revertContext) external {
        emit ReceivedRevert(msg.sender, revertContext);
    }

    function onCall(
        LegacyMessageContext calldata messageContext,
        bytes calldata message
    )
        external
        payable
        returns (bytes memory)
    {
        emit ReceivedOnCall(messageContext.sender, message);
        return "";
    }

    function onCall(
        MessageContext calldata messageContext,
        bytes calldata message
    )
        external
        payable
        returns (bytes memory)
    {
        emit ReceivedOnCallV2(messageContext.sender, messageContext.asset, messageContext.amount, message);
        return "";
    }

    /// @notice Receives ETH.
    receive() external payable { }

    /// @notice Fallback function to receive ETH.
    fallback() external payable { }
}
