// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "../../../contracts/Errors.sol";
import "../../../contracts/zevm/interfaces/ISystem.sol";
import "../../../contracts/zevm/interfaces/IWZETA.sol";
import "../../../contracts/zevm/libraries/GatewayZEVMValidations.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";

import { RevertContext, RevertOptions, Revertable } from "../../../contracts/Revert.sol";
import { CallOptions, IGatewayZEVM } from "../../../contracts/zevm/interfaces/IGatewayZEVM.sol";
import { IZRC20 } from "../../../contracts/zevm/interfaces/IZRC20.sol";
import { MessageContext, UniversalContract } from "../../../contracts/zevm/interfaces/UniversalContract.sol";

/// @title GatewayZEVMUpgradeTest
/// @notice Modified GatewayZEVM contract for testing upgrades
/// @dev The only difference is in event naming
/// @custom:oz-upgrades-from GatewayZEVM
contract GatewayZEVMUpgradeTest is
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

    /// @dev Modified event for testing upgrade.
    event WithdrawnV2(
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
            revert EmptyAddress();
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

    /// @notice Helper function to burn gas fees for ZRC20s withdrawals.
    /// @param zrc20 The address of the ZRC20 token.
    /// @return The gas fee for the withdrawal.
    function _withdrawZRC20(uint256 amount, address zrc20) private returns (uint256) {
        // Use gas limit from zrc20
        return _withdrawZRC20WithGasLimit(amount, zrc20, IZRC20(zrc20).GAS_LIMIT());
    }

    /// @dev Private function to withdraw ZRC20 tokens with gas limit.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param gasLimit Gas limit.
    /// @return The gas fee for the withdrawal.
    function _withdrawZRC20WithGasLimit(uint256 amount, address zrc20, uint256 gasLimit) private returns (uint256) {
        (address gasZRC20, uint256 gasFee) = IZRC20(zrc20).withdrawGasFeeWithGasLimit(gasLimit);
        if (!IZRC20(gasZRC20).transferFrom(msg.sender, PROTOCOL_ADDRESS, gasFee)) {
            revert GasFeeTransferFailed();
        }

        if (!_safeTransferFrom(zrc20, msg.sender, address(this), amount)) {
            revert ZRC20TransferFailed(zrc20, msg.sender, address(this), amount);
        }

        if (!_safeBurn(zrc20, amount)) {
            revert ZRC20BurnFailed(zrc20, amount);
        }

        return gasFee;
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
    function withdraw(
        bytes memory receiver,
        uint256 amount,
        address zrc20,
        RevertOptions calldata revertOptions
    )
        external
        whenNotPaused
    {
        GatewayZEVMValidations.validateWithdrawalParams(receiver, amount, revertOptions);

        uint256 gasFee = _withdrawZRC20WithGasLimit(amount, zrc20, IZRC20(zrc20).GAS_LIMIT());
        emit WithdrawnV2(
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
    //// @param amount The amount of tokens to withdraw.
    //// @param revertOptions Revert options.
    function withdraw(
        bytes memory, /*receiver*/
        uint256, /*amount*/
        uint256, /*chainId*/
        RevertOptions calldata /*revertOptions*/
    )
        external
        view
        whenNotPaused
    {
        // TODO: remove error and comment out code once ZETA supported back
        // https://github.com/zeta-chain/protocol-contracts/issues/394
        // ZETA is not currently supported for withdraws
        revert ZETANotSupported();

        // if (receiver.length == 0) revert ZeroAddress();
        // if (amount == 0) revert InsufficientZetaAmount();
        // if (revertOptions.revertMessage.length > MAX_MESSAGE_SIZE) revert MessageSizeExceeded();

        // _transferZETA(amount, PROTOCOL_ADDRESS);
        // emit Withdrawn(
        //     msg.sender,
        //     chainId,
        //     receiver,
        //     address(zetaToken),
        //     amount,
        //     0,
        //     0,
        //     "",
        //     CallOptions({ gasLimit: 0, isArbitraryCall: true }),
        //     revertOptions
        // );
    }

    /// @notice Withdraw ZETA tokens and call a smart contract on an external chain.
    //// @param receiver The receiver address on the external chain.
    //// @param amount The amount of tokens to withdraw.
    //// @param chainId Chain id of the external chain.
    //// @param message The calldata to pass to the contract call.
    //// @param callOptions Call options including gas limit and arbirtrary call flag.
    //// @param revertOptions Revert options.
    function withdrawAndCall(
        bytes memory, /*receiver*/
        uint256, /*amount*/
        uint256, /*chainId*/
        bytes calldata, /*message*/
        CallOptions calldata, /*callOptions*/
        RevertOptions calldata /*revertOptions*/
    )
        external
        view
        whenNotPaused
    {
        // TODO: remove error and comment out code once ZETA supported back
        // https://github.com/zeta-chain/protocol-contracts/issues/394
        // ZETA is not currently supported for withdraws
        revert ZETANotSupported();

        // if (receiver.length == 0) revert ZeroAddress();
        // if (amount == 0) revert InsufficientZetaAmount();
        // if (callOptions.gasLimit < MIN_GAS_LIMIT) revert InsufficientGasLimit();
        // if (message.length + revertOptions.revertMessage.length > MAX_MESSAGE_SIZE) revert MessageSizeExceeded();

        // _transferZETA(amount, PROTOCOL_ADDRESS);
        // emit WithdrawnAndCalled(
        //     msg.sender, chainId, receiver, address(zetaToken), amount, 0, 0, message, callOptions, revertOptions
        // );
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

        (address gasZRC20, uint256 gasFee) = IZRC20(zrc20).withdrawGasFeeWithGasLimit(callOptions.gasLimit);
        if (!IZRC20(gasZRC20).transferFrom(msg.sender, PROTOCOL_ADDRESS, gasFee)) {
            revert GasFeeTransferFailed();
        }

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

    /// @notice Revert a user-specified contract on ZEVM.
    /// @param target The target contract to call.
    /// @param revertContext Revert context to pass to onRevert.
    function executeRevert(
        address target,
        RevertContext calldata revertContext
    )
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
    function depositAndRevert(
        address zrc20,
        uint256 amount,
        address target,
        RevertContext calldata revertContext
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

        Revertable(target).onRevert(revertContext);
    }

    /// @notice Call onAbort on a user-specified contract on ZEVM.
    /// this function doesn't deposit the asset to the target contract. This operation is done directly by the protocol.
    /// the assets are deposited to the target contract even if onAbort reverts.
    /// @param target The target contract to call.
    /// @param abortContext Abort context to pass to onAbort.
    function executeAbort(
        address target,
        AbortContext calldata abortContext
    )
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

    /// @notice Returns the maximum gas limit allowed.
    /// @return The maximum gas limit.
    function getMaxGasLimit() external pure returns (uint256) {
        return GatewayZEVMValidations.MAX_GAS_LIMIT;
    }

    /// @notice Returns the maximum revert gas limit allowed.
    /// @return The maximum revert gas limit.
    function getMaxRevertGasLimit() external pure returns (uint256) {
        return MAX_REVERT_GAS_LIMIT;
    }
}
