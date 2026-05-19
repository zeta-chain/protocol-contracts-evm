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
	ABI: "[{\"type\":\"function\",\"name\":\"ADDITIONAL_ACTION_FEE_WEI\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testAdditionalActionDisabledReverts\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallSecondActionFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallSecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfCallOnRevertIsTrue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfDestinationIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20SecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20ToCustodyFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20ToCustodyFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallERC20ToCustodySecondActionFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthToTssFailsForSubsequentActions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthWithAmountSecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthWithAmountToTssFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositAndCallEthWithAmountToTssSecondActionFailsIfIncorrectValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20SecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodySecondActionFailsIfExcessEthProvided\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayloadFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTss\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssFailsForSubsequentActions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssWithPayloadFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountSecondActionFailsWithOnlyFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountSecondActionRequiresFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTss\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfAmountIsLessThanMsgValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfAmountIsMoreThanMsgValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfInsufficientFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssSecondActionFailsIfIncorrectValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsLessThanMsgValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfAmountIsMoreThanMsgValue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthWithAmountToTssWithPayloadFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositPauseAllowsAllowlistedERC20Deposits\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositPauseAllowsNativeByDefault\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositPauseAllowsZetaByDefault\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositPauseBlocksNativeDeposits\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositPauseBlocksNonZetaERC20Deposits\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZetaToConnector\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testFeeSystemWithUpdatedFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testMixedActionTypesInSameTransaction\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetDepositAllowedAsset\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetDepositPaused\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetDepositPausedPreservesExplicitZetaBlockAfterRePause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpdateAdditionalActionFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpdateAdditionalActionFeeOnlyAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositedAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedAdditionalActionFee\",\"inputs\":[{\"name\":\"oldFeeWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newFeeWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedDepositAllowedAsset\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedDepositPaused\",\"inputs\":[{\"name\":\"paused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AdditionalActionDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AssetDepositNotAllowed\",\"inputs\":[{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallOnRevertNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExcessETHProvided\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectValueProvided\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientEVMAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFee\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f8054909116909117905569021e19e0c9bab2400000602c553480156039575f80fd5b50620141da80620000495f395ff3fe608060405234801562000010575f80fd5b506004361062000648575f3560e01c806384c59bcf116200033b578063b152693411620001bf578063dc23a35f1162000107578063f1c6b4d311620000ab578063f75fc9691162000083578063f75fc9691462000a19578063f8d37ef21462000a23578063fa7626d41462000a3d575f80fd5b8063f1c6b4d314620009fb578063f2036eda1462000a05578063f6e53a5d1462000a0f575f80fd5b8063e85c5a0711620000df578063e85c5a0714620009dd578063ef17b2ed14620009e7578063f0a635c514620009f1575f80fd5b8063dc23a35f14620009bf578063e20c9f7114620009c9578063e306a97814620009d3575f80fd5b8063ba46ba23116200016f578063c57926c61162000147578063c57926c614620009a1578063c7a1eccb14620009ab578063d86a4a0c14620009b5575f80fd5b8063ba46ba231462000983578063c0fab86d146200098d578063c51a4cbe1462000997575f80fd5b8063b5508aa911620001a3578063b5508aa91462000954578063b966e8fa146200095e578063ba414fa61462000968575f80fd5b8063b15269341462000940578063b2849063146200094a575f80fd5b8063916a17c611620002835780639fd1e5971162000233578063aa030c1c116200020b578063aa030c1c1462000922578063b0464fdc146200092c578063b1409f711462000936575f80fd5b80639fd1e5971462000904578063a00a6fff146200090e578063a0d60b3a1462000918575f80fd5b806395cd0445116200026757806395cd044514620008e65780639a34d8f314620008f05780639acda9fa14620008fa575f80fd5b8063916a17c614620008c357806391a5c81814620008dc575f80fd5b8063895c2bc611620002eb5780638be96f5c11620002c35780638be96f5c14620008a55780639073eafe14620008af5780639078b01c14620008b9575f80fd5b8063895c2bc6146200088757806389b9025614620008915780638aeeb7e7146200089b575f80fd5b806385f5c51c116200031f57806385f5c51c146200086957806386b6e429146200087357806388343a41146200087d575f80fd5b806384c59bcf146200084657806385226c811462000850575f80fd5b8063481daadb11620004cf578063598b7edb11620004175780636aa02e8b11620003bb5780637bb46d4611620003935780637bb46d4614620008285780637e7dfee31462000832578063846b9f7f146200083c575f80fd5b80636aa02e8b146200080a5780636ab1c51614620008145780636c33bacb146200081e575f80fd5b806363d7a41811620003ef57806363d7a41814620007dd5780636459542a14620007e757806366d9a9a014620007f1575f80fd5b8063598b7edb14620007bf5780635d8b85eb14620007c95780635e8fe81f14620007d3575f80fd5b806351da903d116200047f578063545c37451162000457578063545c374514620007a157806358592e3f14620007ab578063586da26714620007b5575f80fd5b806351da903d146200078357806351ee53cb146200078d57806353c9a0a31462000797575f80fd5b806349050a4411620004b357806349050a4414620007655780634a780339146200076f5780634ce25c0a1462000779575f80fd5b8063481daadb146200075157806348452329146200075b575f80fd5b80632cf9572d11620005935780633e5e3c231162000543578063458085f5116200051b578063458085f51462000733578063466f332e146200073d578063478095e51462000747575f80fd5b80633e5e3c2314620007155780633f7286f4146200071f57806343171c121462000729575f80fd5b806332fc1fad116200057757806332fc1fad14620006f757806333ed091c14620007015780633424914f146200070b575f80fd5b80632cf9572d14620006e357806330f7c04f14620006ed575f80fd5b806315ee8f3611620005fb5780631f4f542f11620005d35780631f4f542f14620006b657806320a5ccfe14620006c05780632ade388014620006ca575f80fd5b806315ee8f3614620006805780631b906ef3146200068a5780631ed7831c1462000694575f80fd5b806309a263c1116200062f57806309a263c114620006625780630a9254e4146200066c5780630fc37f5e1462000676575f80fd5b806301a74aee146200064c5780630724d8e31462000658575b5f80fd5b6200065662000a4b565b005b6200065662000c2a565b6200065662000de0565b6200065662000fd8565b6200065662001b53565b6200065662001df7565b6200065662001f84565b6200069e62002309565b604051620006ad9190620108c9565b60405180910390f35b620006566200236b565b620006566200253b565b620006d462002a3c565b604051620006ad919062010944565b6200065662002b84565b6200065662002d9f565b620006566200317c565b62000656620032de565b6200065662003444565b6200069e62003742565b6200069e620037a2565b6200065662003802565b6200065662003a26565b6200065662003c6d565b6200065662003ff1565b62000656620041d9565b6200065662004513565b62000656620047be565b6200065662004a6a565b6200065662004b8b565b6200065662004e69565b6200065662004f80565b62000656620050a1565b62000656620051bd565b620006566200542f565b6200065662005628565b62000656620057e5565b620006566200591c565b6200065662005dd6565b6200065662005fc3565b6200065662006616565b620007fb62006990565b604051620006ad919062010a95565b6200065662006b00565b6200065662006bff565b62000656620071d4565b62000656620073ca565b62000656620074a1565b6200065662007660565b6200065662007721565b6200085a620077c1565b604051620006ad919062010b37565b6200065662007896565b6200065662007a01565b6200065662007b42565b6200065662007cab565b6200065662007dcd565b62000656620080ea565b62000656620081be565b620006566200831d565b6200065662008464565b620008cd6200899e565b604051620006ad919062010bb0565b6200065662008a83565b6200065662008b83565b6200065662008f6a565b6200065662009088565b620006566200915e565b6200065662009302565b6200065662009652565b620006566200979e565b620008cd6200992d565b6200065662009a12565b6200065662009d57565b6200065662009ebe565b6200085a6200a0d2565b620006566200a1a7565b620009726200a3ed565b6040519015158152602001620006ad565b620006566200a4c1565b620006566200a61c565b620006566200a7cd565b620006566200a9b6565b620006566200aa8f565b620006566200ac09565b620006566200af98565b6200069e6200b4df565b620006566200b53f565b620006566200b7b4565b620006566200baa1565b620006566200bd85565b620006566200bea5565b620006566200c078565b620006566200c1d4565b620006566200c2ed565b62000a2e62030d4081565b604051908152602001620006ad565b601f54620009729060ff1681565b6026546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052815160a0810183526103218082525f8284018190528285019190915283519283019093528282526060810191909152919250906080810162000ad8621e8480600162010c76565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162000b749160040162010c8c565b5f604051808303815f87803b15801562000b8c575f80fd5b505af115801562000b9f573d5f803e3d5ffd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb4935062000bf792909116908690869060040162010cfc565b5f604051808303815f87803b15801562000c0f575f80fd5b505af115801562000c22573d5f803e3d5ffd5b505050505050565b6027546020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039182166084820152620186a092919091163190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562000cc4575f80fd5b505af115801562000cd7573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c9062000d279086905f9060289062010e71565b60405180910390a36020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263726ac97c92869262000d84929091169060289060040162010ea7565b5f604051808303818588803b15801562000d9c575f80fd5b505af115801562000daf573d5f803e3d5ffd5b50506027546001600160a01b031631925062000ddb915062000dd49050848462010c76565b826200c38d565b505050565b6020546026546040517f282906ed000000000000000000000000000000000000000000000000000000008152620186a0926001600160a01b039081169263282906ed92859262000e3a921690839060289060040162010eca565b5f604051808303818588803b15801562000e52575f80fd5b505af115801562000e65573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392507f84222ac700000000000000000000000000000000000000000000000000000000915062000eb8905062030d408562010c76565b62000ec762030d408662010c76565b62000ed490600162010c76565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b909216825262000f2e9160040162010c8c565b5f604051808303815f87803b15801562000f46575f80fd5b505af115801562000f59573d5f803e3d5ffd5b50506020546001600160a01b0316915063282906ed905062000f7f62030d408462010c76565b62000f8c90600162010c76565b6026546040516001600160e01b031960e085901b16815262000fc0916001600160a01b031690869060289060040162010eca565b5f604051808303818588803b15801562000c0f575f80fd5b602580547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602680548216611234179055602780549091166156781790556040516200102c90620107fc565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103905ff080158015620010af573d5f803e3d5ffd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161790556027546040519116908190620010fa906201080a565b6001600160a01b03928316815291166020820152604001604051809103905ff0801580156200112b573d5f803e3d5ffd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081178255604080518082018252600e81527f4761746577617945564d2e736f6c0000000000000000000000000000000000006020820152602754602554925190861694810194909452604484019290925290921660648201526200120a91906084015b60408051601f198184030181529190526020810180516001600160e01b03167fc0c53b8b000000000000000000000000000000000000000000000000000000001790526200c40a565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff000000000000000000000000000000000000000090911681178255604080518082018252601081527f4552433230437573746f64792e736f6c000000000000000000000000000000009381019390935260275460255491516024810193909352841660448301529092166064830152620012dc91608401620011c1565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602180549190920483167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055604080518082018252601a81527f5a657461436f6e6e6563746f724e6f6e4e61746976652e736f6c0000000000006020808301919091525460248054602754602554955193871692840192909252851660448301528416606482015291909216608482015262001402919060a40160408051601f198184030181529190526020810180516001600160e01b03167ff8c8765e000000000000000000000000000000000000000000000000000000001790526200c40a565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602280549190920483167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790556027546040517fca669fa700000000000000000000000000000000000000000000000000000000815291166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620014d8575f80fd5b505af1158015620014eb573d5f803e3d5ffd5b5050602480546027546022546040517f15d57fd40000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152908216938101939093521692506315d57fd491506044015f604051808303815f87803b1580156200155c575f80fd5b505af11580156200156f573d5f803e3d5ffd5b50506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d91506044015f604051808303815f87803b158015620015f1575f80fd5b505af115801562001604573d5f803e3d5ffd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d5691506024015f604051808303815f87803b15801562001678575f80fd5b505af11580156200168b573d5f803e3d5ffd5b50506020546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f91506024015f604051808303815f87803b158015620016ef575f80fd5b505af115801562001702573d5f803e3d5ffd5b50506020546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef91506024015f604051808303815f87803b15801562001766575f80fd5b505af115801562001779573d5f803e3d5ffd5b50506021546023546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639b19251a91506024015f604051808303815f87803b158015620017dd575f80fd5b505af1158015620017f0573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156200184f575f80fd5b505af115801562001862573d5f803e3d5ffd5b5050602354602554602c546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810191909152911692506340c10f1991506044015f604051808303815f87803b158015620018d1575f80fd5b505af1158015620018e4573d5f803e3d5ffd5b50506027546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801562001958575f80fd5b505af11580156200196b573d5f803e3d5ffd5b5050602254602554602c546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101919091525f60448201529116925063106e629091506064015f604051808303815f87803b158015620019e0575f80fd5b505af1158015620019f3573d5f803e3d5ffd5b50506040805160a0810182526103218082525f602080840182815284860193845285519182019095528181526060840181905260808401919091528251602880549551151574010000000000000000000000000000000000000000027fffffffffffffffffffffff0000000000000000000000000000000000000000009096166001600160a01b0392831617959095178555915160298054919093167fffffffffffffffffffffffff00000000000000000000000000000000000000009190911617909155909350909150602a9062001acd908262010f68565b50608091909101516003909101556020546040517f7c74425300000000000000000000000000000000000000000000000000000000815262030d4060048201526001600160a01b0390911690637c744253906024015f604051808303815f87803b15801562001b3a575f80fd5b505af115801562001b4d573d5f803e3d5ffd5b50505050565b60208054604080516328ae864d60e21b815290516001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562001b97573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062001bbd919062011030565b62001bca90600162010c76565b67ffffffffffffffff81111562001be55762001be562010ef3565b6040519080825280601f01601f19166020018201604052801562001c10576020820181803683370190505b50602a9062001c20908262010f68565b505f6028600201805462001c349062010d33565b905090505f60205f9054906101000a90046001600160a01b03166001600160a01b031663a2ba19346040518163ffffffff1660e01b8152600401602060405180830381865afa15801562001c8a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062001cb0919062011030565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162001d409160040162010c8c565b5f604051808303815f87803b15801562001d58575f80fd5b505af115801562001d6b573d5f803e3d5ffd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c935060019262001dc392169060289060040162010ea7565b5f604051808303818588803b15801562001ddb575f80fd5b505af115801562001dee573d5f803e3d5ffd5b50505050505050565b60235460205460405163095ea7b360e01b81526001600160a01b039182166004820152620186a0602482018190529261c35092169063095ea7b3906044016020604051808303815f875af115801562001e52573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062001e78919062011048565b506040515f602482015260448101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d9063f28dceb3907f8afe4db400000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b909216825262001f129160040162010c8c565b5f604051808303815f87803b15801562001f2a575f80fd5b505af115801562001f3d573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b09450869362001dc393811692899291169060289060040162011069565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052602754602554915163248e63e160e11b8152600160048201819052602482018190526044820181905260648201529293506001600160a01b03908116319291163190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562002047575f80fd5b505af11580156200205a573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f90620020ac9088905f908990602890620110a1565b60405180910390a36020546026546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263397e375c9288926200210d929091169083908990602890600401620110de565b5f604051808303818588803b15801562002125575f80fd5b505af115801562002138573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063491cc7c2925060840190505f604051808303815f87803b1580156200219e575f80fd5b505af1158015620021b1573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f90620022039088905f908990602890620110a1565b60405180910390a36020546001600160a01b031663397e375c6200222b62030d408762010c76565b6026546040516001600160e01b031960e085901b16815262002261916001600160a01b03169089908990602890600401620110de565b5f604051808303818588803b15801562002279575f80fd5b505af11580156200228c573d5f803e3d5ffd5b50506027546025546001600160a01b0391821631945016319150620022dc905062030d40620022bd88600262011107565b620022c9908762010c76565b620022d5919062010c76565b836200c38d565b62000c2262030d40620022f188600262011107565b620022fd908662011121565b62000dd4919062011121565b606060168054806020026020016040519081016040528092919081815260200182805480156200236157602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831162002342575b5050505050905090565b60235460205460405163095ea7b360e01b81526001600160a01b039182166004820152620186a06024820181905292919091169063095ea7b3906044016020604051808303815f875af1158015620023c5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620023eb919062011048565b506040805160a0810182526103218082525f602080840182905283850192909252835191820190935282815260608201526080810162002430621e8480600162010c76565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391620024cc9160040162010c8c565b5f604051808303815f87803b158015620024e4575f80fd5b505af1158015620024f7573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b0945062000bf7939283169288921690879060040162011137565b6020546040517f543f66a400000000000000000000000000000000000000000000000000000000815260016004820152620186a0916001600160a01b03169063543f66a4906024015f604051808303815f87803b1580156200259b575f80fd5b505af1158015620025ae573d5f803e3d5ffd5b5050602054602480546040517f4f27eb2a0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201525f928101929092529091169250634f27eb2a91506044015f604051808303815f87803b1580156200261c575f80fd5b505af11580156200262f573d5f803e3d5ffd5b5050602054602480546040517f6ae23fb50000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620026c8955092169250636ae23fb591015b602060405180830381865afa1580156200269c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620026c2919062011048565b6200c42c565b6020546040517f543f66a40000000000000000000000000000000000000000000000000000000081525f60048201526001600160a01b039091169063543f66a4906024015f604051808303815f87803b15801562002724575f80fd5b505af115801562002737573d5f803e3d5ffd5b50506020546040517f543f66a4000000000000000000000000000000000000000000000000000000008152600160048201526001600160a01b03909116925063543f66a491506024015f604051808303815f87803b15801562002798575f80fd5b505af1158015620027ab573d5f803e3d5ffd5b5050602054602480546040517f6ae23fb50000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015262002801955092169250636ae23fb5910162002680565b60208054604080517f02befd24000000000000000000000000000000000000000000000000000000008152905162002890936001600160a01b03909316926302befd2492600480820193918290030181865afa15801562002864573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200288a919062011048565b6200c4a4565b6024805460205460405163095ea7b360e01b81526001600160a01b039182166004820152928301849052169063095ea7b3906044016020604051808303815f875af1158015620028e2573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002908919062011048565b5060248054604080516001600160a01b039092168284015280518083039093018352604490910181526020820180516001600160e01b03167f33e23dbb000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916200299a919060040162010c8c565b5f604051808303815f87803b158015620029b2575f80fd5b505af1158015620029c5573d5f803e3d5ffd5b5050602054602654602454604051630102614b60e41b81526001600160a01b03938416955063102614b0945062002a0a93928316928792169060289060040162011069565b5f604051808303815f87803b15801562002a22575f80fd5b505af115801562002a35573d5f803e3d5ffd5b5050505050565b6060601e805480602002602001604051908101604052809291908181526020015f905b8282101562002b7b575f84815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b8282101562002b63578382905f5260205f2001805462002ad19062010d33565b80601f016020809104026020016040519081016040528092919081815260200182805462002aff9062010d33565b801562002b4e5780601f1062002b245761010080835404028352916020019162002b4e565b820191905f5260205f20905b81548152906001019060200180831162002b3057829003601f168201915b50505050508152602001906001019062002ab1565b50505050815250508152602001906001019062002a5f565b50505050905090565b602654604080516001600160a01b0392831660248083018290528351808403909101815260449092018352602080830180516001600160e01b0316630427d73b60e51b1790525492517f282906ed000000000000000000000000000000000000000000000000000000008152620186a09492939092169163282906ed91859162002c179190839060289060040162010eca565b5f604051808303818588803b15801562002c2f575f80fd5b505af115801562002c42573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392507f84222ac700000000000000000000000000000000000000000000000000000000915062002c95905062030d408662010c76565b62002ca462030d408762010c76565b62002cb190600162010c76565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b909216825262002d0b9160040162010c8c565b5f604051808303815f87803b15801562002d23575f80fd5b505af115801562002d36573d5f803e3d5ffd5b50506020546001600160a01b0316915063397e375c905062002d5c62030d408562010c76565b62002d6990600162010c76565b6026546040516001600160e01b031960e085901b16815262001dc3916001600160a01b03169087908790602890600401620110de565b6023546021546040516370a0823160e01b81526001600160a01b039182166004820152620186a0925f9216906370a0823190602401602060405180830381865afa15801562002df0573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002e16919062011030565b905062002e245f826200c38d565b6026546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790526023549054915163095ea7b360e01b81526001600160a01b03928316600482015260248101879052929350169063095ea7b3906044016020604051808303815f875af115801562002ebd573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002ee3919062011048565b506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562002f70575f80fd5b505af115801562002f83573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f9262002fdb928992909116908790602890620110a1565b60405180910390a36020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b789362003041939082169289929091169087906028906004016201116f565b5f604051808303815f87803b15801562003059575f80fd5b505af11580156200306c573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa158015620030bd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620030e3919062011030565b9050620030f184826200c38d565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201525f9291909116906370a0823190602401602060405180830381865afa15801562003140573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062003166919062011030565b905062002a3585602c5462000dd4919062011121565b6020546026546040517f726ac97c000000000000000000000000000000000000000000000000000000008152620186a0926001600160a01b039081169263726ac97c928592620031d492169060289060040162010ea7565b5f604051808303818588803b158015620031ec575f80fd5b505af1158015620031ff573d5f803e3d5ffd5b5050604051630618f58760e51b81527fef564bc9000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063c31eb0e0925060240190505f604051808303815f87803b1580156200326f575f80fd5b505af115801562003282573d5f803e3d5ffd5b50506020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063282906ed935062030d409262000fc09216905f9060289060040162010eca565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052515f602482015261c35060448201819052919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9063f28dceb3907f8afe4db400000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b9092168252620033c09160040162010c8c565b5f604051808303815f87803b158015620033d8575f80fd5b505af1158015620033eb573d5f803e3d5ffd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb49350859262001dc39216908790602890600401620111c7565b60275460255460405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152620186a0926001600160a01b039081163192163190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b158015620034bf575f80fd5b505af1158015620034d2573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c90620035229087905f9060289062010e71565b60405180910390a36020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263282906ed928792620035819290911690839060289060040162010eca565b5f604051808303818588803b15801562003599575f80fd5b505af1158015620035ac573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063491cc7c2925060840190505f604051808303815f87803b15801562003612575f80fd5b505af115801562003625573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c90620036759087905f9060289062010e71565b60405180910390a36020546001600160a01b031663282906ed6200369d62030d408662010c76565b6026546040516001600160e01b031960e085901b168152620036d1916001600160a01b031690889060289060040162010eca565b5f604051808303818588803b158015620036e9575f80fd5b505af1158015620036fc573d5f803e3d5ffd5b50506027546025546001600160a01b03918216319450163191506200372d905062030d40620022bd87600262011107565b62002a3562030d40620022f187600262011107565b606060188054806020026020016040519081016040528092919081815260200182805480156200236157602002820191905f5260205f209081546001600160a01b0316815260019091019060200180831162002342575050505050905090565b606060178054806020026020016040519081016040528092919081815260200182805480156200236157602002820191905f5260205f209081546001600160a01b0316815260019091019060200180831162002342575050505050905090565b6027546020546040517f543f66a400000000000000000000000000000000000000000000000000000000815260016004820152620186a0926001600160a01b039081163192169063543f66a4906024015f604051808303815f87803b1580156200386a575f80fd5b505af11580156200387d573d5f803e3d5ffd5b50506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a4015f604051808303815f87803b1580156200390d575f80fd5b505af115801562003920573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c90620039709086905f9060289062010e71565b60405180910390a36020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263726ac97c928692620039cd929091169060289060040162010ea7565b5f604051808303818588803b158015620039e5575f80fd5b505af1158015620039f8573d5f803e3d5ffd5b505050505062003a22828262003a0f919062010c76565b6027546001600160a01b0316316200c38d565b5050565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790526023549054915163095ea7b360e01b81526001600160a01b03928316600482015260248101869052929350169063095ea7b3906044016020604051808303815f875af115801562003ac4573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062003aea919062011048565b506040805160a0810182526103218082525f602080840182905283850192909252835191820190935282815260608201526080810162003b2f621e8480600162010c76565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162003bcb9160040162010c8c565b5f604051808303815f87803b15801562003be3575f80fd5b505af115801562003bf6573d5f803e3d5ffd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b78945062003c5593928316928992169088908890600401620111fe565b5f604051808303815f87803b15801562001ddb575f80fd5b60208054604080516328ae864d60e21b81529051620186a0935f936002936001600160a01b039091169263a2ba1934926004808401939192918290030181865afa15801562003cbe573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062003ce4919062011030565b62003cf091906201124a565b67ffffffffffffffff81111562003d0b5762003d0b62010ef3565b6040519080825280601f01601f19166020018201604052801562003d36576020820181803683370190505b5060208054604080516328ae864d60e21b815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562003d81573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062003da7919062011030565b62003db391906201124a565b62003dc090600162010c76565b67ffffffffffffffff81111562003ddb5762003ddb62010ef3565b6040519080825280601f01601f19166020018201604052801562003e06576020820181803683370190505b50602a9062003e16908262010f68565b505f6028600201805462003e2a9062010d33565b835162003e38925062010c76565b60208054604080516328ae864d60e21b815290519394505f936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562003e81573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062003ea7919062011030565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162003f379160040162010c8c565b5f604051808303815f87803b15801562003f4f575f80fd5b505af115801562003f62573d5f803e3d5ffd5b50506020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063744b9b8b9350889262003fbb9216908890602890600401620111c7565b5f604051808303818588803b15801562003fd3575f80fd5b505af115801562003fe6573d5f803e3d5ffd5b505050505050505050565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052815160a0810183526103218082525f8284018190528285019190915283519283019093528282526060810191909152919250906080810162004083621e8480600162010c76565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916200411f9160040162010c8c565b5f604051808303815f87803b15801562004137575f80fd5b505af11580156200414a573d5f803e3d5ffd5b50506020546026546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063397e375c93508792620041a492169083908890889060040162011283565b5f604051808303818588803b158015620041bc575f80fd5b505af1158015620041cf573d5f803e3d5ffd5b5050505050505050565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052602554602754915163248e63e160e11b8152600160048201819052602482018190526044820181905260648201529293506001600160a01b03908116319291163190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562004297575f80fd5b505af1158015620042aa573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d97490620042f8908790602890620112c0565b60405180910390a36020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692631becceb49262004354929116908790602890600401620111c7565b5f604051808303815f87803b1580156200436c575f80fd5b505af11580156200437f573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b158015620043e3575f80fd5b505af1158015620043f6573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d9749062004444908790602890620112c0565b60405180910390a36020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692631becceb49262030d4092620044a692909116908890602890600401620111c7565b5f604051808303818588803b158015620044be575f80fd5b505af1158015620044d1573d5f803e3d5ffd5b50506025546027546001600160a01b0391821631945016319150620045009050620022d562030d408662011121565b62002a3562000dd462030d408562010c76565b6026546040516001600160a01b039091166024820152620186a09061c350905f9060440160408051601f19818403018152919052602080820180516001600160e01b0316630427d73b60e51b17905260235490549192506001600160a01b039081169163095ea7b391166200458a86600262011107565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af1158015620045d3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620045f9919062011048565b506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b789362004658939082169289929091169087906028906004016201116f565b5f604051808303815f87803b15801562004670575f80fd5b505af115801562004683573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507f8afe4db400000000000000000000000000000000000000000000000000000000905062030d40620046d5868262010c76565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b90921682526200472f9160040162010c8c565b5f604051808303815f87803b15801562004747575f80fd5b505af11580156200475a573d5f803e3d5ffd5b50506020546001600160a01b0316915063d09e3b789050620047808462030d4062010c76565b6026546023546040516001600160e01b031960e086901b168152620041a4926001600160a01b03908116928a9291169088906028906004016201116f565b620186a05f620047d3600162030d4062011121565b6026546040516001600160a01b0390911660248201529091505f9060440160408051601f19818403018152919052602080820180516001600160e01b0316630427d73b60e51b17905260235490549192506001600160a01b039081169163095ea7b391166200484486600262011107565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af11580156200488d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620048b3919062011048565b506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b789362004912939082169289929091169087906028906004016201116f565b5f604051808303815f87803b1580156200492a575f80fd5b505af11580156200493d573d5f803e3d5ffd5b505060405162030d40602482015260448101859052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507fa458261b00000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b9092168252620049dd9160040162010c8c565b5f604051808303815f87803b158015620049f5575f80fd5b505af115801562004a08573d5f803e3d5ffd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b7894508793620041a4938116928a9291169088906028906004016201116f565b6026546040516001600160a01b0390911660248201525f90819060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fef564bc9000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b15801562004b1d575f80fd5b505af115801562004b30573d5f803e3d5ffd5b50506020546026546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063397e375c9350869262001dc392169083908790602890600401620110de565b60235460205460405163095ea7b360e01b81526001600160a01b039182166004820152620186a06024820181905292919091169063095ea7b3906044016020604051808303815f875af115801562004be5573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004c0b919062011048565b5060208054604080516328ae864d60e21b815290516001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562004c50573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004c76919062011030565b62004c8390600162010c76565b67ffffffffffffffff81111562004c9e5762004c9e62010ef3565b6040519080825280601f01601f19166020018201604052801562004cc9576020820181803683370190505b50602a9062004cd9908262010f68565b505f6028600201805462004ced9062010d33565b905090505f60205f9054906101000a90046001600160a01b03166001600160a01b031663a2ba19346040518163ffffffff1660e01b8152600401602060405180830381865afa15801562004d43573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004d69919062011030565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162004df99160040162010c8c565b5f604051808303815f87803b15801562004e11575f80fd5b505af115801562004e24573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b0945062003c5593928316928992169060289060040162011069565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562004f19575f80fd5b505af115801562004f2c573d5f803e3d5ffd5b50506020546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250631becceb4915062002a0a905f908590602890600401620111c7565b620186a0737109709ecfa91a80626ff3989d68f67f5b1dd12d63f28dceb37f84222ac70000000000000000000000000000000000000000000000000000000062004fcc84600162010c76565b60405160248101919091526044810185905260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b9092168252620050289160040162010c8c565b5f604051808303815f87803b15801562005040575f80fd5b505af115801562005053573d5f803e3d5ffd5b50506020546026546001600160a01b03918216935063282906ed92508491166200507f82600162010c76565b60286040518563ffffffff1660e01b815260040162000fc09392919062010eca565b6026546040516001600160a01b0390911660248201526001905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562005154575f80fd5b505af115801562005167573d5f803e3d5ffd5b50506020546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063744b9b8b9150849062001dc3905f908690602890600401620111c7565b60208054604080516328ae864d60e21b815290516001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562005201573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005227919062011030565b6200523490600162010c76565b67ffffffffffffffff8111156200524f576200524f62010ef3565b6040519080825280601f01601f1916602001820160405280156200527a576020820181803683370190505b50602a906200528a908262010f68565b505f602860020180546200529e9062010d33565b905090505f60205f9054906101000a90046001600160a01b03166001600160a01b031663a2ba19346040518163ffffffff1660e01b8152600401602060405180830381865afa158015620052f4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200531a919062011030565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391620053aa9160040162010c8c565b5f604051808303815f87803b158015620053c2575f80fd5b505af1158015620053d5573d5f803e3d5ffd5b50506020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063282906ed935060019262001dc3921690839060289060040162010eca565b6020546023546040517f6ae23fb50000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015262005484929190911690636ae23fb59060240162002680565b60405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b158015620054e4575f80fd5b505af1158015620054f7573d5f803e3d5ffd5b5050602354604051600181526001600160a01b0390911692507f50708318fef6b3b62fbd0894ea2c5d8fc3ccce4785ed240e4c525ae40bd23332915060200160405180910390a26020546023546040517f4f27eb2a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260016024820152911690634f27eb2a906044015f604051808303815f87803b158015620055a5575f80fd5b505af1158015620055b8573d5f803e3d5ffd5b50506020546023546040517f6ae23fb50000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015262005626945091169150636ae23fb5906024015b602060405180830381865afa15801562002864573d5f803e3d5ffd5b565b60405163248e63e160e11b8152600160048201819052602482018190526044820181905260648201526509184e72a00090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562005690575f80fd5b505af1158015620056a3573d5f803e3d5ffd5b50506040805162030d408152602081018590527fe60a2882aff7841a7d0492b804d46ff2d9b07ea353dbc3791e8276b729e40cd8935001905060405180910390a16020546040517f7c744253000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b0390911690637c744253906024015f604051808303815f87803b15801562005741575f80fd5b505af115801562005754573d5f803e3d5ffd5b505060208054604080517fb01072140000000000000000000000000000000000000000000000000000000081529051620057e295506001600160a01b03909216935063b01072149260048083019391928290030181865afa158015620057bc573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062000dd4919062011030565b50565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529190526020810180516001600160e01b0316630427d73b60e51b17905290505f6200583d600162030d4062011121565b6020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081529293506001600160a01b0391821692631becceb492620058939216908690602890600401620111c7565b5f604051808303815f87803b158015620058ab575f80fd5b505af1158015620058be573d5f803e3d5ffd5b505060405162030d40602482015260448101849052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507fa458261b000000000000000000000000000000000000000000000000000000009060640162003379565b60405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b1580156200597c575f80fd5b505af11580156200598f573d5f803e3d5ffd5b5050602454604051600181526001600160a01b0390911692507f50708318fef6b3b62fbd0894ea2c5d8fc3ccce4785ed240e4c525ae40bd23332915060200160405180910390a260405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562005a36575f80fd5b505af115801562005a49573d5f803e3d5ffd5b5050604051600181525f92507f50708318fef6b3b62fbd0894ea2c5d8fc3ccce4785ed240e4c525ae40bd23332915060200160405180910390a260405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562005ae3575f80fd5b505af115801562005af6573d5f803e3d5ffd5b5050604051600181527fd39ad370883b0cd1a8172b5b006a3ebcaaf65183008c91ffd7655afb74174e579250602001905060405180910390a16020546040517f543f66a4000000000000000000000000000000000000000000000000000000008152600160048201526001600160a01b039091169063543f66a4906024015f604051808303815f87803b15801562005b8c575f80fd5b505af115801562005b9f573d5f803e3d5ffd5b505060208054604080517f02befd24000000000000000000000000000000000000000000000000000000008152905162005c0795506001600160a01b0390921693506302befd249260048083019391928290030181865afa15801562002864573d5f803e3d5ffd5b6020546040517f6ae23fb50000000000000000000000000000000000000000000000000000000081525f600482015262005c54916001600160a01b031690636ae23fb5906024016200560a565b60405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562005cb4575f80fd5b505af115801562005cc7573d5f803e3d5ffd5b50506040515f81527fd39ad370883b0cd1a8172b5b006a3ebcaaf65183008c91ffd7655afb74174e579250602001905060405180910390a16020546040517f543f66a40000000000000000000000000000000000000000000000000000000081525f60048201526001600160a01b039091169063543f66a4906024015f604051808303815f87803b15801562005d5b575f80fd5b505af115801562005d6e573d5f803e3d5ffd5b505060208054604080517f02befd2400000000000000000000000000000000000000000000000000000000815290516200562695506001600160a01b0390921693506302befd249260048083019391928290030181865afa1580156200269c573d5f803e3d5ffd5b6020546040517f543f66a400000000000000000000000000000000000000000000000000000000815260016004820152620186a0916001600160a01b03169063543f66a4906024015f604051808303815f87803b15801562005e36575f80fd5b505af115801562005e49573d5f803e3d5ffd5b50506024805460205460405163095ea7b360e01b81526001600160a01b03918216600482015292830186905216925063095ea7b391506044016020604051808303815f875af115801562005e9f573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005ec5919062011048565b50602054602654602454604051630102614b60e41b81526001600160a01b039384169363102614b09362005f09939082169287929091169060289060040162011069565b5f604051808303815f87803b15801562005f21575f80fd5b505af115801562005f34573d5f803e3d5ffd5b50505050620057e281602c5462005f4c919062011121565b602480546025546040516370a0823160e01b81526001600160a01b0391821660048201529116916370a0823191015b602060405180830381865afa15801562005f97573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005fbd919062011030565b6200c38d565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905260235460215491516370a0823160e01b81526001600160a01b0392831660048201529293505f929116906370a0823190602401602060405180830381865afa1580156200605c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062006082919062011030565b6027546025546023546020549394506001600160a01b03928316319391831631929081169163095ea7b39116620060bb88600262011107565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801562006104573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200612a919062011048565b5060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b1580156200618b575f80fd5b505af11580156200619e573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c92620061f4928b929091169060289062010e71565b60405180910390a3602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b0936200623f93908216928b929091169060289060040162011069565b5f604051808303815f87803b15801562006257575f80fd5b505af11580156200626a573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b158015620062ce575f80fd5b505af1158015620062e1573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f9262006339928b92909116908a90602890620110a1565b60405180910390a36020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b789362030d4093620063a393918316928c929116908b906028906004016201116f565b5f604051808303818588803b158015620063bb575f80fd5b505af1158015620063ce573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063491cc7c2925060840190505f604051808303815f87803b15801562006434575f80fd5b505af115801562006447573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d9749062006495908890602890620112c0565b60405180910390a36020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692631becceb49262030d4092620064f792909116908990602890600401620111c7565b5f604051808303818588803b1580156200650f575f80fd5b505af115801562006522573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9550911692506370a082319150602401602060405180830381865afa15801562006574573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200659a919062011030565b6027546025549192506001600160a01b0390811631911631620065d6620065c389600262011107565b620065cf908862010c76565b846200c38d565b620065f6620065ea62030d40600262011107565b620022d5908762010c76565b620041cf6200660a62030d40600262011107565b62000dd4908662011121565b6023546021546040516370a0823160e01b81526001600160a01b039182166004820152620186a0925f9216906370a0823190602401602060405180830381865afa15801562006667573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200668d919062011030565b90506200669b5f826200c38d565b60235460205460405163095ea7b360e01b81526001600160a01b0391821660048201526024810185905291169063095ea7b3906044016020604051808303815f875af1158015620066ee573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062006714919062011048565b506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015620067a1575f80fd5b505af1158015620067b4573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c926200680a9288929091169060289062010e71565b60405180910390a3602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b09362006855939082169288929091169060289060040162011069565b5f604051808303815f87803b1580156200686d575f80fd5b505af115801562006880573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9450911691506370a0823190602401602060405180830381865afa158015620068d1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620068f7919062011030565b90506200690583826200c38d565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201525f9291909116906370a0823190602401602060405180830381865afa15801562006954573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200697a919062011030565b905062001b4d84602c5462000dd4919062011121565b6060601b805480602002602001604051908101604052809291908181526020015f905b8282101562002b7b578382905f5260205f2090600202016040518060400160405290815f82018054620069e69062010d33565b80601f016020809104026020016040519081016040528092919081815260200182805462006a149062010d33565b801562006a635780601f1062006a395761010080835404028352916020019162006a63565b820191905f5260205f20905b81548152906001019060200180831162006a4557829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801562006ae757602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841162006aa85790505b50505050508152505081526020019060010190620069b3565b620186a0737109709ecfa91a80626ff3989d68f67f5b1dd12d63f28dceb37f84222ac70000000000000000000000000000000000000000000000000000000062006b4c60018562011121565b60405160248101919091526044810185905260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b909216825262006ba89160040162010c8c565b5f604051808303815f87803b15801562006bc0575f80fd5b505af115801562006bd3573d5f803e3d5ffd5b50506020546026546001600160a01b03918216935063282906ed92508491166200507f60018362011121565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905260235460215491516370a0823160e01b81526001600160a01b0392831660048201529293505f929116906370a0823190602401602060405180830381865afa15801562006c98573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062006cbe919062011030565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa15801562006d0e573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062006d34919062011030565b6027546023546020549293506001600160a01b0391821631929082169163095ea7b3911662006d6588600262011107565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801562006dae573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062006dd4919062011048565b5060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b15801562006e35575f80fd5b505af115801562006e48573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f9262006ea0928b92909116908a90602890620110a1565b60405180910390a36020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b789362006f0693908216928b92909116908a906028906004016201116f565b5f604051808303815f87803b15801562006f1e575f80fd5b505af115801562006f31573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562006f95575f80fd5b505af115801562006fa8573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f9262007000928b92909116908a90602890620110a1565b60405180910390a36020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b789362030d40936200706a93918316928c929116908b906028906004016201116f565b5f604051808303818588803b15801562007082575f80fd5b505af115801562007095573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9550911692506370a082319150602401602060405180830381865afa158015620070e7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200710d919062011030565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa1580156200715d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062007183919062011030565b6027549091506001600160a01b031631620071a4620065c389600262011107565b620071c1620071b589600262011107565b620022d5908762011121565b620041cf62000dd462030d408662010c76565b6027546026546040516001600160a01b039182166024820152620186a0929190911631905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b158015620072b4575f80fd5b505af1158015620072c7573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f90620073199087905f908790602890620110a1565b60405180910390a36020546026546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263397e375c9287926200737a929091169083908790602890600401620110de565b5f604051808303818588803b15801562007392575f80fd5b505af1158015620073a5573d5f803e3d5ffd5b50506027546001600160a01b031631925062001b4d915062000dd49050858562010c76565b604051630618f58760e51b81527fef564bc90000000000000000000000000000000000000000000000000000000060048201525f90737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b15801562007437575f80fd5b505af11580156200744a573d5f803e3d5ffd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c9350859262000fc092169060289060040162010ea7565b602654604080516001600160a01b0392831660248083018290528351808403909101815260449092018352602080830180516001600160e01b0316630427d73b60e51b1790525492517f744b9b8b000000000000000000000000000000000000000000000000000000008152620186a09492939092169163744b9b8b9185916200753491908690602890600401620111c7565b5f604051808303818588803b1580156200754c575f80fd5b505af11580156200755f573d5f803e3d5ffd5b50506040805160048082526024820183526020820180516001600160e01b03167f394836a400000000000000000000000000000000000000000000000000000000179052915163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d955063f28dceb39450620075dc935090910162010c8c565b5f604051808303815f87803b158015620075f4575f80fd5b505af115801562007607573d5f803e3d5ffd5b50506020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063744b9b8b9350869262001dc39216908690602890600401620111c7565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015620076cd575f80fd5b505af1158015620076e0573d5f803e3d5ffd5b5050602054602354604051630102614b60e41b81526001600160a01b03928316945063102614b0935062002a0a925f92879291169060289060040162011069565b6026546040516001600160a01b0390911660248201525f90819060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fef564bc9000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401620075dc565b6060601a805480602002602001604051908101604052809291908181526020015f905b8282101562002b7b578382905f5260205f20018054620078049062010d33565b80601f0160208091040260200160405190810160405280929190818152602001828054620078329062010d33565b8015620078815780601f10620078575761010080835404028352916020019162007881565b820191905f5260205f20905b8154815290600101906020018083116200786357829003601f168201915b505050505081526020019060010190620077e4565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529190526020810180516001600160e01b0316630427d73b60e51b1790529050737109709ecfa91a80626ff3989d68f67f5b1dd12d63f28dceb37f84222ac7000000000000000000000000000000000000000000000000000000006200792a85600162010c76565b60405160248101919091526044810186905260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b9092168252620079869160040162010c8c565b5f604051808303815f87803b1580156200799e575f80fd5b505af1158015620079b1573d5f803e3d5ffd5b50506020546026546001600160a01b03918216935063397e375c9250859116620079dd82600162010c76565b8560286040518663ffffffff1660e01b815260040162001dc39493929190620110de565b60235460205460405163095ea7b360e01b81526001600160a01b0391821660048201525f6024820181905292919091169063095ea7b3906044016020604051808303815f875af115801562007a58573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062007a7e919062011048565b50604051630618f58760e51b81527fef564bc9000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b15801562007aea575f80fd5b505af115801562007afd573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b0945062002a0a93928316928792169060289060040162011069565b6040805160a0810182526103218082525f60208084018290528385019290925283519182019093528281526060820152620186a091906080810162007b8c621e8480600162010c76565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162007c289160040162010c8c565b5f604051808303815f87803b15801562007c40575f80fd5b505af115801562007c53573d5f803e3d5ffd5b50506020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063282906ed9350869262001dc392169083908790600401620112e8565b6026546040516001600160a01b0390911660248201526001905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562007d5e575f80fd5b505af115801562007d71573d5f803e3d5ffd5b50506020546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063d09e3b78935062000bf7925f92889291169087906028906004016201116f565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525460235491517f4f27eb2a0000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152600160248201529293501690634f27eb2a906044015f604051808303815f87803b15801562007e7e575f80fd5b505af115801562007e91573d5f803e3d5ffd5b50506020546040517f543f66a4000000000000000000000000000000000000000000000000000000008152600160048201526001600160a01b03909116925063543f66a491506024015f604051808303815f87803b15801562007ef2575f80fd5b505af115801562007f05573d5f803e3d5ffd5b50506023546020546001600160a01b03918216935063095ea7b392501662007f2f85600262011107565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af115801562007f78573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062007f9e919062011048565b50602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b09362007fe2939082169288929091169060289060040162011069565b5f604051808303815f87803b15801562007ffa575f80fd5b505af11580156200800d573d5f803e3d5ffd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b78945062030d40936200807293811692899291169088906028906004016201116f565b5f604051808303818588803b1580156200808a575f80fd5b505af11580156200809d573d5f803e3d5ffd5b505050505062003a22826002620080b5919062011107565b6023546021546040516370a0823160e01b81526001600160a01b0391821660048201529116906370a082319060240162005f7b565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b15801562008157575f80fd5b505af11580156200816a573d5f803e3d5ffd5b50506020546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063726ac97c9150839062000fc0905f9060289060040162010ea7565b602654604080516001600160a01b0392831660248083018290528351808403909101815260449092018352602080830180516001600160e01b0316630427d73b60e51b1790525492517f282906ed000000000000000000000000000000000000000000000000000000008152620186a09492939092169163282906ed918591620082519190839060289060040162010eca565b5f604051808303818588803b15801562008269575f80fd5b505af11580156200827c573d5f803e3d5ffd5b505060405162030d40602482015260448101869052737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392507fa458261b0000000000000000000000000000000000000000000000000000000091506064015b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b909216825262004b059160040162010c8c565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f198184030181529190526020810180516001600160e01b0316630427d73b60e51b1790529050737109709ecfa91a80626ff3989d68f67f5b1dd12d63f28dceb37f84222ac700000000000000000000000000000000000000000000000000000000620083b160018662011121565b60405160248101919091526044810186905260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b90921682526200840d9160040162010c8c565b5f604051808303815f87803b15801562008425575f80fd5b505af115801562008438573d5f803e3d5ffd5b50506020546026546001600160a01b03918216935063397e375c9250859116620079dd60018362011121565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f543f66a4000000000000000000000000000000000000000000000000000000008152600160048201529192506001600160a01b03169063543f66a4906024015f604051808303815f87803b1580156200850a575f80fd5b505af11580156200851d573d5f803e3d5ffd5b50506020546040517f4f27eb2a0000000000000000000000000000000000000000000000000000000081525f6004820181905260248201526001600160a01b039091169250634f27eb2a91506044015f604051808303815f87803b15801562008584575f80fd5b505af115801562008597573d5f803e3d5ffd5b5050604080515f60248083019190915282518083039091018152604490910182526020810180516001600160e01b03167f33e23dbb00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925062008624919060040162010c8c565b5f604051808303815f87803b1580156200863c575f80fd5b505af11580156200864f573d5f803e3d5ffd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c93508692620086a692169060289060040162010ea7565b5f604051808303818588803b158015620086be575f80fd5b505af1158015620086d1573d5f803e3d5ffd5b5050604080515f60248083019190915282518083039091018152604490910182526020810180516001600160e01b03167f33e23dbb00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d945063f28dceb393506200875e925060040162010c8c565b5f604051808303815f87803b15801562008776575f80fd5b505af115801562008789573d5f803e3d5ffd5b50506020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063282906ed93508692620087e2921690839060289060040162010eca565b5f604051808303818588803b158015620087fa575f80fd5b505af11580156200880d573d5f803e3d5ffd5b5050604080515f60248083019190915282518083039091018152604490910182526020810180516001600160e01b03167f33e23dbb00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d945063f28dceb393506200889a925060040162010c8c565b5f604051808303815f87803b158015620088b2575f80fd5b505af1158015620088c5573d5f803e3d5ffd5b50506020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063744b9b8b935086926200891e9216908690602890600401620111c7565b5f604051808303818588803b15801562008936575f80fd5b505af115801562008949573d5f803e3d5ffd5b50506040515f6024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392507f33e23dbb000000000000000000000000000000000000000000000000000000009150604401620082d6565b6060601d805480602002602001604051908101604052809291908181526020015f905b8282101562002b7b575f8481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801562008a6a57602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841162008a2b5790505b50505050508152505081526020019060010190620089c1565b6020546026546040517f726ac97c000000000000000000000000000000000000000000000000000000008152620186a0926001600160a01b039081169263726ac97c92859262008adb92169060289060040162010ea7565b5f604051808303818588803b15801562008af3575f80fd5b505af115801562008b06573d5f803e3d5ffd5b50506040805160048082526024820183526020820180516001600160e01b03167f394836a400000000000000000000000000000000000000000000000000000000179052915163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d955063f28dceb394506200741f935090910162010c8c565b60208054604080516328ae864d60e21b81529051620186a0935f936002936001600160a01b039091169263a2ba1934926004808401939192918290030181865afa15801562008bd4573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062008bfa919062011030565b62008c0691906201124a565b67ffffffffffffffff81111562008c215762008c2162010ef3565b6040519080825280601f01601f19166020018201604052801562008c4c576020820181803683370190505b5060208054604080516328ae864d60e21b815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562008c97573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062008cbd919062011030565b62008cc991906201124a565b62008cd690600162010c76565b67ffffffffffffffff81111562008cf15762008cf162010ef3565b6040519080825280601f01601f19166020018201604052801562008d1c576020820181803683370190505b50602a9062008d2c908262010f68565b5060235460205460405163095ea7b360e01b81526001600160a01b0391821660048201526024810185905291169063095ea7b3906044016020604051808303815f875af115801562008d80573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062008da6919062011048565b505f6028600201805462008dba9062010d33565b835162008dc8925062010c76565b60208054604080516328ae864d60e21b815290519394505f936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562008e11573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062008e37919062011030565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162008ec79160040162010c8c565b5f604051808303815f87803b15801562008edf575f80fd5b505af115801562008ef2573d5f803e3d5ffd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b78945062008f5293928316928a92169089906028906004016201116f565b5f604051808303815f87803b158015620041bc575f80fd5b6026546040516001600160a01b0390911660248201526001905f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156200901d575f80fd5b505af115801562009030573d5f803e3d5ffd5b50506020546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063397e375c9150849062001dc3905f9083908790602890600401620110de565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015620090f5575f80fd5b505af115801562009108573d5f803e3d5ffd5b50506020546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063282906ed9150839062000fc0905f90839060289060040162010eca565b6027546026546040516001600160a01b039182166024820152620186a0929190911631905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b1580156200923e575f80fd5b505af115801562009251573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f90620092a39087905f908790602890620110a1565b60405180910390a36020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263744b9b8b9287926200737a92909116908690602890600401620111c7565b60208054604080516328ae864d60e21b81529051620186a0935f936002936001600160a01b039091169263a2ba1934926004808401939192918290030181865afa15801562009353573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062009379919062011030565b6200938591906201124a565b67ffffffffffffffff811115620093a057620093a062010ef3565b6040519080825280601f01601f191660200182016040528015620093cb576020820181803683370190505b5060208054604080516328ae864d60e21b815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562009416573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200943c919062011030565b6200944891906201124a565b6200945590600162010c76565b67ffffffffffffffff81111562009470576200947062010ef3565b6040519080825280601f01601f1916602001820160405280156200949b576020820181803683370190505b50602a90620094ab908262010f68565b505f60286002018054620094bf9062010d33565b8351620094cd925062010c76565b60208054604080516328ae864d60e21b815290519394505f936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562009516573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200953c919062011030565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb391620095cc9160040162010c8c565b5f604051808303815f87803b158015620095e4575f80fd5b505af1158015620095f7573d5f803e3d5ffd5b50506020546026546040517f397e375c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063397e375c9350889262003fbb92169083908990602890600401620110de565b6027546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201526509184e72a00090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620096ca575f80fd5b505af1158015620096dd573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156200973c575f80fd5b505af11580156200974f573d5f803e3d5ffd5b50506020546040517f7c744253000000000000000000000000000000000000000000000000000000008152600481018590526001600160a01b039091169250637c744253915060240162002a0a565b6026546040516001600160a01b0390911660248201525f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b15801562009870575f80fd5b505af115801562009883573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d97490620098d1908590602890620112c0565b60405180910390a36020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692631becceb49262002a0a929116908590602890600401620111c7565b6060601c805480602002602001604051908101604052809291908181526020015f905b8282101562002b7b575f8481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015620099f957602002820191905f5260205f20905f905b82829054906101000a900460e01b6001600160e01b03191681526020019060040190602082600301049283019260010382029150808411620099ba5790505b5050505050815250508152602001906001019062009950565b60208054604080516328ae864d60e21b815290515f936002936001600160a01b03169263a2ba193492600480830193928290030181865afa15801562009a5a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062009a80919062011030565b62009a8c91906201124a565b67ffffffffffffffff81111562009aa75762009aa762010ef3565b6040519080825280601f01601f19166020018201604052801562009ad2576020820181803683370190505b5060208054604080516328ae864d60e21b815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562009b1d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062009b43919062011030565b62009b4f91906201124a565b62009b5c90600162010c76565b67ffffffffffffffff81111562009b775762009b7762010ef3565b6040519080825280601f01601f19166020018201604052801562009ba2576020820181803683370190505b50602a9062009bb2908262010f68565b505f6028600201805462009bc69062010d33565b835162009bd4925062010c76565b60208054604080516328ae864d60e21b815290519394505f936001600160a01b039092169263a2ba1934926004808401938290030181865afa15801562009c1d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062009c43919062011030565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167f9fcf788e00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162009cd39160040162010c8c565b5f604051808303815f87803b15801562009ceb575f80fd5b505af115801562009cfe573d5f803e3d5ffd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb4935062003c5592909116908790602890600401620111c7565b6040805160a0810182526103218082525f60208084018290528385019290925283519182019093528281526060820152620186a091906080810162009da1621e8480600162010c76565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39162009e3d9160040162010c8c565b5f604051808303815f87803b15801562009e55575f80fd5b505af115801562009e68573d5f803e3d5ffd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c9350869262001dc3921690869060040162011311565b60235460205460405163095ea7b360e01b81526001600160a01b039182166004820152620186a06024820181905292919091169063095ea7b3906044016020604051808303815f875af115801562009f18573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062009f3e919062011048565b506025546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562009faf575f80fd5b505af115801562009fc2573d5f803e3d5ffd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a59042791506024015f604051808303815f87803b1580156200a026575f80fd5b505af11580156200a039573d5f803e3d5ffd5b5050602354604080516001600160a01b0390921660248084019190915281518084039091018152604490920181526020820180516001600160e01b03167f1387a349000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925062007ad2919060040162010c8c565b60606019805480602002602001604051908101604052809291908181526020015f905b8282101562002b7b578382905f5260205f200180546200a1159062010d33565b80601f01602080910402602001604051908101604052809291908181526020018280546200a1439062010d33565b80156200a1925780601f106200a168576101008083540402835291602001916200a192565b820191905f5260205f20905b8154815290600101906020018083116200a17457829003601f168201915b5050505050815260200190600101906200a0f5565b602354602054620186a09161c350916001600160a01b039182169163095ea7b391166200a1d685600262011107565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af11580156200a21f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a245919062011048565b50602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b0936200a289939082169288929091169060289060040162011069565b5f604051808303815f87803b1580156200a2a1575f80fd5b505af11580156200a2b4573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507f8afe4db400000000000000000000000000000000000000000000000000000000905062030d406200a306858262010c76565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b90921682526200a3609160040162010c8c565b5f604051808303815f87803b1580156200a378575f80fd5b505af11580156200a38b573d5f803e3d5ffd5b50506020546001600160a01b0316915063102614b090506200a3b18362030d4062010c76565b6026546023546040516001600160e01b031960e086901b16815262001dc3926001600160a01b0390811692899291169060289060040162011069565b6008545f9060ff16156200a405575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c6564000000000000000000000000000000000000000000000000000060248301525f9163667f9d7090604401602060405180830381865afa1580156200a494573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200a4ba919062011030565b1415905090565b6026546040516001600160a01b0390911660248201525f9060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b179052602880547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000017905551630618f58760e51b81527f19b4bff2000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b1580156200a5b0575f80fd5b505af11580156200a5c3573d5f803e3d5ffd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb4935062002a0a92909116908590602890600401620111c7565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052815160a0810183526103218082525f828401819052828501919091528351928301909352828252606081019190915291925090608081016200a6ae621e8480600162010c76565b90526080810151604080516024810192909252621e848060448084019190915281518084039091018152606490920181526020820180516001600160e01b03167fec86721e000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916200a74a9160040162010c8c565b5f604051808303815f87803b1580156200a762575f80fd5b505af11580156200a775573d5f803e3d5ffd5b50506020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063744b9b8b93508792620041a49216908790879060040162010cfc565b602654604080516001600160a01b0392831660248083018290528351808403909101815260449092018352602080830180516001600160e01b0316630427d73b60e51b1790525492517f1becceb4000000000000000000000000000000000000000000000000000000008152919361c350931691631becceb4916200a85a918690602890600401620111c7565b5f604051808303815f87803b1580156200a872575f80fd5b505af11580156200a885573d5f803e3d5ffd5b50737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507f8afe4db400000000000000000000000000000000000000000000000000000000905062030d406200a8d7858262010c76565b6040516024810192909252604482015260640160408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b90921682526200a9319160040162010c8c565b5f604051808303815f87803b1580156200a949575f80fd5b505af11580156200a95c573d5f803e3d5ffd5b50506020546001600160a01b03169150631becceb490506200a9828362030d4062010c76565b6026546040516001600160e01b031960e085901b16815262001dc3916001600160a01b0316908790602890600401620111c7565b604051630618f58760e51b81527fef564bc90000000000000000000000000000000000000000000000000000000060048201525f90737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b5f604051808303815f87803b1580156200aa23575f80fd5b505af11580156200aa36573d5f803e3d5ffd5b50506020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063282906ed9350859262000fc0921690839060289060040162010eca565b620186a05f6200aaa4600162030d4062011121565b6023546020549192506001600160a01b039081169163095ea7b391166200aacd85600262011107565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af11580156200ab16573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200ab3c919062011048565b50602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b0936200ab80939082169288929091169060289060040162011069565b5f604051808303815f87803b1580156200ab98575f80fd5b505af11580156200abab573d5f803e3d5ffd5b505060405162030d40602482015260448101849052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507fa458261b000000000000000000000000000000000000000000000000000000009060640162001ecb565b6020546040517f7c7442530000000000000000000000000000000000000000000000000000000081525f60048201526001600160a01b0390911690637c744253906024015f604051808303815f87803b1580156200ac65575f80fd5b505af11580156200ac78573d5f803e3d5ffd5b50506026546040516001600160a01b039091166024820152620186a092505f915060440160408051601f19818403018152919052602080820180516001600160e01b0316630427d73b60e51b17905260235490549192506001600160a01b039081169163095ea7b391166200acef85600262011107565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af11580156200ad38573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200ad5e919062011048565b50602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b0936200ada2939082169288929091169060289060040162011069565b5f604051808303815f87803b1580156200adba575f80fd5b505af11580156200adcd573d5f803e3d5ffd5b5050604051630618f58760e51b81527f394836a4000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b1580156200ae3b575f80fd5b505af11580156200ae4e573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b094506200ae9393928316928892169060289060040162011069565b5f604051808303815f87803b1580156200aeab575f80fd5b505af11580156200aebe573d5f803e3d5ffd5b5050604051630618f58760e51b81527f394836a4000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b1580156200af2c575f80fd5b505af11580156200af3f573d5f803e3d5ffd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb4935062000bf792909116908590602890600401620111c7565b6023546021546040516370a0823160e01b81526001600160a01b039182166004820152620186a0925f9216906370a0823190602401602060405180830381865afa1580156200afe9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b00f919062011030565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa1580156200b05f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b085919062011030565b6027546023546020549293506001600160a01b0391821631929082169163095ea7b391166200b0b687600262011107565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af11580156200b0ff573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b125919062011048565b5060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063491cc7c2906084015f604051808303815f87803b1580156200b186575f80fd5b505af11580156200b199573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c926200b1ef928a929091169060289062010e71565b60405180910390a3602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b0936200b23a93908216928a929091169060289060040162011069565b5f604051808303815f87803b1580156200b252575f80fd5b505af11580156200b265573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b1580156200b2c9575f80fd5b505af11580156200b2dc573d5f803e3d5ffd5b50506026546025546023546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c926200b332928a929091169060289062010e71565b60405180910390a3602054602654602354604051630102614b60e41b81526001600160a01b039384169363102614b09362030d40936200b38193918316928b9291169060289060040162011069565b5f604051808303818588803b1580156200b399575f80fd5b505af11580156200b3ac573d5f803e3d5ffd5b50506023546021546040516370a0823160e01b81526001600160a01b0391821660048201525f9550911692506370a082319150602401602060405180830381865afa1580156200b3fe573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b424919062011030565b6023546025546040516370a0823160e01b81526001600160a01b0391821660048201529293505f929116906370a0823190602401602060405180830381865afa1580156200b474573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b49a919062011030565b6027549091506001600160a01b0316316200b4bb620065c388600262011107565b6200b4cc620071b588600262011107565b62001dee62000dd462030d408662010c76565b606060158054806020026020016040519081016040528092919081815260200182805480156200236157602002820191905f5260205f209081546001600160a01b0316815260019091019060200180831162002342575050505050905090565b6024805460205460405163095ea7b360e01b81526001600160a01b039182166004820152620186a093810184905291169063095ea7b3906044016020604051808303815f875af11580156200b596573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b5bc919062011048565b506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b1580156200b649575f80fd5b505af11580156200b65c573d5f803e3d5ffd5b50506026546025546024546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c926200b6b29287929091169060289062010e71565b60405180910390a3602054602654602454604051630102614b60e41b81526001600160a01b039384169363102614b0936200b6fd939082169287929091169060289060040162011069565b5f604051808303815f87803b1580156200b715575f80fd5b505af11580156200b728573d5f803e3d5ffd5b5050602480546025546040516370a0823160e01b81526001600160a01b0391821660048201525f9550911692506370a082319101602060405180830381865afa1580156200b778573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b79e919062011030565b905062003a2282602c5462000dd4919062011121565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790526023549054915163095ea7b360e01b81526001600160a01b03928316600482015260248101869052929350169063095ea7b3906044016020604051808303815f875af11580156200b852573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b878919062011048565b506025546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200b8e9575f80fd5b505af11580156200b8fc573d5f803e3d5ffd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a59042791506024015f604051808303815f87803b1580156200b960575f80fd5b505af11580156200b973573d5f803e3d5ffd5b50506023546040516001600160a01b039091166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507f1387a34900000000000000000000000000000000000000000000000000000000906044015b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e084901b90921682526200ba169160040162010c8c565b5f604051808303815f87803b1580156200ba2e575f80fd5b505af11580156200ba41573d5f803e3d5ffd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b78945062000bf793928316928892169087906028906004016201116f565b6026546040516001600160a01b039091166024820152620186a0905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f543f66a4000000000000000000000000000000000000000000000000000000008152600160048201529192506001600160a01b03169063543f66a4906024015f604051808303815f87803b1580156200bb47575f80fd5b505af11580156200bb5a573d5f803e3d5ffd5b50506023546020546001600160a01b03918216935063095ea7b39250166200bb8485600262011107565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303815f875af11580156200bbcd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200bbf3919062011048565b50602354604080516001600160a01b0390921660248084019190915281518084039091018152604490920181526020820180516001600160e01b03167f33e23dbb000000000000000000000000000000000000000000000000000000001790525163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916200bc89919060040162010c8c565b5f604051808303815f87803b1580156200bca1575f80fd5b505af11580156200bcb4573d5f803e3d5ffd5b5050602054602654602354604051630102614b60e41b81526001600160a01b03938416955063102614b094506200bcf993928316928892169060289060040162011069565b5f604051808303815f87803b1580156200bd11575f80fd5b505af11580156200bd24573d5f803e3d5ffd5b50506023546040516001600160a01b039091166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507f33e23dbb00000000000000000000000000000000000000000000000000000000906044016200b9cf565b6026546040516001600160a01b039091166024820152620186a09061c350905f9060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790526023549054915163095ea7b360e01b81526001600160a01b03928316600482015260248101879052929350169063095ea7b3906044016020604051808303815f875af11580156200be27573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200be4d919062011048565b506040515f602482015260448101839052737109709ecfa91a80626ff3989d68f67f5b1dd12d9063f28dceb3907f8afe4db4000000000000000000000000000000000000000000000000000000009060640162004996565b6020546040517f7c7442530000000000000000000000000000000000000000000000000000000081526509184e72a0006004820181905291620186a0916001600160a01b0390911690637c744253906024015f604051808303815f87803b1580156200bf0f575f80fd5b505af11580156200bf22573d5f803e3d5ffd5b50506027546025546020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b039485163196509284163194509083169263726ac97c9287926200bf8992169060289060040162010ea7565b5f604051808303818588803b1580156200bfa1575f80fd5b505af11580156200bfb4573d5f803e3d5ffd5b50506020546001600160a01b0316925063282906ed91506200bfd99050868662010c76565b6026546040516001600160e01b031960e085901b1681526200c00d916001600160a01b031690889060289060040162010eca565b5f604051808303818588803b1580156200c025575f80fd5b505af11580156200c038573d5f803e3d5ffd5b50506027546025546001600160a01b03918216319450163191506200c066905086620022bd87600262011107565b62000c2286620022f187600262011107565b6027546020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039182166084820152620186a092919091163190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a4015f604051808303815f87803b1580156200c112575f80fd5b505af11580156200c125573d5f803e3d5ffd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c906200c1759086905f9060289062010e71565b60405180910390a36020546026546040517f282906ed0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263282906ed92869262000d849290911690839060289060040162010eca565b6020546026546040517f282906ed000000000000000000000000000000000000000000000000000000008152620186a0926001600160a01b039081169263282906ed9285926200c22e921690839060289060040162010eca565b5f604051808303818588803b1580156200c246575f80fd5b505af11580156200c259573d5f803e3d5ffd5b50506040805162030d406024820152604480820187905282518083039091018152606490910182526020810180516001600160e01b03167fa458261b00000000000000000000000000000000000000000000000000000000179052905163f28dceb360e01b8152737109709ecfa91a80626ff3989d68f67f5b1dd12d945063f28dceb393506200aa0b925060040162010c8c565b6026546040516001600160a01b0390911660248201525f90819060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fef564bc9000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024016200ba16565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c54906044015f6040518083038186803b1580156200c3f7575f80fd5b505afa15801562000c22573d5f803e3d5ffd5b5f6200c41562010818565b6200c4228484836200c4f7565b9150505b92915050565b6040517fa59828850000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063a5982885906024015b5f6040518083038186803b1580156200c491575f80fd5b505afa15801562002a35573d5f803e3d5ffd5b6040517f0c9fd5810000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90630c9fd581906024016200c47a565b5f806200c50585846200c577565b90506200c56c6040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f787900000081525082866040516020016200c55692919062011334565b604051602081830303815290604052856200c584565b9150505b9392505050565b5f6200c57083836200c5b7565b60c0810151515f90156200c5ab576200c5a384848460c001516200c5d5565b90506200c570565b6200c5a384846200c789565b5f6200c5c483836200c87a565b6200c570838360200151846200c584565b5f806200c5e16200c887565b90505f6200c5f086836200c95b565b90505f6200c60882606001518360200151856200ce13565b90505f6200c619838389896200d038565b90505f6200c627826200df7a565b602081015181519192509060030b156200c69f578982604001516040516020016200c6549291906201136e565b60408051601f19818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526200c6969160040162010c8c565b60405180910390fd5b5f6200c6e36040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a2000000000000000000000008152508360016200e150565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d906200c73890849060040162010c8c565b602060405180830381865afa1580156200c754573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200c77a9190620113d3565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081525f908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc925906200c7df90879060040162010c8c565b5f60405180830381865afa1580156200c7fa573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200c8239190810190620114c6565b90505f6200c85482856040516020016200c83f929190620114fc565b6040516020818303038152906040526200e355565b90506001600160a01b0381166200c4225784846040516020016200c65492919062011514565b62003a2282825f6200e366565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c906200c910908490600401620115a8565b5f60405180830381865afa1580156200c92b573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200c9549190810190620115f0565b9250505090565b6200c98e6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b5f737109709ecfa91a80626ff3989d68f67f5b1dd12d90506200c9d96040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6200c9e4856200e473565b60208201525f6200c9f5866200e86c565b90505f836001600160a01b031663d930a0e66040518163ffffffff1660e01b81526004015f60405180830381865afa1580156200ca34573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200ca5d9190810190620115f0565b868385602001516040516020016200ca7994939291906201163a565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291505f906001600160a01b038616906360f9bb11906200cad290859060040162010c8c565b5f60405180830381865afa1580156200caed573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200cb169190810190620115f0565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f6906200cb6090849060040162011712565b602060405180830381865afa1580156200cb7c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200cba2919062011048565b6200cbba57816040516020016200c654919062011765565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906200cc01908490600401620117eb565b5f60405180830381865afa1580156200cc1c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200cc459190810190620115f0565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f6906200cc8e9084906004016201183e565b602060405180830381865afa1580156200ccaa573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200ccd0919062011048565b156200cd67576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906200cd1d9084906004016201183e565b5f60405180830381865afa1580156200cd38573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200cd619190810190620115f0565b60408501525b846001600160a01b03166349c4fac882865f01516040516020016200cd8d919062011891565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016200cdbb929190620118f1565b5f60405180830381865afa1580156200cdd6573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200cdff9190810190620115f0565b606085015250608083015250949350505050565b60408051600480825260a082019092526060915f9190816020015b60608152602001906001900390816200ce2e5790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250815f815181106200ce91576200ce9162011919565b60200260200101819052506040518060400160405280600381526020017f2d726c0000000000000000000000000000000000000000000000000000000000815250816001815181106200cee8576200cee862011919565b6020026020010181905250846040516020016200cf06919062011946565b604051602081830303815290604052816002815181106200cf2b576200cf2b62011919565b6020026020010181905250826040516020016200cf499190620119a6565b604051602081830303815290604052816003815181106200cf6e576200cf6e62011919565b60200260200101819052505f6200cf85826200df7a565b602080820151604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000818501908152825180840184525f80825290860152825180840190935290518252928101929092529192506200d016906040805180820182525f80825260209182015281518083019092528451825280850190820152906200eb01565b6200d02e57856040516020016200c6549190620119e0565b9695505050505050565b60a08101516040805180820182525f80825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d90156200d088565b511590565b6200d201578260200151156200d147576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a4016200c696565b8260c00151156200d201576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a4016200c696565b6040805160ff80825261200082019092525f91816020015b60608152602001906001900390816200d2195790505090505f6040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806200d2769062011a5f565b935060ff16815181106200d28e576200d28e62011919565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e37000000000000000000000000000000000000008152506040516020016200d2e1919062011a80565b6040516020818303038152906040528282806200d2fe9062011a5f565b935060ff16815181106200d316576200d31662011919565b60200260200101819052506040518060400160405280600681526020017f6465706c6f7900000000000000000000000000000000000000000000000000008152508282806200d3659062011a5f565b935060ff16815181106200d37d576200d37d62011919565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d650000000000000000000000000000000000008152508282806200d3cc9062011a5f565b935060ff16815181106200d3e4576200d3e462011919565b602002602001018190525087602001518282806200d4029062011a5f565b935060ff16815181106200d41a576200d41a62011919565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e7472616374506174680000000000000000000000000000000000008152508282806200d4699062011a5f565b935060ff16815181106200d481576200d48162011919565b6020908102919091010152875182826200d49b8162011a5f565b935060ff16815181106200d4b3576200d4b362011919565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e496400000000000000000000000000000000000000000000008152508282806200d5029062011a5f565b935060ff16815181106200d51a576200d51a62011919565b60200260200101819052506200d530466200eb67565b82826200d53d8162011a5f565b935060ff16815181106200d555576200d55562011919565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c6500000000000000000000000000000000008152508282806200d5a49062011a5f565b935060ff16815181106200d5bc576200d5bc62011919565b6020026020010181905250868282806200d5d69062011a5f565b935060ff16815181106200d5ee576200d5ee62011919565b60209081029190910101528551156200d7215760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826200d6428162011a5f565b935060ff16815181106200d65a576200d65a62011919565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d906200d6ac90899060040162010c8c565b5f60405180830381865afa1580156200d6c7573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200d6f09190810190620115f0565b82826200d6fd8162011a5f565b935060ff16815181106200d715576200d71562011919565b60200260200101819052505b8460200151156200d7fd5760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826200d76d8162011a5f565b935060ff16815181106200d785576200d78562011919565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806200d7d49062011a5f565b935060ff16815181106200d7ec576200d7ec62011919565b60200260200101819052506200d9df565b6200d8366200d0838660a001516040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6200d8d35760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826200d87c8162011a5f565b935060ff16815181106200d894576200d89462011919565b60200260200101819052508460a001516040516020016200d8b6919062011946565b6040516020818303038152906040528282806200d7d49062011a5f565b8460c001511580156200d9175750604080890151815180830183525f808252602091820152825180840190935281518352908101908201526200d91590511590565b155b156200d9df5760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826200d95e8162011a5f565b935060ff16815181106200d976576200d97662011919565b60200260200101819052506200d98c886200ec0b565b6040516020016200d99e919062011946565b6040516020818303038152906040528282806200d9bb9062011a5f565b935060ff16815181106200d9d3576200d9d362011919565b60200260200101819052505b604080860151815180830183525f808252602091820152825180840190935281518352908101908201526200da1390511590565b6200dab35760408051808201909152600b81527f2d2d72656c617965724964000000000000000000000000000000000000000000602082015282826200da598162011a5f565b935060ff16815181106200da71576200da7162011919565b602002602001018190525084604001518282806200da8f9062011a5f565b935060ff16815181106200daa7576200daa762011919565b60200260200101819052505b6060850151156200dbde5760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826200daff8162011a5f565b935060ff16815181106200db17576200db1762011919565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e8906024015f60405180830381865afa1580156200db84573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200dbad9190810190620115f0565b82826200dbba8162011a5f565b935060ff16815181106200dbd2576200dbd262011919565b60200260200101819052505b60e085015151156200dc915760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826200dc2b8162011a5f565b935060ff16815181106200dc43576200dc4362011919565b60200260200101819052506200dc608560e001515f01516200eb67565b82826200dc6d8162011a5f565b935060ff16815181106200dc85576200dc8562011919565b60200260200101819052505b60e085015160200151156200dd485760408051808201909152600a81527f2d2d676173507269636500000000000000000000000000000000000000000000602082015282826200dce18162011a5f565b935060ff16815181106200dcf9576200dcf962011919565b60200260200101819052506200dd178560e00151602001516200eb67565b82826200dd248162011a5f565b935060ff16815181106200dd3c576200dd3c62011919565b60200260200101819052505b60e085015160400151156200ddff5760408051808201909152600e81527f2d2d6d6178466565506572476173000000000000000000000000000000000000602082015282826200dd988162011a5f565b935060ff16815181106200ddb0576200ddb062011919565b60200260200101819052506200ddce8560e00151604001516200eb67565b82826200dddb8162011a5f565b935060ff16815181106200ddf3576200ddf362011919565b60200260200101819052505b60e085015160600151156200deb65760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826200de4f8162011a5f565b935060ff16815181106200de67576200de6762011919565b60200260200101819052506200de858560e00151606001516200eb67565b82826200de928162011a5f565b935060ff16815181106200deaa576200deaa62011919565b60200260200101819052505b5f8160ff1667ffffffffffffffff8111156200ded6576200ded662010ef3565b6040519080825280602002602001820160405280156200df0b57816020015b60608152602001906001900390816200def55790505b5090505f5b8260ff168160ff1610156200df6b57838160ff16815181106200df37576200df3762011919565b6020026020010151828260ff16815181106200df57576200df5762011919565b60209081029190910101526001016200df10565b5093505050505b949350505050565b6200dfa160405180606001604052805f60030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d925f91849163d145736c916200e0289186910162011ad9565b5f60405180830381865afa1580156200e043573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200e06c9190810190620115f0565b90505f6200e07b86836200f722565b90505f846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b81526004016200e0ac919062010b37565b5f604051808303815f875af11580156200e0c8573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200e0f1919081019062011b21565b805190915060030b158015906200e10b5750602081015151155b80156200e11b5750604081015151155b156200d02e57815f815181106200e136576200e13662011919565b60200260200101516040516020016200c654919062011bdc565b60605f6200e184856040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925286518252808701908201529091506200e1bc9082905b906200f88a565b156200e323575f6200e23d826200e236846200e22f6200e2028a6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925282518252918201519181019190915290565b906200f8b2565b906200f916565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200e2a29082906200f88a565b156200e30e57604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e30b905b82906200f9a2565b90505b6200e319816200f9c9565b925050506200c570565b82156200e33f5784846040516020016200c65492919062011dbb565b505060408051602081019091525f81526200c570565b5f808251602084015ff09392505050565b8160a00151156200e37657505050565b5f6200e3848484846200fa34565b90505f6200e392826200df7a565b602081015181519192509060030b1580156200e42f5750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e42f906040805180820182525f808252602091820152815180830190925284518252808501908201526200e1b5565b156200e43d57505050505050565b604082015151156200e4605781604001516040516020016200c654919062011e46565b806040516020016200c654919062011e9f565b60605f6200e4a7836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200e50d905b82906200eb01565b156200e58057604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200c570906200e57a90839062010028565b6200f9c9565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e5e3905b8290620100ba565b6001036200e6b457604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e64b906200e303565b50604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200c570906200e57a905b83906200f9a2565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e714906200e505565b156200e85357604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201819052845180860190955292518452830152906200e77e90839062010160565b90505f81600183516200e792919062011121565b815181106200e7a5576200e7a562011919565b602002602001015190506200e84a6200e57a6200e81d6040518060400160405280600581526020017f2e6a736f6e0000000000000000000000000000000000000000000000000000008152506040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925285518252808601908201529062010028565b95945050505050565b826040516020016200c654919062011ef8565b50919050565b60605f6200e8a0836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200e903906200e505565b156200e914576200c570816200f9c9565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200e974906200e5db565b6001036200e9e157604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200c570906200e57a906200e6ac565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200ea41906200e505565b156200e85357604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201819052845180860190955292518452830152906200eaab90839062010160565b90506001815111156200eaed5780600282516200eac9919062011121565b815181106200eadc576200eadc62011919565b602002602001015192505050919050565b50826040516020016200c654919062011ef8565b805182515f9111156200eb1657505f6200c426565b8151835160208501515f92916200eb2d9162010c76565b6200eb39919062011121565b9050826020015181036200eb525760019150506200c426565b82516020840151819020912014905092915050565b60605f6200eb75836201021c565b60010190505f8167ffffffffffffffff8111156200eb97576200eb9762010ef3565b6040519080825280601f01601f1916602001820160405280156200ebc2576020820181803683370190505b5090508181016020015b5f19017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846200ebcc57509392505050565b604081810151815180830183525f808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e53454400000000000000000000000000000000000000000000818401908152855180870187528381528401929092528451808601909552518452908301526060916200ec98905b829062010304565b156200ecd957505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200ed39906200ec90565b156200ed7a57505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200edda906200ec90565b156200ee1b57505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200ee7b906200ec90565b806200eee25750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200eee2906200ec90565b156200ef2357505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200ef83906200ec90565b806200efea5750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200efea906200ec90565b156200f02b57505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200f08b906200ec90565b806200f0f25750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200f0f2906200ec90565b156200f13357505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200f193906200ec90565b806200f1fa5750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200f1fa906200ec90565b156200f23b57505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200f29b906200ec90565b156200f2dc57505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200f33c906200ec90565b156200f37d57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200f3dd906200ec90565b156200f41e57505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200f47e906200ec90565b156200f4bf57505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200f51f906200ec90565b156200f56057505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200f5c0906200ec90565b806200f6275750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200f627906200ec90565b156200f66857505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200f6c8906200ec90565b156200f70957505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b604080840151845191516200c654929060200162011fca565b6060805f5b84518110156200f7b857818582815181106200f747576200f74762011919565b60200260200101516040516020016200f762929190620114fc565b6040516020818303038152906040529150600185516200f783919062011121565b81146200f7af57816040516020016200f79d91906201211c565b60405160208183030381529060405291505b6001016200f727565b50604080516003808252608082019092525f91816020015b60608152602001906001900390816200f7d057905050905083815f815181106200f7fe576200f7fe62011919565b60200260200101819052506040518060400160405280600281526020017f2d63000000000000000000000000000000000000000000000000000000000000815250816001815181106200f855576200f85562011919565b602002602001018190525081816002815181106200f877576200f87762011919565b6020908102919091010152949350505050565b60208083015183518351928401515f936200f8a9929184919062010319565b14159392505050565b604080518082019091525f80825260208201525f6200f8e2845f01518560200151855f015186602001516201044c565b90508360200151816200f8f6919062011121565b845185906200f90790839062011121565b90525060208401525090919050565b604080518082019091525f80825260208201528151835110156200f93c5750816200c426565b60208083015190840151600191146200f9645750815160208481015190840151829020919020145b80156200f99a578251845185906200f97e90839062011121565b90525082516020850180516200f99690839062010c76565b9052505b509192915050565b604080518082019091525f80825260208201526200f9c28383836201058c565b5092915050565b60605f825f015167ffffffffffffffff8111156200f9eb576200f9eb62010ef3565b6040519080825280601f01601f1916602001820160405280156200fa16576020820181803683370190505b5090505f6020820190506200f9c2818560200151865f015162010640565b60605f6200fa416200c887565b6040805160ff80825261200082019092529192505f9190816020015b60608152602001906001900390816200fa5d5790505090505f6040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806200faba9062011a5f565b935060ff16815181106200fad2576200fad262011919565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e33000000000000000000000000000000000000000000000000008152506040516020016200fb25919062012156565b6040516020818303038152906040528282806200fb429062011a5f565b935060ff16815181106200fb5a576200fb5a62011919565b60200260200101819052506040518060400160405280600881526020017f76616c69646174650000000000000000000000000000000000000000000000008152508282806200fba99062011a5f565b935060ff16815181106200fbc1576200fbc162011919565b6020026020010181905250826040516020016200fbdf9190620119a6565b6040516020818303038152906040528282806200fbfc9062011a5f565b935060ff16815181106200fc14576200fc1462011919565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e7472616374000000000000000000000000000000000000000000008152508282806200fc639062011a5f565b935060ff16815181106200fc7b576200fc7b62011919565b60200260200101819052506200fc928784620106c8565b82826200fc9f8162011a5f565b935060ff16815181106200fcb7576200fcb762011919565b6020908102919091010152855151156200fd6f5760408051808201909152600b81527f2d2d7265666572656e6365000000000000000000000000000000000000000000602082015282826200fd0c8162011a5f565b935060ff16815181106200fd24576200fd2462011919565b60200260200101819052506200fd3e865f015184620106c8565b82826200fd4b8162011a5f565b935060ff16815181106200fd63576200fd6362011919565b60200260200101819052505b8560800151156200fde45760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b0000000000000000602082015282826200fdbb8162011a5f565b935060ff16815181106200fdd3576200fdd362011919565b60200260200101819052506200fe50565b84156200fe505760408051808201909152601281527f2d2d726571756972655265666572656e63650000000000000000000000000000602082015282826200fe2c8162011a5f565b935060ff16815181106200fe44576200fe4462011919565b60200260200101819052505b604086015151156200fef75760408051808201909152600d81527f2d2d756e73616665416c6c6f7700000000000000000000000000000000000000602082015282826200fe9d8162011a5f565b935060ff16815181106200feb5576200feb562011919565b602002602001018190525085604001518282806200fed39062011a5f565b935060ff16815181106200feeb576200feeb62011919565b60200260200101819052505b8560600151156200ff675760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d6573000000000000000000000000602082015282826200ff438162011a5f565b935060ff16815181106200ff5b576200ff5b62011919565b60200260200101819052505b5f8160ff1667ffffffffffffffff8111156200ff87576200ff8762010ef3565b6040519080825280602002602001820160405280156200ffbc57816020015b60608152602001906001900390816200ffa65790505b5090505f5b8260ff168160ff1610156201001c57838160ff16815181106200ffe8576200ffe862011919565b6020026020010151828260ff168151811062010008576201000862011919565b60209081029190910101526001016200ffc1565b50979650505050505050565b604080518082019091525f80825260208201528151835110156201004e5750816200c426565b8151835160208501515f9291620100659162010c76565b62010071919062011121565b6020840151909150600190821462010093575082516020840151819020908220145b8015620100b157835185518690620100ad90839062011121565b9052505b50929392505050565b5f80825f0151620100dc855f01518660200151865f015187602001516201044c565b620100e8919062010c76565b90505b83516020850151620100fe919062010c76565b81116200f9c25781620101118162012189565b925050825f01516201014c8560200151836201012e919062011121565b86516201013c919062011121565b83865f015187602001516201044c565b62010158919062010c76565b9050620100eb565b60605f6201016f8484620100ba565b6201017c90600162010c76565b67ffffffffffffffff81111562010197576201019762010ef3565b604051908082528060200260200182016040528015620101cc57816020015b6060815260200190600190039081620101b65790505b5090505f5b81518110156201021457620101eb6200e57a86866200f9a2565b82828151811062010200576201020062011919565b6020908102919091010152600101620101d1565b509392505050565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831062010265577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831062010292576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310620102b157662386f26fc10000830492506010015b6305f5e1008310620102ca576305f5e100830492506008015b6127108310620102df57612710830492506004015b60648310620102f2576064830492506002015b600a83106200c4265760010192915050565b5f6201031183836201070b565b159392505050565b5f80858411620104425760208411620103e2575f84156201036d5760016201034386602062011121565b6201035090600862011107565b6201035d906002620122a0565b62010369919062011121565b1990505b83518116856201037e898962010c76565b6201038a919062011121565b805190935082165b818114620103ca57878411620103af57879450505050506200df72565b83620103bb81620122ad565b94505082845116905062010392565b620103d6878562010c76565b9450505050506200df72565b838320620103f1858862011121565b620103fd908762010c76565b91505b85821062010440578482208082036201042a576201041f868462010c76565b93505050506200df72565b6201043760018462011121565b92505062010400565b505b5092949350505050565b5f83818685116201057557602085116201051b575f8515620104a15760016201047787602062011121565b6201048490600862011107565b62010491906002620122a0565b6201049d919062011121565b1990505b845181165f87620104b38b8b62010c76565b620104bf919062011121565b855190915083165b8281146201050c57818610620104f157620104e38b8b62010c76565b96505050505050506200df72565b85620104fd8162012189565b965050838651169050620104c7565b8596505050505050506200df72565b508383205f905b6201052e868962011121565b821162010573578583208082036201054d57839450505050506200df72565b6201055a60018562010c76565b93505081806201056a9062012189565b92505062010522565b505b62010581878762010c76565b979650505050505050565b604080518082019091525f80825260208201525f620105bc855f01518660200151865f015187602001516201044c565b602080870180519186019190915251909150620105da908262011121565b835284516020860151620105ef919062010c76565b8103620105ff575f855262010637565b835183516201060f919062010c76565b855186906201062090839062011121565b905250835162010631908262010c76565b60208601525b50909392505050565b602081106201068057815183526201065a60208462010c76565b92506201066960208362010c76565b91506201067860208262011121565b905062010640565b5f198115620106b55760016201069883602062011121565b620106a690610100620122a0565b620106b2919062011121565b90505b9151835183169219169190911790915250565b60605f620106d784846200c95b565b8051602080830151604051939450620106f393909101620122c5565b60405160208183030381529060405291505092915050565b815181515f91908111156201071e575081515b602080850151908401515f5b83811015620107eb5782518251808214620107b4575f19602087101562010791576001846201075b89602062011121565b62010767919062010c76565b6201077490600862011107565b62010781906002620122a0565b6201078d919062011121565b1990505b8181168382168181039114620107b15797506200c4269650505050505050565b50505b620107c160208662010c76565b9450620107d060208562010c76565b93505050602081620107e3919062010c76565b90506201072a565b50845186516200d02e919062012304565b610c32806201232783390190565b61124c8062012f5983390190565b6040518060e001604052806060815260200160608152602001606081526020015f151581526020015f151581526020015f151581526020016201085a6201085f565b905290565b6040518061010001604052805f151581526020015f15158152602001606081526020015f801916815260200160608152602001606081526020015f151581526020016201085a60405180608001604052805f81526020015f81526020015f81526020015f81525090565b602080825282518282018190525f918401906040840190835b818110156201090b5783516001600160a01b0316835260209384019390920191600101620108e2565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101562010a43577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b8801810191908801905f5b8181101562010a28577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835262010a1184865162010916565b6020958601959094509290920191600101620109d4565b5091975050506020948501949290920191506001016201096a565b50929695505050505050565b5f8151808452602084019350602083015f5b8281101562010a8b5781516001600160e01b03191686526020958601959091019060010162010a61565b5093949350505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101562010a43577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875262010b02604088018262010916565b905060208201519150868103602088015262010b1f818362010a4f565b96505050602093840193919091019060010162010abb565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101562010a43577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845262010b9a85835162010916565b9450602093840193919091019060010162010b5d565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101562010a43577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015262010c32604087018262010a4f565b955050602093840193919091019060010162010bd6565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156200c426576200c42662010c49565b602081525f6200c570602083018462010916565b6001600160a01b0381511682526020810151151560208301526001600160a01b0360408201511660408301525f606082015160a0606085015262010ce860a085018262010916565b608093840151949093019390935250919050565b6001600160a01b0384168152606060208201525f62010d1f606083018562010916565b82810360408401526200d02e818562010ca0565b600181811c9082168062010d4857607f821691505b6020821081036200e866577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f81546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b0360018301541660408401526002820160a060608501525f815462010dca8162010d33565b8060a0880152600182165f811462010deb576001811462010e265762010e59565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660c089015260c082151560051b890101935062010e59565b845f5260205f205f5b8381101562010e505781548a820160c0015260019091019060200162010e2f565b890160c0019450505b50505060038401546080860152809250505092915050565b8381526001600160a01b0383166020820152608060408201525f608082015260a060608201525f6200e84a60a083018462010d80565b6001600160a01b0383168152604060208201525f6200df72604083018462010d80565b6001600160a01b0384168152826020820152606060408201525f6200e84a606083018462010d80565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b601f82111562000ddb57805f5260205f20601f840160051c8101602085101562010f475750805b601f840160051c820191505b8181101562002a35575f815560010162010f53565b815167ffffffffffffffff81111562010f855762010f8562010ef3565b62010f9d8162010f96845462010d33565b8462010f20565b6020601f82116001811462010fd2575f831562010fba5750848201515b5f19600385901b1c1916600184901b17845562002a35565b5f84815260208120601f198516915b8281101562011003578785015182556020948501946001909201910162010fe1565b50848210156201102157868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f6020828403121562011041575f80fd5b5051919050565b5f6020828403121562011059575f80fd5b815180151581146200c570575f80fd5b6001600160a01b03851681528360208201526001600160a01b0383166040820152608060608201525f6200d02e608083018462010d80565b8481526001600160a01b0384166020820152608060408201525f620110ca608083018562010916565b828103606084015262010581818562010d80565b6001600160a01b0385168152836020820152608060408201525f620110ca608083018562010916565b80820281158282048414176200c426576200c42662010c49565b818103818111156200c426576200c42662010c49565b6001600160a01b03851681528360208201526001600160a01b0383166040820152608060608201525f6200d02e608083018462010ca0565b6001600160a01b03861681528460208201526001600160a01b038416604082015260a060608201525f620111a760a083018562010916565b8281036080840152620111bb818562010d80565b98975050505050505050565b6001600160a01b0384168152606060208201525f620111ea606083018562010916565b82810360408401526200d02e818562010d80565b6001600160a01b03861681528460208201526001600160a01b038416604082015260a060608201525f6201123660a083018562010916565b8281036080840152620111bb818562010ca0565b5f826201127e577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b6001600160a01b0385168152836020820152608060408201525f620112ac608083018562010916565b828103606084015262010581818562010ca0565b604081525f620112d4604083018562010916565b82810360208401526200c56c818562010d80565b6001600160a01b0384168152826020820152606060408201525f6200e84a606083018462010ca0565b6001600160a01b0383168152604060208201525f6200df72604083018462010ca0565b6001600160a01b0383168152604060208201525f6200df72604083018462010916565b5f81518060208401855e5f93019283525090919050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f620113a1601a83018562011357565b7f3a2000000000000000000000000000000000000000000000000000000000000081526200c56c600282018562011357565b5f60208284031215620113e4575f80fd5b81516001600160a01b03811681146200c570575f80fd5b6040516060810167ffffffffffffffff8111828210171562011421576201142162010ef3565b60405290565b5f8067ffffffffffffffff84111562011444576201144462010ef3565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171562011476576201147662010ef3565b6040528381529050808284018510156201148e575f80fd5b8383602083015e5f60208583010152509392505050565b5f82601f830112620114b5575f80fd5b6200c5708383516020850162011427565b5f60208284031215620114d7575f80fd5b815167ffffffffffffffff811115620114ee575f80fd5b6200c42284828501620114a5565b5f6200df726201150d838662011357565b8462011357565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f62011547601a83018562011357565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000815262011579601982018562011357565b7f2200000000000000000000000000000000000000000000000000000000000000815260010195945050505050565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201525f6200c570608083018462010916565b5f6020828403121562011601575f80fd5b815167ffffffffffffffff81111562011618575f80fd5b8201601f8101841362011629575f80fd5b6200c4228482516020840162011427565b5f62011647828762011357565b7f2f00000000000000000000000000000000000000000000000000000000000000815262011679600182018762011357565b90507f2f000000000000000000000000000000000000000000000000000000000000008152620116ad600182018662011357565b90507f2f000000000000000000000000000000000000000000000000000000000000008152620116e1600182018562011357565b7f2e6a736f6e0000000000000000000000000000000000000000000000000000008152600501979650505050505050565b604081525f62011726604083018462010916565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081525f62011798601f83018462011357565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f81527f6d6c00000000000000000000000000000000000000000000000000000000000060208201526022019392505050565b604081525f620117ff604083018462010916565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b604081525f62011852604083018462010916565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081525f620118c4601483018462011357565b7f275d2e6b656363616b32353600000000000000000000000000000000000000008152600c019392505050565b604081525f62011905604083018562010916565b82810360208401526200c56c818562010916565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f220000000000000000000000000000000000000000000000000000000000000081525f62011979600183018462011357565b7f220000000000000000000000000000000000000000000000000000000000000081526001019392505050565b5f620119b3828462011357565b7f2f6275696c642d696e666f0000000000000000000000000000000000000000008152600b019392505050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201525f6200c570604b83018462011357565b5f60ff821660ff810362011a775762011a7762010c49565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201525f6200c570602983018462011357565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201525f6200c570608083018462010916565b5f6020828403121562011b32575f80fd5b815167ffffffffffffffff81111562011b49575f80fd5b82016060818503121562011b5b575f80fd5b62011b65620113fb565b81518060030b811462011b76575f80fd5b8152602082015167ffffffffffffffff81111562011b92575f80fd5b62011ba086828501620114a5565b602083015250604082015167ffffffffffffffff81111562011bc0575f80fd5b62011bce86828501620114a5565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201525f62011c35602183018462011357565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657481527f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960208201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560408201527f616c69666965642070617468206f66207468652062617368206578656375746160608201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960808201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60a08201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960c08201527f6f75722070726f6a65637420287573696e6720666f727761726420736c61736860e08201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101008201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061012082015261013b019392505050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201525f62011e14602183018562011357565b7f2720696e206f75747075743a200000000000000000000000000000000000000081526200c56c600d82018562011357565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201525f6200c570602983018462011357565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201525f6200c570602283018462011357565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081525f62011f2b600e83018462011357565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e7472616381527f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e7360208201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e747260408201527f6163742e6a736f6e00000000000000000000000000000000000000000000000060608201526068019392505050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081525f62011ffd601883018562011357565b7f20696e200000000000000000000000000000000000000000000000000000000081526201202f600482018562011357565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f727465642081527f6c6963656e736520666f7220626c6f636b206578706c6f72657220766572696660208201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f60408201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c60608201527f206f7220736574207468652060736b69704c6963656e73655479706560206f7060808201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060a082015260b70195945050505050565b5f62012129828462011357565b7f200000000000000000000000000000000000000000000000000000000000000081526001019392505050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081525f6200c570601c83018462011357565b5f5f1982036201219d576201219d62010c49565b5060010190565b6001815b6001841115620121e557808504811115620121c757620121c762010c49565b6001841615620121d657908102905b60019390931c928002620121a8565b935093915050565b5f82620121fd575060016200c426565b816201220b57505f6200c426565b81600181146201222457600281146201222f576201224f565b60019150506200c426565b60ff84111562012243576201224362010c49565b50506001821b6200c426565b5060208310610133831016604e8410600b841016171562012274575081810a6200c426565b620122825f198484620121a4565b805f190482111562012298576201229862010c49565b029392505050565b5f6200c5708383620121ed565b5f81620122be57620122be62010c49565b505f190190565b5f620122d2828562011357565b7f3a0000000000000000000000000000000000000000000000000000000000000081526200c56c600182018562011357565b8181035f8312801583831316838312821617156200f9c2576200f9c262010c4956fe608060405234801561000f575f80fd5b50604051610c32380380610c3283398101604081905261002e916100f0565b8181600361003c83826101d9565b50600461004982826101d9565b5050505050610293565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610076575f80fd5b81516001600160401b0381111561008f5761008f610053565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100bd576100bd610053565b6040528181528382016020018510156100d4575f80fd5b8160208501602083015e5f918101602001919091529392505050565b5f8060408385031215610101575f80fd5b82516001600160401b03811115610116575f80fd5b61012285828601610067565b602085015190935090506001600160401b0381111561013f575f80fd5b61014b85828601610067565b9150509250929050565b600181811c9082168061016957607f821691505b60208210810361018757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101d457805f5260205f20601f840160051c810160208510156101b25750805b601f840160051c820191505b818110156101d1575f81556001016101be565b50505b505050565b81516001600160401b038111156101f2576101f2610053565b610206816102008454610155565b8461018d565b6020601f821160018114610238575f83156102215750848201515b5f19600385901b1c1916600184901b1784556101d1565b5f84815260208120601f198516915b828110156102675787850151825560209485019460019092019101610247565b508482101561028457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610992806102a05f395ff3fe608060405234801561000f575f80fd5b50600436106100b9575f3560e01c806340c10f191161007257806395d89b411161005857806395d89b411461017c578063a9059cbb14610184578063dd62ed3e14610197575f80fd5b806340c10f191461013257806370a0823114610147575f80fd5b806318160ddd116100a257806318160ddd146100fe57806323b872dd14610110578063313ce56714610123575f80fd5b806306fdde03146100bd578063095ea7b3146100db575b5f80fd5b6100c56101dc565b6040516100d291906107a5565b60405180910390f35b6100ee6100e9366004610820565b61026c565b60405190151581526020016100d2565b6002545b6040519081526020016100d2565b6100ee61011e366004610848565b610285565b604051601281526020016100d2565b610145610140366004610820565b6102a8565b005b610102610155366004610882565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b6100c56102b6565b6100ee610192366004610820565b6102c5565b6101026101a53660046108a2565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b6060600380546101eb906108d3565b80601f0160208091040260200160405190810160405280929190818152602001828054610217906108d3565b80156102625780601f1061023957610100808354040283529160200191610262565b820191905f5260205f20905b81548152906001019060200180831161024557829003601f168201915b5050505050905090565b5f336102798185856102d2565b60019150505b92915050565b5f336102928582856102e4565b61029d8585856103b6565b506001949350505050565b6102b2828261045f565b5050565b6060600480546101eb906108d3565b5f336102798185856103b6565b6102df83838360016104b9565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103b057818110156103a2576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103b084848484035f6104b9565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610405576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b73ffffffffffffffffffffffffffffffffffffffff8216610454576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b6102df8383836105fe565b73ffffffffffffffffffffffffffffffffffffffff82166104ae576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b6102b25f83836105fe565b73ffffffffffffffffffffffffffffffffffffffff8416610508576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b73ffffffffffffffffffffffffffffffffffffffff8316610557576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610399565b73ffffffffffffffffffffffffffffffffffffffff8085165f90815260016020908152604080832093871683529290522082905580156103b0578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516105f091815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610635578060025f82825461062a9190610924565b909155506106e59050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260208190526040902054818110156106ba576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610399565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661070e57600280548290039055610739565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161079891815260200190565b60405180910390a3505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461081b575f80fd5b919050565b5f8060408385031215610831575f80fd5b61083a836107f8565b946020939093013593505050565b5f805f6060848603121561085a575f80fd5b610863846107f8565b9250610871602085016107f8565b929592945050506040919091013590565b5f60208284031215610892575f80fd5b61089b826107f8565b9392505050565b5f80604083850312156108b3575f80fd5b6108bc836107f8565b91506108ca602084016107f8565b90509250929050565b600181811c908216806108e757607f821691505b60208210810361091e577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b8082018082111561027f577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea2646970667358221220569a538c722848c143f241fcbfd3f113f81549aa32aa89f9a6169cd25cb5399e64736f6c634300081a0033608060405234801561000f575f80fd5b5060405161124c38038061124c83398101604081905261002e9161010e565b604051806040016040528060048152602001635a65746160e01b815250604051806040016040528060048152602001635a45544160e01b815250816003908161007791906101d7565b50600461008482826101d7565b5050506001600160a01b03821615806100a457506001600160a01b038116155b156100c25760405163e6c4247b60e01b815260040160405180910390fd5b600680546001600160a01b039384166001600160a01b03199182161790915560078054929093169116179055610291565b80516001600160a01b0381168114610109575f80fd5b919050565b5f806040838503121561011f575f80fd5b610128836100f3565b9150610136602084016100f3565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061016757607f821691505b60208210810361018557634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101d257805f5260205f20601f840160051c810160208510156101b05750805b601f840160051c820191505b818110156101cf575f81556001016101bc565b50505b505050565b81516001600160401b038111156101f0576101f061013f565b610204816101fe8454610153565b8461018b565b6020601f821160018114610236575f831561021f5750848201515b5f19600385901b1c1916600184901b1784556101cf565b5f84815260208120601f198516915b828110156102655787850151825560209485019460019092019101610245565b508482101561028257868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610fae8061029e5f395ff3fe608060405234801561000f575f80fd5b5060043610610115575f3560e01c806342966c68116100ad57806379cc67901161007d578063a9059cbb11610063578063a9059cbb14610286578063bff9662a14610299578063dd62ed3e146102b9575f80fd5b806379cc67901461026b57806395d89b411461027e575f80fd5b806342966c68146101fb5780635b1125911461020e57806370a082311461022e578063779e3b6314610263575f80fd5b80631e458bee116100e85780631e458bee1461018157806323b872dd14610194578063313ce567146101a7578063328a01d0146101b6575f80fd5b806306fdde0314610119578063095ea7b31461013757806315d57fd41461015a57806318160ddd1461016f575b5f80fd5b6101216102fe565b60405161012e9190610d7a565b60405180910390f35b61014a610145366004610df5565b61038e565b604051901515815260200161012e565b61016d610168366004610e1d565b6103a7565b005b6002545b60405190815260200161012e565b61016d61018f366004610e4e565b610572565b61014a6101a2366004610e7e565b610625565b6040516012815260200161012e565b6007546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161012e565b61016d610209366004610eb8565b610648565b6006546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b61017361023c366004610ecf565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b61016d610655565b61016d610279366004610df5565b610779565b61012161082a565b61014a610294366004610df5565b610839565b6005546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b6101736102c7366004610e1d565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b60606003805461030d90610eef565b80601f016020809104026020016040519081016040528092919081815260200182805461033990610eef565b80156103845780601f1061035b57610100808354040283529160200191610384565b820191905f5260205f20905b81548152906001019060200180831161036757829003601f168201915b5050505050905090565b5f3361039b818585610846565b60019150505b92915050565b60075473ffffffffffffffffffffffffffffffffffffffff1633148015906103e7575060065473ffffffffffffffffffffffffffffffffffffffff163314155b15610425576040517fcdfcef970000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216158061045c575073ffffffffffffffffffffffffffffffffffffffff8116155b15610493576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316811790935560058054918516919092161790556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff910160405180910390a16040805133815273ffffffffffffffffffffffffffffffffffffffff831660208201527f1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c910160405180910390a15050565b60055473ffffffffffffffffffffffffffffffffffffffff1633146105c5576040517f3fe32fba00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b6105cf8383610858565b808373ffffffffffffffffffffffffffffffffffffffff167fc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb8460405161061891815260200190565b60405180910390a3505050565b5f336106328582856108b6565b61063d858585610983565b506001949350505050565b6106523382610a2c565b50565b60075473ffffffffffffffffffffffffffffffffffffffff1633146106a8576040517fe700765e00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b60065473ffffffffffffffffffffffffffffffffffffffff166106f7576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600654600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0910160405180910390a1565b60055473ffffffffffffffffffffffffffffffffffffffff1633146107cc576040517f3fe32fba00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b6107d68282610a86565b8173ffffffffffffffffffffffffffffffffffffffff167f919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b18260405161081e91815260200190565b60405180910390a25050565b60606004805461030d90610eef565b5f3361039b818585610983565b6108538383836001610a9b565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166108a7576040517fec442f050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b6108b25f8383610be0565b5050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461097d578181101561096f576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602481018290526044810183905260640161041c565b61097d84848484035f610a9b565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166109d2576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8216610a21576040517fec442f050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b610853838383610be0565b73ffffffffffffffffffffffffffffffffffffffff8216610a7b576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b6108b2825f83610be0565b610a918233836108b6565b6108b28282610a2c565b73ffffffffffffffffffffffffffffffffffffffff8416610aea576040517fe602df050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8316610b39576040517f94280d620000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8085165f908152600160209081526040808320938716835292905220829055801561097d578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bd291815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c17578060025f828254610c0c9190610f40565b90915550610cc79050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020819052604090205481811015610c9c576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602481018290526044810183905260640161041c565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610cf057600280548290039055610d1b565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161061891815260200190565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610df0575f80fd5b919050565b5f8060408385031215610e06575f80fd5b610e0f83610dcd565b946020939093013593505050565b5f8060408385031215610e2e575f80fd5b610e3783610dcd565b9150610e4560208401610dcd565b90509250929050565b5f805f60608486031215610e60575f80fd5b610e6984610dcd565b95602085013595506040909401359392505050565b5f805f60608486031215610e90575f80fd5b610e9984610dcd565b9250610ea760208501610dcd565b929592945050506040919091013590565b5f60208284031215610ec8575f80fd5b5035919050565b5f60208284031215610edf575f80fd5b610ee882610dcd565b9392505050565b600181811c90821680610f0357607f821691505b602082108103610f3a577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b808201808211156103a1577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea26469706673582212206a41e2cd6fbf39f12e2917f13d39a938ebbfbe1bbc40009f596c89c4dc977dca64736f6c634300081a0033a2646970667358221220cd529ddc800cc53637d5f1328e2a2ab13db4ea27ed6fcc2f37d28137c116f38464736f6c634300081a0033",
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

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestCallFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testCallFailsIfRevertGasLimitExceeded")
}

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestCallFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestCallFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x01a74aee.
//
// Solidity: function testCallFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestCallFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestCallFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
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

// TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x458085f5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded")
}

// TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x458085f5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x458085f5.
//
// Solidity: function testDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallERC20ToCustodyFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
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

// TestDepositAndCallEthFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc0fab86d.
//
// Solidity: function testDepositAndCallEthFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallEthFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallEthFailsIfRevertGasLimitExceeded")
}

// TestDepositAndCallEthFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc0fab86d.
//
// Solidity: function testDepositAndCallEthFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallEthFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xc0fab86d.
//
// Solidity: function testDepositAndCallEthFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallEthFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
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

// TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x478095e5.
//
// Solidity: function testDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded")
}

// TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x478095e5.
//
// Solidity: function testDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x478095e5.
//
// Solidity: function testDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositAndCallEthWithAmountFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
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

// TestDepositERC20ToCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x86b6e429.
//
// Solidity: function testDepositERC20ToCustodyFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfAmountIs0")
}

// TestDepositERC20ToCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x86b6e429.
//
// Solidity: function testDepositERC20ToCustodyFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x86b6e429.
//
// Solidity: function testDepositERC20ToCustodyFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
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

// TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x846b9f7f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfReceiverIsZeroAddress")
}

// TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x846b9f7f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x846b9f7f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x1f4f542f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded")
}

// TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x1f4f542f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x1f4f542f.
//
// Solidity: function testDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
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

// TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xf75fc969.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyWithPayloadFailsIfAmountIs0")
}

// TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xf75fc969.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0xf75fc969.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
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

// TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x895c2bc6.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress")
}

// TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x895c2bc6.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x895c2bc6.
//
// Solidity: function testDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositERC20ToCustodyWithPayloadFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
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

// TestDepositEthToTssFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x7bb46d46.
//
// Solidity: function testDepositEthToTssFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssFailsIfAmountIs0")
}

// TestDepositEthToTssFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x7bb46d46.
//
// Solidity: function testDepositEthToTssFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x7bb46d46.
//
// Solidity: function testDepositEthToTssFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x8aeeb7e7.
//
// Solidity: function testDepositEthToTssFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssFailsIfReceiverIsZeroAddress")
}

// TestDepositEthToTssFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x8aeeb7e7.
//
// Solidity: function testDepositEthToTssFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x8aeeb7e7.
//
// Solidity: function testDepositEthToTssFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xb1526934.
//
// Solidity: function testDepositEthToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssFailsIfRevertGasLimitExceeded")
}

// TestDepositEthToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xb1526934.
//
// Solidity: function testDepositEthToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0xb1526934.
//
// Solidity: function testDepositEthToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
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

// TestDepositEthToTssWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x84c59bcf.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssWithPayloadFailsIfAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssWithPayloadFailsIfAmountIs0")
}

// TestDepositEthToTssWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x84c59bcf.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssWithPayloadFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayloadFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayloadFailsIfAmountIs0 is a paid mutator transaction binding the contract method 0x84c59bcf.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfAmountIs0() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssWithPayloadFailsIfAmountIs0() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayloadFailsIfAmountIs0(&_GatewayEVMInboundTest.TransactOpts)
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

// TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x53c9a0a3.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress")
}

// TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x53c9a0a3.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x53c9a0a3.
//
// Solidity: function testDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthToTssWithPayloadFailsIfReceiverIsZeroAddress(&_GatewayEVMInboundTest.TransactOpts)
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

// TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x88343a41.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactor) TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMInboundTest.contract.Transact(opts, "testDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded")
}

// TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x88343a41.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestSession) TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
}

// TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded is a paid mutator transaction binding the contract method 0x88343a41.
//
// Solidity: function testDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded() returns()
func (_GatewayEVMInboundTest *GatewayEVMInboundTestTransactorSession) TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded() (*types.Transaction, error) {
	return _GatewayEVMInboundTest.Contract.TestDepositEthWithAmountToTssFailsIfRevertGasLimitExceeded(&_GatewayEVMInboundTest.TransactOpts)
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
