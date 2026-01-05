// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayzevmvalidations

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

// GatewayZEVMValidationsMetaData contains all meta data concerning the GatewayZEVMValidations contract.
var GatewayZEVMValidationsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220097a7332c0b3783a9020c207b06a8ef8045aa36846eaa5f399c1b8a44147db7f64736f6c634300081a0033",
}

// GatewayZEVMValidationsABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayZEVMValidationsMetaData.ABI instead.
var GatewayZEVMValidationsABI = GatewayZEVMValidationsMetaData.ABI

// GatewayZEVMValidationsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayZEVMValidationsMetaData.Bin instead.
var GatewayZEVMValidationsBin = GatewayZEVMValidationsMetaData.Bin

// DeployGatewayZEVMValidations deploys a new Ethereum contract, binding an instance of GatewayZEVMValidations to it.
func DeployGatewayZEVMValidations(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayZEVMValidations, error) {
	parsed, err := GatewayZEVMValidationsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayZEVMValidationsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayZEVMValidations{GatewayZEVMValidationsCaller: GatewayZEVMValidationsCaller{contract: contract}, GatewayZEVMValidationsTransactor: GatewayZEVMValidationsTransactor{contract: contract}, GatewayZEVMValidationsFilterer: GatewayZEVMValidationsFilterer{contract: contract}}, nil
}

// GatewayZEVMValidations is an auto generated Go binding around an Ethereum contract.
type GatewayZEVMValidations struct {
	GatewayZEVMValidationsCaller     // Read-only binding to the contract
	GatewayZEVMValidationsTransactor // Write-only binding to the contract
	GatewayZEVMValidationsFilterer   // Log filterer for contract events
}

// GatewayZEVMValidationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayZEVMValidationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMValidationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayZEVMValidationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMValidationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayZEVMValidationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayZEVMValidationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayZEVMValidationsSession struct {
	Contract     *GatewayZEVMValidations // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GatewayZEVMValidationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayZEVMValidationsCallerSession struct {
	Contract *GatewayZEVMValidationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// GatewayZEVMValidationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayZEVMValidationsTransactorSession struct {
	Contract     *GatewayZEVMValidationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// GatewayZEVMValidationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayZEVMValidationsRaw struct {
	Contract *GatewayZEVMValidations // Generic contract binding to access the raw methods on
}

// GatewayZEVMValidationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayZEVMValidationsCallerRaw struct {
	Contract *GatewayZEVMValidationsCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayZEVMValidationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayZEVMValidationsTransactorRaw struct {
	Contract *GatewayZEVMValidationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayZEVMValidations creates a new instance of GatewayZEVMValidations, bound to a specific deployed contract.
func NewGatewayZEVMValidations(address common.Address, backend bind.ContractBackend) (*GatewayZEVMValidations, error) {
	contract, err := bindGatewayZEVMValidations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMValidations{GatewayZEVMValidationsCaller: GatewayZEVMValidationsCaller{contract: contract}, GatewayZEVMValidationsTransactor: GatewayZEVMValidationsTransactor{contract: contract}, GatewayZEVMValidationsFilterer: GatewayZEVMValidationsFilterer{contract: contract}}, nil
}

// NewGatewayZEVMValidationsCaller creates a new read-only instance of GatewayZEVMValidations, bound to a specific deployed contract.
func NewGatewayZEVMValidationsCaller(address common.Address, caller bind.ContractCaller) (*GatewayZEVMValidationsCaller, error) {
	contract, err := bindGatewayZEVMValidations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMValidationsCaller{contract: contract}, nil
}

// NewGatewayZEVMValidationsTransactor creates a new write-only instance of GatewayZEVMValidations, bound to a specific deployed contract.
func NewGatewayZEVMValidationsTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayZEVMValidationsTransactor, error) {
	contract, err := bindGatewayZEVMValidations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMValidationsTransactor{contract: contract}, nil
}

// NewGatewayZEVMValidationsFilterer creates a new log filterer instance of GatewayZEVMValidations, bound to a specific deployed contract.
func NewGatewayZEVMValidationsFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayZEVMValidationsFilterer, error) {
	contract, err := bindGatewayZEVMValidations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayZEVMValidationsFilterer{contract: contract}, nil
}

// bindGatewayZEVMValidations binds a generic wrapper to an already deployed contract.
func bindGatewayZEVMValidations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayZEVMValidationsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMValidations *GatewayZEVMValidationsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMValidations.Contract.GatewayZEVMValidationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMValidations *GatewayZEVMValidationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMValidations.Contract.GatewayZEVMValidationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMValidations *GatewayZEVMValidationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMValidations.Contract.GatewayZEVMValidationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayZEVMValidations *GatewayZEVMValidationsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayZEVMValidations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayZEVMValidations *GatewayZEVMValidationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayZEVMValidations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayZEVMValidations *GatewayZEVMValidationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayZEVMValidations.Contract.contract.Transact(opts, method, params...)
}
