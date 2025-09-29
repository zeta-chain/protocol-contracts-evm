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

// CallOptions is an auto generated low-level Go binding around an user-defined struct.
type CallOptions struct {
	GasLimit        *big.Int
	IsArbitraryCall bool
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

// GatewayZEVMInboundTestMetaData contains all meta data concerning the GatewayZEVMInboundTest contract.
var GatewayZEVMInboundTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testCallFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOpts\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfGasLimitIsBelowMin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfGasLimitTooHigh\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpgradeAndWithdrawZRC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZETAFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfGasLimitTooHigh\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETA\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfNoBalance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithBothDefaultValuesFromRegistry\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithCallOptsFailsIfGasLimitTooHigh\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithCallOptsWithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithDefaultGasLimit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithDefaultProtocolFlatFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfNoBalanceForGasFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfNoBalanceForTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIsAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCallOptsWithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCustomGasLimit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooHigh\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCustomGasLimitFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithMessageFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LowBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZETANotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroGasPrice\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602b575f80fd5b506201a355806200003b5f395ff3fe608060405234801562000010575f80fd5b50600436106200049c575f3560e01c80636f5e27561162000273578063c20049f41162000163578063e20c9f7111620000d3578063eb7a2fac116200009f578063fa7626d41162000083578063fa7626d41462000787578063fbc611c81462000795578063fdad0ad0146200079f575f80fd5b8063eb7a2fac1462000773578063ecf26b30146200077d575f80fd5b8063e20c9f711462000755578063e51c6388146200075f578063e804a40614620004f2578063ea37902f1462000769575f80fd5b8063ceccfab3116200012f578063dc749dd71162000113578063dc749dd71462000737578063dde7e9671462000741578063df0def41146200074b575f80fd5b8063ceccfab31462000723578063d597a27a146200072d575f80fd5b8063c20049f414620006fb578063c3bb79571462000705578063c946d7c0146200070f578063cb8173561462000719575f80fd5b8063a721b2d311620001ff578063b5508aa911620001cb578063ba800c9111620001af578063ba800c9114620006dd578063ba9adeef14620006e7578063bed3e81314620006f1575f80fd5b8063b5508aa914620006b8578063ba414fa614620006c2575f80fd5b8063a721b2d314620005a5578063b0464fdc146200069a578063b152ca4614620006a4578063b51ac07114620006ae575f80fd5b806385226c81116200023f57806385226c81146200065457806386bb7ac1146200066d578063916a17c6146200067757806391dc32e71462000690575f80fd5b80636f5e2756146200062c5780637ae6973014620006365780637ba9b7ad146200064057806383ababa9146200064a575f80fd5b80633828ce8c116200038f5780635efe72a9116200031b57806366d9a9a011620002e75780636d6ce0d011620002cb5780636d6ce0d0146200060e5780636dfcbc5014620006185780636f2154d51462000622575f80fd5b806366d9a9a014620005eb5780636abd223e1462000604575f80fd5b80635efe72a914620005c35780636198fb1914620005cd5780636221b50914620005d757806364002a1f14620005e1575f80fd5b806343181437116200035b57806343181437146200059b5780634ffab9de14620005a5578063564a743514620005af5780635d72228f14620005b9575f80fd5b80633828ce8c14620005735780633e5e3c23146200057d5780633f7286f4146200058757806342752d411462000591575f80fd5b80631e63d2b9116200042b57806322442d7611620003f7578063291f1d8211620003db578063291f1d8214620005465780632ade3880146200055057806336431b3f1462000569575f80fd5b806322442d761462000532578063264b524c146200053c575f80fd5b80631e63d2b914620004f25780631ed7831c14620004fc57806320dee15f146200051e57806321aeb18c1462000528575f80fd5b80631238212c116200046b5780631238212c14620004ca5780631475976614620004d457806316c09ef714620004de5780631b9641bf14620004e8575f80fd5b806301a74aee14620004a057806305b9f04614620004ac5780630a9254e414620004b65780630b5ad28d14620004c0575b5f80fd5b620004aa620007a9565b005b620004aa620009cf565b620004aa62000b76565b620004aa620019c4565b620004aa62001b29565b620004aa62001ec2565b620004aa620025a2565b620004aa62002b1c565b620004aa62002c7b565b6200050662003077565b6040516200051591906200eec9565b60405180910390f35b620004aa620030d9565b620004aa620034c7565b620004aa620035cf565b620004aa6200377d565b620004aa62003a46565b6200055a6200408a565b6040516200051591906200ef44565b620004aa620041d2565b620004aa62004310565b62000506620047e4565b6200050662004844565b620004aa620048a4565b620004aa620049f7565b620004aa62004ae1565b620004aa62004ee9565b620004aa62004f92565b620004aa620052c4565b620004aa62005366565b620004aa62005516565b620004aa62005736565b620005f562005879565b6040516200051591906200f0ad565b620004aa620059e9565b620004aa62005d81565b620004aa62006053565b620004aa62006438565b620004aa62006637565b620004aa62006920565b620004aa62006bdc565b620004aa62006cf7565b6200065e62006e04565b6040516200051591906200f14f565b620004aa62006ed9565b62000681620070dd565b6040516200051591906200f1c8565b620004aa620071c2565b62000681620073ca565b620004aa620074af565b620004aa620074c1565b6200065e620075a8565b620006cc6200767d565b604051901515815260200162000515565b620004aa62007751565b620004aa62007865565b620004aa620079d5565b620004aa62007b36565b620004aa62007cb7565b620004aa62007e5d565b620004aa6200820c565b620004aa620083b8565b620004aa62008794565b620004aa62008915565b620004aa620089b8565b620004aa62008ddc565b6200050662008fd6565b620004aa62009036565b620004aa620094a6565b620004aa62009b17565b620004aa62009bfe565b601f54620006cc9060ff1681565b620004aa62009df2565b620004aa6200a17f565b6027546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052815160a08101835261032180825260018284018190528285019190915283519283019093525f8083526060820192909252929350919060808201906200083a90621e8480906200f28e565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391620008d6916004016200f2a4565b5f604051808303815f87803b158015620008ee575f80fd5b505af115801562000901573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506306cb8983915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526200099c926001600160a01b03909116908790602d9088906004016200f314565b5f604051808303815f87803b158015620009b4575f80fd5b505af1158015620009c7573d5f803e3d5ffd5b505050505050565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529190526020810180516001600160e01b0316630427d73b60e51b179052905062000a266001620186a06200f383565b602d55604051630618f58760e51b81527f98bdb2e0000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b15801562000a94575f80fd5b505af115801562000aa7573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506306cb898391506034015b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835262000b44926001600160a01b03909116908690602d906029906004016200f4d7565b5f604051808303815f87803b15801562000b5c575f80fd5b505af115801562000b6f573d5f803e3d5ffd5b5050505050565b602680547fffffffffffffffffffffffff000000000000000000000000000000000000000090811630179091556027805490911661123417905560405162000bbe906200edd2565b604051809103905ff08015801562000bd8573d5f803e3d5ffd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600f81527f476174657761795a45564d2e736f6c00000000000000000000000000000000006020820152602654915160248101939093529216604482015262000ca9919060640160408051601f198184030181529190526020810180516001600160e01b03167f485cc955000000000000000000000000000000000000000000000000000000001790526200a433565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b039384168102919091179182905560208054919092049092167fffffffffffffffffffffffff0000000000000000000000000000000000000000909216919091178155604051737cce3eb018bf23e1fe2a32692f2c77592d110394915f919062000d459082016200ede0565b601f1982820381018352601f90910116604081905262000d6991906020016200f552565b60405160208183030381529060405290505f808251602084015ff590507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663b4d6c78284836001600160a01b0316803b806020016040519081016040528181525f908060200190933c6040518363ffffffff1660e01b815260040162000dfc9291906200f55f565b5f604051808303815f87803b15801562000e14575f80fd5b505af115801562000e27573d5f803e3d5ffd5b50506026546020546040517fc0c53b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820181905260248201529082166044820152908616925063c0c53b8b91506064015f604051808303815f87803b15801562000e9b575f80fd5b505af115801562000eae573d5f803e3d5ffd5b5050602580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038781169182179092556020546040517fab7b49930000000000000000000000000000000000000000000000000000000081526004810192909252909116925063ab7b499391506024015f604051808303815f87803b15801562000f3f575f80fd5b505af115801562000f52573d5f803e3d5ffd5b505060208054604080517f2722feee00000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169450632722feee935060048082019392918290030181865afa15801562000fb5573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062000fdb91906200f59e565b602880547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039290921691909117905560405162001021906200edee565b604051809103905ff0801580156200103b573d5f803e3d5ffd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161781556028546040517f06447d5600000000000000000000000000000000000000000000000000000000815292166004830152737109709ecfa91a80626ff3989d68f67f5b1dd12d916306447d5691015f604051808303815f87803b158015620010d5575f80fd5b505af1158015620010e8573d5f803e3d5ffd5b505050505f805f604051620010fd906200edfc565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103905ff08015801562001137573d5f803e3d5ffd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392831690811790915560205460405160129360019384935f93919216906200118e906200ee0a565b6200119f969594939291906200f5ba565b604051809103905ff080158015620011b9573d5f803e3d5ffd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081179091556023546040517fee2815ba0000000000000000000000000000000000000000000000000000000081526001600482015260248101929092529091169063ee2815ba906044015f604051808303815f87803b1580156200124e575f80fd5b505af115801562001261573d5f803e3d5ffd5b50506023546040517fa7cb050700000000000000000000000000000000000000000000000000000000815260016004820181905260248201526001600160a01b03909116925063a7cb050791506044015f604051808303815f87803b158015620012c9575f80fd5b505af1158015620012dc573d5f803e3d5ffd5b50506028546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152633b9aca006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d91506044015f604051808303815f87803b1580156200135a575f80fd5b505af11580156200136d573d5f803e3d5ffd5b5050505060225f9054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004015f604051808303818588803b158015620013c0575f80fd5b505af1158015620013d3573d5f803e3d5ffd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303815f875af115801562001445573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200146b91906200f6ae565b506021546026546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e10060248201529116906347e7ef24906044016020604051808303815f875af1158015620014db573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200150191906200f6ae565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156200155d575f80fd5b505af115801562001570573d5f803e3d5ffd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d5691506024015f604051808303815f87803b158015620015e4575f80fd5b505af1158015620015f7573d5f803e3d5ffd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e10060248201529116925063095ea7b391506044016020604051808303815f875af11580156200166a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200169091906200f6ae565b506021546025546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e100602482015291169063095ea7b3906044016020604051808303815f875af115801562001700573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200172691906200f6ae565b5060225f9054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004015f604051808303818588803b15801562001776575f80fd5b505af115801562001789573d5f803e3d5ffd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303815f875af1158015620017fb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200182191906200f6ae565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156200187d575f80fd5b505af115801562001890573d5f803e3d5ffd5b50506040805160a0810182526103218082526001602080840191825283850192835284519081019094525f8085526060840185905260808401528251602980549251151574010000000000000000000000000000000000000000027fffffffffffffffffffffff0000000000000000000000000000000000000000009093166001600160a01b03928316179290921782559151602a8054919093167fffffffffffffffffffffffff000000000000000000000000000000000000000091909116179091559093509150602b906200196890826200f744565b5060809190910151600390910155505060408051808201909152620186a080825260016020909201829052602d55602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909117905550565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b1790525f602d5551630618f58760e51b81527f98bdb2e0000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562001a78575f80fd5b505af115801562001a8b573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b91506034015b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835262000b44926001916001600160a01b0316908790602d906029906004016200f80c565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa15801562001b78573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062001b9e91906200f883565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562001bf8575f80fd5b505af115801562001c0b573d5f803e3d5ffd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f60248201529116925063095ea7b391506044016020604051808303815f875af115801562001c7a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062001ca091906200f6ae565b506027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905280517ff48448140000000000000000000000000000000000000000000000000000000081529051919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f4844814916004808201925f9290919082900301818387803b15801562001d4c575f80fd5b505af115801562001d5f573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f198184030181526021548383018352620186a084525f602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815262001e0c93919289926001600160a01b03909116918891906029906004016200f89b565b5f604051808303815f87803b15801562001e24575f80fd5b505af115801562001e37573d5f803e3d5ffd5b50506021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa15801562001e88573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062001eae91906200f883565b905062001ebc83826200a455565b50505050565b6027546040516001600160a01b03909116602482015260019081905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b17905290516d9876000000000000000000000000918101919091529091505f9060340160408051808303601f190181528282018252600883527f6761734c696d69740000000000000000000000000000000000000000000000006020840152602654915163ca669fa760e01b81526001600160a01b0390921660048301529250737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562001fc1575f80fd5b505af115801562001fd4573d5f803e3d5ffd5b50506025546021546040517fa8f2cb960000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063a8f2cb9693506200202d928992169087906001906004016200f8f3565b5f604051808303815f87803b15801562002045575f80fd5b505af115801562002058573d5f803e3d5ffd5b505050506040518060400160405280600f81526020017f70726f746f636f6c466c6174466565000000000000000000000000000000000081525090505f6003604051602001620020b1919060ff91909116815260200190565b60405160208183030381529060405290505f81806020019051810190620020d991906200f883565b6023546040517fd7fd7afb000000000000000000000000000000000000000000000000000000008152600481018990526001600160a01b039091169063d7fd7afb90602401602060405180830381865afa1580156200213a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200216091906200f883565b602d546200216f91906200f92f565b6200217b91906200f28e565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620021d5575f80fd5b505af1158015620021e8573d5f803e3d5ffd5b50506025546040517f2259e9e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632259e9e591506200223b908990879087906004016200f949565b5f604051808303815f87803b15801562002253575f80fd5b505af115801562002266573d5f803e3d5ffd5b50506026546021546040516370a0823160e01b81526001600160a01b03928316600482018190523194505f93509116906370a0823190602401602060405180830381865afa158015620022bb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620022e191906200f883565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562002370575f80fd5b505af115801562002383573d5f803e3d5ffd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201528b93506001600160a01b0390911691507fd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a9060340160405160208183030381529060405260225f9054906101000a90046001600160a01b03168d888a8060200190518101906200241991906200f883565b8e602d6029604051620024349897969594939291906200f977565b60405180910390a3602080546027546040516001600160a01b039283169363c34da122938e9362002480939116910160609190911b6bffffffffffffffffffffffff1916815260140190565b6040516020818303038152906040528b8b602d60296040518763ffffffff1660e01b8152600401620024b79594939291906200f9f0565b5f604051808303818588803b158015620024cf575f80fd5b505af1158015620024e2573d5f803e3d5ffd5b50506026546001600160a01b03163192506200250e91506200250790508b856200f383565b826200a455565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9291909116906370a0823190602401602060405180830381865afa1580156200255d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200258391906200f883565b9050620025956200250786856200f383565b5050505050505050505050565b6040516d9876000000000000000000000000602082015260019081905f9060340160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562002628575f80fd5b505af11580156200263b573d5f803e3d5ffd5b50506025546021546040517fa8f2cb960000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063a8f2cb96935062002694928792169086906001906004016200f8f3565b5f604051808303815f87803b158015620026ac575f80fd5b505af1158015620026bf573d5f803e3d5ffd5b5050604080518082018252600f81527f70726f746f636f6c466c6174466565000000000000000000000000000000000060208083019190915282516005918101919091529093505f92500160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200276f575f80fd5b505af115801562002782573d5f803e3d5ffd5b50506025546040517f2259e9e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632259e9e59150620027d5908790869086906004016200f949565b5f604051808303815f87803b158015620027ed575f80fd5b505af115801562002800573d5f803e3d5ffd5b505050505f620186a090505f828060200190518101906200282291906200f883565b6023546040517fd7fd7afb000000000000000000000000000000000000000000000000000000008152600481018990526001600160a01b039091169063d7fd7afb90602401602060405180830381865afa15801562002883573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620028a991906200f883565b620028b590846200f92f565b620028c191906200f28e565b6026546021546040516370a0823160e01b81526001600160a01b03928316600482018190529394509231925f92909116906370a0823190602401602060405180830381865afa15801562002917573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200293d91906200f883565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015620029cc575f80fd5b505af1158015620029df573d5f803e3d5ffd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201528b93506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160405160208183030381529060405260225f9054906101000a90046001600160a01b03168d888b80602001905181019062002a7591906200f883565b6040805180820182528d815260016020820152905162002a9e969594939291906029906200fa1e565b60405180910390a3602080546027546040516001600160a01b03928316936322f0ce89938e9362002aea939116910160609190911b6bffffffffffffffffffffffff1916815260140190565b6040516020818303038152906040528b60296040518563ffffffff1660e01b8152600401620024b7939291906200faa3565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562002bcc575f80fd5b505af115801562002bdf573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835262000b44925f916001600160a01b0316908790602d906029906004016200f80c565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa15801562002cca573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002cf091906200f883565b6027546040516001600160a01b0390911660248201529091505f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150620186a0908190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562002dcc575f80fd5b505af115801562002ddf573d5f803e3d5ffd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201525f93506001600160a01b0390911691507fd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a9060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b03909216918b9189918491634d8943bb916004808201926020929091908290030181865afa15801562002eaf573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002ed591906200f883565b6040805180820182528a815260016020820152905162002eff9695949392918d916029906200fad1565b60405180910390a3602080546027546040516001600160a01b0392831693637b15118b9362002f489316910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815260215483830183528684526001602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815262002fb69391928b926001600160a01b03909116918a91906029906004016200f89b565b5f604051808303815f87803b15801562002fce575f80fd5b505af115801562002fe1573d5f803e3d5ffd5b50506021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa15801562003032573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200305891906200f883565b9050620009c7836200306b88886200f383565b6200250791906200f383565b60606016805480602002602001604051908101604052809291908181526020018280548015620030cf57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311620030b0575b5050505050905090565b601f54604080518082018252601a81527f476174657761795a45564d55706772616465546573742e736f6c00000000000060208083019190915282519081019092525f82526026546200313f936001600160a01b0361010090910481169391166200a4d2565b601f546021546026546040516370a0823160e01b81526001600160a01b0391821660048201526101009093048116926001925f9216906370a0823190602401602060405180830381865afa1580156200319a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620031c091906200f883565b6040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b0385166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b1580156200324b575f80fd5b505af11580156200325e573d5f803e3d5ffd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201525f93506001600160a01b0390911691507f5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d979060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b039092169188915f918491634d8943bb916004808201926020929091908290030181865afa1580156200332e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200335491906200f883565b6040805180820182525f81526001602082015290516200337d969594939291906029906200fa1e565b60405180910390a360275460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b03841690637c0dcb5f9060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352620034139287916001600160a01b0316906029906004016200fb37565b5f604051808303815f87803b1580156200342b575f80fd5b505af11580156200343e573d5f803e3d5ffd5b50506021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa1580156200348f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620034b591906200f883565b905062001ebc6200250784846200f383565b604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562003531575f80fd5b505af115801562003544573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506322f0ce8991505f906034015b604051602081830303815290604052600160296040518563ffffffff1660e01b8152600401620035b7939291906200faa3565b5f604051808303818588803b15801562000b5c575f80fd5b60208054604080517f20c612fa00000000000000000000000000000000000000000000000000000000815290516001935f936001600160a01b0316926320c612fa92600480830193928290030181865afa15801562003630573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200365691906200f883565b620036639060016200f28e565b604051630618f58760e51b81527f98bdb2e0000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015620036d0575f80fd5b505af1158015620036e3573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c5356f6d91506034015b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526200099c9287916001600160a01b03169087906029906004016200fb74565b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169263bcbe9365926004808401938290030181865afa158015620037da573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200380091906200f883565b6200380d9060016200f28e565b67ffffffffffffffff8111156200382857620038286200f6cf565b6040519080825280601f01601f19166020018201604052801562003853576020820181803683370190505b50602b906200386390826200f744565b505f6029600201805462003877906200f399565b905090505f60205f9054906101000a90046001600160a01b03166001600160a01b031663bcbe93656040518163ffffffff1660e01b8152600401602060405180830381865afa158015620038cd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620038f391906200f883565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167fcd6f4e6d00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162003983916004016200f2a4565b5f604051808303815f87803b1580156200399b575f80fd5b505af1158015620039ae573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526200099c926002916001600160a01b0316906029906004016200fb37565b6040516d9876000000000000000000000000602082015260019081905f9060340160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562003acc575f80fd5b505af115801562003adf573d5f803e3d5ffd5b50506025546021546040517fa8f2cb960000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063a8f2cb96935062003b38928792169086906001906004016200f8f3565b5f604051808303815f87803b15801562003b50575f80fd5b505af115801562003b63573d5f803e3d5ffd5b5050604080518082018252600881527f6761734c696d69740000000000000000000000000000000000000000000000006020808301919091528251620249f0918101919091529093505f92500160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562003c15575f80fd5b505af115801562003c28573d5f803e3d5ffd5b50506025546040517f2259e9e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632259e9e5915062003c7b908790869086906004016200f949565b5f604051808303815f87803b15801562003c93575f80fd5b505af115801562003ca6573d5f803e3d5ffd5b50506023546040517fd7fd7afb000000000000000000000000000000000000000000000000000000008152600481018890525f93506001600160a01b03909116915063d7fd7afb90602401602060405180830381865afa15801562003d0d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062003d3391906200f883565b8280602001905181019062003d4991906200f883565b62003d5591906200f92f565b6026546021546040516370a0823160e01b81526001600160a01b03928316600482018190529394505f933192849216906370a0823190602401602060405180830381865afa15801562003daa573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062003dd091906200f883565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562003e5f575f80fd5b505af115801562003e72573d5f803e3d5ffd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201528b93506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160405160208183030381529060405260225f9054906101000a90046001600160a01b03168d898960405180604001604052808e80602001905181019062003f1491906200f883565b8152600160209091015260405162003f35969594939291906029906200fa1e565b60405180910390a3602080546027546040516001600160a01b03928316936322f0ce89938e9362003f81939116910160609190911b6bffffffffffffffffffffffff1916815260140190565b6040516020818303038152906040528b60296040518563ffffffff1660e01b815260040162003fb3939291906200faa3565b5f604051808303818588803b15801562003fcb575f80fd5b505af115801562003fde573d5f803e3d5ffd5b50506026546001600160a01b03163192506200400391506200250790508b856200f383565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9291909116906370a0823190602401602060405180830381865afa15801562004052573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200407891906200f883565b9050620025956200250787856200f383565b6060601e805480602002602001604051908101604052809291908181526020015f905b82821015620041c9575f84815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015620041b1578382905f5260205f200180546200411f906200f399565b80601f01602080910402602001604051908101604052809291908181526020018280546200414d906200f399565b80156200419c5780601f1062004172576101008083540402835291602001916200419c565b820191905f5260205f20905b8154815290600101906020018083116200417e57829003601f168201915b505050505081526020019060010190620040ff565b505050508152505081526020019060010190620040ad565b50505050905090565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f7138356f000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562004282575f80fd5b505af115801562004295573d5f803e3d5ffd5b505060208054604080515f808252602154606083018452620186a09583019586528284019190915291517f7b15118b0000000000000000000000000000000000000000000000000000000081526001600160a01b039384169650637b15118b955062000b44949193600193169188916029906004016200f89b565b6040516d9876000000000000000000000000602082015260019081905f9060340160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562004396575f80fd5b505af1158015620043a9573d5f803e3d5ffd5b50506025546021546040517fa8f2cb960000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063a8f2cb96935062004402928792169086906001906004016200f8f3565b5f604051808303815f87803b1580156200441a575f80fd5b505af11580156200442d573d5f803e3d5ffd5b50506023546040517fd7fd7afb00000000000000000000000000000000000000000000000000000000815260048101869052620186a093505f92506001600160a01b039091169063d7fd7afb90602401602060405180830381865afa15801562004499573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620044bf91906200f883565b620044cb90836200f92f565b6026546021546040516370a0823160e01b81526001600160a01b03928316600482018190529394505f933192849216906370a0823190602401602060405180830381865afa15801562004520573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200454691906200f883565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015620045d5575f80fd5b505af1158015620045e8573d5f803e3d5ffd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201528a93506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160408051601f1981840301815260225483830183528b8452600160208501529151620046839391926001600160a01b0316918f918c918c916029906200fa1e565b60405180910390a3602080546027546040516001600160a01b03928316936322f0ce89938d93620046cf939116910160609190911b6bffffffffffffffffffffffff1916815260140190565b6040516020818303038152906040528a60296040518563ffffffff1660e01b815260040162004701939291906200faa3565b5f604051808303818588803b15801562004719575f80fd5b505af11580156200472c573d5f803e3d5ffd5b50506026546001600160a01b03163192506200475191506200250790508a856200f383565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9291909116906370a0823190602401602060405180830381865afa158015620047a0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620047c691906200f883565b9050620047d86200250787856200f383565b50505050505050505050565b60606018805480602002602001604051908101604052809291908181526020018280548015620030cf57602002820191905f5260205f209081546001600160a01b03168152600190910190602001808311620030b0575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015620030cf57602002820191905f5260205f209081546001600160a01b03168152600190910190602001808311620030b0575050505050905090565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562004954575f80fd5b505af115801562004967573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c34da12291505f906034015b604051602081830303815290604052600185602d60296040518763ffffffff1660e01b8152600401620049df9594939291906200f9f0565b5f604051808303818588803b158015620009b4575f80fd5b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f7138356f000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562004aa7575f80fd5b505af115801562004aba573d5f803e3d5ffd5b5050602080546040516001600160a01b03909116935063c34da122925060019101620049a7565b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290515f936002936001600160a01b03169263bcbe936592600480830193928290030181865afa15801562004b42573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004b6891906200f883565b62004b7491906200fbb7565b67ffffffffffffffff81111562004b8f5762004b8f6200f6cf565b6040519080825280601f01601f19166020018201604052801562004bba576020820181803683370190505b5060208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263bcbe9365926004808401938290030181865afa15801562004c1e573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004c4491906200f883565b62004c5091906200fbb7565b62004c5d9060016200f28e565b67ffffffffffffffff81111562004c785762004c786200f6cf565b6040519080825280601f01601f19166020018201604052801562004ca3576020820181803683370190505b50602b9062004cb390826200f744565b505f6029600201805462004cc7906200f399565b835162004cd592506200f28e565b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290519394505f936001600160a01b039092169263bcbe9365926004808401938290030181865afa15801562004d37573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004d5d91906200f883565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167fcd6f4e6d00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162004ded916004016200f2a4565b5f604051808303815f87803b15801562004e05575f80fd5b505af115801562004e18573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835262004eb5926001916001600160a01b0316908990602d906029906004016200f80c565b5f604051808303815f87803b15801562004ecd575f80fd5b505af115801562004ee0573d5f803e3d5ffd5b50505050505050565b604051630618f58760e51b81527f7138356f000000000000000000000000000000000000000000000000000000006004820152600190620249f090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562004f5b575f80fd5b505af115801562004f6e573d5f803e3d5ffd5b5050602080546040516001600160a01b03909116935063c5356f6d92500162003721565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa15801562004fe1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200500791906200f883565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562005061575f80fd5b505af115801562005074573d5f803e3d5ffd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f60248201529116925063095ea7b391506044016020604051808303815f875af1158015620050e3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200510991906200f6ae565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562005165575f80fd5b505af115801562005178573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526200520f9287916001600160a01b0316906029906004016200fb37565b5f604051808303815f87803b15801562005227575f80fd5b505af11580156200523a573d5f803e3d5ffd5b50506021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa1580156200528b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620052b191906200f883565b9050620052bf82826200a455565b505050565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b1790525f602d5551630618f58760e51b81527f98bdb2e0000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e09060240162000a7c565b6027546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562005438575f80fd5b505af11580156200544b573d5f803e3d5ffd5b505060215460265460275460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b039283169450911691507f306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e49060340160408051601f1981840301815290829052620054cd918690602d906029906200fbf0565b60405180910390a3602080546027546040516001600160a01b03928316936306cb89839362000ae59316910160609190911b6bffffffffffffffffffffffff1916815260140190565b6022546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa15801562005565573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200558b91906200f883565b6022546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526101236004820152602481018390529192506001600160a01b03169063a9059cbb906044016020604051808303815f875af1158015620055f6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200561c91906200f6ae565b505f600190507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156200567d575f80fd5b505af115801562005690573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506322f0ce89915085906034016040516020818303038152906040528460296040518563ffffffff1660e01b815260040162005701939291906200faa3565b5f604051808303818588803b15801562005719575f80fd5b505af11580156200572c573d5f803e3d5ffd5b5050505050505050565b6027546040516001600160a01b0390911660248201526001905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b1790525f602d5551630618f58760e51b81527f98bdb2e0000000000000000000000000000000000000000000000000000000006004820152909150600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015620057f0575f80fd5b505af115801562005803573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c34da122915085906034016040516020818303038152906040528486602d60296040518763ffffffff1660e01b8152600401620057019594939291906200f9f0565b6060601b805480602002602001604051908101604052809291908181526020015f905b82821015620041c9578382905f5260205f2090600202016040518060400160405290815f82018054620058cf906200f399565b80601f0160208091040260200160405190810160405280929190818152602001828054620058fd906200f399565b80156200594c5780601f1062005922576101008083540402835291602001916200594c565b820191905f5260205f20905b8154815290600101906020018083116200592e57829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015620059d057602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b03191681526020019060040190602082600301049283019260010382029150808411620059915790505b505050505081525050815260200190600101906200589c565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019262030d40925f929116906370a0823190602401602060405180830381865afa15801562005a3e573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005a6491906200f883565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b0390911660848201529091508290737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562005af5575f80fd5b505af115801562005b08573d5f803e3d5ffd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201525f93506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b03909216918a9188918491634d8943bb916004808201926020929091908290030181865afa15801562005bd8573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005bfe91906200f883565b6040805180820182528c815260016020820152905162005c27969594939291906029906200fa1e565b60405180910390a3602080546027546040516001600160a01b039283169363c5356f6d9362005c709316910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835262005ccc9289916001600160a01b03169089906029906004016200fb74565b5f604051808303815f87803b15801562005ce4575f80fd5b505af115801562005cf7573d5f803e3d5ffd5b50506021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa15801562005d48573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005d6e91906200f883565b905062000b6f826200306b87866200f383565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa15801562005dd0573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005df691906200f883565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562005e50575f80fd5b505af115801562005e63573d5f803e3d5ffd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f60248201529116925063095ea7b391506044016020604051808303815f875af115801562005ed2573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005ef891906200f6ae565b506027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905280517ff48448140000000000000000000000000000000000000000000000000000000081529051919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f4844814916004808201925f9290919082900301818387803b15801562005fa4575f80fd5b505af115801562005fb7573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835262001e0c9288916001600160a01b0316908790602d906029906004016200f80c565b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290515f936002936001600160a01b03169263bcbe936592600480830193928290030181865afa158015620060b4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620060da91906200f883565b620060e691906200fbb7565b67ffffffffffffffff8111156200610157620061016200f6cf565b6040519080825280601f01601f1916602001820160405280156200612c576020820181803683370190505b5060208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263bcbe9365926004808401938290030181865afa15801562006190573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620061b691906200f883565b620061c291906200fbb7565b620061cf9060016200f28e565b67ffffffffffffffff811115620061ea57620061ea6200f6cf565b6040519080825280601f01601f19166020018201604052801562006215576020820181803683370190505b50602b906200622590826200f744565b505f6029600201805462006239906200f399565b83516200624792506200f28e565b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290519394505f936001600160a01b039092169263bcbe9365926004808401938290030181865afa158015620062a9573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620062cf91906200f883565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167fcd6f4e6d00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916200635f916004016200f2a4565b5f604051808303815f87803b15801562006377575f80fd5b505af11580156200638a573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f198184030181526021548383018352620186a084525f602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815262004eb59391926001926001600160a01b03909116918a91906029906004016200f89b565b6027546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052805482517f20c612fa00000000000000000000000000000000000000000000000000000000815292519394505f936001600160a01b03909116926320c612fa9260048083019391928290030181865afa158015620064de573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200650491906200f883565b620065119060016200f28e565b604051630618f58760e51b81527f98bdb2e0000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156200657e575f80fd5b505af115801562006591573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506306cb8983915060340160408051601f1981840301815260215483830183528684526001602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1681526200099c9391926001600160a01b03169188916029906004016200fc44565b60208054604080517fbcbe93650000000000000000000000000000000000000000000000000000000081529051600193620249f0936001600160a01b03169263bcbe936592600480830193928290030181865afa1580156200669b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620066c191906200f883565b620066ce9060016200f28e565b67ffffffffffffffff811115620066e957620066e96200f6cf565b6040519080825280601f01601f19166020018201604052801562006714576020820181803683370190505b50602b906200672490826200f744565b505f6029600201805462006738906200f399565b905090505f60205f9054906101000a90046001600160a01b03166001600160a01b031663bcbe93656040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200678e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620067b491906200f883565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167fcd6f4e6d00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162006844916004016200f2a4565b5f604051808303815f87803b1580156200685c575f80fd5b505af11580156200686f573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c5356f6d915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352620069089289916001600160a01b03169089906029906004016200fb74565b5f604051808303815f87803b15801562005719575f80fd5b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169263bcbe9365926004808401938290030181865afa1580156200697d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620069a391906200f883565b620069b09060016200f28e565b67ffffffffffffffff811115620069cb57620069cb6200f6cf565b6040519080825280601f01601f191660200182016040528015620069f6576020820181803683370190505b50602b9062006a0690826200f744565b505f6029600201805462006a1a906200f399565b905090505f60205f9054906101000a90046001600160a01b03166001600160a01b031663bcbe93656040518163ffffffff1660e01b8152600401602060405180830381865afa15801562006a70573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062006a9691906200f883565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167fcd6f4e6d00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162006b26916004016200f2a4565b5f604051808303815f87803b15801562006b3e575f80fd5b505af115801562006b51573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506322f0ce899150600190603401604051602081830303815290604052600160296040518563ffffffff1660e01b815260040162006bc4939291906200faa3565b5f604051808303818588803b15801562004ecd575f80fd5b604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152620249f090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562006c4b575f80fd5b505af115801562006c5e573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c5356f6d915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835262000b44925f916001600160a01b03169087906029906004016200fb74565b604051630618f58760e51b81527f7138356f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562006d61575f80fd5b505af115801562006d74573d5f803e3d5ffd5b505060208054604080515f8152928301908190526021547f7c0dcb5f000000000000000000000000000000000000000000000000000000009091526001600160a01b039182169450637c0dcb5f935062006dd99291600191166029602484016200fb37565b5f604051808303815f87803b15801562006df1575f80fd5b505af115801562001ebc573d5f803e3d5ffd5b6060601a805480602002602001604051908101604052809291908181526020015f905b82821015620041c9578382905f5260205f2001805462006e47906200f399565b80601f016020809104026020016040519081016040528092919081815260200182805462006e75906200f399565b801562006ec45780601f1062006e9a5761010080835404028352916020019162006ec4565b820191905f5260205f20905b81548152906001019060200180831162006ea657829003601f168201915b50505050508152602001906001019062006e27565b6027546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052805482517f20c612fa00000000000000000000000000000000000000000000000000000000815292519394505f936001600160a01b03909116926320c612fa9260048083019391928290030181865afa15801562006f7f573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062006fa591906200f883565b62006fb29060016200f28e565b604051630618f58760e51b81527f98bdb2e0000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156200701f575f80fd5b505af115801562007032573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f1981840301815260215483830183528684525f602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1681526200099c9391926001926001600160a01b03909116918991906029906004016200f89b565b6060601d805480602002602001604051908101604052809291908181526020015f905b82821015620041c9575f8481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015620071a957602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116200716a5790505b5050505050815250508152602001906001019062007100565b6027546040516001600160a01b03909116602482015260019081905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052815160a08101835261032180825260018284018190528285019190915283519283019093525f8083526060820192909252929350919060808201906200725890621e8480906200f28e565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391620072f4916004016200f2a4565b5f604051808303815f87803b1580156200730c575f80fd5b505af11580156200731f573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c34da122915086906034016040516020818303038152906040528686602d876040518763ffffffff1660e01b8152600401620073949594939291906200fc96565b5f604051808303818588803b158015620073ac575f80fd5b505af1158015620073bf573d5f803e3d5ffd5b505050505050505050565b6060601c805480602002602001604051908101604052809291908181526020015f905b82821015620041c9575f8481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156200749657602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b03191681526020019060040190602082600301049283019260010382029150808411620074575790505b50505050508152505081526020019060010190620073ed565b60015f6200366382620186a06200f383565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f7138356f000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562007571575f80fd5b505af115801562007584573d5f803e3d5ffd5b5050602080546040516001600160a01b0390911693506306cb898392500162000ae5565b60606019805480602002602001604051908101604052809291908181526020015f905b82821015620041c9578382905f5260205f20018054620075eb906200f399565b80601f016020809104026020016040519081016040528092919081815260200182805462007619906200f399565b8015620076685780601f106200763e5761010080835404028352916020019162007668565b820191905f5260205f20905b8154815290600101906020018083116200764a57829003601f168201915b505050505081526020019060010190620075cb565b6008545f9060ff161562007695575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c6564000000000000000000000000000000000000000000000000000060248301525f9163667f9d7090604401602060405180830381865afa15801562007724573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200774a91906200f883565b1415905090565b604051630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015620077bb575f80fd5b505af1158015620077ce573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835262006dd9925f916001600160a01b0316906029906004016200fb37565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f5945ea56000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562007915575f80fd5b505af115801562007928573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f198184030181526021548383018352620186a084525f6020850181905292517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b16815262000b44949293926001600160a01b039092169188916029906004016200f89b565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f98bdb2e0000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562007a85575f80fd5b505af115801562007a98573d5f803e3d5ffd5b5050602080546027546040805160609290921b6bffffffffffffffffffffffff191693820193909352825180820360140181526021546074830185525f6034840181815260549094015293517f7b15118b0000000000000000000000000000000000000000000000000000000081526001600160a01b039384169650637b15118b955062000b4494919360019392169188916029906004016200f89b565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f98bdb2e0000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562007be6575f80fd5b505af115801562007bf9573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152602154838301909252916001916001600160a01b03169086908062007c6c85620186a06200f383565b81525f6020909101526040517fffffffff0000000000000000000000000000000000000000000000000000000060e088901b16815262000b449594939291906029906004016200f89b565b6040805160a081018252610321808252600160208084018290528385019290925283519182019093525f8082526060830191909152906080810162007d00621e8480856200f28e565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162007d9c916004016200f2a4565b5f604051808303815f87803b15801562007db4575f80fd5b505af115801562007dc7573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526200099c9287916001600160a01b03169087906004016200fcc4565b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290515f936002936001600160a01b03169263bcbe936592600480830193928290030181865afa15801562007ebe573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062007ee491906200f883565b62007ef091906200fbb7565b67ffffffffffffffff81111562007f0b5762007f0b6200f6cf565b6040519080825280601f01601f19166020018201604052801562007f36576020820181803683370190505b5060208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263bcbe9365926004808401938290030181865afa15801562007f9a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062007fc091906200f883565b62007fcc91906200fbb7565b62007fd99060016200f28e565b67ffffffffffffffff81111562007ff45762007ff46200f6cf565b6040519080825280601f01601f1916602001820160405280156200801f576020820181803683370190505b50602b906200802f90826200f744565b505f6029600201805462008043906200f399565b83516200805192506200f28e565b60208054604080517fbcbe936500000000000000000000000000000000000000000000000000000000815290519394505f936001600160a01b039092169263bcbe9365926004808401938290030181865afa158015620080b3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620080d991906200f883565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167fcd6f4e6d00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162008169916004016200f2a4565b5f604051808303815f87803b15801562008181575f80fd5b505af115801562008194573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c34da1229150600190603401604051602081830303815290604052600187602d60296040518763ffffffff1660e01b8152600401620057019594939291906200f9f0565b6040805160a081018252610321808252600160208084018290528385019290925283519182019093525f8082526060830191909152620249f0916080810162008259621e8480866200f28e565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391620082f5916004016200f2a4565b5f604051808303815f87803b1580156200830d575f80fd5b505af115801562008320573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c5356f6d915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835262004eb59288916001600160a01b031690889088906004016200fd01565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa15801562008407573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200842d91906200f883565b6021546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526101236004820152602481018390529192506001600160a01b03169063a9059cbb906044016020604051808303815f875af115801562008498573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620084be91906200f6ae565b5060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562008516575f80fd5b505af115801562008529573d5f803e3d5ffd5b50506021546040517ff687d12a000000000000000000000000000000000000000000000000000000008152600a60048201526001600160a01b03909116925063f687d12a91506024015f604051808303815f87803b1580156200858a575f80fd5b505af11580156200859d573d5f803e3d5ffd5b50506021546040517ffc5fecd5000000000000000000000000000000000000000000000000000000008152600a60048201525f93508392506001600160a01b039091169063fc5fecd5906024016040805180830381865afa15801562008605573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200862b91906200fd44565b60208054604080516001600160a01b03868116602483015290921660448301526064808301859052815180840390910181526084909201815291810180516001600160e01b03167f6670111200000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152929450909250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391620086d2916004016200f2a4565b5f604051808303815f87803b158015620086ea575f80fd5b505af1158015620086fd573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352620069089289916001600160a01b0316906029906004016200fb37565b6040805160a081018252610321808252600160208084018290528385019290925283519182019093525f8082526060830191909152829160808101620087de621e8480856200f28e565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916200887a916004016200f2a4565b5f604051808303815f87803b15801562008892575f80fd5b505af1158015620088a5573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506322f0ce899150859060340160405160208183030381529060405285856040518563ffffffff1660e01b815260040162005701939291906200fd71565b604051630618f58760e51b81527f7138356f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156200897f575f80fd5b505af115801562008992573d5f803e3d5ffd5b5050602080546040516001600160a01b0390911693506322f0ce8992505f910162003584565b6022546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa15801562008a07573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062008a2d91906200f883565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa15801562008a7d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062008aa391906200f883565b60285460265460405163ca669fa760e01b81526001600160a01b039182166004820152929350163190600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562008b06575f80fd5b505af115801562008b19573d5f803e3d5ffd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f60248201529116925063095ea7b391506044016020604051808303815f875af115801562008b88573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062008bae91906200f6ae565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562008c0a575f80fd5b505af115801562008c1d573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506322f0ce89915087906034016040516020818303038152906040528460296040518563ffffffff1660e01b815260040162008c8e939291906200faa3565b5f604051808303818588803b15801562008ca6575f80fd5b505af115801562008cb9573d5f803e3d5ffd5b50506022546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9550911692506370a082319150602401602060405180830381865afa15801562008d0b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062008d3191906200f883565b905062008d3f85826200a455565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201525f9291909116906370a0823190602401602060405180830381865afa15801562008d8e573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062008db491906200f883565b905062008dc285826200a455565b60285462004ee09085906001600160a01b0316316200a455565b6027546040516001600160a01b03909116602482015260019081905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052805482517f20c612fa00000000000000000000000000000000000000000000000000000000815292519394505f936001600160a01b03909116926320c612fa9260048083019391928290030181865afa15801562008e87573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062008ead91906200f883565b62008eba9060016200f28e565b604051630618f58760e51b81527f98bdb2e0000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562008f27575f80fd5b505af115801562008f3a573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c34da1229150869060340160408051601f1981840301815282820182528683526001602084015290517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b168152620073949289918991906029906004016200fd9f565b60606015805480602002602001604051908101604052809291908181526020018280548015620030cf57602002820191905f5260205f209081546001600160a01b03168152600190910190602001808311620030b0575050505050905090565b6022546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa15801562009085573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620090ab91906200f883565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa158015620090fb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200912191906200f883565b6028546027546040516001600160a01b0391821660248201529293501631905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052602654905163ca669fa760e01b81526001600160a01b039091166004820152909150600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620091cb575f80fd5b505af1158015620091de573d5f803e3d5ffd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f60248201529116925063095ea7b391506044016020604051808303815f875af11580156200924d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200927391906200f6ae565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b158015620092cf575f80fd5b505af1158015620092e2573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c34da122915088906034016040516020818303038152906040528486602d60296040518763ffffffff1660e01b8152600401620093589594939291906200f9f0565b5f604051808303818588803b15801562009370575f80fd5b505af115801562009383573d5f803e3d5ffd5b50506022546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9550911692506370a082319150602401602060405180830381865afa158015620093d5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620093fb91906200f883565b90506200940986826200a455565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201525f9291909116906370a0823190602401602060405180830381865afa15801562009458573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200947e91906200f883565b90506200948c86826200a455565b6028546200572c9086906001600160a01b0316316200a455565b6040516d98760000000000000000000000006020820152600a906001905f9060340160408051808303601f190181528282018252600883527f6761734c696d697400000000000000000000000000000000000000000000000060208085019190915282516002918101919091529093505f910160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200957e575f80fd5b505af115801562009591573d5f803e3d5ffd5b50506025546021546040517fa8f2cb960000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063a8f2cb969350620095ea928992169088906001906004016200f8f3565b5f604051808303815f87803b15801562009602575f80fd5b505af115801562009615573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801562009670575f80fd5b505af115801562009683573d5f803e3d5ffd5b50506025546040517f2259e9e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632259e9e59150620096d6908790869086906004016200f949565b5f604051808303815f87803b158015620096ee575f80fd5b505af115801562009701573d5f803e3d5ffd5b505050506040518060400160405280600f81526020017f70726f746f636f6c466c6174466565000000000000000000000000000000000081525091505f60036040516020016200975a919060ff91909116815260200190565b60405160208183030381529060405290505f818060200190518101906200978291906200f883565b6023546040517fd7fd7afb000000000000000000000000000000000000000000000000000000008152600481018990526001600160a01b039091169063d7fd7afb90602401602060405180830381865afa158015620097e3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200980991906200f883565b848060200190518101906200981f91906200f883565b6200982b91906200f92f565b6200983791906200f28e565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562009891575f80fd5b505af1158015620098a4573d5f803e3d5ffd5b50506025546040517f2259e9e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632259e9e59150620098f7908990889087906004016200f949565b5f604051808303815f87803b1580156200990f575f80fd5b505af115801562009922573d5f803e3d5ffd5b50506026546021546040516370a0823160e01b81526001600160a01b03928316600482018190523194505f93509116906370a0823190602401602060405180830381865afa15801562009977573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200999d91906200f883565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562009a2c575f80fd5b505af115801562009a3f573d5f803e3d5ffd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201528b93506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160405160208183030381529060405260225f9054906101000a90046001600160a01b03168d888a80602001905181019062009ad591906200f883565b60405180604001604052808e80602001905181019062009af691906200f883565b8152600160209091015260405162002a9e969594939291906029906200fa1e565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f7138356f000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562009bc7575f80fd5b505af115801562009bda573d5f803e3d5ffd5b5050602080546040516001600160a01b039091169350637b15118b92500162001ac9565b6027546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052815160a08101835261032180825260018284018190528285019190915283519283019093525f80835260608201929092529293509190608082019062009c8f90621e8480906200f28e565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162009d2b916004016200f2a4565b5f604051808303815f87803b15801562009d43575f80fd5b505af115801562009d56573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526200099c926001916001600160a01b0316908890602d9089906004016200fdcd565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa15801562009e41573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062009e6791906200f883565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562009ef6575f80fd5b505af115801562009f09573d5f803e3d5ffd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201525f93506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b039092169188915f918491634d8943bb916004808201926020929091908290030181865afa15801562009fd9573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062009fff91906200f883565b6040805180820182525f81526001602082015290516200a028969594939291906029906200fa1e565b60405180910390a3602080546027546040516001600160a01b0392831693637c0dcb5f936200a0719316910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526200a0cb9287916001600160a01b0316906029906004016200fb37565b5f604051808303815f87803b1580156200a0e3575f80fd5b505af11580156200a0f6573d5f803e3d5ffd5b50506021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa1580156200a147573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a16d91906200f883565b9050620052bf6200250784846200f383565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526002925f9216906370a0823190602401602060405180830381865afa1580156200a1ce573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a1f491906200f883565b6021549091506001600160a01b031663a9059cbb6101236200a2186001856200f383565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af11580156200a279573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a29f91906200f6ae565b5060215460265460208054604080516001600160a01b039586166024820181905294861660448201819052959092166064830181905260848084018990528251808503909101815260a4909301825292820180516001600160e01b03167f489ca9b7000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152929392737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916200a35991906004016200f2a4565b5f604051808303815f87803b1580156200a371575f80fd5b505af11580156200a384573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526200a41b928a916001600160a01b0316906029906004016200fb37565b5f604051808303815f87803b158015620073ac575f80fd5b5f6200a43e6200ee18565b6200a44b8484836200a4eb565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c54906044015f6040518083038186803b1580156200a4bf575f80fd5b505afa158015620009c7573d5f803e3d5ffd5b6200a4dc6200ee18565b62000b6f85858584866200a56b565b5f806200a4f985846200a660565b90506200a5606040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f787900000081525082866040516020016200a54a9291906200f55f565b604051602081830303815290604052856200a66d565b9150505b9392505050565b6040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201528190737109709ecfa91a80626ff3989d68f67f5b1dd12d9081906306447d56906024015f604051808303815f87803b1580156200a5db575f80fd5b505af19250505080156200a5ed575060015b6200a606576200a600878787876200a6a0565b62004ee0565b6200a614878787876200a6a0565b806001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156200a64d575f80fd5b505af115801562002595573d5f803e3d5ffd5b5f6200a56483836200a6bc565b60c0810151515f90156200a694576200a68c84848460c001516200a6da565b90506200a564565b6200a68c84846200a88e565b5f6200a6ad84836200a97f565b905062000b6f8582856200a98c565b5f6200a6c983836200ad58565b6200a564838360200151846200a66d565b5f806200a6e66200ad69565b90505f6200a6f586836200ae3d565b90505f6200a70d82606001518360200151856200b2f5565b90505f6200a71e838389896200b51a565b90505f6200a72c826200c45c565b602081015181519192509060030b156200a7a4578982604001516040516020016200a7599291906200fe36565b60408051601f19818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526200a79b916004016200f2a4565b60405180910390fd5b5f6200a7e86040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a2000000000000000000000008152508360016200c632565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d906200a83d9084906004016200f2a4565b602060405180830381865afa1580156200a859573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a87f91906200f59e565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081525f908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc925906200a8e49087906004016200f2a4565b5f60405180830381865afa1580156200a8ff573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200a92891908101906200ff66565b90505f6200a95982856040516020016200a9449291906200ff9c565b6040516020818303038152906040526200c837565b90506001600160a01b0381166200a44b5784846040516020016200a7599291906200ffb4565b5f6200a6c983836200c848565b6040517f667f9d700000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201527fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d905f90829063667f9d7090604401602060405180830381865afa1580156200aa26573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200aa4c91906200f883565b9050806200abf4575f6200aa60866200c856565b604080518082018252600581527f352e302e300000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200aaeb905b6040805180820182525f80825260209182015281518083019092528451825280850190820152906200c93c565b806200aaf757505f8451115b156200ab7b576040517f4f1ef2860000000000000000000000000000000000000000000000000000000081526001600160a01b03871690634f1ef286906200ab4690889088906004016200f55f565b5f604051808303815f87803b1580156200ab5e575f80fd5b505af11580156200ab71573d5f803e3d5ffd5b505050506200abed565b6040517f3659cfe60000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152871690633659cfe6906024015f604051808303815f87803b1580156200abd5575f80fd5b505af11580156200abe8573d5f803e3d5ffd5b505050505b5062000b6f565b805f6200ac01826200c856565b604080518082018252600581527f352e302e300000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200ac64906200aabe565b806200ac7057505f8551115b156200acf6576040517f9623609d0000000000000000000000000000000000000000000000000000000081526001600160a01b03831690639623609d906200acc1908a908a908a9060040162010048565b5f604051808303815f87803b1580156200acd9575f80fd5b505af11580156200acec573d5f803e3d5ffd5b5050505062004ee0565b6040517f99a88ec40000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015287811660248301528316906399a88ec4906044015f604051808303815f87803b1580156200a64d575f80fd5b6200ad6582825f6200c951565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c906200adf29084906004016201007a565b5f60405180830381865afa1580156200ae0d573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200ae369190810190620100c2565b9250505090565b6200ae706040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b5f737109709ecfa91a80626ff3989d68f67f5b1dd12d90506200aebb6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6200aec6856200ca5e565b60208201525f6200aed7866200ce57565b90505f836001600160a01b031663d930a0e66040518163ffffffff1660e01b81526004015f60405180830381865afa1580156200af16573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200af3f9190810190620100c2565b868385602001516040516020016200af5b94939291906201010c565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291505f906001600160a01b038616906360f9bb11906200afb49085906004016200f2a4565b5f60405180830381865afa1580156200afcf573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200aff89190810190620100c2565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f6906200b042908490600401620101e4565b602060405180830381865afa1580156200b05e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b08491906200f6ae565b6200b09c57816040516020016200a759919062010237565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906200b0e3908490600401620102bd565b5f60405180830381865afa1580156200b0fe573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200b1279190810190620100c2565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f6906200b17090849060040162010310565b602060405180830381865afa1580156200b18c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b1b291906200f6ae565b156200b249576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906200b1ff90849060040162010310565b5f60405180830381865afa1580156200b21a573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200b2439190810190620100c2565b60408501525b846001600160a01b03166349c4fac882865f01516040516020016200b26f919062010363565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016200b29d929190620103c3565b5f60405180830381865afa1580156200b2b8573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200b2e19190810190620100c2565b606085015250608083015250949350505050565b60408051600480825260a082019092526060915f9190816020015b60608152602001906001900390816200b3105790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250815f815181106200b373576200b373620103eb565b60200260200101819052506040518060400160405280600381526020017f2d726c0000000000000000000000000000000000000000000000000000000000815250816001815181106200b3ca576200b3ca620103eb565b6020026020010181905250846040516020016200b3e8919062010418565b604051602081830303815290604052816002815181106200b40d576200b40d620103eb565b6020026020010181905250826040516020016200b42b919062010478565b604051602081830303815290604052816003815181106200b450576200b450620103eb565b60200260200101819052505f6200b467826200c45c565b602080820151604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000818501908152825180840184525f80825290860152825180840190935290518252928101929092529192506200b4f8906040805180820182525f80825260209182015281518083019092528451825280850190820152906200d0ec565b6200b51057856040516020016200a7599190620104b2565b9695505050505050565b60a08101516040805180820182525f80825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d90156200b56a565b511590565b6200b6e3578260200151156200b629576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a4016200a79b565b8260c00151156200b6e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a4016200a79b565b6040805160ff80825261200082019092525f91816020015b60608152602001906001900390816200b6fb5790505090505f6040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806200b7589062010531565b935060ff16815181106200b770576200b770620103eb565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e37000000000000000000000000000000000000008152506040516020016200b7c3919062010552565b6040516020818303038152906040528282806200b7e09062010531565b935060ff16815181106200b7f8576200b7f8620103eb565b60200260200101819052506040518060400160405280600681526020017f6465706c6f7900000000000000000000000000000000000000000000000000008152508282806200b8479062010531565b935060ff16815181106200b85f576200b85f620103eb565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d650000000000000000000000000000000000008152508282806200b8ae9062010531565b935060ff16815181106200b8c6576200b8c6620103eb565b602002602001018190525087602001518282806200b8e49062010531565b935060ff16815181106200b8fc576200b8fc620103eb565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e7472616374506174680000000000000000000000000000000000008152508282806200b94b9062010531565b935060ff16815181106200b963576200b963620103eb565b6020908102919091010152875182826200b97d8162010531565b935060ff16815181106200b995576200b995620103eb565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e496400000000000000000000000000000000000000000000008152508282806200b9e49062010531565b935060ff16815181106200b9fc576200b9fc620103eb565b60200260200101819052506200ba12466200d152565b82826200ba1f8162010531565b935060ff16815181106200ba37576200ba37620103eb565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c6500000000000000000000000000000000008152508282806200ba869062010531565b935060ff16815181106200ba9e576200ba9e620103eb565b6020026020010181905250868282806200bab89062010531565b935060ff16815181106200bad0576200bad0620103eb565b60209081029190910101528551156200bc035760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826200bb248162010531565b935060ff16815181106200bb3c576200bb3c620103eb565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d906200bb8e9089906004016200f2a4565b5f60405180830381865afa1580156200bba9573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200bbd29190810190620100c2565b82826200bbdf8162010531565b935060ff16815181106200bbf7576200bbf7620103eb565b60200260200101819052505b8460200151156200bcdf5760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826200bc4f8162010531565b935060ff16815181106200bc67576200bc67620103eb565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806200bcb69062010531565b935060ff16815181106200bcce576200bcce620103eb565b60200260200101819052506200bec1565b6200bd186200b5658660a001516040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6200bdb55760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826200bd5e8162010531565b935060ff16815181106200bd76576200bd76620103eb565b60200260200101819052508460a001516040516020016200bd98919062010418565b6040516020818303038152906040528282806200bcb69062010531565b8460c001511580156200bdf95750604080890151815180830183525f808252602091820152825180840190935281518352908101908201526200bdf790511590565b155b156200bec15760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826200be408162010531565b935060ff16815181106200be58576200be58620103eb565b60200260200101819052506200be6e886200d1f6565b6040516020016200be80919062010418565b6040516020818303038152906040528282806200be9d9062010531565b935060ff16815181106200beb5576200beb5620103eb565b60200260200101819052505b604080860151815180830183525f808252602091820152825180840190935281518352908101908201526200bef590511590565b6200bf955760408051808201909152600b81527f2d2d72656c617965724964000000000000000000000000000000000000000000602082015282826200bf3b8162010531565b935060ff16815181106200bf53576200bf53620103eb565b602002602001018190525084604001518282806200bf719062010531565b935060ff16815181106200bf89576200bf89620103eb565b60200260200101819052505b6060850151156200c0c05760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826200bfe18162010531565b935060ff16815181106200bff9576200bff9620103eb565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e8906024015f60405180830381865afa1580156200c066573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200c08f9190810190620100c2565b82826200c09c8162010531565b935060ff16815181106200c0b4576200c0b4620103eb565b60200260200101819052505b60e085015151156200c1735760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826200c10d8162010531565b935060ff16815181106200c125576200c125620103eb565b60200260200101819052506200c1428560e001515f01516200d152565b82826200c14f8162010531565b935060ff16815181106200c167576200c167620103eb565b60200260200101819052505b60e085015160200151156200c22a5760408051808201909152600a81527f2d2d676173507269636500000000000000000000000000000000000000000000602082015282826200c1c38162010531565b935060ff16815181106200c1db576200c1db620103eb565b60200260200101819052506200c1f98560e00151602001516200d152565b82826200c2068162010531565b935060ff16815181106200c21e576200c21e620103eb565b60200260200101819052505b60e085015160400151156200c2e15760408051808201909152600e81527f2d2d6d6178466565506572476173000000000000000000000000000000000000602082015282826200c27a8162010531565b935060ff16815181106200c292576200c292620103eb565b60200260200101819052506200c2b08560e00151604001516200d152565b82826200c2bd8162010531565b935060ff16815181106200c2d5576200c2d5620103eb565b60200260200101819052505b60e085015160600151156200c3985760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826200c3318162010531565b935060ff16815181106200c349576200c349620103eb565b60200260200101819052506200c3678560e00151606001516200d152565b82826200c3748162010531565b935060ff16815181106200c38c576200c38c620103eb565b60200260200101819052505b5f8160ff1667ffffffffffffffff8111156200c3b8576200c3b86200f6cf565b6040519080825280602002602001820160405280156200c3ed57816020015b60608152602001906001900390816200c3d75790505b5090505f5b8260ff168160ff1610156200c44d57838160ff16815181106200c419576200c419620103eb565b6020026020010151828260ff16815181106200c439576200c439620103eb565b60209081029190910101526001016200c3f2565b5093505050505b949350505050565b6200c48360405180606001604052805f60030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d925f91849163d145736c916200c50a91869101620105ab565b5f60405180830381865afa1580156200c525573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200c54e9190810190620100c2565b90505f6200c55d86836200dd0d565b90505f846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b81526004016200c58e91906200f14f565b5f604051808303815f875af11580156200c5aa573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200c5d39190810190620105f3565b805190915060030b158015906200c5ed5750602081015151155b80156200c5fd5750604081015151155b156200b51057815f815181106200c618576200c618620103eb565b60200260200101516040516020016200a7599190620106ae565b60605f6200c666856040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925286518252808701908201529091506200c69e9082905b906200de75565b156200c805575f6200c71f826200c718846200c7116200c6e48a6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925282518252918201519181019190915290565b906200de9d565b906200df01565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200c7849082906200de75565b156200c7f057604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200c7ed905b82906200df8d565b90505b6200c7fb816200dfb4565b925050506200a564565b82156200c8215784846040516020016200a7599291906201088d565b505060408051602081019091525f81526200a564565b5f808251602084015ff09392505050565b6200ad65828260016200c951565b60408051600481526024810182526020810180516001600160e01b03167fad3cb1cc0000000000000000000000000000000000000000000000000000000017905290516060915f9182916001600160a01b038616916200c8b791906200f552565b5f60405180830381855afa9150503d805f81146200c8f1576040519150601f19603f3d011682016040523d82523d5f602084013e6200c8f6565b606091505b50915091508180156200c90a575060208151115b156200c92657808060200190518101906200c4549190620100c2565b505060408051602081019091525f815292915050565b5f6200c94983836200e01f565b159392505050565b8160a00151156200c96157505050565b5f6200c96f8484846200e110565b90505f6200c97d826200c45c565b602081015181519192509060030b1580156200ca1a5750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200ca1a906040805180820182525f808252602091820152815180830190925284518252808501908201526200c697565b156200ca2857505050505050565b604082015151156200ca4b5781604001516040516020016200a759919062010918565b806040516020016200a759919062010971565b60605f6200ca92836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200caf8905b82906200d0ec565b156200cb6b57604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200a564906200cb659083906200e704565b6200dfb4565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200cbce905b82906200e796565b6001036200cc9f57604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200cc36906200c7e5565b50604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200a564906200cb65905b83906200df8d565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200ccff906200caf0565b156200ce3e57604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201819052845180860190955292518452830152906200cd699083906200e83c565b90505f81600183516200cd7d91906200f383565b815181106200cd90576200cd90620103eb565b602002602001015190506200ce356200cb656200ce086040518060400160405280600581526020017f2e6a736f6e0000000000000000000000000000000000000000000000000000008152506040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f80825260209182015281518083019092528551825280860190820152906200e704565b95945050505050565b826040516020016200a7599190620109ca565b50919050565b60605f6200ce8b836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200ceee906200caf0565b156200ceff576200a564816200dfb4565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200cf5f906200cbc6565b6001036200cfcc57604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200a564906200cb65906200cc97565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d02c906200caf0565b156200ce3e57604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201819052845180860190955292518452830152906200d0969083906200e83c565b90506001815111156200d0d85780600282516200d0b491906200f383565b815181106200d0c7576200d0c7620103eb565b602002602001015192505050919050565b50826040516020016200a7599190620109ca565b805182515f9111156200d10157505f6200a44f565b8151835160208501515f92916200d118916200f28e565b6200d12491906200f383565b9050826020015181036200d13d5760019150506200a44f565b82516020840151819020912014905092915050565b60605f6200d160836200e8f8565b60010190505f8167ffffffffffffffff8111156200d182576200d1826200f6cf565b6040519080825280601f01601f1916602001820160405280156200d1ad576020820181803683370190505b5090508181016020015b5f19017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846200d1b757509392505050565b604081810151815180830183525f808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e53454400000000000000000000000000000000000000000000818401908152855180870187528381528401929092528451808601909552518452908301526060916200d283905b82906200c93c565b156200d2c457505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d324906200d27b565b156200d36557505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d3c5906200d27b565b156200d40657505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d466906200d27b565b806200d4cd5750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d4cd906200d27b565b156200d50e57505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d56e906200d27b565b806200d5d55750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d5d5906200d27b565b156200d61657505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d676906200d27b565b806200d6dd5750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d6dd906200d27b565b156200d71e57505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d77e906200d27b565b806200d7e55750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d7e5906200d27b565b156200d82657505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d886906200d27b565b156200d8c757505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d927906200d27b565b156200d96857505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d9c8906200d27b565b156200da0957505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200da69906200d27b565b156200daaa57505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200db0a906200d27b565b156200db4b57505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200dbab906200d27b565b806200dc125750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200dc12906200d27b565b156200dc5357505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200dcb3906200d27b565b156200dcf457505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b604080840151845191516200a759929060200162010a9c565b6060805f5b84518110156200dda357818582815181106200dd32576200dd32620103eb565b60200260200101516040516020016200dd4d9291906200ff9c565b6040516020818303038152906040529150600185516200dd6e91906200f383565b81146200dd9a57816040516020016200dd88919062010bee565b60405160208183030381529060405291505b6001016200dd12565b50604080516003808252608082019092525f91816020015b60608152602001906001900390816200ddbb57905050905083815f815181106200dde9576200dde9620103eb565b60200260200101819052506040518060400160405280600281526020017f2d63000000000000000000000000000000000000000000000000000000000000815250816001815181106200de40576200de40620103eb565b602002602001018190525081816002815181106200de62576200de62620103eb565b6020908102919091010152949350505050565b60208083015183518351928401515f936200de9492918491906200e9e0565b14159392505050565b604080518082019091525f80825260208201525f6200decd845f01518560200151855f015186602001516200eb13565b90508360200151816200dee191906200f383565b845185906200def29083906200f383565b90525060208401525090919050565b604080518082019091525f80825260208201528151835110156200df275750816200a44f565b60208083015190840151600191146200df4f5750815160208481015190840151829020919020145b80156200df85578251845185906200df699083906200f383565b90525082516020850180516200df819083906200f28e565b9052505b509192915050565b604080518082019091525f80825260208201526200dfad8383836200ec53565b5092915050565b60605f825f015167ffffffffffffffff8111156200dfd6576200dfd66200f6cf565b6040519080825280601f01601f1916602001820160405280156200e001576020820181803683370190505b5090505f6020820190506200dfad818560200151865f01516200ed07565b815181515f91908111156200e032575081515b602080850151908401515f5b838110156200e0ff57825182518082146200e0c8575f1960208710156200e0a5576001846200e06f8960206200f383565b6200e07b91906200f28e565b6200e0889060086200f92f565b6200e09590600262010d24565b6200e0a191906200f383565b1990505b81811683821681810391146200e0c55797506200a44f9650505050505050565b50505b6200e0d56020866200f28e565b94506200e0e46020856200f28e565b935050506020816200e0f791906200f28e565b90506200e03e565b50845186516200b510919062010d31565b60605f6200e11d6200ad69565b6040805160ff80825261200082019092529192505f9190816020015b60608152602001906001900390816200e1395790505090505f6040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806200e1969062010531565b935060ff16815181106200e1ae576200e1ae620103eb565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e33000000000000000000000000000000000000000000000000008152506040516020016200e201919062010d53565b6040516020818303038152906040528282806200e21e9062010531565b935060ff16815181106200e236576200e236620103eb565b60200260200101819052506040518060400160405280600881526020017f76616c69646174650000000000000000000000000000000000000000000000008152508282806200e2859062010531565b935060ff16815181106200e29d576200e29d620103eb565b6020026020010181905250826040516020016200e2bb919062010478565b6040516020818303038152906040528282806200e2d89062010531565b935060ff16815181106200e2f0576200e2f0620103eb565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e7472616374000000000000000000000000000000000000000000008152508282806200e33f9062010531565b935060ff16815181106200e357576200e357620103eb565b60200260200101819052506200e36e87846200ed8f565b82826200e37b8162010531565b935060ff16815181106200e393576200e393620103eb565b6020908102919091010152855151156200e44b5760408051808201909152600b81527f2d2d7265666572656e6365000000000000000000000000000000000000000000602082015282826200e3e88162010531565b935060ff16815181106200e400576200e400620103eb565b60200260200101819052506200e41a865f0151846200ed8f565b82826200e4278162010531565b935060ff16815181106200e43f576200e43f620103eb565b60200260200101819052505b8560800151156200e4c05760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b0000000000000000602082015282826200e4978162010531565b935060ff16815181106200e4af576200e4af620103eb565b60200260200101819052506200e52c565b84156200e52c5760408051808201909152601281527f2d2d726571756972655265666572656e63650000000000000000000000000000602082015282826200e5088162010531565b935060ff16815181106200e520576200e520620103eb565b60200260200101819052505b604086015151156200e5d35760408051808201909152600d81527f2d2d756e73616665416c6c6f7700000000000000000000000000000000000000602082015282826200e5798162010531565b935060ff16815181106200e591576200e591620103eb565b602002602001018190525085604001518282806200e5af9062010531565b935060ff16815181106200e5c7576200e5c7620103eb565b60200260200101819052505b8560600151156200e6435760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d6573000000000000000000000000602082015282826200e61f8162010531565b935060ff16815181106200e637576200e637620103eb565b60200260200101819052505b5f8160ff1667ffffffffffffffff8111156200e663576200e6636200f6cf565b6040519080825280602002602001820160405280156200e69857816020015b60608152602001906001900390816200e6825790505b5090505f5b8260ff168160ff1610156200e6f857838160ff16815181106200e6c4576200e6c4620103eb565b6020026020010151828260ff16815181106200e6e4576200e6e4620103eb565b60209081029190910101526001016200e69d565b50979650505050505050565b604080518082019091525f80825260208201528151835110156200e72a5750816200a44f565b8151835160208501515f92916200e741916200f28e565b6200e74d91906200f383565b602084015190915060019082146200e76f575082516020840151819020908220145b80156200e78d578351855186906200e7899083906200f383565b9052505b50929392505050565b5f80825f01516200e7b8855f01518660200151865f015187602001516200eb13565b6200e7c491906200f28e565b90505b835160208501516200e7da91906200f28e565b81116200dfad57816200e7ed8162010d86565b925050825f01516200e8288560200151836200e80a91906200f383565b86516200e81891906200f383565b83865f015187602001516200eb13565b6200e83491906200f28e565b90506200e7c7565b60605f6200e84b84846200e796565b6200e8589060016200f28e565b67ffffffffffffffff8111156200e873576200e8736200f6cf565b6040519080825280602002602001820160405280156200e8a857816020015b60608152602001906001900390816200e8925790505b5090505f5b81518110156200e8f0576200e8c76200cb6586866200df8d565b8282815181106200e8dc576200e8dc620103eb565b60209081029190910101526001016200e8ad565b509392505050565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106200e941577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106200e96e576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106200e98d57662386f26fc10000830492506010015b6305f5e10083106200e9a6576305f5e100830492506008015b61271083106200e9bb57612710830492506004015b606483106200e9ce576064830492506002015b600a83106200a44f5760010192915050565b5f808584116200eb0957602084116200eaa9575f84156200ea345760016200ea0a8660206200f383565b6200ea179060086200f92f565b6200ea2490600262010d24565b6200ea3091906200f383565b1990505b83518116856200ea4589896200f28e565b6200ea5191906200f383565b805190935082165b8181146200ea91578784116200ea7657879450505050506200c454565b836200ea828162010da1565b9450508284511690506200ea59565b6200ea9d87856200f28e565b9450505050506200c454565b8383206200eab885886200f383565b6200eac490876200f28e565b91505b8582106200eb07578482208082036200eaf1576200eae686846200f28e565b93505050506200c454565b6200eafe6001846200f383565b9250506200eac7565b505b5092949350505050565b5f83818685116200ec3c57602085116200ebe2575f85156200eb685760016200eb3e8760206200f383565b6200eb4b9060086200f92f565b6200eb5890600262010d24565b6200eb6491906200f383565b1990505b845181165f876200eb7a8b8b6200f28e565b6200eb8691906200f383565b855190915083165b8281146200ebd3578186106200ebb8576200ebaa8b8b6200f28e565b96505050505050506200c454565b856200ebc48162010d86565b9650508386511690506200eb8e565b8596505050505050506200c454565b508383205f905b6200ebf586896200f383565b82116200ec3a578583208082036200ec1457839450505050506200c454565b6200ec216001856200f28e565b93505081806200ec319062010d86565b9250506200ebe9565b505b6200ec4887876200f28e565b979650505050505050565b604080518082019091525f80825260208201525f6200ec83855f01518660200151865f015187602001516200eb13565b6020808701805191860191909152519091506200eca190826200f383565b8352845160208601516200ecb691906200f28e565b81036200ecc6575f85526200ecfe565b835183516200ecd691906200f28e565b855186906200ece79083906200f383565b90525083516200ecf890826200f28e565b60208601525b50909392505050565b602081106200ed4757815183526200ed216020846200f28e565b92506200ed306020836200f28e565b91506200ed3f6020826200f383565b90506200ed07565b5f1981156200ed7c5760016200ed5f8360206200f383565b6200ed6d9061010062010d24565b6200ed7991906200f383565b90505b9151835183169219169190911790915250565b60605f6200ed9e84846200ae3d565b80516020808301516040519394506200edba9390910162010db9565b60405160208183030381529060405291505092915050565b610b098062010df983390190565b615048806201190283390190565b6109e8806201694a83390190565b61102e806201733283390190565b611fc0806201836083390190565b6040518060e001604052806060815260200160608152602001606081526020015f151581526020015f151581526020015f151581526020016200ee5a6200ee5f565b905290565b6040518061010001604052805f151581526020015f15158152602001606081526020015f801916815260200160608152602001606081526020015f151581526020016200ee5a60405180608001604052805f81526020015f81526020015f81526020015f81525090565b602080825282518282018190525f918401906040840190835b818110156200ef0b5783516001600160a01b03168352602093840193909201916001016200eee2565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200f043577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b8801810191908801905f5b818110156200f028577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a85030183526200f0118486516200ef16565b60209586019590945092909201916001016200efd4565b5091975050506020948501949290920191506001016200ef6a565b50929695505050505050565b5f8151808452602084019350602083015f5b828110156200f0a35781517fffffffff00000000000000000000000000000000000000000000000000000000168652602095860195909101906001016200f061565b5093949350505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200f043577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051604087526200f11a60408801826200ef16565b90506020820151915086810360208801526200f13781836200f04f565b9650505060209384019391909101906001016200f0d3565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200f043577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184526200f1b28583516200ef16565b945060209384019391909101906001016200f175565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200f043577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b03815116865260208101519050604060208701526200f24a60408701826200f04f565b95505060209384019391909101906001016200f1ee565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156200a44f576200a44f6200f261565b602081525f6200a56460208301846200ef16565b6001600160a01b0381511682526020810151151560208301526001600160a01b0360408201511660408301525f606082015160a060608501526200f30060a08501826200ef16565b608093840151949093019390935250919050565b60c081525f6200f32860c08301886200ef16565b6001600160a01b038716602084015282810360408401526200f34b81876200ef16565b85546060850152600186015460ff1615156080850152905082810360a08401526200f37781856200f2b8565b98975050505050505050565b818103818111156200a44f576200a44f6200f261565b600181811c908216806200f3ae57607f821691505b6020821081036200ce51577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f81546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b0360018301541660408401526002820160a060608501525f81546200f430816200f399565b8060a0880152600182165f81146200f45157600181146200f48c576200f4bf565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660c089015260c082151560051b89010193506200f4bf565b845f5260205f205f5b838110156200f4b65781548a820160c001526001909101906020016200f495565b890160c0019450505b50505060038401546080860152809250505092915050565b60c081525f6200f4eb60c08301886200ef16565b6001600160a01b038716602084015282810360408401526200f50e81876200ef16565b85546060850152600186015460ff161515608085015290505b82810360a08401526200f37781856200f3e6565b5f81518060208401855e5f93019283525090919050565b5f6200a56482846200f53b565b6001600160a01b0383168152604060208201525f6200c45460408301846200ef16565b80516001600160a01b03811681146200f599575f80fd5b919050565b5f602082840312156200f5af575f80fd5b6200a564826200f582565b610100815260056101008201527f544f4b454e000000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f544b4e00000000000000000000000000000000000000000000000000000000006101608201525f6101808201905060ff88166040830152866060830152600386106200f672577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b8560808301528460a08301526200f69460c08301856001600160a01b03169052565b6001600160a01b03831660e0830152979650505050505050565b5f602082840312156200f6bf575f80fd5b815180151581146200a564575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b601f821115620052bf57805f5260205f20601f840160051c810160208510156200f7235750805b601f840160051c820191505b8181101562000b6f575f81556001016200f72f565b815167ffffffffffffffff8111156200f761576200f7616200f6cf565b6200f779816200f77284546200f399565b846200f6fc565b6020601f8211600181146200f7ae575f83156200f7965750848201515b5f19600385901b1c1916600184901b17845562000b6f565b5f84815260208120601f198516915b828110156200f7df57878501518255602094850194600190920191016200f7bd565b50848210156200f7fd57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b60e081525f6200f82060e08301896200ef16565b8760208401526001600160a01b038716604084015282810360608401526200f84981876200ef16565b85546080850152600186015460ff16151560a085015290505b82810360c08401526200f87681856200f3e6565b9998505050505050505050565b5f602082840312156200f894575f80fd5b5051919050565b60e081525f6200f8af60e08301896200ef16565b8760208401526001600160a01b038716604084015282810360608401526200f8d881876200ef16565b855160808501526020860151151560a085015290506200f862565b8481526001600160a01b0384166020820152608060408201525f6200f91c60808301856200ef16565b9050821515606083015295945050505050565b80820281158282048414176200a44f576200a44f6200f261565b838152606060208201525f6200f96360608301856200ef16565b82810360408401526200b51081856200ef16565b61012081525f6200f98d61012083018b6200ef16565b6001600160a01b038a16602084015288604084015287606084015286608084015282810360a08401526200f9c281876200ef16565b855460c0850152600186015460ff16151560e085015290505b8281036101008401526200a87f81856200f3e6565b60c081525f6200fa0460c08301886200ef16565b86602084015282810360408401526200f50e81876200ef16565b61012081525f6200fa3461012083018a6200ef16565b6001600160a01b03891660208401528760408401528660608401528560808401528281038060a08501525f82526200fa7b60c0850187805182526020908101511515910152565b60208101610100850152506200fa9560208201856200f3e6565b9a9950505050505050505050565b606081525f6200fab760608301866200ef16565b84602084015282810360408401526200b51081856200f3e6565b61012081525f6200fae761012083018b6200ef16565b6001600160a01b038a16602084015288604084015287606084015286608084015282810360a08401526200fb1c81876200ef16565b855160c08501526020860151151560e085015290506200f9db565b608081525f6200fb4b60808301876200ef16565b8560208401526001600160a01b038516604084015282810360608401526200ec4881856200f3e6565b60a081525f6200fb8860a08301886200ef16565b8660208401526001600160a01b038616604084015284606084015282810360808401526200f37781856200f3e6565b5f826200fbeb577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b60a081525f6200fc0460a08301876200ef16565b82810360208401526200fc1881876200ef16565b85546040850152600186015460ff1615156060850152905082810360808401526200ec4881856200f3e6565b60c081525f6200fc5860c08301886200ef16565b6001600160a01b038716602084015282810360408401526200fc7b81876200ef16565b8551606085015260208601511515608085015290506200f527565b60c081525f6200fcaa60c08301886200ef16565b86602084015282810360408401526200f34b81876200ef16565b608081525f6200fcd860808301876200ef16565b8560208401526001600160a01b038516604084015282810360608401526200ec4881856200f2b8565b60a081525f6200fd1560a08301886200ef16565b8660208401526001600160a01b038616604084015284606084015282810360808401526200f37781856200f2b8565b5f80604083850312156200fd56575f80fd5b6200fd61836200f582565b6020939093015192949293505050565b606081525f6200fd8560608301866200ef16565b84602084015282810360408401526200b51081856200f2b8565b60c081525f6200fdb360c08301886200ef16565b86602084015282810360408401526200fc7b81876200ef16565b60e081525f6200fde160e08301896200ef16565b8760208401526001600160a01b038716604084015282810360608401526200fe0a81876200ef16565b85546080850152600186015460ff16151560a0850152905082810360c08401526200f87681856200f2b8565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f6200fe69601a8301856200f53b565b7f3a2000000000000000000000000000000000000000000000000000000000000081526200a56060028201856200f53b565b6040516060810167ffffffffffffffff811182821017156200fec1576200fec16200f6cf565b60405290565b5f8067ffffffffffffffff8411156200fee4576200fee46200f6cf565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff821117156200ff16576200ff166200f6cf565b6040528381529050808284018510156200ff2e575f80fd5b8383602083015e5f60208583010152509392505050565b5f82601f8301126200ff55575f80fd5b6200a564838351602085016200fec7565b5f602082840312156200ff77575f80fd5b815167ffffffffffffffff8111156200ff8e575f80fd5b6200a44b848285016200ff45565b5f6200c4546200ffad83866200f53b565b846200f53b565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f6200ffe7601a8301856200f53b565b7f207573696e6720636f6e7374727563746f72206461746120220000000000000081526201001960198201856200f53b565b7f2200000000000000000000000000000000000000000000000000000000000000815260010195945050505050565b6001600160a01b03841681526001600160a01b0383166020820152606060408201525f6200ce3560608301846200ef16565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201525f6200a56460808301846200ef16565b5f60208284031215620100d3575f80fd5b815167ffffffffffffffff811115620100ea575f80fd5b8201601f81018413620100fb575f80fd5b6200a44b848251602084016200fec7565b5f6201011982876200f53b565b7f2f0000000000000000000000000000000000000000000000000000000000000081526201014b60018201876200f53b565b90507f2f0000000000000000000000000000000000000000000000000000000000000081526201017f60018201866200f53b565b90507f2f000000000000000000000000000000000000000000000000000000000000008152620101b360018201856200f53b565b7f2e6a736f6e0000000000000000000000000000000000000000000000000000008152600501979650505050505050565b604081525f620101f860408301846200ef16565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081525f6201026a601f8301846200f53b565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f81527f6d6c00000000000000000000000000000000000000000000000000000000000060208201526022019392505050565b604081525f620102d160408301846200ef16565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b604081525f6201032460408301846200ef16565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081525f6201039660148301846200f53b565b7f275d2e6b656363616b32353600000000000000000000000000000000000000008152600c019392505050565b604081525f620103d760408301856200ef16565b82810360208401526200a56081856200ef16565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f220000000000000000000000000000000000000000000000000000000000000081525f6201044b60018301846200f53b565b7f220000000000000000000000000000000000000000000000000000000000000081526001019392505050565b5f6201048582846200f53b565b7f2f6275696c642d696e666f0000000000000000000000000000000000000000008152600b019392505050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201525f6200a564604b8301846200f53b565b5f60ff821660ff81036201054957620105496200f261565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201525f6200a56460298301846200f53b565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201525f6200a56460808301846200ef16565b5f6020828403121562010604575f80fd5b815167ffffffffffffffff8111156201061b575f80fd5b8201606081850312156201062d575f80fd5b620106376200fe9b565b81518060030b811462010648575f80fd5b8152602082015167ffffffffffffffff81111562010664575f80fd5b62010672868285016200ff45565b602083015250604082015167ffffffffffffffff81111562010692575f80fd5b620106a0868285016200ff45565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201525f6201070760218301846200f53b565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657481527f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960208201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560408201527f616c69666965642070617468206f66207468652062617368206578656375746160608201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960808201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60a08201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960c08201527f6f75722070726f6a65637420287573696e6720666f727761726420736c61736860e08201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101008201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061012082015261013b019392505050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201525f620108e660218301856200f53b565b7f2720696e206f75747075743a200000000000000000000000000000000000000081526200a560600d8201856200f53b565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201525f6200a56460298301846200f53b565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201525f6200a56460228301846200f53b565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081525f620109fd600e8301846200f53b565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e7472616381527f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e7360208201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e747260408201527f6163742e6a736f6e00000000000000000000000000000000000000000000000060608201526068019392505050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081525f62010acf60188301856200f53b565b7f20696e2000000000000000000000000000000000000000000000000000000000815262010b0160048201856200f53b565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f727465642081527f6c6963656e736520666f7220626c6f636b206578706c6f72657220766572696660208201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f60408201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c60608201527f206f7220736574207468652060736b69704c6963656e73655479706560206f7060808201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060a082015260b70195945050505050565b5f62010bfb82846200f53b565b7f200000000000000000000000000000000000000000000000000000000000000081526001019392505050565b6001815b600184111562010c695780850481111562010c4b5762010c4b6200f261565b600184161562010c5a57908102905b60019390931c92800262010c2c565b935093915050565b5f8262010c81575060016200a44f565b8162010c8f57505f6200a44f565b816001811462010ca8576002811462010cb35762010cd3565b60019150506200a44f565b60ff84111562010cc75762010cc76200f261565b50506001821b6200a44f565b5060208310610133831016604e8410600b841016171562010cf8575081810a6200a44f565b62010d065f19848462010c28565b805f190482111562010d1c5762010d1c6200f261565b029392505050565b5f6200a564838362010c71565b8181035f8312801583831316838312821617156200dfad576200dfad6200f261565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081525f6200a564601c8301846200f53b565b5f5f19820362010d9a5762010d9a6200f261565b5060010190565b5f8162010db25762010db26200f261565b505f190190565b5f62010dc682856200f53b565b7f3a0000000000000000000000000000000000000000000000000000000000000081526200a56060018201856200f53b56fe60c0604052600d60809081526c2bb930b83832b21022ba3432b960991b60a0525f9061002b908261010b565b506040805180820190915260048152630ae8aa8960e31b6020820152600190610054908261010b565b506002805460ff1916601217905534801561006d575f80fd5b506101c5565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061009b57607f821691505b6020821081036100b957634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561010657805f5260205f20601f840160051c810160208510156100e45750805b601f840160051c820191505b81811015610103575f81556001016100f0565b50505b505050565b81516001600160401b0381111561012457610124610073565b610138816101328454610087565b846100bf565b6020601f82116001811461016a575f83156101535750848201515b5f19600385901b1c1916600184901b178455610103565b5f84815260208120601f198516915b828110156101995787850151825560209485019460019092019101610179565b50848210156101b657868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610937806101d25f395ff3fe6080604052600436106100bb575f3560e01c8063313ce56711610071578063a9059cbb1161004c578063a9059cbb146101eb578063d0e30db01461020a578063dd62ed3e14610212575f80fd5b8063313ce5671461018157806370a08231146101ac57806395d89b41146101d7575f80fd5b806318160ddd116100a157806318160ddd1461012757806323b872dd146101435780632e1a7d4d14610162575f80fd5b806306fdde03146100ce578063095ea7b3146100f8575f80fd5b366100ca576100c8610248565b005b5f80fd5b3480156100d9575f80fd5b506100e26102a2565b6040516100ef919061071f565b60405180910390f35b348015610103575f80fd5b5061011761011236600461079a565b61032d565b60405190151581526020016100ef565b348015610132575f80fd5b50475b6040519081526020016100ef565b34801561014e575f80fd5b5061011761015d3660046107c2565b6103a6565b34801561016d575f80fd5b506100c861017c3660046107fc565b610628565b34801561018c575f80fd5b5060025461019a9060ff1681565b60405160ff90911681526020016100ef565b3480156101b7575f80fd5b506101356101c6366004610813565b60036020525f908152604090205481565b3480156101e2575f80fd5b506100e26106ff565b3480156101f6575f80fd5b5061011761020536600461079a565b61070c565b6100c8610248565b34801561021d575f80fd5b5061013561022c36600461082c565b600460209081525f928352604080842090915290825290205481565b335f908152600360205260408120805434929061026690849061088a565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b5f80546102ae9061089d565b80601f01602080910402602001604051908101604052809291908181526020018280546102da9061089d565b80156103255780601f106102fc57610100808354040283529160200191610325565b820191905f5260205f20905b81548152906001019060200180831161030857829003601f168201915b505050505081565b335f81815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103949086815260200190565b60405180910390a35060015b92915050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260036020526040812054821115610412576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201525f60248201526044015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff84163314801590610487575073ffffffffffffffffffffffffffffffffffffffff84165f9081526004602090815260408083203384529091529020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14155b156105435773ffffffffffffffffffffffffffffffffffffffff84165f9081526004602090815260408083203384529091529020548211156104fe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201525f6024820152604401610409565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526004602090815260408083203384529091528120805484929061053d9084906108ee565b90915550505b73ffffffffffffffffffffffffffffffffffffffff84165f90815260036020526040812080548492906105779084906108ee565b909155505073ffffffffffffffffffffffffffffffffffffffff83165f90815260036020526040812080548492906105b090849061088a565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161061691815260200190565b60405180910390a35060019392505050565b335f90815260036020526040902054811115610679576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201525f6024820152604401610409565b335f90815260036020526040812080548392906106979084906108ee565b9091555050604051339082156108fc029083905f818181858888f193505050501580156106c6573d5f803e3d5ffd5b5060405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a250565b600180546102ae9061089d565b5f6107183384846103a6565b9392505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610795575f80fd5b919050565b5f80604083850312156107ab575f80fd5b6107b483610772565b946020939093013593505050565b5f805f606084860312156107d4575f80fd5b6107dd84610772565b92506107eb60208501610772565b929592945050506040919091013590565b5f6020828403121561080c575f80fd5b5035919050565b5f60208284031215610823575f80fd5b61071882610772565b5f806040838503121561083d575f80fd5b61084683610772565b915061085460208401610772565b90509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156103a0576103a061085d565b600181811c908216806108b157607f821691505b6020821081036108e8577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b818103818111156103a0576103a061085d56fea2646970667358221220f82a6621bc6ae2f40b7ff1dde0e016bd4b523e6e1df6d8c4401566e37294755f64736f6c634300081a003360a060405230608052348015610013575f80fd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614f4f6100f95f395f81816127fe0152818161282701526129d90152614f4f5ff3fe608060405260043610610291575f3560e01c80638f28397011610165578063bd8fde1c116100c6578063d547741f1161007c578063e9d6c5ba11610062578063e9d6c5ba146107f7578063f354b31f14610828578063f851a44014610847575f80fd5b8063d547741f146107a5578063e63ab1e9146107c4575f80fd5b8063c1bd469f116100ac578063c1bd469f14610746578063cc5ad8b614610767578063d3523ea214610786575f80fd5b8063bd8fde1c146106f4578063c0c53b8b14610727575f80fd5b8063a217fddf1161011b578063a8f2cb9611610101578063a8f2cb961461066e578063aa808c061461068d578063ad3cb1cc146106ac575f80fd5b8063a217fddf14610645578063a3ebd14c14610658575f80fd5b806391d148541161014b57806391d14854146105a057806394cc8683146106035780639ca220dd14610624575f80fd5b80638f283970146105625780639060bda914610581575f80fd5b80633f4ba83a1161020f578063631d62e4116101c55780637066b18d116101ab5780637066b18d146104f5578063804ea334146105215780638456cb591461054e575f80fd5b8063631d62e4146104b75780636e9e2d3f146104d6575f80fd5b806352d1902d116101f557806352d1902d146104405780635c975abb146104545780635cf92c9f1461048a575f80fd5b80633f4ba83a146104195780634f1ef2861461042d575f80fd5b80632259e9e5116102645780632f2ff15d1161024a5780632f2ff15d146103bc5780633500c24b146103db57806336568abe146103fa575f80fd5b80632259e9e514610342578063248a9ca314610361575f80fd5b806301ffc9a7146102955780630c63109e146102c957806310d29b9e1461030057806318d3ce9614610321575b5f80fd5b3480156102a0575f80fd5b506102b46102af366004613eec565b610865565b60405190151581526020015b60405180910390f35b3480156102d4575f80fd5b506001546102e8906001600160a01b031681565b6040516001600160a01b0390911681526020016102c0565b34801561030b575f80fd5b5061031f61031a366004613f7d565b6108fd565b005b34801561032c575f80fd5b506103356109b6565b6040516102c09190614006565b34801561034d575f80fd5b5061031f61035c3660046140c5565b610c47565b34801561036c575f80fd5b506103ae61037b36600461413e565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016102c0565b3480156103c7575f80fd5b5061031f6103d6366004614169565b610cd9565b3480156103e6575f80fd5b5061031f6103f5366004614197565b610d22565b348015610405575f80fd5b5061031f610414366004614169565b610eb4565b348015610424575f80fd5b5061031f610f05565b61031f61043b3660046141df565b610f1a565b34801561044b575f80fd5b506103ae610f39565b34801561045f575f80fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166102b4565b348015610495575f80fd5b506104a96104a43660046142a6565b610f67565b6040516102c09291906142ee565b3480156104c2575f80fd5b5061031f6104d13660046140c5565b61105e565b3480156104e1575f80fd5b5061031f6104f0366004614310565b611104565b348015610500575f80fd5b5061051461050f3660046142a6565b6111c3565b6040516102c091906143e1565b34801561052c575f80fd5b5061054061053b36600461413e565b611288565b6040516102c09291906143f3565b348015610559575f80fd5b5061031f61133d565b34801561056d575f80fd5b5061031f61057c366004614197565b61136f565b34801561058c575f80fd5b5061031f61059b366004614414565b6114bd565b3480156105ab575f80fd5b506102b46105ba366004614169565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561060e575f80fd5b5061061761154b565b6040516102c09190614440565b34801561062f575f80fd5b506106386115a1565b6040516102c09190614482565b348015610650575f80fd5b506103ae5f81565b348015610663575f80fd5b506103ae6207a12081565b348015610679575f80fd5b5061031f61068836600461452d565b61175a565b348015610698575f80fd5b506102e86106a73660046142a6565b6117da565b3480156106b7575f80fd5b506105146040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b3480156106ff575f80fd5b506103ae7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa381565b348015610732575f80fd5b5061031f61074136600461459c565b611819565b348015610751575f80fd5b5061075a611bc1565b6040516102c091906145e4565b348015610772575f80fd5b50600b546102e8906001600160a01b031681565b348015610791575f80fd5b506105146107a03660046140c5565b611ebb565b3480156107b0575f80fd5b5061031f6107bf366004614169565b611f9f565b3480156107cf575f80fd5b506103ae7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b348015610802575f80fd5b50610816610811366004614197565b611fe2565b6040516102c0969594939291906146d9565b348015610833575f80fd5b5061031f610842366004614736565b61222d565b348015610852575f80fd5b505f546102e8906001600160a01b031681565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806108f757507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3610927816122c9565b61092f6122d3565b61093b85858585612331565b6109478585858561247f565b7f6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c60075f8781526020019081526020015f2085856040516109899291906147de565b90815260200160405180910390206001016040516109a791906148bd565b60405180910390a15050505050565b6004546060908067ffffffffffffffff8111156109d5576109d56141b2565b604051908082528060200260200182016040528015610a3157816020015b610a1e60405180608001604052805f1515815260200160608152602001606081526020015f81525090565b8152602001906001900390816109f35790505b5091505f5b81811015610c42575f60048281548110610a5257610a526148cf565b905f5260205f2090600202016040518060400160405290815f8201548152602001600182018054610a82906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054610aae906147ed565b8015610af95780601f10610ad057610100808354040283529160200191610af9565b820191905f5260205f20905b815481529060010190602001808311610adc57829003601f168201915b50505050508152505090505f815f015190505f82602001519050604051806080016040528060075f8581526020019081526020015f2083604051610b3d9190614913565b90815260408051602092819003830190205460ff16151583525f8681526007835281902090519290910191610b73908590614913565b90815260200160405180910390206001018054610b8f906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054610bbb906147ed565b8015610c065780601f10610bdd57610100808354040283529160200191610c06565b820191905f5260205f20905b815481529060010190602001808311610be957829003601f168201915b5050505050815260200182815260200183815250868581518110610c2c57610c2c6148cf565b6020908102919091010152505050600101610a36565b505090565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3610c71816122c9565b610c796122d3565b610c868686868686612501565b610c938686868686612594565b857f40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee63486868686604051610cc99493929190614947565b60405180910390a2505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610d12816122c9565b610d1c8383612611565b50505050565b5f610d2c816122c9565b6001600160a01b038216610d6c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d967ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa383612611565b50610dc17f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a83612611565b50600154610df9907ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3906001600160a01b03166126dd565b50600154610e31907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a906001600160a01b03166126dd565b50600154604080516001600160a01b03928316815291841660208301527f6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279910160405180910390a150600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6001600160a01b0381163314610ef6576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f0082826126dd565b505050565b5f610f0f816122c9565b610f17612781565b50565b610f226127f3565b610f2b826128c3565b610f3582826128cd565b5050565b5f610f426129ce565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f83815260076020526040808220905160609190610f8890869086906147de565b908152604080519182900360209081018320545f898152600790925291902060ff909116935090610fbc90869086906147de565b90815260200160405180910390206001018054610fd8906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611004906147ed565b801561104f5780601f106110265761010080835404028352916020019161104f565b820191905f5260205f20905b81548152906001019060200180831161103257829003601f168201915b50505050509050935093915050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3611088816122c9565b6110906122d3565b61109d8686868686612a30565b6110aa8686868686612d0a565b84846040516110ba9291906147de565b6040518091039020867f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce258185856040516110f4929190614978565b60405180910390a3505050505050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361112e816122c9565b6111366122d3565b6111478a8a8a8a8a8a8a8a8a612d87565b6111588a8a8a8a8a8a8a8a8a6130b3565b896001600160a01b031686866040516111729291906147de565b60405180910390207fa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367848a8d8d6040516111af949392919061498b565b60405180910390a350505050505050505050565b606060065f8581526020019081526020015f2060040183836040516111e99291906147de565b90815260200160405180910390208054611202906147ed565b80601f016020809104026020016040519081016040528092919081815260200182805461122e906147ed565b80156112795780601f1061125057610100808354040283529160200191611279565b820191905f5260205f20905b81548152906001019060200180831161125c57829003601f168201915b505050505090505b9392505050565b5f8181526006602052604090206002810154600390910180546001600160a01b0390921691606091906112ba906147ed565b80601f01602080910402602001604051908101604052809291908181526020018280546112e6906147ed565b80156113315780601f1061130857610100808354040283529160200191611331565b820191905f5260205f20905b81548152906001019060200180831161131457829003601f168201915b50505050509050915091565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a611367816122c9565b610f17613144565b5f611379816122c9565b6001600160a01b0382166113b9576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113c35f83612611565b506113ee7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a83612611565b505f805461140591906001600160a01b03166126dd565b505f5461143c907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a906001600160a01b03166126dd565b505f54604080516001600160a01b03928316815291841660208301527f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f910160405180910390a1505f80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa36114e7816122c9565b6114ef6122d3565b6114f9838361319f565b611503838361328d565b604080516001600160a01b038516815283151560208201527f9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8910160405180910390a1505050565b6060600280548060200260200160405190810160405280929190818152602001828054801561159757602002820191905f5260205f20905b815481526020019060010190808311611583575b5050505050905090565b6003546060908067ffffffffffffffff8111156115c0576115c06141b2565b60405190808252806020026020018201604052801561162e57816020015b604080516080810182525f80825260208083018290529282015260608082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092019101816115de5790505b5091505f5b81811015610c42575f6003828154811061164f5761164f6148cf565b5f918252602080832090910154604080516080810182528285526006808552828620805460ff161515835282860185905260028101546001600160a01b03169383019390935294839052939092526003909101805491935060608301916116b5906147ed565b80601f01602080910402602001604051908101604052809291908181526020018280546116e1906147ed565b801561172c5780601f106117035761010080835404028352916020019161172c565b820191905f5260205f20905b81548152906001019060200180831161170f57829003601f168201915b5050505050815250848381518110611746576117466148cf565b602090810291909101015250600101611633565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3611784816122c9565b61178c6122d3565b611799868686868661330f565b6117a686868686866134bf565b857fc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e83604051610cc9911515815260200190565b5f838152600a602052604080822090516117f790859085906147de565b908152604051908190036020019020546001600160a01b031690509392505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156118635750825b90505f8267ffffffffffffffff16600114801561187f5750303b155b90508115801561188d575080155b156118c4576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156119255784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038816158061194257506001600160a01b038716155b8061195457506001600160a01b038616155b1561198b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61199361353c565b61199b61353c565b6119a3613544565b6119ad5f89612611565b506119d87ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa388612611565b50611a037f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a88612611565b50611a2e7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a89612611565b505f80546001600160a01b038a81167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316178355600180548b8316908416178155600b8054928b16929093169190911790915546808352600660208181526040808620805460ff1916909517855580513060601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016818401528151808203601401815260349091019091529290945290925260030190611af190826149fb565b50600280546001818101909255467f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9091018190556003805492830181555f527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b909101558315611bb75784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b6005546060908067ffffffffffffffff811115611be057611be06141b2565b604051908082528060200260200182016040528015611c5b57816020015b611c486040518060e001604052805f151581526020015f6001600160a01b03168152602001606081526020015f815260200160608152602001606081526020015f60ff1681525090565b815260200190600190039081611bfe5790505b5091505f5b81811015610c42575f60058281548110611c7c57611c7c6148cf565b5f9182526020808320909101546001600160a01b0390811680845260088352604093849020845160e081018652815460ff811615158252610100900490931693830193909352600183018054919550919384019190611cda906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611d06906147ed565b8015611d515780601f10611d2857610100808354040283529160200191611d51565b820191905f5260205f20905b815481529060010190602001808311611d3457829003601f168201915b5050505050815260200160028201548152602001600382018054611d74906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611da0906147ed565b8015611deb5780601f10611dc257610100808354040283529160200191611deb565b820191905f5260205f20905b815481529060010190602001808311611dce57829003601f168201915b50505050508152602001600482018054611e04906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611e30906147ed565b8015611e7b5780601f10611e5257610100808354040283529160200191611e7b565b820191905f5260205f20905b815481529060010190602001808311611e5e57829003601f168201915b50505091835250506005919091015460ff166020909101528451859084908110611ea757611ea76148cf565b602090810291909101015250600101611c60565b606060075f8781526020019081526020015f208585604051611ede9291906147de565b90815260200160405180910390206003018383604051611eff9291906147de565b90815260200160405180910390208054611f18906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611f44906147ed565b8015611f8f5780601f10611f6657610100808354040283529160200191611f8f565b820191905f5260205f20905b815481529060010190602001808311611f7257829003601f168201915b5050505050905095945050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611fd8816122c9565b610d1c83836126dd565b6001600160a01b038082165f908152600860209081526040808320815160e081018352815460ff81161515825261010090049095169285019290925260018201805493946060948694869485948794859490939284019190612043906147ed565b80601f016020809104026020016040519081016040528092919081815260200182805461206f906147ed565b80156120ba5780601f10612091576101008083540402835291602001916120ba565b820191905f5260205f20905b81548152906001019060200180831161209d57829003601f168201915b50505050508152602001600282015481526020016003820180546120dd906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054612109906147ed565b80156121545780601f1061212b57610100808354040283529160200191612154565b820191905f5260205f20905b81548152906001019060200180831161213757829003601f168201915b5050505050815260200160048201805461216d906147ed565b80601f0160208091040260200160405190810160405280929190818152602001828054612199906147ed565b80156121e45780601f106121bb576101008083540402835291602001916121e4565b820191905f5260205f20905b8154815290600101906020018083116121c757829003601f168201915b50505091835250506005919091015460ff16602090910152805160808201516060830151604084015160a085015160c090950151939d929c50909a509850919650945092505050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3612257816122c9565b61225f6122d3565b61226e88888888888888613577565b61227d888888888888886136ca565b877faea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c58888888888886040516122b796959493929190614af4565b60405180910390a25050505050505050565b610f17813361374b565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff161561232f576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f8481526006602052604090205460ff16612380576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018590526024015b60405180910390fd5b5f8290036123bd5782826040517ec10cfd000000000000000000000000000000000000000000000000000000008152600401612377929190614978565b5f848152600760205260409081902090516123db90859085906147de565b908152602001604051809103902060010180546123f7906147ed565b90505f03612437578383836040517f2b4f9c8600000000000000000000000000000000000000000000000000000000815260040161237793929190614b3c565b8060075f8681526020019081526020015f2084846040516124599291906147de565b908152604051908190036020019020805491151560ff1990921691909117905550505050565b5f848484846040516024016124979493929190614b55565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f10d29b9e0000000000000000000000000000000000000000000000000000000017905290506124fa816137d7565b5050505050565b5f8581526006602052604090205460ff1661254b576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612377565b818160065f8881526020019081526020015f2060040186866040516125719291906147de565b9081526020016040518091039020918261258c929190614b81565b505050505050565b5f85858585856040516024016125ae959493929190614c79565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f2259e9e500000000000000000000000000000000000000000000000000000000179052905061258c816137d7565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff166126d4575f848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561268a3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506108f7565b5f9150506108f7565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16156126d4575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506108f7565b612789613876565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061288c57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166128807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b1561232f576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610f35816122c9565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612927575060408051601f3d908101601f1916820190925261292491810190614cb1565b60015b612968576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401612377565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc81146129c4576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401612377565b610f0083836138d1565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461232f576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8581526006602052604090205460ff16612a7a576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612377565b5f839003612ab75783836040517ec10cfd000000000000000000000000000000000000000000000000000000008152600401612377929190614978565b5f819003612af1576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f858152600760205260408082209051612b0e90879087906147de565b90815260200160405180910390206001018054612b2a906147ed565b90501115612b6e5784848484846040517f2b192eab000000000000000000000000000000000000000000000000000000008152600401612377959493929190614c79565b600160075f8781526020019081526020015f208585604051612b919291906147de565b9081526040805160209281900383018120805460ff1916941515949094179093555f888152600790925290208391839190612bcf90889088906147de565b90815260200160405180910390206001019182612bed929190614b81565b50838360075f8881526020019081526020015f208686604051612c119291906147de565b90815260200160405180910390206002019182612c2f929190614b81565b506004604051806040016040528087815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920182905250939094525050835460018181018655948252602091829020845160029092020190815590830151929390929083019150612cae90826149fb565b5050508383604051612cc19291906147de565b6040518091039020857f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce25818484604051612cfb929190614978565b60405180910390a35050505050565b5f8585858585604051602401612d24959493929190614c79565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f631d62e400000000000000000000000000000000000000000000000000000000179052905061258c816137d7565b6001600160a01b038916612dc7576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f879003612e30576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f53796d626f6c2063616e6e6f7420626520656d707479000000000000000000006044820152606401612377565b6001600160a01b038981165f9081526008602052604090205461010090041615612e91576040517f63f4ee1f0000000000000000000000000000000000000000000000000000000081526001600160a01b038a166004820152602401612377565b5f6001600160a01b031660098989604051612ead9291906147de565b908152604051908190036020019020546001600160a01b031614612f015787876040517fb295cac9000000000000000000000000000000000000000000000000000000008152600401612377929190614978565b6001600160a01b0389165f818152600860205260409020805460017fffffffffffffffffffffff000000000000000000000000000000000000000000909116610100909302929092178217815501612f5a858783614b81565b506001600160a01b0389165f90815260086020526040902060028101879055600301612f87888a83614b81565b506001600160a01b0389165f90815260086020526040902060058101805460ff191660ff8416179055600401612fbe838583614b81565b5088600a5f8881526020019081526020015f208686604051612fe19291906147de565b90815260200160405180910390205f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555088600989896040516130259291906147de565b90815260405190819003602001902080546001600160a01b039283167fffffffffffffffffffffffff000000000000000000000000000000000000000091821617909155600580546001810182555f919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180549b9092169a16999099179098555050505050505050565b5f8989898989898989896040516024016130d599989796959493929190614cc8565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6e9e2d3f000000000000000000000000000000000000000000000000000000001790529050613138816137d7565b50505050505050505050565b61314c6122d3565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336127d5565b6001600160a01b0382166131df576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038281165f90815260086020526040902054610100900416613263576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5a52433230206e6f7420726567697374657265640000000000000000000000006044820152606401612377565b6001600160a01b03919091165f908152600860205260409020805460ff1916911515919091179055565b6040516001600160a01b038316602482015281151560448201525f9060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f9060bda9000000000000000000000000000000000000000000000000000000001790529050610f00816137d7565b5f8581526006602052604090205460ff1680156133295750805b15613363576040517fa1452cb000000000000000000000000000000000000000000000000000000000815260048101869052602401612377565b5f8581526006602052604090205460ff1615801561337f575080155b156133b9576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101869052602401612377565b5f858152600660205260409020600201546001600160a01b03161580156133e05750468514155b1561341a57600380546001810182555f919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018590555b5f858152600660205260409020805460ff19168215151781556002810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038716179055600301613476838583614b81565b5080156134b657600280546001810182555f919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace018590556124fa565b6124fa85613926565b5f85858585856040516024016134d9959493929190614d32565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa8f2cb9600000000000000000000000000000000000000000000000000000000179052905061258c816137d7565b61232f6139cb565b61354c6139cb565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff19169055565b5f8781526006602052604090205460ff166135c1576040517f8e6feba500000000000000000000000000000000000000000000000000000000815260048101889052602401612377565b5f8590036135fe5785856040517ec10cfd000000000000000000000000000000000000000000000000000000008152600401612377929190614978565b5f8781526007602052604090819020905161361c90889088906147de565b9081526040519081900360200190205460ff1661366b578686866040517f2b4f9c8600000000000000000000000000000000000000000000000000000000815260040161237793929190614b3c565b818160075f8a81526020019081526020015f20888860405161368e9291906147de565b908152602001604051809103902060030186866040516136af9291906147de565b90815260200160405180910390209182611bb7929190614b81565b5f878787878787876040516024016136e89796959493929190614d6e565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff354b31f000000000000000000000000000000000000000000000000000000001790529050611bb7816137d7565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610f35576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401612377565b5f5b600254811015610f355746600282815481106137f7576137f76148cf565b905f5260205f2001541480613845575060065f6002838154811061381d5761381d6148cf565b905f5260205f20015481526020019081526020015f206003018054613841906147ed565b1590505b61386e5761386e6002828154811061385f5761385f6148cf565b905f5260205f20015483613a32565b6001016137d9565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1661232f576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6138da82613ce4565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561391e57610f008282613d8b565b610f35613dfd565b5f5b600254811015610f35578160028281548110613946576139466148cf565b905f5260205f200154036139c3576002805461396490600190614dbd565b81548110613974576139746148cf565b905f5260205f20015460028281548110613990576139906148cf565b5f9182526020909120015560028054806139ac576139ac614df5565b600190038181905f5260205f20015f905590555050565b600101613928565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661232f576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820182526207a12081525f6020808301829052835160a0810185528281529081018290529283018190526060808401526080830152905f848152600660205260408082206002015490517ffc5fecd50000000000000000000000000000000000000000000000000000000081526207a12060048201526001600160a01b039091169190829063fc5fecd5906024016040805180830381865afa158015613adf573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b039190614e22565b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018290529092506001600160a01b03841691506323b872dd906064016020604051808303815f875af1158015613b70573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b949190614e4e565b613bca576040517f90b8ec1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152602481018390529083169063095ea7b3906044016020604051808303815f875af1158015613c33573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613c579190614e4e565b50600b545f878152600660205260409081902090517f06cb89830000000000000000000000000000000000000000000000000000000081526001600160a01b03909216916306cb898391613cbb9160039091019086908a908a908a90600401614e69565b5f604051808303815f87803b158015613cd2575f80fd5b505af1158015613138573d5f803e3d5ffd5b806001600160a01b03163b5f03613d32576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401612377565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f80846001600160a01b031684604051613da79190614913565b5f60405180830381855af49150503d805f8114613ddf576040519150601f19603f3d011682016040523d82523d5f602084013e613de4565b606091505b5091509150613df4858383613e35565b95945050505050565b341561232f576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606082613e4a57613e4582613eaa565b611281565b8151158015613e6157506001600160a01b0384163b155b15613ea3576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401612377565b5080611281565b805115613eba5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60208284031215613efc575f80fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611281575f80fd5b5f8083601f840112613f3b575f80fd5b50813567ffffffffffffffff811115613f52575f80fd5b602083019150836020828501011115613f69575f80fd5b9250929050565b8015158114610f17575f80fd5b5f805f8060608587031215613f90575f80fd5b84359350602085013567ffffffffffffffff811115613fad575f80fd5b613fb987828801613f2b565b9094509250506040850135613fcd81613f70565b939692955090935050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156140b9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180511515865260208101516080602088015261407c6080880182613fd8565b9050604082015187820360408901526140958282613fd8565b6060938401519890930197909752509450602093840193919091019060010161402c565b50929695505050505050565b5f805f805f606086880312156140d9575f80fd5b85359450602086013567ffffffffffffffff8111156140f6575f80fd5b61410288828901613f2b565b909550935050604086013567ffffffffffffffff811115614121575f80fd5b61412d88828901613f2b565b969995985093965092949392505050565b5f6020828403121561414e575f80fd5b5035919050565b6001600160a01b0381168114610f17575f80fd5b5f806040838503121561417a575f80fd5b82359150602083013561418c81614155565b809150509250929050565b5f602082840312156141a7575f80fd5b813561128181614155565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f80604083850312156141f0575f80fd5b82356141fb81614155565b9150602083013567ffffffffffffffff811115614216575f80fd5b8301601f81018513614226575f80fd5b803567ffffffffffffffff811115614240576142406141b2565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff82111715614270576142706141b2565b604052818152828201602001871015614287575f80fd5b816020840160208301375f602083830101528093505050509250929050565b5f805f604084860312156142b8575f80fd5b83359250602084013567ffffffffffffffff8111156142d5575f80fd5b6142e186828701613f2b565b9497909650939450505050565b8215158152604060208201525f6143086040830184613fd8565b949350505050565b5f805f805f805f805f60c08a8c031215614328575f80fd5b893561433381614155565b985060208a013567ffffffffffffffff81111561434e575f80fd5b61435a8c828d01613f2b565b90995097505060408a0135955060608a013567ffffffffffffffff811115614380575f80fd5b61438c8c828d01613f2b565b90965094505060808a013567ffffffffffffffff8111156143ab575f80fd5b6143b78c828d01613f2b565b90945092505060a08a013560ff811681146143d0575f80fd5b809150509295985092959850929598565b602081525f6112816020830184613fd8565b6001600160a01b0383168152604060208201525f6143086040830184613fd8565b5f8060408385031215614425575f80fd5b823561443081614155565b9150602083013561418c81613f70565b602080825282518282018190525f918401906040840190835b81811015614477578351835260209384019390920191600101614459565b509095945050505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156140b9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805115158652602081015160208701526001600160a01b03604082015116604087015260608101519050608060608701526145176080870182613fd8565b95505060209384019391909101906001016144a8565b5f805f805f60808688031215614541575f80fd5b85359450602086013561455381614155565b9350604086013567ffffffffffffffff81111561456e575f80fd5b61457a88828901613f2b565b909450925050606086013561458e81613f70565b809150509295509295909350565b5f805f606084860312156145ae575f80fd5b83356145b981614155565b925060208401356145c981614155565b915060408401356145d981614155565b809150509250925092565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156140b9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051151586526001600160a01b036020820151166020870152604081015160e0604088015261466d60e0880182613fd8565b905060608201516060880152608082015187820360808901526146908282613fd8565b91505060a082015187820360a08901526146aa8282613fd8565b91505060c082015191506146c360c088018360ff169052565b955050602093840193919091019060010161460a565b861515815260c060208201525f6146f360c0830188613fd8565b866040840152828103606084015261470b8187613fd8565b9050828103608084015261471f8186613fd8565b91505060ff831660a0830152979650505050505050565b5f805f805f805f6080888a03121561474c575f80fd5b87359650602088013567ffffffffffffffff811115614769575f80fd5b6147758a828b01613f2b565b909750955050604088013567ffffffffffffffff811115614794575f80fd5b6147a08a828b01613f2b565b909550935050606088013567ffffffffffffffff8111156147bf575f80fd5b6147cb8a828b01613f2b565b989b979a50959850939692959293505050565b818382375f9101908152919050565b600181811c9082168061480157607f821691505b602082108103614838577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f815461484a816147ed565b8085526001821680156148645760018114614880576148b4565b60ff1983166020870152602082151560051b87010193506148b4565b845f5260205f205f5b838110156148ab5781546020828a010152600182019150602081019050614889565b87016020019450505b50505092915050565b602081525f611281602083018461483e565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81518060208401855e5f93019283525090919050565b5f61128182846148fc565b81835281816020850137505f602082840101525f6020601f19601f840116840101905092915050565b604081525f61495a60408301868861491e565b828103602084015261496d81858761491e565b979650505050505050565b602081525f61430860208301848661491e565b60ff85168152836020820152606060408201525f6149ad60608301848661491e565b9695505050505050565b601f821115610f0057805f5260205f20601f840160051c810160208510156149dc5750805b601f840160051c820191505b818110156124fa575f81556001016149e8565b815167ffffffffffffffff811115614a1557614a156141b2565b614a2981614a2384546147ed565b846149b7565b6020601f821160018114614a7a575f8315614a445750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556124fa565b5f84815260208120601f198516915b82811015614aa95787850151825560209485019460019092019101614a89565b5084821015614ae557868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b606081525f614b0760608301888a61491e565b8281036020840152614b1a81878961491e565b90508281036040840152614b2f81858761491e565b9998505050505050505050565b838152604060208201525f613df460408301848661491e565b848152606060208201525f614b6e60608301858761491e565b9050821515604083015295945050505050565b67ffffffffffffffff831115614b9957614b996141b2565b614bad83614ba783546147ed565b836149b7565b5f601f841160018114614bfd575f8515614bc75750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b1783556124fa565b5f83815260208120601f198716915b82811015614c2c5786850135825560209485019460019092019101614c0c565b5086821015614c67577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b858152606060208201525f614c9260608301868861491e565b8281036040840152614ca581858761491e565b98975050505050505050565b5f60208284031215614cc1575f80fd5b5051919050565b6001600160a01b038a16815260c060208201525f614cea60c083018a8c61491e565b8860408401528281036060840152614d0381888a61491e565b90508281036080840152614d1881868861491e565b91505060ff831660a08301529a9950505050505050505050565b8581526001600160a01b0385166020820152608060408201525f614d5a60808301858761491e565b905082151560608301529695505050505050565b878152608060208201525f614d8760808301888a61491e565b8281036040840152614d9a81878961491e565b90508281036060840152614daf81858761491e565b9a9950505050505050505050565b818103818111156108f7577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b5f8060408385031215614e33575f80fd5b8251614e3e81614155565b6020939093015192949293505050565b5f60208284031215614e5e575f80fd5b815161128181613f70565b60c081525f614e7b60c083018861483e565b6001600160a01b03871660208401528281036040840152614e9c8187613fd8565b90508451606084015260208501511515608084015282810360a08401526001600160a01b0384511681526020840151151560208201526001600160a01b036040850151166040820152606084015160a06060830152614efe60a0830182613fd8565b9050608085015160808301528092505050969550505050505056fea2646970667358221220c8a2b276a66832369b6ff2442dc99da97c81d9f293da94a4ec18c9c59e3da29064736f6c634300081a003360a060405234801561000f575f80fd5b50737cce3eb018bf23e1fe2a32692f2c77592d1103946001600160a01b031663cc5ad8b66040518163ffffffff1660e01b81526004016020604051808303815f875af1158015610061573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100859190610096565b6001600160a01b03166080526100c3565b5f602082840312156100a6575f80fd5b81516001600160a01b03811681146100bc575f80fd5b9392505050565b60805161090e6100da5f395f607c015261090e5ff3fe608060405260043610610062575f3560e01c80637b1039991161003f5780637b10399914610105578063c9028a361461012c578063ebf9b2aa1461013f57005b8063116191b61461006b5780632d4cfb7e146100c75780635bcfd616146100e657005b3661006957005b005b348015610076575f80fd5b5061009e7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100d2575f80fd5b506100696100e13660046102af565b610152565b3480156100f1575f80fd5b50610069610100366004610370565b61018c565b348015610110575f80fd5b5061009e737cce3eb018bf23e1fe2a32692f2c77592d11039481565b61006961013a3660046103f4565b610207565b61006961014d36600461042b565b610236565b7f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db781604051610181919061053b565b60405180910390a150565b606081156101a3576101a08284018461063b565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e6101ce878061072b565b6101de60408a0160208b0161078c565b896040013533866040516101f7969594939291906107a5565b60405180910390a1505050505050565b7fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c481604051610181919061084f565b6060811561024d5761024a8284018461063b565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e610278858061072b565b610288604088016020890161078c565b876040013533866040516102a1969594939291906107a5565b60405180910390a150505050565b5f602082840312156102bf575f80fd5b813567ffffffffffffffff8111156102d5575f80fd5b820160c081850312156102e6575f80fd5b9392505050565b5f606082840312156102fd575f80fd5b50919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610326575f80fd5b919050565b5f8083601f84011261033b575f80fd5b50813567ffffffffffffffff811115610352575f80fd5b602083019150836020828501011115610369575f80fd5b9250929050565b5f805f805f60808688031215610384575f80fd5b853567ffffffffffffffff81111561039a575f80fd5b6103a6888289016102ed565b9550506103b560208701610303565b935060408601359250606086013567ffffffffffffffff8111156103d7575f80fd5b6103e38882890161032b565b969995985093965092949392505050565b5f60208284031215610404575f80fd5b813567ffffffffffffffff81111561041a575f80fd5b8201608081850312156102e6575f80fd5b5f805f6040848603121561043d575f80fd5b833567ffffffffffffffff811115610453575f80fd5b61045f868287016102ed565b935050602084013567ffffffffffffffff81111561047b575f80fd5b6104878682870161032b565b9497909650939450505050565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126104c7575f80fd5b830160208101925035905067ffffffffffffffff8111156104e6575f80fd5b803603821315610369575f80fd5b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b602081525f61054a8384610494565b60c0602085015261055f60e0850182846104f4565b91505073ffffffffffffffffffffffffffffffffffffffff61058360208601610303565b1660408401525f604085013590508060608501525060608401358015158082146105ab575f80fd5b80608086015250505f608085013590508060a0850152506105cf60a0850185610494565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c08601526106048382846104f4565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f6020828403121561064b575f80fd5b813567ffffffffffffffff811115610661575f80fd5b8201601f81018413610671575f80fd5b803567ffffffffffffffff81111561068b5761068b61060e565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156106f7576106f761060e565b60405281815282820160200186101561070e575f80fd5b816020840160208301375f91810160200191909152949350505050565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261075e575f80fd5b83018035915067ffffffffffffffff821115610778575f80fd5b602001915036819003821315610369575f80fd5b5f6020828403121561079c575f80fd5b6102e682610303565b60a081525f6107b860a08301888a6104f4565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff85166060840152828103608084015283518082528060208601602084015e5f6020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff61087183610303565b16602082015273ffffffffffffffffffffffffffffffffffffffff61089860208401610303565b1660408201525f8060408401359050806060840152506108bb6060840184610494565b6080808501526108cf60a0850182846104f4565b9594505050505056fea264697066735822122073b2b7273c8fa2cc0343725b06eef9af8f807b8eda5d1626b08ac364147ff7d464736f6c634300081a003360c060405234801561000f575f80fd5b5060405161102e38038061102e83398101604081905261002e916100d8565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006257604051632b2add3d60e01b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0385811691909117909155828116608052811660a0526040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac5905f90a1505050610118565b80516001600160a01b03811681146100d3575f80fd5b919050565b5f805f606084860312156100ea575f80fd5b6100f3846100bd565b9250610101602085016100bd565b915061010f604085016100bd565b90509250925092565b60805160a051610eee6101405f395f6101dd01525f81816102b001526104510152610eee5ff3fe608060405234801561000f575f80fd5b50600436106100f0575f3560e01c806397770dff11610093578063c63585cc11610063578063c63585cc1461026b578063d7fd7afb1461027e578063d936a012146102ab578063ee2815ba146102d2575f80fd5b806397770dff14610212578063a7cb050714610225578063c39aca3714610238578063c62178ac1461024b575f80fd5b8063513a9c05116100ce578063513a9c0514610183578063569541b9146101b8578063842da36d146101d857806391dd645f146101ff575f80fd5b80630be15547146100f45780631f0e251b146101535780633ce4a5bc14610168575b5f80fd5b610129610102366004610bb9565b60016020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b610166610161366004610bf8565b6102e5565b005b61012973735b14bb79463307aacbed86daf3322b1e6226ab81565b610129610191366004610bb9565b60026020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101299073ffffffffffffffffffffffffffffffffffffffff1681565b6101297f000000000000000000000000000000000000000000000000000000000000000081565b61016661020d366004610c18565b6103f9565b610166610220366004610bf8565b61051b565b610166610233366004610c42565b610628565b610166610246366004610c62565b6106c2565b6004546101299073ffffffffffffffffffffffffffffffffffffffff1681565b610129610279366004610d28565b6108b9565b61029d61028c366004610bb9565b5f6020819052908152604090205481565b60405190815260200161014a565b6101297f000000000000000000000000000000000000000000000000000000000000000081565b6101666102e0366004610c18565b6109ec565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610332576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811661037f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610446576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003545f9061048d907f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff16846108b9565b5f8481526002602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251878152918201529192507f0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e910160405180910390a1505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610568576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166105b5576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e906020016103ee565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610675576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461070f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831673735b14bb79463307aacbed86daf3322b1e6226ab148061075c575073ffffffffffffffffffffffffffffffffffffffff831630145b15610793576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690528616906347e7ef24906044016020604051808303815f875af1158015610805573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108299190610d68565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063de43156e906108849089908990899088908890600401610dce565b5f604051808303815f87803b15801561089b575f80fd5b505af11580156108ad573d5f803e3d5ffd5b50505050505050505050565b5f805f6108c68585610abc565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b16603482015291935091508690604801604051602081830303815290604052805190602001206040516020016109ac9291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a39576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d91016106b6565b5f808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610b23576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610b5d578284610b60565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610bb2576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b5f60208284031215610bc9575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610bf3575f80fd5b919050565b5f60208284031215610c08575f80fd5b610c1182610bd0565b9392505050565b5f8060408385031215610c29575f80fd5b82359150610c3960208401610bd0565b90509250929050565b5f8060408385031215610c53575f80fd5b50508035926020909101359150565b5f805f805f8060a08789031215610c77575f80fd5b863567ffffffffffffffff811115610c8d575f80fd5b87016060818a031215610c9e575f80fd5b9550610cac60208801610bd0565b945060408701359350610cc160608801610bd0565b9250608087013567ffffffffffffffff811115610cdc575f80fd5b8701601f81018913610cec575f80fd5b803567ffffffffffffffff811115610d02575f80fd5b896020828401011115610d13575f80fd5b60208201935080925050509295509295509295565b5f805f60608486031215610d3a575f80fd5b610d4384610bd0565b9250610d5160208501610bd0565b9150610d5f60408501610bd0565b90509250925092565b5f60208284031215610d78575f80fd5b81518015158114610c11575f80fd5b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b608081525f86357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018112610e04575f80fd5b870160208101903567ffffffffffffffff811115610e20575f80fd5b803603821315610e2e575f80fd5b60606080850152610e4360e085018284610d87565b91505073ffffffffffffffffffffffffffffffffffffffff610e6760208a01610bd0565b1660a0840152604088013560c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401528281036060840152610eac818587610d87565b9897505050505050505056fea2646970667358221220ab90e60a29637a57eaea52c9d20d7377af3635b9a7cb793d7eb87995722b46ac64736f6c634300081a003360c060405234801561000f575f80fd5b50604051611fc0380380611fc083398101604081905261002e916101d0565b6001600160a01b038216158061004b57506001600160a01b038116155b156100695760405163d92e233d60e01b815260040160405180910390fd5b60066100758982610315565b5060076100828882610315565b506008805460ff191660ff881617905560808590528360028111156100a9576100a96103cf565b60a08160028111156100bd576100bd6103cf565b9052506001929092555f80546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506103e39350505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261012d575f80fd5b81516001600160401b038111156101465761014661010a565b604051601f8201601f19908116603f011681016001600160401b03811182821017156101745761017461010a565b60405281815283820160200185101561018b575f80fd5b8160208501602083015e5f918101602001919091529392505050565b8051600381106101b5575f80fd5b919050565b80516001600160a01b03811681146101b5575f80fd5b5f805f805f805f80610100898b0312156101e8575f80fd5b88516001600160401b038111156101fd575f80fd5b6102098b828c0161011e565b60208b015190995090506001600160401b03811115610226575f80fd5b6102328b828c0161011e565b975050604089015160ff81168114610248575f80fd5b60608a0151909650945061025e60808a016101a7565b60a08a0151909450925061027460c08a016101ba565b915061028260e08a016101ba565b90509295985092959890939650565b600181811c908216806102a557607f821691505b6020821081036102c357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561031057805f5260205f20601f840160051c810160208510156102ee5750805b601f840160051c820191505b8181101561030d575f81556001016102fa565b50505b505050565b81516001600160401b0381111561032e5761032e61010a565b6103428161033c8454610291565b846102c9565b6020601f821160018114610374575f831561035d5750848201515b5f19600385901b1c1916600184901b17845561030d565b5f84815260208120601f198516915b828110156103a35787850151825560209485019460019092019101610383565b50848210156103c057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52602160045260245ffd5b60805160a051611ba06104205f395f61033901525f81816102e501528181610bbf01528181610cc201528181610ed90152610fdc0152611ba05ff3fe608060405234801561000f575f80fd5b50600436106101b0575f3560e01c806395d89b41116100f3578063ccc7759911610093578063eddeb1231161006e578063eddeb12314610455578063f2441b3214610468578063f687d12a14610487578063fc5fecd51461049a575f80fd5b8063ccc77599146103c9578063d9eeebed146103dc578063dd62ed3e14610410575f80fd5b8063b84c8246116100ce578063b84c82461461037b578063c47f002714610390578063c7012626146103a3578063c835d7cc146103b6575f80fd5b806395d89b411461032c578063a3413d0314610334578063a9059cbb14610368575f80fd5b80633ce4a5bc1161015e5780634d8943bb116101395780634d8943bb146102a257806370a08231146102ab57806385e1f4d0146102e05780638b851b9514610307575f80fd5b80633ce4a5bc1461023c57806342966c681461027c57806347e7ef241461028f575f80fd5b806318160ddd1161018e57806318160ddd1461020c57806323b872dd14610214578063313ce56714610227575f80fd5b806306fdde03146101b4578063091d2788146101d2578063095ea7b3146101e9575b5f80fd5b6101bc6104ad565b6040516101c991906115fb565b60405180910390f35b6101db60015481565b6040519081526020016101c9565b6101fc6101f7366004611638565b61053d565b60405190151581526020016101c9565b6005546101db565b6101fc610222366004611662565b610553565b60085460405160ff90911681526020016101c9565b61025773735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c9565b6101fc61028a3660046116a0565b6105e8565b6101fc61029d366004611638565b6105fb565b6101db60025481565b6101db6102b93660046116b7565b73ffffffffffffffffffffffffffffffffffffffff165f9081526003602052604090205490565b6101db7f000000000000000000000000000000000000000000000000000000000000000081565b60085461025790610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101bc610752565b61035b7f000000000000000000000000000000000000000000000000000000000000000081565b6040516101c991906116d2565b6101fc610376366004611638565b610761565b61038e6103893660046117d3565b61076d565b005b61038e61039e3660046117d3565b6107ca565b6101fc6103b1366004611820565b610823565b61038e6103c43660046116b7565b61096d565b61038e6103d73660046116b7565b610a80565b6103e4610b94565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101c9565b6101db61041e366004611875565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260046020908152604080832093909416825291909152205490565b61038e6104633660046116a0565b610daa565b5f546102579073ffffffffffffffffffffffffffffffffffffffff1681565b61038e6104953660046116a0565b610e2c565b6103e46104a83660046116a0565b610eae565b6060600680546104bc906118ac565b80601f01602080910402602001604051908101604052809291908181526020018280546104e8906118ac565b80156105335780601f1061050a57610100808354040283529160200191610533565b820191905f5260205f20905b81548152906001019060200180831161051657829003601f168201915b5050505050905090565b5f6105493384846110c2565b5060015b92915050565b5f61055f8484846111ca565b73ffffffffffffffffffffffffffffffffffffffff84165f908152600460209081526040808320338452909152902054828110156105c9576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105dd85336105d8868561192a565b6110c2565b506001949350505050565b5f6105f33383611383565b506001919050565b5f3373735b14bb79463307aacbed86daf3322b1e6226ab1480159061063757505f5473ffffffffffffffffffffffffffffffffffffffff163314155b80156106605750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b15610697576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106a183836114c2565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261074191869061193d565b60405180910390a250600192915050565b6060600780546104bc906118ac565b5f6105493384846111ca565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107ba576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076107c682826119aa565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610817576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60066107c682826119aa565b5f805f61082e610b94565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303815f875af11580156108bd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108e19190611ac1565b610917576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109213385611383565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161095a91899189918791611ae0565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109ba576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a07576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610acd576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b1a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610a75565b5f80546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610c24573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c489190611b0e565b905073ffffffffffffffffffffffffffffffffffffffff8116610c97576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610d23573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d479190611b29565b9050805f03610d82576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60025460015483610d949190611b40565b610d9e9190611b57565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610df7576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610a75565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e79576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610a75565b5f80546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610f3e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f629190611b0e565b905073ffffffffffffffffffffffffffffffffffffffff8116610fb1576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa15801561103d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110619190611b29565b9050805f0361109c576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002545f906110ab8784611b40565b6110b59190611b57565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff831661110f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821661115c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8381165f8181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611217576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611264576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f90815260036020526040902054818110156112c3576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112cd828261192a565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260036020526040808220939093559085168152908120805484929061130f908490611b57565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161137591815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166113d0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f908152600360205260409020548181101561142f576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611439828261192a565b73ffffffffffffffffffffffffffffffffffffffff84165f908152600360205260408120919091556005805484929061147390849061192a565b90915550506040518281525f9073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016111bd565b73ffffffffffffffffffffffffffffffffffffffff821661150f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060055f8282546115209190611b57565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f9081526003602052604081208054839290611559908490611b57565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316905f907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61160d60208301846115af565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff81168114611635575f80fd5b50565b5f8060408385031215611649575f80fd5b823561165481611614565b946020939093013593505050565b5f805f60608486031215611674575f80fd5b833561167f81611614565b9250602084013561168f81611614565b929592945050506040919091013590565b5f602082840312156116b0575f80fd5b5035919050565b5f602082840312156116c7575f80fd5b813561160d81611614565b602081016003831061170b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b91905290565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8067ffffffffffffffff84111561175857611758611711565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156117a5576117a5611711565b6040528381529050808284018510156117bc575f80fd5b838360208301375f60208583010152509392505050565b5f602082840312156117e3575f80fd5b813567ffffffffffffffff8111156117f9575f80fd5b8201601f81018413611809575f80fd5b6118188482356020840161173e565b949350505050565b5f8060408385031215611831575f80fd5b823567ffffffffffffffff811115611847575f80fd5b8301601f81018513611857575f80fd5b6118668582356020840161173e565b95602094909401359450505050565b5f8060408385031215611886575f80fd5b823561189181611614565b915060208301356118a181611614565b809150509250929050565b600181811c908216806118c057607f821691505b6020821081036118f7577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561054d5761054d6118fd565b604081525f61194f60408301856115af565b90508260208301529392505050565b601f8211156119a557805f5260205f20601f840160051c810160208510156119835750805b601f840160051c820191505b818110156119a2575f815560010161198f565b50505b505050565b815167ffffffffffffffff8111156119c4576119c4611711565b6119d8816119d284546118ac565b8461195e565b6020601f821160018114611a29575f83156119f35750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556119a2565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611a765787850151825560209485019460019092019101611a56565b5084821015611ab257868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b5f60208284031215611ad1575f80fd5b8151801515811461160d575f80fd5b608081525f611af260808301876115af565b6020830195909552506040810192909252606090910152919050565b5f60208284031215611b1e575f80fd5b815161160d81611614565b5f60208284031215611b39575f80fd5b5051919050565b808202811582820484141761054d5761054d6118fd565b8082018082111561054d5761054d6118fd56fea26469706673582212206130fb621d6dd12e8164e45fefe5a69512cd6f3f9ac39ab60f6f9d44bc45140c64736f6c634300081a0033a264697066735822122092f59a654fdfc11502d8e0dfa9332d3639875baa5405ff0432bb20e2c3a4cedb64736f6c634300081a0033",
}

// GatewayZEVMInboundTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayZEVMInboundTestMetaData.ABI instead.
var GatewayZEVMInboundTestABI = GatewayZEVMInboundTestMetaData.ABI

// GatewayZEVMInboundTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayZEVMInboundTestMetaData.Bin instead.
var GatewayZEVMInboundTestBin = GatewayZEVMInboundTestMetaData.Bin

// DeployGatewayZEVMInboundTest deploys a new Ethereum contract, binding an instance of GatewayZEVMInboundTest to it.
func DeployGatewayZEVMInboundTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayZEVMInboundTest, error) {
	parsed, err := GatewayZEVMInboundTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayZEVMInboundTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayZEVMInboundTest{GatewayZEVMInboundTestCaller: GatewayZEVMInboundTestCaller{contract: contract}, GatewayZEVMInboundTestTransactor: GatewayZEVMInboundTestTransactor{contract: contract}, GatewayZEVMInboundTestFilterer: GatewayZEVMInboundTestFilterer{contract: contract}}, nil
}

// GatewayZEVMInboundTest is an auto generated Go binding around an Ethereum contract.
type GatewayZEVMInboundTest struct {
	GatewayZEVMInboundTestCaller     // Read-only binding to the contract
	GatewayZEVMInboundTestTransactor // Write-only binding to the contract
	GatewayZEVMInboundTestFilterer   // Log filterer for contract events
}

// GatewayZEVMInboundTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayZEVMInboundTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMInboundTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayZEVMInboundTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMInboundTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayZEVMInboundTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMInboundTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayZEVMInboundTestSession struct {
	Contract     *GatewayZEVMInboundTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GatewayZEVMInboundTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayZEVMInboundTestCallerSession struct {
	Contract *GatewayZEVMInboundTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// GatewayZEVMInboundTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayZEVMInboundTestTransactorSession struct {
	Contract     *GatewayZEVMInboundTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// GatewayZEVMInboundTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayZEVMInboundTestRaw struct {
	Contract *GatewayZEVMInboundTest // Generic contract binding to access the raw methods on
}

// GatewayZEVMInboundTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayZEVMInboundTestCallerRaw struct {
	Contract *GatewayZEVMInboundTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayZEVMInboundTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayZEVMInboundTestTransactorRaw struct {
	Contract *GatewayZEVMInboundTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayZEVMInboundTest creates a new instance of GatewayZEVMInboundTest, bound to a specific deployed contract.
func NewGatewayZEVMInboundTest(address common.Address, backend bind.ContractBackend) (*GatewayZEVMInboundTest, error) {
	contract, err := bindGatewayZEVMInboundTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTest{GatewayZEVMInboundTestCaller: GatewayZEVMInboundTestCaller{contract: contract}, GatewayZEVMInboundTestTransactor: GatewayZEVMInboundTestTransactor{contract: contract}, GatewayZEVMInboundTestFilterer: GatewayZEVMInboundTestFilterer{contract: contract}}, nil
}

// NewGatewayZEVMInboundTestCaller creates a new read-only instance of GatewayZEVMInboundTest, bound to a specific deployed contract.
func NewGatewayZEVMInboundTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayZEVMInboundTestCaller, error) {
	contract, err := bindGatewayZEVMInboundTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestCaller{contract: contract}, nil
}

