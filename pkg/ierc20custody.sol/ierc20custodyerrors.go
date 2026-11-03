// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc20custody

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

// IERC20CustodyErrorsMetaData contains all meta data concerning the IERC20CustodyErrors contract.
var IERC20CustodyErrorsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"LegacyMethodsNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// IERC20CustodyErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20CustodyErrorsMetaData.ABI instead.
var IERC20CustodyErrorsABI = IERC20CustodyErrorsMetaData.ABI

// IERC20CustodyErrors is an auto generated Go binding around an Ethereum contract.
type IERC20CustodyErrors struct {
	IERC20CustodyErrorsCaller     // Read-only binding to the contract
	IERC20CustodyErrorsTransactor // Write-only binding to the contract
	IERC20CustodyErrorsFilterer   // Log filterer for contract events
}

// IERC20CustodyErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20CustodyErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20CustodyErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20CustodyErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20CustodyErrorsSession struct {
	Contract     *IERC20CustodyErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IERC20CustodyErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CustodyErrorsCallerSession struct {
	Contract *IERC20CustodyErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IERC20CustodyErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20CustodyErrorsTransactorSession struct {
	Contract     *IERC20CustodyErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IERC20CustodyErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20CustodyErrorsRaw struct {
	Contract *IERC20CustodyErrors // Generic contract binding to access the raw methods on
}

// IERC20CustodyErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CustodyErrorsCallerRaw struct {
	Contract *IERC20CustodyErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20CustodyErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20CustodyErrorsTransactorRaw struct {
	Contract *IERC20CustodyErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20CustodyErrors creates a new instance of IERC20CustodyErrors, bound to a specific deployed contract.
func NewIERC20CustodyErrors(address common.Address, backend bind.ContractBackend) (*IERC20CustodyErrors, error) {
	contract, err := bindIERC20CustodyErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyErrors{IERC20CustodyErrorsCaller: IERC20CustodyErrorsCaller{contract: contract}, IERC20CustodyErrorsTransactor: IERC20CustodyErrorsTransactor{contract: contract}, IERC20CustodyErrorsFilterer: IERC20CustodyErrorsFilterer{contract: contract}}, nil
}

// NewIERC20CustodyErrorsCaller creates a new read-only instance of IERC20CustodyErrors, bound to a specific deployed contract.
func NewIERC20CustodyErrorsCaller(address common.Address, caller bind.ContractCaller) (*IERC20CustodyErrorsCaller, error) {
	contract, err := bindIERC20CustodyErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyErrorsCaller{contract: contract}, nil
}

// NewIERC20CustodyErrorsTransactor creates a new write-only instance of IERC20CustodyErrors, bound to a specific deployed contract.
func NewIERC20CustodyErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20CustodyErrorsTransactor, error) {
	contract, err := bindIERC20CustodyErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyErrorsTransactor{contract: contract}, nil
}

// NewIERC20CustodyErrorsFilterer creates a new log filterer instance of IERC20CustodyErrors, bound to a specific deployed contract.
func NewIERC20CustodyErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20CustodyErrorsFilterer, error) {
	contract, err := bindIERC20CustodyErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyErrorsFilterer{contract: contract}, nil
}

// bindIERC20CustodyErrors binds a generic wrapper to an already deployed contract.
func bindIERC20CustodyErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20CustodyErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20CustodyErrors *IERC20CustodyErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20CustodyErrors.Contract.IERC20CustodyErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20CustodyErrors *IERC20CustodyErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20CustodyErrors.Contract.IERC20CustodyErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20CustodyErrors *IERC20CustodyErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20CustodyErrors.Contract.IERC20CustodyErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20CustodyErrors *IERC20CustodyErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20CustodyErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20CustodyErrors *IERC20CustodyErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20CustodyErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20CustodyErrors *IERC20CustodyErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20CustodyErrors.Contract.contract.Transact(opts, method, params...)
}
