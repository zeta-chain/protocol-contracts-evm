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

// StdInvariantFuzzArtifactSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzArtifactSelector struct {
	Artifact  string
	Selectors [][4]byte
}

// StdInvariantFuzzInterface is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzInterface struct {
	Addr      common.Address
	Artifacts []string
}

// StdInvariantFuzzSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzSelector struct {
	Addr      common.Address
	Selectors [][4]byte
}

// GatewayZEVMOutboundTestMetaData contains all meta data concerning the GatewayZEVMOutboundTest contract.
var GatewayZEVMOutboundTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForDifferentZRC20WithdrawAndCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForDifferentZRC20Withdrawal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForZRC20WithdrawAndCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnGasFeeForZRC20Withdrawal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testBurnProtocolFeesFailsWithInsufficientAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDeposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfSenderNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETA\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndCallUniversalContractFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAAndRevertUniversalContractWhenPaused\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfInsufficientProtocolBalance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZETAFailsWhenPaused\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractIfTargetIsGateway\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZRC20AndCallUniversalContractIfTargetIsProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteAbortUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteFailsIfZRC20IsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertUniversalContractIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContractFailsIfSenderIsNotProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteUniversalContractFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextData\",\"inputs\":[{\"name\":\"origin\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"msgSender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextDataAbort\",\"inputs\":[{\"name\":\"abortContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAbortContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outgoing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextDataRevert\",\"inputs\":[{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroGasPrice\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602b575f80fd5b5062018b12806200003b5f395ff3fe608060405234801561000f575f80fd5b50600436106104c3575f3560e01c80637cec29b011610277578063c35cb5e411610162578063eab7674e116100dd578063f1d98f1b11610093578063f589655d11610079578063f589655d14610770578063fa7626d414610778578063fb339a1c14610785575f80fd5b8063f1d98f1b14610760578063f2a9135814610768575f80fd5b8063ec294d9f116100c3578063ec294d9f14610748578063ee0f4ea114610750578063ef2b539414610758575f80fd5b8063eab7674e14610738578063eb78bd7d14610740575f80fd5b8063da8fd81d11610132578063e20c9f7111610118578063e20c9f71146106f3578063e36c0a23146106fb578063e63ab1e914610703575f80fd5b8063da8fd81d146106e3578063e09bc659146106eb575f80fd5b8063c35cb5e4146106c3578063c8814d2e146106cb578063ca26929c146106d3578063cf2c3d1d146106db575f80fd5b8063996b7675116101f2578063b5508aa9116101c2578063b84a3d2f116101a8578063b84a3d2f146106a3578063b936be8c14610629578063ba414fa6146106ab575f80fd5b8063b5508aa914610693578063b82ce9211461069b575f80fd5b8063996b7675146106735780639c9acd5d1461067b578063ae4662ec14610683578063b0464fdc1461068b575f80fd5b8063884660a311610247578063916a17c61161022d578063916a17c61461065657806396d9d8761461066b57806397f7661f14610629575f80fd5b8063884660a314610646578063890a2d671461064e575f80fd5b80637cec29b0146106195780637f924c4e14610621578063828d267c1461062957806385226c8114610631575f80fd5b806333f7c81a116103b1578063514fcc7f1161032c57806361cfddb7116102e257806366d9a9a0116102c857806366d9a9a0146105f45780636efa04b51461060957806375ca558714610611575f80fd5b806361cfddb7146105e457806362fe2915146105ec575f80fd5b80635b4c90e1116103125780635b4c90e11461053f5780635cec7db5146105d45780636163f8ef146105dc575f80fd5b8063514fcc7f146105c457806358c9987f146105cc575f80fd5b80633f7286f41161038157806348f4fd071161036757806348f4fd07146105ac57806350dd38ca146105b457806351336fb0146105bc575f80fd5b80633f7286f41461059c578063445bbf47146105a4575f80fd5b806333f7c81a1461057c5780633626c616146105845780633ab5b1991461058c5780633e5e3c2314610594575f80fd5b806321d92591116104415780632acb21b4116104115780632ed7cff6116103f75780632ed7cff6146105645780633177f3811461056c578063339bd82814610574575f80fd5b80632acb21b4146105475780632ade38801461054f575f80fd5b806321d92591146105275780632468bc0f1461052f57806327820625146105375780632948df411461053f575f80fd5b80631832cb6e116104965780631c785a141161047c5780631c785a14146104f95780631ed7831c146105015780631ff5bdd71461051f575f80fd5b80631832cb6e146104e9578063195ac143146104f1575f80fd5b8063084fafab146104c75780630a9254e4146104d15780630c4a14b4146104d957806314b7a6da146104e1575b5f80fd5b6104cf61078d565b005b6104cf610954565b6104cf611769565b6104cf61199e565b6104cf611aca565b6104cf611c61565b6104cf611dc7565b6105096126bf565b604051610516919061d905565b60405180910390f35b6104cf61271f565b6104cf61283d565b6104cf61299f565b6104cf612c98565b6104cf612d5b565b6104cf612f01565b6105576130b6565b604051610516919061d97e565b6104cf6131f2565b6104cf613315565b6104cf6139ce565b6104cf613af6565b6104cf613d02565b6104cf613e62565b610509613fef565b61050961404d565b6104cf6140ab565b6104cf6141f5565b6104cf61431c565b6104cf61445b565b6104cf61459c565b6104cf6146d8565b6104cf614882565b6104cf614a58565b6104cf614b9c565b6104cf6150a0565b6105fc615542565b604051610516919061dadf565b6104cf6156bb565b6104cf615a79565b6104cf61627e565b6104cf616462565b6104cf6164b3565b61063961663e565b604051610516919061db7b565b6104cf616709565b6104cf616a04565b61065e616bb5565b604051610516919061dbf0565b6104cf616cab565b6104cf616df4565b6104cf616f3d565b6104cf617069565b61065e61712f565b610639617225565b6104cf6172f0565b6104cf61742a565b6106b3617b43565b6040519015158152602001610516565b6104cf617c13565b6104cf617dc4565b6104cf617ee2565b6104cf618180565b6104cf618334565b6104cf6183f9565b6105096185a8565b6104cf618606565b61072a7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b604051908152602001610516565b6104cf6186cc565b6104cf61880f565b6104cf6189c3565b6104cf618b68565b6104cf618c86565b6104cf618e13565b6104cf618f57565b6104cf6192b3565b601f546106b39060ff1681565b6104cf61950b565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015610819575f80fd5b505af115801561082b573d5f803e3d5ffd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602d60405161085f919061ddaf565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b5f604051808303815f87803b1580156108be575f80fd5b505af11580156108d0573d5f803e3d5ffd5b50506020546024546040517f184b07930000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063184b079393506109259290911690602d9060040161ddc1565b5f604051808303815f87803b15801561093c575f80fd5b505af115801561094e573d5f803e3d5ffd5b50505050565b602680547fffffffffffffffffffffffff000000000000000000000000000000000000000090811630179091556027805490911661123417905560405161099a9061d811565b604051809103905ff0801580156109b3573d5f803e3d5ffd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600f81527f476174657761795a45564d2e736f6c000000000000000000000000000000000060208201526026549151602481019390935292166044820152610a97919060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f485cc9550000000000000000000000000000000000000000000000000000000017905261969a565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b039384168102919091179182905560208054919092049092167fffffffffffffffffffffffff0000000000000000000000000000000000000000909216919091178155604051737cce3eb018bf23e1fe2a32692f2c77592d110394915f9190610b3190820161d81f565b601f1982820381018352601f909101166040819052610b53919060200161ddf9565b60405160208183030381529060405290505f808251602084015ff590507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663b4d6c78284836001600160a01b0316803b806020016040519081016040528181525f908060200190933c6040518363ffffffff1660e01b8152600401610be492919061de04565b5f604051808303815f87803b158015610bfb575f80fd5b505af1158015610c0d573d5f803e3d5ffd5b50506026546020546040517fc0c53b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820181905260248201529082166044820152908616925063c0c53b8b91506064015f604051808303815f87803b158015610c80575f80fd5b505af1158015610c92573d5f803e3d5ffd5b5050602580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038781169182179092556020546040517fab7b49930000000000000000000000000000000000000000000000000000000081526004810192909252909116925063ab7b499391506024015f604051808303815f87803b158015610d22575f80fd5b505af1158015610d34573d5f803e3d5ffd5b505060208054604080517f2722feee00000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169450632722feee935060048082019392918290030181865afa158015610d96573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dba919061de40565b602880547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055604051610dfe9061d82d565b604051809103905ff080158015610e17573d5f803e3d5ffd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161781556028546040517f06447d5600000000000000000000000000000000000000000000000000000000815292166004830152737109709ecfa91a80626ff3989d68f67f5b1dd12d916306447d5691015f604051808303815f87803b158015610eb0575f80fd5b505af1158015610ec2573d5f803e3d5ffd5b505050505f805f604051610ed59061d83b565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103905ff080158015610f0e573d5f803e3d5ffd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392831690811790915560205460405160129360019384935f9391921690610f639061d849565b610f729695949392919061de92565b604051809103905ff080158015610f8b573d5f803e3d5ffd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081179091556023546040517fee2815ba0000000000000000000000000000000000000000000000000000000081526001600482015260248101929092529091169063ee2815ba906044015f604051808303815f87803b15801561101f575f80fd5b505af1158015611031573d5f803e3d5ffd5b50506023546040517fa7cb050700000000000000000000000000000000000000000000000000000000815260016004820181905260248201526001600160a01b03909116925063a7cb050791506044015f604051808303815f87803b158015611098575f80fd5b505af11580156110aa573d5f803e3d5ffd5b50506028546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152633b9aca006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d91506044015f604051808303815f87803b158015611127575f80fd5b505af1158015611139573d5f803e3d5ffd5b5050505060225f9054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004015f604051808303818588803b15801561118b575f80fd5b505af115801561119d573d5f803e3d5ffd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303815f875af115801561120e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611232919061df4b565b506021546026546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a060248201529116906347e7ef24906044016020604051808303815f875af11580156112a0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112c4919061df4b565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561131f575f80fd5b505af1158015611331573d5f803e3d5ffd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d5691506024015f604051808303815f87803b1580156113a4575f80fd5b505af11580156113b6573d5f803e3d5ffd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a060248201529116925063095ea7b391506044016020604051808303815f875af1158015611427573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061144b919061df4b565b5060225f9054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004015f604051808303818588803b15801561149a575f80fd5b505af11580156114ac573d5f803e3d5ffd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303815f875af115801561151d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611541919061df4b565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561159c575f80fd5b505af11580156115ae573d5f803e3d5ffd5b5050604080516080810182526026546001600160a01b0390811682525f602080840182815260018587019081528651928301909652918152606084018190528351602d80549185167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161781559251602e8054919095169116179092559251602f55909350909150603090611645908261dfdb565b50506040805160c0810190915260265460601b6bffffffffffffffffffffffff191660e082015290508060f4810160408051601f198184030181529181529082525f602083810182905260018484018190526060850183905260808501528251908101909252815260a090910152805160319081906116c4908261dfdb565b5060208201516001820180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039092169190911790556040820151600282015560608201516003820180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790556080820151600482015560a08201516005820190611761908261dfdb565b505050505050565b602854602054602480546040517f81bad6f300000000000000000000000000000000000000000000000000000000815260016004820181905292810183905260448101839052606481018390526001600160a01b03918216608482018190529294821631939190911631913190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801561180c575f80fd5b505af115801561181e573d5f803e3d5ffd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602d604051611852919061ddaf565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156118b0575f80fd5b505af11580156118c2573d5f803e3d5ffd5b50506020546024546040517fe90b9e5e0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063e90b9e5e93508892611917921690602d9060040161ddc1565b5f604051808303818588803b15801561192e575f80fd5b505af1158015611940573d5f803e3d5ffd5b50506028546001600160a01b031631925061196891506119629050868661e0c3565b826196b8565b6020546001600160a01b03163161197f84826196b8565b61176161198c878561e0d6565b6024546001600160a01b0316316196b8565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156119f4575f80fd5b505af1158015611a06573d5f803e3d5ffd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015611a5a575f80fd5b505af1158015611a6c573d5f803e3d5ffd5b50506020546021546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600160248201525f60448201529116925063f45346dc9150606401610925565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015611b56575f80fd5b505af1158015611b68573d5f803e3d5ffd5b505050507f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db76031604051611b9c919061e146565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015611bfa575f80fd5b505af1158015611c0c573d5f803e3d5ffd5b50506020546024546040517f2095dedb0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450632095dedb9350610925929091169060319060040161e158565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015611cb7575f80fd5b505af1158015611cc9573d5f803e3d5ffd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015611d36575f80fd5b505af1158015611d48573d5f803e3d5ffd5b50506020546027546040517ff340fa010000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063f340fa0191505f906024015b5f604051808303818588803b158015611dae575f80fd5b505af1158015611dc0573d5f803e3d5ffd5b5050505050565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015611e1d575f80fd5b505af1158015611e2f573d5f803e3d5ffd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611f1a919060040161e179565b5f604051808303815f87803b158015611f31575f80fd5b505af1158015611f43573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b81526004015f604051808303815f87803b158015611f93575f80fd5b505af1158015611fa5573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015611fff575f80fd5b505af1158015612011573d5f803e3d5ffd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506120fc919060040161e179565b5f604051808303815f87803b158015612113575f80fd5b505af1158015612125573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015612175575f80fd5b505af1158015612187573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b1580156121e1575f80fd5b505af11580156121f3573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b81526004015f604051808303815f87803b158015612243575f80fd5b505af1158015612255573d5f803e3d5ffd5b5050604051630618f58760e51b81527fd93c066500000000000000000000000000000000000000000000000000000000600482015260019250737109709ecfa91a80626ff3989d68f67f5b1dd12d915063c31eb0e0906024015f604051808303815f87803b1580156122c5575f80fd5b505af11580156122d7573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015612331575f80fd5b505af1158015612343573d5f803e3d5ffd5b50506020546021546027546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810187905290821660448201529116925063f45346dc91506064015f604051808303815f87803b1580156123b8575f80fd5b505af11580156123ca573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015612424575f80fd5b505af1158015612436573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015612486575f80fd5b505af1158015612498573d5f803e3d5ffd5b50506021546027546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f9450911691506370a08231906024015b602060405180830381865afa158015612502573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612526919061e18b565b90506125325f826196b8565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015612588575f80fd5b505af115801561259a573d5f803e3d5ffd5b50506020546021546027546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810188905290821660448201529116925063f45346dc91506064015f604051808303815f87803b15801561260f575f80fd5b505af1158015612621573d5f803e3d5ffd5b50506021546027546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa15801561268a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906126ae919061e18b565b90506126ba83826196b8565b505050565b6060601680548060200260200160405190810160405280929190818152602001828054801561271557602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116126f7575b5050505050905090565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015612775575f80fd5b505af1158015612787573d5f803e3d5ffd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b1580156127db575f80fd5b505af11580156127ed573d5f803e3d5ffd5b50506020546040517ff340fa010000000000000000000000000000000000000000000000000000000081525f60048201526001600160a01b03909116925063f340fa019150600190602401611d97565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156128a9575f80fd5b505af11580156128bb573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015612915575f80fd5b505af1158015612927573d5f803e3d5ffd5b50506020546028546040517ff340fa010000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063f340fa01915083906024015b5f604051808303818588803b15801561298d575f80fd5b505af1158015611761573d5f803e3d5ffd5b602854602080546024546040516001946001600160a01b0390811631949381163193921631915f916129d1910161e1a2565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff1916608083015291505f90806094810160408051808303601f190181529181529082526028546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015612abb575f80fd5b505af1158015612acd573d5f803e3d5ffd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e9450612b2893506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602854602054612b58936001600160a01b03928316928c92169061e1de565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015612bb6575f80fd5b505af1158015612bc8573d5f803e3d5ffd5b50506020546024546040517f30b103420000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506330b1034293508a92612c1f928792911690889060040161e28d565b5f604051808303818588803b158015612c36575f80fd5b505af1158015612c48573d5f803e3d5ffd5b50506028546001600160a01b0316319250612c6a91506119629050888861e0c3565b6020546001600160a01b031631612c8186826196b8565b612c8e61198c898761e0d6565b5050505050505050565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015612d01575f80fd5b505af1158015612d13573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024016108a7565b5f604051602001612d6b9061e1a2565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa7916024808301925f92919082900301818387803b158015612e17575f80fd5b505af1158015612e29573d5f803e3d5ffd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015612e7d575f80fd5b505af1158015612e8f573d5f803e3d5ffd5b50506020546024546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063bcf7f32b9350612eea9286925f9260019290911690899060040161e2c0565b5f604051808303815f87803b15801561298d575f80fd5b5f604051602001612f119061e1a2565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e0916024808301925f92919082900301818387803b158015612fd9575f80fd5b505af1158015612feb573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015b5f604051808303815f87803b158015613046575f80fd5b505af1158015613058573d5f803e3d5ffd5b50506020546021546024546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063bcf7f32b9450612eea938793811692600192911690899060040161e2c0565b6060601e805480602002602001604051908101604052809291908181526020015f905b828210156131e9575f84815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b828210156131d2578382905f5260205f200180546131479061dc85565b80601f01602080910402602001604051908101604052809291908181526020018280546131739061dc85565b80156131be5780601f10613195576101008083540402835291602001916131be565b820191905f5260205f20905b8154815290600101906020018083116131a157829003601f168201915b50505050508152602001906001019061312a565b5050505081525050815260200190600101906130d9565b50505050905090565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801561324b575f80fd5b505af115801561325d573d5f803e3d5ffd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b1580156132b1575f80fd5b505af11580156132c3573d5f803e3d5ffd5b50506020546040517fe90b9e5e0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063e90b9e5e91508390612976905f90602d9060040161ddc1565b6027546040516001600160a01b0390911660248201526001905f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae7600000000000000000000000000000000000000000000000000000000017905260285490517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b1580156133fc575f80fd5b505af115801561340e573d5f803e3d5ffd5b50506021546040517ff687d12a000000000000000000000000000000000000000000000000000000008152620186a060048201526001600160a01b03909116925063f687d12a91506024015f604051808303815f87803b158015613470575f80fd5b505af1158015613482573d5f803e3d5ffd5b50506021546040517ffc5fecd5000000000000000000000000000000000000000000000000000000008152620186a060048201525f93506001600160a01b03909116915063fc5fecd5906024016040805180830381865afa1580156134e9573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061350d919061e314565b6021546026549193506001600160a01b0390811692506347e7ef249116613534848761e0d6565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af1158015613594573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135b8919061df4b565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015613613575f80fd5b505af1158015613625573d5f803e3d5ffd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d5691506024015f604051808303815f87803b158015613698575f80fd5b505af11580156136aa573d5f803e3d5ffd5b50506021546020546001600160a01b03918216935063095ea7b39250166136d1848761e0d6565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af1158015613731573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613755919061df4b565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156137b0575f80fd5b505af11580156137c2573d5f803e3d5ffd5b505050505f60215f9054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613817573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061383b919061e18b565b602080546027546040519394506001600160a01b0391821693637b15118b936138819392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f198184030181526021548383018352620186a084526001602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1681526138f09391928a926001600160a01b03909116918a919060299060040161e39a565b5f604051808303815f87803b158015613907575f80fd5b505af1158015613919573d5f803e3d5ffd5b505050505f60215f9054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561396e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613992919061e18b565b90505f61399f848761e0d6565b90506117616139ae828561e0c3565b836040518060600160405280603a815260200162018aa3603a9139619734565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015613a24575f80fd5b505af1158015613a36573d5f803e3d5ffd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015613a8a575f80fd5b505af1158015613a9c573d5f803e3d5ffd5b50506020546021546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450639d4ba465935061092592909116906001905f90602d9060040161e407565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015613b4f575f80fd5b505af1158015613b61573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b81526004015f604051808303815f87803b158015613bb1575f80fd5b505af1158015613bc3573d5f803e3d5ffd5b5050604051630618f58760e51b81527fd93c0665000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015613c30575f80fd5b505af1158015613c42573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015b5f604051808303815f87803b158015613c9d575f80fd5b505af1158015613caf573d5f803e3d5ffd5b50506020546027546040517ff340fa010000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063f340fa0191508390602401612976565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015613d5b575f80fd5b505af1158015613d6d573d5f803e3d5ffd5b5050604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015613dda575f80fd5b505af1158015613dec573d5f803e3d5ffd5b50506020546021546028546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810187905290821660448201529116925063f45346dc91506064015b5f604051808303815f87803b158015611dae575f80fd5b5f604051602001613e729061e1a2565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa7916024808301925f92919082900301818387803b158015613f1e575f80fd5b505af1158015613f30573d5f803e3d5ffd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015613f84575f80fd5b505af1158015613f96573d5f803e3d5ffd5b50506020546021546040517fbcf7f32b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063bcf7f32b9350612eea92869216906001905f90899060040161e2c0565b6060601880548060200260200160405190810160405280929190818152602001828054801561271557602002820191905f5260205f209081546001600160a01b031681526001909101906020018083116126f7575050505050905090565b6060601780548060200260200160405190810160405280929190818152602001828054801561271557602002820191905f5260205f209081546001600160a01b031681526001909101906020018083116126f7575050505050905090565b60285460205460275460405163ca669fa760e01b81526001600160a01b039384166004820181905260019490319381163192163190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015614116575f80fd5b505af1158015614128573d5f803e3d5ffd5b50506020546027546040517ff340fa010000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063f340fa01915086906024015f604051808303818588803b15801561418d575f80fd5b505af115801561419f573d5f803e3d5ffd5b50506028546020546027546001600160a01b03928316319550908216319350163190506141d56141cf888861e0c3565b846196b8565b6141df85836196b8565b6141ec611962888661e0d6565b50505050505050565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801561424b575f80fd5b505af115801561425d573d5f803e3d5ffd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b1580156142b1575f80fd5b505af11580156142c3573d5f803e3d5ffd5b50506020546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450639d4ba4659350610925925f92600192911690602d9060040161e407565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015614388575f80fd5b505af115801561439a573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b1580156143f4575f80fd5b505af1158015614406573d5f803e3d5ffd5b50506020546028546040517fe90b9e5e0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063e90b9e5e93508592612976921690602d9060040161ddc1565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156144c4575f80fd5b505af11580156144d6573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015614530575f80fd5b505af1158015614542573d5f803e3d5ffd5b50506020546021546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450639d4ba465935061092592909116906001908590602d9060040161e407565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015614608575f80fd5b505af115801561461a573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015614674575f80fd5b505af1158015614686573d5f803e3d5ffd5b50506020546040517fe90b9e5e0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063e90b9e5e91508390612976908490602d9060040161ddc1565b5f6040516020016146e89061e1a2565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa7916024808301925f92919082900301818387803b158015614794575f80fd5b505af11580156147a6573d5f803e3d5ffd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015614813575f80fd5b505af1158015614825573d5f803e3d5ffd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca379450612eea9387938116925f92911690899060040161e2c0565b6040516001905f906148969060200161e1a2565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e0916024808301925f92919082900301818387803b15801561495e575f80fd5b505af1158015614970573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b1580156149ca575f80fd5b505af11580156149dc573d5f803e3d5ffd5b50506020546040517f30b103420000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506330b1034291508590614a2f9085908590889060040161e28d565b5f604051808303818588803b158015614a46575f80fd5b505af1158015612c8e573d5f803e3d5ffd5b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015614ac1575f80fd5b505af1158015614ad3573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015614b2d575f80fd5b505af1158015614b3f573d5f803e3d5ffd5b50506020546021546028546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba465945061092593928316926001921690602d9060040161e407565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b158015614c0e575f80fd5b505af1158015614c20573d5f803e3d5ffd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018690529116925063095ea7b391506044016020604051808303815f875af1158015614c8f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614cb3919061df4b565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015614d0e575f80fd5b505af1158015614d20573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015614d7a575f80fd5b505af1158015614d8c573d5f803e3d5ffd5b50506021546040517ff687d12a00000000000000000000000000000000000000000000000000000000815262030d4060048201526001600160a01b03909116925063f687d12a91506024015f604051808303815f87803b158015614dee575f80fd5b505af1158015614e00573d5f803e3d5ffd5b50506021546040517ffc5fecd500000000000000000000000000000000000000000000000000000000815262030d4060048201525f93508392506001600160a01b039091169063fc5fecd5906024016040805180830381865afa158015614e69573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614e8d919061e314565b60208054604080516001600160a01b03868116602483015290921660448301526064808301859052815180840390910181526084909201815291810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f667011120000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152929450909250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391614f609160040161e179565b5f604051808303815f87803b158015614f77575f80fd5b505af1158015614f89573d5f803e3d5ffd5b5050602080546027546040805160609290921b6bffffffffffffffffffffffff19168285015280518083036014018152602154600460348501818152605886018552605490950180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f19ff1d21000000000000000000000000000000000000000000000000000000001790528351808501855262030d40815260019781019790975292517f06cb89830000000000000000000000000000000000000000000000000000000081526001600160a01b0395861698506306cb8983975061507796929590911693926029910161e43d565b5f604051808303815f87803b15801561508e575f80fd5b505af11580156141ec573d5f803e3d5ffd5b60265460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156150f9575f80fd5b505af115801561510b573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561515b575f80fd5b505af115801561516d573d5f803e3d5ffd5b5050604051630618f58760e51b81527fd93c0665000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b1580156151da575f80fd5b505af11580156151ec573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015615246575f80fd5b505af1158015615258573d5f803e3d5ffd5b50506020546024546040517fe90b9e5e0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063e90b9e5e935085926152ad921690602d9060040161ddc1565b5f604051808303818588803b1580156152c4575f80fd5b505af11580156152d6573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063ca669fa7925060240190505f604051808303815f87803b158015615332575f80fd5b505af1158015615344573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015615394575f80fd5b505af11580156153a6573d5f803e3d5ffd5b5050602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a4015f604051808303815f87803b158015615436575f80fd5b505af1158015615448573d5f803e3d5ffd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602d60405161547c919061ddaf565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015b5f604051808303815f87803b1580156154db575f80fd5b505af11580156154ed573d5f803e3d5ffd5b50506020546024546040517fe90b9e5e0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063e90b9e5e93508592612976921690602d9060040161ddc1565b6060601b805480602002602001604051908101604052809291908181526020015f905b828210156131e9578382905f5260205f2090600202016040518060400160405290815f820180546155959061dc85565b80601f01602080910402602001604051908101604052809291908181526020018280546155c19061dc85565b801561560c5780601f106155e35761010080835404028352916020019161560c565b820191905f5260205f20905b8154815290600101906020018083116155ef57829003601f168201915b50505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156156a357602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116156505790505b50505050508152505081526020019060010190615565565b602154602480546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f9391909116916370a082319101602060405180830381865afa158015615721573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190615745919061e18b565b90506157515f826196b8565b5f6040516020016157619061e1a2565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff1916608083015291505f90806094810160408051808303601f190181529181529082526028546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801561584b575f80fd5b505af115801561585d573d5f803e3d5ffd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e94506158b893506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526028546020546158e9936001600160a01b0392831692600192169061e1de565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015615947575f80fd5b505af1158015615959573d5f803e3d5ffd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca3794506159b7938793811692600192911690899060040161e2c0565b5f604051808303815f87803b1580156159ce575f80fd5b505af11580156159e0573d5f803e3d5ffd5b5050602154602480546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f9550921692506370a082319101602060405180830381865afa158015615a48573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190615a6c919061e18b565b905061094e6001826196b8565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b158015615ae8575f80fd5b505af1158015615afa573d5f803e3d5ffd5b50506023546020546040515f945060129350600192600292620186a0926001600160a01b039283169290911690615b309061d849565b615b3f9695949392919061e497565b604051809103905ff080158015615b58573d5f803e3d5ffd5b506026546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152606460248201529192508216906347e7ef24906044016020604051808303815f875af1158015615bc4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190615be8919061df4b565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015615c43575f80fd5b505af1158015615c55573d5f803e3d5ffd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d5691506024015f604051808303815f87803b158015615cc8575f80fd5b505af1158015615cda573d5f803e3d5ffd5b50506020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260646024820152908416925063095ea7b391506044016020604051808303815f875af1158015615d47573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190615d6b919061df4b565b506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a0602482015291169063095ea7b3906044016020604051808303815f875af1158015615dd9573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190615dfd919061df4b565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015615e58575f80fd5b505af1158015615e6a573d5f803e3d5ffd5b50506027546040516001600160a01b039091166024820152600192505f915060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae76000000000000000000000000000000000000000000000000000000000179052517ffc5fecd5000000000000000000000000000000000000000000000000000000008152620186a060048201529091505f9081906001600160a01b0386169063fc5fecd5906024016040805180830381865afa158015615f46573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190615f6a919061e314565b6021549193509150615f869083906001600160a01b03166197b0565b602154604080517f18160ddd00000000000000000000000000000000000000000000000000000000815290515f926001600160a01b0316916318160ddd9160048083019260209291908290030181865afa158015615fe6573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061600a919061e18b565b90505f866001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015616049573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061606d919061e18b565b602080546027546040519394506001600160a01b0391821693637b15118b936160b39392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f198184030181528282018252620186a083526001602084015290517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168152616112928b918d918c9160299060040161e39a565b5f604051808303815f87803b158015616129575f80fd5b505af115801561613b573d5f803e3d5ffd5b505050505f60215f9054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015616190573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906161b4919061e18b565b90505f886001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156161f3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190616217919061e18b565b9050616246616226868661e0c3565b836040518060600160405280602b8152602001620189df602b9139619734565b616273616253898561e0c3565b8260405180606001604052806038815260200162018a3060389139619734565b505050505050505050565b5f60405160200161628e9061e1a2565b60408051601f19818403018152606080840190925260205490911b6bffffffffffffffffffffffff1916608083015291505f90806094810160408051808303601f190181529181529082526028546001600160a01b03908116602084015260019282018390526024805492517f81bad6f300000000000000000000000000000000000000000000000000000000815260048101859052908101849052604481018490526064810193909352166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015616378575f80fd5b505af115801561638a573d5f803e3d5ffd5b5050602080546040517fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e94506163e593506001600160a01b03909116910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602854602054616416936001600160a01b0392831692600192169061e1de565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa79060240161302f565b6021546027546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526001925f9216906370a08231906024016124e7565b6040516001905f906164c79060200161e1a2565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa7916024808301925f92919082900301818387803b158015616573575f80fd5b505af1158015616585573d5f803e3d5ffd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b1580156165d9575f80fd5b505af11580156165eb573d5f803e3d5ffd5b50506020546040517f30b103420000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506330b1034291508590614a2f9085905f90889060040161e28d565b6060601a805480602002602001604051908101604052809291908181526020015f905b828210156131e9578382905f5260205f2001805461667e9061dc85565b80601f01602080910402602001604051908101604052809291908181526020018280546166aa9061dc85565b80156166f55780601f106166cc576101008083540402835291602001916166f5565b820191905f5260205f20905b8154815290600101906020018083116166d857829003601f168201915b505050505081526020019060010190616661565b602154602480546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f9391909116916370a082319101602060405180830381865afa15801561676f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190616793919061e18b565b905061679f5f826196b8565b602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801561682b575f80fd5b505af115801561683d573d5f803e3d5ffd5b505050507fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4602d604051616871919061ddaf565b60405180910390a160285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156168cf575f80fd5b505af11580156168e1573d5f803e3d5ffd5b50506020546021546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba465945061693e93928316926001921690602d9060040161e407565b5f604051808303815f87803b158015616955575f80fd5b505af1158015616967573d5f803e3d5ffd5b5050602154602480546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f9550921692506370a082319101602060405180830381865afa1580156169cf573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906169f3919061e18b565b9050616a006001826196b8565b5050565b6040516001905f90616a189060200161e1a2565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e0916024808301925f92919082900301818387803b158015616ae0575f80fd5b505af1158015616af2573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015616b4c575f80fd5b505af1158015616b5e573d5f803e3d5ffd5b50506020546028546040517f30b103420000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506330b1034293508792614a2f928792911690889060040161e28d565b6060601d805480602002602001604051908101604052809291908181526020015f905b828210156131e9575f8481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015616c9357602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411616c405790505b50505050508152505081526020019060010190616bd8565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015616d04575f80fd5b505af1158015616d16573d5f803e3d5ffd5b5050604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015616d83575f80fd5b505af1158015616d95573d5f803e3d5ffd5b50506020546021546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101869052911660448201819052925063f45346dc9150606401613e4b565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015616e4a575f80fd5b505af1158015616e5c573d5f803e3d5ffd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015616ec9575f80fd5b505af1158015616edb573d5f803e3d5ffd5b50506020546021546027546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201525f602482015290821660448201529116925063f45346dc9150606401610925565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015616f93575f80fd5b505af1158015616fa5573d5f803e3d5ffd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015616ff9575f80fd5b505af115801561700b573d5f803e3d5ffd5b50506020546027546040517ff45346dc0000000000000000000000000000000000000000000000000000000081525f6004820152600160248201526001600160a01b0391821660448201529116925063f45346dc9150606401610925565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156170d5575f80fd5b505af11580156170e7573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401613c86565b6060601c805480602002602001604051908101604052809291908181526020015f905b828210156131e9575f8481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561720d57602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116171ba5790505b50505050508152505081526020019060010190617152565b60606019805480602002602001604051908101604052809291908181526020015f905b828210156131e9578382905f5260205f200180546172659061dc85565b80601f01602080910402602001604051908101604052809291908181526020018280546172919061dc85565b80156172dc5780601f106172b3576101008083540402835291602001916172dc565b820191905f5260205f20905b8154815290600101906020018083116172bf57829003601f168201915b505050505081526020019060010190617248565b604051630618f58760e51b81527f82d5d76a000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801561735c575f80fd5b505af115801561736e573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b1580156173c8575f80fd5b505af11580156173da573d5f803e3d5ffd5b50506020546040517ff340fa010000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201819052925063f340fa0191508390602401612976565b6028546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b158015617499575f80fd5b505af11580156174ab573d5f803e3d5ffd5b50506023546020546040515f945060129350600192600292620186a0926001600160a01b0392831692909116906174e19061d849565b6174f09695949392919061e497565b604051809103905ff080158015617509573d5f803e3d5ffd5b506026546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152606460248201529192508216906347e7ef24906044016020604051808303815f875af1158015617575573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190617599919061df4b565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156175f4575f80fd5b505af1158015617606573d5f803e3d5ffd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d5691506024015f604051808303815f87803b158015617679575f80fd5b505af115801561768b573d5f803e3d5ffd5b50506020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260646024820152908416925063095ea7b391506044016020604051808303815f875af11580156176f8573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061771c919061df4b565b506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a0602482015291169063095ea7b3906044016020604051808303815f875af115801561778a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906177ae919061df4b565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015617809575f80fd5b505af115801561781b573d5f803e3d5ffd5b50506040517ffc5fecd5000000000000000000000000000000000000000000000000000000008152620186a06004820152600192505f915081906001600160a01b0385169063fc5fecd5906024016040805180830381865afa158015617883573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906178a7919061e314565b60215491935091506178c39083906001600160a01b03166197b0565b602154604080517f18160ddd00000000000000000000000000000000000000000000000000000000815290515f926001600160a01b0316916318160ddd9160048083019260209291908290030181865afa158015617923573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190617947919061e18b565b90505f856001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015617986573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906179aa919061e18b565b602080546027546040519394506001600160a01b0391821693637c0dcb5f936179f09392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b604051602081830303815290604052878960296040518563ffffffff1660e01b8152600401617a22949392919061e526565b5f604051808303815f87803b158015617a39575f80fd5b505af1158015617a4b573d5f803e3d5ffd5b505050505f60215f9054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015617aa0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190617ac4919061e18b565b90505f876001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015617b03573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190617b27919061e18b565b9050617b36616226868661e0c3565b612c8e616253888561e0c3565b6008545f9060ff1615617b5a575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c6564000000000000000000000000000000000000000000000000000060248301525f9163667f9d7090604401602060405180830381865afa158015617be8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190617c0c919061e18b565b1415905090565b6040516001905f90617c279060200161e1a2565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e0916024808301925f92919082900301818387803b158015617cef575f80fd5b505af1158015617d01573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015617d5b575f80fd5b505af1158015617d6d573d5f803e3d5ffd5b50506020546024546040517f30b103420000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506330b1034293508792614a2f928792911690889060040161e28d565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015617e1a575f80fd5b505af1158015617e2c573d5f803e3d5ffd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015617e80575f80fd5b505af1158015617e92573d5f803e3d5ffd5b50506020546040517f184b07930000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063184b07939150610925905f90602d9060040161ddc1565b6021546027546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa158015617f49573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190617f6d919061e18b565b9050617f795f826196b8565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015617fcf575f80fd5b505af1158015617fe1573d5f803e3d5ffd5b5050604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b15801561804e575f80fd5b505af1158015618060573d5f803e3d5ffd5b50506020546021546027546040517ff45346dc0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810188905290821660448201529116925063f45346dc91506064015f604051808303815f87803b1580156180d5575f80fd5b505af11580156180e7573d5f803e3d5ffd5b50506021546027546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa158015618150573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190618174919061e18b565b90506126ba5f826196b8565b5f6040516020016181909061e1a2565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401526001838301528151630618f58760e51b81527f42c0407e0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e0916024808301925f92919082900301818387803b158015618258575f80fd5b505af115801561826a573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b1580156182c4575f80fd5b505af11580156182d6573d5f803e3d5ffd5b50506020546021546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca379450612eea938793811692600192911690899060040161e2c0565b60285460405163ca669fa760e01b81526001600160a01b0390911660048201525f90737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801561838c575f80fd5b505af115801561839e573d5f803e3d5ffd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024016154c4565b5f6040516020016184099061e1a2565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e0916024808301925f92919082900301818387803b1580156184d1575f80fd5b505af11580156184e3573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801561853d575f80fd5b505af115801561854f573d5f803e3d5ffd5b50506020546021546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca379350612eea92869216906001908690899060040161e2c0565b6060601580548060200260200160405190810160405280929190818152602001828054801561271557602002820191905f5260205f209081546001600160a01b031681526001909101906020018083116126f7575050505050905090565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015618672575f80fd5b505af1158015618684573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024016154c4565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015618722575f80fd5b505af1158015618734573d5f803e3d5ffd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b1580156187a1575f80fd5b505af11580156187b3573d5f803e3d5ffd5b50506020546021546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba465945061092593928316925f921690602d9060040161e407565b5f60405160200161881f9061e1a2565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401526001838301528151630618f58760e51b81527f82d5d76a0000000000000000000000000000000000000000000000000000000060048201529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163c31eb0e0916024808301925f92919082900301818387803b1580156188e7575f80fd5b505af11580156188f9573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015618953575f80fd5b505af1158015618965573d5f803e3d5ffd5b50506020546021546028546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063c39aca379450612eea938793811692600192911690899060040161e2c0565b5f806040516020016189d49061e1a2565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa7916024808301925f92919082900301818387803b158015618a80575f80fd5b505af1158015618a92573d5f803e3d5ffd5b5050604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015618aff575f80fd5b505af1158015618b11573d5f803e3d5ffd5b50506020546021546040517f30b103420000000000000000000000000000000000000000000000000000000081526001600160a01b0392831694506330b1034293508792614a2f928792911690889060040161e28d565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015618bbe575f80fd5b505af1158015618bd0573d5f803e3d5ffd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015618c24575f80fd5b505af1158015618c36573d5f803e3d5ffd5b50506020546040517f2095dedb0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632095dedb9150610925905f9060319060040161e158565b5f604051602001618c969061e1a2565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa7916024808301925f92919082900301818387803b158015618d42575f80fd5b505af1158015618d54573d5f803e3d5ffd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015618da8575f80fd5b505af1158015618dba573d5f803e3d5ffd5b50506020546021546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca379350612eea92869216906001905f90899060040161e2c0565b604051630618f58760e51b81527f42c0407e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015618e7c575f80fd5b505af1158015618e8e573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015618ee8575f80fd5b505af1158015618efa573d5f803e3d5ffd5b50506020546021546024546040517f9d4ba4650000000000000000000000000000000000000000000000000000000081526001600160a01b039384169550639d4ba465945061092593928316926001921690602d9060040161e407565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015618fb0575f80fd5b505af1158015618fc2573d5f803e3d5ffd5b50506021546040517ff687d12a00000000000000000000000000000000000000000000000000000000815261c35060048201526001600160a01b03909116925063f687d12a91506024015f604051808303815f87803b158015619023575f80fd5b505af1158015619035573d5f803e3d5ffd5b50506021546040517ffc5fecd500000000000000000000000000000000000000000000000000000000815261c35060048201525f93506001600160a01b03909116915063fc5fecd5906024016040805180830381865afa15801561909b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906190bf919061e314565b9150505f60215f9054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015619113573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190619137919061e18b565b602080546027546040519394506001600160a01b0391821693637c0dcb5f9361917d9392909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526191d59288916001600160a01b03169060299060040161e526565b5f604051808303815f87803b1580156191ec575f80fd5b505af11580156191fe573d5f803e3d5ffd5b505050505f60215f9054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015619253573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190619277919061e18b565b90505f619284848661e0d6565b9050611dc0619293828561e0c3565b8360405180606001604052806026815260200162018a0a60269139619734565b60285460405163ca669fa760e01b81526001600160a01b03909116600482018190523190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801561930d575f80fd5b505af115801561931f573d5f803e3d5ffd5b5061099992506108fc9150619337905060018461e0c3565b6040518115909202915f818181858888f1935050505015801561935c573d5f803e3d5ffd5b5060275460285460405163ca669fa760e01b81526001600160a01b03918216600482015291163190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156193ba575f80fd5b505af11580156193cc573d5f803e3d5ffd5b50506020546027546040516001600160a01b0391821660248201525f94509116915060029060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff340fa010000000000000000000000000000000000000000000000000000000017905251619457919061ddf9565b5f6040518083038185875af1925050503d805f8114619491576040519150601f19603f3d011682016040523d82523d5f602084013e619496565b606091505b505090506194be815f6040518060600160405280603b815260200162018a68603b9139619811565b60275460408051808201909152601d81527f746172676574206d757374206e6f7420726563656976652066756e647300000060208201526126ba916001600160a01b031631908490619734565b5f60405160200161951b9061e1a2565b60408051601f19818403018152606080840183526020805490911b6bffffffffffffffffffffffff191660808501528251808503607401815260948501845284526028546001600160a01b0316908401819052600184840152825163ca669fa760e01b815260048101919091529151909350737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa7916024808301925f92919082900301818387803b1580156195c7575f80fd5b505af11580156195d9573d5f803e3d5ffd5b5050604051630618f58760e51b8152637138356f60e01b6004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b15801561962d575f80fd5b505af115801561963f573d5f803e3d5ffd5b50506020546024546040517fc39aca370000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063c39aca379350612eea9286925f9260019290911690899060040161e2c0565b5f6196a361d857565b6196ae848483619865565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c54906044015b5f6040518083038186803b158015619722575f80fd5b505afa158015611761573d5f803e3d5ffd5b6040517f88b44c85000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d906388b44c85906197889086908690869060040161e55f565b5f6040518083038186803b15801561979e575f80fd5b505afa1580156141ec573d5f803e3d5ffd5b6040517f515361f60000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015282166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063515361f69060440161970c565b6040517f4db19e7e000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d90634db19e7e906197889086908690869060040161e57d565b5f8061987185846198df565b90506198d46040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f787900000081525082866040516020016198bf92919061de04565b604051602081830303815290604052856198ea565b9150505b9392505050565b5f6198d88383619917565b60c0810151515f901561990d5761990684848460c00151619931565b90506198d8565b6199068484619acf565b5f6199228383619bb4565b6198d8838360200151846198ea565b5f8061993b619bbf565b90505f6199488683619c8e565b90505f61995e826060015183602001518561a117565b90505f61996d8383898961a324565b90505f6199798261b190565b602081015181519192509060030b156199ec578982604001516040516020016199a392919061e59f565b60408051601f19818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526199e39160040161e179565b60405180910390fd5b5f619a2e6040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a20000000000000000000000081525083600161b351565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d90619a8190849060040161e179565b602060405180830381865afa158015619a9c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190619ac0919061de40565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081525f908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc92590619b2390879060040161e179565b5f60405180830381865afa158015619b3d573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052619b64919081019061e6be565b90505f619b918285604051602001619b7d92919061e6f0565b60405160208183030381529060405261b540565b90506001600160a01b0381166196ae5784846040516020016199a392919061e704565b616a0082825f61b551565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c90619c4690849060040161e794565b5f60405180830381865afa158015619c60573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052619c87919081019061e7da565b9250505090565b619cc06040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b5f737109709ecfa91a80626ff3989d68f67f5b1dd12d9050619d0a6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b619d138561b650565b60208201525f619d228661ba29565b90505f836001600160a01b031663d930a0e66040518163ffffffff1660e01b81526004015f60405180830381865afa158015619d60573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052619d87919081019061e7da565b86838560200151604051602001619da1949392919061e81f565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291505f906001600160a01b038616906360f9bb1190619df890859060040161e179565b5f60405180830381865afa158015619e12573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052619e39919081019061e7da565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f690619e8190849060040161e8ef565b602060405180830381865afa158015619e9c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190619ec0919061df4b565b619ed557816040516020016199a3919061e940565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890619f1a90849060040161e9c4565b5f60405180830381865afa158015619f34573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052619f5b919081019061e7da565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f690619fa290849060040161ea15565b602060405180830381865afa158015619fbd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190619fe1919061df4b565b1561a072576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061a02b90849060040161ea15565b5f60405180830381865afa15801561a045573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261a06c919081019061e7da565b60408501525b846001600160a01b03166349c4fac882865f015160405160200161a096919061ea66565b6040516020818303038152906040526040518363ffffffff1660e01b815260040161a0c292919061eac4565b5f60405180830381865afa15801561a0dc573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261a103919081019061e7da565b606085015250608083015250949350505050565b60408051600480825260a082019092526060915f9190816020015b606081526020019060019003908161a1325790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250815f8151811061a1915761a19161eae8565b60200260200101819052506040518060400160405280600381526020017f2d726c00000000000000000000000000000000000000000000000000000000008152508160018151811061a1e55761a1e561eae8565b60200260200101819052508460405160200161a201919061eb15565b6040516020818303038152906040528160028151811061a2235761a22361eae8565b60200260200101819052508260405160200161a23f919061eb73565b6040516020818303038152906040528160038151811061a2615761a26161eae8565b60200260200101819052505f61a2768261b190565b602080820151604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000818501908152825180840184525f808252908601528251808401909352905182529281019290925291925061a305906040805180820182525f808252602091820152815180830190925284518252808501908201529061bca5565b61a31a57856040516020016199a3919061ebab565b9695505050505050565b60a08101516040805180820182525f80825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d901561a373565b511590565b61a4e75782602001511561a42f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a4016199e3565b8260c001511561a4e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a4016199e3565b6040805160ff80825261200082019092525f91816020015b606081526020019060019003908161a4ff5790505090505f6040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061a5599061ec28565b935060ff168151811061a56e5761a56e61eae8565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e370000000000000000000000000000000000000081525060405160200161a5bf919061ec46565b60405160208183030381529060405282828061a5da9061ec28565b935060ff168151811061a5ef5761a5ef61eae8565b60200260200101819052506040518060400160405280600681526020017f6465706c6f79000000000000000000000000000000000000000000000000000081525082828061a63c9061ec28565b935060ff168151811061a6515761a65161eae8565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d6500000000000000000000000000000000000081525082828061a69e9061ec28565b935060ff168151811061a6b35761a6b361eae8565b6020026020010181905250876020015182828061a6cf9061ec28565b935060ff168151811061a6e45761a6e461eae8565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163745061746800000000000000000000000000000000000081525082828061a7319061ec28565b935060ff168151811061a7465761a74661eae8565b60209081029190910101528751828261a75e8161ec28565b935060ff168151811061a7735761a77361eae8565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e4964000000000000000000000000000000000000000000000081525082828061a7c09061ec28565b935060ff168151811061a7d55761a7d561eae8565b602002602001018190525061a7e94661bd03565b828261a7f48161ec28565b935060ff168151811061a8095761a80961eae8565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c65000000000000000000000000000000000081525082828061a8569061ec28565b935060ff168151811061a86b5761a86b61eae8565b60200260200101819052508682828061a8839061ec28565b935060ff168151811061a8985761a89861eae8565b602090810291909101015285511561a9bb5760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f646500000000000000000000006020820152828261a8e98161ec28565b935060ff168151811061a8fe5761a8fe61eae8565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d9061a94e90899060040161e179565b5f60405180830381865afa15801561a968573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261a98f919081019061e7da565b828261a99a8161ec28565b935060ff168151811061a9af5761a9af61eae8565b60200260200101819052505b84602001511561aa8b5760408051808201909152601281527f2d2d766572696679536f75726365436f646500000000000000000000000000006020820152828261aa048161ec28565b935060ff168151811061aa195761aa1961eae8565b60200260200101819052506040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525082828061aa669061ec28565b935060ff168151811061aa7b5761aa7b61eae8565b602002602001018190525061ac50565b61aac261a36e8660a001516040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b61ab555760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261ab058161ec28565b935060ff168151811061ab1a5761ab1a61eae8565b60200260200101819052508460a0015160405160200161ab3a919061eb15565b60405160208183030381529060405282828061aa669061ec28565b8460c0015115801561ab975750604080890151815180830183525f8082526020918201528251808401909352815183529081019082015261ab9590511590565b155b1561ac505760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261abdb8161ec28565b935060ff168151811061abf05761abf061eae8565b602002602001018190525061ac048861bda0565b60405160200161ac14919061eb15565b60405160208183030381529060405282828061ac2f9061ec28565b935060ff168151811061ac445761ac4461eae8565b60200260200101819052505b604080860151815180830183525f8082526020918201528251808401909352815183529081019082015261ac8390511590565b61ad185760408051808201909152600b81527f2d2d72656c6179657249640000000000000000000000000000000000000000006020820152828261acc68161ec28565b935060ff168151811061acdb5761acdb61eae8565b6020026020010181905250846040015182828061acf79061ec28565b935060ff168151811061ad0c5761ad0c61eae8565b60200260200101819052505b60608501511561ae355760408051808201909152600681527f2d2d73616c7400000000000000000000000000000000000000000000000000006020820152828261ad618161ec28565b935060ff168151811061ad765761ad7661eae8565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e8906024015f60405180830381865afa15801561ade2573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261ae09919081019061e7da565b828261ae148161ec28565b935060ff168151811061ae295761ae2961eae8565b60200260200101819052505b60e0850151511561aedb5760408051808201909152600a81527f2d2d6761734c696d6974000000000000000000000000000000000000000000006020820152828261ae7f8161ec28565b935060ff168151811061ae945761ae9461eae8565b602002602001018190525061aeaf8560e001515f015161bd03565b828261aeba8161ec28565b935060ff168151811061aecf5761aecf61eae8565b60200260200101819052505b60e0850151602001511561af855760408051808201909152600a81527f2d2d6761735072696365000000000000000000000000000000000000000000006020820152828261af288161ec28565b935060ff168151811061af3d5761af3d61eae8565b602002602001018190525061af598560e001516020015161bd03565b828261af648161ec28565b935060ff168151811061af795761af7961eae8565b60200260200101819052505b60e0850151604001511561b02f5760408051808201909152600e81527f2d2d6d61784665655065724761730000000000000000000000000000000000006020820152828261afd28161ec28565b935060ff168151811061afe75761afe761eae8565b602002602001018190525061b0038560e001516040015161bd03565b828261b00e8161ec28565b935060ff168151811061b0235761b02361eae8565b60200260200101819052505b60e0850151606001511561b0d95760408051808201909152601681527f2d2d6d61785072696f72697479466565506572476173000000000000000000006020820152828261b07c8161ec28565b935060ff168151811061b0915761b09161eae8565b602002602001018190525061b0ad8560e001516060015161bd03565b828261b0b88161ec28565b935060ff168151811061b0cd5761b0cd61eae8565b60200260200101819052505b5f8160ff1667ffffffffffffffff81111561b0f65761b0f661df6a565b60405190808252806020026020018201604052801561b12957816020015b606081526020019060019003908161b1145790505b5090505f5b8260ff168160ff16101561b18157838160ff168151811061b1515761b15161eae8565b6020026020010151828260ff168151811061b16e5761b16e61eae8565b602090810291909101015260010161b12e565b5093505050505b949350505050565b61b1b660405180606001604052805f60030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d925f91849163d145736c9161b23b9186910161ec9d565b5f60405180830381865afa15801561b255573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261b27c919081019061e7da565b90505f61b289868361c87c565b90505f846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b815260040161b2b8919061db7b565b5f604051808303815f875af115801561b2d3573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261b2fa919081019061ece3565b805190915060030b1580159061b3135750602081015151155b801561b3225750604081015151155b1561a31a57815f8151811061b3395761b33961eae8565b60200260200101516040516020016199a3919061ed92565b60605f61b384856040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f8082526020918201528151808301909252865182528087019082015290915061b3ba9082905b9061c9ce565b1561b512575f61b4348261b42e8461b42861b3fb8a6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925282518252918201519181019190915290565b9061c9f4565b9061ca52565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015290915061b49790829061c9ce565b1561b50057604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261b4fd905b829061cad6565b90505b61b5098161cafb565b925050506198d8565b821561b52b5784846040516020016199a392919061ef6f565b505060408051602081019091525f81526198d8565b5f808251602084015ff09392505050565b8160a001511561b56057505050565b5f61b56c84848461cb60565b90505f61b5788261b190565b602081015181519192509060030b15801561b6125750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261b612906040805180820182525f8082526020918201528151808301909252845182528085019082015261b3b4565b1561b61f57505050505050565b6040820151511561b63f5781604001516040516020016199a3919061eff6565b806040516020016199a3919061f04d565b60605f61b683836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015290915061b6e7905b829061bca5565b1561b75557604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526198d89061b75090839061d0f5565b61cafb565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261b7b6905b829061d17d565b60010361b88157604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261b81b9061b4f6565b50604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526198d89061b750905b839061cad6565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261b8df9061b6e0565b1561ba1257604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082018190528451808601909552925184528301529061b94690839061d211565b90505f816001835161b958919061e0c3565b8151811061b9685761b96861eae8565b6020026020010151905061ba0961b75061b9dd6040518060400160405280600581526020017f2e6a736f6e0000000000000000000000000000000000000000000000000000008152506040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925285518252808601908201529061d0f5565b95945050505050565b826040516020016199a3919061f0a4565b50919050565b60605f61ba5c836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015290915061babd9061b6e0565b1561bacb576198d88161cafb565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261bb299061b7af565b60010361bb9257604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526198d89061b7509061b87a565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261bbf09061b6e0565b1561ba1257604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082018190528451808601909552925184528301529061bc5790839061d211565b905060018151111561bc9357806002825161bc72919061e0c3565b8151811061bc825761bc8261eae8565b602002602001015192505050919050565b50826040516020016199a3919061f0a4565b805182515f91111561bcb857505f6196b2565b8151835160208501515f929161bccd9161e0d6565b61bcd7919061e0c3565b90508260200151810361bcee5760019150506196b2565b82516020840151819020912014905092915050565b60605f61bd0f8361d2bc565b60010190505f8167ffffffffffffffff81111561bd2e5761bd2e61df6a565b6040519080825280601f01601f19166020018201604052801561bd58576020820181803683370190505b5090508181016020015b5f19017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461bd6257509392505050565b604081810151815180830183525f808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e534544000000000000000000000000000000000000000000008184019081528551808701875283815284019290925284518086019095525184529083015260609161be2b905b829061d39d565b1561be6b57505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261bec99061be24565b1561bf0957505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261bf679061be24565b1561bfa757505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c0059061be24565b8061c0695750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c0699061be24565b1561c0a957505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c1079061be24565b8061c16b5750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c16b9061be24565b1561c1ab57505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c2099061be24565b8061c26d5750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c26d9061be24565b1561c2ad57505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c30b9061be24565b8061c36f5750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c36f9061be24565b1561c3af57505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c40d9061be24565b1561c44d57505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c4ab9061be24565b1561c4eb57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c5499061be24565b1561c58957505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c5e79061be24565b1561c62757505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c6859061be24565b1561c6c557505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c7239061be24565b8061c7875750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c7879061be24565b1561c7c757505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261c8259061be24565b1561c86557505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b604080840151845191516199a3929060200161f174565b6060805f5b845181101561c906578185828151811061c89d5761c89d61eae8565b602002602001015160405160200161c8b692919061e6f0565b60405160208183030381529060405291506001855161c8d5919061e0c3565b811461c8fe578160405160200161c8ec919061f2c2565b60405160208183030381529060405291505b60010161c881565b50604080516003808252608082019092525f91816020015b606081526020019060019003908161c91e57905050905083815f8151811061c9485761c94861eae8565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061c99c5761c99c61eae8565b6020026020010181905250818160028151811061c9bb5761c9bb61eae8565b6020908102919091010152949350505050565b60208083015183518351928401515f9361c9eb929184919061d3b0565b14159392505050565b604080518082019091525f80825260208201525f61ca22845f01518560200151855f0151866020015161d4bf565b905083602001518161ca34919061e0c3565b8451859061ca4390839061e0c3565b90525060208401525090919050565b604080518082019091525f808252602082015281518351101561ca765750816196b2565b602080830151908401516001911461ca9d5750815160208481015190840151829020919020145b801561cace5782518451859061cab490839061e0c3565b905250825160208501805161caca90839061e0d6565b9052505b509192915050565b604080518082019091525f808252602082015261caf483838361d5db565b5092915050565b60605f825f015167ffffffffffffffff81111561cb1a5761cb1a61df6a565b6040519080825280601f01601f19166020018201604052801561cb44576020820181803683370190505b5090505f60208201905061caf4818560200151865f015161d681565b60605f61cb6b619bbf565b6040805160ff80825261200082019092529192505f9190816020015b606081526020019060019003908161cb875790505090505f6040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061cbe19061ec28565b935060ff168151811061cbf65761cbf661eae8565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161cc47919061f2fa565b60405160208183030381529060405282828061cc629061ec28565b935060ff168151811061cc775761cc7761eae8565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061ccc49061ec28565b935060ff168151811061ccd95761ccd961eae8565b60200260200101819052508260405160200161ccf5919061eb73565b60405160208183030381529060405282828061cd109061ec28565b935060ff168151811061cd255761cd2561eae8565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061cd729061ec28565b935060ff168151811061cd875761cd8761eae8565b602002602001018190525061cd9c878461d6fa565b828261cda78161ec28565b935060ff168151811061cdbc5761cdbc61eae8565b60209081029190910101528551511561ce675760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261ce0e8161ec28565b935060ff168151811061ce235761ce2361eae8565b602002602001018190525061ce3b865f01518461d6fa565b828261ce468161ec28565b935060ff168151811061ce5b5761ce5b61eae8565b60200260200101819052505b85608001511561ced55760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261ceb08161ec28565b935060ff168151811061cec55761cec561eae8565b602002602001018190525061cf3b565b841561cf3b5760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261cf1a8161ec28565b935060ff168151811061cf2f5761cf2f61eae8565b60200260200101819052505b6040860151511561cfd75760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261cf858161ec28565b935060ff168151811061cf9a5761cf9a61eae8565b6020026020010181905250856040015182828061cfb69061ec28565b935060ff168151811061cfcb5761cfcb61eae8565b60200260200101819052505b85606001511561d0415760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d65730000000000000000000000006020820152828261d0208161ec28565b935060ff168151811061d0355761d03561eae8565b60200260200101819052505b5f8160ff1667ffffffffffffffff81111561d05e5761d05e61df6a565b60405190808252806020026020018201604052801561d09157816020015b606081526020019060019003908161d07c5790505b5090505f5b8260ff168160ff16101561d0e957838160ff168151811061d0b95761d0b961eae8565b6020026020010151828260ff168151811061d0d65761d0d661eae8565b602090810291909101015260010161d096565b50979650505050505050565b604080518082019091525f808252602082015281518351101561d1195750816196b2565b8151835160208501515f929161d12e9161e0d6565b61d138919061e0c3565b6020840151909150600190821461d159575082516020840151819020908220145b801561d1745783518551869061d17090839061e0c3565b9052505b50929392505050565b5f80825f015161d19d855f01518660200151865f0151876020015161d4bf565b61d1a7919061e0d6565b90505b8351602085015161d1bb919061e0d6565b811161caf4578161d1cb8161f32b565b925050825f015161d20085602001518361d1e5919061e0c3565b865161d1f1919061e0c3565b83865f0151876020015161d4bf565b61d20a919061e0d6565b905061d1aa565b60605f61d21e848461d17d565b61d22990600161e0d6565b67ffffffffffffffff81111561d2415761d24161df6a565b60405190808252806020026020018201604052801561d27457816020015b606081526020019060019003908161d25f5790505b5090505f5b815181101561d2b45761d28f61b750868661cad6565b82828151811061d2a15761d2a161eae8565b602090810291909101015260010161d279565b509392505050565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061d304577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061d330576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061d34e57662386f26fc10000830492506010015b6305f5e100831061d366576305f5e100830492506008015b612710831061d37a57612710830492506004015b6064831061d38c576064830492506002015b600a83106196b25760010192915050565b5f61d3a8838361d739565b159392505050565b5f8085841161d4b5576020841161d461575f841561d3f957600161d3d586602061e0c3565b61d3e090600861f343565b61d3eb90600261f43d565b61d3f5919061e0c3565b1990505b835181168561d408898961e0d6565b61d412919061e0c3565b805190935082165b81811461d44c5787841161d434578794505050505061b188565b8361d43e8161f448565b94505082845116905061d41a565b61d456878561e0d6565b94505050505061b188565b83832061d46e858861e0c3565b61d478908761e0d6565b91505b85821061d4b35784822080820361d4a05761d496868461e0d6565b935050505061b188565b61d4ab60018461e0c3565b92505061d47b565b505b5092949350505050565b5f838186851161d5c6576020851161d576575f851561d50957600161d4e587602061e0c3565b61d4f090600861f343565b61d4fb90600261f43d565b61d505919061e0c3565b1990505b845181165f8761d5198b8b61e0d6565b61d523919061e0c3565b855190915083165b82811461d5685781861061d5505761d5438b8b61e0d6565b965050505050505061b188565b8561d55a8161f32b565b96505083865116905061d52b565b85965050505050505061b188565b508383205f905b61d587868961e0c3565b821161d5c45785832080820361d5a3578394505050505061b188565b61d5ae60018561e0d6565b935050818061d5bc9061f32b565b92505061d57d565b505b61d5d0878761e0d6565b979650505050505050565b604080518082019091525f80825260208201525f61d609855f01518660200151865f0151876020015161d4bf565b60208087018051918601919091525190915061d625908261e0c3565b83528451602086015161d638919061e0d6565b810361d646575f855261d678565b8351835161d654919061e0d6565b8551869061d66390839061e0c3565b905250835161d672908261e0d6565b60208601525b50909392505050565b6020811061d6b9578151835261d69860208461e0d6565b925061d6a560208361e0d6565b915061d6b260208261e0c3565b905061d681565b5f19811561d6e757600161d6ce83602061e0c3565b61d6da9061010061f43d565b61d6e4919061e0c3565b90505b9151835183169219169190911790915250565b60605f61d7078484619c8e565b805160208083015160405193945061d7219390910161f45d565b60405160208183030381529060405291505092915050565b815181515f919081111561d74b575081515b602080850151908401515f5b8381101561d802578251825180821461d7d2575f19602087101561d7b15760018461d78389602061e0c3565b61d78d919061e0d6565b61d79890600861f343565b61d7a390600261f43d565b61d7ad919061e0c3565b1990505b818116838216818103911461d7cf5797506196b29650505050505050565b50505b61d7dd60208661e0d6565b945061d7ea60208561e0d6565b9350505060208161d7fb919061e0d6565b905061d757565b508451865161a31a919061f498565b610b09806200f4b883390190565b615048806200ffc183390190565b6109e8806201500983390190565b61102e80620159f183390190565b611fc08062016a1f83390190565b6040518060e001604052806060815260200160608152602001606081526020015f151581526020015f151581526020015f1515815260200161d89761d89c565b905290565b6040518061010001604052805f151581526020015f15158152602001606081526020015f801916815260200160608152602001606081526020015f1515815260200161d89760405180608001604052805f81526020015f81526020015f81526020015f81525090565b602080825282518282018190525f918401906040840190835b8181101561d9455783516001600160a01b031683526020938401939092019160010161d91e565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561da77577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b8801810191908801905f5b8181101561da5d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261da4784865161d950565b602095860195909450929092019160010161da0d565b50919750505060209485019492909201915060010161d9a4565b50929695505050505050565b5f8151808452602084019350602083015f5b8281101561dad55781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161da95565b5093949350505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561da77577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261db49604088018261d950565b905060208201519150868103602088015261db64818361da83565b96505050602093840193919091019060010161db05565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561da77577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845261dbdb85835161d950565b9450602093840193919091019060010161dba1565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561da77577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261dc6f604087018261da83565b955050602093840193919091019060010161dc16565b600181811c9082168061dc9957607f821691505b60208210810361ba23577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f815461dcdc8161dc85565b80855260018216801561dcf6576001811461dd305761dd64565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0083166020870152602082151560051b870101935061dd64565b845f5260205f205f5b8381101561dd5b5781546020828a01015260018201915060208101905061dd39565b87016020019450505b50505092915050565b6001600160a01b0381541682526001600160a01b03600182015416602083015260028101546040830152608060608301525f6198d8608084016003840161dcd0565b602081525f6198d8602083018461dd6d565b6001600160a01b0383168152604060208201525f61b188604083018461dd6d565b5f81518060208401855e5f93019283525090919050565b5f6198d8828461dde2565b6001600160a01b0383168152604060208201525f61b188604083018461d950565b80516001600160a01b038116811461de3b575f80fd5b919050565b5f6020828403121561de50575f80fd5b6198d88261de25565b6003811061de8e577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9052565b610100815260056101008201527f544f4b454e000000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f544b4e00000000000000000000000000000000000000000000000000000000006101608201525f6101808201905060ff8816604083015286606083015261df21608083018761de59565b8460a08301526001600160a01b03841660c083015261d5d060e08301846001600160a01b03169052565b5f6020828403121561df5b575f80fd5b815180151581146198d8575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b601f8211156126ba57805f5260205f20601f840160051c8101602085101561dfbc5750805b601f840160051c820191505b81811015611dc0575f815560010161dfc8565b815167ffffffffffffffff81111561dff55761dff561df6a565b61e0098161e003845461dc85565b8461df97565b6020601f82116001811461e03b575f831561e0245750848201515b5f19600385901b1c1916600184901b178455611dc0565b5f84815260208120601f198516915b8281101561e06a578785015182556020948501946001909201910161e04a565b508482101561e08757868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156196b2576196b261e096565b808201808211156196b2576196b261e096565b60c082525f61e0fb60c084018361dcd0565b6001600160a01b0360018401541660208501526002830154604085015260ff600384015416151560608501526004830154608085015283810360a08501526196ae816005850161dcd0565b602081525f6198d8602083018461e0e9565b6001600160a01b0383168152604060208201525f61b188604083018461e0e9565b602081525f6198d8602083018461d950565b5f6020828403121561e19b575f80fd5b5051919050565b602081525f6196b260208301600581527f68656c6c6f000000000000000000000000000000000000000000000000000000602082015260400190565b60a081525f61e1f060a083018761d950565b6001600160a01b03861660208401528460408401526001600160a01b0384166060840152828103608084015261d5d081600581527f68656c6c6f000000000000000000000000000000000000000000000000000000602082015260400190565b5f81516060845261e264606085018261d950565b90506001600160a01b036020840151166020850152604083015160408501528091505092915050565b606081525f61e29f606083018661e250565b6001600160a01b0385166020840152828103604084015261a31a818561d950565b60a081525f61e2d260a083018861e250565b6001600160a01b03871660208401528560408401526001600160a01b0385166060840152828103608084015261e308818561d950565b98975050505050505050565b5f806040838503121561e325575f80fd5b61e32e8361de25565b6020939093015192949293505050565b5f81546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b03600183015416604084015260a0606084015261e38660a084016002840161dcd0565b600383015460808501528091505092915050565b60e081525f61e3ac60e083018961d950565b8760208401526001600160a01b0387166040840152828103606084015261e3d3818761d950565b855160808501526020860151151560a0850152905082810360c084015261e3fa818561e33e565b9998505050505050505050565b6001600160a01b03851681528360208201526001600160a01b0383166040820152608060608201525f61a31a608083018461dd6d565b60c081525f61e44f60c083018861d950565b6001600160a01b0387166020840152828103604084015261e470818761d950565b85516060850152602086015115156080850152905082810360a084015261e308818561e33e565b610100815260066101008201527f5345434f4e440000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f53454300000000000000000000000000000000000000000000000000000000006101608201525f6101808201905060ff8816604083015286606083015261df21608083018761de59565b608081525f61e538608083018761d950565b8560208401526001600160a01b0385166040840152828103606084015261d5d0818561e33e565b838152826020820152606060408201525f61ba09606083018461d950565b83151581528215156020820152606060408201525f61ba09606083018461d950565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f61e5d0601a83018561dde2565b7f3a2000000000000000000000000000000000000000000000000000000000000081526198d4600282018561dde2565b6040516060810167ffffffffffffffff8111828210171561e6235761e62361df6a565b60405290565b5f8067ffffffffffffffff84111561e6435761e64361df6a565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561e6725761e67261df6a565b60405283815290508082840185101561e689575f80fd5b8383602083015e5f60208583010152509392505050565b5f82601f83011261e6af575f80fd5b6198d88383516020850161e629565b5f6020828403121561e6ce575f80fd5b815167ffffffffffffffff81111561e6e4575f80fd5b6196ae8482850161e6a0565b5f61b18861e6fe838661dde2565b8461dde2565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f61e735601a83018561dde2565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000815261e765601982018561dde2565b7f2200000000000000000000000000000000000000000000000000000000000000815260010195945050505050565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201525f6198d8608083018461d950565b5f6020828403121561e7ea575f80fd5b815167ffffffffffffffff81111561e800575f80fd5b8201601f8101841361e810575f80fd5b6196ae8482516020840161e629565b5f61e82a828761dde2565b7f2f00000000000000000000000000000000000000000000000000000000000000815261e85a600182018761dde2565b90507f2f00000000000000000000000000000000000000000000000000000000000000815261e88c600182018661dde2565b90507f2f00000000000000000000000000000000000000000000000000000000000000815261e8be600182018561dde2565b7f2e6a736f6e0000000000000000000000000000000000000000000000000000008152600501979650505050505050565b604081525f61e901604083018461d950565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081525f61e971601f83018461dde2565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f81527f6d6c00000000000000000000000000000000000000000000000000000000000060208201526022019392505050565b604081525f61e9d6604083018461d950565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b604081525f61ea27604083018461d950565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081525f61ea97601483018461dde2565b7f275d2e6b656363616b32353600000000000000000000000000000000000000008152600c019392505050565b604081525f61ead6604083018561d950565b82810360208401526198d4818561d950565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f220000000000000000000000000000000000000000000000000000000000000081525f61eb46600183018461dde2565b7f220000000000000000000000000000000000000000000000000000000000000081526001019392505050565b5f61eb7e828461dde2565b7f2f6275696c642d696e666f0000000000000000000000000000000000000000008152600b019392505050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201525f6198d8604b83018461dde2565b5f60ff821660ff810361ec3d5761ec3d61e096565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201525f6198d8602983018461dde2565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201525f6198d8608083018461d950565b5f6020828403121561ecf3575f80fd5b815167ffffffffffffffff81111561ed09575f80fd5b82016060818503121561ed1a575f80fd5b61ed2261e600565b81518060030b811461ed32575f80fd5b8152602082015167ffffffffffffffff81111561ed4d575f80fd5b61ed598682850161e6a0565b602083015250604082015167ffffffffffffffff81111561ed78575f80fd5b61ed848682850161e6a0565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201525f61ede9602183018461dde2565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657481527f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960208201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560408201527f616c69666965642070617468206f66207468652062617368206578656375746160608201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960808201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60a08201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960c08201527f6f75722070726f6a65637420287573696e6720666f727761726420736c61736860e08201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101008201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061012082015261013b019392505050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201525f61efc6602183018561dde2565b7f2720696e206f75747075743a200000000000000000000000000000000000000081526198d4600d82018561dde2565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201525f6198d8602983018461dde2565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201525f6198d8602283018461dde2565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081525f61f0d5600e83018461dde2565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e7472616381527f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e7360208201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e747260408201527f6163742e6a736f6e00000000000000000000000000000000000000000000000060608201526068019392505050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081525f61f1a5601883018561dde2565b7f20696e2000000000000000000000000000000000000000000000000000000000815261f1d5600482018561dde2565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f727465642081527f6c6963656e736520666f7220626c6f636b206578706c6f72657220766572696660208201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f60408201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c60608201527f206f7220736574207468652060736b69704c6963656e73655479706560206f7060808201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060a082015260b70195945050505050565b5f61f2cd828461dde2565b7f200000000000000000000000000000000000000000000000000000000000000081526001019392505050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081525f6198d8601c83018461dde2565b5f5f19820361f33c5761f33c61e096565b5060010190565b80820281158282048414176196b2576196b261e096565b6001815b600184111561f3955780850481111561f3795761f37961e096565b600184161561f38757908102905b60019390931c92800261f35e565b935093915050565b5f8261f3ab575060016196b2565b8161f3b757505f6196b2565b816001811461f3cd576002811461f3d75761f3f3565b60019150506196b2565b60ff84111561f3e85761f3e861e096565b50506001821b6196b2565b5060208310610133831016604e8410600b841016171561f416575081810a6196b2565b61f4225f19848461f35a565b805f190482111561f4355761f43561e096565b029392505050565b5f6198d8838361f39d565b5f8161f4565761f45661e096565b505f190190565b5f61f468828561dde2565b7f3a0000000000000000000000000000000000000000000000000000000000000081526198d4600182018561dde2565b8181035f83128015838313168383128216171561caf45761caf461e09656fe60c0604052600d60809081526c2bb930b83832b21022ba3432b960991b60a0525f9061002b908261010b565b506040805180820190915260048152630ae8aa8960e31b6020820152600190610054908261010b565b506002805460ff1916601217905534801561006d575f80fd5b506101c5565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061009b57607f821691505b6020821081036100b957634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561010657805f5260205f20601f840160051c810160208510156100e45750805b601f840160051c820191505b81811015610103575f81556001016100f0565b50505b505050565b81516001600160401b0381111561012457610124610073565b610138816101328454610087565b846100bf565b6020601f82116001811461016a575f83156101535750848201515b5f19600385901b1c1916600184901b178455610103565b5f84815260208120601f198516915b828110156101995787850151825560209485019460019092019101610179565b50848210156101b657868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610937806101d25f395ff3fe6080604052600436106100bb575f3560e01c8063313ce56711610071578063a9059cbb1161004c578063a9059cbb146101eb578063d0e30db01461020a578063dd62ed3e14610212575f80fd5b8063313ce5671461018157806370a08231146101ac57806395d89b41146101d7575f80fd5b806318160ddd116100a157806318160ddd1461012757806323b872dd146101435780632e1a7d4d14610162575f80fd5b806306fdde03146100ce578063095ea7b3146100f8575f80fd5b366100ca576100c8610248565b005b5f80fd5b3480156100d9575f80fd5b506100e26102a2565b6040516100ef919061071f565b60405180910390f35b348015610103575f80fd5b5061011761011236600461079a565b61032d565b60405190151581526020016100ef565b348015610132575f80fd5b50475b6040519081526020016100ef565b34801561014e575f80fd5b5061011761015d3660046107c2565b6103a6565b34801561016d575f80fd5b506100c861017c3660046107fc565b610628565b34801561018c575f80fd5b5060025461019a9060ff1681565b60405160ff90911681526020016100ef565b3480156101b7575f80fd5b506101356101c6366004610813565b60036020525f908152604090205481565b3480156101e2575f80fd5b506100e26106ff565b3480156101f6575f80fd5b5061011761020536600461079a565b61070c565b6100c8610248565b34801561021d575f80fd5b5061013561022c36600461082c565b600460209081525f928352604080842090915290825290205481565b335f908152600360205260408120805434929061026690849061088a565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b5f80546102ae9061089d565b80601f01602080910402602001604051908101604052809291908181526020018280546102da9061089d565b80156103255780601f106102fc57610100808354040283529160200191610325565b820191905f5260205f20905b81548152906001019060200180831161030857829003601f168201915b505050505081565b335f81815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103949086815260200190565b60405180910390a35060015b92915050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260036020526040812054821115610412576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201525f60248201526044015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff84163314801590610487575073ffffffffffffffffffffffffffffffffffffffff84165f9081526004602090815260408083203384529091529020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14155b156105435773ffffffffffffffffffffffffffffffffffffffff84165f9081526004602090815260408083203384529091529020548211156104fe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201525f6024820152604401610409565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526004602090815260408083203384529091528120805484929061053d9084906108ee565b90915550505b73ffffffffffffffffffffffffffffffffffffffff84165f90815260036020526040812080548492906105779084906108ee565b909155505073ffffffffffffffffffffffffffffffffffffffff83165f90815260036020526040812080548492906105b090849061088a565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161061691815260200190565b60405180910390a35060019392505050565b335f90815260036020526040902054811115610679576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201525f6024820152604401610409565b335f90815260036020526040812080548392906106979084906108ee565b9091555050604051339082156108fc029083905f818181858888f193505050501580156106c6573d5f803e3d5ffd5b5060405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a250565b600180546102ae9061089d565b5f6107183384846103a6565b9392505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610795575f80fd5b919050565b5f80604083850312156107ab575f80fd5b6107b483610772565b946020939093013593505050565b5f805f606084860312156107d4575f80fd5b6107dd84610772565b92506107eb60208501610772565b929592945050506040919091013590565b5f6020828403121561080c575f80fd5b5035919050565b5f60208284031215610823575f80fd5b61071882610772565b5f806040838503121561083d575f80fd5b61084683610772565b915061085460208401610772565b90509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156103a0576103a061085d565b600181811c908216806108b157607f821691505b6020821081036108e8577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b818103818111156103a0576103a061085d56fea2646970667358221220f82a6621bc6ae2f40b7ff1dde0e016bd4b523e6e1df6d8c4401566e37294755f64736f6c634300081a003360a060405230608052348015610013575f80fd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614f4f6100f95f395f81816127fe0152818161282701526129d90152614f4f5ff3fe608060405260043610610291575f3560e01c80638f28397011610165578063bd8fde1c116100c6578063d547741f1161007c578063e9d6c5ba11610062578063e9d6c5ba146107f7578063f354b31f14610828578063f851a44014610847575f80fd5b8063d547741f146107a5578063e63ab1e9146107c4575f80fd5b8063c1bd469f116100ac578063c1bd469f14610746578063cc5ad8b614610767578063d3523ea214610786575f80fd5b8063bd8fde1c146106f4578063c0c53b8b14610727575f80fd5b8063a217fddf1161011b578063a8f2cb9611610101578063a8f2cb961461066e578063aa808c061461068d578063ad3cb1cc146106ac575f80fd5b8063a217fddf14610645578063a3ebd14c14610658575f80fd5b806391d148541161014b57806391d14854146105a057806394cc8683146106035780639ca220dd14610624575f80fd5b80638f283970146105625780639060bda914610581575f80fd5b80633f4ba83a1161020f578063631d62e4116101c55780637066b18d116101ab5780637066b18d146104f5578063804ea334146105215780638456cb591461054e575f80fd5b8063631d62e4146104b75780636e9e2d3f146104d6575f80fd5b806352d1902d116101f557806352d1902d146104405780635c975abb146104545780635cf92c9f1461048a575f80fd5b80633f4ba83a146104195780634f1ef2861461042d575f80fd5b80632259e9e5116102645780632f2ff15d1161024a5780632f2ff15d146103bc5780633500c24b146103db57806336568abe146103fa575f80fd5b80632259e9e514610342578063248a9ca314610361575f80fd5b806301ffc9a7146102955780630c63109e146102c957806310d29b9e1461030057806318d3ce9614610321575b5f80fd5b3480156102a0575f80fd5b506102b46102af366004613eec565b610865565b60405190151581526020015b60405180910390f35b3480156102d4575f80fd5b506001546102e8906001600160a01b031681565b6040516001600160a01b0390911681526020016102c0565b34801561030b575f80fd5b5061031f61031a366004613f7d565b6108fd565b005b34801561032c575f80fd5b506103356109b6565b6040516102c09190614006565b34801561034d575f80fd5b5061031f61035c3660046140c5565b610c47565b34801561036c575f80fd5b506103ae61037b36600461413e565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016102c0565b3480156103c7575f80fd5b5061031f6103d6366004614169565b610cd9565b3480156103e6575f80fd5b5061031f6103f5366004614197565b610d22565b348015610405575f80fd5b5061031f610414366004614169565b610eb4565b348015610424575f80fd5b5061031f610f05565b61031f61043b3660046141df565b610f1a565b34801561044b575f80fd5b506103ae610f39565b34801561045f575f80fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166102b4565b348015610495575f80fd5b506104a96104a43660046142a6565b610f67565b6040516102c09291906142ee565b3480156104c2575f80fd5b5061031f6104d13660046140c5565b61105e565b3480156104e1575f80fd5b5061031f6104f0366004614310565b611104565b348015610500575f80fd5b5061051461050f3660046142a6565b6111c3565b6040516102c091906143e1565b34801561052c575f80fd5b5061054061053b36600461413e565b611288565b6040516102c09291906143f3565b348015610559575f80fd5b5061031f61133d565b34801561056d575f80fd5b5061031f61057c366004614197565b61136f565b34801561058c575f80fd5b5061031f61059b366004614414565b6114bd565b3480156105ab575f80fd5b506102b46105ba366004614169565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561060e575f80fd5b5061061761154b565b6040516102c09190614440565b34801561062f575f80fd5b506106386115a1565b6040516102c09190614482565b348015610650575f80fd5b506103ae5f81565b348015610663575f80fd5b506103ae6207a12081565b348015610679575f80fd5b5061031f61068836600461452d565b61175a565b348015610698575f80fd5b506102e86106a73660046142a6565b6117da565b3480156106b7575f80fd5b506105146040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b3480156106ff575f80fd5b506103ae7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa381565b348015610732575f80fd5b5061031f61074136600461459c565b611819565b348015610751575f80fd5b5061075a611bc1565b6040516102c091906145e4565b348015610772575f80fd5b50600b546102e8906001600160a01b031681565b348015610791575f80fd5b506105146107a03660046140c5565b611ebb565b3480156107b0575f80fd5b5061031f6107bf366004614169565b611f9f565b3480156107cf575f80fd5b506103ae7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b348015610802575f80fd5b50610816610811366004614197565b611fe2565b6040516102c0969594939291906146d9565b348015610833575f80fd5b5061031f610842366004614736565b61222d565b348015610852575f80fd5b505f546102e8906001600160a01b031681565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806108f757507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3610927816122c9565b61092f6122d3565b61093b85858585612331565b6109478585858561247f565b7f6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c60075f8781526020019081526020015f2085856040516109899291906147de565b90815260200160405180910390206001016040516109a791906148bd565b60405180910390a15050505050565b6004546060908067ffffffffffffffff8111156109d5576109d56141b2565b604051908082528060200260200182016040528015610a3157816020015b610a1e60405180608001604052805f1515815260200160608152602001606081526020015f81525090565b8152602001906001900390816109f35790505b5091505f5b81811015610c42575f60048281548110610a5257610a526148cf565b905f5260205f2090600202016040518060400160405290815f8201548152602001600182018054610a82906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054610aae906147ed565b8015610af95780601f10610ad057610100808354040283529160200191610af9565b820191905f5260205f20905b815481529060010190602001808311610adc57829003601f168201915b50505050508152505090505f815f015190505f82602001519050604051806080016040528060075f8581526020019081526020015f2083604051610b3d9190614913565b90815260408051602092819003830190205460ff16151583525f8681526007835281902090519290910191610b73908590614913565b90815260200160405180910390206001018054610b8f906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054610bbb906147ed565b8015610c065780601f10610bdd57610100808354040283529160200191610c06565b820191905f5260205f20905b815481529060010190602001808311610be957829003601f168201915b5050505050815260200182815260200183815250868581518110610c2c57610c2c6148cf565b6020908102919091010152505050600101610a36565b505090565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3610c71816122c9565b610c796122d3565b610c868686868686612501565b610c938686868686612594565b857f40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee63486868686604051610cc99493929190614947565b60405180910390a2505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610d12816122c9565b610d1c8383612611565b50505050565b5f610d2c816122c9565b6001600160a01b038216610d6c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d967ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa383612611565b50610dc17f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a83612611565b50600154610df9907ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3906001600160a01b03166126dd565b50600154610e31907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a906001600160a01b03166126dd565b50600154604080516001600160a01b03928316815291841660208301527f6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279910160405180910390a150600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6001600160a01b0381163314610ef6576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f0082826126dd565b505050565b5f610f0f816122c9565b610f17612781565b50565b610f226127f3565b610f2b826128c3565b610f3582826128cd565b5050565b5f610f426129ce565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f83815260076020526040808220905160609190610f8890869086906147de565b908152604080519182900360209081018320545f898152600790925291902060ff909116935090610fbc90869086906147de565b90815260200160405180910390206001018054610fd8906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611004906147ed565b801561104f5780601f106110265761010080835404028352916020019161104f565b820191905f5260205f20905b81548152906001019060200180831161103257829003601f168201915b50505050509050935093915050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3611088816122c9565b6110906122d3565b61109d8686868686612a30565b6110aa8686868686612d0a565b84846040516110ba9291906147de565b6040518091039020867f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce258185856040516110f4929190614978565b60405180910390a3505050505050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361112e816122c9565b6111366122d3565b6111478a8a8a8a8a8a8a8a8a612d87565b6111588a8a8a8a8a8a8a8a8a6130b3565b896001600160a01b031686866040516111729291906147de565b60405180910390207fa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367848a8d8d6040516111af949392919061498b565b60405180910390a350505050505050505050565b606060065f8581526020019081526020015f2060040183836040516111e99291906147de565b90815260200160405180910390208054611202906147ed565b80601f016020809104026020016040519081016040528092919081815260200182805461122e906147ed565b80156112795780601f1061125057610100808354040283529160200191611279565b820191905f5260205f20905b81548152906001019060200180831161125c57829003601f168201915b505050505090505b9392505050565b5f8181526006602052604090206002810154600390910180546001600160a01b0390921691606091906112ba906147ed565b80601f01602080910402602001604051908101604052809291908181526020018280546112e6906147ed565b80156113315780601f1061130857610100808354040283529160200191611331565b820191905f5260205f20905b81548152906001019060200180831161131457829003601f168201915b50505050509050915091565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a611367816122c9565b610f17613144565b5f611379816122c9565b6001600160a01b0382166113b9576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113c35f83612611565b506113ee7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a83612611565b505f805461140591906001600160a01b03166126dd565b505f5461143c907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a906001600160a01b03166126dd565b505f54604080516001600160a01b03928316815291841660208301527f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f910160405180910390a1505f80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa36114e7816122c9565b6114ef6122d3565b6114f9838361319f565b611503838361328d565b604080516001600160a01b038516815283151560208201527f9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8910160405180910390a1505050565b6060600280548060200260200160405190810160405280929190818152602001828054801561159757602002820191905f5260205f20905b815481526020019060010190808311611583575b5050505050905090565b6003546060908067ffffffffffffffff8111156115c0576115c06141b2565b60405190808252806020026020018201604052801561162e57816020015b604080516080810182525f80825260208083018290529282015260608082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092019101816115de5790505b5091505f5b81811015610c42575f6003828154811061164f5761164f6148cf565b5f918252602080832090910154604080516080810182528285526006808552828620805460ff161515835282860185905260028101546001600160a01b03169383019390935294839052939092526003909101805491935060608301916116b5906147ed565b80601f01602080910402602001604051908101604052809291908181526020018280546116e1906147ed565b801561172c5780601f106117035761010080835404028352916020019161172c565b820191905f5260205f20905b81548152906001019060200180831161170f57829003601f168201915b5050505050815250848381518110611746576117466148cf565b602090810291909101015250600101611633565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3611784816122c9565b61178c6122d3565b611799868686868661330f565b6117a686868686866134bf565b857fc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e83604051610cc9911515815260200190565b5f838152600a602052604080822090516117f790859085906147de565b908152604051908190036020019020546001600160a01b031690509392505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156118635750825b90505f8267ffffffffffffffff16600114801561187f5750303b155b90508115801561188d575080155b156118c4576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156119255784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038816158061194257506001600160a01b038716155b8061195457506001600160a01b038616155b1561198b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61199361353c565b61199b61353c565b6119a3613544565b6119ad5f89612611565b506119d87ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa388612611565b50611a037f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a88612611565b50611a2e7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a89612611565b505f80546001600160a01b038a81167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316178355600180548b8316908416178155600b8054928b16929093169190911790915546808352600660208181526040808620805460ff1916909517855580513060601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016818401528151808203601401815260349091019091529290945290925260030190611af190826149fb565b50600280546001818101909255467f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9091018190556003805492830181555f527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b909101558315611bb75784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b6005546060908067ffffffffffffffff811115611be057611be06141b2565b604051908082528060200260200182016040528015611c5b57816020015b611c486040518060e001604052805f151581526020015f6001600160a01b03168152602001606081526020015f815260200160608152602001606081526020015f60ff1681525090565b815260200190600190039081611bfe5790505b5091505f5b81811015610c42575f60058281548110611c7c57611c7c6148cf565b5f9182526020808320909101546001600160a01b0390811680845260088352604093849020845160e081018652815460ff811615158252610100900490931693830193909352600183018054919550919384019190611cda906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611d06906147ed565b8015611d515780601f10611d2857610100808354040283529160200191611d51565b820191905f5260205f20905b815481529060010190602001808311611d3457829003601f168201915b5050505050815260200160028201548152602001600382018054611d74906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611da0906147ed565b8015611deb5780601f10611dc257610100808354040283529160200191611deb565b820191905f5260205f20905b815481529060010190602001808311611dce57829003601f168201915b50505050508152602001600482018054611e04906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611e30906147ed565b8015611e7b5780601f10611e5257610100808354040283529160200191611e7b565b820191905f5260205f20905b815481529060010190602001808311611e5e57829003601f168201915b50505091835250506005919091015460ff166020909101528451859084908110611ea757611ea76148cf565b602090810291909101015250600101611c60565b606060075f8781526020019081526020015f208585604051611ede9291906147de565b90815260200160405180910390206003018383604051611eff9291906147de565b90815260200160405180910390208054611f18906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611f44906147ed565b8015611f8f5780601f10611f6657610100808354040283529160200191611f8f565b820191905f5260205f20905b815481529060010190602001808311611f7257829003601f168201915b5050505050905095945050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611fd8816122c9565b610d1c83836126dd565b6001600160a01b038082165f908152600860209081526040808320815160e081018352815460ff81161515825261010090049095169285019290925260018201805493946060948694869485948794859490939284019190612043906147ed565b80601f016020809104026020016040519081016040528092919081815260200182805461206f906147ed565b80156120ba5780601f10612091576101008083540402835291602001916120ba565b820191905f5260205f20905b81548152906001019060200180831161209d57829003601f168201915b50505050508152602001600282015481526020016003820180546120dd906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054612109906147ed565b80156121545780601f1061212b57610100808354040283529160200191612154565b820191905f5260205f20905b81548152906001019060200180831161213757829003601f168201915b5050505050815260200160048201805461216d906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054612199906147ed565b80156121e45780601f106121bb576101008083540402835291602001916121e4565b820191905f5260205f20905b8154815290600101906020018083116121c757829003601f168201915b50505091835250506005919091015460ff16602090910152805160808201516060830151604084015160a085015160c090950151939d929c50909a509850919650945092505050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3612257816122c9565b61225f6122d3565b61226e88888888888888613577565b61227d888888888888886136ca565b877faea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c58888888888886040516122b796959493929190614af4565b60405180910390a25050505050505050565b610f17813361374b565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff161561232f576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f8481526006602052604090205460ff16612380576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018590526024015b60405180910390fd5b5f8290036123bd5782826040517ec10cfd000000000000000000000000000000000000000000000000000000008152600401612377929190614978565b5f848152600760205260409081902090516123db90859085906147de565b908152602001604051809103902060010180546123f7906147ed565b90505f03612437578383836040517f2b4f9c8600000000000000000000000000000000000000000000000000000000815260040161237793929190614b3c565b8060075f8681526020019081526020015f2084846040516124599291906147de565b908152604051908190036020019020805491151560ff1990921691909117905550505050565b5f848484846040516024016124979493929190614b55565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f10d29b9e0000000000000000000000000000000000000000000000000000000017905290506124fa816137d7565b5050505050565b5f8581526006602052604090205460ff1661254b576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612377565b818160065f8881526020019081526020015f2060040186866040516125719291906147de565b9081526020016040518091039020918261258c929190614b81565b505050505050565b5f85858585856040516024016125ae959493929190614c79565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f2259e9e500000000000000000000000000000000000000000000000000000000179052905061258c816137d7565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff166126d4575f848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561268a3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506108f7565b5f9150506108f7565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16156126d4575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506108f7565b612789613876565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061288c57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166128807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b1561232f576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610f35816122c9565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612927575060408051601f3d908101601f1916820190925261292491810190614cb1565b60015b612968576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401612377565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146129c4576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401612377565b610f0083836138d1565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461232f576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8581526006602052604090205460ff16612a7a576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612377565b5f839003612ab75783836040517ec10cfd000000000000000000000000000000000000000000000000000000008152600401612377929190614978565b5f819003612af1576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f858152600760205260408082209051612b0e90879087906147de565b90815260200160405180910390206001018054612b2a906147ed565b90501115612b6e5784848484846040517f2b192eab000000000000000000000000000000000000000000000000000000008152600401612377959493929190614c79565b600160075f8781526020019081526020015f208585604051612b919291906147de565b9081526040805160209281900383018120805460ff1916941515949094179093555f888152600790925290208391839190612bcf90889088906147de565b90815260200160405180910390206001019182612bed929190614b81565b50838360075f8881526020019081526020015f208686604051612c119291906147de565b90815260200160405180910390206002019182612c2f929190614b81565b506004604051806040016040528087815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920182905250939094525050835460018181018655948252602091829020845160029092020190815590830151929390929083019150612cae90826149fb565b5050508383604051612cc19291906147de565b6040518091039020857f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce25818484604051612cfb929190614978565b60405180910390a35050505050565b5f8585858585604051602401612d24959493929190614c79565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f631d62e400000000000000000000000000000000000000000000000000000000179052905061258c816137d7565b6001600160a01b038916612dc7576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f879003612e30576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f53796d626f6c2063616e6e6f7420626520656d707479000000000000000000006044820152606401612377565b6001600160a01b038981165f9081526008602052604090205461010090041615612e91576040517f63f4ee1f0000000000000000000000000000000000000000000000000000000081526001600160a01b038a166004820152602401612377565b5f6001600160a01b031660098989604051612ead9291906147de565b908152604051908190036020019020546001600160a01b031614612f015787876040517fb295cac9000000000000000000000000000000000000000000000000000000008152600401612377929190614978565b6001600160a01b0389165f818152600860205260409020805460017fffffffffffffffffffffff000000000000000000000000000000000000000000909116610100909302929092178217815501612f5a858783614b81565b506001600160a01b0389165f90815260086020526040902060028101879055600301612f87888a83614b81565b506001600160a01b0389165f90815260086020526040902060058101805460ff191660ff8416179055600401612fbe838583614b81565b5088600a5f8881526020019081526020015f208686604051612fe19291906147de565b90815260200160405180910390205f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555088600989896040516130259291906147de565b90815260405190819003602001902080546001600160a01b039283167fffffffffffffffffffffffff000000000000000000000000000000000000000091821617909155600580546001810182555f919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180549b9092169a16999099179098555050505050505050565b5f8989898989898989896040516024016130d599989796959493929190614cc8565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6e9e2d3f000000000000000000000000000000000000000000000000000000001790529050613138816137d7565b50505050505050505050565b61314c6122d3565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336127d5565b6001600160a01b0382166131df576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038281165f90815260086020526040902054610100900416613263576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5a52433230206e6f7420726567697374657265640000000000000000000000006044820152606401612377565b6001600160a01b03919091165f908152600860205260409020805460ff1916911515919091179055565b6040516001600160a01b038316602482015281151560448201525f9060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f9060bda9000000000000000000000000000000000000000000000000000000001790529050610f00816137d7565b5f8581526006602052604090205460ff1680156133295750805b15613363576040517fa1452cb000000000000000000000000000000000000000000000000000000000815260048101869052602401612377565b5f8581526006602052604090205460ff1615801561337f575080155b156133b9576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612377565b5f858152600660205260409020600201546001600160a01b03161580156133e05750468514155b1561341a57600380546001810182555f919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018590555b5f858152600660205260409020805460ff19168215151781556002810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038716179055600301613476838583614b81565b5080156134b657600280546001810182555f919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace018590556124fa565b6124fa85613926565b5f85858585856040516024016134d9959493929190614d32565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa8f2cb9600000000000000000000000000000000000000000000000000000000179052905061258c816137d7565b61232f6139cb565b61354c6139cb565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff19169055565b5f8781526006602052604090205460ff166135c1576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101889052602401612377565b5f8590036135fe5785856040517ec10cfd000000000000000000000000000000000000000000000000000000008152600401612377929190614978565b5f8781526007602052604090819020905161361c90889088906147de565b9081526040519081900360200190205460ff1661366b578686866040517f2b4f9c8600000000000000000000000000000000000000000000000000000000815260040161237793929190614b3c565b818160075f8a81526020019081526020015f20888860405161368e9291906147de565b908152602001604051809103902060030186866040516136af9291906147de565b90815260200160405180910390209182611bb7929190614b81565b5f878787878787876040516024016136e89796959493929190614d6e565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff354b31f000000000000000000000000000000000000000000000000000000001790529050611bb7816137d7565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610f35576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401612377565b5f5b600254811015610f355746600282815481106137f7576137f76148cf565b905f5260205f2001541480613845575060065f6002838154811061381d5761381d6148cf565b905f5260205f20015481526020019081526020015f206003018054613841906147ed565b1590505b61386e5761386e6002828154811061385f5761385f6148cf565b905f5260205f20015483613a32565b6001016137d9565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1661232f576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6138da82613ce4565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561391e57610f008282613d8b565b610f35613dfd565b5f5b600254811015610f35578160028281548110613946576139466148cf565b905f5260205f200154036139c3576002805461396490600190614dbd565b81548110613974576139746148cf565b905f5260205f20015460028281548110613990576139906148cf565b5f9182526020909120015560028054806139ac576139ac614df5565b600190038181905f5260205f20015f905590555050565b600101613928565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661232f576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820182526207a12081525f6020808301829052835160a0810185528281529081018290529283018190526060808401526080830152905f848152600660205260408082206002015490517ffc5fecd50000000000000000000000000000000000000000000000000000000081526207a12060048201526001600160a01b039091169190829063fc5fecd5906024016040805180830381865afa158015613adf573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b039190614e22565b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018290529092506001600160a01b03841691506323b872dd906064016020604051808303815f875af1158015613b70573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b949190614e4e565b613bca576040517f90b8ec1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018390529083169063095ea7b3906044016020604051808303815f875af1158015613c33573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613c579190614e4e565b50600b545f878152600660205260409081902090517f06cb89830000000000000000000000000000000000000000000000000000000081526001600160a01b03909216916306cb898391613cbb9160039091019086908a908a908a90600401614e69565b5f604051808303815f87803b158015613cd2575f80fd5b505af1158015613138573d5f803e3d5ffd5b806001600160a01b03163b5f03613d32576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401612377565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f80846001600160a01b031684604051613da79190614913565b5f60405180830381855af49150503d805f8114613ddf576040519150601f19603f3d011682016040523d82523d5f602084013e613de4565b606091505b5091509150613df4858383613e35565b95945050505050565b341561232f576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606082613e4a57613e4582613eaa565b611281565b8151158015613e6157506001600160a01b0384163b155b15613ea3576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401612377565b5080611281565b805115613eba5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60208284031215613efc575f80fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611281575f80fd5b5f8083601f840112613f3b575f80fd5b50813567ffffffffffffffff811115613f52575f80fd5b602083019150836020828501011115613f69575f80fd5b9250929050565b8015158114610f17575f80fd5b5f805f8060608587031215613f90575f80fd5b84359350602085013567ffffffffffffffff811115613fad575f80fd5b613fb987828801613f2b565b9094509250506040850135613fcd81613f70565b939692955090935050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156140b9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180511515865260208101516080602088015261407c6080880182613fd8565b9050604082015187820360408901526140958282613fd8565b6060938401519890930197909752509450602093840193919091019060010161402c565b50929695505050505050565b5f805f805f606086880312156140d9575f80fd5b85359450602086013567ffffffffffffffff8111156140f6575f80fd5b61410288828901613f2b565b909550935050604086013567ffffffffffffffff811115614121575f80fd5b61412d88828901613f2b565b969995985093965092949392505050565b5f6020828403121561414e575f80fd5b5035919050565b6001600160a01b0381168114610f17575f80fd5b5f806040838503121561417a575f80fd5b82359150602083013561418c81614155565b809150509250929050565b5f602082840312156141a7575f80fd5b813561128181614155565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f80604083850312156141f0575f80fd5b82356141fb81614155565b9150602083013567ffffffffffffffff811115614216575f80fd5b8301601f81018513614226575f80fd5b803567ffffffffffffffff811115614240576142406141b2565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff82111715614270576142706141b2565b604052818152828201602001871015614287575f80fd5b816020840160208301375f602083830101528093505050509250929050565b5f805f604084860312156142b8575f80fd5b83359250602084013567ffffffffffffffff8111156142d5575f80fd5b6142e186828701613f2b565b9497909650939450505050565b8215158152604060208201525f6143086040830184613fd8565b949350505050565b5f805f805f805f805f60c08a8c031215614328575f80fd5b893561433381614155565b985060208a013567ffffffffffffffff81111561434e575f80fd5b61435a8c828d01613f2b565b90995097505060408a0135955060608a013567ffffffffffffffff811115614380575f80fd5b61438c8c828d01613f2b565b90965094505060808a013567ffffffffffffffff8111156143ab575f80fd5b6143b78c828d01613f2b565b90945092505060a08a013560ff811681146143d0575f80fd5b809150509295985092959850929598565b602081525f6112816020830184613fd8565b6001600160a01b0383168152604060208201525f6143086040830184613fd8565b5f8060408385031215614425575f80fd5b823561443081614155565b9150602083013561418c81613f70565b602080825282518282018190525f918401906040840190835b81811015614477578351835260209384019390920191600101614459565b509095945050505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156140b9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805115158652602081015160208701526001600160a01b03604082015116604087015260608101519050608060608701526145176080870182613fd8565b95505060209384019391909101906001016144a8565b5f805f805f60808688031215614541575f80fd5b85359450602086013561455381614155565b9350604086013567ffffffffffffffff81111561456e575f80fd5b61457a88828901613f2b565b909450925050606086013561458e81613f70565b809150509295509295909350565b5f805f606084860312156145ae575f80fd5b83356145b981614155565b925060208401356145c981614155565b915060408401356145d981614155565b809150509250925092565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156140b9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051151586526001600160a01b036020820151166020870152604081015160e0604088015261466d60e0880182613fd8565b905060608201516060880152608082015187820360808901526146908282613fd8565b91505060a082015187820360a08901526146aa8282613fd8565b91505060c082015191506146c360c088018360ff169052565b955050602093840193919091019060010161460a565b861515815260c060208201525f6146f360c0830188613fd8565b866040840152828103606084015261470b8187613fd8565b9050828103608084015261471f8186613fd8565b91505060ff831660a0830152979650505050505050565b5f805f805f805f6080888a03121561474c575f80fd5b87359650602088013567ffffffffffffffff811115614769575f80fd5b6147758a828b01613f2b565b909750955050604088013567ffffffffffffffff811115614794575f80fd5b6147a08a828b01613f2b565b909550935050606088013567ffffffffffffffff8111156147bf575f80fd5b6147cb8a828b01613f2b565b989b979a50959850939692959293505050565b818382375f9101908152919050565b600181811c9082168061480157607f821691505b602082108103614838577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f815461484a816147ed565b8085526001821680156148645760018114614880576148b4565b60ff1983166020870152602082151560051b87010193506148b4565b845f5260205f205f5b838110156148ab5781546020828a010152600182019150602081019050614889565b87016020019450505b50505092915050565b602081525f611281602083018461483e565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81518060208401855e5f93019283525090919050565b5f61128182846148fc565b81835281816020850137505f602082840101525f6020601f19601f840116840101905092915050565b604081525f61495a60408301868861491e565b828103602084015261496d81858761491e565b979650505050505050565b602081525f61430860208301848661491e565b60ff85168152836020820152606060408201525f6149ad60608301848661491e565b9695505050505050565b601f821115610f0057805f5260205f20601f840160051c810160208510156149dc5750805b601f840160051c820191505b818110156124fa575f81556001016149e8565b815167ffffffffffffffff811115614a1557614a156141b2565b614a2981614a2384546147ed565b846149b7565b6020601f821160018114614a7a575f8315614a445750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556124fa565b5f84815260208120601f198516915b82811015614aa95787850151825560209485019460019092019101614a89565b5084821015614ae557868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b606081525f614b0760608301888a61491e565b8281036020840152614b1a81878961491e565b90508281036040840152614b2f81858761491e565b9998505050505050505050565b838152604060208201525f613df460408301848661491e565b848152606060208201525f614b6e60608301858761491e565b9050821515604083015295945050505050565b67ffffffffffffffff831115614b9957614b996141b2565b614bad83614ba783546147ed565b836149b7565b5f601f841160018114614bfd575f8515614bc75750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b1783556124fa565b5f83815260208120601f198716915b82811015614c2c5786850135825560209485019460019092019101614c0c565b5086821015614c67577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b858152606060208201525f614c9260608301868861491e565b8281036040840152614ca581858761491e565b98975050505050505050565b5f60208284031215614cc1575f80fd5b5051919050565b6001600160a01b038a16815260c060208201525f614cea60c083018a8c61491e565b8860408401528281036060840152614d0381888a61491e565b90508281036080840152614d1881868861491e565b91505060ff831660a08301529a9950505050505050505050565b8581526001600160a01b0385166020820152608060408201525f614d5a60808301858761491e565b905082151560608301529695505050505050565b878152608060208201525f614d8760808301888a61491e565b8281036040840152614d9a81878961491e565b90508281036060840152614daf81858761491e565b9a9950505050505050505050565b818103818111156108f7577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f8060408385031215614e33575f80fd5b8251614e3e81614155565b6020939093015192949293505050565b5f60208284031215614e5e575f80fd5b815161128181613f70565b60c081525f614e7b60c083018861483e565b6001600160a01b03871660208401528281036040840152614e9c8187613fd8565b90508451606084015260208501511515608084015282810360a08401526001600160a01b0384511681526020840151151560208201526001600160a01b036040850151166040820152606084015160a06060830152614efe60a0830182613fd8565b9050608085015160808301528092505050969550505050505056fea2646970667358221220c8a2b276a66832369b6ff2442dc99da97c81d9f293da94a4ec18c9c59e3da29064736f6c634300081a003360a060405234801561000f575f80fd5b50737cce3eb018bf23e1fe2a32692f2c77592d1103946001600160a01b031663cc5ad8b66040518163ffffffff1660e01b81526004016020604051808303815f875af1158015610061573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100859190610096565b6001600160a01b03166080526100c3565b5f602082840312156100a6575f80fd5b81516001600160a01b03811681146100bc575f80fd5b9392505050565b60805161090e6100da5f395f607c015261090e5ff3fe608060405260043610610062575f3560e01c80637b1039991161003f5780637b10399914610105578063c9028a361461012c578063ebf9b2aa1461013f57005b8063116191b61461006b5780632d4cfb7e146100c75780635bcfd616146100e657005b3661006957005b005b348015610076575f80fd5b5061009e7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100d2575f80fd5b506100696100e13660046102af565b610152565b3480156100f1575f80fd5b50610069610100366004610370565b61018c565b348015610110575f80fd5b5061009e737cce3eb018bf23e1fe2a32692f2c77592d11039481565b61006961013a3660046103f4565b610207565b61006961014d36600461042b565b610236565b7f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db781604051610181919061053b565b60405180910390a150565b606081156101a3576101a08284018461063b565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e6101ce878061072b565b6101de60408a0160208b0161078c565b896040013533866040516101f7969594939291906107a5565b60405180910390a1505050505050565b7fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c481604051610181919061084f565b6060811561024d5761024a8284018461063b565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e610278858061072b565b610288604088016020890161078c565b876040013533866040516102a1969594939291906107a5565b60405180910390a150505050565b5f602082840312156102bf575f80fd5b813567ffffffffffffffff8111156102d5575f80fd5b820160c081850312156102e6575f80fd5b9392505050565b5f606082840312156102fd575f80fd5b50919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610326575f80fd5b919050565b5f8083601f84011261033b575f80fd5b50813567ffffffffffffffff811115610352575f80fd5b602083019150836020828501011115610369575f80fd5b9250929050565b5f805f805f60808688031215610384575f80fd5b853567ffffffffffffffff81111561039a575f80fd5b6103a6888289016102ed565b9550506103b560208701610303565b935060408601359250606086013567ffffffffffffffff8111156103d7575f80fd5b6103e38882890161032b565b969995985093965092949392505050565b5f60208284031215610404575f80fd5b813567ffffffffffffffff81111561041a575f80fd5b8201608081850312156102e6575f80fd5b5f805f6040848603121561043d575f80fd5b833567ffffffffffffffff811115610453575f80fd5b61045f868287016102ed565b935050602084013567ffffffffffffffff81111561047b575f80fd5b6104878682870161032b565b9497909650939450505050565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126104c7575f80fd5b830160208101925035905067ffffffffffffffff8111156104e6575f80fd5b803603821315610369575f80fd5b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b602081525f61054a8384610494565b60c0602085015261055f60e0850182846104f4565b91505073ffffffffffffffffffffffffffffffffffffffff61058360208601610303565b1660408401525f604085013590508060608501525060608401358015158082146105ab575f80fd5b80608086015250505f608085013590508060a0850152506105cf60a0850185610494565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c08601526106048382846104f4565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f6020828403121561064b575f80fd5b813567ffffffffffffffff811115610661575f80fd5b8201601f81018413610671575f80fd5b803567ffffffffffffffff81111561068b5761068b61060e565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156106f7576106f761060e565b60405281815282820160200186101561070e575f80fd5b816020840160208301375f91810160200191909152949350505050565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261075e575f80fd5b83018035915067ffffffffffffffff821115610778575f80fd5b602001915036819003821315610369575f80fd5b5f6020828403121561079c575f80fd5b6102e682610303565b60a081525f6107b860a08301888a6104f4565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff85166060840152828103608084015283518082528060208601602084015e5f6020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff61087183610303565b16602082015273ffffffffffffffffffffffffffffffffffffffff61089860208401610303565b1660408201525f8060408401359050806060840152506108bb6060840184610494565b6080808501526108cf60a0850182846104f4565b9594505050505056fea264697066735822122073b2b7273c8fa2cc0343725b06eef9af8f807b8eda5d1626b08ac364147ff7d464736f6c634300081a003360c060405234801561000f575f80fd5b5060405161102e38038061102e83398101604081905261002e916100d8565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006257604051632b2add3d60e01b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0385811691909117909155828116608052811660a0526040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac5905f90a1505050610118565b80516001600160a01b03811681146100d3575f80fd5b919050565b5f805f606084860312156100ea575f80fd5b6100f3846100bd565b9250610101602085016100bd565b915061010f604085016100bd565b90509250925092565b60805160a051610eee6101405f395f6101dd01525f81816102b001526104510152610eee5ff3fe608060405234801561000f575f80fd5b50600436106100f0575f3560e01c806397770dff11610093578063c63585cc11610063578063c63585cc1461026b578063d7fd7afb1461027e578063d936a012146102ab578063ee2815ba146102d2575f80fd5b806397770dff14610212578063a7cb050714610225578063c39aca3714610238578063c62178ac1461024b575f80fd5b8063513a9c05116100ce578063513a9c0514610183578063569541b9146101b8578063842da36d146101d857806391dd645f146101ff575f80fd5b80630be15547146100f45780631f0e251b146101535780633ce4a5bc14610168575b5f80fd5b610129610102366004610bb9565b60016020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b610166610161366004610bf8565b6102e5565b005b61012973735b14bb79463307aacbed86daf3322b1e6226ab81565b610129610191366004610bb9565b60026020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101299073ffffffffffffffffffffffffffffffffffffffff1681565b6101297f000000000000000000000000000000000000000000000000000000000000000081565b61016661020d366004610c18565b6103f9565b610166610220366004610bf8565b61051b565b610166610233366004610c42565b610628565b610166610246366004610c62565b6106c2565b6004546101299073ffffffffffffffffffffffffffffffffffffffff1681565b610129610279366004610d28565b6108b9565b61029d61028c366004610bb9565b5f6020819052908152604090205481565b60405190815260200161014a565b6101297f000000000000000000000000000000000000000000000000000000000000000081565b6101666102e0366004610c18565b6109ec565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610332576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811661037f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610446576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003545f9061048d907f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff16846108b9565b5f8481526002602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251878152918201529192507f0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e910160405180910390a1505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610568576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166105b5576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e906020016103ee565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610675576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461070f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831673735b14bb79463307aacbed86daf3322b1e6226ab148061075c575073ffffffffffffffffffffffffffffffffffffffff831630145b15610793576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690528616906347e7ef24906044016020604051808303815f875af1158015610805573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108299190610d68565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063de43156e906108849089908990899088908890600401610dce565b5f604051808303815f87803b15801561089b575f80fd5b505af11580156108ad573d5f803e3d5ffd5b50505050505050505050565b5f805f6108c68585610abc565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b16603482015291935091508690604801604051602081830303815290604052805190602001206040516020016109ac9291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a39576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d91016106b6565b5f808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610b23576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610b5d578284610b60565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610bb2576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b5f60208284031215610bc9575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610bf3575f80fd5b919050565b5f60208284031215610c08575f80fd5b610c1182610bd0565b9392505050565b5f8060408385031215610c29575f80fd5b82359150610c3960208401610bd0565b90509250929050565b5f8060408385031215610c53575f80fd5b50508035926020909101359150565b5f805f805f8060a08789031215610c77575f80fd5b863567ffffffffffffffff811115610c8d575f80fd5b87016060818a031215610c9e575f80fd5b9550610cac60208801610bd0565b945060408701359350610cc160608801610bd0565b9250608087013567ffffffffffffffff811115610cdc575f80fd5b8701601f81018913610cec575f80fd5b803567ffffffffffffffff811115610d02575f80fd5b896020828401011115610d13575f80fd5b60208201935080925050509295509295509295565b5f805f60608486031215610d3a575f80fd5b610d4384610bd0565b9250610d5160208501610bd0565b9150610d5f60408501610bd0565b90509250925092565b5f60208284031215610d78575f80fd5b81518015158114610c11575f80fd5b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b608081525f86357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018112610e04575f80fd5b870160208101903567ffffffffffffffff811115610e20575f80fd5b803603821315610e2e575f80fd5b60606080850152610e4360e085018284610d87565b91505073ffffffffffffffffffffffffffffffffffffffff610e6760208a01610bd0565b1660a0840152604088013560c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401528281036060840152610eac818587610d87565b9897505050505050505056fea2646970667358221220ab90e60a29637a57eaea52c9d20d7377af3635b9a7cb793d7eb87995722b46ac64736f6c634300081a003360c060405234801561000f575f80fd5b50604051611fc0380380611fc083398101604081905261002e916101d0565b6001600160a01b038216158061004b57506001600160a01b038116155b156100695760405163d92e233d60e01b815260040160405180910390fd5b60066100758982610315565b5060076100828882610315565b506008805460ff191660ff881617905560808590528360028111156100a9576100a96103cf565b60a08160028111156100bd576100bd6103cf565b9052506001929092555f80546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506103e39350505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261012d575f80fd5b81516001600160401b038111156101465761014661010a565b604051601f8201601f19908116603f011681016001600160401b03811182821017156101745761017461010a565b60405281815283820160200185101561018b575f80fd5b8160208501602083015e5f918101602001919091529392505050565b8051600381106101b5575f80fd5b919050565b80516001600160a01b03811681146101b5575f80fd5b5f805f805f805f80610100898b0312156101e8575f80fd5b88516001600160401b038111156101fd575f80fd5b6102098b828c0161011e565b60208b015190995090506001600160401b03811115610226575f80fd5b6102328b828c0161011e565b975050604089015160ff81168114610248575f80fd5b60608a0151909650945061025e60808a016101a7565b60a08a0151909450925061027460c08a016101ba565b915061028260e08a016101ba565b90509295985092959890939650565b600181811c908216806102a557607f821691505b6020821081036102c357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561031057805f5260205f20601f840160051c810160208510156102ee5750805b601f840160051c820191505b8181101561030d575f81556001016102fa565b50505b505050565b81516001600160401b0381111561032e5761032e61010a565b6103428161033c8454610291565b846102c9565b6020601f821160018114610374575f831561035d5750848201515b5f19600385901b1c1916600184901b17845561030d565b5f84815260208120601f198516915b828110156103a35787850151825560209485019460019092019101610383565b50848210156103c057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52602160045260245ffd5b60805160a051611ba06104205f395f61033901525f81816102e501528181610bbf01528181610cc201528181610ed90152610fdc0152611ba05ff3fe608060405234801561000f575f80fd5b50600436106101b0575f3560e01c806395d89b41116100f3578063ccc7759911610093578063eddeb1231161006e578063eddeb12314610455578063f2441b3214610468578063f687d12a14610487578063fc5fecd51461049a575f80fd5b8063ccc77599146103c9578063d9eeebed146103dc578063dd62ed3e14610410575f80fd5b8063b84c8246116100ce578063b84c82461461037b578063c47f002714610390578063c7012626146103a3578063c835d7cc146103b6575f80fd5b806395d89b411461032c578063a3413d0314610334578063a9059cbb14610368575f80fd5b80633ce4a5bc1161015e5780634d8943bb116101395780634d8943bb146102a257806370a08231146102ab57806385e1f4d0146102e05780638b851b9514610307575f80fd5b80633ce4a5bc1461023c57806342966c681461027c57806347e7ef241461028f575f80fd5b806318160ddd1161018e57806318160ddd1461020c57806323b872dd14610214578063313ce56714610227575f80fd5b806306fdde03146101b4578063091d2788146101d2578063095ea7b3146101e9575b5f80fd5b6101bc6104ad565b6040516101c991906115fb565b60405180910390f35b6101db60015481565b6040519081526020016101c9565b6101fc6101f7366004611638565b61053d565b60405190151581526020016101c9565b6005546101db565b6101fc610222366004611662565b610553565b60085460405160ff90911681526020016101c9565b61025773735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c9565b6101fc61028a3660046116a0565b6105e8565b6101fc61029d366004611638565b6105fb565b6101db60025481565b6101db6102b93660046116b7565b73ffffffffffffffffffffffffffffffffffffffff165f9081526003602052604090205490565b6101db7f000000000000000000000000000000000000000000000000000000000000000081565b60085461025790610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101bc610752565b61035b7f000000000000000000000000000000000000000000000000000000000000000081565b6040516101c991906116d2565b6101fc610376366004611638565b610761565b61038e6103893660046117d3565b61076d565b005b61038e61039e3660046117d3565b6107ca565b6101fc6103b1366004611820565b610823565b61038e6103c43660046116b7565b61096d565b61038e6103d73660046116b7565b610a80565b6103e4610b94565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101c9565b6101db61041e366004611875565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260046020908152604080832093909416825291909152205490565b61038e6104633660046116a0565b610daa565b5f546102579073ffffffffffffffffffffffffffffffffffffffff1681565b61038e6104953660046116a0565b610e2c565b6103e46104a83660046116a0565b610eae565b6060600680546104bc906118ac565b80601f01602080910402602001604051908101604052809291908181526020018280546104e8906118ac565b80156105335780601f1061050a57610100808354040283529160200191610533565b820191905f5260205f20905b81548152906001019060200180831161051657829003601f168201915b5050505050905090565b5f6105493384846110c2565b5060015b92915050565b5f61055f8484846111ca565b73ffffffffffffffffffffffffffffffffffffffff84165f908152600460209081526040808320338452909152902054828110156105c9576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105dd85336105d8868561192a565b6110c2565b506001949350505050565b5f6105f33383611383565b506001919050565b5f3373735b14bb79463307aacbed86daf3322b1e6226ab1480159061063757505f5473ffffffffffffffffffffffffffffffffffffffff163314155b80156106605750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b15610697576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106a183836114c2565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261074191869061193d565b60405180910390a250600192915050565b6060600780546104bc906118ac565b5f6105493384846111ca565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107ba576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076107c682826119aa565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610817576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60066107c682826119aa565b5f805f61082e610b94565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303815f875af11580156108bd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108e19190611ac1565b610917576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109213385611383565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161095a91899189918791611ae0565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109ba576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a07576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610acd576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b1a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610a75565b5f80546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610c24573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c489190611b0e565b905073ffffffffffffffffffffffffffffffffffffffff8116610c97576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610d23573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d479190611b29565b9050805f03610d82576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60025460015483610d949190611b40565b610d9e9190611b57565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610df7576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610a75565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e79576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610a75565b5f80546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610f3e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f629190611b0e565b905073ffffffffffffffffffffffffffffffffffffffff8116610fb1576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa15801561103d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110619190611b29565b9050805f0361109c576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002545f906110ab8784611b40565b6110b59190611b57565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff831661110f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821661115c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8381165f8181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611217576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611264576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f90815260036020526040902054818110156112c3576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112cd828261192a565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260036020526040808220939093559085168152908120805484929061130f908490611b57565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161137591815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166113d0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f908152600360205260409020548181101561142f576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611439828261192a565b73ffffffffffffffffffffffffffffffffffffffff84165f908152600360205260408120919091556005805484929061147390849061192a565b90915550506040518281525f9073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016111bd565b73ffffffffffffffffffffffffffffffffffffffff821661150f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060055f8282546115209190611b57565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f9081526003602052604081208054839290611559908490611b57565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316905f907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61160d60208301846115af565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff81168114611635575f80fd5b50565b5f8060408385031215611649575f80fd5b823561165481611614565b946020939093013593505050565b5f805f60608486031215611674575f80fd5b833561167f81611614565b9250602084013561168f81611614565b929592945050506040919091013590565b5f602082840312156116b0575f80fd5b5035919050565b5f602082840312156116c7575f80fd5b813561160d81611614565b602081016003831061170b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b91905290565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8067ffffffffffffffff84111561175857611758611711565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156117a5576117a5611711565b6040528381529050808284018510156117bc575f80fd5b838360208301375f60208583010152509392505050565b5f602082840312156117e3575f80fd5b813567ffffffffffffffff8111156117f9575f80fd5b8201601f81018413611809575f80fd5b6118188482356020840161173e565b949350505050565b5f8060408385031215611831575f80fd5b823567ffffffffffffffff811115611847575f80fd5b8301601f81018513611857575f80fd5b6118668582356020840161173e565b95602094909401359450505050565b5f8060408385031215611886575f80fd5b823561189181611614565b915060208301356118a181611614565b809150509250929050565b600181811c908216806118c057607f821691505b6020821081036118f7577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561054d5761054d6118fd565b604081525f61194f60408301856115af565b90508260208301529392505050565b601f8211156119a557805f5260205f20601f840160051c810160208510156119835750805b601f840160051c820191505b818110156119a2575f815560010161198f565b50505b505050565b815167ffffffffffffffff8111156119c4576119c4611711565b6119d8816119d284546118ac565b8461195e565b6020601f821160018114611a29575f83156119f35750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556119a2565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611a765787850151825560209485019460019092019101611a56565b5084821015611ab257868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b5f60208284031215611ad1575f80fd5b8151801515811461160d575f80fd5b608081525f611af260808301876115af565b6020830195909552506040810192909252606090910152919050565b5f60208284031215611b1e575f80fd5b815161160d81611614565b5f60208284031215611b39575f80fd5b5051919050565b808202811582820484141761054d5761054d6118fd565b8082018082111561054d5761054d6118fd56fea26469706673582212206130fb621d6dd12e8164e45fefe5a69512cd6f3f9ac39ab60f6f9d44bc45140c64736f6c634300081a003347617320666565206e6f74206275726e656420636f72726563746c792066726f6d20676173205a524332305a5243323020746f6b656e732077657265206e6f74206275726e656420636f72726563746c795769746864726177616c20616d6f756e74206e6f74206275726e656420636f72726563746c792066726f6d207365636f6e64205a524332306465706f7369742073686f756c64206661696c2064756520746f20696e73756666696369656e7420455448206f6e2070726f746f636f6c20454f415a5243323020746f6b656e732077657265206e6f74206275726e656420636f72726563746c7920666f72207769746864726177416e6443616c6ca2646970667358221220f57ee5f865f783972d2d445a598b354a753c2a7b0c49bd2166564ed54ef30cc164736f6c634300081a0033",
}

// GatewayZEVMOutboundTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayZEVMOutboundTestMetaData.ABI instead.
var GatewayZEVMOutboundTestABI = GatewayZEVMOutboundTestMetaData.ABI

// GatewayZEVMOutboundTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayZEVMOutboundTestMetaData.Bin instead.
var GatewayZEVMOutboundTestBin = GatewayZEVMOutboundTestMetaData.Bin

// DeployGatewayZEVMOutboundTest deploys a new Ethereum contract, binding an instance of GatewayZEVMOutboundTest to it.
func DeployGatewayZEVMOutboundTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayZEVMOutboundTest, error) {
	parsed, err := GatewayZEVMOutboundTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayZEVMOutboundTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayZEVMOutboundTest{GatewayZEVMOutboundTestCaller: GatewayZEVMOutboundTestCaller{contract: contract}, GatewayZEVMOutboundTestTransactor: GatewayZEVMOutboundTestTransactor{contract: contract}, GatewayZEVMOutboundTestFilterer: GatewayZEVMOutboundTestFilterer{contract: contract}}, nil
}

// GatewayZEVMOutboundTest is an auto generated Go binding around an Ethereum contract.
type GatewayZEVMOutboundTest struct {
	GatewayZEVMOutboundTestCaller     // Read-only binding to the contract
	GatewayZEVMOutboundTestTransactor // Write-only binding to the contract
	GatewayZEVMOutboundTestFilterer   // Log filterer for contract events
}

// GatewayZEVMOutboundTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMOutboundTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMOutboundTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayZEVMOutboundTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMOutboundTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayZEVMOutboundTestSession struct {
	Contract     *GatewayZEVMOutboundTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// GatewayZEVMOutboundTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayZEVMOutboundTestCallerSession struct {
	Contract *GatewayZEVMOutboundTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// GatewayZEVMOutboundTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayZEVMOutboundTestTransactorSession struct {
	Contract     *GatewayZEVMOutboundTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// GatewayZEVMOutboundTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestRaw struct {
	Contract *GatewayZEVMOutboundTest // Generic contract binding to access the raw methods on
}

// GatewayZEVMOutboundTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestCallerRaw struct {
	Contract *GatewayZEVMOutboundTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayZEVMOutboundTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayZEVMOutboundTestTransactorRaw struct {
	Contract *GatewayZEVMOutboundTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayZEVMOutboundTest creates a new instance of GatewayZEVMOutboundTest, bound to a specific deployed contract.
func NewGatewayZEVMOutboundTest(address common.Address, backend bind.ContractBackend) (*GatewayZEVMOutboundTest, error) {
	contract, err := bindGatewayZEVMOutboundTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTest{GatewayZEVMOutboundTestCaller: GatewayZEVMOutboundTestCaller{contract: contract}, GatewayZEVMOutboundTestTransactor: GatewayZEVMOutboundTestTransactor{contract: contract}, GatewayZEVMOutboundTestFilterer: GatewayZEVMOutboundTestFilterer{contract: contract}}, nil
}

// NewGatewayZEVMOutboundTestCaller creates a new read-only instance of GatewayZEVMOutboundTest, bound to a specific deployed contract.
func NewGatewayZEVMOutboundTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayZEVMOutboundTestCaller, error) {
	contract, err := bindGatewayZEVMOutboundTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestCaller{contract: contract}, nil
}

