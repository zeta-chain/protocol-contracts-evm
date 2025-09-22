// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc1363

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

// IERC1363MetaData contains all meta data concerning the IERC1363 contract.
var IERC1363MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approveAndCall\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approveAndCall\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferAndCall\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferAndCall\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFromAndCall\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFromAndCall\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// IERC1363ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1363MetaData.ABI instead.
var IERC1363ABI = IERC1363MetaData.ABI

// IERC1363 is an auto generated Go binding around an Ethereum contract.
type IERC1363 struct {
	IERC1363Caller     // Read-only binding to the contract
	IERC1363Transactor // Write-only binding to the contract
	IERC1363Filterer   // Log filterer for contract events
}

// IERC1363Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1363Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1363Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1363Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1363Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1363Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1363Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1363Session struct {
	Contract     *IERC1363         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC1363CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1363CallerSession struct {
	Contract *IERC1363Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IERC1363TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1363TransactorSession struct {
	Contract     *IERC1363Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IERC1363Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1363Raw struct {
	Contract *IERC1363 // Generic contract binding to access the raw methods on
}

// IERC1363CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1363CallerRaw struct {
	Contract *IERC1363Caller // Generic read-only contract binding to access the raw methods on
}

// IERC1363TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1363TransactorRaw struct {
	Contract *IERC1363Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1363 creates a new instance of IERC1363, bound to a specific deployed contract.
func NewIERC1363(address common.Address, backend bind.ContractBackend) (*IERC1363, error) {
	contract, err := bindIERC1363(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1363{IERC1363Caller: IERC1363Caller{contract: contract}, IERC1363Transactor: IERC1363Transactor{contract: contract}, IERC1363Filterer: IERC1363Filterer{contract: contract}}, nil
}

// NewIERC1363Caller creates a new read-only instance of IERC1363, bound to a specific deployed contract.
func NewIERC1363Caller(address common.Address, caller bind.ContractCaller) (*IERC1363Caller, error) {
	contract, err := bindIERC1363(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1363Caller{contract: contract}, nil
}

// NewIERC1363Transactor creates a new write-only instance of IERC1363, bound to a specific deployed contract.
func NewIERC1363Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC1363Transactor, error) {
	contract, err := bindIERC1363(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1363Transactor{contract: contract}, nil
}

// NewIERC1363Filterer creates a new log filterer instance of IERC1363, bound to a specific deployed contract.
func NewIERC1363Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC1363Filterer, error) {
	contract, err := bindIERC1363(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1363Filterer{contract: contract}, nil
}

// bindIERC1363 binds a generic wrapper to an already deployed contract.
func bindIERC1363(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC1363MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1363 *IERC1363Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1363.Contract.IERC1363Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1363 *IERC1363Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1363.Contract.IERC1363Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1363 *IERC1363Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1363.Contract.IERC1363Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1363 *IERC1363CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1363.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1363 *IERC1363TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1363.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1363 *IERC1363TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1363.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC1363 *IERC1363Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC1363.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC1363 *IERC1363Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC1363.Contract.Allowance(&_IERC1363.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC1363 *IERC1363CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC1363.Contract.Allowance(&_IERC1363.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC1363 *IERC1363Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC1363.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC1363 *IERC1363Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC1363.Contract.BalanceOf(&_IERC1363.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC1363 *IERC1363CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC1363.Contract.BalanceOf(&_IERC1363.CallOpts, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1363 *IERC1363Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC1363.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1363 *IERC1363Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC1363.Contract.SupportsInterface(&_IERC1363.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1363 *IERC1363CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC1363.Contract.SupportsInterface(&_IERC1363.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC1363 *IERC1363Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC1363.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC1363 *IERC1363Session) TotalSupply() (*big.Int, error) {
	return _IERC1363.Contract.TotalSupply(&_IERC1363.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC1363 *IERC1363CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC1363.Contract.TotalSupply(&_IERC1363.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.Approve(&_IERC1363.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.Approve(&_IERC1363.TransactOpts, spender, value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0x3177029f.
//
// Solidity: function approveAndCall(address spender, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Transactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "approveAndCall", spender, value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0x3177029f.
//
// Solidity: function approveAndCall(address spender, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Session) ApproveAndCall(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.ApproveAndCall(&_IERC1363.TransactOpts, spender, value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0x3177029f.
//
// Solidity: function approveAndCall(address spender, uint256 value) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) ApproveAndCall(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.ApproveAndCall(&_IERC1363.TransactOpts, spender, value)
}

// ApproveAndCall0 is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363Transactor) ApproveAndCall0(opts *bind.TransactOpts, spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "approveAndCall0", spender, value, data)
}

// ApproveAndCall0 is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363Session) ApproveAndCall0(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.Contract.ApproveAndCall0(&_IERC1363.TransactOpts, spender, value, data)
}

// ApproveAndCall0 is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) ApproveAndCall0(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.Contract.ApproveAndCall0(&_IERC1363.TransactOpts, spender, value, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.Transfer(&_IERC1363.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.Transfer(&_IERC1363.TransactOpts, to, value)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x1296ee62.
//
// Solidity: function transferAndCall(address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Transactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "transferAndCall", to, value)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x1296ee62.
//
// Solidity: function transferAndCall(address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Session) TransferAndCall(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferAndCall(&_IERC1363.TransactOpts, to, value)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x1296ee62.
//
// Solidity: function transferAndCall(address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) TransferAndCall(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferAndCall(&_IERC1363.TransactOpts, to, value)
}

// TransferAndCall0 is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363Transactor) TransferAndCall0(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "transferAndCall0", to, value, data)
}

// TransferAndCall0 is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363Session) TransferAndCall0(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferAndCall0(&_IERC1363.TransactOpts, to, value, data)
}

// TransferAndCall0 is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) TransferAndCall0(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferAndCall0(&_IERC1363.TransactOpts, to, value, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferFrom(&_IERC1363.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferFrom(&_IERC1363.TransactOpts, from, to, value)
}

// TransferFromAndCall is a paid mutator transaction binding the contract method 0xc1d34b89.
//
// Solidity: function transferFromAndCall(address from, address to, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363Transactor) TransferFromAndCall(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "transferFromAndCall", from, to, value, data)
}

// TransferFromAndCall is a paid mutator transaction binding the contract method 0xc1d34b89.
//
// Solidity: function transferFromAndCall(address from, address to, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363Session) TransferFromAndCall(from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferFromAndCall(&_IERC1363.TransactOpts, from, to, value, data)
}

// TransferFromAndCall is a paid mutator transaction binding the contract method 0xc1d34b89.
//
// Solidity: function transferFromAndCall(address from, address to, uint256 value, bytes data) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) TransferFromAndCall(from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferFromAndCall(&_IERC1363.TransactOpts, from, to, value, data)
}

// TransferFromAndCall0 is a paid mutator transaction binding the contract method 0xd8fbe994.
//
// Solidity: function transferFromAndCall(address from, address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Transactor) TransferFromAndCall0(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.contract.Transact(opts, "transferFromAndCall0", from, to, value)
}

// TransferFromAndCall0 is a paid mutator transaction binding the contract method 0xd8fbe994.
//
// Solidity: function transferFromAndCall(address from, address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363Session) TransferFromAndCall0(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferFromAndCall0(&_IERC1363.TransactOpts, from, to, value)
}

// TransferFromAndCall0 is a paid mutator transaction binding the contract method 0xd8fbe994.
//
// Solidity: function transferFromAndCall(address from, address to, uint256 value) returns(bool)
func (_IERC1363 *IERC1363TransactorSession) TransferFromAndCall0(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC1363.Contract.TransferFromAndCall0(&_IERC1363.TransactOpts, from, to, value)
}

// IERC1363ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC1363 contract.
type IERC1363ApprovalIterator struct {
	Event *IERC1363Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC1363ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1363Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC1363Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC1363ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1363ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1363Approval represents a Approval event raised by the IERC1363 contract.
type IERC1363Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC1363 *IERC1363Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC1363ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC1363.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC1363ApprovalIterator{contract: _IERC1363.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC1363 *IERC1363Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC1363Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC1363.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1363Approval)
				if err := _IERC1363.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC1363 *IERC1363Filterer) ParseApproval(log types.Log) (*IERC1363Approval, error) {
	event := new(IERC1363Approval)
	if err := _IERC1363.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1363TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC1363 contract.
type IERC1363TransferIterator struct {
	Event *IERC1363Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC1363TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1363Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC1363Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC1363TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1363TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1363Transfer represents a Transfer event raised by the IERC1363 contract.
type IERC1363Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC1363 *IERC1363Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC1363TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1363.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC1363TransferIterator{contract: _IERC1363.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC1363 *IERC1363Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC1363Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC1363.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1363Transfer)
				if err := _IERC1363.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC1363 *IERC1363Filterer) ParseTransfer(log types.Log) (*IERC1363Transfer, error) {
	event := new(IERC1363Transfer)
	if err := _IERC1363.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
