// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { CallOptions, CallOptionsV2, IGatewayZEVM } from "./interfaces/IGatewayZEVM.sol";

import { INotSupportedMethods } from "../../contracts/Errors.sol";
import { AbortContext, Abortable, RevertContext, RevertOptions, Revertable } from "../../contracts/Revert.sol";
import "./interfaces/IWZETA.sol";
import { IZRC20 } from "./interfaces/IZRC20.sol";
import { MessageContext, UniversalContract } from "./interfaces/UniversalContract.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

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
    /// @notice Error indicating a zero address was provided.
    error ZeroAddress();

    /// @notice The constant address of the protocol
    address public constant PROTOCOL_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;

    /// @notice The address of the Zeta token.
    address public zetaToken;

    /// @notice New role identifier for pauser role.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    /// @notice Max size of message + revertOptions revert message.
    uint256 public constant MAX_MESSAGE_SIZE = 2880;

    /// @notice Minimum gas limit for a call.
    uint256 public constant MIN_GAS_LIMIT = 100_000;

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
            revert ZeroAddress();
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

    /// @dev Private function to withdraw ZRC20 tokens.
    /// @param amount The amount of tokens to withdraw.
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

        if (!IZRC20(zrc20).transferFrom(msg.sender, address(this), amount)) {
            revert ZRC20TransferFailed();
        }

        if (!IZRC20(zrc20).burn(amount)) revert ZRC20BurnFailed();

        return gasFee;
    }

    /// @dev Private function to transfer ZETA tokens.
    /// @param amount The amount of tokens to transfer.
    /// @param to The address to transfer the tokens to.
    function _transferZETA(uint256 amount, address to) private {
        if (!IWETH9(zetaToken).transferFrom(msg.sender, address(this), amount)) revert FailedZetaSent();
        IWETH9(zetaToken).withdraw(amount);
        (bool sent,) = to.call{ value: amount }("");
        if (!sent) revert FailedZetaSent();
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
        if (receiver.length == 0) revert ZeroAddress();
        if (amount == 0) revert InsufficientZRC20Amount();
        if (revertOptions.revertMessage.length > MAX_MESSAGE_SIZE) revert MessageSizeExceeded();

        uint256 gasFee = _withdrawZRC20(amount, zrc20);
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
        if (receiver.length == 0) revert ZeroAddress();
        if (amount == 0) revert InsufficientZRC20Amount();
        if (gasLimit < MIN_GAS_LIMIT) revert InsufficientGasLimit();
        if (revertOptions.revertMessage.length > MAX_MESSAGE_SIZE) revert MessageSizeExceeded();

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
        if (receiver.length == 0) revert ZeroAddress();
        if (amount == 0) revert InsufficientZRC20Amount();
        if (callOptions.gasLimit < MIN_GAS_LIMIT) revert InsufficientGasLimit();
        if (message.length + revertOptions.revertMessage.length > MAX_MESSAGE_SIZE) revert MessageSizeExceeded();

        // Sui mainnet not supported for now
        require(IZRC20(zrc20).CHAIN_ID() != 105);

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

    /// @notice Withdraw ZRC20 tokens and call a smart contract on an external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param zrc20 The address of the ZRC20 token.
    /// @param message The calldata to pass to the contract call.
    /// @param callOptions Call options including gas limit, arbirtrary call flag and message context version.
    /// @param revertOptions Revert options.
    function withdrawAndCall(
        bytes memory receiver,
        uint256 amount,
        address zrc20,
        bytes calldata message,
        CallOptionsV2 calldata callOptions,
        RevertOptions calldata revertOptions
    )
        external
        whenNotPaused
    {
        if (receiver.length == 0) revert ZeroAddress();
        if (amount == 0) revert InsufficientZRC20Amount();
        if (callOptions.gasLimit < MIN_GAS_LIMIT) revert InsufficientGasLimit();
        if (message.length + revertOptions.revertMessage.length > MAX_MESSAGE_SIZE) revert MessageSizeExceeded();

        uint256 gasFee = _withdrawZRC20WithGasLimit(amount, zrc20, callOptions.gasLimit);
        emit WithdrawnAndCalledV2(
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
        if (callOptions.gasLimit < MIN_GAS_LIMIT) revert InsufficientGasLimit();
        if (message.length + revertOptions.revertMessage.length > MAX_MESSAGE_SIZE) revert MessageSizeExceeded();

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
        if (receiver.length == 0) revert ZeroAddress();

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
        if (zrc20 == address(0) || target == address(0)) revert ZeroAddress();
        if (amount == 0) revert InsufficientZRC20Amount();

        if (target == PROTOCOL_ADDRESS || target == address(this)) revert InvalidTarget();

        if (!IZRC20(zrc20).deposit(target, amount)) revert ZRC20DepositFailed();
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
        if (zrc20 == address(0) || target == address(0)) revert ZeroAddress();

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
        if (zrc20 == address(0) || target == address(0)) revert ZeroAddress();
        if (amount == 0) revert InsufficientZRC20Amount();
        if (target == PROTOCOL_ADDRESS || target == address(this)) revert InvalidTarget();

        if (!IZRC20(zrc20).deposit(target, amount)) revert ZRC20DepositFailed();
        UniversalContract(target).onCall(context, zrc20, amount, message);
    }

    /// @notice Deposit ZETA and call a user-specified contract on ZEVM.
    /// @param context The context of the cross-chain call.
    /// @param amount The amount of tokens to transfer.
    /// @param target The target contract to call.
    /// @param message The calldata to pass to the contract call.
    function depositAndCall(
        MessageContext calldata context,
        uint256 amount,
        address target,
        bytes calldata message
    )
        external
        nonReentrant
        onlyProtocol
        whenNotPaused
    {
        if (target == address(0)) revert ZeroAddress();
        if (amount == 0) revert InsufficientZetaAmount();
        if (target == PROTOCOL_ADDRESS || target == address(this)) revert InvalidTarget();

        _transferZETA(amount, target);
        UniversalContract(target).onCall(context, zetaToken, amount, message);
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
        if (target == address(0)) revert ZeroAddress();

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
        if (zrc20 == address(0) || target == address(0)) revert ZeroAddress();
        if (amount == 0) revert InsufficientZRC20Amount();
        if (target == PROTOCOL_ADDRESS || target == address(this)) revert InvalidTarget();

        if (!IZRC20(zrc20).deposit(target, amount)) revert ZRC20DepositFailed();
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
        if (target == address(0)) revert ZeroAddress();
        Abortable(target).onAbort(abortContext);
    }
}
