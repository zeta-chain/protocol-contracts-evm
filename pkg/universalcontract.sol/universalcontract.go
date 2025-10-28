// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package universalcontract

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

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Sender    []byte
	SenderEVM common.Address
	ChainID   *big.Int
}

// UniversalContractMetaData contains all meta data concerning the UniversalContract contract.
var UniversalContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayZEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICoreRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
}

// UniversalContractABI is the input ABI used to generate the binding from.
// Deprecated: Use UniversalContractMetaData.ABI instead.
var UniversalContractABI = UniversalContractMetaData.ABI

// UniversalContract is an auto generated Go binding around an Ethereum contract.
type UniversalContract struct {
	UniversalContractCaller     // Read-only binding to the contract
	UniversalContractTransactor // Write-only binding to the contract
	UniversalContractFilterer   // Log filterer for contract events
}

// UniversalContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniversalContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniversalContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniversalContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniversalContractSession struct {
	Contract     *UniversalContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UniversalContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniversalContractCallerSession struct {
	Contract *UniversalContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UniversalContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniversalContractTransactorSession struct {
	Contract     *UniversalContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UniversalContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniversalContractRaw struct {
	Contract *UniversalContract // Generic contract binding to access the raw methods on
}

// UniversalContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniversalContractCallerRaw struct {
	Contract *UniversalContractCaller // Generic read-only contract binding to access the raw methods on
}

// UniversalContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniversalContractTransactorRaw struct {
	Contract *UniversalContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniversalContract creates a new instance of UniversalContract, bound to a specific deployed contract.
func NewUniversalContract(address common.Address, backend bind.ContractBackend) (*UniversalContract, error) {
	contract, err := bindUniversalContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniversalContract{UniversalContractCaller: UniversalContractCaller{contract: contract}, UniversalContractTransactor: UniversalContractTransactor{contract: contract}, UniversalContractFilterer: UniversalContractFilterer{contract: contract}}, nil
}

// NewUniversalContractCaller creates a new read-only instance of UniversalContract, bound to a specific deployed contract.
func NewUniversalContractCaller(address common.Address, caller bind.ContractCaller) (*UniversalContractCaller, error) {
	contract, err := bindUniversalContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalContractCaller{contract: contract}, nil
}

// NewUniversalContractTransactor creates a new write-only instance of UniversalContract, bound to a specific deployed contract.
func NewUniversalContractTransactor(address common.Address, transactor bind.ContractTransactor) (*UniversalContractTransactor, error) {
	contract, err := bindUniversalContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalContractTransactor{contract: contract}, nil
}

// NewUniversalContractFilterer creates a new log filterer instance of UniversalContract, bound to a specific deployed contract.
func NewUniversalContractFilterer(address common.Address, filterer bind.ContractFilterer) (*UniversalContractFilterer, error) {
	contract, err := bindUniversalContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniversalContractFilterer{contract: contract}, nil
}

// bindUniversalContract binds a generic wrapper to an already deployed contract.
func bindUniversalContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniversalContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalContract *UniversalContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalContract.Contract.UniversalContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalContract *UniversalContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalContract.Contract.UniversalContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalContract *UniversalContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalContract.Contract.UniversalContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalContract *UniversalContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalContract *UniversalContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalContract *UniversalContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalContract.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_UniversalContract *UniversalContractCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalContract.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_UniversalContract *UniversalContractSession) Gateway() (common.Address, error) {
	return _UniversalContract.Contract.Gateway(&_UniversalContract.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_UniversalContract *UniversalContractCallerSession) Gateway() (common.Address, error) {
	return _UniversalContract.Contract.Gateway(&_UniversalContract.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_UniversalContract *UniversalContractCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalContract.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_UniversalContract *UniversalContractSession) Registry() (common.Address, error) {
	return _UniversalContract.Contract.Registry(&_UniversalContract.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_UniversalContract *UniversalContractCallerSession) Registry() (common.Address, error) {
	return _UniversalContract.Contract.Registry(&_UniversalContract.CallOpts)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_UniversalContract *UniversalContractTransactor) OnCall(opts *bind.TransactOpts, context MessageContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _UniversalContract.contract.Transact(opts, "onCall", context, zrc20, amount, message)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_UniversalContract *UniversalContractSession) OnCall(context MessageContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _UniversalContract.Contract.OnCall(&_UniversalContract.TransactOpts, context, zrc20, amount, message)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) context, address zrc20, uint256 amount, bytes message) returns()
func (_UniversalContract *UniversalContractTransactorSession) OnCall(context MessageContext, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _UniversalContract.Contract.OnCall(&_UniversalContract.TransactOpts, context, zrc20, amount, message)
}
