// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stdmath

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

// StdMathMetaData contains all meta data concerning the StdMath contract.
var StdMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea264697066735822122087206569d1b02bf23091a74b38877226558e035a74d0b20970d0c00233091a6e64736f6c634300081a0033",
}

// StdMathABI is the input ABI used to generate the binding from.
// Deprecated: Use StdMathMetaData.ABI instead.
var StdMathABI = StdMathMetaData.ABI

// StdMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StdMathMetaData.Bin instead.
var StdMathBin = StdMathMetaData.Bin

// DeployStdMath deploys a new Ethereum contract, binding an instance of StdMath to it.
func DeployStdMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StdMath, error) {
	parsed, err := StdMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StdMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StdMath{StdMathCaller: StdMathCaller{contract: contract}, StdMathTransactor: StdMathTransactor{contract: contract}, StdMathFilterer: StdMathFilterer{contract: contract}}, nil
}

// StdMath is an auto generated Go binding around an Ethereum contract.
type StdMath struct {
	StdMathCaller     // Read-only binding to the contract
	StdMathTransactor // Write-only binding to the contract
	StdMathFilterer   // Log filterer for contract events
}

// StdMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdMathSession struct {
	Contract     *StdMath          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdMathCallerSession struct {
	Contract *StdMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StdMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdMathTransactorSession struct {
	Contract     *StdMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StdMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdMathRaw struct {
	Contract *StdMath // Generic contract binding to access the raw methods on
}

// StdMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdMathCallerRaw struct {
	Contract *StdMathCaller // Generic read-only contract binding to access the raw methods on
}

// StdMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdMathTransactorRaw struct {
	Contract *StdMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdMath creates a new instance of StdMath, bound to a specific deployed contract.
func NewStdMath(address common.Address, backend bind.ContractBackend) (*StdMath, error) {
	contract, err := bindStdMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StdMath{StdMathCaller: StdMathCaller{contract: contract}, StdMathTransactor: StdMathTransactor{contract: contract}, StdMathFilterer: StdMathFilterer{contract: contract}}, nil
}

// NewStdMathCaller creates a new read-only instance of StdMath, bound to a specific deployed contract.
func NewStdMathCaller(address common.Address, caller bind.ContractCaller) (*StdMathCaller, error) {
	contract, err := bindStdMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdMathCaller{contract: contract}, nil
}

// NewStdMathTransactor creates a new write-only instance of StdMath, bound to a specific deployed contract.
func NewStdMathTransactor(address common.Address, transactor bind.ContractTransactor) (*StdMathTransactor, error) {
	contract, err := bindStdMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdMathTransactor{contract: contract}, nil
}

// NewStdMathFilterer creates a new log filterer instance of StdMath, bound to a specific deployed contract.
func NewStdMathFilterer(address common.Address, filterer bind.ContractFilterer) (*StdMathFilterer, error) {
	contract, err := bindStdMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdMathFilterer{contract: contract}, nil
}

// bindStdMath binds a generic wrapper to an already deployed contract.
func bindStdMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StdMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdMath *StdMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdMath.Contract.StdMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdMath *StdMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdMath.Contract.StdMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdMath *StdMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdMath.Contract.StdMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdMath *StdMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdMath *StdMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdMath *StdMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdMath.Contract.contract.Transact(opts, method, params...)
}
