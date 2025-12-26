// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { IGatewayZEVM } from "../../contracts/zevm/interfaces/IGatewayZEVM.sol";
import { CallOptions, RevertOptions } from "../../contracts/zevm/interfaces/IGatewayZEVM.sol";
import "../../contracts/zevm/interfaces/IZRC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/// @title SenderZEVM
/// @notice This contract is used just for testing.
/// @dev Provides functions to call a receiver on EVM and to withdraw and call a receiver on EVM.
contract SenderZEVM {
    /// @notice The address of the gateway contract.
    address public gateway;

    /// @notice Error indicating that the approval of tokens failed.
    error ApprovalFailed();

    constructor(address gateway_) {
        gateway = gateway_;
    }

    /// @notice Call a receiver on EVM.
    /// @param receiver The address of the receiver on the external chain.
    /// @param zrc20 Address of zrc20 to pay fees.
    /// @param str A string parameter to pass to the receiver's function.
    /// @param num A numeric parameter to pass to the receiver's function.
    /// @param flag A boolean parameter to pass to the receiver's function.
    /// @dev Encodes the function call and passes it to the gateway.
    function callReceiver(
        bytes memory receiver,
        address zrc20,
        string memory str,
        uint256 num,
        bool flag
    )
        external
    {
        // Encode the function call to the receiver's receivePayable method
        bytes memory message = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);

        RevertOptions memory revertOptions = RevertOptions({
            revertAddress: address(0x321),
            callOnRevert: true,
            abortAddress: address(0x321),
            revertMessage: "",
            onRevertGasLimit: 0
        });

        CallOptions memory callOptions = CallOptions({ gasLimit: 100_000, isArbitraryCall: false });

        IZRC20(zrc20).approve(gateway, 100_000);

        // Pass encoded call to gateway
        IGatewayZEVM(gateway).call(receiver, zrc20, message, callOptions, revertOptions);
    }

    /// @notice Withdraw and call a receiver on EVM.
    /// @param receiver The address of the receiver on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param str A string parameter to pass to the receiver's function.
    /// @param num A numeric parameter to pass to the receiver's function.
    /// @param flag A boolean parameter to pass to the receiver's function.
    /// @dev Approves the gateway to withdraw tokens and encodes the function call to pass to the gateway.
    function withdrawAndCallReceiver(
        bytes memory receiver,
        uint256 amount,
        address zrc20,
        string memory str,
        uint256 num,
        bool flag
    )
        external
    {
        // Encode the function call to the receiver's receivePayable method
        bytes memory message = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);

        // Approve gateway to withdraw
        if (!IZRC20(zrc20).approve(gateway, amount + 100_000)) revert ApprovalFailed();

        RevertOptions memory revertOptions = RevertOptions({
            revertAddress: address(0x321),
            callOnRevert: true,
            abortAddress: address(0x321),
            revertMessage: "",
            onRevertGasLimit: 0
        });

        CallOptions memory callOptions = CallOptions({ gasLimit: 100_000, isArbitraryCall: false });

        // Pass encoded call to gateway
        IGatewayZEVM(gateway).withdrawAndCall(receiver, amount, zrc20, message, callOptions, revertOptions);
    }
}
