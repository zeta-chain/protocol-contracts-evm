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

    /// @notice Emitted when additional action fee for multi-deposits is updated
    /// @param oldFeeWei old fee in wei
    /// @param newFeeWei new fee in wei
    event UpdatedAdditionalActionFee(uint256 oldFeeWei, uint256 newFeeWei);
}

/// @title IGatewayEVMErrors
/// @notice Interface for the errors used in the GatewayEVM contract.
interface IGatewayEVMErrors {
    /// @notice Error for failed execution.
    error ExecutionFailed();

    /// @notice Error for failed deposit.
    error DepositFailed();

    /// @notice Error for insufficient token amount.
    error InsufficientEVMAmount();

    /// @notice Error for zero address input.
    error ZeroAddress();

    /// @notice Error for failed token approval.
    /// @param token The address of the token for which approval failed.
    /// @param spender The address that was supposed to be approved to spend the tokens.
    error ApprovalFailed(address token, address spender);

    /// @notice Error for already initialized custody.
    error CustodyInitialized();

    /// @notice Error for already initialized connector.
    error ConnectorInitialized();

    /// @notice Error when trying to transfer not whitelisted token to custody.
    /// @param token The address of the token that is not whitelisted in custody.
    error NotWhitelistedInCustody(address token);

    /// @notice Error when trying to call onCall method using arbitrary call.
    error NotAllowedToCallOnCall();

    /// @notice Error when trying to call onRevert method using arbitrary call.
    error NotAllowedToCallOnRevert();

    /// @notice Error indicating payload size exceeded in external functions.
    /// @param provided The size of the payload that was provided.
    /// @param maximum The maximum allowed payload size.
    error PayloadSizeExceeded(uint256 provided, uint256 maximum);

    /// @notice Error thrown when fee transfer to TSS address fails.
    /// @dev This error occurs when the low-level call to transfer fees fails.
    error FeeTransferFailed();

    /// @notice Error thrown when insufficient fee is provided for additional actions.
    /// @param required The fee amount required for the action.
    /// @param provided The fee amount actually provided by the caller.
    error InsufficientFee(uint256 required, uint256 provided);

    /// @notice Error thrown when excess ETH is sent for non-ETH operations.
    /// @param required The fee amount required for the action.
    /// @param provided The ETH amount actually provided by the caller.
    error ExcessETHProvided(uint256 required, uint256 provided);

    /// @notice Error thrown when additional action functionality is disabled (fee set to 0).
    error AdditionalActionDisabled();

    /// @notice Error thrown when msg.value doesn't match expected amount + fee.
    /// @param expected The expected value (amount + fee).
    /// @param provided The actual msg.value provided.
    error IncorrectValueProvided(uint256 expected, uint256 provided);
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

    /// @notice Deposits ETH to the TSS address with specified amount.
    /// @param receiver Address of the receiver.
    /// @param amount Amount of ETH to deposit.
    /// @param revertOptions Revert options.
    function deposit(address receiver, uint256 amount, RevertOptions calldata revertOptions) external payable;

    /// @notice Deposits ERC20 tokens to the custody or connector contract.
    /// @param receiver Address of the receiver.
    /// @param amount Amount of tokens to deposit.
    /// @param asset Address of the ERC20 token.
    /// @param revertOptions Revert options.
    function deposit(
        address receiver,
        uint256 amount,
        address asset,
        RevertOptions calldata revertOptions
    )
        external
        payable;

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

    /// @notice Deposits ETH to the TSS address and calls an omnichain smart contract with specified amount.
    /// @param receiver Address of the receiver.
    /// @param amount Amount of ETH to deposit.
    /// @param payload Calldata to pass to the call.
    /// @param revertOptions Revert options.
    function depositAndCall(
        address receiver,
        uint256 amount,
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
        external
        payable;

    /// @notice Calls an omnichain smart contract without asset transfer.
    /// @param receiver Address of the receiver.
    /// @param payload Calldata to pass to the call.
    /// @param revertOptions Revert options.
    function call(address receiver, bytes calldata payload, RevertOptions calldata revertOptions) external payable;
}

/// @notice Message context passed to execute function.
/// @param sender Sender from omnichain contract.
struct MessageContext {
    address sender;
}

/// @notice Interface implemented by contracts receiving authenticated calls.
interface Callable {
    function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
}
