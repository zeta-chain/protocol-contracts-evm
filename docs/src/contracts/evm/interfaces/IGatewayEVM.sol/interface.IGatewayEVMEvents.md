# IGatewayEVMEvents
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/evm/interfaces/IGatewayEVM.sol)

Interface for the events emitted by the GatewayEVM contract.


## Events
### Executed
Emitted when a contract call is executed.


```solidity
event Executed(address indexed destination, uint256 value, bytes data);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`destination`|`address`|The address of the contract called.|
|`value`|`uint256`|The amount of ETH sent with the call.|
|`data`|`bytes`|The calldata passed to the contract call.|

### Reverted
Emitted when a contract call is reverted.


```solidity
event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, RevertContext revertContext);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|The address of the contract called.|
|`token`|`address`|The address of the ERC20 token, empty if gas token|
|`amount`|`uint256`|The amount of ETH sent with the call.|
|`data`|`bytes`|The calldata passed to the contract call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|

### ExecutedWithERC20
Emitted when a contract call with ERC20 tokens is executed.


```solidity
event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|The address of the ERC20 token.|
|`to`|`address`|The address of the contract called.|
|`amount`|`uint256`|The amount of tokens transferred.|
|`data`|`bytes`|The calldata passed to the contract call.|

### Deposited
Emitted when a deposit is made.


```solidity
event Deposited(
    address indexed sender,
    address indexed receiver,
    uint256 amount,
    address asset,
    bytes payload,
    RevertOptions revertOptions
);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|The address of the sender.|
|`receiver`|`address`|The address of the receiver.|
|`amount`|`uint256`|The amount of ETH or tokens deposited.|
|`asset`|`address`|The address of the ERC20 token (zero address if ETH).|
|`payload`|`bytes`|The calldata passed with the deposit. No longer used. Kept to maintain compatibility.|
|`revertOptions`|`RevertOptions`|Revert options.|

### DepositedAndCalled
Emitted when a deposit and call is made.


```solidity
event DepositedAndCalled(
    address indexed sender,
    address indexed receiver,
    uint256 amount,
    address asset,
    bytes payload,
    RevertOptions revertOptions
);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|The address of the sender.|
|`receiver`|`address`|The address of the receiver.|
|`amount`|`uint256`|The amount of ETH or tokens deposited.|
|`asset`|`address`|The address of the ERC20 token (zero address if ETH).|
|`payload`|`bytes`|The calldata passed with the deposit.|
|`revertOptions`|`RevertOptions`|Revert options.|

### Called
Emitted when an omnichain smart contract call is made without asset transfer.


```solidity
event Called(address indexed sender, address indexed receiver, bytes payload, RevertOptions revertOptions);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|The address of the sender.|
|`receiver`|`address`|The address of the receiver.|
|`payload`|`bytes`|The calldata passed to the call.|
|`revertOptions`|`RevertOptions`|Revert options.|

### UpdatedGatewayTSSAddress
Emitted when tss address is updated


```solidity
event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`oldTSSAddress`|`address`|old tss address|
|`newTSSAddress`|`address`|new tss address|

### UpdatedAdditionalActionFee

```solidity
event UpdatedAdditionalActionFee(uint256 oldFeeWei, uint256 newFeeWei);
```

