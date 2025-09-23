# GatewayEVM
[Git Source](https://github.com/zeta-chain/protocol-contracts/blob/main/v2/contracts/evm/GatewayEVM.sol)

**Inherits:**
Initializable, AccessControlUpgradeable, UUPSUpgradeable, [IGatewayEVM](/contracts/evm/interfaces/IGatewayEVM.sol/interface.IGatewayEVM.md), ReentrancyGuardUpgradeable, PausableUpgradeable, [INotSupportedMethods](/contracts/Errors.sol/interface.INotSupportedMethods.md)

The GatewayEVM contract is the endpoint to call smart contracts on external chains.

*The contract doesn't hold any funds and should never have active allowances.*


## State Variables
### custody
The address of the custody contract.


```solidity
address public custody;
```


### tssAddress
The address of the TSS (Threshold Signature Scheme) contract.


```solidity
address public tssAddress;
```


### zetaConnector
The address of the ZetaConnector contract.


```solidity
address public zetaConnector;
```


### zetaToken
The address of the Zeta token contract.


```solidity
address public zetaToken;
```


### additionalActionFeeWei
Fee charged for additional cross-chain actions within the same transaction.

*The first action in a transaction is free, subsequent actions incur this fee.*

*This is configurable by the admin role to allow for fee adjustments.*


```solidity
uint256 public additionalActionFeeWei;
```


### TSS_ROLE
New role identifier for tss role.


```solidity
bytes32 public constant TSS_ROLE = keccak256("TSS_ROLE");
```


### ASSET_HANDLER_ROLE
New role identifier for asset handler role.


```solidity
bytes32 public constant ASSET_HANDLER_ROLE = keccak256("ASSET_HANDLER_ROLE");
```


### PAUSER_ROLE
New role identifier for pauser role.


```solidity
bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
```


### MAX_PAYLOAD_SIZE
Max size of payload + revertOptions revert message.


```solidity
uint256 public constant MAX_PAYLOAD_SIZE = 2880;
```


### _TRANSACTION_ACTION_COUNT_KEY
Storage slot key for tracking transaction action count.

*Uses transient storage (tload/tstore) for gas efficiency.*

*Value 0x01 is used as a unique identifier for this storage slot.*


```solidity
uint256 private constant _TRANSACTION_ACTION_COUNT_KEY = 0x01;
```


## Functions
### constructor

**Note:**
oz-upgrades-unsafe-allow: constructor


```solidity
constructor();
```

### initialize

Initialize with tss address. address of zeta token and admin account set as DEFAULT_ADMIN_ROLE.

*Using admin to authorize upgrades and pause, and tss for tss role.*


```solidity
function initialize(address tssAddress_, address zetaToken_, address admin_) public initializer;
```

### _authorizeUpgrade

*Authorizes the upgrade of the contract, sender must be owner.*


```solidity
function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newImplementation`|`address`|Address of the new implementation.|


### updateTSSAddress

Update tss address


```solidity
function updateTSSAddress(address newTSSAddress) external onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`newTSSAddress`|`address`|new tss address|


### pause

Pause contract.


```solidity
function pause() external onlyRole(PAUSER_ROLE);
```

### unpause

Unpause contract.


```solidity
function unpause() external onlyRole(PAUSER_ROLE);
```

### updateAdditionalActionFee

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


### executeRevert

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


### executeWithERC20

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


### revertWithERC20

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


### deposit

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


### deposit

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


### depositAndCall

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


### depositAndCall

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


### call

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


### setCustody

Sets the custody contract address.


```solidity
function setCustody(address custody_) external onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`custody_`|`address`|Address of the custody contract.|


### setConnector

Sets the connector contract address.


```solidity
function setConnector(address zetaConnector_) external onlyRole(DEFAULT_ADMIN_ROLE);
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`zetaConnector_`|`address`|Address of the connector contract.|


### _resetApproval

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


### _transferFromToAssetHandler

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


### _transferToAssetHandler

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


### _executeArbitraryCall

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


### _executeAuthenticatedCall

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


### _revertIfOnCallOrOnRevert


```solidity
function _revertIfOnCallOrOnRevert(bytes calldata data) private pure;
```

### _processFee

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


### _validateChargedFeeForERC20

Validates fee payment for ERC20 operations (deposit, depositAndCall, call).

*Validates that msg.value equals the required fee (no excess ETH allowed).*


```solidity
function _validateChargedFeeForERC20(uint256 feeCharged) internal view;
```
**Parameters**

|Name|Type|Description|
|----|----|-----------|
|`feeCharged`|`uint256`|The fee amount that was charged.|


### _validateChargedFeeForETHWithAmount

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


### _getNextActionIndex

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


