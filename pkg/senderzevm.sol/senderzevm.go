// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package senderzevm

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

// SenderZEVMMetaData contains all meta data concerning the SenderZEVM contract.
var SenderZEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"gateway_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"callReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawAndCallReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f80fd5b50604051610a57380380610a57833981016040819052602b91604e565b5f80546001600160a01b0319166001600160a01b03929092169190911790556079565b5f60208284031215605d575f80fd5b81516001600160a01b03811681146072575f80fd5b9392505050565b6109d1806100865f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80630abd890514610043578063116191b6146100585780637a34d8bb146100a0575b5f80fd5b610056610051366004610639565b6100b3565b005b5f546100779073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100566100ae3660046106d3565b610320565b5f8383836040516024016100c9939291906107b1565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f97000000000000000000000000000000000000000000000000000000001790525f5490915073ffffffffffffffffffffffffffffffffffffffff8087169163095ea7b3911661017189620186a06107da565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303815f875af11580156101de573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102029190610818565b610238576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160a081018252610321808252600160208084019190915282840191909152825180820184525f80825260608401919091526080830181905283518085018552620186a081529182018190525492517f7b15118b0000000000000000000000000000000000000000000000000000000081529192909173ffffffffffffffffffffffffffffffffffffffff90911690637b15118b906102e8908c908c908c90899088908a906004016108ae565b5f604051808303815f87803b1580156102ff575f80fd5b505af1158015610311573d5f803e3d5ffd5b50505050505050505050505050565b5f838383604051602401610336939291906107b1565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f9700000000000000000000000000000000000000000000000000000000179052815160a08101835261032180825260018284015281840152825180830184525f80825260608301919091526080820181905283518085018552620186a0808252938101829052905493517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff94851660048201526024810193909352939450929188169063095ea7b3906044016020604051808303815f875af115801561046c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104909190610818565b505f546040517f06cb898300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116906306cb8983906104ee908b908b90889087908990600401610928565b5f604051808303815f87803b158015610505575f80fd5b505af1158015610517573d5f803e3d5ffd5b505050505050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112610561575f80fd5b8135602083015f8067ffffffffffffffff84111561058157610581610525565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156105ce576105ce610525565b6040528381529050808284018710156105e5575f80fd5b838360208301375f602085830101528094505050505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610624575f80fd5b919050565b8015158114610636575f80fd5b50565b5f805f805f8060c0878903121561064e575f80fd5b863567ffffffffffffffff811115610664575f80fd5b61067089828a01610552565b9650506020870135945061068660408801610601565b9350606087013567ffffffffffffffff8111156106a1575f80fd5b6106ad89828a01610552565b9350506080870135915060a08701356106c581610629565b809150509295509295509295565b5f805f805f60a086880312156106e7575f80fd5b853567ffffffffffffffff8111156106fd575f80fd5b61070988828901610552565b95505061071860208701610601565b9350604086013567ffffffffffffffff811115610733575f80fd5b61073f88828901610552565b93505060608601359150608086013561075781610629565b809150509295509295909350565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b606081525f6107c36060830186610765565b602083019490945250901515604090910152919050565b80820180821115610812577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b92915050565b5f60208284031215610828575f80fd5b815161083381610629565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff815116825260208101511515602083015273ffffffffffffffffffffffffffffffffffffffff60408201511660408301525f606082015160a0606085015261089a60a0850182610765565b608093840151949093019390935250919050565b60e081525f6108c060e0830189610765565b87602084015273ffffffffffffffffffffffffffffffffffffffff8716604084015282810360608401526108f48187610765565b855160808501526020860151151560a0850152905082810360c084015261091b818561083a565b9998505050505050505050565b60c081525f61093a60c0830188610765565b73ffffffffffffffffffffffffffffffffffffffff8716602084015282810360408401526109688187610765565b85516060850152602086015115156080850152905082810360a084015261098f818561083a565b9897505050505050505056fea2646970667358221220eb5b4bcc51c900188476a67d6bf40d393fbfb65f72412af3fac744d47c8d15a764736f6c634300081a0033",
}

// SenderZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use SenderZEVMMetaData.ABI instead.
var SenderZEVMABI = SenderZEVMMetaData.ABI

// SenderZEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SenderZEVMMetaData.Bin instead.
var SenderZEVMBin = SenderZEVMMetaData.Bin

// DeploySenderZEVM deploys a new Ethereum contract, binding an instance of SenderZEVM to it.
func DeploySenderZEVM(auth *bind.TransactOpts, backend bind.ContractBackend, gateway_ common.Address) (common.Address, *types.Transaction, *SenderZEVM, error) {
	parsed, err := SenderZEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SenderZEVMBin), backend, gateway_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SenderZEVM{SenderZEVMCaller: SenderZEVMCaller{contract: contract}, SenderZEVMTransactor: SenderZEVMTransactor{contract: contract}, SenderZEVMFilterer: SenderZEVMFilterer{contract: contract}}, nil
}

// SenderZEVM is an auto generated Go binding around an Ethereum contract.
type SenderZEVM struct {
	SenderZEVMCaller     // Read-only binding to the contract
	SenderZEVMTransactor // Write-only binding to the contract
	SenderZEVMFilterer   // Log filterer for contract events
}

// SenderZEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenderZEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderZEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenderZEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderZEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenderZEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderZEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenderZEVMSession struct {
	Contract     *SenderZEVM       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenderZEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenderZEVMCallerSession struct {
	Contract *SenderZEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SenderZEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenderZEVMTransactorSession struct {
	Contract     *SenderZEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SenderZEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenderZEVMRaw struct {
	Contract *SenderZEVM // Generic contract binding to access the raw methods on
}

// SenderZEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenderZEVMCallerRaw struct {
	Contract *SenderZEVMCaller // Generic read-only contract binding to access the raw methods on
}

// SenderZEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenderZEVMTransactorRaw struct {
	Contract *SenderZEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSenderZEVM creates a new instance of SenderZEVM, bound to a specific deployed contract.
func NewSenderZEVM(address common.Address, backend bind.ContractBackend) (*SenderZEVM, error) {
	contract, err := bindSenderZEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SenderZEVM{SenderZEVMCaller: SenderZEVMCaller{contract: contract}, SenderZEVMTransactor: SenderZEVMTransactor{contract: contract}, SenderZEVMFilterer: SenderZEVMFilterer{contract: contract}}, nil
}

// NewSenderZEVMCaller creates a new read-only instance of SenderZEVM, bound to a specific deployed contract.
func NewSenderZEVMCaller(address common.Address, caller bind.ContractCaller) (*SenderZEVMCaller, error) {
	contract, err := bindSenderZEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenderZEVMCaller{contract: contract}, nil
}

// NewSenderZEVMTransactor creates a new write-only instance of SenderZEVM, bound to a specific deployed contract.
func NewSenderZEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*SenderZEVMTransactor, error) {
	contract, err := bindSenderZEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenderZEVMTransactor{contract: contract}, nil
}

// NewSenderZEVMFilterer creates a new log filterer instance of SenderZEVM, bound to a specific deployed contract.
func NewSenderZEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*SenderZEVMFilterer, error) {
	contract, err := bindSenderZEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenderZEVMFilterer{contract: contract}, nil
}

// bindSenderZEVM binds a generic wrapper to an already deployed contract.
func bindSenderZEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SenderZEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenderZEVM *SenderZEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenderZEVM.Contract.SenderZEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenderZEVM *SenderZEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenderZEVM.Contract.SenderZEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenderZEVM *SenderZEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenderZEVM.Contract.SenderZEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenderZEVM *SenderZEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenderZEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenderZEVM *SenderZEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenderZEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenderZEVM *SenderZEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenderZEVM.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_SenderZEVM *SenderZEVMCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SenderZEVM.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_SenderZEVM *SenderZEVMSession) Gateway() (common.Address, error) {
	return _SenderZEVM.Contract.Gateway(&_SenderZEVM.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_SenderZEVM *SenderZEVMCallerSession) Gateway() (common.Address, error) {
	return _SenderZEVM.Contract.Gateway(&_SenderZEVM.CallOpts)
}

// CallReceiver is a paid mutator transaction binding the contract method 0x7a34d8bb.
//
// Solidity: function callReceiver(bytes receiver, address zrc20, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMTransactor) CallReceiver(opts *bind.TransactOpts, receiver []byte, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.contract.Transact(opts, "callReceiver", receiver, zrc20, str, num, flag)
}

// CallReceiver is a paid mutator transaction binding the contract method 0x7a34d8bb.
//
// Solidity: function callReceiver(bytes receiver, address zrc20, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMSession) CallReceiver(receiver []byte, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.Contract.CallReceiver(&_SenderZEVM.TransactOpts, receiver, zrc20, str, num, flag)
}

// CallReceiver is a paid mutator transaction binding the contract method 0x7a34d8bb.
//
// Solidity: function callReceiver(bytes receiver, address zrc20, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMTransactorSession) CallReceiver(receiver []byte, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.Contract.CallReceiver(&_SenderZEVM.TransactOpts, receiver, zrc20, str, num, flag)
}

// WithdrawAndCallReceiver is a paid mutator transaction binding the contract method 0x0abd8905.
//
// Solidity: function withdrawAndCallReceiver(bytes receiver, uint256 amount, address zrc20, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMTransactor) WithdrawAndCallReceiver(opts *bind.TransactOpts, receiver []byte, amount *big.Int, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.contract.Transact(opts, "withdrawAndCallReceiver", receiver, amount, zrc20, str, num, flag)
}

// WithdrawAndCallReceiver is a paid mutator transaction binding the contract method 0x0abd8905.
//
// Solidity: function withdrawAndCallReceiver(bytes receiver, uint256 amount, address zrc20, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMSession) WithdrawAndCallReceiver(receiver []byte, amount *big.Int, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.Contract.WithdrawAndCallReceiver(&_SenderZEVM.TransactOpts, receiver, amount, zrc20, str, num, flag)
}

// WithdrawAndCallReceiver is a paid mutator transaction binding the contract method 0x0abd8905.
//
// Solidity: function withdrawAndCallReceiver(bytes receiver, uint256 amount, address zrc20, string str, uint256 num, bool flag) returns()
func (_SenderZEVM *SenderZEVMTransactorSession) WithdrawAndCallReceiver(receiver []byte, amount *big.Int, zrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderZEVM.Contract.WithdrawAndCallReceiver(&_SenderZEVM.TransactOpts, receiver, amount, zrc20, str, num, flag)
}