// NewGatewayZEVMOutboundTestTransactor creates a new write-only instance of GatewayZEVMOutboundTest, bound to a specific deployed contract.
func NewGatewayZEVMOutboundTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayZEVMOutboundTestTransactor, error) {
	contract, err := bindGatewayZEVMOutboundTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestTransactor{contract: contract}, nil
}

// NewGatewayZEVMOutboundTestFilterer creates a new log filterer instance of GatewayZEVMOutboundTest, bound to a specific deployed contract.
func NewGatewayZEVMOutboundTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayZEVMOutboundTestFilterer, error) {
	contract, err := bindGatewayZEVMOutboundTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestFilterer{contract: contract}, nil
}

// bindGatewayZEVMOutboundTest binds a generic wrapper to an already deployed contract.
func bindGatewayZEVMOutboundTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayZEVMOutboundTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMOutboundTest.Contract.GatewayZEVMOutboundTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.GatewayZEVMOutboundTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.GatewayZEVMOutboundTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMOutboundTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.contract.Transact(opts, method, params...)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ISTEST() (bool, error) {
	return _GatewayZEVMOutboundTest.Contract.ISTEST(&_GatewayZEVMOutboundTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ISTEST() (bool, error) {
	return _GatewayZEVMOutboundTest.Contract.ISTEST(&_GatewayZEVMOutboundTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayZEVMOutboundTest.Contract.PAUSERROLE(&_GatewayZEVMOutboundTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayZEVMOutboundTest.Contract.PAUSERROLE(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeArtifacts(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeArtifacts(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeContracts(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeContracts(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeSenders(&_GatewayZEVMOutboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.ExcludeSenders(&_GatewayZEVMOutboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) Failed() (bool, error) {
	return _GatewayZEVMOutboundTest.Contract.Failed(&_GatewayZEVMOutboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) Failed() (bool, error) {
	return _GatewayZEVMOutboundTest.Contract.Failed(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetArtifactSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetArtifactSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetArtifacts(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetArtifacts(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetContracts(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetContracts(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetInterfaces(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetInterfaces(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetSelectors(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMOutboundTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetSenders(&_GatewayZEVMOutboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayZEVMOutboundTest.Contract.TargetSenders(&_GatewayZEVMOutboundTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.SetUp(&_GatewayZEVMOutboundTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.SetUp(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForDifferentZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x75ca5587.
//
// Solidity: function testBurnGasFeeForDifferentZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnGasFeeForDifferentZRC20WithdrawAndCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnGasFeeForDifferentZRC20WithdrawAndCall")
}

// TestBurnGasFeeForDifferentZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x75ca5587.
//
// Solidity: function testBurnGasFeeForDifferentZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnGasFeeForDifferentZRC20WithdrawAndCall() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForDifferentZRC20WithdrawAndCall(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForDifferentZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x75ca5587.
//
// Solidity: function testBurnGasFeeForDifferentZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnGasFeeForDifferentZRC20WithdrawAndCall() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForDifferentZRC20WithdrawAndCall(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForDifferentZRC20Withdrawal is a paid mutator transaction binding the contract method 0xb84a3d2f.
//
// Solidity: function testBurnGasFeeForDifferentZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnGasFeeForDifferentZRC20Withdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnGasFeeForDifferentZRC20Withdrawal")
}

// TestBurnGasFeeForDifferentZRC20Withdrawal is a paid mutator transaction binding the contract method 0xb84a3d2f.
//
// Solidity: function testBurnGasFeeForDifferentZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnGasFeeForDifferentZRC20Withdrawal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForDifferentZRC20Withdrawal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForDifferentZRC20Withdrawal is a paid mutator transaction binding the contract method 0xb84a3d2f.
//
// Solidity: function testBurnGasFeeForDifferentZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnGasFeeForDifferentZRC20Withdrawal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForDifferentZRC20Withdrawal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x3177f381.
//
// Solidity: function testBurnGasFeeForZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnGasFeeForZRC20WithdrawAndCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnGasFeeForZRC20WithdrawAndCall")
}

// TestBurnGasFeeForZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x3177f381.
//
// Solidity: function testBurnGasFeeForZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnGasFeeForZRC20WithdrawAndCall() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForZRC20WithdrawAndCall(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForZRC20WithdrawAndCall is a paid mutator transaction binding the contract method 0x3177f381.
//
// Solidity: function testBurnGasFeeForZRC20WithdrawAndCall() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnGasFeeForZRC20WithdrawAndCall() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForZRC20WithdrawAndCall(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForZRC20Withdrawal is a paid mutator transaction binding the contract method 0xf2a91358.
//
// Solidity: function testBurnGasFeeForZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnGasFeeForZRC20Withdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnGasFeeForZRC20Withdrawal")
}

// TestBurnGasFeeForZRC20Withdrawal is a paid mutator transaction binding the contract method 0xf2a91358.
//
// Solidity: function testBurnGasFeeForZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnGasFeeForZRC20Withdrawal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForZRC20Withdrawal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnGasFeeForZRC20Withdrawal is a paid mutator transaction binding the contract method 0xf2a91358.
//
// Solidity: function testBurnGasFeeForZRC20Withdrawal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnGasFeeForZRC20Withdrawal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnGasFeeForZRC20Withdrawal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnProtocolFeesFailsWithInsufficientAllowance is a paid mutator transaction binding the contract method 0x61cfddb7.
//
// Solidity: function testBurnProtocolFeesFailsWithInsufficientAllowance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestBurnProtocolFeesFailsWithInsufficientAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testBurnProtocolFeesFailsWithInsufficientAllowance")
}

// TestBurnProtocolFeesFailsWithInsufficientAllowance is a paid mutator transaction binding the contract method 0x61cfddb7.
//
// Solidity: function testBurnProtocolFeesFailsWithInsufficientAllowance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestBurnProtocolFeesFailsWithInsufficientAllowance() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnProtocolFeesFailsWithInsufficientAllowance(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestBurnProtocolFeesFailsWithInsufficientAllowance is a paid mutator transaction binding the contract method 0x61cfddb7.
//
// Solidity: function testBurnProtocolFeesFailsWithInsufficientAllowance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestBurnProtocolFeesFailsWithInsufficientAllowance() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestBurnProtocolFeesFailsWithInsufficientAllowance(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDeposit is a paid mutator transaction binding the contract method 0x7f924c4e.
//
// Solidity: function testDeposit() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDeposit")
}

// TestDeposit is a paid mutator transaction binding the contract method 0x7f924c4e.
//
// Solidity: function testDeposit() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDeposit() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDeposit(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDeposit is a paid mutator transaction binding the contract method 0x7f924c4e.
//
// Solidity: function testDeposit() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDeposit() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDeposit(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x884660a3.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContract")
}

// TestDepositAndRevertZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x884660a3.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x884660a3.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway is a paid mutator transaction binding the contract method 0x51336fb0.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway is a paid mutator transaction binding the contract method 0x51336fb0.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway is a paid mutator transaction binding the contract method 0x51336fb0.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol is a paid mutator transaction binding the contract method 0x6163f8ef.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol is a paid mutator transaction binding the contract method 0x6163f8ef.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol is a paid mutator transaction binding the contract method 0x6163f8ef.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xeab7674e.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xeab7674e.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xeab7674e.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xf1d98f1b.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xf1d98f1b.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xf1d98f1b.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x339bd828.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x339bd828.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x339bd828.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x48f4fd07.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress")
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x48f4fd07.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x48f4fd07.
//
// Solidity: function testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x996b7675.
//
// Solidity: function testDepositFailsIfAmountIs0() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfAmountIs0")
}

// TestDepositFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x996b7675.
//
// Solidity: function testDepositFailsIfAmountIs0() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfAmountIs0(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x996b7675.
//
// Solidity: function testDepositFailsIfAmountIs0() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfAmountIs0(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfSenderNotProtocol is a paid mutator transaction binding the contract method 0xca26929c.
//
// Solidity: function testDepositFailsIfSenderNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfSenderNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfSenderNotProtocol")
}

// TestDepositFailsIfSenderNotProtocol is a paid mutator transaction binding the contract method 0xca26929c.
//
// Solidity: function testDepositFailsIfSenderNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfSenderNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfSenderNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfSenderNotProtocol is a paid mutator transaction binding the contract method 0xca26929c.
//
// Solidity: function testDepositFailsIfSenderNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfSenderNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfSenderNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x96d9d876.
//
// Solidity: function testDepositFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfTargetIsGateway")
}

// TestDepositFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x96d9d876.
//
// Solidity: function testDepositFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x96d9d876.
//
// Solidity: function testDepositFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x3626c616.
//
// Solidity: function testDepositFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfTargetIsProtocol")
}

// TestDepositFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x3626c616.
//
// Solidity: function testDepositFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x3626c616.
//
// Solidity: function testDepositFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x14b7a6da.
//
// Solidity: function testDepositFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfTargetIsZeroAddress")
}

// TestDepositFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x14b7a6da.
//
// Solidity: function testDepositFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x14b7a6da.
//
// Solidity: function testDepositFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x9c9acd5d.
//
// Solidity: function testDepositFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositFailsIfZRC20IsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositFailsIfZRC20IsZeroAddress")
}

// TestDepositFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x9c9acd5d.
//
// Solidity: function testDepositFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x9c9acd5d.
//
// Solidity: function testDepositFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositTogglePause is a paid mutator transaction binding the contract method 0x1c785a14.
//
// Solidity: function testDepositTogglePause() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositTogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositTogglePause")
}

// TestDepositTogglePause is a paid mutator transaction binding the contract method 0x1c785a14.
//
// Solidity: function testDepositTogglePause() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositTogglePause() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositTogglePause(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositTogglePause is a paid mutator transaction binding the contract method 0x1c785a14.
//
// Solidity: function testDepositTogglePause() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositTogglePause() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositTogglePause(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETA is a paid mutator transaction binding the contract method 0x445bbf47.
//
// Solidity: function testDepositZETA() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETA")
}

// TestDepositZETA is a paid mutator transaction binding the contract method 0x445bbf47.
//
// Solidity: function testDepositZETA() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETA() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETA(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETA is a paid mutator transaction binding the contract method 0x445bbf47.
//
// Solidity: function testDepositZETA() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETA() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETA(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversal is a paid mutator transaction binding the contract method 0x828d267c.
//
// Solidity: function testDepositZETAAndCallUniversal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversal")
}

// TestDepositZETAAndCallUniversal is a paid mutator transaction binding the contract method 0x828d267c.
//
// Solidity: function testDepositZETAAndCallUniversal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversal is a paid mutator transaction binding the contract method 0x828d267c.
//
// Solidity: function testDepositZETAAndCallUniversal() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversal() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversal(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContract is a paid mutator transaction binding the contract method 0x2468bc0f.
//
// Solidity: function testDepositZETAAndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContract")
}

// TestDepositZETAAndCallUniversalContract is a paid mutator transaction binding the contract method 0x2468bc0f.
//
// Solidity: function testDepositZETAAndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContract is a paid mutator transaction binding the contract method 0x2468bc0f.
//
// Solidity: function testDepositZETAAndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xc35cb5e4.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol")
}

// TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xc35cb5e4.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xc35cb5e4.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero is a paid mutator transaction binding the contract method 0xec294d9f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero")
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero is a paid mutator transaction binding the contract method 0xec294d9f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero is a paid mutator transaction binding the contract method 0xec294d9f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x5cec7db5.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway")
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x5cec7db5.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x5cec7db5.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x890a2d67.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol")
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x890a2d67.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x890a2d67.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x97f7661f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress")
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x97f7661f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x97f7661f.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0xb936be8c.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndCallUniversalContractFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndCallUniversalContractFailsIfZeroAddress")
}

// TestDepositZETAAndCallUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0xb936be8c.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndCallUniversalContractFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndCallUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0xb936be8c.
//
// Solidity: function testDepositZETAAndCallUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndCallUniversalContractFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndCallUniversalContractFailsIfZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContract is a paid mutator transaction binding the contract method 0x0c4a14b4.
//
// Solidity: function testDepositZETAAndRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContract")
}

// TestDepositZETAAndRevertUniversalContract is a paid mutator transaction binding the contract method 0x0c4a14b4.
//
// Solidity: function testDepositZETAAndRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContract is a paid mutator transaction binding the contract method 0x0c4a14b4.
//
// Solidity: function testDepositZETAAndRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xda8fd81d.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero")
}

// TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xda8fd81d.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xda8fd81d.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xe36c0a23.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol")
}

// TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xe36c0a23.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xe36c0a23.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x514fcc7f.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway")
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x514fcc7f.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0x514fcc7f.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x50dd38ca.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol")
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x50dd38ca.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x50dd38ca.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x2ed7cff6.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress")
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x2ed7cff6.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x2ed7cff6.
//
// Solidity: function testDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractWhenPaused is a paid mutator transaction binding the contract method 0x62fe2915.
//
// Solidity: function testDepositZETAAndRevertUniversalContractWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAAndRevertUniversalContractWhenPaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAAndRevertUniversalContractWhenPaused")
}

// TestDepositZETAAndRevertUniversalContractWhenPaused is a paid mutator transaction binding the contract method 0x62fe2915.
//
// Solidity: function testDepositZETAAndRevertUniversalContractWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAAndRevertUniversalContractWhenPaused() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractWhenPaused(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAAndRevertUniversalContractWhenPaused is a paid mutator transaction binding the contract method 0x62fe2915.
//
// Solidity: function testDepositZETAAndRevertUniversalContractWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAAndRevertUniversalContractWhenPaused() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAAndRevertUniversalContractWhenPaused(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x195ac143.
//
// Solidity: function testDepositZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfAmountIsZero")
}

// TestDepositZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x195ac143.
//
// Solidity: function testDepositZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x195ac143.
//
// Solidity: function testDepositZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfInsufficientProtocolBalance is a paid mutator transaction binding the contract method 0xf589655d.
//
// Solidity: function testDepositZETAFailsIfInsufficientProtocolBalance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfInsufficientProtocolBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfInsufficientProtocolBalance")
}

