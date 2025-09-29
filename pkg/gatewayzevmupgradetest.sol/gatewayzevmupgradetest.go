// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayzevmupgradetest

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

// GatewayZEVMUpgradeTestMetaData contains all meta data concerning the GatewayZEVMUpgradeTest contract.
var GatewayZEVMUpgradeTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROTOCOL_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndRevert\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndRevert\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeAbort\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"abortContext\",\"type\":\"tuple\",\"internalType\":\"structAbortContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outgoing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMaxGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getMaxMessageSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getMaxRevertGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getMinGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRegistryAddress\",\"inputs\":[{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallOnRevertNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RevertGasLimitExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroGasPrice\",\"inputs\":[]}]",
	Bin: "0x60a060405230608052348015610013575f80fd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516141556100f95f395f818161226e01528181612297015261244901526141555ff3fe60806040526004361061027a575f3560e01c80637c0dcb5f1161014b578063c34da122116100c6578063e63ab1e91161007c578063edc5b62e11610062578063edc5b62e146107b8578063f340fa01146107cd578063f45346dc146107e0575f80fd5b8063e63ab1e914610772578063e90b9e5e146107a5575f80fd5b8063c5356f6d116100ac578063c5356f6d1461071f578063d547741f1461073e578063e279a72a1461075d575f80fd5b8063c34da122146106ed578063c39aca3714610700575f80fd5b8063a217fddf1161011b578063ad3cb1cc11610101578063ad3cb1cc14610665578063bcbe9365146106ba578063bcf7f32b146106ce575f80fd5b8063a217fddf14610633578063ab7b499314610646575f80fd5b80637c0dcb5f1461057e5780638456cb591461059d57806391d14854146105b15780639d4ba46514610614575f80fd5b80632f2ff15d116101f55780634f1ef286116101ab5780635c975abb116101915780635c975abb1461050a5780637b103999146105405780637b15118b1461055f575f80fd5b80634f1ef286146104e357806352d1902d146104f6575f80fd5b806336568abe116101db57806336568abe146104915780633f4ba83a146104b0578063485cc955146104c4575f80fd5b80632f2ff15d1461045f57806330b103421461047e575f80fd5b806320c612fa1161024a57806322f0ce891161023057806322f0ce89146103d8578063248a9ca3146103eb5780632722feee14610438575f80fd5b806320c612fa1461038357806321e093b1146103a2575f80fd5b806301ffc9a7146102f257806306cb898314610326578063184b0793146103455780632095dedb14610364575f80fd5b366102ee576102876107ff565b5f546001600160a01b031633148015906102b557503373735b14bb79463307aacbed86daf3322b1e6226ab14155b156102ec576040517fb3af013700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b5f80fd5b3480156102fd575f80fd5b5061031161030c366004613146565b61085d565b60405190151581526020015b60405180910390f35b348015610331575f80fd5b506102ec6103403660046132d5565b6108f5565b348015610350575f80fd5b506102ec61035f36600461339c565b610934565b34801561036f575f80fd5b506102ec61037e3660046133e9565b610a39565b34801561038e575f80fd5b50622625a05b60405190815260200161031d565b3480156103ad575f80fd5b505f546103c0906001600160a01b031681565b6040516001600160a01b03909116815260200161031d565b6102ec6103e636600461343c565b610ae4565b3480156103f6575f80fd5b506103946104053660046134a9565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b348015610443575f80fd5b506103c073735b14bb79463307aacbed86daf3322b1e6226ab81565b34801561046a575f80fd5b506102ec6104793660046134c0565b610bc6565b6102ec61048c3660046134f3565b610c0f565b34801561049c575f80fd5b506102ec6104ab3660046134c0565b610d2e565b3480156104bb575f80fd5b506102ec610d7a565b3480156104cf575f80fd5b506102ec6104de36600461356c565b610daf565b6102ec6104f1366004613598565b611001565b348015610501575f80fd5b5061039461101c565b348015610515575f80fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610311565b34801561054b575f80fd5b506001546103c0906001600160a01b031681565b34801561056a575f80fd5b506102ec6105793660046135db565b61104a565b348015610589575f80fd5b506102ec61059836600461369c565b61112a565b3480156105a8575f80fd5b506102ec6112d2565b3480156105bc575f80fd5b506103116105cb3660046134c0565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561061f575f80fd5b506102ec61062e36600461371c565b611304565b34801561063e575f80fd5b506103945f81565b348015610651575f80fd5b506102ec610660366004613778565b61147d565b348015610670575f80fd5b506106ad6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161031d91906137c1565b3480156106c5575f80fd5b50610b40610394565b3480156106d9575f80fd5b506102ec6106e83660046137d3565b611502565b6102ec6106fb36600461386b565b61160c565b34801561070b575f80fd5b506102ec61071a3660046137d3565b6116db565b34801561072a575f80fd5b506102ec6107393660046138c5565b6117b4565b348015610749575f80fd5b506102ec6107583660046134c0565b6118a4565b348015610768575f80fd5b50620186a0610394565b34801561077d575f80fd5b506103947f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6102ec6107b336600461339c565b6118e7565b3480156107c3575f80fd5b50621e8480610394565b6102ec6107db366004613778565b611a02565b3480156107eb575f80fd5b506102ec6107fa36600461394e565b611ab2565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff161561085b576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806108ef57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6108fd6107ff565b610908828285611b83565b61092c8686868661091e3688900388018861399a565b610927876139f0565b611bf5565b505050505050565b61093c611c57565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610989576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109916107ff565b61099a82611cd8565b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063c9028a36906109df908490600401613b1e565b5f604051808303815f87803b1580156109f6575f80fd5b505af1158015610a08573d5f803e3d5ffd5b50505050610a3560017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5050565b610a41611c57565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a8e576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a966107ff565b610a9f82611cd8565b6040517f2d4cfb7e0000000000000000000000000000000000000000000000000000000081526001600160a01b03831690632d4cfb7e906109df908490600401613b8c565b610aec6107ff565b610af4611c57565b610aff833483611d3e565b5f805f610b0c855f611d59565b925092509250610b303473735b14bb79463307aacbed86daf3322b1e6226ab611f3d565b5f54604080518082018252838152600160208201529051879233927f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c92610b8d928c926001600160a01b039091169134918b918b91908e90613cb1565b60405180910390a3505050610bc160017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610bff81611fd8565b610c098383611fe2565b50505050565b610c17611c57565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610c64576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c6c6107ff565b610c8c348473735b14bb79463307aacbed86daf3322b1e6226ab306120cc565b6040517febf9b2aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0384169063ebf9b2aa903490610cd790889087908790600401613d7f565b5f604051808303818588803b158015610cee575f80fd5b505af1158015610d00573d5f803e3d5ffd5b5050505050610c0960017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6001600160a01b0381163314610d70576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610bc182826120e9565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610da481611fd8565b610dac6121ab565b50565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f81158015610df95750825b90505f8267ffffffffffffffff166001148015610e155750303b155b905081158015610e23575080155b15610e5a576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610ebb5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b0387161580610ed857506001600160a01b038616155b15610f0f576040517f7138356f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f1761223b565b610f1f61223b565b610f27612243565b610f2f612253565b610f395f87611fe2565b50610f647f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a87611fe2565b505f80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0389161790558315610ff85784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b611009612263565b61101282612333565b610a35828261233d565b5f61102561243e565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b6110526107ff565b6110608787868686866124a0565b5f61106d878785356124bd565b90505f336001600160a01b03167fd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a8a898b868c6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110db573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110ff9190613da4565b8c8c8c8c60405161111899989796959493929190613dbb565b60405180910390a35050505050505050565b6111326107ff565b61113d848483611d3e565b5f6111a88484856001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa15801561117f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111a39190613da4565b6124bd565b90505f336001600160a01b03167f5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d9787868886896001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611216573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061123a9190613da4565b60405180604001604052808c6001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa158015611281573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112a59190613da4565b815260016020909101526040516112c3969594939291908c90613cb1565b60405180910390a35050505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6112fc81611fd8565b610dac612587565b61130c611c57565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611359576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113616107ff565b61138284848473735b14bb79463307aacbed86daf3322b1e6226ab30612600565b61138d84838561262d565b6113e2576040517f40a143ba0000000000000000000000000000000000000000000000000000000081526001600160a01b03808616600483015283166024820152604481018490526064015b60405180910390fd5b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063c9028a3690611427908490600401613b1e565b5f604051808303815f87803b15801561143e575f80fd5b505af1158015611450573d5f803e3d5ffd5b50505050610c0960017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f61148781611fd8565b6001600160a01b0382166114c7576040517f7138356f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b61150a611c57565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611557576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61155f6107ff565b61156985846126c5565b6040517f5bcfd6160000000000000000000000000000000000000000000000000000000081526001600160a01b03841690635bcfd616906115b69089908990899088908890600401613e44565b5f604051808303815f87803b1580156115cd575f80fd5b505af11580156115df573d5f803e3d5ffd5b5050505061092c60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6116146107ff565b61161c611c57565b61162a8634868686866124a0565b5f80611637878535611d59565b509150915061165a3473735b14bb79463307aacbed86daf3322b1e6226ab611f3d565b5f54604051889133917fd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a916116a8918d916001600160a01b0316903490899089908f908f908f908f90613dbb565b60405180910390a3505061092c60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6116e3611c57565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611730576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6117386107ff565b61175985858573735b14bb79463307aacbed86daf3322b1e6226ab30612600565b61176485848661262d565b611569576040517f40a143ba0000000000000000000000000000000000000000000000000000000081526001600160a01b03808716600483015284166024820152604481018590526064016113d9565b6117bc6107ff565b6117c7858583611d3e565b6117d0826126d7565b5f6117dc8585856124bd565b90505f336001600160a01b03167f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c888789868a6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561184a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061186e9190613da4565b6040805180820182528c8152600160208201529051611894969594939291908c90613cb1565b60405180910390a3505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546118dd81611fd8565b610c0983836120e9565b6118ef611c57565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461193c576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119446107ff565b611964348373735b14bb79463307aacbed86daf3322b1e6226ab306120cc565b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063c9028a369034906119ab908590600401613b1e565b5f604051808303818588803b1580156119c2575f80fd5b505af11580156119d4573d5f803e3d5ffd5b5050505050610a3560017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b611a0a611c57565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611a57576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611a5f6107ff565b611a7f348273735b14bb79463307aacbed86daf3322b1e6226ab306120cc565b611a893482611f3d565b610dac60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611aff576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611b076107ff565b611b2883838373735b14bb79463307aacbed86daf3322b1e6226ab30612600565b611b3383828461262d565b610bc1576040517f40a143ba0000000000000000000000000000000000000000000000000000000081526001600160a01b03808516600483015282166024820152604481018390526064016113d9565b611b8d83356126d7565b611ba581611b9e6060850185613e8a565b9050612721565b621e848082608001351115610bc1576040517fec86721e00000000000000000000000000000000000000000000000000000000815260808301356004820152621e848060248201526044016113d9565b611bfe86612775565b611c0b85835f01516127af565b50846001600160a01b0316336001600160a01b03167f306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e48887878787604051611894959493929190613eeb565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611cd2576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b6001600160a01b038116610dac576040517f7138356f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b611d4783612775565b611d5082612834565b610bc18161286d565b6001546040517f804ea334000000000000000000000000000000000000000000000000000000008152600481018490525f918291829182916001600160a01b039091169063804ea334906024015f60405180830381865afa158015611dc0573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611de79190810190613fda565b5090505f816001600160a01b031663f2441b326040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e27573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e4b919061401e565b6001600160a01b031663d7fd7afb886040518263ffffffff1660e01b8152600401611e7891815260200190565b602060405180830381865afa158015611e93573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611eb79190613da4565b9050805f03611ef2576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855f03611f0557611f0287612922565b95505b611f0e87612a1a565b935083611f1b8783614066565b611f25919061407d565b9450611f318286612b06565b50849150509250925092565b5f816001600160a01b0316836040515f6040518083038185875af1925050503d805f8114611f86576040519150601f19603f3d011682016040523d82523d5f602084013e611f8b565b606091505b5050905080610bc1576040517ff279af7e0000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602481018490526044016113d9565b610dac8133612bb2565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff166120c3575f848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556120793390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506108ef565b5f9150506108ef565b6120d583611cd8565b6120de84612834565b610c09838383612c3e565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16156120c3575f848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506108ef565b6121b3612ca6565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b61085b612d01565b61224b612d01565b61085b612d68565b61225b612d01565b61085b612db9565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806122fc57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166122f07f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b1561085b576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610a3581611fd8565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612397575060408051601f3d908101601f1916820190925261239491810190613da4565b60015b6123d8576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016113d9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612434576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016113d9565b610bc18383612dc1565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461085b576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6124a986612775565b6124b285612834565b61092c828285611b83565b5f806124c984846127af565b90506124d784333088612e16565b61252b576040517f489ca9b70000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152336024820152306044820152606481018690526084016113d9565b6125358486612eb8565b61257d576040517f7112ae770000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602481018690526044016113d9565b90505b9392505050565b61258f6107ff565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583361221d565b61260985611cd8565b61261283611cd8565b61261b84612834565b612626838383612c3e565b5050505050565b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152602482018390525f91908516906347e7ef24906044016020604051808303815f875af19250505080156126b3575060408051601f3d908101601f191682019092526126b091810190614090565b60015b6126be57505f612580565b9050612580565b6126ce82611cd8565b610a3581611cd8565b620186a08110806126ea5750622625a081115b15610dac576040517f98bdb2e000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61272c828461407d565b9050610b40811115610bc1576040517fcd6f4e6d00000000000000000000000000000000000000000000000000000000815260048101829052610b4060248201526044016113d9565b80515f03610dac576040517f7138356f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805f846001600160a01b031663fc5fecd5856040518263ffffffff1660e01b81526004016127e091815260200190565b6040805180830381865afa1580156127fa573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061281e91906140ab565b9150915061282c8282612b06565b949350505050565b805f03610dac576040517f5945ea5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b4061287d6060830183613e8a565b905011156128d2576128926060820182613e8a565b6040517fcd6f4e6d0000000000000000000000000000000000000000000000000000000081526113d99250610b4090600401918252602082015260400190565b621e848081608001351115610dac576040517fec86721e00000000000000000000000000000000000000000000000000000000815260808201356004820152621e848060248201526044016113d9565b600154604080517f7066b18d000000000000000000000000000000000000000000000000000000008152600481018490526024810191909152600860448201527f6761734c696d697400000000000000000000000000000000000000000000000060648201525f91620186a0916001600160a01b0390911690637066b18d906084015f60405180830381865afa9250505080156129e057506040513d5f823e601f3d908101601f191682016040526129dd91908101906140d7565b60015b6129ec57809150612a14565b805115612a0e5780806020019051810190612a079190613da4565b9250612a12565b8192505b505b50919050565b600154604080517f7066b18d000000000000000000000000000000000000000000000000000000008152600481018490526024810191909152600f60448201527f70726f746f636f6c466c6174466565000000000000000000000000000000000060648201525f916001600160a01b031690637066b18d906084015f60405180830381865afa925050508015612ad157506040513d5f823e601f3d908101601f19168201604052612ace91908101906140d7565b60015b612adc57505f919050565b805115612afe5780806020019051810190612af79190613da4565b9150612a14565b505f92915050565b612b1282333084612e16565b612b60576040517f667011120000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152306024820152604481018290526064016113d9565b612b6a8282612eb8565b610a35576040517f7112ae770000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602481018290526044016113d9565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610a35576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044016113d9565b816001600160a01b0316836001600160a01b03161480612c6f5750806001600160a01b0316836001600160a01b0316145b15610bc1576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1661085b576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661085b576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612d70612d01565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b611d18612d01565b612dca82612f47565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115612e0e57610bc18282612fee565b610a35613057565b6040517f23b872dd0000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301528381166024830152604482018390525f91908616906323b872dd906064016020604051808303815f875af1925050508015612ea4575060408051601f3d908101601f19168201909252612ea191810190614090565b60015b612eaf57505f61282c565b95945050505050565b6040517f42966c68000000000000000000000000000000000000000000000000000000008152600481018290525f906001600160a01b038416906342966c68906024016020604051808303815f875af1925050508015612f35575060408051601f3d908101601f19168201909252612f3291810190614090565b60015b612f4057505f6108ef565b90506108ef565b806001600160a01b03163b5f03612f95576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016113d9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f80846001600160a01b03168460405161300a9190614109565b5f60405180830381855af49150503d805f8114613042576040519150601f19603f3d011682016040523d82523d5f602084013e613047565b606091505b5091509150612eaf85838361308f565b341561085b576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060826130a45761309f82613104565b612580565b81511580156130bb57506001600160a01b0384163b155b156130fd576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016113d9565b5080612580565b8051156131145780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60208284031215613156575f80fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114612580575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff811182821017156131db576131db613185565b604052919050565b5f67ffffffffffffffff8211156131fc576131fc613185565b50601f01601f191660200190565b5f82601f830112613219575f80fd5b813561322c613227826131e3565b6131b2565b818152846020838601011115613240575f80fd5b816020850160208301375f918101602001919091529392505050565b6001600160a01b0381168114610dac575f80fd5b5f8083601f840112613280575f80fd5b50813567ffffffffffffffff811115613297575f80fd5b6020830191508360208285010111156132ae575f80fd5b9250929050565b5f60408284031215612a14575f80fd5b5f60a08284031215612a14575f80fd5b5f805f805f8060c087890312156132ea575f80fd5b863567ffffffffffffffff811115613300575f80fd5b61330c89828a0161320a565b965050602087013561331d8161325c565b9450604087013567ffffffffffffffff811115613338575f80fd5b61334489828a01613270565b9095509350613358905088606089016132b5565b915060a087013567ffffffffffffffff811115613373575f80fd5b61337f89828a016132c5565b9150509295509295509295565b5f60808284031215612a14575f80fd5b5f80604083850312156133ad575f80fd5b82356133b88161325c565b9150602083013567ffffffffffffffff8111156133d3575f80fd5b6133df8582860161338c565b9150509250929050565b5f80604083850312156133fa575f80fd5b82356134058161325c565b9150602083013567ffffffffffffffff811115613420575f80fd5b830160c08186031215613431575f80fd5b809150509250929050565b5f805f6060848603121561344e575f80fd5b833567ffffffffffffffff811115613464575f80fd5b6134708682870161320a565b93505060208401359150604084013567ffffffffffffffff811115613493575f80fd5b61349f868287016132c5565b9150509250925092565b5f602082840312156134b9575f80fd5b5035919050565b5f80604083850312156134d1575f80fd5b8235915060208301356134318161325c565b5f60608284031215612a14575f80fd5b5f805f8060608587031215613506575f80fd5b843567ffffffffffffffff81111561351c575f80fd5b613528878288016134e3565b94505060208501356135398161325c565b9250604085013567ffffffffffffffff811115613554575f80fd5b61356087828801613270565b95989497509550505050565b5f806040838503121561357d575f80fd5b82356135888161325c565b915060208301356134318161325c565b5f80604083850312156135a9575f80fd5b82356135b48161325c565b9150602083013567ffffffffffffffff8111156135cf575f80fd5b6133df8582860161320a565b5f805f805f805f60e0888a0312156135f1575f80fd5b873567ffffffffffffffff811115613607575f80fd5b6136138a828b0161320a565b97505060208801359550604088013561362b8161325c565b9450606088013567ffffffffffffffff811115613646575f80fd5b6136528a828b01613270565b909550935061366690508960808a016132b5565b915060c088013567ffffffffffffffff811115613681575f80fd5b61368d8a828b016132c5565b91505092959891949750929550565b5f805f80608085870312156136af575f80fd5b843567ffffffffffffffff8111156136c5575f80fd5b6136d18782880161320a565b9450506020850135925060408501356136e98161325c565b9150606085013567ffffffffffffffff811115613704575f80fd5b613710878288016132c5565b91505092959194509250565b5f805f806080858703121561372f575f80fd5b843561373a8161325c565b93506020850135925060408501356137518161325c565b9150606085013567ffffffffffffffff81111561376c575f80fd5b6137108782880161338c565b5f60208284031215613788575f80fd5b81356125808161325c565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6125806020830184613793565b5f805f805f8060a087890312156137e8575f80fd5b863567ffffffffffffffff8111156137fe575f80fd5b61380a89828a016134e3565b965050602087013561381b8161325c565b94506040870135935060608701356138328161325c565b9250608087013567ffffffffffffffff81111561384d575f80fd5b61385989828a01613270565b979a9699509497509295939492505050565b5f805f805f8060c08789031215613880575f80fd5b863567ffffffffffffffff811115613896575f80fd5b6138a289828a0161320a565b96505060208701359450604087013567ffffffffffffffff811115613338575f80fd5b5f805f805f60a086880312156138d9575f80fd5b853567ffffffffffffffff8111156138ef575f80fd5b6138fb8882890161320a565b9550506020860135935060408601356139138161325c565b925060608601359150608086013567ffffffffffffffff811115613935575f80fd5b613941888289016132c5565b9150509295509295909350565b5f805f60608486031215613960575f80fd5b833561396b8161325c565b92506020840135915060408401356139828161325c565b809150509250925092565b8015158114610dac575f80fd5b5f60408284031280156139ab575f80fd5b506040805190810167ffffffffffffffff811182821017156139cf576139cf613185565b6040528235815260208301356139e48161398d565b60208201529392505050565b5f60a08236031215613a00575f80fd5b60405160a0810167ffffffffffffffff81118282101715613a2357613a23613185565b6040528235613a318161325c565b81526020830135613a418161398d565b60208201526040830135613a548161325c565b6040820152606083013567ffffffffffffffff811115613a72575f80fd5b613a7e3682860161320a565b606083015250608092830135928101929092525090565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613ac8575f80fd5b830160208101925035905067ffffffffffffffff811115613ae7575f80fd5b8036038213156132ae575f80fd5b81835281816020850137505f602082840101525f6020601f19601f840116840101905092915050565b602081525f8235613b2e8161325c565b6001600160a01b0381166020840152506020830135613b4c8161325c565b6001600160a01b0381166040840152505f6040840135905080606084015250613b786060840184613a95565b608080850152612eaf60a085018284613af5565b602081525f613b9b8384613a95565b60c06020850152613bb060e085018284613af5565b9150506020840135613bc18161325c565b6001600160a01b0316604084810191909152840135606080850191909152840135613beb8161398d565b8015156080850152505f608085013590508060a085015250613c1060a0850185613a95565b601f198584030160c0860152613c27838284613af5565b9695505050505050565b5f8135613c3d8161325c565b6001600160a01b031683526020820135613c568161398d565b151560208401526040820135613c6b8161325c565b6001600160a01b03166040840152613c866060830183613a95565b60a06060860152613c9b60a086018284613af5565b6080948501359590940194909452509092915050565b61012081525f613cc561012083018a613793565b6001600160a01b03891660208401528760408401528660608401528560808401528281038060a08501525f8252613d0b60c0850187805182526020908101511515910152565b6020810161010085015250613d236020820185613c31565b9a9950505050505050505050565b5f613d3c8283613a95565b60608552613d4e606086018284613af5565b9150506020830135613d5f8161325c565b6001600160a01b0316602085015260409283013592909301919091525090565b604081525f613d916040830186613d31565b8281036020840152613c27818587613af5565b5f60208284031215613db4575f80fd5b5051919050565b61012081525f613dcf61012083018c613793565b6001600160a01b038b16602084015289604084015288606084015287608084015282810360a0840152613e03818789613af5565b853560c085015290506020850135613e1a8161398d565b151560e0840152828103610100840152613e348185613c31565b9c9b505050505050505050505050565b608081525f613e566080830188613d31565b6001600160a01b03871660208401528560408401528281036060840152613e7e818587613af5565b98975050505050505050565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613ebd575f80fd5b83018035915067ffffffffffffffff821115613ed7575f80fd5b6020019150368190038213156132ae575f80fd5b60a081525f613efd60a0830188613793565b8281036020840152613f10818789613af5565b85516040850152602086015115156060850152905082810360808401526001600160a01b0384511681526020840151151560208201526001600160a01b036040850151166040820152606084015160a06060830152613f7260a0830182613793565b90506080850151608083015280925050509695505050505050565b5f82601f830112613f9c575f80fd5b8151613faa613227826131e3565b818152846020838601011115613fbe575f80fd5b8160208501602083015e5f918101602001919091529392505050565b5f8060408385031215613feb575f80fd5b8251613ff68161325c565b602084015190925067ffffffffffffffff811115614012575f80fd5b6133df85828601613f8d565b5f6020828403121561402e575f80fd5b81516125808161325c565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820281158282048414176108ef576108ef614039565b808201808211156108ef576108ef614039565b5f602082840312156140a0575f80fd5b81516125808161398d565b5f80604083850312156140bc575f80fd5b82516140c78161325c565b6020939093015192949293505050565b5f602082840312156140e7575f80fd5b815167ffffffffffffffff8111156140fd575f80fd5b61282c84828501613f8d565b5f82518060208501845e5f92019182525091905056fea264697066735822122063ed7a9f178420561ab536209a6d477e62f0a1a06e173be627d3331fc810efeb64736f6c634300081a0033",
}

// GatewayZEVMUpgradeTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayZEVMUpgradeTestMetaData.ABI instead.
var GatewayZEVMUpgradeTestABI = GatewayZEVMUpgradeTestMetaData.ABI

// GatewayZEVMUpgradeTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayZEVMUpgradeTestMetaData.Bin instead.
var GatewayZEVMUpgradeTestBin = GatewayZEVMUpgradeTestMetaData.Bin

// DeployGatewayZEVMUpgradeTest deploys a new Ethereum contract, binding an instance of GatewayZEVMUpgradeTest to it.
func DeployGatewayZEVMUpgradeTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayZEVMUpgradeTest, error) {
	parsed, err := GatewayZEVMUpgradeTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayZEVMUpgradeTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayZEVMUpgradeTest{GatewayZEVMUpgradeTestCaller: GatewayZEVMUpgradeTestCaller{contract: contract}, GatewayZEVMUpgradeTestTransactor: GatewayZEVMUpgradeTestTransactor{contract: contract}, GatewayZEVMUpgradeTestFilterer: GatewayZEVMUpgradeTestFilterer{contract: contract}}, nil
}

// GatewayZEVMUpgradeTest is an auto generated Go binding around an Ethereum contract.
type GatewayZEVMUpgradeTest struct {
	GatewayZEVMUpgradeTestCaller     // Read-only binding to the contract
	GatewayZEVMUpgradeTestTransactor // Write-only binding to the contract
	GatewayZEVMUpgradeTestFilterer   // Log filterer for contract events
}

// GatewayZEVMUpgradeTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayZEVMUpgradeTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMUpgradeTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayZEVMUpgradeTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMUpgradeTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayZEVMUpgradeTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMUpgradeTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayZEVMUpgradeTestSession struct {
	Contract     *GatewayZEVMUpgradeTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GatewayZEVMUpgradeTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayZEVMUpgradeTestCallerSession struct {
	Contract *GatewayZEVMUpgradeTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// GatewayZEVMUpgradeTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayZEVMUpgradeTestTransactorSession struct {
	Contract     *GatewayZEVMUpgradeTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// GatewayZEVMUpgradeTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayZEVMUpgradeTestRaw struct {
	Contract *GatewayZEVMUpgradeTest // Generic contract binding to access the raw methods on
}

// GatewayZEVMUpgradeTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayZEVMUpgradeTestCallerRaw struct {
	Contract *GatewayZEVMUpgradeTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayZEVMUpgradeTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayZEVMUpgradeTestTransactorRaw struct {
	Contract *GatewayZEVMUpgradeTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayZEVMUpgradeTest creates a new instance of GatewayZEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayZEVMUpgradeTest(address common.Address, backend bind.ContractBackend) (*GatewayZEVMUpgradeTest, error) {
	contract, err := bindGatewayZEVMUpgradeTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTest{GatewayZEVMUpgradeTestCaller: GatewayZEVMUpgradeTestCaller{contract: contract}, GatewayZEVMUpgradeTestTransactor: GatewayZEVMUpgradeTestTransactor{contract: contract}, GatewayZEVMUpgradeTestFilterer: GatewayZEVMUpgradeTestFilterer{contract: contract}}, nil
}

// NewGatewayZEVMUpgradeTestCaller creates a new read-only instance of GatewayZEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayZEVMUpgradeTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayZEVMUpgradeTestCaller, error) {
	contract, err := bindGatewayZEVMUpgradeTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestCaller{contract: contract}, nil
}

