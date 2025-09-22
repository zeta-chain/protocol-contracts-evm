// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayevmupgradetest

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Sender common.Address
}

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// RevertOptions is an auto generated low-level Go binding around an user-defined struct.
type RevertOptions struct {
	RevertAddress    common.Address
	CallOnRevert     bool
	AbortAddress     common.Address
	RevertMessage    []byte
	OnRevertGasLimit *big.Int
}

// GatewayEVMUpgradeTestMetaData contains all meta data concerning the GatewayEVMUpgradeTest contract.
var GatewayEVMUpgradeTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ASSET_HANDLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PAYLOAD_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"additionalActionFeeWei\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"custody\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"messageContext\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeWithERC20\",\"inputs\":[{\"name\":\"messageContext\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revertWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConnector\",\"inputs\":[{\"name\":\"zetaConnector_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustody\",\"inputs\":[{\"name\":\"custody_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAdditionalActionFee\",\"inputs\":[{\"name\":\"newFeeWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"zetaConnector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositedAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedV2\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedAdditionalActionFee\",\"inputs\":[{\"name\":\"oldFeeWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newFeeWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AdditionalActionDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallOnRevertNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExcessETHProvided\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectValueProvided\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFee\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RevertGasLimitExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZETANotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a060405230608052348015610013575f80fd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516142586100f95f395f8181612a1101528181612a3a0152612bec01526142585ff3fe608060405260043610610277575f3560e01c8063744b9b8b1161014b578063aa0c0fc1116100c6578063cb7ba8e51161007c578063d547741f11610062578063d547741f14610723578063dda79b7514610742578063e63ab1e914610760575f80fd5b8063cb7ba8e5146106fd578063d09e3b7814610710575f80fd5b8063ae7a3a6f116100ac578063ae7a3a6f146106aa578063b0107214146106c9578063c0c53b8b146106de575f80fd5b8063aa0c0fc114610643578063ad3cb1cc14610662575f80fd5b806391d148541161011b578063a217fddf11610101578063a217fddf146105e8578063a2ba1934146105fb578063a783c78914610610575f80fd5b806391d1485414610566578063950837aa146105c9575f80fd5b8063744b9b8b146105015780637bbe9afa146105145780637c744253146105335780638456cb5914610552575f80fd5b806338e22527116101f557806357bec62f116101ab5780635c975abb116101915780635c975abb146104855780635d62c860146104bb578063726ac97c146104ee575f80fd5b806357bec62f146104475780635b11259114610466575f80fd5b80633f4ba83a116101db5780633f4ba83a1461040c5780634f1ef2861461042057806352d1902d14610433575f80fd5b806338e22527146103d9578063397e375c146103f9575f80fd5b806321e093b11161024a578063282906ed11610230578063282906ed146103885780632f2ff15d1461039b57806336568abe146103ba575f80fd5b806321e093b1146102f6578063248a9ca31461032d575f80fd5b806301ffc9a71461027b57806310188aef146102af578063102614b0146102d05780631becceb4146102e3575b5f80fd5b348015610286575f80fd5b5061029a610295366004613707565b610793565b60405190151581526020015b60405180910390f35b3480156102ba575f80fd5b506102ce6102c9366004613761565b61082b565b005b6102ce6102de366004613790565b610905565b6102ce6102f1366004613839565b610aa0565b348015610301575f80fd5b50600354610315906001600160a01b031681565b6040516001600160a01b0390911681526020016102a6565b348015610338575f80fd5b5061037a61034736600461389b565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016102a6565b6102ce6103963660046138b2565b610c33565b3480156103a6575f80fd5b506102ce6103b5366004613905565b610df6565b3480156103c5575f80fd5b506102ce6103d4366004613905565b610e3f565b6103ec6103e736600461393f565b610e90565b6040516102a691906139cb565b6102ce6104073660046139dd565b610fab565b348015610417575f80fd5b506102ce6111d5565b6102ce61042e366004613ae5565b61120a565b34801561043e575f80fd5b5061037a611229565b348015610452575f80fd5b50600254610315906001600160a01b031681565b348015610471575f80fd5b50600154610315906001600160a01b031681565b348015610490575f80fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1661029a565b3480156104c6575f80fd5b5061037a7f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b981565b6102ce6104fc366004613b70565b611257565b6102ce61050f366004613839565b611456565b34801561051f575f80fd5b506102ce61052e366004613bbb565b611654565b34801561053e575f80fd5b506102ce61054d36600461389b565b6118d2565b34801561055d575f80fd5b506102ce611922565b348015610571575f80fd5b5061029a610580366004613905565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156105d4575f80fd5b506102ce6105e3366004613761565b611954565b3480156105f3575f80fd5b5061037a5f81565b348015610606575f80fd5b5061037a610b4081565b34801561061b575f80fd5b5061037a7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b34801561064e575f80fd5b506102ce61065d366004613c46565b611a6a565b34801561066d575f80fd5b506103ec6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b3480156106b5575f80fd5b506102ce6106c4366004613761565b611c05565b3480156106d4575f80fd5b5061037a60045481565b3480156106e9575f80fd5b506102ce6106f8366004613cd8565b611cdd565b6102ce61070b366004613d18565b611f47565b6102ce61071e366004613d86565b61211f565b34801561072e575f80fd5b506102ce61073d366004613905565b61226f565b34801561074d575f80fd5b505f54610315906001600160a01b031681565b34801561076b575f80fd5b5061037a7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061082557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b5f610835816122b2565b6001600160a01b03821661085c5760405163d92e233d60e01b815260040160405180910390fd5b6002546001600160a01b03161561089f576040517f0c8dc01600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108c97f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b9836122bc565b5050600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b61090d6123a6565b825f03610946576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03841661096d5760405163d92e233d60e01b815260040160405180910390fd5b610b4061097d6060830183613e0b565b905011156109db576109926060820182613e0b565b6040517f9fcf788e0000000000000000000000000000000000000000000000000000000081526109d29250610b4090600401918252602082015260400190565b60405180910390fd5b621e848081608001351115610a2b576040517fec86721e00000000000000000000000000000000000000000000000000000000815260808201356004820152621e848060248201526044016109d2565b5f610a34612404565b9050610a3f81612530565b610a4a338486612573565b846001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c868686604051610a9193929190613f7e565b60405180910390a35050505050565b610aa86123a6565b610ab86040820160208301613fb2565b15610aef576040517f19b4bff200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038416610b165760405163d92e233d60e01b815260040160405180910390fd5b5f610b246060830183613e0b565b610b2f915084613fcd565b9050610b40811115610b78576040517f9fcf788e00000000000000000000000000000000000000000000000000000000815260048101829052610b4060248201526044016109d2565b621e848082608001351115610bc8576040517fec86721e00000000000000000000000000000000000000000000000000000000815260808301356004820152621e848060248201526044016109d2565b5f610bd1612404565b9050610bdc81612530565b856001600160a01b0316336001600160a01b03167fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974878787604051610c2393929190614005565b60405180910390a3505050505050565b610c3b6123a6565b815f03610c74576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038316610c9b5760405163d92e233d60e01b815260040160405180910390fd5b610b40610cab6060830183613e0b565b90501115610cc0576109926060820182613e0b565b621e848081608001351115610d10576040517fec86721e00000000000000000000000000000000000000000000000000000000815260808201356004820152621e848060248201526044016109d2565b5f610d19612404565b9050610d258382612690565b6001546040515f916001600160a01b03169085908381818185875af1925050503d805f8114610d6f576040519150601f19603f3d011682016040523d82523d5f602084013e610d74565b606091505b5050905080610daf576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b846001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c865f87604051610a9193929190613f7e565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610e2f816122b2565b610e3983836122bc565b50505050565b6001600160a01b0381163314610e81576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e8b82826126df565b505050565b6060610e9a6127a1565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb610ec4816122b2565b610ecc6123a6565b6001600160a01b038516610ef35760405163d92e233d60e01b815260040160405180910390fd5b60605f610f036020890189613761565b6001600160a01b031603610f2357610f1c868686612822565b9050610f32565b610f2f878787876128d1565b90505b856001600160a01b03167f373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546348787604051610f6f9392919061402a565b60405180910390a2915050610fa360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b949350505050565b610fb36123a6565b835f03610fec576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0385166110135760405163d92e233d60e01b815260040160405180910390fd5b610b406110236060830183613e0b565b61102e915084613fcd565b111561108a576110416060820182613e0b565b61104c915083613fcd565b6040517f9fcf788e0000000000000000000000000000000000000000000000000000000081526004810191909152610b4060248201526044016109d2565b621e8480816080013511156110da576040517fec86721e00000000000000000000000000000000000000000000000000000000815260808201356004820152621e848060248201526044016109d2565b5f6110e3612404565b90506110ef8582612690565b6001546040515f916001600160a01b03169087908381818185875af1925050503d805f8114611139576040519150601f19603f3d011682016040523d82523d5f602084013e61113e565b606091505b5050905080611179576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b866001600160a01b0316336001600160a01b03167fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f885f8989896040516111c4959493929190614043565b60405180910390a350505050505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6111ff816122b2565b611207612976565b50565b611212612a06565b61121b82612ad6565b6112258282612ae0565b5050565b5f611232612be1565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b61125f6123a6565b345f03611298576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382166112bf5760405163d92e233d60e01b815260040160405180910390fd5b610b406112cf6060830183613e0b565b905011156112e4576109926060820182613e0b565b621e848081608001351115611334576040517fec86721e00000000000000000000000000000000000000000000000000000000815260808201356004820152621e848060248201526044016109d2565b5f61133d612c43565b90508015611377576040517f394836a400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040515f916001600160a01b03169034908381818185875af1925050503d805f81146113c1576040519150601f19603f3d011682016040523d82523d5f602084013e6113c6565b606091505b5050905080611401576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b836001600160a01b0316336001600160a01b03167fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c345f8760405161144893929190613f7e565b60405180910390a350505050565b61145e6123a6565b345f03611497576040517f7671265e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0384166114be5760405163d92e233d60e01b815260040160405180910390fd5b610b406114ce6060830183613e0b565b6114d9915084613fcd565b11156114ec576110416060820182613e0b565b621e84808160800135111561153c576040517fec86721e00000000000000000000000000000000000000000000000000000000815260808201356004820152621e848060248201526044016109d2565b5f611545612c43565b9050801561157f576040517f394836a400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040515f916001600160a01b03169034908381818185875af1925050503d805f81146115c9576040519150601f19603f3d011682016040523d82523d5f602084013e6115ce565b606091505b5050905080611609576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b856001600160a01b0316336001600160a01b03167fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f345f898989604051610c23959493929190614043565b61165c6127a1565b7f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b9611686816122b2565b61168e6123a6565b835f036116c7576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0385166116ee5760405163d92e233d60e01b815260040160405180910390fd5b6116f88686612c50565b61172e576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6117426001600160a01b0387168686612d5e565b5f6117506020890189613761565b6001600160a01b03160361176f57611769858484612822565b5061177d565b61177b878685856128d1565b505b6117878686612c50565b6117bd576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f906001600160a01b038816906370a0823190602401602060405180830381865afa15801561181a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061183e9190614089565b90508015611850576118508782612e5e565b856001600160a01b0316876001600160a01b03167f29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b73828787876040516118979392919061402a565b60405180910390a350506118ca60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050505050565b5f6118dc816122b2565b600480549083905560408051828152602081018590527fe60a2882aff7841a7d0492b804d46ff2d9b07ea353dbc3791e8276b729e40cd8910160405180910390a1505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61194c816122b2565b611207613078565b5f61195e816122b2565b6001600160a01b0382166119855760405163d92e233d60e01b815260040160405180910390fd5b6001546119bc907f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb906001600160a01b03166126df565b506119e77f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb836122bc565b50600154604080516001600160a01b03928316815291841660208301527f3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579910160405180910390a150600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b611a726127a1565b7f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b9611a9c816122b2565b611aa46123a6565b845f03611add576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038616611b045760405163d92e233d60e01b815260040160405180910390fd5b611b186001600160a01b03881687876130f1565b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0387169063c9028a3690611b5d9085906004016140fb565b5f604051808303815f87803b158015611b74575f80fd5b505af1158015611b86573d5f803e3d5ffd5b50505050866001600160a01b0316866001600160a01b03167fde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e03587878787604051611bd3949392919061410d565b60405180910390a3506118ca60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f611c0f816122b2565b6001600160a01b038216611c365760405163d92e233d60e01b815260040160405180910390fd5b5f546001600160a01b031615611c78576040517fb337f37800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611ca27f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b9836122bc565b50505f80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b5f611ce6613122565b805490915060ff68010000000000000000820416159067ffffffffffffffff165f81158015611d125750825b90505f8267ffffffffffffffff166001148015611d2e5750303b155b905081158015611d3c575080155b15611d73576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315611dd45784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b0388161580611df157506001600160a01b038716155b15611e0f5760405163d92e233d60e01b815260040160405180910390fd5b611e1761314a565b611e1f613152565b611e2761314a565b611e2f61314a565b611e395f876122bc565b50611e647f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a876122bc565b50611e8f7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb896122bc565b50600180546001600160a01b03808b167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560038054928a16929091169190911790555f6004558315611f3d5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b611f4f6127a1565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb611f79816122b2565b611f816123a6565b6001600160a01b038516611fa85760405163d92e233d60e01b815260040160405180910390fd5b5f856001600160a01b0316346040515f6040518083038185875af1925050503d805f8114611ff1576040519150601f19603f3d011682016040523d82523d5f602084013e611ff6565b606091505b5050905080612031576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0387169063c9028a36906120769086906004016140fb565b5f604051808303815f87803b15801561208d575f80fd5b505af115801561209f573d5f803e3d5ffd5b505050505f6001600160a01b0316866001600160a01b03167fde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035348888886040516120ec949392919061410d565b60405180910390a35050610e3960017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6121276123a6565b845f03612160576040517f951e19ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0386166121875760405163d92e233d60e01b815260040160405180910390fd5b610b406121976060830183613e0b565b6121a2915084613fcd565b11156121b5576110416060820182613e0b565b621e848081608001351115612205576040517fec86721e00000000000000000000000000000000000000000000000000000000815260808201356004820152621e848060248201526044016109d2565b5f61220e612404565b905061221981612530565b612224338688612573565b866001600160a01b0316336001600160a01b03167fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f88888888886040516111c4959493929190614043565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546122a8816122b2565b610e3983836126df565b6112078133613162565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1661239d575f848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556123533390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610825565b5f915050610825565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615612402576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f8061240e612c43565b9050805f0361241e575f91505090565b6004545f03612459576040517f394836a400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004549150813410156124a1576040517fa458261b000000000000000000000000000000000000000000000000000000008152600481018390523460248201526044016109d2565b6001546040515f916001600160a01b03169084908381818185875af1925050503d805f81146124eb576040519150601f19603f3d011682016040523d82523d5f602084013e6124f0565b606091505b505090508061252b576040517f4033e4e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505090565b80341115611207576040517f8afe4db4000000000000000000000000000000000000000000000000000000008152600481018290523460248201526044016109d2565b6003546001600160a01b03908116908316036125bb576040517fe4dd681d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301529091169063d936547e90602401602060405180830381865afa15801561261b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061263f9190614143565b612675576040517fac2175f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54610e8b906001600160a01b0384811691869116846131ee565b5f61269b8284613fcd565b9050803414610e8b576040517f84222ac7000000000000000000000000000000000000000000000000000000008152600481018290523460248201526044016109d2565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff161561239d575f848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610825565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0161281c576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b606061282e8383613227565b5f80856001600160a01b031634868660405161284b92919061415e565b5f6040518083038185875af1925050503d805f8114612885576040519150601f19603f3d011682016040523d82523d5f602084013e61288a565b606091505b5091509150816128c6576040517facfdb44400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9150505b9392505050565b6060836001600160a01b031663676cc054348786866040518563ffffffff1660e01b81526004016129049392919061416d565b5f6040518083038185885af115801561291f573d5f803e3d5ffd5b50505050506040513d5f823e601f3d908101601f191682016040526129479190810190614197565b95945050505050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b61297e613327565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480612a9f57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612a937f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15612402576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611225816122b2565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612b3a575060408051601f3d908101601f19168201909252612b3791810190614089565b60015b612b7b576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016109d2565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612bd7576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016109d2565b610e8b8383613382565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612402576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001805c90818101905d90565b604080516001600160a01b0383811660248301525f604480840182905284518085039091018152606490930184526020830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b300000000000000000000000000000000000000000000000000000000179052925183928392871691612cd9919061420c565b5f604051808303815f865af19150503d805f8114612d12576040519150601f19603f3d011682016040523d82523d5f602084013e612d17565b606091505b509150915081612d2c57600192505050610825565b805115612d53575f81806020019051810190612d489190614143565b935061082592505050565b506001949350505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f095ea7b300000000000000000000000000000000000000000000000000000000179052612ddd84826133d7565b610e39576040516001600160a01b0384811660248301525f6044830152612e5491869182169063095ea7b3906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050613420565b610e398482613420565b6003546001600160a01b0390811690831603612fa5576002546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018390529083169063095ea7b3906044016020604051808303815f875af1158015612edd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f019190614143565b612f37576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546040517f743e0c9b000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b039091169063743e0c9b906024015f604051808303815f87803b158015612f93575f80fd5b505af11580156118ca573d5f803e3d5ffd5b5f546040517fd936547e0000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301529091169063d936547e90602401602060405180830381865afa158015613005573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130299190614143565b61305f576040517fac2175f100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54611225906001600160a01b038481169116836130f1565b6130806123a6565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336129e8565b6040516001600160a01b03838116602483015260448201839052610e8b91859182169063a9059cbb90606401612e0d565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610825565b6124026134a5565b61315a6134a5565b6124026134e3565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16611225576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044016109d2565b6040516001600160a01b038481166024830152838116604483015260648201839052610e399186918216906323b872dd90608401612e0d565b600481106112255781357f98933fac000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216016132ac576040517fed69977500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f36fd75ca000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821601610e8b576040517ff3459a9600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16612402576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61338b826134eb565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156133cf57610e8b8282613592565b6112256135fb565b5f805f8060205f8651602088015f8a5af192503d91505f519050828015613416575081156134085780600114613416565b5f866001600160a01b03163b115b9695505050505050565b5f8060205f8451602086015f885af18061343f576040513d5f823e3d81fd5b50505f513d91508115613456578060011415613463565b6001600160a01b0384163b155b15610e39576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016109d2565b6134ad613633565b612402576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6129506134a5565b806001600160a01b03163b5f03613539576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016109d2565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f80846001600160a01b0316846040516135ae919061420c565b5f60405180830381855af49150503d805f81146135e6576040519150601f19603f3d011682016040523d82523d5f602084013e6135eb565b606091505b5091509150612947858383613651565b3415612402576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61363c613122565b5468010000000000000000900460ff16919050565b60608261366657613661826136c6565b6128ca565b815115801561367d57506001600160a01b0384163b155b156136bf576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016109d2565b50806128ca565b8051156136d557805160208201fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60208284031215613717575f80fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146128ca575f80fd5b80356001600160a01b038116811461375c575f80fd5b919050565b5f60208284031215613771575f80fd5b6128ca82613746565b5f60a0828403121561378a575f80fd5b50919050565b5f805f80608085870312156137a3575f80fd5b6137ac85613746565b9350602085013592506137c160408601613746565b9150606085013567ffffffffffffffff8111156137dc575f80fd5b6137e88782880161377a565b91505092959194509250565b5f8083601f840112613804575f80fd5b50813567ffffffffffffffff81111561381b575f80fd5b602083019150836020828501011115613832575f80fd5b9250929050565b5f805f806060858703121561384c575f80fd5b61385585613746565b9350602085013567ffffffffffffffff811115613870575f80fd5b61387c878288016137f4565b909450925050604085013567ffffffffffffffff8111156137dc575f80fd5b5f602082840312156138ab575f80fd5b5035919050565b5f805f606084860312156138c4575f80fd5b6138cd84613746565b925060208401359150604084013567ffffffffffffffff8111156138ef575f80fd5b6138fb8682870161377a565b9150509250925092565b5f8060408385031215613916575f80fd5b8235915061392660208401613746565b90509250929050565b5f6020828403121561378a575f80fd5b5f805f8060608587031215613952575f80fd5b61395c868661392f565b935061396a60208601613746565b9250604085013567ffffffffffffffff811115613985575f80fd5b613991878288016137f4565b95989497509550505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6128ca602083018461399d565b5f805f805f608086880312156139f1575f80fd5b6139fa86613746565b945060208601359350604086013567ffffffffffffffff811115613a1c575f80fd5b613a28888289016137f4565b909450925050606086013567ffffffffffffffff811115613a47575f80fd5b613a538882890161377a565b9150509295509295909350565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff81118282101715613ab657613ab6613a60565b604052919050565b5f67ffffffffffffffff821115613ad757613ad7613a60565b50601f01601f191660200190565b5f8060408385031215613af6575f80fd5b613aff83613746565b9150602083013567ffffffffffffffff811115613b1a575f80fd5b8301601f81018513613b2a575f80fd5b8035613b3d613b3882613abe565b613a8d565b818152866020838501011115613b51575f80fd5b816020840160208301375f602083830101528093505050509250929050565b5f8060408385031215613b81575f80fd5b613b8a83613746565b9150602083013567ffffffffffffffff811115613ba5575f80fd5b613bb18582860161377a565b9150509250929050565b5f805f805f8060a08789031215613bd0575f80fd5b613bda888861392f565b9550613be860208801613746565b9450613bf660408801613746565b935060608701359250608087013567ffffffffffffffff811115613c18575f80fd5b613c2489828a016137f4565b979a9699509497509295939492505050565b5f6080828403121561378a575f80fd5b5f805f805f8060a08789031215613c5b575f80fd5b613c6487613746565b9550613c7260208801613746565b945060408701359350606087013567ffffffffffffffff811115613c94575f80fd5b613ca089828a016137f4565b909450925050608087013567ffffffffffffffff811115613cbf575f80fd5b613ccb89828a01613c36565b9150509295509295509295565b5f805f60608486031215613cea575f80fd5b613cf384613746565b9250613d0160208501613746565b9150613d0f60408501613746565b90509250925092565b5f805f8060608587031215613d2b575f80fd5b613d3485613746565b9350602085013567ffffffffffffffff811115613d4f575f80fd5b613d5b878288016137f4565b909450925050604085013567ffffffffffffffff811115613d7a575f80fd5b6137e887828801613c36565b5f805f805f8060a08789031215613d9b575f80fd5b613da487613746565b955060208701359450613db960408801613746565b9350606087013567ffffffffffffffff811115613dd4575f80fd5b613de089828a016137f4565b909450925050608087013567ffffffffffffffff811115613dff575f80fd5b613ccb89828a0161377a565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613e3e575f80fd5b83018035915067ffffffffffffffff821115613e58575f80fd5b602001915036819003821315613832575f80fd5b8015158114611207575f80fd5b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613eac575f80fd5b830160208101925035905067ffffffffffffffff811115613ecb575f80fd5b803603821315613832575f80fd5b81835281816020850137505f602082840101525f6020601f19601f840116840101905092915050565b6001600160a01b03613f1382613746565b1682525f6020820135613f2581613e6c565b151560208401526001600160a01b03613f4060408401613746565b166040840152613f536060830183613e79565b60a06060860152613f6860a086018284613ed9565b6080948501359590940194909452509092915050565b8381526001600160a01b0383166020820152608060408201525f608082015260a060608201525f61294760a0830184613f02565b5f60208284031215613fc2575f80fd5b81356128ca81613e6c565b80820180821115610825577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b604081525f614018604083018587613ed9565b82810360208401526134168185613f02565b838152604060208201525f612947604083018486613ed9565b8581526001600160a01b0385166020820152608060408201525f61406b608083018587613ed9565b828103606084015261407d8185613f02565b98975050505050505050565b5f60208284031215614099575f80fd5b5051919050565b6001600160a01b036140b182613746565b1682526001600160a01b036140c860208301613746565b166020830152604081810135908301525f6140e66060830183613e79565b60806060860152612947608086018284613ed9565b602081525f6128ca60208301846140a0565b848152606060208201525f614126606083018587613ed9565b828103604084015261413881856140a0565b979650505050505050565b5f60208284031215614153575f80fd5b81516128ca81613e6c565b818382375f9101908152919050565b6001600160a01b0361417e85613746565b168152604060208201525f612947604083018486613ed9565b5f602082840312156141a7575f80fd5b815167ffffffffffffffff8111156141bd575f80fd5b8201601f810184136141cd575f80fd5b80516141db613b3882613abe565b8181528560208385010111156141ef575f80fd5b8160208401602083015e5f91810160200191909152949350505050565b5f82518060208501845e5f92019182525091905056fea2646970667358221220d7c68a61b0d30441cd810bf786e4a8e8e1709e17be39209945fd7ac8ea5c978f64736f6c634300081a0033",
}

// GatewayEVMUpgradeTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMUpgradeTestMetaData.ABI instead.
var GatewayEVMUpgradeTestABI = GatewayEVMUpgradeTestMetaData.ABI

// GatewayEVMUpgradeTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMUpgradeTestMetaData.Bin instead.
var GatewayEVMUpgradeTestBin = GatewayEVMUpgradeTestMetaData.Bin

// DeployGatewayEVMUpgradeTest deploys a new Ethereum contract, binding an instance of GatewayEVMUpgradeTest to it.
func DeployGatewayEVMUpgradeTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVMUpgradeTest, error) {
	parsed, err := GatewayEVMUpgradeTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMUpgradeTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVMUpgradeTest{GatewayEVMUpgradeTestCaller: GatewayEVMUpgradeTestCaller{contract: contract}, GatewayEVMUpgradeTestTransactor: GatewayEVMUpgradeTestTransactor{contract: contract}, GatewayEVMUpgradeTestFilterer: GatewayEVMUpgradeTestFilterer{contract: contract}}, nil
}

// GatewayEVMUpgradeTest is an auto generated Go binding around an Ethereum contract.
type GatewayEVMUpgradeTest struct {
	GatewayEVMUpgradeTestCaller     // Read-only binding to the contract
	GatewayEVMUpgradeTestTransactor // Write-only binding to the contract
	GatewayEVMUpgradeTestFilterer   // Log filterer for contract events
}

// GatewayEVMUpgradeTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMUpgradeTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMUpgradeTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMUpgradeTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMUpgradeTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMUpgradeTestSession struct {
	Contract     *GatewayEVMUpgradeTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GatewayEVMUpgradeTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMUpgradeTestCallerSession struct {
	Contract *GatewayEVMUpgradeTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// GatewayEVMUpgradeTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMUpgradeTestTransactorSession struct {
	Contract     *GatewayEVMUpgradeTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// GatewayEVMUpgradeTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestRaw struct {
	Contract *GatewayEVMUpgradeTest // Generic contract binding to access the raw methods on
}

// GatewayEVMUpgradeTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestCallerRaw struct {
	Contract *GatewayEVMUpgradeTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMUpgradeTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMUpgradeTestTransactorRaw struct {
	Contract *GatewayEVMUpgradeTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVMUpgradeTest creates a new instance of GatewayEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayEVMUpgradeTest(address common.Address, backend bind.ContractBackend) (*GatewayEVMUpgradeTest, error) {
	contract, err := bindGatewayEVMUpgradeTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTest{GatewayEVMUpgradeTestCaller: GatewayEVMUpgradeTestCaller{contract: contract}, GatewayEVMUpgradeTestTransactor: GatewayEVMUpgradeTestTransactor{contract: contract}, GatewayEVMUpgradeTestFilterer: GatewayEVMUpgradeTestFilterer{contract: contract}}, nil
}

// NewGatewayEVMUpgradeTestCaller creates a new read-only instance of GatewayEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayEVMUpgradeTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMUpgradeTestCaller, error) {
	contract, err := bindGatewayEVMUpgradeTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestCaller{contract: contract}, nil
}

