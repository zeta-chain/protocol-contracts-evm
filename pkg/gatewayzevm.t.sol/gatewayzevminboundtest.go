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
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testCallWithCallOpts\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfGasLimitIsBelowMin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpgradeAndWithdrawZRC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20SolanaFails\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETA\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfNoBalance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithCallOptsWithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfNoBalanceForGasFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfNoBalanceForTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20FailsIsAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20Solana\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCallOptsWithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCustomGasLimit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithMessageFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LowBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZETANotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602b575f80fd5b50620119cb806200003b5f395ff3fe608060405234801561000f575f80fd5b506004361061037d575f3560e01c806373a3db40116101d4578063ba9adeef11610109578063e20c9f71116100a9578063eb7a2fac11610079578063eb7a2fac14610585578063fa7626d41461058d578063fbc611c81461059a578063fdad0ad0146105a2575f80fd5b8063e20c9f711461056d578063e51c638814610575578063e804a406146103b3578063ea37902f1461057d575f80fd5b8063c946d7c0116100e4578063c946d7c01461054d578063ceccfab314610555578063dc749dd71461055d578063dde7e96714610565575f80fd5b8063ba9adeef14610535578063bed3e8131461053d578063c20049f414610545575f80fd5b8063a721b2d311610174578063b51ac0711161014f578063b51ac07114610505578063b5508aa91461050d578063ba414fa614610515578063ba800c911461052d575f80fd5b8063a721b2d3146104ed578063b0464fdc146104f5578063b152ca46146104fd575f80fd5b80637ba9b7ad116101af5780637ba9b7ad146104b357806383ababa9146104bb57806385226c81146104c3578063916a17c6146104d8575f80fd5b806373a3db401461049b57806377def36e146104a35780637ae69730146104ab575f80fd5b80633f7286f4116102b55780636198fb19116102555780636abd223e116102255780636abd223e1461047b5780636d6ce0d0146104835780636dfcbc501461048b5780636f5e275614610493575f80fd5b80636198fb191461044e5780636221b5091461045657806364002a1f1461045e57806366d9a9a014610466575f80fd5b80634ffab9de116102905780634ffab9de1461042e578063564a7435146104365780635d72228f1461043e5780635efe72a914610446575f80fd5b80633f7286f41461041657806342752d411461041e5780634318143714610426575f80fd5b80631ed7831c11610320578063264b524c116102fb578063264b524c146103e95780632ade3880146103f157806336431b3f146104065780633e5e3c231461040e575f80fd5b80631ed7831c146103bb57806320dee15f146103d957806321aeb18c146103e1575f80fd5b80631238212c1161035b5780631238212c1461039b57806314759766146103a35780631b9641bf146103ab5780631e63d2b9146103b3575f80fd5b806305b9f046146103815780630a9254e41461038b5780630b5ad28d14610393575b5f80fd5b6103896105aa565b005b61038961075e565b61038961154c565b6103896116c3565b610389611a77565b610389611c98565b610389611e08565b6103c3612210565b6040516103d0919061bba7565b60405180910390f35b610389612270565b610389612648565b61038961275c565b6103f9612972565b6040516103d0919061bc20565b610389612aae565b6103c3612c12565b6103c3612c70565b610389612cce565b610389612e41565b610389612f6e565b6103896131e8565b61038961330d565b610389613645565b6103896136fb565b6103896138bb565b610389613a9c565b61046e613c6d565b6040516103d0919061bd81565b610389613de6565b610389614169565b61038961445d565b610389614671565b610389614812565b610389614b8b565b610389614d33565b610389614efa565b610389615011565b6104cb6150ef565b6040516103d0919061be1d565b6104e06151ba565b6040516103d0919061be92565b6103896152b0565b6104e06154c4565b6103896155ba565b610389615624565b6104cb61571d565b61051d6157e8565b60405190151581526020016103d0565b6103896158b8565b6103896159c8565b610389615b49565b610389615cbb565b610389615e4b565b6103896160c5565b6103896163ce565b61038961646c565b6103c361688a565b6103896168e8565b610389616d71565b610389616f26565b601f5461051d9060ff1681565b61038961701f565b61038961734e565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905290506106146001620186a061bf54565b602d55604051630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b158015610681575f80fd5b505af1158015610693573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506306cb898391506034015b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835261072e926001600160a01b03909116908690602d9060299060040161c09b565b5f604051808303815f87803b158015610745575f80fd5b505af1158015610757573d5f803e3d5ffd5b5050505050565b602680547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602780549091166112341790556040516107a49061bac1565b604051809103905ff0801580156107bd573d5f803e3d5ffd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600f81527f476174657761795a45564d2e736f6c0000000000000000000000000000000000602082015260265491516024810193909352921660448201526108a1919060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f485cc955000000000000000000000000000000000000000000000000000000001790526174be565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b039384168102919091179182905560208054919092049092167fffffffffffffffffffffffff000000000000000000000000000000000000000090921682178155604080517f2722feee0000000000000000000000000000000000000000000000000000000081529051632722feee926004808401939192918290030181865afa158015610961573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610985919061c104565b602880547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03929092169190911790556040516109c99061bacf565b604051809103905ff0801580156109e2573d5f803e3d5ffd5b50602580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556028546040517f06447d5600000000000000000000000000000000000000000000000000000000815291166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b158015610a7d575f80fd5b505af1158015610a8f573d5f803e3d5ffd5b505050505f805f604051610aa29061badd565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103905ff080158015610adb573d5f803e3d5ffd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392831690811790915560205460405160129360019384935f9391921690610b309061baeb565b610b3f9695949392919061c163565b604051809103905ff080158015610b58573d5f803e3d5ffd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602480546040517fee2815ba000000000000000000000000000000000000000000000000000000008152600160048201529182019290925291169063ee2815ba906044015f604051808303815f87803b158015610beb575f80fd5b505af1158015610bfd573d5f803e3d5ffd5b5050602480546040517fa7cb0507000000000000000000000000000000000000000000000000000000008152600160048201819052928101929092526001600160a01b0316925063a7cb050791506044015f604051808303815f87803b158015610c65575f80fd5b505af1158015610c77573d5f803e3d5ffd5b50506024546020546040516009945061038493506001925f926001600160a01b0391821692911690610ca89061baeb565b610cb79695949392919061c21c565b604051809103905ff080158015610cd0573d5f803e3d5ffd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602480546040517fee2815ba00000000000000000000000000000000000000000000000000000000815261038460048201529182019290925291169063ee2815ba906044015f604051808303815f87803b158015610d64575f80fd5b505af1158015610d76573d5f803e3d5ffd5b5050602480546040517fa7cb050700000000000000000000000000000000000000000000000000000000815261038460048201526001928101929092526001600160a01b0316925063a7cb050791506044015f604051808303815f87803b158015610ddf575f80fd5b505af1158015610df1573d5f803e3d5ffd5b50506028546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152633b9aca006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d91506044015f604051808303815f87803b158015610e6e575f80fd5b505af1158015610e80573d5f803e3d5ffd5b5050505060235f9054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004015f604051808303818588803b158015610ed2575f80fd5b505af1158015610ee4573d5f803e3d5ffd5b50506023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303815f875af1158015610f55573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f79919061c2ab565b506021546026546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e10060248201529116906347e7ef24906044016020604051808303815f875af1158015610fe8573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061100c919061c2ab565b506022546026546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e10060248201529116906347e7ef24906044016020604051808303815f875af115801561107b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061109f919061c2ab565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156110fa575f80fd5b505af115801561110c573d5f803e3d5ffd5b50506026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d5691506024015f604051808303815f87803b15801561117f575f80fd5b505af1158015611191573d5f803e3d5ffd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e10060248201529116925063095ea7b391506044016020604051808303815f875af1158015611203573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611227919061c2ab565b506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e100602482015291169063095ea7b3906044016020604051808303815f875af1158015611296573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112ba919061c2ab565b5060235f9054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004015f604051808303818588803b158015611309575f80fd5b505af115801561131b573d5f803e3d5ffd5b50506023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303815f875af115801561138c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113b0919061c2ab565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561140b575f80fd5b505af115801561141d573d5f803e3d5ffd5b50506040805160a0810182526103218082526001602080840191825283850192835284519081019094525f8085526060840185905260808401528251602980549251151574010000000000000000000000000000000000000000027fffffffffffffffffffffff0000000000000000000000000000000000000000009093166001600160a01b03928316179290921782559151602a8054919093167fffffffffffffffffffffffff000000000000000000000000000000000000000091909116179091559093509150602b906114f3908261c33b565b506080919091015160039091015560408051808201909152620186a080825260016020909201829052602d55602e80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169091179055565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b1790525f602d5551630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b158015611615575f80fd5b505af1158015611627573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b91506034015b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835261072e926001916001600160a01b0316908790602d9060299060040161c3f6565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa158015611711573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611735919061c467565b6026546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156117a7575f80fd5b505af11580156117b9573d5f803e3d5ffd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f60248201529116925063095ea7b391506044016020604051808303815f875af1158015611827573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061184b919061c2ab565b506027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905280517ff48448140000000000000000000000000000000000000000000000000000000081529051919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f4844814916004808201925f9290919082900301818387803b15801561190b575f80fd5b505af115801561191d573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f198184030181526021548383018352620186a084525f602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1681526119c893919289926001600160a01b039091169188919060299060040161c47e565b5f604051808303815f87803b1580156119df575f80fd5b505af11580156119f1573d5f803e3d5ffd5b50506021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa158015611a41573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a65919061c467565b9050611a7183826174dc565b50505050565b6027546040516001600160a01b0390911660248201526001905f9060440160408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b179052602654602754925193945060019384936001600160a01b03928316937f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c93611b38939216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602354611b6f926001600160a01b039091169089905f9081908b90602d9060299061c4d1565b60405180910390a3604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b158015611be1575f80fd5b505af1158015611bf3573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250632810ae639150603401604051602081830303815290604052858486602d60296040518763ffffffff1660e01b8152600401611c679695949392919061c544565b5f6040518083038186803b158015611c7d575f80fd5b505afa158015611c8f573d5f803e3d5ffd5b50505050505050565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015611d5c575f80fd5b505af1158015611d6e573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835261072e925f916001600160a01b0316908790602d9060299060040161c3f6565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa158015611e56573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e7a919061c467565b6027546040516001600160a01b0390911660248201529091505f9060440160408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150620186a0908190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015611f6a575f80fd5b505af1158015611f7c573d5f803e3d5ffd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201525f93506001600160a01b0390911691507fd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a9060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b03909216918b9189918491634d8943bb916004808201926020929091908290030181865afa15801561204b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061206f919061c467565b6040805180820182528a81526001602082015290516120979695949392918d9160299061c574565b60405180910390a3602080546027546040516001600160a01b0392831693637b15118b936120df9316910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815260215483830183528684526001602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815261214b9391928b926001600160a01b03909116918a919060299060040161c47e565b5f604051808303815f87803b158015612162575f80fd5b505af1158015612174573d5f803e3d5ffd5b50506021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa1580156121c4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121e8919061c467565b9050612208836121f8888861bf54565b612202919061bf54565b826174dc565b505050505050565b6060601680548060200260200160405190810160405280929190818152602001828054801561226657602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311612248575b5050505050905090565b601f54604080518082018252601a81527f476174657761795a45564d55706772616465546573742e736f6c00000000000060208083019190915282519081019092525f82526026546122d4936001600160a01b036101009091048116939116617534565b601f546021546026546040516370a0823160e01b81526001600160a01b0391821660048201526101009093048116926001925f9216906370a0823190602401602060405180830381865afa15801561232e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612352919061c467565b6040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b0385166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b1580156123dc575f80fd5b505af11580156123ee573d5f803e3d5ffd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201525f93506001600160a01b0390911691507f5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d979060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b039092169188915f918491634d8943bb916004808201926020929091908290030181865afa1580156124bd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124e1919061c467565b6040805180820182525f81526001602082015290516125089695949392919060299061c5d5565b60405180910390a360275460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b03841690637c0dcb5f9060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835261259c9287916001600160a01b03169060299060040161c655565b5f604051808303815f87803b1580156125b3575f80fd5b505af11580156125c5573d5f803e3d5ffd5b50506021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa158015612615573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612639919061c467565b9050611a71612202848461bf54565b604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156126b1575f80fd5b505af11580156126c3573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef191506034015b6040516020818303038152906040525f600160296040518563ffffffff1660e01b8152600401612734949392919061c68e565b5f6040518083038186803b15801561274a575f80fd5b505afa158015611a71573d5f803e3d5ffd5b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290516001600160a01b03909216926397d340f5926004808401938290030181865afa1580156127b8573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127dc919061c467565b6127e790600161c6be565b67ffffffffffffffff8111156127ff576127ff61c2ca565b6040519080825280601f01601f191660200182016040528015612829576020820181803683370190505b50602b90612837908261c33b565b50604051630618f58760e51b81527f9507fb3d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156128a1575f80fd5b505af11580156128b3573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352612949926002916001600160a01b03169060299060040161c655565b5f604051808303815f87803b158015612960575f80fd5b505af1158015611a71573d5f803e3d5ffd5b6060601e805480602002602001604051908101604052809291908181526020015f905b82821015612aa5575f84815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015612a8e578382905f5260205f20018054612a039061bf67565b80601f0160208091040260200160405190810160405280929190818152602001828054612a2f9061bf67565b8015612a7a5780601f10612a5157610100808354040283529160200191612a7a565b820191905f5260205f20905b815481529060010190602001808311612a5d57829003601f168201915b5050505050815260200190600101906129e6565b505050508152505081526020019060010190612995565b50505050905090565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015612b72575f80fd5b505af1158015612b84573d5f803e3d5ffd5b5050602080546040516001600160a01b039091169350637b15118b9250015b60408051601f198184030181526021548383018352620186a084525f602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815261072e9391926001926001600160a01b039091169188919060299060040161c47e565b6060601880548060200260200160405190810160405280929190818152602001828054801561226657602002820191905f5260205f209081546001600160a01b03168152600190910190602001808311612248575050505050905090565b6060601780548060200260200160405190810160405280929190818152602001828054801561226657602002820191905f5260205f209081546001600160a01b03168152600190910190602001808311612248575050505050905090565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015612d92575f80fd5b505af1158015612da4573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250632810ae6391506034016040516020818303038152906040525f600185602d60296040518763ffffffff1660e01b8152600401612e199695949392919061c544565b5f6040518083038186803b158015612e2f575f80fd5b505afa158015610757573d5f803e3d5ffd5b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015612f05575f80fd5b505af1158015612f17573d5f803e3d5ffd5b5050602080546040516001600160a01b039091169350632810ae639250015b60405160208183030381529060405260018085602d60296040518763ffffffff1660e01b8152600401612e199695949392919061c544565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290515f936002936001600160a01b0316926397d340f592600480830193928290030181865afa158015612fce573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ff2919061c467565b612ffc919061c6d1565b67ffffffffffffffff8111156130145761301461c2ca565b6040519080825280601f01601f19166020018201604052801561303e576020820181803683370190505b5060208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b03909216926397d340f5926004808401938290030181865afa1580156130a1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130c5919061c467565b6130cf919061c6d1565b6130da90600161c6be565b67ffffffffffffffff8111156130f2576130f261c2ca565b6040519080825280601f01601f19166020018201604052801561311c576020820181803683370190505b50602b9061312a908261c33b565b50604051630618f58760e51b81527f9507fb3d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015613194575f80fd5b505af11580156131a6573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b9150603401612ba3565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152600190620249f090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015613259575f80fd5b505af115801561326b573d5f803e3d5ffd5b5050602080546040516001600160a01b03909116935063c5356f6d9250015b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526132e49287916001600160a01b031690879060299060040161c709565b5f604051808303815f87803b1580156132fb575f80fd5b505af1158015612208573d5f803e3d5ffd5b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa15801561335b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061337f919061c467565b6026546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156133f1575f80fd5b505af1158015613403573d5f803e3d5ffd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f60248201529116925063095ea7b391506044016020604051808303815f875af1158015613471573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613495919061c2ab565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156134f0575f80fd5b505af1158015613502573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526135979287916001600160a01b03169060299060040161c655565b5f604051808303815f87803b1580156135ae575f80fd5b505af11580156135c0573d5f803e3d5ffd5b50506021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa158015613610573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613634919061c467565b905061364082826174dc565b505050565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b1790525f602d5551630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e09060240161066a565b6027546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b1580156137e1575f80fd5b505af11580156137f3573d5f803e3d5ffd5b505060215460265460275460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b039283169450911691507f306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e49060340160408051601f1981840301815290829052613873918690602d9060299061c748565b60405180910390a3602080546027546040516001600160a01b03928316936306cb8983936106d19316910160609190911b6bffffffffffffffffffffffff1916815260140190565b6023546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa158015613909573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061392d919061c467565b6023546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526101236004820152602481018390529192506001600160a01b03169063a9059cbb906044016020604051808303815f875af1158015613997573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906139bb919061c2ab565b505f600190507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b158015613a1b575f80fd5b505af1158015613a2d573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef19150603401604051602081830303815290604052858460296040518563ffffffff1660e01b8152600401611c67949392919061c68e565b6027546040516001600160a01b0390911660248201526001905f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905260265490517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015613b6d575f80fd5b505af1158015613b7f573d5f803e3d5ffd5b50506023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f60248201529116925063095ea7b391506044016020604051808303815f875af1158015613bed573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613c11919061c2ab565b505f602d55604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401611bca565b6060601b805480602002602001604051908101604052809291908181526020015f905b82821015612aa5578382905f5260205f2090600202016040518060400160405290815f82018054613cc09061bf67565b80601f0160208091040260200160405190810160405280929190818152602001828054613cec9061bf67565b8015613d375780601f10613d0e57610100808354040283529160200191613d37565b820191905f5260205f20905b815481529060010190602001808311613d1a57829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015613dce57602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411613d7b5790505b50505050508152505081526020019060010190613c90565b6021546026546040516370a0823160e01b81526001600160a01b03918216600482015260019262030d40925f929116906370a0823190602401602060405180830381865afa158015613e3a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613e5e919061c467565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b0390911660848201529091508290737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015613eee575f80fd5b505af1158015613f00573d5f803e3d5ffd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201525f93506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b03909216918a9188918491634d8943bb916004808201926020929091908290030181865afa158015613fcf573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613ff3919061c467565b6040805180820182528c815260016020820152905161401a9695949392919060299061c5d5565b60405180910390a3602080546027546040516001600160a01b039283169363c5356f6d936140629316910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526140bc9289916001600160a01b031690899060299060040161c709565b5f604051808303815f87803b1580156140d3575f80fd5b505af11580156140e5573d5f803e3d5ffd5b50506021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa158015614135573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614159919061c467565b9050610757826121f8878661bf54565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa1580156141b7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906141db919061c467565b6026546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801561424d575f80fd5b505af115801561425f573d5f803e3d5ffd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f60248201529116925063095ea7b391506044016020604051808303815f875af11580156142cd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906142f1919061c2ab565b506027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905280517ff48448140000000000000000000000000000000000000000000000000000000081529051919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f4844814916004808201925f9290919082900301818387803b1580156143b1575f80fd5b505af11580156143c3573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526119c89288916001600160a01b0316908790602d9060299060040161c3f6565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290515f936002936001600160a01b0316926397d340f592600480830193928290030181865afa1580156144bd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906144e1919061c467565b6144eb919061c6d1565b67ffffffffffffffff8111156145035761450361c2ca565b6040519080825280601f01601f19166020018201604052801561452d576020820181803683370190505b5060208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b03909216926397d340f5926004808401938290030181865afa158015614590573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906145b4919061c467565b6145be919061c6d1565b6145c990600161c6be565b67ffffffffffffffff8111156145e1576145e161c2ca565b6040519080825280601f01601f19166020018201604052801561460b576020820181803683370190505b50602b90614619908261c33b565b50604051630618f58760e51b81527f9507fb3d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e09060240161066a565b60208054604080517f97d340f50000000000000000000000000000000000000000000000000000000081529051600193620249f0936001600160a01b0316926397d340f592600480830193928290030181865afa1580156146d4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906146f8919061c467565b61470390600161c6be565b67ffffffffffffffff81111561471b5761471b61c2ca565b6040519080825280601f01601f191660200182016040528015614745576020820181803683370190505b50602b90614753908261c33b565b50604051630618f58760e51b81527f9507fb3d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b1580156147be575f80fd5b505af11580156147d0573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c5356f6d915060340161328a565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa158015614860573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614884919061c467565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015614912575f80fd5b505af1158015614924573d5f803e3d5ffd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201525f93506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160408051601f198184030181528282526022547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b039092169188915f918491634d8943bb916004808201926020929091908290030181865afa1580156149f3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614a17919061c467565b6040805180820182525f8152600160208201529051614a3e9695949392919060299061c5d5565b60405180910390a3602080546027546040516001600160a01b0392831693637c0dcb5f93614a869316910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526022547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352614ade9287916001600160a01b03169060299060040161c655565b5f604051808303815f87803b158015614af5575f80fd5b505af1158015614b07573d5f803e3d5ffd5b50506022546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a08231906024015b602060405180830381865afa158015614b58573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614b7c919061c467565b9050613640612202848461bf54565b6027546040516001600160a01b0390911660248201526001905f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905280517ff48448140000000000000000000000000000000000000000000000000000000081529051919250600191737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f4844814916004808301925f92919082900301818387803b158015614c4f575f80fd5b505af1158015614c61573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f1981840301815260225483830183528684526001602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b168152614d0a93919289926001600160a01b039091169189919060299060040161c47e565b5f604051808303815f87803b158015614d21575f80fd5b505af1158015611c8f573d5f803e3d5ffd5b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290516001600160a01b03909216926397d340f5926004808401938290030181865afa158015614d8f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614db3919061c467565b614dbe90600161c6be565b67ffffffffffffffff811115614dd657614dd661c2ca565b6040519080825280601f01601f191660200182016040528015614e00576020820181803683370190505b50602b90614e0e908261c33b565b50604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015614e78575f80fd5b505af1158015614e8a573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef1915060340160405160208183030381529060405260018060296040518563ffffffff1660e01b8152600401612734949392919061c68e565b604051630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152620249f090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015614f68575f80fd5b505af1158015614f7a573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b0316925063c5356f6d915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835261072e925f916001600160a01b031690879060299060040161c709565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801561507a575f80fd5b505af115801561508c573d5f803e3d5ffd5b505060208054604080515f8152928301908190526021547f7c0dcb5f000000000000000000000000000000000000000000000000000000009091526001600160a01b039182169450637c0dcb5f935061294992916001911660296024840161c655565b6060601a805480602002602001604051908101604052809291908181526020015f905b82821015612aa5578382905f5260205f2001805461512f9061bf67565b80601f016020809104026020016040519081016040528092919081815260200182805461515b9061bf67565b80156151a65780601f1061517d576101008083540402835291602001916151a6565b820191905f5260205f20905b81548152906001019060200180831161518957829003601f168201915b505050505081526020019060010190615112565b6060601d805480602002602001604051908101604052809291908181526020015f905b82821015612aa5575f8481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561529857602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116152455790505b505050505081525050815260200190600101906151dd565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290515f936002936001600160a01b0316926397d340f592600480830193928290030181865afa158015615310573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190615334919061c467565b61533e919061c6d1565b67ffffffffffffffff8111156153565761535661c2ca565b6040519080825280601f01601f191660200182016040528015615380576020820181803683370190505b5060208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b03909216926397d340f5926004808401938290030181865afa1580156153e3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190615407919061c467565b615411919061c6d1565b61541c90600161c6be565b67ffffffffffffffff8111156154345761543461c2ca565b6040519080825280601f01601f19166020018201604052801561545e576020820181803683370190505b50602b9061546c908261c33b565b50604051630618f58760e51b81527f9507fb3d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024016115fe565b6060601c805480602002602001604051908101604052809291908181526020015f905b82821015612aa5575f8481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156155a257602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001906004019060208260030104928301926001038202915080841161554f5790505b505050505081525050815260200190600101906154e7565b60015f6155ca82620186a061bf54565b604051630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024016147a7565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156156e8575f80fd5b505af11580156156fa573d5f803e3d5ffd5b5050602080546040516001600160a01b0390911693506306cb89839250016106d1565b60606019805480602002602001604051908101604052809291908181526020015f905b82821015612aa5578382905f5260205f2001805461575d9061bf67565b80601f01602080910402602001604051908101604052809291908181526020018280546157899061bf67565b80156157d45780601f106157ab576101008083540402835291602001916157d4565b820191905f5260205f20905b8154815290600101906020018083116157b757829003601f168201915b505050505081526020019060010190615740565b6008545f9060ff16156157ff575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c6564000000000000000000000000000000000000000000000000000060248301525f9163667f9d7090604401602060405180830381865afa15801561588d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906158b1919061c467565b1415905090565b604051630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015615921575f80fd5b505af1158015615933573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352612949925f916001600160a01b03169060299060040161c655565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015615a8c575f80fd5b505af1158015615a9e573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f198184030181526021548383018352620186a084525f6020850181905292517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b16815261072e949293926001600160a01b0390921691889160299060040161c47e565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015615c0d575f80fd5b505af1158015615c1f573d5f803e3d5ffd5b5050602080546027546040805160609290921b6bffffffffffffffffffffffff191693820193909352825180820360140181526021546074830185525f6034840181815260549094015293517f7b15118b0000000000000000000000000000000000000000000000000000000081526001600160a01b039384169650637b15118b955061072e949193600193921691889160299060040161c47e565b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015615d7f575f80fd5b505af1158015615d91573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152602154838301909252916001916001600160a01b031690869080615e0285620186a061bf54565b81525f6020909101526040517fffffffff0000000000000000000000000000000000000000000000000000000060e088901b16815261072e95949392919060299060040161c47e565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290515f936002936001600160a01b0316926397d340f592600480830193928290030181865afa158015615eab573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190615ecf919061c467565b615ed9919061c6d1565b67ffffffffffffffff811115615ef157615ef161c2ca565b6040519080825280601f01601f191660200182016040528015615f1b576020820181803683370190505b5060208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b03909216926397d340f5926004808401938290030181865afa158015615f7e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190615fa2919061c467565b615fac919061c6d1565b615fb790600161c6be565b67ffffffffffffffff811115615fcf57615fcf61c2ca565b6040519080825280601f01601f191660200182016040528015615ff9576020820181803683370190505b50602b90616007908261c33b565b50604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015616071575f80fd5b505af1158015616083573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250632810ae639150603401612f36565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa158015616113573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190616137919061c467565b6021546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526101236004820152602481018390529192506001600160a01b03169063a9059cbb906044016020604051808303815f875af11580156161a1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906161c5919061c2ab565b506028546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015616235575f80fd5b505af1158015616247573d5f803e3d5ffd5b50506021546040517ff687d12a000000000000000000000000000000000000000000000000000000008152600a60048201526001600160a01b03909116925063f687d12a91506024015f604051808303815f87803b1580156162a7575f80fd5b505af11580156162b9573d5f803e3d5ffd5b5050604051630618f58760e51b81527ffe382aa7000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015b5f604051808303815f87803b158015616327575f80fd5b505af1158015616339573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526132e49287916001600160a01b03169060299060040161c655565b604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015616437575f80fd5b505af1158015616449573d5f803e3d5ffd5b5050602080546040516001600160a01b0390911693506397a1cef1925001612701565b6023546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa1580156164ba573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906164de919061c467565b6023546020546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa15801561652d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190616551919061c467565b6028546026546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152929350163190600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156165cc575f80fd5b505af11580156165de573d5f803e3d5ffd5b50506023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f60248201529116925063095ea7b391506044016020604051808303815f875af115801561664c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190616670919061c2ab565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156166cb575f80fd5b505af11580156166dd573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef19150603401604051602081830303815290604052878460296040518563ffffffff1660e01b815260040161674c949392919061c68e565b5f6040518083038186803b158015616762575f80fd5b505afa158015616774573d5f803e3d5ffd5b50506023546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa1580156167c4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906167e8919061c467565b90506167f485826174dc565b6023546020546040516370a0823160e01b81526001600160a01b0391821660048201525f9291909116906370a0823190602401602060405180830381865afa158015616842573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190616866919061c467565b905061687285826174dc565b602854611c8f9085906001600160a01b0316316174dc565b6060601580548060200260200160405190810160405280929190818152602001828054801561226657602002820191905f5260205f209081546001600160a01b03168152600190910190602001808311612248575050505050905090565b6023546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa158015616936573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061695a919061c467565b6023546020546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa1580156169a9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906169cd919061c467565b6028546027546040516001600160a01b0391821660248201529293501631905f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905260265490517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015616aa4575f80fd5b505af1158015616ab6573d5f803e3d5ffd5b50506023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f60248201529116925063095ea7b391506044016020604051808303815f875af1158015616b24573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190616b48919061c2ab565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b158015616ba3575f80fd5b505af1158015616bb5573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250632810ae639150603401604051602081830303815290604052888486602d60296040518763ffffffff1660e01b8152600401616c299695949392919061c544565b5f6040518083038186803b158015616c3f575f80fd5b505afa158015616c51573d5f803e3d5ffd5b50506023546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa158015616ca1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190616cc5919061c467565b9050616cd186826174dc565b6023546020546040516370a0823160e01b81526001600160a01b0391821660048201525f9291909116906370a0823190602401602060405180830381865afa158015616d1f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190616d43919061c467565b9050616d4f86826174dc565b602854616d679086906001600160a01b0316316174dc565b5050505050505050565b60265460275460405160609190911b6bffffffffffffffffffffffff19166020820152600191829182916001600160a01b0316907f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160408051601f1981840301815260235483830183525f808552600160208601529251616e0c9492936001600160a01b03909216928a929182919060299061c5d5565b60405180910390a3604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015616e7d575f80fd5b505af1158015616e8f573d5f803e3d5ffd5b50506020805460275460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef19150603401604051602081830303815290604052848460296040518563ffffffff1660e01b8152600401616efe949392919061c68e565b5f6040518083038186803b158015616f14575f80fd5b505afa158015612208573d5f803e3d5ffd5b6027546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015616fea575f80fd5b505af1158015616ffc573d5f803e3d5ffd5b5050602080546040516001600160a01b039091169350637b15118b925001611665565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526001925f9216906370a0823190602401602060405180830381865afa15801561706d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190617091919061c467565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801561711f575f80fd5b505af1158015617131573d5f803e3d5ffd5b505060265460275460405160609190911b6bffffffffffffffffffffffff191660208201525f93506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b039092169188915f918491634d8943bb916004808201926020929091908290030181865afa158015617200573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190617224919061c467565b6040805180820182525f815260016020820152905161724b9695949392919060299061c5d5565b60405180910390a3602080546027546040516001600160a01b0392831693637c0dcb5f936172939316910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526172eb9287916001600160a01b03169060299060040161c655565b5f604051808303815f87803b158015617302575f80fd5b505af1158015617314573d5f803e3d5ffd5b50506021546026546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401614b3d565b6021546026546040516370a0823160e01b81526001600160a01b0391821660048201526002925f9216906370a0823190602401602060405180830381865afa15801561739c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906173c0919061c467565b6021549091506001600160a01b031663a9059cbb6101236173e260018561bf54565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af1158015617442573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190617466919061c2ab565b50604051630618f58760e51b81527ffe382aa7000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401616310565b5f6174c761baf9565b6174d2848483617549565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c5490604401616efe565b61753c61baf9565b61075785858584866175c3565b5f8061755585846176bb565b90506175b86040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f787900000081525082866040516020016175a392919061c796565b604051602081830303815290604052856176c6565b9150505b9392505050565b6040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201528190737109709ecfa91a80626ff3989d68f67f5b1dd12d9081906306447d56906024015f604051808303815f87803b158015617632575f80fd5b505af1925050508015617643575060015b61765857617653878787876176f3565b611c8f565b617664878787876176f3565b806001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561769c575f80fd5b505af11580156176ae573d5f803e3d5ffd5b5050505050505050505050565b5f6175bc838361770b565b60c0810151515f90156176e9576176e284848460c00151617725565b90506175bc565b6176e284846178c3565b5f6176fe84836179a8565b90506107578582856179b3565b5f6177168383617d61565b6175bc838360200151846176c6565b5f8061772f617d70565b90505f61773c8683617e3f565b90505f61775282606001518360200151856182c8565b90505f617761838389896184d5565b90505f61776d82619341565b602081015181519192509060030b156177e05789826040015160405160200161779792919061c7ce565b60408051601f19818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526177d79160040161c82f565b60405180910390fd5b5f6178226040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a200000000000000000000000815250836001619502565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d9061787590849060040161c82f565b602060405180830381865afa158015617890573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906178b4919061c104565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081525f908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc9259061791790879060040161c82f565b5f60405180830381865afa158015617931573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052617958919081019061c8ff565b90505f617985828560405160200161797192919061c931565b6040516020818303038152906040526196f1565b90506001600160a01b0381166174d257848460405160200161779792919061c945565b5f6177168383619702565b6040517f667f9d700000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201527fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d905f90829063667f9d7090604401602060405180830381865afa158015617a4c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190617a70919061c467565b905080617c09575f617a818661970e565b604080518082018252600581527f352e302e300000000000000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152909150617b0a905b6040805180820182525f8082526020918201528151808301909252845182528085019082015290619801565b80617b1557505f8451115b15617b93576040517f4f1ef2860000000000000000000000000000000000000000000000000000000081526001600160a01b03871690634f1ef28690617b61908890889060040161c796565b5f604051808303815f87803b158015617b78575f80fd5b505af1158015617b8a573d5f803e3d5ffd5b50505050617c03565b6040517f3659cfe60000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152871690633659cfe6906024015f604051808303815f87803b158015617bec575f80fd5b505af1158015617bfe573d5f803e3d5ffd5b505050505b50610757565b805f617c148261970e565b604080518082018252600581527f352e302e300000000000000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152909150617c7590617ade565b80617c8057505f8551115b15617d00576040517f9623609d0000000000000000000000000000000000000000000000000000000081526001600160a01b03831690639623609d90617cce908a908a908a9060040161c9d5565b5f604051808303815f87803b158015617ce5575f80fd5b505af1158015617cf7573d5f803e3d5ffd5b50505050611c8f565b6040517f99a88ec40000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015287811660248301528316906399a88ec4906044015f604051808303815f87803b15801561769c575f80fd5b617d6c82825f619814565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c90617df790849060040161ca05565b5f60405180830381865afa158015617e11573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052617e38919081019061ca4b565b9250505090565b617e716040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b5f737109709ecfa91a80626ff3989d68f67f5b1dd12d9050617ebb6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b617ec485619913565b60208201525f617ed386619cec565b90505f836001600160a01b031663d930a0e66040518163ffffffff1660e01b81526004015f60405180830381865afa158015617f11573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052617f38919081019061ca4b565b86838560200151604051602001617f52949392919061ca90565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291505f906001600160a01b038616906360f9bb1190617fa990859060040161c82f565b5f60405180830381865afa158015617fc3573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052617fea919081019061ca4b565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f69061803290849060040161cb60565b602060405180830381865afa15801561804d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190618071919061c2ab565b6180865781604051602001617797919061cbb1565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906180cb90849060040161cc35565b5f60405180830381865afa1580156180e5573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261810c919081019061ca4b565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f69061815390849060040161cc86565b602060405180830381865afa15801561816e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190618192919061c2ab565b15618223576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906181dc90849060040161cc86565b5f60405180830381865afa1580156181f6573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261821d919081019061ca4b565b60408501525b846001600160a01b03166349c4fac882865f0151604051602001618247919061ccd7565b6040516020818303038152906040526040518363ffffffff1660e01b815260040161827392919061cd35565b5f60405180830381865afa15801561828d573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526182b4919081019061ca4b565b606085015250608083015250949350505050565b60408051600480825260a082019092526060915f9190816020015b60608152602001906001900390816182e35790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250815f815181106183425761834261cd59565b60200260200101819052506040518060400160405280600381526020017f2d726c0000000000000000000000000000000000000000000000000000000000815250816001815181106183965761839661cd59565b6020026020010181905250846040516020016183b2919061cd86565b604051602081830303815290604052816002815181106183d4576183d461cd59565b6020026020010181905250826040516020016183f0919061cde4565b604051602081830303815290604052816003815181106184125761841261cd59565b60200260200101819052505f61842782619341565b602080820151604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000818501908152825180840184525f80825290860152825180840190935290518252928101929092529192506184b6906040805180820182525f8082526020918201528151808301909252845182528085019082015290619f68565b6184cb5785604051602001617797919061ce1c565b9695505050505050565b60a08101516040805180820182525f80825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d9015618524565b511590565b618698578260200151156185e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a4016177d7565b8260c0015115618698576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a4016177d7565b6040805160ff80825261200082019092525f91816020015b60608152602001906001900390816186b05790505090505f6040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061870a9061ce99565b935060ff168151811061871f5761871f61cd59565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e3700000000000000000000000000000000000000815250604051602001618770919061ceb7565b60405160208183030381529060405282828061878b9061ce99565b935060ff16815181106187a0576187a061cd59565b60200260200101819052506040518060400160405280600681526020017f6465706c6f7900000000000000000000000000000000000000000000000000008152508282806187ed9061ce99565b935060ff16815181106188025761880261cd59565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d6500000000000000000000000000000000000081525082828061884f9061ce99565b935060ff16815181106188645761886461cd59565b602002602001018190525087602001518282806188809061ce99565b935060ff16815181106188955761889561cd59565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e7472616374506174680000000000000000000000000000000000008152508282806188e29061ce99565b935060ff16815181106188f7576188f761cd59565b60209081029190910101528751828261890f8161ce99565b935060ff16815181106189245761892461cd59565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e496400000000000000000000000000000000000000000000008152508282806189719061ce99565b935060ff16815181106189865761898661cd59565b602002602001018190525061899a46619fc6565b82826189a58161ce99565b935060ff16815181106189ba576189ba61cd59565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c650000000000000000000000000000000000815250828280618a079061ce99565b935060ff1681518110618a1c57618a1c61cd59565b602002602001018190525086828280618a349061ce99565b935060ff1681518110618a4957618a4961cd59565b6020908102919091010152855115618b6c5760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f6465000000000000000000000060208201528282618a9a8161ce99565b935060ff1681518110618aaf57618aaf61cd59565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d90618aff90899060040161c82f565b5f60405180830381865afa158015618b19573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052618b40919081019061ca4b565b8282618b4b8161ce99565b935060ff1681518110618b6057618b6061cd59565b60200260200101819052505b846020015115618c3c5760408051808201909152601281527f2d2d766572696679536f75726365436f6465000000000000000000000000000060208201528282618bb58161ce99565b935060ff1681518110618bca57618bca61cd59565b60200260200101819052506040518060400160405280600581526020017f66616c7365000000000000000000000000000000000000000000000000000000815250828280618c179061ce99565b935060ff1681518110618c2c57618c2c61cd59565b6020026020010181905250618e01565b618c7361851f8660a001516040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b618d065760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282618cb68161ce99565b935060ff1681518110618ccb57618ccb61cd59565b60200260200101819052508460a00151604051602001618ceb919061cd86565b604051602081830303815290604052828280618c179061ce99565b8460c00151158015618d485750604080890151815180830183525f80825260209182015282518084019093528151835290810190820152618d4690511590565b155b15618e015760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282618d8c8161ce99565b935060ff1681518110618da157618da161cd59565b6020026020010181905250618db58861a063565b604051602001618dc5919061cd86565b604051602081830303815290604052828280618de09061ce99565b935060ff1681518110618df557618df561cd59565b60200260200101819052505b604080860151815180830183525f80825260209182015282518084019093528151835290810190820152618e3490511590565b618ec95760408051808201909152600b81527f2d2d72656c61796572496400000000000000000000000000000000000000000060208201528282618e778161ce99565b935060ff1681518110618e8c57618e8c61cd59565b60200260200101819052508460400151828280618ea89061ce99565b935060ff1681518110618ebd57618ebd61cd59565b60200260200101819052505b606085015115618fe65760408051808201909152600681527f2d2d73616c74000000000000000000000000000000000000000000000000000060208201528282618f128161ce99565b935060ff1681518110618f2757618f2761cd59565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e8906024015f60405180830381865afa158015618f93573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052618fba919081019061ca4b565b8282618fc58161ce99565b935060ff1681518110618fda57618fda61cd59565b60200260200101819052505b60e0850151511561908c5760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826190308161ce99565b935060ff16815181106190455761904561cd59565b60200260200101819052506190608560e001515f0151619fc6565b828261906b8161ce99565b935060ff16815181106190805761908061cd59565b60200260200101819052505b60e085015160200151156191365760408051808201909152600a81527f2d2d676173507269636500000000000000000000000000000000000000000000602082015282826190d98161ce99565b935060ff16815181106190ee576190ee61cd59565b602002602001018190525061910a8560e0015160200151619fc6565b82826191158161ce99565b935060ff168151811061912a5761912a61cd59565b60200260200101819052505b60e085015160400151156191e05760408051808201909152600e81527f2d2d6d6178466565506572476173000000000000000000000000000000000000602082015282826191838161ce99565b935060ff16815181106191985761919861cd59565b60200260200101819052506191b48560e0015160400151619fc6565b82826191bf8161ce99565b935060ff16815181106191d4576191d461cd59565b60200260200101819052505b60e0850151606001511561928a5760408051808201909152601681527f2d2d6d61785072696f72697479466565506572476173000000000000000000006020820152828261922d8161ce99565b935060ff16815181106192425761924261cd59565b602002602001018190525061925e8560e0015160600151619fc6565b82826192698161ce99565b935060ff168151811061927e5761927e61cd59565b60200260200101819052505b5f8160ff1667ffffffffffffffff8111156192a7576192a761c2ca565b6040519080825280602002602001820160405280156192da57816020015b60608152602001906001900390816192c55790505b5090505f5b8260ff168160ff16101561933257838160ff16815181106193025761930261cd59565b6020026020010151828260ff168151811061931f5761931f61cd59565b60209081029190910101526001016192df565b5093505050505b949350505050565b61936760405180606001604052805f60030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d925f91849163d145736c916193ec9186910161cf0e565b5f60405180830381865afa158015619406573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261942d919081019061ca4b565b90505f61943a868361ab3f565b90505f846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b8152600401619469919061be1d565b5f604051808303815f875af1158015619484573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526194ab919081019061cf54565b805190915060030b158015906194c45750602081015151155b80156194d35750604081015151155b156184cb57815f815181106194ea576194ea61cd59565b6020026020010151604051602001617797919061d003565b60605f619535856040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f8082526020918201528151808301909252865182528087019082015290915061956b9082905b9061ac91565b156196c3575f6195e5826195df846195d96195ac8a6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925282518252918201519181019190915290565b9061acb7565b9061ad15565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015290915061964890829061ac91565b156196b157604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526196ae905b829061ad99565b90505b6196ba8161adbe565b925050506175bc565b82156196dc57848460405160200161779792919061d1e0565b505060408051602081019091525f81526175bc565b5f808251602084015ff09392505050565b617d6c82826001619814565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fad3cb1cc0000000000000000000000000000000000000000000000000000000017905290516060915f9182916001600160a01b03861691619782919061d267565b5f60405180830381855afa9150503d805f81146197ba576040519150601f19603f3d011682016040523d82523d5f602084013e6197bf565b606091505b50915091508180156197d2575060208151115b156197eb5780806020019051810190619339919061ca4b565b505060408051602081019091525f815292915050565b5f61980c838361ae23565b159392505050565b8160a001511561982357505050565b5f61982f84848461aefb565b90505f61983b82619341565b602081015181519192509060030b1580156198d55750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526198d5906040805180820182525f80825260209182015281518083019092528451825280850190820152619565565b156198e257505050505050565b60408201515115619902578160400151604051602001617797919061d272565b80604051602001617797919061d2c9565b60605f619946836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506199aa905b8290619f68565b15619a1857604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526175bc90619a1390839061b490565b61adbe565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152619a79905b829061b518565b600103619b4457604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152619ade906196a7565b50604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526175bc90619a13905b839061ad99565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152619ba2906199a3565b15619cd557604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820181905284518086019095529251845283015290619c0990839061b5ac565b90505f8160018351619c1b919061bf54565b81518110619c2b57619c2b61cd59565b60200260200101519050619ccc619a13619ca06040518060400160405280600581526020017f2e6a736f6e0000000000000000000000000000000000000000000000000000008152506040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925285518252808601908201529061b490565b95945050505050565b82604051602001617797919061d320565b50919050565b60605f619d1f836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152909150619d80906199a3565b15619d8e576175bc8161adbe565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152619dec90619a72565b600103619e5557604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526175bc90619a1390619b3d565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152619eb3906199a3565b15619cd557604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820181905284518086019095529251845283015290619f1a90839061b5ac565b9050600181511115619f56578060028251619f35919061bf54565b81518110619f4557619f4561cd59565b602002602001015192505050919050565b5082604051602001617797919061d320565b805182515f911115619f7b57505f6174d6565b8151835160208501515f9291619f909161c6be565b619f9a919061bf54565b905082602001518103619fb15760019150506174d6565b82516020840151819020912014905092915050565b60605f619fd28361b657565b60010190505f8167ffffffffffffffff811115619ff157619ff161c2ca565b6040519080825280601f01601f19166020018201604052801561a01b576020820181803683370190505b5090508181016020015b5f19017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461a02557509392505050565b604081810151815180830183525f808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e534544000000000000000000000000000000000000000000008184019081528551808701875283815284019290925284518086019095525184529083015260609161a0ee905b8290619801565b1561a12e57505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a18c9061a0e7565b1561a1cc57505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a22a9061a0e7565b1561a26a57505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a2c89061a0e7565b8061a32c5750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a32c9061a0e7565b1561a36c57505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a3ca9061a0e7565b8061a42e5750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a42e9061a0e7565b1561a46e57505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a4cc9061a0e7565b8061a5305750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a5309061a0e7565b1561a57057505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a5ce9061a0e7565b8061a6325750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a6329061a0e7565b1561a67257505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a6d09061a0e7565b1561a71057505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a76e9061a0e7565b1561a7ae57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a80c9061a0e7565b1561a84c57505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a8aa9061a0e7565b1561a8ea57505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a9489061a0e7565b1561a98857505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261a9e69061a0e7565b8061aa4a5750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261aa4a9061a0e7565b1561aa8a57505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261aae89061a0e7565b1561ab2857505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151617797929060200161d3f0565b6060805f5b845181101561abc9578185828151811061ab605761ab6061cd59565b602002602001015160405160200161ab7992919061c931565b60405160208183030381529060405291506001855161ab98919061bf54565b811461abc1578160405160200161abaf919061d53e565b60405160208183030381529060405291505b60010161ab44565b50604080516003808252608082019092525f91816020015b606081526020019060019003908161abe157905050905083815f8151811061ac0b5761ac0b61cd59565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061ac5f5761ac5f61cd59565b6020026020010181905250818160028151811061ac7e5761ac7e61cd59565b6020908102919091010152949350505050565b60208083015183518351928401515f9361acae929184919061b738565b14159392505050565b604080518082019091525f80825260208201525f61ace5845f01518560200151855f0151866020015161b847565b905083602001518161acf7919061bf54565b8451859061ad0690839061bf54565b90525060208401525090919050565b604080518082019091525f808252602082015281518351101561ad395750816174d6565b602080830151908401516001911461ad605750815160208481015190840151829020919020145b801561ad915782518451859061ad7790839061bf54565b905250825160208501805161ad8d90839061c6be565b9052505b509192915050565b604080518082019091525f808252602082015261adb783838361b963565b5092915050565b60605f825f015167ffffffffffffffff81111561addd5761addd61c2ca565b6040519080825280601f01601f19166020018201604052801561ae07576020820181803683370190505b5090505f60208201905061adb7818560200151865f015161ba09565b815181515f919081111561ae35575081515b602080850151908401515f5b8381101561aeec578251825180821461aebc575f19602087101561ae9b5760018461ae6d89602061bf54565b61ae77919061c6be565b61ae8290600861d576565b61ae8d90600261d670565b61ae97919061bf54565b1990505b818116838216818103911461aeb95797506174d69650505050505050565b50505b61aec760208661c6be565b945061aed460208561c6be565b9350505060208161aee5919061c6be565b905061ae41565b50845186516184cb919061d67b565b60605f61af06617d70565b6040805160ff80825261200082019092529192505f9190816020015b606081526020019060019003908161af225790505090505f6040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061af7c9061ce99565b935060ff168151811061af915761af9161cd59565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161afe2919061d69a565b60405160208183030381529060405282828061affd9061ce99565b935060ff168151811061b0125761b01261cd59565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061b05f9061ce99565b935060ff168151811061b0745761b07461cd59565b60200260200101819052508260405160200161b090919061cde4565b60405160208183030381529060405282828061b0ab9061ce99565b935060ff168151811061b0c05761b0c061cd59565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061b10d9061ce99565b935060ff168151811061b1225761b12261cd59565b602002602001018190525061b137878461ba82565b828261b1428161ce99565b935060ff168151811061b1575761b15761cd59565b60209081029190910101528551511561b2025760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261b1a98161ce99565b935060ff168151811061b1be5761b1be61cd59565b602002602001018190525061b1d6865f01518461ba82565b828261b1e18161ce99565b935060ff168151811061b1f65761b1f661cd59565b60200260200101819052505b85608001511561b2705760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261b24b8161ce99565b935060ff168151811061b2605761b26061cd59565b602002602001018190525061b2d6565b841561b2d65760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261b2b58161ce99565b935060ff168151811061b2ca5761b2ca61cd59565b60200260200101819052505b6040860151511561b3725760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261b3208161ce99565b935060ff168151811061b3355761b33561cd59565b6020026020010181905250856040015182828061b3519061ce99565b935060ff168151811061b3665761b36661cd59565b60200260200101819052505b85606001511561b3dc5760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d65730000000000000000000000006020820152828261b3bb8161ce99565b935060ff168151811061b3d05761b3d061cd59565b60200260200101819052505b5f8160ff1667ffffffffffffffff81111561b3f95761b3f961c2ca565b60405190808252806020026020018201604052801561b42c57816020015b606081526020019060019003908161b4175790505b5090505f5b8260ff168160ff16101561b48457838160ff168151811061b4545761b45461cd59565b6020026020010151828260ff168151811061b4715761b47161cd59565b602090810291909101015260010161b431565b50979650505050505050565b604080518082019091525f808252602082015281518351101561b4b45750816174d6565b8151835160208501515f929161b4c99161c6be565b61b4d3919061bf54565b6020840151909150600190821461b4f4575082516020840151819020908220145b801561b50f5783518551869061b50b90839061bf54565b9052505b50929392505050565b5f80825f015161b538855f01518660200151865f0151876020015161b847565b61b542919061c6be565b90505b8351602085015161b556919061c6be565b811161adb7578161b5668161d6cb565b925050825f015161b59b85602001518361b580919061bf54565b865161b58c919061bf54565b83865f0151876020015161b847565b61b5a5919061c6be565b905061b545565b60605f61b5b9848461b518565b61b5c490600161c6be565b67ffffffffffffffff81111561b5dc5761b5dc61c2ca565b60405190808252806020026020018201604052801561b60f57816020015b606081526020019060019003908161b5fa5790505b5090505f5b815181101561b64f5761b62a619a13868661ad99565b82828151811061b63c5761b63c61cd59565b602090810291909101015260010161b614565b509392505050565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061b69f577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061b6cb576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061b6e957662386f26fc10000830492506010015b6305f5e100831061b701576305f5e100830492506008015b612710831061b71557612710830492506004015b6064831061b727576064830492506002015b600a83106174d65760010192915050565b5f8085841161b83d576020841161b7e9575f841561b78157600161b75d86602061bf54565b61b76890600861d576565b61b77390600261d670565b61b77d919061bf54565b1990505b835181168561b790898961c6be565b61b79a919061bf54565b805190935082165b81811461b7d45787841161b7bc5787945050505050619339565b8361b7c68161d6e3565b94505082845116905061b7a2565b61b7de878561c6be565b945050505050619339565b83832061b7f6858861bf54565b61b800908761c6be565b91505b85821061b83b5784822080820361b8285761b81e868461c6be565b9350505050619339565b61b83360018461bf54565b92505061b803565b505b5092949350505050565b5f838186851161b94e576020851161b8fe575f851561b89157600161b86d87602061bf54565b61b87890600861d576565b61b88390600261d670565b61b88d919061bf54565b1990505b845181165f8761b8a18b8b61c6be565b61b8ab919061bf54565b855190915083165b82811461b8f05781861061b8d85761b8cb8b8b61c6be565b9650505050505050619339565b8561b8e28161d6cb565b96505083865116905061b8b3565b859650505050505050619339565b508383205f905b61b90f868961bf54565b821161b94c5785832080820361b92b5783945050505050619339565b61b93660018561c6be565b935050818061b9449061d6cb565b92505061b905565b505b61b958878761c6be565b979650505050505050565b604080518082019091525f80825260208201525f61b991855f01518660200151865f0151876020015161b847565b60208087018051918601919091525190915061b9ad908261bf54565b83528451602086015161b9c0919061c6be565b810361b9ce575f855261ba00565b8351835161b9dc919061c6be565b8551869061b9eb90839061bf54565b905250835161b9fa908261c6be565b60208601525b50909392505050565b6020811061ba41578151835261ba2060208461c6be565b925061ba2d60208361c6be565b915061ba3a60208261bf54565b905061ba09565b5f19811561ba6f57600161ba5683602061bf54565b61ba629061010061d670565b61ba6c919061bf54565b90505b9151835183169219169190911790915250565b60605f61ba8f8484617e3f565b805160208083015160405193945061baa99390910161d6f8565b60405160208183030381529060405291505092915050565b610b09806200d73483390190565b61076b806200e23d83390190565b61102e806200e9a883390190565b611fc0806200f9d683390190565b6040518060e001604052806060815260200160608152602001606081526020015f151581526020015f151581526020015f1515815260200161bb3961bb3e565b905290565b6040518061010001604052805f151581526020015f15158152602001606081526020015f801916815260200160608152602001606081526020015f1515815260200161bb3960405180608001604052805f81526020015f81526020015f81526020015f81525090565b602080825282518282018190525f918401906040840190835b8181101561bbe75783516001600160a01b031683526020938401939092019160010161bbc0565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561bd19577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b8801810191908801905f5b8181101561bcff577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261bce984865161bbf2565b602095860195909450929092019160010161bcaf565b50919750505060209485019492909201915060010161bc46565b50929695505050505050565b5f8151808452602084019350602083015f5b8281101561bd775781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161bd37565b5093949350505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561bd19577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261bdeb604088018261bbf2565b905060208201519150868103602088015261be06818361bd25565b96505050602093840193919091019060010161bda7565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561bd19577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845261be7d85835161bbf2565b9450602093840193919091019060010161be43565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561bd19577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261bf11604087018261bd25565b955050602093840193919091019060010161beb8565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156174d6576174d661bf27565b600181811c9082168061bf7b57607f821691505b602082108103619ce6577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f81546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b0360018301541660408401526002820160a060608501525f815461bffa8161bf67565b8060a0880152600182165f811461c018576001811461c0525761c083565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660c089015260c082151560051b890101935061c083565b845f5260205f205f5b8381101561c07a5781548a820160c0015260019091019060200161c05b565b890160c0019450505b50505060038401546080860152809250505092915050565b60c081525f61c0ad60c083018861bbf2565b6001600160a01b0387166020840152828103604084015261c0ce818761bbf2565b85546060850152600186015460ff1615156080850152905082810360a084015261c0f8818561bfb2565b98975050505050505050565b5f6020828403121561c114575f80fd5b81516001600160a01b03811681146175bc575f80fd5b6003811061c15f577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b9052565b610100815260056101008201527f544f4b454e000000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f544b4e00000000000000000000000000000000000000000000000000000000006101608201525f6101808201905060ff8816604083015286606083015261c1f2608083018761c12a565b8460a08301526001600160a01b03841660c083015261b95860e08301846001600160a01b03169052565b610100815260066101008201527f534f4c414e410000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f534f4c00000000000000000000000000000000000000000000000000000000006101608201525f6101808201905060ff8816604083015286606083015261c1f2608083018761c12a565b5f6020828403121561c2bb575f80fd5b815180151581146175bc575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b601f82111561364057805f5260205f20601f840160051c8101602085101561c31c5750805b601f840160051c820191505b81811015610757575f815560010161c328565b815167ffffffffffffffff81111561c3555761c35561c2ca565b61c3698161c363845461bf67565b8461c2f7565b6020601f82116001811461c39b575f831561c3845750848201515b5f19600385901b1c1916600184901b178455610757565b5f84815260208120601f198516915b8281101561c3ca578785015182556020948501946001909201910161c3aa565b508482101561c3e757868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b60e081525f61c40860e083018961bbf2565b8760208401526001600160a01b0387166040840152828103606084015261c42f818761bbf2565b85546080850152600186015460ff16151560a085015290505b82810360c084015261c45a818561bfb2565b9998505050505050505050565b5f6020828403121561c477575f80fd5b5051919050565b60e081525f61c49060e083018961bbf2565b8760208401526001600160a01b0387166040840152828103606084015261c4b7818761bbf2565b855160808501526020860151151560a0850152905061c448565b61012081525f61c4e561012083018b61bbf2565b6001600160a01b038a16602084015288604084015287606084015286608084015282810360a084015261c518818761bbf2565b855460c0850152600186015460ff16151560e085015290505b8281036101008401526178b4818561bfb2565b60e081525f61c55660e083018961bbf2565b876020840152866040840152828103606084015261c42f818761bbf2565b61012081525f61c58861012083018b61bbf2565b6001600160a01b038a16602084015288604084015287606084015286608084015282810360a084015261c5bb818761bbf2565b855160c08501526020860151151560e0850152905061c531565b61012081525f61c5e961012083018a61bbf2565b6001600160a01b03891660208401528760408401528660608401528560808401528281038060a08501525f825261c62f60c0850187805182526020908101511515910152565b602081016101008501525061c647602082018561bfb2565b9a9950505050505050505050565b608081525f61c667608083018761bbf2565b8560208401526001600160a01b0385166040840152828103606084015261b958818561bfb2565b608081525f61c6a0608083018761bbf2565b856020840152846040840152828103606084015261b958818561bfb2565b808201808211156174d6576174d661bf27565b5f8261c704577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b60a081525f61c71b60a083018861bbf2565b8660208401526001600160a01b0386166040840152846060840152828103608084015261c0f8818561bfb2565b60a081525f61c75a60a083018761bbf2565b828103602084015261c76c818761bbf2565b85546040850152600186015460ff16151560608501529050828103608084015261b958818561bfb2565b6001600160a01b0383168152604060208201525f619339604083018461bbf2565b5f81518060208401855e5f93019283525090919050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f61c7ff601a83018561c7b7565b7f3a2000000000000000000000000000000000000000000000000000000000000081526175b8600282018561c7b7565b602081525f6175bc602083018461bbf2565b6040516060810167ffffffffffffffff8111828210171561c8645761c86461c2ca565b60405290565b5f8067ffffffffffffffff84111561c8845761c88461c2ca565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561c8b35761c8b361c2ca565b60405283815290508082840185101561c8ca575f80fd5b8383602083015e5f60208583010152509392505050565b5f82601f83011261c8f0575f80fd5b6175bc8383516020850161c86a565b5f6020828403121561c90f575f80fd5b815167ffffffffffffffff81111561c925575f80fd5b6174d28482850161c8e1565b5f61933961c93f838661c7b7565b8461c7b7565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f61c976601a83018561c7b7565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000815261c9a6601982018561c7b7565b7f2200000000000000000000000000000000000000000000000000000000000000815260010195945050505050565b6001600160a01b03841681526001600160a01b0383166020820152606060408201525f619ccc606083018461bbf2565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201525f6175bc608083018461bbf2565b5f6020828403121561ca5b575f80fd5b815167ffffffffffffffff81111561ca71575f80fd5b8201601f8101841361ca81575f80fd5b6174d28482516020840161c86a565b5f61ca9b828761c7b7565b7f2f00000000000000000000000000000000000000000000000000000000000000815261cacb600182018761c7b7565b90507f2f00000000000000000000000000000000000000000000000000000000000000815261cafd600182018661c7b7565b90507f2f00000000000000000000000000000000000000000000000000000000000000815261cb2f600182018561c7b7565b7f2e6a736f6e0000000000000000000000000000000000000000000000000000008152600501979650505050505050565b604081525f61cb72604083018461bbf2565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081525f61cbe2601f83018461c7b7565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f81527f6d6c00000000000000000000000000000000000000000000000000000000000060208201526022019392505050565b604081525f61cc47604083018461bbf2565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b604081525f61cc98604083018461bbf2565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081525f61cd08601483018461c7b7565b7f275d2e6b656363616b32353600000000000000000000000000000000000000008152600c019392505050565b604081525f61cd47604083018561bbf2565b82810360208401526175b8818561bbf2565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f220000000000000000000000000000000000000000000000000000000000000081525f61cdb7600183018461c7b7565b7f220000000000000000000000000000000000000000000000000000000000000081526001019392505050565b5f61cdef828461c7b7565b7f2f6275696c642d696e666f0000000000000000000000000000000000000000008152600b019392505050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201525f6175bc604b83018461c7b7565b5f60ff821660ff810361ceae5761ceae61bf27565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201525f6175bc602983018461c7b7565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201525f6175bc608083018461bbf2565b5f6020828403121561cf64575f80fd5b815167ffffffffffffffff81111561cf7a575f80fd5b82016060818503121561cf8b575f80fd5b61cf9361c841565b81518060030b811461cfa3575f80fd5b8152602082015167ffffffffffffffff81111561cfbe575f80fd5b61cfca8682850161c8e1565b602083015250604082015167ffffffffffffffff81111561cfe9575f80fd5b61cff58682850161c8e1565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201525f61d05a602183018461c7b7565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657481527f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960208201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560408201527f616c69666965642070617468206f66207468652062617368206578656375746160608201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960808201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60a08201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960c08201527f6f75722070726f6a65637420287573696e6720666f727761726420736c61736860e08201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101008201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061012082015261013b019392505050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201525f61d237602183018561c7b7565b7f2720696e206f75747075743a200000000000000000000000000000000000000081526175b8600d82018561c7b7565b5f6175bc828461c7b7565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201525f6175bc602983018461c7b7565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201525f6175bc602283018461c7b7565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081525f61d351600e83018461c7b7565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e7472616381527f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e7360208201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e747260408201527f6163742e6a736f6e00000000000000000000000000000000000000000000000060608201526068019392505050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081525f61d421601883018561c7b7565b7f20696e2000000000000000000000000000000000000000000000000000000000815261d451600482018561c7b7565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f727465642081527f6c6963656e736520666f7220626c6f636b206578706c6f72657220766572696660208201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f60408201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c60608201527f206f7220736574207468652060736b69704c6963656e73655479706560206f7060808201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060a082015260b70195945050505050565b5f61d549828461c7b7565b7f200000000000000000000000000000000000000000000000000000000000000081526001019392505050565b80820281158282048414176174d6576174d661bf27565b6001815b600184111561d5c85780850481111561d5ac5761d5ac61bf27565b600184161561d5ba57908102905b60019390931c92800261d591565b935093915050565b5f8261d5de575060016174d6565b8161d5ea57505f6174d6565b816001811461d600576002811461d60a5761d626565b60019150506174d6565b60ff84111561d61b5761d61b61bf27565b50506001821b6174d6565b5060208310610133831016604e8410600b841016171561d649575081810a6174d6565b61d6555f19848461d58d565b805f190482111561d6685761d66861bf27565b029392505050565b5f6175bc838361d5d0565b8181035f83128015838313168383128216171561adb75761adb761bf27565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081525f6175bc601c83018461c7b7565b5f5f19820361d6dc5761d6dc61bf27565b5060010190565b5f8161d6f15761d6f161bf27565b505f190190565b5f61d703828561c7b7565b7f3a0000000000000000000000000000000000000000000000000000000000000081526175b8600182018561c7b756fe60c0604052600d60809081526c2bb930b83832b21022ba3432b960991b60a0525f9061002b908261010b565b506040805180820190915260048152630ae8aa8960e31b6020820152600190610054908261010b565b506002805460ff1916601217905534801561006d575f80fd5b506101c5565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061009b57607f821691505b6020821081036100b957634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561010657805f5260205f20601f840160051c810160208510156100e45750805b601f840160051c820191505b81811015610103575f81556001016100f0565b50505b505050565b81516001600160401b0381111561012457610124610073565b610138816101328454610087565b846100bf565b6020601f82116001811461016a575f83156101535750848201515b5f19600385901b1c1916600184901b178455610103565b5f84815260208120601f198516915b828110156101995787850151825560209485019460019092019101610179565b50848210156101b657868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610937806101d25f395ff3fe6080604052600436106100bb575f3560e01c8063313ce56711610071578063a9059cbb1161004c578063a9059cbb146101eb578063d0e30db01461020a578063dd62ed3e14610212575f80fd5b8063313ce5671461018157806370a08231146101ac57806395d89b41146101d7575f80fd5b806318160ddd116100a157806318160ddd1461012757806323b872dd146101435780632e1a7d4d14610162575f80fd5b806306fdde03146100ce578063095ea7b3146100f8575f80fd5b366100ca576100c8610248565b005b5f80fd5b3480156100d9575f80fd5b506100e26102a2565b6040516100ef919061071f565b60405180910390f35b348015610103575f80fd5b5061011761011236600461079a565b61032d565b60405190151581526020016100ef565b348015610132575f80fd5b50475b6040519081526020016100ef565b34801561014e575f80fd5b5061011761015d3660046107c2565b6103a6565b34801561016d575f80fd5b506100c861017c3660046107fc565b610628565b34801561018c575f80fd5b5060025461019a9060ff1681565b60405160ff90911681526020016100ef565b3480156101b7575f80fd5b506101356101c6366004610813565b60036020525f908152604090205481565b3480156101e2575f80fd5b506100e26106ff565b3480156101f6575f80fd5b5061011761020536600461079a565b61070c565b6100c8610248565b34801561021d575f80fd5b5061013561022c36600461082c565b600460209081525f928352604080842090915290825290205481565b335f908152600360205260408120805434929061026690849061088a565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b5f80546102ae9061089d565b80601f01602080910402602001604051908101604052809291908181526020018280546102da9061089d565b80156103255780601f106102fc57610100808354040283529160200191610325565b820191905f5260205f20905b81548152906001019060200180831161030857829003601f168201915b505050505081565b335f81815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103949086815260200190565b60405180910390a35060015b92915050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260036020526040812054821115610412576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201525f60248201526044015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff84163314801590610487575073ffffffffffffffffffffffffffffffffffffffff84165f9081526004602090815260408083203384529091529020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14155b156105435773ffffffffffffffffffffffffffffffffffffffff84165f9081526004602090815260408083203384529091529020548211156104fe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201525f6024820152604401610409565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526004602090815260408083203384529091528120805484929061053d9084906108ee565b90915550505b73ffffffffffffffffffffffffffffffffffffffff84165f90815260036020526040812080548492906105779084906108ee565b909155505073ffffffffffffffffffffffffffffffffffffffff83165f90815260036020526040812080548492906105b090849061088a565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161061691815260200190565b60405180910390a35060019392505050565b335f90815260036020526040902054811115610679576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201525f6024820152604401610409565b335f90815260036020526040812080548392906106979084906108ee565b9091555050604051339082156108fc029083905f818181858888f193505050501580156106c6573d5f803e3d5ffd5b5060405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a250565b600180546102ae9061089d565b5f6107183384846103a6565b9392505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610795575f80fd5b919050565b5f80604083850312156107ab575f80fd5b6107b483610772565b946020939093013593505050565b5f805f606084860312156107d4575f80fd5b6107dd84610772565b92506107eb60208501610772565b929592945050506040919091013590565b5f6020828403121561080c575f80fd5b5035919050565b5f60208284031215610823575f80fd5b61071882610772565b5f806040838503121561083d575f80fd5b61084683610772565b915061085460208401610772565b90509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156103a0576103a061085d565b600181811c908216806108b157607f821691505b6020821081036108e8577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b818103818111156103a0576103a061085d56fea2646970667358221220f82a6621bc6ae2f40b7ff1dde0e016bd4b523e6e1df6d8c4401566e37294755f64736f6c634300081a00336080604052348015600e575f80fd5b5061074f8061001c5f395ff3fe608060405260043610610034575f3560e01c80632d4cfb7e1461003d5780635bcfd6161461005c578063c9028a361461007b57005b3661003b57005b005b348015610048575f80fd5b5061003b61005736600461017e565b61009a565b348015610067575f80fd5b5061003b6100763660046101e4565b6100d4565b348015610086575f80fd5b5061003b610095366004610297565b61014f565b7f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7816040516100c9919061037c565b60405180910390a150565b606081156100eb576100e88284018461047c565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e610116878061056c565b61012660408a0160208b016105cd565b8960400135338660405161013f969594939291906105e6565b60405180910390a1505050505050565b7fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4816040516100c99190610690565b5f6020828403121561018e575f80fd5b813567ffffffffffffffff8111156101a4575f80fd5b820160c081850312156101b5575f80fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146101df575f80fd5b919050565b5f805f805f608086880312156101f8575f80fd5b853567ffffffffffffffff81111561020e575f80fd5b86016060818903121561021f575f80fd5b945061022d602087016101bc565b935060408601359250606086013567ffffffffffffffff81111561024f575f80fd5b8601601f8101881361025f575f80fd5b803567ffffffffffffffff811115610275575f80fd5b886020828401011115610286575f80fd5b959894975092955050506020019190565b5f602082840312156102a7575f80fd5b813567ffffffffffffffff8111156102bd575f80fd5b8201608081850312156101b5575f80fd5b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610301575f80fd5b830160208101925035905067ffffffffffffffff811115610320575f80fd5b80360382131561032e575f80fd5b9250929050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b602081525f61038b83846102ce565b60c060208501526103a060e085018284610335565b91505073ffffffffffffffffffffffffffffffffffffffff6103c4602086016101bc565b1660408401525f604085013590508060608501525060608401358015158082146103ec575f80fd5b80608086015250505f608085013590508060a08501525061041060a08501856102ce565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c0860152610445838284610335565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f6020828403121561048c575f80fd5b813567ffffffffffffffff8111156104a2575f80fd5b8201601f810184136104b2575f80fd5b803567ffffffffffffffff8111156104cc576104cc61044f565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156105385761053861044f565b60405281815282820160200186101561054f575f80fd5b816020840160208301375f91810160200191909152949350505050565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261059f575f80fd5b83018035915067ffffffffffffffff8211156105b9575f80fd5b60200191503681900382131561032e575f80fd5b5f602082840312156105dd575f80fd5b6101b5826101bc565b60a081525f6105f960a08301888a610335565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff85166060840152828103608084015283518082528060208601602084015e5f6020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff6106b2836101bc565b16602082015273ffffffffffffffffffffffffffffffffffffffff6106d9602084016101bc565b1660408201525f8060408401359050806060840152506106fc60608401846102ce565b60808085015261071060a085018284610335565b9594505050505056fea264697066735822122067048c1407e27adba772d60723c367ebba0b1435651626fe4a4cb8854fd2848164736f6c634300081a003360c060405234801561000f575f80fd5b5060405161102e38038061102e83398101604081905261002e916100d8565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006257604051632b2add3d60e01b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0385811691909117909155828116608052811660a0526040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac5905f90a1505050610118565b80516001600160a01b03811681146100d3575f80fd5b919050565b5f805f606084860312156100ea575f80fd5b6100f3846100bd565b9250610101602085016100bd565b915061010f604085016100bd565b90509250925092565b60805160a051610eee6101405f395f6101dd01525f81816102b001526104510152610eee5ff3fe608060405234801561000f575f80fd5b50600436106100f0575f3560e01c806397770dff11610093578063c63585cc11610063578063c63585cc1461026b578063d7fd7afb1461027e578063d936a012146102ab578063ee2815ba146102d2575f80fd5b806397770dff14610212578063a7cb050714610225578063c39aca3714610238578063c62178ac1461024b575f80fd5b8063513a9c05116100ce578063513a9c0514610183578063569541b9146101b8578063842da36d146101d857806391dd645f146101ff575f80fd5b80630be15547146100f45780631f0e251b146101535780633ce4a5bc14610168575b5f80fd5b610129610102366004610bb9565b60016020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b610166610161366004610bf8565b6102e5565b005b61012973735b14bb79463307aacbed86daf3322b1e6226ab81565b610129610191366004610bb9565b60026020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101299073ffffffffffffffffffffffffffffffffffffffff1681565b6101297f000000000000000000000000000000000000000000000000000000000000000081565b61016661020d366004610c18565b6103f9565b610166610220366004610bf8565b61051b565b610166610233366004610c42565b610628565b610166610246366004610c62565b6106c2565b6004546101299073ffffffffffffffffffffffffffffffffffffffff1681565b610129610279366004610d28565b6108b9565b61029d61028c366004610bb9565b5f6020819052908152604090205481565b60405190815260200161014a565b6101297f000000000000000000000000000000000000000000000000000000000000000081565b6101666102e0366004610c18565b6109ec565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610332576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811661037f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610446576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003545f9061048d907f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff16846108b9565b5f8481526002602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251878152918201529192507f0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e910160405180910390a1505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610568576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166105b5576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e906020016103ee565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610675576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461070f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831673735b14bb79463307aacbed86daf3322b1e6226ab148061075c575073ffffffffffffffffffffffffffffffffffffffff831630145b15610793576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690528616906347e7ef24906044016020604051808303815f875af1158015610805573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108299190610d68565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063de43156e906108849089908990899088908890600401610dce565b5f604051808303815f87803b15801561089b575f80fd5b505af11580156108ad573d5f803e3d5ffd5b50505050505050505050565b5f805f6108c68585610abc565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b16603482015291935091508690604801604051602081830303815290604052805190602001206040516020016109ac9291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a39576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d91016106b6565b5f808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610b23576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610b5d578284610b60565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610bb2576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b5f60208284031215610bc9575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610bf3575f80fd5b919050565b5f60208284031215610c08575f80fd5b610c1182610bd0565b9392505050565b5f8060408385031215610c29575f80fd5b82359150610c3960208401610bd0565b90509250929050565b5f8060408385031215610c53575f80fd5b50508035926020909101359150565b5f805f805f8060a08789031215610c77575f80fd5b863567ffffffffffffffff811115610c8d575f80fd5b87016060818a031215610c9e575f80fd5b9550610cac60208801610bd0565b945060408701359350610cc160608801610bd0565b9250608087013567ffffffffffffffff811115610cdc575f80fd5b8701601f81018913610cec575f80fd5b803567ffffffffffffffff811115610d02575f80fd5b896020828401011115610d13575f80fd5b60208201935080925050509295509295509295565b5f805f60608486031215610d3a575f80fd5b610d4384610bd0565b9250610d5160208501610bd0565b9150610d5f60408501610bd0565b90509250925092565b5f60208284031215610d78575f80fd5b81518015158114610c11575f80fd5b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b608081525f86357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018112610e04575f80fd5b870160208101903567ffffffffffffffff811115610e20575f80fd5b803603821315610e2e575f80fd5b60606080850152610e4360e085018284610d87565b91505073ffffffffffffffffffffffffffffffffffffffff610e6760208a01610bd0565b1660a0840152604088013560c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401528281036060840152610eac818587610d87565b9897505050505050505056fea264697066735822122038bc8b41b985a9d87c55de5819bed228bedbe526ecdb62ef049429c752d4729164736f6c634300081a003360c060405234801561000f575f80fd5b50604051611fc0380380611fc083398101604081905261002e916101d0565b6001600160a01b038216158061004b57506001600160a01b038116155b156100695760405163d92e233d60e01b815260040160405180910390fd5b60066100758982610315565b5060076100828882610315565b506008805460ff191660ff881617905560808590528360028111156100a9576100a96103cf565b60a08160028111156100bd576100bd6103cf565b9052506001929092555f80546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506103e39350505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261012d575f80fd5b81516001600160401b038111156101465761014661010a565b604051601f8201601f19908116603f011681016001600160401b03811182821017156101745761017461010a565b60405281815283820160200185101561018b575f80fd5b8160208501602083015e5f918101602001919091529392505050565b8051600381106101b5575f80fd5b919050565b80516001600160a01b03811681146101b5575f80fd5b5f805f805f805f80610100898b0312156101e8575f80fd5b88516001600160401b038111156101fd575f80fd5b6102098b828c0161011e565b60208b015190995090506001600160401b03811115610226575f80fd5b6102328b828c0161011e565b975050604089015160ff81168114610248575f80fd5b60608a0151909650945061025e60808a016101a7565b60a08a0151909450925061027460c08a016101ba565b915061028260e08a016101ba565b90509295985092959890939650565b600181811c908216806102a557607f821691505b6020821081036102c357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561031057805f5260205f20601f840160051c810160208510156102ee5750805b601f840160051c820191505b8181101561030d575f81556001016102fa565b50505b505050565b81516001600160401b0381111561032e5761032e61010a565b6103428161033c8454610291565b846102c9565b6020601f821160018114610374575f831561035d5750848201515b5f19600385901b1c1916600184901b17845561030d565b5f84815260208120601f198516915b828110156103a35787850151825560209485019460019092019101610383565b50848210156103c057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52602160045260245ffd5b60805160a051611ba06104205f395f61033901525f81816102e501528181610bbf01528181610cc201528181610ed90152610fdc0152611ba05ff3fe608060405234801561000f575f80fd5b50600436106101b0575f3560e01c806395d89b41116100f3578063ccc7759911610093578063eddeb1231161006e578063eddeb12314610455578063f2441b3214610468578063f687d12a14610487578063fc5fecd51461049a575f80fd5b8063ccc77599146103c9578063d9eeebed146103dc578063dd62ed3e14610410575f80fd5b8063b84c8246116100ce578063b84c82461461037b578063c47f002714610390578063c7012626146103a3578063c835d7cc146103b6575f80fd5b806395d89b411461032c578063a3413d0314610334578063a9059cbb14610368575f80fd5b80633ce4a5bc1161015e5780634d8943bb116101395780634d8943bb146102a257806370a08231146102ab57806385e1f4d0146102e05780638b851b9514610307575f80fd5b80633ce4a5bc1461023c57806342966c681461027c57806347e7ef241461028f575f80fd5b806318160ddd1161018e57806318160ddd1461020c57806323b872dd14610214578063313ce56714610227575f80fd5b806306fdde03146101b4578063091d2788146101d2578063095ea7b3146101e9575b5f80fd5b6101bc6104ad565b6040516101c991906115fb565b60405180910390f35b6101db60015481565b6040519081526020016101c9565b6101fc6101f7366004611638565b61053d565b60405190151581526020016101c9565b6005546101db565b6101fc610222366004611662565b610553565b60085460405160ff90911681526020016101c9565b61025773735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c9565b6101fc61028a3660046116a0565b6105e8565b6101fc61029d366004611638565b6105fb565b6101db60025481565b6101db6102b93660046116b7565b73ffffffffffffffffffffffffffffffffffffffff165f9081526003602052604090205490565b6101db7f000000000000000000000000000000000000000000000000000000000000000081565b60085461025790610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101bc610752565b61035b7f000000000000000000000000000000000000000000000000000000000000000081565b6040516101c991906116d2565b6101fc610376366004611638565b610761565b61038e6103893660046117d3565b61076d565b005b61038e61039e3660046117d3565b6107ca565b6101fc6103b1366004611820565b610823565b61038e6103c43660046116b7565b61096d565b61038e6103d73660046116b7565b610a80565b6103e4610b94565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101c9565b6101db61041e366004611875565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260046020908152604080832093909416825291909152205490565b61038e6104633660046116a0565b610daa565b5f546102579073ffffffffffffffffffffffffffffffffffffffff1681565b61038e6104953660046116a0565b610e2c565b6103e46104a83660046116a0565b610eae565b6060600680546104bc906118ac565b80601f01602080910402602001604051908101604052809291908181526020018280546104e8906118ac565b80156105335780601f1061050a57610100808354040283529160200191610533565b820191905f5260205f20905b81548152906001019060200180831161051657829003601f168201915b5050505050905090565b5f6105493384846110c2565b5060015b92915050565b5f61055f8484846111ca565b73ffffffffffffffffffffffffffffffffffffffff84165f908152600460209081526040808320338452909152902054828110156105c9576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105dd85336105d8868561192a565b6110c2565b506001949350505050565b5f6105f33383611383565b506001919050565b5f3373735b14bb79463307aacbed86daf3322b1e6226ab1480159061063757505f5473ffffffffffffffffffffffffffffffffffffffff163314155b80156106605750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b15610697576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106a183836114c2565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261074191869061193d565b60405180910390a250600192915050565b6060600780546104bc906118ac565b5f6105493384846111ca565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107ba576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076107c682826119aa565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610817576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60066107c682826119aa565b5f805f61082e610b94565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303815f875af11580156108bd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108e19190611ac1565b610917576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109213385611383565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161095a91899189918791611ae0565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109ba576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a07576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610acd576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b1a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610a75565b5f80546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610c24573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c489190611b0e565b905073ffffffffffffffffffffffffffffffffffffffff8116610c97576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610d23573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d479190611b29565b9050805f03610d82576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60025460015483610d949190611b40565b610d9e9190611b57565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610df7576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610a75565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e79576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610a75565b5f80546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610f3e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f629190611b0e565b905073ffffffffffffffffffffffffffffffffffffffff8116610fb1576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa15801561103d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110619190611b29565b9050805f0361109c576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002545f906110ab8784611b40565b6110b59190611b57565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff831661110f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821661115c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8381165f8181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611217576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611264576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f90815260036020526040902054818110156112c3576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112cd828261192a565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260036020526040808220939093559085168152908120805484929061130f908490611b57565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161137591815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166113d0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f908152600360205260409020548181101561142f576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611439828261192a565b73ffffffffffffffffffffffffffffffffffffffff84165f908152600360205260408120919091556005805484929061147390849061192a565b90915550506040518281525f9073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016111bd565b73ffffffffffffffffffffffffffffffffffffffff821661150f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060055f8282546115209190611b57565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f9081526003602052604081208054839290611559908490611b57565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316905f907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61160d60208301846115af565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff81168114611635575f80fd5b50565b5f8060408385031215611649575f80fd5b823561165481611614565b946020939093013593505050565b5f805f60608486031215611674575f80fd5b833561167f81611614565b9250602084013561168f81611614565b929592945050506040919091013590565b5f602082840312156116b0575f80fd5b5035919050565b5f602082840312156116c7575f80fd5b813561160d81611614565b602081016003831061170b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b91905290565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8067ffffffffffffffff84111561175857611758611711565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156117a5576117a5611711565b6040528381529050808284018510156117bc575f80fd5b838360208301375f60208583010152509392505050565b5f602082840312156117e3575f80fd5b813567ffffffffffffffff8111156117f9575f80fd5b8201601f81018413611809575f80fd5b6118188482356020840161173e565b949350505050565b5f8060408385031215611831575f80fd5b823567ffffffffffffffff811115611847575f80fd5b8301601f81018513611857575f80fd5b6118668582356020840161173e565b95602094909401359450505050565b5f8060408385031215611886575f80fd5b823561189181611614565b915060208301356118a181611614565b809150509250929050565b600181811c908216806118c057607f821691505b6020821081036118f7577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561054d5761054d6118fd565b604081525f61194f60408301856115af565b90508260208301529392505050565b601f8211156119a557805f5260205f20601f840160051c810160208510156119835750805b601f840160051c820191505b818110156119a2575f815560010161198f565b50505b505050565b815167ffffffffffffffff8111156119c4576119c4611711565b6119d8816119d284546118ac565b8461195e565b6020601f821160018114611a29575f83156119f35750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556119a2565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611a765787850151825560209485019460019092019101611a56565b5084821015611ab257868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b5f60208284031215611ad1575f80fd5b8151801515811461160d575f80fd5b608081525f611af260808301876115af565b6020830195909552506040810192909252606090910152919050565b5f60208284031215611b1e575f80fd5b815161160d81611614565b5f60208284031215611b39575f80fd5b5051919050565b808202811582820484141761054d5761054d6118fd565b8082018082111561054d5761054d6118fd56fea264697066735822122065c251f5f5642d843f788170d157a28502c6e5d29f4d596c3543cb379ea6b55e64736f6c634300081a0033a26469706673582212207a29470f08e222e1cb5d1f644c377b29e7737d2f2bc9c1b03e1ce39b631fb36764736f6c634300081a0033",
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

