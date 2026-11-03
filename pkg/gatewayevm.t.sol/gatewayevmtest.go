// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayevm

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

// GatewayEVMTestMetaData contains all meta data concerning the GatewayEVMTest contract.
var GatewayEVMTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ASSET_HANDLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testExecuteFailsIfDestinationIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevert\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteRevertFailsIfSenderIsNotTSS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteWithERC20FailsIfNotCustodyOrConnector\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteWithMsgContextFailsIfDestinationIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteWithNonReturnApprovalToken\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testExecuteWithTokenRevertingOnZeroApproval\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNoParams\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNoParamsTogglePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNonPayable\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveOnCallFails\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveOnCallUsingAuthCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceiveOnRevertFails\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForwardCallToReceivePayable\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertWithERC20FailsIfNotCustodyOrConnector\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetConnectorFailsIfSenderIsNotAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetConnectorFailsIfSet\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetConnectorFailsIfZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetCustodyFailsIfSenderIsNotAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetCustodyFailsIfSet\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetCustodyFailsIfZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetDepositAllowedAssetOnlyAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetDepositPausedOnlyAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgrade\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfSenderIsNotTSSUpdater\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testTSSUpgradeFailsIfZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpgradeAndForwardCallToReceivePayable\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpgradePreservesDepositPauseFunctions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositedAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedV2\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrandedFundsRefunded\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unwhitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedAdditionalActionFee\",\"inputs\":[{\"name\":\"oldFeeWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newFeeWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedCustodyTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedDepositAllowedAsset\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedDepositPaused\",\"inputs\":[{\"name\":\"paused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Whitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AdditionalActionDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ArbitraryCallDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AssetDepositNotAllowed\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExcessETHProvided\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectValueProvided\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFee\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x600c8054600160ff199182168117909255601f8054909116909117905560a06040525f608052603180546001600160a01b0319169055348015603f575f80fd5b5062010512806200004f5f395ff3fe608060405234801561000f575f80fd5b506004361061031a575f3560e01c806385226c81116101a8578063ce5871e1116100f3578063e54182191161009e578063f68bd1c011610079578063f68bd1c01461056b578063fa18c09b14610573578063fa7626d41461057b578063fe7bdbb214610588575f80fd5b8063e541821914610534578063e63ab1e91461053c578063e6afc79014610563575f80fd5b8063dd51e82f116100ce578063dd51e82f1461051c578063df23745a14610524578063e20c9f711461052c575f80fd5b8063ce5871e114610504578063cebad2a61461050c578063d38b66cd14610514575f80fd5b8063a783c78911610153578063b5508aa91161012e578063b5508aa9146104dc578063ba414fa6146104e4578063ccf20616146104fc575f80fd5b8063a783c789146104a5578063b0464fdc146104cc578063b124dbf5146104d4575f80fd5b8063a217fddf11610183578063a217fddf1461048e578063a397ffd214610495578063a56f7a4b1461049d575f80fd5b806385226c811461043d57806385f438c114610452578063916a17c614610479575f80fd5b806343fd8c7d1161026857806366d9a9a0116102135780637a380ebf116101ee5780637a380ebf146104255780637d7f772a1461042d5780637ebf744f14610435575f80fd5b806366d9a9a0146104005780636bdd212b1461041557806376b6f88e1461041d575f80fd5b806351010e491161024357806351010e49146103bb57806352ff5939146103c35780635d62c860146103cb575f80fd5b806343fd8c7d146103a357806344671b94146103ab5780634df42da1146103b3575f80fd5b80631855c337116102c85780633801e58f116102a35780633801e58f1461038b5780633e5e3c23146103935780633f7286f41461039b575f80fd5b80631855c337146103505780631ed7831c146103585780632ade388014610376575f80fd5b80630ab10776116102f85780630ab10776146103385780631226c65514610340578063130daf5914610348575f80fd5b806304b974f91461031e578063070f2ad0146103285780630a9254e414610330575b5f80fd5b610326610590565b005b610326610766565b61032661092e565b610326611308565b6103266114a1565b6103266115d9565b61032661176d565b6103606118ab565b60405161036d919061a6cd565b60405180910390f35b61037e61190b565b60405161036d919061a746565b610326611a47565b610360611ea2565b610360611f00565b610326611f5e565b610326612296565b610326612328565b610326612460565b6103266125fc565b6103f27f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b981565b60405190815260200161036d565b610408612b49565b60405161036d919061a8a7565b610326612cad565b610326612e48565b61032661332f565b61032661376f565b610326613a51565b610445613b8f565b60405161036d919061a99d565b6103f27f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b610481613c5a565b60405161036d919061a9af565b6103f25f81565b610326613d3b565b61032661406a565b6103f27f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b610481614188565b610326614269565b6104456144dd565b6104ec6145a8565b604051901515815260200161036d565b610326614678565b610326614796565b6103266148ce565b610326614ad7565b610326614b7f565b6103266152c5565b610360615469565b6103266154c7565b6103f27f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6103266154dd565b61032661571d565b610326615901565b601f546104ec9060ff1681565b610326615c62565b6040805160048082526024820183526020820180516001600160e01b03167f6ed7016900000000000000000000000000000000000000000000000000000000179052602754925163ca669fa760e01b81529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263ca669fa792610622926001600160a01b031691016001600160a01b0391909116815260200190565b5f604051808303815f87803b158015610639575f80fd5b505af115801561064b573d5f803e3d5ffd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b1580156106b8575f80fd5b505af11580156106ca573d5f803e3d5ffd5b5050601f546040517f38e225270000000000000000000000000000000000000000000000000000000081526101009091046001600160a01b031692506338e225279150610720906031905f90869060040161aa44565b5f604051808303815f875af115801561073b573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610762919081019061ab60565b5050565b6027546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b1580156107bc575f80fd5b505af11580156107ce573d5f803e3d5ffd5b5050602754604080516001600160a01b0390921660248301525f60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250610884919060040161ab92565b5f604051808303815f87803b15801561089b575f80fd5b505af11580156108ad573d5f803e3d5ffd5b5050601f546025546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015261010090920416925063950837aa91506024015b5f604051808303815f87803b158015610916575f80fd5b505af1158015610928573d5f803e3d5ffd5b50505050565b602580547fffffffffffffffffffffffff000000000000000000000000000000000000000090811630179091556026805482166112341790556027805482166156781790556028805490911661987617905560405161098c9061a5d9565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103905ff080158015610a0e573d5f803e3d5ffd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556027546040519116908190610a579061a5e7565b6001600160a01b03928316815291166020820152604001604051809103905ff080158015610a87573d5f803e3d5ffd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081178255604080518082018252600e81527f4761746577617945564d2e736f6c0000000000000000000000000000000000006020820152602754602554925190861694810194909452604484019290925290921660648201525f91610b65916084015b60408051601f198184030181529190526020810180516001600160e01b03167fc0c53b8b00000000000000000000000000000000000000000000000000000000179052615ebb565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0384811682029290921792839055604080518082018252601081527f4552433230437573746f64792e736f6c0000000000000000000000000000000060208201526027546025549251939095048416602484015293831660448301529091166064820152919250610c0891608401610b1d565b602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691909117909155604080518082018252601a81527f5a657461436f6e6e6563746f724e6f6e4e61746976652e736f6c0000000000006020820152601f546024805460275460255495516101009094048716928401929092528516604483015284166064820152919092166084820152919250610cfa9160a40160408051601f198184030181529190526020810180516001600160e01b03167ff8c8765e00000000000000000000000000000000000000000000000000000000179052615ebb565b602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038381169190911790915560275460405163ca669fa760e01b815291166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015610d82575f80fd5b505af1158015610d94573d5f803e3d5ffd5b5050602480546027546022546040517f15d57fd40000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152908216938101939093521692506315d57fd491506044015f604051808303815f87803b158015610e04575f80fd5b505af1158015610e16573d5f803e3d5ffd5b50505050604051610e269061a5f5565b604051809103905ff080158015610e3f573d5f803e3d5ffd5b50602080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556027546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d906044015f604051808303815f87803b158015610ee8575f80fd5b505af1158015610efa573d5f803e3d5ffd5b50506025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d5691506024015f604051808303815f87803b158015610f54575f80fd5b505af1158015610f66573d5f803e3d5ffd5b5050601f546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015261010090920416925063ae7a3a6f91506024015f604051808303815f87803b158015610fce575f80fd5b505af1158015610fe0573d5f803e3d5ffd5b5050601f546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526101009092041692506310188aef91506024015f604051808303815f87803b158015611048575f80fd5b505af115801561105a573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156110b8575f80fd5b505af11580156110ca573d5f803e3d5ffd5b50506023546025546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42406024820152911692506340c10f1991506044015f604051808303815f87803b158015611136575f80fd5b505af1158015611148573d5f803e3d5ffd5b50506023546021546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526207a12060248201529116925063a9059cbb91506044016020604051808303815f875af11580156111b9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111dd919061aba4565b506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d906044015f604051808303815f87803b15801561125b575f80fd5b505af115801561126d573d5f803e3d5ffd5b5050604080516080810182526025546001600160a01b03908116825260235481166020808401918252600184860190815285519182019095525f8152606084018190528351602d80549185167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161781559251602e8054919095169116179092559251602f55909350909150603090610928908261ac52565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801561135e575f80fd5b505af1158015611370573d5f803e3d5ffd5b5050602754604080516001600160a01b0390921660248301525f60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250611426919060040161ab92565b5f604051808303815f87803b15801561143d575f80fd5b505af115801561144f573d5f803e3d5ffd5b5050601f546040517f543f66a4000000000000000000000000000000000000000000000000000000008152600160048201526101009091046001600160a01b0316925063543f66a491506024016108ff565b6025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b1580156114f7575f80fd5b505af1158015611509573d5f803e3d5ffd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015611576575f80fd5b505af1158015611588573d5f803e3d5ffd5b5050601f546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081525f60048201526101009091046001600160a01b0316925063ae7a3a6f91506024016108ff565b6040805160048082526024820183526020820180516001600160e01b03167fc9028a3600000000000000000000000000000000000000000000000000000000179052602754925163ca669fa760e01b81529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263ca669fa79261166b926001600160a01b031691016001600160a01b0391909116815260200190565b5f604051808303815f87803b158015611682575f80fd5b505af1158015611694573d5f803e3d5ffd5b5050604051630618f58760e51b81527f69b675aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015611701575f80fd5b505af1158015611713573d5f803e3d5ffd5b5050601f546020546040517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b03610100909304831694506338e225279350610720926031921690869060040161aa44565b6025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b1580156117c3575f80fd5b505af11580156117d5573d5f803e3d5ffd5b5050604051630618f58760e51b81527fb337f378000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015b5f604051808303815f87803b158015611843575f80fd5b505af1158015611855573d5f803e3d5ffd5b5050601f546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015261010090920416925063ae7a3a6f91506024016108ff565b6060601680548060200260200160405190810160405280929190818152602001828054801561190157602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116118e3575b5050505050905090565b6060601e805480602002602001604051908101604052809291908181526020015f905b82821015611a3e575f84815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015611a27578382905f5260205f2001805461199c9061abc3565b80601f01602080910402602001604051908101604052809291908181526020018280546119c89061abc3565b8015611a135780601f106119ea57610100808354040283529160200191611a13565b820191905f5260205f20905b8154815290600101906020018083116119f657829003601f168201915b50505050508152602001906001019061197f565b50505050815250508152602001906001019061192e565b50505050905090565b5f604051611a549061a603565b611a5d9061ad0d565b604051809103905ff080158015611a76573d5f803e3d5ffd5b506025546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424060248201529192508216906340c10f19906044015f604051808303815f87803b158015611adf575f80fd5b505af1158015611af1573d5f803e3d5ffd5b50506025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d5691506024015f604051808303815f87803b158015611b4b575f80fd5b505af1158015611b5d573d5f803e3d5ffd5b50506021546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301529091169250639b19251a91506024015f604051808303815f87803b158015611bbe575f80fd5b505af1158015611bd0573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015611c2e575f80fd5b505af1158015611c40573d5f803e3d5ffd5b50506021546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a0602482018190529350908416915063a9059cbb906044016020604051808303815f875af1158015611cb2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611cd6919061aba4565b506040805160048082526024820183526020820180516001600160e01b03167f6ed7016900000000000000000000000000000000000000000000000000000000179052602154925163ca669fa760e01b81529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263ca669fa792611d69926001600160a01b031691016001600160a01b0391909116815260200190565b5f604051808303815f87803b158015611d80575f80fd5b505af1158015611d92573d5f803e3d5ffd5b5050604051630618f58760e51b81527f69b675aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015611dff575f80fd5b505af1158015611e11573d5f803e3d5ffd5b5050601f546020546040517f7bbe9afa0000000000000000000000000000000000000000000000000000000081526001600160a01b0361010090930483169450637bbe9afa9350611e709260319289929116908890889060040161ad85565b5f604051808303815f87803b158015611e87575f80fd5b505af1158015611e99573d5f803e3d5ffd5b50505050505050565b6060601880548060200260200160405190810160405280929190818152602001828054801561190157602002820191905f5260205f209081546001600160a01b031681526001909101906020018083116118e3575050505050905090565b6060601780548060200260200160405190810160405280929190818152602001828054801561190157602002820191905f5260205f209081546001600160a01b031681526001909101906020018083116118e3575050505050905090565b6040805160018082528183019092525f91816020015b6060815260200190600190039081611f745790505090506040518060400160405280600f81526020017f48656c6c6f2c20466f756e647279210000000000000000000000000000000000815250815f81518110611fd357611fd361adcb565b60209081029190910101526040805160018082528183019092525f91816020016020820280368337019050509050602a815f815181106120155761201561adcb565b60209081029190910101526040516001905f9061203a9085908590859060240161adf8565b60408051601f198184030181529181526020820180516001600160e01b03167ff05b6abf00000000000000000000000000000000000000000000000000000000179052602554905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156120d5575f80fd5b505af11580156120e7573d5f803e3d5ffd5b5050602554604080516001600160a01b0390921660248301527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506121bd919060040161ab92565b5f604051808303815f87803b1580156121d4575f80fd5b505af11580156121e6573d5f803e3d5ffd5b5050601f54604080516020808201835261012382525491517f38e225270000000000000000000000000000000000000000000000000000000081526101009093046001600160a01b0390811695506338e22527945061224d9391921690869060040161ae5b565b5f604051808303815f875af1158015612268573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261228f919081019061ab60565b5050505050565b6040805160048082526024820183526020820180516001600160e01b03167f6ed7016900000000000000000000000000000000000000000000000000000000179052602754925163ca669fa760e01b81529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263ca669fa79261166b926001600160a01b031691016001600160a01b0391909116815260200190565b6025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b15801561237e575f80fd5b505af1158015612390573d5f803e3d5ffd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b1580156123fd575f80fd5b505af115801561240f573d5f803e3d5ffd5b5050601f546040517f950837aa0000000000000000000000000000000000000000000000000000000081525f60048201526101009091046001600160a01b0316925063950837aa91506024016108ff565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152670de0b6b3a7640000905f9060250160408051808303601f190181529082905260275463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156124fc575f80fd5b505af115801561250e573d5f803e3d5ffd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b15801561257b575f80fd5b505af115801561258d573d5f803e3d5ffd5b5050601f546040517fcb7ba8e50000000000000000000000000000000000000000000000000000000081526101009091046001600160a01b0316925063cb7ba8e5915084906125e5905f908690602d9060040161af63565b5f604051808303818588803b158015611e87575f80fd5b601f546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb600482015261432160248201819052915f916101009091046001600160a01b0316906391d1485490604401602060405180830381865afa15801561268c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906126b0919061aba4565b90506126bb81615ed9565b601f546027546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b0391821660248201525f926101009004909116906391d1485490604401602060405180830381865afa15801561274c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612770919061aba4565b905061277b81615f4f565b6025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b1580156127d1575f80fd5b505af11580156127e3573d5f803e3d5ffd5b5050601f546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a4015f604051808303815f87803b158015612876575f80fd5b505af1158015612888573d5f803e3d5ffd5b5050602754604080516001600160a01b03928316815291871660208301527f3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579935001905060405180910390a1601f546040517f950837aa0000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301526101009092049091169063950837aa906024015f604051808303815f87803b158015612937575f80fd5b505af1158015612949573d5f803e3d5ffd5b505050506129cb83601f60019054906101000a90046001600160a01b03166001600160a01b0316635b1125916040518163ffffffff1660e01b8152600401602060405180830381865afa1580156129a2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129c6919061af96565b615fa1565b601f546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b038581166024830152610100909204909116906391d1485490604401602060405180830381865afa158015612a58573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a7c919061aba4565b9150612a8782615f4f565b601f546027546040517f91d148540000000000000000000000000000000000000000000000000000000081527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60048201526001600160a01b03918216602482015261010090920416906391d1485490604401602060405180830381865afa158015612b15573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b39919061aba4565b9050612b4481615ed9565b505050565b6060601b805480602002602001604051908101604052809291908181526020015f905b82821015611a3e578382905f5260205f2090600202016040518060400160405290815f82018054612b9c9061abc3565b80601f0160208091040260200160405190810160405280929190818152602001828054612bc89061abc3565b8015612c135780601f10612bea57610100808354040283529160200191612c13565b820191905f5260205f20905b815481529060010190602001808311612bf657829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015612c9557602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b03191681526020019060040190602082600301049283019260010382029150808411612c575790505b50505050508152505081526020019060010190612b6c565b6040805160048082526024820183526020820180516001600160e01b03167f6ed7016900000000000000000000000000000000000000000000000000000000179052602754925163ca669fa760e01b81529192737109709ecfa91a80626ff3989d68f67f5b1dd12d9263ca669fa792612d3f926001600160a01b031691016001600160a01b0391909116815260200190565b5f604051808303815f87803b158015612d56575f80fd5b505af1158015612d68573d5f803e3d5ffd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015612dd5575f80fd5b505af1158015612de7573d5f803e3d5ffd5b5050601f5460408051602081018252610123815290517f38e225270000000000000000000000000000000000000000000000000000000081526101009092046001600160a01b031693506338e225279250610720915f90869060040161ae5b565b6025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b158015612e9e575f80fd5b505af1158015612eb0573d5f803e3d5ffd5b5050601f546040517f543f66a4000000000000000000000000000000000000000000000000000000008152600160048201526101009091046001600160a01b0316925063543f66a491506024015f604051808303815f87803b158015612f14575f80fd5b505af1158015612f26573d5f803e3d5ffd5b50505050612fa7601f60019054906101000a90046001600160a01b03166001600160a01b03166302befd246040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f7e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612fa2919061aba4565b615f4f565b7f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015613001575f80fd5b505af1158015613013573d5f803e3d5ffd5b5050601f54604080518082018252601981527f4761746577617945564d55706772616465546573742e736f6c0000000000000060208083019190915282519081019092525f825260255461307c95506001600160a01b0361010090940484169450909216616026565b5f601f60019054906101000a90046001600160a01b031690506130d4816001600160a01b03166302befd246040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f7e573d5f803e3d5ffd5b6040517f6ae23fb50000000000000000000000000000000000000000000000000000000081525f6004820152613134906001600160a01b03831690636ae23fb5906024015b602060405180830381865afa158015612f7e573d5f803e3d5ffd5b602480546040517f6ae23fb50000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526131829291841691636ae23fb59101613119565b6025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b1580156131d8575f80fd5b505af11580156131ea573d5f803e3d5ffd5b50506040517f543f66a40000000000000000000000000000000000000000000000000000000081525f60048201526001600160a01b038416925063543f66a491506024015f604051808303815f87803b158015613245575f80fd5b505af1158015613257573d5f803e3d5ffd5b505050506132c3816001600160a01b03166302befd246040518163ffffffff1660e01b8152600401602060405180830381865afa15801561329a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906132be919061aba4565b615ed9565b7f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561331d575f80fd5b505af115801561228f573d5f803e3d5ffd5b601f54604080518082018252601981527f4761746577617945564d55706772616465546573742e736f6c0000000000000060208083019190915282519081019092525f8252602554613393936001600160a01b036101009091048116939116616026565b601f54604080517fdda79b7500000000000000000000000000000000000000000000000000000000815290516101009092046001600160a01b0316915f91839163dda79b75916004808201926020929091908290030181865afa1580156133fc573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613420919061af96565b90505f601f60019054906101000a90046001600160a01b03166001600160a01b0316635b1125916040518163ffffffff1660e01b8152600401602060405180830381865afa158015613474573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613498919061af96565b604080518082018252600f81527f48656c6c6f2c20466f756e6472792100000000000000000000000000000000006020820152905191925090602a90600190670de0b6b3a7640000905f906134f59086908690869060240161afbc565b60408051601f198184030181529181526020820180516001600160e01b03167fe04d4f9700000000000000000000000000000000000000000000000000000000179052602754905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015613590575f80fd5b505af11580156135a2573d5f803e3d5ffd5b5050604051630618f58760e51b81527f69b675aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b15801561360f575f80fd5b505af1158015613621573d5f803e3d5ffd5b50506020546040517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b038c811694506338e22527935086926136779260319290911690879060040161aa44565b5f6040518083038185885af1158015613692573d5f803e3d5ffd5b50505050506040513d5f823e601f3d908101601f191682016040526136ba919081019061ab60565b5061371087601f60019054906101000a90046001600160a01b03166001600160a01b031663dda79b756040518163ffffffff1660e01b8152600401602060405180830381865afa1580156129a2573d5f803e3d5ffd5b61376586601f60019054906101000a90046001600160a01b03166001600160a01b0316635b1125916040518163ffffffff1660e01b8152600401602060405180830381865afa1580156129a2573d5f803e3d5ffd5b5050505050505050565b6040805160018082528183019092525f91816020015b60608152602001906001900390816137855790505090506040518060400160405280600f81526020017f48656c6c6f2c20466f756e647279210000000000000000000000000000000000815250815f815181106137e4576137e461adcb565b60209081029190910101526040805160018082528183019092525f91816020016020820280368337019050509050602a815f815181106138265761382661adcb565b60209081029190910101526040516001905f9061384b9085908590859060240161adf8565b60408051601f198184030181529181526020820180516001600160e01b03167ff05b6abf00000000000000000000000000000000000000000000000000000000179052602554905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156138e6575f80fd5b505af11580156138f8573d5f803e3d5ffd5b5050602554604080516001600160a01b0390921660248301527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506139ce919060040161ab92565b5f604051808303815f87803b1580156139e5575f80fd5b505af11580156139f7573d5f803e3d5ffd5b5050601f546020546040517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b03610100909304831694506338e22527935061224d926031921690869060040161aa44565b6025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b158015613aa7575f80fd5b505af1158015613ab9573d5f803e3d5ffd5b5050604051630618f58760e51b81527f0c8dc016000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015b5f604051808303815f87803b158015613b27575f80fd5b505af1158015613b39573d5f803e3d5ffd5b5050601f546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526101009092041692506310188aef91506024016108ff565b6060601a805480602002602001604051908101604052809291908181526020015f905b82821015611a3e578382905f5260205f20018054613bcf9061abc3565b80601f0160208091040260200160405190810160405280929190818152602001828054613bfb9061abc3565b8015613c465780601f10613c1d57610100808354040283529160200191613c46565b820191905f5260205f20905b815481529060010190602001808311613c2957829003601f168201915b505050505081526020019060010190613bb2565b6060601d805480602002602001604051908101604052809291908181526020015f905b82821015611a3e575f8481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015613d2357602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b03191681526020019060040190602082600301049283019260010382029150808411613ce55790505b50505050508152505081526020019060010190613c7d565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015613dc6575f80fd5b505af1158015613dd8573d5f803e3d5ffd5b5050604080518082018252600181527f31000000000000000000000000000000000000000000000000000000000000006020820152905161012393507fd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a3525019250613e4291849161afe5565b60405180910390a1601f546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015613ed9575f80fd5b505af1158015613eeb573d5f803e3d5ffd5b505060208054604080518082018252600181527f310000000000000000000000000000000000000000000000000000000000000093810193909352516001600160a01b0390911693507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f9250613f62915f9161b006565b60405180910390a260275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015613fc0575f80fd5b505af1158015613fd2573d5f803e3d5ffd5b5050601f5460408051602080820183526001600160a01b038781168352815484518086018652600181527f31000000000000000000000000000000000000000000000000000000000000009381019390935293517f38e22527000000000000000000000000000000000000000000000000000000008152610100909504811696506338e225279550610720949293169160040161ae5b565b6027546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b1580156140c0575f80fd5b505af11580156140d2573d5f803e3d5ffd5b5050602754604080516001600160a01b0390921660248301525f60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061182c919060040161ab92565b6060601c805480602002602001604051908101604052809291908181526020015f905b82821015611a3e575f8481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561425157602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116142135790505b505050505081525050815260200190600101906141ab565b602354602654604051620186a0602482018190526001600160a01b0393841660448301529290911660648201525f9060840160408051601f198184030181529181526020820180516001600160e01b03167f357fc5a200000000000000000000000000000000000000000000000000000000179052602554905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015614336575f80fd5b505af1158015614348573d5f803e3d5ffd5b5050602554604080516001600160a01b0390921660248301527f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b960448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925061441e919060040161ab92565b5f604051808303815f87803b158015614435575f80fd5b505af1158015614447573d5f803e3d5ffd5b5050601f546023546026546040517faa0c0fc10000000000000000000000000000000000000000000000000000000081526101009093046001600160a01b03908116955063aa0c0fc194506144ac939281169291169087908790602d9060040161b01e565b5f604051808303815f87803b1580156144c3575f80fd5b505af11580156144d5573d5f803e3d5ffd5b505050505050565b60606019805480602002602001604051908101604052809291908181526020015f905b82821015611a3e578382905f5260205f2001805461451d9061abc3565b80601f01602080910402602001604051908101604052809291908181526020018280546145499061abc3565b80156145945780601f1061456b57610100808354040283529160200191614594565b820191905f5260205f20905b81548152906001019060200180831161457757829003601f168201915b505050505081526020019060010190614500565b6008545f9060ff16156145bf575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c6564000000000000000000000000000000000000000000000000000060248301525f9163667f9d7090604401602060405180830381865afa15801561464d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614671919061b072565b1415905090565b6027546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b1580156146ce575f80fd5b505af11580156146e0573d5f803e3d5ffd5b5050602754604080516001600160a01b0390921660248301525f60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250613b10919060040161ab92565b6025546040516303223eab60e11b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b1580156147ec575f80fd5b505af11580156147fe573d5f803e3d5ffd5b5050604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b15801561486b575f80fd5b505af115801561487d573d5f803e3d5ffd5b5050601f546040517f10188aef0000000000000000000000000000000000000000000000000000000081525f60048201526101009091046001600160a01b031692506310188aef91506024016108ff565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152670de0b6b3a7640000905f9060250160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801561496a575f80fd5b505af115801561497c573d5f803e3d5ffd5b5050602554604080516001600160a01b0390921660248301527f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250614a52919060040161ab92565b5f604051808303815f87803b158015614a69575f80fd5b505af1158015614a7b573d5f803e3d5ffd5b5050601f546020546040517fcb7ba8e50000000000000000000000000000000000000000000000000000000081526001600160a01b036101009093048316945063cb7ba8e5935086926125e59216908690602d9060040161af63565b604080516020810182525f8082529151614af691607b9160240161afe5565b60408051601f198184030181529181526020820180516001600160e01b03167f676cc05400000000000000000000000000000000000000000000000000000000179052602754905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa79060240161166b565b60285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015614bd5575f80fd5b505af1158015614be7573d5f803e3d5ffd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250614cbd919060040161ab92565b5f604051808303815f87803b158015614cd4575f80fd5b505af1158015614ce6573d5f803e3d5ffd5b50505050601f60019054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b81526004015f604051808303815f87803b158015614d37575f80fd5b505af1158015614d49573d5f803e3d5ffd5b505060285460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015614da3575f80fd5b505af1158015614db5573d5f803e3d5ffd5b5050602854604080516001600160a01b0390921660248301527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250614e8b919060040161ab92565b5f604051808303815f87803b158015614ea2575f80fd5b505af1158015614eb4573d5f803e3d5ffd5b50505050601f60019054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015614f05575f80fd5b505af1158015614f17573d5f803e3d5ffd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015614f71575f80fd5b505af1158015614f83573d5f803e3d5ffd5b50505050601f60019054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b81526004015f604051808303815f87803b158015614fd4575f80fd5b505af1158015614fe6573d5f803e3d5ffd5b5050604080516004808252602480830184526020830180516001600160e01b03167f6ed70169000000000000000000000000000000000000000000000000000000001790529251630618f58760e51b81527fd93c06650000000000000000000000000000000000000000000000000000000091810191909152909350737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091015f604051808303815f87803b158015615097575f80fd5b505af11580156150a9573d5f803e3d5ffd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015615103575f80fd5b505af1158015615115573d5f803e3d5ffd5b5050601f546020546040517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b03610100909304831694506338e22527935061516f926031921690869060040161aa44565b5f604051808303815f875af115801561518a573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526151b1919081019061ab60565b5060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015615208575f80fd5b505af115801561521a573d5f803e3d5ffd5b50505050601f60019054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561526b575f80fd5b505af115801561527d573d5f803e3d5ffd5b505060275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa7915060240161166b565b60275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801561531b575f80fd5b505af115801561532d573d5f803e3d5ffd5b5050602754604080516001600160a01b0390921660248301525f60448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506153e3919060040161ab92565b5f604051808303815f87803b1580156153fa575f80fd5b505af115801561540c573d5f803e3d5ffd5b5050601f546023546040517f4f27eb2a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260016024820152610100909204169250634f27eb2a91506044016108ff565b6060601580548060200260200160405190810160405280929190818152602001828054801561190157602002820191905f5260205f209081546001600160a01b031681526001909101906020018083116118e3575050505050905090565b5f6040516154d49061a611565b611a5d9061b089565b602354602654604051620186a0602482018190526001600160a01b0393841660448301529290911660648201525f9060840160408051601f198184030181529181526020820180516001600160e01b03167f357fc5a200000000000000000000000000000000000000000000000000000000179052602554905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156155aa575f80fd5b505af11580156155bc573d5f803e3d5ffd5b5050602554604080516001600160a01b0390921660248301527f584a0b16e9f616d90ccec14a0b852c19aceccfd3d60699398a57dce2b0de01b960448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fe2517d3f00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250615692919060040161ab92565b5f604051808303815f87803b1580156156a9575f80fd5b505af11580156156bb573d5f803e3d5ffd5b5050601f546023546026546040517f7bbe9afa0000000000000000000000000000000000000000000000000000000081526001600160a01b0361010090940484169550637bbe9afa94506144ac9360319381169216908890889060040161ad85565b6040805160018082528183019092525f91816020015b60608152602001906001900390816157335790505090506040518060400160405280600f81526020017f48656c6c6f2c20466f756e647279210000000000000000000000000000000000815250815f815181106157925761579261adcb565b60209081029190910101526040805160018082528183019092525f91816020016020820280368337019050509050602a815f815181106157d4576157d461adcb565b60209081029190910101526040516001905f906157f99085908590859060240161adf8565b60408051601f198184030181529181526020820180516001600160e01b03167ff05b6abf00000000000000000000000000000000000000000000000000000000179052602754905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015615894575f80fd5b505af11580156158a6573d5f803e3d5ffd5b5050604051630618f58760e51b81527f69b675aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024016139ce565b6040517f68656c6c6f0000000000000000000000000000000000000000000000000000006020820152670de0b6b3a7640000905f9060250160408051601f198184030181529190526020549091506001600160a01b031631615963815f61603b565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b1580156159ee575f80fd5b505af1158015615a00573d5f803e3d5ffd5b505050507f689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b601f60019054906101000a90046001600160a01b0316602d604051615a4b92919061b101565b60405180910390a1601f546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015615ae2575f80fd5b505af1158015615af4573d5f803e3d5ffd5b50506020546040515f93506001600160a01b0390911691507fde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e03590615b4690670de0b6b3a7640000908790602d9061b122565b60405180910390a360275460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015615ba4575f80fd5b505af1158015615bb6573d5f803e3d5ffd5b5050601f546020546040517fcb7ba8e50000000000000000000000000000000000000000000000000000000081526001600160a01b036101009093048316945063cb7ba8e593508792615c129216908790602d9060040161af63565b5f604051808303818588803b158015615c29575f80fd5b505af1158015615c3b573d5f803e3d5ffd5b50506020546001600160a01b03163192506109289150829050670de0b6b3a764000061603b565b60408051808201909152600f81527f48656c6c6f2c20466f756e64727921000000000000000000000000000000000060208083019190915254602a90600190670de0b6b3a764000090615cc0905f906001600160a01b03163161603b565b5f848484604051602401615cd69392919061afbc565b60408051601f198184030181529181526020820180516001600160e01b03167fe04d4f9700000000000000000000000000000000000000000000000000000000179052602754905163ca669fa760e01b81526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015615d71575f80fd5b505af1158015615d83573d5f803e3d5ffd5b5050604051630618f58760e51b81527f69b675aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015615df0575f80fd5b505af1158015615e02573d5f803e3d5ffd5b5050601f546020546040517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b03610100909304831694506338e2252793508692615e5f92603192911690879060040161aa44565b5f6040518083038185885af1158015615e7a573d5f803e3d5ffd5b50505050506040513d5f823e601f3d908101601f19168201604052615ea2919081019061ab60565b5060205461228f905f906001600160a01b03163161603b565b5f615ec461a61f565b615ecf848483616093565b9150505b92915050565b6040517fa59828850000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063a5982885906024015b5f6040518083038186803b158015615f3d575f80fd5b505afa15801561228f573d5f803e3d5ffd5b6040517f0c9fd5810000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90630c9fd58190602401615f27565b6040517f515361f60000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015282166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063515361f6906044015b5f6040518083038186803b158015616014575f80fd5b505afa1580156144d5573d5f803e3d5ffd5b61602e61a61f565b61228f858585848661610d565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c5490604401615ffe565b5f8061609f85846161ec565b90506161026040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f787900000081525082866040516020016160ed92919061afe5565b604051602081830303815290604052856161f7565b9150505b9392505050565b6040516303223eab60e11b81526001600160a01b03821660048201528190737109709ecfa91a80626ff3989d68f67f5b1dd12d9081906306447d56906024015f604051808303815f87803b158015616163575f80fd5b505af1925050508015616174575060015b6161895761618487878787616224565b611e99565b61619587878787616224565b806001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156161cd575f80fd5b505af11580156161df573d5f803e3d5ffd5b5050505050505050505050565b5f616106838361623c565b60c0810151515f901561621a5761621384848460c00151616256565b9050616106565b61621384846163f4565b5f61622f84836164d9565b905061228f8582856164e4565b5f6162478383616892565b616106838360200151846161f7565b5f8061626061689d565b90505f61626d868361696c565b90505f6162838260600151836020015185616df5565b90505f61629283838989617002565b90505f61629e82617e6e565b602081015181519192509060030b15616311578982604001516040516020016162c892919061b151565b60408051601f19818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526163089160040161ab92565b60405180910390fd5b5f6163536040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a20000000000000000000000081525083600161802f565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d906163a690849060040161ab92565b602060405180830381865afa1580156163c1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906163e5919061af96565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081525f908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc9259061644890879060040161ab92565b5f60405180830381865afa158015616462573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052616489919081019061ab60565b90505f6164b682856040516020016164a292919061b1b2565b60405160208183030381529060405261821e565b90506001600160a01b038116615ecf5784846040516020016162c892919061b1c6565b5f616247838361822f565b6040517f667f9d700000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201527fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d905f90829063667f9d7090604401602060405180830381865afa15801561657d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906165a1919061b072565b90508061673a575f6165b28661823b565b604080518082018252600581527f352e302e300000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015290915061663b905b6040805180820182525f8082526020918201528151808301909252845182528085019082015290618319565b8061664657505f8451115b156166c4576040517f4f1ef2860000000000000000000000000000000000000000000000000000000081526001600160a01b03871690634f1ef28690616692908890889060040161afe5565b5f604051808303815f87803b1580156166a9575f80fd5b505af11580156166bb573d5f803e3d5ffd5b50505050616734565b6040517f3659cfe60000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152871690633659cfe6906024015f604051808303815f87803b15801561671d575f80fd5b505af115801561672f573d5f803e3d5ffd5b505050505b5061228f565b805f6167458261823b565b604080518082018252600581527f352e302e300000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506167a69061660f565b806167b157505f8551115b15616831576040517f9623609d0000000000000000000000000000000000000000000000000000000081526001600160a01b03831690639623609d906167ff908a908a908a9060040161b256565b5f604051808303815f87803b158015616816575f80fd5b505af1158015616828573d5f803e3d5ffd5b50505050611e99565b6040517f99a88ec40000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015287811660248301528316906399a88ec4906044015f604051808303815f87803b1580156161cd575f80fd5b61076282825f61832c565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c9061692490849060040161b286565b5f60405180830381865afa15801561693e573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052616965919081019061b2cc565b9250505090565b61699e6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b5f737109709ecfa91a80626ff3989d68f67f5b1dd12d90506169e86040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6169f18561842b565b60208201525f616a0086618804565b90505f836001600160a01b031663d930a0e66040518163ffffffff1660e01b81526004015f60405180830381865afa158015616a3e573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052616a65919081019061b2cc565b86838560200151604051602001616a7f949392919061b311565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291505f906001600160a01b038616906360f9bb1190616ad690859060040161ab92565b5f60405180830381865afa158015616af0573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052616b17919081019061b2cc565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f690616b5f90849060040161b3e1565b602060405180830381865afa158015616b7a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190616b9e919061aba4565b616bb357816040516020016162c8919061b432565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890616bf890849060040161b4b6565b5f60405180830381865afa158015616c12573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052616c39919081019061b2cc565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f690616c8090849060040161b507565b602060405180830381865afa158015616c9b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190616cbf919061aba4565b15616d50576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890616d0990849060040161b507565b5f60405180830381865afa158015616d23573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052616d4a919081019061b2cc565b60408501525b846001600160a01b03166349c4fac882865f0151604051602001616d74919061b558565b6040516020818303038152906040526040518363ffffffff1660e01b8152600401616da092919061b5b6565b5f60405180830381865afa158015616dba573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052616de1919081019061b2cc565b606085015250608083015250949350505050565b60408051600480825260a082019092526060915f9190816020015b6060815260200190600190039081616e105790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250815f81518110616e6f57616e6f61adcb565b60200260200101819052506040518060400160405280600381526020017f2d726c000000000000000000000000000000000000000000000000000000000081525081600181518110616ec357616ec361adcb565b602002602001018190525084604051602001616edf919061b5da565b60405160208183030381529060405281600281518110616f0157616f0161adcb565b602002602001018190525082604051602001616f1d919061b638565b60405160208183030381529060405281600381518110616f3f57616f3f61adcb565b60200260200101819052505f616f5482617e6e565b602080820151604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000818501908152825180840184525f8082529086015282518084019093529051825292810192909252919250616fe3906040805180820182525f8082526020918201528151808301909252845182528085019082015290618a80565b616ff857856040516020016162c8919061b670565b9695505050505050565b60a08101516040805180820182525f80825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d9015617051565b511590565b6171c55782602001511561710d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401616308565b8260c00151156171c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401616308565b6040805160ff80825261200082019092525f91816020015b60608152602001906001900390816171dd5790505090505f6040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806172379061b71a565b935060ff168151811061724c5761724c61adcb565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e370000000000000000000000000000000000000081525060405160200161729d919061b738565b6040516020818303038152906040528282806172b89061b71a565b935060ff16815181106172cd576172cd61adcb565b60200260200101819052506040518060400160405280600681526020017f6465706c6f79000000000000000000000000000000000000000000000000000081525082828061731a9061b71a565b935060ff168151811061732f5761732f61adcb565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d6500000000000000000000000000000000000081525082828061737c9061b71a565b935060ff16815181106173915761739161adcb565b602002602001018190525087602001518282806173ad9061b71a565b935060ff16815181106173c2576173c261adcb565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163745061746800000000000000000000000000000000000081525082828061740f9061b71a565b935060ff16815181106174245761742461adcb565b60209081029190910101528751828261743c8161b71a565b935060ff16815181106174515761745161adcb565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e4964000000000000000000000000000000000000000000000081525082828061749e9061b71a565b935060ff16815181106174b3576174b361adcb565b60200260200101819052506174c746618ade565b82826174d28161b71a565b935060ff16815181106174e7576174e761adcb565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c6500000000000000000000000000000000008152508282806175349061b71a565b935060ff16815181106175495761754961adcb565b6020026020010181905250868282806175619061b71a565b935060ff16815181106175765761757661adcb565b60209081029190910101528551156176995760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826175c78161b71a565b935060ff16815181106175dc576175dc61adcb565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d9061762c90899060040161ab92565b5f60405180830381865afa158015617646573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261766d919081019061b2cc565b82826176788161b71a565b935060ff168151811061768d5761768d61adcb565b60200260200101819052505b8460200151156177695760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826176e28161b71a565b935060ff16815181106176f7576176f761adcb565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806177449061b71a565b935060ff16815181106177595761775961adcb565b602002602001018190525061792e565b6177a061704c8660a001516040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6178335760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826177e38161b71a565b935060ff16815181106177f8576177f861adcb565b60200260200101819052508460a00151604051602001617818919061b5da565b6040516020818303038152906040528282806177449061b71a565b8460c001511580156178755750604080890151815180830183525f8082526020918201528251808401909352815183529081019082015261787390511590565b155b1561792e5760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826178b98161b71a565b935060ff16815181106178ce576178ce61adcb565b60200260200101819052506178e288618b7b565b6040516020016178f2919061b5da565b60405160208183030381529060405282828061790d9061b71a565b935060ff16815181106179225761792261adcb565b60200260200101819052505b604080860151815180830183525f8082526020918201528251808401909352815183529081019082015261796190511590565b6179f65760408051808201909152600b81527f2d2d72656c617965724964000000000000000000000000000000000000000000602082015282826179a48161b71a565b935060ff16815181106179b9576179b961adcb565b602002602001018190525084604001518282806179d59061b71a565b935060ff16815181106179ea576179ea61adcb565b60200260200101819052505b606085015115617b135760408051808201909152600681527f2d2d73616c74000000000000000000000000000000000000000000000000000060208201528282617a3f8161b71a565b935060ff1681518110617a5457617a5461adcb565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e8906024015f60405180830381865afa158015617ac0573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052617ae7919081019061b2cc565b8282617af28161b71a565b935060ff1681518110617b0757617b0761adcb565b60200260200101819052505b60e08501515115617bb95760408051808201909152600a81527f2d2d6761734c696d69740000000000000000000000000000000000000000000060208201528282617b5d8161b71a565b935060ff1681518110617b7257617b7261adcb565b6020026020010181905250617b8d8560e001515f0151618ade565b8282617b988161b71a565b935060ff1681518110617bad57617bad61adcb565b60200260200101819052505b60e08501516020015115617c635760408051808201909152600a81527f2d2d67617350726963650000000000000000000000000000000000000000000060208201528282617c068161b71a565b935060ff1681518110617c1b57617c1b61adcb565b6020026020010181905250617c378560e0015160200151618ade565b8282617c428161b71a565b935060ff1681518110617c5757617c5761adcb565b60200260200101819052505b60e08501516040015115617d0d5760408051808201909152600e81527f2d2d6d617846656550657247617300000000000000000000000000000000000060208201528282617cb08161b71a565b935060ff1681518110617cc557617cc561adcb565b6020026020010181905250617ce18560e0015160400151618ade565b8282617cec8161b71a565b935060ff1681518110617d0157617d0161adcb565b60200260200101819052505b60e08501516060015115617db75760408051808201909152601681527f2d2d6d61785072696f726974794665655065724761730000000000000000000060208201528282617d5a8161b71a565b935060ff1681518110617d6f57617d6f61adcb565b6020026020010181905250617d8b8560e0015160600151618ade565b8282617d968161b71a565b935060ff1681518110617dab57617dab61adcb565b60200260200101819052505b5f8160ff1667ffffffffffffffff811115617dd457617dd461aa75565b604051908082528060200260200182016040528015617e0757816020015b6060815260200190600190039081617df25790505b5090505f5b8260ff168160ff161015617e5f57838160ff1681518110617e2f57617e2f61adcb565b6020026020010151828260ff1681518110617e4c57617e4c61adcb565b6020908102919091010152600101617e0c565b5093505050505b949350505050565b617e9460405180606001604052805f60030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d925f91849163d145736c91617f199186910161b78f565b5f60405180830381865afa158015617f33573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052617f5a919081019061b2cc565b90505f617f678683619657565b90505f846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b8152600401617f96919061a99d565b5f604051808303815f875af1158015617fb1573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052617fd8919081019061b7d5565b805190915060030b15801590617ff15750602081015151155b80156180005750604081015151155b15616ff857815f815181106180175761801761adcb565b60200260200101516040516020016162c8919061b884565b60605f618062856040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925286518252808701908201529091506180989082905b906197a9565b156181f0575f6181128261810c846181066180d98a6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925282518252918201519181019190915290565b906197cf565b9061982d565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506181759082906197a9565b156181de57604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526181db905b82906198b1565b90505b6181e7816198d6565b92505050616106565b82156182095784846040516020016162c892919061ba61565b505060408051602081019091525f8152616106565b5f808251602084015ff09392505050565b6107628282600161832c565b60408051600481526024810182526020810180516001600160e01b03167fad3cb1cc0000000000000000000000000000000000000000000000000000000017905290516060915f9182916001600160a01b0386169161829a919061bae8565b5f60405180830381855afa9150503d805f81146182d2576040519150601f19603f3d011682016040523d82523d5f602084013e6182d7565b606091505b50915091508180156182ea575060208151115b156183035780806020019051810190617e66919061b2cc565b505060408051602081019091525f815292915050565b5f618324838361993b565b159392505050565b8160a001511561833b57505050565b5f618347848484619a13565b90505f61835382617e6e565b602081015181519192509060030b1580156183ed5750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526183ed906040805180820182525f80825260209182015281518083019092528451825280850190820152618092565b156183fa57505050505050565b6040820151511561841a5781604001516040516020016162c8919061baf3565b806040516020016162c8919061bb4a565b60605f61845e836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506184c2905b8290618a80565b1561853057604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526161069061852b908390619fa8565b6198d6565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152618591905b829061a030565b60010361865c57604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526185f6906181d4565b50604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526161069061852b905b83906198b1565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526186ba906184bb565b156187ed57604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082018190528451808601909552925184528301529061872190839061a0c4565b90505f8160018351618733919061bba1565b815181106187435761874361adcb565b602002602001015190506187e461852b6187b86040518060400160405280600581526020017f2e6a736f6e0000000000000000000000000000000000000000000000000000008152506040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f8082526020918201528151808301909252855182528086019082015290619fa8565b95945050505050565b826040516020016162c8919061bbb4565b50919050565b60605f618837836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152909150618898906184bb565b156188a657616106816198d6565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526189049061858a565b60010361896d57604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526161069061852b90618655565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526189cb906184bb565b156187ed57604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820181905284518086019095529251845283015290618a3290839061a0c4565b9050600181511115618a6e578060028251618a4d919061bba1565b81518110618a5d57618a5d61adcb565b602002602001015192505050919050565b50826040516020016162c8919061bbb4565b805182515f911115618a9357505f615ed3565b8151835160208501515f9291618aa89161bc84565b618ab2919061bba1565b905082602001518103618ac9576001915050615ed3565b82516020840151819020912014905092915050565b60605f618aea8361a16f565b60010190505f8167ffffffffffffffff811115618b0957618b0961aa75565b6040519080825280601f01601f191660200182016040528015618b33576020820181803683370190505b5090508181016020015b5f19017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084618b3d57509392505050565b604081810151815180830183525f808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e5345440000000000000000000000000000000000000000000081840190815285518087018752838152840192909252845180860190955251845290830152606091618c06905b8290618319565b15618c4657505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152618ca490618bff565b15618ce457505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152618d4290618bff565b15618d8257505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152618de090618bff565b80618e445750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152618e4490618bff565b15618e8457505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152618ee290618bff565b80618f465750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152618f4690618bff565b15618f8657505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f8082529082015283518085019094529151835290820152618fe490618bff565b806190485750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261904890618bff565b1561908857505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526190e690618bff565b8061914a5750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261914a90618bff565b1561918a57505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526191e890618bff565b1561922857505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261928690618bff565b156192c657505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261932490618bff565b1561936457505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526193c290618bff565b1561940257505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261946090618bff565b156194a057505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526194fe90618bff565b806195625750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261956290618bff565b156195a257505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201528351808501909452915183529082015261960090618bff565b1561964057505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b604080840151845191516162c8929060200161bc97565b6060805f5b84518110156196e157818582815181106196785761967861adcb565b602002602001015160405160200161969192919061b1b2565b6040516020818303038152906040529150600185516196b0919061bba1565b81146196d957816040516020016196c7919061bde5565b60405160208183030381529060405291505b60010161965c565b50604080516003808252608082019092525f91816020015b60608152602001906001900390816196f957905050905083815f815181106197235761972361adcb565b60200260200101819052506040518060400160405280600281526020017f2d63000000000000000000000000000000000000000000000000000000000000815250816001815181106197775761977761adcb565b602002602001018190525081816002815181106197965761979661adcb565b6020908102919091010152949350505050565b60208083015183518351928401515f936197c6929184919061a250565b14159392505050565b604080518082019091525f80825260208201525f6197fd845f01518560200151855f0151866020015161a35f565b905083602001518161980f919061bba1565b8451859061981e90839061bba1565b90525060208401525090919050565b604080518082019091525f8082526020820152815183511015619851575081615ed3565b60208083015190840151600191146198785750815160208481015190840151829020919020145b80156198a95782518451859061988f90839061bba1565b90525082516020850180516198a590839061bc84565b9052505b509192915050565b604080518082019091525f80825260208201526198cf83838361a47b565b5092915050565b60605f825f015167ffffffffffffffff8111156198f5576198f561aa75565b6040519080825280601f01601f19166020018201604052801561991f576020820181803683370190505b5090505f6020820190506198cf818560200151865f015161a521565b815181515f919081111561994d575081515b602080850151908401515f5b83811015619a0457825182518082146199d4575f1960208710156199b35760018461998589602061bba1565b61998f919061bc84565b61999a90600861be1d565b6199a590600261bf17565b6199af919061bba1565b1990505b81811683821681810391146199d1579750615ed39650505050505050565b50505b6199df60208661bc84565b94506199ec60208561bc84565b935050506020816199fd919061bc84565b9050619959565b5084518651616ff8919061bf22565b60605f619a1e61689d565b6040805160ff80825261200082019092529192505f9190816020015b6060815260200190600190039081619a3a5790505090505f6040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280619a949061b71a565b935060ff1681518110619aa957619aa961adcb565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e3300000000000000000000000000000000000000000000000000815250604051602001619afa919061bf41565b604051602081830303815290604052828280619b159061b71a565b935060ff1681518110619b2a57619b2a61adcb565b60200260200101819052506040518060400160405280600881526020017f76616c6964617465000000000000000000000000000000000000000000000000815250828280619b779061b71a565b935060ff1681518110619b8c57619b8c61adcb565b602002602001018190525082604051602001619ba8919061b638565b604051602081830303815290604052828280619bc39061b71a565b935060ff1681518110619bd857619bd861adcb565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e747261637400000000000000000000000000000000000000000000815250828280619c259061b71a565b935060ff1681518110619c3a57619c3a61adcb565b6020026020010181905250619c4f878461a59a565b8282619c5a8161b71a565b935060ff1681518110619c6f57619c6f61adcb565b602090810291909101015285515115619d1a5760408051808201909152600b81527f2d2d7265666572656e636500000000000000000000000000000000000000000060208201528282619cc18161b71a565b935060ff1681518110619cd657619cd661adcb565b6020026020010181905250619cee865f01518461a59a565b8282619cf98161b71a565b935060ff1681518110619d0e57619d0e61adcb565b60200260200101819052505b856080015115619d885760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b000000000000000060208201528282619d638161b71a565b935060ff1681518110619d7857619d7861adcb565b6020026020010181905250619dee565b8415619dee5760408051808201909152601281527f2d2d726571756972655265666572656e6365000000000000000000000000000060208201528282619dcd8161b71a565b935060ff1681518110619de257619de261adcb565b60200260200101819052505b60408601515115619e8a5760408051808201909152600d81527f2d2d756e73616665416c6c6f770000000000000000000000000000000000000060208201528282619e388161b71a565b935060ff1681518110619e4d57619e4d61adcb565b60200260200101819052508560400151828280619e699061b71a565b935060ff1681518110619e7e57619e7e61adcb565b60200260200101819052505b856060015115619ef45760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d657300000000000000000000000060208201528282619ed38161b71a565b935060ff1681518110619ee857619ee861adcb565b60200260200101819052505b5f8160ff1667ffffffffffffffff811115619f1157619f1161aa75565b604051908082528060200260200182016040528015619f4457816020015b6060815260200190600190039081619f2f5790505b5090505f5b8260ff168160ff161015619f9c57838160ff1681518110619f6c57619f6c61adcb565b6020026020010151828260ff1681518110619f8957619f8961adcb565b6020908102919091010152600101619f49565b50979650505050505050565b604080518082019091525f8082526020820152815183511015619fcc575081615ed3565b8151835160208501515f9291619fe19161bc84565b619feb919061bba1565b6020840151909150600190821461a00c575082516020840151819020908220145b801561a0275783518551869061a02390839061bba1565b9052505b50929392505050565b5f80825f015161a050855f01518660200151865f0151876020015161a35f565b61a05a919061bc84565b90505b8351602085015161a06e919061bc84565b81116198cf578161a07e8161bf72565b925050825f015161a0b385602001518361a098919061bba1565b865161a0a4919061bba1565b83865f0151876020015161a35f565b61a0bd919061bc84565b905061a05d565b60605f61a0d1848461a030565b61a0dc90600161bc84565b67ffffffffffffffff81111561a0f45761a0f461aa75565b60405190808252806020026020018201604052801561a12757816020015b606081526020019060019003908161a1125790505b5090505f5b815181101561a1675761a14261852b86866198b1565b82828151811061a1545761a15461adcb565b602090810291909101015260010161a12c565b509392505050565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061a1b7577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061a1e3576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061a20157662386f26fc10000830492506010015b6305f5e100831061a219576305f5e100830492506008015b612710831061a22d57612710830492506004015b6064831061a23f576064830492506002015b600a8310615ed35760010192915050565b5f8085841161a355576020841161a301575f841561a29957600161a27586602061bba1565b61a28090600861be1d565b61a28b90600261bf17565b61a295919061bba1565b1990505b835181168561a2a8898961bc84565b61a2b2919061bba1565b805190935082165b81811461a2ec5787841161a2d45787945050505050617e66565b8361a2de8161bf8a565b94505082845116905061a2ba565b61a2f6878561bc84565b945050505050617e66565b83832061a30e858861bba1565b61a318908761bc84565b91505b85821061a3535784822080820361a3405761a336868461bc84565b9350505050617e66565b61a34b60018461bba1565b92505061a31b565b505b5092949350505050565b5f838186851161a466576020851161a416575f851561a3a957600161a38587602061bba1565b61a39090600861be1d565b61a39b90600261bf17565b61a3a5919061bba1565b1990505b845181165f8761a3b98b8b61bc84565b61a3c3919061bba1565b855190915083165b82811461a4085781861061a3f05761a3e38b8b61bc84565b9650505050505050617e66565b8561a3fa8161bf72565b96505083865116905061a3cb565b859650505050505050617e66565b508383205f905b61a427868961bba1565b821161a4645785832080820361a4435783945050505050617e66565b61a44e60018561bc84565b935050818061a45c9061bf72565b92505061a41d565b505b61a470878761bc84565b979650505050505050565b604080518082019091525f80825260208201525f61a4a9855f01518660200151865f0151876020015161a35f565b60208087018051918601919091525190915061a4c5908261bba1565b83528451602086015161a4d8919061bc84565b810361a4e6575f855261a518565b8351835161a4f4919061bc84565b8551869061a50390839061bba1565b905250835161a512908261bc84565b60208601525b50909392505050565b6020811061a559578151835261a53860208461bc84565b925061a54560208361bc84565b915061a55260208261bba1565b905061a521565b5f19811561a58757600161a56e83602061bba1565b61a57a9061010061bf17565b61a584919061bba1565b90505b9151835183169219169190911790915250565b60605f61a5a7848461696c565b805160208083015160405193945061a5c19390910161bf9f565b60405160208183030381529060405291505092915050565b610c32806200bfdb83390190565b61124c806200cc0d83390190565b610eaf806200de5983390190565b610b8d806200ed0883390190565b610c48806200f89583390190565b6040518060e001604052806060815260200160608152602001606081526020015f151581526020015f151581526020015f1515815260200161a65f61a664565b905290565b6040518061010001604052805f151581526020015f15158152602001606081526020015f801916815260200160608152602001606081526020015f1515815260200161a65f60405180608001604052805f81526020015f81526020015f81526020015f81525090565b602080825282518282018190525f918401906040840190835b8181101561a70d5783516001600160a01b031683526020938401939092019160010161a6e6565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561a83f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b8801810191908801905f5b8181101561a825577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261a80f84865161a718565b602095860195909450929092019160010161a7d5565b50919750505060209485019492909201915060010161a76c565b50929695505050505050565b5f8151808452602084019350602083015f5b8281101561a89d5781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161a85d565b5093949350505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561a83f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261a911604088018261a718565b905060208201519150868103602088015261a92c818361a84b565b96505050602093840193919091019060010161a8cd565b5f82825180855260208501945060208160051b830101602085015f5b8381101561a99157601f1985840301885261a97b83835161a718565b602098890198909350919091019060010161a95f565b50909695505050505050565b602081525f616106602083018461a943565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561a83f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261aa2e604087018261a84b565b955050602093840193919091019060010161a9d5565b6001600160a01b0384541681526001600160a01b0383166020820152606060408201525f6187e4606083018461a718565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040516060810167ffffffffffffffff8111828210171561aac55761aac561aa75565b60405290565b5f8067ffffffffffffffff84111561aae55761aae561aa75565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561ab145761ab1461aa75565b60405283815290508082840185101561ab2b575f80fd5b8383602083015e5f60208583010152509392505050565b5f82601f83011261ab51575f80fd5b6161068383516020850161aacb565b5f6020828403121561ab70575f80fd5b815167ffffffffffffffff81111561ab86575f80fd5b615ecf8482850161ab42565b602081525f616106602083018461a718565b5f6020828403121561abb4575f80fd5b81518015158114616106575f80fd5b600181811c9082168061abd757607f821691505b6020821081036187fe577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b601f821115612b4457805f5260205f20601f840160051c8101602085101561ac335750805b601f840160051c820191505b8181101561228f575f815560010161ac3f565b815167ffffffffffffffff81111561ac6c5761ac6c61aa75565b61ac808161ac7a845461abc3565b8461ac0e565b6020601f82116001811461acb2575f831561ac9b5750848201515b5f19600385901b1c1916600184901b17845561228f565b5f84815260208120601f198516915b8281101561ace1578785015182556020948501946001909201910161acc1565b508482101561acfe57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b604081525f61ad4960408301600481527f5553445400000000000000000000000000000000000000000000000000000000602082015260400190565b828103602084015261610681600481527f5553445400000000000000000000000000000000000000000000000000000000602082015260400190565b6001600160a01b0386541681526001600160a01b03851660208201526001600160a01b038416604082015282606082015260a060808201525f61a47060a083018461a718565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b606081525f61ae0a606083018661a943565b82810360208401528085518083526020830191506020870192505f5b8181101561ae4457835183526020938401939092019160010161ae26565b505080925050508215156040830152949350505050565b6001600160a01b0384511681526001600160a01b0383166020820152606060408201525f6187e4606083018461a718565b6001600160a01b0381541682526001600160a01b036001820154166020830152600281015460408301525f60038201608060608501525f815461aece8161abc3565b806080880152600182165f811461aeec576001811461af265761af57565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660a089015260a082151560051b890101935061af57565b845f5260205f205f5b8381101561af4e5781548a820160a0015260019091019060200161af2f565b890160a0019450505b50919695505050505050565b6001600160a01b0384168152606060208201525f61af84606083018561a718565b8281036040840152616ff8818561ae8c565b5f6020828403121561afa6575f80fd5b81516001600160a01b0381168114616106575f80fd5b606081525f61afce606083018661a718565b602083019490945250901515604090910152919050565b6001600160a01b0383168152604060208201525f617e66604083018461a718565b828152604060208201525f617e66604083018461a718565b6001600160a01b03861681526001600160a01b038516602082015283604082015260a060608201525f61b05460a083018561a718565b828103608084015261b066818561ae8c565b98975050505050505050565b5f6020828403121561b082575f80fd5b5051919050565b604081525f61b0c560408301600381527f424e420000000000000000000000000000000000000000000000000000000000602082015260400190565b828103602084015261610681600381527f424e420000000000000000000000000000000000000000000000000000000000602082015260400190565b6001600160a01b0383168152604060208201525f617e66604083018461ae8c565b838152606060208201525f61af84606083018561a718565b5f81518060208401855e5f93019283525090919050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f61b182601a83018561b13a565b7f3a200000000000000000000000000000000000000000000000000000000000008152616102600282018561b13a565b5f617e6661b1c0838661b13a565b8461b13a565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f61b1f7601a83018561b13a565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000815261b227601982018561b13a565b7f2200000000000000000000000000000000000000000000000000000000000000815260010195945050505050565b6001600160a01b03841681526001600160a01b0383166020820152606060408201525f6187e4606083018461a718565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201525f616106608083018461a718565b5f6020828403121561b2dc575f80fd5b815167ffffffffffffffff81111561b2f2575f80fd5b8201601f8101841361b302575f80fd5b615ecf8482516020840161aacb565b5f61b31c828761b13a565b7f2f00000000000000000000000000000000000000000000000000000000000000815261b34c600182018761b13a565b90507f2f00000000000000000000000000000000000000000000000000000000000000815261b37e600182018661b13a565b90507f2f00000000000000000000000000000000000000000000000000000000000000815261b3b0600182018561b13a565b7f2e6a736f6e0000000000000000000000000000000000000000000000000000008152600501979650505050505050565b604081525f61b3f3604083018461a718565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081525f61b463601f83018461b13a565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f81527f6d6c00000000000000000000000000000000000000000000000000000000000060208201526022019392505050565b604081525f61b4c8604083018461a718565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b604081525f61b519604083018461a718565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081525f61b589601483018461b13a565b7f275d2e6b656363616b32353600000000000000000000000000000000000000008152600c019392505050565b604081525f61b5c8604083018561a718565b8281036020840152616102818561a718565b7f220000000000000000000000000000000000000000000000000000000000000081525f61b60b600183018461b13a565b7f220000000000000000000000000000000000000000000000000000000000000081526001019392505050565b5f61b643828461b13a565b7f2f6275696c642d696e666f0000000000000000000000000000000000000000008152600b019392505050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201525f616106604b83018461b13a565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f60ff821660ff810361b72f5761b72f61b6ed565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201525f616106602983018461b13a565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201525f616106608083018461a718565b5f6020828403121561b7e5575f80fd5b815167ffffffffffffffff81111561b7fb575f80fd5b82016060818503121561b80c575f80fd5b61b81461aaa2565b81518060030b811461b824575f80fd5b8152602082015167ffffffffffffffff81111561b83f575f80fd5b61b84b8682850161ab42565b602083015250604082015167ffffffffffffffff81111561b86a575f80fd5b61b8768682850161ab42565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201525f61b8db602183018461b13a565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657481527f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960208201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560408201527f616c69666965642070617468206f66207468652062617368206578656375746160608201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960808201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60a08201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960c08201527f6f75722070726f6a65637420287573696e6720666f727761726420736c61736860e08201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101008201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061012082015261013b019392505050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201525f61bab8602183018561b13a565b7f2720696e206f75747075743a20000000000000000000000000000000000000008152616102600d82018561b13a565b5f616106828461b13a565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201525f616106602983018461b13a565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201525f616106602283018461b13a565b81810381811115615ed357615ed361b6ed565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081525f61bbe5600e83018461b13a565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e7472616381527f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e7360208201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e747260408201527f6163742e6a736f6e00000000000000000000000000000000000000000000000060608201526068019392505050565b80820180821115615ed357615ed361b6ed565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081525f61bcc8601883018561b13a565b7f20696e2000000000000000000000000000000000000000000000000000000000815261bcf8600482018561b13a565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f727465642081527f6c6963656e736520666f7220626c6f636b206578706c6f72657220766572696660208201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f60408201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c60608201527f206f7220736574207468652060736b69704c6963656e73655479706560206f7060808201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060a082015260b70195945050505050565b5f61bdf0828461b13a565b7f200000000000000000000000000000000000000000000000000000000000000081526001019392505050565b8082028115828204841417615ed357615ed361b6ed565b6001815b600184111561be6f5780850481111561be535761be5361b6ed565b600184161561be6157908102905b60019390931c92800261be38565b935093915050565b5f8261be8557506001615ed3565b8161be9157505f615ed3565b816001811461bea7576002811461beb15761becd565b6001915050615ed3565b60ff84111561bec25761bec261b6ed565b50506001821b615ed3565b5060208310610133831016604e8410600b841016171561bef0575081810a615ed3565b61befc5f19848461be34565b805f190482111561bf0f5761bf0f61b6ed565b029392505050565b5f616106838361be77565b8181035f8312801583831316838312821617156198cf576198cf61b6ed565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081525f616106601c83018461b13a565b5f5f19820361bf835761bf8361b6ed565b5060010190565b5f8161bf985761bf9861b6ed565b505f190190565b5f61bfaa828561b13a565b7f3a000000000000000000000000000000000000000000000000000000000000008152616102600182018561b13a56fe608060405234801561000f575f80fd5b50604051610c32380380610c3283398101604081905261002e916100f0565b8181600361003c83826101d9565b50600461004982826101d9565b5050505050610293565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610076575f80fd5b81516001600160401b0381111561008f5761008f610053565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100bd576100bd610053565b6040528181528382016020018510156100d4575f80fd5b8160208501602083015e5f918101602001919091529392505050565b5f8060408385031215610101575f80fd5b82516001600160401b03811115610116575f80fd5b61012285828601610067565b602085015190935090506001600160401b0381111561013f575f80fd5b61014b85828601610067565b9150509250929050565b600181811c9082168061016957607f821691505b60208210810361018757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101d457805f5260205f20601f840160051c810160208510156101b25750805b601f840160051c820191505b818110156101d1575f81556001016101be565b50505b505050565b81516001600160401b038111156101f2576101f2610053565b610206816102008454610155565b8461018d565b6020601f821160018114610238575f83156102215750848201515b5f19600385901b1c1916600184901b1784556101d1565b5f84815260208120601f198516915b828110156102675787850151825560209485019460019092019101610247565b508482101561028457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610992806102a05f395ff3fe608060405234801561000f575f80fd5b50600436106100b9575f3560e01c806340c10f191161007257806395d89b411161005857806395d89b411461017c578063a9059cbb14610184578063dd62ed3e14610197575f80fd5b806340c10f191461013257806370a0823114610147575f80fd5b806318160ddd116100a257806318160ddd146100fe57806323b872dd14610110578063313ce56714610123575f80fd5b806306fdde03146100bd578063095ea7b3146100db575b5f80fd5b6100c56101dc565b6040516100d291906107a5565b60405180910390f35b6100ee6100e9366004610820565b61026c565b60405190151581526020016100d2565b6002545b6040519081526020016100d2565b6100ee61011e366004610848565b610285565b604051601281526020016100d2565b610145610140366004610820565b6102a8565b005b610102610155366004610882565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b6100c56102b6565b6100ee610192366004610820565b6102c5565b6101026101a53660046108a2565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b6060600380546101eb906108d3565b80601f0160208091040260200160405190810160405280929190818152602001828054610217906108d3565b80156102625780601f1061023957610100808354040283529160200191610262565b820191905f5260205f20905b81548152906001019060200180831161024557829003601f168201915b5050505050905090565b5f336102798185856102d2565b60019150505b92915050565b5f336102928582856102e4565b61029d8585856103b6565b506001949350505050565b6102b2828261045f565b5050565b6060600480546101eb906108d3565b5f336102798185856103b6565b6102df83838360016104b9565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103b057818110156103a2576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103b084848484035f6104b9565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610405576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b73ffffffffffffffffffffffffffffffffffffffff8216610454576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b6102df8383836105fe565b73ffffffffffffffffffffffffffffffffffffffff82166104ae576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b6102b25f83836105fe565b73ffffffffffffffffffffffffffffffffffffffff8416610508576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b73ffffffffffffffffffffffffffffffffffffffff8316610557576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b73ffffffffffffffffffffffffffffffffffffffff8085165f90815260016020908152604080832093871683529290522082905580156103b0578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105f091815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610635578060025f82825461062a9190610924565b909155506106e59050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260208190526040902054818110156106ba576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610399565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661070e57600280548290039055610739565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161079891815260200190565b60405180910390a3505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461081b575f80fd5b919050565b5f8060408385031215610831575f80fd5b61083a836107f8565b946020939093013593505050565b5f805f6060848603121561085a575f80fd5b610863846107f8565b9250610871602085016107f8565b929592945050506040919091013590565b5f60208284031215610892575f80fd5b61089b826107f8565b9392505050565b5f80604083850312156108b3575f80fd5b6108bc836107f8565b91506108ca602084016107f8565b90509250929050565b600181811c908216806108e757607f821691505b60208210810361091e577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b8082018082111561027f577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea2646970667358221220569a538c722848c143f241fcbfd3f113f81549aa32aa89f9a6169cd25cb5399e64736f6c634300081a0033608060405234801561000f575f80fd5b5060405161124c38038061124c83398101604081905261002e9161010e565b604051806040016040528060048152602001635a65746160e01b815250604051806040016040528060048152602001635a45544160e01b815250816003908161007791906101d7565b50600461008482826101d7565b5050506001600160a01b03821615806100a457506001600160a01b038116155b156100c25760405163e6c4247b60e01b815260040160405180910390fd5b600680546001600160a01b039384166001600160a01b03199182161790915560078054929093169116179055610291565b80516001600160a01b0381168114610109575f80fd5b919050565b5f806040838503121561011f575f80fd5b610128836100f3565b9150610136602084016100f3565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061016757607f821691505b60208210810361018557634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101d257805f5260205f20601f840160051c810160208510156101b05750805b601f840160051c820191505b818110156101cf575f81556001016101bc565b50505b505050565b81516001600160401b038111156101f0576101f061013f565b610204816101fe8454610153565b8461018b565b6020601f821160018114610236575f831561021f5750848201515b5f19600385901b1c1916600184901b1784556101cf565b5f84815260208120601f198516915b828110156102655787850151825560209485019460019092019101610245565b508482101561028257868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610fae8061029e5f395ff3fe608060405234801561000f575f80fd5b5060043610610115575f3560e01c806342966c68116100ad57806379cc67901161007d578063a9059cbb11610063578063a9059cbb14610286578063bff9662a14610299578063dd62ed3e146102b9575f80fd5b806379cc67901461026b57806395d89b411461027e575f80fd5b806342966c68146101fb5780635b1125911461020e57806370a082311461022e578063779e3b6314610263575f80fd5b80631e458bee116100e85780631e458bee1461018157806323b872dd14610194578063313ce567146101a7578063328a01d0146101b6575f80fd5b806306fdde0314610119578063095ea7b31461013757806315d57fd41461015a57806318160ddd1461016f575b5f80fd5b6101216102fe565b60405161012e9190610d7a565b60405180910390f35b61014a610145366004610df5565b61038e565b604051901515815260200161012e565b61016d610168366004610e1d565b6103a7565b005b6002545b60405190815260200161012e565b61016d61018f366004610e4e565b610572565b61014a6101a2366004610e7e565b610625565b6040516012815260200161012e565b6007546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161012e565b61016d610209366004610eb8565b610648565b6006546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b61017361023c366004610ecf565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b61016d610655565b61016d610279366004610df5565b610779565b61012161082a565b61014a610294366004610df5565b610839565b6005546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b6101736102c7366004610e1d565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b60606003805461030d90610eef565b80601f016020809104026020016040519081016040528092919081815260200182805461033990610eef565b80156103845780601f1061035b57610100808354040283529160200191610384565b820191905f5260205f20905b81548152906001019060200180831161036757829003601f168201915b5050505050905090565b5f3361039b818585610846565b60019150505b92915050565b60075473ffffffffffffffffffffffffffffffffffffffff1633148015906103e7575060065473ffffffffffffffffffffffffffffffffffffffff163314155b15610425576040517fcdfcef970000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216158061045c575073ffffffffffffffffffffffffffffffffffffffff8116155b15610493576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316811790935560058054918516919092161790556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff910160405180910390a16040805133815273ffffffffffffffffffffffffffffffffffffffff831660208201527f1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c910160405180910390a15050565b60055473ffffffffffffffffffffffffffffffffffffffff1633146105c5576040517f3fe32fba00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b6105cf8383610858565b808373ffffffffffffffffffffffffffffffffffffffff167fc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb8460405161061891815260200190565b60405180910390a3505050565b5f336106328582856108b6565b61063d858585610983565b506001949350505050565b6106523382610a2c565b50565b60075473ffffffffffffffffffffffffffffffffffffffff1633146106a8576040517fe700765e00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b60065473ffffffffffffffffffffffffffffffffffffffff166106f7576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600654600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0910160405180910390a1565b60055473ffffffffffffffffffffffffffffffffffffffff1633146107cc576040517f3fe32fba00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b6107d68282610a86565b8173ffffffffffffffffffffffffffffffffffffffff167f919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b18260405161081e91815260200190565b60405180910390a25050565b60606004805461030d90610eef565b5f3361039b818585610983565b6108538383836001610a9b565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166108a7576040517fec442f050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b6108b25f8383610be0565b5050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461097d578181101561096f576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602481018290526044810183905260640161041c565b61097d84848484035f610a9b565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166109d2576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8216610a21576040517fec442f050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b610853838383610be0565b73ffffffffffffffffffffffffffffffffffffffff8216610a7b576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b6108b2825f83610be0565b610a918233836108b6565b6108b28282610a2c565b73ffffffffffffffffffffffffffffffffffffffff8416610aea576040517fe602df050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8316610b39576040517f94280d620000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8085165f908152600160209081526040808320938716835292905220829055801561097d578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bd291815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c17578060025f828254610c0c9190610f40565b90915550610cc79050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020819052604090205481811015610c9c576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602481018290526044810183905260640161041c565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610cf057600280548290039055610d1b565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161061891815260200190565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610df0575f80fd5b919050565b5f8060408385031215610e06575f80fd5b610e0f83610dcd565b946020939093013593505050565b5f8060408385031215610e2e575f80fd5b610e3783610dcd565b9150610e4560208401610dcd565b90509250929050565b5f805f60608486031215610e60575f80fd5b610e6984610dcd565b95602085013595506040909401359392505050565b5f805f60608486031215610e90575f80fd5b610e9984610dcd565b9250610ea760208501610dcd565b929592945050506040919091013590565b5f60208284031215610ec8575f80fd5b5035919050565b5f60208284031215610edf575f80fd5b610ee882610dcd565b9392505050565b600181811c90821680610f0357607f821691505b602082108103610f3a577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b808201808211156103a1577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea26469706673582212206a41e2cd6fbf39f12e2917f13d39a938ebbfbe1bbc40009f596c89c4dc977dca64736f6c634300081a00336080604052348015600e575f80fd5b5060015f55610e8f806100205f395ff3fe60806040526004361061006d575f3560e01c8063c51316911161004a578063c5131691146100d2578063c9028a36146100f1578063e04d4f9714610110578063f05b6abf1461012357005b8063357fc5a214610076578063676cc054146100955780636ed70169146100be57005b3661007457005b005b348015610081575f80fd5b506100746100903660046106f4565b610142565b6100a86100a336600461072d565b6101d7565b6040516100b591906107fe565b60405180910390f35b3480156100c9575f80fd5b50610074610237565b3480156100dd575f80fd5b506100746100ec3660046106f4565b61026c565b3480156100fc575f80fd5b5061007461010b366004610810565b610344565b61007461011e366004610965565b610380565b34801561012e575f80fd5b5061007461013d366004610a49565b6103c4565b61014a6103f9565b61016c73ffffffffffffffffffffffffffffffffffffffff831633838661043a565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101d260015f55565b505050565b60607fd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a3525016102076020860186610b2b565b848460405161021893929190610b8b565b60405180910390a15060408051602081019091525f81525b9392505050565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b6102746103f9565b5f610280600285610bc3565b9050805f036102bb576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102dd73ffffffffffffffffffffffffffffffffffffffff841633848461043a565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101d260015f55565b7f689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b3382604051610375929190610bfb565b60405180910390a150565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa33348585856040516103b7959493929190610ce9565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146338484846040516103b79493929190610d70565b60025f5403610434576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f55565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526104cf9085906104d5565b50505050565b5f6104f673ffffffffffffffffffffffffffffffffffffffff84168361056e565b905080515f1415801561051a5750808060200190518101906105189190610e28565b155b156101d2576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b606061023083835f845f808573ffffffffffffffffffffffffffffffffffffffff16848660405161059f9190610e43565b5f6040518083038185875af1925050503d805f81146105d9576040519150601f19603f3d011682016040523d82523d5f602084013e6105de565b606091505b50915091506105ee8683836105f8565b9695505050505050565b60608261060d5761060882610687565b610230565b8151158015610631575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610680576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610565565b5080610230565b8051156106975780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff811681146106ef575f80fd5b919050565b5f805f60608486031215610706575f80fd5b83359250610716602085016106cc565b9150610724604085016106cc565b90509250925092565b5f805f8385036040811215610740575f80fd5b602081121561074d575f80fd5b50839250602084013567ffffffffffffffff81111561076a575f80fd5b8401601f8101861361077a575f80fd5b803567ffffffffffffffff811115610790575f80fd5b8660208284010111156107a1575f80fd5b939660209190910195509293505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61023060208301846107b2565b5f60208284031215610820575f80fd5b813567ffffffffffffffff811115610836575f80fd5b820160808185031215610230575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156108bb576108bb610847565b604052919050565b5f82601f8301126108d2575f80fd5b813567ffffffffffffffff8111156108ec576108ec610847565b61091d60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610874565b818152846020838601011115610931575f80fd5b816020850160208301375f918101602001919091529392505050565b80151581146106c9575f80fd5b80356106ef8161094d565b5f805f60608486031215610977575f80fd5b833567ffffffffffffffff81111561098d575f80fd5b610999868287016108c3565b9350506020840135915060408401356109b18161094d565b809150509250925092565b5f67ffffffffffffffff8211156109d5576109d5610847565b5060051b60200190565b5f82601f8301126109ee575f80fd5b8135610a016109fc826109bc565b610874565b8082825260208201915060208360051b860101925085831115610a22575f80fd5b602085015b83811015610a3f578035835260209283019201610a27565b5095945050505050565b5f805f60608486031215610a5b575f80fd5b833567ffffffffffffffff811115610a71575f80fd5b8401601f81018613610a81575f80fd5b8035610a8f6109fc826109bc565b8082825260208201915060208360051b850101925088831115610ab0575f80fd5b602084015b83811015610af157803567ffffffffffffffff811115610ad3575f80fd5b610ae28b6020838901016108c3565b84525060209283019201610ab5565b509550505050602084013567ffffffffffffffff811115610b10575f80fd5b610b1c868287016109df565b9250506107246040850161095a565b5f60208284031215610b3b575f80fd5b610230826106cc565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff84168152604060208201525f610bba604083018486610b44565b95945050505050565b5f82610bf6577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610c39836106cc565b16604082015273ffffffffffffffffffffffffffffffffffffffff610c60602084016106cc565b1660608201525f80604084013590508060808401525060608301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610caa575f80fd5b830160208101903567ffffffffffffffff811115610cc6575f80fd5b803603821315610cd4575f80fd5b608060a08501526105ee60c085018284610b44565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201525f610d1d60a08301866107b2565b6060830194909452509015156080909101529392505050565b5f8151808452602084019350602083015f5b82811015610d66578151865260209586019590910190600101610d48565b5093949350505050565b5f6080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b8601019250602088015f5b82811015610e01577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610dec8583516107b2565b94506020938401939190910190600101610db2565b505050508281036040840152610e178186610d36565b915050610bba606083018415159052565b5f60208284031215610e38575f80fd5b81516102308161094d565b5f82518060208501845e5f92019182525091905056fea26469706673582212202de8b3c05d226b98a1d998ad50c0d95ac2707978098d8ed21140d553a74cf0a664736f6c634300081a003360806040526002805460ff1916601217905534801561001c575f80fd5b50604051610b8d380380610b8d83398101604081905261003b916100f8565b5f61004683826101e1565b50600161005382826101e1565b50505061029b565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261007e575f80fd5b81516001600160401b038111156100975761009761005b565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c5576100c561005b565b6040528181528382016020018510156100dc575f80fd5b8160208501602083015e5f918101602001919091529392505050565b5f8060408385031215610109575f80fd5b82516001600160401b0381111561011e575f80fd5b61012a8582860161006f565b602085015190935090506001600160401b03811115610147575f80fd5b6101538582860161006f565b9150509250929050565b600181811c9082168061017157607f821691505b60208210810361018f57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101dc57805f5260205f20601f840160051c810160208510156101ba5750805b601f840160051c820191505b818110156101d9575f81556001016101c6565b50505b505050565b81516001600160401b038111156101fa576101fa61005b565b61020e81610208845461015d565b84610195565b6020601f821160018114610240575f83156102295750848201515b5f19600385901b1c1916600184901b1784556101d9565b5f84815260208120601f198516915b8281101561026f578785015182556020948501946001909201910161024f565b508482101561028c57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6108e5806102a85f395ff3fe608060405234801561000f575f80fd5b50600436106100b9575f3560e01c806340c10f191161007257806395d89b411161005857806395d89b411461017b578063a9059cbb14610183578063dd62ed3e14610196575f80fd5b806340c10f191461014957806370a082311461015c575f80fd5b806318160ddd116100a257806318160ddd146100f057806323b872dd14610107578063313ce5671461012a575f80fd5b806306fdde03146100bd578063095ea7b3146100db575b5f80fd5b6100c56101c0565b6040516100d291906106dd565b60405180910390f35b6100ee6100e9366004610758565b61024b565b005b6100f960035481565b6040519081526020016100d2565b61011a610115366004610780565b6102b7565b60405190151581526020016100d2565b6002546101379060ff1681565b60405160ff90911681526020016100d2565b6100ee610157366004610758565b61050c565b6100f961016a3660046107ba565b60046020525f908152604090205481565b6100c56105a5565b61011a610191366004610758565b6105b2565b6100f96101a43660046107da565b600560209081525f928352604080842090915290825290205481565b5f80546101cc9061080b565b80601f01602080910402602001604051908101604052809291908181526020018280546101f89061080b565b80156102435780601f1061021a57610100808354040283529160200191610243565b820191905f5260205f20905b81548152906001019060200180831161022657829003601f168201915b505050505081565b335f81815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff871680855290835292819020859055518481529192917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526004602052604081205482111561034a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e636500000000000000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff84165f9081526005602090815260408083203384529091529020548211156103e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e73756666696369656e7420616c6c6f77616e6365000000000000000000006044820152606401610341565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526004602052604081208054849290610417908490610889565b909155505073ffffffffffffffffffffffffffffffffffffffff83165f908152600460205260408120805484929061045090849061089c565b909155505073ffffffffffffffffffffffffffffffffffffffff84165f90815260056020908152604080832033845290915281208054849290610494908490610889565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516104fa91815260200190565b60405180910390a35060019392505050565b8060035f82825461051d919061089c565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f908152600460205260408120805483929061055690849061089c565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316905f907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016102ab565b600180546101cc9061080b565b335f9081526004602052604081205482111561062a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e63650000000000000000000000006044820152606401610341565b335f9081526004602052604081208054849290610648908490610889565b909155505073ffffffffffffffffffffffffffffffffffffffff83165f908152600460205260408120805484929061068190849061089c565b909155505060405182815273ffffffffffffffffffffffffffffffffffffffff84169033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35060015b92915050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610753575f80fd5b919050565b5f8060408385031215610769575f80fd5b61077283610730565b946020939093013593505050565b5f805f60608486031215610792575f80fd5b61079b84610730565b92506107a960208501610730565b929592945050506040919091013590565b5f602082840312156107ca575f80fd5b6107d382610730565b9392505050565b5f80604083850312156107eb575f80fd5b6107f483610730565b915061080260208401610730565b90509250929050565b600181811c9082168061081f57607f821691505b602082108103610856577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156106d7576106d761085c565b808201808211156106d7576106d761085c56fea26469706673582212200cfd2a3b3ac61410f384f5158efa9e45b9e8b44cb7226c0b479f845f3c50d43664736f6c634300081a0033608060405234801561000f575f80fd5b50604051610c48380380610c4883398101604081905261002e916100f4565b81818181600361003e83826101dd565b50600461004b82826101dd565b50505050505050610297565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261007a575f80fd5b81516001600160401b0381111561009357610093610057565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c1576100c1610057565b6040528181528382016020018510156100d8575f80fd5b8160208501602083015e5f918101602001919091529392505050565b5f8060408385031215610105575f80fd5b82516001600160401b0381111561011a575f80fd5b6101268582860161006b565b602085015190935090506001600160401b03811115610143575f80fd5b61014f8582860161006b565b9150509250929050565b600181811c9082168061016d57607f821691505b60208210810361018b57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101d857805f5260205f20601f840160051c810160208510156101b65750805b601f840160051c820191505b818110156101d5575f81556001016101c2565b50505b505050565b81516001600160401b038111156101f6576101f6610057565b61020a816102048454610159565b84610191565b6020601f82116001811461023c575f83156102255750848201515b5f19600385901b1c1916600184901b1784556101d5565b5f84815260208120601f198516915b8281101561026b578785015182556020948501946001909201910161024b565b508482101561028857868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6109a4806102a45f395ff3fe608060405234801561000f575f80fd5b50600436106100b9575f3560e01c806340c10f191161007257806395d89b411161005857806395d89b411461017c578063a9059cbb14610184578063dd62ed3e14610197575f80fd5b806340c10f191461013257806370a0823114610147575f80fd5b806318160ddd116100a257806318160ddd146100fe57806323b872dd14610110578063313ce56714610123575f80fd5b806306fdde03146100bd578063095ea7b3146100db575b5f80fd5b6100c56101dc565b6040516100d291906107be565b60405180910390f35b6100ee6100e9366004610839565b61026c565b60405190151581526020016100d2565b6002545b6040519081526020016100d2565b6100ee61011e366004610861565b61028b565b604051601281526020016100d2565b610145610140366004610839565b6102ae565b005b61010261015536600461089b565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b6100c56102bc565b6100ee610192366004610839565b6102cb565b6101026101a53660046108b4565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b6060600380546101eb906108e5565b80601f0160208091040260200160405190810160405280929190818152602001828054610217906108e5565b80156102625780601f1061023957610100808354040283529160200191610262565b820191905f5260205f20905b81548152906001019060200180831161024557829003601f168201915b5050505050905090565b5f808211610278575f80fd5b61028283836102e2565b90505b92915050565b5f336102988582856102ef565b6102a38585856103c1565b506001949350505050565b6102b8828261046f565b5050565b6060600480546101eb906108e5565b5f336102d88185856103c1565b5060019392505050565b5f336102d88185856104c9565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bb57818110156103ad576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bb84848484035f6104d2565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610410576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016103a4565b73ffffffffffffffffffffffffffffffffffffffff821661045f576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016103a4565b61046a838383610617565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166104be576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016103a4565b6102b85f8383610617565b61046a83838360015b73ffffffffffffffffffffffffffffffffffffffff8416610521576040517fe602df050000000000000000000000000000000000000000000000000000000081525f60048201526024016103a4565b73ffffffffffffffffffffffffffffffffffffffff8316610570576040517f94280d620000000000000000000000000000000000000000000000000000000081525f60048201526024016103a4565b73ffffffffffffffffffffffffffffffffffffffff8085165f90815260016020908152604080832093871683529290522082905580156103bb578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060991815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064e578060025f8282546106439190610936565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260208190526040902054818110156106d3576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a4565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610752565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b191815260200190565b60405180910390a3505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610834575f80fd5b919050565b5f806040838503121561084a575f80fd5b61085383610811565b946020939093013593505050565b5f805f60608486031215610873575f80fd5b61087c84610811565b925061088a60208501610811565b929592945050506040919091013590565b5f602082840312156108ab575f80fd5b61028282610811565b5f80604083850312156108c5575f80fd5b6108ce83610811565b91506108dc60208401610811565b90509250929050565b600181811c908216806108f957607f821691505b602082108103610930577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b80820180821115610285577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea2646970667358221220791da0190b5ac204bcf523447fbcff4426806bdcf439d40982804ed6b0e7a3f464736f6c634300081a0033a2646970667358221220787864b44706e381fbde8948f84cc37cd611d2ceb0adcf5f9246d1b37447ef8964736f6c634300081a0033",
}

// GatewayEVMTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMTestMetaData.ABI instead.
var GatewayEVMTestABI = GatewayEVMTestMetaData.ABI

// GatewayEVMTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMTestMetaData.Bin instead.
var GatewayEVMTestBin = GatewayEVMTestMetaData.Bin

// DeployGatewayEVMTest deploys a new Ethereum contract, binding an instance of GatewayEVMTest to it.
func DeployGatewayEVMTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVMTest, error) {
	parsed, err := GatewayEVMTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVMTest{GatewayEVMTestCaller: GatewayEVMTestCaller{contract: contract}, GatewayEVMTestTransactor: GatewayEVMTestTransactor{contract: contract}, GatewayEVMTestFilterer: GatewayEVMTestFilterer{contract: contract}}, nil
}

// GatewayEVMTest is an auto generated Go binding around an Ethereum contract.
type GatewayEVMTest struct {
	GatewayEVMTestCaller     // Read-only binding to the contract
	GatewayEVMTestTransactor // Write-only binding to the contract
	GatewayEVMTestFilterer   // Log filterer for contract events
}

// GatewayEVMTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMTestSession struct {
	Contract     *GatewayEVMTest   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayEVMTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMTestCallerSession struct {
	Contract *GatewayEVMTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GatewayEVMTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMTestTransactorSession struct {
	Contract     *GatewayEVMTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GatewayEVMTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMTestRaw struct {
	Contract *GatewayEVMTest // Generic contract binding to access the raw methods on
}

// GatewayEVMTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMTestCallerRaw struct {
	Contract *GatewayEVMTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMTestTransactorRaw struct {
	Contract *GatewayEVMTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVMTest creates a new instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTest(address common.Address, backend bind.ContractBackend) (*GatewayEVMTest, error) {
	contract, err := bindGatewayEVMTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTest{GatewayEVMTestCaller: GatewayEVMTestCaller{contract: contract}, GatewayEVMTestTransactor: GatewayEVMTestTransactor{contract: contract}, GatewayEVMTestFilterer: GatewayEVMTestFilterer{contract: contract}}, nil
}

// NewGatewayEVMTestCaller creates a new read-only instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMTestCaller, error) {
	contract, err := bindGatewayEVMTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestCaller{contract: contract}, nil
}

// NewGatewayEVMTestTransactor creates a new write-only instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMTestTransactor, error) {
	contract, err := bindGatewayEVMTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestTransactor{contract: contract}, nil
}

// NewGatewayEVMTestFilterer creates a new log filterer instance of GatewayEVMTest, bound to a specific deployed contract.
func NewGatewayEVMTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMTestFilterer, error) {
	contract, err := bindGatewayEVMTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestFilterer{contract: contract}, nil
}

