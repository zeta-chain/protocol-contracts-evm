// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { IGatewayZEVM } from "../../contracts/zevm/interfaces/IGatewayZEVM.sol";
import { RevertOptions } from "../../contracts/zevm/interfaces/IGatewayZEVM.sol";
import "../../contracts/zevm/interfaces/IZRC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/// @title SenderZEVM
/// @notice This contract is used just for testing.
/// @dev Provides a function to withdraw to a receiver on EVM.
contract SenderZEVM {
    /// @notice The address of the gateway contract.
    address public gateway;

    /// @notice Error indicating that the approval of tokens failed.
    error ApprovalFailed();

    constructor(address gateway_) {
        gateway = gateway_;
    }

    /// @notice Withdraw to a receiver on EVM.
    /// @param receiver The address of the receiver on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    /// @dev Approves the gateway to withdraw tokens.
    function withdrawReceiver(bytes memory receiver, uint256 amount, address zrc20) external {
        // Approve gateway to withdraw
        if (!IZRC20(zrc20).approve(gateway, amount + 100_000)) revert ApprovalFailed();

        RevertOptions memory revertOptions = RevertOptions({
            revertAddress: address(0x321),
            callOnRevert: true,
            abortAddress: address(0x321),
            revertMessage: "",
            onRevertGasLimit: 0
        });

        IGatewayZEVM(gateway).withdraw(receiver, amount, zrc20, revertOptions);
    }
}
