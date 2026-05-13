// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { TestERC20 } from "./TestERC20.sol";

/// @title FeeOnTransferERC20
/// @notice ERC20 test token that charges a transfer fee.
contract FeeOnTransferERC20 is TestERC20 {
    uint256 public immutable FEE_BPS;

    constructor(string memory name, string memory symbol, uint256 feeBps_) TestERC20(name, symbol) {
        FEE_BPS = feeBps_;
    }

    function _update(address from, address to, uint256 value) internal override {
        if (from == address(0) || to == address(0) || FEE_BPS == 0) {
            super._update(from, to, value);
            return;
        }

        uint256 fee = (value * FEE_BPS) / 10_000;
        super._update(from, to, value - fee);
        super._update(from, address(0xdead), fee);
    }
}
