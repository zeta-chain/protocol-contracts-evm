// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc20custody

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
	Sender common.Address
	Asset  common.Address
	Amount *big.Int
}

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// IERC20CustodyMetaData contains all meta data concerning the IERC20Custody contract.
var IERC20CustodyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"whitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"messageContext\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndRevert\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unwhitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedCustodyTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Whitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"LegacyMethodsNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// IERC20CustodyABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20CustodyMetaData.ABI instead.
var IERC20CustodyABI = IERC20CustodyMetaData.ABI

// IERC20Custody is an auto generated Go binding around an Ethereum contract.
type IERC20Custody struct {
	IERC20CustodyCaller     // Read-only binding to the contract
	IERC20CustodyTransactor // Write-only binding to the contract
	IERC20CustodyFilterer   // Log filterer for contract events
}

// IERC20CustodyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20CustodyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20CustodyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20CustodyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20CustodySession struct {
	Contract     *IERC20Custody    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CustodyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CustodyCallerSession struct {
	Contract *IERC20CustodyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IERC20CustodyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20CustodyTransactorSession struct {
	Contract     *IERC20CustodyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IERC20CustodyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20CustodyRaw struct {
	Contract *IERC20Custody // Generic contract binding to access the raw methods on
}

// IERC20CustodyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CustodyCallerRaw struct {
	Contract *IERC20CustodyCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20CustodyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20CustodyTransactorRaw struct {
	Contract *IERC20CustodyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Custody creates a new instance of IERC20Custody, bound to a specific deployed contract.
func NewIERC20Custody(address common.Address, backend bind.ContractBackend) (*IERC20Custody, error) {
	contract, err := bindIERC20Custody(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Custody{IERC20CustodyCaller: IERC20CustodyCaller{contract: contract}, IERC20CustodyTransactor: IERC20CustodyTransactor{contract: contract}, IERC20CustodyFilterer: IERC20CustodyFilterer{contract: contract}}, nil
}

// NewIERC20CustodyCaller creates a new read-only instance of IERC20Custody, bound to a specific deployed contract.
func NewIERC20CustodyCaller(address common.Address, caller bind.ContractCaller) (*IERC20CustodyCaller, error) {
	contract, err := bindIERC20Custody(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyCaller{contract: contract}, nil
}

// NewIERC20CustodyTransactor creates a new write-only instance of IERC20Custody, bound to a specific deployed contract.
func NewIERC20CustodyTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20CustodyTransactor, error) {
	contract, err := bindIERC20Custody(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyTransactor{contract: contract}, nil
}

// NewIERC20CustodyFilterer creates a new log filterer instance of IERC20Custody, bound to a specific deployed contract.
func NewIERC20CustodyFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20CustodyFilterer, error) {
	contract, err := bindIERC20Custody(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyFilterer{contract: contract}, nil
}

// bindIERC20Custody binds a generic wrapper to an already deployed contract.
func bindIERC20Custody(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20CustodyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Custody *IERC20CustodyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Custody.Contract.IERC20CustodyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Custody *IERC20CustodyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Custody.Contract.IERC20CustodyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Custody *IERC20CustodyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Custody.Contract.IERC20CustodyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Custody *IERC20CustodyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Custody.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Custody *IERC20CustodyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Custody.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Custody *IERC20CustodyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Custody.Contract.contract.Transact(opts, method, params...)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address token) view returns(bool)
func (_IERC20Custody *IERC20CustodyCaller) Whitelisted(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _IERC20Custody.contract.Call(opts, &out, "whitelisted", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address token) view returns(bool)
func (_IERC20Custody *IERC20CustodySession) Whitelisted(token common.Address) (bool, error) {
	return _IERC20Custody.Contract.Whitelisted(&_IERC20Custody.CallOpts, token)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address token) view returns(bool)
func (_IERC20Custody *IERC20CustodyCallerSession) Whitelisted(token common.Address) (bool, error) {
	return _IERC20Custody.Contract.Whitelisted(&_IERC20Custody.CallOpts, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_IERC20Custody *IERC20CustodyTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Custody.contract.Transact(opts, "withdraw", token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_IERC20Custody *IERC20CustodySession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Custody.Contract.Withdraw(&_IERC20Custody.TransactOpts, token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_IERC20Custody *IERC20CustodyTransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Custody.Contract.Withdraw(&_IERC20Custody.TransactOpts, token, to, amount)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0xa1c432b9.
//
// Solidity: function withdrawAndCall((address,address,uint256) messageContext, address token, address to, uint256 amount, bytes data) returns()
func (_IERC20Custody *IERC20CustodyTransactor) WithdrawAndCall(opts *bind.TransactOpts, messageContext MessageContext, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC20Custody.contract.Transact(opts, "withdrawAndCall", messageContext, token, to, amount, data)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0xa1c432b9.
//
// Solidity: function withdrawAndCall((address,address,uint256) messageContext, address token, address to, uint256 amount, bytes data) returns()
func (_IERC20Custody *IERC20CustodySession) WithdrawAndCall(messageContext MessageContext, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC20Custody.Contract.WithdrawAndCall(&_IERC20Custody.TransactOpts, messageContext, token, to, amount, data)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0xa1c432b9.
//
// Solidity: function withdrawAndCall((address,address,uint256) messageContext, address token, address to, uint256 amount, bytes data) returns()
func (_IERC20Custody *IERC20CustodyTransactorSession) WithdrawAndCall(messageContext MessageContext, token common.Address, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC20Custody.Contract.WithdrawAndCall(&_IERC20Custody.TransactOpts, messageContext, token, to, amount, data)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x99a3c356.
//
// Solidity: function withdrawAndRevert(address token, address to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_IERC20Custody *IERC20CustodyTransactor) WithdrawAndRevert(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _IERC20Custody.contract.Transact(opts, "withdrawAndRevert", token, to, amount, data, revertContext)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x99a3c356.
//
// Solidity: function withdrawAndRevert(address token, address to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_IERC20Custody *IERC20CustodySession) WithdrawAndRevert(token common.Address, to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _IERC20Custody.Contract.WithdrawAndRevert(&_IERC20Custody.TransactOpts, token, to, amount, data, revertContext)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x99a3c356.
//
// Solidity: function withdrawAndRevert(address token, address to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_IERC20Custody *IERC20CustodyTransactorSession) WithdrawAndRevert(token common.Address, to common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _IERC20Custody.Contract.WithdrawAndRevert(&_IERC20Custody.TransactOpts, token, to, amount, data, revertContext)
}

// IERC20CustodyDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the IERC20Custody contract.
type IERC20CustodyDepositedIterator struct {
	Event *IERC20CustodyDeposited // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyDeposited)
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
		it.Event = new(IERC20CustodyDeposited)
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
func (it *IERC20CustodyDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyDeposited represents a Deposited event raised by the IERC20Custody contract.
type IERC20CustodyDeposited struct {
	Recipient []byte
	Asset     common.Address
	Amount    *big.Int
	Message   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_IERC20Custody *IERC20CustodyFilterer) FilterDeposited(opts *bind.FilterOpts, asset []common.Address) (*IERC20CustodyDepositedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _IERC20Custody.contract.FilterLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyDepositedIterator{contract: _IERC20Custody.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_IERC20Custody *IERC20CustodyFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *IERC20CustodyDeposited, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _IERC20Custody.contract.WatchLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyDeposited)
				if err := _IERC20Custody.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_IERC20Custody *IERC20CustodyFilterer) ParseDeposited(log types.Log) (*IERC20CustodyDeposited, error) {
	event := new(IERC20CustodyDeposited)
	if err := _IERC20Custody.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyUnwhitelistedIterator is returned from FilterUnwhitelisted and is used to iterate over the raw logs and unpacked data for Unwhitelisted events raised by the IERC20Custody contract.
type IERC20CustodyUnwhitelistedIterator struct {
	Event *IERC20CustodyUnwhitelisted // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyUnwhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyUnwhitelisted)
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
		it.Event = new(IERC20CustodyUnwhitelisted)
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
func (it *IERC20CustodyUnwhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyUnwhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyUnwhitelisted represents a Unwhitelisted event raised by the IERC20Custody contract.
type IERC20CustodyUnwhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnwhitelisted is a free log retrieval operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_IERC20Custody *IERC20CustodyFilterer) FilterUnwhitelisted(opts *bind.FilterOpts, token []common.Address) (*IERC20CustodyUnwhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20Custody.contract.FilterLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyUnwhitelistedIterator{contract: _IERC20Custody.contract, event: "Unwhitelisted", logs: logs, sub: sub}, nil
}

// WatchUnwhitelisted is a free log subscription operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_IERC20Custody *IERC20CustodyFilterer) WatchUnwhitelisted(opts *bind.WatchOpts, sink chan<- *IERC20CustodyUnwhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20Custody.contract.WatchLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyUnwhitelisted)
				if err := _IERC20Custody.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
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

// ParseUnwhitelisted is a log parse operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_IERC20Custody *IERC20CustodyFilterer) ParseUnwhitelisted(log types.Log) (*IERC20CustodyUnwhitelisted, error) {
	event := new(IERC20CustodyUnwhitelisted)
	if err := _IERC20Custody.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyUpdatedCustodyTSSAddressIterator is returned from FilterUpdatedCustodyTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedCustodyTSSAddress events raised by the IERC20Custody contract.
type IERC20CustodyUpdatedCustodyTSSAddressIterator struct {
	Event *IERC20CustodyUpdatedCustodyTSSAddress // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyUpdatedCustodyTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyUpdatedCustodyTSSAddress)
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
		it.Event = new(IERC20CustodyUpdatedCustodyTSSAddress)
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
func (it *IERC20CustodyUpdatedCustodyTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyUpdatedCustodyTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyUpdatedCustodyTSSAddress represents a UpdatedCustodyTSSAddress event raised by the IERC20Custody contract.
type IERC20CustodyUpdatedCustodyTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCustodyTSSAddress is a free log retrieval operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_IERC20Custody *IERC20CustodyFilterer) FilterUpdatedCustodyTSSAddress(opts *bind.FilterOpts) (*IERC20CustodyUpdatedCustodyTSSAddressIterator, error) {

	logs, sub, err := _IERC20Custody.contract.FilterLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyUpdatedCustodyTSSAddressIterator{contract: _IERC20Custody.contract, event: "UpdatedCustodyTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedCustodyTSSAddress is a free log subscription operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_IERC20Custody *IERC20CustodyFilterer) WatchUpdatedCustodyTSSAddress(opts *bind.WatchOpts, sink chan<- *IERC20CustodyUpdatedCustodyTSSAddress) (event.Subscription, error) {

	logs, sub, err := _IERC20Custody.contract.WatchLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyUpdatedCustodyTSSAddress)
				if err := _IERC20Custody.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
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

// ParseUpdatedCustodyTSSAddress is a log parse operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_IERC20Custody *IERC20CustodyFilterer) ParseUpdatedCustodyTSSAddress(log types.Log) (*IERC20CustodyUpdatedCustodyTSSAddress, error) {
	event := new(IERC20CustodyUpdatedCustodyTSSAddress)
	if err := _IERC20Custody.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the IERC20Custody contract.
type IERC20CustodyWhitelistedIterator struct {
	Event *IERC20CustodyWhitelisted // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyWhitelisted)
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
		it.Event = new(IERC20CustodyWhitelisted)
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
func (it *IERC20CustodyWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyWhitelisted represents a Whitelisted event raised by the IERC20Custody contract.
type IERC20CustodyWhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_IERC20Custody *IERC20CustodyFilterer) FilterWhitelisted(opts *bind.FilterOpts, token []common.Address) (*IERC20CustodyWhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20Custody.contract.FilterLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyWhitelistedIterator{contract: _IERC20Custody.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_IERC20Custody *IERC20CustodyFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *IERC20CustodyWhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20Custody.contract.WatchLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyWhitelisted)
				if err := _IERC20Custody.contract.UnpackLog(event, "Whitelisted", log); err != nil {
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

// ParseWhitelisted is a log parse operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_IERC20Custody *IERC20CustodyFilterer) ParseWhitelisted(log types.Log) (*IERC20CustodyWhitelisted, error) {
	event := new(IERC20CustodyWhitelisted)
	if err := _IERC20Custody.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the IERC20Custody contract.
type IERC20CustodyWithdrawnIterator struct {
	Event *IERC20CustodyWithdrawn // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyWithdrawn)
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
		it.Event = new(IERC20CustodyWithdrawn)
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
func (it *IERC20CustodyWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyWithdrawn represents a Withdrawn event raised by the IERC20Custody contract.
type IERC20CustodyWithdrawn struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_IERC20Custody *IERC20CustodyFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*IERC20CustodyWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20Custody.contract.FilterLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyWithdrawnIterator{contract: _IERC20Custody.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_IERC20Custody *IERC20CustodyFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *IERC20CustodyWithdrawn, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20Custody.contract.WatchLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyWithdrawn)
				if err := _IERC20Custody.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_IERC20Custody *IERC20CustodyFilterer) ParseWithdrawn(log types.Log) (*IERC20CustodyWithdrawn, error) {
	event := new(IERC20CustodyWithdrawn)
	if err := _IERC20Custody.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the IERC20Custody contract.
type IERC20CustodyWithdrawnAndCalledIterator struct {
	Event *IERC20CustodyWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyWithdrawnAndCalled)
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
		it.Event = new(IERC20CustodyWithdrawnAndCalled)
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
func (it *IERC20CustodyWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the IERC20Custody contract.
type IERC20CustodyWithdrawnAndCalled struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_IERC20Custody *IERC20CustodyFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*IERC20CustodyWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20Custody.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyWithdrawnAndCalledIterator{contract: _IERC20Custody.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_IERC20Custody *IERC20CustodyFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *IERC20CustodyWithdrawnAndCalled, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20Custody.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyWithdrawnAndCalled)
				if err := _IERC20Custody.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_IERC20Custody *IERC20CustodyFilterer) ParseWithdrawnAndCalled(log types.Log) (*IERC20CustodyWithdrawnAndCalled, error) {
	event := new(IERC20CustodyWithdrawnAndCalled)
	if err := _IERC20Custody.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the IERC20Custody contract.
type IERC20CustodyWithdrawnAndRevertedIterator struct {
	Event *IERC20CustodyWithdrawnAndReverted // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyWithdrawnAndReverted)
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
		it.Event = new(IERC20CustodyWithdrawnAndReverted)
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
func (it *IERC20CustodyWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the IERC20Custody contract.
type IERC20CustodyWithdrawnAndReverted struct {
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_IERC20Custody *IERC20CustodyFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*IERC20CustodyWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20Custody.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyWithdrawnAndRevertedIterator{contract: _IERC20Custody.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_IERC20Custody *IERC20CustodyFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *IERC20CustodyWithdrawnAndReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20Custody.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyWithdrawnAndReverted)
				if err := _IERC20Custody.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
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

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_IERC20Custody *IERC20CustodyFilterer) ParseWithdrawnAndReverted(log types.Log) (*IERC20CustodyWithdrawnAndReverted, error) {
	event := new(IERC20CustodyWithdrawnAndReverted)
	if err := _IERC20Custody.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
