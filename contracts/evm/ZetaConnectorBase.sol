// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import { RevertContext } from "../../contracts/Revert.sol";
import {
    IGatewayEVM,
    IGatewayEVMErrors,
    IGatewayEVMEvents,
    MessageContext
} from "../../contracts/evm/interfaces/IGatewayEVM.sol";
import "../../contracts/evm/interfaces/IZetaConnector.sol";

/// @title ZetaConnectorBase
/// @notice Abstract base contract for ZetaConnector.
/// @dev This contract implements basic functionality for handling tokens and interacting with the Gateway contract.
abstract contract ZetaConnectorBase is
    Initializable,
    UUPSUpgradeable,
    IZetaConnectorEvents,
    ReentrancyGuardUpgradeable,
    PausableUpgradeable,
    AccessControlUpgradeable
{
    using SafeERC20 for IERC20;

    /// @notice Error indicating that a zero address was provided.
    error ZeroAddress();

    /// @notice The Gateway contract used for executing cross-chain calls.
    IGatewayEVM public gateway;
    /// @notice The address of the Zeta token.
    address public zetaToken;
    /// @notice The address of the TSS (Threshold Signature Scheme) contract.
    address public tssAddress;

    /// @notice New role identifier for withdrawer role.
    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
    /// @notice New role identifier for pauser role.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    /// @notice New role identifier for tss role.
    bytes32 public constant TSS_ROLE = keccak256("TSS_ROLE");

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /// @notice Initializer for ZetaConnectors.
    /// @dev Set admin as default admin and pauser, and tssAddress as tss role.
    function initialize(address gateway_, address zetaToken_, address tssAddress_, address admin_)
        public
        virtual
        onlyInitializing
    {
        if (gateway_ == address(0) || zetaToken_ == address(0) || tssAddress_ == address(0) || admin_ == address(0)) {
            revert ZeroAddress();
        }

        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();
        __AccessControl_init();
        __Pausable_init();

        gateway = IGatewayEVM(gateway_);
        zetaToken = zetaToken_;
        tssAddress = tssAddress_;

        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(WITHDRAWER_ROLE, tssAddress_);
        _grantRole(TSS_ROLE, tssAddress_);
        _grantRole(PAUSER_ROLE, admin_);
    }

    /// @dev Authorizes the upgrade of the contract, sender must be owner.
    /// @param newImplementation Address of the new implementation.
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) { }

    /// @notice Update tss address
    /// @param newTSSAddress new tss address
    function updateTSSAddress(address newTSSAddress) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (newTSSAddress == address(0)) revert ZeroAddress();

        _revokeRole(WITHDRAWER_ROLE, tssAddress);
        _revokeRole(TSS_ROLE, tssAddress);

        _grantRole(WITHDRAWER_ROLE, newTSSAddress);
        _grantRole(TSS_ROLE, newTSSAddress);

        emit UpdatedZetaConnectorTSSAddress(tssAddress, newTSSAddress);

        tssAddress = newTSSAddress;
    }

    /// @notice Pause contract.
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /// @notice Unpause contract.
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /// @notice Handle received tokens.
    /// @param amount The amount of tokens received.
    function deposit(uint256 amount) external virtual;
}