// NewGatewayEVMUpgradeTestTransactor creates a new write-only instance of GatewayEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayEVMUpgradeTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMUpgradeTestTransactor, error) {
	contract, err := bindGatewayEVMUpgradeTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestTransactor{contract: contract}, nil
}

// NewGatewayEVMUpgradeTestFilterer creates a new log filterer instance of GatewayEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayEVMUpgradeTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMUpgradeTestFilterer, error) {
	contract, err := bindGatewayEVMUpgradeTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestFilterer{contract: contract}, nil
}

// bindGatewayEVMUpgradeTest binds a generic wrapper to an already deployed contract.
func bindGatewayEVMUpgradeTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMUpgradeTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMUpgradeTest.Contract.GatewayEVMUpgradeTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.GatewayEVMUpgradeTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.GatewayEVMUpgradeTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMUpgradeTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.contract.Transact(opts, method, params...)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) ASSETHANDLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "ASSET_HANDLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.ASSETHANDLERROLE(&_GatewayEVMUpgradeTest.CallOpts)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.ASSETHANDLERROLE(&_GatewayEVMUpgradeTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.DEFAULTADMINROLE(&_GatewayEVMUpgradeTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.DEFAULTADMINROLE(&_GatewayEVMUpgradeTest.CallOpts)
}

