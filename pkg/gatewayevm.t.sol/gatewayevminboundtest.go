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

// GatewayEVMInboundTestMetaData contains all meta data concerning the GatewayEVMInboundTest contract.
var GatewayEVMInboundTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ADDITIONAL_ACTION_FEE_WEI\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testAdditionalActionDisabledReverts\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallNotAffectedByDepositPause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallSecondActionFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallSecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfCallOnRevertIsTrue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfDestinationIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20SecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20ToCustodyFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20ToCustodyFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthToTssFailsForSubsequentActions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthWithAmountSecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthWithAmountToTssFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20SecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodySecondActionFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTss\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssFailsForSubsequentActions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountSecondActionFailsWithOnlyFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountSecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTss\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositPauseAllowsAllowlistedERC20Deposits\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositPauseAllowsNativeByDefault\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositPauseAllowsZetaByDefault\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositPauseBlocksNativeDeposits\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositPauseBlocksNonZetaERC20Deposits\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZetaToConnector\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testFeeSystemWithUpdatedFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testMixedActionTypesInSameTransaction\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositERC20ToCustodyIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositERC20ToCustodyIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositERC20ToCustodyWithPayloadIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssWithPayloadIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetDepositAllowedAsset\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetDepositPaused\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetDepositPausedPreservesExplicitZetaBlockAfterRePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpdateAdditionalActionFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpdateAdditionalActionFeeOnlyAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositedAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedAdditionalActionFee\",\"inputs\":[{\"name\":\"oldFeeWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newFeeWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedDepositAllowedAsset\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedDepositPaused\",\"inputs\":[{\"name\":\"paused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AdditionalActionDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ArbitraryCallDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AssetDepositNotAllowed\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallOnRevertNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExcessETHProvided\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectValueProvided\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFee\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZETANotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f8054909116909117905569021e19e0c9bab2400000602c553480156039575f80fd5b506201312f80620000495f395ff3fe608060405234801562000010575f80fd5b506004361062000600575f3560e01c80637e7dfee31162000317578063b5508aa911620001a7578063ddd03d5711620000fb578063f0a635c511620000ab578063f6e53a5d1162000083578063f6e53a5d1462000995578063f8d37ef2146200099f578063fa7626d414620009b9575f80fd5b8063f0a635c51462000977578063f1c6b4d31462000981578063f2036eda146200098b575f80fd5b8063e306a97811620000df578063e306a9781462000959578063e85c5a071462000963578063ef17b2ed146200096d575f80fd5b8063ddd03d571462000945578063e20c9f71146200094f575f80fd5b8063c57926c61162000157578063cdef966f116200013b578063cdef966f1462000927578063d86a4a0c1462000931578063dc23a35f146200093b575f80fd5b8063c57926c61462000913578063c7a1eccb146200091d575f80fd5b8063ba414fa6116200018b578063ba414fa614620008e4578063ba46ba2314620008ff578063c51a4cbe1462000909575f80fd5b8063b5508aa914620008d0578063b966e8fa14620008da575f80fd5b806395cd0445116200026b578063a00a6fff116200021b578063b0464fdc11620001f3578063b0464fdc14620008b2578063b1409f7114620008bc578063b284906314620008c6575f80fd5b8063a00a6fff1462000894578063a0d60b3a146200089e578063aa030c1c14620008a8575f80fd5b80639acda9fa116200024f5780639acda9fa14620008765780639d96310a14620008805780639fd1e597146200088a575f80fd5b806395cd044514620008625780639a34d8f3146200086c575f80fd5b80638f5a422411620002c75780639078b01c11620002ab5780639078b01c1462000835578063916a17c6146200083f57806391a5c8181462000858575f80fd5b80638f5a422414620008215780639073eafe146200082b575f80fd5b806385f5c51c11620002fb57806385f5c51c146200080357806389b90256146200080d5780638be96f5c1462000817575f80fd5b80637e7dfee314620007e057806385226c8114620007ea575f80fd5b8063466f332e1162000493578063586da26711620003e757806363d7a41811620003975780636aa02e8b116200036f5780636aa02e8b14620007c25780636ab1c51614620007cc5780636c33bacb14620007d6575f80fd5b806363d7a41814620007955780636459542a146200079f57806366d9a9a014620007a9575f80fd5b80635c4013d011620003cb5780635c4013d014620007775780635d8b85eb14620007815780635e8fe81f146200078b575f80fd5b8063586da2671462000763578063598b7edb146200076d575f80fd5b80634ce25c0a116200044357806351ee53cb116200042757806351ee53cb1462000745578063545c3745146200074f57806358592e3f1462000759575f80fd5b80634ce25c0a146200073157806351da903d146200073b575f80fd5b806348452329116200047757806348452329146200071357806349050a44146200071d5780634a7803391462000727575f80fd5b8063466f332e14620006ff578063481daadb1462000709575f80fd5b80632a5dcf36116200055757806333ed091c11620005075780633f7286f411620004df5780633f7286f414620006e157806341b8fec614620006eb57806343171c1214620006f5575f80fd5b806333ed091c14620006c35780633424914f14620006cd5780633e5e3c2314620006d7575f80fd5b80632cf9572d116200053b5780632cf9572d14620006a557806330f7c04f14620006af57806332fc1fad14620006b9575f80fd5b80632a5dcf3614620006825780632ade3880146200068c575f80fd5b806315ee8f3611620005b35780631b906ef311620005975780631b906ef3146200064c5780631ed7831c146200065657806320a5ccfe1462000678575f80fd5b806315ee8f36146200063857806318e141381462000642575f80fd5b806309b21ccb11620005e757806309b21ccb146200061a5780630a9254e414620006245780630fc37f5e146200062e575f80fd5b80630724d8e3146200060457806309a263c11462000610575b5f80fd5b6200060e620009c7565b005b6200060e62000b7d565b6200060e62000d90565b6200060e62000eca565b6200060e62001a45565b6200060e62001c33565b6200060e62001e0d565b6200060e62001f33565b62000660620022b8565b6040516200066f91906200f9be565b60405180910390f35b6200060e6200231a565b6200060e62002833565b6200069662002952565b6040516200066f91906200fa39565b6200060e62002a9a565b6200060e62002cb5565b6200060e620030ab565b6200060e6200320d565b6200060e62003373565b6200066062003671565b62000660620036d1565b6200060e62003731565b6200060e620037f2565b6200060e62003a16565b6200060e62003c4a565b6200060e62003f84565b6200060e62004264565b6200060e62004510565b6200060e62004631565b6200060e62004875565b6200060e6200498c565b6200060e62004aad565b6200060e62004d51565b6200060e62004f4a565b6200060e62005107565b6200060e6200523e565b6200060e62005312565b6200060e620057cc565b6200060e6200597f565b6200060e62005fd2565b620007b362006365565b6040516200066f91906200fb8a565b6200060e620064d5565b6200060e620065d4565b6200060e62006ba9565b6200060e62006d9f565b620007f462006ef3565b6040516200066f91906200fc2c565b6200060e62006fc8565b6200060e62007133565b6200060e62007492565b6200060e620075f1565b6200060e620076c8565b6200060e6200780f565b6200084962007d94565b6040516200066f91906200fca5565b6200060e62007e79565b6200060e62007f92565b6200060e62008259565b6200060e62008377565b6200060e6200844d565b6200060e62008569565b6200060e6200870d565b6200060e62008af7565b6200060e62008c43565b6200084962008dd3565b6200060e62008eb8565b6200060e62009162565b620007f462009352565b6200060e62009427565b620008ee6200966d565b60405190151581526020016200066f565b6200060e62009741565b6200060e6200981e565b6200060e62009a07565b6200060e62009ae0565b6200060e62009c5a565b6200060e62009d49565b6200060e6200a0d8565b6200060e6200a61f565b620006606200a751565b6200060e6200a7b1565b6200060e6200a966565b6200060e6200ab9a565b6200060e6200aee8565b6200060e6200b021565b6200060e6200b1f4565b6200060e6200b350565b620009aa62030d4081565b6040519081526020016200066f565b601f54620008ee9060ff1681565b6027546020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039182166084820152620186a092919091163190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562000a61575f80fd5b505af115801562000a74573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c9062000ac49086905f906028906200fe7c565b60405180910390a36020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263726ac97c92869262000b2192909116906028906004016200feb2565b5f604051808303818588803b15801562000b39575f80fd5b505af115801562000b4c573d5f803e3d5ffd5b50506027546001600160a01b031631925062000b78915062000b71905084846200ff02565b826200b482565b505050565b6020546026546040517f282906ed000000000000000000000000000000000000000000000000000000008152620186a0926001600160a01b039081169263282906ed92859262000bd792169083906028906004016200ff18565b5f604051808303818588803b15801562000bef575f80fd5b505af115801562000c02573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392507f84222ac700000000000000000000000000000000000000000000000000000000915062000c55905062030d40856200ff02565b62000c6462030d40866200ff02565b62000c719060016200ff02565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b909216825262000ccb916004016200ff41565b5f604051808303815f87803b15801562000ce3575f80fd5b505af115801562000cf6573d5f803e3d5ffd5b50506020546001600160a01b0316915063282906ed905062000d1c62030d40846200ff02565b62000d299060016200ff02565b6026546040516001600160e01b031960e085901b16815262000d5d916001600160a01b03169086906028906004016200ff18565b5f604051808303818588803b15801562000d75575f80fd5b505af115801562000d88573d5f803e3d5ffd5b505050505050565b6026546040516001600160a01b0390911660248201526001905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562000e43575f80fd5b505af115801562000e56573d5f803e3d5ffd5b50506020546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063d09e3b78935062000eb2925f92889291169087906028906004016200ff55565b5f604051808303815f87803b15801562000d75575f80fd5b602580547fffffffffffffffffffffffff000000000000000000000000000000000000000090811630179091556026805482166112341790556027805490911661567817905560405162000f1e906200f8f1565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103905ff08015801562000fa1573d5f803e3d5ffd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316179055602754604051911690819062000fec906200f8ff565b6001600160a01b03928316815291166020820152604001604051809103905ff0801580156200101d573d5f803e3d5ffd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081178255604080518082018252600e81527f4761746577617945564d2e736f6c000000000000000000000000000000000000602082015260275460255492519086169481019490945260448401929092529092166064820152620010fc91906084015b60408051601f198184030181529190526020810180516001600160e01b03167fc0c53b8b000000000000000000000000000000000000000000000000000000001790526200b4ff565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff000000000000000000000000000000000000000090911681178255604080518082018252601081527f4552433230437573746f64792e736f6c000000000000000000000000000000009381019390935260275460255491516024810193909352841660448301529092166064830152620011ce91608401620010b3565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602180549190920483167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055604080518082018252601a81527f5a657461436f6e6e6563746f724e6f6e4e61746976652e736f6c00000000000060208083019190915254602480546027546025549551938716928401929092528516604483015284166064820152919092166084820152620012f4919060a40160408051601f198184030181529190526020810180516001600160e01b03167ff8c8765e000000000000000000000000000000000000000000000000000000001790526200b4ff565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602280549190920483167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790556027546040517fca669fa700000000000000000000000000000000000000000000000000000000815291166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620013ca575f80fd5b505af1158015620013dd573d5f803e3d5ffd5b5050602480546027546022546040517f15d57fd40000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152908216938101939093521692506315d57fd491506044015f604051808303815f87803b1580156200144e575f80fd5b505af115801562001461573d5f803e3d5ffd5b50506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d91506044015f604051808303815f87803b158015620014e3575f80fd5b505af1158015620014f6573d5f803e3d5ffd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d5691506024015f604051808303815f87803b1580156200156a575f80fd5b505af11580156200157d573d5f803e3d5ffd5b50506020546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f91506024015f604051808303815f87803b158015620015e1575f80fd5b505af1158015620015f4573d5f803e3d5ffd5b50506020546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef91506024015f604051808303815f87803b15801562001658575f80fd5b505af11580156200166b573d5f803e3d5ffd5b50506021546023546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639b19251a91506024015f604051808303815f87803b158015620016cf575f80fd5b505af1158015620016e2573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562001741575f80fd5b505af115801562001754573d5f803e3d5ffd5b5050602354602554602c546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810191909152911692506340c10f1991506044015f604051808303815f87803b158015620017c3575f80fd5b505af1158015620017d6573d5f803e3d5ffd5b50506027546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b1580156200184a575f80fd5b505af11580156200185d573d5f803e3d5ffd5b5050602254602554602c546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101919091525f60448201529116925063106e629091506064015f604051808303815f87803b158015620018d2575f80fd5b505af1158015620018e5573d5f803e3d5ffd5b50506040805160a0810182526103218082525f602080840182815284860193845285519182019095528181526060840181905260808401919091528251602880549551151574010000000000000000000000000000000000000000027fffffffffffffffffffffff0000000000000000000000000000000000000000009096166001600160a01b0392831617959095178555915160298054919093167fffffffffffffffffffffffff00000000000000000000000000000000000000009190911617909155909350909150602a90620019bf908262010022565b50608091909101516003909101556020546040517f7c74425300000000000000000000000000000000000000000000000000000000815262030d4060048201526001600160a01b0390911690637c744253906024015f604051808303815f87803b15801562001a2c575f80fd5b505af115801562001a3f573d5f803e3d5ffd5b50505050565b60208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562001aa2573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062001ac89190620100ea565b62001ad59060016200ff02565b67ffffffffffffffff81111562001af05762001af06200ffad565b6040519080825280601f01601f19166020018201604052801562001b1b576020820181803683370190505b50602a9062001b2b908262010022565b50604051630618f58760e51b81527f386691aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562001b96575f80fd5b505af115801562001ba9573d5f803e3d5ffd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c935060019262001c019216906028906004016200feb2565b5f604051808303818588803b15801562001c19575f80fd5b505af115801562001c2c573d5f803e3d5ffd5b5050505050565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a0602482018190529261c35092169063095ea7b3906044016020604051808303815f875af115801562001ca7573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062001ccd919062010102565b506040515f602482015260448101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d9063f28dceb3907f8afe4db400000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b909216825262001d67916004016200ff41565b5f604051808303815f87803b15801562001d7f575f80fd5b505af115801562001d92573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b09450869362001dd993811692899291169060289060040162010123565b5f604051808303818588803b15801562001df1575f80fd5b505af115801562001e04573d5f803e3d5ffd5b50505050505050565b6026546040516001600160a01b0390911660248201525f90819060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f951e19ed000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b15801562001ec0575f80fd5b505af115801562001ed3573d5f803e3d5ffd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b78945062000eb293928316928892169087906028906004016200ff55565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052602754602554915163248e63e160e11b8152600160048201819052602482018190526044820181905260648201529293506001600160a01b03908116319291163190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562001ff6575f80fd5b505af115801562002009573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f906200205b9088905f9089906028906201015b565b60405180910390a36020546026546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263397e375c928892620020bc92909116908390899060289060040162010198565b5f604051808303818588803b158015620020d4575f80fd5b505af1158015620020e7573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063491cc7c2925060840190505f604051808303815f87803b1580156200214d575f80fd5b505af115801562002160573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f90620021b29088905f9089906028906201015b565b60405180910390a36020546001600160a01b031663397e375c620021da62030d40876200ff02565b6026546040516001600160e01b031960e085901b16815262002210916001600160a01b0316908990899060289060040162010198565b5f604051808303818588803b15801562002228575f80fd5b505af11580156200223b573d5f803e3d5ffd5b50506027546025546001600160a01b03918216319450163191506200228b905062030d406200226c886002620101c1565b6200227890876200ff02565b6200228491906200ff02565b836200b482565b62000d8862030d40620022a0886002620101c1565b620022ac9086620101db565b62000b719190620101db565b606060168054806020026020016040519081016040528092919081815260200182805480156200231057602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311620022f1575b5050505050905090565b6020546040517f543f66a400000000000000000000000000000000000000000000000000000000815260016004820152620186a0916001600160a01b03169063543f66a4906024015f604051808303815f87803b1580156200237a575f80fd5b505af11580156200238d573d5f803e3d5ffd5b5050602054602480546040517f4f27eb2a0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f928101929092529091169250634f27eb2a91506044015f604051808303815f87803b158015620023fb575f80fd5b505af11580156200240e573d5f803e3d5ffd5b5050602054602480546040517f6ae23fb50000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620024a7955092169250636ae23fb591015b602060405180830381865afa1580156200247b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620024a1919062010102565b6200b521565b6020546040517f543f66a40000000000000000000000000000000000000000000000000000000081525f60048201526001600160a01b039091169063543f66a4906024015f604051808303815f87803b15801562002503575f80fd5b505af115801562002516573d5f803e3d5ffd5b50506020546040517f543f66a4000000000000000000000000000000000000000000000000000000008152600160048201526001600160a01b03909116925063543f66a491506024015f604051808303815f87803b15801562002577575f80fd5b505af11580156200258a573d5f803e3d5ffd5b5050602054602480546040517f6ae23fb50000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620025e0955092169250636ae23fb591016200245f565b60208054604080517f02befd2400000000000000000000000000000000000000000000000000000000815290516200266f936001600160a01b03909316926302befd2492600480820193918290030181865afa15801562002643573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002669919062010102565b6200b599565b602480546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152928301849052169063095ea7b3906044016020604051808303815f875af1158015620026da573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002700919062010102565b5060248054604080516001600160a01b039092168284015280518083039093018352604490910181526020820180516001600160e01b03167f33e23dbb00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391620027ab91906004016200ff41565b5f604051808303815f87803b158015620027c3575f80fd5b505af1158015620027d6573d5f803e3d5ffd5b5050602054602654602454604051630102614b60e41b81526001600160a01b03938416955063102614b094506200281b93928316928792169060289060040162010123565b5f604051808303815f87803b15801562001c19575f80fd5b6026546040516001600160a01b0390911660248201525f90819060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f7671265e000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b158015620028e6575f80fd5b505af1158015620028f9573d5f803e3d5ffd5b50506020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063744b9b8b9350869262001dd99216908690602890600401620101f1565b6060601e805480602002602001604051908101604052809291908181526020015f905b8282101562002a91575f84815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b8282101562002a79578382905f5260205f20018054620029e7906200fd3e565b80601f016020809104026020016040519081016040528092919081815260200182805462002a15906200fd3e565b801562002a645780601f1062002a3a5761010080835404028352916020019162002a64565b820191905f5260205f20905b81548152906001019060200180831162002a4657829003601f168201915b505050505081526020019060010190620029c7565b50505050815250508152602001906001019062002975565b50505050905090565b602654604080516001600160a01b0392831660248083018290528351808403909101815260449092018352602080830180516001600160e01b0316630427d73b60e51b1790525492517f282906ed000000000000000000000000000000000000000000000000000000008152620186a09492939092169163282906ed91859162002b2d919083906028906004016200ff18565b5f604051808303818588803b15801562002b45575f80fd5b505af115801562002b58573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392507f84222ac700000000000000000000000000000000000000000000000000000000915062002bab905062030d40866200ff02565b62002bba62030d40876200ff02565b62002bc79060016200ff02565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b909216825262002c21916004016200ff41565b5f604051808303815f87803b15801562002c39575f80fd5b505af115801562002c4c573d5f803e3d5ffd5b50506020546001600160a01b0316915063397e375c905062002c7262030d40856200ff02565b62002c7f9060016200ff02565b6026546040516001600160e01b031960e085901b16815262001dd9916001600160a01b0316908790879060289060040162010198565b6023546021546040516370a0823160e01b81526001600160a01b039182166004820152620186a0925f9216906370a0823190602401602060405180830381865afa15801562002d06573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002d2c9190620100ea565b905062002d3a5f826200b482565b6026546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052602354905491517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101879052929350169063095ea7b3906044016020604051808303815f875af115801562002dec573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002e12919062010102565b506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562002e9f575f80fd5b505af115801562002eb2573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f9262002f0a9289929091169087906028906201015b565b60405180910390a36020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b789362002f70939082169289929091169087906028906004016200ff55565b5f604051808303815f87803b15801562002f88575f80fd5b505af115801562002f9b573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa15801562002fec573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620030129190620100ea565b90506200302084826200b482565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201525f9291909116906370a0823190602401602060405180830381865afa1580156200306f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620030959190620100ea565b905062001c2c85602c5462000b719190620101db565b6020546026546040517f726ac97c000000000000000000000000000000000000000000000000000000008152620186a0926001600160a01b039081169263726ac97c928592620031039216906028906004016200feb2565b5f604051808303818588803b1580156200311b575f80fd5b505af11580156200312e573d5f803e3d5ffd5b5050604051630618f58760e51b81527f7671265e000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063c31eb0e0925060240190505f604051808303815f87803b1580156200319e575f80fd5b505af1158015620031b1573d5f803e3d5ffd5b50506020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063282906ed935062030d409262000d5d9216905f906028906004016200ff18565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052515f602482015261c35060448201819052919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9063f28dceb3907f8afe4db400000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b9092168252620032ef916004016200ff41565b5f604051808303815f87803b15801562003307575f80fd5b505af11580156200331a573d5f803e3d5ffd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb49350859262001dd99216908790602890600401620101f1565b60275460255460405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152620186a0926001600160a01b039081163192163190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b158015620033ee575f80fd5b505af115801562003401573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c90620034519087905f906028906200fe7c565b60405180910390a36020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263282906ed928792620034b0929091169083906028906004016200ff18565b5f604051808303818588803b158015620034c8575f80fd5b505af1158015620034db573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063491cc7c2925060840190505f604051808303815f87803b15801562003541575f80fd5b505af115801562003554573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c90620035a49087905f906028906200fe7c565b60405180910390a36020546001600160a01b031663282906ed620035cc62030d40866200ff02565b6026546040516001600160e01b031960e085901b16815262003600916001600160a01b03169088906028906004016200ff18565b5f604051808303818588803b15801562003618575f80fd5b505af11580156200362b573d5f803e3d5ffd5b50506027546025546001600160a01b03918216319450163191506200365c905062030d406200226c876002620101c1565b62001c2c62030d40620022a0876002620101c1565b606060188054806020026020016040519081016040528092919081815260200182805480156200231057602002820191905f5260205f209081546001600160a01b03168152600190910190602001808311620022f1575050505050905090565b606060178054806020026020016040519081016040528092919081815260200182805480156200231057602002820191905f5260205f209081546001600160a01b03168152600190910190602001808311620022f1575050505050905090565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156200379e575f80fd5b505af1158015620037b1573d5f803e3d5ffd5b5050602054602354604051630102614b60e41b81526001600160a01b03928316945063102614b093506200281b925f92879291169060289060040162010123565b6027546020546040517f543f66a400000000000000000000000000000000000000000000000000000000815260016004820152620186a0926001600160a01b039081163192169063543f66a4906024015f604051808303815f87803b1580156200385a575f80fd5b505af11580156200386d573d5f803e3d5ffd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a4015f604051808303815f87803b158015620038fd575f80fd5b505af115801562003910573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c90620039609086905f906028906200fe7c565b60405180910390a36020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263726ac97c928692620039bd92909116906028906004016200feb2565b5f604051808303818588803b158015620039d5575f80fd5b505af1158015620039e8573d5f803e3d5ffd5b505050505062003a128282620039ff91906200ff02565b6027546001600160a01b0316316200b482565b5050565b60208054604080517fa2ba19340000000000000000000000000000000000000000000000000000000081529051620186a0935f936002936001600160a01b039091169263a2ba1934926004808401939192918290030181865afa15801562003a80573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062003aa69190620100ea565b62003ab2919062010228565b67ffffffffffffffff81111562003acd5762003acd6200ffad565b6040519080825280601f01601f19166020018201604052801562003af8576020820181803683370190505b5060208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562003b5c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062003b829190620100ea565b62003b8e919062010228565b62003b9b9060016200ff02565b67ffffffffffffffff81111562003bb65762003bb66200ffad565b6040519080825280601f01601f19166020018201604052801562003be1576020820181803683370190505b50602a9062003bf1908262010022565b50604051630618f58760e51b81527f386691aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401620028ce565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052602554602754915163248e63e160e11b8152600160048201819052602482018190526044820181905260648201529293506001600160a01b03908116319291163190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562003d08575f80fd5b505af115801562003d1b573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d9749062003d6990879060289062010261565b60405180910390a36020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692631becceb49262003dc5929116908790602890600401620101f1565b5f604051808303815f87803b15801562003ddd575f80fd5b505af115801562003df0573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562003e54575f80fd5b505af115801562003e67573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d9749062003eb590879060289062010261565b60405180910390a36020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692631becceb49262030d409262003f1792909116908890602890600401620101f1565b5f604051808303818588803b15801562003f2f575f80fd5b505af115801562003f42573d5f803e3d5ffd5b50506025546027546001600160a01b039182163194501631915062003f7190506200228462030d4086620101db565b62001c2c62000b7162030d40856200ff02565b6026546040516001600160a01b039091166024820152620186a09061c350905f9060440160408051601f19818403018152919052602080820180516001600160e01b0316630427d73b60e51b17905260235490549192506001600160a01b039081169163095ea7b3911662003ffb866002620101c1565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801562004044573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200406a919062010102565b506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b7893620040c9939082169289929091169087906028906004016200ff55565b5f604051808303815f87803b158015620040e1575f80fd5b505af1158015620040f4573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507f8afe4db400000000000000000000000000000000000000000000000000000000905062030d406200414686826200ff02565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b9092168252620041a0916004016200ff41565b5f604051808303815f87803b158015620041b8575f80fd5b505af1158015620041cb573d5f803e3d5ffd5b50506020546001600160a01b0316915063d09e3b789050620041f18462030d406200ff02565b6026546023546040516001600160e01b031960e086901b1681526200422f926001600160a01b03908116928a9291169088906028906004016200ff55565b5f604051808303818588803b15801562004247575f80fd5b505af11580156200425a573d5f803e3d5ffd5b5050505050505050565b620186a05f62004279600162030d40620101db565b6026546040516001600160a01b0390911660248201529091505f9060440160408051601f19818403018152919052602080820180516001600160e01b0316630427d73b60e51b17905260235490549192506001600160a01b039081169163095ea7b39116620042ea866002620101c1565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801562004333573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004359919062010102565b506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b7893620043b8939082169289929091169087906028906004016200ff55565b5f604051808303815f87803b158015620043d0575f80fd5b505af1158015620043e3573d5f803e3d5ffd5b505060405162030d40602482015260448101859052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507fa458261b00000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b909216825262004483916004016200ff41565b5f604051808303815f87803b1580156200449b575f80fd5b505af1158015620044ae573d5f803e3d5ffd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b78945087936200422f938116928a9291169088906028906004016200ff55565b6026546040516001600160a01b0390911660248201525f90819060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f7671265e000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b158015620045c3575f80fd5b505af1158015620045d6573d5f803e3d5ffd5b50506020546026546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063397e375c9350869262001dd99216908390879060289060040162010198565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a06024820181905292919091169063095ea7b3906044016020604051808303815f875af1158015620046a4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620046ca919062010102565b5060208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562004728573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200474e9190620100ea565b6200475b9060016200ff02565b67ffffffffffffffff8111156200477657620047766200ffad565b6040519080825280601f01601f191660200182016040528015620047a1576020820181803683370190505b50602a90620047b1908262010022565b50604051630618f58760e51b81527f386691aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b1580156200481d575f80fd5b505af115801562004830573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b094506200281b93928316928792169060289060040162010123565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562004925575f80fd5b505af115801562004938573d5f803e3d5ffd5b50506020546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250631becceb491506200281b905f908590602890600401620101f1565b620186a0737109709ecfa91a80626ff3989d68f67f5b1dd12d63f28dceb37f84222ac700000000000000000000000000000000000000000000000000000000620049d88460016200ff02565b60405160248101919091526044810185905260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b909216825262004a34916004016200ff41565b5f604051808303815f87803b15801562004a4c575f80fd5b505af115801562004a5f573d5f803e3d5ffd5b50506020546026546001600160a01b03918216935063282906ed925084911662004a8b8260016200ff02565b60286040518563ffffffff1660e01b815260040162000d5d939291906200ff18565b60208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562004b0a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004b309190620100ea565b62004b3d9060016200ff02565b67ffffffffffffffff81111562004b585762004b586200ffad565b6040519080825280601f01601f19166020018201604052801562004b83576020820181803683370190505b50602a9062004b93908262010022565b505f6028600201805462004ba7906200fd3e565b905090505f60205f9054906101000a90046001600160a01b03166001600160a01b031663a2ba19346040518163ffffffff1660e01b8152600401602060405180830381865afa15801562004bfd573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004c239190620100ea565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f386691aa0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162004ccc916004016200ff41565b5f604051808303815f87803b15801562004ce4575f80fd5b505af115801562004cf7573d5f803e3d5ffd5b50506020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063282906ed935060019262001dd992169083906028906004016200ff18565b6020546023546040517f6ae23fb50000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015262004da6929190911690636ae23fb5906024016200245f565b60405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562004e06575f80fd5b505af115801562004e19573d5f803e3d5ffd5b5050602354604051600181526001600160a01b0390911692507f50708318fef6b3b62fbd0894ea2c5d8fc3ccce4785ed240e4c525ae40bd23332915060200160405180910390a26020546023546040517f4f27eb2a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260016024820152911690634f27eb2a906044015f604051808303815f87803b15801562004ec7575f80fd5b505af115801562004eda573d5f803e3d5ffd5b50506020546023546040517f6ae23fb50000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015262004f48945091169150636ae23fb5906024015b602060405180830381865afa15801562002643573d5f803e3d5ffd5b565b60405163248e63e160e11b8152600160048201819052602482018190526044820181905260648201526509184e72a00090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562004fb2575f80fd5b505af115801562004fc5573d5f803e3d5ffd5b50506040805162030d408152602081018590527fe60a2882aff7841a7d0492b804d46ff2d9b07ea353dbc3791e8276b729e40cd8935001905060405180910390a16020546040517f7c744253000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b0390911690637c744253906024015f604051808303815f87803b15801562005063575f80fd5b505af115801562005076573d5f803e3d5ffd5b505060208054604080517fb010721400000000000000000000000000000000000000000000000000000000815290516200510495506001600160a01b03909216935063b01072149260048083019391928290030181865afa158015620050de573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062000b719190620100ea565b50565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529190526020810180516001600160e01b0316630427d73b60e51b17905290505f6200515f600162030d40620101db565b6020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081529293506001600160a01b0391821692631becceb492620051b59216908690602890600401620101f1565b5f604051808303815f87803b158015620051cd575f80fd5b505af1158015620051e0573d5f803e3d5ffd5b505060405162030d40602482015260448101849052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507fa458261b0000000000000000000000000000000000000000000000000000000090606401620032a8565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015620052ab575f80fd5b505af1158015620052be573d5f803e3d5ffd5b50506020546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063726ac97c9150839062000d5d905f906028906004016200feb2565b60405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562005372575f80fd5b505af115801562005385573d5f803e3d5ffd5b5050602454604051600181526001600160a01b0390911692507f50708318fef6b3b62fbd0894ea2c5d8fc3ccce4785ed240e4c525ae40bd23332915060200160405180910390a260405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b1580156200542c575f80fd5b505af11580156200543f573d5f803e3d5ffd5b5050604051600181525f92507f50708318fef6b3b62fbd0894ea2c5d8fc3ccce4785ed240e4c525ae40bd23332915060200160405180910390a260405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b158015620054d9575f80fd5b505af1158015620054ec573d5f803e3d5ffd5b5050604051600181527fd39ad370883b0cd1a8172b5b006a3ebcaaf65183008c91ffd7655afb74174e579250602001905060405180910390a16020546040517f543f66a4000000000000000000000000000000000000000000000000000000008152600160048201526001600160a01b039091169063543f66a4906024015f604051808303815f87803b15801562005582575f80fd5b505af115801562005595573d5f803e3d5ffd5b505060208054604080517f02befd240000000000000000000000000000000000000000000000000000000081529051620055fd95506001600160a01b0390921693506302befd249260048083019391928290030181865afa15801562002643573d5f803e3d5ffd5b6020546040517f6ae23fb50000000000000000000000000000000000000000000000000000000081525f60048201526200564a916001600160a01b031690636ae23fb59060240162004f2c565b60405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b158015620056aa575f80fd5b505af1158015620056bd573d5f803e3d5ffd5b50506040515f81527fd39ad370883b0cd1a8172b5b006a3ebcaaf65183008c91ffd7655afb74174e579250602001905060405180910390a16020546040517f543f66a40000000000000000000000000000000000000000000000000000000081525f60048201526001600160a01b039091169063543f66a4906024015f604051808303815f87803b15801562005751575f80fd5b505af115801562005764573d5f803e3d5ffd5b505060208054604080517f02befd24000000000000000000000000000000000000000000000000000000008152905162004f4895506001600160a01b0390921693506302befd249260048083019391928290030181865afa1580156200247b573d5f803e3d5ffd5b6020546040517f543f66a400000000000000000000000000000000000000000000000000000000815260016004820152620186a0916001600160a01b03169063543f66a4906024015f604051808303815f87803b1580156200582c575f80fd5b505af11580156200583f573d5f803e3d5ffd5b5050602054602480546040517f6ae23fb50000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015262005895955092169250636ae23fb5910162004f2c565b602480546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152928301849052169063095ea7b3906044016020604051808303815f875af115801562005900573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005926919062010102565b50604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401620027ab565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905260235460215491516370a0823160e01b81526001600160a01b0392831660048201529293505f929116906370a0823190602401602060405180830381865afa15801562005a18573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005a3e9190620100ea565b6027546025546023546020549394506001600160a01b03928316319391831631929081169163095ea7b3911662005a77886002620101c1565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801562005ac0573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005ae6919062010102565b5060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562005b47575f80fd5b505af115801562005b5a573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c9262005bb0928b92909116906028906200fe7c565b60405180910390a3602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b09362005bfb93908216928b929091169060289060040162010123565b5f604051808303815f87803b15801562005c13575f80fd5b505af115801562005c26573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562005c8a575f80fd5b505af115801562005c9d573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f9262005cf5928b92909116908a906028906201015b565b60405180910390a36020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b789362030d409362005d5f93918316928c929116908b906028906004016200ff55565b5f604051808303818588803b15801562005d77575f80fd5b505af115801562005d8a573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063491cc7c2925060840190505f604051808303815f87803b15801562005df0575f80fd5b505af115801562005e03573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d9749062005e5190889060289062010261565b60405180910390a36020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692631becceb49262030d409262005eb392909116908990602890600401620101f1565b5f604051808303818588803b15801562005ecb575f80fd5b505af115801562005ede573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9550911692506370a082319150602401602060405180830381865afa15801562005f30573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005f569190620100ea565b6027546025549192506001600160a01b039081163191163162005f9262005f7f896002620101c1565b62005f8b90886200ff02565b846200b482565b62005fb262005fa662030d406002620101c1565b6200228490876200ff02565b6200425a62005fc662030d406002620101c1565b62000b719086620101db565b6023546021546040516370a0823160e01b81526001600160a01b039182166004820152620186a0925f9216906370a0823190602401602060405180830381865afa15801562006023573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620060499190620100ea565b9050620060575f826200b482565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526024810185905291169063095ea7b3906044016020604051808303815f875af1158015620060c3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620060e9919062010102565b506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562006176575f80fd5b505af115801562006189573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c92620061df928892909116906028906200fe7c565b60405180910390a3602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b0936200622a939082169288929091169060289060040162010123565b5f604051808303815f87803b15801562006242575f80fd5b505af115801562006255573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa158015620062a6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620062cc9190620100ea565b9050620062da83826200b482565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201525f9291909116906370a0823190602401602060405180830381865afa15801562006329573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200634f9190620100ea565b905062001a3f84602c5462000b719190620101db565b6060601b805480602002602001604051908101604052809291908181526020015f905b8282101562002a91578382905f5260205f2090600202016040518060400160405290815f82018054620063bb906200fd3e565b80601f0160208091040260200160405190810160405280929190818152602001828054620063e9906200fd3e565b8015620064385780601f106200640e5761010080835404028352916020019162006438565b820191905f5260205f20905b8154815290600101906020018083116200641a57829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015620064bc57602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116200647d5790505b5050505050815250508152602001906001019062006388565b620186a0737109709ecfa91a80626ff3989d68f67f5b1dd12d63f28dceb37f84222ac70000000000000000000000000000000000000000000000000000000062006521600185620101db565b60405160248101919091526044810185905260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b90921682526200657d916004016200ff41565b5f604051808303815f87803b15801562006595575f80fd5b505af1158015620065a8573d5f803e3d5ffd5b50506020546026546001600160a01b03918216935063282906ed925084911662004a8b600183620101db565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905260235460215491516370a0823160e01b81526001600160a01b0392831660048201529293505f929116906370a0823190602401602060405180830381865afa1580156200666d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620066939190620100ea565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa158015620066e3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620067099190620100ea565b6027546023546020549293506001600160a01b0391821631929082169163095ea7b391166200673a886002620101c1565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801562006783573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620067a9919062010102565b5060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b1580156200680a575f80fd5b505af11580156200681d573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f9262006875928b92909116908a906028906201015b565b60405180910390a36020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b7893620068db93908216928b92909116908a906028906004016200ff55565b5f604051808303815f87803b158015620068f3575f80fd5b505af115801562006906573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b1580156200696a575f80fd5b505af11580156200697d573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f92620069d5928b92909116908a906028906201015b565b60405180910390a36020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b789362030d409362006a3f93918316928c929116908b906028906004016200ff55565b5f604051808303818588803b15801562006a57575f80fd5b505af115801562006a6a573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9550911692506370a082319150602401602060405180830381865afa15801562006abc573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062006ae29190620100ea565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa15801562006b32573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062006b589190620100ea565b6027549091506001600160a01b03163162006b7962005f7f896002620101c1565b62006b9662006b8a896002620101c1565b620022849087620101db565b6200425a62000b7162030d40866200ff02565b6027546026546040516001600160a01b039182166024820152620186a0929190911631905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562006c89575f80fd5b505af115801562006c9c573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f9062006cee9087905f9087906028906201015b565b60405180910390a36020546026546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263397e375c92879262006d4f92909116908390879060289060040162010198565b5f604051808303818588803b15801562006d67575f80fd5b505af115801562006d7a573d5f803e3d5ffd5b50506027546001600160a01b031631925062001a3f915062000b71905085856200ff02565b602654604080516001600160a01b0392831660248083018290528351808403909101815260449092018352602080830180516001600160e01b0316630427d73b60e51b1790525492517f744b9b8b000000000000000000000000000000000000000000000000000000008152620186a09492939092169163744b9b8b91859162006e3291908690602890600401620101f1565b5f604051808303818588803b15801562006e4a575f80fd5b505af115801562006e5d573d5f803e3d5ffd5b50506040805160048082526024820183526020820180516001600160e01b03167f394836a40000000000000000000000000000000000000000000000000000000017905291517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d955063f28dceb39450620028ce93509091016200ff41565b6060601a805480602002602001604051908101604052809291908181526020015f905b8282101562002a91578382905f5260205f2001805462006f36906200fd3e565b80601f016020809104026020016040519081016040528092919081815260200182805462006f64906200fd3e565b801562006fb35780601f1062006f895761010080835404028352916020019162006fb3565b820191905f5260205f20905b81548152906001019060200180831162006f9557829003601f168201915b50505050508152602001906001019062006f16565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529190526020810180516001600160e01b0316630427d73b60e51b1790529050737109709ecfa91a80626ff3989d68f67f5b1dd12d63f28dceb37f84222ac7000000000000000000000000000000000000000000000000000000006200705c8560016200ff02565b60405160248101919091526044810186905260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b9092168252620070b8916004016200ff41565b5f604051808303815f87803b158015620070d0575f80fd5b505af1158015620070e3573d5f803e3d5ffd5b50506020546026546001600160a01b03918216935063397e375c92508591166200710f8260016200ff02565b8560286040518663ffffffff1660e01b815260040162001dd9949392919062010198565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525460235491517f4f27eb2a0000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152600160248201529293501690634f27eb2a906044015f604051808303815f87803b158015620071e4575f80fd5b505af1158015620071f7573d5f803e3d5ffd5b50506020546040517f543f66a4000000000000000000000000000000000000000000000000000000008152600160048201526001600160a01b03909116925063543f66a491506024015f604051808303815f87803b15801562007258575f80fd5b505af11580156200726b573d5f803e3d5ffd5b50506023546020546001600160a01b03918216935063095ea7b392501662007295856002620101c1565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af1158015620072de573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062007304919062010102565b50602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b09362007348939082169288929091169060289060040162010123565b5f604051808303815f87803b15801562007360575f80fd5b505af115801562007373573d5f803e3d5ffd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b78945062030d4093620073d893811692899291169088906028906004016200ff55565b5f604051808303818588803b158015620073f0575f80fd5b505af115801562007403573d5f803e3d5ffd5b505050505062003a128260026200741b9190620101c1565b6023546021546040516370a0823160e01b81526001600160a01b0391821660048201529116906370a0823190602401602060405180830381865afa15801562007466573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200748c9190620100ea565b6200b482565b602654604080516001600160a01b0392831660248083018290528351808403909101815260449092018352602080830180516001600160e01b0316630427d73b60e51b1790525492517f282906ed000000000000000000000000000000000000000000000000000000008152620186a09492939092169163282906ed91859162007525919083906028906004016200ff18565b5f604051808303818588803b1580156200753d575f80fd5b505af115801562007550573d5f803e3d5ffd5b505060405162030d40602482015260448101869052737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392507fa458261b0000000000000000000000000000000000000000000000000000000091506064015b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b9092168252620045ab916004016200ff41565b604051630618f58760e51b81527f7671265e0000000000000000000000000000000000000000000000000000000060048201525f90737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b1580156200765e575f80fd5b505af115801562007671573d5f803e3d5ffd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c9350859262000d5d9216906028906004016200feb2565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529190526020810180516001600160e01b0316630427d73b60e51b1790529050737109709ecfa91a80626ff3989d68f67f5b1dd12d63f28dceb37f84222ac7000000000000000000000000000000000000000000000000000000006200775c600186620101db565b60405160248101919091526044810186905260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b9092168252620077b8916004016200ff41565b5f604051808303815f87803b158015620077d0575f80fd5b505af1158015620077e3573d5f803e3d5ffd5b50506020546026546001600160a01b03918216935063397e375c92508591166200710f600183620101db565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f543f66a4000000000000000000000000000000000000000000000000000000008152600160048201529192506001600160a01b03169063543f66a4906024015f604051808303815f87803b158015620078b5575f80fd5b505af1158015620078c8573d5f803e3d5ffd5b50506020546040517f4f27eb2a0000000000000000000000000000000000000000000000000000000081525f6004820181905260248201526001600160a01b039091169250634f27eb2a91506044015f604051808303815f87803b1580156200792f575f80fd5b505af115801562007942573d5f803e3d5ffd5b5050604080515f60248083019190915282518083039091018152604490910182526020810180516001600160e01b03167f33e23dbb0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250620079e891906004016200ff41565b5f604051808303815f87803b15801562007a00575f80fd5b505af115801562007a13573d5f803e3d5ffd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c9350869262007a6a9216906028906004016200feb2565b5f604051808303818588803b15801562007a82575f80fd5b505af115801562007a95573d5f803e3d5ffd5b5050604080515f60248083019190915282518083039091018152604490910182526020810180516001600160e01b03167f33e23dbb0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d945063f28dceb3935062007b3b92506004016200ff41565b5f604051808303815f87803b15801562007b53575f80fd5b505af115801562007b66573d5f803e3d5ffd5b50506020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063282906ed9350869262007bbf92169083906028906004016200ff18565b5f604051808303818588803b15801562007bd7575f80fd5b505af115801562007bea573d5f803e3d5ffd5b5050604080515f60248083019190915282518083039091018152604490910182526020810180516001600160e01b03167f33e23dbb0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d945063f28dceb3935062007c9092506004016200ff41565b5f604051808303815f87803b15801562007ca8575f80fd5b505af115801562007cbb573d5f803e3d5ffd5b50506020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063744b9b8b9350869262007d149216908690602890600401620101f1565b5f604051808303818588803b15801562007d2c575f80fd5b505af115801562007d3f573d5f803e3d5ffd5b50506040515f6024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392507f33e23dbb000000000000000000000000000000000000000000000000000000009150604401620075aa565b6060601d805480602002602001604051908101604052809291908181526020015f905b8282101562002a91575f8481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801562007e6057602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841162007e215790505b5050505050815250508152602001906001019062007db7565b6020546026546040517f726ac97c000000000000000000000000000000000000000000000000000000008152620186a0926001600160a01b039081169263726ac97c92859262007ed19216906028906004016200feb2565b5f604051808303818588803b15801562007ee9575f80fd5b505af115801562007efc573d5f803e3d5ffd5b50506040805160048082526024820183526020820180516001600160e01b03167f394836a40000000000000000000000000000000000000000000000000000000017905291517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d955063f28dceb394506200764693509091016200ff41565b60208054604080517fa2ba19340000000000000000000000000000000000000000000000000000000081529051620186a0935f936002936001600160a01b039091169263a2ba1934926004808401939192918290030181865afa15801562007ffc573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620080229190620100ea565b6200802e919062010228565b67ffffffffffffffff8111156200804957620080496200ffad565b6040519080825280601f01601f19166020018201604052801562008074576020820181803683370190505b5060208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa158015620080d8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620080fe9190620100ea565b6200810a919062010228565b620081179060016200ff02565b67ffffffffffffffff8111156200813257620081326200ffad565b6040519080825280601f01601f1916602001820160405280156200815d576020820181803683370190505b50602a906200816d908262010022565b506023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526024810185905291169063095ea7b3906044016020604051808303815f875af1158015620081da573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062008200919062010102565b50604051630618f58760e51b81527f386691aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e09060240162001ea8565b6026546040516001600160a01b0390911660248201526001905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156200830c575f80fd5b505af11580156200831f573d5f803e3d5ffd5b50506020546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063397e375c9150849062001dd9905f908390879060289060040162010198565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015620083e4575f80fd5b505af1158015620083f7573d5f803e3d5ffd5b50506020546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063282906ed9150839062000d5d905f9083906028906004016200ff18565b6026546040516001600160a01b0390911660248201526001905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562008500575f80fd5b505af115801562008513573d5f803e3d5ffd5b50506020546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063744b9b8b9150849062001dd9905f908690602890600401620101f1565b6027546026546040516001600160a01b039182166024820152620186a0929190911631905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562008649575f80fd5b505af11580156200865c573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f90620086ae9087905f9087906028906201015b565b60405180910390a36020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263744b9b8b92879262006d4f92909116908690602890600401620101f1565b60208054604080517fa2ba19340000000000000000000000000000000000000000000000000000000081529051620186a0935f936002936001600160a01b039091169263a2ba1934926004808401939192918290030181865afa15801562008777573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200879d9190620100ea565b620087a9919062010228565b67ffffffffffffffff811115620087c457620087c46200ffad565b6040519080825280601f01601f191660200182016040528015620087ef576020820181803683370190505b5060208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562008853573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620088799190620100ea565b62008885919062010228565b620088929060016200ff02565b67ffffffffffffffff811115620088ad57620088ad6200ffad565b6040519080825280601f01601f191660200182016040528015620088d8576020820181803683370190505b50602a90620088e8908262010022565b505f60286002018054620088fc906200fd3e565b83516200890a92506200ff02565b60208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290519394505f936001600160a01b039092169263a2ba1934926004808401938290030181865afa1580156200896c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620089929190620100ea565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f386691aa0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162008a3b916004016200ff41565b5f604051808303815f87803b15801562008a53575f80fd5b505af115801562008a66573d5f803e3d5ffd5b50506020546026546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063397e375c9350889262008ac19216908390899060289060040162010198565b5f604051808303818588803b15801562008ad9575f80fd5b505af115801562008aec573d5f803e3d5ffd5b505050505050505050565b6027546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201526509184e72a00090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562008b6f575f80fd5b505af115801562008b82573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562008be1575f80fd5b505af115801562008bf4573d5f803e3d5ffd5b50506020546040517f7c744253000000000000000000000000000000000000000000000000000000008152600481018590526001600160a01b039091169250637c74425391506024016200281b565b6026546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015b5f604051808303815f87803b15801562008d16575f80fd5b505af115801562008d29573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d9749062008d7790859060289062010261565b60405180910390a36020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692631becceb4926200281b929116908590602890600401620101f1565b6060601c805480602002602001604051908101604052809291908181526020015f905b8282101562002a91575f8481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801562008e9f57602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841162008e605790505b5050505050815250508152602001906001019062008df6565b60208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290515f936002936001600160a01b03169263a2ba193492600480830193928290030181865afa15801562008f19573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062008f3f9190620100ea565b62008f4b919062010228565b67ffffffffffffffff81111562008f665762008f666200ffad565b6040519080825280601f01601f19166020018201604052801562008f91576020820181803683370190505b5060208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562008ff5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200901b9190620100ea565b62009027919062010228565b620090349060016200ff02565b67ffffffffffffffff8111156200904f576200904f6200ffad565b6040519080825280601f01601f1916602001820160405280156200907a576020820181803683370190505b50602a906200908a908262010022565b50604051630618f58760e51b81527f386691aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b158015620090f6575f80fd5b505af115801562009109573d5f803e3d5ffd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb493506200281b92909116908590602890600401620101f1565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a06024820181905292919091169063095ea7b3906044016020604051808303815f875af1158015620091d5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620091fb919062010102565b506025546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200926c575f80fd5b505af11580156200927f573d5f803e3d5ffd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a59042791506024015f604051808303815f87803b158015620092e3575f80fd5b505af1158015620092f6573d5f803e3d5ffd5b5050604051630618f58760e51b81527fac2175f1000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e0915060240162004805565b60606019805480602002602001604051908101604052809291908181526020015f905b8282101562002a91578382905f5260205f2001805462009395906200fd3e565b80601f0160208091040260200160405190810160405280929190818152602001828054620093c3906200fd3e565b8015620094125780601f10620093e85761010080835404028352916020019162009412565b820191905f5260205f20905b815481529060010190602001808311620093f457829003601f168201915b50505050508152602001906001019062009375565b602354602054620186a09161c350916001600160a01b039182169163095ea7b3911662009456856002620101c1565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af11580156200949f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620094c5919062010102565b50602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b09362009509939082169288929091169060289060040162010123565b5f604051808303815f87803b15801562009521575f80fd5b505af115801562009534573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507f8afe4db400000000000000000000000000000000000000000000000000000000905062030d406200958685826200ff02565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b9092168252620095e0916004016200ff41565b5f604051808303815f87803b158015620095f8575f80fd5b505af11580156200960b573d5f803e3d5ffd5b50506020546001600160a01b0316915063102614b09050620096318362030d406200ff02565b6026546023546040516001600160e01b031960e086901b16815262001dd9926001600160a01b0390811692899291169060289060040162010123565b6008545f9060ff161562009685575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c6564000000000000000000000000000000000000000000000000000060248301525f9163667f9d7090604401602060405180830381865afa15801562009714573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200973a9190620100ea565b1415905090565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052602880547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000017905551630618f58760e51b81527f19b4bff2000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401620090de565b602654604080516001600160a01b0392831660248083018290528351808403909101815260449092018352602080830180516001600160e01b0316630427d73b60e51b1790525492517f1becceb4000000000000000000000000000000000000000000000000000000008152919361c350931691631becceb491620098ab918690602890600401620101f1565b5f604051808303815f87803b158015620098c3575f80fd5b505af1158015620098d6573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507f8afe4db400000000000000000000000000000000000000000000000000000000905062030d406200992885826200ff02565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b909216825262009982916004016200ff41565b5f604051808303815f87803b1580156200999a575f80fd5b505af1158015620099ad573d5f803e3d5ffd5b50506020546001600160a01b03169150631becceb49050620099d38362030d406200ff02565b6026546040516001600160e01b031960e085901b16815262001dd9916001600160a01b0316908790602890600401620101f1565b604051630618f58760e51b81527f7671265e0000000000000000000000000000000000000000000000000000000060048201525f90737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b15801562009a74575f80fd5b505af115801562009a87573d5f803e3d5ffd5b50506020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063282906ed9350859262000d5d92169083906028906004016200ff18565b620186a05f62009af5600162030d40620101db565b6023546020549192506001600160a01b039081169163095ea7b3911662009b1e856002620101c1565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801562009b67573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062009b8d919062010102565b50602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b09362009bd1939082169288929091169060289060040162010123565b5f604051808303815f87803b15801562009be9575f80fd5b505af115801562009bfc573d5f803e3d5ffd5b505060405162030d40602482015260448101849052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507fa458261b000000000000000000000000000000000000000000000000000000009060640162001d20565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f6024820181905292919091169063095ea7b3906044016020604051808303815f875af115801562009cca573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062009cf0919062010102565b50604051630618f58760e51b81527f951e19ed000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e09060240162004805565b6020546040517f7c7442530000000000000000000000000000000000000000000000000000000081525f60048201526001600160a01b0390911690637c744253906024015f604051808303815f87803b15801562009da5575f80fd5b505af115801562009db8573d5f803e3d5ffd5b50506026546040516001600160a01b039091166024820152620186a092505f915060440160408051601f19818403018152919052602080820180516001600160e01b0316630427d73b60e51b17905260235490549192506001600160a01b039081169163095ea7b3911662009e2f856002620101c1565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801562009e78573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062009e9e919062010102565b50602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b09362009ee2939082169288929091169060289060040162010123565b5f604051808303815f87803b15801562009efa575f80fd5b505af115801562009f0d573d5f803e3d5ffd5b5050604051630618f58760e51b81527f394836a4000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b15801562009f7b575f80fd5b505af115801562009f8e573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b0945062009fd393928316928892169060289060040162010123565b5f604051808303815f87803b15801562009feb575f80fd5b505af115801562009ffe573d5f803e3d5ffd5b5050604051630618f58760e51b81527f394836a4000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b1580156200a06c575f80fd5b505af11580156200a07f573d5f803e3d5ffd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb4935062000eb292909116908590602890600401620101f1565b6023546021546040516370a0823160e01b81526001600160a01b039182166004820152620186a0925f9216906370a0823190602401602060405180830381865afa1580156200a129573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a14f9190620100ea565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa1580156200a19f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a1c59190620100ea565b6027546023546020549293506001600160a01b0391821631929082169163095ea7b391166200a1f6876002620101c1565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af11580156200a23f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a265919062010102565b5060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b1580156200a2c6575f80fd5b505af11580156200a2d9573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c926200a32f928a92909116906028906200fe7c565b60405180910390a3602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b0936200a37a93908216928a929091169060289060040162010123565b5f604051808303815f87803b1580156200a392575f80fd5b505af11580156200a3a5573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b1580156200a409575f80fd5b505af11580156200a41c573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c926200a472928a92909116906028906200fe7c565b60405180910390a3602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b09362030d40936200a4c193918316928b9291169060289060040162010123565b5f604051808303818588803b1580156200a4d9575f80fd5b505af11580156200a4ec573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9550911692506370a082319150602401602060405180830381865afa1580156200a53e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a5649190620100ea565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa1580156200a5b4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a5da9190620100ea565b6027549091506001600160a01b0316316200a5fb62005f7f886002620101c1565b6200a60c62006b8a886002620101c1565b62001e0462000b7162030d40866200ff02565b6026546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f543f66a4000000000000000000000000000000000000000000000000000000008152600160048201529192506001600160a01b03169063543f66a4906024015f604051808303815f87803b1580156200a6c0575f80fd5b505af11580156200a6d3573d5f803e3d5ffd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a40162008cfe565b606060158054806020026020016040519081016040528092919081815260200182805480156200231057602002820191905f5260205f209081546001600160a01b03168152600190910190602001808311620022f1575050505050905090565b602480546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a093810184905291169063095ea7b3906044016020604051808303815f875af11580156200a821573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a847919062010102565b50604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156200a8b2575f80fd5b505af11580156200a8c5573d5f803e3d5ffd5b50506026546025546024546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c926200a91b928792909116906028906200fe7c565b60405180910390a3602054602654602454604051630102614b60e41b81526001600160a01b039384169363102614b0936200281b939082169287929091169060289060040162010123565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052602354905491517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101869052929350169063095ea7b3906044016020604051808303815f875af11580156200aa1d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200aa43919062010102565b506025546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200aab4575f80fd5b505af11580156200aac7573d5f803e3d5ffd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a59042791506024015f604051808303815f87803b1580156200ab2b575f80fd5b505af11580156200ab3e573d5f803e3d5ffd5b5050604051630618f58760e51b81527fac2175f1000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e0915060240162001ea8565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f543f66a4000000000000000000000000000000000000000000000000000000008152600160048201529192506001600160a01b03169063543f66a4906024015f604051808303815f87803b1580156200ac40575f80fd5b505af11580156200ac53573d5f803e3d5ffd5b50506023546020546001600160a01b03918216935063095ea7b39250166200ac7d856002620101c1565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af11580156200acc6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200acec919062010102565b50602354604080516001600160a01b0390921660248084019190915281518084039091018152604490920181526020820180516001600160e01b03167f33e23dbb00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916200ad9b91906004016200ff41565b5f604051808303815f87803b1580156200adb3575f80fd5b505af11580156200adc6573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b094506200ae0b93928316928892169060289060040162010123565b5f604051808303815f87803b1580156200ae23575f80fd5b505af11580156200ae36573d5f803e3d5ffd5b5050602354604080516001600160a01b0390921660248084019190915281518084039091018152604490920181526020820180516001600160e01b03167f33e23dbb00000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925062001ea891906004016200ff41565b6026546040516001600160a01b039091166024820152620186a09061c350905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052602354905491517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101879052929350169063095ea7b3906044016020604051808303815f875af11580156200afa3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200afc9919062010102565b506040515f602482015260448101839052737109709ecfa91a80626ff3989d68f67f5b1dd12d9063f28dceb3907f8afe4db400000000000000000000000000000000000000000000000000000000906064016200443c565b6020546040517f7c7442530000000000000000000000000000000000000000000000000000000081526509184e72a0006004820181905291620186a0916001600160a01b0390911690637c744253906024015f604051808303815f87803b1580156200b08b575f80fd5b505af11580156200b09e573d5f803e3d5ffd5b50506027546025546020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b039485163196509284163194509083169263726ac97c9287926200b1059216906028906004016200feb2565b5f604051808303818588803b1580156200b11d575f80fd5b505af11580156200b130573d5f803e3d5ffd5b50506020546001600160a01b0316925063282906ed91506200b155905086866200ff02565b6026546040516001600160e01b031960e085901b1681526200b189916001600160a01b03169088906028906004016200ff18565b5f604051808303818588803b1580156200b1a1575f80fd5b505af11580156200b1b4573d5f803e3d5ffd5b50506027546025546001600160a01b03918216319450163191506200b1e29050866200226c876002620101c1565b62000d8886620022a0876002620101c1565b6027546020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039182166084820152620186a092919091163190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b1580156200b28e575f80fd5b505af11580156200b2a1573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c906200b2f19086905f906028906200fe7c565b60405180910390a36020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263282906ed92869262000b21929091169083906028906004016200ff18565b6020546026546040517f282906ed000000000000000000000000000000000000000000000000000000008152620186a0926001600160a01b039081169263282906ed9285926200b3aa92169083906028906004016200ff18565b5f604051808303818588803b1580156200b3c2575f80fd5b505af11580156200b3d5573d5f803e3d5ffd5b50506040805162030d406024820152604480820187905282518083039091018152606490910182526020810180516001600160e01b03167fa458261b0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d945063f28dceb3935062009a5c92506004016200ff41565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c54906044015f6040518083038186803b1580156200b4ec575f80fd5b505afa15801562000d88573d5f803e3d5ffd5b5f6200b50a6200f90d565b6200b5178484836200b5ec565b9150505b92915050565b6040517fa59828850000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063a5982885906024015b5f6040518083038186803b1580156200b586575f80fd5b505afa15801562001c2c573d5f803e3d5ffd5b6040517f0c9fd5810000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90630c9fd581906024016200b56f565b5f806200b5fa85846200b66c565b90506200b6616040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f787900000081525082866040516020016200b64b92919062010289565b604051602081830303815290604052856200b679565b9150505b9392505050565b5f6200b66583836200b6ac565b60c0810151515f90156200b6a0576200b69884848460c001516200b6ca565b90506200b665565b6200b69884846200b87e565b5f6200b6b983836200b96f565b6200b665838360200151846200b679565b5f806200b6d66200b97c565b90505f6200b6e586836200ba50565b90505f6200b6fd82606001518360200151856200bf08565b90505f6200b70e838389896200c12d565b90505f6200b71c826200d06f565b602081015181519192509060030b156200b794578982604001516040516020016200b749929190620102c3565b60408051601f19818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526200b78b916004016200ff41565b60405180910390fd5b5f6200b7d86040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a2000000000000000000000008152508360016200d245565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d906200b82d9084906004016200ff41565b602060405180830381865afa1580156200b849573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b86f919062010328565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081525f908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc925906200b8d49087906004016200ff41565b5f60405180830381865afa1580156200b8ef573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200b91891908101906201041b565b90505f6200b94982856040516020016200b93492919062010451565b6040516020818303038152906040526200d44a565b90506001600160a01b0381166200b5175784846040516020016200b74992919062010469565b62003a1282825f6200d45b565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c906200ba05908490600401620104fd565b5f60405180830381865afa1580156200ba20573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200ba49919081019062010545565b9250505090565b6200ba836040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b5f737109709ecfa91a80626ff3989d68f67f5b1dd12d90506200bace6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6200bad9856200d568565b60208201525f6200baea866200d961565b90505f836001600160a01b031663d930a0e66040518163ffffffff1660e01b81526004015f60405180830381865afa1580156200bb29573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200bb52919081019062010545565b868385602001516040516020016200bb6e94939291906201058f565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291505f906001600160a01b038616906360f9bb11906200bbc79085906004016200ff41565b5f60405180830381865afa1580156200bbe2573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200bc0b919081019062010545565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f6906200bc5590849060040162010667565b602060405180830381865afa1580156200bc71573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200bc97919062010102565b6200bcaf57816040516020016200b7499190620106ba565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906200bcf690849060040162010740565b5f60405180830381865afa1580156200bd11573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200bd3a919081019062010545565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f6906200bd8390849060040162010793565b602060405180830381865afa1580156200bd9f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200bdc5919062010102565b156200be5c576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906200be1290849060040162010793565b5f60405180830381865afa1580156200be2d573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200be56919081019062010545565b60408501525b846001600160a01b03166349c4fac882865f01516040516020016200be829190620107e6565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016200beb092919062010846565b5f60405180830381865afa1580156200becb573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200bef4919081019062010545565b606085015250608083015250949350505050565b60408051600480825260a082019092526060915f9190816020015b60608152602001906001900390816200bf235790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250815f815181106200bf86576200bf866201086e565b60200260200101819052506040518060400160405280600381526020017f2d726c0000000000000000000000000000000000000000000000000000000000815250816001815181106200bfdd576200bfdd6201086e565b6020026020010181905250846040516020016200bffb91906201089b565b604051602081830303815290604052816002815181106200c020576200c0206201086e565b6020026020010181905250826040516020016200c03e9190620108fb565b604051602081830303815290604052816003815181106200c063576200c0636201086e565b60200260200101819052505f6200c07a826200d06f565b602080820151604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000818501908152825180840184525f80825290860152825180840190935290518252928101929092529192506200c10b906040805180820182525f80825260209182015281518083019092528451825280850190820152906200dbf6565b6200c12357856040516020016200b749919062010935565b9695505050505050565b60a08101516040805180820182525f80825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d90156200c17d565b511590565b6200c2f6578260200151156200c23c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a4016200b78b565b8260c00151156200c2f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a4016200b78b565b6040805160ff80825261200082019092525f91816020015b60608152602001906001900390816200c30e5790505090505f6040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806200c36b90620109b4565b935060ff16815181106200c383576200c3836201086e565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e37000000000000000000000000000000000000008152506040516020016200c3d69190620109d5565b6040516020818303038152906040528282806200c3f390620109b4565b935060ff16815181106200c40b576200c40b6201086e565b60200260200101819052506040518060400160405280600681526020017f6465706c6f7900000000000000000000000000000000000000000000000000008152508282806200c45a90620109b4565b935060ff16815181106200c472576200c4726201086e565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d650000000000000000000000000000000000008152508282806200c4c190620109b4565b935060ff16815181106200c4d9576200c4d96201086e565b602002602001018190525087602001518282806200c4f790620109b4565b935060ff16815181106200c50f576200c50f6201086e565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e7472616374506174680000000000000000000000000000000000008152508282806200c55e90620109b4565b935060ff16815181106200c576576200c5766201086e565b6020908102919091010152875182826200c59081620109b4565b935060ff16815181106200c5a8576200c5a86201086e565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e496400000000000000000000000000000000000000000000008152508282806200c5f790620109b4565b935060ff16815181106200c60f576200c60f6201086e565b60200260200101819052506200c625466200dc5c565b82826200c63281620109b4565b935060ff16815181106200c64a576200c64a6201086e565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c6500000000000000000000000000000000008152508282806200c69990620109b4565b935060ff16815181106200c6b1576200c6b16201086e565b6020026020010181905250868282806200c6cb90620109b4565b935060ff16815181106200c6e3576200c6e36201086e565b60209081029190910101528551156200c8165760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826200c73781620109b4565b935060ff16815181106200c74f576200c74f6201086e565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d906200c7a19089906004016200ff41565b5f60405180830381865afa1580156200c7bc573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200c7e5919081019062010545565b82826200c7f281620109b4565b935060ff16815181106200c80a576200c80a6201086e565b60200260200101819052505b8460200151156200c8f25760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826200c86281620109b4565b935060ff16815181106200c87a576200c87a6201086e565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806200c8c990620109b4565b935060ff16815181106200c8e1576200c8e16201086e565b60200260200101819052506200cad4565b6200c92b6200c1788660a001516040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6200c9c85760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826200c97181620109b4565b935060ff16815181106200c989576200c9896201086e565b60200260200101819052508460a001516040516020016200c9ab91906201089b565b6040516020818303038152906040528282806200c8c990620109b4565b8460c001511580156200ca0c5750604080890151815180830183525f808252602091820152825180840190935281518352908101908201526200ca0a90511590565b155b156200cad45760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826200ca5381620109b4565b935060ff16815181106200ca6b576200ca6b6201086e565b60200260200101819052506200ca81886200dd00565b6040516020016200ca9391906201089b565b6040516020818303038152906040528282806200cab090620109b4565b935060ff16815181106200cac8576200cac86201086e565b60200260200101819052505b604080860151815180830183525f808252602091820152825180840190935281518352908101908201526200cb0890511590565b6200cba85760408051808201909152600b81527f2d2d72656c617965724964000000000000000000000000000000000000000000602082015282826200cb4e81620109b4565b935060ff16815181106200cb66576200cb666201086e565b602002602001018190525084604001518282806200cb8490620109b4565b935060ff16815181106200cb9c576200cb9c6201086e565b60200260200101819052505b6060850151156200ccd35760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826200cbf481620109b4565b935060ff16815181106200cc0c576200cc0c6201086e565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e8906024015f60405180830381865afa1580156200cc79573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200cca2919081019062010545565b82826200ccaf81620109b4565b935060ff16815181106200ccc7576200ccc76201086e565b60200260200101819052505b60e085015151156200cd865760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826200cd2081620109b4565b935060ff16815181106200cd38576200cd386201086e565b60200260200101819052506200cd558560e001515f01516200dc5c565b82826200cd6281620109b4565b935060ff16815181106200cd7a576200cd7a6201086e565b60200260200101819052505b60e085015160200151156200ce3d5760408051808201909152600a81527f2d2d676173507269636500000000000000000000000000000000000000000000602082015282826200cdd681620109b4565b935060ff16815181106200cdee576200cdee6201086e565b60200260200101819052506200ce0c8560e00151602001516200dc5c565b82826200ce1981620109b4565b935060ff16815181106200ce31576200ce316201086e565b60200260200101819052505b60e085015160400151156200cef45760408051808201909152600e81527f2d2d6d6178466565506572476173000000000000000000000000000000000000602082015282826200ce8d81620109b4565b935060ff16815181106200cea5576200cea56201086e565b60200260200101819052506200cec38560e00151604001516200dc5c565b82826200ced081620109b4565b935060ff16815181106200cee8576200cee86201086e565b60200260200101819052505b60e085015160600151156200cfab5760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826200cf4481620109b4565b935060ff16815181106200cf5c576200cf5c6201086e565b60200260200101819052506200cf7a8560e00151606001516200dc5c565b82826200cf8781620109b4565b935060ff16815181106200cf9f576200cf9f6201086e565b60200260200101819052505b5f8160ff1667ffffffffffffffff8111156200cfcb576200cfcb6200ffad565b6040519080825280602002602001820160405280156200d00057816020015b60608152602001906001900390816200cfea5790505b5090505f5b8260ff168160ff1610156200d06057838160ff16815181106200d02c576200d02c6201086e565b6020026020010151828260ff16815181106200d04c576200d04c6201086e565b60209081029190910101526001016200d005565b5093505050505b949350505050565b6200d09660405180606001604052805f60030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d925f91849163d145736c916200d11d9186910162010a2e565b5f60405180830381865afa1580156200d138573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200d161919081019062010545565b90505f6200d17086836200e817565b90505f846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b81526004016200d1a191906200fc2c565b5f604051808303815f875af11580156200d1bd573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200d1e6919081019062010a76565b805190915060030b158015906200d2005750602081015151155b80156200d2105750604081015151155b156200c12357815f815181106200d22b576200d22b6201086e565b60200260200101516040516020016200b749919062010b31565b60605f6200d279856040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925286518252808701908201529091506200d2b19082905b906200e97f565b156200d418575f6200d332826200d32b846200d3246200d2f78a6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925282518252918201519181019190915290565b906200e9a7565b906200ea0b565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200d3979082906200e97f565b156200d40357604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d400905b82906200ea97565b90505b6200d40e816200eabe565b925050506200b665565b82156200d4345784846040516020016200b74992919062010d10565b505060408051602081019091525f81526200b665565b5f808251602084015ff09392505050565b8160a00151156200d46b57505050565b5f6200d4798484846200eb29565b90505f6200d487826200d06f565b602081015181519192509060030b1580156200d5245750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d524906040805180820182525f808252602091820152815180830190925284518252808501908201526200d2aa565b156200d53257505050505050565b604082015151156200d5555781604001516040516020016200b749919062010d9b565b806040516020016200b749919062010df4565b60605f6200d59c836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200d602905b82906200dbf6565b156200d67557604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200b665906200d66f9083906200f11d565b6200eabe565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d6d8905b82906200f1af565b6001036200d7a957604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d740906200d3f8565b50604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200b665906200d66f905b83906200ea97565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d809906200d5fa565b156200d94857604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201819052845180860190955292518452830152906200d8739083906200f255565b90505f81600183516200d8879190620101db565b815181106200d89a576200d89a6201086e565b602002602001015190506200d93f6200d66f6200d9126040518060400160405280600581526020017f2e6a736f6e0000000000000000000000000000000000000000000000000000008152506040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f80825260209182015281518083019092528551825280860190820152906200f11d565b95945050505050565b826040516020016200b749919062010e4d565b50919050565b60605f6200d995836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200d9f8906200d5fa565b156200da09576200b665816200eabe565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200da69906200d6d0565b6001036200dad657604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200b665906200d66f906200d7a1565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200db36906200d5fa565b156200d94857604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201819052845180860190955292518452830152906200dba09083906200f255565b90506001815111156200dbe25780600282516200dbbe9190620101db565b815181106200dbd1576200dbd16201086e565b602002602001015192505050919050565b50826040516020016200b749919062010e4d565b805182515f9111156200dc0b57505f6200b51b565b8151835160208501515f92916200dc22916200ff02565b6200dc2e9190620101db565b9050826020015181036200dc475760019150506200b51b565b82516020840151819020912014905092915050565b60605f6200dc6a836200f311565b60010190505f8167ffffffffffffffff8111156200dc8c576200dc8c6200ffad565b6040519080825280601f01601f1916602001820160405280156200dcb7576020820181803683370190505b5090508181016020015b5f19017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846200dcc157509392505050565b604081810151815180830183525f808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e53454400000000000000000000000000000000000000000000818401908152855180870187528381528401929092528451808601909552518452908301526060916200dd8d905b82906200f3f9565b156200ddce57505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200de2e906200dd85565b156200de6f57505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200decf906200dd85565b156200df1057505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200df70906200dd85565b806200dfd75750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200dfd7906200dd85565b156200e01857505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e078906200dd85565b806200e0df5750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e0df906200dd85565b156200e12057505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e180906200dd85565b806200e1e75750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e1e7906200dd85565b156200e22857505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e288906200dd85565b806200e2ef5750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e2ef906200dd85565b156200e33057505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e390906200dd85565b156200e3d157505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e431906200dd85565b156200e47257505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e4d2906200dd85565b156200e51357505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e573906200dd85565b156200e5b457505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e614906200dd85565b156200e65557505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e6b5906200dd85565b806200e71c5750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e71c906200dd85565b156200e75d57505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e7bd906200dd85565b156200e7fe57505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b604080840151845191516200b749929060200162010f1f565b6060805f5b84518110156200e8ad57818582815181106200e83c576200e83c6201086e565b60200260200101516040516020016200e85792919062010451565b6040516020818303038152906040529150600185516200e8789190620101db565b81146200e8a457816040516020016200e892919062011071565b60405160208183030381529060405291505b6001016200e81c565b50604080516003808252608082019092525f91816020015b60608152602001906001900390816200e8c557905050905083815f815181106200e8f3576200e8f36201086e565b60200260200101819052506040518060400160405280600281526020017f2d63000000000000000000000000000000000000000000000000000000000000815250816001815181106200e94a576200e94a6201086e565b602002602001018190525081816002815181106200e96c576200e96c6201086e565b6020908102919091010152949350505050565b60208083015183518351928401515f936200e99e92918491906200f40e565b14159392505050565b604080518082019091525f80825260208201525f6200e9d7845f01518560200151855f015186602001516200f541565b90508360200151816200e9eb9190620101db565b845185906200e9fc908390620101db565b90525060208401525090919050565b604080518082019091525f80825260208201528151835110156200ea315750816200b51b565b60208083015190840151600191146200ea595750815160208481015190840151829020919020145b80156200ea8f578251845185906200ea73908390620101db565b90525082516020850180516200ea8b9083906200ff02565b9052505b509192915050565b604080518082019091525f80825260208201526200eab78383836200f681565b5092915050565b60605f825f015167ffffffffffffffff8111156200eae0576200eae06200ffad565b6040519080825280601f01601f1916602001820160405280156200eb0b576020820181803683370190505b5090505f6020820190506200eab7818560200151865f01516200f735565b60605f6200eb366200b97c565b6040805160ff80825261200082019092529192505f9190816020015b60608152602001906001900390816200eb525790505090505f6040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806200ebaf90620109b4565b935060ff16815181106200ebc7576200ebc76201086e565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e33000000000000000000000000000000000000000000000000008152506040516020016200ec1a9190620110ab565b6040516020818303038152906040528282806200ec3790620109b4565b935060ff16815181106200ec4f576200ec4f6201086e565b60200260200101819052506040518060400160405280600881526020017f76616c69646174650000000000000000000000000000000000000000000000008152508282806200ec9e90620109b4565b935060ff16815181106200ecb6576200ecb66201086e565b6020026020010181905250826040516020016200ecd49190620108fb565b6040516020818303038152906040528282806200ecf190620109b4565b935060ff16815181106200ed09576200ed096201086e565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e7472616374000000000000000000000000000000000000000000008152508282806200ed5890620109b4565b935060ff16815181106200ed70576200ed706201086e565b60200260200101819052506200ed8787846200f7bd565b82826200ed9481620109b4565b935060ff16815181106200edac576200edac6201086e565b6020908102919091010152855151156200ee645760408051808201909152600b81527f2d2d7265666572656e6365000000000000000000000000000000000000000000602082015282826200ee0181620109b4565b935060ff16815181106200ee19576200ee196201086e565b60200260200101819052506200ee33865f0151846200f7bd565b82826200ee4081620109b4565b935060ff16815181106200ee58576200ee586201086e565b60200260200101819052505b8560800151156200eed95760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b0000000000000000602082015282826200eeb081620109b4565b935060ff16815181106200eec8576200eec86201086e565b60200260200101819052506200ef45565b84156200ef455760408051808201909152601281527f2d2d726571756972655265666572656e63650000000000000000000000000000602082015282826200ef2181620109b4565b935060ff16815181106200ef39576200ef396201086e565b60200260200101819052505b604086015151156200efec5760408051808201909152600d81527f2d2d756e73616665416c6c6f7700000000000000000000000000000000000000602082015282826200ef9281620109b4565b935060ff16815181106200efaa576200efaa6201086e565b602002602001018190525085604001518282806200efc890620109b4565b935060ff16815181106200efe0576200efe06201086e565b60200260200101819052505b8560600151156200f05c5760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d6573000000000000000000000000602082015282826200f03881620109b4565b935060ff16815181106200f050576200f0506201086e565b60200260200101819052505b5f8160ff1667ffffffffffffffff8111156200f07c576200f07c6200ffad565b6040519080825280602002602001820160405280156200f0b157816020015b60608152602001906001900390816200f09b5790505b5090505f5b8260ff168160ff1610156200f11157838160ff16815181106200f0dd576200f0dd6201086e565b6020026020010151828260ff16815181106200f0fd576200f0fd6201086e565b60209081029190910101526001016200f0b6565b50979650505050505050565b604080518082019091525f80825260208201528151835110156200f1435750816200b51b565b8151835160208501515f92916200f15a916200ff02565b6200f1669190620101db565b602084015190915060019082146200f188575082516020840151819020908220145b80156200f1a6578351855186906200f1a2908390620101db565b9052505b50929392505050565b5f80825f01516200f1d1855f01518660200151865f015187602001516200f541565b6200f1dd91906200ff02565b90505b835160208501516200f1f391906200ff02565b81116200eab757816200f20681620110de565b925050825f01516200f2418560200151836200f2239190620101db565b86516200f2319190620101db565b83865f015187602001516200f541565b6200f24d91906200ff02565b90506200f1e0565b60605f6200f26484846200f1af565b6200f2719060016200ff02565b67ffffffffffffffff8111156200f28c576200f28c6200ffad565b6040519080825280602002602001820160405280156200f2c157816020015b60608152602001906001900390816200f2ab5790505b5090505f5b81518110156200f309576200f2e06200d66f86866200ea97565b8282815181106200f2f5576200f2f56201086e565b60209081029190910101526001016200f2c6565b509392505050565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106200f35a577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106200f387576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106200f3a657662386f26fc10000830492506010015b6305f5e10083106200f3bf576305f5e100830492506008015b61271083106200f3d457612710830492506004015b606483106200f3e7576064830492506002015b600a83106200b51b5760010192915050565b5f6200f40683836200f800565b159392505050565b5f808584116200f53757602084116200f4d7575f84156200f4625760016200f438866020620101db565b6200f445906008620101c1565b6200f452906002620111f5565b6200f45e9190620101db565b1990505b83518116856200f47389896200ff02565b6200f47f9190620101db565b805190935082165b8181146200f4bf578784116200f4a457879450505050506200d067565b836200f4b08162011202565b9450508284511690506200f487565b6200f4cb87856200ff02565b9450505050506200d067565b8383206200f4e68588620101db565b6200f4f290876200ff02565b91505b8582106200f535578482208082036200f51f576200f51486846200ff02565b93505050506200d067565b6200f52c600184620101db565b9250506200f4f5565b505b5092949350505050565b5f83818685116200f66a57602085116200f610575f85156200f5965760016200f56c876020620101db565b6200f579906008620101c1565b6200f586906002620111f5565b6200f5929190620101db565b1990505b845181165f876200f5a88b8b6200ff02565b6200f5b49190620101db565b855190915083165b8281146200f601578186106200f5e6576200f5d88b8b6200ff02565b96505050505050506200d067565b856200f5f281620110de565b9650508386511690506200f5bc565b8596505050505050506200d067565b508383205f905b6200f6238689620101db565b82116200f668578583208082036200f64257839450505050506200d067565b6200f64f6001856200ff02565b93505081806200f65f90620110de565b9250506200f617565b505b6200f67687876200ff02565b979650505050505050565b604080518082019091525f80825260208201525f6200f6b1855f01518660200151865f015187602001516200f541565b6020808701805191860191909152519091506200f6cf9082620101db565b8352845160208601516200f6e491906200ff02565b81036200f6f4575f85526200f72c565b835183516200f70491906200ff02565b855186906200f715908390620101db565b90525083516200f72690826200ff02565b60208601525b50909392505050565b602081106200f77557815183526200f74f6020846200ff02565b92506200f75e6020836200ff02565b91506200f76d602082620101db565b90506200f735565b5f1981156200f7aa5760016200f78d836020620101db565b6200f79b90610100620111f5565b6200f7a79190620101db565b90505b9151835183169219169190911790915250565b60605f6200f7cc84846200ba50565b80516020808301516040519394506200f7e8939091016201121a565b60405160208183030381529060405291505092915050565b815181515f91908111156200f813575081515b602080850151908401515f5b838110156200f8e057825182518082146200f8a9575f1960208710156200f886576001846200f850896020620101db565b6200f85c91906200ff02565b6200f869906008620101c1565b6200f876906002620111f5565b6200f8829190620101db565b1990505b81811683821681810391146200f8a65797506200b51b9650505050505050565b50505b6200f8b66020866200ff02565b94506200f8c56020856200ff02565b935050506020816200f8d891906200ff02565b90506200f81f565b50845186516200c123919062011259565b610c32806201127c83390190565b61124c8062011eae83390190565b6040518060e001604052806060815260200160608152602001606081526020015f151581526020015f151581526020015f151581526020016200f94f6200f954565b905290565b6040518061010001604052805f151581526020015f15158152602001606081526020015f801916815260200160608152602001606081526020015f151581526020016200f94f60405180608001604052805f81526020015f81526020015f81526020015f81525090565b602080825282518282018190525f918401906040840190835b818110156200fa005783516001600160a01b03168352602093840193909201916001016200f9d7565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200fb38577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b8801810191908801905f5b818110156200fb1d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a85030183526200fb068486516200fa0b565b60209586019590945092909201916001016200fac9565b5091975050506020948501949290920191506001016200fa5f565b50929695505050505050565b5f8151808452602084019350602083015f5b828110156200fb805781516001600160e01b0319168652602095860195909101906001016200fb56565b5093949350505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200fb38577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051604087526200fbf760408801826200fa0b565b90506020820151915086810360208801526200fc1481836200fb44565b9650505060209384019391909101906001016200fbb0565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200fb38577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184526200fc8f8583516200fa0b565b945060209384019391909101906001016200fc52565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200fb38577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b03815116865260208101519050604060208701526200fd2760408701826200fb44565b95505060209384019391909101906001016200fccb565b600181811c908216806200fd5357607f821691505b6020821081036200d95b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f81546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b0360018301541660408401526002820160a060608501525f81546200fdd5816200fd3e565b8060a0880152600182165f81146200fdf657600181146200fe31576200fe64565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660c089015260c082151560051b89010193506200fe64565b845f5260205f205f5b838110156200fe5b5781548a820160c001526001909101906020016200fe3a565b890160c0019450505b50505060038401546080860152809250505092915050565b8381526001600160a01b0383166020820152608060408201525f608082015260a060608201525f6200d93f60a08301846200fd8b565b6001600160a01b0383168152604060208201525f6200d06760408301846200fd8b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156200b51b576200b51b6200fed5565b6001600160a01b0384168152826020820152606060408201525f6200d93f60608301846200fd8b565b602081525f6200b66560208301846200fa0b565b6001600160a01b03861681528460208201526001600160a01b038416604082015260a060608201525f6200ff8d60a08301856200fa0b565b82810360808401526200ffa181856200fd8b565b98975050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b601f82111562000b7857805f5260205f20601f840160051c81016020851015620100015750805b601f840160051c820191505b8181101562001c2c575f81556001016201000d565b815167ffffffffffffffff8111156201003f576201003f6200ffad565b62010057816201005084546200fd3e565b846200ffda565b6020601f8211600181146201008c575f8315620100745750848201515b5f19600385901b1c1916600184901b17845562001c2c565b5f84815260208120601f198516915b82811015620100bd57878501518255602094850194600190920191016201009b565b5084821015620100db57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f60208284031215620100fb575f80fd5b5051919050565b5f6020828403121562010113575f80fd5b815180151581146200b665575f80fd5b6001600160a01b03851681528360208201526001600160a01b0383166040820152608060608201525f6200c12360808301846200fd8b565b8481526001600160a01b0384166020820152608060408201525f6201018460808301856200fa0b565b82810360608401526200f67681856200fd8b565b6001600160a01b0385168152836020820152608060408201525f6201018460808301856200fa0b565b80820281158282048414176200b51b576200b51b6200fed5565b818103818111156200b51b576200b51b6200fed5565b6001600160a01b0384168152606060208201525f6201021460608301856200fa0b565b82810360408401526200c12381856200fd8b565b5f826201025c577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b604081525f6201027560408301856200fa0b565b82810360208401526200b66181856200fd8b565b6001600160a01b0383168152604060208201525f6200d06760408301846200fa0b565b5f81518060208401855e5f93019283525090919050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f620102f6601a830185620102ac565b7f3a2000000000000000000000000000000000000000000000000000000000000081526200b6616002820185620102ac565b5f6020828403121562010339575f80fd5b81516001600160a01b03811681146200b665575f80fd5b6040516060810167ffffffffffffffff811182821017156201037657620103766200ffad565b60405290565b5f8067ffffffffffffffff8411156201039957620103996200ffad565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff82111715620103cb57620103cb6200ffad565b604052838152905080828401851015620103e3575f80fd5b8383602083015e5f60208583010152509392505050565b5f82601f8301126201040a575f80fd5b6200b665838351602085016201037c565b5f602082840312156201042c575f80fd5b815167ffffffffffffffff81111562010443575f80fd5b6200b51784828501620103fa565b5f6200d067620104628386620102ac565b84620102ac565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f6201049c601a830185620102ac565b7f207573696e6720636f6e7374727563746f7220646174612022000000000000008152620104ce6019820185620102ac565b7f2200000000000000000000000000000000000000000000000000000000000000815260010195945050505050565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201525f6200b66560808301846200fa0b565b5f6020828403121562010556575f80fd5b815167ffffffffffffffff8111156201056d575f80fd5b8201601f810184136201057e575f80fd5b6200b517848251602084016201037c565b5f6201059c8287620102ac565b7f2f000000000000000000000000000000000000000000000000000000000000008152620105ce6001820187620102ac565b90507f2f000000000000000000000000000000000000000000000000000000000000008152620106026001820186620102ac565b90507f2f000000000000000000000000000000000000000000000000000000000000008152620106366001820185620102ac565b7f2e6a736f6e0000000000000000000000000000000000000000000000000000008152600501979650505050505050565b604081525f6201067b60408301846200fa0b565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081525f620106ed601f830184620102ac565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f81527f6d6c00000000000000000000000000000000000000000000000000000000000060208201526022019392505050565b604081525f6201075460408301846200fa0b565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b604081525f620107a760408301846200fa0b565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081525f620108196014830184620102ac565b7f275d2e6b656363616b32353600000000000000000000000000000000000000008152600c019392505050565b604081525f6201085a60408301856200fa0b565b82810360208401526200b66181856200fa0b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f220000000000000000000000000000000000000000000000000000000000000081525f620108ce6001830184620102ac565b7f220000000000000000000000000000000000000000000000000000000000000081526001019392505050565b5f620109088284620102ac565b7f2f6275696c642d696e666f0000000000000000000000000000000000000000008152600b019392505050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201525f6200b665604b830184620102ac565b5f60ff821660ff8103620109cc57620109cc6200fed5565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201525f6200b6656029830184620102ac565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201525f6200b66560808301846200fa0b565b5f6020828403121562010a87575f80fd5b815167ffffffffffffffff81111562010a9e575f80fd5b82016060818503121562010ab0575f80fd5b62010aba62010350565b81518060030b811462010acb575f80fd5b8152602082015167ffffffffffffffff81111562010ae7575f80fd5b62010af586828501620103fa565b602083015250604082015167ffffffffffffffff81111562010b15575f80fd5b62010b2386828501620103fa565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201525f62010b8a6021830184620102ac565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657481527f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960208201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560408201527f616c69666965642070617468206f66207468652062617368206578656375746160608201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960808201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60a08201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960c08201527f6f75722070726f6a65637420287573696e6720666f727761726420736c61736860e08201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101008201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061012082015261013b019392505050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201525f62010d696021830185620102ac565b7f2720696e206f75747075743a200000000000000000000000000000000000000081526200b661600d820185620102ac565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201525f6200b6656029830184620102ac565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201525f6200b6656022830184620102ac565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081525f62010e80600e830184620102ac565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e7472616381527f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e7360208201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e747260408201527f6163742e6a736f6e00000000000000000000000000000000000000000000000060608201526068019392505050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081525f62010f526018830185620102ac565b7f20696e2000000000000000000000000000000000000000000000000000000000815262010f846004820185620102ac565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f727465642081527f6c6963656e736520666f7220626c6f636b206578706c6f72657220766572696660208201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f60408201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c60608201527f206f7220736574207468652060736b69704c6963656e73655479706560206f7060808201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060a082015260b70195945050505050565b5f6201107e8284620102ac565b7f200000000000000000000000000000000000000000000000000000000000000081526001019392505050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081525f6200b665601c830184620102ac565b5f5f198203620110f257620110f26200fed5565b5060010190565b6001815b60018411156201113a578085048111156201111c576201111c6200fed5565b60018416156201112b57908102905b60019390931c928002620110fd565b935093915050565b5f8262011152575060016200b51b565b816201116057505f6200b51b565b81600181146201117957600281146201118457620111a4565b60019150506200b51b565b60ff8411156201119857620111986200fed5565b50506001821b6200b51b565b5060208310610133831016604e8410600b8410161715620111c9575081810a6200b51b565b620111d75f198484620110f9565b805f1904821115620111ed57620111ed6200fed5565b029392505050565b5f6200b665838362011142565b5f816201121357620112136200fed5565b505f190190565b5f620112278285620102ac565b7f3a0000000000000000000000000000000000000000000000000000000000000081526200b6616001820185620102ac565b8181035f8312801583831316838312821617156200eab7576200eab76200fed556fe608060405234801561000f575f80fd5b50604051610c32380380610c3283398101604081905261002e916100f0565b8181600361003c83826101d9565b50600461004982826101d9565b5050505050610293565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610076575f80fd5b81516001600160401b0381111561008f5761008f610053565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100bd576100bd610053565b6040528181528382016020018510156100d4575f80fd5b8160208501602083015e5f918101602001919091529392505050565b5f8060408385031215610101575f80fd5b82516001600160401b03811115610116575f80fd5b61012285828601610067565b602085015190935090506001600160401b0381111561013f575f80fd5b61014b85828601610067565b9150509250929050565b600181811c9082168061016957607f821691505b60208210810361018757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101d457805f5260205f20601f840160051c810160208510156101b25750805b601f840160051c820191505b818110156101d1575f81556001016101be565b50505b505050565b81516001600160401b038111156101f2576101f2610053565b610206816102008454610155565b8461018d565b6020601f821160018114610238575f83156102215750848201515b5f19600385901b1c1916600184901b1784556101d1565b5f84815260208120601f198516915b828110156102675787850151825560209485019460019092019101610247565b508482101561028457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610992806102a05f395ff3fe608060405234801561000f575f80fd5b50600436106100b9575f3560e01c806340c10f191161007257806395d89b411161005857806395d89b411461017c578063a9059cbb14610184578063dd62ed3e14610197575f80fd5b806340c10f191461013257806370a0823114610147575f80fd5b806318160ddd116100a257806318160ddd146100fe57806323b872dd14610110578063313ce56714610123575f80fd5b806306fdde03146100bd578063095ea7b3146100db575b5f80fd5b6100c56101dc565b6040516100d291906107a5565b60405180910390f35b6100ee6100e9366004610820565b61026c565b60405190151581526020016100d2565b6002545b6040519081526020016100d2565b6100ee61011e366004610848565b610285565b604051601281526020016100d2565b610145610140366004610820565b6102a8565b005b610102610155366004610882565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b6100c56102b6565b6100ee610192366004610820565b6102c5565b6101026101a53660046108a2565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b6060600380546101eb906108d3565b80601f0160208091040260200160405190810160405280929190818152602001828054610217906108d3565b80156102625780601f1061023957610100808354040283529160200191610262565b820191905f5260205f20905b81548152906001019060200180831161024557829003601f168201915b5050505050905090565b5f336102798185856102d2565b60019150505b92915050565b5f336102928582856102e4565b61029d8585856103b6565b506001949350505050565b6102b2828261045f565b5050565b6060600480546101eb906108d3565b5f336102798185856103b6565b6102df83838360016104b9565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103b057818110156103a2576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103b084848484035f6104b9565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610405576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b73ffffffffffffffffffffffffffffffffffffffff8216610454576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b6102df8383836105fe565b73ffffffffffffffffffffffffffffffffffffffff82166104ae576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b6102b25f83836105fe565b73ffffffffffffffffffffffffffffffffffffffff8416610508576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b73ffffffffffffffffffffffffffffffffffffffff8316610557576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b73ffffffffffffffffffffffffffffffffffffffff8085165f90815260016020908152604080832093871683529290522082905580156103b0578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105f091815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610635578060025f82825461062a9190610924565b909155506106e59050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260208190526040902054818110156106ba576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610399565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661070e57600280548290039055610739565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161079891815260200190565b60405180910390a3505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461081b575f80fd5b919050565b5f8060408385031215610831575f80fd5b61083a836107f8565b946020939093013593505050565b5f805f6060848603121561085a575f80fd5b610863846107f8565b9250610871602085016107f8565b929592945050506040919091013590565b5f60208284031215610892575f80fd5b61089b826107f8565b9392505050565b5f80604083850312156108b3575f80fd5b6108bc836107f8565b91506108ca602084016107f8565b90509250929050565b600181811c908216806108e757607f821691505b60208210810361091e577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b8082018082111561027f577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea2646970667358221220569a538c722848c143f241fcbfd3f113f81549aa32aa89f9a6169cd25cb5399e64736f6c634300081a0033608060405234801561000f575f80fd5b5060405161124c38038061124c83398101604081905261002e9161010e565b604051806040016040528060048152602001635a65746160e01b815250604051806040016040528060048152602001635a45544160e01b815250816003908161007791906101d7565b50600461008482826101d7565b5050506001600160a01b03821615806100a457506001600160a01b038116155b156100c25760405163e6c4247b60e01b815260040160405180910390fd5b600680546001600160a01b039384166001600160a01b03199182161790915560078054929093169116179055610291565b80516001600160a01b0381168114610109575f80fd5b919050565b5f806040838503121561011f575f80fd5b610128836100f3565b9150610136602084016100f3565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061016757607f821691505b60208210810361018557634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101d257805f5260205f20601f840160051c810160208510156101b05750805b601f840160051c820191505b818110156101cf575f81556001016101bc565b50505b505050565b81516001600160401b038111156101f0576101f061013f565b610204816101fe8454610153565b8461018b565b6020601f821160018114610236575f831561021f5750848201515b5f19600385901b1c1916600184901b1784556101cf565b5f84815260208120601f198516915b828110156102655787850151825560209485019460019092019101610245565b508482101561028257868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610fae8061029e5f395ff3fe608060405234801561000f575f80fd5b5060043610610115575f3560e01c806342966c68116100ad57806379cc67901161007d578063a9059cbb11610063578063a9059cbb14610286578063bff9662a14610299578063dd62ed3e146102b9575f80fd5b806379cc67901461026b57806395d89b411461027e575f80fd5b806342966c68146101fb5780635b1125911461020e57806370a082311461022e578063779e3b6314610263575f80fd5b80631e458bee116100e85780631e458bee1461018157806323b872dd14610194578063313ce567146101a7578063328a01d0146101b6575f80fd5b806306fdde0314610119578063095ea7b31461013757806315d57fd41461015a57806318160ddd1461016f575b5f80fd5b6101216102fe565b60405161012e9190610d7a565b60405180910390f35b61014a610145366004610df5565b61038e565b604051901515815260200161012e565b61016d610168366004610e1d565b6103a7565b005b6002545b60405190815260200161012e565b61016d61018f366004610e4e565b610572565b61014a6101a2366004610e7e565b610625565b6040516012815260200161012e565b6007546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161012e565b61016d610209366004610eb8565b610648565b6006546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b61017361023c366004610ecf565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b61016d610655565b61016d610279366004610df5565b610779565b61012161082a565b61014a610294366004610df5565b610839565b6005546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b6101736102c7366004610e1d565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b60606003805461030d90610eef565b80601f016020809104026020016040519081016040528092919081815260200182805461033990610eef565b80156103845780601f1061035b57610100808354040283529160200191610384565b820191905f5260205f20905b81548152906001019060200180831161036757829003601f168201915b5050505050905090565b5f3361039b818585610846565b60019150505b92915050565b60075473ffffffffffffffffffffffffffffffffffffffff1633148015906103e7575060065473ffffffffffffffffffffffffffffffffffffffff163314155b15610425576040517fcdfcef970000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216158061045c575073ffffffffffffffffffffffffffffffffffffffff8116155b15610493576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316811790935560058054918516919092161790556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff910160405180910390a16040805133815273ffffffffffffffffffffffffffffffffffffffff831660208201527f1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c910160405180910390a15050565b60055473ffffffffffffffffffffffffffffffffffffffff1633146105c5576040517f3fe32fba00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b6105cf8383610858565b808373ffffffffffffffffffffffffffffffffffffffff167fc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb8460405161061891815260200190565b60405180910390a3505050565b5f336106328582856108b6565b61063d858585610983565b506001949350505050565b6106523382610a2c565b50565b60075473ffffffffffffffffffffffffffffffffffffffff1633146106a8576040517fe700765e00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b60065473ffffffffffffffffffffffffffffffffffffffff166106f7576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600654600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0910160405180910390a1565b60055473ffffffffffffffffffffffffffffffffffffffff1633146107cc576040517f3fe32fba00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b6107d68282610a86565b8173ffffffffffffffffffffffffffffffffffffffff167f919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b18260405161081e91815260200190565b60405180910390a25050565b60606004805461030d90610eef565b5f3361039b818585610983565b6108538383836001610a9b565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166108a7576040517fec442f050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b6108b25f8383610be0565b5050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461097d578181101561096f576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602481018290526044810183905260640161041c565b61097d84848484035f610a9b565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166109d2576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8216610a21576040517fec442f050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b610853838383610be0565b73ffffffffffffffffffffffffffffffffffffffff8216610a7b576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b6108b2825f83610be0565b610a918233836108b6565b6108b28282610a2c565b73ffffffffffffffffffffffffffffffffffffffff8416610aea576040517fe602df050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8316610b39576040517f94280d620000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8085165f908152600160209081526040808320938716835292905220829055801561097d578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bd291815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c17578060025f828254610c0c9190610f40565b90915550610cc79050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020819052604090205481811015610c9c576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602481018290526044810183905260640161041c565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610cf057600280548290039055610d1b565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161061891815260200190565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610df0575f80fd5b919050565b5f8060408385031215610e06575f80fd5b610e0f83610dcd565b946020939093013593505050565b5f8060408385031215610e2e575f80fd5b610e3783610dcd565b9150610e4560208401610dcd565b90509250929050565b5f805f60608486031215610e60575f80fd5b610e6984610dcd565b95602085013595506040909401359392505050565b5f805f60608486031215610e90575f80fd5b610e9984610dcd565b9250610ea760208501610dcd565b929592945050506040919091013590565b5f60208284031215610ec8575f80fd5b5035919050565b5f60208284031215610edf575f80fd5b610ee882610dcd565b9392505050565b600181811c90821680610f0357607f821691505b602082108103610f3a577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b808201808211156103a1577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea26469706673582212206a41e2cd6fbf39f12e2917f13d39a938ebbfbe1bbc40009f596c89c4dc977dca64736f6c634300081a0033a26469706673582212209e5e05d46b27207b04b7e144c96445f5c812046f4e183596645b2cecff63335664736f6c634300081a0033",
}

// GatewayEVMInboundTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMInboundTestMetaData.ABI instead.
var GatewayEVMInboundTestABI = GatewayEVMInboundTestMetaData.ABI

// GatewayEVMInboundTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMInboundTestMetaData.Bin instead.
var GatewayEVMInboundTestBin = GatewayEVMInboundTestMetaData.Bin

// DeployGatewayEVMInboundTest deploys a new Ethereum contract, binding an instance of GatewayEVMInboundTest to it.
func DeployGatewayEVMInboundTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVMInboundTest, error) {
	parsed, err := GatewayEVMInboundTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMInboundTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVMInboundTest{GatewayEVMInboundTestCaller: GatewayEVMInboundTestCaller{contract: contract}, GatewayEVMInboundTestTransactor: GatewayEVMInboundTestTransactor{contract: contract}, GatewayEVMInboundTestFilterer: GatewayEVMInboundTestFilterer{contract: contract}}, nil
}

// GatewayEVMInboundTest is an auto generated Go binding around an Ethereum contract.
type GatewayEVMInboundTest struct {
	GatewayEVMInboundTestCaller     // Read-only binding to the contract
	GatewayEVMInboundTestTransactor // Write-only binding to the contract
	GatewayEVMInboundTestFilterer   // Log filterer for contract events
}

// GatewayEVMInboundTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMInboundTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMInboundTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMInboundTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMInboundTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMInboundTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMInboundTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMInboundTestSession struct {
	Contract     *GatewayEVMInboundTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GatewayEVMInboundTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMInboundTestCallerSession struct {
	Contract *GatewayEVMInboundTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// GatewayEVMInboundTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMInboundTestTransactorSession struct {
	Contract     *GatewayEVMInboundTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// GatewayEVMInboundTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMInboundTestRaw struct {
	Contract *GatewayEVMInboundTest // Generic contract binding to access the raw methods on
}

// GatewayEVMInboundTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMInboundTestCallerRaw struct {
	Contract *GatewayEVMInboundTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMInboundTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMInboundTestTransactorRaw struct {
	Contract *GatewayEVMInboundTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVMInboundTest creates a new instance of GatewayEVMInboundTest, bound to a specific deployed contract.
func NewGatewayEVMInboundTest(address common.Address, backend bind.ContractBackend) (*GatewayEVMInboundTest, error) {
	contract, err := bindGatewayEVMInboundTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTest{GatewayEVMInboundTestCaller: GatewayEVMInboundTestCaller{contract: contract}, GatewayEVMInboundTestTransactor: GatewayEVMInboundTestTransactor{contract: contract}, GatewayEVMInboundTestFilterer: GatewayEVMInboundTestFilterer{contract: contract}}, nil
}

// NewGatewayEVMInboundTestCaller creates a new read-only instance of GatewayEVMInboundTest, bound to a specific deployed contract.
func NewGatewayEVMInboundTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMInboundTestCaller, error) {
	contract, err := bindGatewayEVMInboundTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestCaller{contract: contract}, nil
}

