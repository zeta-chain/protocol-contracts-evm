// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package core

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

// CoreMetaData contains all meta data concerning the Core contract.
var CoreMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f522f1786830eb0cf35af3a240bea9bb3fe52bd935eff3296c1ec6d77e18316064736f6c634300081a0033",
}

// CoreABI is the input ABI used to generate the binding from.
// Deprecated: Use CoreMetaData.ABI instead.
var CoreABI = CoreMetaData.ABI

// CoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CoreMetaData.Bin instead.
var CoreBin = CoreMetaData.Bin

// DeployCore deploys a new Ethereum contract, binding an instance of Core to it.
func DeployCore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Core, error) {
	parsed, err := CoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Core{CoreCaller: CoreCaller{contract: contract}, CoreTransactor: CoreTransactor{contract: contract}, CoreFilterer: CoreFilterer{contract: contract}}, nil
}

// Core is an auto generated Go binding around an Ethereum contract.
type Core struct {
	CoreCaller     // Read-only binding to the contract
	CoreTransactor // Write-only binding to the contract
	CoreFilterer   // Log filterer for contract events
}

// CoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoreSession struct {
	Contract     *Core             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoreCallerSession struct {
	Contract *CoreCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoreTransactorSession struct {
	Contract     *CoreTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoreRaw struct {
	Contract *Core // Generic contract binding to access the raw methods on
}

// CoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoreCallerRaw struct {
	Contract *CoreCaller // Generic read-only contract binding to access the raw methods on
}

// CoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoreTransactorRaw struct {
	Contract *CoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCore creates a new instance of Core, bound to a specific deployed contract.
func NewCore(address common.Address, backend bind.ContractBackend) (*Core, error) {
	contract, err := bindCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Core{CoreCaller: CoreCaller{contract: contract}, CoreTransactor: CoreTransactor{contract: contract}, CoreFilterer: CoreFilterer{contract: contract}}, nil
}

// NewCoreCaller creates a new read-only instance of Core, bound to a specific deployed contract.
func NewCoreCaller(address common.Address, caller bind.ContractCaller) (*CoreCaller, error) {
	contract, err := bindCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoreCaller{contract: contract}, nil
}

// NewCoreTransactor creates a new write-only instance of Core, bound to a specific deployed contract.
func NewCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*CoreTransactor, error) {
	contract, err := bindCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoreTransactor{contract: contract}, nil
}

// NewCoreFilterer creates a new log filterer instance of Core, bound to a specific deployed contract.
func NewCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*CoreFilterer, error) {
	contract, err := bindCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoreFilterer{contract: contract}, nil
}

// bindCore binds a generic wrapper to an already deployed contract.
func bindCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Core *CoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Core.Contract.CoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Core *CoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.Contract.CoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Core *CoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Core.Contract.CoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Core *CoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Core.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Core *CoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Core.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Core *CoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Core.Contract.contract.Transact(opts, method, params...)
}
