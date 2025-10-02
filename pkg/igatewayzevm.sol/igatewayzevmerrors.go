// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igatewayzevm

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

// IGatewayZEVMErrorsMetaData contains all meta data concerning the IGatewayZEVMErrors contract.
var IGatewayZEVMErrorsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[]}]",
}

// IGatewayZEVMErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IGatewayZEVMErrorsMetaData.ABI instead.
var IGatewayZEVMErrorsABI = IGatewayZEVMErrorsMetaData.ABI

// IGatewayZEVMErrors is an auto generated Go binding around an Ethereum contract.
type IGatewayZEVMErrors struct {
	IGatewayZEVMErrorsCaller     // Read-only binding to the contract
	IGatewayZEVMErrorsTransactor // Write-only binding to the contract
	IGatewayZEVMErrorsFilterer   // Log filterer for contract events
}

// IGatewayZEVMErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGatewayZEVMErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayZEVMErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGatewayZEVMErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayZEVMErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGatewayZEVMErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayZEVMErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGatewayZEVMErrorsSession struct {
	Contract     *IGatewayZEVMErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IGatewayZEVMErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGatewayZEVMErrorsCallerSession struct {
	Contract *IGatewayZEVMErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IGatewayZEVMErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGatewayZEVMErrorsTransactorSession struct {
	Contract     *IGatewayZEVMErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IGatewayZEVMErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGatewayZEVMErrorsRaw struct {
	Contract *IGatewayZEVMErrors // Generic contract binding to access the raw methods on
}

// IGatewayZEVMErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGatewayZEVMErrorsCallerRaw struct {
	Contract *IGatewayZEVMErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IGatewayZEVMErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGatewayZEVMErrorsTransactorRaw struct {
	Contract *IGatewayZEVMErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGatewayZEVMErrors creates a new instance of IGatewayZEVMErrors, bound to a specific deployed contract.
func NewIGatewayZEVMErrors(address common.Address, backend bind.ContractBackend) (*IGatewayZEVMErrors, error) {
	contract, err := bindIGatewayZEVMErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMErrors{IGatewayZEVMErrorsCaller: IGatewayZEVMErrorsCaller{contract: contract}, IGatewayZEVMErrorsTransactor: IGatewayZEVMErrorsTransactor{contract: contract}, IGatewayZEVMErrorsFilterer: IGatewayZEVMErrorsFilterer{contract: contract}}, nil
}

// NewIGatewayZEVMErrorsCaller creates a new read-only instance of IGatewayZEVMErrors, bound to a specific deployed contract.
func NewIGatewayZEVMErrorsCaller(address common.Address, caller bind.ContractCaller) (*IGatewayZEVMErrorsCaller, error) {
	contract, err := bindIGatewayZEVMErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMErrorsCaller{contract: contract}, nil
}

// NewIGatewayZEVMErrorsTransactor creates a new write-only instance of IGatewayZEVMErrors, bound to a specific deployed contract.
func NewIGatewayZEVMErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IGatewayZEVMErrorsTransactor, error) {
	contract, err := bindIGatewayZEVMErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMErrorsTransactor{contract: contract}, nil
}

// NewIGatewayZEVMErrorsFilterer creates a new log filterer instance of IGatewayZEVMErrors, bound to a specific deployed contract.
func NewIGatewayZEVMErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IGatewayZEVMErrorsFilterer, error) {
	contract, err := bindIGatewayZEVMErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGatewayZEVMErrorsFilterer{contract: contract}, nil
}

// bindIGatewayZEVMErrors binds a generic wrapper to an already deployed contract.
func bindIGatewayZEVMErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGatewayZEVMErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayZEVMErrors *IGatewayZEVMErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayZEVMErrors.Contract.IGatewayZEVMErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayZEVMErrors *IGatewayZEVMErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayZEVMErrors.Contract.IGatewayZEVMErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayZEVMErrors *IGatewayZEVMErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayZEVMErrors.Contract.IGatewayZEVMErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayZEVMErrors *IGatewayZEVMErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayZEVMErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayZEVMErrors *IGatewayZEVMErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayZEVMErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayZEVMErrors *IGatewayZEVMErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayZEVMErrors.Contract.contract.Transact(opts, method, params...)
}
