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

// IZRC20MetaData contains all meta data concerning the IZRC20 contract.
var IZRC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROTOCOL_FLAT_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"newName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSymbol\",\"inputs\":[{\"name\":\"newSymbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawGasFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawGasFeeWithGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
}

// IZRC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IZRC20MetaData.ABI instead.
var IZRC20ABI = IZRC20MetaData.ABI

// IZRC20 is an auto generated Go binding around an Ethereum contract.
type IZRC20 struct {
	IZRC20Caller     // Read-only binding to the contract
	IZRC20Transactor // Write-only binding to the contract
	IZRC20Filterer   // Log filterer for contract events
}

// IZRC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IZRC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZRC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IZRC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZRC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZRC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZRC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZRC20Session struct {
	Contract     *IZRC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IZRC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZRC20CallerSession struct {
	Contract *IZRC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IZRC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZRC20TransactorSession struct {
	Contract     *IZRC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IZRC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IZRC20Raw struct {
	Contract *IZRC20 // Generic contract binding to access the raw methods on
}

// IZRC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZRC20CallerRaw struct {
	Contract *IZRC20Caller // Generic read-only contract binding to access the raw methods on
}

// IZRC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZRC20TransactorRaw struct {
	Contract *IZRC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIZRC20 creates a new instance of IZRC20, bound to a specific deployed contract.
func NewIZRC20(address common.Address, backend bind.ContractBackend) (*IZRC20, error) {
	contract, err := bindIZRC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZRC20{IZRC20Caller: IZRC20Caller{contract: contract}, IZRC20Transactor: IZRC20Transactor{contract: contract}, IZRC20Filterer: IZRC20Filterer{contract: contract}}, nil
}

// NewIZRC20Caller creates a new read-only instance of IZRC20, bound to a specific deployed contract.
func NewIZRC20Caller(address common.Address, caller bind.ContractCaller) (*IZRC20Caller, error) {
	contract, err := bindIZRC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZRC20Caller{contract: contract}, nil
}

// NewIZRC20Transactor creates a new write-only instance of IZRC20, bound to a specific deployed contract.
func NewIZRC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IZRC20Transactor, error) {
	contract, err := bindIZRC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZRC20Transactor{contract: contract}, nil
}

// NewIZRC20Filterer creates a new log filterer instance of IZRC20, bound to a specific deployed contract.
func NewIZRC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IZRC20Filterer, error) {
	contract, err := bindIZRC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZRC20Filterer{contract: contract}, nil
}