// NewGatewayEVMInboundTestTransactor creates a new write-only instance of GatewayEVMInboundTest, bound to a specific deployed contract.
func NewGatewayEVMInboundTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMInboundTestTransactor, error) {
	contract, err := bindGatewayEVMInboundTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestTransactor{contract: contract}, nil
}

// NewGatewayEVMInboundTestFilterer creates a new log filterer instance of GatewayEVMInboundTest, bound to a specific deployed contract.
func NewGatewayEVMInboundTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMInboundTestFilterer, error) {
	contract, err := bindGatewayEVMInboundTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestFilterer{contract: contract}, nil
}

// bindGatewayEVMInboundTest binds a generic wrapper to an already deployed contract.
func bindGatewayEVMInboundTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMInboundTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMInboundTest.Contract.GatewayEVMInboundTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.GatewayEVMInboundTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.GatewayEVMInboundTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMInboundTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.contract.Transact(opts, method, params...)
}

// ADDITIONALACTIONFEEWEI is a free data retrieval call binding the contract method 0xf8d37ef2.
//
// Solidity: function ADDITIONAL_ACTION_FEE_WEI() view returns(uint256)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ADDITIONALACTIONFEEWEI(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "ADDITIONAL_ACTION_FEE_WEI")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ADDITIONALACTIONFEEWEI is a free data retrieval call binding the contract method 0xf8d37ef2.
//
// Solidity: function ADDITIONAL_ACTION_FEE_WEI() view returns(uint256)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ADDITIONALACTIONFEEWEI() (*big.Int, error) {
	return _GatewayEVMInboundTest.Contract.ADDITIONALACTIONFEEWEI(&_GatewayEVMInboundTest.CallOpts)
}

