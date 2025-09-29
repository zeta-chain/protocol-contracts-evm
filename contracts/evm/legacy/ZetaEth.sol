// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @dev ZetaEth.sol is an implementation of OpenZeppelin's ERC20
 * @notice Ethereum is the origin and native chain of the ZETA token deployment (native)
 */
contract ZetaEth is ERC20("Zeta", "ZETA") {
    constructor(address creator, uint256 initialSupply) {
        _mint(creator, initialSupply * (10 ** uint256(decimals())));
    }
}
