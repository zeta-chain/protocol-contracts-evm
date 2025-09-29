// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

/// @notice Struct containing revert options
/// @param revertAddress Address to receive revert.
/// @param callOnRevert Flag if onRevert hook should be called.
/// @param abortAddress Address to receive funds if aborted.
/// @param revertMessage Arbitrary data sent back in onRevert.
/// @param onRevertGasLimit Gas limit for revert tx, unused on GatewayZEVM methods
struct RevertOptions {
    address revertAddress;
    bool callOnRevert;
    address abortAddress;
    bytes revertMessage;
    uint256 onRevertGasLimit;
}

/// @notice Struct containing revert context passed to onRevert.
/// @param sender Address of account that initiated smart contract call.
/// @param asset Address of asset. On a connected chain, it contains the fungible
/// token address or is empty if it's a gas token. On ZetaChain, it contains the
/// address of the ZRC20.
/// @param amount Amount specified with the transaction.
/// @param revertMessage Arbitrary data sent back in onRevert.
struct RevertContext {
    address sender;
    address asset;
    uint256 amount;
    bytes revertMessage;
}

/// @notice Struct containing abort context passed to onAbort.
/// @param sender Address of account that initiated smart contract call.
/// bytes is used as the crosschain transaction can be initiated from a non-EVM chain.
/// @param asset Address of asset. On a connected chain, it contains the fungible
/// token address or is empty if it's a gas token. On ZetaChain, it contains the
/// address of the ZRC20.
/// @param amount Amount specified with the transaction.
/// @param outgoing Flag to indicate if the crosschain transaction was outgoing: from ZetaChain to connected chain.
/// if false, the transaction was incoming: from connected chain to ZetaChain.
/// @param chainID Chain ID of the connected chain.
/// @param revertMessage Arbitrary data specified in the RevertOptions object when initating the crosschain transaction.
struct AbortContext {
    bytes sender;
    address asset;
    uint256 amount;
    bool outgoing;
    uint256 chainID;
    bytes revertMessage;
}

/// @title Revertable
/// @notice Interface for contracts that support revertable calls.
interface Revertable {
    /// @notice Called when a revertable call is made.
    /// @param revertContext Revert context to pass to onRevert.
    function onRevert(RevertContext calldata revertContext) external;
}

/// @title Abortable
/// @notice Interface for contracts that support abortable calls.
interface Abortable {
    /// @notice Called when a revertable call is aborted.
    /// @param abortContext Abort context to pass to onAbort.
    function onAbort(AbortContext calldata abortContext) external;
}