// bindIZRC20 binds a generic wrapper to an already deployed contract.
func bindIZRC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZRC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZRC20 *IZRC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZRC20.Contract.IZRC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZRC20 *IZRC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZRC20.Contract.IZRC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZRC20 *IZRC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZRC20.Contract.IZRC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZRC20 *IZRC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZRC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZRC20 *IZRC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZRC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZRC20 *IZRC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZRC20.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_IZRC20 *IZRC20Caller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZRC20.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_IZRC20 *IZRC20Session) CHAINID() (*big.Int, error) {
	return _IZRC20.Contract.CHAINID(&_IZRC20.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_IZRC20 *IZRC20CallerSession) CHAINID() (*big.Int, error) {
	return _IZRC20.Contract.CHAINID(&_IZRC20.CallOpts)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_IZRC20 *IZRC20Caller) GASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZRC20.contract.Call(opts, &out, "GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_IZRC20 *IZRC20Session) GASLIMIT() (*big.Int, error) {
	return _IZRC20.Contract.GASLIMIT(&_IZRC20.CallOpts)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_IZRC20 *IZRC20CallerSession) GASLIMIT() (*big.Int, error) {
	return _IZRC20.Contract.GASLIMIT(&_IZRC20.CallOpts)
}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_IZRC20 *IZRC20Caller) PROTOCOLFLATFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZRC20.contract.Call(opts, &out, "PROTOCOL_FLAT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_IZRC20 *IZRC20Session) PROTOCOLFLATFEE() (*big.Int, error) {
	return _IZRC20.Contract.PROTOCOLFLATFEE(&_IZRC20.CallOpts)
}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_IZRC20 *IZRC20CallerSession) PROTOCOLFLATFEE() (*big.Int, error) {
	return _IZRC20.Contract.PROTOCOLFLATFEE(&_IZRC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZRC20 *IZRC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IZRC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZRC20 *IZRC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IZRC20.Contract.Allowance(&_IZRC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZRC20 *IZRC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IZRC20.Contract.Allowance(&_IZRC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZRC20 *IZRC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IZRC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZRC20 *IZRC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IZRC20.Contract.BalanceOf(&_IZRC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZRC20 *IZRC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IZRC20.Contract.BalanceOf(&_IZRC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZRC20 *IZRC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZRC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZRC20 *IZRC20Session) TotalSupply() (*big.Int, error) {
	return _IZRC20.Contract.TotalSupply(&_IZRC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZRC20 *IZRC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IZRC20.Contract.TotalSupply(&_IZRC20.CallOpts)
}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_IZRC20 *IZRC20Caller) WithdrawGasFee(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _IZRC20.contract.Call(opts, &out, "withdrawGasFee")

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
func (_IZRC20 *IZRC20Session) WithdrawGasFee() (common.Address, *big.Int, error) {
	return _IZRC20.Contract.WithdrawGasFee(&_IZRC20.CallOpts)
}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_IZRC20 *IZRC20CallerSession) WithdrawGasFee() (common.Address, *big.Int, error) {
	return _IZRC20.Contract.WithdrawGasFee(&_IZRC20.CallOpts)
}

// WithdrawGasFeeWithGasLimit is a free data retrieval call binding the contract method 0xfc5fecd5.
//
// Solidity: function withdrawGasFeeWithGasLimit(uint256 gasLimit) view returns(address, uint256)
func (_IZRC20 *IZRC20Caller) WithdrawGasFeeWithGasLimit(opts *bind.CallOpts, gasLimit *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _IZRC20.contract.Call(opts, &out, "withdrawGasFeeWithGasLimit", gasLimit)

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
func (_IZRC20 *IZRC20Session) WithdrawGasFeeWithGasLimit(gasLimit *big.Int) (common.Address, *big.Int, error) {
	return _IZRC20.Contract.WithdrawGasFeeWithGasLimit(&_IZRC20.CallOpts, gasLimit)
}

// WithdrawGasFeeWithGasLimit is a free data retrieval call binding the contract method 0xfc5fecd5.
//
// Solidity: function withdrawGasFeeWithGasLimit(uint256 gasLimit) view returns(address, uint256)
func (_IZRC20 *IZRC20CallerSession) WithdrawGasFeeWithGasLimit(gasLimit *big.Int) (common.Address, *big.Int, error) {
	return _IZRC20.Contract.WithdrawGasFeeWithGasLimit(&_IZRC20.CallOpts, gasLimit)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.Contract.Approve(&_IZRC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.Contract.Approve(&_IZRC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_IZRC20 *IZRC20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_IZRC20 *IZRC20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.Contract.Burn(&_IZRC20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_IZRC20 *IZRC20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.Contract.Burn(&_IZRC20.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20Transactor) Deposit(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.contract.Transact(opts, "deposit", to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20Session) Deposit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.Contract.Deposit(&_IZRC20.TransactOpts, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20TransactorSession) Deposit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.Contract.Deposit(&_IZRC20.TransactOpts, to, amount)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_IZRC20 *IZRC20Transactor) SetName(opts *bind.TransactOpts, newName string) (*types.Transaction, error) {
	return _IZRC20.contract.Transact(opts, "setName", newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_IZRC20 *IZRC20Session) SetName(newName string) (*types.Transaction, error) {
	return _IZRC20.Contract.SetName(&_IZRC20.TransactOpts, newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_IZRC20 *IZRC20TransactorSession) SetName(newName string) (*types.Transaction, error) {
	return _IZRC20.Contract.SetName(&_IZRC20.TransactOpts, newName)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_IZRC20 *IZRC20Transactor) SetSymbol(opts *bind.TransactOpts, newSymbol string) (*types.Transaction, error) {
	return _IZRC20.contract.Transact(opts, "setSymbol", newSymbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_IZRC20 *IZRC20Session) SetSymbol(newSymbol string) (*types.Transaction, error) {
	return _IZRC20.Contract.SetSymbol(&_IZRC20.TransactOpts, newSymbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_IZRC20 *IZRC20TransactorSession) SetSymbol(newSymbol string) (*types.Transaction, error) {
	return _IZRC20.Contract.SetSymbol(&_IZRC20.TransactOpts, newSymbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.Contract.Transfer(&_IZRC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.Contract.Transfer(&_IZRC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.Contract.TransferFrom(&_IZRC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.Contract.TransferFrom(&_IZRC20.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20Transactor) Withdraw(opts *bind.TransactOpts, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20Session) Withdraw(to []byte, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.Contract.Withdraw(&_IZRC20.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_IZRC20 *IZRC20TransactorSession) Withdraw(to []byte, amount *big.Int) (*types.Transaction, error) {
	return _IZRC20.Contract.Withdraw(&_IZRC20.TransactOpts, to, amount)
}