// TestDepositZETAFailsIfInsufficientProtocolBalance is a paid mutator transaction binding the contract method 0xf589655d.
//
// Solidity: function testDepositZETAFailsIfInsufficientProtocolBalance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfInsufficientProtocolBalance() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfInsufficientProtocolBalance(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfInsufficientProtocolBalance is a paid mutator transaction binding the contract method 0xf589655d.
//
// Solidity: function testDepositZETAFailsIfInsufficientProtocolBalance() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfInsufficientProtocolBalance() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfInsufficientProtocolBalance(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xae4662ec.
//
// Solidity: function testDepositZETAFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfSenderIsNotProtocol")
}

// TestDepositZETAFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xae4662ec.
//
// Solidity: function testDepositZETAFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xae4662ec.
//
// Solidity: function testDepositZETAFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0xb82ce921.
//
// Solidity: function testDepositZETAFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfTargetIsGateway")
}

// TestDepositZETAFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0xb82ce921.
//
// Solidity: function testDepositZETAFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsGateway is a paid mutator transaction binding the contract method 0xb82ce921.
//
// Solidity: function testDepositZETAFailsIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x21d92591.
//
// Solidity: function testDepositZETAFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfTargetIsProtocol")
}

// TestDepositZETAFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x21d92591.
//
// Solidity: function testDepositZETAFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsProtocol is a paid mutator transaction binding the contract method 0x21d92591.
//
// Solidity: function testDepositZETAFailsIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x1ff5bdd7.
//
// Solidity: function testDepositZETAFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsIfTargetIsZeroAddress")
}

// TestDepositZETAFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x1ff5bdd7.
//
// Solidity: function testDepositZETAFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x1ff5bdd7.
//
// Solidity: function testDepositZETAFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsWhenPaused is a paid mutator transaction binding the contract method 0x33f7c81a.
//
// Solidity: function testDepositZETAFailsWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZETAFailsWhenPaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZETAFailsWhenPaused")
}

// TestDepositZETAFailsWhenPaused is a paid mutator transaction binding the contract method 0x33f7c81a.
//
// Solidity: function testDepositZETAFailsWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZETAFailsWhenPaused() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsWhenPaused(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZETAFailsWhenPaused is a paid mutator transaction binding the contract method 0x33f7c81a.
//
// Solidity: function testDepositZETAFailsWhenPaused() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZETAFailsWhenPaused() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZETAFailsWhenPaused(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x6efa04b5.
//
// Solidity: function testDepositZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContract")
}

// TestDepositZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x6efa04b5.
//
// Solidity: function testDepositZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContract is a paid mutator transaction binding the contract method 0x6efa04b5.
//
// Solidity: function testDepositZRC20AndCallUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x58c9987f.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero")
}

// TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x58c9987f.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x58c9987f.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfAmountIsZero(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xcf2c3d1d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol")
}

// TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xcf2c3d1d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0xcf2c3d1d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xef2b5394.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress")
}

// TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xef2b5394.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xef2b5394.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0xfb339a1c.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress")
}

// TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0xfb339a1c.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0xfb339a1c.
//
// Solidity: function testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsGateway is a paid mutator transaction binding the contract method 0xe09bc659.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractIfTargetIsGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractIfTargetIsGateway")
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsGateway is a paid mutator transaction binding the contract method 0xe09bc659.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsGateway is a paid mutator transaction binding the contract method 0xe09bc659.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsGateway() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractIfTargetIsGateway() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractIfTargetIsGateway(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol is a paid mutator transaction binding the contract method 0xeb78bd7d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testDepositZRC20AndCallUniversalContractIfTargetIsProtocol")
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol is a paid mutator transaction binding the contract method 0xeb78bd7d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol is a paid mutator transaction binding the contract method 0xeb78bd7d.
//
// Solidity: function testDepositZRC20AndCallUniversalContractIfTargetIsProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestDepositZRC20AndCallUniversalContractIfTargetIsProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteAbortUniversalContract is a paid mutator transaction binding the contract method 0x1832cb6e.
//
// Solidity: function testExecuteAbortUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteAbortUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteAbortUniversalContract")
}