// NewGatewayZEVMUpgradeTestTransactor creates a new write-only instance of GatewayZEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayZEVMUpgradeTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayZEVMUpgradeTestTransactor, error) {
	contract, err := bindGatewayZEVMUpgradeTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestTransactor{contract: contract}, nil
}

// NewGatewayZEVMUpgradeTestFilterer creates a new log filterer instance of GatewayZEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayZEVMUpgradeTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayZEVMUpgradeTestFilterer, error) {
	contract, err := bindGatewayZEVMUpgradeTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestFilterer{contract: contract}, nil
}

// bindGatewayZEVMUpgradeTest binds a generic wrapper to an already deployed contract.
func bindGatewayZEVMUpgradeTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayZEVMUpgradeTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMUpgradeTest.Contract.GatewayZEVMUpgradeTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.GatewayZEVMUpgradeTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.GatewayZEVMUpgradeTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMUpgradeTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.DEFAULTADMINROLE(&_GatewayZEVMUpgradeTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.DEFAULTADMINROLE(&_GatewayZEVMUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.PAUSERROLE(&_GatewayZEVMUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.PAUSERROLE(&_GatewayZEVMUpgradeTest.CallOpts)
}

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) PROTOCOLADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "PROTOCOL_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) PROTOCOLADDRESS() (common.Address, error) {
	return _GatewayZEVMUpgradeTest.Contract.PROTOCOLADDRESS(&_GatewayZEVMUpgradeTest.CallOpts)
}

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) PROTOCOLADDRESS() (common.Address, error) {
	return _GatewayZEVMUpgradeTest.Contract.PROTOCOLADDRESS(&_GatewayZEVMUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayZEVMUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_GatewayZEVMUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayZEVMUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_GatewayZEVMUpgradeTest.CallOpts)
}

