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

// UpgradesMetaData contains all meta data concerning the Upgrades contract.
var UpgradesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea264697066735822122094b1a108024a01ad7a868f482b504e48145cd66cf3821a5fce281a3ab24a5b7664736f6c634300081a0033",
}

// UpgradesABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradesMetaData.ABI instead.
var UpgradesABI = UpgradesMetaData.ABI

// UpgradesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UpgradesMetaData.Bin instead.
var UpgradesBin = UpgradesMetaData.Bin

// DeployUpgrades deploys a new Ethereum contract, binding an instance of Upgrades to it.
func DeployUpgrades(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Upgrades, error) {
	parsed, err := UpgradesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UpgradesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Upgrades{UpgradesCaller: UpgradesCaller{contract: contract}, UpgradesTransactor: UpgradesTransactor{contract: contract}, UpgradesFilterer: UpgradesFilterer{contract: contract}}, nil
}

// Upgrades is an auto generated Go binding around an Ethereum contract.
type Upgrades struct {
	UpgradesCaller     // Read-only binding to the contract
	UpgradesTransactor // Write-only binding to the contract
	UpgradesFilterer   // Log filterer for contract events
}

// UpgradesCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradesSession struct {
	Contract     *Upgrades         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradesCallerSession struct {
	Contract *UpgradesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// UpgradesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradesTransactorSession struct {
	Contract     *UpgradesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UpgradesRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradesRaw struct {
	Contract *Upgrades // Generic contract binding to access the raw methods on
}

// UpgradesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradesCallerRaw struct {
	Contract *UpgradesCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradesTransactorRaw struct {
	Contract *UpgradesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgrades creates a new instance of Upgrades, bound to a specific deployed contract.
func NewUpgrades(address common.Address, backend bind.ContractBackend) (*Upgrades, error) {
	contract, err := bindUpgrades(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Upgrades{UpgradesCaller: UpgradesCaller{contract: contract}, UpgradesTransactor: UpgradesTransactor{contract: contract}, UpgradesFilterer: UpgradesFilterer{contract: contract}}, nil
}

// NewUpgradesCaller creates a new read-only instance of Upgrades, bound to a specific deployed contract.
func NewUpgradesCaller(address common.Address, caller bind.ContractCaller) (*UpgradesCaller, error) {
	contract, err := bindUpgrades(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradesCaller{contract: contract}, nil
}

// NewUpgradesTransactor creates a new write-only instance of Upgrades, bound to a specific deployed contract.
func NewUpgradesTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradesTransactor, error) {
	contract, err := bindUpgrades(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradesTransactor{contract: contract}, nil
}

// NewUpgradesFilterer creates a new log filterer instance of Upgrades, bound to a specific deployed contract.
func NewUpgradesFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradesFilterer, error) {
	contract, err := bindUpgrades(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradesFilterer{contract: contract}, nil
}

// bindUpgrades binds a generic wrapper to an already deployed contract.
func bindUpgrades(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpgradesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Upgrades *UpgradesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Upgrades.Contract.UpgradesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Upgrades *UpgradesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Upgrades.Contract.UpgradesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Upgrades *UpgradesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Upgrades.Contract.UpgradesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Upgrades *UpgradesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Upgrades.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Upgrades *UpgradesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Upgrades.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Upgrades *UpgradesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Upgrades.Contract.contract.Transact(opts, method, params...)
}
