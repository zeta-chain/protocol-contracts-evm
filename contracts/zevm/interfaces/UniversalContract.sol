// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { RevertContext } from "../../../contracts/Revert.sol";

/// @custom:deprecated should be removed once v2 SystemContract is not used anymore.
/// MessageContext should be used
struct zContext {
    bytes origin;
    address sender;
    uint256 chainID;
}

/// @custom:deprecated should be removed once v2 SystemContract is not used anymore.
/// UniversalContract should be used
interface zContract {
    function onCrossChainCall(
        zContext calldata context,
        address zrc20,
        uint256 amount,
        bytes calldata message
    )
        external;
}

struct MessageContext {
    bytes origin;
    address sender;
    uint256 chainID;
}

interface UniversalContract {
    function onCall(MessageContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
}
