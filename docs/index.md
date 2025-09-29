

## GatewayEVM
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/GatewayEVM.sol)

The GatewayEVM contract is the endpoint to call smart contracts on external chains.

*The contract doesn't hold any funds and should never have active allowances.*


### State Variables
#### custody
The address of the custody contract.


```solidity
address public custody;
```


#### tssAddress
The address of the TSS (Threshold Signature Scheme) contract.


```solidity
address public tssAddress;
```


#### zetaConnector
The address of the ZetaConnector contract.


```solidity
address public zetaConnector;
```


#### zetaToken
The address of the Zeta token contract.


```solidity
address public zetaToken;
```


#### additionalActionFeeWei
Fee charged for additional cross-chain actions within the same transaction.

*The first action in a transaction is free, subsequent actions incur this fee.*

*This is configurable by the admin role to allow for fee adjustments.*


```solidity
uint256 public additionalActionFeeWei;
```


#### TSS_ROLE
New role identifier for tss role.


```solidity
bytes32 public constant TSS_ROLE = keccak256("TSS_ROLE");
```


#### ASSET_HANDLER_ROLE
New role identifier for asset handler role.


```solidity
bytes32 public constant ASSET_HANDLER_ROLE = keccak256("ASSET_HANDLER_ROLE");
```


#### PAUSER_ROLE
New role identifier for pauser role.


```solidity
bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
```


#### MAX_PAYLOAD_SIZE
Max size of payload + revertOptions revert message.


```solidity
uint256 public constant MAX_PAYLOAD_SIZE = 2880;
```


#### _TRANSACTION_ACTION_COUNT_KEY
Storage slot key for tracking transaction action count.

*Uses transient storage (tload/tstore) for gas efficiency.*

*Value 0x01 is used as a unique identifier for this storage slot.*


```solidity
uint256 private constant _TRANSACTION_ACTION_COUNT_KEY = 0x01;
```


### Functions
#### constructor

**Note:**
oz-upgrades-unsafe-allow: constructor


```solidity
constructor();
```

#### initialize

Initialize with tss address. address of zeta token and admin account set as DEFAULT_ADMIN_ROLE.

*Using admin to authorize upgrades and pause, and tss for tss role.*


```solidity
function initialize(address tssAddress_, address zetaToken_, address admin_) public initializer;
```

#### _authorizeUpgrade

*Authorizes the upgrade of the contract, sender must be owner.*


```solidity
function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newImplementation`|`address`|Address of the new implementation.|


#### updateTSSAddress

Update tss address


```solidity
function updateTSSAddress(address newTSSAddress) external onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newTSSAddress`|`address`|new tss address|


#### pause

Pause contract.


```solidity
function pause() external onlyRole(PAUSER_ROLE);
```

#### unpause

Unpause contract.


```solidity
function unpause() external onlyRole(PAUSER_ROLE);
```

#### updateAdditionalActionFee

Update the additional action fee.

*Only callable by admin role. This allows for fee adjustments based on network conditions.*

*Setting fee to 0 disables additional action fees entirely.*

*Fee should be adjusted based on the chain's native token decimals.*


```solidity
function updateAdditionalActionFee(uint256 newFeeWei) external onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newFeeWei`|`uint256`|The new fee amount in wei for additional actions in the same transaction.|


#### executeRevert

Transfers msg.value to destination contract and executes it's onRevert function.

*This function can only be called by the TSS address and it is payable.*


```solidity
function executeRevert(
    address destination,
    bytes calldata data,
    RevertContext calldata revertContext
)
    public
    payable
    nonReentrant
    onlyRole(TSS_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`destination`|`address`|Address to call.|
|`data`|`bytes`|Calldata to pass to the call.|
|`revertContext`|`RevertContext`||


#### execute

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
    nonReentrant
    onlyRole(TSS_ROLE)
    whenNotPaused
    returns (bytes memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`messageContext`|`MessageContext`|Message context containing sender.|
|`destination`|`address`|Address to call.|
|`data`|`bytes`|Calldata to pass to the call.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes`|The result of the call.|


#### executeWithERC20

Executes a call to a destination contract using ERC20 tokens.

*This function can only be called by the custody or connector address.
It uses the ERC20 allowance system, resetting gateway allowance at the end.*


```solidity
function executeWithERC20(
    MessageContext calldata messageContext,
    address token,
    address to,
    uint256 amount,
    bytes calldata data
)
    public
    nonReentrant
    onlyRole(ASSET_HANDLER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`messageContext`|`MessageContext`|Message context containing sender.|
|`token`|`address`|Address of the ERC20 token.|
|`to`|`address`|Address of the contract to call.|
|`amount`|`uint256`|Amount of tokens to transfer.|
|`data`|`bytes`|Calldata to pass to the call.|


#### revertWithERC20

Directly transfers ERC20 tokens and calls onRevert.

*This function can only be called by the custody or connector address.*


```solidity
function revertWithERC20(
    address token,
    address to,
    uint256 amount,
    bytes calldata data,
    RevertContext calldata revertContext
)
    external
    nonReentrant
    onlyRole(ASSET_HANDLER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|Address of the ERC20 token.|
|`to`|`address`|Address of the contract to call.|
|`amount`|`uint256`|Amount of tokens to transfer.|
|`data`|`bytes`|Calldata to pass to the call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|


#### deposit

Deposits ETH to the TSS address.

*This function only works for the first action in a transaction (backward compatibility).*

*For subsequent actions, use the overloaded version with amount parameter.*


```solidity
function deposit(address receiver, RevertOptions calldata revertOptions) external payable whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`revertOptions`|`RevertOptions`|Revert options.|


#### deposit

Deposits ETH to the TSS address with specified amount.

*msg.value must equal amount + required fee for the action.*


```solidity
function deposit(
    address receiver,
    uint256 amount,
    RevertOptions calldata revertOptions
)
    external
    payable
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`amount`|`uint256`|Amount of ETH to deposit (excluding fees).|
|`revertOptions`|`RevertOptions`|Revert options.|


#### deposit

Deposits ERC20 tokens to the custody or connector contract.


```solidity
function deposit(
    address receiver,
    uint256 amount,
    address asset,
    RevertOptions calldata revertOptions
)
    external
    payable
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`amount`|`uint256`|Amount of tokens to deposit.|
|`asset`|`address`|Address of the ERC20 token.|
|`revertOptions`|`RevertOptions`|Revert options.|


#### depositAndCall

Deposits ETH to the TSS address and calls an omnichain smart contract.

*This function only works for the first action in a transaction (backward compatibility).*

*For subsequent actions, use the overloaded version with amount parameter.*


```solidity
function depositAndCall(
    address receiver,
    bytes calldata payload,
    RevertOptions calldata revertOptions
)
    external
    payable
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`payload`|`bytes`|Calldata to pass to the call.|
|`revertOptions`|`RevertOptions`|Revert options.|


#### depositAndCall

Deposits ETH to the TSS address and calls an omnichain smart contract with specified amount.

*msg.value must equal amount + required fee for the action.*


```solidity
function depositAndCall(
    address receiver,
    uint256 amount,
    bytes calldata payload,
    RevertOptions calldata revertOptions
)
    external
    payable
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`amount`|`uint256`|Amount of ETH to deposit (excluding fees).|
|`payload`|`bytes`|Calldata to pass to the call.|
|`revertOptions`|`RevertOptions`|Revert options.|


#### depositAndCall

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
    payable
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`amount`|`uint256`|Amount of tokens to deposit.|
|`asset`|`address`|Address of the ERC20 token.|
|`payload`|`bytes`|Calldata to pass to the call.|
|`revertOptions`|`RevertOptions`|Revert options.|


#### call

Calls an omnichain smart contract without asset transfer.


```solidity
function call(
    address receiver,
    bytes calldata payload,
    RevertOptions calldata revertOptions
)
    external
    payable
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`payload`|`bytes`|Calldata to pass to the call.|
|`revertOptions`|`RevertOptions`|Revert options.|


#### setCustody

Sets the custody contract address.


```solidity
function setCustody(address custody_) external onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`custody_`|`address`|Address of the custody contract.|


#### setConnector

Sets the connector contract address.


```solidity
function setConnector(address zetaConnector_) external onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zetaConnector_`|`address`|Address of the connector contract.|


#### _resetApproval

Resets the approval of a token for a specified address.
This is used to ensure that the approval is set to zero before setting it to a new value.


```solidity
function _resetApproval(address token, address to) private returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|Address of the ERC20 token.|
|`to`|`address`|Address to reset the approval for.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|True if the approval reset was successful or if the token reverts on zero approval.|


#### _transferFromToAssetHandler

*Transfers tokens from the sender to the asset handler.
This function handles the transfer of tokens to either the connector or custody contract based on the asset
type.*


```solidity
function _transferFromToAssetHandler(address from, address token, uint256 amount) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`from`|`address`|Address of the sender.|
|`token`|`address`|Address of the ERC20 token.|
|`amount`|`uint256`|Amount of tokens to transfer.|


#### _transferToAssetHandler

*Transfers tokens to the asset handler.
This function handles the transfer of tokens to either the connector or custody contract based on the asset
type.*


```solidity
function _transferToAssetHandler(address token, uint256 amount) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|Address of the ERC20 token.|
|`amount`|`uint256`|Amount of tokens to transfer.|


#### _executeArbitraryCall

*Private function to execute an arbitrary call to a destination address.*


```solidity
function _executeArbitraryCall(address destination, bytes calldata data) private returns (bytes memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`destination`|`address`|Address to call.|
|`data`|`bytes`|Calldata to pass to the call.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes`|The result of the call.|


#### _executeAuthenticatedCall

*Private function to execute an authenticated call to a destination address.*


```solidity
function _executeAuthenticatedCall(
    MessageContext calldata messageContext,
    address destination,
    bytes calldata data
)
    private
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


#### _revertIfOnCallOrOnRevert


```solidity
function _revertIfOnCallOrOnRevert(bytes calldata data) private pure;
```

#### _processFee

Processes fee collection for cross-chain actions within a transaction.

*The first action in a transaction is free, subsequent actions incur ADDITIONAL_ACTION_FEE_WEI.*

*If fee is 0, the entire functionality is disabled and will revert.*


```solidity
function _processFee() internal returns (uint256);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|The fee amount actually charged (0 for first action, ADDITIONAL_ACTION_FEE_WEI for subsequent actions).|


#### _validateChargedFeeForERC20

Validates fee payment for ERC20 operations (deposit, depositAndCall, call).

*Validates that msg.value equals the required fee (no excess ETH allowed).*


```solidity
function _validateChargedFeeForERC20(uint256 feeCharged) internal view;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`feeCharged`|`uint256`|The fee amount that was charged.|


#### _validateChargedFeeForETHWithAmount

Validates fee payment for ETH operations with specified amount.

*Validates that msg.value equals amount + feeCharged.*


```solidity
function _validateChargedFeeForETHWithAmount(uint256 amount, uint256 feeCharged) internal view;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amount`|`uint256`|The amount to deposit (excluding fees).|
|`feeCharged`|`uint256`|The fee amount that was charged.|


#### _getNextActionIndex

Gets and increments the transaction action counter using transient storage.

*Uses assembly for gas efficiency with tload/tstore operations.*

*Transient storage is transaction-scoped and automatically cleared after each transaction.*


```solidity
function _getNextActionIndex() internal returns (uint256 currentIndex);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`currentIndex`|`uint256`|The current action index within the transaction (0-based).|




## GatewayZEVM
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/GatewayZEVM.sol)

The GatewayZEVM contract is the endpoint to call smart contracts on omnichain.

*The contract doesn't hold any funds and should never have active allowances.*


### State Variables
#### PROTOCOL_ADDRESS
The constant address of the protocol


```solidity
address public constant PROTOCOL_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;
```


#### PAUSER_ROLE
New role identifier for pauser role.


```solidity
bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
```


#### zetaToken
The address of the Zeta token.


```solidity
address public zetaToken;
```


#### registry
The address of the registry contract on ZetaChain


```solidity
address public registry;
```


### Functions
#### onlyProtocol

*Only protocol address allowed modifier.*


```solidity
modifier onlyProtocol();
```

#### constructor

**Note:**
oz-upgrades-unsafe-allow: constructor


```solidity
constructor();
```

#### initialize

Initialize with address of zeta token and admin account set as DEFAULT_ADMIN_ROLE.

*Using admin to authorize upgrades and pause.*


```solidity
function initialize(address zetaToken_, address admin_) public initializer;
```

#### _authorizeUpgrade

*Authorizes the upgrade of the contract.*


```solidity
function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newImplementation`|`address`|The address of the new implementation.|


#### receive

*Receive function to receive ZETA from WETH9.withdraw().*


```solidity
receive() external payable whenNotPaused;
```

#### pause

Pause contract.


```solidity
function pause() external onlyRole(PAUSER_ROLE);
```

#### unpause

Unpause contract.


```solidity
function unpause() external onlyRole(PAUSER_ROLE);
```

#### setRegistryAddress

Set registry address, callable only by DEFAULT_ADMIN_ROLE.


```solidity
function setRegistryAddress(address _registry) external onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`_registry`|`address`|The address of the registry contract.|


#### _safeTransferFrom

Helper function to safely execute transferFrom


```solidity
function _safeTransferFrom(address zrc20, address from, address to, uint256 amount) private returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The ZRC20 token address|
|`from`|`address`|The sender address|
|`to`|`address`|The recipient address|
|`amount`|`uint256`|The amount to transfer|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|True if the transfer was successful, false otherwise.|


#### _safeBurn


```solidity
function _safeBurn(address zrc20, uint256 amount) private returns (bool);
```

#### _safeDeposit


```solidity
function _safeDeposit(address zrc20, address target, uint256 amount) private returns (bool);
```

#### _burnProtocolFees

Helper function to burn gas fees.


```solidity
function _burnProtocolFees(address gasZRC20, uint256 gasFee) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`gasZRC20`|`address`|The address of the gas ZRC20 token.|
|`gasFee`|`uint256`|The amount of gasZRC20 that should be burned.|


#### _burnZRC20ProtocolFees

Helper function to burn gas fees for ZRC20s withdrawals.


```solidity
function _burnZRC20ProtocolFees(address zrc20, uint256 gasLimit) private returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`gasLimit`|`uint256`|Gas limit.|


#### _withdrawZRC20WithGasLimit

*Private function to withdraw ZRC20 tokens with gas limit.*


```solidity
function _withdrawZRC20WithGasLimit(uint256 amount, address zrc20, uint256 gasLimit) private returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amount`|`uint256`|The amount of tokens to withdraw.|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`gasLimit`|`uint256`|Gas limit.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|The gas fee for the withdrawal.|


#### _getGasLimitForZETATransfer

*Helper function to get gas limit for the ZETA transfer to the external chain.*


```solidity
function _getGasLimitForZETATransfer(uint256 chainId) private view returns (uint256 gasLimit);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|Chain id of the external chain.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`gasLimit`|`uint256`|The gas limit.|


#### _getProtocolFlatFeeFromRegistry

*Helper function to get protocol flat fee for the external chain.*


```solidity
function _getProtocolFlatFeeFromRegistry(uint256 chainId) private view returns (uint256 protocolFlatFee);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|Chain id of the external chain.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`protocolFlatFee`|`uint256`|The protocol flat fee.|


#### _computeAndPayFeesForZETAWithdrawals

*Helper function to burn gas fees for ZETA withdrawals.*


```solidity
function _computeAndPayFeesForZETAWithdrawals(
    uint256 chainId,
    uint256 gasLimit
)
    private
    returns (uint256 gasFee, uint256 protocolFlatFee, uint256 gasLimit_);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|Chain id of the external chain.|
|`gasLimit`|`uint256`|The gas limit.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`gasFee`|`uint256`|The gas fee for the withdrawal.|
|`protocolFlatFee`|`uint256`|The protocol flat fee.|
|`gasLimit_`|`uint256`|The gas limit used for the withdrawal.|


#### _transferZETA

*Private function to transfer ZETA tokens.*


```solidity
function _transferZETA(uint256 amount, address to) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amount`|`uint256`|The amount of tokens to transfer.|
|`to`|`address`|The address to transfer the tokens to.|


#### withdraw

Withdraw ZRC20 tokens to an external chain.


```solidity
function withdraw(
    bytes memory receiver,
    uint256 amount,
    address zrc20,
    RevertOptions calldata revertOptions
)
    external
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver address on the external chain.|
|`amount`|`uint256`|The amount of tokens to withdraw.|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`revertOptions`|`RevertOptions`|Revert options.|


#### withdraw

Withdraw ZRC20 tokens to an external chain with custom gas limit.

*Use this function for simple gas ZRC20 withdrawals to the receivers that are
either smart contract accounts or smart contracts with custom receive/fallback implementations.*


```solidity
function withdraw(
    bytes memory receiver,
    uint256 amount,
    address zrc20,
    uint256 gasLimit,
    RevertOptions calldata revertOptions
)
    external
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver address on the external chain.|
|`amount`|`uint256`|The amount of tokens to withdraw.|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`gasLimit`|`uint256`|The custom gas limit for the withdrawal (must be >= MIN_GAS_LIMIT).|
|`revertOptions`|`RevertOptions`|Revert options.|


#### withdrawAndCall

Withdraw ZRC20 tokens and call a smart contract on an external chain.


```solidity
function withdrawAndCall(
    bytes memory receiver,
    uint256 amount,
    address zrc20,
    bytes calldata message,
    CallOptions calldata callOptions,
    RevertOptions calldata revertOptions
)
    external
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver address on the external chain.|
|`amount`|`uint256`|The amount of tokens to withdraw.|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`message`|`bytes`|The calldata to pass to the contract call.|
|`callOptions`|`CallOptions`|Call options including gas limit and arbirtrary call flag.|
|`revertOptions`|`RevertOptions`|Revert options.|