// bindGatewayEVMTest binds a generic wrapper to an already deployed contract.
func bindGatewayEVMTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMTest *GatewayEVMTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMTest.Contract.GatewayEVMTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMTest *GatewayEVMTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.GatewayEVMTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMTest *GatewayEVMTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.GatewayEVMTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMTest *GatewayEVMTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMTest *GatewayEVMTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMTest *GatewayEVMTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.contract.Transact(opts, method, params...)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) ASSETHANDLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "ASSET_HANDLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.ASSETHANDLERROLE(&_GatewayEVMTest.CallOpts)
}

// ASSETHANDLERROLE is a free data retrieval call binding the contract method 0x5d62c860.
//
// Solidity: function ASSET_HANDLER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ASSETHANDLERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.ASSETHANDLERROLE(&_GatewayEVMTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.DEFAULTADMINROLE(&_GatewayEVMTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.DEFAULTADMINROLE(&_GatewayEVMTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestSession) ISTEST() (bool, error) {
	return _GatewayEVMTest.Contract.ISTEST(&_GatewayEVMTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ISTEST() (bool, error) {
	return _GatewayEVMTest.Contract.ISTEST(&_GatewayEVMTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.PAUSERROLE(&_GatewayEVMTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.PAUSERROLE(&_GatewayEVMTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) TSSROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.TSSROLE(&_GatewayEVMTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TSSROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.TSSROLE(&_GatewayEVMTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestSession) WITHDRAWERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.WITHDRAWERROLE(&_GatewayEVMTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _GatewayEVMTest.Contract.WITHDRAWERROLE(&_GatewayEVMTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.ExcludeArtifacts(&_GatewayEVMTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.ExcludeArtifacts(&_GatewayEVMTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeContracts(&_GatewayEVMTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeContracts(&_GatewayEVMTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.ExcludeSelectors(&_GatewayEVMTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.ExcludeSelectors(&_GatewayEVMTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMTest *GatewayEVMTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeSenders(&_GatewayEVMTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.ExcludeSenders(&_GatewayEVMTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestSession) Failed() (bool, error) {
	return _GatewayEVMTest.Contract.Failed(&_GatewayEVMTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) Failed() (bool, error) {
	return _GatewayEVMTest.Contract.Failed(&_GatewayEVMTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMTest.Contract.TargetArtifactSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMTest.Contract.TargetArtifactSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.TargetArtifacts(&_GatewayEVMTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMTest.Contract.TargetArtifacts(&_GatewayEVMTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetContracts(&_GatewayEVMTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetContracts(&_GatewayEVMTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMTest.Contract.TargetInterfaces(&_GatewayEVMTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMTest.Contract.TargetInterfaces(&_GatewayEVMTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.TargetSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMTest.Contract.TargetSelectors(&_GatewayEVMTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMTest *GatewayEVMTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetSenders(&_GatewayEVMTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMTest *GatewayEVMTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMTest.Contract.TargetSenders(&_GatewayEVMTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.SetUp(&_GatewayEVMTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.SetUp(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x04b974f9.
//
// Solidity: function testExecuteFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteFailsIfDestinationIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteFailsIfDestinationIsZeroAddress")
}

// TestExecuteFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x04b974f9.
//
// Solidity: function testExecuteFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteFailsIfDestinationIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x04b974f9.
//
// Solidity: function testExecuteFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteFailsIfDestinationIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevert is a paid mutator transaction binding the contract method 0xfa18c09b.
//
// Solidity: function testExecuteRevert() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteRevert(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteRevert")
}

// TestExecuteRevert is a paid mutator transaction binding the contract method 0xfa18c09b.
//
// Solidity: function testExecuteRevert() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteRevert() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevert(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevert is a paid mutator transaction binding the contract method 0xfa18c09b.
//
// Solidity: function testExecuteRevert() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteRevert() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevert(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51010e49.
//
// Solidity: function testExecuteRevertFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteRevertFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteRevertFailsIfReceiverIsZeroAddress")
}

// TestExecuteRevertFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51010e49.
//
// Solidity: function testExecuteRevertFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteRevertFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfReceiverIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x51010e49.
//
// Solidity: function testExecuteRevertFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteRevertFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfReceiverIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0xcebad2a6.
//
// Solidity: function testExecuteRevertFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteRevertFailsIfSenderIsNotTSS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteRevertFailsIfSenderIsNotTSS")
}

// TestExecuteRevertFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0xcebad2a6.
//
// Solidity: function testExecuteRevertFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteRevertFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteRevertFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0xcebad2a6.
//
// Solidity: function testExecuteRevertFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteRevertFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteRevertFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xe6afc790.
//
// Solidity: function testExecuteWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteWithERC20FailsIfNotCustodyOrConnector(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteWithERC20FailsIfNotCustodyOrConnector")
}

// TestExecuteWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xe6afc790.
//
// Solidity: function testExecuteWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xe6afc790.
//
// Solidity: function testExecuteWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x6bdd212b.
//
// Solidity: function testExecuteWithMsgContextFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteWithMsgContextFailsIfDestinationIsZeroAddress")
}

// TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x6bdd212b.
//
// Solidity: function testExecuteWithMsgContextFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x6bdd212b.
//
// Solidity: function testExecuteWithMsgContextFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithMsgContextFailsIfDestinationIsZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithNonReturnApprovalToken is a paid mutator transaction binding the contract method 0x3801e58f.
//
// Solidity: function testExecuteWithNonReturnApprovalToken() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteWithNonReturnApprovalToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteWithNonReturnApprovalToken")
}

// TestExecuteWithNonReturnApprovalToken is a paid mutator transaction binding the contract method 0x3801e58f.
//
// Solidity: function testExecuteWithNonReturnApprovalToken() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteWithNonReturnApprovalToken() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithNonReturnApprovalToken(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithNonReturnApprovalToken is a paid mutator transaction binding the contract method 0x3801e58f.
//
// Solidity: function testExecuteWithNonReturnApprovalToken() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteWithNonReturnApprovalToken() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithNonReturnApprovalToken(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithTokenRevertingOnZeroApproval is a paid mutator transaction binding the contract method 0xe5418219.
//
// Solidity: function testExecuteWithTokenRevertingOnZeroApproval() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestExecuteWithTokenRevertingOnZeroApproval(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testExecuteWithTokenRevertingOnZeroApproval")
}

// TestExecuteWithTokenRevertingOnZeroApproval is a paid mutator transaction binding the contract method 0xe5418219.
//
// Solidity: function testExecuteWithTokenRevertingOnZeroApproval() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestExecuteWithTokenRevertingOnZeroApproval() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithTokenRevertingOnZeroApproval(&_GatewayEVMTest.TransactOpts)
}

// TestExecuteWithTokenRevertingOnZeroApproval is a paid mutator transaction binding the contract method 0xe5418219.
//
// Solidity: function testExecuteWithTokenRevertingOnZeroApproval() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestExecuteWithTokenRevertingOnZeroApproval() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestExecuteWithTokenRevertingOnZeroApproval(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParams is a paid mutator transaction binding the contract method 0x44671b94.
//
// Solidity: function testForwardCallToReceiveNoParams() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNoParams(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNoParams")
}

// TestForwardCallToReceiveNoParams is a paid mutator transaction binding the contract method 0x44671b94.
//
// Solidity: function testForwardCallToReceiveNoParams() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNoParams() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParams(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParams is a paid mutator transaction binding the contract method 0x44671b94.
//
// Solidity: function testForwardCallToReceiveNoParams() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNoParams() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParams(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParamsTogglePause is a paid mutator transaction binding the contract method 0xdd51e82f.
//
// Solidity: function testForwardCallToReceiveNoParamsTogglePause() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNoParamsTogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNoParamsTogglePause")
}

// TestForwardCallToReceiveNoParamsTogglePause is a paid mutator transaction binding the contract method 0xdd51e82f.
//
// Solidity: function testForwardCallToReceiveNoParamsTogglePause() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNoParamsTogglePause() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParamsTogglePause(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNoParamsTogglePause is a paid mutator transaction binding the contract method 0xdd51e82f.
//
// Solidity: function testForwardCallToReceiveNoParamsTogglePause() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNoParamsTogglePause() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNoParamsTogglePause(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayable is a paid mutator transaction binding the contract method 0xf68bd1c0.
//
// Solidity: function testForwardCallToReceiveNonPayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNonPayable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNonPayable")
}

// TestForwardCallToReceiveNonPayable is a paid mutator transaction binding the contract method 0xf68bd1c0.
//
// Solidity: function testForwardCallToReceiveNonPayable() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNonPayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayable(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayable is a paid mutator transaction binding the contract method 0xf68bd1c0.
//
// Solidity: function testForwardCallToReceiveNonPayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNonPayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayable(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x7d7f772a.
//
// Solidity: function testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS")
}

// TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x7d7f772a.
//
// Solidity: function testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x7d7f772a.
//
// Solidity: function testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x43fd8c7d.
//
// Solidity: function testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS")
}

// TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x43fd8c7d.
//
// Solidity: function testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS is a paid mutator transaction binding the contract method 0x43fd8c7d.
//
// Solidity: function testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnCallFails is a paid mutator transaction binding the contract method 0xd38b66cd.
//
// Solidity: function testForwardCallToReceiveOnCallFails() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveOnCallFails(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveOnCallFails")
}

// TestForwardCallToReceiveOnCallFails is a paid mutator transaction binding the contract method 0xd38b66cd.
//
// Solidity: function testForwardCallToReceiveOnCallFails() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveOnCallFails() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnCallFails(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnCallFails is a paid mutator transaction binding the contract method 0xd38b66cd.
//
// Solidity: function testForwardCallToReceiveOnCallFails() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveOnCallFails() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnCallFails(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnCallUsingAuthCall is a paid mutator transaction binding the contract method 0xa397ffd2.
//
// Solidity: function testForwardCallToReceiveOnCallUsingAuthCall() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveOnCallUsingAuthCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveOnCallUsingAuthCall")
}

// TestForwardCallToReceiveOnCallUsingAuthCall is a paid mutator transaction binding the contract method 0xa397ffd2.
//
// Solidity: function testForwardCallToReceiveOnCallUsingAuthCall() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveOnCallUsingAuthCall() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnCallUsingAuthCall(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnCallUsingAuthCall is a paid mutator transaction binding the contract method 0xa397ffd2.
//
// Solidity: function testForwardCallToReceiveOnCallUsingAuthCall() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveOnCallUsingAuthCall() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnCallUsingAuthCall(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnRevertFails is a paid mutator transaction binding the contract method 0x130daf59.
//
// Solidity: function testForwardCallToReceiveOnRevertFails() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceiveOnRevertFails(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceiveOnRevertFails")
}

// TestForwardCallToReceiveOnRevertFails is a paid mutator transaction binding the contract method 0x130daf59.
//
// Solidity: function testForwardCallToReceiveOnRevertFails() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceiveOnRevertFails() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnRevertFails(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceiveOnRevertFails is a paid mutator transaction binding the contract method 0x130daf59.
//
// Solidity: function testForwardCallToReceiveOnRevertFails() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceiveOnRevertFails() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceiveOnRevertFails(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0xfe7bdbb2.
//
// Solidity: function testForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestForwardCallToReceivePayable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testForwardCallToReceivePayable")
}

// TestForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0xfe7bdbb2.
//
// Solidity: function testForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestForwardCallToReceivePayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceivePayable(&_GatewayEVMTest.TransactOpts)
}

// TestForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0xfe7bdbb2.
//
// Solidity: function testForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestForwardCallToReceivePayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestForwardCallToReceivePayable(&_GatewayEVMTest.TransactOpts)
}

// TestRevertWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xb124dbf5.
//
// Solidity: function testRevertWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestRevertWithERC20FailsIfNotCustodyOrConnector(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testRevertWithERC20FailsIfNotCustodyOrConnector")
}

// TestRevertWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xb124dbf5.
//
// Solidity: function testRevertWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestRevertWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestRevertWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestRevertWithERC20FailsIfNotCustodyOrConnector is a paid mutator transaction binding the contract method 0xb124dbf5.
//
// Solidity: function testRevertWithERC20FailsIfNotCustodyOrConnector() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestRevertWithERC20FailsIfNotCustodyOrConnector() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestRevertWithERC20FailsIfNotCustodyOrConnector(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xccf20616.
//
// Solidity: function testSetConnectorFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetConnectorFailsIfSenderIsNotAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetConnectorFailsIfSenderIsNotAdmin")
}

// TestSetConnectorFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xccf20616.
//
// Solidity: function testSetConnectorFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetConnectorFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xccf20616.
//
// Solidity: function testSetConnectorFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetConnectorFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSet is a paid mutator transaction binding the contract method 0x7ebf744f.
//
// Solidity: function testSetConnectorFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetConnectorFailsIfSet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetConnectorFailsIfSet")
}

// TestSetConnectorFailsIfSet is a paid mutator transaction binding the contract method 0x7ebf744f.
//
// Solidity: function testSetConnectorFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetConnectorFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfSet is a paid mutator transaction binding the contract method 0x7ebf744f.
//
// Solidity: function testSetConnectorFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetConnectorFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfZero is a paid mutator transaction binding the contract method 0xce5871e1.
//
// Solidity: function testSetConnectorFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetConnectorFailsIfZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetConnectorFailsIfZero")
}

// TestSetConnectorFailsIfZero is a paid mutator transaction binding the contract method 0xce5871e1.
//
// Solidity: function testSetConnectorFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetConnectorFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestSetConnectorFailsIfZero is a paid mutator transaction binding the contract method 0xce5871e1.
//
// Solidity: function testSetConnectorFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetConnectorFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetConnectorFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xa56f7a4b.
//
// Solidity: function testSetCustodyFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetCustodyFailsIfSenderIsNotAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetCustodyFailsIfSenderIsNotAdmin")
}

// TestSetCustodyFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xa56f7a4b.
//
// Solidity: function testSetCustodyFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetCustodyFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSenderIsNotAdmin is a paid mutator transaction binding the contract method 0xa56f7a4b.
//
// Solidity: function testSetCustodyFailsIfSenderIsNotAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetCustodyFailsIfSenderIsNotAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSenderIsNotAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSet is a paid mutator transaction binding the contract method 0x1855c337.
//
// Solidity: function testSetCustodyFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetCustodyFailsIfSet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetCustodyFailsIfSet")
}

// TestSetCustodyFailsIfSet is a paid mutator transaction binding the contract method 0x1855c337.
//
// Solidity: function testSetCustodyFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetCustodyFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfSet is a paid mutator transaction binding the contract method 0x1855c337.
//
// Solidity: function testSetCustodyFailsIfSet() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetCustodyFailsIfSet() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfSet(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfZero is a paid mutator transaction binding the contract method 0x1226c655.
//
// Solidity: function testSetCustodyFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetCustodyFailsIfZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetCustodyFailsIfZero")
}

// TestSetCustodyFailsIfZero is a paid mutator transaction binding the contract method 0x1226c655.
//
// Solidity: function testSetCustodyFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetCustodyFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestSetCustodyFailsIfZero is a paid mutator transaction binding the contract method 0x1226c655.
//
// Solidity: function testSetCustodyFailsIfZero() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetCustodyFailsIfZero() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetCustodyFailsIfZero(&_GatewayEVMTest.TransactOpts)
}

