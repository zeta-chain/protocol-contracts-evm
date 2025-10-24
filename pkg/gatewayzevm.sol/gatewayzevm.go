// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayzevm

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

// AbortContext is an auto generated low-level Go binding around an user-defined struct.
type AbortContext struct {
	Sender        []byte
	Asset         common.Address
	Amount        *big.Int
	Outgoing      bool
	ChainID       *big.Int
	RevertMessage []byte
}

// CallOptions is an auto generated low-level Go binding around an user-defined struct.
type CallOptions struct {
	GasLimit        *big.Int
	IsArbitraryCall bool
}

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Sender    []byte
	SenderEVM common.Address
	ChainID   *big.Int
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

// GatewayZEVMMetaData contains all meta data concerning the GatewayZEVM contract.
var GatewayZEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_MESSAGE_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROTOCOL_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndRevert\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeAbort\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"abortContext\",\"type\":\"tuple\",\"internalType\":\"structAbortContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outgoing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallOnRevertNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZETANotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a060405230608052348015610013575f80fd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051613e6d6100f95f395f8181612487015281816124b001526126850152613e6d5ff3fe6080604052600436106101de575f3560e01c80637b15118b116100fd578063a217fddf11610092578063c5356f6d11610062578063c5356f6d1461066a578063d547741f14610689578063e63ab1e9146106a8578063f45346dc146106db575f80fd5b8063a217fddf146105c4578063ad3cb1cc146105d7578063bcf7f32b1461062c578063c39aca371461064b575f80fd5b806391d14854116100cd57806391d148541461051357806397a1cef11461057657806397d340f5146105905780639d4ba465146105a5575f80fd5b80637b15118b146104ab5780637c0dcb5f146104ca5780637ce1ffeb146104e95780638456cb59146104ff575f80fd5b80632810ae6311610173578063485cc95511610143578063485cc9551461042f5780634f1ef2861461044e57806352d1902d146104615780635c975abb14610475575f80fd5b80632810ae63146103be5780632f2ff15d146103dd57806336568abe146103fc5780633f4ba83a1461041b575f80fd5b806321501a95116101ae57806321501a95146102e757806321e093b114610306578063248a9ca31461033c5780632722feee14610397575f80fd5b806301ffc9a71461025657806306cb89831461028a578063184b0793146102a95780632095dedb146102c8575f80fd5b36610252576101eb6106fa565b5f546001600160a01b0316331480159061021957503373735b14bb79463307aacbed86daf3322b1e6226ab14155b15610250576040517fb3af013700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b5f80fd5b348015610261575f80fd5b50610275610270366004612f4d565b610758565b60405190151581526020015b60405180910390f35b348015610295575f80fd5b506102506102a43660046130e1565b6107f0565b3480156102b4575f80fd5b506102506102c33660046131a8565b6108b5565b3480156102d3575f80fd5b506102506102e23660046131f5565b6109f1565b3480156102f2575f80fd5b50610250610301366004613258565b610ad3565b348015610311575f80fd5b505f54610324906001600160a01b031681565b6040516001600160a01b039091168152602001610281565b348015610347575f80fd5b506103896103563660046132de565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b604051908152602001610281565b3480156103a2575f80fd5b5061032473735b14bb79463307aacbed86daf3322b1e6226ab81565b3480156103c9575f80fd5b506102506103d83660046132f5565b610cca565b3480156103e8575f80fd5b506102506103f73660046133ad565b610d04565b348015610407575f80fd5b506102506104163660046133ad565b610d4d565b348015610426575f80fd5b50610250610d9e565b34801561043a575f80fd5b506102506104493660046133d0565b610dd3565b61025061045c3660046133fc565b611025565b34801561046c575f80fd5b50610389611040565b348015610480575f80fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610275565b3480156104b6575f80fd5b506102506104c536600461343f565b61106e565b3480156104d5575f80fd5b506102506104e43660046134aa565b6112af565b3480156104f4575f80fd5b50610389620186a081565b34801561050a575f80fd5b506102506114a9565b34801561051e575f80fd5b5061027561052d3660046133ad565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b348015610581575f80fd5b506102506103d836600461352a565b34801561059b575f80fd5b50610389610b4081565b3480156105b0575f80fd5b506102506105bf366004613589565b6114db565b3480156105cf575f80fd5b506103895f81565b3480156105e2575f80fd5b5061061f6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516102819190613631565b348015610637575f80fd5b50610250610646366004613643565b611789565b348015610656575f80fd5b50610250610665366004613643565b6118dd565b348015610675575f80fd5b506102506106843660046136db565b611af0565b348015610694575f80fd5b506102506106a33660046133ad565b611cc6565b3480156106b3575f80fd5b506103897f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b3480156106e6575f80fd5b506102506106f5366004613764565b611d09565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615610756576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806107ea57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6107f86106fa565b620186a082351015610836576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b4061084660608301836137a3565b610851915085613804565b1115610889576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108ad8686868661089f36889003880188613849565b6108a88761389f565b611f14565b505050505050565b6108bd611fa7565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461090a576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109126106fa565b6001600160a01b038216610952576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063c9028a36906109979084906004016139eb565b5f604051808303815f87803b1580156109ae575f80fd5b505af11580156109c0573d5f803e3d5ffd5b505050506109ed60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5050565b6109f9611fa7565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a46576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a4e6106fa565b6001600160a01b038216610a8e576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f2d4cfb7e0000000000000000000000000000000000000000000000000000000081526001600160a01b03831690632d4cfb7e90610997908490600401613a59565b610adb611fa7565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610b28576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b306106fa565b6001600160a01b038316610b70576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b835f03610ba9576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831673735b14bb79463307aacbed86daf3322b1e6226ab1480610bdc57506001600160a01b03831630145b15610c13576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c1d848461204e565b5f546040517f5bcfd6160000000000000000000000000000000000000000000000000000000081526001600160a01b0380861692635bcfd61692610c6d928a921690899088908890600401613b1c565b5f604051808303815f87803b158015610c84575f80fd5b505af1158015610c96573d5f803e3d5ffd5b50505050610cc360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5050505050565b610cd26106fa565b6040517fe4dd681d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610d3d8161220e565b610d478383612218565b50505050565b6001600160a01b0381163314610d8f576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d998282612302565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610dc88161220e565b610dd06123c4565b50565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f81158015610e1d5750825b90505f8267ffffffffffffffff166001148015610e395750303b155b905081158015610e47575080155b15610e7e576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610edf5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b0387161580610efc57506001600160a01b038616155b15610f33576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f3b612454565b610f43612454565b610f4b61245c565b610f5361246c565b610f5d5f87612218565b50610f887f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a87612218565b505f80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038916179055831561101c5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b61102d61247c565b6110368261254c565b6109ed8282612556565b5f61104961267a565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b6110766106fa565b86515f036110b0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855f036110e9576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620186a082351015611127576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b4061113760608301836137a3565b611142915085613804565b111561117a576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b846001600160a01b03166385e1f4d06040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111b6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111da9190613b98565b6069036111e5575f80fd5b5f6111f2878785356126dc565b90505f336001600160a01b03167fd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a8a898b868c6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611260573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112849190613b98565b8c8c8c8c60405161129d99989796959493929190613c2f565b60405180910390a35050505050505050565b6112b76106fa565b83515f036112f1576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b825f0361132a576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b4061133a60608301836137a3565b90501115611374576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61137f848461286f565b90505f336001600160a01b03167f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c87868886896001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113ed573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114119190613b98565b60405180604001604052808c6001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa158015611458573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061147c9190613b98565b8152600160209091015260405161149a969594939291908c90613cb8565b60405180910390a35050505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6114d38161220e565b610dd06128da565b6114e3611fa7565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611530576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115386106fa565b6001600160a01b038416158061155557506001600160a01b038216155b1561158c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b825f036115c5576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03821673735b14bb79463307aacbed86daf3322b1e6226ab14806115f857506001600160a01b03821630145b1561162f576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152602482018590528516906347e7ef24906044016020604051808303815f875af1158015611694573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116b89190613d38565b6116ee576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063c9028a36906117339084906004016139eb565b5f604051808303815f87803b15801561174a575f80fd5b505af115801561175c573d5f803e3d5ffd5b50505050610d4760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b611791611fa7565b3373735b14bb79463307aacbed86daf3322b1e6226ab146117de576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6117e66106fa565b6001600160a01b038516158061180357506001600160a01b038316155b1561183a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f5bcfd6160000000000000000000000000000000000000000000000000000000081526001600160a01b03841690635bcfd616906118879089908990899088908890600401613b1c565b5f604051808303815f87803b15801561189e575f80fd5b505af11580156118b0573d5f803e3d5ffd5b505050506108ad60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6118e5611fa7565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611932576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61193a6106fa565b6001600160a01b038516158061195757506001600160a01b038316155b1561198e576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b835f036119c7576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831673735b14bb79463307aacbed86daf3322b1e6226ab14806119fa57506001600160a01b03831630145b15611a31576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038481166004830152602482018690528616906347e7ef24906044016020604051808303815f875af1158015611a96573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611aba9190613d38565b61183a576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611af86106fa565b84515f03611b32576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b835f03611b6b576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620186a0821015611ba8576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b40611bb860608301836137a3565b90501115611bf2576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611bfe8585856126dc565b90505f336001600160a01b03167f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c888789868a6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c6c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c909190613b98565b6040805180820182528c8152600160208201529051611cb6969594939291908c90613cb8565b60405180910390a3505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611cff8161220e565b610d478383612302565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611d56576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611d5e6106fa565b6001600160a01b0383161580611d7b57506001600160a01b038116155b15611db2576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f03611deb576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03811673735b14bb79463307aacbed86daf3322b1e6226ab1480611e1e57506001600160a01b03811630145b15611e55576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038281166004830152602482018490528416906347e7ef24906044016020604051808303815f875af1158015611eba573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ede9190613d38565b610d99576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b85515f03611f4e576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611f5b85835f0151612953565b50846001600160a01b0316336001600160a01b03167f306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e48887878787604051611cb6959493929190613d53565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01612022576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490526001600160a01b03909116906323b872dd906064016020604051808303815f875af11580156120ba573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120de9190613d38565b612114576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f546040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b0390911690632e1a7d4d906024015f604051808303815f87803b15801561216f575f80fd5b505af1158015612181573d5f803e3d5ffd5b505050505f816001600160a01b0316836040515f6040518083038185875af1925050503d805f81146121ce576040519150601f19603f3d011682016040523d82523d5f602084013e6121d3565b606091505b5050905080610d99576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610dd08133612b49565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff166122f9575f848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556122af3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506107ea565b5f9150506107ea565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16156122f9575f848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506107ea565b6123cc612bd5565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b610756612c30565b612464612c30565b610756612c97565b612474612c30565b610756612ce8565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061251557507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166125097f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15610756576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6109ed8161220e565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156125ce575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526125cb91810190613b98565b60015b612614576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612670576040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004810182905260240161260b565b610d998383612cf0565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610756576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f806126e88484612953565b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018790529091506001600160a01b038516906323b872dd906064016020604051808303815f875af1158015612754573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127789190613d38565b6127ae576040517f4dd9ee8d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f42966c68000000000000000000000000000000000000000000000000000000008152600481018690526001600160a01b038516906342966c68906024016020604051808303815f875af115801561280b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061282f9190613d38565b612865576040517f2c77e05c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b90505b9392505050565b5f6128688383846001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128b1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128d59190613b98565b6126dc565b6128e26106fa565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612436565b5f805f846001600160a01b031663fc5fecd5856040518263ffffffff1660e01b815260040161298491815260200190565b6040805180830381865afa15801561299e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129c29190613df5565b6040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810182905291935091506001600160a01b038316906323b872dd906064016020604051808303815f875af1158015612a30573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a549190613d38565b612a8a576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f42966c68000000000000000000000000000000000000000000000000000000008152600481018290526001600160a01b038316906342966c68906024016020604051808303815f875af1158015612ae7573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b0b9190613d38565b612b41576040517f2c77e05c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b949350505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff166109ed576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024810183905260440161260b565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610756576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610756576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612c9f612c30565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b612028612c30565b612cf982612d45565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612d3d57610d998282612dec565b6109ed612e5e565b806001600160a01b03163b5f03612d93576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161260b565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f80846001600160a01b031684604051612e089190613e21565b5f60405180830381855af49150503d805f8114612e40576040519150601f19603f3d011682016040523d82523d5f602084013e612e45565b606091505b5091509150612e55858383612e96565b95945050505050565b3415610756576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606082612eab57612ea682612f0b565b612868565b8151158015612ec257506001600160a01b0384163b155b15612f04576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161260b565b5080612868565b805115612f1b5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60208284031215612f5d575f80fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114612868575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112612fc8575f80fd5b813567ffffffffffffffff811115612fe257612fe2612f8c565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810167ffffffffffffffff8111828210171561302f5761302f612f8c565b604052818152838201602001851015613046575f80fd5b816020850160208301375f918101602001919091529392505050565b6001600160a01b0381168114610dd0575f80fd5b5f8083601f840112613086575f80fd5b50813567ffffffffffffffff81111561309d575f80fd5b6020830191508360208285010111156130b4575f80fd5b9250929050565b5f604082840312156130cb575f80fd5b50919050565b5f60a082840312156130cb575f80fd5b5f805f805f8060c087890312156130f6575f80fd5b863567ffffffffffffffff81111561310c575f80fd5b61311889828a01612fb9565b965050602087013561312981613062565b9450604087013567ffffffffffffffff811115613144575f80fd5b61315089828a01613076565b9095509350613164905088606089016130bb565b915060a087013567ffffffffffffffff81111561317f575f80fd5b61318b89828a016130d1565b9150509295509295509295565b5f608082840312156130cb575f80fd5b5f80604083850312156131b9575f80fd5b82356131c481613062565b9150602083013567ffffffffffffffff8111156131df575f80fd5b6131eb85828601613198565b9150509250929050565b5f8060408385031215613206575f80fd5b823561321181613062565b9150602083013567ffffffffffffffff81111561322c575f80fd5b830160c0818603121561323d575f80fd5b809150509250929050565b5f606082840312156130cb575f80fd5b5f805f805f6080868803121561326c575f80fd5b853567ffffffffffffffff811115613282575f80fd5b61328e88828901613248565b9550506020860135935060408601356132a681613062565b9250606086013567ffffffffffffffff8111156132c1575f80fd5b6132cd88828901613076565b969995985093965092949392505050565b5f602082840312156132ee575f80fd5b5035919050565b5f805f805f805f60e0888a03121561330b575f80fd5b873567ffffffffffffffff811115613321575f80fd5b61332d8a828b01612fb9565b9750506020880135955060408801359450606088013567ffffffffffffffff811115613357575f80fd5b6133638a828b01613076565b909550935061337790508960808a016130bb565b915060c088013567ffffffffffffffff811115613392575f80fd5b61339e8a828b016130d1565b91505092959891949750929550565b5f80604083850312156133be575f80fd5b82359150602083013561323d81613062565b5f80604083850312156133e1575f80fd5b82356133ec81613062565b9150602083013561323d81613062565b5f806040838503121561340d575f80fd5b823561341881613062565b9150602083013567ffffffffffffffff811115613433575f80fd5b6131eb85828601612fb9565b5f805f805f805f60e0888a031215613455575f80fd5b873567ffffffffffffffff81111561346b575f80fd5b6134778a828b01612fb9565b97505060208801359550604088013561348f81613062565b9450606088013567ffffffffffffffff811115613357575f80fd5b5f805f80608085870312156134bd575f80fd5b843567ffffffffffffffff8111156134d3575f80fd5b6134df87828801612fb9565b9450506020850135925060408501356134f781613062565b9150606085013567ffffffffffffffff811115613512575f80fd5b61351e878288016130d1565b91505092959194509250565b5f805f806080858703121561353d575f80fd5b843567ffffffffffffffff811115613553575f80fd5b61355f87828801612fb9565b9450506020850135925060408501359150606085013567ffffffffffffffff811115613512575f80fd5b5f805f806080858703121561359c575f80fd5b84356135a781613062565b93506020850135925060408501356135be81613062565b9150606085013567ffffffffffffffff8111156135d9575f80fd5b61351e87828801613198565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61286860208301846135e5565b5f805f805f8060a08789031215613658575f80fd5b863567ffffffffffffffff81111561366e575f80fd5b61367a89828a01613248565b965050602087013561368b81613062565b94506040870135935060608701356136a281613062565b9250608087013567ffffffffffffffff8111156136bd575f80fd5b6136c989828a01613076565b979a9699509497509295939492505050565b5f805f805f60a086880312156136ef575f80fd5b853567ffffffffffffffff811115613705575f80fd5b61371188828901612fb9565b95505060208601359350604086013561372981613062565b925060608601359150608086013567ffffffffffffffff81111561374b575f80fd5b613757888289016130d1565b9150509295509295909350565b5f805f60608486031215613776575f80fd5b833561378181613062565b925060208401359150604084013561379881613062565b809150509250925092565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126137d6575f80fd5b83018035915067ffffffffffffffff8211156137f0575f80fd5b6020019150368190038213156130b4575f80fd5b808201808211156107ea577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8015158114610dd0575f80fd5b5f604082840312801561385a575f80fd5b506040805190810167ffffffffffffffff8111828210171561387e5761387e612f8c565b6040528235815260208301356138938161383c565b60208201529392505050565b5f60a082360312156138af575f80fd5b60405160a0810167ffffffffffffffff811182821017156138d2576138d2612f8c565b60405282356138e081613062565b815260208301356138f08161383c565b6020820152604083013561390381613062565b6040820152606083013567ffffffffffffffff811115613921575f80fd5b61392d36828601612fb9565b606083015250608092830135928101929092525090565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613977575f80fd5b830160208101925035905067ffffffffffffffff811115613996575f80fd5b8036038213156130b4575f80fd5b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b602081525f82356139fb81613062565b6001600160a01b0381166020840152506020830135613a1981613062565b6001600160a01b0381166040840152505f6040840135905080606084015250613a456060840184613944565b608080850152612e5560a0850182846139a4565b602081525f613a688384613944565b60c06020850152613a7d60e0850182846139a4565b9150506020840135613a8e81613062565b6001600160a01b0316604084810191909152840135606080850191909152840135613ab88161383c565b8015156080850152505f608085013590508060a085015250613add60a0850185613944565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c0860152613b128382846139a4565b9695505050505050565b608081525f613b2b8788613944565b60606080850152613b4060e0850182846139a4565b9150506020880135613b5181613062565b6001600160a01b0390811660a085015260408981013560c0860152908816602085015283018690528281036060840152613b8c8185876139a4565b98975050505050505050565b5f60208284031215613ba8575f80fd5b5051919050565b5f8135613bbb81613062565b6001600160a01b031683526020820135613bd48161383c565b151560208401526040820135613be981613062565b6001600160a01b03166040840152613c046060830183613944565b60a06060860152613c1960a0860182846139a4565b6080948501359590940194909452509092915050565b61012081525f613c4361012083018c6135e5565b6001600160a01b038b16602084015289604084015288606084015287608084015282810360a0840152613c778187896139a4565b853560c085015290506020850135613c8e8161383c565b151560e0840152828103610100840152613ca88185613baf565b9c9b505050505050505050505050565b61012081525f613ccc61012083018a6135e5565b6001600160a01b03891660208401528760408401528660608401528560808401528281038060a08501525f8252613d1260c0850187805182526020908101511515910152565b6020810161010085015250613d2a6020820185613baf565b9a9950505050505050505050565b5f60208284031215613d48575f80fd5b81516128688161383c565b60a081525f613d6560a08301886135e5565b8281036020840152613d788187896139a4565b85516040850152602086015115156060850152905082810360808401526001600160a01b0384511681526020840151151560208201526001600160a01b036040850151166040820152606084015160a06060830152613dda60a08301826135e5565b90506080850151608083015280925050509695505050505050565b5f8060408385031215613e06575f80fd5b8251613e1181613062565b6020939093015192949293505050565b5f82518060208501845e5f92019182525091905056fea2646970667358221220b5e8be7f4c42c777a791a61cc01378a6d40c25ddd6e480572b9f97ee3a93465a64736f6c634300081a0033",
}

// GatewayZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayZEVMMetaData.ABI instead.
var GatewayZEVMABI = GatewayZEVMMetaData.ABI

// GatewayZEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayZEVMMetaData.Bin instead.
var GatewayZEVMBin = GatewayZEVMMetaData.Bin

// DeployGatewayZEVM deploys a new Ethereum contract, binding an instance of GatewayZEVM to it.
func DeployGatewayZEVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayZEVM, error) {
	parsed, err := GatewayZEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayZEVMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayZEVM{GatewayZEVMCaller: GatewayZEVMCaller{contract: contract}, GatewayZEVMTransactor: GatewayZEVMTransactor{contract: contract}, GatewayZEVMFilterer: GatewayZEVMFilterer{contract: contract}}, nil
}

// GatewayZEVM is an auto generated Go binding around an Ethereum contract.
type GatewayZEVM struct {
	GatewayZEVMCaller     // Read-only binding to the contract
	GatewayZEVMTransactor // Write-only binding to the contract
	GatewayZEVMFilterer   // Log filterer for contract events
}

// GatewayZEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayZEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayZEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayZEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayZEVMSession struct {
	Contract     *GatewayZEVM      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayZEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayZEVMCallerSession struct {
	Contract *GatewayZEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GatewayZEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayZEVMTransactorSession struct {
	Contract     *GatewayZEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GatewayZEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayZEVMRaw struct {
	Contract *GatewayZEVM // Generic contract binding to access the raw methods on
}

// GatewayZEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayZEVMCallerRaw struct {
	Contract *GatewayZEVMCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayZEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayZEVMTransactorRaw struct {
	Contract *GatewayZEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayZEVM creates a new instance of GatewayZEVM, bound to a specific deployed contract.
func NewGatewayZEVM(address common.Address, backend bind.ContractBackend) (*GatewayZEVM, error) {
	contract, err := bindGatewayZEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVM{GatewayZEVMCaller: GatewayZEVMCaller{contract: contract}, GatewayZEVMTransactor: GatewayZEVMTransactor{contract: contract}, GatewayZEVMFilterer: GatewayZEVMFilterer{contract: contract}}, nil
}

// NewGatewayZEVMCaller creates a new read-only instance of GatewayZEVM, bound to a specific deployed contract.
func NewGatewayZEVMCaller(address common.Address, caller bind.ContractCaller) (*GatewayZEVMCaller, error) {
	contract, err := bindGatewayZEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMCaller{contract: contract}, nil
}

// NewGatewayZEVMTransactor creates a new write-only instance of GatewayZEVM, bound to a specific deployed contract.
func NewGatewayZEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayZEVMTransactor, error) {
	contract, err := bindGatewayZEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMTransactor{contract: contract}, nil
}

// NewGatewayZEVMFilterer creates a new log filterer instance of GatewayZEVM, bound to a specific deployed contract.
func NewGatewayZEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayZEVMFilterer, error) {
	contract, err := bindGatewayZEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMFilterer{contract: contract}, nil
}

// bindGatewayZEVM binds a generic wrapper to an already deployed contract.
func bindGatewayZEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayZEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVM *GatewayZEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVM.Contract.GatewayZEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVM *GatewayZEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.GatewayZEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVM *GatewayZEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.GatewayZEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVM *GatewayZEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVM *GatewayZEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVM *GatewayZEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayZEVM.Contract.DEFAULTADMINROLE(&_GatewayZEVM.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayZEVM.Contract.DEFAULTADMINROLE(&_GatewayZEVM.CallOpts)
}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVM *GatewayZEVMCaller) MAXMESSAGESIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "MAX_MESSAGE_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVM *GatewayZEVMSession) MAXMESSAGESIZE() (*big.Int, error) {
	return _GatewayZEVM.Contract.MAXMESSAGESIZE(&_GatewayZEVM.CallOpts)
}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVM *GatewayZEVMCallerSession) MAXMESSAGESIZE() (*big.Int, error) {
	return _GatewayZEVM.Contract.MAXMESSAGESIZE(&_GatewayZEVM.CallOpts)
}

