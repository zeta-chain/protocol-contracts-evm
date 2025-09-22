// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc1967utils

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

// ERC1967UtilsMetaData contains all meta data concerning the ERC1967Utils contract.
var ERC1967UtilsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"ERC1967InvalidAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidBeacon\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]}]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea264697066735822122027aeb13eab0b487628dd3f9b05a6b0cb3625c061cd235953f9da355b60d4763164736f6c634300081a0033",
}

// ERC1967UtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1967UtilsMetaData.ABI instead.
var ERC1967UtilsABI = ERC1967UtilsMetaData.ABI

// ERC1967UtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC1967UtilsMetaData.Bin instead.
var ERC1967UtilsBin = ERC1967UtilsMetaData.Bin

// DeployERC1967Utils deploys a new Ethereum contract, binding an instance of ERC1967Utils to it.
func DeployERC1967Utils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC1967Utils, error) {
	parsed, err := ERC1967UtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC1967UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC1967Utils{ERC1967UtilsCaller: ERC1967UtilsCaller{contract: contract}, ERC1967UtilsTransactor: ERC1967UtilsTransactor{contract: contract}, ERC1967UtilsFilterer: ERC1967UtilsFilterer{contract: contract}}, nil
}

// ERC1967Utils is an auto generated Go binding around an Ethereum contract.
type ERC1967Utils struct {
	ERC1967UtilsCaller     // Read-only binding to the contract
	ERC1967UtilsTransactor // Write-only binding to the contract
	ERC1967UtilsFilterer   // Log filterer for contract events
}

// ERC1967UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1967UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1967UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1967UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1967UtilsSession struct {
	Contract     *ERC1967Utils     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1967UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1967UtilsCallerSession struct {
	Contract *ERC1967UtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC1967UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1967UtilsTransactorSession struct {
	Contract     *ERC1967UtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC1967UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1967UtilsRaw struct {
	Contract *ERC1967Utils // Generic contract binding to access the raw methods on
}

// ERC1967UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1967UtilsCallerRaw struct {
	Contract *ERC1967UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1967UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1967UtilsTransactorRaw struct {
	Contract *ERC1967UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1967Utils creates a new instance of ERC1967Utils, bound to a specific deployed contract.
func NewERC1967Utils(address common.Address, backend bind.ContractBackend) (*ERC1967Utils, error) {
	contract, err := bindERC1967Utils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1967Utils{ERC1967UtilsCaller: ERC1967UtilsCaller{contract: contract}, ERC1967UtilsTransactor: ERC1967UtilsTransactor{contract: contract}, ERC1967UtilsFilterer: ERC1967UtilsFilterer{contract: contract}}, nil
}

// NewERC1967UtilsCaller creates a new read-only instance of ERC1967Utils, bound to a specific deployed contract.
func NewERC1967UtilsCaller(address common.Address, caller bind.ContractCaller) (*ERC1967UtilsCaller, error) {
	contract, err := bindERC1967Utils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UtilsCaller{contract: contract}, nil
}

// NewERC1967UtilsTransactor creates a new write-only instance of ERC1967Utils, bound to a specific deployed contract.
func NewERC1967UtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1967UtilsTransactor, error) {
	contract, err := bindERC1967Utils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UtilsTransactor{contract: contract}, nil
}

// NewERC1967UtilsFilterer creates a new log filterer instance of ERC1967Utils, bound to a specific deployed contract.
func NewERC1967UtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1967UtilsFilterer, error) {
	contract, err := bindERC1967Utils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1967UtilsFilterer{contract: contract}, nil
}

// bindERC1967Utils binds a generic wrapper to an already deployed contract.
func bindERC1967Utils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1967UtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967Utils *ERC1967UtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967Utils.Contract.ERC1967UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967Utils *ERC1967UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967Utils.Contract.ERC1967UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967Utils *ERC1967UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967Utils.Contract.ERC1967UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967Utils *ERC1967UtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967Utils *ERC1967UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967Utils *ERC1967UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967Utils.Contract.contract.Transact(opts, method, params...)
}
