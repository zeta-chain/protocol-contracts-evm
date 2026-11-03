// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coreregistry

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

// CallOptions is an auto generated low-level Go binding around an user-defined struct.
type CallOptions struct {
	GasLimit        *big.Int
	IsArbitraryCall bool
}

// RevertOptions is an auto generated low-level Go binding around an user-defined struct.
type RevertOptions struct {
	RevertAddress    common.Address
	CallOnRevert     bool
	AbortAddress     common.Address
	RevertMessage    []byte
	OnRevertGasLimit *big.Int
}

// MockGatewayZEVMMetaData contains all meta data concerning the MockGatewayZEVM contract.
var MockGatewayZEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"PROTOCOL_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"WithdrawAndCallEmitted\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f80fd5b506105378061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80632722feee146100385780637b15118b1461007c575b5f80fd5b61005373735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b61008f61008a36600461019b565b610091565b005b7f712fc301a93269775bb67d50a844ab29eb58b69625ff0faaeeb042c10e166e608786868686866040516100ca96959493929190610448565b60405180910390a150505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461012b575f80fd5b919050565b5f8083601f840112610140575f80fd5b50813567ffffffffffffffff811115610157575f80fd5b60208301915083602082850101111561016e575f80fd5b9250929050565b5f60408284031215610185575f80fd5b50919050565b5f60a08284031215610185575f80fd5b5f805f805f805f60e0888a0312156101b1575f80fd5b873567ffffffffffffffff8111156101c7575f80fd5b8801601f81018a136101d7575f80fd5b803567ffffffffffffffff8111156101f1576101f16100db565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561025d5761025d6100db565b6040528181528282016020018c1015610274575f80fd5b816020840160208301375f60209282018301529850890135965061029c905060408901610108565b9450606088013567ffffffffffffffff8111156102b7575f80fd5b6102c38a828b01610130565b90955093506102d790508960808a01610175565b915060c088013567ffffffffffffffff8111156102f2575f80fd5b6102fe8a828b0161018b565b91505092959891949750929550565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8035801515811461012b575f80fd5b73ffffffffffffffffffffffffffffffffffffffff61038182610108565b16825261039060208201610354565b1515602083015273ffffffffffffffffffffffffffffffffffffffff6103b860408301610108565b1660408301525f60608201357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18336030181126103f3575f80fd5b820160208101903567ffffffffffffffff81111561040f575f80fd5b80360382131561041d575f80fd5b60a0606086015261043260a08601828461030d565b6080948501359590940194909452509092915050565b60c081525f87518060c08401528060208a0160e085015e5f60e082850101527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8201168301905073ffffffffffffffffffffffffffffffffffffffff8816602084015260e08382030160408401526104c660e08201878961030d565b8535606085015290506104db60208601610354565b1515608084015282810360a08401526104f48185610363565b999850505050505050505056fea26469706673582212207d2e053f93597a8778f1211d86892abfae7d37b30d7b6c4c6bb01117d97205c864736f6c634300081a0033",
}

// MockGatewayZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use MockGatewayZEVMMetaData.ABI instead.
var MockGatewayZEVMABI = MockGatewayZEVMMetaData.ABI

// MockGatewayZEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockGatewayZEVMMetaData.Bin instead.
var MockGatewayZEVMBin = MockGatewayZEVMMetaData.Bin

// DeployMockGatewayZEVM deploys a new Ethereum contract, binding an instance of MockGatewayZEVM to it.
func DeployMockGatewayZEVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockGatewayZEVM, error) {
	parsed, err := MockGatewayZEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockGatewayZEVMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockGatewayZEVM{MockGatewayZEVMCaller: MockGatewayZEVMCaller{contract: contract}, MockGatewayZEVMTransactor: MockGatewayZEVMTransactor{contract: contract}, MockGatewayZEVMFilterer: MockGatewayZEVMFilterer{contract: contract}}, nil
}

// MockGatewayZEVM is an auto generated Go binding around an Ethereum contract.
type MockGatewayZEVM struct {
	MockGatewayZEVMCaller     // Read-only binding to the contract
	MockGatewayZEVMTransactor // Write-only binding to the contract
	MockGatewayZEVMFilterer   // Log filterer for contract events
}

// MockGatewayZEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockGatewayZEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockGatewayZEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockGatewayZEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockGatewayZEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockGatewayZEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockGatewayZEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockGatewayZEVMSession struct {
	Contract     *MockGatewayZEVM  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockGatewayZEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockGatewayZEVMCallerSession struct {
	Contract *MockGatewayZEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MockGatewayZEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockGatewayZEVMTransactorSession struct {
	Contract     *MockGatewayZEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MockGatewayZEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockGatewayZEVMRaw struct {
	Contract *MockGatewayZEVM // Generic contract binding to access the raw methods on
}

// MockGatewayZEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockGatewayZEVMCallerRaw struct {
	Contract *MockGatewayZEVMCaller // Generic read-only contract binding to access the raw methods on
}

// MockGatewayZEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockGatewayZEVMTransactorRaw struct {
	Contract *MockGatewayZEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockGatewayZEVM creates a new instance of MockGatewayZEVM, bound to a specific deployed contract.
func NewMockGatewayZEVM(address common.Address, backend bind.ContractBackend) (*MockGatewayZEVM, error) {
	contract, err := bindMockGatewayZEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockGatewayZEVM{MockGatewayZEVMCaller: MockGatewayZEVMCaller{contract: contract}, MockGatewayZEVMTransactor: MockGatewayZEVMTransactor{contract: contract}, MockGatewayZEVMFilterer: MockGatewayZEVMFilterer{contract: contract}}, nil
}

// NewMockGatewayZEVMCaller creates a new read-only instance of MockGatewayZEVM, bound to a specific deployed contract.
func NewMockGatewayZEVMCaller(address common.Address, caller bind.ContractCaller) (*MockGatewayZEVMCaller, error) {
	contract, err := bindMockGatewayZEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockGatewayZEVMCaller{contract: contract}, nil
}

// NewMockGatewayZEVMTransactor creates a new write-only instance of MockGatewayZEVM, bound to a specific deployed contract.
func NewMockGatewayZEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*MockGatewayZEVMTransactor, error) {
	contract, err := bindMockGatewayZEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockGatewayZEVMTransactor{contract: contract}, nil
}

// NewMockGatewayZEVMFilterer creates a new log filterer instance of MockGatewayZEVM, bound to a specific deployed contract.
func NewMockGatewayZEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*MockGatewayZEVMFilterer, error) {
	contract, err := bindMockGatewayZEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockGatewayZEVMFilterer{contract: contract}, nil
}