// TestSetDepositAllowedAssetOnlyAdmin is a paid mutator transaction binding the contract method 0xdf23745a.
//
// Solidity: function testSetDepositAllowedAssetOnlyAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetDepositAllowedAssetOnlyAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetDepositAllowedAssetOnlyAdmin")
}

// TestSetDepositAllowedAssetOnlyAdmin is a paid mutator transaction binding the contract method 0xdf23745a.
//
// Solidity: function testSetDepositAllowedAssetOnlyAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetDepositAllowedAssetOnlyAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetDepositAllowedAssetOnlyAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetDepositAllowedAssetOnlyAdmin is a paid mutator transaction binding the contract method 0xdf23745a.
//
// Solidity: function testSetDepositAllowedAssetOnlyAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetDepositAllowedAssetOnlyAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetDepositAllowedAssetOnlyAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetDepositPausedOnlyAdmin is a paid mutator transaction binding the contract method 0x0ab10776.
//
// Solidity: function testSetDepositPausedOnlyAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestSetDepositPausedOnlyAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testSetDepositPausedOnlyAdmin")
}

// TestSetDepositPausedOnlyAdmin is a paid mutator transaction binding the contract method 0x0ab10776.
//
// Solidity: function testSetDepositPausedOnlyAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestSetDepositPausedOnlyAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetDepositPausedOnlyAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestSetDepositPausedOnlyAdmin is a paid mutator transaction binding the contract method 0x0ab10776.
//
// Solidity: function testSetDepositPausedOnlyAdmin() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestSetDepositPausedOnlyAdmin() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestSetDepositPausedOnlyAdmin(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestTSSUpgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testTSSUpgrade")
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestTSSUpgrade() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgrade(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgrade is a paid mutator transaction binding the contract method 0x52ff5939.
//
// Solidity: function testTSSUpgrade() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestTSSUpgrade() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgrade(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testTSSUpgradeFailsIfSenderIsNotTSSUpdater")
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfSenderIsNotTSSUpdater is a paid mutator transaction binding the contract method 0x070f2ad0.
//
// Solidity: function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestTSSUpgradeFailsIfSenderIsNotTSSUpdater() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfSenderIsNotTSSUpdater(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestTSSUpgradeFailsIfZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testTSSUpgradeFailsIfZeroAddress")
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestTSSUpgradeFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestTSSUpgradeFailsIfZeroAddress is a paid mutator transaction binding the contract method 0x4df42da1.
//
// Solidity: function testTSSUpgradeFailsIfZeroAddress() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestTSSUpgradeFailsIfZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestTSSUpgradeFailsIfZeroAddress(&_GatewayEVMTest.TransactOpts)
}

// TestUpgradeAndForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0x7a380ebf.
//
// Solidity: function testUpgradeAndForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestUpgradeAndForwardCallToReceivePayable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testUpgradeAndForwardCallToReceivePayable")
}

// TestUpgradeAndForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0x7a380ebf.
//
// Solidity: function testUpgradeAndForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestUpgradeAndForwardCallToReceivePayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestUpgradeAndForwardCallToReceivePayable(&_GatewayEVMTest.TransactOpts)
}

// TestUpgradeAndForwardCallToReceivePayable is a paid mutator transaction binding the contract method 0x7a380ebf.
//
// Solidity: function testUpgradeAndForwardCallToReceivePayable() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestUpgradeAndForwardCallToReceivePayable() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestUpgradeAndForwardCallToReceivePayable(&_GatewayEVMTest.TransactOpts)
}

// TestUpgradePreservesDepositPauseFunctions is a paid mutator transaction binding the contract method 0x76b6f88e.
//
// Solidity: function testUpgradePreservesDepositPauseFunctions() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactor) TestUpgradePreservesDepositPauseFunctions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMTest.contract.Transact(opts, "testUpgradePreservesDepositPauseFunctions")
}

// TestUpgradePreservesDepositPauseFunctions is a paid mutator transaction binding the contract method 0x76b6f88e.
//
// Solidity: function testUpgradePreservesDepositPauseFunctions() returns()
func (_GatewayEVMTest *GatewayEVMTestSession) TestUpgradePreservesDepositPauseFunctions() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestUpgradePreservesDepositPauseFunctions(&_GatewayEVMTest.TransactOpts)
}