// TestExecuteAbortUniversalContract is a paid mutator transaction binding the contract method 0x1832cb6e.
//
// Solidity: function testExecuteAbortUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteAbortUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteAbortUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteAbortUniversalContract is a paid mutator transaction binding the contract method 0x1832cb6e.
//
// Solidity: function testExecuteAbortUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteAbortUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteAbortUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xee0f4ea1.
//
// Solidity: function testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress")
}

// TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xee0f4ea1.
//
// Solidity: function testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xee0f4ea1.
//
// Solidity: function testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteAbortUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x3ab5b199.
//
// Solidity: function testExecuteFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteFailsIfTargetIsZeroAddress")
}

// TestExecuteFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x3ab5b199.
//
// Solidity: function testExecuteFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0x3ab5b199.
//
// Solidity: function testExecuteFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x2948df41.
//
// Solidity: function testExecuteFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteFailsIfZRC20IsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteFailsIfZRC20IsZeroAddress")
}

// TestExecuteFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x2948df41.
//
// Solidity: function testExecuteFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteFailsIfZRC20IsZeroAddress is a paid mutator transaction binding the contract method 0x2948df41.
//
// Solidity: function testExecuteFailsIfZRC20IsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteFailsIfZRC20IsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteFailsIfZRC20IsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContract is a paid mutator transaction binding the contract method 0x084fafab.
//
// Solidity: function testExecuteRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteRevertUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteRevertUniversalContract")
}

// TestExecuteRevertUniversalContract is a paid mutator transaction binding the contract method 0x084fafab.
//
// Solidity: function testExecuteRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteRevertUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContract is a paid mutator transaction binding the contract method 0x084fafab.
//
// Solidity: function testExecuteRevertUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteRevertUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xc8814d2e.
//
// Solidity: function testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress")
}

// TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xc8814d2e.
//
// Solidity: function testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress is a paid mutator transaction binding the contract method 0xc8814d2e.
//
// Solidity: function testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContractFailsIfTargetIsZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContractIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x27820625.
//
// Solidity: function testExecuteRevertUniversalContractIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteRevertUniversalContractIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteRevertUniversalContractIfSenderIsNotProtocol")
}

// TestExecuteRevertUniversalContractIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x27820625.
//
// Solidity: function testExecuteRevertUniversalContractIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteRevertUniversalContractIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContractIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteRevertUniversalContractIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x27820625.
//
// Solidity: function testExecuteRevertUniversalContractIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteRevertUniversalContractIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteRevertUniversalContractIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContract is a paid mutator transaction binding the contract method 0x7cec29b0.
//
// Solidity: function testExecuteUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteUniversalContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteUniversalContract")
}

// TestExecuteUniversalContract is a paid mutator transaction binding the contract method 0x7cec29b0.
//
// Solidity: function testExecuteUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContract is a paid mutator transaction binding the contract method 0x7cec29b0.
//
// Solidity: function testExecuteUniversalContract() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteUniversalContract() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContract(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x2acb21b4.
//
// Solidity: function testExecuteUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteUniversalContractFailsIfSenderIsNotProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteUniversalContractFailsIfSenderIsNotProtocol")
}

// TestExecuteUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x2acb21b4.
//
// Solidity: function testExecuteUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContractFailsIfSenderIsNotProtocol is a paid mutator transaction binding the contract method 0x2acb21b4.
//
// Solidity: function testExecuteUniversalContractFailsIfSenderIsNotProtocol() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteUniversalContractFailsIfSenderIsNotProtocol() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContractFailsIfSenderIsNotProtocol(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x5b4c90e1.
//
// Solidity: function testExecuteUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactor) TestExecuteUniversalContractFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.contract.Transact(opts, "testExecuteUniversalContractFailsIfZeroAddress")
}

// TestExecuteUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x5b4c90e1.
//
// Solidity: function testExecuteUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestSession) TestExecuteUniversalContractFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContractFailsIfZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// TestExecuteUniversalContractFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x5b4c90e1.
//
// Solidity: function testExecuteUniversalContractFailsIfZeroAddress() returns()
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestTransactorSession) TestExecuteUniversalContractFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMOutboundTest.Contract.TestExecuteUniversalContractFailsIfZeroAddress(&_GatewayZEVMOutboundTest.TransactOpts)
}

// GatewayZEVMOutboundTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestCalledIterator struct {
	Event *GatewayZEVMOutboundTestCalled // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestCalled)
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
		it.Event = new(GatewayZEVMOutboundTestCalled)
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
func (it *GatewayZEVMOutboundTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestCalled represents a Called event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestCalled struct {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, zrc20 []common.Address) (*GatewayZEVMOutboundTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestCalledIterator{contract: _GatewayZEVMOutboundTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestCalled, sender []common.Address, zrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestCalled)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "Called", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseCalled(log types.Log) (*GatewayZEVMOutboundTestCalled, error) {
	event := new(GatewayZEVMOutboundTestCalled)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestContextDataIterator is returned from FilterContextData and is used to iterate over the raw logs and unpacked data for ContextData events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataIterator struct {
	Event *GatewayZEVMOutboundTestContextData // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestContextDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestContextData)
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
		it.Event = new(GatewayZEVMOutboundTestContextData)
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
func (it *GatewayZEVMOutboundTestContextDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestContextDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestContextData represents a ContextData event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextData struct {
	Origin    []byte
	Sender    common.Address
	ChainID   *big.Int
	MsgSender common.Address
	Message   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContextData is a free log retrieval operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterContextData(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestContextDataIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestContextDataIterator{contract: _GatewayZEVMOutboundTest.contract, event: "ContextData", logs: logs, sub: sub}, nil
}

// WatchContextData is a free log subscription operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchContextData(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestContextData) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestContextData)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextData", log); err != nil {
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

// ParseContextData is a log parse operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseContextData(log types.Log) (*GatewayZEVMOutboundTestContextData, error) {
	event := new(GatewayZEVMOutboundTestContextData)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestContextDataAbortIterator is returned from FilterContextDataAbort and is used to iterate over the raw logs and unpacked data for ContextDataAbort events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataAbortIterator struct {
	Event *GatewayZEVMOutboundTestContextDataAbort // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestContextDataAbortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestContextDataAbort)
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
		it.Event = new(GatewayZEVMOutboundTestContextDataAbort)
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
func (it *GatewayZEVMOutboundTestContextDataAbortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestContextDataAbortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestContextDataAbort represents a ContextDataAbort event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataAbort struct {
	AbortContext AbortContext
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContextDataAbort is a free log retrieval operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterContextDataAbort(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestContextDataAbortIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "ContextDataAbort")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestContextDataAbortIterator{contract: _GatewayZEVMOutboundTest.contract, event: "ContextDataAbort", logs: logs, sub: sub}, nil
}

// WatchContextDataAbort is a free log subscription operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchContextDataAbort(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestContextDataAbort) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "ContextDataAbort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestContextDataAbort)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextDataAbort", log); err != nil {
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

// ParseContextDataAbort is a log parse operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseContextDataAbort(log types.Log) (*GatewayZEVMOutboundTestContextDataAbort, error) {
	event := new(GatewayZEVMOutboundTestContextDataAbort)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextDataAbort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestContextDataRevertIterator is returned from FilterContextDataRevert and is used to iterate over the raw logs and unpacked data for ContextDataRevert events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataRevertIterator struct {
	Event *GatewayZEVMOutboundTestContextDataRevert // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestContextDataRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestContextDataRevert)
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
		it.Event = new(GatewayZEVMOutboundTestContextDataRevert)
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
func (it *GatewayZEVMOutboundTestContextDataRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestContextDataRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestContextDataRevert represents a ContextDataRevert event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestContextDataRevert struct {
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContextDataRevert is a free log retrieval operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterContextDataRevert(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestContextDataRevertIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestContextDataRevertIterator{contract: _GatewayZEVMOutboundTest.contract, event: "ContextDataRevert", logs: logs, sub: sub}, nil
}

// WatchContextDataRevert is a free log subscription operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchContextDataRevert(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestContextDataRevert) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestContextDataRevert)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
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

// ParseContextDataRevert is a log parse operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseContextDataRevert(log types.Log) (*GatewayZEVMOutboundTestContextDataRevert, error) {
	event := new(GatewayZEVMOutboundTestContextDataRevert)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestWithdrawnIterator struct {
	Event *GatewayZEVMOutboundTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestWithdrawn)
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
		it.Event = new(GatewayZEVMOutboundTestWithdrawn)
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
func (it *GatewayZEVMOutboundTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestWithdrawn represents a Withdrawn event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestWithdrawn struct {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMOutboundTestWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestWithdrawnIterator{contract: _GatewayZEVMOutboundTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestWithdrawn)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseWithdrawn(log types.Log) (*GatewayZEVMOutboundTestWithdrawn, error) {
	event := new(GatewayZEVMOutboundTestWithdrawn)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestWithdrawnAndCalledIterator struct {
	Event *GatewayZEVMOutboundTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestWithdrawnAndCalled)
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
		it.Event = new(GatewayZEVMOutboundTestWithdrawnAndCalled)
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
func (it *GatewayZEVMOutboundTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestWithdrawnAndCalled struct {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMOutboundTestWithdrawnAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestWithdrawnAndCalledIterator{contract: _GatewayZEVMOutboundTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestWithdrawnAndCalled, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestWithdrawnAndCalled)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayZEVMOutboundTestWithdrawnAndCalled, error) {
	event := new(GatewayZEVMOutboundTestWithdrawnAndCalled)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogIterator struct {
	Event *GatewayZEVMOutboundTestLog // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLog)
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
		it.Event = new(GatewayZEVMOutboundTestLog)
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
func (it *GatewayZEVMOutboundTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLog represents a Log event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLog)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLog(log types.Log) (*GatewayZEVMOutboundTestLog, error) {
	event := new(GatewayZEVMOutboundTestLog)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogAddressIterator struct {
	Event *GatewayZEVMOutboundTestLogAddress // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogAddress)
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
		it.Event = new(GatewayZEVMOutboundTestLogAddress)
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
func (it *GatewayZEVMOutboundTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogAddress represents a LogAddress event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogAddressIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogAddressIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogAddress)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
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

// ParseLogAddress is a log parse operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogAddress(log types.Log) (*GatewayZEVMOutboundTestLogAddress, error) {
	event := new(GatewayZEVMOutboundTestLogAddress)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArrayIterator struct {
	Event *GatewayZEVMOutboundTestLogArray // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogArray)
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
		it.Event = new(GatewayZEVMOutboundTestLogArray)
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
func (it *GatewayZEVMOutboundTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogArray represents a LogArray event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogArrayIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogArrayIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogArray)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
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

// ParseLogArray is a log parse operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogArray(log types.Log) (*GatewayZEVMOutboundTestLogArray, error) {
	event := new(GatewayZEVMOutboundTestLogArray)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray0Iterator struct {
	Event *GatewayZEVMOutboundTestLogArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogArray0)
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
		it.Event = new(GatewayZEVMOutboundTestLogArray0)
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
func (it *GatewayZEVMOutboundTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogArray0 represents a LogArray0 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogArray0Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogArray0)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
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

// ParseLogArray0 is a log parse operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogArray0(log types.Log) (*GatewayZEVMOutboundTestLogArray0, error) {
	event := new(GatewayZEVMOutboundTestLogArray0)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray1Iterator struct {
	Event *GatewayZEVMOutboundTestLogArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogArray1)
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
		it.Event = new(GatewayZEVMOutboundTestLogArray1)
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
func (it *GatewayZEVMOutboundTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogArray1 represents a LogArray1 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogArray1Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogArray1)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
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

// ParseLogArray1 is a log parse operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogArray1(log types.Log) (*GatewayZEVMOutboundTestLogArray1, error) {
	event := new(GatewayZEVMOutboundTestLogArray1)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogBytesIterator struct {
	Event *GatewayZEVMOutboundTestLogBytes // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogBytes)
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
		it.Event = new(GatewayZEVMOutboundTestLogBytes)
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
func (it *GatewayZEVMOutboundTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogBytes represents a LogBytes event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogBytesIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogBytesIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogBytes)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
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

// ParseLogBytes is a log parse operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogBytes(log types.Log) (*GatewayZEVMOutboundTestLogBytes, error) {
	event := new(GatewayZEVMOutboundTestLogBytes)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogBytes32Iterator struct {
	Event *GatewayZEVMOutboundTestLogBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogBytes32)
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
		it.Event = new(GatewayZEVMOutboundTestLogBytes32)
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
func (it *GatewayZEVMOutboundTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogBytes32 represents a LogBytes32 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogBytes32Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogBytes32)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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

// ParseLogBytes32 is a log parse operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogBytes32(log types.Log) (*GatewayZEVMOutboundTestLogBytes32, error) {
	event := new(GatewayZEVMOutboundTestLogBytes32)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogIntIterator struct {
	Event *GatewayZEVMOutboundTestLogInt // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogInt)
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
		it.Event = new(GatewayZEVMOutboundTestLogInt)
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
func (it *GatewayZEVMOutboundTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogInt represents a LogInt event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogIntIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogIntIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogInt)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
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

// ParseLogInt is a log parse operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogInt(log types.Log) (*GatewayZEVMOutboundTestLogInt, error) {
	event := new(GatewayZEVMOutboundTestLogInt)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedAddressIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedAddress)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedAddress)
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
func (it *GatewayZEVMOutboundTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedAddressIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedAddress)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
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

// ParseLogNamedAddress is a log parse operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayZEVMOutboundTestLogNamedAddress, error) {
	event := new(GatewayZEVMOutboundTestLogNamedAddress)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArrayIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedArray // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedArray)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedArray)
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
func (it *GatewayZEVMOutboundTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedArray represents a LogNamedArray event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedArrayIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedArray)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
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

// ParseLogNamedArray is a log parse operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayZEVMOutboundTestLogNamedArray, error) {
	event := new(GatewayZEVMOutboundTestLogNamedArray)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray0Iterator struct {
	Event *GatewayZEVMOutboundTestLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedArray0)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedArray0)
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
func (it *GatewayZEVMOutboundTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedArray0Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedArray0)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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

// ParseLogNamedArray0 is a log parse operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayZEVMOutboundTestLogNamedArray0, error) {
	event := new(GatewayZEVMOutboundTestLogNamedArray0)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray1Iterator struct {
	Event *GatewayZEVMOutboundTestLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedArray1)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedArray1)
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
func (it *GatewayZEVMOutboundTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedArray1Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedArray1)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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

// ParseLogNamedArray1 is a log parse operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayZEVMOutboundTestLogNamedArray1, error) {
	event := new(GatewayZEVMOutboundTestLogNamedArray1)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedBytesIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedBytes)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedBytes)
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
func (it *GatewayZEVMOutboundTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedBytesIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedBytes)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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

// ParseLogNamedBytes is a log parse operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayZEVMOutboundTestLogNamedBytes, error) {
	event := new(GatewayZEVMOutboundTestLogNamedBytes)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedBytes32Iterator struct {
	Event *GatewayZEVMOutboundTestLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedBytes32)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedBytes32)
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
func (it *GatewayZEVMOutboundTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedBytes32Iterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedBytes32)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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

// ParseLogNamedBytes32 is a log parse operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayZEVMOutboundTestLogNamedBytes32, error) {
	event := new(GatewayZEVMOutboundTestLogNamedBytes32)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedDecimalIntIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedDecimalInt)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedDecimalInt)
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
func (it *GatewayZEVMOutboundTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedDecimalIntIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedDecimalInt)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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

// ParseLogNamedDecimalInt is a log parse operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayZEVMOutboundTestLogNamedDecimalInt, error) {
	event := new(GatewayZEVMOutboundTestLogNamedDecimalInt)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedDecimalUintIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedDecimalUint)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedDecimalUint)
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
func (it *GatewayZEVMOutboundTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedDecimalUintIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedDecimalUint)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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

// ParseLogNamedDecimalUint is a log parse operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayZEVMOutboundTestLogNamedDecimalUint, error) {
	event := new(GatewayZEVMOutboundTestLogNamedDecimalUint)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedIntIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedInt // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedInt)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedInt)
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
func (it *GatewayZEVMOutboundTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedInt represents a LogNamedInt event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedIntIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedInt)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
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

// ParseLogNamedInt is a log parse operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayZEVMOutboundTestLogNamedInt, error) {
	event := new(GatewayZEVMOutboundTestLogNamedInt)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedStringIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedString // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedString)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedString)
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
func (it *GatewayZEVMOutboundTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedString represents a LogNamedString event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedStringIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedString)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
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

// ParseLogNamedString is a log parse operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedString(log types.Log) (*GatewayZEVMOutboundTestLogNamedString, error) {
	event := new(GatewayZEVMOutboundTestLogNamedString)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedUintIterator struct {
	Event *GatewayZEVMOutboundTestLogNamedUint // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogNamedUint)
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
		it.Event = new(GatewayZEVMOutboundTestLogNamedUint)
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
func (it *GatewayZEVMOutboundTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogNamedUint represents a LogNamedUint event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogNamedUintIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogNamedUint)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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

// ParseLogNamedUint is a log parse operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayZEVMOutboundTestLogNamedUint, error) {
	event := new(GatewayZEVMOutboundTestLogNamedUint)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogStringIterator struct {
	Event *GatewayZEVMOutboundTestLogString // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogString)
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
		it.Event = new(GatewayZEVMOutboundTestLogString)
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
func (it *GatewayZEVMOutboundTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogString represents a LogString event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogStringIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogStringIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogString)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
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

// ParseLogString is a log parse operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogString(log types.Log) (*GatewayZEVMOutboundTestLogString, error) {
	event := new(GatewayZEVMOutboundTestLogString)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogUintIterator struct {
	Event *GatewayZEVMOutboundTestLogUint // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogUint)
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
		it.Event = new(GatewayZEVMOutboundTestLogUint)
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
func (it *GatewayZEVMOutboundTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogUint represents a LogUint event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogUintIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogUintIterator{contract: _GatewayZEVMOutboundTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogUint)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
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

// ParseLogUint is a log parse operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogUint(log types.Log) (*GatewayZEVMOutboundTestLogUint, error) {
	event := new(GatewayZEVMOutboundTestLogUint)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMOutboundTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogsIterator struct {
	Event *GatewayZEVMOutboundTestLogs // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMOutboundTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMOutboundTestLogs)
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
		it.Event = new(GatewayZEVMOutboundTestLogs)
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
func (it *GatewayZEVMOutboundTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMOutboundTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMOutboundTestLogs represents a Logs event raised by the GatewayZEVMOutboundTest contract.
type GatewayZEVMOutboundTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayZEVMOutboundTestLogsIterator, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMOutboundTestLogsIterator{contract: _GatewayZEVMOutboundTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayZEVMOutboundTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMOutboundTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMOutboundTestLogs)
				if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "logs", log); err != nil {
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

// ParseLogs is a log parse operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayZEVMOutboundTest *GatewayZEVMOutboundTestFilterer) ParseLogs(log types.Log) (*GatewayZEVMOutboundTestLogs, error) {
	event := new(GatewayZEVMOutboundTestLogs)
	if err := _GatewayZEVMOutboundTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
