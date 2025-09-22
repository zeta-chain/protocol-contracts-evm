# IGatewayEVMErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/evm/interfaces/IGatewayEVM.sol)

Interface for the errors used in the GatewayEVM contract.


## Errors
### ExecutionFailed
Error for failed execution.


```solidity
error ExecutionFailed();
```

### DepositFailed
Error for failed deposit.


```solidity
error DepositFailed();
```

### InsufficientETHAmount
Error for insufficient ETH amount.


```solidity
error InsufficientETHAmount();
```

### InsufficientERC20Amount
Error for insufficient ERC20 token amount.


```solidity
error InsufficientERC20Amount();
```

### ZeroAddress
Error for zero address input.


```solidity
error ZeroAddress();
```

### ApprovalFailed
Error for failed token approval.


```solidity
error ApprovalFailed();
```

### CustodyInitialized
Error for already initialized custody.


```solidity
error CustodyInitialized();
```

### ConnectorInitialized
Error for already initialized connector.


```solidity
error ConnectorInitialized();
```

### NotWhitelistedInCustody
Error when trying to transfer not whitelisted token to custody.


```solidity
error NotWhitelistedInCustody();
```

### NotAllowedToCallOnCall
Error when trying to call onCall method using arbitrary call.


```solidity
error NotAllowedToCallOnCall();
```

### NotAllowedToCallOnRevert
Error when trying to call onRevert method using arbitrary call.


```solidity
error NotAllowedToCallOnRevert();
```

### PayloadSizeExceeded
Error indicating payload size exceeded in external functions.


```solidity
error PayloadSizeExceeded(uint256 provided, uint256 maximum);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`provided`|`uint256`|The size of the payload that was provided.|
|`maximum`|`uint256`|The maximum allowed payload size.|

### FeeTransferFailed
Error thrown when fee transfer to TSS address fails.

*This error occurs when the low-level call to transfer fees fails.*


```solidity
error FeeTransferFailed();
```

### InsufficientFee
Error thrown when insufficient fee is provided for additional actions.


```solidity
error InsufficientFee(uint256 required, uint256 provided);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`required`|`uint256`|The fee amount required for the action.|
|`provided`|`uint256`|The fee amount actually provided by the caller.|

### ExcessETHProvided
Error thrown when excess ETH is sent for non-ETH operations.


```solidity
error ExcessETHProvided(uint256 required, uint256 provided);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`required`|`uint256`|The fee amount required for the action.|
|`provided`|`uint256`|The ETH amount actually provided by the caller.|

### AdditionalActionDisabled
Error thrown when additional action functionality is disabled (fee set to 0).


```solidity
error AdditionalActionDisabled();
```

### IncorrectValueProvided
Error thrown when msg.value doesn't match expected amount + fee.


```solidity
error IncorrectValueProvided(uint256 expected, uint256 provided);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`expected`|`uint256`|The expected value (amount + fee).|
|`provided`|`uint256`|The actual msg.value provided.|

