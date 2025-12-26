// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izrc20

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

// IZRC20MetadataMetaData contains all meta data concerning the IZRC20Metadata contract.
var IZRC20MetadataMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROTOCOL_FLAT_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"newName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSymbol\",\"inputs\":[{\"name\":\"newSymbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawGasFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawGasFeeWithGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
}

// IZRC20MetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use IZRC20MetadataMetaData.ABI instead.
var IZRC20MetadataABI = IZRC20MetadataMetaData.ABI

// IZRC20Metadata is an auto generated Go binding around an Ethereum contract.
type IZRC20Metadata struct {
	IZRC20MetadataCaller     // Read-only binding to the contract
	IZRC20MetadataTransactor // Write-only binding to the contract
	IZRC20MetadataFilterer   // Log filterer for contract events
}

// IZRC20MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZRC20MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZRC20MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZRC20MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZRC20MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZRC20MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZRC20MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZRC20MetadataSession struct {
	Contract     *IZRC20Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IZRC20MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZRC20MetadataCallerSession struct {
	Contract *IZRC20MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IZRC20MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZRC20MetadataTransactorSession struct {
	Contract     *IZRC20MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IZRC20MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZRC20MetadataRaw struct {
	Contract *IZRC20Metadata // Generic contract binding to access the raw methods on
}

// IZRC20MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZRC20MetadataCallerRaw struct {
	Contract *IZRC20MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IZRC20MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZRC20MetadataTransactorRaw struct {
	Contract *IZRC20MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZRC20Metadata creates a new instance of IZRC20Metadata, bound to a specific deployed contract.
func NewIZRC20Metadata(address common.Address, backend bind.ContractBackend) (*IZRC20Metadata, error) {
	contract, err := bindIZRC20Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZRC20Metadata{IZRC20MetadataCaller: IZRC20MetadataCaller{contract: contract}, IZRC20MetadataTransactor: IZRC20MetadataTransactor{contract: contract}, IZRC20MetadataFilterer: IZRC20MetadataFilterer{contract: contract}}, nil
}

// NewIZRC20MetadataCaller creates a new read-only instance of IZRC20Metadata, bound to a specific deployed contract.
func NewIZRC20MetadataCaller(address common.Address, caller bind.ContractCaller) (*IZRC20MetadataCaller, error) {
	contract, err := bindIZRC20Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZRC20MetadataCaller{contract: contract}, nil
}

// NewIZRC20MetadataTransactor creates a new write-only instance of IZRC20Metadata, bound to a specific deployed contract.
func NewIZRC20MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IZRC20MetadataTransactor, error) {
	contract, err := bindIZRC20Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZRC20MetadataTransactor{contract: contract}, nil
}

// NewIZRC20MetadataFilterer creates a new log filterer instance of IZRC20Metadata, bound to a specific deployed contract.
func NewIZRC20MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IZRC20MetadataFilterer, error) {
	contract, err := bindIZRC20Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZRC20MetadataFilterer{contract: contract}, nil
}

// bindIZRC20Metadata binds a generic wrapper to an already deployed contract.
func bindIZRC20Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZRC20MetadataMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZRC20Metadata *IZRC20MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZRC20Metadata.Contract.IZRC20MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZRC20Metadata *IZRC20MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.IZRC20MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZRC20Metadata *IZRC20MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.IZRC20MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZRC20Metadata *IZRC20MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZRC20Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZRC20Metadata *IZRC20MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZRC20Metadata *IZRC20MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataCaller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZRC20Metadata.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataSession) CHAINID() (*big.Int, error) {
	return _IZRC20Metadata.Contract.CHAINID(&_IZRC20Metadata.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataCallerSession) CHAINID() (*big.Int, error) {
	return _IZRC20Metadata.Contract.CHAINID(&_IZRC20Metadata.CallOpts)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataCaller) GASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZRC20Metadata.contract.Call(opts, &out, "GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataSession) GASLIMIT() (*big.Int, error) {
	return _IZRC20Metadata.Contract.GASLIMIT(&_IZRC20Metadata.CallOpts)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataCallerSession) GASLIMIT() (*big.Int, error) {
	return _IZRC20Metadata.Contract.GASLIMIT(&_IZRC20Metadata.CallOpts)
}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataCaller) PROTOCOLFLATFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZRC20Metadata.contract.Call(opts, &out, "PROTOCOL_FLAT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataSession) PROTOCOLFLATFEE() (*big.Int, error) {
	return _IZRC20Metadata.Contract.PROTOCOLFLATFEE(&_IZRC20Metadata.CallOpts)
}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataCallerSession) PROTOCOLFLATFEE() (*big.Int, error) {
	return _IZRC20Metadata.Contract.PROTOCOLFLATFEE(&_IZRC20Metadata.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IZRC20Metadata.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IZRC20Metadata.Contract.Allowance(&_IZRC20Metadata.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IZRC20Metadata.Contract.Allowance(&_IZRC20Metadata.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IZRC20Metadata.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IZRC20Metadata.Contract.BalanceOf(&_IZRC20Metadata.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IZRC20Metadata.Contract.BalanceOf(&_IZRC20Metadata.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IZRC20Metadata *IZRC20MetadataCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IZRC20Metadata.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IZRC20Metadata *IZRC20MetadataSession) Decimals() (uint8, error) {
	return _IZRC20Metadata.Contract.Decimals(&_IZRC20Metadata.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IZRC20Metadata *IZRC20MetadataCallerSession) Decimals() (uint8, error) {
	return _IZRC20Metadata.Contract.Decimals(&_IZRC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IZRC20Metadata *IZRC20MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IZRC20Metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IZRC20Metadata *IZRC20MetadataSession) Name() (string, error) {
	return _IZRC20Metadata.Contract.Name(&_IZRC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IZRC20Metadata *IZRC20MetadataCallerSession) Name() (string, error) {
	return _IZRC20Metadata.Contract.Name(&_IZRC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IZRC20Metadata *IZRC20MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IZRC20Metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IZRC20Metadata *IZRC20MetadataSession) Symbol() (string, error) {
	return _IZRC20Metadata.Contract.Symbol(&_IZRC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IZRC20Metadata *IZRC20MetadataCallerSession) Symbol() (string, error) {
	return _IZRC20Metadata.Contract.Symbol(&_IZRC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZRC20Metadata.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataSession) TotalSupply() (*big.Int, error) {
	return _IZRC20Metadata.Contract.TotalSupply(&_IZRC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZRC20Metadata *IZRC20MetadataCallerSession) TotalSupply() (*big.Int, error) {
	return _IZRC20Metadata.Contract.TotalSupply(&_IZRC20Metadata.CallOpts)
}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_IZRC20Metadata *IZRC20MetadataCaller) WithdrawGasFee(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _IZRC20Metadata.contract.Call(opts, &out, "withdrawGasFee")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_IZRC20Metadata *IZRC20MetadataSession) WithdrawGasFee() (common.Address, *big.Int, error) {
	return _IZRC20Metadata.Contract.WithdrawGasFee(&_IZRC20Metadata.CallOpts)
}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_IZRC20Metadata *IZRC20MetadataCallerSession) WithdrawGasFee() (common.Address, *big.Int, error) {
	return _IZRC20Metadata.Contract.WithdrawGasFee(&_IZRC20Metadata.CallOpts)
}

// WithdrawGasFeeWithGasLimit is a free data retrieval call binding the contract method 0xfc5fecd5.
//
// Solidity: function withdrawGasFeeWithGasLimit(uint256 gasLimit) view returns(address, uint256)
func (_IZRC20Metadata *IZRC20MetadataCaller) WithdrawGasFeeWithGasLimit(opts *bind.CallOpts, gasLimit *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _IZRC20Metadata.contract.Call(opts, &out, "withdrawGasFeeWithGasLimit", gasLimit)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// WithdrawGasFeeWithGasLimit is a free data retrieval call binding the contract method 0xfc5fecd5.
//
// Solidity: function withdrawGasFeeWithGasLimit(uint256 gasLimit) view returns(address, uint256)
func (_IZRC20Metadata *IZRC20MetadataSession) WithdrawGasFeeWithGasLimit(gasLimit *big.Int) (common.Address, *big.Int, error) {
	return _IZRC20Metadata.Contract.WithdrawGasFeeWithGasLimit(&_IZRC20Metadata.CallOpts, gasLimit)
}

// WithdrawGasFeeWithGasLimit is a free data retrieval call binding the contract method 0xfc5fecd5.
//
// Solidity: function withdrawGasFeeWithGasLimit(uint256 gasLimit) view returns(address, uint256)
func (_IZRC20Metadata *IZRC20MetadataCallerSession) WithdrawGasFeeWithGasLimit(gasLimit *big.Int) (common.Address, *big.Int, error) {
	return _IZRC20Metadata.Contract.WithdrawGasFeeWithGasLimit(&_IZRC20Metadata.CallOpts, gasLimit)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.Approve(&_IZRC20Metadata.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.Approve(&_IZRC20Metadata.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.Burn(&_IZRC20Metadata.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.Burn(&_IZRC20Metadata.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataTransactor) Deposit(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.contract.Transact(opts, "deposit", to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataSession) Deposit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.Deposit(&_IZRC20Metadata.TransactOpts, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataTransactorSession) Deposit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.Deposit(&_IZRC20Metadata.TransactOpts, to, amount)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_IZRC20Metadata *IZRC20MetadataTransactor) SetName(opts *bind.TransactOpts, newName string) (*types.Transaction, error) {
	return _IZRC20Metadata.contract.Transact(opts, "setName", newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_IZRC20Metadata *IZRC20MetadataSession) SetName(newName string) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.SetName(&_IZRC20Metadata.TransactOpts, newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_IZRC20Metadata *IZRC20MetadataTransactorSession) SetName(newName string) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.SetName(&_IZRC20Metadata.TransactOpts, newName)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_IZRC20Metadata *IZRC20MetadataTransactor) SetSymbol(opts *bind.TransactOpts, newSymbol string) (*types.Transaction, error) {
	return _IZRC20Metadata.contract.Transact(opts, "setSymbol", newSymbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_IZRC20Metadata *IZRC20MetadataSession) SetSymbol(newSymbol string) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.SetSymbol(&_IZRC20Metadata.TransactOpts, newSymbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_IZRC20Metadata *IZRC20MetadataTransactorSession) SetSymbol(newSymbol string) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.SetSymbol(&_IZRC20Metadata.TransactOpts, newSymbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.Transfer(&_IZRC20Metadata.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.Transfer(&_IZRC20Metadata.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.TransferFrom(&_IZRC20Metadata.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.TransferFrom(&_IZRC20Metadata.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataTransactor) Withdraw(opts *bind.TransactOpts, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataSession) Withdraw(to []byte, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.Withdraw(&_IZRC20Metadata.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_IZRC20Metadata *IZRC20MetadataTransactorSession) Withdraw(to []byte, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20Metadata.Contract.Withdraw(&_IZRC20Metadata.TransactOpts, to, amount)
}
