// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package panic

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

// PanicMetaData contains all meta data concerning the Panic contract.
var PanicMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220156113f4d64d0b3e47f47ea13f4da4ec55a7b071f6e7740a102b010275ede12364736f6c634300081a0033",
}

// PanicABI is the input ABI used to generate the binding from.
// Deprecated: Use PanicMetaData.ABI instead.
var PanicABI = PanicMetaData.ABI

// PanicBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PanicMetaData.Bin instead.
var PanicBin = PanicMetaData.Bin

// DeployPanic deploys a new Ethereum contract, binding an instance of Panic to it.
func DeployPanic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Panic, error) {
	parsed, err := PanicMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PanicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Panic{PanicCaller: PanicCaller{contract: contract}, PanicTransactor: PanicTransactor{contract: contract}, PanicFilterer: PanicFilterer{contract: contract}}, nil
}

// Panic is an auto generated Go binding around an Ethereum contract.
type Panic struct {
	PanicCaller     // Read-only binding to the contract
	PanicTransactor // Write-only binding to the contract
	PanicFilterer   // Log filterer for contract events
}

// PanicCaller is an auto generated read-only Go binding around an Ethereum contract.
type PanicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PanicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PanicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PanicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PanicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PanicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PanicSession struct {
	Contract     *Panic            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PanicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PanicCallerSession struct {
	Contract *PanicCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PanicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PanicTransactorSession struct {
	Contract     *PanicTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PanicRaw is an auto generated low-level Go binding around an Ethereum contract.
type PanicRaw struct {
	Contract *Panic // Generic contract binding to access the raw methods on
}

// PanicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PanicCallerRaw struct {
	Contract *PanicCaller // Generic read-only contract binding to access the raw methods on
}

// PanicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PanicTransactorRaw struct {
	Contract *PanicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPanic creates a new instance of Panic, bound to a specific deployed contract.
func NewPanic(address common.Address, backend bind.ContractBackend) (*Panic, error) {
	contract, err := bindPanic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Panic{PanicCaller: PanicCaller{contract: contract}, PanicTransactor: PanicTransactor{contract: contract}, PanicFilterer: PanicFilterer{contract: contract}}, nil
}

// NewPanicCaller creates a new read-only instance of Panic, bound to a specific deployed contract.
func NewPanicCaller(address common.Address, caller bind.ContractCaller) (*PanicCaller, error) {
	contract, err := bindPanic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PanicCaller{contract: contract}, nil
}

// NewPanicTransactor creates a new write-only instance of Panic, bound to a specific deployed contract.
func NewPanicTransactor(address common.Address, transactor bind.ContractTransactor) (*PanicTransactor, error) {
	contract, err := bindPanic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PanicTransactor{contract: contract}, nil
}

// NewPanicFilterer creates a new log filterer instance of Panic, bound to a specific deployed contract.
func NewPanicFilterer(address common.Address, filterer bind.ContractFilterer) (*PanicFilterer, error) {
	contract, err := bindPanic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PanicFilterer{contract: contract}, nil
}

// bindPanic binds a generic wrapper to an already deployed contract.
func bindPanic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PanicMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Panic *PanicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Panic.Contract.PanicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Panic *PanicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Panic.Contract.PanicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Panic *PanicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Panic.Contract.PanicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Panic *PanicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Panic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Panic *PanicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Panic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Panic *PanicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Panic.Contract.contract.Transact(opts, method, params...)
}