#### withdraw

Withdraw ZETA tokens to an external chain.


```solidity
function withdraw(
    bytes memory receiver,
    uint256 chainId,
    RevertOptions calldata revertOptions
)
    external
    payable
    whenNotPaused
    nonReentrant;
```

#### withdrawAndCall

Withdraw ZETA tokens and call a smart contract on an external chain.


```solidity
function withdrawAndCall(
    bytes memory receiver,
    uint256 chainId,
    bytes calldata message,
    CallOptions calldata callOptions,
    RevertOptions calldata revertOptions
)
    external
    payable
    whenNotPaused
    nonReentrant;
```

#### call

Call a smart contract on an external chain without asset transfer.


```solidity
function call(
    bytes memory receiver,
    address zrc20,
    bytes calldata message,
    CallOptions calldata callOptions,
    RevertOptions calldata revertOptions
)
    external
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver address on the external chain.|
|`zrc20`|`address`|Address of zrc20 to pay fees.|
|`message`|`bytes`|The calldata to pass to the contract call.|
|`callOptions`|`CallOptions`|Call options including gas limit and arbirtrary call flag.|
|`revertOptions`|`RevertOptions`|Revert options.|


#### _call


```solidity
function _call(
    bytes memory receiver,
    address zrc20,
    bytes calldata message,
    CallOptions memory callOptions,
    RevertOptions memory revertOptions
)
    private;
```

#### deposit

Deposit foreign coins into ZRC20.


```solidity
function deposit(address zrc20, uint256 amount, address target) external onlyProtocol whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`amount`|`uint256`|The amount of tokens to deposit.|
|`target`|`address`|The target address to receive the deposited tokens.|


#### deposit

Deposit native ZETA.


```solidity
function deposit(address target) external payable nonReentrant onlyProtocol whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`target`|`address`|The target address to receive the ZETA.|


#### execute

Execute a user-specified contract on ZEVM.


```solidity
function execute(
    MessageContext calldata context,
    address zrc20,
    uint256 amount,
    address target,
    bytes calldata message
)
    external
    nonReentrant
    onlyProtocol
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`context`|`MessageContext`|The context of the cross-chain call.|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`amount`|`uint256`|The amount of tokens to transfer.|
|`target`|`address`|The target contract to call.|
|`message`|`bytes`|The calldata to pass to the contract call.|


#### depositAndCall

Deposit foreign coins into ZRC20 and call a user-specified contract on ZEVM.


```solidity
function depositAndCall(
    MessageContext calldata context,
    address zrc20,
    uint256 amount,
    address target,
    bytes calldata message
)
    external
    nonReentrant
    onlyProtocol
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`context`|`MessageContext`|The context of the cross-chain call.|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`amount`|`uint256`|The amount of tokens to transfer.|
|`target`|`address`|The target contract to call.|
|`message`|`bytes`|The calldata to pass to the contract call.|


#### depositAndCall

Deposit native ZETA and call a user-specified contract on ZEVM.


```solidity
function depositAndCall(
    MessageContext calldata context,
    address target,
    bytes calldata message
)
    external
    payable
    nonReentrant
    onlyProtocol
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`context`|`MessageContext`|The context of the cross-chain call.|
|`target`|`address`|The target contract to call.|
|`message`|`bytes`|The calldata to pass to the contract call.|


#### executeRevert

Revert a user-specified contract on ZEVM.


```solidity
function executeRevert(
    address target,
    RevertContext calldata revertContext
)
    external
    nonReentrant
    onlyProtocol
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`target`|`address`|The target contract to call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|


#### depositAndRevert

Deposit foreign coins into ZRC20 and revert a user-specified contract on ZEVM.


```solidity
function depositAndRevert(
    address zrc20,
    uint256 amount,
    address target,
    RevertContext calldata revertContext
)
    external
    nonReentrant
    onlyProtocol
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`amount`|`uint256`|The amount of tokens to revert.|
|`target`|`address`|The target contract to call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|


#### depositAndRevert

Deposit native ZETA and revert a user-specified contract on ZEVM.


```solidity
function depositAndRevert(
    address target,
    RevertContext calldata revertContext
)
    external
    payable
    nonReentrant
    onlyProtocol
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`target`|`address`|The target contract to call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|


#### executeAbort

Call onAbort on a user-specified contract on ZEVM.
this function doesn't deposit the asset to the target contract. This operation is done directly by the protocol.
the assets are deposited to the target contract even if onAbort reverts.


```solidity
function executeAbort(
    address target,
    AbortContext calldata abortContext
)
    external
    nonReentrant
    onlyProtocol
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`target`|`address`|The target contract to call.|
|`abortContext`|`AbortContext`|Abort context to pass to onAbort.|


#### getMaxMessageSize

Returns the maximum message size.


```solidity
function getMaxMessageSize() external pure returns (uint256);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|The maximum message size.|


#### getMinGasLimit

Returns the minimum gas limit allowed.


```solidity
function getMinGasLimit() external pure returns (uint256);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|The minimum gas limit.|


#### getMaxRevertGasLimit

Returns the maximum revert gas limit allowed.


```solidity
function getMaxRevertGasLimit() external pure returns (uint256);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|The maximum revert gas limit.|




## INotSupportedMethods
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/Errors.sol)

Interface for contracts that with non supported methods.


### Errors
#### CallOnRevertNotSupported

```solidity
error CallOnRevertNotSupported();
```



## ERC20Custody
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/ERC20Custody.sol)

Holds the ERC20 tokens deposited on ZetaChain and includes functionality to call a contract.

*This contract does not call smart contracts directly, it passes through the Gateway contract.*


### State Variables
#### gateway
Gateway contract.


```solidity
IGatewayEVM public gateway;
```


#### whitelisted
Mapping of whitelisted tokens => true/false.


```solidity
mapping(address => bool) public whitelisted;
```


#### tssAddress
The address of the TSS (Threshold Signature Scheme) contract.


```solidity
address public tssAddress;
```


#### supportsLegacy
Used to flag if contract supports legacy methods (eg. deposit).


```solidity
bool public supportsLegacy;
```


#### PAUSER_ROLE
New role identifier for pauser role.


```solidity
bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
```


#### WITHDRAWER_ROLE
New role identifier for withdrawer role.


```solidity
bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
```


#### WHITELISTER_ROLE
New role identifier for whitelister role.


```solidity
bytes32 public constant WHITELISTER_ROLE = keccak256("WHITELISTER_ROLE");
```


### Functions
#### initialize

Initializer for ERC20Custody.

*Set admin as default admin and pauser, and tssAddress as tss role.*


```solidity
function initialize(address gateway_, address tssAddress_, address admin_) public initializer;
```

#### _authorizeUpgrade

*Authorizes the upgrade of the contract, sender must be owner.*


```solidity
function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newImplementation`|`address`|Address of the new implementation.|


#### pause

Pause contract.


```solidity
function pause() external onlyRole(PAUSER_ROLE);
```

#### unpause

Unpause contract.


```solidity
function unpause() external onlyRole(PAUSER_ROLE);
```

#### updateTSSAddress

Update tss address


```solidity
function updateTSSAddress(address newTSSAddress) external onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newTSSAddress`|`address`|new tss address|


#### setSupportsLegacy

Unpause contract.


```solidity
function setSupportsLegacy(bool _supportsLegacy) external onlyRole(DEFAULT_ADMIN_ROLE);
```

#### whitelist

Whitelist ERC20 token.


```solidity
function whitelist(address token) external onlyRole(WHITELISTER_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|address of ERC20 token|


#### unwhitelist

Unwhitelist ERC20 token.


```solidity
function unwhitelist(address token) external onlyRole(WHITELISTER_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|address of ERC20 token|


#### withdraw

Withdraw directly transfers the tokens to the destination address without contract call.

*This function can only be called by the TSS address.*


```solidity
function withdraw(
    address to,
    address token,
    uint256 amount
)
    external
    nonReentrant
    onlyRole(WITHDRAWER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|Destination address for the tokens.|
|`token`|`address`|Address of the ERC20 token.|
|`amount`|`uint256`|Amount of tokens to withdraw.|


#### withdrawAndCall

WithdrawAndCall transfers tokens to Gateway and call a contract through the Gateway.

*This function can only be called by the TSS address.*


```solidity
function withdrawAndCall(
    MessageContext calldata messageContext,
    address to,
    address token,
    uint256 amount,
    bytes calldata data
)
    public
    nonReentrant
    onlyRole(WITHDRAWER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`messageContext`|`MessageContext`|Message context containing sender.|
|`to`|`address`|Address of the contract to call.|
|`token`|`address`|Address of the ERC20 token.|
|`amount`|`uint256`|Amount of tokens to withdraw.|
|`data`|`bytes`|Calldata to pass to the contract call.|


#### withdrawAndRevert

WithdrawAndRevert transfers tokens to Gateway and call a contract with a revert functionality through
the Gateway.

*This function can only be called by the TSS address.*


```solidity
function withdrawAndRevert(
    address to,
    address token,
    uint256 amount,
    bytes calldata data,
    RevertContext calldata revertContext
)
    public
    nonReentrant
    onlyRole(WITHDRAWER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|Address of the contract to call.|
|`token`|`address`|Address of the ERC20 token.|
|`amount`|`uint256`|Amount of tokens to withdraw.|
|`data`|`bytes`|Calldata to pass to the contract call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|


#### deposit

Deposits asset to custody and pay fee in zeta erc20.

**Note:**
deprecated: This method is deprecated.


```solidity
function deposit(
    bytes calldata recipient,
    IERC20 asset,
    uint256 amount,
    bytes calldata message
)
    external
    nonReentrant
    whenNotPaused;
```



## IERC20Custody
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/interfaces/IERC20Custody.sol)


### Functions
#### whitelisted

Mapping of whitelisted tokens => true/false.


```solidity
function whitelisted(address token) external view returns (bool);
```

#### withdraw

Withdraw directly transfers the tokens to the destination address without contract call.

*This function can only be called by the TSS address.*


```solidity
function withdraw(address token, address to, uint256 amount) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|Address of the ERC20 token.|
|`to`|`address`|Destination address for the tokens.|
|`amount`|`uint256`|Amount of tokens to withdraw.|


#### withdrawAndCall

WithdrawAndCall transfers tokens to Gateway and call a contract through the Gateway.

*This function can only be called by the TSS address.*


```solidity
function withdrawAndCall(
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
|`messageContext`|`MessageContext`|Message context containing sender.|
|`token`|`address`|Address of the ERC20 token.|
|`to`|`address`|Address of the contract to call.|
|`amount`|`uint256`|Amount of tokens to withdraw.|
|`data`|`bytes`|Calldata to pass to the contract call.|


#### withdrawAndRevert

WithdrawAndRevert transfers tokens to Gateway and call a contract with a revert functionality through
the Gateway.

*This function can only be called by the TSS address.*


```solidity
function withdrawAndRevert(
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
|`token`|`address`|Address of the ERC20 token.|
|`to`|`address`|Address of the contract to call.|
|`amount`|`uint256`|Amount of tokens to withdraw.|
|`data`|`bytes`|Calldata to pass to the contract call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|




## IERC20CustodyErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/interfaces/IERC20Custody.sol)

Interface for the errors used in the ERC20 custody contract.


### Errors
#### ZeroAddress
Error for zero address input.


```solidity
error ZeroAddress();
```

#### NotWhitelisted
Error for not whitelisted ERC20 token


```solidity
error NotWhitelisted();
```

#### LegacyMethodsNotSupported
Error for calling not supported legacy methods.


```solidity
error LegacyMethodsNotSupported();
```



## IERC20CustodyEvents
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/interfaces/IERC20Custody.sol)

Interface for the events emitted by the ERC20 custody contract.


### Events
#### Withdrawn
Emitted when tokens are withdrawn.


```solidity
event Withdrawn(address indexed to, address indexed token, uint256 amount);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|The address receiving the tokens.|
|`token`|`address`|The address of the ERC20 token.|
|`amount`|`uint256`|The amount of tokens withdrawn.|

#### WithdrawnAndCalled
Emitted when tokens are withdrawn and a contract call is made.


```solidity
event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|The address receiving the tokens.|
|`token`|`address`|The address of the ERC20 token.|
|`amount`|`uint256`|The amount of tokens withdrawn.|
|`data`|`bytes`|The calldata passed to the contract call.|

#### WithdrawnAndReverted
Emitted when tokens are withdrawn and a revertable contract call is made.


```solidity
event WithdrawnAndReverted(
    address indexed to, address indexed token, uint256 amount, bytes data, RevertContext revertContext
);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|The address receiving the tokens.|
|`token`|`address`|The address of the ERC20 token.|
|`amount`|`uint256`|The amount of tokens withdrawn.|
|`data`|`bytes`|The calldata passed to the contract call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|

#### Whitelisted
Emitted when ERC20 token is whitelisted


```solidity
event Whitelisted(address indexed token);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|address of ERC20 token.|

#### Unwhitelisted
Emitted when ERC20 token is unwhitelisted


```solidity
event Unwhitelisted(address indexed token);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|address of ERC20 token.|

#### Deposited
Emitted in legacy deposit method.


```solidity
event Deposited(bytes recipient, IERC20 indexed asset, uint256 amount, bytes message);
```

#### UpdatedCustodyTSSAddress
Emitted when tss address is updated


```solidity
event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`oldTSSAddress`|`address`|old tss address|
|`newTSSAddress`|`address`|new tss address|



## Callable
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/interfaces/IGatewayEVM.sol)

Interface implemented by contracts receiving authenticated calls.


### Functions
#### onCall


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable returns (bytes memory);
```



## IGatewayEVM
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/interfaces/IGatewayEVM.sol)

Interface for the GatewayEVM contract.


### Functions
#### executeWithERC20

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


#### executeRevert

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


#### execute

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


#### revertWithERC20

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


#### deposit

Deposits ETH to the TSS address.


```solidity
function deposit(address receiver, RevertOptions calldata revertOptions) external payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|Address of the receiver.|
|`revertOptions`|`RevertOptions`|Revert options.|


#### deposit

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


#### deposit

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


#### depositAndCall

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


#### depositAndCall

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


#### depositAndCall

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


#### call

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




## IGatewayEVMErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/interfaces/IGatewayEVM.sol)

Interface for the errors used in the GatewayEVM contract.


### Errors
#### ExecutionFailed
Error for failed execution.


```solidity
error ExecutionFailed();
```

#### DepositFailed
Error for failed deposit.


```solidity
error DepositFailed();
```

#### InsufficientEVMAmount
Error for insufficient token amount.


```solidity
error InsufficientEVMAmount();
```

#### ZeroAddress
Error for zero address input.


```solidity
error ZeroAddress();
```

#### ApprovalFailed
Error for failed token approval.


```solidity
error ApprovalFailed(address token, address spender);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|The address of the token for which approval failed.|
|`spender`|`address`|The address that was supposed to be approved to spend the tokens.|

#### CustodyInitialized
Error for already initialized custody.


```solidity
error CustodyInitialized();
```

#### ConnectorInitialized
Error for already initialized connector.


```solidity
error ConnectorInitialized();
```

#### NotWhitelistedInCustody
Error when trying to transfer not whitelisted token to custody.


```solidity
error NotWhitelistedInCustody(address token);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|The address of the token that is not whitelisted in custody.|

#### NotAllowedToCallOnCall
Error when trying to call onCall method using arbitrary call.


```solidity
error NotAllowedToCallOnCall();
```

#### NotAllowedToCallOnRevert
Error when trying to call onRevert method using arbitrary call.


```solidity
error NotAllowedToCallOnRevert();
```

#### PayloadSizeExceeded
Error indicating payload size exceeded in external functions.


```solidity
error PayloadSizeExceeded(uint256 provided, uint256 maximum);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`provided`|`uint256`|The size of the payload that was provided.|
|`maximum`|`uint256`|The maximum allowed payload size.|

#### FeeTransferFailed
Error thrown when fee transfer to TSS address fails.

*This error occurs when the low-level call to transfer fees fails.*


```solidity
error FeeTransferFailed();
```

#### InsufficientFee
Error thrown when insufficient fee is provided for additional actions.


```solidity
error InsufficientFee(uint256 required, uint256 provided);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`required`|`uint256`|The fee amount required for the action.|
|`provided`|`uint256`|The fee amount actually provided by the caller.|

#### ExcessETHProvided
Error thrown when excess ETH is sent for non-ETH operations.


```solidity
error ExcessETHProvided(uint256 required, uint256 provided);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`required`|`uint256`|The fee amount required for the action.|
|`provided`|`uint256`|The ETH amount actually provided by the caller.|

#### AdditionalActionDisabled
Error thrown when additional action functionality is disabled (fee set to 0).


```solidity
error AdditionalActionDisabled();
```

#### IncorrectValueProvided
Error thrown when msg.value doesn't match expected amount + fee.


```solidity
error IncorrectValueProvided(uint256 expected, uint256 provided);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`expected`|`uint256`|The expected value (amount + fee).|
|`provided`|`uint256`|The actual msg.value provided.|



## IGatewayEVMEvents
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/interfaces/IGatewayEVM.sol)

Interface for the events emitted by the GatewayEVM contract.


### Events
#### Executed
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

#### Reverted
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

#### ExecutedWithERC20
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

#### Deposited
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

#### DepositedAndCalled
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

#### Called
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

#### UpdatedGatewayTSSAddress
Emitted when tss address is updated


```solidity
event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`oldTSSAddress`|`address`|old tss address|
|`newTSSAddress`|`address`|new tss address|

#### UpdatedAdditionalActionFee
Emitted when additional action fee for multi-deposits is updated


```solidity
event UpdatedAdditionalActionFee(uint256 oldFeeWei, uint256 newFeeWei);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`oldFeeWei`|`uint256`|old fee in wei|
|`newFeeWei`|`uint256`|new fee in wei|



## MessageContext
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/interfaces/IGatewayEVM.sol)

Message context passed to execute function.


```solidity
struct MessageContext {
    address sender;
}
```

**Properties**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|Sender from omnichain contract.|



## IRegistry
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/interfaces/IRegistry.sol)


### Structs
#### ChainMetadataEntry
Structure for metadata entries used during bootstrapping


```solidity
struct ChainMetadataEntry {
    uint256 chainId;
    string key;
    bytes value;
}
```

#### ContractConfigEntry
Structure for contract configuration entries used during bootstrapping


```solidity
struct ContractConfigEntry {
    uint256 chainId;
    string contractType;
    string key;
    bytes value;
}
```



## IZetaConnectorEvents
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/interfaces/IZetaConnector.sol)

Interface for the events emitted by the ZetaConnector contracts.


### Events
#### Withdrawn
Emitted when tokens are withdrawn.


```solidity
event Withdrawn(address indexed to, uint256 amount);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|The address to which the tokens are withdrawn.|
|`amount`|`uint256`|The amount of tokens withdrawn.|

#### WithdrawnAndCalled
Emitted when tokens are withdrawn and a contract is called.


```solidity
event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|The address to which the tokens are withdrawn.|
|`amount`|`uint256`|The amount of tokens withdrawn.|
|`data`|`bytes`|The calldata passed to the contract call.|

#### WithdrawnAndReverted
Emitted when tokens are withdrawn and a contract is called with a revert callback.


```solidity
event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, RevertContext revertContext);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|The address to which the tokens are withdrawn.|
|`amount`|`uint256`|The amount of tokens withdrawn.|
|`data`|`bytes`|The calldata passed to the contract call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|

#### UpdatedZetaConnectorTSSAddress
Emitted when tss address is updated


```solidity
event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`oldTSSAddress`|`address`|old tss address|
|`newTSSAddress`|`address`|new tss address|



## IZetaNonEthNew
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/interfaces/IZetaNonEthNew.sol)

IZetaNonEthNew is a mintable / burnable version of IERC20.


### Functions
#### burnFrom

Burns the specified amount of tokens from the specified account.

*Emits a {Transfer} event with `to` set to the zero address.*


```solidity
function burnFrom(address account, uint256 amount) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`account`|`address`|The address of the account from which tokens will be burned.|
|`amount`|`uint256`|The amount of tokens to burn.|


#### mint

Mints the specified amount of tokens to the specified account.

*Emits a {Transfer} event with `from` set to the zero address.*


```solidity
function mint(address mintee, uint256 value, bytes32 internalSendHash) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`mintee`|`address`|The address of the account to which tokens will be minted.|
|`value`|`uint256`|The amount of tokens to mint.|
|`internalSendHash`|`bytes32`|A hash used for internal tracking of the minting transaction.|




## ConnectorErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/legacy/ConnectorErrors.sol)

*Interface with connector custom errors*


### Errors
#### CallerIsNotPauser

```solidity
error CallerIsNotPauser(address caller);
```

#### CallerIsNotTss

```solidity
error CallerIsNotTss(address caller);
```

#### CallerIsNotTssUpdater

```solidity
error CallerIsNotTssUpdater(address caller);
```

#### CallerIsNotTssOrUpdater

```solidity
error CallerIsNotTssOrUpdater(address caller);
```

#### ZetaTransferError

```solidity
error ZetaTransferError();
```

#### ExceedsMaxSupply

```solidity
error ExceedsMaxSupply(uint256 maxSupply);
```



## IZetaNonEthInterface
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/legacy/IZetaNonEthInterface.sol)

*IZetaNonEthInterface.sol is a mintable / burnable version of IERC20*


### Functions
#### burnFrom


```solidity
function burnFrom(address account, uint256 amount) external;
```

#### mint


```solidity
function mint(address mintee, uint256 value, bytes32 internalSendHash) external;
```



## ZetaConnectorBase
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/legacy/ZetaConnector.base.sol)

*Main abstraction of ZetaConnector.
This contract manages interactions between TSS and different chains.
There's an instance of this contract on each chain supported by ZetaChain.*


### State Variables
#### zetaToken

```solidity
address public immutable zetaToken;
```


#### pauserAddress
*Multisig contract to pause incoming transactions.
The responsibility of pausing outgoing transactions is left to the protocol for more flexibility.*


```solidity
address public pauserAddress;
```


#### tssAddress
*Collectively held by ZetaChain validators.*


```solidity
address public tssAddress;
```


#### tssAddressUpdater
*This address will start pointing to a multisig contract, then it will become the TSS address itself.*


```solidity
address public tssAddressUpdater;
```


### Functions
#### constructor

*Constructor requires initial addresses.
zetaToken address is the only immutable one, while others can be updated.*


```solidity
constructor(address zetaToken_, address tssAddress_, address tssAddressUpdater_, address pauserAddress_);
```

#### onlyPauser

*Modifier to restrict actions to pauser address.*


```solidity
modifier onlyPauser();
```

#### onlyTssAddress

*Modifier to restrict actions to TSS address.*


```solidity
modifier onlyTssAddress();
```

#### onlyTssUpdater

*Modifier to restrict actions to TSS updater address.*


```solidity
modifier onlyTssUpdater();
```

#### updatePauserAddress

*Update the pauser address. The only address allowed to do that is the current pauser.*


```solidity
function updatePauserAddress(address pauserAddress_) external onlyPauser;
```

#### updateTssAddress

*Update the TSS address. The address can be updated by the TSS updater or the TSS address itself.*


```solidity
function updateTssAddress(address tssAddress_) external;
```

#### renounceTssAddressUpdater

*Changes the ownership of tssAddressUpdater to be the one held by the ZetaChain TSS Signer nodes.*


```solidity
function renounceTssAddressUpdater() external onlyTssUpdater;
```

#### pause

*Pause the input (send) transactions.*


```solidity
function pause() external onlyPauser;
```

#### unpause

*Unpause the contract to allow transactions again.*


```solidity
function unpause() external onlyPauser;
```

#### send

*Entrypoint to send data and value through ZetaChain.*


```solidity
function send(ZetaInterfaces.SendInput calldata input) external virtual;
```

#### onReceive

*Handler to receive data from other chain.
This method can be called only by TSS. Access validation is in implementation.*


```solidity
function onReceive(
    bytes calldata zetaTxSenderAddress,
    uint256 sourceChainId,
    address destinationAddress,
    uint256 zetaValue,
    bytes calldata message,
    bytes32 internalSendHash
)
    external
    virtual;
```

#### onRevert

*Handler to receive errors from other chain.
This method can be called only by TSS. Access validation is in implementation.*


```solidity
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
    virtual;
```

### Events
#### ZetaSent

```solidity
event ZetaSent(
    address sourceTxOriginAddress,
    address indexed zetaTxSenderAddress,
    uint256 indexed destinationChainId,
    bytes destinationAddress,
    uint256 zetaValueAndGas,
    uint256 destinationGasLimit,
    bytes message,
    bytes zetaParams
);
```

#### ZetaReceived

```solidity
event ZetaReceived(
    bytes zetaTxSenderAddress,
    uint256 indexed sourceChainId,
    address indexed destinationAddress,
    uint256 zetaValue,
    bytes message,
    bytes32 indexed internalSendHash
);
```

#### ZetaReverted

```solidity
event ZetaReverted(
    address zetaTxSenderAddress,
    uint256 sourceChainId,
    uint256 indexed destinationChainId,
    bytes destinationAddress,
    uint256 remainingZetaValue,
    bytes message,
    bytes32 indexed internalSendHash
);
```

#### TSSAddressUpdated

```solidity
event TSSAddressUpdated(address callerAddress, address newTssAddress);
```

#### TSSAddressUpdaterUpdated

```solidity
event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress);
```

#### PauserAddressUpdated

```solidity
event PauserAddressUpdated(address callerAddress, address newTssAddress);
```



## ZetaConnectorEth
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/legacy/ZetaConnector.eth.sol)

*ETH implementation of ZetaConnector.
This contract manages interactions between TSS and different chains.
This version is only for Ethereum network because in the other chains we mint and burn and in this one we lock and
unlock.*


### Functions
#### constructor


```solidity
constructor(
    address zetaToken_,
    address tssAddress_,
    address tssAddressUpdater_,
    address pauserAddress_
)
    ZetaConnectorBase(zetaToken_, tssAddress_, tssAddressUpdater_, pauserAddress_);