// NewGatewayZEVMInboundTestTransactor creates a new write-only instance of GatewayZEVMInboundTest, bound to a specific deployed contract.
func NewGatewayZEVMInboundTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayZEVMInboundTestTransactor, error) {
	contract, err := bindGatewayZEVMInboundTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestTransactor{contract: contract}, nil
}

// NewGatewayZEVMInboundTestFilterer creates a new log filterer instance of GatewayZEVMInboundTest, bound to a specific deployed contract.
func NewGatewayZEVMInboundTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayZEVMInboundTestFilterer, error) {
	contract, err := bindGatewayZEVMInboundTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestFilterer{contract: contract}, nil
}

// bindGatewayZEVMInboundTest binds a generic wrapper to an already deployed contract.
func bindGatewayZEVMInboundTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayZEVMInboundTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMInboundTest.Contract.GatewayZEVMInboundTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.GatewayZEVMInboundTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.GatewayZEVMInboundTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMInboundTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.contract.Transact(opts, method, params...)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) ISTEST() (bool, error) {
	return _GatewayZEVMInboundTest.Contract.ISTEST(&_GatewayZEVMInboundTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) ISTEST() (bool, error) {
	return _GatewayZEVMInboundTest.Contract.ISTEST(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeArtifacts(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeArtifacts(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeContracts(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeContracts(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeSelectors(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeSelectors(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeSenders(&_GatewayZEVMInboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.ExcludeSenders(&_GatewayZEVMInboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) Failed() (bool, error) {
	return _GatewayZEVMInboundTest.Contract.Failed(&_GatewayZEVMInboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) Failed() (bool, error) {
	return _GatewayZEVMInboundTest.Contract.Failed(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayZEVMInboundTest.Contract.TargetArtifactSelectors(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayZEVMInboundTest.Contract.TargetArtifactSelectors(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayZEVMInboundTest.Contract.TargetArtifacts(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayZEVMInboundTest.Contract.TargetArtifacts(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.TargetContracts(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.TargetContracts(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayZEVMInboundTest.Contract.TargetInterfaces(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayZEVMInboundTest.Contract.TargetInterfaces(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMInboundTest.Contract.TargetSelectors(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayZEVMInboundTest.Contract.TargetSelectors(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayZEVMInboundTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.TargetSenders(&_GatewayZEVMInboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayZEVMInboundTest.Contract.TargetSenders(&_GatewayZEVMInboundTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.SetUp(&_GatewayZEVMInboundTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.SetUp(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestCallFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testCallFailsIfRevertGasLimitExceeded")
}

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestCallFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallFailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestCallFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallFailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOpts is a paid mutator transaction binding the contract method 0x6198fb19.
//
// Solidity: function testCallWithCallOpts() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestCallWithCallOpts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testCallWithCallOpts")
}

// TestCallWithCallOpts is a paid mutator transaction binding the contract method 0x6198fb19.
//
// Solidity: function testCallWithCallOpts() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestCallWithCallOpts() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOpts(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOpts is a paid mutator transaction binding the contract method 0x6198fb19.
//
// Solidity: function testCallWithCallOpts() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestCallWithCallOpts() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOpts(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0x05b9f046.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsBelowMin() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestCallWithCallOptsFailsIfGasLimitIsBelowMin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testCallWithCallOptsFailsIfGasLimitIsBelowMin")
}

// TestCallWithCallOptsFailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0x05b9f046.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsBelowMin() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestCallWithCallOptsFailsIfGasLimitIsBelowMin() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfGasLimitIsBelowMin(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0x05b9f046.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsBelowMin() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestCallWithCallOptsFailsIfGasLimitIsBelowMin() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfGasLimitIsBelowMin(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x5efe72a9.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestCallWithCallOptsFailsIfGasLimitIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testCallWithCallOptsFailsIfGasLimitIsZero")
}

// TestCallWithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x5efe72a9.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestCallWithCallOptsFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x5efe72a9.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestCallWithCallOptsFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfGasLimitTooHigh is a paid mutator transaction binding the contract method 0x6f2154d5.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitTooHigh() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestCallWithCallOptsFailsIfGasLimitTooHigh(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testCallWithCallOptsFailsIfGasLimitTooHigh")
}

// TestCallWithCallOptsFailsIfGasLimitTooHigh is a paid mutator transaction binding the contract method 0x6f2154d5.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitTooHigh() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestCallWithCallOptsFailsIfGasLimitTooHigh() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfGasLimitTooHigh(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfGasLimitTooHigh is a paid mutator transaction binding the contract method 0x6f2154d5.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitTooHigh() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestCallWithCallOptsFailsIfGasLimitTooHigh() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfGasLimitTooHigh(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x6dfcbc50.
//
// Solidity: function testCallWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestCallWithCallOptsFailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testCallWithCallOptsFailsIfMessageSizeExceeded")
}

// TestCallWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x6dfcbc50.
//
// Solidity: function testCallWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestCallWithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x6dfcbc50.
//
// Solidity: function testCallWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestCallWithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xb51ac071.
//
// Solidity: function testCallWithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestCallWithCallOptsFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testCallWithCallOptsFailsIfReceiverIsZeroAddress")
}

// TestCallWithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xb51ac071.
//
// Solidity: function testCallWithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestCallWithCallOptsFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xb51ac071.
//
// Solidity: function testCallWithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestCallWithCallOptsFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestCallWithCallOptsFailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestUpgradeAndWithdrawZRC20 is a paid mutator transaction binding the contract method 0x20dee15f.
//
// Solidity: function testUpgradeAndWithdrawZRC20() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestUpgradeAndWithdrawZRC20(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testUpgradeAndWithdrawZRC20")
}

// TestUpgradeAndWithdrawZRC20 is a paid mutator transaction binding the contract method 0x20dee15f.
//
// Solidity: function testUpgradeAndWithdrawZRC20() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestUpgradeAndWithdrawZRC20() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestUpgradeAndWithdrawZRC20(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestUpgradeAndWithdrawZRC20 is a paid mutator transaction binding the contract method 0x20dee15f.
//
// Solidity: function testUpgradeAndWithdrawZRC20() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestUpgradeAndWithdrawZRC20() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestUpgradeAndWithdrawZRC20(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x91dc32e7.
//
// Solidity: function testWithdrawAndCallZETAFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZETAFailsIfRevertGasLimitExceeded")
}

// TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x91dc32e7.
//
// Solidity: function testWithdrawAndCallZETAFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x91dc32e7.
//
// Solidity: function testWithdrawAndCallZETAFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAFailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x43181437.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress")
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x43181437.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x43181437.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x42752d41.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero")
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x42752d41.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x42752d41.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xc946d7c0.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded")
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xc946d7c0.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xc946d7c0.
//
// Solidity: function testWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xba9adeef.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20FailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20FailsIfAmountIsZero")
}

// TestWithdrawAndCallZRC20FailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xba9adeef.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20FailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xba9adeef.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20FailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0xc20049f4.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin")
}

// TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0xc20049f4.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0xc20049f4.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0xbed3e813.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20FailsIfGasLimitIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20FailsIfGasLimitIsZero")
}

// TestWithdrawAndCallZRC20FailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0xbed3e813.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20FailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0xbed3e813.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20FailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfGasLimitTooHigh is a paid mutator transaction binding the contract method 0x86bb7ac1.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitTooHigh() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20FailsIfGasLimitTooHigh(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20FailsIfGasLimitTooHigh")
}

// TestWithdrawAndCallZRC20FailsIfGasLimitTooHigh is a paid mutator transaction binding the contract method 0x86bb7ac1.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitTooHigh() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20FailsIfGasLimitTooHigh() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfGasLimitTooHigh(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfGasLimitTooHigh is a paid mutator transaction binding the contract method 0x86bb7ac1.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfGasLimitTooHigh() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20FailsIfGasLimitTooHigh() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfGasLimitTooHigh(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x4ffab9de.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20FailsIfMessageSizeExceeded")
}

// TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x4ffab9de.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x4ffab9de.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x36431b3f.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress")
}

// TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x36431b3f.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x36431b3f.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xecf26b30.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded")
}

// TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xecf26b30.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xecf26b30.
//
// Solidity: function testWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20FailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x1b9641bf.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero")
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x1b9641bf.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x1b9641bf.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x0b5ad28d.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero")
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x0b5ad28d.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x0b5ad28d.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xa721b2d3.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded")
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xa721b2d3.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xa721b2d3.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xeb7a2fac.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress")
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xeb7a2fac.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xeb7a2fac.
//
// Solidity: function testWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETA is a paid mutator transaction binding the contract method 0xea37902f.
//
// Solidity: function testWithdrawZETA() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETA")
}

// TestWithdrawZETA is a paid mutator transaction binding the contract method 0xea37902f.
//
// Solidity: function testWithdrawZETA() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETA() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETA(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETA is a paid mutator transaction binding the contract method 0xea37902f.
//
// Solidity: function testWithdrawZETA() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETA() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETA(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x21aeb18c.
//
// Solidity: function testWithdrawZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAFailsIfAmountIsZero")
}

// TestWithdrawZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x21aeb18c.
//
// Solidity: function testWithdrawZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x21aeb18c.
//
// Solidity: function testWithdrawZETAFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x7ae69730.
//
// Solidity: function testWithdrawZETAFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAFailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAFailsIfMessageSizeExceeded")
}

// TestWithdrawZETAFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x7ae69730.
//
// Solidity: function testWithdrawZETAFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x7ae69730.
//
// Solidity: function testWithdrawZETAFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xdde7e967.
//
// Solidity: function testWithdrawZETAFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAFailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAFailsIfNoAllowance")
}

// TestWithdrawZETAFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xdde7e967.
//
// Solidity: function testWithdrawZETAFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xdde7e967.
//
// Solidity: function testWithdrawZETAFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfNoBalance is a paid mutator transaction binding the contract method 0x6221b509.
//
// Solidity: function testWithdrawZETAFailsIfNoBalance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAFailsIfNoBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAFailsIfNoBalance")
}

// TestWithdrawZETAFailsIfNoBalance is a paid mutator transaction binding the contract method 0x6221b509.
//
// Solidity: function testWithdrawZETAFailsIfNoBalance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAFailsIfNoBalance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfNoBalance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfNoBalance is a paid mutator transaction binding the contract method 0x6221b509.
//
// Solidity: function testWithdrawZETAFailsIfNoBalance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAFailsIfNoBalance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfNoBalance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xdc749dd7.
//
// Solidity: function testWithdrawZETAFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAFailsIfReceiverIsZeroAddress")
}

// TestWithdrawZETAFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xdc749dd7.
//
// Solidity: function testWithdrawZETAFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xdc749dd7.
//
// Solidity: function testWithdrawZETAFailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xd597a27a.
//
// Solidity: function testWithdrawZETAFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAFailsIfRevertGasLimitExceeded")
}

// TestWithdrawZETAFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xd597a27a.
//
// Solidity: function testWithdrawZETAFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xd597a27a.
//
// Solidity: function testWithdrawZETAFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAFailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithBothDefaultValuesFromRegistry is a paid mutator transaction binding the contract method 0x3828ce8c.
//
// Solidity: function testWithdrawZETAWithBothDefaultValuesFromRegistry() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAWithBothDefaultValuesFromRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAWithBothDefaultValuesFromRegistry")
}

// TestWithdrawZETAWithBothDefaultValuesFromRegistry is a paid mutator transaction binding the contract method 0x3828ce8c.
//
// Solidity: function testWithdrawZETAWithBothDefaultValuesFromRegistry() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAWithBothDefaultValuesFromRegistry() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithBothDefaultValuesFromRegistry(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithBothDefaultValuesFromRegistry is a paid mutator transaction binding the contract method 0x3828ce8c.
//
// Solidity: function testWithdrawZETAWithBothDefaultValuesFromRegistry() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAWithBothDefaultValuesFromRegistry() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithBothDefaultValuesFromRegistry(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithCallOptsFailsIfGasLimitTooHigh is a paid mutator transaction binding the contract method 0xdf0def41.
//
// Solidity: function testWithdrawZETAWithCallOptsFailsIfGasLimitTooHigh() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAWithCallOptsFailsIfGasLimitTooHigh(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAWithCallOptsFailsIfGasLimitTooHigh")
}

// TestWithdrawZETAWithCallOptsFailsIfGasLimitTooHigh is a paid mutator transaction binding the contract method 0xdf0def41.
//
// Solidity: function testWithdrawZETAWithCallOptsFailsIfGasLimitTooHigh() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAWithCallOptsFailsIfGasLimitTooHigh() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithCallOptsFailsIfGasLimitTooHigh(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithCallOptsFailsIfGasLimitTooHigh is a paid mutator transaction binding the contract method 0xdf0def41.
//
// Solidity: function testWithdrawZETAWithCallOptsFailsIfGasLimitTooHigh() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAWithCallOptsFailsIfGasLimitTooHigh() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithCallOptsFailsIfGasLimitTooHigh(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithCallOptsWithMessage is a paid mutator transaction binding the contract method 0x14759766.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAWithCallOptsWithMessage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAWithCallOptsWithMessage")
}

// TestWithdrawZETAWithCallOptsWithMessage is a paid mutator transaction binding the contract method 0x14759766.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAWithCallOptsWithMessage() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithCallOptsWithMessage(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithCallOptsWithMessage is a paid mutator transaction binding the contract method 0x14759766.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAWithCallOptsWithMessage() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithCallOptsWithMessage(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x64002a1f.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero")
}

// TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x64002a1f.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x64002a1f.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xe51c6388.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance")
}

// TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xe51c6388.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xe51c6388.
//
// Solidity: function testWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithDefaultGasLimit is a paid mutator transaction binding the contract method 0x16c09ef7.
//
// Solidity: function testWithdrawZETAWithDefaultGasLimit() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAWithDefaultGasLimit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAWithDefaultGasLimit")
}

// TestWithdrawZETAWithDefaultGasLimit is a paid mutator transaction binding the contract method 0x16c09ef7.
//
// Solidity: function testWithdrawZETAWithDefaultGasLimit() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAWithDefaultGasLimit() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithDefaultGasLimit(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithDefaultGasLimit is a paid mutator transaction binding the contract method 0x16c09ef7.
//
// Solidity: function testWithdrawZETAWithDefaultGasLimit() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAWithDefaultGasLimit() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithDefaultGasLimit(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithDefaultProtocolFlatFee is a paid mutator transaction binding the contract method 0x291f1d82.
//
// Solidity: function testWithdrawZETAWithDefaultProtocolFlatFee() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZETAWithDefaultProtocolFlatFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZETAWithDefaultProtocolFlatFee")
}

// TestWithdrawZETAWithDefaultProtocolFlatFee is a paid mutator transaction binding the contract method 0x291f1d82.
//
// Solidity: function testWithdrawZETAWithDefaultProtocolFlatFee() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZETAWithDefaultProtocolFlatFee() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithDefaultProtocolFlatFee(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZETAWithDefaultProtocolFlatFee is a paid mutator transaction binding the contract method 0x291f1d82.
//
// Solidity: function testWithdrawZETAWithDefaultProtocolFlatFee() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZETAWithDefaultProtocolFlatFee() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZETAWithDefaultProtocolFlatFee(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20 is a paid mutator transaction binding the contract method 0xfbc611c8.
//
// Solidity: function testWithdrawZRC20() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20")
}

