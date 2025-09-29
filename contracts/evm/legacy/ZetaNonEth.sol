// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";

import "./ZetaErrors.sol";

import "./IZetaNonEthInterface.sol";

/**
 * @notice On non-native (non-Ethereum) chains, ZETA tokens are minted and burned after the initial deployment on
 * Ethereum.
 */
contract ZetaNonEth is IZetaNonEthInterface, ERC20Burnable, ZetaErrors {
    address public connectorAddress;

    /**
     * @dev Collectively held by Zeta blockchain validators
     */
    address public tssAddress;

    /**
     * @dev Initially a multi-sig, eventually held by Zeta blockchain validators (via renounceTssAddressUpdater)
     */
    address public tssAddressUpdater;

    event Minted(address indexed mintee, uint256 amount, bytes32 indexed internalSendHash);

    event Burnt(address indexed burnee, uint256 amount);

    event TSSAddressUpdated(address callerAddress, address newTssAddress);

    event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress);

    event ConnectorAddressUpdated(address callerAddress, address newConnectorAddress);

    constructor(address tssAddress_, address tssAddressUpdater_) ERC20("Zeta", "ZETA") {
        if (tssAddress_ == address(0) || tssAddressUpdater_ == address(0)) revert InvalidAddress();

        tssAddress = tssAddress_;
        tssAddressUpdater = tssAddressUpdater_;
    }

    function updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) external {
        if (msg.sender != tssAddressUpdater && msg.sender != tssAddress) revert CallerIsNotTssOrUpdater(msg.sender);
        if (tssAddress_ == address(0) || connectorAddress_ == address(0)) revert InvalidAddress();

        tssAddress = tssAddress_;
        connectorAddress = connectorAddress_;

        emit TSSAddressUpdated(msg.sender, tssAddress_);
        emit ConnectorAddressUpdated(msg.sender, connectorAddress_);
    }

    /**
     * @dev Sets tssAddressUpdater to be tssAddress
     */
    function renounceTssAddressUpdater() external {
        if (msg.sender != tssAddressUpdater) revert CallerIsNotTssUpdater(msg.sender);
        if (tssAddress == address(0)) revert InvalidAddress();

        tssAddressUpdater = tssAddress;
        emit TSSAddressUpdaterUpdated(msg.sender, tssAddress);
    }

    function mint(address mintee, uint256 value, bytes32 internalSendHash) external override {
        /**
         * @dev Only Connector can mint. Minting requires burning the equivalent amount on another chain
         */
        if (msg.sender != connectorAddress) revert CallerIsNotConnector(msg.sender);

        _mint(mintee, value);

        emit Minted(mintee, value, internalSendHash);
    }

    function burnFrom(address account, uint256 amount) public override(IZetaNonEthInterface, ERC20Burnable) {
        /**
         * @dev Only Connector can burn.
         */
        if (msg.sender != connectorAddress) revert CallerIsNotConnector(msg.sender);

        ERC20Burnable.burnFrom(account, amount);

        emit Burnt(account, amount);
    }
}
