// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package safeconsole

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

// SafeconsoleMetaData contains all meta data concerning the Safeconsole contract.
var SafeconsoleMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207b079e6fb84aa97be24cf6564cf0912ffa73bbf21e832216c586952a580d3bd664736f6c634300081a0033",
}

// SafeconsoleABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeconsoleMetaData.ABI instead.
var SafeconsoleABI = SafeconsoleMetaData.ABI

// SafeconsoleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeconsoleMetaData.Bin instead.
var SafeconsoleBin = SafeconsoleMetaData.Bin

// DeploySafeconsole deploys a new Ethereum contract, binding an instance of Safeconsole to it.
func DeploySafeconsole(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Safeconsole, error) {
	parsed, err := SafeconsoleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeconsoleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Safeconsole{SafeconsoleCaller: SafeconsoleCaller{contract: contract}, SafeconsoleTransactor: SafeconsoleTransactor{contract: contract}, SafeconsoleFilterer: SafeconsoleFilterer{contract: contract}}, nil
}

// Safeconsole is an auto generated Go binding around an Ethereum contract.
type Safeconsole struct {
	SafeconsoleCaller     // Read-only binding to the contract
	SafeconsoleTransactor // Write-only binding to the contract
	SafeconsoleFilterer   // Log filterer for contract events
}

// SafeconsoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeconsoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeconsoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeconsoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeconsoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeconsoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeconsoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeconsoleSession struct {
	Contract     *Safeconsole      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeconsoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeconsoleCallerSession struct {
	Contract *SafeconsoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SafeconsoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeconsoleTransactorSession struct {
	Contract     *SafeconsoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SafeconsoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeconsoleRaw struct {
	Contract *Safeconsole // Generic contract binding to access the raw methods on
}

// SafeconsoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeconsoleCallerRaw struct {
	Contract *SafeconsoleCaller // Generic read-only contract binding to access the raw methods on
}

// SafeconsoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeconsoleTransactorRaw struct {
	Contract *SafeconsoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeconsole creates a new instance of Safeconsole, bound to a specific deployed contract.
func NewSafeconsole(address common.Address, backend bind.ContractBackend) (*Safeconsole, error) {
	contract, err := bindSafeconsole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Safeconsole{SafeconsoleCaller: SafeconsoleCaller{contract: contract}, SafeconsoleTransactor: SafeconsoleTransactor{contract: contract}, SafeconsoleFilterer: SafeconsoleFilterer{contract: contract}}, nil
}

// NewSafeconsoleCaller creates a new read-only instance of Safeconsole, bound to a specific deployed contract.
func NewSafeconsoleCaller(address common.Address, caller bind.ContractCaller) (*SafeconsoleCaller, error) {
	contract, err := bindSafeconsole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeconsoleCaller{contract: contract}, nil
}

// NewSafeconsoleTransactor creates a new write-only instance of Safeconsole, bound to a specific deployed contract.
func NewSafeconsoleTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeconsoleTransactor, error) {
	contract, err := bindSafeconsole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeconsoleTransactor{contract: contract}, nil
}

// NewSafeconsoleFilterer creates a new log filterer instance of Safeconsole, bound to a specific deployed contract.
func NewSafeconsoleFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeconsoleFilterer, error) {
	contract, err := bindSafeconsole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeconsoleFilterer{contract: contract}, nil
}

// bindSafeconsole binds a generic wrapper to an already deployed contract.
func bindSafeconsole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeconsoleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safeconsole *SafeconsoleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Safeconsole.Contract.SafeconsoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safeconsole *SafeconsoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safeconsole.Contract.SafeconsoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safeconsole *SafeconsoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safeconsole.Contract.SafeconsoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Safeconsole *SafeconsoleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Safeconsole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Safeconsole *SafeconsoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Safeconsole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Safeconsole *SafeconsoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Safeconsole.Contract.contract.Transact(opts, method, params...)
}
