// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package upgradeablebeacon

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

// UpgradeableBeaconMetaData contains all meta data concerning the UpgradeableBeacon contract.
var UpgradeableBeaconMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"implementation_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BeaconInvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161054538038061054583398101604081905261002f91610165565b806001600160a01b03811661005f57604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b61006881610079565b50610072826100c9565b5050610198565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b806001600160a01b03163b6000036100ff5760405163211eb15960e21b81526001600160a01b0382166004820152602401610056565b600180546001600160a01b0319166001600160a01b0383169081179091556040517fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b80516001600160a01b038116811461016057600080fd5b919050565b6000806040838503121561017857600080fd5b61018183610149565b915061018f60208401610149565b90509250929050565b61039e806101a76000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063715018a611610050578063715018a6146100c45780638da5cb5b146100cc578063f2fde38b146100ea57600080fd5b80633659cfe61461006c5780635c60da1b14610081575b600080fd5b61007f61007a36600461032b565b6100fd565b005b60015473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b61007f610111565b60005473ffffffffffffffffffffffffffffffffffffffff1661009b565b61007f6100f836600461032b565b610125565b61010561018b565b61010e816101de565b50565b61011961018b565b61012360006102b6565b565b61012d61018b565b73ffffffffffffffffffffffffffffffffffffffff8116610182576040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600060048201526024015b60405180910390fd5b61010e816102b6565b60005473ffffffffffffffffffffffffffffffffffffffff163314610123576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610179565b8073ffffffffffffffffffffffffffffffffffffffff163b600003610247576040517f847ac56400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610179565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561033d57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461036157600080fd5b939250505056fea2646970667358221220d0b48b2ca205e02c363d927b12afa619fc0b60e74d3acae0413fda2eb832329464736f6c634300081a0033",
}

// UpgradeableBeaconABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradeableBeaconMetaData.ABI instead.
var UpgradeableBeaconABI = UpgradeableBeaconMetaData.ABI

// UpgradeableBeaconBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UpgradeableBeaconMetaData.Bin instead.
var UpgradeableBeaconBin = UpgradeableBeaconMetaData.Bin

// DeployUpgradeableBeacon deploys a new Ethereum contract, binding an instance of UpgradeableBeacon to it.
func DeployUpgradeableBeacon(auth *bind.TransactOpts, backend bind.ContractBackend, implementation_ common.Address, initialOwner common.Address) (common.Address, *types.Transaction, *UpgradeableBeacon, error) {
	parsed, err := UpgradeableBeaconMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UpgradeableBeaconBin), backend, implementation_, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UpgradeableBeacon{UpgradeableBeaconCaller: UpgradeableBeaconCaller{contract: contract}, UpgradeableBeaconTransactor: UpgradeableBeaconTransactor{contract: contract}, UpgradeableBeaconFilterer: UpgradeableBeaconFilterer{contract: contract}}, nil
}

// UpgradeableBeacon is an auto generated Go binding around an Ethereum contract.
type UpgradeableBeacon struct {
	UpgradeableBeaconCaller     // Read-only binding to the contract
	UpgradeableBeaconTransactor // Write-only binding to the contract
	UpgradeableBeaconFilterer   // Log filterer for contract events
}

// UpgradeableBeaconCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeableBeaconCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableBeaconTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeableBeaconTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableBeaconFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeableBeaconFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableBeaconSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeableBeaconSession struct {
	Contract     *UpgradeableBeacon // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UpgradeableBeaconCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeableBeaconCallerSession struct {
	Contract *UpgradeableBeaconCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UpgradeableBeaconTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeableBeaconTransactorSession struct {
	Contract     *UpgradeableBeaconTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UpgradeableBeaconRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeableBeaconRaw struct {
	Contract *UpgradeableBeacon // Generic contract binding to access the raw methods on
}

// UpgradeableBeaconCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeableBeaconCallerRaw struct {
	Contract *UpgradeableBeaconCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeableBeaconTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeableBeaconTransactorRaw struct {
	Contract *UpgradeableBeaconTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeableBeacon creates a new instance of UpgradeableBeacon, bound to a specific deployed contract.
func NewUpgradeableBeacon(address common.Address, backend bind.ContractBackend) (*UpgradeableBeacon, error) {
	contract, err := bindUpgradeableBeacon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeableBeacon{UpgradeableBeaconCaller: UpgradeableBeaconCaller{contract: contract}, UpgradeableBeaconTransactor: UpgradeableBeaconTransactor{contract: contract}, UpgradeableBeaconFilterer: UpgradeableBeaconFilterer{contract: contract}}, nil
}

// NewUpgradeableBeaconCaller creates a new read-only instance of UpgradeableBeacon, bound to a specific deployed contract.
func NewUpgradeableBeaconCaller(address common.Address, caller bind.ContractCaller) (*UpgradeableBeaconCaller, error) {
	contract, err := bindUpgradeableBeacon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableBeaconCaller{contract: contract}, nil
}

// NewUpgradeableBeaconTransactor creates a new write-only instance of UpgradeableBeacon, bound to a specific deployed contract.
func NewUpgradeableBeaconTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeableBeaconTransactor, error) {
	contract, err := bindUpgradeableBeacon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableBeaconTransactor{contract: contract}, nil
}

// NewUpgradeableBeaconFilterer creates a new log filterer instance of UpgradeableBeacon, bound to a specific deployed contract.
func NewUpgradeableBeaconFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeableBeaconFilterer, error) {
	contract, err := bindUpgradeableBeacon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeableBeaconFilterer{contract: contract}, nil
}