// TestWithdrawZRC20 is a paid mutator transaction binding the contract method 0xfbc611c8.
//
// Solidity: function testWithdrawZRC20() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20 is a paid mutator transaction binding the contract method 0xfbc611c8.
//
// Solidity: function testWithdrawZRC20() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x264b524c.
//
// Solidity: function testWithdrawZRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20FailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20FailsIfMessageSizeExceeded")
}

// TestWithdrawZRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x264b524c.
//
// Solidity: function testWithdrawZRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20FailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x264b524c.
//
// Solidity: function testWithdrawZRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20FailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfNoAllowance is a paid mutator transaction binding the contract method 0x5d72228f.
//
// Solidity: function testWithdrawZRC20FailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20FailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20FailsIfNoAllowance")
}

// TestWithdrawZRC20FailsIfNoAllowance is a paid mutator transaction binding the contract method 0x5d72228f.
//
// Solidity: function testWithdrawZRC20FailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20FailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfNoAllowance is a paid mutator transaction binding the contract method 0x5d72228f.
//
// Solidity: function testWithdrawZRC20FailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20FailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfNoBalanceForGasFee is a paid mutator transaction binding the contract method 0xceccfab3.
//
// Solidity: function testWithdrawZRC20FailsIfNoBalanceForGasFee() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20FailsIfNoBalanceForGasFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20FailsIfNoBalanceForGasFee")
}

