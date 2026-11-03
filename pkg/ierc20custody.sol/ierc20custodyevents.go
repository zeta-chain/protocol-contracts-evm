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

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// IERC20CustodyEventsMetaData contains all meta data concerning the IERC20CustodyEvents contract.
var IERC20CustodyEventsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unwhitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedCustodyTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Whitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false}]",
}

// IERC20CustodyEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20CustodyEventsMetaData.ABI instead.
var IERC20CustodyEventsABI = IERC20CustodyEventsMetaData.ABI

// IERC20CustodyEvents is an auto generated Go binding around an Ethereum contract.
type IERC20CustodyEvents struct {
	IERC20CustodyEventsCaller     // Read-only binding to the contract
	IERC20CustodyEventsTransactor // Write-only binding to the contract
	IERC20CustodyEventsFilterer   // Log filterer for contract events
}

// IERC20CustodyEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20CustodyEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20CustodyEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20CustodyEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20CustodyEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20CustodyEventsSession struct {
	Contract     *IERC20CustodyEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IERC20CustodyEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CustodyEventsCallerSession struct {
	Contract *IERC20CustodyEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IERC20CustodyEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20CustodyEventsTransactorSession struct {
	Contract     *IERC20CustodyEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IERC20CustodyEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20CustodyEventsRaw struct {
	Contract *IERC20CustodyEvents // Generic contract binding to access the raw methods on
}

// IERC20CustodyEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CustodyEventsCallerRaw struct {
	Contract *IERC20CustodyEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20CustodyEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20CustodyEventsTransactorRaw struct {
	Contract *IERC20CustodyEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20CustodyEvents creates a new instance of IERC20CustodyEvents, bound to a specific deployed contract.
func NewIERC20CustodyEvents(address common.Address, backend bind.ContractBackend) (*IERC20CustodyEvents, error) {
	contract, err := bindIERC20CustodyEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEvents{IERC20CustodyEventsCaller: IERC20CustodyEventsCaller{contract: contract}, IERC20CustodyEventsTransactor: IERC20CustodyEventsTransactor{contract: contract}, IERC20CustodyEventsFilterer: IERC20CustodyEventsFilterer{contract: contract}}, nil
}

// NewIERC20CustodyEventsCaller creates a new read-only instance of IERC20CustodyEvents, bound to a specific deployed contract.
func NewIERC20CustodyEventsCaller(address common.Address, caller bind.ContractCaller) (*IERC20CustodyEventsCaller, error) {
	contract, err := bindIERC20CustodyEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsCaller{contract: contract}, nil
}

// NewIERC20CustodyEventsTransactor creates a new write-only instance of IERC20CustodyEvents, bound to a specific deployed contract.
func NewIERC20CustodyEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20CustodyEventsTransactor, error) {
	contract, err := bindIERC20CustodyEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsTransactor{contract: contract}, nil
}

// NewIERC20CustodyEventsFilterer creates a new log filterer instance of IERC20CustodyEvents, bound to a specific deployed contract.
func NewIERC20CustodyEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20CustodyEventsFilterer, error) {
	contract, err := bindIERC20CustodyEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsFilterer{contract: contract}, nil
}

// bindIERC20CustodyEvents binds a generic wrapper to an already deployed contract.
func bindIERC20CustodyEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20CustodyEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20CustodyEvents *IERC20CustodyEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20CustodyEvents.Contract.IERC20CustodyEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20CustodyEvents *IERC20CustodyEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20CustodyEvents.Contract.IERC20CustodyEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20CustodyEvents *IERC20CustodyEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20CustodyEvents.Contract.IERC20CustodyEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20CustodyEvents *IERC20CustodyEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20CustodyEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20CustodyEvents *IERC20CustodyEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20CustodyEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20CustodyEvents *IERC20CustodyEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20CustodyEvents.Contract.contract.Transact(opts, method, params...)
}

