// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stdstorage

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

// StdStorageSafeMetaData contains all meta data concerning the StdStorageSafe contract.
var StdStorageSafeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"SlotFound\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"fsig\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"keysHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"slot\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WARNING_UninitedSlot\",\"inputs\":[{\"name\":\"who\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b9ae532ec0e0263d093c73fe11e27f7c536ea19536560fc3285dfbb15cceacc764736f6c634300081a0033",
}

// StdStorageSafeABI is the input ABI used to generate the binding from.
// Deprecated: Use StdStorageSafeMetaData.ABI instead.
var StdStorageSafeABI = StdStorageSafeMetaData.ABI

// StdStorageSafeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StdStorageSafeMetaData.Bin instead.
var StdStorageSafeBin = StdStorageSafeMetaData.Bin

// DeployStdStorageSafe deploys a new Ethereum contract, binding an instance of StdStorageSafe to it.
func DeployStdStorageSafe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StdStorageSafe, error) {
	parsed, err := StdStorageSafeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StdStorageSafeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StdStorageSafe{StdStorageSafeCaller: StdStorageSafeCaller{contract: contract}, StdStorageSafeTransactor: StdStorageSafeTransactor{contract: contract}, StdStorageSafeFilterer: StdStorageSafeFilterer{contract: contract}}, nil
}

// StdStorageSafe is an auto generated Go binding around an Ethereum contract.
type StdStorageSafe struct {
	StdStorageSafeCaller     // Read-only binding to the contract
	StdStorageSafeTransactor // Write-only binding to the contract
	StdStorageSafeFilterer   // Log filterer for contract events
}

// StdStorageSafeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdStorageSafeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdStorageSafeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdStorageSafeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdStorageSafeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdStorageSafeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdStorageSafeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdStorageSafeSession struct {
	Contract     *StdStorageSafe   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdStorageSafeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdStorageSafeCallerSession struct {
	Contract *StdStorageSafeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StdStorageSafeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdStorageSafeTransactorSession struct {
	Contract     *StdStorageSafeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StdStorageSafeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdStorageSafeRaw struct {
	Contract *StdStorageSafe // Generic contract binding to access the raw methods on
}

// StdStorageSafeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdStorageSafeCallerRaw struct {
	Contract *StdStorageSafeCaller // Generic read-only contract binding to access the raw methods on
}

// StdStorageSafeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdStorageSafeTransactorRaw struct {
	Contract *StdStorageSafeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdStorageSafe creates a new instance of StdStorageSafe, bound to a specific deployed contract.
func NewStdStorageSafe(address common.Address, backend bind.ContractBackend) (*StdStorageSafe, error) {
	contract, err := bindStdStorageSafe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StdStorageSafe{StdStorageSafeCaller: StdStorageSafeCaller{contract: contract}, StdStorageSafeTransactor: StdStorageSafeTransactor{contract: contract}, StdStorageSafeFilterer: StdStorageSafeFilterer{contract: contract}}, nil
}

// NewStdStorageSafeCaller creates a new read-only instance of StdStorageSafe, bound to a specific deployed contract.
func NewStdStorageSafeCaller(address common.Address, caller bind.ContractCaller) (*StdStorageSafeCaller, error) {
	contract, err := bindStdStorageSafe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdStorageSafeCaller{contract: contract}, nil
}

// NewStdStorageSafeTransactor creates a new write-only instance of StdStorageSafe, bound to a specific deployed contract.
func NewStdStorageSafeTransactor(address common.Address, transactor bind.ContractTransactor) (*StdStorageSafeTransactor, error) {
	contract, err := bindStdStorageSafe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdStorageSafeTransactor{contract: contract}, nil
}

// NewStdStorageSafeFilterer creates a new log filterer instance of StdStorageSafe, bound to a specific deployed contract.
func NewStdStorageSafeFilterer(address common.Address, filterer bind.ContractFilterer) (*StdStorageSafeFilterer, error) {
	contract, err := bindStdStorageSafe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdStorageSafeFilterer{contract: contract}, nil
}

// bindStdStorageSafe binds a generic wrapper to an already deployed contract.
func bindStdStorageSafe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StdStorageSafeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdStorageSafe *StdStorageSafeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdStorageSafe.Contract.StdStorageSafeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdStorageSafe *StdStorageSafeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdStorageSafe.Contract.StdStorageSafeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdStorageSafe *StdStorageSafeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdStorageSafe.Contract.StdStorageSafeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdStorageSafe *StdStorageSafeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdStorageSafe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdStorageSafe *StdStorageSafeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdStorageSafe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdStorageSafe *StdStorageSafeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdStorageSafe.Contract.contract.Transact(opts, method, params...)
}