// ADDITIONALACTIONFEEWEI is a free data retrieval call binding the contract method 0xf8d37ef2.
//
// Solidity: function ADDITIONAL_ACTION_FEE_WEI() view returns(uint256)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ADDITIONALACTIONFEEWEI() (*big.Int, error) {
	return _GatewayEVMInboundTest.Contract.ADDITIONALACTIONFEEWEI(&_GatewayEVMInboundTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ISTEST() (bool, error) {
	return _GatewayEVMInboundTest.Contract.ISTEST(&_GatewayEVMInboundTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ISTEST() (bool, error) {
	return _GatewayEVMInboundTest.Contract.ISTEST(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeArtifacts(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeArtifacts(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeContracts(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeContracts(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeSenders(&_GatewayEVMInboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.ExcludeSenders(&_GatewayEVMInboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) Failed() (bool, error) {
	return _GatewayEVMInboundTest.Contract.Failed(&_GatewayEVMInboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) Failed() (bool, error) {
	return _GatewayEVMInboundTest.Contract.Failed(&_GatewayEVMInboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMInboundTest.Contract.TargetArtifactSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMInboundTest.Contract.TargetArtifactSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMInboundTest.Contract.TargetArtifacts(&_GatewayEVMInboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMInboundTest.Contract.TargetArtifacts(&_GatewayEVMInboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.TargetContracts(&_GatewayEVMInboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.TargetContracts(&_GatewayEVMInboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMInboundTest.Contract.TargetInterfaces(&_GatewayEVMInboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMInboundTest.Contract.TargetInterfaces(&_GatewayEVMInboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMInboundTest.Contract.TargetSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMInboundTest.Contract.TargetSelectors(&_GatewayEVMInboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMInboundTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.TargetSenders(&_GatewayEVMInboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMInboundTest.Contract.TargetSenders(&_GatewayEVMInboundTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.SetUp(&_GatewayEVMInboundTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.SetUp(&_GatewayEVMInboundTest.TransactOpts)
}

// TestAdditionalActionDisabledReverts is a paid mutator transaction binding the contract method 0xd86a4a0c.
//
// Solidity: function testAdditionalActionDisabledReverts() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestAdditionalActionDisabledReverts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testAdditionalActionDisabledReverts")
}

// TestAdditionalActionDisabledReverts is a paid mutator transaction binding the contract method 0xd86a4a0c.
//
// Solidity: function testAdditionalActionDisabledReverts() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestAdditionalActionDisabledReverts() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestAdditionalActionDisabledReverts(&_GatewayEVMInboundTest.TransactOpts)
}

// TestAdditionalActionDisabledReverts is a paid mutator transaction binding the contract method 0xd86a4a0c.
//
// Solidity: function testAdditionalActionDisabledReverts() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestAdditionalActionDisabledReverts() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestAdditionalActionDisabledReverts(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x33ed091c.
//
// Solidity: function testCallFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallFailsIfExcessEthProvided(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallFailsIfExcessEthProvided")
}

// TestCallFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x33ed091c.
//
// Solidity: function testCallFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x33ed091c.
//
// Solidity: function testCallFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x598b7edb.
//
// Solidity: function testCallFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallFailsIfInsufficientFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallFailsIfInsufficientFee")
}

// TestCallFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x598b7edb.
//
// Solidity: function testCallFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x598b7edb.
//
// Solidity: function testCallFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallNotAffectedByDepositPause is a paid mutator transaction binding the contract method 0xddd03d57.
//
// Solidity: function testCallNotAffectedByDepositPause() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallNotAffectedByDepositPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallNotAffectedByDepositPause")
}

// TestCallNotAffectedByDepositPause is a paid mutator transaction binding the contract method 0xddd03d57.
//
// Solidity: function testCallNotAffectedByDepositPause() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallNotAffectedByDepositPause() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallNotAffectedByDepositPause(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallNotAffectedByDepositPause is a paid mutator transaction binding the contract method 0xddd03d57.
//
// Solidity: function testCallNotAffectedByDepositPause() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallNotAffectedByDepositPause() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallNotAffectedByDepositPause(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallSecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xc51a4cbe.
//
// Solidity: function testCallSecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallSecondActionFailsIfExcessEthProvided(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallSecondActionFailsIfExcessEthProvided")
}

// TestCallSecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xc51a4cbe.
//
// Solidity: function testCallSecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallSecondActionFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallSecondActionFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallSecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xc51a4cbe.
//
// Solidity: function testCallSecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallSecondActionFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallSecondActionFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x481daadb.
//
// Solidity: function testCallSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallSecondActionRequiresFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallSecondActionRequiresFee")
}

// TestCallSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x481daadb.
//
// Solidity: function testCallSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallSecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallSecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x481daadb.
//
// Solidity: function testCallSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallSecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallSecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayload is a paid mutator transaction binding the contract method 0xaa030c1c.
//
// Solidity: function testCallWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallWithPayload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallWithPayload")
}

// TestCallWithPayload is a paid mutator transaction binding the contract method 0xaa030c1c.
//
// Solidity: function testCallWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayload is a paid mutator transaction binding the contract method 0xaa030c1c.
//
// Solidity: function testCallWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfCallOnRevertIsTrue is a paid mutator transaction binding the contract method 0xba46ba23.
//
// Solidity: function testCallWithPayloadFailsIfCallOnRevertIsTrue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallWithPayloadFailsIfCallOnRevertIsTrue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallWithPayloadFailsIfCallOnRevertIsTrue")
}

// TestCallWithPayloadFailsIfCallOnRevertIsTrue is a paid mutator transaction binding the contract method 0xba46ba23.
//
// Solidity: function testCallWithPayloadFailsIfCallOnRevertIsTrue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallWithPayloadFailsIfCallOnRevertIsTrue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfCallOnRevertIsTrue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfCallOnRevertIsTrue is a paid mutator transaction binding the contract method 0xba46ba23.
//
// Solidity: function testCallWithPayloadFailsIfCallOnRevertIsTrue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallWithPayloadFailsIfCallOnRevertIsTrue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfCallOnRevertIsTrue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x51da903d.
//
// Solidity: function testCallWithPayloadFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallWithPayloadFailsIfDestinationIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallWithPayloadFailsIfDestinationIsZeroAddress")
}

// TestCallWithPayloadFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x51da903d.
//
// Solidity: function testCallWithPayloadFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallWithPayloadFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfDestinationIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfDestinationIsZeroAddress is a paid mutator transaction binding the contract method 0x51da903d.
//
// Solidity: function testCallWithPayloadFailsIfDestinationIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallWithPayloadFailsIfDestinationIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfDestinationIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xb1409f71.
//
// Solidity: function testCallWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallWithPayloadFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallWithPayloadFailsIfPayloadSizeExceeded")
}

// TestCallWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xb1409f71.
//
// Solidity: function testCallWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xb1409f71.
//
// Solidity: function testCallWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20SecondActionRequiresFee is a paid mutator transaction binding the contract method 0x6ab1c516.
//
// Solidity: function testDepositAndCallERC20SecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallERC20SecondActionRequiresFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallERC20SecondActionRequiresFee")
}

// TestDepositAndCallERC20SecondActionRequiresFee is a paid mutator transaction binding the contract method 0x6ab1c516.
//
// Solidity: function testDepositAndCallERC20SecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallERC20SecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20SecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20SecondActionRequiresFee is a paid mutator transaction binding the contract method 0x6ab1c516.
//
// Solidity: function testDepositAndCallERC20SecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallERC20SecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20SecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xf0a635c5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallERC20ToCustodyFailsIfExcessEthProvided")
}

// TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xf0a635c5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xf0a635c5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodyFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x49050a44.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallERC20ToCustodyFailsIfInsufficientFee")
}

// TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x49050a44.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x49050a44.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodyFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x48452329.
//
// Solidity: function testDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided")
}

// TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x48452329.
//
// Solidity: function testDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x48452329.
//
// Solidity: function testDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthToTssFailsForSubsequentActions is a paid mutator transaction binding the contract method 0x7e7dfee3.
//
// Solidity: function testDepositAndCallEthToTssFailsForSubsequentActions() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallEthToTssFailsForSubsequentActions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallEthToTssFailsForSubsequentActions")
}

// TestDepositAndCallEthToTssFailsForSubsequentActions is a paid mutator transaction binding the contract method 0x7e7dfee3.
//
// Solidity: function testDepositAndCallEthToTssFailsForSubsequentActions() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallEthToTssFailsForSubsequentActions() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthToTssFailsForSubsequentActions(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthToTssFailsForSubsequentActions is a paid mutator transaction binding the contract method 0x7e7dfee3.
//
// Solidity: function testDepositAndCallEthToTssFailsForSubsequentActions() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallEthToTssFailsForSubsequentActions() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthToTssFailsForSubsequentActions(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x1b906ef3.
//
// Solidity: function testDepositAndCallEthWithAmountSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallEthWithAmountSecondActionRequiresFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallEthWithAmountSecondActionRequiresFee")
}

// TestDepositAndCallEthWithAmountSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x1b906ef3.
//
// Solidity: function testDepositAndCallEthWithAmountSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallEthWithAmountSecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountSecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x1b906ef3.
//
// Solidity: function testDepositAndCallEthWithAmountSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallEthWithAmountSecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountSecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x8be96f5c.
//
// Solidity: function testDepositAndCallEthWithAmountToTssFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallEthWithAmountToTssFailsIfInsufficientFee")
}

// TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x8be96f5c.
//
// Solidity: function testDepositAndCallEthWithAmountToTssFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0x8be96f5c.
//
// Solidity: function testDepositAndCallEthWithAmountToTssFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountToTssFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue is a paid mutator transaction binding the contract method 0x2cf9572d.
//
// Solidity: function testDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue")
}

// TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue is a paid mutator transaction binding the contract method 0x2cf9572d.
//
// Solidity: function testDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue is a paid mutator transaction binding the contract method 0x2cf9572d.
//
// Solidity: function testDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20SecondActionRequiresFee is a paid mutator transaction binding the contract method 0xdc23a35f.
//
// Solidity: function testDepositERC20SecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20SecondActionRequiresFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20SecondActionRequiresFee")
}

// TestDepositERC20SecondActionRequiresFee is a paid mutator transaction binding the contract method 0xdc23a35f.
//
// Solidity: function testDepositERC20SecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20SecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20SecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20SecondActionRequiresFee is a paid mutator transaction binding the contract method 0xdc23a35f.
//
// Solidity: function testDepositERC20SecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20SecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20SecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustody is a paid mutator transaction binding the contract method 0x6459542a.
//
// Solidity: function testDepositERC20ToCustody() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustody(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustody")
}

// TestDepositERC20ToCustody is a paid mutator transaction binding the contract method 0x6459542a.
//
// Solidity: function testDepositERC20ToCustody() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustody() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustody(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustody is a paid mutator transaction binding the contract method 0x6459542a.
//
// Solidity: function testDepositERC20ToCustody() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustody() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustody(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x15ee8f36.
//
// Solidity: function testDepositERC20ToCustodyFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfExcessEthProvided(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfExcessEthProvided")
}