// IERC20CustodyEventsDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsDepositedIterator struct {
	Event *IERC20CustodyEventsDeposited // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyEventsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyEventsDeposited)
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
		it.Event = new(IERC20CustodyEventsDeposited)
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
func (it *IERC20CustodyEventsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyEventsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyEventsDeposited represents a Deposited event raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsDeposited struct {
	Recipient []byte
	Asset     common.Address
	Amount    *big.Int
	Message   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) FilterDeposited(opts *bind.FilterOpts, asset []common.Address) (*IERC20CustodyEventsDepositedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.FilterLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsDepositedIterator{contract: _IERC20CustodyEvents.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *IERC20CustodyEventsDeposited, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.WatchLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyEventsDeposited)
				if err := _IERC20CustodyEvents.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) ParseDeposited(log types.Log) (*IERC20CustodyEventsDeposited, error) {
	event := new(IERC20CustodyEventsDeposited)
	if err := _IERC20CustodyEvents.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyEventsUnwhitelistedIterator is returned from FilterUnwhitelisted and is used to iterate over the raw logs and unpacked data for Unwhitelisted events raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsUnwhitelistedIterator struct {
	Event *IERC20CustodyEventsUnwhitelisted // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyEventsUnwhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyEventsUnwhitelisted)
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
		it.Event = new(IERC20CustodyEventsUnwhitelisted)
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
func (it *IERC20CustodyEventsUnwhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyEventsUnwhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyEventsUnwhitelisted represents a Unwhitelisted event raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsUnwhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnwhitelisted is a free log retrieval operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) FilterUnwhitelisted(opts *bind.FilterOpts, token []common.Address) (*IERC20CustodyEventsUnwhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.FilterLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsUnwhitelistedIterator{contract: _IERC20CustodyEvents.contract, event: "Unwhitelisted", logs: logs, sub: sub}, nil
}

// WatchUnwhitelisted is a free log subscription operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) WatchUnwhitelisted(opts *bind.WatchOpts, sink chan<- *IERC20CustodyEventsUnwhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.WatchLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyEventsUnwhitelisted)
				if err := _IERC20CustodyEvents.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
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
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) ParseUnwhitelisted(log types.Log) (*IERC20CustodyEventsUnwhitelisted, error) {
	event := new(IERC20CustodyEventsUnwhitelisted)
	if err := _IERC20CustodyEvents.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyEventsUpdatedCustodyTSSAddressIterator is returned from FilterUpdatedCustodyTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedCustodyTSSAddress events raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsUpdatedCustodyTSSAddressIterator struct {
	Event *IERC20CustodyEventsUpdatedCustodyTSSAddress // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyEventsUpdatedCustodyTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyEventsUpdatedCustodyTSSAddress)
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
		it.Event = new(IERC20CustodyEventsUpdatedCustodyTSSAddress)
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
func (it *IERC20CustodyEventsUpdatedCustodyTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyEventsUpdatedCustodyTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyEventsUpdatedCustodyTSSAddress represents a UpdatedCustodyTSSAddress event raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsUpdatedCustodyTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCustodyTSSAddress is a free log retrieval operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) FilterUpdatedCustodyTSSAddress(opts *bind.FilterOpts) (*IERC20CustodyEventsUpdatedCustodyTSSAddressIterator, error) {

	logs, sub, err := _IERC20CustodyEvents.contract.FilterLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsUpdatedCustodyTSSAddressIterator{contract: _IERC20CustodyEvents.contract, event: "UpdatedCustodyTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedCustodyTSSAddress is a free log subscription operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) WatchUpdatedCustodyTSSAddress(opts *bind.WatchOpts, sink chan<- *IERC20CustodyEventsUpdatedCustodyTSSAddress) (event.Subscription, error) {

	logs, sub, err := _IERC20CustodyEvents.contract.WatchLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyEventsUpdatedCustodyTSSAddress)
				if err := _IERC20CustodyEvents.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
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
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) ParseUpdatedCustodyTSSAddress(log types.Log) (*IERC20CustodyEventsUpdatedCustodyTSSAddress, error) {
	event := new(IERC20CustodyEventsUpdatedCustodyTSSAddress)
	if err := _IERC20CustodyEvents.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyEventsWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsWhitelistedIterator struct {
	Event *IERC20CustodyEventsWhitelisted // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyEventsWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyEventsWhitelisted)
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
		it.Event = new(IERC20CustodyEventsWhitelisted)
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
func (it *IERC20CustodyEventsWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyEventsWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyEventsWhitelisted represents a Whitelisted event raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsWhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) FilterWhitelisted(opts *bind.FilterOpts, token []common.Address) (*IERC20CustodyEventsWhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.FilterLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsWhitelistedIterator{contract: _IERC20CustodyEvents.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *IERC20CustodyEventsWhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.WatchLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyEventsWhitelisted)
				if err := _IERC20CustodyEvents.contract.UnpackLog(event, "Whitelisted", log); err != nil {
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
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) ParseWhitelisted(log types.Log) (*IERC20CustodyEventsWhitelisted, error) {
	event := new(IERC20CustodyEventsWhitelisted)
	if err := _IERC20CustodyEvents.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyEventsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsWithdrawnIterator struct {
	Event *IERC20CustodyEventsWithdrawn // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyEventsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyEventsWithdrawn)
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
		it.Event = new(IERC20CustodyEventsWithdrawn)
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
func (it *IERC20CustodyEventsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyEventsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyEventsWithdrawn represents a Withdrawn event raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsWithdrawn struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*IERC20CustodyEventsWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.FilterLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsWithdrawnIterator{contract: _IERC20CustodyEvents.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *IERC20CustodyEventsWithdrawn, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.WatchLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyEventsWithdrawn)
				if err := _IERC20CustodyEvents.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) ParseWithdrawn(log types.Log) (*IERC20CustodyEventsWithdrawn, error) {
	event := new(IERC20CustodyEventsWithdrawn)
	if err := _IERC20CustodyEvents.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyEventsWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsWithdrawnAndCalledIterator struct {
	Event *IERC20CustodyEventsWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyEventsWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyEventsWithdrawnAndCalled)
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
		it.Event = new(IERC20CustodyEventsWithdrawnAndCalled)
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
func (it *IERC20CustodyEventsWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyEventsWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyEventsWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsWithdrawnAndCalled struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*IERC20CustodyEventsWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsWithdrawnAndCalledIterator{contract: _IERC20CustodyEvents.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *IERC20CustodyEventsWithdrawnAndCalled, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyEventsWithdrawnAndCalled)
				if err := _IERC20CustodyEvents.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) ParseWithdrawnAndCalled(log types.Log) (*IERC20CustodyEventsWithdrawnAndCalled, error) {
	event := new(IERC20CustodyEventsWithdrawnAndCalled)
	if err := _IERC20CustodyEvents.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20CustodyEventsWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsWithdrawnAndRevertedIterator struct {
	Event *IERC20CustodyEventsWithdrawnAndReverted // Event containing the contract specifics and raw log

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
func (it *IERC20CustodyEventsWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20CustodyEventsWithdrawnAndReverted)
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
		it.Event = new(IERC20CustodyEventsWithdrawnAndReverted)
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
func (it *IERC20CustodyEventsWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20CustodyEventsWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20CustodyEventsWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the IERC20CustodyEvents contract.
type IERC20CustodyEventsWithdrawnAndReverted struct {
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
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*IERC20CustodyEventsWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IERC20CustodyEventsWithdrawnAndRevertedIterator{contract: _IERC20CustodyEvents.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *IERC20CustodyEventsWithdrawnAndReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IERC20CustodyEvents.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20CustodyEventsWithdrawnAndReverted)
				if err := _IERC20CustodyEvents.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
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
func (_IERC20CustodyEvents *IERC20CustodyEventsFilterer) ParseWithdrawnAndReverted(log types.Log) (*IERC20CustodyEventsWithdrawnAndReverted, error) {
	event := new(IERC20CustodyEventsWithdrawnAndReverted)
	if err := _IERC20CustodyEvents.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
