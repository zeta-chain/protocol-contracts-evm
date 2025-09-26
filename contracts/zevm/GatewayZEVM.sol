// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { CallOptions, IGatewayZEVM } from "./interfaces/IGatewayZEVM.sol";

import { INotSupportedMethods } from "../../contracts/Errors.sol";
import "../../contracts/Revert.sol";
import "./interfaces/IWZETA.sol";
import { IZRC20 } from "./interfaces/IZRC20.sol";
import { MessageContext, UniversalContract } from "./interfaces/UniversalContract.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import { ICoreRegistry } from "./interfaces/ICoreRegistry.sol";
import { ISystem } from "./interfaces/ISystem.sol";
import { GatewayZEVMValidations } from "./libraries/GatewayZEVMValidations.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";

/// @title GatewayZEVM
/// @notice The GatewayZEVM contract is the endpoint to call smart contracts on omnichain.
/// @dev The contract doesn't hold any funds and should never have active allowances.
contract GatewayZEVM is
    IGatewayZEVM,
    Initializable,
    AccessControlUpgradeable,
    UUPSUpgradeable,
    ReentrancyGuardUpgradeable,
    PausableUpgradeable,
    INotSupportedMethods
{
    using GatewayZEVMValidations for *;

    /// @notice The constant address of the protocol
    address public constant PROTOCOL_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;

    /// @notice New role identifier for pauser role.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    /// @notice The address of the Zeta token.
    address public zetaToken;

    /// @notice The address of the registry contract on ZetaChain
    address public registry;

    /// @dev Only protocol address allowed modifier.
    modifier onlyProtocol() {
        if (msg.sender != PROTOCOL_ADDRESS) {
            revert CallerIsNotProtocol();
        }
        _;
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /// @notice Initialize with address of zeta token and admin account set as DEFAULT_ADMIN_ROLE.
    /// @dev Using admin to authorize upgrades and pause.
    function initialize(address zetaToken_, address admin_) public initializer {
        if (zetaToken_ == address(0) || admin_ == address(0)) {
            revert GatewayZEVMValidations.EmptyAddress();
        }
        __UUPSUpgradeable_init();
        __AccessControl_init();
        __Pausable_init();
        __ReentrancyGuard_init();

        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(PAUSER_ROLE, admin_);
        zetaToken = zetaToken_;
    }

    /// @dev Authorizes the upgrade of the contract.
    /// @param newImplementation The address of the new implementation.
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) { }

    /// @dev Receive function to receive ZETA from WETH9.withdraw().
    receive() external payable whenNotPaused {
        if (msg.sender != zetaToken && msg.sender != PROTOCOL_ADDRESS) revert OnlyWZETAOrProtocol();
    }

    /// @notice Pause contract.
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /// @notice Unpause contract.
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /// @notice Set registry address, callable only by DEFAULT_ADMIN_ROLE.
    /// @param _registry The address of the registry contract.
    function setRegistryAddress(address _registry) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (_registry == address(0)) revert GatewayZEVMValidations.EmptyAddress();
        registry = _registry;
    }

    /// @notice Helper function to safely execute transferFrom
    /// @param zrc20 The ZRC20 token address
    /// @param from The sender address
    /// @param to The recipient address
    /// @param amount The amount to transfer
    /// @return True if the transfer was successful, false otherwise.
    function _safeTransferFrom(address zrc20, address from, address to, uint256 amount) private returns (bool) {
        try IZRC20(zrc20).transferFrom(from, to, amount) returns (bool success) {
            return success;
        } catch {
            return false;
        }
    }

    // @notice Helper function to safely burn ZRC20 tokens
    // @param zrc20 The ZRC20 token address
    // @param amount The amount to burn
    // @return True if the burn was successful, false otherwise
    function _safeBurn(address zrc20, uint256 amount) private returns (bool) {
        try IZRC20(zrc20).burn(amount) returns (bool success) {
            return success;
        } catch {
            return false;
        }
    }

    // @notice Helper function to safely deposit
    // @param zrc20 The ZRC20 token address
    // @param amount The target address to receive the deposited tokens
    // @param amount The amount to deposit
    // @return True if the deposit was successful, false otherwise
    function _safeDeposit(address zrc20, address target, uint256 amount) private returns (bool) {
        try IZRC20(zrc20).deposit(target, amount) returns (bool success) {
            return success;
        } catch {
            return false;
        }
    }

    /// @notice Helper function to burn gas fees.
    /// @param gasZRC20 The address of the gas ZRC20 token.
    /// @param gasFee The amount of gasZRC20 that should be burned.
    function _burnProtocolFees(address gasZRC20, uint256 gasFee) private {
        if (!_safeTransferFrom(gasZRC20, msg.sender, address(this), gasFee)) {
            revert GasFeeTransferFailed(gasZRC20, address(this), gasFee);
        }

        if (!_safeBurn(gasZRC20, gasFee)) {
            revert ZRC20BurnFailed(gasZRC20, gasFee);
        }
    }

    /// @notice Helper function to burn gas fees for ZRC20s withdrawals.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param gasLimit Gas limit.
    function _burnZRC20ProtocolFees(address zrc20, uint256 gasLimit) private returns (uint256) {
        (address gasZRC20, uint256 gasFee) = IZRC20(zrc20).withdrawGasFeeWithGasLimit(gasLimit);
        _burnProtocolFees(gasZRC20, gasFee);
        return gasFee;
    }

    /// @dev Private function to withdraw ZRC20 tokens with gas limit.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param gasLimit Gas limit.
    /// @return The gas fee for the withdrawal.
    function _withdrawZRC20WithGasLimit(uint256 amount, address zrc20, uint256 gasLimit) private returns (uint256) {
        uint256 gasFee = _burnZRC20ProtocolFees(zrc20, gasLimit);

        if (!_safeTransferFrom(zrc20, msg.sender, address(this), amount)) {
            revert ZRC20TransferFailed(zrc20, msg.sender, address(this), amount);
        }

        if (!_safeBurn(zrc20, amount)) {
            revert ZRC20BurnFailed(zrc20, amount);
        }

        return gasFee;
    }

    /// @dev Helper function to get gas limit for the ZETA transfer to the external chain.
    /// @param chainId Chain id of the external chain.
    /// @return gasLimit The gas limit.
    function _getGasLimitForZETATransfer(uint256 chainId) private view returns (uint256 gasLimit) {
        // default gas limit for ZETA transfer
        uint256 DEFAULT_GAS_LIMIT = 100_000;
        // fetch gasLimit for the external chain
        try ICoreRegistry(registry).getChainMetadata(chainId, "gasLimit") returns (bytes memory _gasLimit) {
            if (_gasLimit.length > 0) {
                gasLimit = abi.decode(_gasLimit, (uint256));
            } else {
                gasLimit = DEFAULT_GAS_LIMIT;
            }
        } catch {
            gasLimit = DEFAULT_GAS_LIMIT;
        }
    }

    /// @dev Helper function to get protocol flat fee for the external chain.
    /// @param chainId Chain id of the external chain.
    /// @return protocolFlatFee The protocol flat fee.
    function _getProtocolFlatFeeFromRegistry(uint256 chainId) private view returns (uint256 protocolFlatFee) {
        // fetch protocolFlatFee for the external chain
        try ICoreRegistry(registry)
            .getChainMetadata(chainId, "protocolFlatFee") returns (bytes memory _protocolFlatFee) {
            if (_protocolFlatFee.length > 0) {
                protocolFlatFee = abi.decode(_protocolFlatFee, (uint256));
            } else {
                protocolFlatFee = 0;
            }
        } catch {
            protocolFlatFee = 0;
        }
    }

    /// @dev Helper function to burn gas fees for ZETA withdrawals.
    /// @param chainId Chain id of the external chain.
    /// @param gasLimit The gas limit.
    /// @return gasFee The gas fee for the withdrawal.
    /// @return protocolFlatFee The protocol flat fee.
    /// @return gasLimit_ The gas limit used for the withdrawal.
    function _computeAndPayFeesForZETAWithdrawals(uint256 chainId, uint256 gasLimit)
        private
        returns (uint256 gasFee, uint256 protocolFlatFee, uint256 gasLimit_)
    {
        // get the gas ZRC20 token address for the external chain
        (address gasZRC20,) = ICoreRegistry(registry).getChainInfo(chainId);
        // get the current gas price for the external chain
        uint256 gasPrice = ISystem(IZRC20(gasZRC20).SYSTEM_CONTRACT_ADDRESS()).gasPriceByChainId(chainId);
        if (gasPrice == 0) {
            revert ZeroGasPrice();
        }
        // if gasLimit is not provided, get the gasLimit for ZETA transfer to the external chain
        if (gasLimit == 0) {
            gasLimit = _getGasLimitForZETATransfer(chainId);
        }
        // get the protocol flat fee for the external chain
        protocolFlatFee = _getProtocolFlatFeeFromRegistry(chainId);
        // calculate gas fee
        gasFee = gasPrice * gasLimit + protocolFlatFee;
        // burn protocol fees
        _burnProtocolFees(gasZRC20, gasFee);
        return (gasFee, protocolFlatFee, gasLimit);
    }

    /// @dev Private function to transfer ZETA tokens.
    /// @param amount The amount of tokens to transfer.
    /// @param to The address to transfer the tokens to.
    function _transferZETA(uint256 amount, address to) private {
        (bool sent,) = to.call{ value: amount }("");
        if (!sent) revert FailedZetaSent(to, amount);
    }

    /// @notice Withdraw ZRC20 tokens to an external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param revertOptions Revert options.
    function withdraw(bytes memory receiver, uint256 amount, address zrc20, RevertOptions calldata revertOptions)
        external
        whenNotPaused
    {
        GatewayZEVMValidations.validateWithdrawalParams(receiver, amount, revertOptions);

        uint256 gasFee = _withdrawZRC20WithGasLimit(amount, zrc20, IZRC20(zrc20).GAS_LIMIT());
        emit Withdrawn(
            msg.sender,
            0,
            receiver,
            zrc20,
            amount,
            gasFee,
            IZRC20(zrc20).PROTOCOL_FLAT_FEE(),
            "",
            CallOptions({ gasLimit: IZRC20(zrc20).GAS_LIMIT(), isArbitraryCall: true }),
            revertOptions
        );
    }

    /// @notice Withdraw ZRC20 tokens to an external chain with custom gas limit.
    /// @dev Use this function for simple gas ZRC20 withdrawals to the receivers that are
    ///      either smart contract accounts or smart contracts with custom receive/fallback implementations.
    /// @param receiver The receiver address on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param gasLimit The custom gas limit for the withdrawal (must be >= MIN_GAS_LIMIT).
    /// @param revertOptions Revert options.
    function withdraw(
        bytes memory receiver,
        uint256 amount,
        address zrc20,
        uint256 gasLimit,
        RevertOptions calldata revertOptions
    )
        external
        whenNotPaused
    {
        GatewayZEVMValidations.validateWithdrawalParams(receiver, amount, revertOptions);
        GatewayZEVMValidations.validateGasLimit(gasLimit);

        uint256 gasFee = _withdrawZRC20WithGasLimit(amount, zrc20, gasLimit);
        emit Withdrawn(
            msg.sender,
            0,
            receiver,
            zrc20,
            amount,
            gasFee,
            IZRC20(zrc20).PROTOCOL_FLAT_FEE(),
            "",
            CallOptions({ gasLimit: gasLimit, isArbitraryCall: true }),
            revertOptions
        );
    }

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
        external
        whenNotPaused
    {
        GatewayZEVMValidations.validateWithdrawalAndCallParams(receiver, amount, message, callOptions, revertOptions);

        uint256 gasFee = _withdrawZRC20WithGasLimit(amount, zrc20, callOptions.gasLimit);
        emit WithdrawnAndCalled(
            msg.sender,
            0,
            receiver,
            zrc20,
            amount,
            gasFee,
            IZRC20(zrc20).PROTOCOL_FLAT_FEE(),
            message,
            callOptions,
            revertOptions
        );
    }

    /// @notice Withdraw ZETA tokens to an external chain.
    //// @param receiver The receiver address on the external chain.
    //// @param chainId Chain id of the external chain.
    //// @param revertOptions Revert options.
    function withdraw(bytes memory receiver, uint256 chainId, RevertOptions calldata revertOptions)
        external
        payable
        whenNotPaused
        nonReentrant
    {
        GatewayZEVMValidations.validateWithdrawalParams(receiver, msg.value, revertOptions);
        (uint256 gasFee, uint256 protocolFlatFee, uint256 gasLimit) = _computeAndPayFeesForZETAWithdrawals(chainId, 0);
        _transferZETA(msg.value, PROTOCOL_ADDRESS);

        emit Withdrawn(
            msg.sender,
            chainId,
            receiver,
            address(zetaToken),
            msg.value,
            gasFee,
            protocolFlatFee,
            "",
            CallOptions({ gasLimit: gasLimit, isArbitraryCall: true }),
            revertOptions
        );
    }

    /// @notice Withdraw ZETA tokens and call a smart contract on an external chain.
    //// @param receiver The receiver address on the external chain.
    //// @param chainId Chain id of the external chain.
    //// @param message The calldata to pass to the contract call.
    //// @param callOptions Call options including gas limit and arbirtrary call flag.
    //// @param revertOptions Revert options.
    function withdrawAndCall(
        bytes memory receiver,
        uint256 chainId,
        bytes calldata message,
        CallOptions calldata callOptions,
        RevertOptions calldata revertOptions
    )
        external
        payable
        whenNotPaused
        nonReentrant
    {
        GatewayZEVMValidations.validateWithdrawalAndCallParams(receiver, msg.value, message, callOptions, revertOptions);
        (uint256 gasFee, uint256 protocolFlatFee,) = _computeAndPayFeesForZETAWithdrawals(chainId, callOptions.gasLimit);
        _transferZETA(msg.value, PROTOCOL_ADDRESS);

        emit WithdrawnAndCalled(
            msg.sender,
            chainId,
            receiver,
            address(zetaToken),
            msg.value,
            gasFee,
            protocolFlatFee,
            message,
            callOptions,
            revertOptions
        );
    }

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
        external
        whenNotPaused
    {
        GatewayZEVMValidations.validateCallAndRevertOptions(callOptions, revertOptions, message.length);
        _call(receiver, zrc20, message, callOptions, revertOptions);
    }

    function _call(
        bytes memory receiver,
        address zrc20,
        bytes calldata message,
        CallOptions memory callOptions,
        RevertOptions memory revertOptions
    )
        private
    {
        GatewayZEVMValidations.validateReceiver(receiver);

        _burnZRC20ProtocolFees(zrc20, callOptions.gasLimit);
        emit Called(msg.sender, zrc20, receiver, message, callOptions, revertOptions);
    }

    /// @notice Deposit foreign coins into ZRC20.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param amount The amount of tokens to deposit.
    /// @param target The target address to receive the deposited tokens.
    function deposit(address zrc20, uint256 amount, address target) external onlyProtocol whenNotPaused {
        GatewayZEVMValidations.validateDepositParams(zrc20, amount, target, PROTOCOL_ADDRESS, address(this));

        if (!_safeDeposit(zrc20, target, amount)) {
            revert ZRC20DepositFailed(zrc20, target, amount);
        }
    }

    /// @notice Deposit native ZETA.
    /// @param target The target address to receive the ZETA.
    function deposit(address target) external payable nonReentrant onlyProtocol whenNotPaused {
        GatewayZEVMValidations.validateZetaDepositParams(msg.value, target, PROTOCOL_ADDRESS, address(this));
        _transferZETA(msg.value, target);
    }

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
        external
        nonReentrant
        onlyProtocol
        whenNotPaused
    {
        GatewayZEVMValidations.validateExecuteParams(zrc20, target);
        UniversalContract(target).onCall(context, zrc20, amount, message);
    }

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
        external
        nonReentrant
        onlyProtocol
        whenNotPaused
    {
        GatewayZEVMValidations.validateDepositParams(zrc20, amount, target, PROTOCOL_ADDRESS, address(this));
        if (!_safeDeposit(zrc20, target, amount)) {
            revert ZRC20DepositFailed(zrc20, target, amount);
        }

        UniversalContract(target).onCall(context, zrc20, amount, message);
    }

    /// @notice Deposit native ZETA and call a user-specified contract on ZEVM.
    /// @param context The context of the cross-chain call.
    /// @param target The target contract to call.
    /// @param message The calldata to pass to the contract call.
    function depositAndCall(MessageContext calldata context, address target, bytes calldata message)
        external
        payable
        nonReentrant
        onlyProtocol
        whenNotPaused
    {
        GatewayZEVMValidations.validateZetaDepositParams(msg.value, target, PROTOCOL_ADDRESS, address(this));

        UniversalContract(target).onCall{ value: msg.value }(context, message);
    }

    /// @notice Revert a user-specified contract on ZEVM.
    /// @param target The target contract to call.
    /// @param revertContext Revert context to pass to onRevert.
    function executeRevert(address target, RevertContext calldata revertContext)
        external
        nonReentrant
        onlyProtocol
        whenNotPaused
    {
        GatewayZEVMValidations.validateNonZeroAddress(target);
        Revertable(target).onRevert(revertContext);
    }

    /// @notice Deposit foreign coins into ZRC20 and revert a user-specified contract on ZEVM.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param amount The amount of tokens to revert.
    /// @param target The target contract to call.
    /// @param revertContext Revert context to pass to onRevert.
    function depositAndRevert(address zrc20, uint256 amount, address target, RevertContext calldata revertContext)
        external
        nonReentrant
        onlyProtocol
        whenNotPaused
    {
        GatewayZEVMValidations.validateDepositParams(zrc20, amount, target, PROTOCOL_ADDRESS, address(this));
        if (!_safeDeposit(zrc20, target, amount)) {
            revert ZRC20DepositFailed(zrc20, target, amount);
        }

        Revertable(target).onRevert(revertContext);
    }

    /// @notice Deposit native ZETA and revert a user-specified contract on ZEVM.
    /// @param target The target contract to call.
    /// @param revertContext Revert context to pass to onRevert.
    function depositAndRevert(address target, RevertContext calldata revertContext)
        external
        payable
        nonReentrant
        onlyProtocol
        whenNotPaused
    {
        GatewayZEVMValidations.validateZetaDepositParams(msg.value, target, PROTOCOL_ADDRESS, address(this));

        Revertable(target).onRevert{ value: msg.value }(revertContext);
    }

    /// @notice Call onAbort on a user-specified contract on ZEVM.
    /// this function doesn't deposit the asset to the target contract. This operation is done directly by the protocol.
    /// the assets are deposited to the target contract even if onAbort reverts.
    /// @param target The target contract to call.
    /// @param abortContext Abort context to pass to onAbort.
    function executeAbort(address target, AbortContext calldata abortContext)
        external
        nonReentrant
        onlyProtocol
        whenNotPaused
    {
        GatewayZEVMValidations.validateNonZeroAddress(target);
        Abortable(target).onAbort(abortContext);
    }

    /// @notice Returns the maximum message size.
    /// @return The maximum message size.
    function getMaxMessageSize() external pure returns (uint256) {
        return GatewayZEVMValidations.MAX_MESSAGE_SIZE;
    }

    /// @notice Returns the minimum gas limit allowed.
    /// @return The minimum gas limit.
    function getMinGasLimit() external pure returns (uint256) {
        return GatewayZEVMValidations.MIN_GAS_LIMIT;
    }

    /// @notice Returns the maximum revert gas limit allowed.
    /// @return The maximum revert gas limit.
    function getMaxRevertGasLimit() external pure returns (uint256) {
        return MAX_REVERT_GAS_LIMIT;
    }
}