// MINGASLIMIT is a free data retrieval call binding the contract method 0x7ce1ffeb.
//
// Solidity: function MIN_GAS_LIMIT() view returns(uint256)
func (_GatewayZEVM *GatewayZEVMCaller) MINGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "MIN_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINGASLIMIT is a free data retrieval call binding the contract method 0x7ce1ffeb.
//
// Solidity: function MIN_GAS_LIMIT() view returns(uint256)
func (_GatewayZEVM *GatewayZEVMSession) MINGASLIMIT() (*big.Int, error) {
	return _GatewayZEVM.Contract.MINGASLIMIT(&_GatewayZEVM.CallOpts)
}

// MINGASLIMIT is a free data retrieval call binding the contract method 0x7ce1ffeb.
//
// Solidity: function MIN_GAS_LIMIT() view returns(uint256)
func (_GatewayZEVM *GatewayZEVMCallerSession) MINGASLIMIT() (*big.Int, error) {
	return _GatewayZEVM.Contract.MINGASLIMIT(&_GatewayZEVM.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayZEVM.Contract.PAUSERROLE(&_GatewayZEVM.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayZEVM.Contract.PAUSERROLE(&_GatewayZEVM.CallOpts)
}

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_GatewayZEVM *GatewayZEVMCaller) PROTOCOLADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "PROTOCOL_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_GatewayZEVM *GatewayZEVMSession) PROTOCOLADDRESS() (common.Address, error) {
	return _GatewayZEVM.Contract.PROTOCOLADDRESS(&_GatewayZEVM.CallOpts)
}

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_GatewayZEVM *GatewayZEVMCallerSession) PROTOCOLADDRESS() (common.Address, error) {
	return _GatewayZEVM.Contract.PROTOCOLADDRESS(&_GatewayZEVM.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayZEVM *GatewayZEVMCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayZEVM *GatewayZEVMSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayZEVM.Contract.UPGRADEINTERFACEVERSION(&_GatewayZEVM.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayZEVM *GatewayZEVMCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayZEVM.Contract.UPGRADEINTERFACEVERSION(&_GatewayZEVM.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayZEVM.Contract.GetRoleAdmin(&_GatewayZEVM.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayZEVM.Contract.GetRoleAdmin(&_GatewayZEVM.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayZEVM *GatewayZEVMCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayZEVM *GatewayZEVMSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayZEVM.Contract.HasRole(&_GatewayZEVM.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayZEVM *GatewayZEVMCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayZEVM.Contract.HasRole(&_GatewayZEVM.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayZEVM *GatewayZEVMCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayZEVM *GatewayZEVMSession) Paused() (bool, error) {
	return _GatewayZEVM.Contract.Paused(&_GatewayZEVM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayZEVM *GatewayZEVMCallerSession) Paused() (bool, error) {
	return _GatewayZEVM.Contract.Paused(&_GatewayZEVM.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayZEVM.Contract.ProxiableUUID(&_GatewayZEVM.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVM *GatewayZEVMCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayZEVM.Contract.ProxiableUUID(&_GatewayZEVM.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayZEVM *GatewayZEVMCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayZEVM *GatewayZEVMSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayZEVM.Contract.SupportsInterface(&_GatewayZEVM.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayZEVM *GatewayZEVMCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayZEVM.Contract.SupportsInterface(&_GatewayZEVM.CallOpts, interfaceId)
}

// Withdraw0 is a free data retrieval call binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes , uint256 , uint256 , (address,bool,address,bytes,uint256) ) view returns()
func (_GatewayZEVM *GatewayZEVMCaller) Withdraw0(opts *bind.CallOpts, arg0 []byte, arg1 *big.Int, arg2 *big.Int, arg3 RevertOptions) error {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "withdraw0", arg0, arg1, arg2, arg3)

	if err != nil {
		return err
	}

	return err

}

// Withdraw0 is a free data retrieval call binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes , uint256 , uint256 , (address,bool,address,bytes,uint256) ) view returns()
func (_GatewayZEVM *GatewayZEVMSession) Withdraw0(arg0 []byte, arg1 *big.Int, arg2 *big.Int, arg3 RevertOptions) error {
	return _GatewayZEVM.Contract.Withdraw0(&_GatewayZEVM.CallOpts, arg0, arg1, arg2, arg3)
}

// Withdraw0 is a free data retrieval call binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes , uint256 , uint256 , (address,bool,address,bytes,uint256) ) view returns()
func (_GatewayZEVM *GatewayZEVMCallerSession) Withdraw0(arg0 []byte, arg1 *big.Int, arg2 *big.Int, arg3 RevertOptions) error {
	return _GatewayZEVM.Contract.Withdraw0(&_GatewayZEVM.CallOpts, arg0, arg1, arg2, arg3)
}

// WithdrawAndCall is a free data retrieval call binding the contract method 0x2810ae63.
//
// Solidity: function withdrawAndCall(bytes , uint256 , uint256 , bytes , (uint256,bool) , (address,bool,address,bytes,uint256) ) view returns()
func (_GatewayZEVM *GatewayZEVMCaller) WithdrawAndCall(opts *bind.CallOpts, arg0 []byte, arg1 *big.Int, arg2 *big.Int, arg3 []byte, arg4 CallOptions, arg5 RevertOptions) error {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "withdrawAndCall", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// WithdrawAndCall is a free data retrieval call binding the contract method 0x2810ae63.
//
// Solidity: function withdrawAndCall(bytes , uint256 , uint256 , bytes , (uint256,bool) , (address,bool,address,bytes,uint256) ) view returns()
func (_GatewayZEVM *GatewayZEVMSession) WithdrawAndCall(arg0 []byte, arg1 *big.Int, arg2 *big.Int, arg3 []byte, arg4 CallOptions, arg5 RevertOptions) error {
	return _GatewayZEVM.Contract.WithdrawAndCall(&_GatewayZEVM.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// WithdrawAndCall is a free data retrieval call binding the contract method 0x2810ae63.
//
// Solidity: function withdrawAndCall(bytes , uint256 , uint256 , bytes , (uint256,bool) , (address,bool,address,bytes,uint256) ) view returns()
func (_GatewayZEVM *GatewayZEVMCallerSession) WithdrawAndCall(arg0 []byte, arg1 *big.Int, arg2 *big.Int, arg3 []byte, arg4 CallOptions, arg5 RevertOptions) error {
	return _GatewayZEVM.Contract.WithdrawAndCall(&_GatewayZEVM.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayZEVM *GatewayZEVMCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayZEVM.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayZEVM *GatewayZEVMSession) ZetaToken() (common.Address, error) {
	return _GatewayZEVM.Contract.ZetaToken(&_GatewayZEVM.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayZEVM *GatewayZEVMCallerSession) ZetaToken() (common.Address, error) {
	return _GatewayZEVM.Contract.ZetaToken(&_GatewayZEVM.CallOpts)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Call(opts *bind.TransactOpts, receiver []byte, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "call", receiver, zrc20, message, callOptions, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMSession) Call(receiver []byte, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Call(&_GatewayZEVM.TransactOpts, receiver, zrc20, message, callOptions, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Call(receiver []byte, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Call(&_GatewayZEVM.TransactOpts, receiver, zrc20, message, callOptions, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Deposit(opts *bind.TransactOpts, zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "deposit", zrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVM *GatewayZEVMSession) Deposit(zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Deposit(&_GatewayZEVM.TransactOpts, zrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Deposit(zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Deposit(&_GatewayZEVM.TransactOpts, zrc20, amount, target)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) DepositAndCall(opts *bind.TransactOpts, context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "depositAndCall", context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMSession) DepositAndCall(context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndCall(&_GatewayZEVM.TransactOpts, context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) DepositAndCall(context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndCall(&_GatewayZEVM.TransactOpts, context, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) DepositAndCall0(opts *bind.TransactOpts, context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "depositAndCall0", context, zrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMSession) DepositAndCall0(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndCall0(&_GatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) DepositAndCall0(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndCall0(&_GatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) DepositAndRevert(opts *bind.TransactOpts, zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "depositAndRevert", zrc20, amount, target, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMSession) DepositAndRevert(zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndRevert(&_GatewayZEVM.TransactOpts, zrc20, amount, target, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) DepositAndRevert(zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.DepositAndRevert(&_GatewayZEVM.TransactOpts, zrc20, amount, target, revertContext)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Execute(opts *bind.TransactOpts, context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "execute", context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMSession) Execute(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Execute(&_GatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Execute(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Execute(&_GatewayZEVM.TransactOpts, context, zrc20, amount, target, message)
}

// ExecuteAbort is a paid mutator transaction binding the contract method 0x2095dedb.
//
// Solidity: function executeAbort(address target, (bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) ExecuteAbort(opts *bind.TransactOpts, target common.Address, abortContext AbortContext) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "executeAbort", target, abortContext)
}

// ExecuteAbort is a paid mutator transaction binding the contract method 0x2095dedb.
//
// Solidity: function executeAbort(address target, (bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_GatewayZEVM *GatewayZEVMSession) ExecuteAbort(target common.Address, abortContext AbortContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.ExecuteAbort(&_GatewayZEVM.TransactOpts, target, abortContext)
}

// ExecuteAbort is a paid mutator transaction binding the contract method 0x2095dedb.
//
// Solidity: function executeAbort(address target, (bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) ExecuteAbort(target common.Address, abortContext AbortContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.ExecuteAbort(&_GatewayZEVM.TransactOpts, target, abortContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) ExecuteRevert(opts *bind.TransactOpts, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "executeRevert", target, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMSession) ExecuteRevert(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.ExecuteRevert(&_GatewayZEVM.TransactOpts, target, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) ExecuteRevert(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.ExecuteRevert(&_GatewayZEVM.TransactOpts, target, revertContext)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayZEVM *GatewayZEVMSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.GrantRole(&_GatewayZEVM.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.GrantRole(&_GatewayZEVM.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address zetaToken_, address admin_) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Initialize(opts *bind.TransactOpts, zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "initialize", zetaToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address zetaToken_, address admin_) returns()
func (_GatewayZEVM *GatewayZEVMSession) Initialize(zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Initialize(&_GatewayZEVM.TransactOpts, zetaToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address zetaToken_, address admin_) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Initialize(zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Initialize(&_GatewayZEVM.TransactOpts, zetaToken_, admin_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayZEVM *GatewayZEVMSession) Pause() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Pause(&_GatewayZEVM.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Pause() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Pause(&_GatewayZEVM.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayZEVM *GatewayZEVMSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.RenounceRole(&_GatewayZEVM.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.RenounceRole(&_GatewayZEVM.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayZEVM *GatewayZEVMSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.RevokeRole(&_GatewayZEVM.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.RevokeRole(&_GatewayZEVM.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayZEVM *GatewayZEVMSession) Unpause() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Unpause(&_GatewayZEVM.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Unpause() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Unpause(&_GatewayZEVM.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVM *GatewayZEVMTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVM *GatewayZEVMSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.UpgradeToAndCall(&_GatewayZEVM.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.UpgradeToAndCall(&_GatewayZEVM.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Withdraw(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "withdraw", receiver, amount, zrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Withdraw(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Withdraw(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, revertOptions)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xc5356f6d.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Withdraw1(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "withdraw1", receiver, amount, zrc20, gasLimit, revertOptions)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xc5356f6d.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMSession) Withdraw1(receiver []byte, amount *big.Int, zrc20 common.Address, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Withdraw1(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, gasLimit, revertOptions)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xc5356f6d.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Withdraw1(receiver []byte, amount *big.Int, zrc20 common.Address, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Withdraw1(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, gasLimit, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactor) WithdrawAndCall0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.contract.Transact(opts, "withdrawAndCall0", receiver, amount, zrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMSession) WithdrawAndCall0(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.WithdrawAndCall0(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) WithdrawAndCall0(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVM.Contract.WithdrawAndCall0(&_GatewayZEVM.TransactOpts, receiver, amount, zrc20, message, callOptions, revertOptions)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayZEVM *GatewayZEVMTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVM.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayZEVM *GatewayZEVMSession) Receive() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Receive(&_GatewayZEVM.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayZEVM *GatewayZEVMTransactorSession) Receive() (*types.Transaction, error) {
	return _GatewayZEVM.Contract.Receive(&_GatewayZEVM.TransactOpts)
}

// GatewayZEVMCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayZEVM contract.
type GatewayZEVMCalledIterator struct {
	Event *GatewayZEVMCalled // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMCalled)
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
		it.Event = new(GatewayZEVMCalled)
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
func (it *GatewayZEVMCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMCalled represents a Called event raised by the GatewayZEVM contract.
type GatewayZEVMCalled struct {
	Sender        common.Address
	Zrc20         common.Address
	Receiver      []byte
	Message       []byte
	CallOptions   CallOptions
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, zrc20 []common.Address) (*GatewayZEVMCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMCalledIterator{contract: _GatewayZEVM.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMCalled, sender []common.Address, zrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMCalled)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Called", log); err != nil {
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

// ParseCalled is a log parse operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) ParseCalled(log types.Log) (*GatewayZEVMCalled, error) {
	event := new(GatewayZEVMCalled)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayZEVM contract.
type GatewayZEVMInitializedIterator struct {
	Event *GatewayZEVMInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInitialized)
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
		it.Event = new(GatewayZEVMInitialized)
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
func (it *GatewayZEVMInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInitialized represents a Initialized event raised by the GatewayZEVM contract.
type GatewayZEVMInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayZEVMInitializedIterator, error) {

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInitializedIterator{contract: _GatewayZEVM.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInitialized)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseInitialized(log types.Log) (*GatewayZEVMInitialized, error) {
	event := new(GatewayZEVMInitialized)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GatewayZEVM contract.
type GatewayZEVMPausedIterator struct {
	Event *GatewayZEVMPaused // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMPaused)
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
		it.Event = new(GatewayZEVMPaused)
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
func (it *GatewayZEVMPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMPaused represents a Paused event raised by the GatewayZEVM contract.
type GatewayZEVMPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterPaused(opts *bind.FilterOpts) (*GatewayZEVMPausedIterator, error) {

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMPausedIterator{contract: _GatewayZEVM.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GatewayZEVMPaused) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMPaused)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParsePaused(log types.Log) (*GatewayZEVMPaused, error) {
	event := new(GatewayZEVMPaused)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GatewayZEVM contract.
type GatewayZEVMRoleAdminChangedIterator struct {
	Event *GatewayZEVMRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMRoleAdminChanged)
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
		it.Event = new(GatewayZEVMRoleAdminChanged)
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
func (it *GatewayZEVMRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMRoleAdminChanged represents a RoleAdminChanged event raised by the GatewayZEVM contract.
type GatewayZEVMRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GatewayZEVMRoleAdminChangedIterator, error) {

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

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMRoleAdminChangedIterator{contract: _GatewayZEVM.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GatewayZEVMRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMRoleAdminChanged)
				if err := _GatewayZEVM.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseRoleAdminChanged(log types.Log) (*GatewayZEVMRoleAdminChanged, error) {
	event := new(GatewayZEVMRoleAdminChanged)
	if err := _GatewayZEVM.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GatewayZEVM contract.
type GatewayZEVMRoleGrantedIterator struct {
	Event *GatewayZEVMRoleGranted // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMRoleGranted)
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
		it.Event = new(GatewayZEVMRoleGranted)
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
func (it *GatewayZEVMRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMRoleGranted represents a RoleGranted event raised by the GatewayZEVM contract.
type GatewayZEVMRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayZEVMRoleGrantedIterator, error) {

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

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMRoleGrantedIterator{contract: _GatewayZEVM.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GatewayZEVMRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMRoleGranted)
				if err := _GatewayZEVM.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseRoleGranted(log types.Log) (*GatewayZEVMRoleGranted, error) {
	event := new(GatewayZEVMRoleGranted)
	if err := _GatewayZEVM.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GatewayZEVM contract.
type GatewayZEVMRoleRevokedIterator struct {
	Event *GatewayZEVMRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMRoleRevoked)
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
		it.Event = new(GatewayZEVMRoleRevoked)
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
func (it *GatewayZEVMRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMRoleRevoked represents a RoleRevoked event raised by the GatewayZEVM contract.
type GatewayZEVMRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayZEVMRoleRevokedIterator, error) {

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

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMRoleRevokedIterator{contract: _GatewayZEVM.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GatewayZEVMRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMRoleRevoked)
				if err := _GatewayZEVM.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseRoleRevoked(log types.Log) (*GatewayZEVMRoleRevoked, error) {
	event := new(GatewayZEVMRoleRevoked)
	if err := _GatewayZEVM.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GatewayZEVM contract.
type GatewayZEVMUnpausedIterator struct {
	Event *GatewayZEVMUnpaused // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUnpaused)
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
		it.Event = new(GatewayZEVMUnpaused)
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
func (it *GatewayZEVMUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUnpaused represents a Unpaused event raised by the GatewayZEVM contract.
type GatewayZEVMUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GatewayZEVMUnpausedIterator, error) {

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUnpausedIterator{contract: _GatewayZEVM.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUnpaused) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUnpaused)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseUnpaused(log types.Log) (*GatewayZEVMUnpaused, error) {
	event := new(GatewayZEVMUnpaused)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayZEVM contract.
type GatewayZEVMUpgradedIterator struct {
	Event *GatewayZEVMUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgraded)
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
		it.Event = new(GatewayZEVMUpgraded)
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
func (it *GatewayZEVMUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgraded represents a Upgraded event raised by the GatewayZEVM contract.
type GatewayZEVMUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayZEVMUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradedIterator{contract: _GatewayZEVM.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgraded)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GatewayZEVM *GatewayZEVMFilterer) ParseUpgraded(log types.Log) (*GatewayZEVMUpgraded, error) {
	event := new(GatewayZEVMUpgraded)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayZEVM contract.
type GatewayZEVMWithdrawnIterator struct {
	Event *GatewayZEVMWithdrawn // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMWithdrawn)
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
		it.Event = new(GatewayZEVMWithdrawn)
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
func (it *GatewayZEVMWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMWithdrawn represents a Withdrawn event raised by the GatewayZEVM contract.
type GatewayZEVMWithdrawn struct {
	Sender          common.Address
	ChainId         *big.Int
	Receiver        []byte
	Zrc20           common.Address
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	CallOptions     CallOptions
	RevertOptions   RevertOptions
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMWithdrawnIterator{contract: _GatewayZEVM.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayZEVMWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMWithdrawn)
				if err := _GatewayZEVM.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) ParseWithdrawn(log types.Log) (*GatewayZEVMWithdrawn, error) {
	event := new(GatewayZEVMWithdrawn)
	if err := _GatewayZEVM.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayZEVM contract.
type GatewayZEVMWithdrawnAndCalledIterator struct {
	Event *GatewayZEVMWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMWithdrawnAndCalled)
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
		it.Event = new(GatewayZEVMWithdrawnAndCalled)
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
func (it *GatewayZEVMWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayZEVM contract.
type GatewayZEVMWithdrawnAndCalled struct {
	Sender          common.Address
	ChainId         *big.Int
	Receiver        []byte
	Zrc20           common.Address
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	CallOptions     CallOptions
	RevertOptions   RevertOptions
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMWithdrawnAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVM.contract.FilterLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMWithdrawnAndCalledIterator{contract: _GatewayZEVM.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMWithdrawnAndCalled, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVM.contract.WatchLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMWithdrawnAndCalled)
				if err := _GatewayZEVM.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVM *GatewayZEVMFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayZEVMWithdrawnAndCalled, error) {
	event := new(GatewayZEVMWithdrawnAndCalled)
	if err := _GatewayZEVM.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
