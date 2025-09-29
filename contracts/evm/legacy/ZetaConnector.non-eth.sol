// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "./IZetaNonEthInterface.sol";
import "./ZetaConnector.base.sol";
import "./ZetaInterfaces.sol";

/**
 * @dev Non ETH implementation of ZetaConnector.
 * This contract manages interactions between TSS and different chains.
 * This version is for every chain but Etherum network because in the other chains we mint and burn and in Etherum we
 * lock and unlock
 */
contract ZetaConnectorNonEth is ZetaConnectorBase {
    uint256 public maxSupply = 2 ** 256 - 1;

    event MaxSupplyUpdated(address callerAddress, uint256 newMaxSupply);

    constructor(
        address zetaTokenAddress_,
        address tssAddress_,
        address tssAddressUpdater_,
        address pauserAddress_
    )
        ZetaConnectorBase(zetaTokenAddress_, tssAddress_, tssAddressUpdater_, pauserAddress_)
    { }

    function getLockedAmount() external view returns (uint256) {
        return IZetaNonEthInterface(zetaToken).balanceOf(address(this));
    }

    function setMaxSupply(uint256 maxSupply_) external onlyTssAddress {
        maxSupply = maxSupply_;
        emit MaxSupplyUpdated(msg.sender, maxSupply_);
    }

    /**
     * @dev Entry point to send data to protocol
     * This call burn the token and emit an event with all the data needed by the protocol
     */
    function send(ZetaInterfaces.SendInput calldata input) external override whenNotPaused {
        IZetaNonEthInterface(zetaToken).burnFrom(msg.sender, input.zetaValueAndGas);

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
     * This method can be called only by TSS.
     * Transfer the Zeta tokens to destination and calls onZetaMessage if it's needed.
     * To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain.
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
        override
        onlyTssAddress
    {
        if (zetaValue + IZetaNonEthInterface(zetaToken).totalSupply() > maxSupply) revert ExceedsMaxSupply(maxSupply);
        IZetaNonEthInterface(zetaToken).mint(destinationAddress, zetaValue, internalSendHash);

        if (message.length > 0) {
            ZetaReceiver(destinationAddress).onZetaMessage(
                ZetaInterfaces.ZetaMessage(zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message)
            );
        }

        emit ZetaReceived(zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash);
    }

    /**
     * @dev Handler to receive errors from other chain.
     * This method can be called only by TSS.
     * Transfer the Zeta tokens to destination and calls onZetaRevert if it's needed.
     * To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain.
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
        override
        whenNotPaused
        onlyTssAddress
    {
        if (remainingZetaValue + IZetaNonEthInterface(zetaToken).totalSupply() > maxSupply) {
            revert ExceedsMaxSupply(maxSupply);
        }
        IZetaNonEthInterface(zetaToken).mint(zetaTxSenderAddress, remainingZetaValue, internalSendHash);

        if (message.length > 0) {
            ZetaReceiver(zetaTxSenderAddress).onZetaRevert(
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