// GetMaxGasLimit is a free data retrieval call binding the contract method 0x20c612fa.
//
// Solidity: function getMaxGasLimit() pure returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) GetMaxGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "getMaxGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxGasLimit is a free data retrieval call binding the contract method 0x20c612fa.
//
// Solidity: function getMaxGasLimit() pure returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) GetMaxGasLimit() (*big.Int, error) {
	return _GatewayZEVMUpgradeTest.Contract.GetMaxGasLimit(&_GatewayZEVMUpgradeTest.CallOpts)
}

// GetMaxGasLimit is a free data retrieval call binding the contract method 0x20c612fa.
//
// Solidity: function getMaxGasLimit() pure returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) GetMaxGasLimit() (*big.Int, error) {
	return _GatewayZEVMUpgradeTest.Contract.GetMaxGasLimit(&_GatewayZEVMUpgradeTest.CallOpts)
}

// GetMaxMessageSize is a free data retrieval call binding the contract method 0xbcbe9365.
//
// Solidity: function getMaxMessageSize() pure returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) GetMaxMessageSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "getMaxMessageSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxMessageSize is a free data retrieval call binding the contract method 0xbcbe9365.
//
// Solidity: function getMaxMessageSize() pure returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) GetMaxMessageSize() (*big.Int, error) {
	return _GatewayZEVMUpgradeTest.Contract.GetMaxMessageSize(&_GatewayZEVMUpgradeTest.CallOpts)
}