// TestWithdrawAndCallZRC20SolanaFails is a paid mutator transaction binding the contract method 0x77def36e.
//
// Solidity: function testWithdrawAndCallZRC20SolanaFails() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawAndCallZRC20SolanaFails(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallZRC20SolanaFails")
}

// TestWithdrawAndCallZRC20SolanaFails is a paid mutator transaction binding the contract method 0x77def36e.
//
// Solidity: function testWithdrawAndCallZRC20SolanaFails() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawAndCallZRC20SolanaFails() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20SolanaFails(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallZRC20SolanaFails is a paid mutator transaction binding the contract method 0x77def36e.
//
// Solidity: function testWithdrawAndCallZRC20SolanaFails() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawAndCallZRC20SolanaFails() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawAndCallZRC20SolanaFails(&_GatewayZEVMInboundTest.TransactOpts)
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

// TestWithdrawZRC20Solana is a paid mutator transaction binding the contract method 0x73a3db40.
//
// Solidity: function testWithdrawZRC20Solana() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactor) TestWithdrawZRC20Solana(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.contract.Transact(opts, "testWithdrawZRC20Solana")
}

// TestWithdrawZRC20Solana is a paid mutator transaction binding the contract method 0x73a3db40.
//
// Solidity: function testWithdrawZRC20Solana() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestSession) TestWithdrawZRC20Solana() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20Solana(&_GatewayZEVMInboundTest.TransactOpts)
}

// TestWithdrawZRC20Solana is a paid mutator transaction binding the contract method 0x73a3db40.
//
// Solidity: function testWithdrawZRC20Solana() returns()
func (_GatewayZEVMInboundTest *GatewayZEVMInboundTestTransactorSession) TestWithdrawZRC20Solana() (*types.Transaction, error) {
	return _GatewayZEVMInboundTest.Contract.TestWithdrawZRC20Solana(&_GatewayZEVMInboundTest.TransactOpts)
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