// TestUpgradePreservesDepositPauseFunctions is a paid mutator transaction binding the contract method 0x76b6f88e.
//
// Solidity: function testUpgradePreservesDepositPauseFunctions() returns()
func (_GatewayEVMTest *GatewayEVMTestTransactorSession) TestUpgradePreservesDepositPauseFunctions() (*types.Transaction, error) {
	return _GatewayEVMTest.Contract.TestUpgradePreservesDepositPauseFunctions(&_GatewayEVMTest.TransactOpts)
}

// GatewayEVMTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayEVMTest contract.
type GatewayEVMTestCalledIterator struct {
	Event *GatewayEVMTestCalled // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestCalled)
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
		it.Event = new(GatewayEVMTestCalled)
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
func (it *GatewayEVMTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestCalled represents a Called event raised by the GatewayEVMTest contract.
type GatewayEVMTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestCalledIterator{contract: _GatewayEVMTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestCalled)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Called", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseCalled(log types.Log) (*GatewayEVMTestCalled, error) {
	event := new(GatewayEVMTestCalled)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GatewayEVMTest contract.
type GatewayEVMTestDepositedIterator struct {
	Event *GatewayEVMTestDeposited // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestDeposited)
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
		it.Event = new(GatewayEVMTestDeposited)
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
func (it *GatewayEVMTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestDeposited represents a Deposited event raised by the GatewayEVMTest contract.
type GatewayEVMTestDeposited struct {
	Recipient []byte
	Asset     common.Address
	Amount    *big.Int
	Message   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterDeposited(opts *bind.FilterOpts, asset []common.Address) (*GatewayEVMTestDepositedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestDepositedIterator{contract: _GatewayEVMTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestDeposited, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestDeposited)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseDeposited(log types.Log) (*GatewayEVMTestDeposited, error) {
	event := new(GatewayEVMTestDeposited)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestDeposited0Iterator is returned from FilterDeposited0 and is used to iterate over the raw logs and unpacked data for Deposited0 events raised by the GatewayEVMTest contract.
type GatewayEVMTestDeposited0Iterator struct {
	Event *GatewayEVMTestDeposited0 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestDeposited0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestDeposited0)
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
		it.Event = new(GatewayEVMTestDeposited0)
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
func (it *GatewayEVMTestDeposited0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestDeposited0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestDeposited0 represents a Deposited0 event raised by the GatewayEVMTest contract.
type GatewayEVMTestDeposited0 struct {
	Sender        common.Address
	Receiver      common.Address
	Amount        *big.Int
	Asset         common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposited0 is a free log retrieval operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterDeposited0(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMTestDeposited0Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Deposited0", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestDeposited0Iterator{contract: _GatewayEVMTest.contract, event: "Deposited0", logs: logs, sub: sub}, nil
}

// WatchDeposited0 is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchDeposited0(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestDeposited0, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Deposited0", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestDeposited0)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited0", log); err != nil {
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

// ParseDeposited0 is a log parse operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseDeposited0(log types.Log) (*GatewayEVMTestDeposited0, error) {
	event := new(GatewayEVMTestDeposited0)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Deposited0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestDepositedAndCalledIterator is returned from FilterDepositedAndCalled and is used to iterate over the raw logs and unpacked data for DepositedAndCalled events raised by the GatewayEVMTest contract.
type GatewayEVMTestDepositedAndCalledIterator struct {
	Event *GatewayEVMTestDepositedAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestDepositedAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestDepositedAndCalled)
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
		it.Event = new(GatewayEVMTestDepositedAndCalled)
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
func (it *GatewayEVMTestDepositedAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestDepositedAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestDepositedAndCalled represents a DepositedAndCalled event raised by the GatewayEVMTest contract.
type GatewayEVMTestDepositedAndCalled struct {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterDepositedAndCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMTestDepositedAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestDepositedAndCalledIterator{contract: _GatewayEVMTest.contract, event: "DepositedAndCalled", logs: logs, sub: sub}, nil
}

// WatchDepositedAndCalled is a free log subscription operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchDepositedAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestDepositedAndCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestDepositedAndCalled)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseDepositedAndCalled(log types.Log) (*GatewayEVMTestDepositedAndCalled, error) {
	event := new(GatewayEVMTestDepositedAndCalled)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedIterator struct {
	Event *GatewayEVMTestExecuted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestExecuted)
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
		it.Event = new(GatewayEVMTestExecuted)
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
func (it *GatewayEVMTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestExecuted represents a Executed event raised by the GatewayEVMTest contract.
type GatewayEVMTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestExecutedIterator{contract: _GatewayEVMTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestExecuted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Executed", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseExecuted(log types.Log) (*GatewayEVMTestExecuted, error) {
	event := new(GatewayEVMTestExecuted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestExecutedV2Iterator is returned from FilterExecutedV2 and is used to iterate over the raw logs and unpacked data for ExecutedV2 events raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedV2Iterator struct {
	Event *GatewayEVMTestExecutedV2 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestExecutedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestExecutedV2)
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
		it.Event = new(GatewayEVMTestExecutedV2)
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
func (it *GatewayEVMTestExecutedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestExecutedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestExecutedV2 represents a ExecutedV2 event raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedV2 struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutedV2 is a free log retrieval operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterExecutedV2(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMTestExecutedV2Iterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ExecutedV2", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestExecutedV2Iterator{contract: _GatewayEVMTest.contract, event: "ExecutedV2", logs: logs, sub: sub}, nil
}

// WatchExecutedV2 is a free log subscription operation binding the contract event 0x373df382b9c587826f3de13f16d67f8d99f28ee947fc0924c6ef2d6d2c7e8546.
//
// Solidity: event ExecutedV2(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchExecutedV2(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestExecutedV2, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ExecutedV2", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestExecutedV2)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ExecutedV2", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseExecutedV2(log types.Log) (*GatewayEVMTestExecutedV2, error) {
	event := new(GatewayEVMTestExecutedV2)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ExecutedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedWithERC20Iterator struct {
	Event *GatewayEVMTestExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestExecutedWithERC20)
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
		it.Event = new(GatewayEVMTestExecutedWithERC20)
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
func (it *GatewayEVMTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVMTest contract.
type GatewayEVMTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestExecutedWithERC20Iterator{contract: _GatewayEVMTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestExecutedWithERC20)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMTestExecutedWithERC20, error) {
	event := new(GatewayEVMTestExecutedWithERC20)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedERC20Iterator struct {
	Event *GatewayEVMTestReceivedERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedERC20)
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
		it.Event = new(GatewayEVMTestReceivedERC20)
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
func (it *GatewayEVMTestReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedERC20 represents a ReceivedERC20 event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*GatewayEVMTestReceivedERC20Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedERC20Iterator{contract: _GatewayEVMTest.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedERC20)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedERC20(log types.Log) (*GatewayEVMTestReceivedERC20, error) {
	event := new(GatewayEVMTestReceivedERC20)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNoParamsIterator struct {
	Event *GatewayEVMTestReceivedNoParams // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedNoParams)
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
		it.Event = new(GatewayEVMTestReceivedNoParams)
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
func (it *GatewayEVMTestReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedNoParams represents a ReceivedNoParams event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*GatewayEVMTestReceivedNoParamsIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedNoParamsIterator{contract: _GatewayEVMTest.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedNoParams)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedNoParams(log types.Log) (*GatewayEVMTestReceivedNoParams, error) {
	event := new(GatewayEVMTestReceivedNoParams)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNonPayableIterator struct {
	Event *GatewayEVMTestReceivedNonPayable // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedNonPayable)
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
		it.Event = new(GatewayEVMTestReceivedNonPayable)
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
func (it *GatewayEVMTestReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedNonPayable represents a ReceivedNonPayable event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*GatewayEVMTestReceivedNonPayableIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedNonPayableIterator{contract: _GatewayEVMTest.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedNonPayable)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedNonPayable(log types.Log) (*GatewayEVMTestReceivedNonPayable, error) {
	event := new(GatewayEVMTestReceivedNonPayable)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedOnCallIterator is returned from FilterReceivedOnCall and is used to iterate over the raw logs and unpacked data for ReceivedOnCall events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedOnCallIterator struct {
	Event *GatewayEVMTestReceivedOnCall // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestReceivedOnCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedOnCall)
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
		it.Event = new(GatewayEVMTestReceivedOnCall)
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
func (it *GatewayEVMTestReceivedOnCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedOnCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedOnCall represents a ReceivedOnCall event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedOnCall struct {
	Sender  common.Address
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceivedOnCall is a free log retrieval operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedOnCall(opts *bind.FilterOpts) (*GatewayEVMTestReceivedOnCallIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedOnCallIterator{contract: _GatewayEVMTest.contract, event: "ReceivedOnCall", logs: logs, sub: sub}, nil
}

// WatchReceivedOnCall is a free log subscription operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedOnCall(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedOnCall) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedOnCall)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedOnCall(log types.Log) (*GatewayEVMTestReceivedOnCall, error) {
	event := new(GatewayEVMTestReceivedOnCall)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedPayableIterator struct {
	Event *GatewayEVMTestReceivedPayable // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedPayable)
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
		it.Event = new(GatewayEVMTestReceivedPayable)
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
func (it *GatewayEVMTestReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedPayable represents a ReceivedPayable event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedPayable struct {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*GatewayEVMTestReceivedPayableIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedPayableIterator{contract: _GatewayEVMTest.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedPayable)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedPayable(log types.Log) (*GatewayEVMTestReceivedPayable, error) {
	event := new(GatewayEVMTestReceivedPayable)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedRevertIterator struct {
	Event *GatewayEVMTestReceivedRevert // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReceivedRevert)
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
		it.Event = new(GatewayEVMTestReceivedRevert)
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
func (it *GatewayEVMTestReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReceivedRevert represents a ReceivedRevert event raised by the GatewayEVMTest contract.
type GatewayEVMTestReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*GatewayEVMTestReceivedRevertIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestReceivedRevertIterator{contract: _GatewayEVMTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReceivedRevert)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReceivedRevert(log types.Log) (*GatewayEVMTestReceivedRevert, error) {
	event := new(GatewayEVMTestReceivedRevert)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the GatewayEVMTest contract.
type GatewayEVMTestRevertedIterator struct {
	Event *GatewayEVMTestReverted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestReverted)
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
		it.Event = new(GatewayEVMTestReverted)
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
func (it *GatewayEVMTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestReverted represents a Reverted event raised by the GatewayEVMTest contract.
type GatewayEVMTestReverted struct {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestRevertedIterator{contract: _GatewayEVMTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestReverted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Reverted", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseReverted(log types.Log) (*GatewayEVMTestReverted, error) {
	event := new(GatewayEVMTestReverted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestStrandedFundsRefundedIterator is returned from FilterStrandedFundsRefunded and is used to iterate over the raw logs and unpacked data for StrandedFundsRefunded events raised by the GatewayEVMTest contract.
type GatewayEVMTestStrandedFundsRefundedIterator struct {
	Event *GatewayEVMTestStrandedFundsRefunded // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestStrandedFundsRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestStrandedFundsRefunded)
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
		it.Event = new(GatewayEVMTestStrandedFundsRefunded)
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
func (it *GatewayEVMTestStrandedFundsRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestStrandedFundsRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestStrandedFundsRefunded represents a StrandedFundsRefunded event raised by the GatewayEVMTest contract.
type GatewayEVMTestStrandedFundsRefunded struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStrandedFundsRefunded is a free log retrieval operation binding the contract event 0x7b73729b322c3c294385b7454f9895d7975135d24a54310b0e4542962bd089ab.
//
// Solidity: event StrandedFundsRefunded(address indexed to, address indexed token, uint256 amount)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterStrandedFundsRefunded(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestStrandedFundsRefundedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "StrandedFundsRefunded", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestStrandedFundsRefundedIterator{contract: _GatewayEVMTest.contract, event: "StrandedFundsRefunded", logs: logs, sub: sub}, nil
}

// WatchStrandedFundsRefunded is a free log subscription operation binding the contract event 0x7b73729b322c3c294385b7454f9895d7975135d24a54310b0e4542962bd089ab.
//
// Solidity: event StrandedFundsRefunded(address indexed to, address indexed token, uint256 amount)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchStrandedFundsRefunded(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestStrandedFundsRefunded, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "StrandedFundsRefunded", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestStrandedFundsRefunded)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "StrandedFundsRefunded", log); err != nil {
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

// ParseStrandedFundsRefunded is a log parse operation binding the contract event 0x7b73729b322c3c294385b7454f9895d7975135d24a54310b0e4542962bd089ab.
//
// Solidity: event StrandedFundsRefunded(address indexed to, address indexed token, uint256 amount)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseStrandedFundsRefunded(log types.Log) (*GatewayEVMTestStrandedFundsRefunded, error) {
	event := new(GatewayEVMTestStrandedFundsRefunded)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "StrandedFundsRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUnwhitelistedIterator is returned from FilterUnwhitelisted and is used to iterate over the raw logs and unpacked data for Unwhitelisted events raised by the GatewayEVMTest contract.
type GatewayEVMTestUnwhitelistedIterator struct {
	Event *GatewayEVMTestUnwhitelisted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestUnwhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUnwhitelisted)
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
		it.Event = new(GatewayEVMTestUnwhitelisted)
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
func (it *GatewayEVMTestUnwhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUnwhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUnwhitelisted represents a Unwhitelisted event raised by the GatewayEVMTest contract.
type GatewayEVMTestUnwhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnwhitelisted is a free log retrieval operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUnwhitelisted(opts *bind.FilterOpts, token []common.Address) (*GatewayEVMTestUnwhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUnwhitelistedIterator{contract: _GatewayEVMTest.contract, event: "Unwhitelisted", logs: logs, sub: sub}, nil
}

// WatchUnwhitelisted is a free log subscription operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUnwhitelisted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUnwhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUnwhitelisted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
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

// ParseUnwhitelisted is a log parse operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUnwhitelisted(log types.Log) (*GatewayEVMTestUnwhitelisted, error) {
	event := new(GatewayEVMTestUnwhitelisted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUpdatedAdditionalActionFeeIterator is returned from FilterUpdatedAdditionalActionFee and is used to iterate over the raw logs and unpacked data for UpdatedAdditionalActionFee events raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedAdditionalActionFeeIterator struct {
	Event *GatewayEVMTestUpdatedAdditionalActionFee // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestUpdatedAdditionalActionFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUpdatedAdditionalActionFee)
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
		it.Event = new(GatewayEVMTestUpdatedAdditionalActionFee)
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
func (it *GatewayEVMTestUpdatedAdditionalActionFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUpdatedAdditionalActionFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUpdatedAdditionalActionFee represents a UpdatedAdditionalActionFee event raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedAdditionalActionFee struct {
	OldFeeWei *big.Int
	NewFeeWei *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAdditionalActionFee is a free log retrieval operation binding the contract event 0xe60a2882aff7841a7d0492b804d46ff2d9b07ea353dbc3791e8276b729e40cd8.
//
// Solidity: event UpdatedAdditionalActionFee(uint256 oldFeeWei, uint256 newFeeWei)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUpdatedAdditionalActionFee(opts *bind.FilterOpts) (*GatewayEVMTestUpdatedAdditionalActionFeeIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "UpdatedAdditionalActionFee")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUpdatedAdditionalActionFeeIterator{contract: _GatewayEVMTest.contract, event: "UpdatedAdditionalActionFee", logs: logs, sub: sub}, nil
}

// WatchUpdatedAdditionalActionFee is a free log subscription operation binding the contract event 0xe60a2882aff7841a7d0492b804d46ff2d9b07ea353dbc3791e8276b729e40cd8.
//
// Solidity: event UpdatedAdditionalActionFee(uint256 oldFeeWei, uint256 newFeeWei)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUpdatedAdditionalActionFee(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUpdatedAdditionalActionFee) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "UpdatedAdditionalActionFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUpdatedAdditionalActionFee)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedAdditionalActionFee", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUpdatedAdditionalActionFee(log types.Log) (*GatewayEVMTestUpdatedAdditionalActionFee, error) {
	event := new(GatewayEVMTestUpdatedAdditionalActionFee)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedAdditionalActionFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUpdatedCustodyTSSAddressIterator is returned from FilterUpdatedCustodyTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedCustodyTSSAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedCustodyTSSAddressIterator struct {
	Event *GatewayEVMTestUpdatedCustodyTSSAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestUpdatedCustodyTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUpdatedCustodyTSSAddress)
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
		it.Event = new(GatewayEVMTestUpdatedCustodyTSSAddress)
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
func (it *GatewayEVMTestUpdatedCustodyTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUpdatedCustodyTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUpdatedCustodyTSSAddress represents a UpdatedCustodyTSSAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedCustodyTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCustodyTSSAddress is a free log retrieval operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUpdatedCustodyTSSAddress(opts *bind.FilterOpts) (*GatewayEVMTestUpdatedCustodyTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUpdatedCustodyTSSAddressIterator{contract: _GatewayEVMTest.contract, event: "UpdatedCustodyTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedCustodyTSSAddress is a free log subscription operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUpdatedCustodyTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUpdatedCustodyTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUpdatedCustodyTSSAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
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

// ParseUpdatedCustodyTSSAddress is a log parse operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUpdatedCustodyTSSAddress(log types.Log) (*GatewayEVMTestUpdatedCustodyTSSAddress, error) {
	event := new(GatewayEVMTestUpdatedCustodyTSSAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUpdatedDepositAllowedAssetIterator is returned from FilterUpdatedDepositAllowedAsset and is used to iterate over the raw logs and unpacked data for UpdatedDepositAllowedAsset events raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedDepositAllowedAssetIterator struct {
	Event *GatewayEVMTestUpdatedDepositAllowedAsset // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestUpdatedDepositAllowedAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUpdatedDepositAllowedAsset)
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
		it.Event = new(GatewayEVMTestUpdatedDepositAllowedAsset)
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
func (it *GatewayEVMTestUpdatedDepositAllowedAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUpdatedDepositAllowedAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUpdatedDepositAllowedAsset represents a UpdatedDepositAllowedAsset event raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedDepositAllowedAsset struct {
	Asset   common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedDepositAllowedAsset is a free log retrieval operation binding the contract event 0x50708318fef6b3b62fbd0894ea2c5d8fc3ccce4785ed240e4c525ae40bd23332.
//
// Solidity: event UpdatedDepositAllowedAsset(address indexed asset, bool allowed)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUpdatedDepositAllowedAsset(opts *bind.FilterOpts, asset []common.Address) (*GatewayEVMTestUpdatedDepositAllowedAssetIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "UpdatedDepositAllowedAsset", assetRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUpdatedDepositAllowedAssetIterator{contract: _GatewayEVMTest.contract, event: "UpdatedDepositAllowedAsset", logs: logs, sub: sub}, nil
}

// WatchUpdatedDepositAllowedAsset is a free log subscription operation binding the contract event 0x50708318fef6b3b62fbd0894ea2c5d8fc3ccce4785ed240e4c525ae40bd23332.
//
// Solidity: event UpdatedDepositAllowedAsset(address indexed asset, bool allowed)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUpdatedDepositAllowedAsset(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUpdatedDepositAllowedAsset, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "UpdatedDepositAllowedAsset", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUpdatedDepositAllowedAsset)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedDepositAllowedAsset", log); err != nil {
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

// ParseUpdatedDepositAllowedAsset is a log parse operation binding the contract event 0x50708318fef6b3b62fbd0894ea2c5d8fc3ccce4785ed240e4c525ae40bd23332.
//
// Solidity: event UpdatedDepositAllowedAsset(address indexed asset, bool allowed)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUpdatedDepositAllowedAsset(log types.Log) (*GatewayEVMTestUpdatedDepositAllowedAsset, error) {
	event := new(GatewayEVMTestUpdatedDepositAllowedAsset)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedDepositAllowedAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUpdatedDepositPausedIterator is returned from FilterUpdatedDepositPaused and is used to iterate over the raw logs and unpacked data for UpdatedDepositPaused events raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedDepositPausedIterator struct {
	Event *GatewayEVMTestUpdatedDepositPaused // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestUpdatedDepositPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUpdatedDepositPaused)
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
		it.Event = new(GatewayEVMTestUpdatedDepositPaused)
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
func (it *GatewayEVMTestUpdatedDepositPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUpdatedDepositPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUpdatedDepositPaused represents a UpdatedDepositPaused event raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedDepositPaused struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedDepositPaused is a free log retrieval operation binding the contract event 0xd39ad370883b0cd1a8172b5b006a3ebcaaf65183008c91ffd7655afb74174e57.
//
// Solidity: event UpdatedDepositPaused(bool paused)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUpdatedDepositPaused(opts *bind.FilterOpts) (*GatewayEVMTestUpdatedDepositPausedIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "UpdatedDepositPaused")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUpdatedDepositPausedIterator{contract: _GatewayEVMTest.contract, event: "UpdatedDepositPaused", logs: logs, sub: sub}, nil
}

// WatchUpdatedDepositPaused is a free log subscription operation binding the contract event 0xd39ad370883b0cd1a8172b5b006a3ebcaaf65183008c91ffd7655afb74174e57.
//
// Solidity: event UpdatedDepositPaused(bool paused)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUpdatedDepositPaused(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUpdatedDepositPaused) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "UpdatedDepositPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUpdatedDepositPaused)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedDepositPaused", log); err != nil {
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

// ParseUpdatedDepositPaused is a log parse operation binding the contract event 0xd39ad370883b0cd1a8172b5b006a3ebcaaf65183008c91ffd7655afb74174e57.
//
// Solidity: event UpdatedDepositPaused(bool paused)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUpdatedDepositPaused(log types.Log) (*GatewayEVMTestUpdatedDepositPaused, error) {
	event := new(GatewayEVMTestUpdatedDepositPaused)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedDepositPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedGatewayTSSAddressIterator struct {
	Event *GatewayEVMTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestUpdatedGatewayTSSAddress)
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
		it.Event = new(GatewayEVMTestUpdatedGatewayTSSAddress)
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
func (it *GatewayEVMTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestUpdatedGatewayTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*GatewayEVMTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestUpdatedGatewayTSSAddressIterator{contract: _GatewayEVMTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestUpdatedGatewayTSSAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*GatewayEVMTestUpdatedGatewayTSSAddress, error) {
	event := new(GatewayEVMTestUpdatedGatewayTSSAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the GatewayEVMTest contract.
type GatewayEVMTestWhitelistedIterator struct {
	Event *GatewayEVMTestWhitelisted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWhitelisted)
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
		it.Event = new(GatewayEVMTestWhitelisted)
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
func (it *GatewayEVMTestWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWhitelisted represents a Whitelisted event raised by the GatewayEVMTest contract.
type GatewayEVMTestWhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWhitelisted(opts *bind.FilterOpts, token []common.Address) (*GatewayEVMTestWhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWhitelistedIterator{contract: _GatewayEVMTest.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWhitelisted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Whitelisted", log); err != nil {
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

// ParseWhitelisted is a log parse operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWhitelisted(log types.Log) (*GatewayEVMTestWhitelisted, error) {
	event := new(GatewayEVMTestWhitelisted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnIterator struct {
	Event *GatewayEVMTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWithdrawn)
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
		it.Event = new(GatewayEVMTestWithdrawn)
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
func (it *GatewayEVMTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWithdrawn represents a Withdrawn event raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawn struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWithdrawnIterator{contract: _GatewayEVMTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWithdrawn, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWithdrawn)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWithdrawn(log types.Log) (*GatewayEVMTestWithdrawn, error) {
	event := new(GatewayEVMTestWithdrawn)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndCalledIterator struct {
	Event *GatewayEVMTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWithdrawnAndCalled)
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
		it.Event = new(GatewayEVMTestWithdrawnAndCalled)
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
func (it *GatewayEVMTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndCalled struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWithdrawnAndCalledIterator{contract: _GatewayEVMTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWithdrawnAndCalled, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWithdrawnAndCalled)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayEVMTestWithdrawnAndCalled, error) {
	event := new(GatewayEVMTestWithdrawnAndCalled)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndRevertedIterator struct {
	Event *GatewayEVMTestWithdrawnAndReverted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestWithdrawnAndReverted)
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
		it.Event = new(GatewayEVMTestWithdrawnAndReverted)
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
func (it *GatewayEVMTestWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the GatewayEVMTest contract.
type GatewayEVMTestWithdrawnAndReverted struct {
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMTestWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestWithdrawnAndRevertedIterator{contract: _GatewayEVMTest.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestWithdrawnAndReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestWithdrawnAndReverted)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
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

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseWithdrawnAndReverted(log types.Log) (*GatewayEVMTestWithdrawnAndReverted, error) {
	event := new(GatewayEVMTestWithdrawnAndReverted)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogIterator struct {
	Event *GatewayEVMTestLog // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLog)
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
		it.Event = new(GatewayEVMTestLog)
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
func (it *GatewayEVMTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLog represents a Log event raised by the GatewayEVMTest contract.
type GatewayEVMTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayEVMTestLogIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogIterator{contract: _GatewayEVMTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLog)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLog(log types.Log) (*GatewayEVMTestLog, error) {
	event := new(GatewayEVMTestLog)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogAddressIterator struct {
	Event *GatewayEVMTestLogAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogAddress)
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
		it.Event = new(GatewayEVMTestLogAddress)
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
func (it *GatewayEVMTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogAddress represents a LogAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayEVMTestLogAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogAddressIterator{contract: _GatewayEVMTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_address", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogAddress(log types.Log) (*GatewayEVMTestLogAddress, error) {
	event := new(GatewayEVMTestLogAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArrayIterator struct {
	Event *GatewayEVMTestLogArray // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogArray)
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
		it.Event = new(GatewayEVMTestLogArray)
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
func (it *GatewayEVMTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogArray represents a LogArray event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayEVMTestLogArrayIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogArrayIterator{contract: _GatewayEVMTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogArray)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogArray(log types.Log) (*GatewayEVMTestLogArray, error) {
	event := new(GatewayEVMTestLogArray)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray0Iterator struct {
	Event *GatewayEVMTestLogArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogArray0)
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
		it.Event = new(GatewayEVMTestLogArray0)
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
func (it *GatewayEVMTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogArray0 represents a LogArray0 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayEVMTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogArray0Iterator{contract: _GatewayEVMTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogArray0)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array0", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogArray0(log types.Log) (*GatewayEVMTestLogArray0, error) {
	event := new(GatewayEVMTestLogArray0)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray1Iterator struct {
	Event *GatewayEVMTestLogArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogArray1)
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
		it.Event = new(GatewayEVMTestLogArray1)
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
func (it *GatewayEVMTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogArray1 represents a LogArray1 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayEVMTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogArray1Iterator{contract: _GatewayEVMTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogArray1)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array1", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogArray1(log types.Log) (*GatewayEVMTestLogArray1, error) {
	event := new(GatewayEVMTestLogArray1)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytesIterator struct {
	Event *GatewayEVMTestLogBytes // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogBytes)
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
		it.Event = new(GatewayEVMTestLogBytes)
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
func (it *GatewayEVMTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogBytes represents a LogBytes event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayEVMTestLogBytesIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogBytesIterator{contract: _GatewayEVMTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogBytes)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogBytes(log types.Log) (*GatewayEVMTestLogBytes, error) {
	event := new(GatewayEVMTestLogBytes)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytes32Iterator struct {
	Event *GatewayEVMTestLogBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogBytes32)
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
		it.Event = new(GatewayEVMTestLogBytes32)
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
func (it *GatewayEVMTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogBytes32 represents a LogBytes32 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayEVMTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogBytes32Iterator{contract: _GatewayEVMTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogBytes32)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogBytes32(log types.Log) (*GatewayEVMTestLogBytes32, error) {
	event := new(GatewayEVMTestLogBytes32)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogIntIterator struct {
	Event *GatewayEVMTestLogInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogInt)
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
		it.Event = new(GatewayEVMTestLogInt)
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
func (it *GatewayEVMTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogInt represents a LogInt event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayEVMTestLogIntIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogIntIterator{contract: _GatewayEVMTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogInt)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_int", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogInt(log types.Log) (*GatewayEVMTestLogInt, error) {
	event := new(GatewayEVMTestLogInt)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedAddressIterator struct {
	Event *GatewayEVMTestLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedAddress)
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
		it.Event = new(GatewayEVMTestLogNamedAddress)
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
func (it *GatewayEVMTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedAddressIterator{contract: _GatewayEVMTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedAddress)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayEVMTestLogNamedAddress, error) {
	event := new(GatewayEVMTestLogNamedAddress)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArrayIterator struct {
	Event *GatewayEVMTestLogNamedArray // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedArray)
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
		it.Event = new(GatewayEVMTestLogNamedArray)
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
func (it *GatewayEVMTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedArray represents a LogNamedArray event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedArrayIterator{contract: _GatewayEVMTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedArray)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayEVMTestLogNamedArray, error) {
	event := new(GatewayEVMTestLogNamedArray)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray0Iterator struct {
	Event *GatewayEVMTestLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedArray0)
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
		it.Event = new(GatewayEVMTestLogNamedArray0)
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
func (it *GatewayEVMTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedArray0Iterator{contract: _GatewayEVMTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedArray0)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayEVMTestLogNamedArray0, error) {
	event := new(GatewayEVMTestLogNamedArray0)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray1Iterator struct {
	Event *GatewayEVMTestLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedArray1)
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
		it.Event = new(GatewayEVMTestLogNamedArray1)
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
func (it *GatewayEVMTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedArray1Iterator{contract: _GatewayEVMTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedArray1)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayEVMTestLogNamedArray1, error) {
	event := new(GatewayEVMTestLogNamedArray1)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytesIterator struct {
	Event *GatewayEVMTestLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedBytes)
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
		it.Event = new(GatewayEVMTestLogNamedBytes)
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
func (it *GatewayEVMTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedBytesIterator{contract: _GatewayEVMTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedBytes)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayEVMTestLogNamedBytes, error) {
	event := new(GatewayEVMTestLogNamedBytes)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytes32Iterator struct {
	Event *GatewayEVMTestLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedBytes32)
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
		it.Event = new(GatewayEVMTestLogNamedBytes32)
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
func (it *GatewayEVMTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedBytes32Iterator{contract: _GatewayEVMTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedBytes32)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayEVMTestLogNamedBytes32, error) {
	event := new(GatewayEVMTestLogNamedBytes32)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalIntIterator struct {
	Event *GatewayEVMTestLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedDecimalInt)
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
		it.Event = new(GatewayEVMTestLogNamedDecimalInt)
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
func (it *GatewayEVMTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedDecimalIntIterator{contract: _GatewayEVMTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedDecimalInt)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayEVMTestLogNamedDecimalInt, error) {
	event := new(GatewayEVMTestLogNamedDecimalInt)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalUintIterator struct {
	Event *GatewayEVMTestLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedDecimalUint)
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
		it.Event = new(GatewayEVMTestLogNamedDecimalUint)
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
func (it *GatewayEVMTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedDecimalUintIterator{contract: _GatewayEVMTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedDecimalUint)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayEVMTestLogNamedDecimalUint, error) {
	event := new(GatewayEVMTestLogNamedDecimalUint)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedIntIterator struct {
	Event *GatewayEVMTestLogNamedInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedInt)
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
		it.Event = new(GatewayEVMTestLogNamedInt)
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
func (it *GatewayEVMTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedInt represents a LogNamedInt event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedIntIterator{contract: _GatewayEVMTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedInt)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayEVMTestLogNamedInt, error) {
	event := new(GatewayEVMTestLogNamedInt)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedStringIterator struct {
	Event *GatewayEVMTestLogNamedString // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedString)
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
		it.Event = new(GatewayEVMTestLogNamedString)
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
func (it *GatewayEVMTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedString represents a LogNamedString event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedStringIterator{contract: _GatewayEVMTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedString)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedString(log types.Log) (*GatewayEVMTestLogNamedString, error) {
	event := new(GatewayEVMTestLogNamedString)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedUintIterator struct {
	Event *GatewayEVMTestLogNamedUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogNamedUint)
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
		it.Event = new(GatewayEVMTestLogNamedUint)
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
func (it *GatewayEVMTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogNamedUint represents a LogNamedUint event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayEVMTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogNamedUintIterator{contract: _GatewayEVMTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogNamedUint)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayEVMTestLogNamedUint, error) {
	event := new(GatewayEVMTestLogNamedUint)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogStringIterator struct {
	Event *GatewayEVMTestLogString // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogString)
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
		it.Event = new(GatewayEVMTestLogString)
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
func (it *GatewayEVMTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogString represents a LogString event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayEVMTestLogStringIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogStringIterator{contract: _GatewayEVMTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogString)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_string", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogString(log types.Log) (*GatewayEVMTestLogString, error) {
	event := new(GatewayEVMTestLogString)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogUintIterator struct {
	Event *GatewayEVMTestLogUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogUint)
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
		it.Event = new(GatewayEVMTestLogUint)
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
func (it *GatewayEVMTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogUint represents a LogUint event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayEVMTestLogUintIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogUintIterator{contract: _GatewayEVMTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogUint)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "log_uint", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogUint(log types.Log) (*GatewayEVMTestLogUint, error) {
	event := new(GatewayEVMTestLogUint)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayEVMTest contract.
type GatewayEVMTestLogsIterator struct {
	Event *GatewayEVMTestLogs // Event containing the contract specifics and raw log

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
func (it *GatewayEVMTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMTestLogs)
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
		it.Event = new(GatewayEVMTestLogs)
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
func (it *GatewayEVMTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMTestLogs represents a Logs event raised by the GatewayEVMTest contract.
type GatewayEVMTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayEVMTestLogsIterator, error) {

	logs, sub, err := _GatewayEVMTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMTestLogsIterator{contract: _GatewayEVMTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMTest *GatewayEVMTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayEVMTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMTestLogs)
				if err := _GatewayEVMTest.contract.UnpackLog(event, "logs", log); err != nil {
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
func (_GatewayEVMTest *GatewayEVMTestFilterer) ParseLogs(log types.Log) (*GatewayEVMTestLogs, error) {
	event := new(GatewayEVMTestLogs)
	if err := _GatewayEVMTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
