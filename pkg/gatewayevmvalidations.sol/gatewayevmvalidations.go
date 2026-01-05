// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayevmvalidations

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

// GatewayEVMValidationsMetaData contains all meta data concerning the GatewayEVMValidations contract.
var GatewayEVMValidationsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea264697066735822122029e3b7a322110c347802890846867bb6a2c8b9179a2a685fe2bd1f3f3eadab8d64736f6c634300081a0033",
}

// GatewayEVMValidationsABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMValidationsMetaData.ABI instead.
var GatewayEVMValidationsABI = GatewayEVMValidationsMetaData.ABI

// GatewayEVMValidationsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMValidationsMetaData.Bin instead.
var GatewayEVMValidationsBin = GatewayEVMValidationsMetaData.Bin

// DeployGatewayEVMValidations deploys a new Ethereum contract, binding an instance of GatewayEVMValidations to it.
func DeployGatewayEVMValidations(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVMValidations, error) {
	parsed, err := GatewayEVMValidationsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMValidationsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVMValidations{GatewayEVMValidationsCaller: GatewayEVMValidationsCaller{contract: contract}, GatewayEVMValidationsTransactor: GatewayEVMValidationsTransactor{contract: contract}, GatewayEVMValidationsFilterer: GatewayEVMValidationsFilterer{contract: contract}}, nil
}

// GatewayEVMValidations is an auto generated Go binding around an Ethereum contract.
type GatewayEVMValidations struct {
	GatewayEVMValidationsCaller     // Read-only binding to the contract
	GatewayEVMValidationsTransactor // Write-only binding to the contract
	GatewayEVMValidationsFilterer   // Log filterer for contract events
}

// GatewayEVMValidationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMValidationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMValidationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMValidationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMValidationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMValidationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMValidationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMValidationsSession struct {
	Contract     *GatewayEVMValidations // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GatewayEVMValidationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMValidationsCallerSession struct {
	Contract *GatewayEVMValidationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// GatewayEVMValidationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMValidationsTransactorSession struct {
	Contract     *GatewayEVMValidationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// GatewayEVMValidationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMValidationsRaw struct {
	Contract *GatewayEVMValidations // Generic contract binding to access the raw methods on
}

// GatewayEVMValidationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMValidationsCallerRaw struct {
	Contract *GatewayEVMValidationsCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMValidationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMValidationsTransactorRaw struct {
	Contract *GatewayEVMValidationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVMValidations creates a new instance of GatewayEVMValidations, bound to a specific deployed contract.
func NewGatewayEVMValidations(address common.Address, backend bind.ContractBackend) (*GatewayEVMValidations, error) {
	contract, err := bindGatewayEVMValidations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMValidations{GatewayEVMValidationsCaller: GatewayEVMValidationsCaller{contract: contract}, GatewayEVMValidationsTransactor: GatewayEVMValidationsTransactor{contract: contract}, GatewayEVMValidationsFilterer: GatewayEVMValidationsFilterer{contract: contract}}, nil
}

// NewGatewayEVMValidationsCaller creates a new read-only instance of GatewayEVMValidations, bound to a specific deployed contract.
func NewGatewayEVMValidationsCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMValidationsCaller, error) {
	contract, err := bindGatewayEVMValidations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMValidationsCaller{contract: contract}, nil
}

// NewGatewayEVMValidationsTransactor creates a new write-only instance of GatewayEVMValidations, bound to a specific deployed contract.
func NewGatewayEVMValidationsTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMValidationsTransactor, error) {
	contract, err := bindGatewayEVMValidations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMValidationsTransactor{contract: contract}, nil
}

// NewGatewayEVMValidationsFilterer creates a new log filterer instance of GatewayEVMValidations, bound to a specific deployed contract.
func NewGatewayEVMValidationsFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMValidationsFilterer, error) {
	contract, err := bindGatewayEVMValidations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMValidationsFilterer{contract: contract}, nil
}

// bindGatewayEVMValidations binds a generic wrapper to an already deployed contract.
func bindGatewayEVMValidations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMValidationsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMValidations *GatewayEVMValidationsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMValidations.Contract.GatewayEVMValidationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMValidations *GatewayEVMValidationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMValidations.Contract.GatewayEVMValidationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMValidations *GatewayEVMValidationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMValidations.Contract.GatewayEVMValidationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMValidations *GatewayEVMValidationsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMValidations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMValidations *GatewayEVMValidationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMValidations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMValidations *GatewayEVMValidationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMValidations.Contract.contract.Transact(opts, method, params...)
}
