// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izetanonethinterface

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

// IZetaNonEthInterfaceMetaData contains all meta data concerning the IZetaNonEthInterface contract.
var IZetaNonEthInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"mintee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// IZetaNonEthInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use IZetaNonEthInterfaceMetaData.ABI instead.
var IZetaNonEthInterfaceABI = IZetaNonEthInterfaceMetaData.ABI

// IZetaNonEthInterface is an auto generated Go binding around an Ethereum contract.
type IZetaNonEthInterface struct {
	IZetaNonEthInterfaceCaller     // Read-only binding to the contract
	IZetaNonEthInterfaceTransactor // Write-only binding to the contract
	IZetaNonEthInterfaceFilterer   // Log filterer for contract events
}

// IZetaNonEthInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZetaNonEthInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaNonEthInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZetaNonEthInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaNonEthInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZetaNonEthInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZetaNonEthInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZetaNonEthInterfaceSession struct {
	Contract     *IZetaNonEthInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IZetaNonEthInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZetaNonEthInterfaceCallerSession struct {
	Contract *IZetaNonEthInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IZetaNonEthInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZetaNonEthInterfaceTransactorSession struct {
	Contract     *IZetaNonEthInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IZetaNonEthInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZetaNonEthInterfaceRaw struct {
	Contract *IZetaNonEthInterface // Generic contract binding to access the raw methods on
}

// IZetaNonEthInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZetaNonEthInterfaceCallerRaw struct {
	Contract *IZetaNonEthInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// IZetaNonEthInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZetaNonEthInterfaceTransactorRaw struct {
	Contract *IZetaNonEthInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZetaNonEthInterface creates a new instance of IZetaNonEthInterface, bound to a specific deployed contract.
func NewIZetaNonEthInterface(address common.Address, backend bind.ContractBackend) (*IZetaNonEthInterface, error) {
	contract, err := bindIZetaNonEthInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZetaNonEthInterface{IZetaNonEthInterfaceCaller: IZetaNonEthInterfaceCaller{contract: contract}, IZetaNonEthInterfaceTransactor: IZetaNonEthInterfaceTransactor{contract: contract}, IZetaNonEthInterfaceFilterer: IZetaNonEthInterfaceFilterer{contract: contract}}, nil
}

// NewIZetaNonEthInterfaceCaller creates a new read-only instance of IZetaNonEthInterface, bound to a specific deployed contract.
func NewIZetaNonEthInterfaceCaller(address common.Address, caller bind.ContractCaller) (*IZetaNonEthInterfaceCaller, error) {
	contract, err := bindIZetaNonEthInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaNonEthInterfaceCaller{contract: contract}, nil
}

// NewIZetaNonEthInterfaceTransactor creates a new write-only instance of IZetaNonEthInterface, bound to a specific deployed contract.
func NewIZetaNonEthInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*IZetaNonEthInterfaceTransactor, error) {
	contract, err := bindIZetaNonEthInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZetaNonEthInterfaceTransactor{contract: contract}, nil
}

// NewIZetaNonEthInterfaceFilterer creates a new log filterer instance of IZetaNonEthInterface, bound to a specific deployed contract.
func NewIZetaNonEthInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*IZetaNonEthInterfaceFilterer, error) {
	contract, err := bindIZetaNonEthInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZetaNonEthInterfaceFilterer{contract: contract}, nil
}

