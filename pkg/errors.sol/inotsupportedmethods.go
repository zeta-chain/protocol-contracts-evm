// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package errors

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

// INotSupportedMethodsMetaData contains all meta data concerning the INotSupportedMethods contract.
var INotSupportedMethodsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"CallOnRevertNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZETANotSupported\",\"inputs\":[]}]",
}

// INotSupportedMethodsABI is the input ABI used to generate the binding from.
// Deprecated: Use INotSupportedMethodsMetaData.ABI instead.
var INotSupportedMethodsABI = INotSupportedMethodsMetaData.ABI

// INotSupportedMethods is an auto generated Go binding around an Ethereum contract.
type INotSupportedMethods struct {
	INotSupportedMethodsCaller     // Read-only binding to the contract
	INotSupportedMethodsTransactor // Write-only binding to the contract
	INotSupportedMethodsFilterer   // Log filterer for contract events
}

// INotSupportedMethodsCaller is an auto generated read-only Go binding around an Ethereum contract.
type INotSupportedMethodsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INotSupportedMethodsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INotSupportedMethodsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INotSupportedMethodsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INotSupportedMethodsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INotSupportedMethodsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INotSupportedMethodsSession struct {
	Contract     *INotSupportedMethods // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// INotSupportedMethodsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INotSupportedMethodsCallerSession struct {
	Contract *INotSupportedMethodsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// INotSupportedMethodsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INotSupportedMethodsTransactorSession struct {
	Contract     *INotSupportedMethodsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// INotSupportedMethodsRaw is an auto generated low-level Go binding around an Ethereum contract.
type INotSupportedMethodsRaw struct {
	Contract *INotSupportedMethods // Generic contract binding to access the raw methods on
}

// INotSupportedMethodsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INotSupportedMethodsCallerRaw struct {
	Contract *INotSupportedMethodsCaller // Generic read-only contract binding to access the raw methods on
}

// INotSupportedMethodsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INotSupportedMethodsTransactorRaw struct {
	Contract *INotSupportedMethodsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINotSupportedMethods creates a new instance of INotSupportedMethods, bound to a specific deployed contract.
func NewINotSupportedMethods(address common.Address, backend bind.ContractBackend) (*INotSupportedMethods, error) {
	contract, err := bindINotSupportedMethods(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INotSupportedMethods{INotSupportedMethodsCaller: INotSupportedMethodsCaller{contract: contract}, INotSupportedMethodsTransactor: INotSupportedMethodsTransactor{contract: contract}, INotSupportedMethodsFilterer: INotSupportedMethodsFilterer{contract: contract}}, nil
}

// NewINotSupportedMethodsCaller creates a new read-only instance of INotSupportedMethods, bound to a specific deployed contract.
func NewINotSupportedMethodsCaller(address common.Address, caller bind.ContractCaller) (*INotSupportedMethodsCaller, error) {
	contract, err := bindINotSupportedMethods(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INotSupportedMethodsCaller{contract: contract}, nil
}

// NewINotSupportedMethodsTransactor creates a new write-only instance of INotSupportedMethods, bound to a specific deployed contract.
func NewINotSupportedMethodsTransactor(address common.Address, transactor bind.ContractTransactor) (*INotSupportedMethodsTransactor, error) {
	contract, err := bindINotSupportedMethods(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INotSupportedMethodsTransactor{contract: contract}, nil
}

// NewINotSupportedMethodsFilterer creates a new log filterer instance of INotSupportedMethods, bound to a specific deployed contract.
func NewINotSupportedMethodsFilterer(address common.Address, filterer bind.ContractFilterer) (*INotSupportedMethodsFilterer, error) {
	contract, err := bindINotSupportedMethods(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INotSupportedMethodsFilterer{contract: contract}, nil
}

// bindINotSupportedMethods binds a generic wrapper to an already deployed contract.
func bindINotSupportedMethods(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := INotSupportedMethodsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INotSupportedMethods *INotSupportedMethodsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INotSupportedMethods.Contract.INotSupportedMethodsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INotSupportedMethods *INotSupportedMethodsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INotSupportedMethods.Contract.INotSupportedMethodsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INotSupportedMethods *INotSupportedMethodsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INotSupportedMethods.Contract.INotSupportedMethodsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INotSupportedMethods *INotSupportedMethodsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INotSupportedMethods.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INotSupportedMethods *INotSupportedMethodsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INotSupportedMethods.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INotSupportedMethods *INotSupportedMethodsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INotSupportedMethods.Contract.contract.Transact(opts, method, params...)
}
