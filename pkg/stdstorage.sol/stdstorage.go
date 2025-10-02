// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stdstorage

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

// StdStorageMetaData contains all meta data concerning the StdStorage contract.
var StdStorageMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220afa38f0721251312b4eea523f3dfb8d89a388704f5684cd97977d50f08a0472c64736f6c634300081a0033",
}

// StdStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StdStorageMetaData.ABI instead.
var StdStorageABI = StdStorageMetaData.ABI

// StdStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StdStorageMetaData.Bin instead.
var StdStorageBin = StdStorageMetaData.Bin

// DeployStdStorage deploys a new Ethereum contract, binding an instance of StdStorage to it.
func DeployStdStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StdStorage, error) {
	parsed, err := StdStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StdStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StdStorage{StdStorageCaller: StdStorageCaller{contract: contract}, StdStorageTransactor: StdStorageTransactor{contract: contract}, StdStorageFilterer: StdStorageFilterer{contract: contract}}, nil
}

// StdStorage is an auto generated Go binding around an Ethereum contract.
type StdStorage struct {
	StdStorageCaller     // Read-only binding to the contract
	StdStorageTransactor // Write-only binding to the contract
	StdStorageFilterer   // Log filterer for contract events
}

// StdStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdStorageSession struct {
	Contract     *StdStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdStorageCallerSession struct {
	Contract *StdStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StdStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdStorageTransactorSession struct {
	Contract     *StdStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StdStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdStorageRaw struct {
	Contract *StdStorage // Generic contract binding to access the raw methods on
}

// StdStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdStorageCallerRaw struct {
	Contract *StdStorageCaller // Generic read-only contract binding to access the raw methods on
}

// StdStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdStorageTransactorRaw struct {
	Contract *StdStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdStorage creates a new instance of StdStorage, bound to a specific deployed contract.
func NewStdStorage(address common.Address, backend bind.ContractBackend) (*StdStorage, error) {
	contract, err := bindStdStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StdStorage{StdStorageCaller: StdStorageCaller{contract: contract}, StdStorageTransactor: StdStorageTransactor{contract: contract}, StdStorageFilterer: StdStorageFilterer{contract: contract}}, nil
}

// NewStdStorageCaller creates a new read-only instance of StdStorage, bound to a specific deployed contract.
func NewStdStorageCaller(address common.Address, caller bind.ContractCaller) (*StdStorageCaller, error) {
	contract, err := bindStdStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdStorageCaller{contract: contract}, nil
}

// NewStdStorageTransactor creates a new write-only instance of StdStorage, bound to a specific deployed contract.
func NewStdStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StdStorageTransactor, error) {
	contract, err := bindStdStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdStorageTransactor{contract: contract}, nil
}

// NewStdStorageFilterer creates a new log filterer instance of StdStorage, bound to a specific deployed contract.
func NewStdStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StdStorageFilterer, error) {
	contract, err := bindStdStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdStorageFilterer{contract: contract}, nil
}

// bindStdStorage binds a generic wrapper to an already deployed contract.
func bindStdStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StdStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdStorage *StdStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdStorage.Contract.StdStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdStorage *StdStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdStorage.Contract.StdStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdStorage *StdStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdStorage.Contract.StdStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdStorage *StdStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdStorage *StdStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdStorage *StdStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdStorage.Contract.contract.Transact(opts, method, params...)
}