// MAXPAYLOADSIZE is a free data retrieval call binding the contract method 0xa2ba1934.
//
// Solidity: function MAX_PAYLOAD_SIZE() view returns(uint256)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) MAXPAYLOADSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "MAX_PAYLOAD_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPAYLOADSIZE is a free data retrieval call binding the contract method 0xa2ba1934.
//
// Solidity: function MAX_PAYLOAD_SIZE() view returns(uint256)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) MAXPAYLOADSIZE() (*big.Int, error) {
	return _GatewayEVMUpgradeTest.Contract.MAXPAYLOADSIZE(&_GatewayEVMUpgradeTest.CallOpts)
}

// MAXPAYLOADSIZE is a free data retrieval call binding the contract method 0xa2ba1934.
//
// Solidity: function MAX_PAYLOAD_SIZE() view returns(uint256)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) MAXPAYLOADSIZE() (*big.Int, error) {
	return _GatewayEVMUpgradeTest.Contract.MAXPAYLOADSIZE(&_GatewayEVMUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.PAUSERROLE(&_GatewayEVMUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.PAUSERROLE(&_GatewayEVMUpgradeTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) TSSROLE() ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.TSSROLE(&_GatewayEVMUpgradeTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) TSSROLE() ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.TSSROLE(&_GatewayEVMUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayEVMUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_GatewayEVMUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayEVMUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_GatewayEVMUpgradeTest.CallOpts)
}

// AdditionalActionFeeWei is a free data retrieval call binding the contract method 0xb0107214.
//
// Solidity: function additionalActionFeeWei() view returns(uint256)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) AdditionalActionFeeWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "additionalActionFeeWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdditionalActionFeeWei is a free data retrieval call binding the contract method 0xb0107214.
//
// Solidity: function additionalActionFeeWei() view returns(uint256)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) AdditionalActionFeeWei() (*big.Int, error) {
	return _GatewayEVMUpgradeTest.Contract.AdditionalActionFeeWei(&_GatewayEVMUpgradeTest.CallOpts)
}

// AdditionalActionFeeWei is a free data retrieval call binding the contract method 0xb0107214.
//
// Solidity: function additionalActionFeeWei() view returns(uint256)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) AdditionalActionFeeWei() (*big.Int, error) {
	return _GatewayEVMUpgradeTest.Contract.AdditionalActionFeeWei(&_GatewayEVMUpgradeTest.CallOpts)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) Custody(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "custody")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Custody() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.Custody(&_GatewayEVMUpgradeTest.CallOpts)
}

// Custody is a free data retrieval call binding the contract method 0xdda79b75.
//
// Solidity: function custody() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) Custody() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.Custody(&_GatewayEVMUpgradeTest.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.GetRoleAdmin(&_GatewayEVMUpgradeTest.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.GetRoleAdmin(&_GatewayEVMUpgradeTest.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayEVMUpgradeTest.Contract.HasRole(&_GatewayEVMUpgradeTest.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayEVMUpgradeTest.Contract.HasRole(&_GatewayEVMUpgradeTest.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Paused() (bool, error) {
	return _GatewayEVMUpgradeTest.Contract.Paused(&_GatewayEVMUpgradeTest.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) Paused() (bool, error) {
	return _GatewayEVMUpgradeTest.Contract.Paused(&_GatewayEVMUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.ProxiableUUID(&_GatewayEVMUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayEVMUpgradeTest.Contract.ProxiableUUID(&_GatewayEVMUpgradeTest.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayEVMUpgradeTest.Contract.SupportsInterface(&_GatewayEVMUpgradeTest.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayEVMUpgradeTest.Contract.SupportsInterface(&_GatewayEVMUpgradeTest.CallOpts, interfaceId)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) TssAddress() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.TssAddress(&_GatewayEVMUpgradeTest.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) TssAddress() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.TssAddress(&_GatewayEVMUpgradeTest.CallOpts)
}

// ZetaConnector is a free data retrieval call binding the contract method 0x57bec62f.
//
// Solidity: function zetaConnector() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) ZetaConnector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "zetaConnector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaConnector is a free data retrieval call binding the contract method 0x57bec62f.
//
// Solidity: function zetaConnector() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) ZetaConnector() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.ZetaConnector(&_GatewayEVMUpgradeTest.CallOpts)
}

// ZetaConnector is a free data retrieval call binding the contract method 0x57bec62f.
//
// Solidity: function zetaConnector() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) ZetaConnector() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.ZetaConnector(&_GatewayEVMUpgradeTest.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayEVMUpgradeTest.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) ZetaToken() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.ZetaToken(&_GatewayEVMUpgradeTest.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestCallerSession) ZetaToken() (common.Address, error) {
	return _GatewayEVMUpgradeTest.Contract.ZetaToken(&_GatewayEVMUpgradeTest.CallOpts)
}

// Call is a paid mutator transaction binding the contract method 0x1becceb4.
//
// Solidity: function call(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Call(opts *bind.TransactOpts, receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "call", receiver, payload, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x1becceb4.
//
// Solidity: function call(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Call(receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Call(&_GatewayEVMUpgradeTest.TransactOpts, receiver, payload, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x1becceb4.
//
// Solidity: function call(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Call(receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Call(&_GatewayEVMUpgradeTest.TransactOpts, receiver, payload, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0x102614b0.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Deposit(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, asset common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "deposit", receiver, amount, asset, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0x102614b0.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Deposit(receiver common.Address, amount *big.Int, asset common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Deposit(&_GatewayEVMUpgradeTest.TransactOpts, receiver, amount, asset, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0x102614b0.
//
// Solidity: function deposit(address receiver, uint256 amount, address asset, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Deposit(receiver common.Address, amount *big.Int, asset common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Deposit(&_GatewayEVMUpgradeTest.TransactOpts, receiver, amount, asset, revertOptions)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x282906ed.
//
// Solidity: function deposit(address receiver, uint256 amount, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Deposit0(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "deposit0", receiver, amount, revertOptions)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x282906ed.
//
// Solidity: function deposit(address receiver, uint256 amount, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Deposit0(receiver common.Address, amount *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Deposit0(&_GatewayEVMUpgradeTest.TransactOpts, receiver, amount, revertOptions)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x282906ed.
//
// Solidity: function deposit(address receiver, uint256 amount, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Deposit0(receiver common.Address, amount *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Deposit0(&_GatewayEVMUpgradeTest.TransactOpts, receiver, amount, revertOptions)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x726ac97c.
//
// Solidity: function deposit(address receiver, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Deposit1(opts *bind.TransactOpts, receiver common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "deposit1", receiver, revertOptions)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x726ac97c.
//
// Solidity: function deposit(address receiver, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Deposit1(receiver common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Deposit1(&_GatewayEVMUpgradeTest.TransactOpts, receiver, revertOptions)
}

