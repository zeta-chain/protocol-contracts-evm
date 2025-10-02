// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package versions

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

// VersionsMetaData contains all meta data concerning the Versions contract.
var VersionsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122057b47491dea85370ad9620e69892fad0dc6e9b8e4abae7bf5e07b7e52c4eef8764736f6c634300081a0033",
}

// VersionsABI is the input ABI used to generate the binding from.
// Deprecated: Use VersionsMetaData.ABI instead.
var VersionsABI = VersionsMetaData.ABI

// VersionsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VersionsMetaData.Bin instead.
var VersionsBin = VersionsMetaData.Bin

// DeployVersions deploys a new Ethereum contract, binding an instance of Versions to it.
func DeployVersions(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Versions, error) {
	parsed, err := VersionsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VersionsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Versions{VersionsCaller: VersionsCaller{contract: contract}, VersionsTransactor: VersionsTransactor{contract: contract}, VersionsFilterer: VersionsFilterer{contract: contract}}, nil
}

// Versions is an auto generated Go binding around an Ethereum contract.
type Versions struct {
	VersionsCaller     // Read-only binding to the contract
	VersionsTransactor // Write-only binding to the contract
	VersionsFilterer   // Log filterer for contract events
}

// VersionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type VersionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VersionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VersionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VersionsSession struct {
	Contract     *Versions         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VersionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VersionsCallerSession struct {
	Contract *VersionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VersionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VersionsTransactorSession struct {
	Contract     *VersionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VersionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type VersionsRaw struct {
	Contract *Versions // Generic contract binding to access the raw methods on
}

// VersionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VersionsCallerRaw struct {
	Contract *VersionsCaller // Generic read-only contract binding to access the raw methods on
}

// VersionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VersionsTransactorRaw struct {
	Contract *VersionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVersions creates a new instance of Versions, bound to a specific deployed contract.
func NewVersions(address common.Address, backend bind.ContractBackend) (*Versions, error) {
	contract, err := bindVersions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Versions{VersionsCaller: VersionsCaller{contract: contract}, VersionsTransactor: VersionsTransactor{contract: contract}, VersionsFilterer: VersionsFilterer{contract: contract}}, nil
}

// NewVersionsCaller creates a new read-only instance of Versions, bound to a specific deployed contract.
func NewVersionsCaller(address common.Address, caller bind.ContractCaller) (*VersionsCaller, error) {
	contract, err := bindVersions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VersionsCaller{contract: contract}, nil
}

// NewVersionsTransactor creates a new write-only instance of Versions, bound to a specific deployed contract.
func NewVersionsTransactor(address common.Address, transactor bind.ContractTransactor) (*VersionsTransactor, error) {
	contract, err := bindVersions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VersionsTransactor{contract: contract}, nil
}

// NewVersionsFilterer creates a new log filterer instance of Versions, bound to a specific deployed contract.
func NewVersionsFilterer(address common.Address, filterer bind.ContractFilterer) (*VersionsFilterer, error) {
	contract, err := bindVersions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VersionsFilterer{contract: contract}, nil
}

// bindVersions binds a generic wrapper to an already deployed contract.
func bindVersions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VersionsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Versions *VersionsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Versions.Contract.VersionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Versions *VersionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Versions.Contract.VersionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Versions *VersionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Versions.Contract.VersionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Versions *VersionsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Versions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Versions *VersionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Versions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Versions *VersionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Versions.Contract.contract.Transact(opts, method, params...)
}
