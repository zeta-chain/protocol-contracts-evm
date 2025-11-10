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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_MESSAGE_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROTOCOL_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndRevert\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalledV2\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"version\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a060405230608052348015610013575f80fd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516145516100f95f395f8181612d1101528181612d3a0152612f0f01526145515ff3fe6080604052600436106101f8575f3560e01c80634f1ef2861161011757806397d340f5116100ac578063bcf7f32b1161007c578063d547741f11610062578063d547741f146106d0578063e63ab1e9146106ef578063f45346dc14610722575f80fd5b8063bcf7f32b14610692578063c39aca37146106b1575f80fd5b806397d340f5146105f65780639d4ba4651461060b578063a217fddf1461062a578063ad3cb1cc1461063d575f80fd5b80637c0dcb5f116100e75780637c0dcb5f146105415780638456cb591461056057806391d148541461057457806397a1cef1146105d7575f80fd5b80634f1ef286146104c557806352d1902d146104d85780635c975abb146104ec5780637b15118b14610522575f80fd5b80632722feee1161018d57806338da73201161015d57806338da7320146104545780633b283933146104735780633f4ba83a14610492578063485cc955146104a6575f80fd5b80632722feee146103d05780632810ae63146103f75780632f2ff15d1461041657806336568abe14610435575f80fd5b80631cb5ea75116101c85780631cb5ea751461030157806321501a951461032057806321e093b11461033f578063248a9ca314610375575f80fd5b806301ffc9a714610270578063048ae42c146102a457806306cb8983146102c3578063184b0793146102e2575f80fd5b3661026c57610205610741565b5f546001600160a01b0316331480159061023357503373735b14bb79463307aacbed86daf3322b1e6226ab14155b1561026a576040517fb3af013700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b5f80fd5b34801561027b575f80fd5b5061028f61028a36600461344e565b61079f565b60405190151581526020015b60405180910390f35b3480156102af575f80fd5b5061026a6102be3660046135d2565b610837565b3480156102ce575f80fd5b5061026a6102dd36600461369a565b610a30565b3480156102ed575f80fd5b5061026a6102fc366004613761565b610b22565b34801561030c575f80fd5b5061026a61031b3660046137ae565b610c0c565b34801561032b575f80fd5b5061026a61033a366004613853565b610cd1565b34801561034a575f80fd5b505f5461035d906001600160a01b031681565b6040516001600160a01b03909116815260200161029b565b348015610380575f80fd5b506103c261038f3660046138d9565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b60405190815260200161029b565b3480156103db575f80fd5b5061035d73735b14bb79463307aacbed86daf3322b1e6226ab81565b348015610402575f80fd5b5061026a6104113660046138f0565b610e7e565b348015610421575f80fd5b5061026a61043036600461398d565b611012565b348015610440575f80fd5b5061026a61044f36600461398d565b61105b565b34801561045f575f80fd5b5061026a61046e3660046139bb565b6110ac565b34801561047e575f80fd5b5061026a61048d366004613a85565b611268565b34801561049d575f80fd5b5061026a6113d6565b3480156104b1575f80fd5b5061026a6104c0366004613b11565b61140b565b61026a6104d3366004613b3d565b611643565b3480156104e3575f80fd5b506103c2611662565b3480156104f7575f80fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1661028f565b34801561052d575f80fd5b5061026a61053c366004613b80565b611690565b34801561054c575f80fd5b5061026a61055b366004613beb565b61183f565b34801561056b575f80fd5b5061026a611a4b565b34801561057f575f80fd5b5061028f61058e36600461398d565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156105e2575f80fd5b5061026a6105f1366004613c6b565b611a7d565b348015610601575f80fd5b506103c261040081565b348015610616575f80fd5b5061026a610625366004613cca565b611bde565b348015610635575f80fd5b506103c25f81565b348015610648575f80fd5b506106856040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161029b9190613d72565b34801561069d575f80fd5b5061026a6106ac366004613d84565b611e48565b3480156106bc575f80fd5b5061026a6106cb366004613d84565b611f5a565b3480156106db575f80fd5b5061026a6106ea36600461398d565b61214c565b3480156106fa575f80fd5b506103c27f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b34801561072d575f80fd5b5061026a61073c366004613e1c565b61218f565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff161561079d576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061083157507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b61083f612381565b610847610741565b86515f036108685760405163d92e233d60e01b815260040160405180910390fd5b855f036108a1576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f036108da576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104006108ea6060830183613e5b565b6108f5915085613ebc565b111561092d576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610939878785612402565b90505f336001600160a01b03167f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c8a898b868c6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109a7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109cb9190613ef4565b6040805180820182528c81526001602082015290516109f59695949392918f918f91908e9061403f565b60405180910390a350610a2760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50505050505050565b610a38612381565b610a40610741565b81355f03610a7a576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610400610a8a6060830183613e5b565b610a95915085613ebc565b1115610acd576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610af186868686610ae3368890038801886140c0565b610aec87614116565b6126f8565b610b1a60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610b6f576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b77610741565b6001600160a01b038216610b9e5760405163d92e233d60e01b815260040160405180910390fd5b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063c9028a3690610be39084906004016141bb565b5f604051808303815f87803b158015610bfa575f80fd5b505af1158015610b1a573d5f803e3d5ffd5b610c14612381565b610c1c610741565b815f03610c55576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610400610c656060830183613e5b565b610c70915085613ebc565b1115610ca8576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610af18686868660405180604001604052808881526020016001151581525086610aec90614116565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610d1e576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d26610741565b6001600160a01b038316610d4d5760405163d92e233d60e01b815260040160405180910390fd5b835f03610d86576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831673735b14bb79463307aacbed86daf3322b1e6226ab1480610db957506001600160a01b03831630145b15610df0576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610dfa84846128d8565b5f546040517f5bcfd6160000000000000000000000000000000000000000000000000000000081526001600160a01b0380861692635bcfd61692610e4a928a921690899088908890600401614229565b5f604051808303815f87803b158015610e61575f80fd5b505af1158015610e73573d5f803e3d5ffd5b505050505050505050565b610e86612381565b610e8e610741565b86515f03610eaf5760405163d92e233d60e01b815260040160405180910390fd5b855f03610ee8576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81355f03610f22576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610400610f326060830183613e5b565b610f3d915085613ebc565b1115610f75576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f938673735b14bb79463307aacbed86daf3322b1e6226ab6128d8565b5f8054604051879233927f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c92610fe1928d926001600160a01b0316918d919081908d908d908d908d906142c4565b60405180910390a3610a2760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461104b81612a98565b6110558383612aa2565b50505050565b6001600160a01b038116331461109d576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110a78282612b8c565b505050565b6110b4610741565b87515f036110d55760405163d92e233d60e01b815260040160405180910390fd5b865f0361110e576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81355f03611148576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104006111586060830183613e5b565b611163915086613ebc565b111561119b576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6111a888888535612402565b90505f336001600160a01b03167f2e65c177543a837db5898e269ed0185a1f88d5cea587de5c980d481591751d9c8b8a8c868d6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611216573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061123a9190613ef4565b8d8d8d8d8d6040516112559a9998979695949392919061431b565b60405180910390a3505050505050505050565b611270612381565b611278610741565b85515f036112995760405163d92e233d60e01b815260040160405180910390fd5b845f036112d2576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104006112e26060830183613e5b565b6112ed915084613ebc565b1115611325576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113438573735b14bb79463307aacbed86daf3322b1e6226ab6128d8565b5f8054604080518082018252838152600160208201529051879333937f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c936113a5938d936001600160a01b03909316928d92909182918d918d91908d9061403f565b60405180910390a3610b1a60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61140081612a98565b611408612c4e565b50565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156114555750825b90505f8267ffffffffffffffff1660011480156114715750303b155b90508115801561147f575080155b156114b6576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156115175784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038716158061153457506001600160a01b038616155b156115525760405163d92e233d60e01b815260040160405180910390fd5b61155a612cde565b611562612cde565b61156a612ce6565b611572612cf6565b61157c5f87612aa2565b506115a77f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a87612aa2565b505f80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0389161790558315610a275784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050505050565b61164b612d06565b61165482612dd6565b61165e8282612de0565b5050565b5f61166b612f04565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b611698612381565b6116a0610741565b86515f036116c15760405163d92e233d60e01b815260040160405180910390fd5b855f036116fa576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81355f03611734576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104006117446060830183613e5b565b61174f915085613ebc565b1115611787576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61179487878535612402565b90505f336001600160a01b03167f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c8a898b868c6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611802573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118269190613ef4565b8c8c8c8c6040516109f5999897969594939291906142c4565b611847612381565b61184f610741565b83515f036118705760405163d92e233d60e01b815260040160405180910390fd5b825f036118a9576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104006118b96060830183613e5b565b905011156118f3576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6118fe8484612f66565b90505f336001600160a01b03167f5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d9787868886896001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561196c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119909190613ef4565b60405180604001604052808c6001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119d7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119fb9190613ef4565b81526001602090910152604051611a19969594939291908c9061439c565b60405180910390a35061105560017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a611a7581612a98565b611408612fd1565b611a85612381565b611a8d610741565b83515f03611aae5760405163d92e233d60e01b815260040160405180910390fd5b825f03611ae7576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610400611af76060830183613e5b565b90501115611b31576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611b4f8373735b14bb79463307aacbed86daf3322b1e6226ab6128d8565b5f8054604080518082018252838152600160208201529051859333937f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c93611bad938b936001600160a01b03909316928b9290918291908b9061439c565b60405180910390a361105560017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611c2b576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c33610741565b6001600160a01b0384161580611c5057506001600160a01b038216155b15611c6e5760405163d92e233d60e01b815260040160405180910390fd5b825f03611ca7576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03821673735b14bb79463307aacbed86daf3322b1e6226ab1480611cda57506001600160a01b03821630145b15611d11576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152602482018590528516906347e7ef24906044016020604051808303815f875af1158015611d76573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d9a919061441c565b611dd0576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063c9028a3690611e159084906004016141bb565b5f604051808303815f87803b158015611e2c575f80fd5b505af1158015611e3e573d5f803e3d5ffd5b5050505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611e95576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611e9d610741565b6001600160a01b0385161580611eba57506001600160a01b038316155b15611ed85760405163d92e233d60e01b815260040160405180910390fd5b6040517f5bcfd6160000000000000000000000000000000000000000000000000000000081526001600160a01b03841690635bcfd61690611f259089908990899088908890600401614229565b5f604051808303815f87803b158015611f3c575f80fd5b505af1158015611f4e573d5f803e3d5ffd5b50505050505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611fa7576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611faf610741565b6001600160a01b0385161580611fcc57506001600160a01b038316155b15611fea5760405163d92e233d60e01b815260040160405180910390fd5b835f03612023576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831673735b14bb79463307aacbed86daf3322b1e6226ab148061205657506001600160a01b03831630145b1561208d576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038481166004830152602482018690528616906347e7ef24906044016020604051808303815f875af11580156120f2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612116919061441c565b611ed8576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461218581612a98565b6110558383612b8c565b3373735b14bb79463307aacbed86daf3322b1e6226ab146121dc576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6121e4610741565b6001600160a01b038316158061220157506001600160a01b038116155b1561221f5760405163d92e233d60e01b815260040160405180910390fd5b815f03612258576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03811673735b14bb79463307aacbed86daf3322b1e6226ab148061228b57506001600160a01b03811630145b156122c2576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038281166004830152602482018490528416906347e7ef24906044016020604051808303815f875af1158015612327573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061234b919061441c565b6110a7576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016123fc576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b5f805f846001600160a01b031663fc5fecd5856040518263ffffffff1660e01b815260040161243391815260200190565b6040805180830381865afa15801561244d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124719190614437565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab60248201526044810182905291935091506001600160a01b038316906323b872dd906064016020604051808303815f875af11580156124f3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612517919061441c565b61254d576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018790526001600160a01b038616906323b872dd906064016020604051808303815f875af11580156125b6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125da919061441c565b612610576040517f4dd9ee8d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f42966c68000000000000000000000000000000000000000000000000000000008152600481018790526001600160a01b038616906342966c68906024016020604051808303815f875af115801561266d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612691919061441c565b6126c7576040517f2c77e05c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9150505b9392505050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b85515f036127195760405163d92e233d60e01b815260040160405180910390fd5b81516040517ffc5fecd500000000000000000000000000000000000000000000000000000000815260048101919091525f9081906001600160a01b0388169063fc5fecd5906024016040805180830381865afa15801561277b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061279f9190614437565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab60248201526044810182905291935091506001600160a01b038316906323b872dd906064016020604051808303815f875af1158015612821573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612845919061441c565b61287b576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b866001600160a01b0316336001600160a01b03167f306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e48a898989896040516128c6959493929190614463565b60405180910390a35050505050505050565b5f546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490526001600160a01b03909116906323b872dd906064016020604051808303815f875af1158015612944573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612968919061441c565b61299e576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f546040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b0390911690632e1a7d4d906024015f604051808303815f87803b1580156129f9575f80fd5b505af1158015612a0b573d5f803e3d5ffd5b505050505f816001600160a01b0316836040515f6040518083038185875af1925050503d805f8114612a58576040519150601f19603f3d011682016040523d82523d5f602084013e612a5d565b606091505b50509050806110a7576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611408813361304a565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612b83575f848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055612b393390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610831565b5f915050610831565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612b83575f848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610831565b612c566130d6565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b61079d613131565b612cee613131565b61079d613198565b612cfe613131565b61079d6131e9565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480612d9f57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612d937f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b1561079d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61165e81612a98565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612e58575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252612e5591810190613ef4565b60015b612e9e576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612efa576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401612e95565b6110a783836131f1565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461079d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6126cb8383846001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa158015612fa8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612fcc9190613ef4565b612402565b612fd9610741565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612cc0565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff1661165e576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401612e95565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1661079d576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661079d576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6131a0613131565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6126d2613131565b6131fa82613246565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561323e576110a782826132ed565b61165e61335f565b806001600160a01b03163b5f03613294576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401612e95565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f80846001600160a01b0316846040516133099190614505565b5f60405180830381855af49150503d805f8114613341576040519150601f19603f3d011682016040523d82523d5f602084013e613346565b606091505b5091509150613356858383613397565b95945050505050565b341561079d576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060826133ac576133a78261340c565b6126cb565b81511580156133c357506001600160a01b0384163b155b15613405576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401612e95565b50806126cb565b80511561341c5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6020828403121561345e575f80fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146126cb575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f8301126134c9575f80fd5b813567ffffffffffffffff8111156134e3576134e361348d565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810167ffffffffffffffff811182821017156135305761353061348d565b604052818152838201602001851015613547575f80fd5b816020850160208301375f918101602001919091529392505050565b6001600160a01b0381168114611408575f80fd5b5f8083601f840112613587575f80fd5b50813567ffffffffffffffff81111561359e575f80fd5b6020830191508360208285010111156135b5575f80fd5b9250929050565b5f60a082840312156135cc575f80fd5b50919050565b5f805f805f805f60c0888a0312156135e8575f80fd5b873567ffffffffffffffff8111156135fe575f80fd5b61360a8a828b016134ba565b97505060208801359550604088013561362281613563565b9450606088013567ffffffffffffffff81111561363d575f80fd5b6136498a828b01613577565b9095509350506080880135915060a088013567ffffffffffffffff81111561366f575f80fd5b61367b8a828b016135bc565b91505092959891949750929550565b5f604082840312156135cc575f80fd5b5f805f805f8060c087890312156136af575f80fd5b863567ffffffffffffffff8111156136c5575f80fd5b6136d189828a016134ba565b96505060208701356136e281613563565b9450604087013567ffffffffffffffff8111156136fd575f80fd5b61370989828a01613577565b909550935061371d9050886060890161368a565b915060a087013567ffffffffffffffff811115613738575f80fd5b61374489828a016135bc565b9150509295509295509295565b5f608082840312156135cc575f80fd5b5f8060408385031215613772575f80fd5b823561377d81613563565b9150602083013567ffffffffffffffff811115613798575f80fd5b6137a485828601613751565b9150509250929050565b5f805f805f8060a087890312156137c3575f80fd5b863567ffffffffffffffff8111156137d9575f80fd5b6137e589828a016134ba565b96505060208701356137f681613563565b9450604087013567ffffffffffffffff811115613811575f80fd5b61381d89828a01613577565b90955093505060608701359150608087013567ffffffffffffffff811115613738575f80fd5b5f606082840312156135cc575f80fd5b5f805f805f60808688031215613867575f80fd5b853567ffffffffffffffff81111561387d575f80fd5b61388988828901613843565b9550506020860135935060408601356138a181613563565b9250606086013567ffffffffffffffff8111156138bc575f80fd5b6138c888828901613577565b969995985093965092949392505050565b5f602082840312156138e9575f80fd5b5035919050565b5f805f805f805f60e0888a031215613906575f80fd5b873567ffffffffffffffff81111561391c575f80fd5b6139288a828b016134ba565b9750506020880135955060408801359450606088013567ffffffffffffffff811115613952575f80fd5b61395e8a828b01613577565b909550935061397290508960808a0161368a565b915060c088013567ffffffffffffffff81111561366f575f80fd5b5f806040838503121561399e575f80fd5b8235915060208301356139b081613563565b809150509250929050565b5f805f805f805f80610100898b0312156139d3575f80fd5b883567ffffffffffffffff8111156139e9575f80fd5b6139f58b828c016134ba565b985050602089013596506040890135613a0d81613563565b9550606089013567ffffffffffffffff811115613a28575f80fd5b613a348b828c01613577565b90965094505060808901359250613a4e8a60a08b0161368a565b915060e089013567ffffffffffffffff811115613a69575f80fd5b613a758b828c016135bc565b9150509295985092959890939650565b5f805f805f8060a08789031215613a9a575f80fd5b863567ffffffffffffffff811115613ab0575f80fd5b613abc89828a016134ba565b9650506020870135945060408701359350606087013567ffffffffffffffff811115613ae6575f80fd5b613af289828a01613577565b909450925050608087013567ffffffffffffffff811115613738575f80fd5b5f8060408385031215613b22575f80fd5b8235613b2d81613563565b915060208301356139b081613563565b5f8060408385031215613b4e575f80fd5b8235613b5981613563565b9150602083013567ffffffffffffffff811115613b74575f80fd5b6137a4858286016134ba565b5f805f805f805f60e0888a031215613b96575f80fd5b873567ffffffffffffffff811115613bac575f80fd5b613bb88a828b016134ba565b975050602088013595506040880135613bd081613563565b9450606088013567ffffffffffffffff811115613952575f80fd5b5f805f8060808587031215613bfe575f80fd5b843567ffffffffffffffff811115613c14575f80fd5b613c20878288016134ba565b945050602085013592506040850135613c3881613563565b9150606085013567ffffffffffffffff811115613c53575f80fd5b613c5f878288016135bc565b91505092959194509250565b5f805f8060808587031215613c7e575f80fd5b843567ffffffffffffffff811115613c94575f80fd5b613ca0878288016134ba565b9450506020850135925060408501359150606085013567ffffffffffffffff811115613c53575f80fd5b5f805f8060808587031215613cdd575f80fd5b8435613ce881613563565b9350602085013592506040850135613cff81613563565b9150606085013567ffffffffffffffff811115613d1a575f80fd5b613c5f87828801613751565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6126cb6020830184613d26565b5f805f805f8060a08789031215613d99575f80fd5b863567ffffffffffffffff811115613daf575f80fd5b613dbb89828a01613843565b9650506020870135613dcc81613563565b9450604087013593506060870135613de381613563565b9250608087013567ffffffffffffffff811115613dfe575f80fd5b613e0a89828a01613577565b979a9699509497509295939492505050565b5f805f60608486031215613e2e575f80fd5b8335613e3981613563565b9250602084013591506040840135613e5081613563565b809150509250925092565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613e8e575f80fd5b83018035915067ffffffffffffffff821115613ea8575f80fd5b6020019150368190038213156135b5575f80fd5b80820180821115610831577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f60208284031215613f04575f80fd5b5051919050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8015158114611408575f80fd5b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613f92575f80fd5b830160208101925035905067ffffffffffffffff811115613fb1575f80fd5b8036038213156135b5575f80fd5b5f8135613fcb81613563565b6001600160a01b031683526020820135613fe481613f52565b151560208401526040820135613ff981613563565b6001600160a01b031660408401526140146060830183613f5f565b60a0606086015261402960a086018284613f0b565b6080948501359590940194909452509092915050565b61012081525f61405361012083018c613d26565b6001600160a01b038b16602084015289604084015288606084015287608084015282810360a0840152614087818789613f0b565b855160c08501526020860151151560e085015290505b8281036101008401526140b08185613fbf565b9c9b505050505050505050505050565b5f60408284031280156140d1575f80fd5b506040805190810167ffffffffffffffff811182821017156140f5576140f561348d565b60405282358152602083013561410a81613f52565b60208201529392505050565b5f60a08236031215614126575f80fd5b60405160a0810167ffffffffffffffff811182821017156141495761414961348d565b604052823561415781613563565b8152602083013561416781613f52565b6020820152604083013561417a81613563565b6040820152606083013567ffffffffffffffff811115614198575f80fd5b6141a4368286016134ba565b606083015250608092830135928101929092525090565b602081525f82356141cb81613563565b6001600160a01b03811660208401525060208301356141e981613563565b6001600160a01b0381166040840152505f60408401359050806060840152506142156060840184613f5f565b60808085015261335660a085018284613f0b565b608081525f6142388788613f5f565b6060608085015261424d60e085018284613f0b565b915050602088013561425e81613563565b6001600160a01b0390811660a085015260408981013560c0860152908816602085015283018690528281036060840152614299818587613f0b565b98975050505050505050565b8035825260208101356142b781613f52565b8015156020840152505050565b61012081525f6142d861012083018c613d26565b6001600160a01b038b16602084015289604084015288606084015287608084015282810360a084015261430c818789613f0b565b905061409d60c08401866142a5565b61014081525f61432f61014083018d613d26565b6001600160a01b038c1660208401528a604084015289606084015288608084015282810360a084015261436381888a613f0b565b90508560c084015261437860e08401866142a5565b82810361012084015261438b8185613fbf565b9d9c50505050505050505050505050565b61012081525f6143b061012083018a613d26565b6001600160a01b03891660208401528760408401528660608401528560808401528281038060a08501525f82526143f660c0850187805182526020908101511515910152565b602081016101008501525061440e6020820185613fbf565b9a9950505050505050505050565b5f6020828403121561442c575f80fd5b81516126cb81613f52565b5f8060408385031215614448575f80fd5b825161445381613563565b6020939093015192949293505050565b60a081525f61447560a0830188613d26565b8281036020840152614488818789613f0b565b85516040850152602086015115156060850152905082810360808401526001600160a01b0384511681526020840151151560208201526001600160a01b036040850151166040820152606084015160a060608301526144ea60a0830182613d26565b90506080850151608083015280925050509695505050505050565b5f82518060208501845e5f92019182525091905056fea2646970667358221220638792f33c7cea673ab585d6c89bc2a575545acc727c226b00fbd6e2d837963664736f6c634300081a0033",
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

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCaller) MAXMESSAGESIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayZEVMUpgradeTest.contract.Call(opts, &out, "MAX_MESSAGE_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) MAXMESSAGESIZE() (*big.Int, error) {
	return _GatewayZEVMUpgradeTest.Contract.MAXMESSAGESIZE(&_GatewayZEVMUpgradeTest.CallOpts)
}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestCallerSession) MAXMESSAGESIZE() (*big.Int, error) {
	return _GatewayZEVMUpgradeTest.Contract.MAXMESSAGESIZE(&_GatewayZEVMUpgradeTest.CallOpts)
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

// Call0 is a paid mutator transaction binding the contract method 0x1cb5ea75.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Call0(opts *bind.TransactOpts, receiver []byte, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "call0", receiver, zrc20, message, gasLimit, revertOptions)
}

