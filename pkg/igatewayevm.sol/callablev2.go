// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igatewayevm

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

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Sender common.Address
	Asset  common.Address
	Amount *big.Int
}

// CallableV2MetaData contains all meta data concerning the CallableV2 contract.
var CallableV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"onCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"}]",
}

// CallableV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use CallableV2MetaData.ABI instead.
var CallableV2ABI = CallableV2MetaData.ABI

// CallableV2 is an auto generated Go binding around an Ethereum contract.
type CallableV2 struct {
	CallableV2Caller     // Read-only binding to the contract
	CallableV2Transactor // Write-only binding to the contract
	CallableV2Filterer   // Log filterer for contract events
}

// CallableV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type CallableV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallableV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CallableV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallableV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallableV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallableV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallableV2Session struct {
	Contract     *CallableV2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallableV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallableV2CallerSession struct {
	Contract *CallableV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CallableV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallableV2TransactorSession struct {
	Contract     *CallableV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CallableV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type CallableV2Raw struct {
	Contract *CallableV2 // Generic contract binding to access the raw methods on
}

// CallableV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallableV2CallerRaw struct {
	Contract *CallableV2Caller // Generic read-only contract binding to access the raw methods on
}

// CallableV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallableV2TransactorRaw struct {
	Contract *CallableV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCallableV2 creates a new instance of CallableV2, bound to a specific deployed contract.
func NewCallableV2(address common.Address, backend bind.ContractBackend) (*CallableV2, error) {
	contract, err := bindCallableV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CallableV2{CallableV2Caller: CallableV2Caller{contract: contract}, CallableV2Transactor: CallableV2Transactor{contract: contract}, CallableV2Filterer: CallableV2Filterer{contract: contract}}, nil
}

// NewCallableV2Caller creates a new read-only instance of CallableV2, bound to a specific deployed contract.
func NewCallableV2Caller(address common.Address, caller bind.ContractCaller) (*CallableV2Caller, error) {
	contract, err := bindCallableV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallableV2Caller{contract: contract}, nil
}

// NewCallableV2Transactor creates a new write-only instance of CallableV2, bound to a specific deployed contract.
func NewCallableV2Transactor(address common.Address, transactor bind.ContractTransactor) (*CallableV2Transactor, error) {
	contract, err := bindCallableV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallableV2Transactor{contract: contract}, nil
}

// NewCallableV2Filterer creates a new log filterer instance of CallableV2, bound to a specific deployed contract.
func NewCallableV2Filterer(address common.Address, filterer bind.ContractFilterer) (*CallableV2Filterer, error) {
	contract, err := bindCallableV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallableV2Filterer{contract: contract}, nil
}

// bindCallableV2 binds a generic wrapper to an already deployed contract.
func bindCallableV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CallableV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallableV2 *CallableV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CallableV2.Contract.CallableV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallableV2 *CallableV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallableV2.Contract.CallableV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallableV2 *CallableV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallableV2.Contract.CallableV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallableV2 *CallableV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CallableV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallableV2 *CallableV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallableV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallableV2 *CallableV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallableV2.Contract.contract.Transact(opts, method, params...)
}

// OnCall is a paid mutator transaction binding the contract method 0x01a64767.
//
// Solidity: function onCall((address,address,uint256) context, bytes message) payable returns(bytes)
func (_CallableV2 *CallableV2Transactor) OnCall(opts *bind.TransactOpts, context MessageContext, message []byte) (*types.Transaction, error) {
	return _CallableV2.contract.Transact(opts, "onCall", context, message)
}

// OnCall is a paid mutator transaction binding the contract method 0x01a64767.
//
// Solidity: function onCall((address,address,uint256) context, bytes message) payable returns(bytes)
func (_CallableV2 *CallableV2Session) OnCall(context MessageContext, message []byte) (*types.Transaction, error) {
	return _CallableV2.Contract.OnCall(&_CallableV2.TransactOpts, context, message)
}

// OnCall is a paid mutator transaction binding the contract method 0x01a64767.
//
// Solidity: function onCall((address,address,uint256) context, bytes message) payable returns(bytes)
func (_CallableV2 *CallableV2TransactorSession) OnCall(context MessageContext, message []byte) (*types.Transaction, error) {
	return _CallableV2.Contract.OnCall(&_CallableV2.TransactOpts, context, message)
}
