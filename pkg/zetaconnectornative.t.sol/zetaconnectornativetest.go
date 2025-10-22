// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectornative

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

// ZetaConnectorNativeTestMetaData contains all meta data concerning the ZetaConnectorNativeTest contract.
var ZetaConnectorNativeTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testTSSUpgrade\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfSenderIsNotTSSUpdater\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpgradeAndWithdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveERC20Partial\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveNoParams\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveOnCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiveOnCallTNotAllowedWithArbitraryCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevert\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndRevertFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawFailsIfSenderIsNotWithdrawer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositedAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedZetaConnectorTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x600c8054600160ff199182168117909255601f8054909116909117905560a06040526000608052602c80546001600160a01b0319169055348015604157600080fd5b5061d708806100516000396000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c8063a217fddf1161010f578063ccb0e3f2116100a2578063e20c9f7111610071578063e20c9f711461037d578063e63ab1e914610385578063fa7626d4146103ac578063fe574f84146103b957600080fd5b8063ccb0e3f21461035d578063d509b16c14610365578063dcf7d0371461036d578063de1cb76c1461037557600080fd5b8063b0a64d03116100de578063b0a64d031461032d578063b5508aa914610335578063ba414fa61461033d578063c19099721461035557600080fd5b8063a217fddf146102ee578063a783c789146102f6578063af298bb11461031d578063b0464fdc1461032557600080fd5b80634df42da11161018757806385226c811161015657806385226c811461028757806385f438c11461029c578063916a17c6146102d157806395665330146102e657600080fd5b80634df42da11461025a57806352ff59391461026257806366d9a9a01461026a578063828320141461027f57600080fd5b80633cba0107116101c35780633cba01071461023a5780633e5e3c23146102425780633f7286f41461024a578063493465581461025257600080fd5b8063070f2ad0146101f55780630a9254e4146101ff5780631ed7831c146102075780632ade388014610225575b600080fd5b6101fd6103c1565b005b6101fd6105c1565b61020f610e12565b60405161021c919061a13d565b60405180910390f35b61022d610e74565b60405161021c919061a1d9565b6101fd610fb6565b61020f611748565b61020f6117a8565b6101fd611808565b6101fd611e1b565b6101fd611f8b565b6102726127c1565b60405161021c919061a33f565b6101fd612943565b61028f612ba3565b60405161021c919061a3dd565b6102c37f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b60405190815260200161021c565b6102d9612c73565b60405161021c919061a454565b6101fd612d6e565b6102c3600081565b6102c37f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b6101fd612fb7565b6102d96133ea565b6101fd6134e5565b61028f6139bf565b610345613a8f565b604051901515815260200161021c565b6101fd613b63565b6101fd613dd4565b6101fd6148e1565b6101fd61491f565b6101fd614f98565b61020f6155aa565b6102c37f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b601f546103459060ff1681565b6101fd61560a565b6026546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b15801561043357600080fd5b505af1158015610447573d6000803e3d6000fd5b5050602654604080516001600160a01b039092166024830152600060448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250610513919060040161a4eb565b600060405180830381600087803b15801561052d57600080fd5b505af1158015610541573d6000803e3d6000fd5b5050602254602480546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529216935063950837aa9250015b600060405180830381600087803b1580156105a757600080fd5b505af11580156105bb573d6000803e3d6000fd5b50505050565b602480547fffffffffffffffffffffffff000000000000000000000000000000000000000090811630179091556025805482166112341790556026805482166156781790556027805490911661987617905560405161061f9061a06a565b604080825260049082018190527f7a6574610000000000000000000000000000000000000000000000000000000060608301526080602083018190528201527f5a4554410000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f0801580156106a3573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600e81527f4761746577617945564d2e736f6c000000000000000000000000000000000000602082015260265460248054935191861690820152604481019390935292166064820152600091610795916084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b00000000000000000000000000000000000000000000000000000000179052615823565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0384811682029290921792839055604080518082018252601081527f4552433230437573746f64792e736f6c000000000000000000000000000000006020820152602654602480549351949096048516958401959095529383166044830152909116606482015291925061083b91608401610738565b602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691909117909155604080518082018252601781527f5a657461436f6e6e6563746f724e61746976652e736f6c0000000000000000006020820152601f546023546026546024805495516101009094048716908401529085166044830152841660648201529190921660848201529192506109409160a40160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff8c8765e00000000000000000000000000000000000000000000000000000000179052615823565b602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383161790556040519091506109829061a077565b604051809103906000f08015801561099e573d6000803e3d6000fd5b50602080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556026546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b158015610a4a57600080fd5b505af1158015610a5e573d6000803e3d6000fd5b5050602480546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d93506306447d56925001600060405180830381600087803b158015610ad357600080fd5b505af1158015610ae7573d6000803e3d6000fd5b5050601f546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015261010090920416925063ae7a3a6f9150602401600060405180830381600087803b158015610b5257600080fd5b505af1158015610b66573d6000803e3d6000fd5b5050601f546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526101009092041692506310188aef9150602401600060405180830381600087803b158015610bd157600080fd5b505af1158015610be5573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610c4757600080fd5b505af1158015610c5b573d6000803e3d6000fd5b50506023546022546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152624c4b406024820152911692506340c10f199150604401600060405180830381600087803b158015610cca57600080fd5b505af1158015610cde573d6000803e3d6000fd5b50506026546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b158015610d6257600080fd5b505af1158015610d76573d6000803e3d6000fd5b5050604080516080810182526024546001600160a01b039081168252602354811660208084019182526001848601908152855191820190955260008152606084018190528351602880549185167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316178155925160298054919095169116179092559251602a55909350909150602b906105bb908261a5c6565b60606016805480602002602001604051908101604052809291908181526020018280548015610e6a57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e4c575b5050505050905090565b6060601e805480602002602001604051908101604052809291908181526020016000905b82821015610fad57600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015610f96578382906000526020600020018054610f099061a52d565b80601f0160208091040260200160405190810160405280929190818152602001828054610f359061a52d565b8015610f825780601f10610f5757610100808354040283529160200191610f82565b820191906000526020600020905b815481529060010190602001808311610f6557829003601f168201915b505050505081526020019060010190610eea565b505050508152505081526020019060010190610e98565b50505050905090565b602354602554604051620186a0602482018190526001600160a01b039384166044830152929091166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905260235460255491516370a0823160e01b81526001600160a01b0392831660048201529293506000929116906370a0823190602401602060405180830381865afa158015611094573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110b8919061a685565b90506110c5816000615842565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015611116573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061113a919061a685565b601f54604080516001600160a01b036101009093048316602482015260448082018a905282518083039091018152606490910182526020810180517fa9059cbb000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90911617905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba39261121592911690600090869060040161a69e565b600060405180830381600087803b15801561122f57600080fd5b505af1158015611243573d6000803e3d6000fd5b50506020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156112bc57600080fd5b505af11580156112d0573d6000803e3d6000fd5b5050601f54602354602554604080516101009094046001600160a01b039081168552602085018d9052928316908401521660608201527f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609250608001905060405180910390a16022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156113ab57600080fd5b505af11580156113bf573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d9150611404908990889061a6c6565b60405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561146557600080fd5b505af1158015611479573d6000803e3d6000fd5b50506022546020546040517f6fb9a7af0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450636fb9a7af93506114d292602c9216908b908a908c9060040161a6df565b600060405180830381600087803b1580156114ec57600080fd5b505af1158015611500573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a08231906024015b602060405180830381865afa158015611554573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611578919061a685565b90506115848188615842565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa1580156115d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115f9919061a685565b905061160e816116098a8761a758565b615842565b602354601f546020546040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081526101009092046001600160a01b0390811660048401529081166024830152600092169063dd62ed3e90604401602060405180830381865afa158015611686573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116aa919061a685565b90506116b7816000615842565b602354601f546040516370a0823160e01b81526101009091046001600160a01b03908116600483015260009216906370a0823190602401602060405180830381865afa15801561170b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061172f919061a685565b905061173c816000615842565b50505050505050505050565b60606018805480602002602001604051908101604052809291908181526020018280548015610e6a576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610e4c575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015610e6a576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610e4c575050505050905090565b6040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f6ed701690000000000000000000000000000000000000000000000000000000017905260235460255493516370a0823160e01b8152620186a0946000949385936001600160a01b03908116936370a08231936118aa93921691016001600160a01b0391909116815260200190565b602060405180830381865afa1580156118c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118eb919061a685565b90506118f8816000615842565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015611949573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061196d919061a685565b601f54604080516001600160a01b036101009093048316602482015260448082018a905282518083039091018152606490910182526020810180517fa9059cbb000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90911617905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba392611a4892911690600090869060040161a69e565b600060405180830381600087803b158015611a6257600080fd5b505af1158015611a76573d6000803e3d6000fd5b50506020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015611aef57600080fd5b505af1158015611b03573d6000803e3d6000fd5b5050601f546040516101009091046001600160a01b031681527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09250602001905060405180910390a16022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611bc157600080fd5b505af1158015611bd5573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d9150611c1a908990889061a6c6565b60405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611c7b57600080fd5b505af1158015611c8f573d6000803e3d6000fd5b50506022546020546040517f6fb9a7af0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450636fb9a7af9350611ce892602c9216908b908a908c9060040161a6df565b600060405180830381600087803b158015611d0257600080fd5b505af1158015611d16573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015611d69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d8d919061a685565b9050611d9a816000615842565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015611deb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e0f919061a685565b905061160e8185615842565b602480546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d916306447d569101600060405180830381600087803b158015611e8c57600080fd5b505af1158015611ea0573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b158015611f2957600080fd5b505af1158015611f3d573d6000803e3d6000fd5b50506022546040517f950837aa000000000000000000000000000000000000000000000000000000008152600060048201526001600160a01b03909116925063950837aa915060240161058d565b6022546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4600482015261432160248201819052916000916001600160a01b03909116906391d1485490604401602060405180830381865afa15801561201a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061203e919061a76b565b9050612049816158c2565b6022546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b03848116602483015260009216906391d1485490604401602060405180830381865afa1580156120d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120f7919061a76b565b9050612102816158c2565b6022546026546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460048201526001600160a01b03918216602482015260009291909116906391d1485490604401602060405180830381865afa158015612192573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121b6919061a76b565b90506121c18161593c565b6022546026546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b03918216602482015260009291909116906391d1485490604401602060405180830381865afa158015612251573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612275919061a76b565b90506122808161593c565b602480546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d916306447d569101600060405180830381600087803b1580156122f157600080fd5b505af1158015612305573d6000803e3d6000fd5b50506022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561237e57600080fd5b505af1158015612392573d6000803e3d6000fd5b5050602654604080516001600160a01b03928316815291891660208301527f33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e935001905060405180910390a16022546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0387811660048301529091169063950837aa90602401600060405180830381600087803b15801561243e57600080fd5b505af1158015612452573d6000803e3d6000fd5b505050506124d685602260009054906101000a90046001600160a01b03166001600160a01b0316635b1125916040518163ffffffff1660e01b8152600401602060405180830381865afa1580156124ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124d1919061a78d565b61598e565b6022546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460048201526001600160a01b038781166024830152909116906391d1485490604401602060405180830381865afa15801561255f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612583919061a76b565b935061258e8461593c565b6022546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b038781166024830152909116906391d1485490604401602060405180830381865afa158015612617573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061263b919061a76b565b92506126468361593c565b6022546026546040517f91d148540000000000000000000000000000000000000000000000000000000081527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460048201526001600160a01b0391821660248201529116906391d1485490604401602060405180830381865afa1580156126d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126f5919061a76b565b9150612700826158c2565b6022546026546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b0391821660248201529116906391d1485490604401602060405180830381865afa15801561278b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127af919061a76b565b90506127ba816158c2565b5050505050565b6060601b805480602002602001604051908101604052809291908181526020016000905b82821015610fad57838290600052602060002090600202016040518060400160405290816000820180546128189061a52d565b80601f01602080910402602001604051908101604052809291908181526020018280546128449061a52d565b80156128915780601f1061286657610100808354040283529160200191612891565b820191906000526020600020905b81548152906001019060200180831161287457829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561292b57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116128d85790505b505050505081525050815260200190600101906127e5565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a090600090819060250160408051808303601f19018152908290526024805463ca669fa760e01b84526001600160a01b03166004840152909250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa79101600060405180830381600087803b1580156129e057600080fd5b505af11580156129f4573d6000803e3d6000fd5b505060248054604080516001600160a01b03909216928201929092527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250612ae3919060040161a4eb565b600060405180830381600087803b158015612afd57600080fd5b505af1158015612b11573d6000803e3d6000fd5b50506022546020546040517f6f8728ad0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450636f8728ad9350612b6c929091169087908690889060289060040161a893565b600060405180830381600087803b158015612b8657600080fd5b505af1158015612b9a573d6000803e3d6000fd5b50505050505050565b6060601a805480602002602001604051908101604052809291908181526020016000905b82821015610fad578382906000526020600020018054612be69061a52d565b80601f0160208091040260200160405190810160405280929190818152602001828054612c129061a52d565b8015612c5f5780601f10612c3457610100808354040283529160200191612c5f565b820191906000526020600020905b815481529060010190602001808311612c4257829003601f168201915b505050505081526020019060010190612bc7565b6060601d805480602002602001604051908101604052809291908181526020016000905b82821015610fad5760008481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015612d5657602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411612d035790505b50505050508152505081526020019060010190612c97565b604080518082018252600181527f310000000000000000000000000000000000000000000000000000000000000060208201529051620186a09160009161012391908390612dc2908490849060240161a8df565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f676cc05400000000000000000000000000000000000000000000000000000000179052517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fed699775000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015612ea057600080fd5b505af1158015612eb4573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b158015612f1157600080fd5b505af1158015612f25573d6000803e3d6000fd5b50506022546020546040517f6fb9a7af0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450636fb9a7af9350612f7e92602c9216908a9087908b9060040161a6df565b600060405180830381600087803b158015612f9857600080fd5b505af1158015612fac573d6000803e3d6000fd5b505050505050505050565b602280546040805160608101909152828152613001926001600160a01b039092169161d6b160208301396040805160208101909152600081526024546001600160a01b03166159ef565b6022546023546025546040516370a0823160e01b81526001600160a01b03918216600482015292811692620186a09260009283929116906370a0823190602401602060405180830381865afa15801561305e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613082919061a685565b905061308f816000615842565b6025546040516001600160a01b0390911660248201526044810184905260009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260235490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba39161316f916001600160a01b039190911690600090869060040161a69e565b600060405180830381600087803b15801561318957600080fd5b505af115801561319d573d6000803e3d6000fd5b50506040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b0388166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561321257600080fd5b505af1158015613226573d6000803e3d6000fd5b50506025546040518781526001600160a01b0390911692507f3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9915060200160405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156132c557600080fd5b505af11580156132d9573d6000803e3d6000fd5b50506025546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526024810188905260448101879052908816925063106e62909150606401600060405180830381600087803b15801561334b57600080fd5b505af115801561335f573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156133b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133d6919061a685565b90506133e28186615842565b505050505050565b6060601c805480602002602001604051908101604052809291908181526020016000905b82821015610fad5760008481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156134cd57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001906004019060208260030104928301926001038202915080841161347a5790505b5050505050815250508152602001906001019061340e565b604080518082018252600181527f3100000000000000000000000000000000000000000000000000000000000000602082015260235460255492516370a0823160e01b81526001600160a01b039384166004820152620186a0936000936101239390928592909116906370a0823190602401602060405180830381865afa158015613574573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613598919061a685565b90506135a5816000615842565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa1580156135f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061361a919061a685565b6020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561369257600080fd5b505af11580156136a6573d6000803e3d6000fd5b505050507fd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a35250184846040516136db92919061a8df565b60405180910390a16022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561375857600080fd5b505af115801561376c573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d91506137b1908990879061a6c6565b60405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561381257600080fd5b505af1158015613826573d6000803e3d6000fd5b505060225460408051602080820183526001600160a01b038a81168352905492517f6fb9a7af0000000000000000000000000000000000000000000000000000000081529381169550636fb9a7af945061388c93919216908b9089908c9060040161a901565b600060405180830381600087803b1580156138a657600080fd5b505af11580156138ba573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa15801561390d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613931919061a685565b905061393e816000615842565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa15801561398f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139b3919061a685565b905061160e8184615842565b60606019805480602002602001604051908101604052809291908181526020016000905b82821015610fad578382906000526020600020018054613a029061a52d565b80601f0160208091040260200160405190810160405280929190818152602001828054613a2e9061a52d565b8015613a7b5780601f10613a5057610100808354040283529160200191613a7b565b820191906000526020600020905b815481529060010190602001808311613a5e57829003601f168201915b5050505050815260200190600101906139e3565b60085460009060ff1615613aa7575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015613b38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b5c919061a685565b1415905090565b602354602554604051620186a0602482018190526001600160a01b039384166044830152929091166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f357fc5a20000000000000000000000000000000000000000000000000000000017905260248054915163ca669fa760e01b81526001600160a01b039092166004830152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa79101600060405180830381600087803b158015613c4a57600080fd5b505af1158015613c5e573d6000803e3d6000fd5b505060248054604080516001600160a01b03909216928201929092527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250613d4d919060040161a4eb565b600060405180830381600087803b158015613d6757600080fd5b505af1158015613d7b573d6000803e3d6000fd5b50506022546020546040517f6fb9a7af0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450636fb9a7af9350612b6c92602c92169088908790899060040161a6df565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152620186a090600090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613e3557600080fd5b505af1158015613e49573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250613f34919060040161a4eb565b600060405180830381600087803b158015613f4e57600080fd5b505af1158015613f62573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b158015613fb657600080fd5b505af1158015613fca573d6000803e3d6000fd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561402757600080fd5b505af115801561403b573d6000803e3d6000fd5b5050602754604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250614126919060040161a4eb565b600060405180830381600087803b15801561414057600080fd5b505af1158015614154573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156141a857600080fd5b505af11580156141bc573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561421957600080fd5b505af115801561422d573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561428157600080fd5b505af1158015614295573d6000803e3d6000fd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd93c0665000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e09150602401600060405180830381600087803b15801561431e57600080fd5b505af1158015614332573d6000803e3d6000fd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561438f57600080fd5b505af11580156143a3573d6000803e3d6000fd5b50506022546025546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101879052604481018690529116925063106e62909150606401600060405180830381600087803b15801561441757600080fd5b505af115801561442b573d6000803e3d6000fd5b50506024805460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063ca669fa7925001600060405180830381600087803b15801561448757600080fd5b505af115801561449b573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156144ef57600080fd5b505af1158015614503573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a08231906024015b602060405180830381865afa158015614557573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061457b919061a685565b9050614588816000615842565b6025546040516001600160a01b0390911660248201526044810184905260009060640160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260235490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba391614668916001600160a01b039190911690600090869060040161a69e565b600060405180830381600087803b15801561468257600080fd5b505af1158015614696573d6000803e3d6000fd5b50506022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561470f57600080fd5b505af1158015614723573d6000803e3d6000fd5b50506025546040518781526001600160a01b0390911692507f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5915060200160405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156147c257600080fd5b505af11580156147d6573d6000803e3d6000fd5b50506022546025546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101899052604481018890529116925063106e62909150606401600060405180830381600087803b15801561484a57600080fd5b505af115801561485e573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156148b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906148d5919061a685565b90506127ba8186615842565b6023546025546040516370a0823160e01b81526001600160a01b039182166004820152620186a09260009283929116906370a082319060240161453a565b602354602554604051620186a0602482018190526001600160a01b039384166044830152929091166064820152600090819060840160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc51316910000000000000000000000000000000000000000000000000000000017905260235460255491516370a0823160e01b81526001600160a01b0392831660048201529293506000929116906370a0823190602401602060405180830381865afa1580156149fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614a21919061a685565b9050614a2e816000615842565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015614a7f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614aa3919061a685565b601f54604080516001600160a01b036101009093048316602482015260448082018a905282518083039091018152606490910182526020810180517fa9059cbb000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90911617905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba392614b7e92911690600090869060040161a69e565b600060405180830381600087803b158015614b9857600080fd5b505af1158015614bac573d6000803e3d6000fd5b50506020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015614c2557600080fd5b505af1158015614c39573d6000803e3d6000fd5b5050601f547f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60925061010090046001600160a01b03169050614c7c60028961a939565b602354602554604080516001600160a01b03958616815260208101949094529184168383015292909216606082015290519081900360800190a16022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015614d2b57600080fd5b505af1158015614d3f573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d9150614d84908990889061a6c6565b60405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015614de557600080fd5b505af1158015614df9573d6000803e3d6000fd5b50506022546020546040517f6fb9a7af0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450636fb9a7af9350614e5292602c9216908b908a908c9060040161a6df565b600060405180830381600087803b158015614e6c57600080fd5b505af1158015614e80573d6000803e3d6000fd5b50506023546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015614ed3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ef7919061a685565b9050614f088161160960028a61a939565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015614f59573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614f7d919061a685565b905061160e81614f8e60028b61a939565b611609908761a758565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152620186a090600090819060250160408051808303601f19018152908290526023546020546370a0823160e01b84526001600160a01b0390811660048501529193506000929116906370a0823190602401602060405180830381865afa15801561502e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615052919061a685565b905061505f816000615842565b6023546022546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa1580156150b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906150d4919061a685565b601f54604080516001600160a01b036101009093048316602482015260448082018a905282518083039091018152606490910182526020810180517fa9059cbb000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90911617905260235491517ff30c7ba300000000000000000000000000000000000000000000000000000000815293945092737109709ecfa91a80626ff3989d68f67f5b1dd12d9263f30c7ba3926151af92911690600090869060040161a69e565b600060405180830381600087803b1580156151c957600080fd5b505af11580156151dd573d6000803e3d6000fd5b50506020546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561525657600080fd5b505af115801561526a573d6000803e3d6000fd5b505050507f689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b601f60019054906101000a90046001600160a01b031660286040516152b592919061a974565b60405180910390a1601f546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561533657600080fd5b505af115801561534a573d6000803e3d6000fd5b50506023546020546040516001600160a01b039283169450911691507fde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e03590615398908a90899060289061a996565b60405180910390a36022546040516381bad6f360e01b8152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561541557600080fd5b505af1158015615429573d6000803e3d6000fd5b50506020546040516001600160a01b0390911692507f5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff09150615471908990889060289061a996565b60405180910390a260265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156154d257600080fd5b505af11580156154e6573d6000803e3d6000fd5b50506022546020546040517f6f8728ad0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450636f8728ad935061554192909116908a9089908b9060289060040161a893565b600060405180830381600087803b15801561555b57600080fd5b505af115801561556f573d6000803e3d6000fd5b50506023546020546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401611537565b60606015805480602002602001604051908101604052809291908181526020018280548015610e6a576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610e4c575050505050905090565b6024805460405163ca669fa760e01b81526001600160a01b039091166004820152620186a091600091737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa79101600060405180830381600087803b15801561566a57600080fd5b505af115801561567e573d6000803e3d6000fd5b505060248054604080516001600160a01b03909216928201929092527f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e460448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe2517d3f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061576d919060040161a4eb565b600060405180830381600087803b15801561578757600080fd5b505af115801561579b573d6000803e3d6000fd5b50506022546025546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260248101879052604481018690529116925063106e62909150606401600060405180830381600087803b15801561580f57600080fd5b505af11580156133e2573d6000803e3d6000fd5b600061582d61a084565b615838848483615a04565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c54906044015b60006040518083038186803b1580156158ae57600080fd5b505afa1580156133e2573d6000803e3d6000fd5b6040517fa59828850000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063a5982885906024015b60006040518083038186803b15801561592857600080fd5b505afa1580156127ba573d6000803e3d6000fd5b6040517f0c9fd5810000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90630c9fd58190602401615910565b6040517f515361f60000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015282166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063515361f690604401615896565b6159f761a084565b6127ba8585858486615a7f565b600080615a118584615b7f565b9050615a746040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f78790000008152508286604051602001615a5f92919061a8df565b60405160208183030381529060405285615b8b565b9150505b9392505050565b6040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201528190737109709ecfa91a80626ff3989d68f67f5b1dd12d9081906306447d5690602401600060405180830381600087803b158015615af157600080fd5b505af1925050508015615b02575060015b615b1757615b1287878787615bb9565b612b9a565b615b2387878787615bb9565b806001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015615b5e57600080fd5b505af1158015615b72573d6000803e3d6000fd5b5050505050505050505050565b6000615a788383615bd2565b60c08101515160009015615baf57615ba884848460c00151615bed565b9050615a78565b615ba88484615d93565b6000615bc58483615e7e565b90506127ba858285615e8a565b6000615bde8383616254565b615a7883836020015184615b8b565b600080615bf8616264565b90506000615c068683616337565b90506000615c1d82606001518360200151856167dd565b90506000615c2d838389896169ef565b90506000615c3a8261786c565b602081015181519192509060030b15615cad57898260400151604051602001615c6492919061a9c1565b60408051601f19818403018152908290527f08c379a0000000000000000000000000000000000000000000000000000000008252615ca49160040161a4eb565b60405180910390fd5b6000615cf06040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a200000000000000000000000815250836001617a3b565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d90615d4390849060040161a4eb565b602060405180830381865afa158015615d60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615d84919061a78d565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc92590615de890879060040161a4eb565b600060405180830381865afa158015615e05573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615e2d919081019061aafb565b90506000615e5b8285604051602001615e4792919061ab30565b604051602081830303815290604052617c3b565b90506001600160a01b038116615838578484604051602001615c6492919061ab5f565b6000615bde8383617c4e565b6040517f667f9d700000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201527fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90600090829063667f9d7090604401602060405180830381865afa158015615f26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615f4a919061a685565b9050806160f1576000615f5c86617c5a565b604080518082018252600581527f352e302e3000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150615fe7905b60408051808201825260008082526020918201528151808301909252845182528085019082015290617d52565b80615ff3575060008451115b15616076576040517f4f1ef2860000000000000000000000000000000000000000000000000000000081526001600160a01b03871690634f1ef2869061603f908890889060040161a8df565b600060405180830381600087803b15801561605957600080fd5b505af115801561606d573d6000803e3d6000fd5b505050506160eb565b6040517f3659cfe60000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152871690633659cfe690602401600060405180830381600087803b1580156160d257600080fd5b505af11580156160e6573d6000803e3d6000fd5b505050505b506127ba565b8060006160fd82617c5a565b604080518082018252600581527f352e302e300000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061615f90615fba565b8061616b575060008551115b156161f0576040517f9623609d0000000000000000000000000000000000000000000000000000000081526001600160a01b03831690639623609d906161b9908a908a908a9060040161ac0a565b600060405180830381600087803b1580156161d357600080fd5b505af11580156161e7573d6000803e3d6000fd5b50505050612b9a565b6040517f99a88ec40000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015287811660248301528316906399a88ec490604401600060405180830381600087803b158015615b5e57600080fd5b61626082826000617d66565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c906162eb90849060040161ac3b565b600060405180830381865afa158015616308573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616330919081019061ac82565b9250505090565b6163696040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d90506163b46040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6163bd85617e69565b602082015260006163cd8661824e565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa15801561640f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616437919081019061ac82565b86838560200151604051602001616451949392919061accb565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb11906164a990859060040161a4eb565b600060405180830381865afa1580156164c6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526164ee919081019061ac82565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f69061653690849060040161adcf565b602060405180830381865afa158015616553573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616577919061a76b565b61658c5781604051602001615c64919061ae21565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906165d190849060040161aeb3565b600060405180830381865afa1580156165ee573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616616919081019061ac82565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f69061665d90849060040161af05565b602060405180830381865afa15801561667a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061669e919061a76b565b15616733576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906166e890849060040161af05565b600060405180830381865afa158015616705573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261672d919081019061ac82565b60408501525b846001600160a01b03166349c4fac8828660000151604051602001616758919061af57565b6040516020818303038152906040526040518363ffffffff1660e01b815260040161678492919061afc3565b600060405180830381865afa1580156167a1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526167c9919081019061ac82565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b60608152602001906001900390816167f95790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250816000815181106168595761685961afe8565b60200260200101819052506040518060400160405280600381526020017f2d726c0000000000000000000000000000000000000000000000000000000000815250816001815181106168ad576168ad61afe8565b6020026020010181905250846040516020016168c9919061b017565b604051602081830303815290604052816002815181106168eb576168eb61afe8565b602002602001018190525082604051602001616907919061b083565b604051602081830303815290604052816003815181106169295761692961afe8565b6020026020010181905250600061693f8261786c565b602080820151604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000081850190815282518084018452600080825290860152825180840190935290518252928101929092529192506169d090604080518082018252600080825260209182015281518083019092528451825280850190820152906184d1565b6169e55785604051602001615c64919061b0c4565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d9015616a3f565b511590565b616bb357826020015115616afb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401615ca4565b8260c0015115616bb3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401615ca4565b6040805160ff8082526120008201909252600091816020015b6060815260200190600190039081616bcc57905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280616c279061b155565b935060ff1681518110616c3c57616c3c61afe8565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e3700000000000000000000000000000000000000815250604051602001616c8d919061b174565b604051602081830303815290604052828280616ca89061b155565b935060ff1681518110616cbd57616cbd61afe8565b60200260200101819052506040518060400160405280600681526020017f6465706c6f790000000000000000000000000000000000000000000000000000815250828280616d0a9061b155565b935060ff1681518110616d1f57616d1f61afe8565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d65000000000000000000000000000000000000815250828280616d6c9061b155565b935060ff1681518110616d8157616d8161afe8565b60200260200101819052508760200151828280616d9d9061b155565b935060ff1681518110616db257616db261afe8565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e747261637450617468000000000000000000000000000000000000815250828280616dff9061b155565b935060ff1681518110616e1457616e1461afe8565b602090810291909101015287518282616e2c8161b155565b935060ff1681518110616e4157616e4161afe8565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e49640000000000000000000000000000000000000000000000815250828280616e8e9061b155565b935060ff1681518110616ea357616ea361afe8565b6020026020010181905250616eb746618532565b8282616ec28161b155565b935060ff1681518110616ed757616ed761afe8565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c650000000000000000000000000000000000815250828280616f249061b155565b935060ff1681518110616f3957616f3961afe8565b602002602001018190525086828280616f519061b155565b935060ff1681518110616f6657616f6661afe8565b602090810291909101015285511561708d5760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f6465000000000000000000000060208201528282616fb78161b155565b935060ff1681518110616fcc57616fcc61afe8565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d9061701c90899060040161a4eb565b600060405180830381865afa158015617039573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617061919081019061ac82565b828261706c8161b155565b935060ff16815181106170815761708161afe8565b60200260200101819052505b84602001511561715d5760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826170d68161b155565b935060ff16815181106170eb576170eb61afe8565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806171389061b155565b935060ff168151811061714d5761714d61afe8565b6020026020010181905250617324565b617195616a3a8660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6172285760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826171d88161b155565b935060ff16815181106171ed576171ed61afe8565b60200260200101819052508460a0015160405160200161720d919061b017565b6040516020818303038152906040528282806171389061b155565b8460c0015115801561726b57506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261726990511590565b155b156173245760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826172af8161b155565b935060ff16815181106172c4576172c461afe8565b60200260200101819052506172d8886185d2565b6040516020016172e8919061b017565b6040516020818303038152906040528282806173039061b155565b935060ff16815181106173185761731861afe8565b60200260200101819052505b6040808601518151808301835260008082526020918201528251808401909352815183529081019082015261735890511590565b6173ed5760408051808201909152600b81527f2d2d72656c6179657249640000000000000000000000000000000000000000006020820152828261739b8161b155565b935060ff16815181106173b0576173b061afe8565b602002602001018190525084604001518282806173cc9061b155565b935060ff16815181106173e1576173e161afe8565b60200260200101819052505b60608501511561750e5760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826174368161b155565b935060ff168151811061744b5761744b61afe8565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa1580156174ba573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526174e2919081019061ac82565b82826174ed8161b155565b935060ff16815181106175025761750261afe8565b60200260200101819052505b60e085015151156175b55760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826175588161b155565b935060ff168151811061756d5761756d61afe8565b60200260200101819052506175898560e0015160000151618532565b82826175948161b155565b935060ff16815181106175a9576175a961afe8565b60200260200101819052505b60e0850151602001511561765f5760408051808201909152600a81527f2d2d676173507269636500000000000000000000000000000000000000000000602082015282826176028161b155565b935060ff16815181106176175761761761afe8565b60200260200101819052506176338560e0015160200151618532565b828261763e8161b155565b935060ff16815181106176535761765361afe8565b60200260200101819052505b60e085015160400151156177095760408051808201909152600e81527f2d2d6d6178466565506572476173000000000000000000000000000000000000602082015282826176ac8161b155565b935060ff16815181106176c1576176c161afe8565b60200260200101819052506176dd8560e0015160400151618532565b82826176e88161b155565b935060ff16815181106176fd576176fd61afe8565b60200260200101819052505b60e085015160600151156177b35760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826177568161b155565b935060ff168151811061776b5761776b61afe8565b60200260200101819052506177878560e0015160600151618532565b82826177928161b155565b935060ff16815181106177a7576177a761afe8565b60200260200101819052505b60008160ff1667ffffffffffffffff8111156177d1576177d161a4fe565b60405190808252806020026020018201604052801561780457816020015b60608152602001906001900390816177ef5790505b50905060005b8260ff168160ff16101561785d57838160ff168151811061782d5761782d61afe8565b6020026020010151828260ff168151811061784a5761784a61afe8565b602090810291909101015260010161780a565b5093505050505b949350505050565b6178936040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c916179199186910161b1df565b600060405180830381865afa158015617936573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261795e919081019061ac82565b9050600061796c86836190c1565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b815260040161799c919061a3dd565b6000604051808303816000875af11580156179bb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526179e3919081019061b226565b805190915060030b158015906179fc5750602081015151155b8015617a0b5750604081015151155b156169e55781600081518110617a2357617a2361afe8565b6020026020010151604051602001615c64919061b2dc565b60606000617a708560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600080825260209182015281518083019092528651825280870190820152909150617aa79082905b90619216565b15617c04576000617b2482617b1e84617b18617aea8a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b9061923d565b9061929f565b604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150617b88908290619216565b15617bf257604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617bef905b8290619324565b90505b617bfb8161934a565b92505050615a78565b8215617c1d578484604051602001615c6492919061b4c8565b5050604080516020810190915260008152615a78565b509392505050565b6000808251602084016000f09392505050565b61626082826001617d66565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fad3cb1cc00000000000000000000000000000000000000000000000000000000179052905160609160009182916001600160a01b03861691617ccf919061b56f565b600060405180830381855afa9150503d8060008114617d0a576040519150601f19603f3d011682016040523d82523d6000602084013e617d0f565b606091505b5091509150818015617d22575060208151115b15617d3b5780806020019051810190617864919061ac82565b505060408051602081019091526000815292915050565b6000617d5e83836193b3565b159392505050565b8160a0015115617d7557505050565b6000617d8284848461948e565b90506000617d8f8261786c565b602081015181519192509060030b158015617e2b5750604080518082018252600781527f535543434553530000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617e2b90604080518082018252600080825260209182015281518083019092528451825280850190820152617aa1565b15617e3857505050505050565b60408201515115617e58578160400151604051602001615c64919061b58b565b80604051602001615c64919061b5e9565b60606000617e9e8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150617f03905b82906184d1565b15617f7257604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615a7890617f6d908390619a29565b61934a565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152617fd4905b8290619ab3565b6001036180a157604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261803a90617be8565b50604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615a7890617f6d905b8390619324565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261810090617efc565b1561823757604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290618168908390619b4d565b90506000816001835161817b919061a758565b8151811061818b5761818b61afe8565b6020026020010151905061822e617f6d6182016040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60408051808201825260008082526020918201528151808301909252855182528086019082015290619a29565b95945050505050565b82604051602001615c64919061b654565b50919050565b606060006182838360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506182e590617efc565b156182f357615a788161934a565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261835290617fcd565b6001036183bc57604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615a7890617f6d9061809a565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261841b90617efc565b1561823757604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290618483908390619b4d565b90506001815111156184bf57806002825161849e919061a758565b815181106184ae576184ae61afe8565b602002602001015192505050919050565b5082604051602001615c64919061b654565b8051825160009111156184e65750600061583c565b815183516020850151600092916184fc9161b732565b618506919061a758565b90508260200151810361851d57600191505061583c565b82516020840151819020912014905092915050565b6060600061853f83619bf2565b600101905060008167ffffffffffffffff81111561855f5761855f61a4fe565b6040519080825280601f01601f191660200182016040528015618589576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461859357509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e534544000000000000000000000000000000000000000000008184019081528551808701875283815284019290925284518086019095525184529083015260609161865e905b8290617d52565b1561869e57505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e73650000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526186fd90618657565b1561873d57505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261879c90618657565b156187dc57505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261883b90618657565b806188a05750604080518082018252601081527f47504c2d322e302d6f722d6c6174657200000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526188a090618657565b156188e057505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261893f90618657565b806189a45750604080518082018252601081527f47504c2d332e302d6f722d6c6174657200000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526189a490618657565b156189e457505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618a4390618657565b80618aa85750604080518082018252601181527f4c47504c2d322e312d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618aa890618657565b15618ae857505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618b4790618657565b80618bac5750604080518082018252601181527f4c47504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618bac90618657565b15618bec57505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c61757365000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618c4b90618657565b15618c8b57505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c61757365000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618cea90618657565b15618d2a57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618d8990618657565b15618dc957505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618e2890618657565b15618e6857505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e300000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618ec790618657565b15618f0757505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618f6690618657565b80618fcb5750604080518082018252601181527f4147504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618fcb90618657565b1561900b57505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261906a90618657565b156190aa57505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151615c64929060200161b745565b60608060005b845181101561914c57818582815181106190e3576190e361afe8565b60200260200101516040516020016190fc92919061ab30565b60405160208183030381529060405291506001855161911b919061a758565b81146191445781604051602001619132919061b8ae565b60405160208183030381529060405291505b6001016190c7565b5060408051600380825260808201909252600091816020015b606081526020019060019003908161916557905050905083816000815181106191905761919061afe8565b60200260200101819052506040518060400160405280600281526020017f2d63000000000000000000000000000000000000000000000000000000000000815250816001815181106191e4576191e461afe8565b602002602001018190525081816002815181106192035761920361afe8565b6020908102919091010152949350505050565b60208083015183518351928401516000936192349291849190619cd4565b14159392505050565b6040805180820190915260008082526020820152600061926f8460000151856020015185600001518660200151619de5565b9050836020015181619281919061a758565b8451859061929090839061a758565b90525060208401525090919050565b60408051808201909152600080825260208201528151835110156192c457508161583c565b60208083015190840151600191146192eb5750815160208481015190840151829020919020145b801561931c5782518451859061930290839061a758565b905250825160208501805161931890839061b732565b9052505b509192915050565b6040805180820190915260008082526020820152619343838383619f05565b5092915050565b60606000826000015167ffffffffffffffff81111561936b5761936b61a4fe565b6040519080825280601f01601f191660200182016040528015619395576020820181803683370190505b50905060006020820190506193438185602001518660000151619fb0565b81518151600091908111156193c6575081515b6020808501519084015160005b8381101561947f578251825180821461944f57600019602087101561942e5760018461940089602061a758565b61940a919061b732565b61941590600861b8ef565b61942090600261b9ed565b61942a919061a758565b1990505b818116838216818103911461944c57975061583c9650505050505050565b50505b61945a60208661b732565b945061946760208561b732565b93505050602081619478919061b732565b90506193d3565b50845186516169e5919061b9f9565b6060600061949a616264565b6040805160ff808252612000820190925291925060009190816020015b60608152602001906001900390816194b757905050905060006040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806195129061b155565b935060ff16815181106195275761952761afe8565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e3300000000000000000000000000000000000000000000000000815250604051602001619578919061ba19565b6040516020818303038152906040528282806195939061b155565b935060ff16815181106195a8576195a861afe8565b60200260200101819052506040518060400160405280600881526020017f76616c69646174650000000000000000000000000000000000000000000000008152508282806195f59061b155565b935060ff168151811061960a5761960a61afe8565b602002602001018190525082604051602001619626919061b083565b6040516020818303038152906040528282806196419061b155565b935060ff16815181106196565761965661afe8565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e7472616374000000000000000000000000000000000000000000008152508282806196a39061b155565b935060ff16815181106196b8576196b861afe8565b60200260200101819052506196cd878461a02a565b82826196d88161b155565b935060ff16815181106196ed576196ed61afe8565b6020908102919091010152855151156197995760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261973f8161b155565b935060ff16815181106197545761975461afe8565b602002602001018190525061976d86600001518461a02a565b82826197788161b155565b935060ff168151811061978d5761978d61afe8565b60200260200101819052505b8560800151156198075760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b0000000000000000602082015282826197e28161b155565b935060ff16815181106197f7576197f761afe8565b602002602001018190525061986d565b841561986d5760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261984c8161b155565b935060ff16815181106198615761986161afe8565b60200260200101819052505b604086015151156199095760408051808201909152600d81527f2d2d756e73616665416c6c6f7700000000000000000000000000000000000000602082015282826198b78161b155565b935060ff16815181106198cc576198cc61afe8565b602002602001018190525085604001518282806198e89061b155565b935060ff16815181106198fd576198fd61afe8565b60200260200101819052505b8560600151156199735760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d6573000000000000000000000000602082015282826199528161b155565b935060ff16815181106199675761996761afe8565b60200260200101819052505b60008160ff1667ffffffffffffffff8111156199915761999161a4fe565b6040519080825280602002602001820160405280156199c457816020015b60608152602001906001900390816199af5790505b50905060005b8260ff168160ff161015619a1d57838160ff16815181106199ed576199ed61afe8565b6020026020010151828260ff1681518110619a0a57619a0a61afe8565b60209081029190910101526001016199ca565b50979650505050505050565b6040805180820190915260008082526020820152815183511015619a4e57508161583c565b81518351602085015160009291619a649161b732565b619a6e919061a758565b60208401519091506001908214619a8f575082516020840151819020908220145b8015619aaa57835185518690619aa690839061a758565b9052505b50929392505050565b6000808260000151619ad78560000151866020015186600001518760200151619de5565b619ae1919061b732565b90505b83516020850151619af5919061b732565b81116193435781619b058161ba5e565b9250508260000151619b3c856020015183619b20919061a758565b8651619b2c919061a758565b8386600001518760200151619de5565b619b46919061b732565b9050619ae4565b60606000619b5b8484619ab3565b619b6690600161b732565b67ffffffffffffffff811115619b7e57619b7e61a4fe565b604051908082528060200260200182016040528015619bb157816020015b6060815260200190600190039081619b9c5790505b50905060005b8151811015617c3357619bcd617f6d8686619324565b828281518110619bdf57619bdf61afe8565b6020908102919091010152600101619bb7565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310619c3b577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310619c67576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310619c8557662386f26fc10000830492506010015b6305f5e1008310619c9d576305f5e100830492506008015b6127108310619cb157612710830492506004015b60648310619cc3576064830492506002015b600a831061583c5760010192915050565b600080858411619ddb5760208411619d875760008415619d1f576001619cfb86602061a758565b619d0690600861b8ef565b619d1190600261b9ed565b619d1b919061a758565b1990505b8351811685619d2e898961b732565b619d38919061a758565b805190935082165b818114619d7257878411619d5a5787945050505050617864565b83619d648161ba78565b945050828451169050619d40565b619d7c878561b732565b945050505050617864565b838320619d94858861a758565b619d9e908761b732565b91505b858210619dd957848220808203619dc657619dbc868461b732565b9350505050617864565b619dd160018461a758565b925050619da1565b505b5092949350505050565b60008381868511619ef05760208511619e9f5760008515619e31576001619e0d87602061a758565b619e1890600861b8ef565b619e2390600261b9ed565b619e2d919061a758565b1990505b84518116600087619e428b8b61b732565b619e4c919061a758565b855190915083165b828114619e9157818610619e7957619e6c8b8b61b732565b9650505050505050617864565b85619e838161ba5e565b965050838651169050619e54565b859650505050505050617864565b508383206000905b619eb1868961a758565b8211619eee57858320808203619ecd5783945050505050617864565b619ed860018561b732565b9350508180619ee69061ba5e565b925050619ea7565b505b619efa878761b732565b979650505050505050565b60408051808201909152600080825260208201526000619f378560000151866020015186600001518760200151619de5565b602080870180519186019190915251909150619f53908261a758565b835284516020860151619f66919061b732565b8103619f755760008552619fa7565b83518351619f83919061b732565b85518690619f9290839061a758565b9052508351619fa1908261b732565b60208601525b50909392505050565b60208110619fe85781518352619fc760208461b732565b9250619fd460208361b732565b9150619fe160208261a758565b9050619fb0565b600019811561a017576001619ffe83602061a758565b61a00a9061010061b9ed565b61a014919061a758565b90505b9151835183169219169190911790915250565b6060600061a0388484616337565b805160208083015160405193945061a0529390910161ba8f565b60405160208183030381529060405291505092915050565b610c9f8061bae883390190565b610f2a8061c78783390190565b6040518060e0016040528060608152602001606081526020016060815260200160001515815260200160001515815260200160001515815260200161a0c761a0cc565b905290565b6040518061010001604052806000151581526020016000151581526020016060815260200160008019168152602001606081526020016060815260200160001515815260200161a0c76040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b8181101561a17e5783516001600160a01b031683526020938401939092019160010161a157565b509095945050505050565b60005b8381101561a1a457818101518382015260200161a18c565b50506000910152565b6000815180845261a1c581602086016020860161a189565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561a2d5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b8181101561a2bb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261a2a584865161a1ad565b602095860195909450929092019160010161a26b565b50919750505060209485019492909201915060010161a201565b50929695505050505050565b600081518084526020840193506020830160005b8281101561a3355781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161a2f5565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561a2d5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261a3ab604088018261a1ad565b905060208201519150868103602088015261a3c6818361a2e1565b96505050602093840193919091019060010161a367565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561a2d5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845261a43f85835161a1ad565b9450602093840193919091019060010161a405565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561a2d5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261a4d5604087018261a2e1565b955050602093840193919091019060010161a47c565b602081526000615a78602083018461a1ad565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c9082168061a54157607f821691505b602082108103618248577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f82111561a5c157806000526020600020601f840160051c8101602085101561a5a15750805b601f840160051c820191505b818110156127ba576000815560010161a5ad565b505050565b815167ffffffffffffffff81111561a5e05761a5e061a4fe565b61a5f48161a5ee845461a52d565b8461a57a565b6020601f82116001811461a628576000831561a6105750848201515b600019600385901b1c1916600184901b1784556127ba565b600084815260208120601f198516915b8281101561a658578785015182556020948501946001909201910161a638565b508482101561a6765786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b60006020828403121561a69757600080fd5b5051919050565b6001600160a01b038416815282602082015260606040820152600061822e606083018461a1ad565b828152604060208201526000617864604083018461a1ad565b6001600160a01b0386541681526001600160a01b038516602082015283604082015260a06060820152600061a71760a083018561a1ad565b90508260808301529695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561583c5761583c61a729565b60006020828403121561a77d57600080fd5b81518015158114615a7857600080fd5b60006020828403121561a79f57600080fd5b81516001600160a01b0381168114615a7857600080fd5b6001600160a01b0381541682526001600160a01b03600182015416602083015260028101546040830152600060038201608060608501526000815461a7fa8161a52d565b806080880152600182166000811461a819576001811461a8535761a887565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660a089015260a082151560051b890101935061a887565b84600052602060002060005b8381101561a87e5781548a820160a0015260019091019060200161a85f565b890160a0019450505b50919695505050505050565b6001600160a01b038616815284602082015260a06040820152600061a8bb60a083018661a1ad565b846060840152828103608084015261a8d3818561a7b6565b98975050505050505050565b6001600160a01b0383168152604060208201526000617864604083018461a1ad565b6001600160a01b0386511681526001600160a01b038516602082015283604082015260a06060820152600061a71760a083018561a1ad565b60008261a96f577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6001600160a01b0383168152604060208201526000617864604083018461a7b6565b83815260606020820152600061a9af606083018561a1ad565b82810360408401526169e5818561a7b6565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161a9f981601a85016020880161a189565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161aa3681601c84016020880161a189565b01601c01949350505050565b6040516060810167ffffffffffffffff8111828210171561aa655761aa6561a4fe565b60405290565b60008067ffffffffffffffff84111561aa865761aa8661a4fe565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561aab55761aab561a4fe565b60405283815290508082840185101561aacd57600080fd5b617c3384602083018561a189565b600082601f83011261aaec57600080fd5b615a788383516020850161aa6b565b60006020828403121561ab0d57600080fd5b815167ffffffffffffffff81111561ab2457600080fd5b6158388482850161aadb565b6000835161ab4281846020880161a189565b83519083019061ab5681836020880161a189565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161ab9781601a85016020880161a189565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161abd481603384016020880161a189565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b6001600160a01b03841681526001600160a01b038316602082015260606040820152600061822e606083018461a1ad565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201526000615a78608083018461a1ad565b60006020828403121561ac9457600080fd5b815167ffffffffffffffff81111561acab57600080fd5b8201601f8101841361acbc57600080fd5b6158388482516020840161aa6b565b6000855161acdd818460208a0161a189565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161ad17816001840160208a0161a189565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161ad5581600284016020890161a189565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161ad9781600284016020880161a189565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061ade2604083018461a1ad565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161ae5981601f85016020870161a189565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061aec6604083018461a1ad565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061af18604083018461a1ad565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161af8f81601485016020870161a189565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061afd6604083018561a1ad565b8281036020840152615a74818561a1ad565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161b04f81600185016020870161a189565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161b09581846020870161a189565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161b14881604b85016020870161a189565b91909101604b0192915050565b600060ff821660ff810361b16b5761b16b61a729565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161b1d281602985016020870161a189565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201526000615a78608083018461a1ad565b60006020828403121561b23857600080fd5b815167ffffffffffffffff81111561b24f57600080fd5b82016060818503121561b26157600080fd5b61b26961aa42565b81518060030b811461b27a57600080fd5b8152602082015167ffffffffffffffff81111561b29657600080fd5b61b2a28682850161aadb565b602083015250604082015167ffffffffffffffff81111561b2c257600080fd5b61b2ce8682850161aadb565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161b33a81602185016020870161a189565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161b52681602185016020880161a189565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161b56381602e84016020880161a189565b01602e01949350505050565b6000825161b58181846020870161a189565b9190910192915050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161b1d281602985016020870161a189565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161b64781602285016020870161a189565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161b68c81600e85016020870161a189565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b8082018082111561583c5761583c61a729565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161b77d81601885016020880161a189565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161b7ba81601c84016020880161a189565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161b8c081846020870161a189565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b808202811582820484141761583c5761583c61a729565b6001815b600184111561b9415780850481111561b9255761b92561a729565b600184161561b93357908102905b60019390931c92800261b90a565b935093915050565b60008261b9585750600161583c565b8161b9655750600061583c565b816001811461b97b576002811461b9855761b9a1565b600191505061583c565b60ff84111561b9965761b99661a729565b50506001821b61583c565b5060208310610133831016604e8410600b841016171561b9c4575081810a61583c565b61b9d1600019848461b906565b806000190482111561b9e55761b9e561a729565b029392505050565b6000615a78838361b949565b81810360008312801583831316838312821617156193435761934361a729565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161ba5181601c85016020870161a189565b91909101601c0192915050565b6000600019820361ba715761ba7161a729565b5060010190565b60008161ba875761ba8761a729565b506000190190565b6000835161baa181846020880161a189565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161badb81600184016020880161a189565b0160010194935050505056fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220837c7d9916de10b20cdb18567d8c7679613426bbd7b0b72548d8000a412f307b64736f6c634300081a00336080604052348015600f57600080fd5b506001600055610f06806100246000396000f3fe60806040526004361061006e5760003560e01c8063c51316911161004b578063c5131691146100d5578063c9028a36146100f5578063e04d4f9714610115578063f05b6abf1461012857005b8063357fc5a214610077578063676cc054146100975780636ed70169146100c057005b3661007557005b005b34801561008357600080fd5b50610075610092366004610709565b610148565b6100aa6100a5366004610745565b6101de565b6040516100b79190610840565b60405180910390f35b3480156100cc57600080fd5b5061007561023f565b3480156100e157600080fd5b506100756100f0366004610709565b610274565b34801561010157600080fd5b50610075610110366004610853565b61034f565b6100756101233660046109b3565b61038b565b34801561013457600080fd5b50610075610143366004610a9f565b6103cf565b610150610404565b61017273ffffffffffffffffffffffffffffffffffffffff8316338386610447565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101d96001600055565b505050565b60607fd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a35250161020e6020860186610b89565b848460405161021f93929190610bed565b60405180910390a1506040805160208101909152600081525b9392505050565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61027c610404565b6000610289600285610c26565b9050806000036102c5576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102e773ffffffffffffffffffffffffffffffffffffffff8416338484610447565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101d96001600055565b7f689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b3382604051610380929190610c61565b60405180910390a150565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa33348585856040516103c2959493929190610d53565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146338484846040516103c29493929190610ddd565b600260005403610440576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526104dc9085906104e2565b50505050565b600061050473ffffffffffffffffffffffffffffffffffffffff84168361057d565b905080516000141580156105295750808060200190518101906105279190610e97565b155b156101d9576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b606061023883836000846000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105b09190610eb4565b60006040518083038185875af1925050503d80600081146105ed576040519150601f19603f3d011682016040523d82523d6000602084013e6105f2565b606091505b509150915061060286838361060c565b9695505050505050565b6060826106215761061c8261069b565b610238565b8151158015610645575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610694576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610574565b5080610238565b8051156106ab5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff8116811461070457600080fd5b919050565b60008060006060848603121561071e57600080fd5b8335925061072e602085016106e0565b915061073c604085016106e0565b90509250925092565b6000806000838503604081121561075b57600080fd5b602081121561076957600080fd5b50839250602084013567ffffffffffffffff81111561078757600080fd5b8401601f8101861361079857600080fd5b803567ffffffffffffffff8111156107af57600080fd5b8660208284010111156107c157600080fd5b939660209190910195509293505050565b60005b838110156107ed5781810151838201526020016107d5565b50506000910152565b6000815180845261080e8160208601602086016107d2565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061023860208301846107f6565b60006020828403121561086557600080fd5b813567ffffffffffffffff81111561087c57600080fd5b82016080818503121561023857600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156109045761090461088e565b604052919050565b600082601f83011261091d57600080fd5b813567ffffffffffffffff8111156109375761093761088e565b61096860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016108bd565b81815284602083860101111561097d57600080fd5b816020850160208301376000918101602001919091529392505050565b80151581146106dd57600080fd5b80356107048161099a565b6000806000606084860312156109c857600080fd5b833567ffffffffffffffff8111156109df57600080fd5b6109eb8682870161090c565b935050602084013591506040840135610a038161099a565b809150509250925092565b600067ffffffffffffffff821115610a2857610a2861088e565b5060051b60200190565b600082601f830112610a4357600080fd5b8135610a56610a5182610a0e565b6108bd565b8082825260208201915060208360051b860101925085831115610a7857600080fd5b602085015b83811015610a95578035835260209283019201610a7d565b5095945050505050565b600080600060608486031215610ab457600080fd5b833567ffffffffffffffff811115610acb57600080fd5b8401601f81018613610adc57600080fd5b8035610aea610a5182610a0e565b8082825260208201915060208360051b850101925088831115610b0c57600080fd5b602084015b83811015610b4e57803567ffffffffffffffff811115610b3057600080fd5b610b3f8b60208389010161090c565b84525060209283019201610b11565b509550505050602084013567ffffffffffffffff811115610b6e57600080fd5b610b7a86828701610a32565b92505061073c604085016109a8565b600060208284031215610b9b57600080fd5b610238826106e0565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff84168152604060208201526000610c1d604083018486610ba4565b95945050505050565b600082610c5c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610c9f836106e0565b16604082015273ffffffffffffffffffffffffffffffffffffffff610cc6602084016106e0565b166060820152600080604084013590508060808401525060608301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610d1257600080fd5b830160208101903567ffffffffffffffff811115610d2f57600080fd5b803603821315610d3e57600080fd5b608060a085015261060260c085018284610ba4565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610d8860a08301866107f6565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610dd3578151865260209586019590910190600101610db5565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610e70577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610e5b8583516107f6565b94506020938401939190910190600101610e21565b505050508281036040840152610e868186610da1565b915050610c1d606083018415159052565b600060208284031215610ea957600080fd5b81516102388161099a565b60008251610ec68184602087016107d2565b919091019291505056fea2646970667358221220ac505b52c7a260db2019d521951b58b1ba27fef2e4c4759aa1fa75878c87363664736f6c634300081a00335a657461436f6e6e6563746f724e617469766555706772616465546573742e736f6ca2646970667358221220abf0b75f06ba1103471bc9bc14e67504f42acf43f8c6d972b6d4e1e218569ca264736f6c634300081a0033",
}

// ZetaConnectorNativeTestABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNativeTestMetaData.ABI instead.
var ZetaConnectorNativeTestABI = ZetaConnectorNativeTestMetaData.ABI

// ZetaConnectorNativeTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNativeTestMetaData.Bin instead.
var ZetaConnectorNativeTestBin = ZetaConnectorNativeTestMetaData.Bin

// DeployZetaConnectorNativeTest deploys a new Ethereum contract, binding an instance of ZetaConnectorNativeTest to it.
func DeployZetaConnectorNativeTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZetaConnectorNativeTest, error) {
	parsed, err := ZetaConnectorNativeTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNativeTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorNativeTest{ZetaConnectorNativeTestCaller: ZetaConnectorNativeTestCaller{contract: contract}, ZetaConnectorNativeTestTransactor: ZetaConnectorNativeTestTransactor{contract: contract}, ZetaConnectorNativeTestFilterer: ZetaConnectorNativeTestFilterer{contract: contract}}, nil
}

// ZetaConnectorNativeTest is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNativeTest struct {
	ZetaConnectorNativeTestCaller     // Read-only binding to the contract
	ZetaConnectorNativeTestTransactor // Write-only binding to the contract
	ZetaConnectorNativeTestFilterer   // Log filterer for contract events
}

// ZetaConnectorNativeTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNativeTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNativeTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNativeTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNativeTestSession struct {
	Contract     *ZetaConnectorNativeTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ZetaConnectorNativeTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNativeTestCallerSession struct {
	Contract *ZetaConnectorNativeTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ZetaConnectorNativeTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNativeTestTransactorSession struct {
	Contract     *ZetaConnectorNativeTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ZetaConnectorNativeTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNativeTestRaw struct {
	Contract *ZetaConnectorNativeTest // Generic contract binding to access the raw methods on
}

// ZetaConnectorNativeTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNativeTestCallerRaw struct {
	Contract *ZetaConnectorNativeTestCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNativeTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNativeTestTransactorRaw struct {
	Contract *ZetaConnectorNativeTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNativeTest creates a new instance of ZetaConnectorNativeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeTest(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNativeTest, error) {
	contract, err := bindZetaConnectorNativeTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTest{ZetaConnectorNativeTestCaller: ZetaConnectorNativeTestCaller{contract: contract}, ZetaConnectorNativeTestTransactor: ZetaConnectorNativeTestTransactor{contract: contract}, ZetaConnectorNativeTestFilterer: ZetaConnectorNativeTestFilterer{contract: contract}}, nil
}

// NewZetaConnectorNativeTestCaller creates a new read-only instance of ZetaConnectorNativeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeTestCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNativeTestCaller, error) {
	contract, err := bindZetaConnectorNativeTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestCaller{contract: contract}, nil
}

// NewZetaConnectorNativeTestTransactor creates a new write-only instance of ZetaConnectorNativeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNativeTestTransactor, error) {
	contract, err := bindZetaConnectorNativeTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestTransactor{contract: contract}, nil
}

// NewZetaConnectorNativeTestFilterer creates a new log filterer instance of ZetaConnectorNativeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNativeTestFilterer, error) {
	contract, err := bindZetaConnectorNativeTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestFilterer{contract: contract}, nil
}