// Call0 is a paid mutator transaction binding the contract method 0x1cb5ea75.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Call0(receiver []byte, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Call0(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, zrc20, message, gasLimit, revertOptions)
}

// Call0 is a paid mutator transaction binding the contract method 0x1cb5ea75.
//
// Solidity: function call(bytes receiver, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Call0(receiver []byte, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Call0(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, zrc20, message, gasLimit, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Deposit(opts *bind.TransactOpts, zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "deposit", zrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Deposit(zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Deposit(&_GatewayZEVMUpgradeTest.TransactOpts, zrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address zrc20, uint256 amount, address target) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Deposit(zrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Deposit(&_GatewayZEVMUpgradeTest.TransactOpts, zrc20, amount, target)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) DepositAndCall(opts *bind.TransactOpts, context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "depositAndCall", context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) DepositAndCall(context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) DepositAndCall(context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.DepositAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, context, amount, target, message)
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

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Withdraw(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdraw", receiver, amount, zrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Withdraw(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address zrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Withdraw(receiver []byte, amount *big.Int, zrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Withdraw(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) Withdraw0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdraw0", receiver, amount, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) Withdraw0(receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Withdraw0(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) Withdraw0(receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.Withdraw0(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, chainId, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x048ae42c.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) WithdrawAndCall(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall", receiver, amount, zrc20, message, gasLimit, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x048ae42c.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) WithdrawAndCall(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, message, gasLimit, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x048ae42c.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) WithdrawAndCall(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, message, gasLimit, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x2810ae63.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) WithdrawAndCall0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall0", receiver, amount, chainId, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x2810ae63.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) WithdrawAndCall0(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall0(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, chainId, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x2810ae63.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) WithdrawAndCall0(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall0(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, chainId, message, callOptions, revertOptions)
}

// WithdrawAndCall1 is a paid mutator transaction binding the contract method 0x38da7320.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, uint256 version, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) WithdrawAndCall1(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, version *big.Int, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall1", receiver, amount, zrc20, message, version, callOptions, revertOptions)
}

// WithdrawAndCall1 is a paid mutator transaction binding the contract method 0x38da7320.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, uint256 version, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) WithdrawAndCall1(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, version *big.Int, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall1(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, message, version, callOptions, revertOptions)
}

// WithdrawAndCall1 is a paid mutator transaction binding the contract method 0x38da7320.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, uint256 version, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) WithdrawAndCall1(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, version *big.Int, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall1(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, message, version, callOptions, revertOptions)
}

