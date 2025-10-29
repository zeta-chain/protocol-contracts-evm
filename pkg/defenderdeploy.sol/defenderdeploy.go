// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package defenderdeploy

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

// DefenderDeployMetaData contains all meta data concerning the DefenderDeploy contract.
var DefenderDeployMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea26469706673582212205865f67a4298dd8cc48d16bcda38169d82fd94d815119b9a9e4132d597636f0c64736f6c634300081a0033",
}

// DefenderDeployABI is the input ABI used to generate the binding from.
// Deprecated: Use DefenderDeployMetaData.ABI instead.
var DefenderDeployABI = DefenderDeployMetaData.ABI

// DefenderDeployBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DefenderDeployMetaData.Bin instead.
var DefenderDeployBin = DefenderDeployMetaData.Bin

// DeployDefenderDeploy deploys a new Ethereum contract, binding an instance of DefenderDeploy to it.
func DeployDefenderDeploy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DefenderDeploy, error) {
	parsed, err := DefenderDeployMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DefenderDeployBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DefenderDeploy{DefenderDeployCaller: DefenderDeployCaller{contract: contract}, DefenderDeployTransactor: DefenderDeployTransactor{contract: contract}, DefenderDeployFilterer: DefenderDeployFilterer{contract: contract}}, nil
}

// DefenderDeploy is an auto generated Go binding around an Ethereum contract.
type DefenderDeploy struct {
	DefenderDeployCaller     // Read-only binding to the contract
	DefenderDeployTransactor // Write-only binding to the contract
	DefenderDeployFilterer   // Log filterer for contract events
}

// DefenderDeployCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefenderDeployCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefenderDeployTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefenderDeployTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefenderDeployFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefenderDeployFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefenderDeploySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefenderDeploySession struct {
	Contract     *DefenderDeploy   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DefenderDeployCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefenderDeployCallerSession struct {
	Contract *DefenderDeployCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DefenderDeployTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefenderDeployTransactorSession struct {
	Contract     *DefenderDeployTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DefenderDeployRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefenderDeployRaw struct {
	Contract *DefenderDeploy // Generic contract binding to access the raw methods on
}

// DefenderDeployCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefenderDeployCallerRaw struct {
	Contract *DefenderDeployCaller // Generic read-only contract binding to access the raw methods on
}

// DefenderDeployTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefenderDeployTransactorRaw struct {
	Contract *DefenderDeployTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefenderDeploy creates a new instance of DefenderDeploy, bound to a specific deployed contract.
func NewDefenderDeploy(address common.Address, backend bind.ContractBackend) (*DefenderDeploy, error) {
	contract, err := bindDefenderDeploy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefenderDeploy{DefenderDeployCaller: DefenderDeployCaller{contract: contract}, DefenderDeployTransactor: DefenderDeployTransactor{contract: contract}, DefenderDeployFilterer: DefenderDeployFilterer{contract: contract}}, nil
}

// NewDefenderDeployCaller creates a new read-only instance of DefenderDeploy, bound to a specific deployed contract.
func NewDefenderDeployCaller(address common.Address, caller bind.ContractCaller) (*DefenderDeployCaller, error) {
	contract, err := bindDefenderDeploy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefenderDeployCaller{contract: contract}, nil
}

// NewDefenderDeployTransactor creates a new write-only instance of DefenderDeploy, bound to a specific deployed contract.
func NewDefenderDeployTransactor(address common.Address, transactor bind.ContractTransactor) (*DefenderDeployTransactor, error) {
	contract, err := bindDefenderDeploy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefenderDeployTransactor{contract: contract}, nil
}

// NewDefenderDeployFilterer creates a new log filterer instance of DefenderDeploy, bound to a specific deployed contract.
func NewDefenderDeployFilterer(address common.Address, filterer bind.ContractFilterer) (*DefenderDeployFilterer, error) {
	contract, err := bindDefenderDeploy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefenderDeployFilterer{contract: contract}, nil
}

// bindDefenderDeploy binds a generic wrapper to an already deployed contract.
func bindDefenderDeploy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DefenderDeployMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefenderDeploy *DefenderDeployRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefenderDeploy.Contract.DefenderDeployCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefenderDeploy *DefenderDeployRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefenderDeploy.Contract.DefenderDeployTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefenderDeploy *DefenderDeployRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefenderDeploy.Contract.DefenderDeployTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefenderDeploy *DefenderDeployCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefenderDeploy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefenderDeploy *DefenderDeployTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefenderDeploy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefenderDeploy *DefenderDeployTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefenderDeploy.Contract.contract.Transact(opts, method, params...)
}
