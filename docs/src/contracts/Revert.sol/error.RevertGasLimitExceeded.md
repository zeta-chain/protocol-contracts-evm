# RevertGasLimitExceeded
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/Revert.sol)

Error indicating revert gas limit exceeds maximum allowed


```solidity
error RevertGasLimitExceeded(uint256 provided, uint256 maximum);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`provided`|`uint256`|The gas limit provided for revert operation.|
|`maximum`|`uint256`|The maximum allowed gas limit for revert operation.|

