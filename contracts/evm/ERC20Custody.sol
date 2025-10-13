// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { IERC20Custody } from "./interfaces/IERC20Custody.sol";
import { IGatewayEVM, MessageContext } from "./interfaces/IGatewayEVM.sol";

import { RevertContext } from "../../contracts/Revert.sol";

import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/// @title ERC20Custody
/// @notice Holds the ERC20 tokens deposited on ZetaChain and includes functionality to call a contract.
/// @dev This contract does not call smart contracts directly, it passes through the Gateway contract.
contract ERC20Custody is
    Initializable,
    UUPSUpgradeable,
    IERC20Custody,
    ReentrancyGuardUpgradeable,
    AccessControlUpgradeable,
    PausableUpgradeable
{
    using SafeERC20 for IERC20;

    /// @notice Gateway contract.
    IGatewayEVM public gateway;
    /// @notice Mapping of whitelisted tokens => true/false.
    mapping(address => bool) public whitelisted;
    /// @notice The address of the TSS (Threshold Signature Scheme) contract.
    address public tssAddress;
    /// @notice Used to flag if contract supports legacy methods (eg. deposit).
    bool public supportsLegacy;
    /// @notice New role identifier for pauser role.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    /// @notice New role identifier for withdrawer role.
    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
    /// @notice New role identifier for whitelister role.
    bytes32 public constant WHITELISTER_ROLE = keccak256("WHITELISTER_ROLE");

    /// @notice Initializer for ERC20Custody.
    /// @dev Set admin as default admin and pauser, and tssAddress as tss role.
    function initialize(address gateway_, address tssAddress_, address admin_) public initializer {
        if (gateway_ == address(0) || tssAddress_ == address(0) || admin_ == address(0)) {
            revert ZeroAddress();
        }

        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();
        __AccessControl_init();
        __Pausable_init();

        gateway = IGatewayEVM(gateway_);
        tssAddress = tssAddress_;
        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(PAUSER_ROLE, admin_);
        _grantRole(PAUSER_ROLE, tssAddress_);
        _grantRole(WITHDRAWER_ROLE, tssAddress_);
        _grantRole(WHITELISTER_ROLE, admin_);
        _grantRole(WHITELISTER_ROLE, tssAddress_);
    }

    /// @dev Authorizes the upgrade of the contract, sender must be owner.
    /// @param newImplementation Address of the new implementation.
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) { }

    /// @notice Pause contract.
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /// @notice Unpause contract.
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /// @notice Update tss address
    /// @param newTSSAddress new tss address
    function updateTSSAddress(address newTSSAddress) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (newTSSAddress == address(0)) revert ZeroAddress();

        _revokeRole(WITHDRAWER_ROLE, tssAddress);
        _revokeRole(WHITELISTER_ROLE, tssAddress);

        _grantRole(WITHDRAWER_ROLE, newTSSAddress);
        _grantRole(WHITELISTER_ROLE, newTSSAddress);

        emit UpdatedCustodyTSSAddress(tssAddress, newTSSAddress);

        tssAddress = newTSSAddress;
    }

    /// @notice Unpause contract.
    function setSupportsLegacy(bool _supportsLegacy) external onlyRole(DEFAULT_ADMIN_ROLE) {
        supportsLegacy = _supportsLegacy;
    }

    /// @notice Whitelist ERC20 token.
    /// @param token address of ERC20 token
    function whitelist(address token) external onlyRole(WHITELISTER_ROLE) {
        if (token == address(0)) revert ZeroAddress();
        whitelisted[token] = true;
        emit Whitelisted(token);
    }

    /// @notice Unwhitelist ERC20 token.
    /// @param token address of ERC20 token
    function unwhitelist(address token) external onlyRole(WHITELISTER_ROLE) {
        if (token == address(0)) revert ZeroAddress();
        whitelisted[token] = false;
        emit Unwhitelisted(token);
    }

    /// @notice Withdraw directly transfers the tokens to the destination address without contract call.
    /// @dev This function can only be called by the TSS address.
    /// @param to Destination address for the tokens.
    /// @param token Address of the ERC20 token.
    /// @param amount Amount of tokens to withdraw.
    function withdraw(
        address to,
        address token,
        uint256 amount
    )
        external
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        if (!whitelisted[token]) revert NotWhitelisted();

        IERC20(token).safeTransfer(to, amount);

        emit Withdrawn(to, token, amount);
    }

    /// @notice WithdrawAndCall transfers tokens to Gateway and call a contract through the Gateway.
    /// @dev This function can only be called by the TSS address.
    /// @param messageContext Message context containing sender.
    /// @param to Address of the contract to call.
    /// @param token Address of the ERC20 token.
    /// @param amount Amount of tokens to withdraw.
    /// @param data Calldata to pass to the contract call.
    function withdrawAndCall(
        MessageContext calldata messageContext,
        address to,
        address token,
        uint256 amount,
        bytes calldata data
    )
        public
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        if (!whitelisted[token]) revert NotWhitelisted();

        // Transfer the tokens to the Gateway contract
        IERC20(token).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.executeWithERC20(messageContext, token, to, amount, data);

        emit WithdrawnAndCalled(to, token, amount, data);
    }

    /// @notice WithdrawAndRevert transfers tokens to Gateway and call a contract with a revert functionality through
    /// the Gateway.
    /// @dev This function can only be called by the TSS address.
    /// @param to Address of the contract to call.
    /// @param token Address of the ERC20 token.
    /// @param amount Amount of tokens to withdraw.
    /// @param data Calldata to pass to the contract call.
    /// @param revertContext Revert context to pass to onRevert.
    function withdrawAndRevert(
        address to,
        address token,
        uint256 amount,
        bytes calldata data,
        RevertContext calldata revertContext
    )
        public
        nonReentrant
        onlyRole(WITHDRAWER_ROLE)
        whenNotPaused
    {
        if (!whitelisted[token]) revert NotWhitelisted();

        // Transfer the tokens to the Gateway contract
        IERC20(token).safeTransfer(address(gateway), amount);

        // Forward the call to the Gateway contract
        gateway.revertWithERC20(token, to, amount, data, revertContext);

        emit WithdrawnAndReverted(to, token, amount, data, revertContext);
    }

    /// @notice Deposits asset to custody and pay fee in zeta erc20.
    /// @custom:deprecated This method is deprecated.
    function deposit(
        bytes calldata recipient,
        IERC20 asset,
        uint256 amount,
        bytes calldata message
    )
        external
        nonReentrant
        whenNotPaused
    {
        if (!supportsLegacy) revert LegacyMethodsNotSupported();
        if (!whitelisted[address(asset)]) revert NotWhitelisted();
        uint256 oldBalance = asset.balanceOf(address(this));
        asset.safeTransferFrom(msg.sender, address(this), amount);
        // In case if there is a fee on a token transfer, we might not receive a full expected amount
        // and we need to correctly process that, we subtract an old balance from a new balance, which should be an
        // actual received amount.
        emit Deposited(recipient, asset, asset.balanceOf(address(this)) - oldBalance, message);
    }
}