// Deposit1 is a paid mutator transaction binding the contract method 0x726ac97c.
//
// Solidity: function deposit(address receiver, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Deposit1(receiver common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Deposit1(&_GatewayEVMUpgradeTest.TransactOpts, receiver, revertOptions)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x397e375c.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) DepositAndCall(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "depositAndCall", receiver, amount, payload, revertOptions)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x397e375c.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) DepositAndCall(receiver common.Address, amount *big.Int, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.DepositAndCall(&_GatewayEVMUpgradeTest.TransactOpts, receiver, amount, payload, revertOptions)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x397e375c.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) DepositAndCall(receiver common.Address, amount *big.Int, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.DepositAndCall(&_GatewayEVMUpgradeTest.TransactOpts, receiver, amount, payload, revertOptions)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0x744b9b8b.
//
// Solidity: function depositAndCall(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) DepositAndCall0(opts *bind.TransactOpts, receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "depositAndCall0", receiver, payload, revertOptions)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0x744b9b8b.
//
// Solidity: function depositAndCall(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) DepositAndCall0(receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.DepositAndCall0(&_GatewayEVMUpgradeTest.TransactOpts, receiver, payload, revertOptions)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0x744b9b8b.
//
// Solidity: function depositAndCall(address receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) DepositAndCall0(receiver common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.DepositAndCall0(&_GatewayEVMUpgradeTest.TransactOpts, receiver, payload, revertOptions)
}

// DepositAndCall1 is a paid mutator transaction binding the contract method 0xd09e3b78.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) DepositAndCall1(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, asset common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "depositAndCall1", receiver, amount, asset, payload, revertOptions)
}

// DepositAndCall1 is a paid mutator transaction binding the contract method 0xd09e3b78.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) DepositAndCall1(receiver common.Address, amount *big.Int, asset common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.DepositAndCall1(&_GatewayEVMUpgradeTest.TransactOpts, receiver, amount, asset, payload, revertOptions)
}

// DepositAndCall1 is a paid mutator transaction binding the contract method 0xd09e3b78.
//
// Solidity: function depositAndCall(address receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) DepositAndCall1(receiver common.Address, amount *big.Int, asset common.Address, payload []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.DepositAndCall1(&_GatewayEVMUpgradeTest.TransactOpts, receiver, amount, asset, payload, revertOptions)
}

// Execute is a paid mutator transaction binding the contract method 0x38e22527.
//
// Solidity: function execute((address) messageContext, address destination, bytes data) payable returns(bytes)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Execute(opts *bind.TransactOpts, messageContext MessageContext, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "execute", messageContext, destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x38e22527.
//
// Solidity: function execute((address) messageContext, address destination, bytes data) payable returns(bytes)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Execute(messageContext MessageContext, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Execute(&_GatewayEVMUpgradeTest.TransactOpts, messageContext, destination, data)
}

// Execute is a paid mutator transaction binding the contract method 0x38e22527.
//
// Solidity: function execute((address) messageContext, address destination, bytes data) payable returns(bytes)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Execute(messageContext MessageContext, destination common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Execute(&_GatewayEVMUpgradeTest.TransactOpts, messageContext, destination, data)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0xcb7ba8e5.
//
// Solidity: function executeRevert(address destination, bytes data, (address,address,uint256,bytes) revertContext) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) ExecuteRevert(opts *bind.TransactOpts, destination common.Address, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "executeRevert", destination, data, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0xcb7ba8e5.
//
// Solidity: function executeRevert(address destination, bytes data, (address,address,uint256,bytes) revertContext) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) ExecuteRevert(destination common.Address, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.ExecuteRevert(&_GatewayEVMUpgradeTest.TransactOpts, destination, data, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0xcb7ba8e5.
//
// Solidity: function executeRevert(address destination, bytes data, (address,address,uint256,bytes) revertContext) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) ExecuteRevert(destination common.Address, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.ExecuteRevert(&_GatewayEVMUpgradeTest.TransactOpts, destination, data, revertContext)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x7bbe9afa.
//
// Solidity: function executeWithERC20((address) messageContext, address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) ExecuteWithERC20(opts *bind.TransactOpts, messageContext MessageContext, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "executeWithERC20", messageContext, token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x7bbe9afa.
//
// Solidity: function executeWithERC20((address) messageContext, address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) ExecuteWithERC20(messageContext MessageContext, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.ExecuteWithERC20(&_GatewayEVMUpgradeTest.TransactOpts, messageContext, token, to, amount, data)
}

// ExecuteWithERC20 is a paid mutator transaction binding the contract method 0x7bbe9afa.
//
// Solidity: function executeWithERC20((address) messageContext, address token, address to, uint256 amount, bytes data) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) ExecuteWithERC20(messageContext MessageContext, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.ExecuteWithERC20(&_GatewayEVMUpgradeTest.TransactOpts, messageContext, token, to, amount, data)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.GrantRole(&_GatewayEVMUpgradeTest.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.GrantRole(&_GatewayEVMUpgradeTest.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address tssAddress_, address zetaToken_, address admin_) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Initialize(opts *bind.TransactOpts, tssAddress_ common.Address, zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "initialize", tssAddress_, zetaToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address tssAddress_, address zetaToken_, address admin_) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Initialize(tssAddress_ common.Address, zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Initialize(&_GatewayEVMUpgradeTest.TransactOpts, tssAddress_, zetaToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address tssAddress_, address zetaToken_, address admin_) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Initialize(tssAddress_ common.Address, zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Initialize(&_GatewayEVMUpgradeTest.TransactOpts, tssAddress_, zetaToken_, admin_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Pause() (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Pause(&_GatewayEVMUpgradeTest.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Pause() (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Pause(&_GatewayEVMUpgradeTest.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.RenounceRole(&_GatewayEVMUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.RenounceRole(&_GatewayEVMUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xaa0c0fc1.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) RevertWithERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "revertWithERC20", token, to, amount, data, revertContext)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xaa0c0fc1.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) RevertWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.RevertWithERC20(&_GatewayEVMUpgradeTest.TransactOpts, token, to, amount, data, revertContext)
}

