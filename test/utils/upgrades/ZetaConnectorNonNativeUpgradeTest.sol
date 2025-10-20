// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "../../../contracts/evm/ZetaConnectorBase.sol";
import "../../../contracts/evm/interfaces/IZetaNonEthNew.sol";

/// @title ZetaConnectorNonNativeUpgradeTest
/// @notice Modified ZetaConnectorNonNative contract for testing upgrades
/// @dev The only difference is in event naming
/// @custom:oz-upgrades-from ZetaConnectorNonNative
contract ZetaConnectorNonNativeUpgradeTest is ZetaConnectorBase {
    /// @notice Event triggered when max supply is updated.
    /// @param maxSupply New max supply.
    event MaxSupplyUpdated(uint256 maxSupply);

    error ExceedsMaxSupply();

    /// @notice Max supply for minting.
    uint256 public maxSupply;

    /// @dev Modified event for testing upgrade.
    event WithdrawnV2(address indexed to, uint256 amount);

    function initialize(address gateway_, address zetaToken_, address tssAddress_, address admin_)
        public
        override
        initializer
    {
        super.initialize(gateway_, zetaToken_, tssAddress_, admin_);

        maxSupply = type(uint256).max;
    }

    /// @notice Set max supply for minting.
    /// @param maxSupply_ New max supply.
    /// @dev This function can only be called by the TSS address.
    function setMaxSupply(uint256 maxSupply_) external onlyRole(TSS_ROLE) whenNotPaused {
        maxSupply = maxSupply_;
        emit MaxSupplyUpdated(maxSupply_);
    }

    /// @notice Withdraw tokens to a specified address.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param internalSendHash A hash used for internal tracking of the transaction.
    /// @dev This function can only be called by the TSS address.
    function withdraw(address to, uint256 amount, bytes32 internalSendHash)
        external
        override
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        _mintTo(to, amount, internalSendHash);
        emit WithdrawnV2(to, amount);
    }

    /// @notice Withdraw tokens and call a contract through Gateway.
    /// @param messageContext Message context containing sender.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param data The calldata to pass to the contract call.
    /// @param internalSendHash A hash used for internal tracking of the transaction.
    /// @dev This function can only be called by the TSS address, and mints if supply is not reached.
    function withdrawAndCall(
        MessageContext calldata messageContext,
        address to,
        uint256 amount,
        bytes calldata data,
        bytes32 internalSendHash
    )
        external
        override
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        // Mint zetaToken to the Gateway contract
        _mintTo(address(gateway), amount, internalSendHash);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(messageContext, address(zetaToken), to, amount, data);

        emit WithdrawnAndCalled(to, amount, data);
    }

    /// @notice Withdraw tokens and call a contract with a revert callback through Gateway.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param data The calldata to pass to the contract call.
    /// @param internalSendHash A hash used for internal tracking of the transaction.
    /// @dev This function can only be called by the TSS address, and mints if supply is not reached.
    /// @param revertContext Revert context to pass to onRevert.
    function withdrawAndRevert(
        address to,
        uint256 amount,
        bytes calldata data,
        bytes32 internalSendHash,
        RevertContext calldata revertContext
    )
        external
        override
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        // Mint zetaToken to the Gateway contract
        _mintTo(address(gateway), amount, internalSendHash);

        // Forward the call to the Gateway contract
        gateway.revertWithERC20(address(zetaToken), to, amount, data, revertContext);

        emit WithdrawnAndReverted(to, amount, data, revertContext);
    }

    /// @notice Handle received tokens and burn them.
    /// @param amount The amount of tokens received.
    function receiveTokens(uint256 amount) external override whenNotPaused {
        IZetaNonEthNew(zetaToken).burnFrom(msg.sender, amount);
    }

    /// @dev mints to provided account and checks if totalSupply will be exceeded
    function _mintTo(address to, uint256 amount, bytes32 internalSendHash) private {
        if (amount + IERC20(zetaToken).totalSupply() > maxSupply) revert ExceedsMaxSupply();
        IZetaNonEthNew(zetaToken).mint(address(to), amount, internalSendHash);
    }
}
