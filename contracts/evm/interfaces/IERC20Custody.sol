// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { RevertContext } from "../../../contracts/Revert.sol";

import { MessageContextV2 } from "./IGatewayEVM.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/// @title IERC20CustodyEvents
/// @notice Interface for the events emitted by the ERC20 custody contract.
interface IERC20CustodyEvents {
    /// @notice Emitted when tokens are withdrawn.
    /// @param token The address of the ERC20 token.
    /// @param to The address receiving the tokens.
    /// @param amount The amount of tokens withdrawn.
    event Withdrawn(address indexed to, address indexed token, uint256 amount);

    /// @notice Emitted when tokens are withdrawn and a contract call is made.
    /// @param to The address receiving the tokens.
    /// @param token The address of the ERC20 token.
    /// @param amount The amount of tokens withdrawn.
    /// @param data The calldata passed to the contract call.
    event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data);

    /// @notice Emitted when tokens are withdrawn and a revertable contract call is made.
    /// @param to The address receiving the tokens.
    /// @param token The address of the ERC20 token.
    /// @param amount The amount of tokens withdrawn.
    /// @param data The calldata passed to the contract call.
    /// @param revertContext Revert context to pass to onRevert.
    event WithdrawnAndReverted(
        address indexed to, address indexed token, uint256 amount, bytes data, RevertContext revertContext
    );

    /// @notice Emitted when ERC20 token is whitelisted
    /// @param token address of ERC20 token.
    event Whitelisted(address indexed token);

    /// @notice Emitted when ERC20 token is unwhitelisted
    /// @param token address of ERC20 token.
    event Unwhitelisted(address indexed token);

    /// @notice Emitted in legacy deposit method.
    event Deposited(bytes recipient, IERC20 indexed asset, uint256 amount, bytes message);

    /// @notice Emitted when tss address is updated
    /// @param oldTSSAddress old tss address
    /// @param newTSSAddress new tss address
    event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress);
}

/// @title IERC20CustodyErrors
/// @notice Interface for the errors used in the ERC20 custody contract.
interface IERC20CustodyErrors {
    /// @notice Error for zero address input.
    error ZeroAddress();
    /// @notice Error for not whitelisted ERC20 token
    error NotWhitelisted();
    /// @notice Error for calling not supported legacy methods.
    error LegacyMethodsNotSupported();
}

interface IERC20Custody is IERC20CustodyEvents, IERC20CustodyErrors {
    /// @notice Mapping of whitelisted tokens => true/false.
    function whitelisted(address token) external view returns (bool);

    /// @notice Withdraw directly transfers the tokens to the destination address without contract call.
    /// @dev This function can only be called by the TSS address.
    /// @param token Address of the ERC20 token.
    /// @param to Destination address for the tokens.
    /// @param amount Amount of tokens to withdraw.
    function withdraw(address token, address to, uint256 amount) external;

    /// @notice WithdrawAndCall transfers tokens to Gateway and call a contract through the Gateway.
    /// @dev This function can only be called by the TSS address.
    /// @param messageContext Message context containing sender.
    /// @param token Address of the ERC20 token.
    /// @param to Address of the contract to call.
    /// @param amount Amount of tokens to withdraw.
    /// @param data Calldata to pass to the contract call.
    function withdrawAndCall(
        MessageContextV2 calldata messageContext,
        address token,
        address to,
        uint256 amount,
        bytes calldata data
    )
        external;

    /// @notice WithdrawAndRevert transfers tokens to Gateway and call a contract with a revert functionality through
    /// the Gateway.
    /// @dev This function can only be called by the TSS address.
    /// @param token Address of the ERC20 token.
    /// @param to Address of the contract to call.
    /// @param amount Amount of tokens to withdraw.
    /// @param data Calldata to pass to the contract call.
    /// @param revertContext Revert context to pass to onRevert.
    function withdrawAndRevert(
        address token,
        address to,
        uint256 amount,
        bytes calldata data,
        RevertContext calldata revertContext
    )
        external;
}
