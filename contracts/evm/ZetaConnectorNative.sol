// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "./ZetaConnectorBase.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/// @title ZetaConnectorNative
/// @notice Implementation of ZetaConnectorBase for native token handling.
/// @dev This contract directly transfers Zeta tokens and interacts with the Gateway contract.
contract ZetaConnectorNative is ZetaConnectorBase {
    using SafeERC20 for IERC20;

    function initialize(address gateway_, address zetaToken_, address tssAddress_, address admin_)
        public
        override
        initializer
    {
        super.initialize(gateway_, zetaToken_, tssAddress_, admin_);
    }

    /// @notice Withdraw tokens to a specified address.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    //// @param internalSendHash A hash used for internal tracking of the transaction (not used currently
    // https://github.com/zeta-chain/protocol-contracts-evm/issues/398)
    /// @dev This function can only be called by the TSS address.
    function withdraw(
        address to,
        uint256 amount,
        bytes32 /*internalSendHash*/
    )
        external
        override
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        IERC20(zetaToken).safeTransfer(to, amount);
        emit Withdrawn(to, amount);
    }

    /// @notice Withdraw tokens and call a contract through Gateway.
    /// @param messageContext Message context containing sender.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param data The calldata to pass to the contract call.
    //// @param internalSendHash A hash used for internal tracking of the transaction (not used currently
    // https://github.com/zeta-chain/protocol-contracts-evm/issues/398)
    /// @dev This function can only be called by the TSS address.
    function withdrawAndCall(
        MessageContext calldata messageContext,
        address to,
        uint256 amount,
        bytes calldata data,
        bytes32 /*internalSendHash*/
    )
        external
        override
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        // Transfer zetaToken to the Gateway contract
        IERC20(zetaToken).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(messageContext, address(zetaToken), to, amount, data);

        emit WithdrawnAndCalled(to, amount, data);
    }

    /// @notice Withdraw tokens and call a contract with a revert callback through Gateway.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param data The calldata to pass to the contract call.
    //// @param internalSendHash A hash used for internal tracking of the transaction (not used currently
    // https://github.com/zeta-chain/protocol-contracts-evm/issues/398)
    /// @dev This function can only be called by the TSS address.
    /// @param revertContext Revert context to pass to onRevert.
    function withdrawAndRevert(
        address to,
        uint256 amount,
        bytes calldata data,
        bytes32, /*internalSendHash*/
        RevertContext calldata revertContext
    )
        external
        override
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        // Transfer zetaToken to the Gateway contract
        IERC20(zetaToken).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.revertWithERC20(address(zetaToken), to, amount, data, revertContext);

        emit WithdrawnAndReverted(to, amount, data, revertContext);
    }

    /// @notice Handle received tokens.
    /// @param amount The amount of tokens received.
    function receiveTokens(uint256 amount) external override whenNotPaused {
        IERC20(zetaToken).safeTransferFrom(msg.sender, address(this), amount);
    }
}
