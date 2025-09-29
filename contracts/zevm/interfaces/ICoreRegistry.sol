// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "../../helpers/interfaces/IBaseRegistry.sol";

interface ICoreRegistry is IBaseRegistry {
    function gatewayZEVM() external returns (address);
}
