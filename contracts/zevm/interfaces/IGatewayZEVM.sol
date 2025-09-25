// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "../../../contracts/Revert.sol";
import "./UniversalContract.sol";

/// @title IGatewayZEVMEvents
/// @notice Interface for the events emitted by the GatewayZEVM contract.
interface IGatewayZEVMEvents {
    /// @notice Emitted when a cross-chain call is made.
    /// @param sender The address of the sender.
    /// @param zrc20 Address of zrc20 to pay fees.
    /// @param receiver The receiver address on the external chain.
    /// @param message The calldata passed to the contract call.
    /// @param callOptions Call options including gas limit and arbirtrary call flag.
    /// @param revertOptions Revert options.
    event Called(
        address indexed sender,
        address indexed zrc20,
        bytes receiver,
        bytes message,
        CallOptions callOptions,
        RevertOptions revertOptions
    );

    /// @notice Emitted when a withdrawal is made.
    /// @param sender The address from which the tokens are withdrawn.
    /// @param chainId Chain id of external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param value The amount of tokens withdrawn.
    /// @param gasfee The gas fee for the withdrawal.
    /// @param protocolFlatFee The protocol flat fee for the withdrawal.
    /// @param message The calldata passed with the withdraw. No longer used. Kept to maintain compatibility.
    /// @param callOptions Call options including gas limit and arbirtrary call flag.
    /// @param revertOptions Revert options.
    event Withdrawn(
        address indexed sender,
        uint256 indexed chainId,
        bytes receiver,
        address zrc20,
        uint256 value,
        uint256 gasfee,
        uint256 protocolFlatFee,
        bytes message,
        CallOptions callOptions,
        RevertOptions revertOptions
    );

    /// @notice Emitted when a withdraw and call is made.
    /// @param sender The address from which the tokens are withdrawn.
    /// @param chainId Chain id of external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param value The amount of tokens withdrawn.
    /// @param gasfee The gas fee for the withdrawal.
    /// @param protocolFlatFee The protocol flat fee for the withdrawal.
    /// @param message The calldata passed to the contract call.
    /// @param callOptions Call options including gas limit and arbirtrary call flag.
    /// @param revertOptions Revert options.
    event WithdrawnAndCalled(
        address indexed sender,
        uint256 indexed chainId,
        bytes receiver,
        address zrc20,
        uint256 value,
        uint256 gasfee,
        uint256 protocolFlatFee,
        bytes message,
        CallOptions callOptions,
        RevertOptions revertOptions
    );
}

/// @title IGatewayZEVMErrors
/// @notice Interface for the errors used in the GatewayZEVM contract.
interface IGatewayZEVMErrors {
    /// @notice Error indicating a empty address was provided.
    error EmptyAddress();

    /// @notice Error indicating a withdrawal failure.
    /// @param token The address of the token that failed to withdraw.
    /// @param recipient The address that was supposed to receive the tokens.
    /// @param amount The amount of tokens that failed to withdraw.
    error WithdrawalFailed(address token, address recipient, uint256 amount);

    /// @notice Error indicating an insufficient token amount.
    error InsufficientAmount();

    /// @notice Error indicating a failure to burn ZRC20 tokens.
    /// @param zrc20 The address of the ZRC20 token that failed to burn.
    /// @param amount The amount of tokens that failed to burn.
    error ZRC20BurnFailed(address zrc20, uint256 amount);

    /// @notice Error indicating a failure to transfer ZRC20 tokens.
    /// @param zrc20 The address of the ZRC20 token that failed to transfer.
    /// @param from The address sending the tokens.
    /// @param to The address receiving the tokens.
    /// @param amount The amount of tokens that failed to transfer.
    error ZRC20TransferFailed(address zrc20, address from, address to, uint256 amount);

    /// @notice Error indicating a failure to deposit ZRC20 tokens.
    /// @param zrc20 The address of the ZRC20 token that failed to deposit.
    /// @param to The address that was supposed to receive the deposit.
    /// @param amount The amount of tokens that failed to deposit.
    error ZRC20DepositFailed(address zrc20, address to, uint256 amount);

    /// @notice Error indicating a failure to transfer gas fee.
    /// @param token The address of the token used for gas fee.
    /// @param to The address that was supposed to receive the gas fee.
    /// @param amount The amount of gas fee that failed to transfer.
    error GasFeeTransferFailed(address token, address to, uint256 amount);

    /// @notice Error indicating that the caller is not the protocol account.
    error CallerIsNotProtocol();

    /// @notice Error indicating an invalid target address.
    error InvalidTarget();

    /// @notice Error indicating a failure to send ZETA tokens.
    /// @param recipient The address that was supposed to receive the ZETA tokens.
    /// @param amount The amount of ZETA tokens that failed to send.
    error FailedZetaSent(address recipient, uint256 amount);

    /// @notice Error indicating that only WZETA or the protocol address can call the function.
    error OnlyWZETAOrProtocol();