// TestWithdrawZRC20FailsIfNoBalanceForGasFee is a paid mutator transaction binding the contract method 0xceccfab3.
//
// Solidity: function testWithdrawZRC20FailsIfNoBalanceForGasFee() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20FailsIfNoBalanceForGasFee() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfNoBalanceForGasFee(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfNoBalanceForGasFee is a paid mutator transaction binding the contract method 0xceccfab3.
//
// Solidity: function testWithdrawZRC20FailsIfNoBalanceForGasFee() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20FailsIfNoBalanceForGasFee() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfNoBalanceForGasFee(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfNoBalanceForTransfer is a paid mutator transaction binding the contract method 0xfdad0ad0.
//
// Solidity: function testWithdrawZRC20FailsIfNoBalanceForTransfer() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20FailsIfNoBalanceForTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20FailsIfNoBalanceForTransfer")
}

// TestWithdrawZRC20FailsIfNoBalanceForTransfer is a paid mutator transaction binding the contract method 0xfdad0ad0.
//
// Solidity: function testWithdrawZRC20FailsIfNoBalanceForTransfer() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20FailsIfNoBalanceForTransfer() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfNoBalanceForTransfer(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfNoBalanceForTransfer is a paid mutator transaction binding the contract method 0xfdad0ad0.
//
// Solidity: function testWithdrawZRC20FailsIfNoBalanceForTransfer() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20FailsIfNoBalanceForTransfer() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfNoBalanceForTransfer(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x83ababa9.
//
// Solidity: function testWithdrawZRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20FailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20FailsIfReceiverIsZeroAddress")
}

// TestWithdrawZRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x83ababa9.
//
// Solidity: function testWithdrawZRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20FailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x83ababa9.
//
// Solidity: function testWithdrawZRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20FailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfReceiverIsZeroAddress(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc3bb7957.
//
// Solidity: function testWithdrawZRC20FailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20FailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20FailsIfRevertGasLimitExceeded")
}

// TestWithdrawZRC20FailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc3bb7957.
//
// Solidity: function testWithdrawZRC20FailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20FailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc3bb7957.
//
// Solidity: function testWithdrawZRC20FailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20FailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIsAmountIs0 is a paid mutator transaction binding the contract method 0xba800c91.
//
// Solidity: function testWithdrawZRC20FailsIsAmountIs0() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20FailsIsAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20FailsIsAmountIs0")
}

// TestWithdrawZRC20FailsIsAmountIs0 is a paid mutator transaction binding the contract method 0xba800c91.
//
// Solidity: function testWithdrawZRC20FailsIsAmountIs0() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20FailsIsAmountIs0() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIsAmountIs0(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20FailsIsAmountIs0 is a paid mutator transaction binding the contract method 0xba800c91.
//
// Solidity: function testWithdrawZRC20FailsIsAmountIs0() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20FailsIsAmountIs0() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20FailsIsAmountIs0(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCallOptsWithMessage is a paid mutator transaction binding the contract method 0xe804a406.
//
// Solidity: function testWithdrawZRC20WithCallOptsWithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithCallOptsWithMessage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithCallOptsWithMessage")
}

// TestWithdrawZRC20WithCallOptsWithMessage is a paid mutator transaction binding the contract method 0xe804a406.
//
// Solidity: function testWithdrawZRC20WithCallOptsWithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithCallOptsWithMessage() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCallOptsWithMessage(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCallOptsWithMessage is a paid mutator transaction binding the contract method 0xe804a406.
//
// Solidity: function testWithdrawZRC20WithCallOptsWithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithCallOptsWithMessage() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCallOptsWithMessage(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCustomGasLimit is a paid mutator transaction binding the contract method 0x6abd223e.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimit() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithCustomGasLimit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithCustomGasLimit")
}

// TestWithdrawZRC20WithCustomGasLimit is a paid mutator transaction binding the contract method 0x6abd223e.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimit() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithCustomGasLimit() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCustomGasLimit(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCustomGasLimit is a paid mutator transaction binding the contract method 0x6abd223e.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimit() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithCustomGasLimit() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCustomGasLimit(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x7ba9b7ad.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero")
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x7ba9b7ad.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x7ba9b7ad.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooHigh is a paid mutator transaction binding the contract method 0x22442d76.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooHigh() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooHigh(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooHigh")
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooHigh is a paid mutator transaction binding the contract method 0x22442d76.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooHigh() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooHigh() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooHigh(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooHigh is a paid mutator transaction binding the contract method 0x22442d76.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooHigh() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooHigh() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooHigh(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow is a paid mutator transaction binding the contract method 0xb152ca46.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow")
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow is a paid mutator transaction binding the contract method 0xb152ca46.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow is a paid mutator transaction binding the contract method 0xb152ca46.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x6f5e2756.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded")
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x6f5e2756.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x6f5e2756.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty is a paid mutator transaction binding the contract method 0x564a7435.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty")
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty is a paid mutator transaction binding the contract method 0x564a7435.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty is a paid mutator transaction binding the contract method 0x564a7435.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xcb817356.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithCustomGasLimitFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithCustomGasLimitFailsIfRevertGasLimitExceeded")
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xcb817356.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithCustomGasLimitFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCustomGasLimitFailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithCustomGasLimitFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xcb817356.
//
// Solidity: function testWithdrawZRC20WithCustomGasLimitFailsIfRevertGasLimitExceeded() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithCustomGasLimitFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithCustomGasLimitFailsIfRevertGasLimitExceeded(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithMessage is a paid mutator transaction binding the contract method 0x1e63d2b9.
//
// Solidity: function testWithdrawZRC20WithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithMessage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithMessage")
}

// TestWithdrawZRC20WithMessage is a paid mutator transaction binding the contract method 0x1e63d2b9.
//
// Solidity: function testWithdrawZRC20WithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithMessage() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithMessage(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithMessage is a paid mutator transaction binding the contract method 0x1e63d2b9.
//
// Solidity: function testWithdrawZRC20WithMessage() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithMessage() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithMessage(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x1238212c.
//
// Solidity: function testWithdrawZRC20WithMessageFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithMessageFailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithMessageFailsIfNoAllowance")
}

// TestWithdrawZRC20WithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x1238212c.
//
// Solidity: function testWithdrawZRC20WithMessageFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithMessageFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithMessageFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x1238212c.
//
// Solidity: function testWithdrawZRC20WithMessageFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithMessageFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithMessageFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x6d6ce0d0.
//
// Solidity: function testWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance")
}

// TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x6d6ce0d0.
//
// Solidity: function testWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x6d6ce0d0.
//
// Solidity: function testWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance(&_GatewayZEVMInboundTest.TransactOpts)
}

// GatewayZEVMInboundTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestCalledIterator struct {
	Event *GatewayZEVMInboundTestCalled // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestCalled)
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
		it.Event = new(GatewayZEVMInboundTestCalled)
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
func (it *GatewayZEVMInboundTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestCalled represents a Called event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestCalled struct {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, zrc20 []common.Address) (*GatewayZEVMInboundTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestCalledIterator{contract: _GatewayZEVMInboundTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestCalled, sender []common.Address, zrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "Called", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestCalled)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "Called", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseCalled(log types.Log) (*GatewayZEVMInboundTestCalled, error) {
	event := new(GatewayZEVMInboundTestCalled)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestWithdrawnIterator struct {
	Event *GatewayZEVMInboundTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestWithdrawn)
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
		it.Event = new(GatewayZEVMInboundTestWithdrawn)
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
func (it *GatewayZEVMInboundTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestWithdrawn represents a Withdrawn event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestWithdrawn struct {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMInboundTestWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestWithdrawnIterator{contract: _GatewayZEVMInboundTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestWithdrawn)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseWithdrawn(log types.Log) (*GatewayZEVMInboundTestWithdrawn, error) {
	event := new(GatewayZEVMInboundTestWithdrawn)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestWithdrawnAndCalledIterator struct {
	Event *GatewayZEVMInboundTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestWithdrawnAndCalled)
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
		it.Event = new(GatewayZEVMInboundTestWithdrawnAndCalled)
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
func (it *GatewayZEVMInboundTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestWithdrawnAndCalled struct {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMInboundTestWithdrawnAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestWithdrawnAndCalledIterator{contract: _GatewayZEVMInboundTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestWithdrawnAndCalled, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestWithdrawnAndCalled)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayZEVMInboundTestWithdrawnAndCalled, error) {
	event := new(GatewayZEVMInboundTestWithdrawnAndCalled)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestWithdrawnV2Iterator is returned from FilterWithdrawnV2 and is used to iterate over the raw logs and unpacked data for WithdrawnV2 events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestWithdrawnV2Iterator struct {
	Event *GatewayZEVMInboundTestWithdrawnV2 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestWithdrawnV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestWithdrawnV2)
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
		it.Event = new(GatewayZEVMInboundTestWithdrawnV2)
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
func (it *GatewayZEVMInboundTestWithdrawnV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestWithdrawnV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestWithdrawnV2 represents a WithdrawnV2 event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestWithdrawnV2 struct {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterWithdrawnV2(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayZEVMInboundTestWithdrawnV2Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "WithdrawnV2", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestWithdrawnV2Iterator{contract: _GatewayZEVMInboundTest.contract, event: "WithdrawnV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnV2 is a free log subscription operation binding the contract event 0x5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d97.
//
// Solidity: event WithdrawnV2(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchWithdrawnV2(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestWithdrawnV2, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "WithdrawnV2", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestWithdrawnV2)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseWithdrawnV2(log types.Log) (*GatewayZEVMInboundTestWithdrawnV2, error) {
	event := new(GatewayZEVMInboundTestWithdrawnV2)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogIterator struct {
	Event *GatewayZEVMInboundTestLog // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLog)
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
		it.Event = new(GatewayZEVMInboundTestLog)
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
func (it *GatewayZEVMInboundTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLog represents a Log event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogIterator{contract: _GatewayZEVMInboundTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLog)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLog(log types.Log) (*GatewayZEVMInboundTestLog, error) {
	event := new(GatewayZEVMInboundTestLog)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogAddressIterator struct {
	Event *GatewayZEVMInboundTestLogAddress // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogAddress)
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
		it.Event = new(GatewayZEVMInboundTestLogAddress)
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
func (it *GatewayZEVMInboundTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogAddress represents a LogAddress event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogAddressIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogAddressIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogAddress)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogAddress(log types.Log) (*GatewayZEVMInboundTestLogAddress, error) {
	event := new(GatewayZEVMInboundTestLogAddress)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogArrayIterator struct {
	Event *GatewayZEVMInboundTestLogArray // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogArray)
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
		it.Event = new(GatewayZEVMInboundTestLogArray)
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
func (it *GatewayZEVMInboundTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogArray represents a LogArray event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogArrayIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogArrayIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogArray)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogArray(log types.Log) (*GatewayZEVMInboundTestLogArray, error) {
	event := new(GatewayZEVMInboundTestLogArray)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogArray0Iterator struct {
	Event *GatewayZEVMInboundTestLogArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogArray0)
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
		it.Event = new(GatewayZEVMInboundTestLogArray0)
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
func (it *GatewayZEVMInboundTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogArray0 represents a LogArray0 event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogArray0Iterator{contract: _GatewayZEVMInboundTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogArray0)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogArray0(log types.Log) (*GatewayZEVMInboundTestLogArray0, error) {
	event := new(GatewayZEVMInboundTestLogArray0)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogArray1Iterator struct {
	Event *GatewayZEVMInboundTestLogArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogArray1)
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
		it.Event = new(GatewayZEVMInboundTestLogArray1)
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
func (it *GatewayZEVMInboundTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogArray1 represents a LogArray1 event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogArray1Iterator{contract: _GatewayZEVMInboundTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogArray1)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogArray1(log types.Log) (*GatewayZEVMInboundTestLogArray1, error) {
	event := new(GatewayZEVMInboundTestLogArray1)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogBytesIterator struct {
	Event *GatewayZEVMInboundTestLogBytes // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogBytes)
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
		it.Event = new(GatewayZEVMInboundTestLogBytes)
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
func (it *GatewayZEVMInboundTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogBytes represents a LogBytes event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogBytesIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogBytesIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogBytes)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogBytes(log types.Log) (*GatewayZEVMInboundTestLogBytes, error) {
	event := new(GatewayZEVMInboundTestLogBytes)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogBytes32Iterator struct {
	Event *GatewayZEVMInboundTestLogBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogBytes32)
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
		it.Event = new(GatewayZEVMInboundTestLogBytes32)
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
func (it *GatewayZEVMInboundTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogBytes32 represents a LogBytes32 event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogBytes32Iterator{contract: _GatewayZEVMInboundTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogBytes32)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogBytes32(log types.Log) (*GatewayZEVMInboundTestLogBytes32, error) {
	event := new(GatewayZEVMInboundTestLogBytes32)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogIntIterator struct {
	Event *GatewayZEVMInboundTestLogInt // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogInt)
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
		it.Event = new(GatewayZEVMInboundTestLogInt)
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
func (it *GatewayZEVMInboundTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogInt represents a LogInt event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogIntIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogIntIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogInt)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogInt(log types.Log) (*GatewayZEVMInboundTestLogInt, error) {
	event := new(GatewayZEVMInboundTestLogInt)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedAddressIterator struct {
	Event *GatewayZEVMInboundTestLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedAddress)
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
		it.Event = new(GatewayZEVMInboundTestLogNamedAddress)
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
func (it *GatewayZEVMInboundTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedAddressIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedAddress)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayZEVMInboundTestLogNamedAddress, error) {
	event := new(GatewayZEVMInboundTestLogNamedAddress)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedArrayIterator struct {
	Event *GatewayZEVMInboundTestLogNamedArray // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedArray)
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
		it.Event = new(GatewayZEVMInboundTestLogNamedArray)
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
func (it *GatewayZEVMInboundTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedArray represents a LogNamedArray event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedArrayIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedArray)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayZEVMInboundTestLogNamedArray, error) {
	event := new(GatewayZEVMInboundTestLogNamedArray)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedArray0Iterator struct {
	Event *GatewayZEVMInboundTestLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedArray0)
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
		it.Event = new(GatewayZEVMInboundTestLogNamedArray0)
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
func (it *GatewayZEVMInboundTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedArray0Iterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedArray0)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayZEVMInboundTestLogNamedArray0, error) {
	event := new(GatewayZEVMInboundTestLogNamedArray0)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedArray1Iterator struct {
	Event *GatewayZEVMInboundTestLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedArray1)
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
		it.Event = new(GatewayZEVMInboundTestLogNamedArray1)
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
func (it *GatewayZEVMInboundTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedArray1Iterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedArray1)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayZEVMInboundTestLogNamedArray1, error) {
	event := new(GatewayZEVMInboundTestLogNamedArray1)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedBytesIterator struct {
	Event *GatewayZEVMInboundTestLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedBytes)
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
		it.Event = new(GatewayZEVMInboundTestLogNamedBytes)
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
func (it *GatewayZEVMInboundTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedBytesIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedBytes)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayZEVMInboundTestLogNamedBytes, error) {
	event := new(GatewayZEVMInboundTestLogNamedBytes)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedBytes32Iterator struct {
	Event *GatewayZEVMInboundTestLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedBytes32)
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
		it.Event = new(GatewayZEVMInboundTestLogNamedBytes32)
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
func (it *GatewayZEVMInboundTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedBytes32Iterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedBytes32)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayZEVMInboundTestLogNamedBytes32, error) {
	event := new(GatewayZEVMInboundTestLogNamedBytes32)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedDecimalIntIterator struct {
	Event *GatewayZEVMInboundTestLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedDecimalInt)
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
		it.Event = new(GatewayZEVMInboundTestLogNamedDecimalInt)
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
func (it *GatewayZEVMInboundTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedDecimalIntIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedDecimalInt)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayZEVMInboundTestLogNamedDecimalInt, error) {
	event := new(GatewayZEVMInboundTestLogNamedDecimalInt)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedDecimalUintIterator struct {
	Event *GatewayZEVMInboundTestLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedDecimalUint)
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
		it.Event = new(GatewayZEVMInboundTestLogNamedDecimalUint)
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
func (it *GatewayZEVMInboundTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedDecimalUintIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedDecimalUint)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayZEVMInboundTestLogNamedDecimalUint, error) {
	event := new(GatewayZEVMInboundTestLogNamedDecimalUint)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedIntIterator struct {
	Event *GatewayZEVMInboundTestLogNamedInt // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedInt)
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
		it.Event = new(GatewayZEVMInboundTestLogNamedInt)
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
func (it *GatewayZEVMInboundTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedInt represents a LogNamedInt event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedIntIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedInt)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayZEVMInboundTestLogNamedInt, error) {
	event := new(GatewayZEVMInboundTestLogNamedInt)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedStringIterator struct {
	Event *GatewayZEVMInboundTestLogNamedString // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedString)
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
		it.Event = new(GatewayZEVMInboundTestLogNamedString)
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
func (it *GatewayZEVMInboundTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedString represents a LogNamedString event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedStringIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedString)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedString(log types.Log) (*GatewayZEVMInboundTestLogNamedString, error) {
	event := new(GatewayZEVMInboundTestLogNamedString)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedUintIterator struct {
	Event *GatewayZEVMInboundTestLogNamedUint // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogNamedUint)
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
		it.Event = new(GatewayZEVMInboundTestLogNamedUint)
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
func (it *GatewayZEVMInboundTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogNamedUint represents a LogNamedUint event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogNamedUintIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogNamedUint)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayZEVMInboundTestLogNamedUint, error) {
	event := new(GatewayZEVMInboundTestLogNamedUint)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogStringIterator struct {
	Event *GatewayZEVMInboundTestLogString // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogString)
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
		it.Event = new(GatewayZEVMInboundTestLogString)
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
func (it *GatewayZEVMInboundTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogString represents a LogString event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogStringIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogStringIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogString)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogString(log types.Log) (*GatewayZEVMInboundTestLogString, error) {
	event := new(GatewayZEVMInboundTestLogString)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogUintIterator struct {
	Event *GatewayZEVMInboundTestLogUint // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogUint)
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
		it.Event = new(GatewayZEVMInboundTestLogUint)
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
func (it *GatewayZEVMInboundTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogUint represents a LogUint event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogUintIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogUintIterator{contract: _GatewayZEVMInboundTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogUint)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogUint(log types.Log) (*GatewayZEVMInboundTestLogUint, error) {
	event := new(GatewayZEVMInboundTestLogUint)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayZEVMInboundTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogsIterator struct {
	Event *GatewayZEVMInboundTestLogs // Event containing the contract specifics and raw log

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
func (it *GatewayZEVMInboundTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayZEVMInboundTestLogs)
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
		it.Event = new(GatewayZEVMInboundTestLogs)
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
func (it *GatewayZEVMInboundTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayZEVMInboundTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayZEVMInboundTestLogs represents a Logs event raised by the GatewayZEVMInboundTest contract.
type GatewayZEVMInboundTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayZEVMInboundTestLogsIterator, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMInboundTestLogsIterator{contract: _GatewayZEVMInboundTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayZEVMInboundTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayZEVMInboundTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayZEVMInboundTestLogs)
				if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "logs", log); err != nil {
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
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestFilterer) ParseLogs(log types.Log) (*GatewayZEVMInboundTestLogs, error) {
	event := new(GatewayZEVMInboundTestLogs)
	if err := _GatewayZEVMInboundTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