// bindIZetaNonEthInterface binds a generic wrapper to an already deployed contract.
func bindIZetaNonEthInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZetaNonEthInterfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaNonEthInterface *IZetaNonEthInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaNonEthInterface.Contract.IZetaNonEthInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaNonEthInterface *IZetaNonEthInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaNonEthInterface.Contract.IZetaNonEthInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaNonEthInterface *IZetaNonEthInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaNonEthInterface.Contract.IZetaNonEthInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZetaNonEthInterface *IZetaNonEthInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZetaNonEthInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZetaNonEthInterface *IZetaNonEthInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZetaNonEthInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZetaNonEthInterface *IZetaNonEthInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZetaNonEthInterface.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IZetaNonEthInterface.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IZetaNonEthInterface.Contract.Allowance(&_IZetaNonEthInterface.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IZetaNonEthInterface.Contract.Allowance(&_IZetaNonEthInterface.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IZetaNonEthInterface.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IZetaNonEthInterface.Contract.BalanceOf(&_IZetaNonEthInterface.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IZetaNonEthInterface.Contract.BalanceOf(&_IZetaNonEthInterface.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZetaNonEthInterface.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceSession) TotalSupply() (*big.Int, error) {
	return _IZetaNonEthInterface.Contract.TotalSupply(&_IZetaNonEthInterface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceCallerSession) TotalSupply() (*big.Int, error) {
	return _IZetaNonEthInterface.Contract.TotalSupply(&_IZetaNonEthInterface.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthInterface.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthInterface.Contract.Approve(&_IZetaNonEthInterface.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthInterface.Contract.Approve(&_IZetaNonEthInterface.TransactOpts, spender, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IZetaNonEthInterface *IZetaNonEthInterfaceTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthInterface.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IZetaNonEthInterface *IZetaNonEthInterfaceSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthInterface.Contract.BurnFrom(&_IZetaNonEthInterface.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IZetaNonEthInterface *IZetaNonEthInterfaceTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthInterface.Contract.BurnFrom(&_IZetaNonEthInterface.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_IZetaNonEthInterface *IZetaNonEthInterfaceTransactor) Mint(opts *bind.TransactOpts, mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _IZetaNonEthInterface.contract.Transact(opts, "mint", mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_IZetaNonEthInterface *IZetaNonEthInterfaceSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _IZetaNonEthInterface.Contract.Mint(&_IZetaNonEthInterface.TransactOpts, mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_IZetaNonEthInterface *IZetaNonEthInterfaceTransactorSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _IZetaNonEthInterface.Contract.Mint(&_IZetaNonEthInterface.TransactOpts, mintee, value, internalSendHash)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthInterface.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthInterface.Contract.Transfer(&_IZetaNonEthInterface.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthInterface.Contract.Transfer(&_IZetaNonEthInterface.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthInterface.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthInterface.Contract.TransferFrom(&_IZetaNonEthInterface.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IZetaNonEthInterface.Contract.TransferFrom(&_IZetaNonEthInterface.TransactOpts, from, to, value)
}

// IZetaNonEthInterfaceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IZetaNonEthInterface contract.
type IZetaNonEthInterfaceApprovalIterator struct {
	Event *IZetaNonEthInterfaceApproval // Event containing the contract specifics and raw log

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
func (it *IZetaNonEthInterfaceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaNonEthInterfaceApproval)
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
		it.Event = new(IZetaNonEthInterfaceApproval)
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
func (it *IZetaNonEthInterfaceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaNonEthInterfaceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaNonEthInterfaceApproval represents a Approval event raised by the IZetaNonEthInterface contract.
type IZetaNonEthInterfaceApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IZetaNonEthInterfaceApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IZetaNonEthInterface.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IZetaNonEthInterfaceApprovalIterator{contract: _IZetaNonEthInterface.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IZetaNonEthInterfaceApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IZetaNonEthInterface.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaNonEthInterfaceApproval)
				if err := _IZetaNonEthInterface.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IZetaNonEthInterface *IZetaNonEthInterfaceFilterer) ParseApproval(log types.Log) (*IZetaNonEthInterfaceApproval, error) {
	event := new(IZetaNonEthInterfaceApproval)
	if err := _IZetaNonEthInterface.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZetaNonEthInterfaceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IZetaNonEthInterface contract.
type IZetaNonEthInterfaceTransferIterator struct {
	Event *IZetaNonEthInterfaceTransfer // Event containing the contract specifics and raw log

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
func (it *IZetaNonEthInterfaceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZetaNonEthInterfaceTransfer)
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
		it.Event = new(IZetaNonEthInterfaceTransfer)
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
func (it *IZetaNonEthInterfaceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZetaNonEthInterfaceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZetaNonEthInterfaceTransfer represents a Transfer event raised by the IZetaNonEthInterface contract.
type IZetaNonEthInterfaceTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IZetaNonEthInterfaceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaNonEthInterface.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IZetaNonEthInterfaceTransferIterator{contract: _IZetaNonEthInterface.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IZetaNonEthInterface *IZetaNonEthInterfaceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IZetaNonEthInterfaceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZetaNonEthInterface.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZetaNonEthInterfaceTransfer)
				if err := _IZetaNonEthInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IZetaNonEthInterface *IZetaNonEthInterfaceFilterer) ParseTransfer(log types.Log) (*IZetaNonEthInterfaceTransfer, error) {
	event := new(IZetaNonEthInterfaceTransfer)
	if err := _IZetaNonEthInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
