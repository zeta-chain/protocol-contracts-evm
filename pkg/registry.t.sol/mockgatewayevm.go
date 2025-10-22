// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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
}

// MockGatewayEVMMetaData contains all meta data concerning the MockGatewayEVM contract.
var MockGatewayEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"onCall\",\"inputs\":[{\"name\":\"messageContext\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallEmitted\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"messageContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f80fd5b506102518061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c8063676cc0541461002d575b5f80fd5b61004061003b3660046100ab565b610056565b60405161004d9190610130565b60405180910390f35b60607f13e79c97a555ca94839a23ea9ea3c2702c15dd006ffae3e0cf3a166c0fb2a3bd3384848760405161008d9493929190610183565b60405180910390a15060408051602081019091525f81529392505050565b5f805f83850360408112156100be575f80fd5b60208112156100cb575f80fd5b50839250602084013567ffffffffffffffff8111156100e8575f80fd5b8401601f810186136100f8575f80fd5b803567ffffffffffffffff81111561010e575f80fd5b86602082840101111561011f575f80fd5b939660209190910195509293505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b73ffffffffffffffffffffffffffffffffffffffff8516815260606020820152826060820152828460808301375f608084830101525f60807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8601168301019050823573ffffffffffffffffffffffffffffffffffffffff811680821461020a575f80fd5b80604085015250509594505050505056fea26469706673582212207c777e42557f8cc55c9ab7922984439515806b7837066bc68aa43057ca252d8264736f6c634300081a0033",
}

// MockGatewayEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use MockGatewayEVMMetaData.ABI instead.
var MockGatewayEVMABI = MockGatewayEVMMetaData.ABI

// MockGatewayEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockGatewayEVMMetaData.Bin instead.
var MockGatewayEVMBin = MockGatewayEVMMetaData.Bin

// DeployMockGatewayEVM deploys a new Ethereum contract, binding an instance of MockGatewayEVM to it.
func DeployMockGatewayEVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockGatewayEVM, error) {
	parsed, err := MockGatewayEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockGatewayEVMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockGatewayEVM{MockGatewayEVMCaller: MockGatewayEVMCaller{contract: contract}, MockGatewayEVMTransactor: MockGatewayEVMTransactor{contract: contract}, MockGatewayEVMFilterer: MockGatewayEVMFilterer{contract: contract}}, nil
}

// MockGatewayEVM is an auto generated Go binding around an Ethereum contract.
type MockGatewayEVM struct {
	MockGatewayEVMCaller     // Read-only binding to the contract
	MockGatewayEVMTransactor // Write-only binding to the contract
	MockGatewayEVMFilterer   // Log filterer for contract events
}

// MockGatewayEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockGatewayEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockGatewayEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockGatewayEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockGatewayEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockGatewayEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockGatewayEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockGatewayEVMSession struct {
	Contract     *MockGatewayEVM   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockGatewayEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockGatewayEVMCallerSession struct {
	Contract *MockGatewayEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MockGatewayEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockGatewayEVMTransactorSession struct {
	Contract     *MockGatewayEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MockGatewayEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockGatewayEVMRaw struct {
	Contract *MockGatewayEVM // Generic contract binding to access the raw methods on
}

// MockGatewayEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockGatewayEVMCallerRaw struct {
	Contract *MockGatewayEVMCaller // Generic read-only contract binding to access the raw methods on
}

// MockGatewayEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockGatewayEVMTransactorRaw struct {
	Contract *MockGatewayEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockGatewayEVM creates a new instance of MockGatewayEVM, bound to a specific deployed contract.
func NewMockGatewayEVM(address common.Address, backend bind.ContractBackend) (*MockGatewayEVM, error) {
	contract, err := bindMockGatewayEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockGatewayEVM{MockGatewayEVMCaller: MockGatewayEVMCaller{contract: contract}, MockGatewayEVMTransactor: MockGatewayEVMTransactor{contract: contract}, MockGatewayEVMFilterer: MockGatewayEVMFilterer{contract: contract}}, nil
}

// NewMockGatewayEVMCaller creates a new read-only instance of MockGatewayEVM, bound to a specific deployed contract.
func NewMockGatewayEVMCaller(address common.Address, caller bind.ContractCaller) (*MockGatewayEVMCaller, error) {
	contract, err := bindMockGatewayEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockGatewayEVMCaller{contract: contract}, nil
}

// NewMockGatewayEVMTransactor creates a new write-only instance of MockGatewayEVM, bound to a specific deployed contract.
func NewMockGatewayEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*MockGatewayEVMTransactor, error) {
	contract, err := bindMockGatewayEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockGatewayEVMTransactor{contract: contract}, nil
}

// NewMockGatewayEVMFilterer creates a new log filterer instance of MockGatewayEVM, bound to a specific deployed contract.
func NewMockGatewayEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*MockGatewayEVMFilterer, error) {
	contract, err := bindMockGatewayEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockGatewayEVMFilterer{contract: contract}, nil
}