// RevertWithERC20 is a paid mutator transaction binding the contract method 0xaa0c0fc1.
//
// Solidity: function revertWithERC20(address token, address to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) RevertWithERC20(token common.Address, to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.RevertWithERC20(&_GatewayEVMUpgradeTest.TransactOpts, token, to, amount, data, revertContext)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.RevokeRole(&_GatewayEVMUpgradeTest.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.RevokeRole(&_GatewayEVMUpgradeTest.TransactOpts, role, account)
}

// SetConnector is a paid mutator transaction binding the contract method 0x10188aef.
//
// Solidity: function setConnector(address zetaConnector_) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) SetConnector(opts *bind.TransactOpts, zetaConnector_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "setConnector", zetaConnector_)
}

// SetConnector is a paid mutator transaction binding the contract method 0x10188aef.
//
// Solidity: function setConnector(address zetaConnector_) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) SetConnector(zetaConnector_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.SetConnector(&_GatewayEVMUpgradeTest.TransactOpts, zetaConnector_)
}

// SetConnector is a paid mutator transaction binding the contract method 0x10188aef.
//
// Solidity: function setConnector(address zetaConnector_) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) SetConnector(zetaConnector_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.SetConnector(&_GatewayEVMUpgradeTest.TransactOpts, zetaConnector_)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address custody_) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) SetCustody(opts *bind.TransactOpts, custody_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "setCustody", custody_)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address custody_) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) SetCustody(custody_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.SetCustody(&_GatewayEVMUpgradeTest.TransactOpts, custody_)
}

// SetCustody is a paid mutator transaction binding the contract method 0xae7a3a6f.
//
// Solidity: function setCustody(address custody_) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) SetCustody(custody_ common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.SetCustody(&_GatewayEVMUpgradeTest.TransactOpts, custody_)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) Unpause() (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Unpause(&_GatewayEVMUpgradeTest.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) Unpause() (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.Unpause(&_GatewayEVMUpgradeTest.TransactOpts)
}

// UpdateAdditionalActionFee is a paid mutator transaction binding the contract method 0x7c744253.
//
// Solidity: function updateAdditionalActionFee(uint256 newFeeWei) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) UpdateAdditionalActionFee(opts *bind.TransactOpts, newFeeWei *big.Int) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "updateAdditionalActionFee", newFeeWei)
}

// UpdateAdditionalActionFee is a paid mutator transaction binding the contract method 0x7c744253.
//
// Solidity: function updateAdditionalActionFee(uint256 newFeeWei) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) UpdateAdditionalActionFee(newFeeWei *big.Int) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.UpdateAdditionalActionFee(&_GatewayEVMUpgradeTest.TransactOpts, newFeeWei)
}

// UpdateAdditionalActionFee is a paid mutator transaction binding the contract method 0x7c744253.
//
// Solidity: function updateAdditionalActionFee(uint256 newFeeWei) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) UpdateAdditionalActionFee(newFeeWei *big.Int) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.UpdateAdditionalActionFee(&_GatewayEVMUpgradeTest.TransactOpts, newFeeWei)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) UpdateTSSAddress(opts *bind.TransactOpts, newTSSAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "updateTSSAddress", newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.UpdateTSSAddress(&_GatewayEVMUpgradeTest.TransactOpts, newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.UpdateTSSAddress(&_GatewayEVMUpgradeTest.TransactOpts, newTSSAddress)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.UpgradeToAndCall(&_GatewayEVMUpgradeTest.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayEVMUpgradeTest.Contract.UpgradeToAndCall(&_GatewayEVMUpgradeTest.TransactOpts, newImplementation, data)
}

// GatewayEVMUpgradeTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestCalledIterator struct {
	Event *GatewayEVMUpgradeTestCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestCalled represents a Called event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMUpgradeTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestCalledIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestCalled)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Called", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCalled is a log parse operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseCalled(log types.Log) (*GatewayEVMUpgradeTestCalled, error) {
	event := new(GatewayEVMUpgradeTestCalled)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestDepositedIterator struct {
	Event *GatewayEVMUpgradeTestDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestDeposited represents a Deposited event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestDeposited struct {
	Sender        common.Address
	Receiver      common.Address
	Amount        *big.Int
	Asset         common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterDeposited(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMUpgradeTestDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestDepositedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestDeposited, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestDeposited)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseDeposited(log types.Log) (*GatewayEVMUpgradeTestDeposited, error) {
	event := new(GatewayEVMUpgradeTestDeposited)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestDepositedAndCalledIterator is returned from FilterDepositedAndCalled and is used to iterate over the raw logs and unpacked data for DepositedAndCalled events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestDepositedAndCalledIterator struct {
	Event *GatewayEVMUpgradeTestDepositedAndCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestDepositedAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestDepositedAndCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestDepositedAndCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestDepositedAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestDepositedAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestDepositedAndCalled represents a DepositedAndCalled event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestDepositedAndCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Amount        *big.Int
	Asset         common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDepositedAndCalled is a free log retrieval operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterDepositedAndCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMUpgradeTestDepositedAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestDepositedAndCalledIterator{contract: _GatewayEVMUpgradeTest.contract, event: "DepositedAndCalled", logs: logs, sub: sub}, nil
}

// WatchDepositedAndCalled is a free log subscription operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchDepositedAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestDepositedAndCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestDepositedAndCalled)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositedAndCalled is a log parse operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseDepositedAndCalled(log types.Log) (*GatewayEVMUpgradeTestDepositedAndCalled, error) {
	event := new(GatewayEVMUpgradeTestDepositedAndCalled)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecutedIterator struct {
	Event *GatewayEVMUpgradeTestExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestExecuted represents a Executed event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMUpgradeTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestExecutedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestExecuted)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecuted is a log parse operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseExecuted(log types.Log) (*GatewayEVMUpgradeTestExecuted, error) {
	event := new(GatewayEVMUpgradeTestExecuted)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestExecutedV2Iterator is returned from FilterExecutedV2 and is used to iterate over the raw logs and unpacked data for ExecutedV2 events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecutedV2Iterator struct {
	Event *GatewayEVMUpgradeTestExecutedV2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestExecutedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestExecutedV2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestExecutedV2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestExecutedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestExecutedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestExecutedV2 represents a ExecutedV2 event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecutedV2 struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutedV2 is a free log retrieval operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterExecutedV2(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMUpgradeTestExecutedV2Iterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "ExecutedV2", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestExecutedV2Iterator{contract: _GatewayEVMUpgradeTest.contract, event: "ExecutedV2", logs: logs, sub: sub}, nil
}

// WatchExecutedV2 is a free log subscription operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchExecutedV2(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestExecutedV2, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "ExecutedV2", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestExecutedV2)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "ExecutedV2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutedV2 is a log parse operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseExecutedV2(log types.Log) (*GatewayEVMUpgradeTestExecutedV2, error) {
	event := new(GatewayEVMUpgradeTestExecutedV2)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "ExecutedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecutedWithERC20Iterator struct {
	Event *GatewayEVMUpgradeTestExecutedWithERC20 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestExecutedWithERC20)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestExecutedWithERC20)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMUpgradeTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestExecutedWithERC20Iterator{contract: _GatewayEVMUpgradeTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestExecutedWithERC20)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutedWithERC20 is a log parse operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMUpgradeTestExecutedWithERC20, error) {
	event := new(GatewayEVMUpgradeTestExecutedWithERC20)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestInitializedIterator struct {
	Event *GatewayEVMUpgradeTestInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestInitialized represents a Initialized event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayEVMUpgradeTestInitializedIterator, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestInitializedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestInitialized)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseInitialized(log types.Log) (*GatewayEVMUpgradeTestInitialized, error) {
	event := new(GatewayEVMUpgradeTestInitialized)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestPausedIterator struct {
	Event *GatewayEVMUpgradeTestPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestPaused represents a Paused event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterPaused(opts *bind.FilterOpts) (*GatewayEVMUpgradeTestPausedIterator, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestPausedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestPaused) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestPaused)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParsePaused(log types.Log) (*GatewayEVMUpgradeTestPaused, error) {
	event := new(GatewayEVMUpgradeTestPaused)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestRevertedIterator struct {
	Event *GatewayEVMUpgradeTestReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestReverted represents a Reverted event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestReverted struct {
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReverted is a free log retrieval operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMUpgradeTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestRevertedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestReverted)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Reverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReverted is a log parse operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseReverted(log types.Log) (*GatewayEVMUpgradeTestReverted, error) {
	event := new(GatewayEVMUpgradeTestReverted)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestRoleAdminChangedIterator struct {
	Event *GatewayEVMUpgradeTestRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestRoleAdminChanged represents a RoleAdminChanged event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GatewayEVMUpgradeTestRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestRoleAdminChangedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestRoleAdminChanged)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseRoleAdminChanged(log types.Log) (*GatewayEVMUpgradeTestRoleAdminChanged, error) {
	event := new(GatewayEVMUpgradeTestRoleAdminChanged)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestRoleGrantedIterator struct {
	Event *GatewayEVMUpgradeTestRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestRoleGranted represents a RoleGranted event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayEVMUpgradeTestRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestRoleGrantedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestRoleGranted)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseRoleGranted(log types.Log) (*GatewayEVMUpgradeTestRoleGranted, error) {
	event := new(GatewayEVMUpgradeTestRoleGranted)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestRoleRevokedIterator struct {
	Event *GatewayEVMUpgradeTestRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestRoleRevoked represents a RoleRevoked event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayEVMUpgradeTestRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestRoleRevokedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestRoleRevoked)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseRoleRevoked(log types.Log) (*GatewayEVMUpgradeTestRoleRevoked, error) {
	event := new(GatewayEVMUpgradeTestRoleRevoked)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestUnpausedIterator struct {
	Event *GatewayEVMUpgradeTestUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestUnpaused represents a Unpaused event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GatewayEVMUpgradeTestUnpausedIterator, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestUnpausedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestUnpaused) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestUnpaused)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseUnpaused(log types.Log) (*GatewayEVMUpgradeTestUnpaused, error) {
	event := new(GatewayEVMUpgradeTestUnpaused)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestUpdatedAdditionalActionFeeIterator is returned from FilterUpdatedAdditionalActionFee and is used to iterate over the raw logs and unpacked data for UpdatedAdditionalActionFee events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestUpdatedAdditionalActionFeeIterator struct {
	Event *GatewayEVMUpgradeTestUpdatedAdditionalActionFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestUpdatedAdditionalActionFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestUpdatedAdditionalActionFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestUpdatedAdditionalActionFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestUpdatedAdditionalActionFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestUpdatedAdditionalActionFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestUpdatedAdditionalActionFee represents a UpdatedAdditionalActionFee event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestUpdatedAdditionalActionFee struct {
	OldFeeWei *big.Int
	NewFeeWei *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAdditionalActionFee is a free log retrieval operation binding the contract event 0xe60a2882aff7841a7d0492b804d46ff2d9b07ea353dbc3791e8276b729e40cd8.
//
// Solidity: event UpdatedAdditionalActionFee(uint256 oldFeeWei, uint256 newFeeWei)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterUpdatedAdditionalActionFee(opts *bind.FilterOpts) (*GatewayEVMUpgradeTestUpdatedAdditionalActionFeeIterator, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "UpdatedAdditionalActionFee")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestUpdatedAdditionalActionFeeIterator{contract: _GatewayEVMUpgradeTest.contract, event: "UpdatedAdditionalActionFee", logs: logs, sub: sub}, nil
}

// WatchUpdatedAdditionalActionFee is a free log subscription operation binding the contract event 0xe60a2882aff7841a7d0492b804d46ff2d9b07ea353dbc3791e8276b729e40cd8.
//
// Solidity: event UpdatedAdditionalActionFee(uint256 oldFeeWei, uint256 newFeeWei)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchUpdatedAdditionalActionFee(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestUpdatedAdditionalActionFee) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "UpdatedAdditionalActionFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestUpdatedAdditionalActionFee)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "UpdatedAdditionalActionFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedAdditionalActionFee is a log parse operation binding the contract event 0xe60a2882aff7841a7d0492b804d46ff2d9b07ea353dbc3791e8276b729e40cd8.
//
// Solidity: event UpdatedAdditionalActionFee(uint256 oldFeeWei, uint256 newFeeWei)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseUpdatedAdditionalActionFee(log types.Log) (*GatewayEVMUpgradeTestUpdatedAdditionalActionFee, error) {
	event := new(GatewayEVMUpgradeTestUpdatedAdditionalActionFee)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "UpdatedAdditionalActionFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestUpdatedGatewayTSSAddressIterator struct {
	Event *GatewayEVMUpgradeTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestUpdatedGatewayTSSAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestUpdatedGatewayTSSAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestUpdatedGatewayTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*GatewayEVMUpgradeTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestUpdatedGatewayTSSAddressIterator{contract: _GatewayEVMUpgradeTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestUpdatedGatewayTSSAddress)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedGatewayTSSAddress is a log parse operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*GatewayEVMUpgradeTestUpdatedGatewayTSSAddress, error) {
	event := new(GatewayEVMUpgradeTestUpdatedGatewayTSSAddress)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMUpgradeTestUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestUpgradedIterator struct {
	Event *GatewayEVMUpgradeTestUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayEVMUpgradeTestUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMUpgradeTestUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayEVMUpgradeTestUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayEVMUpgradeTestUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMUpgradeTestUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMUpgradeTestUpgraded represents a Upgraded event raised by the GatewayEVMUpgradeTest contract.
type GatewayEVMUpgradeTestUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayEVMUpgradeTestUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMUpgradeTestUpgradedIterator{contract: _GatewayEVMUpgradeTest.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayEVMUpgradeTestUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayEVMUpgradeTest.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMUpgradeTestUpgraded)
				if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayEVMUpgradeTest *GatewayEVMUpgradeTestFilterer) ParseUpgraded(log types.Log) (*GatewayEVMUpgradeTestUpgraded, error) {
	event := new(GatewayEVMUpgradeTestUpgraded)
	if err := _GatewayEVMUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