```

#### getLockedAmount


```solidity
function getLockedAmount() external view returns (uint256);
```

#### send

*Entrypoint to send data through ZetaChain
This call locks the token on the contract and emits an event with all the data needed by the protocol.*


```solidity
function send(ZetaInterfaces.SendInput calldata input) external override whenNotPaused;
```

#### onReceive

*Handler to receive data from other chain.
This method can be called only by TSS.
Transfers the Zeta tokens to destination and calls onZetaMessage if it's needed.*


```solidity
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
    onlyTssAddress;
```

#### onRevert

*Handler to receive errors from other chain.
This method can be called only by TSS.
Transfers the Zeta tokens to destination and calls onZetaRevert if it's needed.*


```solidity
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
    onlyTssAddress;
```



## ZetaConnectorNonEth
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/legacy/ZetaConnector.non-eth.sol)

*Non ETH implementation of ZetaConnector.
This contract manages interactions between TSS and different chains.
This version is for every chain but Etherum network because in the other chains we mint and burn and in Etherum we
lock and unlock*


### State Variables
#### maxSupply

```solidity
uint256 public maxSupply = 2 ** 256 - 1;
```


### Functions
#### constructor


```solidity
constructor(
    address zetaTokenAddress_,
    address tssAddress_,
    address tssAddressUpdater_,
    address pauserAddress_
)
    ZetaConnectorBase(zetaTokenAddress_, tssAddress_, tssAddressUpdater_, pauserAddress_);
```

#### getLockedAmount


```solidity
function getLockedAmount() external view returns (uint256);
```

#### setMaxSupply


```solidity
function setMaxSupply(uint256 maxSupply_) external onlyTssAddress;
```

#### send

*Entry point to send data to protocol
This call burn the token and emit an event with all the data needed by the protocol*


```solidity
function send(ZetaInterfaces.SendInput calldata input) external override whenNotPaused;
```

#### onReceive

*Handler to receive data from other chain.
This method can be called only by TSS.
Transfer the Zeta tokens to destination and calls onZetaMessage if it's needed.
To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain.*


```solidity
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
    onlyTssAddress;
```

#### onRevert

*Handler to receive errors from other chain.
This method can be called only by TSS.
Transfer the Zeta tokens to destination and calls onZetaRevert if it's needed.
To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain.*


```solidity
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
    onlyTssAddress;
```

### Events
#### MaxSupplyUpdated

```solidity
event MaxSupplyUpdated(address callerAddress, uint256 newMaxSupply);
```



## ZetaErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/legacy/ZetaErrors.sol)

*Common custom errors*


### Errors
#### CallerIsNotTss

```solidity
error CallerIsNotTss(address caller);
```

#### CallerIsNotConnector

```solidity
error CallerIsNotConnector(address caller);
```

#### CallerIsNotTssUpdater

```solidity
error CallerIsNotTssUpdater(address caller);
```

#### CallerIsNotTssOrUpdater

```solidity
error CallerIsNotTssOrUpdater(address caller);
```

#### InvalidAddress

```solidity
error InvalidAddress();
```

#### ZetaTransferError

```solidity
error ZetaTransferError();
```



## ZetaEth
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/legacy/ZetaEth.sol)

Ethereum is the origin and native chain of the ZETA token deployment (native)

*ZetaEth.sol is an implementation of OpenZeppelin's ERC20*


### Functions
#### constructor


```solidity
constructor(address creator, uint256 initialSupply);
```



## ZetaCommonErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/legacy/ZetaInterfaces.sol)


### Errors
#### InvalidAddress

```solidity
error InvalidAddress();
```



## ZetaConnector
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/legacy/ZetaInterfaces.sol)


### Functions
#### send

*Sending value and data cross-chain is as easy as calling connector.send(SendInput)*


```solidity
function send(ZetaInterfaces.SendInput calldata input) external;
```



## ZetaInterfaces
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/legacy/ZetaInterfaces.sol)


### Structs
#### SendInput
*Use SendInput to interact with the Connector: connector.send(SendInput)*


```solidity
struct SendInput {
    uint256 destinationChainId;
    bytes destinationAddress;
    uint256 destinationGasLimit;
    bytes message;
    uint256 zetaValueAndGas;
    bytes zetaParams;
}
```

#### ZetaMessage
*Our Connector calls onZetaMessage with this struct as argument*


```solidity
struct ZetaMessage {
    bytes zetaTxSenderAddress;
    uint256 sourceChainId;
    address destinationAddress;
    uint256 zetaValue;
    bytes message;
}
```

#### ZetaRevert
*Our Connector calls onZetaRevert with this struct as argument*


```solidity
struct ZetaRevert {
    address zetaTxSenderAddress;
    uint256 sourceChainId;
    bytes destinationAddress;
    uint256 destinationChainId;
    uint256 remainingZetaValue;
    bytes message;
}
```



## ZetaReceiver
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/legacy/ZetaInterfaces.sol)


### Functions
#### onZetaMessage

*onZetaMessage is called when a cross-chain message reaches a contract*


```solidity
function onZetaMessage(ZetaInterfaces.ZetaMessage calldata zetaMessage) external;
```

#### onZetaRevert

*onZetaRevert is called when a cross-chain message reverts.
It's useful to rollback to the original state*


```solidity
function onZetaRevert(ZetaInterfaces.ZetaRevert calldata zetaRevert) external;
```



## ZetaTokenConsumer
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/legacy/ZetaInterfaces.sol)

*ZetaTokenConsumer makes it easier to handle the following situations:
- Getting Zeta using native coin (to pay for destination gas while using `connector.send`)
- Getting Zeta using a token (to pay for destination gas while using `connector.send`)
- Getting native coin using Zeta (to return unused destination gas when `onZetaRevert` is executed)
- Getting a token using Zeta (to return unused destination gas when `onZetaRevert` is executed)*

*The interface can be implemented using different strategies, like UniswapV2, UniswapV3, etc*


### Functions
#### getZetaFromEth


```solidity
function getZetaFromEth(address destinationAddress, uint256 minAmountOut) external payable returns (uint256);
```

#### getZetaFromToken


```solidity
function getZetaFromToken(
    address destinationAddress,
    uint256 minAmountOut,
    address inputToken,
    uint256 inputTokenAmount
)
    external
    returns (uint256);
```

#### getEthFromZeta


```solidity
function getEthFromZeta(
    address destinationAddress,
    uint256 minAmountOut,
    uint256 zetaTokenAmount
)
    external
    returns (uint256);
```

#### getTokenFromZeta


```solidity
function getTokenFromZeta(
    address destinationAddress,
    uint256 minAmountOut,
    address outputToken,
    uint256 zetaTokenAmount
)
    external
    returns (uint256);
```

#### hasZetaLiquidity


```solidity
function hasZetaLiquidity() external view returns (bool);
```

### Events
#### EthExchangedForZeta

```solidity
event EthExchangedForZeta(uint256 amountIn, uint256 amountOut);
```

#### TokenExchangedForZeta

```solidity
event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut);
```

#### ZetaExchangedForEth

```solidity
event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut);
```

#### ZetaExchangedForToken

```solidity
event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut);
```



## ZetaNonEth
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/legacy/ZetaNonEth.sol)

On non-native (non-Ethereum) chains, ZETA tokens are minted and burned after the initial deployment on
Ethereum.


### State Variables
#### connectorAddress

```solidity
address public connectorAddress;
```


#### tssAddress
*Collectively held by Zeta blockchain validators*


```solidity
address public tssAddress;
```


#### tssAddressUpdater
*Initially a multi-sig, eventually held by Zeta blockchain validators (via renounceTssAddressUpdater)*


```solidity
address public tssAddressUpdater;
```


### Functions
#### constructor


```solidity
constructor(address tssAddress_, address tssAddressUpdater_) ERC20("Zeta", "ZETA");
```

#### updateTssAndConnectorAddresses


```solidity
function updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) external;
```

#### renounceTssAddressUpdater

*Sets tssAddressUpdater to be tssAddress*


```solidity
function renounceTssAddressUpdater() external;
```

#### mint


```solidity
function mint(address mintee, uint256 value, bytes32 internalSendHash) external override;
```

#### burnFrom


```solidity
function burnFrom(address account, uint256 amount) public override(IZetaNonEthInterface, ERC20Burnable);
```

### Events
#### Minted

```solidity
event Minted(address indexed mintee, uint256 amount, bytes32 indexed internalSendHash);
```

#### Burnt

```solidity
event Burnt(address indexed burnee, uint256 amount);
```

#### TSSAddressUpdated

```solidity
event TSSAddressUpdated(address callerAddress, address newTssAddress);
```

#### TSSAddressUpdaterUpdated

```solidity
event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress);
```

#### ConnectorAddressUpdated

```solidity
event ConnectorAddressUpdated(address callerAddress, address newConnectorAddress);
```



## GatewayEVMValidations
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/libraries/GatewayEVMValidations.sol)

Library containing validation functions for GatewayEVM contract.

*This library provides common validation logic used across GatewayEVM contract.*


### State Variables
#### MAX_PAYLOAD_SIZE
Maximum payload size constant.


```solidity
uint256 internal constant MAX_PAYLOAD_SIZE = 2880;
```


### Functions
#### validateNonZeroAddress

*Validates that an address is not zero.*


```solidity
function validateNonZeroAddress(address addr) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`addr`|`address`|The address to validate.|


