// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc1967utils

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

// ERC1967UtilsMetaData contains all meta data concerning the ERC1967Utils contract.
var ERC1967UtilsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC1967InvalidAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidBeacon\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]}]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220532640c8bd776ca2798ca402dcee6bfb6199d2e0db09530bfe1981c8d71c8f6a64736f6c634300081a0033",
}

// ERC1967UtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1967UtilsMetaData.ABI instead.
var ERC1967UtilsABI = ERC1967UtilsMetaData.ABI

// ERC1967UtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC1967UtilsMetaData.Bin instead.
var ERC1967UtilsBin = ERC1967UtilsMetaData.Bin

// DeployERC1967Utils deploys a new Ethereum contract, binding an instance of ERC1967Utils to it.
func DeployERC1967Utils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC1967Utils, error) {
	parsed, err := ERC1967UtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC1967UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC1967Utils{ERC1967UtilsCaller: ERC1967UtilsCaller{contract: contract}, ERC1967UtilsTransactor: ERC1967UtilsTransactor{contract: contract}, ERC1967UtilsFilterer: ERC1967UtilsFilterer{contract: contract}}, nil
}

// ERC1967Utils is an auto generated Go binding around an Ethereum contract.
type ERC1967Utils struct {
	ERC1967UtilsCaller     // Read-only binding to the contract
	ERC1967UtilsTransactor // Write-only binding to the contract
	ERC1967UtilsFilterer   // Log filterer for contract events
}

// ERC1967UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1967UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1967UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1967UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1967UtilsSession struct {
	Contract     *ERC1967Utils     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1967UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1967UtilsCallerSession struct {
	Contract *ERC1967UtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC1967UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1967UtilsTransactorSession struct {
	Contract     *ERC1967UtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC1967UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1967UtilsRaw struct {
	Contract *ERC1967Utils // Generic contract binding to access the raw methods on
}

// ERC1967UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1967UtilsCallerRaw struct {
	Contract *ERC1967UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1967UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1967UtilsTransactorRaw struct {
	Contract *ERC1967UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1967Utils creates a new instance of ERC1967Utils, bound to a specific deployed contract.
func NewERC1967Utils(address common.Address, backend bind.ContractBackend) (*ERC1967Utils, error) {
	contract, err := bindERC1967Utils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1967Utils{ERC1967UtilsCaller: ERC1967UtilsCaller{contract: contract}, ERC1967UtilsTransactor: ERC1967UtilsTransactor{contract: contract}, ERC1967UtilsFilterer: ERC1967UtilsFilterer{contract: contract}}, nil
}

// NewERC1967UtilsCaller creates a new read-only instance of ERC1967Utils, bound to a specific deployed contract.
func NewERC1967UtilsCaller(address common.Address, caller bind.ContractCaller) (*ERC1967UtilsCaller, error) {
	contract, err := bindERC1967Utils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UtilsCaller{contract: contract}, nil
}

// NewERC1967UtilsTransactor creates a new write-only instance of ERC1967Utils, bound to a specific deployed contract.
func NewERC1967UtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1967UtilsTransactor, error) {
	contract, err := bindERC1967Utils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UtilsTransactor{contract: contract}, nil
}

// NewERC1967UtilsFilterer creates a new log filterer instance of ERC1967Utils, bound to a specific deployed contract.
func NewERC1967UtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1967UtilsFilterer, error) {
	contract, err := bindERC1967Utils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1967UtilsFilterer{contract: contract}, nil
}

// bindERC1967Utils binds a generic wrapper to an already deployed contract.
func bindERC1967Utils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1967UtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967Utils *ERC1967UtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967Utils.Contract.ERC1967UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967Utils *ERC1967UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967Utils.Contract.ERC1967UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967Utils *ERC1967UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967Utils.Contract.ERC1967UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967Utils *ERC1967UtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967Utils *ERC1967UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967Utils *ERC1967UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967Utils.Contract.contract.Transact(opts, method, params...)
}

// ERC1967UtilsAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ERC1967Utils contract.
type ERC1967UtilsAdminChangedIterator struct {
	Event *ERC1967UtilsAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC1967UtilsAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UtilsAdminChanged)
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
		it.Event = new(ERC1967UtilsAdminChanged)
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
func (it *ERC1967UtilsAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UtilsAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UtilsAdminChanged represents a AdminChanged event raised by the ERC1967Utils contract.
type ERC1967UtilsAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967Utils *ERC1967UtilsFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ERC1967UtilsAdminChangedIterator, error) {

	logs, sub, err := _ERC1967Utils.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ERC1967UtilsAdminChangedIterator{contract: _ERC1967Utils.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967Utils *ERC1967UtilsFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC1967UtilsAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ERC1967Utils.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UtilsAdminChanged)
				if err := _ERC1967Utils.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967Utils *ERC1967UtilsFilterer) ParseAdminChanged(log types.Log) (*ERC1967UtilsAdminChanged, error) {
	event := new(ERC1967UtilsAdminChanged)
	if err := _ERC1967Utils.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UtilsBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ERC1967Utils contract.
type ERC1967UtilsBeaconUpgradedIterator struct {
	Event *ERC1967UtilsBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967UtilsBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UtilsBeaconUpgraded)
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
		it.Event = new(ERC1967UtilsBeaconUpgraded)
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
func (it *ERC1967UtilsBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UtilsBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UtilsBeaconUpgraded represents a BeaconUpgraded event raised by the ERC1967Utils contract.
type ERC1967UtilsBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967Utils *ERC1967UtilsFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ERC1967UtilsBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967Utils.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967UtilsBeaconUpgradedIterator{contract: _ERC1967Utils.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967Utils *ERC1967UtilsFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967UtilsBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967Utils.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UtilsBeaconUpgraded)
				if err := _ERC1967Utils.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967Utils *ERC1967UtilsFilterer) ParseBeaconUpgraded(log types.Log) (*ERC1967UtilsBeaconUpgraded, error) {
	event := new(ERC1967UtilsBeaconUpgraded)
	if err := _ERC1967Utils.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UtilsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC1967Utils contract.
type ERC1967UtilsUpgradedIterator struct {
	Event *ERC1967UtilsUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967UtilsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UtilsUpgraded)
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
		it.Event = new(ERC1967UtilsUpgraded)
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
func (it *ERC1967UtilsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UtilsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UtilsUpgraded represents a Upgraded event raised by the ERC1967Utils contract.
type ERC1967UtilsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Utils *ERC1967UtilsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ERC1967UtilsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967Utils.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967UtilsUpgradedIterator{contract: _ERC1967Utils.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Utils *ERC1967UtilsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967UtilsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967Utils.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UtilsUpgraded)
				if err := _ERC1967Utils.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Utils *ERC1967UtilsFilterer) ParseUpgraded(log types.Log) (*ERC1967UtilsUpgraded, error) {
	event := new(ERC1967UtilsUpgraded)
	if err := _ERC1967Utils.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