// StdStorageSafeSlotFoundIterator is returned from FilterSlotFound and is used to iterate over the raw logs and unpacked data for SlotFound events raised by the StdStorageSafe contract.
type StdStorageSafeSlotFoundIterator struct {
	Event *StdStorageSafeSlotFound // Event containing the contract specifics and raw log

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
func (it *StdStorageSafeSlotFoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StdStorageSafeSlotFound)
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
		it.Event = new(StdStorageSafeSlotFound)
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
func (it *StdStorageSafeSlotFoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StdStorageSafeSlotFoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StdStorageSafeSlotFound represents a SlotFound event raised by the StdStorageSafe contract.
type StdStorageSafeSlotFound struct {
	Who      common.Address
	Fsig     [4]byte
	KeysHash [32]byte
	Slot     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSlotFound is a free log retrieval operation binding the contract event 0x9c9555b1e3102e3cf48f427d79cb678f5d9bd1ed0ad574389461e255f95170ed.
//
// Solidity: event SlotFound(address who, bytes4 fsig, bytes32 keysHash, uint256 slot)
func (_StdStorageSafe *StdStorageSafeFilterer) FilterSlotFound(opts *bind.FilterOpts) (*StdStorageSafeSlotFoundIterator, error) {

	logs, sub, err := _StdStorageSafe.contract.FilterLogs(opts, "SlotFound")
	if err != nil {
		return nil, err
	}
	return &StdStorageSafeSlotFoundIterator{contract: _StdStorageSafe.contract, event: "SlotFound", logs: logs, sub: sub}, nil
}

// WatchSlotFound is a free log subscription operation binding the contract event 0x9c9555b1e3102e3cf48f427d79cb678f5d9bd1ed0ad574389461e255f95170ed.
//
// Solidity: event SlotFound(address who, bytes4 fsig, bytes32 keysHash, uint256 slot)
func (_StdStorageSafe *StdStorageSafeFilterer) WatchSlotFound(opts *bind.WatchOpts, sink chan<- *StdStorageSafeSlotFound) (event.Subscription, error) {

	logs, sub, err := _StdStorageSafe.contract.WatchLogs(opts, "SlotFound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StdStorageSafeSlotFound)
				if err := _StdStorageSafe.contract.UnpackLog(event, "SlotFound", log); err != nil {
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

// ParseSlotFound is a log parse operation binding the contract event 0x9c9555b1e3102e3cf48f427d79cb678f5d9bd1ed0ad574389461e255f95170ed.
//
// Solidity: event SlotFound(address who, bytes4 fsig, bytes32 keysHash, uint256 slot)
func (_StdStorageSafe *StdStorageSafeFilterer) ParseSlotFound(log types.Log) (*StdStorageSafeSlotFound, error) {
	event := new(StdStorageSafeSlotFound)
	if err := _StdStorageSafe.contract.UnpackLog(event, "SlotFound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StdStorageSafeWARNINGUninitedSlotIterator is returned from FilterWARNINGUninitedSlot and is used to iterate over the raw logs and unpacked data for WARNINGUninitedSlot events raised by the StdStorageSafe contract.
type StdStorageSafeWARNINGUninitedSlotIterator struct {
	Event *StdStorageSafeWARNINGUninitedSlot // Event containing the contract specifics and raw log

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
func (it *StdStorageSafeWARNINGUninitedSlotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StdStorageSafeWARNINGUninitedSlot)
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
		it.Event = new(StdStorageSafeWARNINGUninitedSlot)
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
func (it *StdStorageSafeWARNINGUninitedSlotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StdStorageSafeWARNINGUninitedSlotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StdStorageSafeWARNINGUninitedSlot represents a WARNINGUninitedSlot event raised by the StdStorageSafe contract.
type StdStorageSafeWARNINGUninitedSlot struct {
	Who  common.Address
	Slot *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWARNINGUninitedSlot is a free log retrieval operation binding the contract event 0x080fc4a96620c4462e705b23f346413fe3796bb63c6f8d8591baec0e231577a5.
//
// Solidity: event WARNING_UninitedSlot(address who, uint256 slot)
func (_StdStorageSafe *StdStorageSafeFilterer) FilterWARNINGUninitedSlot(opts *bind.FilterOpts) (*StdStorageSafeWARNINGUninitedSlotIterator, error) {

	logs, sub, err := _StdStorageSafe.contract.FilterLogs(opts, "WARNING_UninitedSlot")
	if err != nil {
		return nil, err
	}
	return &StdStorageSafeWARNINGUninitedSlotIterator{contract: _StdStorageSafe.contract, event: "WARNING_UninitedSlot", logs: logs, sub: sub}, nil
}

// WatchWARNINGUninitedSlot is a free log subscription operation binding the contract event 0x080fc4a96620c4462e705b23f346413fe3796bb63c6f8d8591baec0e231577a5.
//
// Solidity: event WARNING_UninitedSlot(address who, uint256 slot)
func (_StdStorageSafe *StdStorageSafeFilterer) WatchWARNINGUninitedSlot(opts *bind.WatchOpts, sink chan<- *StdStorageSafeWARNINGUninitedSlot) (event.Subscription, error) {

	logs, sub, err := _StdStorageSafe.contract.WatchLogs(opts, "WARNING_UninitedSlot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StdStorageSafeWARNINGUninitedSlot)
				if err := _StdStorageSafe.contract.UnpackLog(event, "WARNING_UninitedSlot", log); err != nil {
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

// ParseWARNINGUninitedSlot is a log parse operation binding the contract event 0x080fc4a96620c4462e705b23f346413fe3796bb63c6f8d8591baec0e231577a5.
//
// Solidity: event WARNING_UninitedSlot(address who, uint256 slot)
func (_StdStorageSafe *StdStorageSafeFilterer) ParseWARNINGUninitedSlot(log types.Log) (*StdStorageSafeWARNINGUninitedSlot, error) {
	event := new(StdStorageSafeWARNINGUninitedSlot)
	if err := _StdStorageSafe.contract.UnpackLog(event, "WARNING_UninitedSlot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