#### validateAmount

*Validates that amount is not zero.*


```solidity
function validateAmount(uint256 amount) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amount`|`uint256`|The amount to validate.|


#### validatePayloadSize

*Validates payload size constraints.*


```solidity
function validatePayloadSize(uint256 payloadLength, uint256 revertMessageLength) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`payloadLength`|`uint256`|The length of the main payload.|
|`revertMessageLength`|`uint256`|The length of the revert message.|


#### validateGasLimitRevertOptions

*Validates on revert gas limit.*


```solidity
function validateGasLimitRevertOptions(RevertOptions calldata revertOptions) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`revertOptions`|`RevertOptions`|The revert options to validate.|


#### validateRevertMessageLength

*Validates on revert message length.*


```solidity
function validateRevertMessageLength(RevertOptions calldata revertOptions) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`revertOptions`|`RevertOptions`|The revert options to validate.|


#### validateCallRevertOptions

*Validates revert options for call operations (includes callOnRevert check).*


```solidity
function validateCallRevertOptions(RevertOptions calldata revertOptions) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`revertOptions`|`RevertOptions`|The revert options to validate.|


#### validateDepositParams

*Validates standard deposit parameters.*


```solidity
function validateDepositParams(address receiver, uint256 amount, RevertOptions calldata revertOptions) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|The receiver address.|
|`amount`|`uint256`|The amount to deposit.|
|`revertOptions`|`RevertOptions`|The revert options.|


#### validateDepositAndCallParams

*Validates deposit and call parameters.*


```solidity
function validateDepositAndCallParams(
    address receiver,
    uint256 amount,
    bytes calldata payload,
    RevertOptions calldata revertOptions
)
    internal
    pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|The receiver address.|
|`amount`|`uint256`|The amount to deposit.|
|`payload`|`bytes`|The payload to send.|
|`revertOptions`|`RevertOptions`|The revert options.|


#### validateCallParams

*Validates call parameters.*


```solidity
function validateCallParams(
    address receiver,
    bytes calldata payload,
    RevertOptions calldata revertOptions
)
    internal
    pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`address`|The receiver address.|
|`payload`|`bytes`|The payload to send.|
|`revertOptions`|`RevertOptions`|The revert options.|




## Registry
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/Registry.sol)

Satellite registry contract for connected chains, receiving updates from CoreRegistry.

*This contract is deployed on every connected chain and maintains a synchronized view of the registry.*


### State Variables
#### GATEWAY_ROLE
Identifier for the gateway role


```solidity
bytes32 public constant GATEWAY_ROLE = keccak256("GATEWAY_ROLE");
```


#### gatewayEVM
GatewayEVM contract that will call this contract with messages from CoreRegistry


```solidity
IGatewayEVM public gatewayEVM;
```


#### coreRegistry
Represents the address of the CoreRegistry contract on the ZetaChain


```solidity
address public coreRegistry;
```


### Functions
#### onlyRegistry

Restricts function calls to only be made by this contract itself

*Only registry address allowed modifier.*

*This is used to ensure functions receiving cross-chain messages can only be called through
the onCall function using a self-call pattern, preventing direct external calls to these functions*


```solidity
modifier onlyRegistry();
```

#### initialize

Initialize the Registry contract


```solidity
function initialize(
    address admin_,
    address registryManager_,
    address gatewayEVM_,
    address coreRegistry_
)
    public
    initializer;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`admin_`|`address`|Address with DEFAULT_ADMIN_ROLE, authorized for upgrades and pausing actions|
|`registryManager_`|`address`|Address with REGISTRY_MANAGER_ROLE, authorized for all registry write actions.|
|`gatewayEVM_`|`address`|Address of the GatewayEVM contract for cross-chain messaging|
|`coreRegistry_`|`address`|Address of the CoreRegistry contract deployed on ZetaChain|


#### onCall

onCall is called by the GatewayEVM when a cross-chain message is received


```solidity
function onCall(
    MessageContext calldata context,
    bytes calldata data
)
    external
    onlyRole(GATEWAY_ROLE)
    whenNotPaused
    returns (bytes memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`context`|`MessageContext`|Information about the cross-chain message|
|`data`|`bytes`|The encoded function call to execute|


#### changeChainStatus

Changes status of the chain to activated/deactivated

*Only callable through onCall from CoreRegistry*


```solidity
function changeChainStatus(
    uint256 chainId,
    address gasZRC20,
    bytes calldata registry,
    bool activation
)
    external
    onlyRegistry
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain being activated/deactivated.|
|`gasZRC20`|`address`|The address of the ZRC20 token that represents gas token for the chain.|
|`registry`|`bytes`|Address of the Registry contract on the connected chain.|
|`activation`|`bool`|Whether activate or deactivate the chain|


#### updateChainMetadata

Updates chain metadata, only for the active chains

*Only callable through onCall from CoreRegistry*


```solidity
function updateChainMetadata(
    uint256 chainId,
    string calldata key,
    bytes calldata value
)
    external
    onlyRegistry
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain|
|`key`|`string`|The metadata key to update|
|`value`|`bytes`|The new value for the metadata|


#### registerContract

Registers a new contract address for a specific chain

*Only callable through onCall from CoreRegistry*


```solidity
function registerContract(
    uint256 chainId,
    string calldata contractType,
    bytes calldata addressBytes
)
    external
    onlyRegistry
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`|The type of the contract (e.g., "connector", "gateway")|
|`addressBytes`|`bytes`|The address of the contract|


#### updateContractConfiguration

Updates contract configuration

*Only callable through onCall from CoreRegistry*


```solidity
function updateContractConfiguration(
    uint256 chainId,
    string calldata contractType,
    string calldata key,
    bytes calldata value
)
    external
    onlyRegistry
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`|The type of the contract|
|`key`|`string`|The configuration key to update|
|`value`|`bytes`|The new value for the configuration|


#### setContractActive

Sets a contract's active status

*Only callable through onCall from CoreRegistry*


```solidity
function setContractActive(uint256 chainId, string calldata contractType, bool active) external onlyRegistry;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`|The type of the contract|
|`active`|`bool`|Whether the contract should be active|


#### registerZRC20Token

Registers a new ZRC20 token in the registry

*Only callable through onCall from CoreRegistry*


```solidity
function registerZRC20Token(
    address address_,
    string calldata symbol,
    uint256 originChainId,
    bytes calldata originAddress,
    string calldata coinType,
    uint8 decimals
)
    external
    onlyRegistry
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token on ZetaChain|
|`symbol`|`string`|The symbol of the token|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain|
|`coinType`|`string`|The type of the original coin|
|`decimals`|`uint8`|The number of decimals the token uses|


#### setZRC20TokenActive

Updates ZRC20 token active status

*Only callable through onCall from CoreRegistry*


```solidity
function setZRC20TokenActive(address address_, bool active) external onlyRegistry whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token|
|`active`|`bool`|Whether the token should be active|


#### bootstrapChains

Bootstrap the registry with chain data

*This function can only be called by an address with the REGISTRY_MANAGER_ROLE.*


```solidity
function bootstrapChains(
    ChainInfoDTO[] calldata chains,
    ChainMetadataEntry[] calldata metadataEntries
)
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chains`|`ChainInfoDTO[]`|Array of chain data structures to bootstrap|
|`metadataEntries`|`ChainMetadataEntry[]`|Array of chain metadata entries|


#### bootstrapContracts

Bootstrap the registry with contract data

*This function can only be called by an address with the REGISTRY_MANAGER_ROLE.*


```solidity
function bootstrapContracts(
    ContractInfoDTO[] calldata contracts,
    ContractConfigEntry[] calldata configEntries
)
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`contracts`|`ContractInfoDTO[]`|Array of contract data structures to bootstrap|
|`configEntries`|`ContractConfigEntry[]`|Array of contract configuration entries|


#### bootstrapZRC20Tokens

Bootstrap the registry with ZRC20 token data

*This function can only be called by an address with the REGISTRY_MANAGER_ROLE.*


```solidity
function bootstrapZRC20Tokens(ZRC20Info[] calldata tokens) external onlyRole(REGISTRY_MANAGER_ROLE) whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`tokens`|`ZRC20Info[]`|Array of ZRC20 token data structures to bootstrap|




## ZetaConnectorBase
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/ZetaConnectorBase.sol)

Abstract base contract for ZetaConnector.

*This contract implements basic functionality for handling tokens and interacting with the Gateway contract.*


### State Variables
#### gateway
The Gateway contract used for executing cross-chain calls.


```solidity
IGatewayEVM public gateway;
```


#### zetaToken
The address of the Zeta token.


```solidity
address public zetaToken;
```


#### tssAddress
The address of the TSS (Threshold Signature Scheme) contract.


```solidity
address public tssAddress;
```


#### WITHDRAWER_ROLE
New role identifier for withdrawer role.


```solidity
bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
```


#### PAUSER_ROLE
New role identifier for pauser role.


```solidity
bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
```


#### TSS_ROLE
New role identifier for tss role.


```solidity
bytes32 public constant TSS_ROLE = keccak256("TSS_ROLE");
```


### Functions
#### constructor

**Note:**
oz-upgrades-unsafe-allow: constructor


```solidity
constructor();
```

#### initialize

Initializer for ZetaConnectors.

*Set admin as default admin and pauser, and tssAddress as tss role.*


```solidity
function initialize(
    address gateway_,
    address zetaToken_,
    address tssAddress_,
    address admin_
)
    public
    virtual
    onlyInitializing;
```

#### _authorizeUpgrade

*Authorizes the upgrade of the contract, sender must be owner.*


```solidity
function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newImplementation`|`address`|Address of the new implementation.|


#### updateTSSAddress

Update tss address


```solidity
function updateTSSAddress(address newTSSAddress) external onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newTSSAddress`|`address`|new tss address|


#### pause

Pause contract.


```solidity
function pause() external onlyRole(PAUSER_ROLE);
```

#### unpause

Unpause contract.


```solidity
function unpause() external onlyRole(PAUSER_ROLE);
```

#### deposit

Handle received tokens.


```solidity
function deposit(uint256 amount) external virtual;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amount`|`uint256`|The amount of tokens received.|


### Errors
#### ZeroAddress
Error indicating that a zero address was provided.


```solidity
error ZeroAddress();
```



## ZetaConnectorNative
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/ZetaConnectorNative.sol)

Implementation of ZetaConnectorBase for native token handling.

*This contract directly transfers Zeta tokens and interacts with the Gateway contract.*


### Functions
#### initialize

Initializer for ZetaConnectorNative.


```solidity
function initialize(
    address gateway_,
    address zetaToken_,
    address tssAddress_,
    address admin_
)
    public
    override
    initializer;
```

#### withdraw

Withdraw tokens to a specified address.

*This function can only be called by the TSS address.*


```solidity
function withdraw(address to, uint256 amount) external nonReentrant onlyRole(WITHDRAWER_ROLE) whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|The address to withdraw tokens to.|
|`amount`|`uint256`|The amount of tokens to withdraw.|


#### withdrawAndCall

Withdraw tokens and call a contract through Gateway.

*This function can only be called by the TSS address.*


```solidity
function withdrawAndCall(
    MessageContext calldata messageContext,
    address to,
    uint256 amount,
    bytes calldata data
)
    external
    nonReentrant
    onlyRole(WITHDRAWER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`messageContext`|`MessageContext`|Message context containing sender.|
|`to`|`address`|The address to withdraw tokens to.|
|`amount`|`uint256`|The amount of tokens to withdraw.|
|`data`|`bytes`|The calldata to pass to the contract call.|


#### withdrawAndRevert

Withdraw tokens and call a contract with a revert callback through Gateway.

*This function can only be called by the TSS address.*


```solidity
function withdrawAndRevert(
    address to,
    uint256 amount,
    bytes calldata data,
    RevertContext calldata revertContext
)
    external
    nonReentrant
    onlyRole(WITHDRAWER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|The address to withdraw tokens to.|
|`amount`|`uint256`|The amount of tokens to withdraw.|
|`data`|`bytes`|The calldata to pass to the contract call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|


#### deposit

Handle received tokens.


```solidity
function deposit(uint256 amount) external override whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amount`|`uint256`|The amount of tokens received.|




## ZetaConnectorNonNative
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/evm/ZetaConnectorNonNative.sol)

Implementation of ZetaConnectorBase for non-native token handling.

*This contract mints and burns Zeta tokens and interacts with the Gateway contract.*


### State Variables
#### maxSupply
Max supply for minting.


```solidity
uint256 public maxSupply;
```


### Functions
#### initialize

Initializer for ZetaConnectorNonNative.


```solidity
function initialize(
    address gateway_,
    address zetaToken_,
    address tssAddress_,
    address admin_
)
    public
    override
    initializer;
```

#### setMaxSupply

Set max supply for minting.

*This function can only be called by the TSS address.*


