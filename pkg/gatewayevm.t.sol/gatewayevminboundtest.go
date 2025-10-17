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
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testCallWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfCallOnRevertIsTrue\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfDestinationIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustody\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTss\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssWithPayload\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDepositZetaToConnector\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositERC20ToCustodyIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositERC20ToCustodyIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositERC20ToCustodyWithPayloadIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssIfPayloadSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssWithPayloadIfAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositedAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallOnRevertNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZETANotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055620f4240602c55348015603357600080fd5b5061be41806100436000396000f3fe608060405234801561001057600080fd5b506004361061025c5760003560e01c806385226c8111610145578063b2849063116100bd578063cdef966f1161008c578063e306a97811610071578063e306a978146103c5578063e85c5a07146103cd578063fa7626d4146103d557600080fd5b8063cdef966f146103b5578063e20c9f71146103bd57600080fd5b8063b284906314610385578063b5508aa91461038d578063ba414fa614610395578063ba46ba23146103ad57600080fd5b80639d96310a11610114578063aa030c1c116100f9578063aa030c1c1461036d578063b0464fdc14610375578063b1409f711461037d57600080fd5b80639d96310a1461035d5780639fd1e5971461036557600080fd5b806385226c81146103235780638f5a422414610338578063916a17c61461034057806395cd04451461035557600080fd5b80633e5e3c23116101d85780634ce25c0a116101a75780635c4013d01161018c5780635c4013d0146102fe5780636459542a1461030657806366d9a9a01461030e57600080fd5b80634ce25c0a146102ee57806351da903d146102f657600080fd5b80633e5e3c23146102ce5780633f7286f4146102d657806341b8fec6146102de578063466f332e146102e657600080fd5b806318e141381161022f5780632a5dcf36116102145780632a5dcf36146102a95780632ade3880146102b157806330f7c04f146102c657600080fd5b806318e14138146102835780631ed7831c1461028b57600080fd5b80630724d8e31461026157806309b21ccb1461026b5780630a9254e4146102735780630fc37f5e1461027b575b600080fd5b6102696103e2565b005b610269610596565b61026961071f565b61026961126a565b610269611452565b6102936115a8565b6040516102a09190618541565b60405180910390f35b61026961160a565b6102b961178f565b6040516102a091906185dd565b6102696118d1565b610293611d3d565b610293611d9d565b610269611dfd565b610269611ef3565b610269612115565b61026961236a565b6102696124b2565b6102696125a1565b61031661299a565b6040516102a09190618743565b61032b612b1c565b6040516102a091906187e1565b610269612bec565b610348612cc4565b6040516102a09190618858565b610269612dbf565b610269613074565b6102696131c1565b6102696133e3565b6103486135a0565b61026961369b565b610269613935565b61032b613b2a565b61039d613bfa565b60405190151581526020016102a0565b610269613cce565b610269613dd9565b610293613ec8565b610269613f28565b6102696140f5565b601f5461039d9060ff1681565b6027546020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039182166084820152620186a092919091163190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561047e57600080fd5b505af1158015610492573d6000803e3d6000fd5b50506026546025546040516001600160a01b039283169450911691507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c906104e1908690600090602890618a2b565b60405180910390a36020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263726ac97c92869261053c9290911690602890600401618a61565b6000604051808303818588803b15801561055557600080fd5b505af1158015610569573d6000803e3d6000fd5b50506027546001600160a01b0316319250610591915061058b90508484618ab2565b8261435d565b505050565b6026546040516001600160a01b03909116602482015260019060009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae7600000000000000000000000000000000000000000000000000000000017905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561067a57600080fd5b505af115801561068e573d6000803e3d6000fd5b50506020546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063d09e3b7893506106e99260009288929116908790602890600401618ac5565b600060405180830381600087803b15801561070357600080fd5b505af1158015610717573d6000803e3d6000fd5b505050505050565b602580547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602680548216611234179055602780549091166156781790556040516107719061846e565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f0801580156107f6573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316179055602754604051911690819061083f9061847b565b6001600160a01b03928316815291166020820152604001604051809103906000f080158015610872573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081178255604080518082018252600e81527f4761746577617945564d2e736f6c00000000000000000000000000000000000060208201526027546025549251908616948101949094526044840192909252909216606482015261096491906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b000000000000000000000000000000000000000000000000000000001790526143dc565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff000000000000000000000000000000000000000090911681178255604080518082018252601081527f4552433230437573746f64792e736f6c000000000000000000000000000000009381019390935260275460255491516024810193909352841660448301529092166064830152610a3491608401610907565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602180549190920483167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055604080518082018252601a81527f5a657461436f6e6e6563746f724e6f6e4e61746976652e736f6c00000000000060208083019190915254602480546027546025549551938716928401929092528516604483015284166064820152919092166084820152610b6d919060a40160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff8c8765e000000000000000000000000000000000000000000000000000000001790526143dc565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602280549190920483167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790556027546040517fca669fa700000000000000000000000000000000000000000000000000000000815291166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015610c4557600080fd5b505af1158015610c59573d6000803e3d6000fd5b5050602480546027546022546040517f15d57fd40000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152908216938101939093521692506315d57fd49150604401600060405180830381600087803b158015610ccc57600080fd5b505af1158015610ce0573d6000803e3d6000fd5b50506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b158015610d6457600080fd5b505af1158015610d78573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015610dee57600080fd5b505af1158015610e02573d6000803e3d6000fd5b50506020546021546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201529116925063ae7a3a6f9150602401600060405180830381600087803b158015610e6857600080fd5b505af1158015610e7c573d6000803e3d6000fd5b50506020546022546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911692506310188aef9150602401600060405180830381600087803b158015610ee257600080fd5b505af1158015610ef6573d6000803e3d6000fd5b50506021546023546040517f9b19251a0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639b19251a9150602401600060405180830381600087803b158015610f5c57600080fd5b505af1158015610f70573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610fd257600080fd5b505af1158015610fe6573d6000803e3d6000fd5b5050602354602554602c546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810191909152911692506340c10f199150604401600060405180830381600087803b15801561105757600080fd5b505af115801561106b573d6000803e3d6000fd5b50506027546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156110e157600080fd5b505af11580156110f5573d6000803e3d6000fd5b5050602254602554602c546040517f106e62900000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201526024810191909152600060448201529116925063106e62909150606401600060405180830381600087803b15801561116d57600080fd5b505af1158015611181573d6000803e3d6000fd5b50506040805160a0810182526103218082526000602080840182815284860193845285519182019095528181526060840181905260808401919091528251602880549551151574010000000000000000000000000000000000000000027fffffffffffffffffffffff0000000000000000000000000000000000000000009096166001600160a01b0392831617959095178555915160298054919093167fffffffffffffffffffffffff00000000000000000000000000000000000000009190911617909155909350909150602a9061125a9082618b90565b5060808201518160030155905050565b60208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169263a2ba1934926004808401938290030181865afa1580156112c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112ec9190618c4f565b6112f7906001618ab2565b67ffffffffffffffff81111561130f5761130f618b1a565b6040519080825280601f01601f191660200182016040528015611339576020820181803683370190505b50602a906113479082618b90565b50604051630618f58760e51b81527f386691aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156113b457600080fd5b505af11580156113c8573d6000803e3d6000fd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c935060019261141e921690602890600401618a61565b6000604051808303818588803b15801561143757600080fd5b505af115801561144b573d6000803e3d6000fd5b5050505050565b6026546040516001600160a01b039091166024820152600090819060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae7600000000000000000000000000000000000000000000000000000000017905251630618f58760e51b81527f951e19ed000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b600060405180830381600087803b15801561153657600080fd5b505af115801561154a573d6000803e3d6000fd5b50506020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063d09e3b7894506106e99392831692889216908790602890600401618ac5565b6060601680548060200260200160405190810160405280929190818152602001828054801561160057602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116115e2575b5050505050905090565b6026546040516001600160a01b039091166024820152600090819060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae7600000000000000000000000000000000000000000000000000000000017905251630618f58760e51b81527f7671265e000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b600060405180830381600087803b1580156116ee57600080fd5b505af1158015611702573d6000803e3d6000fd5b50506020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063744b9b8b935086926117599216908690602890600401618c68565b6000604051808303818588803b15801561177257600080fd5b505af1158015611786573d6000803e3d6000fd5b50505050505050565b6060601e805480602002602001604051908101604052809291908181526020016000905b828210156118c857600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b828210156118b1578382906000526020600020018054611824906188ef565b80601f0160208091040260200160405190810160405280929190818152602001828054611850906188ef565b801561189d5780601f106118725761010080835404028352916020019161189d565b820191906000526020600020905b81548152906001019060200180831161188057829003601f168201915b505050505081526020019060010190611805565b5050505081525050815260200190600101906117b3565b50505050905090565b6023546021546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a09260009216906370a0823190602401602060405180830381865afa15801561193d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119619190618c4f565b905061196e60008261435d565b6026546040516001600160a01b03909116602482015260009060440160408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae76000000000000000000000000000000000000000000000000000000000179052602354905491517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101879052929350169063095ea7b3906044016020604051808303816000875af1158015611a51573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a759190618c9c565b506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611b0457600080fd5b505af1158015611b18573d6000803e3d6000fd5b50506026546025546023546040516001600160a01b03938416955091831693507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f92611b6e928992909116908790602890618cbe565b60405180910390a36020546026546023546040517fd09e3b780000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363d09e3b7893611bd293908216928992909116908790602890600401618ac5565b600060405180830381600087803b158015611bec57600080fd5b505af1158015611c00573d6000803e3d6000fd5b50506023546021546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015611c6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c909190618c4f565b9050611c9c848261435d565b6023546025546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015611d06573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d2a9190618c4f565b905061144b85602c5461058b9190618cf8565b60606018805480602002602001604051908101604052809291908181526020018280548015611600576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116115e2575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015611600576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116115e2575050505050905090565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015611e6c57600080fd5b505af1158015611e80573d6000803e3d6000fd5b50506020546023546040517f102614b00000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063102614b09350611ed9926000928792911690602890600401618d0b565b600060405180830381600087803b15801561143757600080fd5b60208054604080517fa2ba19340000000000000000000000000000000000000000000000000000000081529051620186a0936000936002936001600160a01b039091169263a2ba1934926004808401939192918290030181865afa158015611f5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f839190618c4f565b611f8d9190618d42565b67ffffffffffffffff811115611fa557611fa5618b1a565b6040519080825280601f01601f191660200182016040528015611fcf576020820181803683370190505b5060208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa158015612034573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120589190618c4f565b6120629190618d42565b61206d906001618ab2565b67ffffffffffffffff81111561208557612085618b1a565b6040519080825280601f01601f1916602001820160405280156120af576020820181803683370190505b50602a906120bd9082618b90565b50604051630618f58760e51b81527f386691aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024016116d4565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a06024820181905292919091169063095ea7b3906044016020604051808303816000875af115801561218a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121ae9190618c9c565b5060208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169263a2ba1934926004808401938290030181865afa15801561220d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122319190618c4f565b61223c906001618ab2565b67ffffffffffffffff81111561225457612254618b1a565b6040519080825280601f01601f19166020018201604052801561227e576020820181803683370190505b50602a9061228c9082618b90565b50604051630618f58760e51b81527f386691aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b600060405180830381600087803b1580156122fa57600080fd5b505af115801561230e573d6000803e3d6000fd5b50506020546026546023546040517f102614b00000000000000000000000000000000000000000000000000000000081526001600160a01b03938416955063102614b09450611ed9939283169287921690602890600401618d0b565b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae7600000000000000000000000000000000000000000000000000000000017905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561244b57600080fd5b505af115801561245f573d6000803e3d6000fd5b50506020546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250631becceb49150611ed9906000908590602890600401618c68565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561252157600080fd5b505af1158015612535573d6000803e3d6000fd5b50506020546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063726ac97c9150839061258890600090602890600401618a61565b6000604051808303818588803b15801561070357600080fd5b6023546021546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a09260009216906370a0823190602401602060405180830381865afa15801561260d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126319190618c4f565b905061263e60008261435d565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526024810185905291169063095ea7b3906044016020604051808303816000875af11580156126ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126d09190618c9c565b506020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561275f57600080fd5b505af1158015612773573d6000803e3d6000fd5b50506026546025546023546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c926127c792889290911690602890618a2b565b60405180910390a36020546026546023546040517f102614b00000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363102614b0936128299390821692889290911690602890600401618d0b565b600060405180830381600087803b15801561284357600080fd5b505af1158015612857573d6000803e3d6000fd5b50506023546021546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156128c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128e79190618c4f565b90506128f3838261435d565b6023546025546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa15801561295d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129819190618c4f565b905061299484602c5461058b9190618cf8565b50505050565b6060601b805480602002602001604051908101604052809291908181526020016000905b828210156118c857838290600052602060002090600202016040518060400160405290816000820180546129f1906188ef565b80601f0160208091040260200160405190810160405280929190818152602001828054612a1d906188ef565b8015612a6a5780601f10612a3f57610100808354040283529160200191612a6a565b820191906000526020600020905b815481529060010190602001808311612a4d57829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015612b0457602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411612ab15790505b505050505081525050815260200190600101906129be565b6060601a805480602002602001604051908101604052809291908181526020016000905b828210156118c8578382906000526020600020018054612b5f906188ef565b80601f0160208091040260200160405190810160405280929190818152602001828054612b8b906188ef565b8015612bd85780601f10612bad57610100808354040283529160200191612bd8565b820191906000526020600020905b815481529060010190602001808311612bbb57829003601f168201915b505050505081526020019060010190612b40565b604051630618f58760e51b81527f7671265e000000000000000000000000000000000000000000000000000000006004820152600090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015612c5b57600080fd5b505af1158015612c6f573d6000803e3d6000fd5b50506020546026546040517f726ac97c0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316945063726ac97c93508592612588921690602890600401618a61565b6060601d805480602002602001604051908101604052809291908181526020016000905b828210156118c85760008481526020908190206040805180820182526002860290920180546001600160a01b03168352600181018054835181870281018701909452808452939491938583019392830182828015612da757602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411612d545790505b50505050508152505081526020019060010190612ce8565b60208054604080517fa2ba19340000000000000000000000000000000000000000000000000000000081529051620186a0936000936002936001600160a01b039091169263a2ba1934926004808401939192918290030181865afa158015612e2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e4f9190618c4f565b612e599190618d42565b67ffffffffffffffff811115612e7157612e71618b1a565b6040519080825280601f01601f191660200182016040528015612e9b576020820181803683370190505b5060208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa158015612f00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f249190618c4f565b612f2e9190618d42565b612f39906001618ab2565b67ffffffffffffffff811115612f5157612f51618b1a565b6040519080825280601f01601f191660200182016040528015612f7b576020820181803683370190505b50602a90612f899082618b90565b506023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526024810185905291169063095ea7b3906044016020604051808303816000875af1158015612ff8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061301c9190618c9c565b50604051630618f58760e51b81527f386691aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e09060240161151c565b6026546040516001600160a01b03909116602482015260019060009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae7600000000000000000000000000000000000000000000000000000000017905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561315857600080fd5b505af115801561316c573d6000803e3d6000fd5b50506020546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063744b9b8b91508490611759906000908690602890600401618c68565b6027546026546040516001600160a01b039182166024820152620186a09291909116319060009060440160408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae760000000000000000000000000000000000000000000000000000000001790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156132d257600080fd5b505af11580156132e6573d6000803e3d6000fd5b50506026546025546040516001600160a01b039283169450911691507fa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f906133379087906000908790602890618cbe565b60405180910390a36020546026546040517f744b9b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263744b9b8b92879261339492909116908690602890600401618c68565b6000604051808303818588803b1580156133ad57600080fd5b505af11580156133c1573d6000803e3d6000fd5b50506027546001600160a01b0316319250612994915061058b90508585618ab2565b6026546040516001600160a01b03909116602482015260009060440160408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae760000000000000000000000000000000000000000000000000000000001790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156134e657600080fd5b505af11580156134fa573d6000803e3d6000fd5b50506026546025546040516001600160a01b039283169450911691507fd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d97490613546908590602890618d7d565b60405180910390a36020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692631becceb492611ed9929116908590602890600401618c68565b6060601c805480602002602001604051908101604052809291908181526020016000905b828210156118c85760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561368357602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116136305790505b505050505081525050815260200190600101906135c4565b60208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290516000936002936001600160a01b03169263a2ba193492600480830193928290030181865afa1580156136fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137229190618c4f565b61372c9190618d42565b67ffffffffffffffff81111561374457613744618b1a565b6040519080825280601f01601f19166020018201604052801561376e576020820181803683370190505b5060208054604080517fa2ba193400000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b039092169263a2ba1934926004808401938290030181865afa1580156137d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137f79190618c4f565b6138019190618d42565b61380c906001618ab2565b67ffffffffffffffff81111561382457613824618b1a565b6040519080825280601f01601f19166020018201604052801561384e576020820181803683370190505b50602a9061385c9082618b90565b50604051630618f58760e51b81527f386691aa000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b600060405180830381600087803b1580156138ca57600080fd5b505af11580156138de573d6000803e3d6000fd5b50506020546026546040517f1becceb40000000000000000000000000000000000000000000000000000000081526001600160a01b039283169450631becceb49350611ed992909116908590602890600401618c68565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a06024820181905292919091169063095ea7b3906044016020604051808303816000875af11580156139aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139ce9190618c9c565b506025546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613a4157600080fd5b505af1158015613a55573d6000803e3d6000fd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a5904279150602401600060405180830381600087803b158015613abb57600080fd5b505af1158015613acf573d6000803e3d6000fd5b5050604051630618f58760e51b81527fac2175f1000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024016122e0565b60606019805480602002602001604051908101604052809291908181526020016000905b828210156118c8578382906000526020600020018054613b6d906188ef565b80601f0160208091040260200160405190810160405280929190818152602001828054613b99906188ef565b8015613be65780601f10613bbb57610100808354040283529160200191613be6565b820191906000526020600020905b815481529060010190602001808311613bc957829003601f168201915b505050505081526020019060010190613b4e565b60085460009060ff1615613c12575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015613ca3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613cc79190618c4f565b1415905090565b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae76000000000000000000000000000000000000000000000000000000000179052602880547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000017905551630618f58760e51b81527f19b4bff2000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024016138b0565b6023546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260006024820181905292919091169063095ea7b3906044016020604051808303816000875af1158015613e4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e709190618c9c565b50604051630618f58760e51b81527f951e19ed000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024016122e0565b60606015805480602002602001604051908101604052809291908181526020018280548015611600576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116115e2575050505050905090565b602480546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620186a093810184905291169063095ea7b3906044016020604051808303816000875af1158015613f9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613fbe9190618c9c565b50604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561402b57600080fd5b505af115801561403f573d6000803e3d6000fd5b50506026546025546024546040516001600160a01b03938416955091831693507fc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c9261409392879290911690602890618a2b565b60405180910390a36020546026546024546040517f102614b00000000000000000000000000000000000000000000000000000000081526001600160a01b039384169363102614b093611ed99390821692879290911690602890600401618d0b565b6026546040516001600160a01b039091166024820152620186a09060009060440160408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f84fae76000000000000000000000000000000000000000000000000000000000179052602354905491517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482015260248101869052929350169063095ea7b3906044016020604051808303816000875af11580156141dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906142019190618c9c565b506025546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561427457600080fd5b505af1158015614288573d6000803e3d6000fd5b50506021546023546040517f9a5904270000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015291169250639a5904279150602401600060405180830381600087803b1580156142ee57600080fd5b505af1158015614302573d6000803e3d6000fd5b5050604051630618f58760e51b81527fac2175f1000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e0915060240161151c565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c549060440160006040518083038186803b1580156143c857600080fd5b505afa158015610717573d6000803e3d6000fd5b60006143e6618488565b6143f18484836143fb565b9150505b92915050565b6000806144088584614476565b905061446b6040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f78790000008152508286604051602001614456929190618da2565b60405160208183030381529060405285614482565b9150505b9392505050565b600061446f83836144b0565b60c081015151600090156144a65761449f84848460c001516144cb565b905061446f565b61449f8484614671565b60006144bc838361475c565b61446f83836020015184614482565b6000806144d661476c565b905060006144e4868361483f565b905060006144fb8260600151836020015185614ce5565b9050600061450b83838989614ef7565b9050600061451882615d74565b602081015181519192509060030b1561458b57898260400151604051602001614542929190618dc4565b60408051601f19818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261458291600401618e45565b60405180910390fd5b60006145ce6040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a200000000000000000000000815250836001615f43565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d90614621908490600401618e45565b602060405180830381865afa15801561463e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906146629190618e58565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc925906146c6908790600401618e45565b600060405180830381865afa1580156146e3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261470b9190810190618f3a565b905060006147398285604051602001614725929190618f6f565b604051602081830303815290604052616143565b90506001600160a01b0381166143f1578484604051602001614542929190618f9e565b61476882826000616156565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c906147f3908490600401619049565b600060405180830381865afa158015614810573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526148389190810190619090565b9250505090565b6148716040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d90506148bc6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6148c585616259565b602082015260006148d58661663e565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa158015614917573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261493f9190810190619090565b8683856020015160405160200161495994939291906190d9565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb11906149b1908590600401618e45565b600060405180830381865afa1580156149ce573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526149f69190810190619090565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f690614a3e9084906004016191dd565b602060405180830381865afa158015614a5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614a7f9190618c9c565b614a945781604051602001614542919061922f565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890614ad99084906004016192c1565b600060405180830381865afa158015614af6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052614b1e9190810190619090565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f690614b65908490600401619313565b602060405180830381865afa158015614b82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ba69190618c9c565b15614c3b576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890614bf0908490600401619313565b600060405180830381865afa158015614c0d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052614c359190810190619090565b60408501525b846001600160a01b03166349c4fac8828660000151604051602001614c609190619365565b6040516020818303038152906040526040518363ffffffff1660e01b8152600401614c8c9291906193d1565b600060405180830381865afa158015614ca9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052614cd19190810190619090565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b6060815260200190600190039081614d015790505090506040518060400160405280600481526020017f677265700000000000000000000000000000000000000000000000000000000081525081600081518110614d6157614d616193f6565b60200260200101819052506040518060400160405280600381526020017f2d726c000000000000000000000000000000000000000000000000000000000081525081600181518110614db557614db56193f6565b602002602001018190525084604051602001614dd19190619425565b60405160208183030381529060405281600281518110614df357614df36193f6565b602002602001018190525082604051602001614e0f9190619491565b60405160208183030381529060405281600381518110614e3157614e316193f6565b60200260200101819052506000614e4782615d74565b602080820151604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000008185019081528251808401845260008082529086015282518084019093529051825292810192909252919250614ed890604080518082018252600080825260209182015281518083019092528451825280850190820152906168c1565b614eed578560405160200161454291906194d2565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d9015614f47565b511590565b6150bb57826020015115615003576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401614582565b8260c00151156150bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401614582565b6040805160ff8082526120008201909252600091816020015b60608152602001906001900390816150d457905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061512f90619563565b935060ff1681518110615144576151446193f6565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e37000000000000000000000000000000000000008152506040516020016151959190619582565b6040516020818303038152906040528282806151b090619563565b935060ff16815181106151c5576151c56193f6565b60200260200101819052506040518060400160405280600681526020017f6465706c6f79000000000000000000000000000000000000000000000000000081525082828061521290619563565b935060ff1681518110615227576152276193f6565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d6500000000000000000000000000000000000081525082828061527490619563565b935060ff1681518110615289576152896193f6565b602002602001018190525087602001518282806152a590619563565b935060ff16815181106152ba576152ba6193f6565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163745061746800000000000000000000000000000000000081525082828061530790619563565b935060ff168151811061531c5761531c6193f6565b60209081029190910101528751828261533481619563565b935060ff1681518110615349576153496193f6565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e4964000000000000000000000000000000000000000000000081525082828061539690619563565b935060ff16815181106153ab576153ab6193f6565b60200260200101819052506153bf46616922565b82826153ca81619563565b935060ff16815181106153df576153df6193f6565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c65000000000000000000000000000000000081525082828061542c90619563565b935060ff1681518110615441576154416193f6565b60200260200101819052508682828061545990619563565b935060ff168151811061546e5761546e6193f6565b60209081029190910101528551156155955760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826154bf81619563565b935060ff16815181106154d4576154d46193f6565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d90615524908990600401618e45565b600060405180830381865afa158015615541573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526155699190810190619090565b828261557481619563565b935060ff1681518110615589576155896193f6565b60200260200101819052505b8460200151156156655760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826155de81619563565b935060ff16815181106155f3576155f36193f6565b60200260200101819052506040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525082828061564090619563565b935060ff1681518110615655576156556193f6565b602002602001018190525061582c565b61569d614f428660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6157305760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826156e081619563565b935060ff16815181106156f5576156f56193f6565b60200260200101819052508460a001516040516020016157159190619425565b60405160208183030381529060405282828061564090619563565b8460c0015115801561577357506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261577190511590565b155b1561582c5760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826157b781619563565b935060ff16815181106157cc576157cc6193f6565b60200260200101819052506157e0886169c2565b6040516020016157f09190619425565b60405160208183030381529060405282828061580b90619563565b935060ff1681518110615820576158206193f6565b60200260200101819052505b6040808601518151808301835260008082526020918201528251808401909352815183529081019082015261586090511590565b6158f55760408051808201909152600b81527f2d2d72656c617965724964000000000000000000000000000000000000000000602082015282826158a381619563565b935060ff16815181106158b8576158b86193f6565b602002602001018190525084604001518282806158d490619563565b935060ff16815181106158e9576158e96193f6565b60200260200101819052505b606085015115615a165760408051808201909152600681527f2d2d73616c7400000000000000000000000000000000000000000000000000006020820152828261593e81619563565b935060ff1681518110615953576159536193f6565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa1580156159c2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526159ea9190810190619090565b82826159f581619563565b935060ff1681518110615a0a57615a0a6193f6565b60200260200101819052505b60e08501515115615abd5760408051808201909152600a81527f2d2d6761734c696d69740000000000000000000000000000000000000000000060208201528282615a6081619563565b935060ff1681518110615a7557615a756193f6565b6020026020010181905250615a918560e0015160000151616922565b8282615a9c81619563565b935060ff1681518110615ab157615ab16193f6565b60200260200101819052505b60e08501516020015115615b675760408051808201909152600a81527f2d2d67617350726963650000000000000000000000000000000000000000000060208201528282615b0a81619563565b935060ff1681518110615b1f57615b1f6193f6565b6020026020010181905250615b3b8560e0015160200151616922565b8282615b4681619563565b935060ff1681518110615b5b57615b5b6193f6565b60200260200101819052505b60e08501516040015115615c115760408051808201909152600e81527f2d2d6d617846656550657247617300000000000000000000000000000000000060208201528282615bb481619563565b935060ff1681518110615bc957615bc96193f6565b6020026020010181905250615be58560e0015160400151616922565b8282615bf081619563565b935060ff1681518110615c0557615c056193f6565b60200260200101819052505b60e08501516060015115615cbb5760408051808201909152601681527f2d2d6d61785072696f726974794665655065724761730000000000000000000060208201528282615c5e81619563565b935060ff1681518110615c7357615c736193f6565b6020026020010181905250615c8f8560e0015160600151616922565b8282615c9a81619563565b935060ff1681518110615caf57615caf6193f6565b60200260200101819052505b60008160ff1667ffffffffffffffff811115615cd957615cd9618b1a565b604051908082528060200260200182016040528015615d0c57816020015b6060815260200190600190039081615cf75790505b50905060005b8260ff168160ff161015615d6557838160ff1681518110615d3557615d356193f6565b6020026020010151828260ff1681518110615d5257615d526193f6565b6020908102919091010152600101615d12565b5093505050505b949350505050565b615d9b6040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c91615e21918691016195ed565b600060405180830381865afa158015615e3e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615e669190810190619090565b90506000615e7486836174b1565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b8152600401615ea491906187e1565b6000604051808303816000875af1158015615ec3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052615eeb9190810190619634565b805190915060030b15801590615f045750602081015151155b8015615f135750604081015151155b15614eed5781600081518110615f2b57615f2b6193f6565b602002602001015160405160200161454291906196ea565b60606000615f788560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600080825260209182015281518083019092528651825280870190820152909150615faf9082905b90617606565b1561610c57600061602c8261602684616020615ff28a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b9061762d565b9061768f565b604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150616090908290617606565b156160fa57604080518082018252600181527f0a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526160f7905b8290617714565b90505b6161038161773a565b9250505061446f565b82156161255784846040516020016145429291906198d6565b505060408051602081019091526000815261446f565b509392505050565b6000808251602084016000f09392505050565b8160a001511561616557505050565b60006161728484846177a3565b9050600061617f82615d74565b602081015181519192509060030b15801561621b5750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261621b90604080518082018252600080825260209182015281518083019092528451825280850190820152615fa9565b1561622857505050505050565b60408201515115616248578160400151604051602001614542919061997d565b8060405160200161454291906199db565b6060600061628e8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506162f3905b82906168c1565b1561636257604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261446f9061635d908390617d3e565b61773a565b604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526163c4905b8290617dc8565b60010361649157604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261642a906160f0565b50604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261446f9061635d905b8390617714565b604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526164f0906162ec565b1561662757604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290616558908390617e62565b90506000816001835161656b9190618cf8565b8151811061657b5761657b6193f6565b6020026020010151905061661e61635d6165f16040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b60408051808201825260008082526020918201528151808301909252855182528086019082015290617d3e565b95945050505050565b826040516020016145429190619a46565b50919050565b606060006166738360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c00000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506166d5906162ec565b156166e35761446f8161773a565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616742906163bd565b6001036167ac57604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261446f9061635d9061648a565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261680b906162ec565b1561662757604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290616873908390617e62565b90506001815111156168af57806002825161688e9190618cf8565b8151811061689e5761689e6193f6565b602002602001015192505050919050565b50826040516020016145429190619a46565b8051825160009111156168d6575060006143f5565b815183516020850151600092916168ec91618ab2565b6168f69190618cf8565b90508260200151810361690d5760019150506143f5565b82516020840151819020912014905092915050565b6060600061692f83617f07565b600101905060008167ffffffffffffffff81111561694f5761694f618b1a565b6040519080825280601f01601f191660200182016040528015616979576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461698357509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e5345440000000000000000000000000000000000000000000081840190815285518087018752838152840192909252845180860190955251845290830152606091616a4e905b8290617fe9565b15616a8e57505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e7365000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616aed90616a47565b15616b2d57505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d4954000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616b8c90616a47565b15616bcc57505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c79000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616c2b90616a47565b80616c905750604080518082018252601081527f47504c2d322e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616c9090616a47565b15616cd057505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c79000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616d2f90616a47565b80616d945750604080518082018252601081527f47504c2d332e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616d9490616a47565b15616dd457505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616e3390616a47565b80616e985750604080518082018252601181527f4c47504c2d322e312d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616e9890616a47565b15616ed857505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616f3790616a47565b80616f9c5750604080518082018252601181527f4c47504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616f9c90616a47565b15616fdc57505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261703b90616a47565b1561707b57505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c617573650000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526170da90616a47565b1561711a57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261717990616a47565b156171b957505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261721890616a47565b1561725857505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e3000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526172b790616a47565b156172f757505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261735690616a47565b806173bb5750604080518082018252601181527f4147504c2d332e302d6f722d6c61746572000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526173bb90616a47565b156173fb57505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261745a90616a47565b1561749a57505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b604080840151845191516145429290602001619b24565b60608060005b845181101561753c57818582815181106174d3576174d36193f6565b60200260200101516040516020016174ec929190618f6f565b60405160208183030381529060405291506001855161750b9190618cf8565b811461753457816040516020016175229190619c8d565b60405160208183030381529060405291505b6001016174b7565b5060408051600380825260808201909252600091816020015b60608152602001906001900390816175555790505090508381600081518110617580576175806193f6565b60200260200101819052506040518060400160405280600281526020017f2d63000000000000000000000000000000000000000000000000000000000000815250816001815181106175d4576175d46193f6565b602002602001018190525081816002815181106175f3576175f36193f6565b6020908102919091010152949350505050565b60208083015183518351928401516000936176249291849190617ffd565b14159392505050565b6040805180820190915260008082526020820152600061765f846000015185602001518560000151866020015161810e565b90508360200151816176719190618cf8565b84518590617680908390618cf8565b90525060208401525090919050565b60408051808201909152600080825260208201528151835110156176b45750816143f5565b60208083015190840151600191146176db5750815160208481015190840151829020919020145b801561770c578251845185906176f2908390618cf8565b9052508251602085018051617708908390618ab2565b9052505b509192915050565b604080518082019091526000808252602082015261773383838361822e565b5092915050565b60606000826000015167ffffffffffffffff81111561775b5761775b618b1a565b6040519080825280601f01601f191660200182016040528015617785576020820181803683370190505b509050600060208201905061773381856020015186600001516182d9565b606060006177af61476c565b6040805160ff808252612000820190925291925060009190816020015b60608152602001906001900390816177cc57905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061782790619563565b935060ff168151811061783c5761783c6193f6565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161788d9190619cce565b6040516020818303038152906040528282806178a890619563565b935060ff16815181106178bd576178bd6193f6565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061790a90619563565b935060ff168151811061791f5761791f6193f6565b60200260200101819052508260405160200161793b9190619491565b60405160208183030381529060405282828061795690619563565b935060ff168151811061796b5761796b6193f6565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e7472616374000000000000000000000000000000000000000000008152508282806179b890619563565b935060ff16815181106179cd576179cd6193f6565b60200260200101819052506179e28784618353565b82826179ed81619563565b935060ff1681518110617a0257617a026193f6565b602090810291909101015285515115617aae5760408051808201909152600b81527f2d2d7265666572656e636500000000000000000000000000000000000000000060208201528282617a5481619563565b935060ff1681518110617a6957617a696193f6565b6020026020010181905250617a82866000015184618353565b8282617a8d81619563565b935060ff1681518110617aa257617aa26193f6565b60200260200101819052505b856080015115617b1c5760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b000000000000000060208201528282617af781619563565b935060ff1681518110617b0c57617b0c6193f6565b6020026020010181905250617b82565b8415617b825760408051808201909152601281527f2d2d726571756972655265666572656e6365000000000000000000000000000060208201528282617b6181619563565b935060ff1681518110617b7657617b766193f6565b60200260200101819052505b60408601515115617c1e5760408051808201909152600d81527f2d2d756e73616665416c6c6f770000000000000000000000000000000000000060208201528282617bcc81619563565b935060ff1681518110617be157617be16193f6565b60200260200101819052508560400151828280617bfd90619563565b935060ff1681518110617c1257617c126193f6565b60200260200101819052505b856060015115617c885760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d657300000000000000000000000060208201528282617c6781619563565b935060ff1681518110617c7c57617c7c6193f6565b60200260200101819052505b60008160ff1667ffffffffffffffff811115617ca657617ca6618b1a565b604051908082528060200260200182016040528015617cd957816020015b6060815260200190600190039081617cc45790505b50905060005b8260ff168160ff161015617d3257838160ff1681518110617d0257617d026193f6565b6020026020010151828260ff1681518110617d1f57617d1f6193f6565b6020908102919091010152600101617cdf565b50979650505050505050565b6040805180820190915260008082526020820152815183511015617d635750816143f5565b81518351602085015160009291617d7991618ab2565b617d839190618cf8565b60208401519091506001908214617da4575082516020840151819020908220145b8015617dbf57835185518690617dbb908390618cf8565b9052505b50929392505050565b6000808260000151617dec856000015186602001518660000151876020015161810e565b617df69190618ab2565b90505b83516020850151617e0a9190618ab2565b81116177335781617e1a81619d13565b9250508260000151617e51856020015183617e359190618cf8565b8651617e419190618cf8565b838660000151876020015161810e565b617e5b9190618ab2565b9050617df9565b60606000617e708484617dc8565b617e7b906001618ab2565b67ffffffffffffffff811115617e9357617e93618b1a565b604051908082528060200260200182016040528015617ec657816020015b6060815260200190600190039081617eb15790505b50905060005b815181101561613b57617ee261635d8686617714565b828281518110617ef457617ef46193f6565b6020908102919091010152600101617ecc565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310617f50577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310617f7c576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310617f9a57662386f26fc10000830492506010015b6305f5e1008310617fb2576305f5e100830492506008015b6127108310617fc657612710830492506004015b60648310617fd8576064830492506002015b600a83106143f55760010192915050565b6000617ff58383618393565b159392505050565b60008085841161810457602084116180b05760008415618048576001618024866020618cf8565b61802f906008619d2d565b61803a906002619e2b565b6180449190618cf8565b1990505b83518116856180578989618ab2565b6180619190618cf8565b805190935082165b81811461809b578784116180835787945050505050615d6c565b8361808d81619e37565b945050828451169050618069565b6180a58785618ab2565b945050505050615d6c565b8383206180bd8588618cf8565b6180c79087618ab2565b91505b858210618102578482208082036180ef576180e58684618ab2565b9350505050615d6c565b6180fa600184618cf8565b9250506180ca565b505b5092949350505050565b6000838186851161821957602085116181c8576000851561815a576001618136876020618cf8565b618141906008619d2d565b61814c906002619e2b565b6181569190618cf8565b1990505b8451811660008761816b8b8b618ab2565b6181759190618cf8565b855190915083165b8281146181ba578186106181a2576181958b8b618ab2565b9650505050505050615d6c565b856181ac81619d13565b96505083865116905061817d565b859650505050505050615d6c565b508383206000905b6181da8689618cf8565b8211618217578583208082036181f65783945050505050615d6c565b618201600185618ab2565b935050818061820f90619d13565b9250506181d0565b505b6182238787618ab2565b979650505050505050565b60408051808201909152600080825260208201526000618260856000015186602001518660000151876020015161810e565b60208087018051918601919091525190915061827c9082618cf8565b83528451602086015161828f9190618ab2565b810361829e57600085526182d0565b835183516182ac9190618ab2565b855186906182bb908390618cf8565b90525083516182ca9082618ab2565b60208601525b50909392505050565b6020811061831157815183526182f0602084618ab2565b92506182fd602083618ab2565b915061830a602082618cf8565b90506182d9565b6000198115618340576001618327836020618cf8565b61833390610100619e2b565b61833d9190618cf8565b90505b9151835183169219169190911790915250565b60606000618361848461483f565b805160208083015160405193945061837b93909101619e4e565b60405160208183030381529060405291505092915050565b81518151600091908111156183a6575081515b6020808501519084015160005b8381101561845f578251825180821461842f57600019602087101561840e576001846183e0896020618cf8565b6183ea9190618ab2565b6183f5906008619d2d565b618400906002619e2b565b61840a9190618cf8565b1990505b818116838216818103911461842c5797506143f59650505050505050565b50505b61843a602086618ab2565b9450618447602085618ab2565b935050506020816184589190618ab2565b90506183b3565b5084518651614eed9190619ea6565b610c9f80619ec783390190565b6112a68061ab6683390190565b6040518060e001604052806060815260200160608152602001606081526020016000151581526020016000151581526020016000151581526020016184cb6184d0565b905290565b604051806101000160405280600015158152602001600015158152602001606081526020016000801916815260200160608152602001606081526020016000151581526020016184cb6040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b818110156185825783516001600160a01b031683526020938401939092019160010161855b565b509095945050505050565b60005b838110156185a8578181015183820152602001618590565b50506000910152565b600081518084526185c981602086016020860161858d565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156186d9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b818110156186bf577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a85030183526186a98486516185b1565b602095860195909450929092019160010161866f565b509197505050602094850194929092019150600101618605565b50929695505050505050565b600081518084526020840193506020830160005b828110156187395781517fffffffff00000000000000000000000000000000000000000000000000000000168652602095860195909101906001016186f9565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156186d9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051604087526187af60408801826185b1565b90506020820151915086810360208801526187ca81836186e5565b96505050602093840193919091019060010161876b565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156186d9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184526188438583516185b1565b94506020938401939190910190600101618809565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156186d9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b03815116865260208101519050604060208701526188d960408701826186e5565b9550506020938401939190910190600101618880565b600181811c9082168061890357607f821691505b602082108103616638577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600081546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b0360018301541660408401526002820160a0606085015260008154618986816188ef565b8060a088015260018216600081146189a557600181146189df57618a13565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660c089015260c082151560051b8901019350618a13565b84600052602060002060005b83811015618a0a5781548a820160c001526001909101906020016189eb565b890160c0019450505b50505060038401546080860152809250505092915050565b8381526001600160a01b0383166020820152608060408201526000608082015260a06060820152600061661e60a083018461893c565b6001600160a01b0383168152604060208201526000615d6c604083018461893c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156143f5576143f5618a83565b6001600160a01b03861681528460208201526001600160a01b038416604082015260a060608201526000618afc60a08301856185b1565b8281036080840152618b0e818561893c565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82111561059157806000526020600020601f840160051c81016020851015618b705750805b601f840160051c820191505b8181101561144b5760008155600101618b7c565b815167ffffffffffffffff811115618baa57618baa618b1a565b618bbe81618bb884546188ef565b84618b49565b6020601f821160018114618bf25760008315618bda5750848201515b600019600385901b1c1916600184901b17845561144b565b600084815260208120601f198516915b82811015618c225787850151825560209485019460019092019101618c02565b5084821015618c405786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b600060208284031215618c6157600080fd5b5051919050565b6001600160a01b0384168152606060208201526000618c8a60608301856185b1565b8281036040840152614eed818561893c565b600060208284031215618cae57600080fd5b8151801515811461446f57600080fd5b8481526001600160a01b0384166020820152608060408201526000618ce660808301856185b1565b8281036060840152618223818561893c565b818103818111156143f5576143f5618a83565b6001600160a01b03851681528360208201526001600160a01b0383166040820152608060608201526000614eed608083018461893c565b600082618d78577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b604081526000618d9060408301856185b1565b828103602084015261446b818561893c565b6001600160a01b0383168152604060208201526000615d6c60408301846185b1565b7f4661696c656420746f206465706c6f7920636f6e747261637420000000000000815260008351618dfc81601a85016020880161858d565b7f3a20000000000000000000000000000000000000000000000000000000000000601a918401918201528351618e3981601c84016020880161858d565b01601c01949350505050565b60208152600061446f60208301846185b1565b600060208284031215618e6a57600080fd5b81516001600160a01b038116811461446f57600080fd5b6040516060810167ffffffffffffffff81118282101715618ea457618ea4618b1a565b60405290565b60008067ffffffffffffffff841115618ec557618ec5618b1a565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff82111715618ef457618ef4618b1a565b604052838152905080828401851015618f0c57600080fd5b61613b84602083018561858d565b600082601f830112618f2b57600080fd5b61446f83835160208501618eaa565b600060208284031215618f4c57600080fd5b815167ffffffffffffffff811115618f6357600080fd5b6143f184828501618f1a565b60008351618f8181846020880161858d565b835190830190618f9581836020880161858d565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e747261637420000000000000815260008351618fd681601a85016020880161858d565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161901381603384016020880161858d565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f5554000000000000000000000000000000000000000000606082015260806020820152600061446f60808301846185b1565b6000602082840312156190a257600080fd5b815167ffffffffffffffff8111156190b957600080fd5b8201601f810184136190ca57600080fd5b6143f184825160208401618eaa565b600085516190eb818460208a0161858d565b7f2f000000000000000000000000000000000000000000000000000000000000009083019081528551619125816001840160208a0161858d565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161916381600284016020890161858d565b6001818301019150507f2f00000000000000000000000000000000000000000000000000000000000000600182015283516191a581600284016020880161858d565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b6040815260006191f060408301846185b1565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161926781601f85016020870161858d565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b6040815260006192d460408301846185b1565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061932660408301846185b1565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161939d81601485016020870161858d565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b6040815260006193e460408301856185b1565b828103602084015261446b81856185b1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161945d81600185016020870161858d565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b600082516194a381846020870161858d565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161955681604b85016020870161858d565b91909101604b0192915050565b600060ff821660ff810361957957619579618a83565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c694000000000000000000000000000000000000000000000006020820152600082516195e081602985016020870161858d565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f5041544800000000000000000000606082015260806020820152600061446f60808301846185b1565b60006020828403121561964657600080fd5b815167ffffffffffffffff81111561965d57600080fd5b82016060818503121561966f57600080fd5b619677618e81565b81518060030b811461968857600080fd5b8152602082015167ffffffffffffffff8111156196a457600080fd5b6196b086828501618f1a565b602083015250604082015167ffffffffffffffff8111156196d057600080fd5b6196dc86828501618f1a565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161974881602185016020870161858d565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161993481602185016020880161858d565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161997181602e84016020880161858d565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a2000000000000000000000000000000000000000000000006020820152600082516195e081602985016020870161858d565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a000000000000000000000000000000000000000000000000000000000000602082015260008251619a3981602285016020870161858d565b9190910160220192915050565b7f436f6e7472616374206e616d6520000000000000000000000000000000000000815260008251619a7e81600e85016020870161858d565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e746966696572200000000000000000815260008351619b5c81601885016020880161858d565b7f20696e20000000000000000000000000000000000000000000000000000000006018918401918201528351619b9981601c84016020880161858d565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b60008251619c9f81846020870161858d565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f72654000000000815260008251619d0681601c85016020870161858d565b91909101601c0192915050565b60006000198203619d2657619d26618a83565b5060010190565b80820281158282048414176143f5576143f5618a83565b6001815b6001841115619d7f57808504811115619d6357619d63618a83565b6001841615619d7157908102905b60019390931c928002619d48565b935093915050565b600082619d96575060016143f5565b81619da3575060006143f5565b8160018114619db95760028114619dc357619ddf565b60019150506143f5565b60ff841115619dd457619dd4618a83565b50506001821b6143f5565b5060208310610133831016604e8410600b8410161715619e02575081810a6143f5565b619e0f6000198484619d44565b8060001904821115619e2357619e23618a83565b029392505050565b600061446f8383619d87565b600081619e4657619e46618a83565b506000190190565b60008351619e6081846020880161858d565b7f3a000000000000000000000000000000000000000000000000000000000000009083019081528351619e9a81600184016020880161858d565b01600101949350505050565b818103600083128015838313168383128216171561773357617733618a8356fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220837c7d9916de10b20cdb18567d8c7679613426bbd7b0b72548d8000a412f307b64736f6c634300081a0033608060405234801561001057600080fd5b506040516112a63803806112a683398101604081905261002f91610110565b604051806040016040528060048152602001635a65746160e01b815250604051806040016040528060048152602001635a45544160e01b815250816003908161007891906101e2565b50600461008582826101e2565b5050506001600160a01b03821615806100a557506001600160a01b038116155b156100c35760405163e6c4247b60e01b815260040160405180910390fd5b600680546001600160a01b039384166001600160a01b031991821617909155600780549290931691161790556102a0565b80516001600160a01b038116811461010b57600080fd5b919050565b6000806040838503121561012357600080fd5b61012c836100f4565b915061013a602084016100f4565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061016d57607f821691505b60208210810361018d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101dd57806000526020600020601f840160051c810160208510156101ba5750805b601f840160051c820191505b818110156101da57600081556001016101c6565b50505b505050565b81516001600160401b038111156101fb576101fb610143565b61020f816102098454610159565b84610193565b6020601f821160018114610243576000831561022b5750848201515b600019600385901b1c1916600184901b1784556101da565b600084815260208120601f198516915b828110156102735787850151825560209485019460019092019101610253565b50848210156102915786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610ff7806102af6000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806342966c68116100b257806379cc679011610081578063a9059cbb11610066578063a9059cbb1461028e578063bff9662a146102a1578063dd62ed3e146102c157600080fd5b806379cc67901461027357806395d89b411461028657600080fd5b806342966c68146102025780635b1125911461021557806370a0823114610235578063779e3b631461026b57600080fd5b80631e458bee116100ee5780631e458bee1461018857806323b872dd1461019b578063313ce567146101ae578063328a01d0146101bd57600080fd5b806306fdde0314610120578063095ea7b31461013e57806315d57fd41461016157806318160ddd14610176575b600080fd5b610128610307565b6040516101359190610d97565b60405180910390f35b61015161014c366004610e2c565b610399565b6040519015158152602001610135565b61017461016f366004610e56565b6103b3565b005b6002545b604051908152602001610135565b610174610196366004610e89565b61057e565b6101516101a9366004610ebc565b610631565b60405160128152602001610135565b6007546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610135565b610174610210366004610ef9565b610655565b6006546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a610243366004610f12565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b610174610662565b610174610281366004610e2c565b610786565b610128610837565b61015161029c366004610e2c565b610846565b6005546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a6102cf366004610e56565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b60606003805461031690610f34565b80601f016020809104026020016040519081016040528092919081815260200182805461034290610f34565b801561038f5780601f106103645761010080835404028352916020019161038f565b820191906000526020600020905b81548152906001019060200180831161037257829003601f168201915b5050505050905090565b6000336103a7818585610854565b60019150505b92915050565b60075473ffffffffffffffffffffffffffffffffffffffff1633148015906103f3575060065473ffffffffffffffffffffffffffffffffffffffff163314155b15610431576040517fcdfcef970000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82161580610468575073ffffffffffffffffffffffffffffffffffffffff8116155b1561049f576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316811790935560058054918516919092161790556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff910160405180910390a16040805133815273ffffffffffffffffffffffffffffffffffffffff831660208201527f1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c910160405180910390a15050565b60055473ffffffffffffffffffffffffffffffffffffffff1633146105d1576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6105db8383610866565b808373ffffffffffffffffffffffffffffffffffffffff167fc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb8460405161062491815260200190565b60405180910390a3505050565b60003361063f8582856108c6565b61064a858585610995565b506001949350505050565b61065f3382610a40565b50565b60075473ffffffffffffffffffffffffffffffffffffffff1633146106b5576040517fe700765e000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b60065473ffffffffffffffffffffffffffffffffffffffff16610704576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600654600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0910160405180910390a1565b60055473ffffffffffffffffffffffffffffffffffffffff1633146107d9576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6107e38282610a9c565b8173ffffffffffffffffffffffffffffffffffffffff167f919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b18260405161082b91815260200190565b60405180910390a25050565b60606004805461031690610f34565b6000336103a7818585610995565b6108618383836001610ab1565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166108b6576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c260008383610bf9565b5050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461098f5781811015610980576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610428565b61098f84848484036000610ab1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166109e5576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8216610a35576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b610861838383610bf9565b73ffffffffffffffffffffffffffffffffffffffff8216610a90576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c282600083610bf9565b610aa78233836108c6565b6108c28282610a40565b73ffffffffffffffffffffffffffffffffffffffff8416610b01576040517fe602df0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8316610b51576040517f94280d6200000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600160209081526040808320938716835292905220829055801561098f578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610beb91815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c31578060026000828254610c269190610f87565b90915550610ce39050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610cb7576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610428565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610d0c57600280548290039055610d38565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161062491815260200190565b602081526000825180602084015260005b81811015610dc55760208186018101516040868401015201610da8565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610e2757600080fd5b919050565b60008060408385031215610e3f57600080fd5b610e4883610e03565b946020939093013593505050565b60008060408385031215610e6957600080fd5b610e7283610e03565b9150610e8060208401610e03565b90509250929050565b600080600060608486031215610e9e57600080fd5b610ea784610e03565b95602085013595506040909401359392505050565b600080600060608486031215610ed157600080fd5b610eda84610e03565b9250610ee860208501610e03565b929592945050506040919091013590565b600060208284031215610f0b57600080fd5b5035919050565b600060208284031215610f2457600080fd5b610f2d82610e03565b9392505050565b600181811c90821680610f4857607f821691505b602082108103610f81577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b808201808211156103ad577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220fec2a89bc2f37ba9f63e9d7e47fa9a0d2b02402309f8905f94cc44fb9e5fbba064736f6c634300081a0033a264697066735822122026c2ee184095c5d8f3430f20cec395af72b5a44ec6f7988a85c24716e7f4e26f64736f6c634300081a0033",
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
