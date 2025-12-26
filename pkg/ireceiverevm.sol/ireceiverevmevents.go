// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ireceiverevm

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

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// IReceiverEVMEventsMetaData contains all meta data concerning the IReceiverEVMEvents contract.
var IReceiverEVMEventsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCallV2\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false}]",
}

// IReceiverEVMEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IReceiverEVMEventsMetaData.ABI instead.
var IReceiverEVMEventsABI = IReceiverEVMEventsMetaData.ABI

// IReceiverEVMEvents is an auto generated Go binding around an Ethereum contract.
type IReceiverEVMEvents struct {
	IReceiverEVMEventsCaller     // Read-only binding to the contract
	IReceiverEVMEventsTransactor // Write-only binding to the contract
	IReceiverEVMEventsFilterer   // Log filterer for contract events
}

// IReceiverEVMEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IReceiverEVMEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReceiverEVMEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IReceiverEVMEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReceiverEVMEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IReceiverEVMEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReceiverEVMEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IReceiverEVMEventsSession struct {
	Contract     *IReceiverEVMEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IReceiverEVMEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IReceiverEVMEventsCallerSession struct {
	Contract *IReceiverEVMEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IReceiverEVMEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IReceiverEVMEventsTransactorSession struct {
	Contract     *IReceiverEVMEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IReceiverEVMEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IReceiverEVMEventsRaw struct {
	Contract *IReceiverEVMEvents // Generic contract binding to access the raw methods on
}

// IReceiverEVMEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IReceiverEVMEventsCallerRaw struct {
	Contract *IReceiverEVMEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IReceiverEVMEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IReceiverEVMEventsTransactorRaw struct {
	Contract *IReceiverEVMEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIReceiverEVMEvents creates a new instance of IReceiverEVMEvents, bound to a specific deployed contract.
func NewIReceiverEVMEvents(address common.Address, backend bind.ContractBackend) (*IReceiverEVMEvents, error) {
	contract, err := bindIReceiverEVMEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IReceiverEVMEvents{IReceiverEVMEventsCaller: IReceiverEVMEventsCaller{contract: contract}, IReceiverEVMEventsTransactor: IReceiverEVMEventsTransactor{contract: contract}, IReceiverEVMEventsFilterer: IReceiverEVMEventsFilterer{contract: contract}}, nil
}

// NewIReceiverEVMEventsCaller creates a new read-only instance of IReceiverEVMEvents, bound to a specific deployed contract.
func NewIReceiverEVMEventsCaller(address common.Address, caller bind.ContractCaller) (*IReceiverEVMEventsCaller, error) {
	contract, err := bindIReceiverEVMEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IReceiverEVMEventsCaller{contract: contract}, nil
}

// NewIReceiverEVMEventsTransactor creates a new write-only instance of IReceiverEVMEvents, bound to a specific deployed contract.
func NewIReceiverEVMEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IReceiverEVMEventsTransactor, error) {
	contract, err := bindIReceiverEVMEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IReceiverEVMEventsTransactor{contract: contract}, nil
}

// NewIReceiverEVMEventsFilterer creates a new log filterer instance of IReceiverEVMEvents, bound to a specific deployed contract.
func NewIReceiverEVMEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IReceiverEVMEventsFilterer, error) {
	contract, err := bindIReceiverEVMEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IReceiverEVMEventsFilterer{contract: contract}, nil
}

// bindIReceiverEVMEvents binds a generic wrapper to an already deployed contract.
func bindIReceiverEVMEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IReceiverEVMEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReceiverEVMEvents *IReceiverEVMEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReceiverEVMEvents.Contract.IReceiverEVMEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReceiverEVMEvents *IReceiverEVMEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReceiverEVMEvents.Contract.IReceiverEVMEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReceiverEVMEvents *IReceiverEVMEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReceiverEVMEvents.Contract.IReceiverEVMEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReceiverEVMEvents *IReceiverEVMEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReceiverEVMEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReceiverEVMEvents *IReceiverEVMEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReceiverEVMEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReceiverEVMEvents *IReceiverEVMEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReceiverEVMEvents.Contract.contract.Transact(opts, method, params...)
}

// IReceiverEVMEventsReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the IReceiverEVMEvents contract.
type IReceiverEVMEventsReceivedERC20Iterator struct {
	Event *IReceiverEVMEventsReceivedERC20 // Event containing the contract specifics and raw log

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
func (it *IReceiverEVMEventsReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReceiverEVMEventsReceivedERC20)
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
		it.Event = new(IReceiverEVMEventsReceivedERC20)
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
func (it *IReceiverEVMEventsReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReceiverEVMEventsReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReceiverEVMEventsReceivedERC20 represents a ReceivedERC20 event raised by the IReceiverEVMEvents contract.
type IReceiverEVMEventsReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*IReceiverEVMEventsReceivedERC20Iterator, error) {

	logs, sub, err := _IReceiverEVMEvents.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &IReceiverEVMEventsReceivedERC20Iterator{contract: _IReceiverEVMEvents.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *IReceiverEVMEventsReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _IReceiverEVMEvents.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReceiverEVMEventsReceivedERC20)
				if err := _IReceiverEVMEvents.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
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

// ParseReceivedERC20 is a log parse operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) ParseReceivedERC20(log types.Log) (*IReceiverEVMEventsReceivedERC20, error) {
	event := new(IReceiverEVMEventsReceivedERC20)
	if err := _IReceiverEVMEvents.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IReceiverEVMEventsReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the IReceiverEVMEvents contract.
type IReceiverEVMEventsReceivedNoParamsIterator struct {
	Event *IReceiverEVMEventsReceivedNoParams // Event containing the contract specifics and raw log

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
func (it *IReceiverEVMEventsReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReceiverEVMEventsReceivedNoParams)
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
		it.Event = new(IReceiverEVMEventsReceivedNoParams)
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
func (it *IReceiverEVMEventsReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReceiverEVMEventsReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReceiverEVMEventsReceivedNoParams represents a ReceivedNoParams event raised by the IReceiverEVMEvents contract.
type IReceiverEVMEventsReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*IReceiverEVMEventsReceivedNoParamsIterator, error) {

	logs, sub, err := _IReceiverEVMEvents.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &IReceiverEVMEventsReceivedNoParamsIterator{contract: _IReceiverEVMEvents.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *IReceiverEVMEventsReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _IReceiverEVMEvents.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReceiverEVMEventsReceivedNoParams)
				if err := _IReceiverEVMEvents.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
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

// ParseReceivedNoParams is a log parse operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) ParseReceivedNoParams(log types.Log) (*IReceiverEVMEventsReceivedNoParams, error) {
	event := new(IReceiverEVMEventsReceivedNoParams)
	if err := _IReceiverEVMEvents.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IReceiverEVMEventsReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the IReceiverEVMEvents contract.
type IReceiverEVMEventsReceivedNonPayableIterator struct {
	Event *IReceiverEVMEventsReceivedNonPayable // Event containing the contract specifics and raw log

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
func (it *IReceiverEVMEventsReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReceiverEVMEventsReceivedNonPayable)
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
		it.Event = new(IReceiverEVMEventsReceivedNonPayable)
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
func (it *IReceiverEVMEventsReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReceiverEVMEventsReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReceiverEVMEventsReceivedNonPayable represents a ReceivedNonPayable event raised by the IReceiverEVMEvents contract.
type IReceiverEVMEventsReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*IReceiverEVMEventsReceivedNonPayableIterator, error) {

	logs, sub, err := _IReceiverEVMEvents.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &IReceiverEVMEventsReceivedNonPayableIterator{contract: _IReceiverEVMEvents.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *IReceiverEVMEventsReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _IReceiverEVMEvents.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReceiverEVMEventsReceivedNonPayable)
				if err := _IReceiverEVMEvents.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
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

// ParseReceivedNonPayable is a log parse operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) ParseReceivedNonPayable(log types.Log) (*IReceiverEVMEventsReceivedNonPayable, error) {
	event := new(IReceiverEVMEventsReceivedNonPayable)
	if err := _IReceiverEVMEvents.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IReceiverEVMEventsReceivedOnCallIterator is returned from FilterReceivedOnCall and is used to iterate over the raw logs and unpacked data for ReceivedOnCall events raised by the IReceiverEVMEvents contract.
type IReceiverEVMEventsReceivedOnCallIterator struct {
	Event *IReceiverEVMEventsReceivedOnCall // Event containing the contract specifics and raw log

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
func (it *IReceiverEVMEventsReceivedOnCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReceiverEVMEventsReceivedOnCall)
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
		it.Event = new(IReceiverEVMEventsReceivedOnCall)
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
func (it *IReceiverEVMEventsReceivedOnCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReceiverEVMEventsReceivedOnCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReceiverEVMEventsReceivedOnCall represents a ReceivedOnCall event raised by the IReceiverEVMEvents contract.
type IReceiverEVMEventsReceivedOnCall struct {
	Sender  common.Address
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceivedOnCall is a free log retrieval operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) FilterReceivedOnCall(opts *bind.FilterOpts) (*IReceiverEVMEventsReceivedOnCallIterator, error) {

	logs, sub, err := _IReceiverEVMEvents.contract.FilterLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return &IReceiverEVMEventsReceivedOnCallIterator{contract: _IReceiverEVMEvents.contract, event: "ReceivedOnCall", logs: logs, sub: sub}, nil
}

// WatchReceivedOnCall is a free log subscription operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) WatchReceivedOnCall(opts *bind.WatchOpts, sink chan<- *IReceiverEVMEventsReceivedOnCall) (event.Subscription, error) {

	logs, sub, err := _IReceiverEVMEvents.contract.WatchLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReceiverEVMEventsReceivedOnCall)
				if err := _IReceiverEVMEvents.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
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

// ParseReceivedOnCall is a log parse operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) ParseReceivedOnCall(log types.Log) (*IReceiverEVMEventsReceivedOnCall, error) {
	event := new(IReceiverEVMEventsReceivedOnCall)
	if err := _IReceiverEVMEvents.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IReceiverEVMEventsReceivedOnCallV2Iterator is returned from FilterReceivedOnCallV2 and is used to iterate over the raw logs and unpacked data for ReceivedOnCallV2 events raised by the IReceiverEVMEvents contract.
type IReceiverEVMEventsReceivedOnCallV2Iterator struct {
	Event *IReceiverEVMEventsReceivedOnCallV2 // Event containing the contract specifics and raw log

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
func (it *IReceiverEVMEventsReceivedOnCallV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReceiverEVMEventsReceivedOnCallV2)
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
		it.Event = new(IReceiverEVMEventsReceivedOnCallV2)
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
func (it *IReceiverEVMEventsReceivedOnCallV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReceiverEVMEventsReceivedOnCallV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReceiverEVMEventsReceivedOnCallV2 represents a ReceivedOnCallV2 event raised by the IReceiverEVMEvents contract.
type IReceiverEVMEventsReceivedOnCallV2 struct {
	Sender  common.Address
	Asset   common.Address
	Amount  *big.Int
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceivedOnCallV2 is a free log retrieval operation binding the contract event 0x77c91f89244a04200aa9bae5695cacb5a2b6894a33d119ee69184324139feb9c.
//
// Solidity: event ReceivedOnCallV2(address sender, address asset, uint256 amount, bytes message)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) FilterReceivedOnCallV2(opts *bind.FilterOpts) (*IReceiverEVMEventsReceivedOnCallV2Iterator, error) {

	logs, sub, err := _IReceiverEVMEvents.contract.FilterLogs(opts, "ReceivedOnCallV2")
	if err != nil {
		return nil, err
	}
	return &IReceiverEVMEventsReceivedOnCallV2Iterator{contract: _IReceiverEVMEvents.contract, event: "ReceivedOnCallV2", logs: logs, sub: sub}, nil
}

// WatchReceivedOnCallV2 is a free log subscription operation binding the contract event 0x77c91f89244a04200aa9bae5695cacb5a2b6894a33d119ee69184324139feb9c.
//
// Solidity: event ReceivedOnCallV2(address sender, address asset, uint256 amount, bytes message)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) WatchReceivedOnCallV2(opts *bind.WatchOpts, sink chan<- *IReceiverEVMEventsReceivedOnCallV2) (event.Subscription, error) {

	logs, sub, err := _IReceiverEVMEvents.contract.WatchLogs(opts, "ReceivedOnCallV2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReceiverEVMEventsReceivedOnCallV2)
				if err := _IReceiverEVMEvents.contract.UnpackLog(event, "ReceivedOnCallV2", log); err != nil {
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

// ParseReceivedOnCallV2 is a log parse operation binding the contract event 0x77c91f89244a04200aa9bae5695cacb5a2b6894a33d119ee69184324139feb9c.
//
// Solidity: event ReceivedOnCallV2(address sender, address asset, uint256 amount, bytes message)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) ParseReceivedOnCallV2(log types.Log) (*IReceiverEVMEventsReceivedOnCallV2, error) {
	event := new(IReceiverEVMEventsReceivedOnCallV2)
	if err := _IReceiverEVMEvents.contract.UnpackLog(event, "ReceivedOnCallV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IReceiverEVMEventsReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the IReceiverEVMEvents contract.
type IReceiverEVMEventsReceivedPayableIterator struct {
	Event *IReceiverEVMEventsReceivedPayable // Event containing the contract specifics and raw log

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
func (it *IReceiverEVMEventsReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReceiverEVMEventsReceivedPayable)
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
		it.Event = new(IReceiverEVMEventsReceivedPayable)
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
func (it *IReceiverEVMEventsReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReceiverEVMEventsReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReceiverEVMEventsReceivedPayable represents a ReceivedPayable event raised by the IReceiverEVMEvents contract.
type IReceiverEVMEventsReceivedPayable struct {
	Sender common.Address
	Value  *big.Int
	Str    string
	Num    *big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedPayable is a free log retrieval operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*IReceiverEVMEventsReceivedPayableIterator, error) {

	logs, sub, err := _IReceiverEVMEvents.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &IReceiverEVMEventsReceivedPayableIterator{contract: _IReceiverEVMEvents.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *IReceiverEVMEventsReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _IReceiverEVMEvents.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReceiverEVMEventsReceivedPayable)
				if err := _IReceiverEVMEvents.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
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

// ParseReceivedPayable is a log parse operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) ParseReceivedPayable(log types.Log) (*IReceiverEVMEventsReceivedPayable, error) {
	event := new(IReceiverEVMEventsReceivedPayable)
	if err := _IReceiverEVMEvents.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IReceiverEVMEventsReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the IReceiverEVMEvents contract.
type IReceiverEVMEventsReceivedRevertIterator struct {
	Event *IReceiverEVMEventsReceivedRevert // Event containing the contract specifics and raw log

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
func (it *IReceiverEVMEventsReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IReceiverEVMEventsReceivedRevert)
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
		it.Event = new(IReceiverEVMEventsReceivedRevert)
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
func (it *IReceiverEVMEventsReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IReceiverEVMEventsReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IReceiverEVMEventsReceivedRevert represents a ReceivedRevert event raised by the IReceiverEVMEvents contract.
type IReceiverEVMEventsReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*IReceiverEVMEventsReceivedRevertIterator, error) {

	logs, sub, err := _IReceiverEVMEvents.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &IReceiverEVMEventsReceivedRevertIterator{contract: _IReceiverEVMEvents.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *IReceiverEVMEventsReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _IReceiverEVMEvents.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IReceiverEVMEventsReceivedRevert)
				if err := _IReceiverEVMEvents.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
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

// ParseReceivedRevert is a log parse operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_IReceiverEVMEvents *IReceiverEVMEventsFilterer) ParseReceivedRevert(log types.Log) (*IReceiverEVMEventsReceivedRevert, error) {
	event := new(IReceiverEVMEventsReceivedRevert)
	if err := _IReceiverEVMEvents.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
