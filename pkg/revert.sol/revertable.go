// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package revert

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

// RevertableMetaData contains all meta data concerning the Revertable contract.
var RevertableMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"onRevert\",\"inputs\":[{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// RevertableABI is the input ABI used to generate the binding from.
// Deprecated: Use RevertableMetaData.ABI instead.
var RevertableABI = RevertableMetaData.ABI

// Revertable is an auto generated Go binding around an Ethereum contract.
type Revertable struct {
	RevertableCaller     // Read-only binding to the contract
	RevertableTransactor // Write-only binding to the contract
	RevertableFilterer   // Log filterer for contract events
}

// RevertableCaller is an auto generated read-only Go binding around an Ethereum contract.
type RevertableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RevertableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RevertableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RevertableSession struct {
	Contract     *Revertable       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RevertableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RevertableCallerSession struct {
	Contract *RevertableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RevertableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RevertableTransactorSession struct {
	Contract     *RevertableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RevertableRaw is an auto generated low-level Go binding around an Ethereum contract.
type RevertableRaw struct {
	Contract *Revertable // Generic contract binding to access the raw methods on
}

// RevertableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RevertableCallerRaw struct {
	Contract *RevertableCaller // Generic read-only contract binding to access the raw methods on
}

// RevertableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RevertableTransactorRaw struct {
	Contract *RevertableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRevertable creates a new instance of Revertable, bound to a specific deployed contract.
func NewRevertable(address common.Address, backend bind.ContractBackend) (*Revertable, error) {
	contract, err := bindRevertable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Revertable{RevertableCaller: RevertableCaller{contract: contract}, RevertableTransactor: RevertableTransactor{contract: contract}, RevertableFilterer: RevertableFilterer{contract: contract}}, nil
}

// NewRevertableCaller creates a new read-only instance of Revertable, bound to a specific deployed contract.
func NewRevertableCaller(address common.Address, caller bind.ContractCaller) (*RevertableCaller, error) {
	contract, err := bindRevertable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RevertableCaller{contract: contract}, nil
}

// NewRevertableTransactor creates a new write-only instance of Revertable, bound to a specific deployed contract.
func NewRevertableTransactor(address common.Address, transactor bind.ContractTransactor) (*RevertableTransactor, error) {
	contract, err := bindRevertable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RevertableTransactor{contract: contract}, nil
}

// NewRevertableFilterer creates a new log filterer instance of Revertable, bound to a specific deployed contract.
func NewRevertableFilterer(address common.Address, filterer bind.ContractFilterer) (*RevertableFilterer, error) {
	contract, err := bindRevertable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RevertableFilterer{contract: contract}, nil
}

// bindRevertable binds a generic wrapper to an already deployed contract.
func bindRevertable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RevertableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Revertable *RevertableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Revertable.Contract.RevertableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Revertable *RevertableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Revertable.Contract.RevertableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Revertable *RevertableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Revertable.Contract.RevertableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Revertable *RevertableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Revertable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Revertable *RevertableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Revertable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Revertable *RevertableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Revertable.Contract.contract.Transact(opts, method, params...)
}

// OnRevert is a paid mutator transaction binding the contract method 0xc9028a36.
//
// Solidity: function onRevert((address,address,uint256,bytes) revertContext) returns()
func (_Revertable *RevertableTransactor) OnRevert(opts *bind.TransactOpts, revertContext RevertContext) (*types.Transaction, error) {
	return _Revertable.contract.Transact(opts, "onRevert", revertContext)
}

// OnRevert is a paid mutator transaction binding the contract method 0xc9028a36.
//
// Solidity: function onRevert((address,address,uint256,bytes) revertContext) returns()
func (_Revertable *RevertableSession) OnRevert(revertContext RevertContext) (*types.Transaction, error) {
	return _Revertable.Contract.OnRevert(&_Revertable.TransactOpts, revertContext)
}

// OnRevert is a paid mutator transaction binding the contract method 0xc9028a36.
//
// Solidity: function onRevert((address,address,uint256,bytes) revertContext) returns()
func (_Revertable *RevertableTransactorSession) OnRevert(revertContext RevertContext) (*types.Transaction, error) {
	return _Revertable.Contract.OnRevert(&_Revertable.TransactOpts, revertContext)
}