```solidity
function setMaxSupply(uint256 maxSupply_) external onlyRole(TSS_ROLE) whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`maxSupply_`|`uint256`|New max supply.|


#### withdraw

Withdraw tokens to a specified address.

*This function can only be called by the TSS address.*


```solidity
function withdraw(
    address to,
    uint256 amount,
    bytes32 internalSendHash
)
    external
    nonReentrant
    onlyRole(WITHDRAWER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|The address to withdraw tokens to.|
|`amount`|`uint256`|The amount of tokens to withdraw.|
|`internalSendHash`|`bytes32`|A hash used for internal tracking of the transaction.|


#### withdrawAndCall

Withdraw tokens and call a contract through Gateway.

*This function can only be called by the TSS address, and mints if supply is not reached.*


```solidity
function withdrawAndCall(
    MessageContext calldata messageContext,
    address to,
    uint256 amount,
    bytes calldata data,
    bytes32 internalSendHash
)
    external
    nonReentrant
    onlyRole(WITHDRAWER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`messageContext`|`MessageContext`|Message context containing sender.|
|`to`|`address`|The address to withdraw tokens to.|
|`amount`|`uint256`|The amount of tokens to withdraw.|
|`data`|`bytes`|The calldata to pass to the contract call.|
|`internalSendHash`|`bytes32`|A hash used for internal tracking of the transaction.|


#### withdrawAndRevert

Withdraw tokens and call a contract with a revert callback through Gateway.

*This function can only be called by the TSS address, and mints if supply is not reached.*


```solidity
function withdrawAndRevert(
    address to,
    uint256 amount,
    bytes calldata data,
    bytes32 internalSendHash,
    RevertContext calldata revertContext
)
    external
    nonReentrant
    onlyRole(WITHDRAWER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`|The address to withdraw tokens to.|
|`amount`|`uint256`|The amount of tokens to withdraw.|
|`data`|`bytes`|The calldata to pass to the contract call.|
|`internalSendHash`|`bytes32`|A hash used for internal tracking of the transaction.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|


#### deposit

Handle received tokens and burn them.


```solidity
function deposit(uint256 amount) external override whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amount`|`uint256`|The amount of tokens received.|


#### _mintTo

*mints to provided account and checks if totalSupply will be exceeded*


```solidity
function _mintTo(address to, uint256 amount, bytes32 internalSendHash) private;
```

### Events
#### MaxSupplyUpdated
Event triggered when max supply is updated.


```solidity
event MaxSupplyUpdated(uint256 maxSupply);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`maxSupply`|`uint256`|New max supply.|

### Errors
#### ExceedsMaxSupply

```solidity
error ExceedsMaxSupply();
```



## BaseRegistry
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/helpers/BaseRegistry.sol)


### State Variables
#### PAUSER_ROLE
New role identifier for pauser role.


```solidity
bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
```


#### REGISTRY_MANAGER_ROLE
New role identifier for registry manager role.


```solidity
bytes32 public constant REGISTRY_MANAGER_ROLE = keccak256("REGISTRY_MANAGER_ROLE");
```


#### admin
Address with DEFAULT_ADMIN_ROLE, authorized for upgrades and pausing actions.


```solidity
address public admin;
```


#### registryManager
Address with REGISTRY_MANAGER_ROLE, authorized for all registry write actions.


```solidity
address public registryManager;
```


#### _activeChains
Active chains in the registry.


```solidity
uint256[] internal _activeChains;
```


#### _allChains
Array of all chain IDs in the registry (active and inactive).


```solidity
uint256[] internal _allChains;
```


#### _allContracts
Array to store all contracts as chainId and contractType pairs.


```solidity
ContractIdentifier[] internal _allContracts;
```


#### _allZRC20Addresses
Array of all ZRC20 token addresses.


```solidity
address[] internal _allZRC20Addresses;
```


#### _chains
Maps chain IDs to their information.


```solidity
mapping(uint256 => ChainInfo) internal _chains;
```


#### _contracts
Maps chain ID -> contract type -> ContractInfo


```solidity
mapping(uint256 => mapping(string => ContractInfo)) internal _contracts;
```


#### _zrc20Tokens
Maps ZRC20 token address to their information


```solidity
mapping(address => ZRC20Info) internal _zrc20Tokens;
```


#### _zrc20SymbolToAddress
Maps token symbol to ZRC20 address.


```solidity
mapping(string => address) internal _zrc20SymbolToAddress;
```


#### _originAssetToZRC20
Maps origin chain ID and origin address to ZRC20 token address.


```solidity
mapping(uint256 => mapping(bytes => address)) internal _originAssetToZRC20;
```


### Functions
#### constructor

**Note:**
oz-upgrades-unsafe-allow: constructor


```solidity
constructor();
```

#### _authorizeUpgrade

*Authorizes the upgrade of the contract, sender must be admin.*


```solidity
function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newImplementation`|`address`|Address of the new implementation,|


#### pause

Pause contract.


```solidity
function pause() external onlyRole(PAUSER_ROLE);
```

#### unpause

Unpause contract.


```solidity
function unpause() external onlyRole(DEFAULT_ADMIN_ROLE);
```

#### changeAdmin

Changes the admin address and transfers DEFAULT_ADMIN_ROLE and PAUSER_ROLE.

*Only callable by current admin.*


```solidity
function changeAdmin(address newAdmin) external onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newAdmin`|`address`|The address of the new admin.|


#### changeRegistryManager

Changes the registry manager address and transfers REGISTRY_MANAGER_ROLE and PAUSER_ROLE.

*Only callable by admin.*


```solidity
function changeRegistryManager(address newRegistryManager) external onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newRegistryManager`|`address`|The address of the new registry manager.|


#### _changeChainStatus

Changes status of the chain to activated/deactivated.


```solidity
function _changeChainStatus(uint256 chainId, address gasZRC20, bytes calldata registry, bool activation) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain to activate.|
|`gasZRC20`|`address`|The address of the ZRC20 token that represents gas token for the chain.|
|`registry`|`bytes`|Address of the Registry contract on the connected chain.|
|`activation`|`bool`|Whether activate or deactivate the chain|


#### _updateChainMetadata

Updates chain metadata, only for the active chains.


```solidity
function _updateChainMetadata(uint256 chainId, string calldata key, bytes calldata value) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|
|`key`|`string`|The metadata key to update.|
|`value`|`bytes`|The new value for the metadata.|


#### _registerContract

Registers a new contract address for a specific chain.


```solidity
function _registerContract(uint256 chainId, string calldata contractType, bytes calldata addressBytes) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract (e.g., "connector", "gateway").|
|`addressBytes`|`bytes`|The bytes representation of the non-EVM address.|


#### _updateContractConfiguration

Updates contract configuration.


```solidity
function _updateContractConfiguration(
    uint256 chainId,
    string calldata contractType,
    string calldata key,
    bytes calldata value
)
    internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`key`|`string`|The configuration key to update.|
|`value`|`bytes`|The new value for the configuration.|


#### _setContractActive

Sets a contract's active status


```solidity
function _setContractActive(uint256 chainId, string calldata contractType, bool active) internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`active`|`bool`|Whether the contract should be active.|


#### _registerZRC20Token

Registers a new ZRC20 token in the registry.


```solidity
function _registerZRC20Token(
    address address_,
    string calldata symbol,
    uint256 originChainId,
    bytes calldata originAddress,
    string calldata coinType,
    uint8 decimals
)
    internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token on ZetaChain.|
|`symbol`|`string`|The symbol of the token.|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists.|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain.|
|`coinType`|`string`|The type of the original coin.|
|`decimals`|`uint8`|The number of decimals the token uses.|


#### _setZRC20TokenActive

Updates ZRC20 token active status.


```solidity
function _setZRC20TokenActive(address address_, bool active) internal;
```

#### getChainInfo

Gets information about a specific chain.


```solidity
function getChainInfo(uint256 chainId) external view returns (address gasZRC20, bytes memory registry);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`gasZRC20`|`address`|The address of the ZRC20 token that represents gas token for the chain.|
|`registry`|`bytes`|The registry address deployed on the chain.|


#### getChainMetadata

Gets chain-specific metadata


```solidity
function getChainMetadata(uint256 chainId, string calldata key) external view returns (bytes memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain|
|`key`|`string`|The metadata key to retrieve|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes`|The value of the requested metadata|


#### getContractInfo

Gets information about a specific contract


```solidity
function getContractInfo(
    uint256 chainId,
    string calldata contractType
)
    external
    view
    returns (bool active, bytes memory addressBytes);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`|The type of the contract|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`active`|`bool`|Whether the contract is active|
|`addressBytes`|`bytes`|The address of the contract|


#### getContractConfiguration

Gets contract-specific configuration


```solidity
function getContractConfiguration(
    uint256 chainId,
    string calldata contractType,
    string calldata key
)
    external
    view
    returns (bytes memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`|The type of the contract|
|`key`|`string`|The configuration key to retrieve|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes`|The value of the requested configuration|


#### getZRC20TokenInfo

Gets information about a specific ZRC20 token


```solidity
function getZRC20TokenInfo(address address_)
    external
    view
    returns (
        bool active,
        string memory symbol,
        uint256 originChainId,
        bytes memory originAddress,
        string memory coinType,
        uint8 decimals
    );
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`active`|`bool`|Whether the token is active|
|`symbol`|`string`|The symbol of the token|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain|
|`coinType`|`string`|The type of the original coin|
|`decimals`|`uint8`|The number of decimals the token uses|


#### getZRC20AddressByForeignAsset

Gets the ZRC20 token address for a specific asset on a foreign chain.


```solidity
function getZRC20AddressByForeignAsset(
    uint256 originChainId,
    bytes calldata originAddress
)
    external
    view
    returns (address);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`originChainId`|`uint256`|The ID of the foreign chain|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`address`|The address of the corresponding ZRC20 token on ZetaChain.|


#### getActiveChains

Gets all active chains in the registry.


```solidity
function getActiveChains() external view returns (uint256[] memory);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256[]`|Array of chain IDs for all active chains.|


#### getAllChains

Returns information for all chains (active and inactive) in the registry.


```solidity
function getAllChains() external view returns (ChainInfoDTO[] memory chainsInfo);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`chainsInfo`|`ChainInfoDTO[]`|Array of ChainInfoDTO structs containing information about all chains.|


#### getAllContracts

Returns information for all contracts in the registry.


```solidity
function getAllContracts() external view returns (ContractInfoDTO[] memory contractsInfo);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`contractsInfo`|`ContractInfoDTO[]`|Array of ContractInfoDTO structs containing information about all contracts.|


#### getAllZRC20Tokens

Returns information for all ZRC20 tokens in the registry.


```solidity
function getAllZRC20Tokens() external view returns (ZRC20Info[] memory tokensInfo);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`tokensInfo`|`ZRC20Info[]`|Array of ZRC20Info structs containing information about all ZRC20 tokens.|


#### _removeFromActiveChains

Removes a chain ID from the active chains array.


```solidity
function _removeFromActiveChains(uint256 chainId) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain to remove.|




## IBaseRegistry
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/helpers/interfaces/IBaseRegistry.sol)

Interface for the BaseRegistry contract.


### Functions
#### changeChainStatus

Changes status of the chain to activated/deactivated.


```solidity
function changeChainStatus(uint256 chainId, address gasZRC20, bytes calldata registry, bool activation) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain to activate.|
|`gasZRC20`|`address`|The address of the ZRC20 token that represents gas token for the chain.|
|`registry`|`bytes`||
|`activation`|`bool`|Whether activate or deactivate a chain|


#### updateChainMetadata

Updates chain metadata.


```solidity
function updateChainMetadata(uint256 chainId, string calldata key, bytes calldata value) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|
|`key`|`string`|The metadata key to update.|
|`value`|`bytes`|The new value for the metadata.|


#### registerContract

Registers a new contract address for a specific chain.


```solidity
function registerContract(uint256 chainId, string calldata contractType, bytes calldata addressBytes) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract (e.g., "connector", "gateway").|
|`addressBytes`|`bytes`|The bytes representation of the non-EVM address.|


#### updateContractConfiguration

Updates contract configuration.


```solidity
function updateContractConfiguration(
    uint256 chainId,
    string calldata contractType,
    string calldata key,
    bytes calldata value
)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`key`|`string`|The configuration key to update.|
|`value`|`bytes`|The new value for the configuration.|


#### setContractActive


```solidity
function setContractActive(uint256 chainId, string calldata contractType, bool active) external;
```

#### registerZRC20Token

Registers a new ZRC20 token in the registry.


```solidity
function registerZRC20Token(
    address address_,
    string calldata symbol,
    uint256 originChainId,
    bytes calldata originAddress,
    string calldata coinType,
    uint8 decimals
)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token on ZetaChain.|
|`symbol`|`string`|The symbol of the token.|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists.|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain.|
|`coinType`|`string`||
|`decimals`|`uint8`||


#### setZRC20TokenActive

Updates ZRC20 token information.


```solidity
function setZRC20TokenActive(address address_, bool active) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token.|
|`active`|`bool`|Whether the token should be active.|


#### getChainInfo

Gets information about a specific chain.


```solidity
function getChainInfo(uint256 chainId) external view returns (address gasZRC20, bytes memory registry);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`gasZRC20`|`address`|The address of the ZRC20 token that represents gas token for the chain.|
|`registry`|`bytes`|The registry address deployed on the chain.|


#### getChainMetadata

Gets chain-specific metadata.


```solidity
function getChainMetadata(uint256 chainId, string calldata key) external view returns (bytes memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|
|`key`|`string`|The metadata key to retrieve.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes`|The value of the requested metadata.|


#### getContractInfo

Gets information about a specific contract.


```solidity
function getContractInfo(
    uint256 chainId,
    string calldata contractType
)
    external
    view
    returns (bool active, bytes memory addressBytes);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`active`|`bool`|Whether the contract is active.|
|`addressBytes`|`bytes`|The address of the contract.|


#### getContractConfiguration

Gets contract-specific configuration.


```solidity
function getContractConfiguration(
    uint256 chainId,
    string calldata contractType,
    string calldata key
)
    external
    view
    returns (bytes memory);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`key`|`string`|The configuration key to retrieve.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bytes`|The value of the requested configuration.|


#### getZRC20TokenInfo

Gets information about a specific ZRC20 token.


```solidity
function getZRC20TokenInfo(address address_)
    external
    view
    returns (
        bool active,
        string memory symbol,
        uint256 originChainId,
        bytes memory originAddress,
        string memory coinType,
        uint8 decimals
    );
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`active`|`bool`|Whether the token is active.|
|`symbol`|`string`|The symbol of the token|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists.|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain.|
|`coinType`|`string`|The type of the original coin.|
|`decimals`|`uint8`|The number of decimals the token uses.|


#### getZRC20AddressByForeignAsset

Gets the ZRC20 token address for a specific asset on a foreign chain.


```solidity
function getZRC20AddressByForeignAsset(
    uint256 originChainId,
    bytes calldata originAddress
)
    external
    view
    returns (address);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`originChainId`|`uint256`|The ID of the foreign chain.|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain.|

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`address`|The address of the corresponding ZRC20 token on ZetaChain.|


#### getActiveChains

Gets all active chains in the registry.


```solidity
function getActiveChains() external view returns (uint256[] memory);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256[]`|Array of chain IDs for all active chains.|


#### getAllChains

Returns information for all chains (active and inactive) in the registry.


```solidity
function getAllChains() external view returns (ChainInfoDTO[] memory);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`ChainInfoDTO[]`|chainsInfo Array of ChainInfoDTO structs containing information about all chains.|


#### getAllContracts

Returns information for all contracts in the registry.


```solidity
function getAllContracts() external view returns (ContractInfoDTO[] memory);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`ContractInfoDTO[]`|contractsInfo Array of ContractInfoDTO structs containing information about all contracts.|


#### getAllZRC20Tokens

Gets all active chains in the registry.


```solidity
function getAllZRC20Tokens() external view returns (ZRC20Info[] memory);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`ZRC20Info[]`|tokensInfo Array of ZRC20Info structs containing information about all ZRC20 tokens.|




## IBaseRegistryErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/helpers/interfaces/IBaseRegistry.sol)

Interface for the errors used by the BaseRegistry contract.


### Errors
#### ZeroAddress
Error thrown when a zero address is provided where a non-zero address is required.


```solidity
error ZeroAddress();
```

#### InvalidSender
Error thrown when the sender is invalid


```solidity
error InvalidSender();
```

#### TransferFailed
Error thrown when a ZRC20 token transfer failed.


```solidity
error TransferFailed();
```

#### ChainActive
Error thrown when a chain is already active.


```solidity
error ChainActive(uint256 chainId);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain that is already active.|

#### ChainNonActive
Error thrown when a chain is not active.


```solidity
error ChainNonActive(uint256 chainId);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain that is not active.|

#### InvalidContractType
Error thrown when a contract type is invalid.


```solidity
error InvalidContractType(string message);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`message`|`string`|Describes why error happened|

#### ContractAlreadyRegistered
Error thrown when a contract is already registered.


```solidity
error ContractAlreadyRegistered(uint256 chainId, string contractType, bytes addressBytes);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|
|`contractType`|`string`|The type of the contract.|
|`addressBytes`|`bytes`|The address of the contract.|

#### ContractNotFound
Error thrown when a contract is not found in the registry.


```solidity
error ContractNotFound(uint256 chainId, string contractType);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain,|
|`contractType`|`string`|The type of the contract.|

#### ZRC20AlreadyRegistered
Error thrown when a ZRC20 token is already registered.


```solidity
error ZRC20AlreadyRegistered(address address_);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token.|

#### ZRC20SymbolAlreadyInUse
Error thrown when a ZRC20 token symbol is already in use.


```solidity
error ZRC20SymbolAlreadyInUse(string symbol);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`symbol`|`string`|The symbol that is already in use.|



## IBaseRegistryEvents
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/helpers/interfaces/IBaseRegistry.sol)

Interface for the events emitted by the BaseRegistry contract.


### Events
#### ChainStatusChanged
Emitted when a chain status has changed.


```solidity
event ChainStatusChanged(uint256 indexed chainId, bool newStatus);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|
|`newStatus`|`bool`|The new chain status (is active or not).|

#### ChainMetadataUpdated
Emitted when a chain metadata is set.


```solidity
event ChainMetadataUpdated(uint256 indexed chainId, string key, bytes value);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|
|`key`|`string`|The metadata key to update.|
|`value`|`bytes`|The new value for the metadata.|

#### ContractRegistered
Emitted when a new contract is registered.


```solidity
event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract (e.g. "connector", "gateway", "tss").|
|`addressBytes`|`bytes`|The contract address in bytes representation.|

#### ContractStatusChanged
Emitted when a contract status has changed.


```solidity
event ContractStatusChanged(bytes addressBytes);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`addressBytes`|`bytes`|The contract address in bytes representation.|

#### ContractConfigurationUpdated
Emitted when a contract configuration is updated.


```solidity
event ContractConfigurationUpdated(uint256 indexed chainId, string contractType, string key, bytes value);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`key`|`string`|The configuration key to update.|
|`value`|`bytes`|The new value for the configuration.|

#### ZRC20TokenRegistered
Emitted when a ZRC20 token is registered.


```solidity
event ZRC20TokenRegistered(
    bytes indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol
);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`originAddress`|`bytes`|The address of the asset on its native chain.|
|`address_`|`address`|The address of the ZRC20 token on ZetaChain.|
|`decimals`|`uint8`|The number of decimals the token uses.|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists.|
|`symbol`|`string`|The symbol of the token.|

#### ZRC20TokenUpdated
Emitted when a ZRC20 token is updated.


```solidity
event ZRC20TokenUpdated(address address_, bool active);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token.|
|`active`|`bool`|Whether the token should be active.|

#### AdminChanged
Emitted when admin address is changed.


```solidity
event AdminChanged(address oldAdmin, address newAdmin);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`oldAdmin`|`address`|The previous admin address.|
|`newAdmin`|`address`|The new admin address.|

#### RegistryManagerChanged
Emitted when registry manager address is changed.


```solidity
event RegistryManagerChanged(address oldRegistryManager, address newRegistryManager);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`oldRegistryManager`|`address`|The previous registry manager address.|
|`newRegistryManager`|`address`|The new registry manager address.|



## ChainInfo
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/helpers/interfaces/IBaseRegistry.sol)

Structure that contains information about a chain.


```solidity
struct ChainInfo {
    bool active;
    uint256 chainId;
    address gasZRC20;
    bytes registry;
    mapping(string => bytes) metadata;
}
```



## ChainInfoDTO
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/helpers/interfaces/IBaseRegistry.sol)

Structure that contains information about a chain, used for data retrieving.


```solidity
struct ChainInfoDTO {
    bool active;
    uint256 chainId;
    address gasZRC20;
    bytes registry;
}
```



## ContractIdentifier
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/helpers/interfaces/IBaseRegistry.sol)

Each entry consists of: chainId (uint256) and contractType (string)


```solidity
struct ContractIdentifier {
    uint256 chainId;
    string contractType;
}
```



## ContractInfo
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/helpers/interfaces/IBaseRegistry.sol)

Structure that contains information about a contract registered in the system.


```solidity
struct ContractInfo {
    bool active;
    bytes addressBytes;
    string contractType;
    mapping(string => bytes) configuration;
}
```



## ContractInfoDTO
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/helpers/interfaces/IBaseRegistry.sol)

Structure that contains information about a contract registered in the system, used for data retrieving.


```solidity
struct ContractInfoDTO {
    bool active;
    bytes addressBytes;
    string contractType;
    uint256 chainId;
}
```



## ZRC20Info
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/helpers/interfaces/IBaseRegistry.sol)

Structure that contains information about a ZRC20 token.


```solidity
struct ZRC20Info {
    bool active;
    address address_;
    bytes originAddress;
    uint256 originChainId;
    string symbol;
    string coinType;
    uint8 decimals;
}
```



## Constants
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/Revert.sol)

#### MAX_REVERT_GAS_LIMIT

```solidity
uint256 constant MAX_REVERT_GAS_LIMIT = 2_000_000;
```



## RevertGasLimitExceeded
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/Revert.sol)

Error indicating revert gas limit exceeds maximum allowed


```solidity
error RevertGasLimitExceeded(uint256 provided, uint256 maximum);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`provided`|`uint256`|The gas limit provided for revert operation.|
|`maximum`|`uint256`|The maximum allowed gas limit for revert operation.|



## Abortable
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/Revert.sol)

Interface for contracts that support abortable calls.


### Functions
#### onAbort

Called when a revertable call is aborted.


```solidity
function onAbort(AbortContext calldata abortContext) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`abortContext`|`AbortContext`|Abort context to pass to onAbort.|




## Revertable
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/Revert.sol)

Interface for contracts that support revertable calls.


### Functions
#### onRevert

Called when a revertable call is made.


```solidity
function onRevert(RevertContext calldata revertContext) external payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|




## AbortContext
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/Revert.sol)