// bindMockGatewayEVM binds a generic wrapper to an already deployed contract.
func bindMockGatewayEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockGatewayEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockGatewayEVM *MockGatewayEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockGatewayEVM.Contract.MockGatewayEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockGatewayEVM *MockGatewayEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockGatewayEVM.Contract.MockGatewayEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockGatewayEVM *MockGatewayEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockGatewayEVM.Contract.MockGatewayEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockGatewayEVM *MockGatewayEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockGatewayEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockGatewayEVM *MockGatewayEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockGatewayEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockGatewayEVM *MockGatewayEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockGatewayEVM.Contract.contract.Transact(opts, method, params...)
}

// OnCall is a paid mutator transaction binding the contract method 0x676cc054.
//
// Solidity: function onCall((address) messageContext, bytes data) returns(bytes)
func (_MockGatewayEVM *MockGatewayEVMTransactor) OnCall(opts *bind.TransactOpts, messageContext MessageContext, data []byte) (*types.Transaction, error) {
	return _MockGatewayEVM.contract.Transact(opts, "onCall", messageContext, data)
}

// OnCall is a paid mutator transaction binding the contract method 0x676cc054.
//
// Solidity: function onCall((address) messageContext, bytes data) returns(bytes)
func (_MockGatewayEVM *MockGatewayEVMSession) OnCall(messageContext MessageContext, data []byte) (*types.Transaction, error) {
	return _MockGatewayEVM.Contract.OnCall(&_MockGatewayEVM.TransactOpts, messageContext, data)
}

// OnCall is a paid mutator transaction binding the contract method 0x676cc054.
//
// Solidity: function onCall((address) messageContext, bytes data) returns(bytes)
func (_MockGatewayEVM *MockGatewayEVMTransactorSession) OnCall(messageContext MessageContext, data []byte) (*types.Transaction, error) {
	return _MockGatewayEVM.Contract.OnCall(&_MockGatewayEVM.TransactOpts, messageContext, data)
}

// MockGatewayEVMCallEmittedIterator is returned from FilterCallEmitted and is used to iterate over the raw logs and unpacked data for CallEmitted events raised by the MockGatewayEVM contract.
type MockGatewayEVMCallEmittedIterator struct {
	Event *MockGatewayEVMCallEmitted // Event containing the contract specifics and raw log

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
func (it *MockGatewayEVMCallEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockGatewayEVMCallEmitted)
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
		it.Event = new(MockGatewayEVMCallEmitted)
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
func (it *MockGatewayEVMCallEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockGatewayEVMCallEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockGatewayEVMCallEmitted represents a CallEmitted event raised by the MockGatewayEVM contract.
type MockGatewayEVMCallEmitted struct {
	Receiver       common.Address
	Message        []byte
	MessageContext MessageContext
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCallEmitted is a free log retrieval operation binding the contract event 0x13e79c97a555ca94839a23ea9ea3c2702c15dd006ffae3e0cf3a166c0fb2a3bd.
//
// Solidity: event CallEmitted(address receiver, bytes message, (address) messageContext)
func (_MockGatewayEVM *MockGatewayEVMFilterer) FilterCallEmitted(opts *bind.FilterOpts) (*MockGatewayEVMCallEmittedIterator, error) {

	logs, sub, err := _MockGatewayEVM.contract.FilterLogs(opts, "CallEmitted")
	if err != nil {
		return nil, err
	}
	return &MockGatewayEVMCallEmittedIterator{contract: _MockGatewayEVM.contract, event: "CallEmitted", logs: logs, sub: sub}, nil
}

// WatchCallEmitted is a free log subscription operation binding the contract event 0x13e79c97a555ca94839a23ea9ea3c2702c15dd006ffae3e0cf3a166c0fb2a3bd.
//
// Solidity: event CallEmitted(address receiver, bytes message, (address) messageContext)
func (_MockGatewayEVM *MockGatewayEVMFilterer) WatchCallEmitted(opts *bind.WatchOpts, sink chan<- *MockGatewayEVMCallEmitted) (event.Subscription, error) {

	logs, sub, err := _MockGatewayEVM.contract.WatchLogs(opts, "CallEmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockGatewayEVMCallEmitted)
				if err := _MockGatewayEVM.contract.UnpackLog(event, "CallEmitted", log); err != nil {
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

// ParseCallEmitted is a log parse operation binding the contract event 0x13e79c97a555ca94839a23ea9ea3c2702c15dd006ffae3e0cf3a166c0fb2a3bd.
//
// Solidity: event CallEmitted(address receiver, bytes message, (address) messageContext)
func (_MockGatewayEVM *MockGatewayEVMFilterer) ParseCallEmitted(log types.Log) (*MockGatewayEVMCallEmitted, error) {
	event := new(MockGatewayEVMCallEmitted)
	if err := _MockGatewayEVM.contract.UnpackLog(event, "CallEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