// WithdrawAndCall2 is a paid mutator transaction binding the contract method 0x3b283933.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) WithdrawAndCall2(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall2", receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall2 is a paid mutator transaction binding the contract method 0x3b283933.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) WithdrawAndCall2(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall2(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall2 is a paid mutator transaction binding the contract method 0x3b283933.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) WithdrawAndCall2(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall2(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall3 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) WithdrawAndCall3(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall3", receiver, amount, zrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall3 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) WithdrawAndCall3(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall3(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall3 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) WithdrawAndCall3(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall3(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, message, callOptions, revertOptions)
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

// GatewayZEVMUpgradeTestWithdrawnAndCalledV2Iterator is returned from FilterWithdrawnAndCalledV2 and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalledV2 events raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestWithdrawnAndCalledV2Iterator struct {
	Event *GatewayZEVMUpgradeTestWithdrawnAndCalledV2 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMUpgradeTestWithdrawnAndCalledV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMUpgradeTestWithdrawnAndCalledV2)
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
		it.Event = new(GatewayZEVMUpgradeTestWithdrawnAndCalledV2)
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
func (it *GatewayZEVMUpgradeTestWithdrawnAndCalledV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMUpgradeTestWithdrawnAndCalledV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMUpgradeTestWithdrawnAndCalledV2 represents a WithdrawnAndCalledV2 event raised by the GatewayZEVMUpgradeTest contract.
type GatewayZEVMUpgradeTestWithdrawnAndCalledV2 struct {
	Sender          common.Address
	ChainId         *big.Int
	Receiver        []byte
	Zrc20           common.Address
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	Version         *big.Int
	CallOptions     CallOptions
	RevertOptions   RevertOptions
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalledV2 is a free log retrieval operation binding the contract event 0x2e65c177543a837db5898e269ed0185a1f88d5cea587de5c980d481591751d9c.
//
// Solidity: event WithdrawnAndCalledV2(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 version, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) FilterWithdrawnAndCalledV2(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMUpgradeTestWithdrawnAndCalledV2Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.FilterLogs(opts, "WithdrawnAndCalledV2", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMUpgradeTestWithdrawnAndCalledV2Iterator{contract: _GatewayZEVMUpgradeTest.contract, event: "WithdrawnAndCalledV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalledV2 is a free log subscription operation binding the contract event 0x2e65c177543a837db5898e269ed0185a1f88d5cea587de5c980d481591751d9c.
//
// Solidity: event WithdrawnAndCalledV2(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 version, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) WatchWithdrawnAndCalledV2(opts *bind.WatchOpts, sink chan<- *GatewayZEVMUpgradeTestWithdrawnAndCalledV2, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMUpgradeTest.contract.WatchLogs(opts, "WithdrawnAndCalledV2", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMUpgradeTestWithdrawnAndCalledV2)
				if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "WithdrawnAndCalledV2", log); err != nil {
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

// ParseWithdrawnAndCalledV2 is a log parse operation binding the contract event 0x2e65c177543a837db5898e269ed0185a1f88d5cea587de5c980d481591751d9c.
//
// Solidity: event WithdrawnAndCalledV2(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, uint256 version, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestFilterer) ParseWithdrawnAndCalledV2(log types.Log) (*GatewayZEVMUpgradeTestWithdrawnAndCalledV2, error) {
	event := new(GatewayZEVMUpgradeTestWithdrawnAndCalledV2)
	if err := _GatewayZEVMUpgradeTest.contract.UnpackLog(event, "WithdrawnAndCalledV2", log); err != nil {
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