// TestDepositERC20ToCustodyFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x15ee8f36.
//
// Solidity: function testDepositERC20ToCustodyFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0x15ee8f36.
//
// Solidity: function testDepositERC20ToCustodyFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0xc7a1eccb.
//
// Solidity: function testDepositERC20ToCustodyFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfInsufficientFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfInsufficientFee")
}

// TestDepositERC20ToCustodyFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0xc7a1eccb.
//
// Solidity: function testDepositERC20ToCustodyFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0xc7a1eccb.
//
// Solidity: function testDepositERC20ToCustodyFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x4ce25c0a.
//
// Solidity: function testDepositERC20ToCustodyFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfPayloadSizeExceeded")
}

// TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x4ce25c0a.
//
// Solidity: function testDepositERC20ToCustodyFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x4ce25c0a.
//
// Solidity: function testDepositERC20ToCustodyFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xb2849063.
//
// Solidity: function testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted")
}

// TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xb2849063.
//
// Solidity: function testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xb2849063.
//
// Solidity: function testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xb966e8fa.
//
// Solidity: function testDepositERC20ToCustodySecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodySecondActionFailsIfExcessEthProvided")
}

// TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xb966e8fa.
//
// Solidity: function testDepositERC20ToCustodySecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided is a paid mutator transaction binding the contract method 0xb966e8fa.
//
// Solidity: function testDepositERC20ToCustodySecondActionFailsIfExcessEthProvided() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodySecondActionFailsIfExcessEthProvided(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayload is a paid mutator transaction binding the contract method 0x30f7c04f.
//
// Solidity: function testDepositERC20ToCustodyWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyWithPayload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyWithPayload")
}

// TestDepositERC20ToCustodyWithPayload is a paid mutator transaction binding the contract method 0x30f7c04f.
//
// Solidity: function testDepositERC20ToCustodyWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayload is a paid mutator transaction binding the contract method 0x30f7c04f.
//
// Solidity: function testDepositERC20ToCustodyWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x95cd0445.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded")
}

// TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x95cd0445.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x95cd0445.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xe85c5a07.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted")
}

// TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xe85c5a07.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted is a paid mutator transaction binding the contract method 0xe85c5a07.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTss is a paid mutator transaction binding the contract method 0x0724d8e3.
//
// Solidity: function testDepositEthToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTss(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTss")
}

// TestDepositEthToTss is a paid mutator transaction binding the contract method 0x0724d8e3.
//
// Solidity: function testDepositEthToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTss() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTss(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTss is a paid mutator transaction binding the contract method 0x0724d8e3.
//
// Solidity: function testDepositEthToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTss() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTss(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsForSubsequentActions is a paid mutator transaction binding the contract method 0x91a5c818.
//
// Solidity: function testDepositEthToTssFailsForSubsequentActions() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssFailsForSubsequentActions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssFailsForSubsequentActions")
}

// TestDepositEthToTssFailsForSubsequentActions is a paid mutator transaction binding the contract method 0x91a5c818.
//
// Solidity: function testDepositEthToTssFailsForSubsequentActions() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssFailsForSubsequentActions() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsForSubsequentActions(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsForSubsequentActions is a paid mutator transaction binding the contract method 0x91a5c818.
//
// Solidity: function testDepositEthToTssFailsForSubsequentActions() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssFailsForSubsequentActions() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsForSubsequentActions(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayload is a paid mutator transaction binding the contract method 0x9fd1e597.
//
// Solidity: function testDepositEthToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssWithPayload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssWithPayload")
}

// TestDepositEthToTssWithPayload is a paid mutator transaction binding the contract method 0x9fd1e597.
//
// Solidity: function testDepositEthToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayload is a paid mutator transaction binding the contract method 0x9fd1e597.
//
// Solidity: function testDepositEthToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x466f332e.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded")
}

// TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x466f332e.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x466f332e.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountSecondActionFailsWithOnlyFee is a paid mutator transaction binding the contract method 0x32fc1fad.
//
// Solidity: function testDepositEthWithAmountSecondActionFailsWithOnlyFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountSecondActionFailsWithOnlyFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountSecondActionFailsWithOnlyFee")
}

// TestDepositEthWithAmountSecondActionFailsWithOnlyFee is a paid mutator transaction binding the contract method 0x32fc1fad.
//
// Solidity: function testDepositEthWithAmountSecondActionFailsWithOnlyFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountSecondActionFailsWithOnlyFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountSecondActionFailsWithOnlyFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountSecondActionFailsWithOnlyFee is a paid mutator transaction binding the contract method 0x32fc1fad.
//
// Solidity: function testDepositEthWithAmountSecondActionFailsWithOnlyFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountSecondActionFailsWithOnlyFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountSecondActionFailsWithOnlyFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x3424914f.
//
// Solidity: function testDepositEthWithAmountSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountSecondActionRequiresFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountSecondActionRequiresFee")
}

// TestDepositEthWithAmountSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x3424914f.
//
// Solidity: function testDepositEthWithAmountSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountSecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountSecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountSecondActionRequiresFee is a paid mutator transaction binding the contract method 0x3424914f.
//
// Solidity: function testDepositEthWithAmountSecondActionRequiresFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountSecondActionRequiresFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountSecondActionRequiresFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTss is a paid mutator transaction binding the contract method 0xf2036eda.
//
// Solidity: function testDepositEthWithAmountToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTss(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTss")
}

// TestDepositEthWithAmountToTss is a paid mutator transaction binding the contract method 0xf2036eda.
//
// Solidity: function testDepositEthWithAmountToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTss() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTss(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTss is a paid mutator transaction binding the contract method 0xf2036eda.
//
// Solidity: function testDepositEthWithAmountToTss() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTss() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTss(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xc57926c6.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssFailsIfAmountIs0")
}

// TestDepositEthWithAmountToTssFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xc57926c6.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xc57926c6.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue is a paid mutator transaction binding the contract method 0x6aa02e8b.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue")
}

// TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue is a paid mutator transaction binding the contract method 0x6aa02e8b.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue is a paid mutator transaction binding the contract method 0x6aa02e8b.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue is a paid mutator transaction binding the contract method 0x51ee53cb.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue")
}

// TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue is a paid mutator transaction binding the contract method 0x51ee53cb.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue is a paid mutator transaction binding the contract method 0x51ee53cb.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0xf6e53a5d.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssFailsIfInsufficientFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssFailsIfInsufficientFee")
}

// TestDepositEthWithAmountToTssFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0xf6e53a5d.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfInsufficientFee is a paid mutator transaction binding the contract method 0xf6e53a5d.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfInsufficientFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssFailsIfInsufficientFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfInsufficientFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x545c3745.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssFailsIfPayloadSizeExceeded")
}

// TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x545c3745.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x545c3745.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9acda9fa.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress")
}

// TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9acda9fa.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9acda9fa.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue is a paid mutator transaction binding the contract method 0x09a263c1.
//
// Solidity: function testDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue")
}

// TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue is a paid mutator transaction binding the contract method 0x09a263c1.
//
// Solidity: function testDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue is a paid mutator transaction binding the contract method 0x09a263c1.
//
// Solidity: function testDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayload is a paid mutator transaction binding the contract method 0x6c33bacb.
//
// Solidity: function testDepositEthWithAmountToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssWithPayload(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssWithPayload")
}

// TestDepositEthWithAmountToTssWithPayload is a paid mutator transaction binding the contract method 0x6c33bacb.
//
// Solidity: function testDepositEthWithAmountToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayload is a paid mutator transaction binding the contract method 0x6c33bacb.
//
// Solidity: function testDepositEthWithAmountToTssWithPayload() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssWithPayload() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayload(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x4a780339.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0")
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x4a780339.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x4a780339.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue is a paid mutator transaction binding the contract method 0x9073eafe.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue")
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue is a paid mutator transaction binding the contract method 0x9073eafe.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue is a paid mutator transaction binding the contract method 0x9073eafe.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue is a paid mutator transaction binding the contract method 0x85f5c51c.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue")
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue is a paid mutator transaction binding the contract method 0x85f5c51c.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue is a paid mutator transaction binding the contract method 0x85f5c51c.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xa00a6fff.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded")
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xa00a6fff.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0xa00a6fff.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9a34d8f3.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress")
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9a34d8f3.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9a34d8f3.
//
// Solidity: function testDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositPauseAllowsAllowlistedERC20Deposits is a paid mutator transaction binding the contract method 0x89b90256.
//
// Solidity: function testDepositPauseAllowsAllowlistedERC20Deposits() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositPauseAllowsAllowlistedERC20Deposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositPauseAllowsAllowlistedERC20Deposits")
}

// TestDepositPauseAllowsAllowlistedERC20Deposits is a paid mutator transaction binding the contract method 0x89b90256.
//
// Solidity: function testDepositPauseAllowsAllowlistedERC20Deposits() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositPauseAllowsAllowlistedERC20Deposits() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositPauseAllowsAllowlistedERC20Deposits(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositPauseAllowsAllowlistedERC20Deposits is a paid mutator transaction binding the contract method 0x89b90256.
//
// Solidity: function testDepositPauseAllowsAllowlistedERC20Deposits() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositPauseAllowsAllowlistedERC20Deposits() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositPauseAllowsAllowlistedERC20Deposits(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositPauseAllowsNativeByDefault is a paid mutator transaction binding the contract method 0x43171c12.
//
// Solidity: function testDepositPauseAllowsNativeByDefault() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositPauseAllowsNativeByDefault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositPauseAllowsNativeByDefault")
}

// TestDepositPauseAllowsNativeByDefault is a paid mutator transaction binding the contract method 0x43171c12.
//
// Solidity: function testDepositPauseAllowsNativeByDefault() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositPauseAllowsNativeByDefault() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositPauseAllowsNativeByDefault(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositPauseAllowsNativeByDefault is a paid mutator transaction binding the contract method 0x43171c12.
//
// Solidity: function testDepositPauseAllowsNativeByDefault() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositPauseAllowsNativeByDefault() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositPauseAllowsNativeByDefault(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositPauseAllowsZetaByDefault is a paid mutator transaction binding the contract method 0x5e8fe81f.
//
// Solidity: function testDepositPauseAllowsZetaByDefault() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositPauseAllowsZetaByDefault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositPauseAllowsZetaByDefault")
}

// TestDepositPauseAllowsZetaByDefault is a paid mutator transaction binding the contract method 0x5e8fe81f.
//
// Solidity: function testDepositPauseAllowsZetaByDefault() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositPauseAllowsZetaByDefault() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositPauseAllowsZetaByDefault(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositPauseAllowsZetaByDefault is a paid mutator transaction binding the contract method 0x5e8fe81f.
//
// Solidity: function testDepositPauseAllowsZetaByDefault() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositPauseAllowsZetaByDefault() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositPauseAllowsZetaByDefault(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositPauseBlocksNativeDeposits is a paid mutator transaction binding the contract method 0x9078b01c.
//
// Solidity: function testDepositPauseBlocksNativeDeposits() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositPauseBlocksNativeDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositPauseBlocksNativeDeposits")
}

// TestDepositPauseBlocksNativeDeposits is a paid mutator transaction binding the contract method 0x9078b01c.
//
// Solidity: function testDepositPauseBlocksNativeDeposits() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositPauseBlocksNativeDeposits() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositPauseBlocksNativeDeposits(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositPauseBlocksNativeDeposits is a paid mutator transaction binding the contract method 0x9078b01c.
//
// Solidity: function testDepositPauseBlocksNativeDeposits() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositPauseBlocksNativeDeposits() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositPauseBlocksNativeDeposits(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositPauseBlocksNonZetaERC20Deposits is a paid mutator transaction binding the contract method 0xef17b2ed.
//
// Solidity: function testDepositPauseBlocksNonZetaERC20Deposits() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositPauseBlocksNonZetaERC20Deposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositPauseBlocksNonZetaERC20Deposits")
}

// TestDepositPauseBlocksNonZetaERC20Deposits is a paid mutator transaction binding the contract method 0xef17b2ed.
//
// Solidity: function testDepositPauseBlocksNonZetaERC20Deposits() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositPauseBlocksNonZetaERC20Deposits() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositPauseBlocksNonZetaERC20Deposits(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositPauseBlocksNonZetaERC20Deposits is a paid mutator transaction binding the contract method 0xef17b2ed.
//
// Solidity: function testDepositPauseBlocksNonZetaERC20Deposits() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositPauseBlocksNonZetaERC20Deposits() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositPauseBlocksNonZetaERC20Deposits(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositZetaToConnector is a paid mutator transaction binding the contract method 0xe306a978.
//
// Solidity: function testDepositZetaToConnector() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositZetaToConnector(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositZetaToConnector")
}

// TestDepositZetaToConnector is a paid mutator transaction binding the contract method 0xe306a978.
//
// Solidity: function testDepositZetaToConnector() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositZetaToConnector() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositZetaToConnector(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositZetaToConnector is a paid mutator transaction binding the contract method 0xe306a978.
//
// Solidity: function testDepositZetaToConnector() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositZetaToConnector() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositZetaToConnector(&_GatewayEVMInboundTest.TransactOpts)
}

// TestFeeSystemWithUpdatedFee is a paid mutator transaction binding the contract method 0xf1c6b4d3.
//
// Solidity: function testFeeSystemWithUpdatedFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestFeeSystemWithUpdatedFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testFeeSystemWithUpdatedFee")
}

// TestFeeSystemWithUpdatedFee is a paid mutator transaction binding the contract method 0xf1c6b4d3.
//
// Solidity: function testFeeSystemWithUpdatedFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestFeeSystemWithUpdatedFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestFeeSystemWithUpdatedFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestFeeSystemWithUpdatedFee is a paid mutator transaction binding the contract method 0xf1c6b4d3.
//
// Solidity: function testFeeSystemWithUpdatedFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestFeeSystemWithUpdatedFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestFeeSystemWithUpdatedFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestMixedActionTypesInSameTransaction is a paid mutator transaction binding the contract method 0x63d7a418.
//
// Solidity: function testMixedActionTypesInSameTransaction() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestMixedActionTypesInSameTransaction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testMixedActionTypesInSameTransaction")
}

// TestMixedActionTypesInSameTransaction is a paid mutator transaction binding the contract method 0x63d7a418.
//
// Solidity: function testMixedActionTypesInSameTransaction() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestMixedActionTypesInSameTransaction() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestMixedActionTypesInSameTransaction(&_GatewayEVMInboundTest.TransactOpts)
}

// TestMixedActionTypesInSameTransaction is a paid mutator transaction binding the contract method 0x63d7a418.
//
// Solidity: function testMixedActionTypesInSameTransaction() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestMixedActionTypesInSameTransaction() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestMixedActionTypesInSameTransaction(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyIfAmountIs0 is a paid mutator transaction binding the contract method 0xcdef966f.
//
// Solidity: function testRevertDepositERC20ToCustodyIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositERC20ToCustodyIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositERC20ToCustodyIfAmountIs0")
}

// TestRevertDepositERC20ToCustodyIfAmountIs0 is a paid mutator transaction binding the contract method 0xcdef966f.
//
// Solidity: function testRevertDepositERC20ToCustodyIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositERC20ToCustodyIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyIfAmountIs0 is a paid mutator transaction binding the contract method 0xcdef966f.
//
// Solidity: function testRevertDepositERC20ToCustodyIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositERC20ToCustodyIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x41b8fec6.
//
// Solidity: function testRevertDepositERC20ToCustodyIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositERC20ToCustodyIfReceiverIsZeroAddress")
}

// TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x41b8fec6.
//
// Solidity: function testRevertDepositERC20ToCustodyIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x41b8fec6.
//
// Solidity: function testRevertDepositERC20ToCustodyIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0 is a paid mutator transaction binding the contract method 0x18e14138.
//
// Solidity: function testRevertDepositERC20ToCustodyWithPayloadIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositERC20ToCustodyWithPayloadIfAmountIs0")
}

// TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0 is a paid mutator transaction binding the contract method 0x18e14138.
//
// Solidity: function testRevertDepositERC20ToCustodyWithPayloadIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0 is a paid mutator transaction binding the contract method 0x18e14138.
//
// Solidity: function testRevertDepositERC20ToCustodyWithPayloadIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyWithPayloadIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x09b21ccb.
//
// Solidity: function testRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress")
}

// TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x09b21ccb.
//
// Solidity: function testRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x09b21ccb.
//
// Solidity: function testRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssIfAmountIs0 is a paid mutator transaction binding the contract method 0x8f5a4224.
//
// Solidity: function testRevertDepositEthToTssIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositEthToTssIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositEthToTssIfAmountIs0")
}

// TestRevertDepositEthToTssIfAmountIs0 is a paid mutator transaction binding the contract method 0x8f5a4224.
//
// Solidity: function testRevertDepositEthToTssIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositEthToTssIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssIfAmountIs0 is a paid mutator transaction binding the contract method 0x8f5a4224.
//
// Solidity: function testRevertDepositEthToTssIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositEthToTssIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x0fc37f5e.
//
// Solidity: function testRevertDepositEthToTssIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositEthToTssIfPayloadSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositEthToTssIfPayloadSizeExceeded")
}

// TestRevertDepositEthToTssIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x0fc37f5e.
//
// Solidity: function testRevertDepositEthToTssIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositEthToTssIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssIfPayloadSizeExceeded is a paid mutator transaction binding the contract method 0x0fc37f5e.
//
// Solidity: function testRevertDepositEthToTssIfPayloadSizeExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositEthToTssIfPayloadSizeExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssIfPayloadSizeExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x5c4013d0.
//
// Solidity: function testRevertDepositEthToTssIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositEthToTssIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositEthToTssIfReceiverIsZeroAddress")
}

// TestRevertDepositEthToTssIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x5c4013d0.
//
// Solidity: function testRevertDepositEthToTssIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositEthToTssIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x5c4013d0.
//
// Solidity: function testRevertDepositEthToTssIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositEthToTssIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssWithPayloadIfAmountIs0 is a paid mutator transaction binding the contract method 0x2a5dcf36.
//
// Solidity: function testRevertDepositEthToTssWithPayloadIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositEthToTssWithPayloadIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositEthToTssWithPayloadIfAmountIs0")
}

// TestRevertDepositEthToTssWithPayloadIfAmountIs0 is a paid mutator transaction binding the contract method 0x2a5dcf36.
//
// Solidity: function testRevertDepositEthToTssWithPayloadIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositEthToTssWithPayloadIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssWithPayloadIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssWithPayloadIfAmountIs0 is a paid mutator transaction binding the contract method 0x2a5dcf36.
//
// Solidity: function testRevertDepositEthToTssWithPayloadIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositEthToTssWithPayloadIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssWithPayloadIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9d96310a.
//
// Solidity: function testRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress")
}

// TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9d96310a.
//
// Solidity: function testRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x9d96310a.
//
// Solidity: function testRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestSetDepositAllowedAsset is a paid mutator transaction binding the contract method 0x58592e3f.
//
// Solidity: function testSetDepositAllowedAsset() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestSetDepositAllowedAsset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testSetDepositAllowedAsset")
}

// TestSetDepositAllowedAsset is a paid mutator transaction binding the contract method 0x58592e3f.
//
// Solidity: function testSetDepositAllowedAsset() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestSetDepositAllowedAsset() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestSetDepositAllowedAsset(&_GatewayEVMInboundTest.TransactOpts)
}

// TestSetDepositAllowedAsset is a paid mutator transaction binding the contract method 0x58592e3f.
//
// Solidity: function testSetDepositAllowedAsset() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestSetDepositAllowedAsset() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestSetDepositAllowedAsset(&_GatewayEVMInboundTest.TransactOpts)
}

// TestSetDepositPaused is a paid mutator transaction binding the contract method 0x5d8b85eb.
//
// Solidity: function testSetDepositPaused() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestSetDepositPaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testSetDepositPaused")
}

// TestSetDepositPaused is a paid mutator transaction binding the contract method 0x5d8b85eb.
//
// Solidity: function testSetDepositPaused() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestSetDepositPaused() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestSetDepositPaused(&_GatewayEVMInboundTest.TransactOpts)
}

// TestSetDepositPaused is a paid mutator transaction binding the contract method 0x5d8b85eb.
//
// Solidity: function testSetDepositPaused() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestSetDepositPaused() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestSetDepositPaused(&_GatewayEVMInboundTest.TransactOpts)
}

// TestSetDepositPausedPreservesExplicitZetaBlockAfterRePause is a paid mutator transaction binding the contract method 0x20a5ccfe.
//
// Solidity: function testSetDepositPausedPreservesExplicitZetaBlockAfterRePause() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestSetDepositPausedPreservesExplicitZetaBlockAfterRePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testSetDepositPausedPreservesExplicitZetaBlockAfterRePause")
}

// TestSetDepositPausedPreservesExplicitZetaBlockAfterRePause is a paid mutator transaction binding the contract method 0x20a5ccfe.
//
// Solidity: function testSetDepositPausedPreservesExplicitZetaBlockAfterRePause() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestSetDepositPausedPreservesExplicitZetaBlockAfterRePause() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestSetDepositPausedPreservesExplicitZetaBlockAfterRePause(&_GatewayEVMInboundTest.TransactOpts)
}

// TestSetDepositPausedPreservesExplicitZetaBlockAfterRePause is a paid mutator transaction binding the contract method 0x20a5ccfe.
//
// Solidity: function testSetDepositPausedPreservesExplicitZetaBlockAfterRePause() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestSetDepositPausedPreservesExplicitZetaBlockAfterRePause() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestSetDepositPausedPreservesExplicitZetaBlockAfterRePause(&_GatewayEVMInboundTest.TransactOpts)
}

// TestUpdateAdditionalActionFee is a paid mutator transaction binding the contract method 0x586da267.
//
// Solidity: function testUpdateAdditionalActionFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestUpdateAdditionalActionFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testUpdateAdditionalActionFee")
}

// TestUpdateAdditionalActionFee is a paid mutator transaction binding the contract method 0x586da267.
//
// Solidity: function testUpdateAdditionalActionFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestUpdateAdditionalActionFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestUpdateAdditionalActionFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestUpdateAdditionalActionFee is a paid mutator transaction binding the contract method 0x586da267.
//
// Solidity: function testUpdateAdditionalActionFee() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestUpdateAdditionalActionFee() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestUpdateAdditionalActionFee(&_GatewayEVMInboundTest.TransactOpts)
}

// TestUpdateAdditionalActionFeeOnlyAdmin is a paid mutator transaction binding the contract method 0xa0d60b3a.
//
// Solidity: function testUpdateAdditionalActionFeeOnlyAdmin() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestUpdateAdditionalActionFeeOnlyAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testUpdateAdditionalActionFeeOnlyAdmin")
}

// TestUpdateAdditionalActionFeeOnlyAdmin is a paid mutator transaction binding the contract method 0xa0d60b3a.
//
// Solidity: function testUpdateAdditionalActionFeeOnlyAdmin() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestUpdateAdditionalActionFeeOnlyAdmin() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestUpdateAdditionalActionFeeOnlyAdmin(&_GatewayEVMInboundTest.TransactOpts)
}

// TestUpdateAdditionalActionFeeOnlyAdmin is a paid mutator transaction binding the contract method 0xa0d60b3a.
//
// Solidity: function testUpdateAdditionalActionFeeOnlyAdmin() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestUpdateAdditionalActionFeeOnlyAdmin() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestUpdateAdditionalActionFeeOnlyAdmin(&_GatewayEVMInboundTest.TransactOpts)
}

// GatewayEVMInboundTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestCalledIterator struct {
	Event *GatewayEVMInboundTestCalled // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestCalled)
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
		it.Event = new(GatewayEVMInboundTestCalled)
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
func (it *GatewayEVMInboundTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestCalled represents a Called event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMInboundTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestCalledIterator{contract: _GatewayEVMInboundTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestCalled)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Called", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseCalled(log types.Log) (*GatewayEVMInboundTestCalled, error) {
	event := new(GatewayEVMInboundTestCalled)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestDepositedIterator struct {
	Event *GatewayEVMInboundTestDeposited // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestDeposited)
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
		it.Event = new(GatewayEVMInboundTestDeposited)
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
func (it *GatewayEVMInboundTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestDeposited represents a Deposited event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestDeposited struct {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterDeposited(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMInboundTestDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestDepositedIterator{contract: _GatewayEVMInboundTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestDeposited, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestDeposited)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseDeposited(log types.Log) (*GatewayEVMInboundTestDeposited, error) {
	event := new(GatewayEVMInboundTestDeposited)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestDepositedAndCalledIterator is returned from FilterDepositedAndCalled and is used to iterate over the raw logs and unpacked data for DepositedAndCalled events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestDepositedAndCalledIterator struct {
	Event *GatewayEVMInboundTestDepositedAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestDepositedAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestDepositedAndCalled)
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
		it.Event = new(GatewayEVMInboundTestDepositedAndCalled)
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
func (it *GatewayEVMInboundTestDepositedAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestDepositedAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestDepositedAndCalled represents a DepositedAndCalled event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestDepositedAndCalled struct {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterDepositedAndCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMInboundTestDepositedAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestDepositedAndCalledIterator{contract: _GatewayEVMInboundTest.contract, event: "DepositedAndCalled", logs: logs, sub: sub}, nil
}

// WatchDepositedAndCalled is a free log subscription operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchDepositedAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestDepositedAndCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestDepositedAndCalled)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseDepositedAndCalled(log types.Log) (*GatewayEVMInboundTestDepositedAndCalled, error) {
	event := new(GatewayEVMInboundTestDepositedAndCalled)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestExecutedIterator struct {
	Event *GatewayEVMInboundTestExecuted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestExecuted)
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
		it.Event = new(GatewayEVMInboundTestExecuted)
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
func (it *GatewayEVMInboundTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestExecuted represents a Executed event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMInboundTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestExecutedIterator{contract: _GatewayEVMInboundTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestExecuted)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Executed", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseExecuted(log types.Log) (*GatewayEVMInboundTestExecuted, error) {
	event := new(GatewayEVMInboundTestExecuted)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestExecutedWithERC20Iterator struct {
	Event *GatewayEVMInboundTestExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestExecutedWithERC20)
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
		it.Event = new(GatewayEVMInboundTestExecutedWithERC20)
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
func (it *GatewayEVMInboundTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMInboundTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestExecutedWithERC20Iterator{contract: _GatewayEVMInboundTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestExecutedWithERC20)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMInboundTestExecutedWithERC20, error) {
	event := new(GatewayEVMInboundTestExecutedWithERC20)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedERC20Iterator struct {
	Event *GatewayEVMInboundTestReceivedERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedERC20)
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
		it.Event = new(GatewayEVMInboundTestReceivedERC20)
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
func (it *GatewayEVMInboundTestReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedERC20 represents a ReceivedERC20 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedERC20Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedERC20Iterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedERC20)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedERC20(log types.Log) (*GatewayEVMInboundTestReceivedERC20, error) {
	event := new(GatewayEVMInboundTestReceivedERC20)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedNoParamsIterator struct {
	Event *GatewayEVMInboundTestReceivedNoParams // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedNoParams)
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
		it.Event = new(GatewayEVMInboundTestReceivedNoParams)
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
func (it *GatewayEVMInboundTestReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedNoParams represents a ReceivedNoParams event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedNoParamsIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedNoParamsIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedNoParams)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedNoParams(log types.Log) (*GatewayEVMInboundTestReceivedNoParams, error) {
	event := new(GatewayEVMInboundTestReceivedNoParams)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedNonPayableIterator struct {
	Event *GatewayEVMInboundTestReceivedNonPayable // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedNonPayable)
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
		it.Event = new(GatewayEVMInboundTestReceivedNonPayable)
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
func (it *GatewayEVMInboundTestReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedNonPayable represents a ReceivedNonPayable event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedNonPayableIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedNonPayableIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedNonPayable)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedNonPayable(log types.Log) (*GatewayEVMInboundTestReceivedNonPayable, error) {
	event := new(GatewayEVMInboundTestReceivedNonPayable)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedOnCallIterator is returned from FilterReceivedOnCall and is used to iterate over the raw logs and unpacked data for ReceivedOnCall events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedOnCallIterator struct {
	Event *GatewayEVMInboundTestReceivedOnCall // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestReceivedOnCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedOnCall)
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
		it.Event = new(GatewayEVMInboundTestReceivedOnCall)
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
func (it *GatewayEVMInboundTestReceivedOnCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedOnCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedOnCall represents a ReceivedOnCall event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedOnCall struct {
	Sender  common.Address
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceivedOnCall is a free log retrieval operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedOnCall(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedOnCallIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedOnCallIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedOnCall", logs: logs, sub: sub}, nil
}

// WatchReceivedOnCall is a free log subscription operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedOnCall(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedOnCall) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedOnCall)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedOnCall(log types.Log) (*GatewayEVMInboundTestReceivedOnCall, error) {
	event := new(GatewayEVMInboundTestReceivedOnCall)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedPayableIterator struct {
	Event *GatewayEVMInboundTestReceivedPayable // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedPayable)
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
		it.Event = new(GatewayEVMInboundTestReceivedPayable)
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
func (it *GatewayEVMInboundTestReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedPayable represents a ReceivedPayable event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedPayable struct {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedPayableIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedPayableIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedPayable)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedPayable(log types.Log) (*GatewayEVMInboundTestReceivedPayable, error) {
	event := new(GatewayEVMInboundTestReceivedPayable)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedRevertIterator struct {
	Event *GatewayEVMInboundTestReceivedRevert // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReceivedRevert)
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
		it.Event = new(GatewayEVMInboundTestReceivedRevert)
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
func (it *GatewayEVMInboundTestReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReceivedRevert represents a ReceivedRevert event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*GatewayEVMInboundTestReceivedRevertIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestReceivedRevertIterator{contract: _GatewayEVMInboundTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReceivedRevert)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReceivedRevert(log types.Log) (*GatewayEVMInboundTestReceivedRevert, error) {
	event := new(GatewayEVMInboundTestReceivedRevert)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestRevertedIterator struct {
	Event *GatewayEVMInboundTestReverted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestReverted)
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
		it.Event = new(GatewayEVMInboundTestReverted)
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
func (it *GatewayEVMInboundTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestReverted represents a Reverted event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestReverted struct {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMInboundTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestRevertedIterator{contract: _GatewayEVMInboundTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestReverted)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Reverted", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseReverted(log types.Log) (*GatewayEVMInboundTestReverted, error) {
	event := new(GatewayEVMInboundTestReverted)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestUpdatedAdditionalActionFeeIterator is returned from FilterUpdatedAdditionalActionFee and is used to iterate over the raw logs and unpacked data for UpdatedAdditionalActionFee events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestUpdatedAdditionalActionFeeIterator struct {
	Event *GatewayEVMInboundTestUpdatedAdditionalActionFee // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestUpdatedAdditionalActionFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestUpdatedAdditionalActionFee)
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
		it.Event = new(GatewayEVMInboundTestUpdatedAdditionalActionFee)
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
func (it *GatewayEVMInboundTestUpdatedAdditionalActionFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestUpdatedAdditionalActionFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestUpdatedAdditionalActionFee represents a UpdatedAdditionalActionFee event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestUpdatedAdditionalActionFee struct {
	OldFeeWei *big.Int
	NewFeeWei *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAdditionalActionFee is a free log retrieval operation binding the contract event 0xe60a2882aff7841a7d0492b804d46ff2d9b07ea353dbc3791e8276b729e40cd8.
//
// Solidity: event UpdatedAdditionalActionFee(uint256 oldFeeWei, uint256 newFeeWei)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterUpdatedAdditionalActionFee(opts *bind.FilterOpts) (*GatewayEVMInboundTestUpdatedAdditionalActionFeeIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "UpdatedAdditionalActionFee")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestUpdatedAdditionalActionFeeIterator{contract: _GatewayEVMInboundTest.contract, event: "UpdatedAdditionalActionFee", logs: logs, sub: sub}, nil
}

// WatchUpdatedAdditionalActionFee is a free log subscription operation binding the contract event 0xe60a2882aff7841a7d0492b804d46ff2d9b07ea353dbc3791e8276b729e40cd8.
//
// Solidity: event UpdatedAdditionalActionFee(uint256 oldFeeWei, uint256 newFeeWei)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchUpdatedAdditionalActionFee(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestUpdatedAdditionalActionFee) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "UpdatedAdditionalActionFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestUpdatedAdditionalActionFee)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "UpdatedAdditionalActionFee", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseUpdatedAdditionalActionFee(log types.Log) (*GatewayEVMInboundTestUpdatedAdditionalActionFee, error) {
	event := new(GatewayEVMInboundTestUpdatedAdditionalActionFee)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "UpdatedAdditionalActionFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestUpdatedDepositAllowedAssetIterator is returned from FilterUpdatedDepositAllowedAsset and is used to iterate over the raw logs and unpacked data for UpdatedDepositAllowedAsset events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestUpdatedDepositAllowedAssetIterator struct {
	Event *GatewayEVMInboundTestUpdatedDepositAllowedAsset // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestUpdatedDepositAllowedAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestUpdatedDepositAllowedAsset)
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
		it.Event = new(GatewayEVMInboundTestUpdatedDepositAllowedAsset)
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
func (it *GatewayEVMInboundTestUpdatedDepositAllowedAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestUpdatedDepositAllowedAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestUpdatedDepositAllowedAsset represents a UpdatedDepositAllowedAsset event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestUpdatedDepositAllowedAsset struct {
	Asset   common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedDepositAllowedAsset is a free log retrieval operation binding the contract event 0x50708318fef6b3b62fbd0894ea2c5d8fc3ccce4785ed240e4c525ae40bd23332.
//
// Solidity: event UpdatedDepositAllowedAsset(address indexed asset, bool allowed)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterUpdatedDepositAllowedAsset(opts *bind.FilterOpts, asset []common.Address) (*GatewayEVMInboundTestUpdatedDepositAllowedAssetIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "UpdatedDepositAllowedAsset", assetRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestUpdatedDepositAllowedAssetIterator{contract: _GatewayEVMInboundTest.contract, event: "UpdatedDepositAllowedAsset", logs: logs, sub: sub}, nil
}

// WatchUpdatedDepositAllowedAsset is a free log subscription operation binding the contract event 0x50708318fef6b3b62fbd0894ea2c5d8fc3ccce4785ed240e4c525ae40bd23332.
//
// Solidity: event UpdatedDepositAllowedAsset(address indexed asset, bool allowed)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchUpdatedDepositAllowedAsset(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestUpdatedDepositAllowedAsset, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "UpdatedDepositAllowedAsset", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestUpdatedDepositAllowedAsset)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "UpdatedDepositAllowedAsset", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseUpdatedDepositAllowedAsset(log types.Log) (*GatewayEVMInboundTestUpdatedDepositAllowedAsset, error) {
	event := new(GatewayEVMInboundTestUpdatedDepositAllowedAsset)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "UpdatedDepositAllowedAsset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestUpdatedDepositPausedIterator is returned from FilterUpdatedDepositPaused and is used to iterate over the raw logs and unpacked data for UpdatedDepositPaused events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestUpdatedDepositPausedIterator struct {
	Event *GatewayEVMInboundTestUpdatedDepositPaused // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestUpdatedDepositPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestUpdatedDepositPaused)
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
		it.Event = new(GatewayEVMInboundTestUpdatedDepositPaused)
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
func (it *GatewayEVMInboundTestUpdatedDepositPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestUpdatedDepositPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestUpdatedDepositPaused represents a UpdatedDepositPaused event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestUpdatedDepositPaused struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedDepositPaused is a free log retrieval operation binding the contract event 0xd39ad370883b0cd1a8172b5b006a3ebcaaf65183008c91ffd7655afb74174e57.
//
// Solidity: event UpdatedDepositPaused(bool paused)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterUpdatedDepositPaused(opts *bind.FilterOpts) (*GatewayEVMInboundTestUpdatedDepositPausedIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "UpdatedDepositPaused")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestUpdatedDepositPausedIterator{contract: _GatewayEVMInboundTest.contract, event: "UpdatedDepositPaused", logs: logs, sub: sub}, nil
}

// WatchUpdatedDepositPaused is a free log subscription operation binding the contract event 0xd39ad370883b0cd1a8172b5b006a3ebcaaf65183008c91ffd7655afb74174e57.
//
// Solidity: event UpdatedDepositPaused(bool paused)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchUpdatedDepositPaused(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestUpdatedDepositPaused) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "UpdatedDepositPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestUpdatedDepositPaused)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "UpdatedDepositPaused", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseUpdatedDepositPaused(log types.Log) (*GatewayEVMInboundTestUpdatedDepositPaused, error) {
	event := new(GatewayEVMInboundTestUpdatedDepositPaused)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "UpdatedDepositPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator struct {
	Event *GatewayEVMInboundTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestUpdatedGatewayTSSAddress)
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
		it.Event = new(GatewayEVMInboundTestUpdatedGatewayTSSAddress)
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
func (it *GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestUpdatedGatewayTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestUpdatedGatewayTSSAddressIterator{contract: _GatewayEVMInboundTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestUpdatedGatewayTSSAddress)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*GatewayEVMInboundTestUpdatedGatewayTSSAddress, error) {
	event := new(GatewayEVMInboundTestUpdatedGatewayTSSAddress)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogIterator struct {
	Event *GatewayEVMInboundTestLog // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLog)
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
		it.Event = new(GatewayEVMInboundTestLog)
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
func (it *GatewayEVMInboundTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLog represents a Log event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogIterator{contract: _GatewayEVMInboundTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLog)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLog(log types.Log) (*GatewayEVMInboundTestLog, error) {
	event := new(GatewayEVMInboundTestLog)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogAddressIterator struct {
	Event *GatewayEVMInboundTestLogAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogAddress)
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
		it.Event = new(GatewayEVMInboundTestLogAddress)
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
func (it *GatewayEVMInboundTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogAddress represents a LogAddress event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogAddressIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogAddressIterator{contract: _GatewayEVMInboundTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogAddress)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogAddress(log types.Log) (*GatewayEVMInboundTestLogAddress, error) {
	event := new(GatewayEVMInboundTestLogAddress)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArrayIterator struct {
	Event *GatewayEVMInboundTestLogArray // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogArray)
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
		it.Event = new(GatewayEVMInboundTestLogArray)
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
func (it *GatewayEVMInboundTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogArray represents a LogArray event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogArrayIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogArrayIterator{contract: _GatewayEVMInboundTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogArray)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogArray(log types.Log) (*GatewayEVMInboundTestLogArray, error) {
	event := new(GatewayEVMInboundTestLogArray)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray0Iterator struct {
	Event *GatewayEVMInboundTestLogArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogArray0)
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
		it.Event = new(GatewayEVMInboundTestLogArray0)
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
func (it *GatewayEVMInboundTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogArray0 represents a LogArray0 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogArray0Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogArray0)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogArray0(log types.Log) (*GatewayEVMInboundTestLogArray0, error) {
	event := new(GatewayEVMInboundTestLogArray0)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray1Iterator struct {
	Event *GatewayEVMInboundTestLogArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogArray1)
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
		it.Event = new(GatewayEVMInboundTestLogArray1)
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
func (it *GatewayEVMInboundTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogArray1 represents a LogArray1 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogArray1Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogArray1)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogArray1(log types.Log) (*GatewayEVMInboundTestLogArray1, error) {
	event := new(GatewayEVMInboundTestLogArray1)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogBytesIterator struct {
	Event *GatewayEVMInboundTestLogBytes // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogBytes)
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
		it.Event = new(GatewayEVMInboundTestLogBytes)
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
func (it *GatewayEVMInboundTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogBytes represents a LogBytes event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogBytesIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogBytesIterator{contract: _GatewayEVMInboundTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogBytes)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogBytes(log types.Log) (*GatewayEVMInboundTestLogBytes, error) {
	event := new(GatewayEVMInboundTestLogBytes)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogBytes32Iterator struct {
	Event *GatewayEVMInboundTestLogBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogBytes32)
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
		it.Event = new(GatewayEVMInboundTestLogBytes32)
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
func (it *GatewayEVMInboundTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogBytes32 represents a LogBytes32 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogBytes32Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogBytes32)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogBytes32(log types.Log) (*GatewayEVMInboundTestLogBytes32, error) {
	event := new(GatewayEVMInboundTestLogBytes32)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogIntIterator struct {
	Event *GatewayEVMInboundTestLogInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogInt)
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
		it.Event = new(GatewayEVMInboundTestLogInt)
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
func (it *GatewayEVMInboundTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogInt represents a LogInt event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogIntIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogIntIterator{contract: _GatewayEVMInboundTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogInt)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogInt(log types.Log) (*GatewayEVMInboundTestLogInt, error) {
	event := new(GatewayEVMInboundTestLogInt)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedAddressIterator struct {
	Event *GatewayEVMInboundTestLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedAddress)
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
		it.Event = new(GatewayEVMInboundTestLogNamedAddress)
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
func (it *GatewayEVMInboundTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedAddressIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedAddress)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayEVMInboundTestLogNamedAddress, error) {
	event := new(GatewayEVMInboundTestLogNamedAddress)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArrayIterator struct {
	Event *GatewayEVMInboundTestLogNamedArray // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedArray)
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
		it.Event = new(GatewayEVMInboundTestLogNamedArray)
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
func (it *GatewayEVMInboundTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedArray represents a LogNamedArray event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedArrayIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedArray)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayEVMInboundTestLogNamedArray, error) {
	event := new(GatewayEVMInboundTestLogNamedArray)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray0Iterator struct {
	Event *GatewayEVMInboundTestLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedArray0)
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
		it.Event = new(GatewayEVMInboundTestLogNamedArray0)
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
func (it *GatewayEVMInboundTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedArray0Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedArray0)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayEVMInboundTestLogNamedArray0, error) {
	event := new(GatewayEVMInboundTestLogNamedArray0)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray1Iterator struct {
	Event *GatewayEVMInboundTestLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedArray1)
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
		it.Event = new(GatewayEVMInboundTestLogNamedArray1)
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
func (it *GatewayEVMInboundTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedArray1Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedArray1)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayEVMInboundTestLogNamedArray1, error) {
	event := new(GatewayEVMInboundTestLogNamedArray1)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedBytesIterator struct {
	Event *GatewayEVMInboundTestLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedBytes)
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
		it.Event = new(GatewayEVMInboundTestLogNamedBytes)
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
func (it *GatewayEVMInboundTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedBytesIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedBytes)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayEVMInboundTestLogNamedBytes, error) {
	event := new(GatewayEVMInboundTestLogNamedBytes)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedBytes32Iterator struct {
	Event *GatewayEVMInboundTestLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedBytes32)
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
		it.Event = new(GatewayEVMInboundTestLogNamedBytes32)
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
func (it *GatewayEVMInboundTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedBytes32Iterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedBytes32)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayEVMInboundTestLogNamedBytes32, error) {
	event := new(GatewayEVMInboundTestLogNamedBytes32)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedDecimalIntIterator struct {
	Event *GatewayEVMInboundTestLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedDecimalInt)
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
		it.Event = new(GatewayEVMInboundTestLogNamedDecimalInt)
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
func (it *GatewayEVMInboundTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedDecimalIntIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedDecimalInt)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayEVMInboundTestLogNamedDecimalInt, error) {
	event := new(GatewayEVMInboundTestLogNamedDecimalInt)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedDecimalUintIterator struct {
	Event *GatewayEVMInboundTestLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedDecimalUint)
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
		it.Event = new(GatewayEVMInboundTestLogNamedDecimalUint)
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
func (it *GatewayEVMInboundTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedDecimalUintIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedDecimalUint)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayEVMInboundTestLogNamedDecimalUint, error) {
	event := new(GatewayEVMInboundTestLogNamedDecimalUint)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedIntIterator struct {
	Event *GatewayEVMInboundTestLogNamedInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedInt)
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
		it.Event = new(GatewayEVMInboundTestLogNamedInt)
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
func (it *GatewayEVMInboundTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedInt represents a LogNamedInt event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedIntIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedInt)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayEVMInboundTestLogNamedInt, error) {
	event := new(GatewayEVMInboundTestLogNamedInt)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedStringIterator struct {
	Event *GatewayEVMInboundTestLogNamedString // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedString)
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
		it.Event = new(GatewayEVMInboundTestLogNamedString)
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
func (it *GatewayEVMInboundTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedString represents a LogNamedString event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedStringIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedString)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedString(log types.Log) (*GatewayEVMInboundTestLogNamedString, error) {
	event := new(GatewayEVMInboundTestLogNamedString)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedUintIterator struct {
	Event *GatewayEVMInboundTestLogNamedUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogNamedUint)
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
		it.Event = new(GatewayEVMInboundTestLogNamedUint)
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
func (it *GatewayEVMInboundTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogNamedUint represents a LogNamedUint event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogNamedUintIterator{contract: _GatewayEVMInboundTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogNamedUint)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayEVMInboundTestLogNamedUint, error) {
	event := new(GatewayEVMInboundTestLogNamedUint)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogStringIterator struct {
	Event *GatewayEVMInboundTestLogString // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogString)
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
		it.Event = new(GatewayEVMInboundTestLogString)
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
func (it *GatewayEVMInboundTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogString represents a LogString event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogStringIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogStringIterator{contract: _GatewayEVMInboundTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogString)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogString(log types.Log) (*GatewayEVMInboundTestLogString, error) {
	event := new(GatewayEVMInboundTestLogString)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogUintIterator struct {
	Event *GatewayEVMInboundTestLogUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogUint)
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
		it.Event = new(GatewayEVMInboundTestLogUint)
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
func (it *GatewayEVMInboundTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogUint represents a LogUint event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogUintIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogUintIterator{contract: _GatewayEVMInboundTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogUint)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogUint(log types.Log) (*GatewayEVMInboundTestLogUint, error) {
	event := new(GatewayEVMInboundTestLogUint)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMInboundTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogsIterator struct {
	Event *GatewayEVMInboundTestLogs // Event containing the contract specifics and raw log

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
func (it *GatewayEVMInboundTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMInboundTestLogs)
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
		it.Event = new(GatewayEVMInboundTestLogs)
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
func (it *GatewayEVMInboundTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMInboundTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMInboundTestLogs represents a Logs event raised by the GatewayEVMInboundTest contract.
type GatewayEVMInboundTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayEVMInboundTestLogsIterator, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMInboundTestLogsIterator{contract: _GatewayEVMInboundTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayEVMInboundTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMInboundTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMInboundTestLogs)
				if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "logs", log); err != nil {
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
func (_GatewayEVMInboundTest *GatewayEVMInboundTestFilterer) ParseLogs(log types.Log) (*GatewayEVMInboundTestLogs, error) {
	event := new(GatewayEVMInboundTestLogs)
	if err := _GatewayEVMInboundTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
