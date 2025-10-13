// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "../../../contracts/Revert.sol";

/// @title IGatewayEVMEvents
/// @notice Interface for the events emitted by the GatewayEVM contract.
interface IGatewayEVMEvents {
    /// @notice Emitted when a contract call is executed.
    /// @param destination The address of the contract called.
    /// @param value The amount of ETH sent with the call.
    /// @param data The calldata passed to the contract call.
    event Executed(address indexed destination, uint256 value, bytes data);

    /// @notice Emitted when a contract call is reverted.
    /// @param to The address of the contract called.
    /// @param token The address of the ERC20 token, empty if gas token
    /// @param amount The amount of ETH sent with the call.
    /// @param data The calldata passed to the contract call.
    /// @param revertContext Revert context to pass to onRevert.
    event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, RevertContext revertContext);

    /// @notice Emitted when a contract call with ERC20 tokens is executed.
    /// @param token The address of the ERC20 token.
    /// @param to The address of the contract called.
    /// @param amount The amount of tokens transferred.
    /// @param data The calldata passed to the contract call.
    event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data);

    /// @notice Emitted when a deposit is made.
    /// @param sender The address of the sender.
    /// @param receiver The address of the receiver.
    /// @param amount The amount of ETH or tokens deposited.
    /// @param asset The address of the ERC20 token (zero address if ETH).
    /// @param payload The calldata passed with the deposit. No longer used. Kept to maintain compatibility.
    /// @param revertOptions Revert options.
    event Deposited(
        address indexed sender,
        address indexed receiver,
        uint256 amount,
        address asset,
        bytes payload,
        RevertOptions revertOptions
    );

    /// @notice Emitted when a deposit and call is made.
    /// @param sender The address of the sender.
    /// @param receiver The address of the receiver.
    /// @param amount The amount of ETH or tokens deposited.
    /// @param asset The address of the ERC20 token (zero address if ETH).
    /// @param payload The calldata passed with the deposit.
    /// @param revertOptions Revert options.
    event DepositedAndCalled(
        address indexed sender,
        address indexed receiver,
        uint256 amount,
        address asset,
        bytes payload,
        RevertOptions revertOptions
    );

    /// @notice Emitted when an omnichain smart contract call is made without asset transfer.
    /// @param sender The address of the sender.
    /// @param receiver The address of the receiver.
    /// @param payload The calldata passed to the call.
    /// @param revertOptions Revert options.
    event Called(address indexed sender, address indexed receiver, bytes payload, RevertOptions revertOptions);

    /// @notice Emitted when tss address is updated
    /// @param oldTSSAddress old tss address
    /// @param newTSSAddress new tss address
    event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress);
}

/// @title IGatewayEVMErrors
/// @notice Interface for the errors used in the GatewayEVM contract.
interface IGatewayEVMErrors {
    /// @notice Error for failed execution.
    error ExecutionFailed();

    /// @notice Error for failed deposit.
    error DepositFailed();

    /// @notice Error for insufficient ETH amount.
    error InsufficientETHAmount();

    /// @notice Error for insufficient ERC20 token amount.
    error InsufficientERC20Amount();

    /// @notice Error for zero address input.
    error ZeroAddress();

    /// @notice Error for failed token approval.
    error ApprovalFailed();

    /// @notice Error for already initialized custody.
    error CustodyInitialized();

    /// @notice Error for already initialized connector.
    error ConnectorInitialized();

    /// @notice Error when trying to transfer not whitelisted token to custody.
    error NotWhitelistedInCustody();

    /// @notice Error when trying to call onCall method using arbitrary call.
    error NotAllowedToCallOnCall();

    /// @notice Error when trying to call onRevert method using arbitrary call.
    error NotAllowedToCallOnRevert();

    /// @notice Error indicating payload size exceeded in external functions.
    error PayloadSizeExceeded();
}