    /// @notice Error indicating an invalid gas limit.
    error InvalidGasLimit();

    /// @notice Error indicating message size exceeded in external functions.
    /// @param provided The size of the message that was provided.
    /// @param maximum The maximum allowed message size.
    error MessageSizeExceeded(uint256 provided, uint256 maximum);

    /// @notice Error indicating an invalid gas price.
    error ZeroGasPrice();
}

/// @title IGatewayZEVM
/// @notice Interface for the GatewayZEVM contract.
/// @dev Defines functions for cross-chain interactions and token handling.
interface IGatewayZEVM is IGatewayZEVMErrors, IGatewayZEVMEvents {
    /// @notice Withdraw ZRC20 tokens to an external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param revertOptions Revert options.
    function withdraw(
        bytes memory receiver,
        uint256 amount,
        address zrc20,
        RevertOptions calldata revertOptions
    )
        external;

    /// @notice Withdraw ZETA tokens to an external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param revertOptions Revert options.
    function withdraw(bytes memory receiver, uint256 chainId, RevertOptions calldata revertOptions) external payable;

    /// @notice Withdraw ZRC20 tokens and call a smart contract on an external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param message The calldata to pass to the contract call.
    /// @param callOptions Call options including gas limit and arbirtrary call flag.
    /// @param revertOptions Revert options.
    function withdrawAndCall(
        bytes memory receiver,
        uint256 amount,
        address zrc20,
        bytes calldata message,
        CallOptions calldata callOptions,
        RevertOptions calldata revertOptions
    )
        external;

    /// @notice Withdraw ZETA tokens and call a smart contract on an external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param chainId Chain id of the external chain.
    /// @param message The calldata to pass to the contract call.
    /// @param callOptions Call options including gas limit and arbirtrary call flag.
    /// @param revertOptions Revert options.
    function withdrawAndCall(
        bytes memory receiver,
        uint256 chainId,
        bytes calldata message,
        CallOptions calldata callOptions,
        RevertOptions calldata revertOptions
    )
        external
        payable;

    /// @notice Call a smart contract on an external chain without asset transfer.
    /// @param receiver The receiver address on the external chain.
    /// @param zrc20 Address of zrc20 to pay fees.
    /// @param message The calldata to pass to the contract call.
    /// @param callOptions Call options including gas limit and arbirtrary call flag.
    /// @param revertOptions Revert options.
    function call(
        bytes memory receiver,
        address zrc20,
        bytes calldata message,
        CallOptions calldata callOptions,
        RevertOptions calldata revertOptions
    )
        external;

    /// @notice Deposit foreign coins into ZRC20.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param amount The amount of tokens to deposit.
    /// @param target The target address to receive the deposited tokens.
    function deposit(address zrc20, uint256 amount, address target) external;

    /// @notice Deposit native ZETA.
    /// @param target The target address to receive the ZETA.
    function deposit(address target) external payable;

    /// @notice Execute a user-specified contract on ZEVM.
    /// @param context The context of the cross-chain call.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param amount The amount of tokens to transfer.
    /// @param target The target contract to call.
    /// @param message The calldata to pass to the contract call.
    function execute(
        MessageContext calldata context,
        address zrc20,
        uint256 amount,
        address target,
        bytes calldata message
    )
        external;

    /// @notice Deposit foreign coins into ZRC20 and call a user-specified contract on ZEVM.
    /// @param context The context of the cross-chain call.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param amount The amount of tokens to transfer.
    /// @param target The target contract to call.
    /// @param message The calldata to pass to the contract call.
    function depositAndCall(
        MessageContext calldata context,
        address zrc20,
        uint256 amount,
        address target,
        bytes calldata message
    )
        external;

    /// @notice Deposit native ZETA and call a user-specified contract on ZEVM.
    /// @param context The context of the cross-chain call.
    /// @param target The target contract to call.
    /// @param message The calldata to pass to the contract call.
    function depositAndCall(MessageContext calldata context, address target, bytes calldata message) external payable;

    /// @notice Revert a user-specified contract on ZEVM.
    /// @param target The target contract to call.
    /// @param revertContext Revert context to pass to onRevert.
    function executeRevert(address target, RevertContext calldata revertContext) external;

    /// @notice Deposit foreign coins into ZRC20 and revert a user-specified contract on ZEVM.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param amount The amount of tokens to revert.
    /// @param target The target contract to call.
    /// @param revertContext Revert context to pass to onRevert.
    function depositAndRevert(
        address zrc20,
        uint256 amount,
        address target,
        RevertContext calldata revertContext
    )
        external;
}

/// @notice CallOptions struct passed to call and withdrawAndCall functions.
/// @param gasLimit Gas limit.
/// @param isArbitraryCall Indicates if call should be arbitrary or authenticated.
struct CallOptions {
    uint256 gasLimit;
    bool isArbitraryCall;
}