Struct containing abort context passed to onAbort.


```solidity
struct AbortContext {
    bytes sender;
    address asset;
    uint256 amount;
    bool outgoing;
    uint256 chainID;
    bytes revertMessage;
}
```

**Properties**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`bytes`|Address of account that initiated smart contract call. bytes is used as the crosschain transaction can be initiated from a non-EVM chain.|
|`asset`|`address`|Address of asset. On a connected chain, it contains the fungible token address or is empty if it's a gas token. On ZetaChain, it contains the address of the ZRC20.|
|`amount`|`uint256`|Amount specified with the transaction.|
|`outgoing`|`bool`|Flag to indicate if the crosschain transaction was outgoing: from ZetaChain to connected chain. if false, the transaction was incoming: from connected chain to ZetaChain.|
|`chainID`|`uint256`|Chain ID of the connected chain.|
|`revertMessage`|`bytes`|Arbitrary data specified in the RevertOptions object when initating the crosschain transaction.|



## RevertContext
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/Revert.sol)

Struct containing revert context passed to onRevert.


```solidity
struct RevertContext {
    address sender;
    address asset;
    uint256 amount;
    bytes revertMessage;
}
```

**Properties**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|Address of account that initiated smart contract call.|
|`asset`|`address`|Address of asset. On a connected chain, it contains the fungible token address or is empty if it's a gas token. On ZetaChain, it contains the address of the ZRC20.|
|`amount`|`uint256`|Amount specified with the transaction.|
|`revertMessage`|`bytes`|Arbitrary data sent back in onRevert.|



## RevertOptions
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/Revert.sol)

Struct containing revert options


```solidity
struct RevertOptions {
    address revertAddress;
    bool callOnRevert;
    address abortAddress;
    bytes revertMessage;
    uint256 onRevertGasLimit;
}
```

**Properties**

|Name|Type|Description|
|----|----|-----------|
|`revertAddress`|`address`|Address to receive revert.|
|`callOnRevert`|`bool`|Flag if onRevert hook should be called.|
|`abortAddress`|`address`|Address to receive funds if aborted.|
|`revertMessage`|`bytes`|Arbitrary data sent back in onRevert.|
|`onRevertGasLimit`|`uint256`|Gas limit for revert tx, unused on GatewayZEVM methods|



## CoreRegistry
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/CoreRegistry.sol)

Central registry for ZetaChain, managing chain info, ZRC20 data, and contract addresses across all chains.

*The contract doesn't hold any funds and should never have active allowances.*


### State Variables
#### CROSS_CHAIN_GAS_LIMIT
Cross-chain message gas limit


```solidity
uint256 public constant CROSS_CHAIN_GAS_LIMIT = 500_000;
```


#### gatewayZEVM
Instance of the GatewayZEVM contract for cross-chain communication


```solidity
IGatewayZEVM public gatewayZEVM;
```


### Functions
#### initialize

Initialize the CoreRegistry contract.


```solidity
function initialize(address admin_, address registryManager_, address gatewayZEVM_) public initializer;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`admin_`|`address`|Address with DEFAULT_ADMIN_ROLE, authorized for upgrades and pausing actions.|
|`registryManager_`|`address`|Address with REGISTRY_MANAGER_ROLE, authorized for all registry write actions.|
|`gatewayZEVM_`|`address`|Address of the GatewayZEVM contract for cross-chain messaging|


#### changeChainStatus

Changes status of the chain to activated/deactivated.


```solidity
function changeChainStatus(
    uint256 chainId,
    address gasZRC20,
    bytes calldata registry,
    bool activation
)
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain to activate.|
|`gasZRC20`|`address`|The address of the ZRC20 token that represents gas token for the chain.|
|`registry`|`bytes`|Address of the Registry contract on the connected chain.|
|`activation`|`bool`|Whether activate or deactivate the chain|


#### updateChainMetadata

Updates chain metadata, only for the active chains.


```solidity
function updateChainMetadata(
    uint256 chainId,
    string calldata key,
    bytes calldata value
)
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain.|
|`key`|`string`|The metadata key to update.|
|`value`|`bytes`|The new value for the metadata.|


#### registerContract

Registers a new contract address for a specific chain.


```solidity
function registerContract(
    uint256 chainId,
    string calldata contractType,
    bytes calldata addressBytes
)
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract (e.g., "connector", "gateway").|
|`addressBytes`|`bytes`|The bytes representation of the non-EVM address.|


#### updateContractConfiguration

Updates contract configuration.


```solidity
function updateContractConfiguration(
    uint256 chainId,
    string calldata contractType,
    string calldata key,
    bytes calldata value
)
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`key`|`string`|The configuration key to update.|
|`value`|`bytes`|The new value for the configuration.|


#### setContractActive

Sets a contract's active status


```solidity
function setContractActive(
    uint256 chainId,
    string calldata contractType,
    bool active
)
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed.|
|`contractType`|`string`|The type of the contract.|
|`active`|`bool`|Whether the contract should be active.|


#### registerZRC20Token

Registers a new ZRC20 token in the registry.


```solidity
function registerZRC20Token(
    address address_,
    string calldata symbol,
    uint256 originChainId,
    bytes calldata originAddress,
    string calldata coinType,
    uint8 decimals
)
    external
    onlyRole(REGISTRY_MANAGER_ROLE)
    whenNotPaused;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token on ZetaChain.|
|`symbol`|`string`|The symbol of the token.|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists.|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain.|
|`coinType`|`string`|The type of the original coin.|
|`decimals`|`uint8`|The number of decimals the token uses.|


#### setZRC20TokenActive

Updates ZRC20 token active status.


```solidity
function setZRC20TokenActive(address address_, bool active) external onlyRole(REGISTRY_MANAGER_ROLE) whenNotPaused;
```

#### _broadcastChainActivation

Broadcast chain activation update to all satellite registries.


```solidity
function _broadcastChainActivation(
    uint256 chainId,
    address gasZRC20,
    bytes calldata registry,
    bool activation
)
    internal;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain being activated/deactivated.|
|`gasZRC20`|`address`|The address of the ZRC20 token that represents gas token for the chain.|
|`registry`|`bytes`|Address of the Registry contract on the connected chain.|
|`activation`|`bool`|Whether activate or deactivate the chain|


#### _broadcastChainMetadataUpdate

Broadcast chain metadata to all satellite registries


```solidity
function _broadcastChainMetadataUpdate(uint256 chainId, string calldata key, bytes calldata value) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain whose metadata is being updated|
|`key`|`string`|The metadata key being updated|
|`value`|`bytes`|The new value for the metadata|


#### _broadcastContractRegistration

Broadcast contract registration to all satellite registries

contractType The type of the contract

addressBytes The bytes representation of the non-EVM address


```solidity
function _broadcastContractRegistration(
    uint256 chainId,
    string calldata contractType,
    bytes calldata addressBytes
)
    private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`||
|`addressBytes`|`bytes`||


#### _broadcastContractConfigUpdate

Broadcast contract configuration update to all satellite registries

contractType The type of the contract

key The configuration key being updated

value The new value for the configuration


```solidity
function _broadcastContractConfigUpdate(
    uint256 chainId,
    string calldata contractType,
    string calldata key,
    bytes calldata value
)
    private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`||
|`key`|`string`||
|`value`|`bytes`||


#### _broadcastContractStatusUpdate

Broadcast contract status update to all satellite registries

contractType The type of the contract

active Whether the contract should be active


```solidity
function _broadcastContractStatusUpdate(uint256 chainId, string calldata contractType, bool active) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainId`|`uint256`|The ID of the chain where the contract is deployed|
|`contractType`|`string`||
|`active`|`bool`||


#### _broadcastZRC20Registration

Broadcast ZRC20 token registration to all satellite registries


```solidity
function _broadcastZRC20Registration(
    address address_,
    string calldata symbol,
    uint256 originChainId,
    bytes calldata originAddress,
    string calldata coinType,
    uint8 decimals
)
    private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token on ZetaChain|
|`symbol`|`string`|The symbol of the token|
|`originChainId`|`uint256`|The ID of the foreign chain where the original asset exists|
|`originAddress`|`bytes`|The address or identifier of the asset on its native chain|
|`coinType`|`string`|The type of the original coin|
|`decimals`|`uint8`|The number of decimals the token uses|


#### _broadcastZRC20Update

Broadcast ZRC20 token update to all satellite registries


```solidity
function _broadcastZRC20Update(address address_, bool active) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`address_`|`address`|The address of the ZRC20 token|
|`active`|`bool`|Whether the token should be active|


#### _broadcastToAllChains

Generic function to broadcast encoded messages to all satellite registries


```solidity
function _broadcastToAllChains(bytes memory encodedMessage) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`encodedMessage`|`bytes`|The fully encoded function call to broadcast|


#### _sendCrossChainMessage

Sends a cross-chain message to the Registry contract on a target chain.


```solidity
function _sendCrossChainMessage(uint256 targetChainId, bytes memory message) private;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`targetChainId`|`uint256`|The ID of the chain to send the message to.|
|`message`|`bytes`|The encoded function call to execute on the target chain.|




## ICoreRegistry
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/ICoreRegistry.sol)


### Functions
#### gatewayZEVM


```solidity
function gatewayZEVM() external returns (address);
```



## IGatewayZEVM
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/IGatewayZEVM.sol)

Interface for the GatewayZEVM contract.

*Defines functions for cross-chain interactions and token handling.*


### Functions
#### withdraw

Withdraw ZRC20 tokens to an external chain.


```solidity
function withdraw(
    bytes memory receiver,
    uint256 amount,
    address zrc20,
    RevertOptions calldata revertOptions
)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver address on the external chain.|
|`amount`|`uint256`|The amount of tokens to withdraw.|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`revertOptions`|`RevertOptions`|Revert options.|


#### withdraw

Withdraw ZETA tokens to an external chain.


```solidity
function withdraw(bytes memory receiver, uint256 chainId, RevertOptions calldata revertOptions) external payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver address on the external chain.|
|`chainId`|`uint256`||
|`revertOptions`|`RevertOptions`|Revert options.|


#### withdrawAndCall

Withdraw ZRC20 tokens and call a smart contract on an external chain.


```solidity
function withdrawAndCall(
    bytes memory receiver,
    uint256 amount,
    address zrc20,
    bytes calldata message,
    CallOptions calldata callOptions,
    RevertOptions calldata revertOptions
)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver address on the external chain.|
|`amount`|`uint256`|The amount of tokens to withdraw.|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`message`|`bytes`|The calldata to pass to the contract call.|
|`callOptions`|`CallOptions`|Call options including gas limit and arbirtrary call flag.|
|`revertOptions`|`RevertOptions`|Revert options.|


#### withdrawAndCall

Withdraw ZETA tokens and call a smart contract on an external chain.


```solidity
function withdrawAndCall(
    bytes memory receiver,
    uint256 chainId,
    bytes calldata message,
    CallOptions calldata callOptions,
    RevertOptions calldata revertOptions
)
    external
    payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver address on the external chain.|
|`chainId`|`uint256`|Chain id of the external chain.|
|`message`|`bytes`|The calldata to pass to the contract call.|
|`callOptions`|`CallOptions`|Call options including gas limit and arbirtrary call flag.|
|`revertOptions`|`RevertOptions`|Revert options.|


#### call

Call a smart contract on an external chain without asset transfer.


```solidity
function call(
    bytes memory receiver,
    address zrc20,
    bytes calldata message,
    CallOptions calldata callOptions,
    RevertOptions calldata revertOptions
)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver address on the external chain.|
|`zrc20`|`address`|Address of zrc20 to pay fees.|
|`message`|`bytes`|The calldata to pass to the contract call.|
|`callOptions`|`CallOptions`|Call options including gas limit and arbirtrary call flag.|
|`revertOptions`|`RevertOptions`|Revert options.|


#### deposit

Deposit foreign coins into ZRC20.


```solidity
function deposit(address zrc20, uint256 amount, address target) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`amount`|`uint256`|The amount of tokens to deposit.|
|`target`|`address`|The target address to receive the deposited tokens.|


#### deposit

Deposit native ZETA.


```solidity
function deposit(address target) external payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`target`|`address`|The target address to receive the ZETA.|


#### execute

Execute a user-specified contract on ZEVM.


```solidity
function execute(
    MessageContext calldata context,
    address zrc20,
    uint256 amount,
    address target,
    bytes calldata message
)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`context`|`MessageContext`|The context of the cross-chain call.|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`amount`|`uint256`|The amount of tokens to transfer.|
|`target`|`address`|The target contract to call.|
|`message`|`bytes`|The calldata to pass to the contract call.|


#### depositAndCall

Deposit foreign coins into ZRC20 and call a user-specified contract on ZEVM.


```solidity
function depositAndCall(
    MessageContext calldata context,
    address zrc20,
    uint256 amount,
    address target,
    bytes calldata message
)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`context`|`MessageContext`|The context of the cross-chain call.|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`amount`|`uint256`|The amount of tokens to transfer.|
|`target`|`address`|The target contract to call.|
|`message`|`bytes`|The calldata to pass to the contract call.|


#### depositAndCall

Deposit native ZETA and call a user-specified contract on ZEVM.


```solidity
function depositAndCall(MessageContext calldata context, address target, bytes calldata message) external payable;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`context`|`MessageContext`|The context of the cross-chain call.|
|`target`|`address`|The target contract to call.|
|`message`|`bytes`|The calldata to pass to the contract call.|


#### executeRevert

Revert a user-specified contract on ZEVM.


```solidity
function executeRevert(address target, RevertContext calldata revertContext) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`target`|`address`|The target contract to call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|


#### depositAndRevert

Deposit foreign coins into ZRC20 and revert a user-specified contract on ZEVM.


```solidity
function depositAndRevert(
    address zrc20,
    uint256 amount,
    address target,
    RevertContext calldata revertContext
)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`amount`|`uint256`|The amount of tokens to revert.|
|`target`|`address`|The target contract to call.|
|`revertContext`|`RevertContext`|Revert context to pass to onRevert.|




## IGatewayZEVMErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/IGatewayZEVM.sol)

Interface for the errors used in the GatewayZEVM contract.


### Errors
#### WithdrawalFailed
Error indicating a withdrawal failure.


```solidity
error WithdrawalFailed(address token, address recipient, uint256 amount);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|The address of the token that failed to withdraw.|
|`recipient`|`address`|The address that was supposed to receive the tokens.|
|`amount`|`uint256`|The amount of tokens that failed to withdraw.|

#### InsufficientAmount
Error indicating an insufficient token amount.


```solidity
error InsufficientAmount();
```

#### ZRC20BurnFailed
Error indicating a failure to burn ZRC20 tokens.


```solidity
error ZRC20BurnFailed(address zrc20, uint256 amount);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The address of the ZRC20 token that failed to burn.|
|`amount`|`uint256`|The amount of tokens that failed to burn.|

#### ZRC20TransferFailed
Error indicating a failure to transfer ZRC20 tokens.


```solidity
error ZRC20TransferFailed(address zrc20, address from, address to, uint256 amount);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The address of the ZRC20 token that failed to transfer.|
|`from`|`address`|The address sending the tokens.|
|`to`|`address`|The address receiving the tokens.|
|`amount`|`uint256`|The amount of tokens that failed to transfer.|

#### ZRC20DepositFailed
Error indicating a failure to deposit ZRC20 tokens.


```solidity
error ZRC20DepositFailed(address zrc20, address to, uint256 amount);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The address of the ZRC20 token that failed to deposit.|
|`to`|`address`|The address that was supposed to receive the deposit.|
|`amount`|`uint256`|The amount of tokens that failed to deposit.|

#### GasFeeTransferFailed
Error indicating a failure to transfer gas fee.


```solidity
error GasFeeTransferFailed(address token, address to, uint256 amount);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`token`|`address`|The address of the token used for gas fee.|
|`to`|`address`|The address that was supposed to receive the gas fee.|
|`amount`|`uint256`|The amount of gas fee that failed to transfer.|

