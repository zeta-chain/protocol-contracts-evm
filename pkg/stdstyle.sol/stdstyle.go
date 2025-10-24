// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stdstyle

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

// StdStyleMetaData contains all meta data concerning the StdStyle contract.
var StdStyleMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea26469706673582212201846a05a05b69c65f328c792504e7ae392cb809f0f3d640d405842253040963a64736f6c634300081a0033",
}

// StdStyleABI is the input ABI used to generate the binding from.
// Deprecated: Use StdStyleMetaData.ABI instead.
var StdStyleABI = StdStyleMetaData.ABI

// StdStyleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StdStyleMetaData.Bin instead.
var StdStyleBin = StdStyleMetaData.Bin

// DeployStdStyle deploys a new Ethereum contract, binding an instance of StdStyle to it.
func DeployStdStyle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StdStyle, error) {
	parsed, err := StdStyleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StdStyleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StdStyle{StdStyleCaller: StdStyleCaller{contract: contract}, StdStyleTransactor: StdStyleTransactor{contract: contract}, StdStyleFilterer: StdStyleFilterer{contract: contract}}, nil
}

// StdStyle is an auto generated Go binding around an Ethereum contract.
type StdStyle struct {
	StdStyleCaller     // Read-only binding to the contract
	StdStyleTransactor // Write-only binding to the contract
	StdStyleFilterer   // Log filterer for contract events
}

// StdStyleCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdStyleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdStyleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdStyleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdStyleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdStyleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdStyleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdStyleSession struct {
	Contract     *StdStyle         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdStyleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdStyleCallerSession struct {
	Contract *StdStyleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StdStyleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdStyleTransactorSession struct {
	Contract     *StdStyleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StdStyleRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdStyleRaw struct {
	Contract *StdStyle // Generic contract binding to access the raw methods on
}

// StdStyleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdStyleCallerRaw struct {
	Contract *StdStyleCaller // Generic read-only contract binding to access the raw methods on
}

// StdStyleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdStyleTransactorRaw struct {
	Contract *StdStyleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdStyle creates a new instance of StdStyle, bound to a specific deployed contract.
func NewStdStyle(address common.Address, backend bind.ContractBackend) (*StdStyle, error) {
	contract, err := bindStdStyle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StdStyle{StdStyleCaller: StdStyleCaller{contract: contract}, StdStyleTransactor: StdStyleTransactor{contract: contract}, StdStyleFilterer: StdStyleFilterer{contract: contract}}, nil
}

// NewStdStyleCaller creates a new read-only instance of StdStyle, bound to a specific deployed contract.
func NewStdStyleCaller(address common.Address, caller bind.ContractCaller) (*StdStyleCaller, error) {
	contract, err := bindStdStyle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdStyleCaller{contract: contract}, nil
}

// NewStdStyleTransactor creates a new write-only instance of StdStyle, bound to a specific deployed contract.
func NewStdStyleTransactor(address common.Address, transactor bind.ContractTransactor) (*StdStyleTransactor, error) {
	contract, err := bindStdStyle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdStyleTransactor{contract: contract}, nil
}

// NewStdStyleFilterer creates a new log filterer instance of StdStyle, bound to a specific deployed contract.
func NewStdStyleFilterer(address common.Address, filterer bind.ContractFilterer) (*StdStyleFilterer, error) {
	contract, err := bindStdStyle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdStyleFilterer{contract: contract}, nil
}

// bindStdStyle binds a generic wrapper to an already deployed contract.
func bindStdStyle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StdStyleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdStyle *StdStyleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdStyle.Contract.StdStyleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdStyle *StdStyleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdStyle.Contract.StdStyleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdStyle *StdStyleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdStyle.Contract.StdStyleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdStyle *StdStyleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdStyle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdStyle *StdStyleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdStyle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdStyle *StdStyleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdStyle.Contract.contract.Transact(opts, method, params...)
}