// GetMaxMessageSize is a free data retrieval call binding the contract method 0xbcbe9365.
//
// Solidity: function getMaxMessageSize() pure returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) GetMaxMessageSize() (*big.Int, error) {
	return _GatewayZEVMUpgradeTest.Contract.GetMaxMessageSize(&_GatewayZEVMUpgradeTest.CallOpts)
}

// GetMaxRevertGasLimit is a free data retrieval call binding the contract method 0xedc5b62e.
//
// Solidity: function getMaxRevertGasLimit() pure returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) GetMaxRevertGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "getMaxRevertGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxRevertGasLimit is a free data retrieval call binding the contract method 0xedc5b62e.
//
// Solidity: function getMaxRevertGasLimit() pure returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) GetMaxRevertGasLimit() (*big.Int, error) {
	return _GatewayZEVMUpgradeTest.Contract.GetMaxRevertGasLimit(&_GatewayZEVMUpgradeTest.CallOpts)
}

// GetMaxRevertGasLimit is a free data retrieval call binding the contract method 0xedc5b62e.
//
// Solidity: function getMaxRevertGasLimit() pure returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) GetMaxRevertGasLimit() (*big.Int, error) {
	return _GatewayZEVMUpgradeTest.Contract.GetMaxRevertGasLimit(&_GatewayZEVMUpgradeTest.CallOpts)
}

// GetMinGasLimit is a free data retrieval call binding the contract method 0xe279a72a.
//
// Solidity: function getMinGasLimit() pure returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) GetMinGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "getMinGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinGasLimit is a free data retrieval call binding the contract method 0xe279a72a.
//
// Solidity: function getMinGasLimit() pure returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) GetMinGasLimit() (*big.Int, error) {
	return _GatewayZEVMUpgradeTest.Contract.GetMinGasLimit(&_GatewayZEVMUpgradeTest.CallOpts)
}