#### CallerIsNotProtocol
Error indicating that the caller is not the protocol account.


```solidity
error CallerIsNotProtocol();
```

#### InvalidTarget
Error indicating an invalid target address.


```solidity
error InvalidTarget();
```

#### FailedZetaSent
Error indicating a failure to send ZETA tokens.


```solidity
error FailedZetaSent(address recipient, uint256 amount);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`recipient`|`address`|The address that was supposed to receive the ZETA tokens.|
|`amount`|`uint256`|The amount of ZETA tokens that failed to send.|

#### OnlyWZETAOrProtocol
Error indicating that only WZETA or the protocol address can call the function.


```solidity
error OnlyWZETAOrProtocol();
```

#### InsufficientGasLimit
Error indicating an insufficient gas limit.


```solidity
error InsufficientGasLimit();
```

#### MessageSizeExceeded
Error indicating message size exceeded in external functions.


```solidity
error MessageSizeExceeded(uint256 provided, uint256 maximum);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`provided`|`uint256`|The size of the message that was provided.|
|`maximum`|`uint256`|The maximum allowed message size.|

#### ZeroGasPrice
Error indicating an invalid gas price.


```solidity
error ZeroGasPrice();
```



## IGatewayZEVMEvents
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/IGatewayZEVM.sol)

Interface for the events emitted by the GatewayZEVM contract.


### Events
#### Called
Emitted when a cross-chain call is made.


```solidity
event Called(
    address indexed sender,
    address indexed zrc20,
    bytes receiver,
    bytes message,
    CallOptions callOptions,
    RevertOptions revertOptions
);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|The address of the sender.|
|`zrc20`|`address`|Address of zrc20 to pay fees.|
|`receiver`|`bytes`|The receiver address on the external chain.|
|`message`|`bytes`|The calldata passed to the contract call.|
|`callOptions`|`CallOptions`|Call options including gas limit and arbirtrary call flag.|
|`revertOptions`|`RevertOptions`|Revert options.|

#### Withdrawn
Emitted when a withdrawal is made.


```solidity
event Withdrawn(
    address indexed sender,
    uint256 indexed chainId,
    bytes receiver,
    address zrc20,
    uint256 value,
    uint256 gasfee,
    uint256 protocolFlatFee,
    bytes message,
    CallOptions callOptions,
    RevertOptions revertOptions
);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|The address from which the tokens are withdrawn.|
|`chainId`|`uint256`|Chain id of external chain.|
|`receiver`|`bytes`|The receiver address on the external chain.|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`value`|`uint256`|The amount of tokens withdrawn.|
|`gasfee`|`uint256`|The gas fee for the withdrawal.|
|`protocolFlatFee`|`uint256`|The protocol flat fee for the withdrawal.|
|`message`|`bytes`|The calldata passed with the withdraw. No longer used. Kept to maintain compatibility.|
|`callOptions`|`CallOptions`|Call options including gas limit and arbirtrary call flag.|
|`revertOptions`|`RevertOptions`|Revert options.|

#### WithdrawnAndCalled
Emitted when a withdraw and call is made.


```solidity
event WithdrawnAndCalled(
    address indexed sender,
    uint256 indexed chainId,
    bytes receiver,
    address zrc20,
    uint256 value,
    uint256 gasfee,
    uint256 protocolFlatFee,
    bytes message,
    CallOptions callOptions,
    RevertOptions revertOptions
);
```

**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`|The address from which the tokens are withdrawn.|
|`chainId`|`uint256`|Chain id of external chain.|
|`receiver`|`bytes`|The receiver address on the external chain.|
|`zrc20`|`address`|The address of the ZRC20 token.|
|`value`|`uint256`|The amount of tokens withdrawn.|
|`gasfee`|`uint256`|The gas fee for the withdrawal.|
|`protocolFlatFee`|`uint256`|The protocol flat fee for the withdrawal.|
|`message`|`bytes`|The calldata passed to the contract call.|
|`callOptions`|`CallOptions`|Call options including gas limit and arbirtrary call flag.|
|`revertOptions`|`RevertOptions`|Revert options.|



## CallOptions
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/IGatewayZEVM.sol)

CallOptions struct passed to call and withdrawAndCall functions.


```solidity
struct CallOptions {
    uint256 gasLimit;
    bool isArbitraryCall;
}
```

**Properties**

|Name|Type|Description|
|----|----|-----------|
|`gasLimit`|`uint256`|Gas limit.|
|`isArbitraryCall`|`bool`|Indicates if call should be arbitrary or authenticated.|



## ISystem
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/ISystem.sol)

Interface for the System contract.

*Defines functions for system contract callable by fungible module.*


### Functions
#### FUNGIBLE_MODULE_ADDRESS


```solidity
function FUNGIBLE_MODULE_ADDRESS() external view returns (address);
```

#### wZetaContractAddress


```solidity
function wZetaContractAddress() external view returns (address);
```

#### uniswapv2FactoryAddress


```solidity
function uniswapv2FactoryAddress() external view returns (address);
```

#### gasPriceByChainId


```solidity
function gasPriceByChainId(uint256 chainID) external view returns (uint256);
```

#### gasCoinZRC20ByChainId


```solidity
function gasCoinZRC20ByChainId(uint256 chainID) external view returns (address);
```

#### gasZetaPoolByChainId


```solidity
function gasZetaPoolByChainId(uint256 chainID) external view returns (address);
```



## IWETH9
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/IWZETA.sol)

Interface for the Weth9 contract.


### Functions
#### totalSupply


```solidity
function totalSupply() external view returns (uint256);
```

#### balanceOf


```solidity
function balanceOf(address owner) external view returns (uint256);
```

#### allowance


```solidity
function allowance(address owner, address spender) external view returns (uint256);
```

#### approve


```solidity
function approve(address spender, uint256 wad) external returns (bool);
```

#### transfer


```solidity
function transfer(address to, uint256 wad) external returns (bool);
```

#### transferFrom


```solidity
function transferFrom(address from, address to, uint256 wad) external returns (bool);
```

#### deposit


```solidity
function deposit() external payable;
```

#### withdraw


```solidity
function withdraw(uint256 wad) external;
```

### Events
#### Approval

```solidity
event Approval(address indexed owner, address indexed spender, uint256 value);
```

#### Transfer

```solidity
event Transfer(address indexed from, address indexed to, uint256 value);
```

#### Deposit

```solidity
event Deposit(address indexed dst, uint256 wad);
```

#### Withdrawal

```solidity
event Withdrawal(address indexed src, uint256 wad);
```



## CoinType
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/IZRC20.sol)

*Coin types for ZRC20. Zeta value should not be used.*


```solidity
enum CoinType {
    Zeta,
    Gas,
    ERC20
}
```



## IZRC20
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/IZRC20.sol)

Interface for the ZRC20 token contract.


### Functions
#### totalSupply


```solidity
function totalSupply() external view returns (uint256);
```

#### balanceOf


```solidity
function balanceOf(address account) external view returns (uint256);
```

#### transfer


```solidity
function transfer(address recipient, uint256 amount) external returns (bool);
```

#### allowance


```solidity
function allowance(address owner, address spender) external view returns (uint256);
```

#### approve


```solidity
function approve(address spender, uint256 amount) external returns (bool);
```

#### transferFrom


```solidity
function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
```

#### deposit


```solidity
function deposit(address to, uint256 amount) external returns (bool);
```

#### burn


```solidity
function burn(uint256 amount) external returns (bool);
```

#### withdraw


```solidity
function withdraw(bytes memory to, uint256 amount) external returns (bool);
```

#### withdrawGasFee


```solidity
function withdrawGasFee() external view returns (address, uint256);
```

#### withdrawGasFeeWithGasLimit


```solidity
function withdrawGasFeeWithGasLimit(uint256 gasLimit) external view returns (address, uint256);
```

#### PROTOCOL_FLAT_FEE

*Name is in upper case to maintain compatibility with ZRC20.sol v1*


```solidity
function PROTOCOL_FLAT_FEE() external view returns (uint256);
```

#### GAS_LIMIT

*Name is in upper case to maintain compatibility with ZRC20.sol v1*


```solidity
function GAS_LIMIT() external view returns (uint256);
```

#### SYSTEM_CONTRACT_ADDRESS

*Name is in upper case to maintain compatibility with ZRC20.sol v1*


```solidity
function SYSTEM_CONTRACT_ADDRESS() external view returns (address);
```

#### setName


```solidity
function setName(string memory newName) external;
```

#### setSymbol


```solidity
function setSymbol(string memory newSymbol) external;
```



## IZRC20Metadata
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/IZRC20.sol)

Interface for the ZRC20 metadata.


### Functions
#### name


```solidity
function name() external view returns (string memory);
```

#### symbol


```solidity
function symbol() external view returns (string memory);
```

#### decimals


```solidity
function decimals() external view returns (uint8);
```



## ZRC20Events
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/IZRC20.sol)

Interface for the ZRC20 events.


### Events
#### Transfer

```solidity
event Transfer(address indexed from, address indexed to, uint256 value);
```

#### Approval

```solidity
event Approval(address indexed owner, address indexed spender, uint256 value);
```

#### Deposit

```solidity
event Deposit(bytes from, address indexed to, uint256 value);
```

#### Withdrawal

```solidity
event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee);
```

#### UpdatedSystemContract

```solidity
event UpdatedSystemContract(address systemContract);
```

#### UpdatedGateway

```solidity
event UpdatedGateway(address gateway);
```

#### UpdatedGasLimit

```solidity
event UpdatedGasLimit(uint256 gasLimit);
```

#### UpdatedProtocolFlatFee

```solidity
event UpdatedProtocolFlatFee(uint256 protocolFlatFee);
```



## UniversalContract
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/UniversalContract.sol)

Abstract contract for contracts that can receive cross-chain calls on ZetaChain.

*Contracts extending this abstract contract can handle incoming cross-chain messages
and execute logic based on the provided context, token, and message payload.*


### State Variables
#### registry
Reference to the ZetaChain Registry contract


```solidity
ICoreRegistry public constant registry = ICoreRegistry(0x7CCE3Eb018bf23e1FE2a32692f2C77592D110394);
```


#### gateway
Reference to the ZetaChain Gateway contract


```solidity
IGatewayZEVM public immutable gateway;
```


### Functions
#### onlyGateway

Restricts function access to only the gateway contract

*Used on functions that process cross-chain messages to ensure they're only called through the Gateway,
where message validation occurs.
Important for security in functions like `onCall()` and `onRevert()` that handle incoming cross-chain
operations.*


```solidity
modifier onlyGateway();
```

#### constructor

Initializes the contract by retrieving the gateway address from the registry

*Fetches the gateway contract address for the current chain from the registry.
If the gateway is not active or not found, the gateway will remain uninitialized (address(0)).*


```solidity
constructor();
```

#### onCall

Function to handle cross-chain calls with native ZETA transfers


```solidity
function onCall(MessageContext calldata context, bytes calldata message) external payable virtual;
```

#### onCall

Function to handle cross-chain calls with ZRC20 token transfers


```solidity
function onCall(
    MessageContext calldata context,
    address zrc20,
    uint256 amount,
    bytes calldata message
)
    external
    virtual;
```

### Errors
#### Unauthorized
Error thrown when a function is called by an unauthorized address


```solidity
error Unauthorized();
```



## zContract
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/UniversalContract.sol)

**Note:**
deprecated: should be removed once v2 SystemContract is not used anymore.
UniversalContract should be used


### Functions
#### onCrossChainCall


```solidity
function onCrossChainCall(zContext calldata context, address zrc20, uint256 amount, bytes calldata message) external;
```



## MessageContext
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/UniversalContract.sol)

Provides contextual information when executing a cross-chain call on ZetaChain.

*This struct helps identify the sender of the message across different blockchain environments.*


```solidity
struct MessageContext {
    bytes sender;
    address senderEVM;
    uint256 chainID;
}
```



## zContext
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/interfaces/UniversalContract.sol)

**Note:**
deprecated: should be removed once v2 SystemContract is not used anymore.
MessageContext should be used


```solidity
struct zContext {
    bytes origin;
    address sender;
    uint256 chainID;
}
```



## ZetaConnectorZEVM
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/legacy/ZetaConnectorZEVM.sol)


### State Variables
#### wzeta
WZETA token address.


```solidity
address public wzeta;
```


#### FUNGIBLE_MODULE_ADDRESS
Fungible module address.


```solidity
address public constant FUNGIBLE_MODULE_ADDRESS = payable(0x735b14BB79463307AAcBED86DAf3322B1e6226aB);
```


### Functions
#### onlyFungibleModule

*Modifier to restrict actions to fungible module.*


```solidity
modifier onlyFungibleModule();
```

#### constructor


```solidity
constructor(address wzeta_);
```

#### receive

*Receive function to receive ZETA from WETH9.withdraw().*


```solidity
receive() external payable;
```

#### setWzetaAddress


```solidity
function setWzetaAddress(address wzeta_) external onlyFungibleModule;
```

#### send

*Sends ZETA and bytes messages (to execute it) crosschain.*


```solidity
function send(ZetaInterfaces.SendInput calldata input) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`input`|`ZetaInterfaces.SendInput`||


#### onReceive

*Handler to receive data from other chain.
This method can be called only by Fungible Module.
Transfer the Zeta tokens to destination and calls onZetaMessage if it's needed.
To perform the transfer wrap the new tokens*


```solidity
function onReceive(
    bytes calldata zetaTxSenderAddress,
    uint256 sourceChainId,
    address destinationAddress,
    uint256 zetaValue,
    bytes calldata message,
    bytes32 internalSendHash
)
    external
    payable
    onlyFungibleModule;
```

#### onRevert

*Handler to receive errors from other chain.
This method can be called only by Fungible Module.
Transfer the Zeta tokens to destination and calls onZetaRevert if it's needed.*


```solidity
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
    payable
    onlyFungibleModule;
```

### Events
#### SetWZETA

```solidity
event SetWZETA(address wzeta_);
```

#### ZetaSent

```solidity
event ZetaSent(
    address sourceTxOriginAddress,
    address indexed zetaTxSenderAddress,
    uint256 indexed destinationChainId,
    bytes destinationAddress,
    uint256 zetaValueAndGas,
    uint256 destinationGasLimit,
    bytes message,
    bytes zetaParams
);
```

#### ZetaReceived

```solidity
event ZetaReceived(
    bytes zetaTxSenderAddress,
    uint256 indexed sourceChainId,
    address indexed destinationAddress,
    uint256 zetaValue,
    bytes message,
    bytes32 indexed internalSendHash
);
```

#### ZetaReverted

```solidity
event ZetaReverted(
    address zetaTxSenderAddress,
    uint256 sourceChainId,
    uint256 indexed destinationChainId,
    bytes destinationAddress,
    uint256 remainingZetaValue,
    bytes message,
    bytes32 indexed internalSendHash
);
```

### Errors
#### OnlyWZETAOrFungible
Contract custom errors.


```solidity
error OnlyWZETAOrFungible();
```

#### WZETATransferFailed

```solidity
error WZETATransferFailed();
```

#### OnlyFungibleModule

```solidity
error OnlyFungibleModule();
```

#### FailedZetaSent

```solidity
error FailedZetaSent();
```

#### WrongValue

```solidity
error WrongValue();
```



## ZetaInterfaces
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/legacy/ZetaConnectorZEVM.sol)


### Structs
#### SendInput
*Use SendInput to interact with the Connector: connector.send(SendInput)*


```solidity
struct SendInput {
    uint256 destinationChainId;
    bytes destinationAddress;
    uint256 destinationGasLimit;
    bytes message;
    uint256 zetaValueAndGas;
    bytes zetaParams;
}
```

#### ZetaMessage
*Our Connector calls onZetaMessage with this struct as argument*


```solidity
struct ZetaMessage {
    bytes zetaTxSenderAddress;
    uint256 sourceChainId;
    address destinationAddress;
    uint256 zetaValue;
    bytes message;
}
```

#### ZetaRevert
*Our Connector calls onZetaRevert with this struct as argument*


```solidity
struct ZetaRevert {
    address zetaTxSenderAddress;
    uint256 sourceChainId;
    bytes destinationAddress;
    uint256 destinationChainId;
    uint256 remainingZetaValue;
    bytes message;
}
```



## ZetaReceiver
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/legacy/ZetaConnectorZEVM.sol)


### Functions
#### onZetaMessage

*onZetaMessage is called when a cross-chain message reaches a contract*


```solidity
function onZetaMessage(ZetaInterfaces.ZetaMessage calldata zetaMessage) external;
```

#### onZetaRevert

*onZetaRevert is called when a cross-chain message reverts.
It's useful to rollback to the original state*


```solidity
function onZetaRevert(ZetaInterfaces.ZetaRevert calldata zetaRevert) external;
```



## GatewayZEVMValidations
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/libraries/GatewayZEVMValidations.sol)

Library containing validation functions for GatewayZEVM contract

*This library provides common validation logic used across GatewayZEVM contract*


### State Variables
#### MAX_MESSAGE_SIZE
Maximum message size constant


```solidity
uint256 internal constant MAX_MESSAGE_SIZE = 2880;
```


#### MIN_GAS_LIMIT
Minimum gas limit constant


```solidity
uint256 internal constant MIN_GAS_LIMIT = 100_000;
```


