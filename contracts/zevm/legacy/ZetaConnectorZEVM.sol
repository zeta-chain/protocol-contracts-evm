// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../interfaces/IWZETA.sol";

interface ZetaInterfaces {
    /**
     * @dev Use SendInput to interact with the Connector: connector.send(SendInput)
     */
    struct SendInput {
        /// @dev Chain id of the destination chain. More about chain ids
        /// https://docs.zetachain.com/learn/glossary#chain-id
        uint256 destinationChainId;
        /// @dev Address receiving the message on the destination chain (expressed in bytes since it can be non-EVM)
        bytes destinationAddress;
        /// @dev Gas limit for the destination chain's transaction
        uint256 destinationGasLimit;
        /// @dev An encoded, arbitrary message to be parsed by the destination contract
        bytes message;
        /// @dev ZETA to be sent cross-chain + ZetaChain gas fees + destination chain gas fees (expressed in ZETA)
        uint256 zetaValueAndGas;
        /// @dev Optional parameters for the ZetaChain protocol
        bytes zetaParams;
    }

    /**
     * @dev Our Connector calls onZetaMessage with this struct as argument
     */
    struct ZetaMessage {
        bytes zetaTxSenderAddress;
        uint256 sourceChainId;
        address destinationAddress;
        /// @dev Remaining ZETA from zetaValueAndGas after subtracting ZetaChain gas fees and destination gas fees
        uint256 zetaValue;
        bytes message;
    }

    /**
     * @dev Our Connector calls onZetaRevert with this struct as argument
     */
    struct ZetaRevert {
        address zetaTxSenderAddress;
        uint256 sourceChainId;
        bytes destinationAddress;
        uint256 destinationChainId;
        /// @dev Equals to: zetaValueAndGas - ZetaChain gas fees - destination chain gas fees - source chain revert tx
        /// gas fees
        uint256 remainingZetaValue;
        bytes message;
    }
}

interface ZetaReceiver {
    /**
     * @dev onZetaMessage is called when a cross-chain message reaches a contract
     */
    function onZetaMessage(ZetaInterfaces.ZetaMessage calldata zetaMessage) external;

    /**
     * @dev onZetaRevert is called when a cross-chain message reverts.
     * It's useful to rollback to the original state
     */
    function onZetaRevert(ZetaInterfaces.ZetaRevert calldata zetaRevert) external;
}

contract ZetaConnectorZEVM {
    /// @notice WZETA token address.
    address public wzeta;

    /// @notice Contract custom errors.
    error OnlyWZETAOrFungible();
    error WZETATransferFailed();
    error OnlyFungibleModule();
    error FailedZetaSent();
    error WrongValue();

    /// @notice Fungible module address.
    address public constant FUNGIBLE_MODULE_ADDRESS = payable(0x735b14BB79463307AAcBED86DAf3322B1e6226aB);

    event SetWZETA(address wzeta_);

    event ZetaSent(
        address sourceTxOriginAddress,
        address indexed zetaTxSenderAddress,
        uint256 indexed destinationChainId,
        bytes destinationAddress,
        uint256 zetaValueAndGas,
        uint256 destinationGasLimit,
        bytes message,
        bytes zetaParams
    );

    event ZetaReceived(
        bytes zetaTxSenderAddress,
        uint256 indexed sourceChainId,
        address indexed destinationAddress,
        uint256 zetaValue,
        bytes message,
        bytes32 indexed internalSendHash
    );

    event ZetaReverted(
        address zetaTxSenderAddress,
        uint256 sourceChainId,
        uint256 indexed destinationChainId,
        bytes destinationAddress,
        uint256 remainingZetaValue,
        bytes message,
        bytes32 indexed internalSendHash
    );

    /**
     * @dev Modifier to restrict actions to fungible module.
     */
    modifier onlyFungibleModule() {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert OnlyFungibleModule();
        _;
    }

    constructor(address wzeta_) {
        wzeta = wzeta_;
    }

    /// @dev Receive function to receive ZETA from WETH9.withdraw().
    receive() external payable {
        if (msg.sender != wzeta && msg.sender != FUNGIBLE_MODULE_ADDRESS) revert OnlyWZETAOrFungible();
    }

    function setWzetaAddress(address wzeta_) external onlyFungibleModule {
        wzeta = wzeta_;
        emit SetWZETA(wzeta_);
    }

    /**
     * @dev Sends ZETA and bytes messages (to execute it) crosschain.
     * @param input, SendInput struct, checkout above.
     */
    function send(ZetaInterfaces.SendInput calldata input) external {
        // Transfer wzeta to "fungible" module, which will be burnt by the protocol post processing via hooks.
        if (!IWETH9(wzeta).transferFrom(msg.sender, address(this), input.zetaValueAndGas)) {
            revert WZETATransferFailed();
        }
        IWETH9(wzeta).withdraw(input.zetaValueAndGas);
        (bool sent,) = FUNGIBLE_MODULE_ADDRESS.call{ value: input.zetaValueAndGas }("");
        if (!sent) revert FailedZetaSent();
        emit ZetaSent(
            tx.origin,
            msg.sender,
            input.destinationChainId,
            input.destinationAddress,
            input.zetaValueAndGas,
            input.destinationGasLimit,
            input.message,
            input.zetaParams
        );
    }

    /**
     * @dev Handler to receive data from other chain.
     * This method can be called only by Fungible Module.
     * Transfer the Zeta tokens to destination and calls onZetaMessage if it's needed.
     * To perform the transfer wrap the new tokens
     */
    function onReceive(
        bytes calldata zetaTxSenderAddress,
        uint256 sourceChainId,
        address destinationAddress,
        uint256 zetaValue,
        bytes calldata message,
        bytes32 internalSendHash
    )
        external
        payable
        onlyFungibleModule
    {
        if (msg.value != zetaValue) revert WrongValue();
        IWETH9(wzeta).deposit{ value: zetaValue }();
        if (!IWETH9(wzeta).transferFrom(address(this), destinationAddress, zetaValue)) revert WZETATransferFailed();

        if (message.length > 0) {
            ZetaReceiver(destinationAddress)
                .onZetaMessage(
                    ZetaInterfaces.ZetaMessage(
                        zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message
                    )
                );
        }

        emit ZetaReceived(zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash);
    }

    /**
     * @dev Handler to receive errors from other chain.
     * This method can be called only by Fungible Module.
     * Transfer the Zeta tokens to destination and calls onZetaRevert if it's needed.
     */
    function onRevert(
        address zetaTxSenderAddress,
        uint256 sourceChainId,
        bytes calldata destinationAddress,
        uint256 destinationChainId,
        uint256 remainingZetaValue,
        bytes calldata message,
        bytes32 internalSendHash
    )
        external
        payable
        onlyFungibleModule
    {
        if (msg.value != remainingZetaValue) revert WrongValue();
        IWETH9(wzeta).deposit{ value: remainingZetaValue }();
        if (!IWETH9(wzeta).transferFrom(address(this), zetaTxSenderAddress, remainingZetaValue)) {
            revert WZETATransferFailed();
        }

        if (message.length > 0) {
            ZetaReceiver(zetaTxSenderAddress)
                .onZetaRevert(
                    ZetaInterfaces.ZetaRevert(
                        zetaTxSenderAddress,
                        sourceChainId,
                        destinationAddress,
                        destinationChainId,
                        remainingZetaValue,
                        message
                    )
                );
        }

        emit ZetaReverted(
            zetaTxSenderAddress,
            sourceChainId,
            destinationChainId,
            destinationAddress,
            remainingZetaValue,
            message,
            internalSendHash
        );
    }
}