// bindZetaConnectorNativeTest binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNativeTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNativeTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNativeTest.Contract.ZetaConnectorNativeTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.ZetaConnectorNativeTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.ZetaConnectorNativeTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNativeTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.DEFAULTADMINROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.DEFAULTADMINROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) ISTEST() (bool, error) {
	return _ZetaConnectorNativeTest.Contract.ISTEST(&_ZetaConnectorNativeTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) ISTEST() (bool, error) {
	return _ZetaConnectorNativeTest.Contract.ISTEST(&_ZetaConnectorNativeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.PAUSERROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.PAUSERROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.TSSROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.TSSROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.WITHDRAWERROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeTest.Contract.WITHDRAWERROLE(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) ExcludeArtifacts() ([]string, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeArtifacts(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeArtifacts(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) ExcludeContracts() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeContracts(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeContracts(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeSelectors(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeSelectors(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) ExcludeSenders() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeSenders(&_ZetaConnectorNativeTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.ExcludeSenders(&_ZetaConnectorNativeTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) Failed() (bool, error) {
	return _ZetaConnectorNativeTest.Contract.Failed(&_ZetaConnectorNativeTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) Failed() (bool, error) {
	return _ZetaConnectorNativeTest.Contract.Failed(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _ZetaConnectorNativeTest.Contract.TargetArtifactSelectors(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _ZetaConnectorNativeTest.Contract.TargetArtifactSelectors(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TargetArtifacts() ([]string, error) {
	return _ZetaConnectorNativeTest.Contract.TargetArtifacts(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) TargetArtifacts() ([]string, error) {
	return _ZetaConnectorNativeTest.Contract.TargetArtifacts(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TargetContracts() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.TargetContracts(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.TargetContracts(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _ZetaConnectorNativeTest.Contract.TargetInterfaces(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _ZetaConnectorNativeTest.Contract.TargetInterfaces(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ZetaConnectorNativeTest.Contract.TargetSelectors(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _ZetaConnectorNativeTest.Contract.TargetSelectors(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNativeTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TargetSenders() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.TargetSenders(&_ZetaConnectorNativeTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _ZetaConnectorNativeTest.Contract.TargetSenders(&_ZetaConnectorNativeTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) SetUp() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.SetUp(&_ZetaConnectorNativeTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.SetUp(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestTSSUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testTSSUpgrade")
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestTSSUpgrade() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestTSSUpgrade(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestTSSUpgrade() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestTSSUpgrade(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testTSSUpgradeFailsIfSenderIsNotTSSUpdater")
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestTSSUpgradeFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testTSSUpgradeFailsIfZeroAddress")
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestTSSUpgradeFailsIfZeroAddress() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestTSSUpgradeFailsIfZeroAddress(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestTSSUpgradeFailsIfZeroAddress() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestTSSUpgradeFailsIfZeroAddress(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestUpgradeAndWithdraw is a paid mutator transaction binding the contract method 0xaf298bb1.
//
// Solidity: function testUpgradeAndWithdraw() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestUpgradeAndWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testUpgradeAndWithdraw")
}

// TestUpgradeAndWithdraw is a paid mutator transaction binding the contract method 0xaf298bb1.
//
// Solidity: function testUpgradeAndWithdraw() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestUpgradeAndWithdraw() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestUpgradeAndWithdraw(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestUpgradeAndWithdraw is a paid mutator transaction binding the contract method 0xaf298bb1.
//
// Solidity: function testUpgradeAndWithdraw() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestUpgradeAndWithdraw() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestUpgradeAndWithdraw(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdraw is a paid mutator transaction binding the contract method 0xd509b16c.
//
// Solidity: function testWithdraw() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdraw")
}

// TestWithdraw is a paid mutator transaction binding the contract method 0xd509b16c.
//
// Solidity: function testWithdraw() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdraw() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdraw(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdraw is a paid mutator transaction binding the contract method 0xd509b16c.
//
// Solidity: function testWithdraw() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdraw() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdraw(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20 is a paid mutator transaction binding the contract method 0x3cba0107.
//
// Solidity: function testWithdrawAndCallReceiveERC20() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawAndCallReceiveERC20(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawAndCallReceiveERC20")
}

// TestWithdrawAndCallReceiveERC20 is a paid mutator transaction binding the contract method 0x3cba0107.
//
// Solidity: function testWithdrawAndCallReceiveERC20() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawAndCallReceiveERC20() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveERC20(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20 is a paid mutator transaction binding the contract method 0x3cba0107.
//
// Solidity: function testWithdrawAndCallReceiveERC20() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawAndCallReceiveERC20() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveERC20(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xc1909972.
//
// Solidity: function testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer")
}

// TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xc1909972.
//
// Solidity: function testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xc1909972.
//
// Solidity: function testWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveERC20FailsIfSenderIsNotWithdrawer(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20Partial is a paid mutator transaction binding the contract method 0xdcf7d037.
//
// Solidity: function testWithdrawAndCallReceiveERC20Partial() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawAndCallReceiveERC20Partial(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawAndCallReceiveERC20Partial")
}

// TestWithdrawAndCallReceiveERC20Partial is a paid mutator transaction binding the contract method 0xdcf7d037.
//
// Solidity: function testWithdrawAndCallReceiveERC20Partial() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawAndCallReceiveERC20Partial() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveERC20Partial(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveERC20Partial is a paid mutator transaction binding the contract method 0xdcf7d037.
//
// Solidity: function testWithdrawAndCallReceiveERC20Partial() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawAndCallReceiveERC20Partial() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveERC20Partial(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveNoParams is a paid mutator transaction binding the contract method 0x49346558.
//
// Solidity: function testWithdrawAndCallReceiveNoParams() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawAndCallReceiveNoParams(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawAndCallReceiveNoParams")
}

// TestWithdrawAndCallReceiveNoParams is a paid mutator transaction binding the contract method 0x49346558.
//
// Solidity: function testWithdrawAndCallReceiveNoParams() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawAndCallReceiveNoParams() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveNoParams(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveNoParams is a paid mutator transaction binding the contract method 0x49346558.
//
// Solidity: function testWithdrawAndCallReceiveNoParams() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawAndCallReceiveNoParams() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveNoParams(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveOnCall is a paid mutator transaction binding the contract method 0xb0a64d03.
//
// Solidity: function testWithdrawAndCallReceiveOnCall() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawAndCallReceiveOnCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawAndCallReceiveOnCall")
}

// TestWithdrawAndCallReceiveOnCall is a paid mutator transaction binding the contract method 0xb0a64d03.
//
// Solidity: function testWithdrawAndCallReceiveOnCall() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawAndCallReceiveOnCall() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveOnCall(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveOnCall is a paid mutator transaction binding the contract method 0xb0a64d03.
//
// Solidity: function testWithdrawAndCallReceiveOnCall() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawAndCallReceiveOnCall() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveOnCall(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveOnCallTNotAllowedWithArbitraryCall is a paid mutator transaction binding the contract method 0x95665330.
//
// Solidity: function testWithdrawAndCallReceiveOnCallTNotAllowedWithArbitraryCall() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawAndCallReceiveOnCallTNotAllowedWithArbitraryCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawAndCallReceiveOnCallTNotAllowedWithArbitraryCall")
}

// TestWithdrawAndCallReceiveOnCallTNotAllowedWithArbitraryCall is a paid mutator transaction binding the contract method 0x95665330.
//
// Solidity: function testWithdrawAndCallReceiveOnCallTNotAllowedWithArbitraryCall() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawAndCallReceiveOnCallTNotAllowedWithArbitraryCall() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveOnCallTNotAllowedWithArbitraryCall(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndCallReceiveOnCallTNotAllowedWithArbitraryCall is a paid mutator transaction binding the contract method 0x95665330.
//
// Solidity: function testWithdrawAndCallReceiveOnCallTNotAllowedWithArbitraryCall() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawAndCallReceiveOnCallTNotAllowedWithArbitraryCall() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndCallReceiveOnCallTNotAllowedWithArbitraryCall(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndRevert is a paid mutator transaction binding the contract method 0xde1cb76c.
//
// Solidity: function testWithdrawAndRevert() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawAndRevert(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawAndRevert")
}

// TestWithdrawAndRevert is a paid mutator transaction binding the contract method 0xde1cb76c.
//
// Solidity: function testWithdrawAndRevert() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawAndRevert() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndRevert(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndRevert is a paid mutator transaction binding the contract method 0xde1cb76c.
//
// Solidity: function testWithdrawAndRevert() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawAndRevert() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndRevert(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x82832014.
//
// Solidity: function testWithdrawAndRevertFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawAndRevertFailsIfSenderIsNotWithdrawer")
}

// TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x82832014.
//
// Solidity: function testWithdrawAndRevertFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0x82832014.
//
// Solidity: function testWithdrawAndRevertFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawAndRevertFailsIfSenderIsNotWithdrawer(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xfe574f84.
//
// Solidity: function testWithdrawFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawFailsIfSenderIsNotWithdrawer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawFailsIfSenderIsNotWithdrawer")
}

// TestWithdrawFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xfe574f84.
//
// Solidity: function testWithdrawFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawFailsIfSenderIsNotWithdrawer(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawFailsIfSenderIsNotWithdrawer is a paid mutator transaction binding the contract method 0xfe574f84.
//
// Solidity: function testWithdrawFailsIfSenderIsNotWithdrawer() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawFailsIfSenderIsNotWithdrawer() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawFailsIfSenderIsNotWithdrawer(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawTogglePause is a paid mutator transaction binding the contract method 0xccb0e3f2.
//
// Solidity: function testWithdrawTogglePause() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactor) TestWithdrawTogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.contract.Transact(opts, "testWithdrawTogglePause")
}

// TestWithdrawTogglePause is a paid mutator transaction binding the contract method 0xccb0e3f2.
//
// Solidity: function testWithdrawTogglePause() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestSession) TestWithdrawTogglePause() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawTogglePause(&_ZetaConnectorNativeTest.TransactOpts)
}

// TestWithdrawTogglePause is a paid mutator transaction binding the contract method 0xccb0e3f2.
//
// Solidity: function testWithdrawTogglePause() returns()
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestTransactorSession) TestWithdrawTogglePause() (*types.Transaction, error) {
	return _ZetaConnectorNativeTest.Contract.TestWithdrawTogglePause(&_ZetaConnectorNativeTest.TransactOpts)
}

// ZetaConnectorNativeTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestCalledIterator struct {
	Event *ZetaConnectorNativeTestCalled // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestCalled)
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
		it.Event = new(ZetaConnectorNativeTestCalled)
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
func (it *ZetaConnectorNativeTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestCalled represents a Called event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*ZetaConnectorNativeTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestCalledIterator{contract: _ZetaConnectorNativeTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestCalled)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Called", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseCalled(log types.Log) (*ZetaConnectorNativeTestCalled, error) {
	event := new(ZetaConnectorNativeTestCalled)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestDepositedIterator struct {
	Event *ZetaConnectorNativeTestDeposited // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestDeposited)
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
		it.Event = new(ZetaConnectorNativeTestDeposited)
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
func (it *ZetaConnectorNativeTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestDeposited represents a Deposited event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestDeposited struct {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterDeposited(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*ZetaConnectorNativeTestDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestDepositedIterator{contract: _ZetaConnectorNativeTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestDeposited, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestDeposited)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseDeposited(log types.Log) (*ZetaConnectorNativeTestDeposited, error) {
	event := new(ZetaConnectorNativeTestDeposited)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestDepositedAndCalledIterator is returned from FilterDepositedAndCalled and is used to iterate over the raw logs and unpacked data for DepositedAndCalled events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestDepositedAndCalledIterator struct {
	Event *ZetaConnectorNativeTestDepositedAndCalled // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestDepositedAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestDepositedAndCalled)
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
		it.Event = new(ZetaConnectorNativeTestDepositedAndCalled)
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
func (it *ZetaConnectorNativeTestDepositedAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestDepositedAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestDepositedAndCalled represents a DepositedAndCalled event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestDepositedAndCalled struct {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterDepositedAndCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*ZetaConnectorNativeTestDepositedAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestDepositedAndCalledIterator{contract: _ZetaConnectorNativeTest.contract, event: "DepositedAndCalled", logs: logs, sub: sub}, nil
}

// WatchDepositedAndCalled is a free log subscription operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchDepositedAndCalled(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestDepositedAndCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestDepositedAndCalled)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseDepositedAndCalled(log types.Log) (*ZetaConnectorNativeTestDepositedAndCalled, error) {
	event := new(ZetaConnectorNativeTestDepositedAndCalled)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestExecutedIterator struct {
	Event *ZetaConnectorNativeTestExecuted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestExecuted)
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
		it.Event = new(ZetaConnectorNativeTestExecuted)
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
func (it *ZetaConnectorNativeTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestExecuted represents a Executed event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*ZetaConnectorNativeTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestExecutedIterator{contract: _ZetaConnectorNativeTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestExecuted)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Executed", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseExecuted(log types.Log) (*ZetaConnectorNativeTestExecuted, error) {
	event := new(ZetaConnectorNativeTestExecuted)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestExecutedWithERC20Iterator struct {
	Event *ZetaConnectorNativeTestExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestExecutedWithERC20)
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
		it.Event = new(ZetaConnectorNativeTestExecutedWithERC20)
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
func (it *ZetaConnectorNativeTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*ZetaConnectorNativeTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestExecutedWithERC20Iterator{contract: _ZetaConnectorNativeTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestExecutedWithERC20)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseExecutedWithERC20(log types.Log) (*ZetaConnectorNativeTestExecutedWithERC20, error) {
	event := new(ZetaConnectorNativeTestExecutedWithERC20)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedERC20Iterator struct {
	Event *ZetaConnectorNativeTestReceivedERC20 // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestReceivedERC20)
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
		it.Event = new(ZetaConnectorNativeTestReceivedERC20)
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
func (it *ZetaConnectorNativeTestReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestReceivedERC20 represents a ReceivedERC20 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*ZetaConnectorNativeTestReceivedERC20Iterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestReceivedERC20Iterator{contract: _ZetaConnectorNativeTest.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestReceivedERC20)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
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

// ParseReceivedERC20 is a log parse operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseReceivedERC20(log types.Log) (*ZetaConnectorNativeTestReceivedERC20, error) {
	event := new(ZetaConnectorNativeTestReceivedERC20)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedNoParamsIterator struct {
	Event *ZetaConnectorNativeTestReceivedNoParams // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestReceivedNoParams)
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
		it.Event = new(ZetaConnectorNativeTestReceivedNoParams)
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
func (it *ZetaConnectorNativeTestReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestReceivedNoParams represents a ReceivedNoParams event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*ZetaConnectorNativeTestReceivedNoParamsIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestReceivedNoParamsIterator{contract: _ZetaConnectorNativeTest.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestReceivedNoParams)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
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

// ParseReceivedNoParams is a log parse operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseReceivedNoParams(log types.Log) (*ZetaConnectorNativeTestReceivedNoParams, error) {
	event := new(ZetaConnectorNativeTestReceivedNoParams)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedNonPayableIterator struct {
	Event *ZetaConnectorNativeTestReceivedNonPayable // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestReceivedNonPayable)
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
		it.Event = new(ZetaConnectorNativeTestReceivedNonPayable)
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
func (it *ZetaConnectorNativeTestReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestReceivedNonPayable represents a ReceivedNonPayable event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*ZetaConnectorNativeTestReceivedNonPayableIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestReceivedNonPayableIterator{contract: _ZetaConnectorNativeTest.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestReceivedNonPayable)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
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

// ParseReceivedNonPayable is a log parse operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseReceivedNonPayable(log types.Log) (*ZetaConnectorNativeTestReceivedNonPayable, error) {
	event := new(ZetaConnectorNativeTestReceivedNonPayable)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestReceivedOnCallIterator is returned from FilterReceivedOnCall and is used to iterate over the raw logs and unpacked data for ReceivedOnCall events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedOnCallIterator struct {
	Event *ZetaConnectorNativeTestReceivedOnCall // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestReceivedOnCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestReceivedOnCall)
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
		it.Event = new(ZetaConnectorNativeTestReceivedOnCall)
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
func (it *ZetaConnectorNativeTestReceivedOnCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestReceivedOnCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestReceivedOnCall represents a ReceivedOnCall event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedOnCall struct {
	Sender  common.Address
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceivedOnCall is a free log retrieval operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterReceivedOnCall(opts *bind.FilterOpts) (*ZetaConnectorNativeTestReceivedOnCallIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestReceivedOnCallIterator{contract: _ZetaConnectorNativeTest.contract, event: "ReceivedOnCall", logs: logs, sub: sub}, nil
}

// WatchReceivedOnCall is a free log subscription operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchReceivedOnCall(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestReceivedOnCall) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestReceivedOnCall)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
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

// ParseReceivedOnCall is a log parse operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseReceivedOnCall(log types.Log) (*ZetaConnectorNativeTestReceivedOnCall, error) {
	event := new(ZetaConnectorNativeTestReceivedOnCall)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedPayableIterator struct {
	Event *ZetaConnectorNativeTestReceivedPayable // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestReceivedPayable)
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
		it.Event = new(ZetaConnectorNativeTestReceivedPayable)
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
func (it *ZetaConnectorNativeTestReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestReceivedPayable represents a ReceivedPayable event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedPayable struct {
	Sender common.Address
	Value  *big.Int
	Str    string
	Num    *big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedPayable is a free log retrieval operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*ZetaConnectorNativeTestReceivedPayableIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestReceivedPayableIterator{contract: _ZetaConnectorNativeTest.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestReceivedPayable)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
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

// ParseReceivedPayable is a log parse operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseReceivedPayable(log types.Log) (*ZetaConnectorNativeTestReceivedPayable, error) {
	event := new(ZetaConnectorNativeTestReceivedPayable)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedRevertIterator struct {
	Event *ZetaConnectorNativeTestReceivedRevert // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestReceivedRevert)
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
		it.Event = new(ZetaConnectorNativeTestReceivedRevert)
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
func (it *ZetaConnectorNativeTestReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestReceivedRevert represents a ReceivedRevert event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*ZetaConnectorNativeTestReceivedRevertIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestReceivedRevertIterator{contract: _ZetaConnectorNativeTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestReceivedRevert)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
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

// ParseReceivedRevert is a log parse operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseReceivedRevert(log types.Log) (*ZetaConnectorNativeTestReceivedRevert, error) {
	event := new(ZetaConnectorNativeTestReceivedRevert)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestRevertedIterator struct {
	Event *ZetaConnectorNativeTestReverted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestReverted)
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
		it.Event = new(ZetaConnectorNativeTestReverted)
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
func (it *ZetaConnectorNativeTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestReverted represents a Reverted event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestReverted struct {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ZetaConnectorNativeTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestRevertedIterator{contract: _ZetaConnectorNativeTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestReverted)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Reverted", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseReverted(log types.Log) (*ZetaConnectorNativeTestReverted, error) {
	event := new(ZetaConnectorNativeTestReverted)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestUpdatedGatewayTSSAddressIterator struct {
	Event *ZetaConnectorNativeTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestUpdatedGatewayTSSAddress)
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
		it.Event = new(ZetaConnectorNativeTestUpdatedGatewayTSSAddress)
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
func (it *ZetaConnectorNativeTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestUpdatedGatewayTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*ZetaConnectorNativeTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestUpdatedGatewayTSSAddressIterator{contract: _ZetaConnectorNativeTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestUpdatedGatewayTSSAddress)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*ZetaConnectorNativeTestUpdatedGatewayTSSAddress, error) {
	event := new(ZetaConnectorNativeTestUpdatedGatewayTSSAddress)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddressIterator is returned from FilterUpdatedZetaConnectorTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedZetaConnectorTSSAddress events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddressIterator struct {
	Event *ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress)
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
		it.Event = new(ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress)
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
func (it *ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress represents a UpdatedZetaConnectorTSSAddress event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedZetaConnectorTSSAddress is a free log retrieval operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterUpdatedZetaConnectorTSSAddress(opts *bind.FilterOpts) (*ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddressIterator{contract: _ZetaConnectorNativeTest.contract, event: "UpdatedZetaConnectorTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedZetaConnectorTSSAddress is a free log subscription operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchUpdatedZetaConnectorTSSAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "UpdatedZetaConnectorTSSAddress", log); err != nil {
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

// ParseUpdatedZetaConnectorTSSAddress is a log parse operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseUpdatedZetaConnectorTSSAddress(log types.Log) (*ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress, error) {
	event := new(ZetaConnectorNativeTestUpdatedZetaConnectorTSSAddress)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "UpdatedZetaConnectorTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawnIterator struct {
	Event *ZetaConnectorNativeTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestWithdrawn)
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
		it.Event = new(ZetaConnectorNativeTestWithdrawn)
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
func (it *ZetaConnectorNativeTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestWithdrawn represents a Withdrawn event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeTestWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestWithdrawnIterator{contract: _ZetaConnectorNativeTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestWithdrawn)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseWithdrawn(log types.Log) (*ZetaConnectorNativeTestWithdrawn, error) {
	event := new(ZetaConnectorNativeTestWithdrawn)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawnAndCalledIterator struct {
	Event *ZetaConnectorNativeTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestWithdrawnAndCalled)
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
		it.Event = new(ZetaConnectorNativeTestWithdrawnAndCalled)
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
func (it *ZetaConnectorNativeTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawnAndCalled struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeTestWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestWithdrawnAndCalledIterator{contract: _ZetaConnectorNativeTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestWithdrawnAndCalled, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestWithdrawnAndCalled)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*ZetaConnectorNativeTestWithdrawnAndCalled, error) {
	event := new(ZetaConnectorNativeTestWithdrawnAndCalled)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawnAndRevertedIterator struct {
	Event *ZetaConnectorNativeTestWithdrawnAndReverted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestWithdrawnAndReverted)
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
		it.Event = new(ZetaConnectorNativeTestWithdrawnAndReverted)
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
func (it *ZetaConnectorNativeTestWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawnAndReverted struct {
	To            common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeTestWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestWithdrawnAndRevertedIterator{contract: _ZetaConnectorNativeTest.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestWithdrawnAndReverted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestWithdrawnAndReverted)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
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

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseWithdrawnAndReverted(log types.Log) (*ZetaConnectorNativeTestWithdrawnAndReverted, error) {
	event := new(ZetaConnectorNativeTestWithdrawnAndReverted)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestWithdrawnV2Iterator is returned from FilterWithdrawnV2 and is used to iterate over the raw logs and unpacked data for WithdrawnV2 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawnV2Iterator struct {
	Event *ZetaConnectorNativeTestWithdrawnV2 // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestWithdrawnV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestWithdrawnV2)
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
		it.Event = new(ZetaConnectorNativeTestWithdrawnV2)
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
func (it *ZetaConnectorNativeTestWithdrawnV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestWithdrawnV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestWithdrawnV2 represents a WithdrawnV2 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestWithdrawnV2 struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnV2 is a free log retrieval operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterWithdrawnV2(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeTestWithdrawnV2Iterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "WithdrawnV2", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestWithdrawnV2Iterator{contract: _ZetaConnectorNativeTest.contract, event: "WithdrawnV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnV2 is a free log subscription operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchWithdrawnV2(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestWithdrawnV2, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "WithdrawnV2", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestWithdrawnV2)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
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

// ParseWithdrawnV2 is a log parse operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseWithdrawnV2(log types.Log) (*ZetaConnectorNativeTestWithdrawnV2, error) {
	event := new(ZetaConnectorNativeTestWithdrawnV2)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogIterator struct {
	Event *ZetaConnectorNativeTestLog // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLog)
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
		it.Event = new(ZetaConnectorNativeTestLog)
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
func (it *ZetaConnectorNativeTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLog represents a Log event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLog(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogIterator{contract: _ZetaConnectorNativeTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLog) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLog)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLog(log types.Log) (*ZetaConnectorNativeTestLog, error) {
	event := new(ZetaConnectorNativeTestLog)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogAddressIterator struct {
	Event *ZetaConnectorNativeTestLogAddress // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogAddress)
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
		it.Event = new(ZetaConnectorNativeTestLogAddress)
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
func (it *ZetaConnectorNativeTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogAddress represents a LogAddress event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogAddressIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogAddress)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_address", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogAddress(log types.Log) (*ZetaConnectorNativeTestLogAddress, error) {
	event := new(ZetaConnectorNativeTestLogAddress)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogArrayIterator struct {
	Event *ZetaConnectorNativeTestLogArray // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogArray)
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
		it.Event = new(ZetaConnectorNativeTestLogArray)
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
func (it *ZetaConnectorNativeTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogArray represents a LogArray event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogArrayIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogArrayIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogArray) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogArray)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_array", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogArray(log types.Log) (*ZetaConnectorNativeTestLogArray, error) {
	event := new(ZetaConnectorNativeTestLogArray)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogArray0Iterator struct {
	Event *ZetaConnectorNativeTestLogArray0 // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogArray0)
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
		it.Event = new(ZetaConnectorNativeTestLogArray0)
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
func (it *ZetaConnectorNativeTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogArray0 represents a LogArray0 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogArray0Iterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogArray0Iterator{contract: _ZetaConnectorNativeTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogArray0)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_array0", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogArray0(log types.Log) (*ZetaConnectorNativeTestLogArray0, error) {
	event := new(ZetaConnectorNativeTestLogArray0)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogArray1Iterator struct {
	Event *ZetaConnectorNativeTestLogArray1 // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogArray1)
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
		it.Event = new(ZetaConnectorNativeTestLogArray1)
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
func (it *ZetaConnectorNativeTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogArray1 represents a LogArray1 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogArray1Iterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogArray1Iterator{contract: _ZetaConnectorNativeTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogArray1)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_array1", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogArray1(log types.Log) (*ZetaConnectorNativeTestLogArray1, error) {
	event := new(ZetaConnectorNativeTestLogArray1)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogBytesIterator struct {
	Event *ZetaConnectorNativeTestLogBytes // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogBytes)
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
		it.Event = new(ZetaConnectorNativeTestLogBytes)
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
func (it *ZetaConnectorNativeTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogBytes represents a LogBytes event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogBytesIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogBytesIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogBytes)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogBytes(log types.Log) (*ZetaConnectorNativeTestLogBytes, error) {
	event := new(ZetaConnectorNativeTestLogBytes)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogBytes32Iterator struct {
	Event *ZetaConnectorNativeTestLogBytes32 // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogBytes32)
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
		it.Event = new(ZetaConnectorNativeTestLogBytes32)
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
func (it *ZetaConnectorNativeTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogBytes32 represents a LogBytes32 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogBytes32Iterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogBytes32Iterator{contract: _ZetaConnectorNativeTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogBytes32)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogBytes32(log types.Log) (*ZetaConnectorNativeTestLogBytes32, error) {
	event := new(ZetaConnectorNativeTestLogBytes32)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogIntIterator struct {
	Event *ZetaConnectorNativeTestLogInt // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogInt)
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
		it.Event = new(ZetaConnectorNativeTestLogInt)
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
func (it *ZetaConnectorNativeTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogInt represents a LogInt event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogIntIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogIntIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogInt) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogInt)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_int", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogInt(log types.Log) (*ZetaConnectorNativeTestLogInt, error) {
	event := new(ZetaConnectorNativeTestLogInt)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedAddressIterator struct {
	Event *ZetaConnectorNativeTestLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedAddress)
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
		it.Event = new(ZetaConnectorNativeTestLogNamedAddress)
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
func (it *ZetaConnectorNativeTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedAddress represents a LogNamedAddress event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedAddressIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedAddress)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedAddress(log types.Log) (*ZetaConnectorNativeTestLogNamedAddress, error) {
	event := new(ZetaConnectorNativeTestLogNamedAddress)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedArrayIterator struct {
	Event *ZetaConnectorNativeTestLogNamedArray // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedArray)
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
		it.Event = new(ZetaConnectorNativeTestLogNamedArray)
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
func (it *ZetaConnectorNativeTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedArray represents a LogNamedArray event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedArrayIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedArrayIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedArray)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedArray(log types.Log) (*ZetaConnectorNativeTestLogNamedArray, error) {
	event := new(ZetaConnectorNativeTestLogNamedArray)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedArray0Iterator struct {
	Event *ZetaConnectorNativeTestLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedArray0)
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
		it.Event = new(ZetaConnectorNativeTestLogNamedArray0)
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
func (it *ZetaConnectorNativeTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedArray0 represents a LogNamedArray0 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedArray0Iterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedArray0Iterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedArray0)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedArray0(log types.Log) (*ZetaConnectorNativeTestLogNamedArray0, error) {
	event := new(ZetaConnectorNativeTestLogNamedArray0)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedArray1Iterator struct {
	Event *ZetaConnectorNativeTestLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedArray1)
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
		it.Event = new(ZetaConnectorNativeTestLogNamedArray1)
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
func (it *ZetaConnectorNativeTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedArray1 represents a LogNamedArray1 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedArray1Iterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedArray1Iterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedArray1)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedArray1(log types.Log) (*ZetaConnectorNativeTestLogNamedArray1, error) {
	event := new(ZetaConnectorNativeTestLogNamedArray1)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedBytesIterator struct {
	Event *ZetaConnectorNativeTestLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedBytes)
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
		it.Event = new(ZetaConnectorNativeTestLogNamedBytes)
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
func (it *ZetaConnectorNativeTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedBytes represents a LogNamedBytes event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedBytesIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedBytesIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedBytes)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedBytes(log types.Log) (*ZetaConnectorNativeTestLogNamedBytes, error) {
	event := new(ZetaConnectorNativeTestLogNamedBytes)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedBytes32Iterator struct {
	Event *ZetaConnectorNativeTestLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedBytes32)
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
		it.Event = new(ZetaConnectorNativeTestLogNamedBytes32)
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
func (it *ZetaConnectorNativeTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedBytes32Iterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedBytes32)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedBytes32(log types.Log) (*ZetaConnectorNativeTestLogNamedBytes32, error) {
	event := new(ZetaConnectorNativeTestLogNamedBytes32)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedDecimalIntIterator struct {
	Event *ZetaConnectorNativeTestLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedDecimalInt)
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
		it.Event = new(ZetaConnectorNativeTestLogNamedDecimalInt)
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
func (it *ZetaConnectorNativeTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedDecimalIntIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedDecimalInt)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*ZetaConnectorNativeTestLogNamedDecimalInt, error) {
	event := new(ZetaConnectorNativeTestLogNamedDecimalInt)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedDecimalUintIterator struct {
	Event *ZetaConnectorNativeTestLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedDecimalUint)
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
		it.Event = new(ZetaConnectorNativeTestLogNamedDecimalUint)
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
func (it *ZetaConnectorNativeTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedDecimalUintIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedDecimalUint)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*ZetaConnectorNativeTestLogNamedDecimalUint, error) {
	event := new(ZetaConnectorNativeTestLogNamedDecimalUint)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedIntIterator struct {
	Event *ZetaConnectorNativeTestLogNamedInt // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedInt)
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
		it.Event = new(ZetaConnectorNativeTestLogNamedInt)
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
func (it *ZetaConnectorNativeTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedInt represents a LogNamedInt event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedIntIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedIntIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedInt)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedInt(log types.Log) (*ZetaConnectorNativeTestLogNamedInt, error) {
	event := new(ZetaConnectorNativeTestLogNamedInt)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedStringIterator struct {
	Event *ZetaConnectorNativeTestLogNamedString // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedString)
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
		it.Event = new(ZetaConnectorNativeTestLogNamedString)
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
func (it *ZetaConnectorNativeTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedString represents a LogNamedString event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedStringIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedStringIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedString)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedString(log types.Log) (*ZetaConnectorNativeTestLogNamedString, error) {
	event := new(ZetaConnectorNativeTestLogNamedString)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedUintIterator struct {
	Event *ZetaConnectorNativeTestLogNamedUint // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogNamedUint)
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
		it.Event = new(ZetaConnectorNativeTestLogNamedUint)
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
func (it *ZetaConnectorNativeTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogNamedUint represents a LogNamedUint event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogNamedUintIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogNamedUintIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogNamedUint)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogNamedUint(log types.Log) (*ZetaConnectorNativeTestLogNamedUint, error) {
	event := new(ZetaConnectorNativeTestLogNamedUint)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogStringIterator struct {
	Event *ZetaConnectorNativeTestLogString // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogString)
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
		it.Event = new(ZetaConnectorNativeTestLogString)
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
func (it *ZetaConnectorNativeTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogString represents a LogString event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogString(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogStringIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogStringIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogString) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogString)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_string", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogString(log types.Log) (*ZetaConnectorNativeTestLogString, error) {
	event := new(ZetaConnectorNativeTestLogString)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogUintIterator struct {
	Event *ZetaConnectorNativeTestLogUint // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogUint)
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
		it.Event = new(ZetaConnectorNativeTestLogUint)
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
func (it *ZetaConnectorNativeTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogUint represents a LogUint event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogUintIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogUintIterator{contract: _ZetaConnectorNativeTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogUint) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogUint)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_uint", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogUint(log types.Log) (*ZetaConnectorNativeTestLogUint, error) {
	event := new(ZetaConnectorNativeTestLogUint)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogsIterator struct {
	Event *ZetaConnectorNativeTestLogs // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeTestLogs)
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
		it.Event = new(ZetaConnectorNativeTestLogs)
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
func (it *ZetaConnectorNativeTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeTestLogs represents a Logs event raised by the ZetaConnectorNativeTest contract.
type ZetaConnectorNativeTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) FilterLogs(opts *bind.FilterOpts) (*ZetaConnectorNativeTestLogsIterator, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeTestLogsIterator{contract: _ZetaConnectorNativeTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeTestLogs) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeTestLogs)
				if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "logs", log); err != nil {
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
func (_ZetaConnectorNativeTest *ZetaConnectorNativeTestFilterer) ParseLogs(log types.Log) (*ZetaConnectorNativeTestLogs, error) {
	event := new(ZetaConnectorNativeTestLogs)
	if err := _ZetaConnectorNativeTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