// bindUpgradeableBeacon binds a generic wrapper to an already deployed contract.
func bindUpgradeableBeacon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpgradeableBeaconMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeableBeacon *UpgradeableBeaconRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeableBeacon.Contract.UpgradeableBeaconCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeableBeacon *UpgradeableBeaconRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableBeacon.Contract.UpgradeableBeaconTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeableBeacon *UpgradeableBeaconRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeableBeacon.Contract.UpgradeableBeaconTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeableBeacon *UpgradeableBeaconCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeableBeacon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeableBeacon *UpgradeableBeaconTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableBeacon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeableBeacon *UpgradeableBeaconTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeableBeacon.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_UpgradeableBeacon *UpgradeableBeaconCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpgradeableBeacon.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_UpgradeableBeacon *UpgradeableBeaconSession) Implementation() (common.Address, error) {
	return _UpgradeableBeacon.Contract.Implementation(&_UpgradeableBeacon.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_UpgradeableBeacon *UpgradeableBeaconCallerSession) Implementation() (common.Address, error) {
	return _UpgradeableBeacon.Contract.Implementation(&_UpgradeableBeacon.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradeableBeacon *UpgradeableBeaconCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpgradeableBeacon.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradeableBeacon *UpgradeableBeaconSession) Owner() (common.Address, error) {
	return _UpgradeableBeacon.Contract.Owner(&_UpgradeableBeacon.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradeableBeacon *UpgradeableBeaconCallerSession) Owner() (common.Address, error) {
	return _UpgradeableBeacon.Contract.Owner(&_UpgradeableBeacon.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradeableBeacon *UpgradeableBeaconTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableBeacon.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradeableBeacon *UpgradeableBeaconSession) RenounceOwnership() (*types.Transaction, error) {
	return _UpgradeableBeacon.Contract.RenounceOwnership(&_UpgradeableBeacon.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradeableBeacon *UpgradeableBeaconTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UpgradeableBeacon.Contract.RenounceOwnership(&_UpgradeableBeacon.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradeableBeacon *UpgradeableBeaconTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UpgradeableBeacon.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradeableBeacon *UpgradeableBeaconSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UpgradeableBeacon.Contract.TransferOwnership(&_UpgradeableBeacon.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradeableBeacon *UpgradeableBeaconTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UpgradeableBeacon.Contract.TransferOwnership(&_UpgradeableBeacon.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UpgradeableBeacon *UpgradeableBeaconTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _UpgradeableBeacon.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UpgradeableBeacon *UpgradeableBeaconSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UpgradeableBeacon.Contract.UpgradeTo(&_UpgradeableBeacon.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UpgradeableBeacon *UpgradeableBeaconTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UpgradeableBeacon.Contract.UpgradeTo(&_UpgradeableBeacon.TransactOpts, newImplementation)
}

// UpgradeableBeaconOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UpgradeableBeacon contract.
type UpgradeableBeaconOwnershipTransferredIterator struct {
	Event *UpgradeableBeaconOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UpgradeableBeaconOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableBeaconOwnershipTransferred)
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
		it.Event = new(UpgradeableBeaconOwnershipTransferred)
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
func (it *UpgradeableBeaconOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableBeaconOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableBeaconOwnershipTransferred represents a OwnershipTransferred event raised by the UpgradeableBeacon contract.
type UpgradeableBeaconOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradeableBeacon *UpgradeableBeaconFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UpgradeableBeaconOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UpgradeableBeacon.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableBeaconOwnershipTransferredIterator{contract: _UpgradeableBeacon.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradeableBeacon *UpgradeableBeaconFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UpgradeableBeaconOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UpgradeableBeacon.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableBeaconOwnershipTransferred)
				if err := _UpgradeableBeacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradeableBeacon *UpgradeableBeaconFilterer) ParseOwnershipTransferred(log types.Log) (*UpgradeableBeaconOwnershipTransferred, error) {
	event := new(UpgradeableBeaconOwnershipTransferred)
	if err := _UpgradeableBeacon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableBeaconUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UpgradeableBeacon contract.
type UpgradeableBeaconUpgradedIterator struct {
	Event *UpgradeableBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *UpgradeableBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableBeaconUpgraded)
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
		it.Event = new(UpgradeableBeaconUpgraded)
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
func (it *UpgradeableBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableBeaconUpgraded represents a Upgraded event raised by the UpgradeableBeacon contract.
type UpgradeableBeaconUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeableBeacon *UpgradeableBeaconFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UpgradeableBeaconUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UpgradeableBeacon.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableBeaconUpgradedIterator{contract: _UpgradeableBeacon.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UpgradeableBeacon *UpgradeableBeaconFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UpgradeableBeaconUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UpgradeableBeacon.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableBeaconUpgraded)
				if err := _UpgradeableBeacon.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_UpgradeableBeacon *UpgradeableBeaconFilterer) ParseUpgraded(log types.Log) (*UpgradeableBeaconUpgraded, error) {
	event := new(UpgradeableBeaconUpgraded)
	if err := _UpgradeableBeacon.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