// bindMockGatewayZEVM binds a generic wrapper to an already deployed contract.
func bindMockGatewayZEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockGatewayZEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockGatewayZEVM *MockGatewayZEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockGatewayZEVM.Contract.MockGatewayZEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockGatewayZEVM *MockGatewayZEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockGatewayZEVM.Contract.MockGatewayZEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockGatewayZEVM *MockGatewayZEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockGatewayZEVM.Contract.MockGatewayZEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockGatewayZEVM *MockGatewayZEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockGatewayZEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockGatewayZEVM *MockGatewayZEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockGatewayZEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockGatewayZEVM *MockGatewayZEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockGatewayZEVM.Contract.contract.Transact(opts, method, params...)
}

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_MockGatewayZEVM *MockGatewayZEVMCaller) PROTOCOLADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockGatewayZEVM.contract.Call(opts, &out, "PROTOCOL_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_MockGatewayZEVM *MockGatewayZEVMSession) PROTOCOLADDRESS() (common.Address, error) {
	return _MockGatewayZEVM.Contract.PROTOCOLADDRESS(&_MockGatewayZEVM.CallOpts)
}

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_MockGatewayZEVM *MockGatewayZEVMCallerSession) PROTOCOLADDRESS() (common.Address, error) {
	return _MockGatewayZEVM.Contract.PROTOCOLADDRESS(&_MockGatewayZEVM.CallOpts)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 , address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_MockGatewayZEVM *MockGatewayZEVMTransactor) WithdrawAndCall(opts *bind.TransactOpts, receiver []byte, arg1 *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _MockGatewayZEVM.contract.Transact(opts, "withdrawAndCall", receiver, arg1, zrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 , address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_MockGatewayZEVM *MockGatewayZEVMSession) WithdrawAndCall(receiver []byte, arg1 *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _MockGatewayZEVM.Contract.WithdrawAndCall(&_MockGatewayZEVM.TransactOpts, receiver, arg1, zrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 , address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_MockGatewayZEVM *MockGatewayZEVMTransactorSession) WithdrawAndCall(receiver []byte, arg1 *big.Int, zrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _MockGatewayZEVM.Contract.WithdrawAndCall(&_MockGatewayZEVM.TransactOpts, receiver, arg1, zrc20, message, callOptions, revertOptions)
}

// MockGatewayZEVMWithdrawAndCallEmittedIterator is returned from FilterWithdrawAndCallEmitted and is used to iterate over the raw logs and unpacked data for WithdrawAndCallEmitted events raised by the MockGatewayZEVM contract.
type MockGatewayZEVMWithdrawAndCallEmittedIterator struct {
	Event *MockGatewayZEVMWithdrawAndCallEmitted // Event containing the contract specifics and raw log

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
func (it *MockGatewayZEVMWithdrawAndCallEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockGatewayZEVMWithdrawAndCallEmitted)
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
		it.Event = new(MockGatewayZEVMWithdrawAndCallEmitted)
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
func (it *MockGatewayZEVMWithdrawAndCallEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockGatewayZEVMWithdrawAndCallEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockGatewayZEVMWithdrawAndCallEmitted represents a WithdrawAndCallEmitted event raised by the MockGatewayZEVM contract.
type MockGatewayZEVMWithdrawAndCallEmitted struct {
	Receiver      []byte
	Zrc20         common.Address
	Message       []byte
	CallOptions   CallOptions
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawAndCallEmitted is a free log retrieval operation binding the contract event 0x712fc301a93269775bb67d50a844ab29eb58b69625ff0faaeeb042c10e166e60.
//
// Solidity: event WithdrawAndCallEmitted(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_MockGatewayZEVM *MockGatewayZEVMFilterer) FilterWithdrawAndCallEmitted(opts *bind.FilterOpts) (*MockGatewayZEVMWithdrawAndCallEmittedIterator, error) {

	logs, sub, err := _MockGatewayZEVM.contract.FilterLogs(opts, "WithdrawAndCallEmitted")
	if err != nil {
		return nil, err
	}
	return &MockGatewayZEVMWithdrawAndCallEmittedIterator{contract: _MockGatewayZEVM.contract, event: "WithdrawAndCallEmitted", logs: logs, sub: sub}, nil
}

// WatchWithdrawAndCallEmitted is a free log subscription operation binding the contract event 0x712fc301a93269775bb67d50a844ab29eb58b69625ff0faaeeb042c10e166e60.
//
// Solidity: event WithdrawAndCallEmitted(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_MockGatewayZEVM *MockGatewayZEVMFilterer) WatchWithdrawAndCallEmitted(opts *bind.WatchOpts, sink chan<- *MockGatewayZEVMWithdrawAndCallEmitted) (event.Subscription, error) {

	logs, sub, err := _MockGatewayZEVM.contract.WatchLogs(opts, "WithdrawAndCallEmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockGatewayZEVMWithdrawAndCallEmitted)
				if err := _MockGatewayZEVM.contract.UnpackLog(event, "WithdrawAndCallEmitted", log); err != nil {
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

// ParseWithdrawAndCallEmitted is a log parse operation binding the contract event 0x712fc301a93269775bb67d50a844ab29eb58b69625ff0faaeeb042c10e166e60.
//
// Solidity: event WithdrawAndCallEmitted(bytes receiver, address zrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_MockGatewayZEVM *MockGatewayZEVMFilterer) ParseWithdrawAndCallEmitted(log types.Log) (*MockGatewayZEVMWithdrawAndCallEmitted, error) {
	event := new(MockGatewayZEVMWithdrawAndCallEmitted)
	if err := _MockGatewayZEVM.contract.UnpackLog(event, "WithdrawAndCallEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