// GetMinGasLimit is a free data retrieval call binding the contract method 0xe279a72a.
//
// Solidity: function getMinGasLimit() pure returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) GetMinGasLimit() (*big.Int, error) {
	return _GatewayZEVMUpgradeTest.Contract.GetMinGasLimit(&_GatewayZEVMUpgradeTest.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.GetRoleAdmin(&_GatewayZEVMUpgradeTest.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.GetRoleAdmin(&_GatewayZEVMUpgradeTest.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayZEVMUpgradeTest.Contract.HasRole(&_GatewayZEVMUpgradeTest.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayZEVMUpgradeTest.Contract.HasRole(&_GatewayZEVMUpgradeTest.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Paused() (bool, error) {
	return _GatewayZEVMUpgradeTest.Contract.Paused(&_GatewayZEVMUpgradeTest.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) Paused() (bool, error) {
	return _GatewayZEVMUpgradeTest.Contract.Paused(&_GatewayZEVMUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.ProxiableUUID(&_GatewayZEVMUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayZEVMUpgradeTest.Contract.ProxiableUUID(&_GatewayZEVMUpgradeTest.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Registry() (common.Address, error) {
	return _GatewayZEVMUpgradeTest.Contract.Registry(&_GatewayZEVMUpgradeTest.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) Registry() (common.Address, error) {
	return _GatewayZEVMUpgradeTest.Contract.Registry(&_GatewayZEVMUpgradeTest.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayZEVMUpgradeTest.Contract.SupportsInterface(&_GatewayZEVMUpgradeTest.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayZEVMUpgradeTest.Contract.SupportsInterface(&_GatewayZEVMUpgradeTest.CallOpts, interfaceId)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) ZetaToken() (common.Address, error) {
	return _GatewayZEVMUpgradeTest.Contract.ZetaToken(&_GatewayZEVMUpgradeTest.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) ZetaToken() (common.Address, error) {
	return _GatewayZEVMUpgradeTest.Contract.ZetaToken(&_GatewayZEVMUpgradeTest.CallOpts)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Call(opts *bind.TransactOpts, receiver []byte, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "call", receiver, zrc20, message, callOptions, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Call(receiver []byte, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Call(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, zrc20, message, callOptions, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Call(receiver []byte, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Call(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, zrc20, message, callOptions, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address target) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Deposit(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "deposit", target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address target) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Deposit(target common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Deposit(&_GatewayZEVMUpgradeTest.TransactOpts, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address target) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Deposit(target common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Deposit(&_GatewayZEVMUpgradeTest.TransactOpts, target)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Deposit0(opts *bind.TransactOpts, zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "deposit0", zrc20, amount, target)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Deposit0(zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Deposit0(&_GatewayZEVMUpgradeTest.TransactOpts, zrc20, amount, target)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Deposit0(zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Deposit0(&_GatewayZEVMUpgradeTest.TransactOpts, zrc20, amount, target)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x30b10342.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address target, bytes message) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) DepositAndCall(opts *bind.TransactOpts, context MessageContext, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "depositAndCall", context, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x30b10342.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address target, bytes message) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) DepositAndCall(context MessageContext, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, context, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x30b10342.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address target, bytes message) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) DepositAndCall(context MessageContext, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, context, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) DepositAndCall0(opts *bind.TransactOpts, context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "depositAndCall0", context, zrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) DepositAndCall0(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndCall0(&_GatewayZEVMUpgradeTest.TransactOpts, context, zrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) DepositAndCall0(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndCall0(&_GatewayZEVMUpgradeTest.TransactOpts, context, zrc20, amount, target, message)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) DepositAndRevert(opts *bind.TransactOpts, zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "depositAndRevert", zrc20, amount, target, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) DepositAndRevert(zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndRevert(&_GatewayZEVMUpgradeTest.TransactOpts, zrc20, amount, target, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address zrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) DepositAndRevert(zrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndRevert(&_GatewayZEVMUpgradeTest.TransactOpts, zrc20, amount, target, revertContext)
}

// DepositAndRevert0 is a paid mutator transaction binding the contract method 0xe90b9e5e.
//
// Solidity: function depositAndRevert(address target, (address,address,uint256,bytes) revertContext) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) DepositAndRevert0(opts *bind.TransactOpts, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "depositAndRevert0", target, revertContext)
}

// DepositAndRevert0 is a paid mutator transaction binding the contract method 0xe90b9e5e.
//
// Solidity: function depositAndRevert(address target, (address,address,uint256,bytes) revertContext) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) DepositAndRevert0(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndRevert0(&_GatewayZEVMUpgradeTest.TransactOpts, target, revertContext)
}

// DepositAndRevert0 is a paid mutator transaction binding the contract method 0xe90b9e5e.
//
// Solidity: function depositAndRevert(address target, (address,address,uint256,bytes) revertContext) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) DepositAndRevert0(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndRevert0(&_GatewayZEVMUpgradeTest.TransactOpts, target, revertContext)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Execute(opts *bind.TransactOpts, context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "execute", context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Execute(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Execute(&_GatewayZEVMUpgradeTest.TransactOpts, context, zrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address zrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Execute(context MessageContext, zrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Execute(&_GatewayZEVMUpgradeTest.TransactOpts, context, zrc20, amount, target, message)
}

// ExecuteAbort is a paid mutator transaction binding the contract method 0x2095dedb.
//
// Solidity: function executeAbort(address target, (bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) ExecuteAbort(opts *bind.TransactOpts, target common.Address, abortContext AbortContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "executeAbort", target, abortContext)
}

// ExecuteAbort is a paid mutator transaction binding the contract method 0x2095dedb.
//
// Solidity: function executeAbort(address target, (bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) ExecuteAbort(target common.Address, abortContext AbortContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.ExecuteAbort(&_GatewayZEVMUpgradeTest.TransactOpts, target, abortContext)
}

// ExecuteAbort is a paid mutator transaction binding the contract method 0x2095dedb.
//
// Solidity: function executeAbort(address target, (bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) ExecuteAbort(target common.Address, abortContext AbortContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.ExecuteAbort(&_GatewayZEVMUpgradeTest.TransactOpts, target, abortContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) ExecuteRevert(opts *bind.TransactOpts, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "executeRevert", target, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) ExecuteRevert(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.ExecuteRevert(&_GatewayZEVMUpgradeTest.TransactOpts, target, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) ExecuteRevert(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.ExecuteRevert(&_GatewayZEVMUpgradeTest.TransactOpts, target, revertContext)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.GrantRole(&_GatewayZEVMUpgradeTest.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.GrantRole(&_GatewayZEVMUpgradeTest.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address zetaToken_, address admin_) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Initialize(opts *bind.TransactOpts, zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "initialize", zetaToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address zetaToken_, address admin_) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Initialize(zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Initialize(&_GatewayZEVMUpgradeTest.TransactOpts, zetaToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address zetaToken_, address admin_) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Initialize(zetaToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Initialize(&_GatewayZEVMUpgradeTest.TransactOpts, zetaToken_, admin_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Pause() (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Pause(&_GatewayZEVMUpgradeTest.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Pause() (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Pause(&_GatewayZEVMUpgradeTest.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.RenounceRole(&_GatewayZEVMUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.RenounceRole(&_GatewayZEVMUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.RevokeRole(&_GatewayZEVMUpgradeTest.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.RevokeRole(&_GatewayZEVMUpgradeTest.TransactOpts, role, account)
}

// SetRegistryAddress is a paid mutator transaction binding the contract method 0xab7b4993.
//
// Solidity: function setRegistryAddress(address _registry) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) SetRegistryAddress(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "setRegistryAddress", _registry)
}

// SetRegistryAddress is a paid mutator transaction binding the contract method 0xab7b4993.
//
// Solidity: function setRegistryAddress(address _registry) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) SetRegistryAddress(_registry common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.SetRegistryAddress(&_GatewayZEVMUpgradeTest.TransactOpts, _registry)
}

// SetRegistryAddress is a paid mutator transaction binding the contract method 0xab7b4993.
//
// Solidity: function setRegistryAddress(address _registry) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) SetRegistryAddress(_registry common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.SetRegistryAddress(&_GatewayZEVMUpgradeTest.TransactOpts, _registry)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Unpause() (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Unpause(&_GatewayZEVMUpgradeTest.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Unpause() (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Unpause(&_GatewayZEVMUpgradeTest.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.UpgradeToAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.UpgradeToAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x22f0ce89.
//
// Solidity: function withdraw(bytes receiver, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Withdraw(opts *bind.TransactOpts, receiver []byte, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdraw", receiver, chainId, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x22f0ce89.
//
// Solidity: function withdraw(bytes receiver, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Withdraw(receiver []byte, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Withdraw(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, chainId, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x22f0ce89.
//
// Solidity: function withdraw(bytes receiver, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Withdraw(receiver []byte, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Withdraw(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Withdraw0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdraw0", receiver, amount, zrc20, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Withdraw0(receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Withdraw0(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Withdraw0(receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Withdraw0(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, revertOptions)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xc5356f6d.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Withdraw1(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdraw1", receiver, amount, zrc20, gasLimit, revertOptions)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xc5356f6d.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Withdraw1(receiver []byte, amount *big.Int, zrc20 common.Address, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Withdraw1(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, gasLimit, revertOptions)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xc5356f6d.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Withdraw1(receiver []byte, amount *big.Int, zrc20 common.Address, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Withdraw1(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, gasLimit, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) WithdrawAndCall(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall", receiver, amount, zrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) WithdrawAndCall(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) WithdrawAndCall(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0xc34da122.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) WithdrawAndCall0(opts *bind.TransactOpts, receiver []byte, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall0", receiver, chainId, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0xc34da122.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) WithdrawAndCall0(receiver []byte, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall0(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, chainId, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0xc34da122.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) WithdrawAndCall0(receiver []byte, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall0(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, chainId, message, callOptions, revertOptions)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Receive() (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Receive(&_GatewayZEVMUpgradeTest.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Receive() (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Receive(&_GatewayZEVMUpgradeTest.TransactOpts)
}

// GatewayZEVMUpgradeTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestCalledIterator struct {
	Event *GatewayZEVMUpgradeTestCalled // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestCalled)
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
		it.Event = new(GatewayZEVMUpgradeTestCalled)
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
func (it *GatewayZEVMUpgradeTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestCalled represents a Called event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestCalled struct {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, zrc20 []common.Address) (*GatewayZEVMUpgradeTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestCalledIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestCalled, sender []common.Address, zrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestCalled)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Called", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseCalled(log types.Log) (*GatewayZEVMUpgradeTestCalled, error) {
	event := new(GatewayZEVMUpgradeTestCalled)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestInitializedIterator struct {
	Event *GatewayZEVMUpgradeTestInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestInitialized)
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
		it.Event = new(GatewayZEVMUpgradeTestInitialized)
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
func (it *GatewayZEVMUpgradeTestInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestInitialized represents a Initialized event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayZEVMUpgradeTestInitializedIterator, error) {

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestInitializedIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestInitialized)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseInitialized(log types.Log) (*GatewayZEVMUpgradeTestInitialized, error) {
	event := new(GatewayZEVMUpgradeTestInitialized)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestPausedIterator struct {
	Event *GatewayZEVMUpgradeTestPaused // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestPaused)
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
		it.Event = new(GatewayZEVMUpgradeTestPaused)
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
func (it *GatewayZEVMUpgradeTestPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestPaused represents a Paused event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterPaused(opts *bind.FilterOpts) (*GatewayZEVMUpgradeTestPausedIterator, error) {

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestPausedIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestPaused) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestPaused)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParsePaused(log types.Log) (*GatewayZEVMUpgradeTestPaused, error) {
	event := new(GatewayZEVMUpgradeTestPaused)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestRoleAdminChangedIterator struct {
	Event *GatewayZEVMUpgradeTestRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestRoleAdminChanged)
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
		it.Event = new(GatewayZEVMUpgradeTestRoleAdminChanged)
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
func (it *GatewayZEVMUpgradeTestRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestRoleAdminChanged represents a RoleAdminChanged event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GatewayZEVMUpgradeTestRoleAdminChangedIterator, error) {

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

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestRoleAdminChangedIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestRoleAdminChanged)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseRoleAdminChanged(log types.Log) (*GatewayZEVMUpgradeTestRoleAdminChanged, error) {
	event := new(GatewayZEVMUpgradeTestRoleAdminChanged)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestRoleGrantedIterator struct {
	Event *GatewayZEVMUpgradeTestRoleGranted // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestRoleGranted)
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
		it.Event = new(GatewayZEVMUpgradeTestRoleGranted)
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
func (it *GatewayZEVMUpgradeTestRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestRoleGranted represents a RoleGranted event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayZEVMUpgradeTestRoleGrantedIterator, error) {

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

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestRoleGrantedIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestRoleGranted)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseRoleGranted(log types.Log) (*GatewayZEVMUpgradeTestRoleGranted, error) {
	event := new(GatewayZEVMUpgradeTestRoleGranted)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestRoleRevokedIterator struct {
	Event *GatewayZEVMUpgradeTestRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestRoleRevoked)
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
		it.Event = new(GatewayZEVMUpgradeTestRoleRevoked)
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
func (it *GatewayZEVMUpgradeTestRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestRoleRevoked represents a RoleRevoked event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayZEVMUpgradeTestRoleRevokedIterator, error) {

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

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestRoleRevokedIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestRoleRevoked)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseRoleRevoked(log types.Log) (*GatewayZEVMUpgradeTestRoleRevoked, error) {
	event := new(GatewayZEVMUpgradeTestRoleRevoked)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestUnpausedIterator struct {
	Event *GatewayZEVMUpgradeTestUnpaused // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestUnpaused)
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
		it.Event = new(GatewayZEVMUpgradeTestUnpaused)
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
func (it *GatewayZEVMUpgradeTestUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestUnpaused represents a Unpaused event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GatewayZEVMUpgradeTestUnpausedIterator, error) {

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestUnpausedIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestUnpaused) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestUnpaused)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseUnpaused(log types.Log) (*GatewayZEVMUpgradeTestUnpaused, error) {
	event := new(GatewayZEVMUpgradeTestUnpaused)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestUpgradedIterator struct {
	Event *GatewayZEVMUpgradeTestUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestUpgraded)
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
		it.Event = new(GatewayZEVMUpgradeTestUpgraded)
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
func (it *GatewayZEVMUpgradeTestUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestUpgraded represents a Upgraded event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayZEVMUpgradeTestUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestUpgradedIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestUpgraded)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseUpgraded(log types.Log) (*GatewayZEVMUpgradeTestUpgraded, error) {
	event := new(GatewayZEVMUpgradeTestUpgraded)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestWithdrawnIterator struct {
	Event *GatewayZEVMUpgradeTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestWithdrawn)
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
		it.Event = new(GatewayZEVMUpgradeTestWithdrawn)
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
func (it *GatewayZEVMUpgradeTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestWithdrawn represents a Withdrawn event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestWithdrawn struct {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMUpgradeTestWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestWithdrawnIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestWithdrawn)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseWithdrawn(log types.Log) (*GatewayZEVMUpgradeTestWithdrawn, error) {
	event := new(GatewayZEVMUpgradeTestWithdrawn)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestWithdrawnAndCalledIterator struct {
	Event *GatewayZEVMUpgradeTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestWithdrawnAndCalled)
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
		it.Event = new(GatewayZEVMUpgradeTestWithdrawnAndCalled)
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
func (it *GatewayZEVMUpgradeTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestWithdrawnAndCalled struct {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMUpgradeTestWithdrawnAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestWithdrawnAndCalledIterator{contract: _GatewayZEVMUpgradeTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestWithdrawnAndCalled, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestWithdrawnAndCalled)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayZEVMUpgradeTestWithdrawnAndCalled, error) {
	event := new(GatewayZEVMUpgradeTestWithdrawnAndCalled)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMUpgradeTestWithdrawnV2Iterator is returned from FilterWithdrawnV2 and is used to iterate over the raw logs and unpacked data for WithdrawnV2 events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestWithdrawnV2Iterator struct {
	Event *GatewayZEVMUpgradeTestWithdrawnV2 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestWithdrawnV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestWithdrawnV2)
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
		it.Event = new(GatewayZEVMUpgradeTestWithdrawnV2)
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
func (it *GatewayZEVMUpgradeTestWithdrawnV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestWithdrawnV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestWithdrawnV2 represents a WithdrawnV2 event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestWithdrawnV2 struct {
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

// FilterWithdrawnV2 is a free log retrieval operation binding the contract event 0x5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d97.
//
// Solidity: event WithdrawnV2(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterWithdrawnV2(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMUpgradeTestWithdrawnV2Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "WithdrawnV2", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestWithdrawnV2Iterator{contract: _GatewayZEVMUpgradeTest.contract, event: "WithdrawnV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnV2 is a free log subscription operation binding the contract event 0x5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d97.
//
// Solidity: event WithdrawnV2(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchWithdrawnV2(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestWithdrawnV2, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "WithdrawnV2", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestWithdrawnV2)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
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

// ParseWithdrawnV2 is a log parse operation binding the contract event 0x5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d97.
//
// Solidity: event WithdrawnV2(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseWithdrawnV2(log types.Log) (*GatewayZEVMUpgradeTestWithdrawnV2, error) {
	event := new(GatewayZEVMUpgradeTestWithdrawnV2)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
