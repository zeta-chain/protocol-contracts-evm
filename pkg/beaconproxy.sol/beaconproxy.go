// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package beaconproxy

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

// BeaconProxyMetaData contains all meta data concerning the BeaconProxy contract.
var BeaconProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidBeacon\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]}]",
	Bin: "0x60a060405260405161059d38038061059d83398101604081905261002291610354565b61002c828261003e565b506001600160a01b0316608052610445565b610047826100fb565b6040516001600160a01b038316907f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e905f90a28051156100ef576100ea826001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100e49190610416565b82610209565b505050565b6100f761027c565b5050565b806001600160a01b03163b5f0361013557604051631933b43b60e21b81526001600160a01b03821660048201526024015b60405180910390fd5b807fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5080546001600160a01b0319166001600160a01b0392831617905560408051635c60da1b60e01b815290515f92841691635c60da1b9160048083019260209291908290030181865afa1580156101ae573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101d29190610416565b9050806001600160a01b03163b5f036100f757604051634c9c8ce360e01b81526001600160a01b038216600482015260240161012c565b60605f80846001600160a01b031684604051610225919061042f565b5f60405180830381855af49150503d805f811461025d576040519150601f19603f3d011682016040523d82523d5f602084013e610262565b606091505b50909250905061027385838361029d565b95945050505050565b341561029b5760405163b398979f60e01b815260040160405180910390fd5b565b6060826102b2576102ad826102fc565b6102f5565b81511580156102c957506001600160a01b0384163b155b156102f257604051639996b31560e01b81526001600160a01b038516600482015260240161012c565b50805b9392505050565b80511561030c5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811461033b575f80fd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f8060408385031215610365575f80fd5b61036e83610325565b60208401519092506001600160401b03811115610389575f80fd5b8301601f81018513610399575f80fd5b80516001600160401b038111156103b2576103b2610340565b604051601f8201601f19908116603f011681016001600160401b03811182821017156103e0576103e0610340565b6040528181528282016020018710156103f7575f80fd5b8160208401602083015e5f602083830101528093505050509250929050565b5f60208284031215610426575f80fd5b6102f582610325565b5f82518060208501845e5f920191825250919050565b60805161014161045c5f395f602301526101415ff3fe608060405261000c61000e565b005b61001e610019610020565b6100b3565b565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561008a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100ae91906100d1565b905090565b365f80375f80365f845af43d5f803e8080156100cd573d5ff35b3d5ffd5b5f602082840312156100e1575f80fd5b815173ffffffffffffffffffffffffffffffffffffffff81168114610104575f80fd5b939250505056fea2646970667358221220e2bd6184836185cdf4783b293cfd1a94001b0b622d04c0a4ab615b47fb6236a564736f6c634300081a0033",
}

// BeaconProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use BeaconProxyMetaData.ABI instead.
var BeaconProxyABI = BeaconProxyMetaData.ABI

// BeaconProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BeaconProxyMetaData.Bin instead.
var BeaconProxyBin = BeaconProxyMetaData.Bin

// DeployBeaconProxy deploys a new Ethereum contract, binding an instance of BeaconProxy to it.
func DeployBeaconProxy(auth *bind.TransactOpts, backend bind.ContractBackend, beacon common.Address, data []byte) (common.Address, *types.Transaction, *BeaconProxy, error) {
	parsed, err := BeaconProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BeaconProxyBin), backend, beacon, data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BeaconProxy{BeaconProxyCaller: BeaconProxyCaller{contract: contract}, BeaconProxyTransactor: BeaconProxyTransactor{contract: contract}, BeaconProxyFilterer: BeaconProxyFilterer{contract: contract}}, nil
}

// BeaconProxy is an auto generated Go binding around an Ethereum contract.
type BeaconProxy struct {
	BeaconProxyCaller     // Read-only binding to the contract
	BeaconProxyTransactor // Write-only binding to the contract
	BeaconProxyFilterer   // Log filterer for contract events
}

// BeaconProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeaconProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeaconProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeaconProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeaconProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeaconProxySession struct {
	Contract     *BeaconProxy      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeaconProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeaconProxyCallerSession struct {
	Contract *BeaconProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BeaconProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeaconProxyTransactorSession struct {
	Contract     *BeaconProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BeaconProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeaconProxyRaw struct {
	Contract *BeaconProxy // Generic contract binding to access the raw methods on
}

// BeaconProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeaconProxyCallerRaw struct {
	Contract *BeaconProxyCaller // Generic read-only contract binding to access the raw methods on
}

// BeaconProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeaconProxyTransactorRaw struct {
	Contract *BeaconProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeaconProxy creates a new instance of BeaconProxy, bound to a specific deployed contract.
func NewBeaconProxy(address common.Address, backend bind.ContractBackend) (*BeaconProxy, error) {
	contract, err := bindBeaconProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeaconProxy{BeaconProxyCaller: BeaconProxyCaller{contract: contract}, BeaconProxyTransactor: BeaconProxyTransactor{contract: contract}, BeaconProxyFilterer: BeaconProxyFilterer{contract: contract}}, nil
}

// NewBeaconProxyCaller creates a new read-only instance of BeaconProxy, bound to a specific deployed contract.
func NewBeaconProxyCaller(address common.Address, caller bind.ContractCaller) (*BeaconProxyCaller, error) {
	contract, err := bindBeaconProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconProxyCaller{contract: contract}, nil
}

// NewBeaconProxyTransactor creates a new write-only instance of BeaconProxy, bound to a specific deployed contract.
func NewBeaconProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*BeaconProxyTransactor, error) {
	contract, err := bindBeaconProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeaconProxyTransactor{contract: contract}, nil
}

// NewBeaconProxyFilterer creates a new log filterer instance of BeaconProxy, bound to a specific deployed contract.
func NewBeaconProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*BeaconProxyFilterer, error) {
	contract, err := bindBeaconProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeaconProxyFilterer{contract: contract}, nil
}

// bindBeaconProxy binds a generic wrapper to an already deployed contract.
func bindBeaconProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeaconProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeaconProxy *BeaconProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconProxy.Contract.BeaconProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeaconProxy *BeaconProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeaconProxy.Contract.BeaconProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeaconProxy *BeaconProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeaconProxy.Contract.BeaconProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeaconProxy *BeaconProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeaconProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeaconProxy *BeaconProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeaconProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeaconProxy *BeaconProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeaconProxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BeaconProxy *BeaconProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _BeaconProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BeaconProxy *BeaconProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BeaconProxy.Contract.Fallback(&_BeaconProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_BeaconProxy *BeaconProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _BeaconProxy.Contract.Fallback(&_BeaconProxy.TransactOpts, calldata)
}

// BeaconProxyBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the BeaconProxy contract.
type BeaconProxyBeaconUpgradedIterator struct {
	Event *BeaconProxyBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BeaconProxyBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeaconProxyBeaconUpgraded)
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
		it.Event = new(BeaconProxyBeaconUpgraded)
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
func (it *BeaconProxyBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeaconProxyBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeaconProxyBeaconUpgraded represents a BeaconUpgraded event raised by the BeaconProxy contract.
type BeaconProxyBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BeaconProxy *BeaconProxyFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BeaconProxyBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BeaconProxy.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BeaconProxyBeaconUpgradedIterator{contract: _BeaconProxy.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BeaconProxy *BeaconProxyFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BeaconProxyBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BeaconProxy.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeaconProxyBeaconUpgraded)
				if err := _BeaconProxy.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_BeaconProxy *BeaconProxyFilterer) ParseBeaconUpgraded(log types.Log) (*BeaconProxyBeaconUpgraded, error) {
	event := new(BeaconProxyBeaconUpgraded)
	if err := _BeaconProxy.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
