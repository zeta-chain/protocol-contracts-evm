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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_MESSAGE_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROTOCOL_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndRevert\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516143726100fd60003960008181612bbd01528181612be60152612dbc01526143726000f3fe6080604052600436106101e75760003560e01c806352d1902d116101025780639d4ba46511610095578063c39aca3711610064578063c39aca37146106a2578063d547741f146106c2578063e63ab1e9146106e2578063f45346dc1461071657600080fd5b80639d4ba465146105f7578063a217fddf14610617578063ad3cb1cc1461062c578063bcf7f32b1461068257600080fd5b80638456cb59116100d15780638456cb591461054757806391d148541461055c57806397a1cef1146105c157806397d340f5146105e157600080fd5b806352d1902d146104bb5780635c975abb146104d05780637b15118b146105075780637c0dcb5f1461052757600080fd5b80632722feee1161017a5780633b283933116101495780633b283933146104535780633f4ba83a14610473578063485cc955146104885780634f1ef286146104a857600080fd5b80632722feee146103cb5780632810ae63146103f35780632f2ff15d1461041357806336568abe1461043357600080fd5b80631cb5ea75116101b65780631cb5ea75146102f657806321501a951461031657806321e093b114610336578063248a9ca31461036e57600080fd5b806301ffc9a714610261578063048ae42c1461029657806306cb8983146102b6578063184b0793146102d657600080fd5b3661025c576101f4610736565b6000546001600160a01b0316331480159061022357503373735b14bb79463307aacbed86daf3322b1e6226ab14155b1561025a576040517fb3af013700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b34801561026d57600080fd5b5061028161027c366004613305565b610794565b60405190151581526020015b60405180910390f35b3480156102a257600080fd5b5061025a6102b1366004613499565b61082d565b3480156102c257600080fd5b5061025a6102d136600461356b565b610a2d565b3480156102e257600080fd5b5061025a6102f136600461363b565b610b20565b34801561030257600080fd5b5061025a61031136600461368b565b610c0f565b34801561032257600080fd5b5061025a610331366004613739565b610cd5565b34801561034257600080fd5b50600054610356906001600160a01b031681565b6040516001600160a01b03909116815260200161028d565b34801561037a57600080fd5b506103bd6103893660046137c5565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b60405190815260200161028d565b3480156103d757600080fd5b5061035673735b14bb79463307aacbed86daf3322b1e6226ab81565b3480156103ff57600080fd5b5061025a61040e3660046137de565b610e89565b34801561041f57600080fd5b5061025a61042e366004613883565b611021565b34801561043f57600080fd5b5061025a61044e366004613883565b61106b565b34801561045f57600080fd5b5061025a61046e3660046138b3565b6110bc565b34801561047f57600080fd5b5061025a61122d565b34801561049457600080fd5b5061025a6104a3366004613946565b611262565b61025a6104b6366004613974565b61149e565b3480156104c757600080fd5b506103bd6114bd565b3480156104dc57600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610281565b34801561051357600080fd5b5061025a6105223660046139ba565b6114ec565b34801561053357600080fd5b5061025a610542366004613a2c565b6116a2565b34801561055357600080fd5b5061025a6118b6565b34801561056857600080fd5b50610281610577366004613883565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156105cd57600080fd5b5061025a6105dc366004613ab1565b6118e8565b3480156105ed57600080fd5b506103bd61040081565b34801561060357600080fd5b5061025a610612366004613b15565b611a4c565b34801561062357600080fd5b506103bd600081565b34801561063857600080fd5b506106756040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161028d9190613be3565b34801561068e57600080fd5b5061025a61069d366004613bf6565b611cbf565b3480156106ae57600080fd5b5061025a6106bd366004613bf6565b611dd6565b3480156106ce57600080fd5b5061025a6106dd366004613883565b611fcc565b3480156106ee57600080fd5b506103bd7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b34801561072257600080fd5b5061025a610731366004613c94565b612010565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615610792576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061082757507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b610835612206565b61083d610736565b865160000361085f5760405163d92e233d60e01b815260040160405180910390fd5b85600003610899576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816000036108d3576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104006108e36060830183613cd6565b6108ee915085613d3b565b1115610926576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610933878785612287565b90506000336001600160a01b03167f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c8a898b868c6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109c89190613d75565b6040805180820182528c81526001602082015290516109f29695949392918f918f91908e90613eca565b60405180910390a350610a2460017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50505050505050565b610a35612206565b610a3d610736565b8135600003610a78576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610400610a886060830183613cd6565b610a93915085613d3b565b1115610acb576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610aef86868686610ae136889003880188613f4c565b610aea87613fa4565b61258a565b610b1860017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610b6d576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b75610736565b6001600160a01b038216610b9c5760405163d92e233d60e01b815260040160405180910390fd5b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063c9028a3690610be190849060040161404c565b600060405180830381600087803b158015610bfb57600080fd5b505af1158015610b18573d6000803e3d6000fd5b610c17612206565b610c1f610736565b81600003610c59576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610400610c696060830183613cd6565b610c74915085613d3b565b1115610cac576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610aef8686868660405180604001604052808881526020016001151581525086610aea90613fa4565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610d22576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d2a610736565b6001600160a01b038316610d515760405163d92e233d60e01b815260040160405180910390fd5b83600003610d8b576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831673735b14bb79463307aacbed86daf3322b1e6226ab1480610dbe57506001600160a01b03831630145b15610df5576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610dff8484612771565b6000546040517f5bcfd6160000000000000000000000000000000000000000000000000000000081526001600160a01b0380861692635bcfd61692610e50928a9216908990889088906004016140bc565b600060405180830381600087803b158015610e6a57600080fd5b505af1158015610e7e573d6000803e3d6000fd5b505050505050505050565b610e91612206565b610e99610736565b8651600003610ebb5760405163d92e233d60e01b815260040160405180910390fd5b85600003610ef5576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8135600003610f30576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610400610f406060830183613cd6565b610f4b915085613d3b565b1115610f83576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fa18673735b14bb79463307aacbed86daf3322b1e6226ab612771565b60008054604051879233927f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c92610ff0928d926001600160a01b0316918d919081908d908d908d908d90614158565b60405180910390a3610a2460017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461105b8161293f565b6110658383612949565b50505050565b6001600160a01b03811633146110ad576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110b78282612a36565b505050565b6110c4612206565b6110cc610736565b85516000036110ee5760405163d92e233d60e01b815260040160405180910390fd5b84600003611128576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104006111386060830183613cd6565b611143915084613d3b565b111561117b576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111998573735b14bb79463307aacbed86daf3322b1e6226ab612771565b60008054604080518082018252838152600160208201529051879333937f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c936111fc938d936001600160a01b03909316928d92909182918d918d91908d90613eca565b60405180910390a3610b1860017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6112578161293f565b61125f612afa565b50565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156112ad5750825b905060008267ffffffffffffffff1660011480156112ca5750303b155b9050811580156112d8575080155b1561130f576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156113705784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038716158061138d57506001600160a01b038616155b156113ab5760405163d92e233d60e01b815260040160405180910390fd5b6113b3612b8a565b6113bb612b8a565b6113c3612b92565b6113cb612ba2565b6113d6600087612949565b506114017f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a87612949565b50600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0389161790558315610a245784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050505050565b6114a6612bb2565b6114af82612c82565b6114b98282612c8d565b5050565b60006114c7612db1565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b6114f4612206565b6114fc610736565b865160000361151e5760405163d92e233d60e01b815260040160405180910390fd5b85600003611558576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8135600003611593576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104006115a36060830183613cd6565b6115ae915085613d3b565b11156115e6576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006115f487878535612287565b90506000336001600160a01b03167f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c8a898b868c6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611665573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116899190613d75565b8c8c8c8c6040516109f299989796959493929190614158565b6116aa612206565b6116b2610736565b83516000036116d45760405163d92e233d60e01b815260040160405180910390fd5b8260000361170e576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61040061171e6060830183613cd6565b90501115611758576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006117648484612e13565b90506000336001600160a01b03167f5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d9787868886896001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117f99190613d75565b60405180604001604052808c6001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa158015611842573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118669190613d75565b81526001602090910152604051611884969594939291908c906141b0565b60405180910390a35061106560017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6118e08161293f565b61125f612e81565b6118f0612206565b6118f8610736565b835160000361191a5760405163d92e233d60e01b815260040160405180910390fd5b82600003611954576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104006119646060830183613cd6565b9050111561199e576040517f9507fb3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119bc8373735b14bb79463307aacbed86daf3322b1e6226ab612771565b60008054604080518082018252838152600160208201529051859333937f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c93611a1b938b936001600160a01b03909316928b9290918291908b906141b0565b60405180910390a361106560017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611a99576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611aa1610736565b6001600160a01b0384161580611abe57506001600160a01b038216155b15611adc5760405163d92e233d60e01b815260040160405180910390fd5b82600003611b16576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03821673735b14bb79463307aacbed86daf3322b1e6226ab1480611b4957506001600160a01b03821630145b15611b80576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152602482018590528516906347e7ef24906044016020604051808303816000875af1158015611be8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c0c9190614232565b611c42576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063c9028a3690611c8790849060040161404c565b600060405180830381600087803b158015611ca157600080fd5b505af1158015611cb5573d6000803e3d6000fd5b5050505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611d0c576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611d14610736565b6001600160a01b0385161580611d3157506001600160a01b038316155b15611d4f5760405163d92e233d60e01b815260040160405180910390fd5b6040517f5bcfd6160000000000000000000000000000000000000000000000000000000081526001600160a01b03841690635bcfd61690611d9c90899089908990889088906004016140bc565b600060405180830381600087803b158015611db657600080fd5b505af1158015611dca573d6000803e3d6000fd5b50505050505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611e23576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611e2b610736565b6001600160a01b0385161580611e4857506001600160a01b038316155b15611e665760405163d92e233d60e01b815260040160405180910390fd5b83600003611ea0576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831673735b14bb79463307aacbed86daf3322b1e6226ab1480611ed357506001600160a01b03831630145b15611f0a576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038481166004830152602482018690528616906347e7ef24906044016020604051808303816000875af1158015611f72573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f969190614232565b611d4f576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546120068161293f565b6110658383612a36565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461205d576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612065610736565b6001600160a01b038316158061208257506001600160a01b038116155b156120a05760405163d92e233d60e01b815260040160405180910390fd5b816000036120da576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03811673735b14bb79463307aacbed86daf3322b1e6226ab148061210d57506001600160a01b03811630145b15612144576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038281166004830152602482018490528416906347e7ef24906044016020604051808303816000875af11580156121ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121d09190614232565b6110b7576040517f47d19fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01612281576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b6000806000846001600160a01b031663fc5fecd5856040518263ffffffff1660e01b81526004016122ba91815260200190565b6040805180830381865afa1580156122d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122fa919061424f565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab60248201526044810182905291935091506001600160a01b038316906323b872dd906064016020604051808303816000875af115801561237f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123a39190614232565b6123d9576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018790526001600160a01b038616906323b872dd906064016020604051808303816000875af1158015612445573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124699190614232565b61249f576040517f4dd9ee8d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f42966c68000000000000000000000000000000000000000000000000000000008152600481018790526001600160a01b038616906342966c68906024016020604051808303816000875af11580156124ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125239190614232565b612559576040517f2c77e05c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9150505b9392505050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b85516000036125ac5760405163d92e233d60e01b815260040160405180910390fd5b81516040517ffc5fecd5000000000000000000000000000000000000000000000000000000008152600481019190915260009081906001600160a01b0388169063fc5fecd5906024016040805180830381865afa158015612611573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612635919061424f565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab60248201526044810182905291935091506001600160a01b038316906323b872dd906064016020604051808303816000875af11580156126ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126de9190614232565b612714576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b866001600160a01b0316336001600160a01b03167f306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e48a8989898960405161275f95949392919061427d565b60405180910390a35050505050505050565b6000546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490526001600160a01b03909116906323b872dd906064016020604051808303816000875af11580156127e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128059190614232565b61283b576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000546040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b0390911690632e1a7d4d90602401600060405180830381600087803b15801561289a57600080fd5b505af11580156128ae573d6000803e3d6000fd5b505050506000816001600160a01b03168360405160006040518083038185875af1925050503d80600081146128ff576040519150601f19603f3d011682016040523d82523d6000602084013e612904565b606091505b50509050806110b7576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61125f8133612efa565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612a2c576000848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556129e23390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610827565b6000915050610827565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612a2c576000848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610827565b612b02612f87565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b610792612fe2565b612b9a612fe2565b610792613049565b612baa612fe2565b61079261309a565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480612c4b57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612c3f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15610792576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006114b98161293f565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612d05575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252612d0291810190613d75565b60015b612d4b576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612da7576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401612d42565b6110b783836130a2565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610792576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061255d8383846001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e7c9190613d75565b612287565b612e89610736565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612b6c565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff166114b9576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401612d42565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610792576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610792576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613051612fe2565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b612564612fe2565b6130ab826130f8565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156130f0576110b782826131a0565b6114b9613216565b806001600160a01b03163b600003613147576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401612d42565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516131bd9190614320565b600060405180830381855af49150503d80600081146131f8576040519150601f19603f3d011682016040523d82523d6000602084013e6131fd565b606091505b509150915061320d85838361324e565b95945050505050565b3415610792576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060826132635761325e826132c3565b61255d565b815115801561327a57506001600160a01b0384163b155b156132bc576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401612d42565b508061255d565b8051156132d35780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561331757600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461255d57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261338757600080fd5b813567ffffffffffffffff8111156133a1576133a1613347565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810167ffffffffffffffff811182821017156133ee576133ee613347565b60405281815283820160200185101561340657600080fd5b816020850160208301376000918101602001919091529392505050565b6001600160a01b038116811461125f57600080fd5b60008083601f84011261344a57600080fd5b50813567ffffffffffffffff81111561346257600080fd5b60208301915083602082850101111561347a57600080fd5b9250929050565b600060a0828403121561349357600080fd5b50919050565b600080600080600080600060c0888a0312156134b457600080fd5b873567ffffffffffffffff8111156134cb57600080fd5b6134d78a828b01613376565b9750506020880135955060408801356134ef81613423565b9450606088013567ffffffffffffffff81111561350b57600080fd5b6135178a828b01613438565b9095509350506080880135915060a088013567ffffffffffffffff81111561353e57600080fd5b61354a8a828b01613481565b91505092959891949750929550565b60006040828403121561349357600080fd5b60008060008060008060c0878903121561358457600080fd5b863567ffffffffffffffff81111561359b57600080fd5b6135a789828a01613376565b96505060208701356135b881613423565b9450604087013567ffffffffffffffff8111156135d457600080fd5b6135e089828a01613438565b90955093506135f490508860608901613559565b915060a087013567ffffffffffffffff81111561361057600080fd5b61361c89828a01613481565b9150509295509295509295565b60006080828403121561349357600080fd5b6000806040838503121561364e57600080fd5b823561365981613423565b9150602083013567ffffffffffffffff81111561367557600080fd5b61368185828601613629565b9150509250929050565b60008060008060008060a087890312156136a457600080fd5b863567ffffffffffffffff8111156136bb57600080fd5b6136c789828a01613376565b96505060208701356136d881613423565b9450604087013567ffffffffffffffff8111156136f457600080fd5b61370089828a01613438565b90955093505060608701359150608087013567ffffffffffffffff81111561361057600080fd5b60006060828403121561349357600080fd5b60008060008060006080868803121561375157600080fd5b853567ffffffffffffffff81111561376857600080fd5b61377488828901613727565b95505060208601359350604086013561378c81613423565b9250606086013567ffffffffffffffff8111156137a857600080fd5b6137b488828901613438565b969995985093965092949392505050565b6000602082840312156137d757600080fd5b5035919050565b600080600080600080600060e0888a0312156137f957600080fd5b873567ffffffffffffffff81111561381057600080fd5b61381c8a828b01613376565b9750506020880135955060408801359450606088013567ffffffffffffffff81111561384757600080fd5b6138538a828b01613438565b909550935061386790508960808a01613559565b915060c088013567ffffffffffffffff81111561353e57600080fd5b6000806040838503121561389657600080fd5b8235915060208301356138a881613423565b809150509250929050565b60008060008060008060a087890312156138cc57600080fd5b863567ffffffffffffffff8111156138e357600080fd5b6138ef89828a01613376565b9650506020870135945060408701359350606087013567ffffffffffffffff81111561391a57600080fd5b61392689828a01613438565b909450925050608087013567ffffffffffffffff81111561361057600080fd5b6000806040838503121561395957600080fd5b823561396481613423565b915060208301356138a881613423565b6000806040838503121561398757600080fd5b823561399281613423565b9150602083013567ffffffffffffffff8111156139ae57600080fd5b61368185828601613376565b600080600080600080600060e0888a0312156139d557600080fd5b873567ffffffffffffffff8111156139ec57600080fd5b6139f88a828b01613376565b975050602088013595506040880135613a1081613423565b9450606088013567ffffffffffffffff81111561384757600080fd5b60008060008060808587031215613a4257600080fd5b843567ffffffffffffffff811115613a5957600080fd5b613a6587828801613376565b945050602085013592506040850135613a7d81613423565b9150606085013567ffffffffffffffff811115613a9957600080fd5b613aa587828801613481565b91505092959194509250565b60008060008060808587031215613ac757600080fd5b843567ffffffffffffffff811115613ade57600080fd5b613aea87828801613376565b9450506020850135925060408501359150606085013567ffffffffffffffff811115613a9957600080fd5b60008060008060808587031215613b2b57600080fd5b8435613b3681613423565b9350602085013592506040850135613b4d81613423565b9150606085013567ffffffffffffffff811115613b6957600080fd5b613aa587828801613629565b60005b83811015613b90578181015183820152602001613b78565b50506000910152565b60008151808452613bb1816020860160208601613b75565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061255d6020830184613b99565b60008060008060008060a08789031215613c0f57600080fd5b863567ffffffffffffffff811115613c2657600080fd5b613c3289828a01613727565b9650506020870135613c4381613423565b9450604087013593506060870135613c5a81613423565b9250608087013567ffffffffffffffff811115613c7657600080fd5b613c8289828a01613438565b979a9699509497509295939492505050565b600080600060608486031215613ca957600080fd5b8335613cb481613423565b9250602084013591506040840135613ccb81613423565b809150509250925092565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613d0b57600080fd5b83018035915067ffffffffffffffff821115613d2657600080fd5b60200191503681900382131561347a57600080fd5b80820180821115610827577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600060208284031215613d8757600080fd5b5051919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b801515811461125f57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613e1a57600080fd5b830160208101925035905067ffffffffffffffff811115613e3a57600080fd5b80360382131561347a57600080fd5b60008135613e5681613423565b6001600160a01b031683526020820135613e6f81613dd7565b151560208401526040820135613e8481613423565b6001600160a01b03166040840152613e9f6060830183613de5565b60a06060860152613eb460a086018284613d8e565b6080948501359590940194909452509092915050565b61012081526000613edf61012083018c613b99565b6001600160a01b038b16602084015289604084015288606084015287608084015282810360a0840152613f13818789613d8e565b855160c08501526020860151151560e085015290505b828103610100840152613f3c8185613e49565b9c9b505050505050505050505050565b60006040828403128015613f5f57600080fd5b506040805190810167ffffffffffffffff81118282101715613f8357613f83613347565b604052823581526020830135613f9881613dd7565b60208201529392505050565b600060a08236031215613fb657600080fd5b60405160a0810167ffffffffffffffff81118282101715613fd957613fd9613347565b6040528235613fe781613423565b81526020830135613ff781613dd7565b6020820152604083013561400a81613423565b6040820152606083013567ffffffffffffffff81111561402957600080fd5b61403536828601613376565b606083015250608092830135928101929092525090565b602081526000823561405d81613423565b6001600160a01b038116602084015250602083013561407b81613423565b6001600160a01b038116604084015250600060408401359050806060840152506140a86060840184613de5565b60808085015261320d60a085018284613d8e565b6080815260006140cc8788613de5565b606060808501526140e160e085018284613d8e565b91505060208801356140f281613423565b6001600160a01b0390811660a085015260408981013560c086015290881660208501528301869052828103606084015261412d818587613d8e565b98975050505050505050565b80358252602081013561414b81613dd7565b8015156020840152505050565b6101208152600061416d61012083018c613b99565b6001600160a01b038b16602084015289604084015288606084015287608084015282810360a08401526141a1818789613d8e565b9050613f2960c0840186614139565b610120815260006141c561012083018a613b99565b6001600160a01b03891660208401528760408401528660608401528560808401528281038060a08501526000825261420c60c0850187805182526020908101511515910152565b60208101610100850152506142246020820185613e49565b9a9950505050505050505050565b60006020828403121561424457600080fd5b815161255d81613dd7565b6000806040838503121561426257600080fd5b825161426d81613423565b6020939093015192949293505050565b60a08152600061429060a0830188613b99565b82810360208401526142a3818789613d8e565b85516040850152602086015115156060850152905082810360808401526001600160a01b0384511681526020840151151560208201526001600160a01b036040850151166040820152606084015160a0606083015261430560a0830182613b99565b90506080850151608083015280925050509695505050505050565b60008251614332818460208701613b75565b919091019291505056fea2646970667358221220749248cf6694aaa069fbb6ebb425546a93e6128189aa1fc6d00f745fe826651f64736f6c634300081a0033",
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

// WithdrawAndCall1 is a paid mutator transaction binding the contract method 0x3b283933.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) WithdrawAndCall1(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall1", receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall1 is a paid mutator transaction binding the contract method 0x3b283933.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) WithdrawAndCall1(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall1(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall1 is a paid mutator transaction binding the contract method 0x3b283933.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) WithdrawAndCall1(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall1(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall2 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactor) WithdrawAndCall2(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall2", receiver, amount, zrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall2 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestSession) WithdrawAndCall2(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall2(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall2 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayZEVMUpgradeTest *GatewayZEVMUpgradeTestTransactorSession) WithdrawAndCall2(receiver []byte, amount *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayZEVMUpgradeTest.Contract.WithdrawAndCall2(&_GatewayZEVMUpgradeTest.TransactOpts, receiver, amount, zrc20, message, callOptions, revertOptions)
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