### Functions
#### validateNonZeroAddress

*Validates that an address is not zero*


```solidity
function validateNonZeroAddress(address addr) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`addr`|`address`|The address to validate|


#### validateReceiver

*Validates that receiver bytes are not empty*


```solidity
function validateReceiver(bytes memory receiver) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver bytes to validate|


#### validateAmount

*Validates that amount is not zero*


```solidity
function validateAmount(uint256 amount) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amount`|`uint256`|The amount to validate|


#### validateGasLimit

*Validates that gas limit meets minimum requirement*


```solidity
function validateGasLimit(uint256 gasLimit) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`gasLimit`|`uint256`|The gas limit to validate|


#### validateTarget

*Validates that target address is not restricted*


```solidity
function validateTarget(address target, address protocolAddress, address contractAddress) private pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`target`|`address`|The target address to validate|
|`protocolAddress`|`address`|The protocol address to check against|
|`contractAddress`|`address`|The contract address to check against|


#### validateMessageSize

*Validates message size constraints*


```solidity
function validateMessageSize(uint256 messageLength, uint256 revertMessageLength) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`messageLength`|`uint256`|The length of the main message|
|`revertMessageLength`|`uint256`|The length of the revert message|


#### validateRevertOptions

*Validates revert options*


```solidity
function validateRevertOptions(RevertOptions calldata revertOptions) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`revertOptions`|`RevertOptions`|The revert options to validate|


#### validateCallAndRevertOptions

*Validates call options and revert options together*


```solidity
function validateCallAndRevertOptions(
    CallOptions calldata callOptions,
    RevertOptions calldata revertOptions,
    uint256 messageLength
)
    internal
    pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`callOptions`|`CallOptions`|The call options to validate|
|`revertOptions`|`RevertOptions`|The revert options to validate|
|`messageLength`|`uint256`|The length of the message|


#### validateWithdrawalParams

*Validates standard withdrawal parameters*


```solidity
function validateWithdrawalParams(
    bytes memory receiver,
    uint256 amount,
    RevertOptions calldata revertOptions
)
    internal
    pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver address|
|`amount`|`uint256`|The amount to withdraw|
|`revertOptions`|`RevertOptions`|The revert options|


#### validateWithdrawalAndCallParams

*Validates withdrawal and call parameters*


```solidity
function validateWithdrawalAndCallParams(
    bytes memory receiver,
    uint256 amount,
    bytes calldata message,
    CallOptions calldata callOptions,
    RevertOptions calldata revertOptions
)
    internal
    pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`receiver`|`bytes`|The receiver address|
|`amount`|`uint256`|The amount to withdraw|
|`message`|`bytes`|The message to send|
|`callOptions`|`CallOptions`|The call options|
|`revertOptions`|`RevertOptions`|The revert options|


#### validateDepositParams

*Validates deposit parameters*


```solidity
function validateDepositParams(
    address zrc20,
    uint256 amount,
    address target,
    address protocolAddress,
    address contractAddress
)
    internal
    pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The ZRC20 token address|
|`amount`|`uint256`|The amount to deposit|
|`target`|`address`|The target address|
|`protocolAddress`|`address`|The protocol address|
|`contractAddress`|`address`|The contract address|


#### validateExecuteParams

*Validates execute parameters*


```solidity
function validateExecuteParams(address zrc20, address target) internal pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zrc20`|`address`|The ZRC20 token address|
|`target`|`address`|The target address|


#### validateZetaDepositParams

*Validates ZETA deposit and call parameters*


```solidity
function validateZetaDepositParams(
    uint256 amount,
    address target,
    address protocolAddress,
    address contractAddress
)
    internal
    pure;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amount`|`uint256`|The amount to deposit|
|`target`|`address`|The target address|
|`protocolAddress`|`address`|The protocol address|
|`contractAddress`|`address`|The contract address|


### Errors
#### EmptyAddress
Error indicating a empty address was provided.


```solidity
error EmptyAddress();
```



## SystemContract
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/SystemContract.sol)

*The system contract it's called by the protocol to interact with the blockchain.
Also includes a lot of tools to make easier to interact with ZetaChain.*


### State Variables
#### gasPriceByChainId
Map to know the gas price of each chain given a chain id.


```solidity
mapping(uint256 => uint256) public gasPriceByChainId;
```


#### gasCoinZRC20ByChainId
Map to know the ZRC20 address of a token given a chain id, ex zETH, zBNB etc.


```solidity
mapping(uint256 => address) public gasCoinZRC20ByChainId;
```


#### gasZetaPoolByChainId

```solidity
mapping(uint256 => address) public gasZetaPoolByChainId;
```


#### FUNGIBLE_MODULE_ADDRESS
Fungible address is always the same, it's on protocol level.


```solidity
address public constant FUNGIBLE_MODULE_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;
```


#### uniswapv2FactoryAddress
Uniswap V2 addresses.


```solidity
address public immutable uniswapv2FactoryAddress;
```


#### uniswapv2Router02Address

```solidity
address public immutable uniswapv2Router02Address;
```


#### wZetaContractAddress
Address of the wrapped ZETA to interact with Uniswap V2.


```solidity
address public wZetaContractAddress;
```


#### zetaConnectorZEVMAddress
Address of ZEVM Zeta Connector.


```solidity
address public zetaConnectorZEVMAddress;
```


### Functions
#### constructor

*Only fungible module can deploy a system contract.*


```solidity
constructor(address wzeta_, address uniswapv2Factory_, address uniswapv2Router02_);
```

#### depositAndCall

*Deposit foreign coins into ZRC20 and call user specified contract on zEVM.*


```solidity
function depositAndCall(
    zContext calldata context,
    address zrc20,
    uint256 amount,
    address target,
    bytes calldata message
)
    external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`context`|`zContext`||
|`zrc20`|`address`||
|`amount`|`uint256`||
|`target`|`address`||
|`message`|`bytes`||


#### sortTokens

*Sort token addresses lexicographically. Used to handle return values from pairs sorted in the order.*


```solidity
function sortTokens(address tokenA, address tokenB) internal pure returns (address token0, address token1);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`tokenA`|`address`||
|`tokenB`|`address`||

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`token0`|`address`|token1, returns sorted token addresses,.|
|`token1`|`address`||


#### uniswapv2PairFor

*Calculates the CREATE2 address for a pair without making any external calls.*


```solidity
function uniswapv2PairFor(address factory, address tokenA, address tokenB) public pure returns (address pair);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`factory`|`address`||
|`tokenA`|`address`||
|`tokenB`|`address`||

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`pair`|`address`|tokens pair address.|


#### setGasPrice

*Fungible module updates the gas price oracle periodically.*


```solidity
function setGasPrice(uint256 chainID, uint256 price) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainID`|`uint256`||
|`price`|`uint256`||


#### setGasCoinZRC20

*Setter for gasCoinZRC20ByChainId map.*


```solidity
function setGasCoinZRC20(uint256 chainID, address zrc20) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainID`|`uint256`||
|`zrc20`|`address`||


#### setGasZetaPool

*Set the pool wzeta/erc20 address.*


```solidity
function setGasZetaPool(uint256 chainID, address erc20) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`chainID`|`uint256`||
|`erc20`|`address`||


#### setWZETAContractAddress

*Setter for wrapped ZETA address.*


```solidity
function setWZETAContractAddress(address addr) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`addr`|`address`||


#### setConnectorZEVMAddress

*Setter for zetaConnector ZEVM Address*


```solidity
function setConnectorZEVMAddress(address addr) external;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`addr`|`address`||


### Events
#### SystemContractDeployed
Custom SystemContract errors.


```solidity
event SystemContractDeployed();
```

#### SetGasPrice

```solidity
event SetGasPrice(uint256, uint256);
```

#### SetGasCoin

```solidity
event SetGasCoin(uint256, address);
```

#### SetGasZetaPool

```solidity
event SetGasZetaPool(uint256, address);
```

#### SetWZeta

```solidity
event SetWZeta(address);
```

#### SetConnectorZEVM

```solidity
event SetConnectorZEVM(address);
```



## SystemContractErrors
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/SystemContract.sol)

*Custom errors for SystemContract*


### Errors
#### CallerIsNotFungibleModule

```solidity
error CallerIsNotFungibleModule();
```

#### InvalidTarget

```solidity
error InvalidTarget();
```

#### CantBeIdenticalAddresses

```solidity
error CantBeIdenticalAddresses();
```

#### CantBeZeroAddress

```solidity
error CantBeZeroAddress();
```

#### ZeroAddress

```solidity
error ZeroAddress();
```



## WETH9
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/WZETA.sol)


### State Variables
#### name

```solidity
string public name = "Wrapped Ether";
```


#### symbol

```solidity
string public symbol = "WETH";
```


#### decimals

```solidity
uint8 public decimals = 18;
```


#### balanceOf

```solidity
mapping(address => uint256) public balanceOf;
```


#### allowance

```solidity
mapping(address => mapping(address => uint256)) public allowance;
```


### Functions
#### receive


```solidity
receive() external payable;
```

#### deposit


```solidity
function deposit() public payable;
```

#### withdraw


```solidity
function withdraw(uint256 wad) public;
```

#### totalSupply


```solidity
function totalSupply() public view returns (uint256);
```

#### approve


```solidity
function approve(address guy, uint256 wad) public returns (bool);
```

#### transfer


```solidity
function transfer(address dst, uint256 wad) public returns (bool);
```

#### transferFrom


```solidity
function transferFrom(address src, address dst, uint256 wad) public returns (bool);
```

### Events
#### Approval

```solidity
event Approval(address indexed src, address indexed guy, uint256 wad);
```

#### Transfer

```solidity
event Transfer(address indexed src, address indexed dst, uint256 wad);
```

#### Deposit

```solidity
event Deposit(address indexed dst, uint256 wad);
```

#### Withdrawal

```solidity
event Withdrawal(address indexed src, uint256 wad);
```



## ZRC20
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/ZRC20.sol)


### State Variables
#### FUNGIBLE_MODULE_ADDRESS
Fungible address is always the same, maintained at the protocol level


```solidity
address public constant FUNGIBLE_MODULE_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;
```


#### CHAIN_ID
Chain id.abi


```solidity
uint256 public immutable CHAIN_ID;
```


#### COIN_TYPE
Coin type, checkout Interfaces.sol.


```solidity
CoinType public immutable COIN_TYPE;
```


#### SYSTEM_CONTRACT_ADDRESS
System contract address.

*Name is in upper case to maintain compatibility with ZRC20.sol v1*


```solidity
address public SYSTEM_CONTRACT_ADDRESS;
```


#### GAS_LIMIT
Gas limit.

*Name is in upper case to maintain compatibility with ZRC20.sol v1*


```solidity
uint256 public GAS_LIMIT;
```


#### PROTOCOL_FLAT_FEE
Protocol flat fee.

*Name is in upper case to maintain compatibility with ZRC20.sol v1*


```solidity
uint256 public override PROTOCOL_FLAT_FEE;
```


#### _balances

```solidity
mapping(address => uint256) private _balances;
```


#### _allowances

```solidity
mapping(address => mapping(address => uint256)) private _allowances;
```


#### _totalSupply

```solidity
uint256 private _totalSupply;
```


#### _name

```solidity
string private _name;
```


#### _symbol

```solidity
string private _symbol;
```


#### _decimals

```solidity
uint8 private _decimals;
```


#### gatewayAddress
Gateway contract address.

*This variable is added at last position to maintain storage layout with ZRC20.sol v1*


```solidity
address public gatewayAddress;
```


### Functions
#### _msgSender


```solidity
function _msgSender() internal view virtual returns (address);
```

#### onlyFungible

*Only fungible module modifier.*


```solidity
modifier onlyFungible();
```

#### constructor

*The only one allowed to deploy new ZRC20 is fungible address.*


```solidity
constructor(
    string memory name_,
    string memory symbol_,
    uint8 decimals_,
    uint256 chainid_,
    CoinType coinType_,
    uint256 gasLimit_,
    address systemContractAddress_,
    address gatewayAddress_
);
```

#### name

*ZRC20 name*


```solidity
function name() public view virtual override returns (string memory);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`string`|name as string|


#### setName

*Name can be updated by fungible module account.*


```solidity
function setName(string memory newName) external override onlyFungible;
```

#### setSymbol

*Symbol can be updated by fungible module account.*


```solidity
function setSymbol(string memory newSymbol) external override onlyFungible;
```

#### symbol

*ZRC20 symbol.*


```solidity
function symbol() public view virtual override returns (string memory);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`string`|symbol as string.|


#### decimals

*ZRC20 decimals.*


```solidity
function decimals() public view virtual override returns (uint8);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint8`|returns uint8 decimals.|


#### totalSupply

*ZRC20 total supply.*


```solidity
function totalSupply() public view virtual override returns (uint256);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|returns uint256 total supply.|


#### balanceOf

*Returns ZRC20 balance of an account.*


```solidity
function balanceOf(address account) public view virtual override returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`account`|`address`||

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|uint256 account balance.|


#### transfer

*Returns ZRC20 balance of an account.*


```solidity
function transfer(address recipient, uint256 amount) public virtual override returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`recipient`|`address`||
|`amount`|`uint256`||

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|true/false if transfer succeeded/failed.|


#### allowance

*Returns token allowance from owner to spender.*


```solidity
function allowance(address owner, address spender) public view virtual override returns (uint256);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`owner`|`address`||
|`spender`|`address`||

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`uint256`|uint256 allowance.|


#### approve

*Approves amount transferFrom for spender.*


```solidity
function approve(address spender, uint256 amount) public virtual override returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`spender`|`address`||
|`amount`|`uint256`||

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|true/false if succeeded/failed.|


#### transferFrom

*Transfers tokens from sender to recipient.*


```solidity
function transferFrom(address sender, address recipient, uint256 amount) public virtual override returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`sender`|`address`||
|`recipient`|`address`||
|`amount`|`uint256`||

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|true/false if succeeded/failed.|


#### burn

*Burns an amount of tokens.*


```solidity
function burn(uint256 amount) external override returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`amount`|`uint256`||

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|true/false if succeeded/failed.|


#### _transfer


```solidity
function _transfer(address sender, address recipient, uint256 amount) internal virtual;
```

#### _mint


```solidity
function _mint(address account, uint256 amount) internal virtual;
```

#### _burn


```solidity
function _burn(address account, uint256 amount) internal virtual;
```

#### _approve


```solidity
function _approve(address owner, address spender, uint256 amount) internal virtual;
```

#### deposit

*Deposits corresponding tokens from external chain, only callable by Fungible module.*


```solidity
function deposit(address to, uint256 amount) external override returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`address`||
|`amount`|`uint256`||

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|true/false if succeeded/failed.|


#### withdrawGasFee

*Withdraws gas fees.*


```solidity
function withdrawGasFee() public view override returns (address, uint256);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`address`|returns the ZRC20 address for gas on the same chain of this ZRC20, and calculates the gas fee for withdraw()|
|`<none>`|`uint256`||


#### withdrawGasFeeWithGasLimit

*Withdraws gas fees with specified gasLimit*


```solidity
function withdrawGasFeeWithGasLimit(uint256 gasLimit) public view override returns (address, uint256);
```
**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`address`|returns the ZRC20 address for gas on the same chain of this ZRC20, and calculates the gas fee for withdraw()|
|`<none>`|`uint256`||


#### withdraw

*Withraws ZRC20 tokens to external chains, this function causes cctx module to send out outbound tx to the
outbound chain
this contract should be given enough allowance of the gas ZRC20 to pay for outbound tx gas fee.*


```solidity
function withdraw(bytes memory to, uint256 amount) external override returns (bool);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`to`|`bytes`||
|`amount`|`uint256`||

**Returns**

|Name|Type|Description|
|----|----|-----------|
|`<none>`|`bool`|true/false if succeeded/failed.|


#### updateSystemContractAddress

*Updates system contract address. Can only be updated by the fungible module.*


```solidity
function updateSystemContractAddress(address addr) external onlyFungible;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`addr`|`address`||


#### updateGatewayAddress

*Updates gateway contract address. Can only be updated by the fungible module.*


```solidity
function updateGatewayAddress(address addr) external onlyFungible;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`addr`|`address`||


#### updateGasLimit

*Updates gas limit. Can only be updated by the fungible module.*


```solidity
function updateGasLimit(uint256 gasLimit_) external onlyFungible;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`gasLimit_`|`uint256`||


#### updateProtocolFlatFee

*Updates protocol flat fee. Can only be updated by the fungible module.*


```solidity
function updateProtocolFlatFee(uint256 protocolFlatFee_) external onlyFungible;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`protocolFlatFee_`|`uint256`||




## ZRC20Errors
[Git Source](https://github.com/zeta-chain/protocol-contracts-evm/blob/bb2aa0f0936267e47f637d52fd190e8b250b4c93/contracts/zevm/ZRC20.sol)

*Custom errors for ZRC20*


### Errors
#### CallerIsNotFungibleModule

```solidity
error CallerIsNotFungibleModule();
```

#### InvalidSender

```solidity
error InvalidSender();
```

#### GasFeeTransferFailed

```solidity
error GasFeeTransferFailed();
```

#### ZeroGasCoin

```solidity
error ZeroGasCoin();
```

#### ZeroGasPrice

```solidity
error ZeroGasPrice();
```

#### LowAllowance

```solidity
error LowAllowance();
```

#### LowBalance

```solidity
error LowBalance();
```

#### ZeroAddress

```solidity
error ZeroAddress();
```

