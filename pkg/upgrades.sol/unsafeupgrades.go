// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package upgrades

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

// UnsafeUpgradesMetaData contains all meta data concerning the UnsafeUpgrades contract.
var UnsafeUpgradesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122022aae055d0ed5f7a6246fce7acdb58528c976f867a57654c7a2cc577450d54ec64736f6c634300081a0033",
}

// UnsafeUpgradesABI is the input ABI used to generate the binding from.
// Deprecated: Use UnsafeUpgradesMetaData.ABI instead.
var UnsafeUpgradesABI = UnsafeUpgradesMetaData.ABI

// UnsafeUpgradesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UnsafeUpgradesMetaData.Bin instead.
var UnsafeUpgradesBin = UnsafeUpgradesMetaData.Bin

// DeployUnsafeUpgrades deploys a new Ethereum contract, binding an instance of UnsafeUpgrades to it.
func DeployUnsafeUpgrades(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UnsafeUpgrades, error) {
	parsed, err := UnsafeUpgradesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UnsafeUpgradesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UnsafeUpgrades{UnsafeUpgradesCaller: UnsafeUpgradesCaller{contract: contract}, UnsafeUpgradesTransactor: UnsafeUpgradesTransactor{contract: contract}, UnsafeUpgradesFilterer: UnsafeUpgradesFilterer{contract: contract}}, nil
}

// UnsafeUpgrades is an auto generated Go binding around an Ethereum contract.
type UnsafeUpgrades struct {
	UnsafeUpgradesCaller     // Read-only binding to the contract
	UnsafeUpgradesTransactor // Write-only binding to the contract
	UnsafeUpgradesFilterer   // Log filterer for contract events
}

// UnsafeUpgradesCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnsafeUpgradesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnsafeUpgradesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnsafeUpgradesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnsafeUpgradesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnsafeUpgradesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnsafeUpgradesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnsafeUpgradesSession struct {
	Contract     *UnsafeUpgrades   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnsafeUpgradesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnsafeUpgradesCallerSession struct {
	Contract *UnsafeUpgradesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UnsafeUpgradesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnsafeUpgradesTransactorSession struct {
	Contract     *UnsafeUpgradesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UnsafeUpgradesRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnsafeUpgradesRaw struct {
	Contract *UnsafeUpgrades // Generic contract binding to access the raw methods on
}

// UnsafeUpgradesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnsafeUpgradesCallerRaw struct {
	Contract *UnsafeUpgradesCaller // Generic read-only contract binding to access the raw methods on
}

// UnsafeUpgradesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnsafeUpgradesTransactorRaw struct {
	Contract *UnsafeUpgradesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnsafeUpgrades creates a new instance of UnsafeUpgrades, bound to a specific deployed contract.
func NewUnsafeUpgrades(address common.Address, backend bind.ContractBackend) (*UnsafeUpgrades, error) {
	contract, err := bindUnsafeUpgrades(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnsafeUpgrades{UnsafeUpgradesCaller: UnsafeUpgradesCaller{contract: contract}, UnsafeUpgradesTransactor: UnsafeUpgradesTransactor{contract: contract}, UnsafeUpgradesFilterer: UnsafeUpgradesFilterer{contract: contract}}, nil
}

// NewUnsafeUpgradesCaller creates a new read-only instance of UnsafeUpgrades, bound to a specific deployed contract.
func NewUnsafeUpgradesCaller(address common.Address, caller bind.ContractCaller) (*UnsafeUpgradesCaller, error) {
	contract, err := bindUnsafeUpgrades(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnsafeUpgradesCaller{contract: contract}, nil
}

// NewUnsafeUpgradesTransactor creates a new write-only instance of UnsafeUpgrades, bound to a specific deployed contract.
func NewUnsafeUpgradesTransactor(address common.Address, transactor bind.ContractTransactor) (*UnsafeUpgradesTransactor, error) {
	contract, err := bindUnsafeUpgrades(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnsafeUpgradesTransactor{contract: contract}, nil
}

// NewUnsafeUpgradesFilterer creates a new log filterer instance of UnsafeUpgrades, bound to a specific deployed contract.
func NewUnsafeUpgradesFilterer(address common.Address, filterer bind.ContractFilterer) (*UnsafeUpgradesFilterer, error) {
	contract, err := bindUnsafeUpgrades(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnsafeUpgradesFilterer{contract: contract}, nil
}

// bindUnsafeUpgrades binds a generic wrapper to an already deployed contract.
func bindUnsafeUpgrades(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UnsafeUpgradesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnsafeUpgrades *UnsafeUpgradesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnsafeUpgrades.Contract.UnsafeUpgradesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnsafeUpgrades *UnsafeUpgradesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnsafeUpgrades.Contract.UnsafeUpgradesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnsafeUpgrades *UnsafeUpgradesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnsafeUpgrades.Contract.UnsafeUpgradesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnsafeUpgrades *UnsafeUpgradesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnsafeUpgrades.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnsafeUpgrades *UnsafeUpgradesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnsafeUpgrades.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnsafeUpgrades *UnsafeUpgradesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnsafeUpgrades.Contract.contract.Transact(opts, method, params...)
}
