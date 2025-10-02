// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package defender

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

// DefenderMetaData contains all meta data concerning the Defender contract.
var DefenderMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122038a7bb1c9ed76aaf600e4208ed9890b102c2ca44168a4400e14317b321922f3b64736f6c634300081a0033",
}

// DefenderABI is the input ABI used to generate the binding from.
// Deprecated: Use DefenderMetaData.ABI instead.
var DefenderABI = DefenderMetaData.ABI

// DefenderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DefenderMetaData.Bin instead.
var DefenderBin = DefenderMetaData.Bin

// DeployDefender deploys a new Ethereum contract, binding an instance of Defender to it.
func DeployDefender(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Defender, error) {
	parsed, err := DefenderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DefenderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Defender{DefenderCaller: DefenderCaller{contract: contract}, DefenderTransactor: DefenderTransactor{contract: contract}, DefenderFilterer: DefenderFilterer{contract: contract}}, nil
}

// Defender is an auto generated Go binding around an Ethereum contract.
type Defender struct {
	DefenderCaller     // Read-only binding to the contract
	DefenderTransactor // Write-only binding to the contract
	DefenderFilterer   // Log filterer for contract events
}

// DefenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefenderSession struct {
	Contract     *Defender         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DefenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefenderCallerSession struct {
	Contract *DefenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DefenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefenderTransactorSession struct {
	Contract     *DefenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DefenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefenderRaw struct {
	Contract *Defender // Generic contract binding to access the raw methods on
}

// DefenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefenderCallerRaw struct {
	Contract *DefenderCaller // Generic read-only contract binding to access the raw methods on
}

// DefenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefenderTransactorRaw struct {
	Contract *DefenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefender creates a new instance of Defender, bound to a specific deployed contract.
func NewDefender(address common.Address, backend bind.ContractBackend) (*Defender, error) {
	contract, err := bindDefender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Defender{DefenderCaller: DefenderCaller{contract: contract}, DefenderTransactor: DefenderTransactor{contract: contract}, DefenderFilterer: DefenderFilterer{contract: contract}}, nil
}

// NewDefenderCaller creates a new read-only instance of Defender, bound to a specific deployed contract.
func NewDefenderCaller(address common.Address, caller bind.ContractCaller) (*DefenderCaller, error) {
	contract, err := bindDefender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefenderCaller{contract: contract}, nil
}

// NewDefenderTransactor creates a new write-only instance of Defender, bound to a specific deployed contract.
func NewDefenderTransactor(address common.Address, transactor bind.ContractTransactor) (*DefenderTransactor, error) {
	contract, err := bindDefender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefenderTransactor{contract: contract}, nil
}

// NewDefenderFilterer creates a new log filterer instance of Defender, bound to a specific deployed contract.
func NewDefenderFilterer(address common.Address, filterer bind.ContractFilterer) (*DefenderFilterer, error) {
	contract, err := bindDefender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefenderFilterer{contract: contract}, nil
}

// bindDefender binds a generic wrapper to an already deployed contract.
func bindDefender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DefenderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Defender *DefenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Defender.Contract.DefenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Defender *DefenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Defender.Contract.DefenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Defender *DefenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Defender.Contract.DefenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Defender *DefenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Defender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Defender *DefenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Defender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Defender *DefenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Defender.Contract.contract.Transact(opts, method, params...)
}