/// @title IGatewayEVM
/// @notice Interface for the GatewayEVM contract.
interface IGatewayEVM is IGatewayEVMErrors, IGatewayEVMEvents {
    /// @notice Executes a call to a contract using ERC20 tokens.
    /// @param messageContext Message context containing sender and arbitrary call flag.
    /// @param token The address of the ERC20 token.
    /// @param to The address of the contract to call.
    /// @param amount The amount of tokens to transfer.
    /// @param data The calldata to pass to the contract call.
    function executeWithERC20(
        MessageContext calldata messageContext,
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    )
        external;

    /// @notice Transfers msg.value to destination contract and executes it's onRevert function.
    /// @dev This function can only be called by the TSS address and it is payable.
    /// @param destination Address to call.
    /// @param data Calldata to pass to the call.
    /// @param revertContext Revert context to pass to onRevert.
    function executeRevert(
        address destination,
        bytes calldata data,
        RevertContext calldata revertContext
    )
        external
        payable;

    /// @notice Executes a call to a destination address without ERC20 tokens.
    /// @dev This function can only be called by the TSS address and it is payable.
    /// @param messageContext Message context containing sender and arbitrary call flag.
    /// @param destination Address to call.
    /// @param data Calldata to pass to the call.
    /// @return The result of the call.
    function execute(
        MessageContext calldata messageContext,
        address destination,
        bytes calldata data
    )
        external
        payable
        returns (bytes memory);

    /// @notice Executes a revertable call to a contract using ERC20 tokens.
    /// @param token The address of the ERC20 token.
    /// @param to The address of the contract to call.
    /// @param amount The amount of tokens to transfer.
    /// @param data The calldata to pass to the contract call.
    /// @param revertContext Revert context to pass to onRevert.
    function revertWithERC20(
        address token,
        address to,
        uint256 amount,
        bytes calldata data,
        RevertContext calldata revertContext
    )
        external;

    /// @notice Deposits ETH to the TSS address.
    /// @param receiver Address of the receiver.
    /// @param revertOptions Revert options.
    function deposit(address receiver, RevertOptions calldata revertOptions) external payable;

    /// @notice Deposits ERC20 tokens to the custody or connector contract.
    /// @param receiver Address of the receiver.
    /// @param amount Amount of tokens to deposit.
    /// @param asset Address of the ERC20 token.
    /// @param revertOptions Revert options.
    function deposit(address receiver, uint256 amount, address asset, RevertOptions calldata revertOptions) external;

    /// @notice Deposits ETH to the TSS address and calls an omnichain smart contract.
    /// @param receiver Address of the receiver.
    /// @param payload Calldata to pass to the call.
    /// @param revertOptions Revert options.
    function depositAndCall(
        address receiver,
        bytes calldata payload,
        RevertOptions calldata revertOptions
    )
        external
        payable;

    /// @notice Deposits ERC20 tokens to the custody or connector contract and calls an omnichain smart contract.
    /// @param receiver Address of the receiver.
    /// @param amount Amount of tokens to deposit.
    /// @param asset Address of the ERC20 token.
    /// @param payload Calldata to pass to the call.
    /// @param revertOptions Revert options.
    function depositAndCall(
        address receiver,
        uint256 amount,
        address asset,
        bytes calldata payload,
        RevertOptions calldata revertOptions
    )
        external;

    /// @notice Calls an omnichain smart contract without asset transfer.
    /// @param receiver Address of the receiver.
    /// @param payload Calldata to pass to the call.
    /// @param revertOptions Revert options.
    function call(address receiver, bytes calldata payload, RevertOptions calldata revertOptions) external;
}

/// @notice Message context passed to execute function.
/// @param sender Sender from omnichain contract.
struct LegacyMessageContext {
    address sender;
}

/// @notice Interface implemented by contracts receiving authenticated calls.
interface Callable {
    function onCall(LegacyMessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
}

/// @notice Message context passed to execute function.
/// @param sender Sender from omnichain contract.
/// @param asset The address of the asset.
/// @param amount The amount of the asset.
struct MessageContext {
    address sender;
    address asset;
    uint256 amount;
}

/// @notice Interface implemented by contracts receiving authenticated calls with new MessageContext.
interface CallableV2 {
    function onCall(
        MessageContext calldata context,
        bytes calldata message
    )
        external
        payable
        returns (bytes memory);
}
