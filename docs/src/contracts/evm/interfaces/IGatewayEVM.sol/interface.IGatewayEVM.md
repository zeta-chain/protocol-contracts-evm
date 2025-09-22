# IGatewayEVM
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/evm/interfaces/IGatewayEVM.sol)

**Inherits:**
[IGatewayEVMErrors](/contracts/evm/interfaces/IGatewayEVM.sol/interface.IGatewayEVMErrors.md), [IGatewayEVMEvents](/contracts/evm/interfaces/IGatewayEVM.sol/interface.IGatewayEVMEvents.md)

Interface for the GatewayEVM contract.


## Functions
### executeWithERC20

Executes a call to a contract using ERC20 tokens.


```solidity
function executeWithERC20(
    MessageContext calldata messageContext,
    address token,
    address to,
    uint256 amount,
    bytes calldata data
)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`messageContext`|`MessageContext`|Message context containing sender and arbitrary call flag.|
|`token`|`address`|The address of the ERC20 token.|
|`to`|`address`|The address of the contract to call.|
|`amount`|`uint256`|The amount of tokens to transfer.|
|`data`|`bytes`|The calldata to pass to the contract call.|


### executeRevert

Transfers msg.value to destination contract and executes it's onRevert function.

*This function can only be called by the TSS address and it is payable.*


```solidity
function executeRevert(
    address destination,
    bytes calldata data,
    RevertContext calldata revertContext
)
    external
    payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`destination`|`address`|Address to call.|
|`data`|`bytes`|Calldata to pass to the call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|


### execute

Executes a call to a destination address without ERC20 tokens.

*This function can only be called by the TSS address and it is payable.*


```solidity
function execute(
    MessageContext calldata messageContext,
    address destination,
    bytes calldata data
)
    external
    payable
    returns (bytes memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`messageContext`|`MessageContext`|Message context containing sender and arbitrary call flag.|
|`destination`|`address`|Address to call.|
|`data`|`bytes`|Calldata to pass to the call.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes`|The result of the call.|


### revertWithERC20

Executes a revertable call to a contract using ERC20 tokens.


```solidity
function revertWithERC20(
    address token,
    address to,
    uint256 amount,
    bytes calldata data,
    RevertContext calldata revertContext
)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|The address of the ERC20 token.|
|`to`|`address`|The address of the contract to call.|
|`amount`|`uint256`|The amount of tokens to transfer.|
|`data`|`bytes`|The calldata to pass to the contract call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|


### deposit

Deposits ETH to the TSS address.


```solidity
function deposit(address receiver, RevertOptions calldata revertOptions) external payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`revertOptions`|`RevertOptions`|Revert options.|


### deposit

Deposits ETH to the TSS address with specified amount.


```solidity
function deposit(address receiver, uint256 amount, RevertOptions calldata revertOptions) external payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`amount`|`uint256`|Amount of ETH to deposit.|
|`revertOptions`|`RevertOptions`|Revert options.|


### deposit

Deposits ERC20 tokens to the custody or connector contract.


```solidity
function deposit(
    address receiver,
    uint256 amount,
    address asset,
    RevertOptions calldata revertOptions
)
    external
    payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`amount`|`uint256`|Amount of tokens to deposit.|
|`asset`|`address`|Address of the ERC20 token.|
|`revertOptions`|`RevertOptions`|Revert options.|


### depositAndCall

Deposits ETH to the TSS address and calls an omnichain smart contract.


```solidity
function depositAndCall(
    address receiver,
    bytes calldata payload,
    RevertOptions calldata revertOptions
)
    external
    payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`payload`|`bytes`|Calldata to pass to the call.|
|`revertOptions`|`RevertOptions`|Revert options.|


### depositAndCall

Deposits ETH to the TSS address and calls an omnichain smart contract with specified amount.


```solidity
function depositAndCall(
    address receiver,
    uint256 amount,
    bytes calldata payload,
    RevertOptions calldata revertOptions
)
    external
    payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`amount`|`uint256`|Amount of ETH to deposit.|
|`payload`|`bytes`|Calldata to pass to the call.|
|`revertOptions`|`RevertOptions`|Revert options.|


### depositAndCall

Deposits ERC20 tokens to the custody or connector contract and calls an omnichain smart contract.


```solidity
function depositAndCall(
    address receiver,
    uint256 amount,
    address asset,
    bytes calldata payload,
    RevertOptions calldata revertOptions
)
    external
    payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`amount`|`uint256`|Amount of tokens to deposit.|
|`asset`|`address`|Address of the ERC20 token.|
|`payload`|`bytes`|Calldata to pass to the call.|
|`revertOptions`|`RevertOptions`|Revert options.|


### call

Calls an omnichain smart contract without asset transfer.


```solidity
function call(address receiver, bytes calldata payload, RevertOptions calldata revertOptions) external payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`payload`|`bytes`|Calldata to pass to the call.|
|`revertOptions`|`RevertOptions`|Revert options.|


