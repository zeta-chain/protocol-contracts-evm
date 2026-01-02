// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testuniversalcontract

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

// AbortContext is an auto generated low-level Go binding around an user-defined struct.
type AbortContext struct {
	Sender        []byte
	Asset         common.Address
	Amount        *big.Int
	Outgoing      bool
	ChainID       *big.Int
	RevertMessage []byte
}

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Sender    []byte
	SenderEVM common.Address
	ChainID   *big.Int
}

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// TestUniversalContractMetaData contains all meta data concerning the TestUniversalContract contract.
var TestUniversalContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayZEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onAbort\",\"inputs\":[{\"name\":\"abortContext\",\"type\":\"tuple\",\"internalType\":\"structAbortContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outgoing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"onRevert\",\"inputs\":[{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"registry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICoreRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ContextData\",\"inputs\":[{\"name\":\"origin\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"msgSender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextDataAbort\",\"inputs\":[{\"name\":\"abortContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAbortContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outgoing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContextDataRevert\",\"inputs\":[{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f80fd5b50737cce3eb018bf23e1fe2a32692f2c77592d1103946001600160a01b031663cc5ad8b66040518163ffffffff1660e01b81526004016020604051808303815f875af1158015610061573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100859190610096565b6001600160a01b03166080526100c3565b5f602082840312156100a6575f80fd5b81516001600160a01b03811681146100bc575f80fd5b9392505050565b60805161090e6100da5f395f607c015261090e5ff3fe608060405260043610610062575f3560e01c80637b1039991161003f5780637b10399914610105578063c9028a361461012c578063ebf9b2aa1461013f57005b8063116191b61461006b5780632d4cfb7e146100c75780635bcfd616146100e657005b3661006957005b005b348015610076575f80fd5b5061009e7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100d2575f80fd5b506100696100e13660046102af565b610152565b3480156100f1575f80fd5b50610069610100366004610370565b61018c565b348015610110575f80fd5b5061009e737cce3eb018bf23e1fe2a32692f2c77592d11039481565b61006961013a3660046103f4565b610207565b61006961014d36600461042b565b610236565b7f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db781604051610181919061053b565b60405180910390a150565b606081156101a3576101a08284018461063b565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e6101ce878061072b565b6101de60408a0160208b0161078c565b896040013533866040516101f7969594939291906107a5565b60405180910390a1505050505050565b7fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c481604051610181919061084f565b6060811561024d5761024a8284018461063b565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e610278858061072b565b610288604088016020890161078c565b876040013533866040516102a1969594939291906107a5565b60405180910390a150505050565b5f602082840312156102bf575f80fd5b813567ffffffffffffffff8111156102d5575f80fd5b820160c081850312156102e6575f80fd5b9392505050565b5f606082840312156102fd575f80fd5b50919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610326575f80fd5b919050565b5f8083601f84011261033b575f80fd5b50813567ffffffffffffffff811115610352575f80fd5b602083019150836020828501011115610369575f80fd5b9250929050565b5f805f805f60808688031215610384575f80fd5b853567ffffffffffffffff81111561039a575f80fd5b6103a6888289016102ed565b9550506103b560208701610303565b935060408601359250606086013567ffffffffffffffff8111156103d7575f80fd5b6103e38882890161032b565b969995985093965092949392505050565b5f60208284031215610404575f80fd5b813567ffffffffffffffff81111561041a575f80fd5b8201608081850312156102e6575f80fd5b5f805f6040848603121561043d575f80fd5b833567ffffffffffffffff811115610453575f80fd5b61045f868287016102ed565b935050602084013567ffffffffffffffff81111561047b575f80fd5b6104878682870161032b565b9497909650939450505050565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126104c7575f80fd5b830160208101925035905067ffffffffffffffff8111156104e6575f80fd5b803603821315610369575f80fd5b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b602081525f61054a8384610494565b60c0602085015261055f60e0850182846104f4565b91505073ffffffffffffffffffffffffffffffffffffffff61058360208601610303565b1660408401525f604085013590508060608501525060608401358015158082146105ab575f80fd5b80608086015250505f608085013590508060a0850152506105cf60a0850185610494565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c08601526106048382846104f4565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f6020828403121561064b575f80fd5b813567ffffffffffffffff811115610661575f80fd5b8201601f81018413610671575f80fd5b803567ffffffffffffffff81111561068b5761068b61060e565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156106f7576106f761060e565b60405281815282820160200186101561070e575f80fd5b816020840160208301375f91810160200191909152949350505050565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261075e575f80fd5b83018035915067ffffffffffffffff821115610778575f80fd5b602001915036819003821315610369575f80fd5b5f6020828403121561079c575f80fd5b6102e682610303565b60a081525f6107b860a08301888a6104f4565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff85166060840152828103608084015283518082528060208601602084015e5f6020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff61087183610303565b16602082015273ffffffffffffffffffffffffffffffffffffffff61089860208401610303565b1660408201525f8060408401359050806060840152506108bb6060840184610494565b6080808501526108cf60a0850182846104f4565b9594505050505056fea264697066735822122073b2b7273c8fa2cc0343725b06eef9af8f807b8eda5d1626b08ac364147ff7d464736f6c634300081a0033",
}

// TestUniversalContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TestUniversalContractMetaData.ABI instead.
var TestUniversalContractABI = TestUniversalContractMetaData.ABI

// TestUniversalContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestUniversalContractMetaData.Bin instead.
var TestUniversalContractBin = TestUniversalContractMetaData.Bin

// DeployTestUniversalContract deploys a new Ethereum contract, binding an instance of TestUniversalContract to it.
func DeployTestUniversalContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestUniversalContract, error) {
	parsed, err := TestUniversalContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestUniversalContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestUniversalContract{TestUniversalContractCaller: TestUniversalContractCaller{contract: contract}, TestUniversalContractTransactor: TestUniversalContractTransactor{contract: contract}, TestUniversalContractFilterer: TestUniversalContractFilterer{contract: contract}}, nil
}

// TestUniversalContract is an auto generated Go binding around an Ethereum contract.
type TestUniversalContract struct {
	TestUniversalContractCaller     // Read-only binding to the contract
	TestUniversalContractTransactor // Write-only binding to the contract
	TestUniversalContractFilterer   // Log filterer for contract events
}

// TestUniversalContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestUniversalContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUniversalContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestUniversalContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUniversalContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestUniversalContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestUniversalContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestUniversalContractSession struct {
	Contract     *TestUniversalContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestUniversalContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestUniversalContractCallerSession struct {
	Contract *TestUniversalContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TestUniversalContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestUniversalContractTransactorSession struct {
	Contract     *TestUniversalContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TestUniversalContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestUniversalContractRaw struct {
	Contract *TestUniversalContract // Generic contract binding to access the raw methods on
}

// TestUniversalContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestUniversalContractCallerRaw struct {
	Contract *TestUniversalContractCaller // Generic read-only contract binding to access the raw methods on
}

// TestUniversalContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestUniversalContractTransactorRaw struct {
	Contract *TestUniversalContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestUniversalContract creates a new instance of TestUniversalContract, bound to a specific deployed contract.
func NewTestUniversalContract(address common.Address, backend bind.ContractBackend) (*TestUniversalContract, error) {
	contract, err := bindTestUniversalContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestUniversalContract{TestUniversalContractCaller: TestUniversalContractCaller{contract: contract}, TestUniversalContractTransactor: TestUniversalContractTransactor{contract: contract}, TestUniversalContractFilterer: TestUniversalContractFilterer{contract: contract}}, nil
}

// NewTestUniversalContractCaller creates a new read-only instance of TestUniversalContract, bound to a specific deployed contract.
func NewTestUniversalContractCaller(address common.Address, caller bind.ContractCaller) (*TestUniversalContractCaller, error) {
	contract, err := bindTestUniversalContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestUniversalContractCaller{contract: contract}, nil
}

// NewTestUniversalContractTransactor creates a new write-only instance of TestUniversalContract, bound to a specific deployed contract.
func NewTestUniversalContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestUniversalContractTransactor, error) {
	contract, err := bindTestUniversalContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestUniversalContractTransactor{contract: contract}, nil
}

// NewTestUniversalContractFilterer creates a new log filterer instance of TestUniversalContract, bound to a specific deployed contract.
func NewTestUniversalContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TestUniversalContractFilterer, error) {
	contract, err := bindTestUniversalContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestUniversalContractFilterer{contract: contract}, nil
}

// bindTestUniversalContract binds a generic wrapper to an already deployed contract.
func bindTestUniversalContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestUniversalContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestUniversalContract *TestUniversalContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestUniversalContract.Contract.TestUniversalContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestUniversalContract *TestUniversalContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.TestUniversalContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestUniversalContract *TestUniversalContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.TestUniversalContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestUniversalContract *TestUniversalContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestUniversalContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestUniversalContract *TestUniversalContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestUniversalContract *TestUniversalContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_TestUniversalContract *TestUniversalContractCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestUniversalContract.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_TestUniversalContract *TestUniversalContractSession) Gateway() (common.Address, error) {
	return _TestUniversalContract.Contract.Gateway(&_TestUniversalContract.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_TestUniversalContract *TestUniversalContractCallerSession) Gateway() (common.Address, error) {
	return _TestUniversalContract.Contract.Gateway(&_TestUniversalContract.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_TestUniversalContract *TestUniversalContractCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestUniversalContract.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_TestUniversalContract *TestUniversalContractSession) Registry() (common.Address, error) {
	return _TestUniversalContract.Contract.Registry(&_TestUniversalContract.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_TestUniversalContract *TestUniversalContractCallerSession) Registry() (common.Address, error) {
	return _TestUniversalContract.Contract.Registry(&_TestUniversalContract.CallOpts)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_TestUniversalContract *TestUniversalContractTransactor) OnAbort(opts *bind.TransactOpts, abortContext AbortContext) (*types.Transaction, error) {
	return _TestUniversalContract.contract.Transact(opts, "onAbort", abortContext)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_TestUniversalContract *TestUniversalContractSession) OnAbort(abortContext AbortContext) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.OnAbort(&_TestUniversalContract.TransactOpts, abortContext)
}

// OnAbort is a paid mutator transaction binding the contract method 0x2d4cfb7e.
//
// Solidity: function onAbort((bytes,address,uint256,bool,uint256,bytes) abortContext) returns()
func (_TestUniversalContract *TestUniversalContractTransactorSession) OnAbort(abortContext AbortContext) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.OnAbort(&_TestUniversalContract.TransactOpts, abortContext)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) context, address , uint256 , bytes message) returns()
func (_TestUniversalContract *TestUniversalContractTransactor) OnCall(opts *bind.TransactOpts, context MessageContext, arg1 common.Address, arg2 *big.Int, message []byte) (*types.Transaction, error) {
	return _TestUniversalContract.contract.Transact(opts, "onCall", context, arg1, arg2, message)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) context, address , uint256 , bytes message) returns()
func (_TestUniversalContract *TestUniversalContractSession) OnCall(context MessageContext, arg1 common.Address, arg2 *big.Int, message []byte) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.OnCall(&_TestUniversalContract.TransactOpts, context, arg1, arg2, message)
}

// OnCall is a paid mutator transaction binding the contract method 0x5bcfd616.
//
// Solidity: function onCall((bytes,address,uint256) context, address , uint256 , bytes message) returns()
func (_TestUniversalContract *TestUniversalContractTransactorSession) OnCall(context MessageContext, arg1 common.Address, arg2 *big.Int, message []byte) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.OnCall(&_TestUniversalContract.TransactOpts, context, arg1, arg2, message)
}

// OnCall0 is a paid mutator transaction binding the contract method 0xebf9b2aa.
//
// Solidity: function onCall((bytes,address,uint256) context, bytes message) payable returns()
func (_TestUniversalContract *TestUniversalContractTransactor) OnCall0(opts *bind.TransactOpts, context MessageContext, message []byte) (*types.Transaction, error) {
	return _TestUniversalContract.contract.Transact(opts, "onCall0", context, message)
}

// OnCall0 is a paid mutator transaction binding the contract method 0xebf9b2aa.
//
// Solidity: function onCall((bytes,address,uint256) context, bytes message) payable returns()
func (_TestUniversalContract *TestUniversalContractSession) OnCall0(context MessageContext, message []byte) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.OnCall0(&_TestUniversalContract.TransactOpts, context, message)
}

// OnCall0 is a paid mutator transaction binding the contract method 0xebf9b2aa.
//
// Solidity: function onCall((bytes,address,uint256) context, bytes message) payable returns()
func (_TestUniversalContract *TestUniversalContractTransactorSession) OnCall0(context MessageContext, message []byte) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.OnCall0(&_TestUniversalContract.TransactOpts, context, message)
}

// OnRevert is a paid mutator transaction binding the contract method 0xc9028a36.
//
// Solidity: function onRevert((address,address,uint256,bytes) revertContext) payable returns()
func (_TestUniversalContract *TestUniversalContractTransactor) OnRevert(opts *bind.TransactOpts, revertContext RevertContext) (*types.Transaction, error) {
	return _TestUniversalContract.contract.Transact(opts, "onRevert", revertContext)
}

// OnRevert is a paid mutator transaction binding the contract method 0xc9028a36.
//
// Solidity: function onRevert((address,address,uint256,bytes) revertContext) payable returns()
func (_TestUniversalContract *TestUniversalContractSession) OnRevert(revertContext RevertContext) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.OnRevert(&_TestUniversalContract.TransactOpts, revertContext)
}

// OnRevert is a paid mutator transaction binding the contract method 0xc9028a36.
//
// Solidity: function onRevert((address,address,uint256,bytes) revertContext) payable returns()
func (_TestUniversalContract *TestUniversalContractTransactorSession) OnRevert(revertContext RevertContext) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.OnRevert(&_TestUniversalContract.TransactOpts, revertContext)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestUniversalContract *TestUniversalContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TestUniversalContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestUniversalContract *TestUniversalContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.Fallback(&_TestUniversalContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TestUniversalContract *TestUniversalContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TestUniversalContract.Contract.Fallback(&_TestUniversalContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestUniversalContract *TestUniversalContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestUniversalContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestUniversalContract *TestUniversalContractSession) Receive() (*types.Transaction, error) {
	return _TestUniversalContract.Contract.Receive(&_TestUniversalContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestUniversalContract *TestUniversalContractTransactorSession) Receive() (*types.Transaction, error) {
	return _TestUniversalContract.Contract.Receive(&_TestUniversalContract.TransactOpts)
}

// TestUniversalContractContextDataIterator is returned from FilterContextData and is used to iterate over the raw logs and unpacked data for ContextData events raised by the TestUniversalContract contract.
type TestUniversalContractContextDataIterator struct {
	Event *TestUniversalContractContextData // Event containing the contract specifics and raw log

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
func (it *TestUniversalContractContextDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestUniversalContractContextData)
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
		it.Event = new(TestUniversalContractContextData)
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
func (it *TestUniversalContractContextDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestUniversalContractContextDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestUniversalContractContextData represents a ContextData event raised by the TestUniversalContract contract.
type TestUniversalContractContextData struct {
	Origin    []byte
	Sender    common.Address
	ChainID   *big.Int
	MsgSender common.Address
	Message   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContextData is a free log retrieval operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestUniversalContract *TestUniversalContractFilterer) FilterContextData(opts *bind.FilterOpts) (*TestUniversalContractContextDataIterator, error) {

	logs, sub, err := _TestUniversalContract.contract.FilterLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return &TestUniversalContractContextDataIterator{contract: _TestUniversalContract.contract, event: "ContextData", logs: logs, sub: sub}, nil
}

// WatchContextData is a free log subscription operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestUniversalContract *TestUniversalContractFilterer) WatchContextData(opts *bind.WatchOpts, sink chan<- *TestUniversalContractContextData) (event.Subscription, error) {

	logs, sub, err := _TestUniversalContract.contract.WatchLogs(opts, "ContextData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestUniversalContractContextData)
				if err := _TestUniversalContract.contract.UnpackLog(event, "ContextData", log); err != nil {
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

// ParseContextData is a log parse operation binding the contract event 0xcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e.
//
// Solidity: event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message)
func (_TestUniversalContract *TestUniversalContractFilterer) ParseContextData(log types.Log) (*TestUniversalContractContextData, error) {
	event := new(TestUniversalContractContextData)
	if err := _TestUniversalContract.contract.UnpackLog(event, "ContextData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestUniversalContractContextDataAbortIterator is returned from FilterContextDataAbort and is used to iterate over the raw logs and unpacked data for ContextDataAbort events raised by the TestUniversalContract contract.
type TestUniversalContractContextDataAbortIterator struct {
	Event *TestUniversalContractContextDataAbort // Event containing the contract specifics and raw log

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
func (it *TestUniversalContractContextDataAbortIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestUniversalContractContextDataAbort)
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
		it.Event = new(TestUniversalContractContextDataAbort)
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
func (it *TestUniversalContractContextDataAbortIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestUniversalContractContextDataAbortIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestUniversalContractContextDataAbort represents a ContextDataAbort event raised by the TestUniversalContract contract.
type TestUniversalContractContextDataAbort struct {
	AbortContext AbortContext
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContextDataAbort is a free log retrieval operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_TestUniversalContract *TestUniversalContractFilterer) FilterContextDataAbort(opts *bind.FilterOpts) (*TestUniversalContractContextDataAbortIterator, error) {

	logs, sub, err := _TestUniversalContract.contract.FilterLogs(opts, "ContextDataAbort")
	if err != nil {
		return nil, err
	}
	return &TestUniversalContractContextDataAbortIterator{contract: _TestUniversalContract.contract, event: "ContextDataAbort", logs: logs, sub: sub}, nil
}

// WatchContextDataAbort is a free log subscription operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_TestUniversalContract *TestUniversalContractFilterer) WatchContextDataAbort(opts *bind.WatchOpts, sink chan<- *TestUniversalContractContextDataAbort) (event.Subscription, error) {

	logs, sub, err := _TestUniversalContract.contract.WatchLogs(opts, "ContextDataAbort")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestUniversalContractContextDataAbort)
				if err := _TestUniversalContract.contract.UnpackLog(event, "ContextDataAbort", log); err != nil {
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

// ParseContextDataAbort is a log parse operation binding the contract event 0x7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7.
//
// Solidity: event ContextDataAbort((bytes,address,uint256,bool,uint256,bytes) abortContext)
func (_TestUniversalContract *TestUniversalContractFilterer) ParseContextDataAbort(log types.Log) (*TestUniversalContractContextDataAbort, error) {
	event := new(TestUniversalContractContextDataAbort)
	if err := _TestUniversalContract.contract.UnpackLog(event, "ContextDataAbort", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestUniversalContractContextDataRevertIterator is returned from FilterContextDataRevert and is used to iterate over the raw logs and unpacked data for ContextDataRevert events raised by the TestUniversalContract contract.
type TestUniversalContractContextDataRevertIterator struct {
	Event *TestUniversalContractContextDataRevert // Event containing the contract specifics and raw log

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
func (it *TestUniversalContractContextDataRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestUniversalContractContextDataRevert)
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
		it.Event = new(TestUniversalContractContextDataRevert)
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
func (it *TestUniversalContractContextDataRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestUniversalContractContextDataRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestUniversalContractContextDataRevert represents a ContextDataRevert event raised by the TestUniversalContract contract.
type TestUniversalContractContextDataRevert struct {
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContextDataRevert is a free log retrieval operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_TestUniversalContract *TestUniversalContractFilterer) FilterContextDataRevert(opts *bind.FilterOpts) (*TestUniversalContractContextDataRevertIterator, error) {

	logs, sub, err := _TestUniversalContract.contract.FilterLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return &TestUniversalContractContextDataRevertIterator{contract: _TestUniversalContract.contract, event: "ContextDataRevert", logs: logs, sub: sub}, nil
}

// WatchContextDataRevert is a free log subscription operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_TestUniversalContract *TestUniversalContractFilterer) WatchContextDataRevert(opts *bind.WatchOpts, sink chan<- *TestUniversalContractContextDataRevert) (event.Subscription, error) {

	logs, sub, err := _TestUniversalContract.contract.WatchLogs(opts, "ContextDataRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestUniversalContractContextDataRevert)
				if err := _TestUniversalContract.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
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

// ParseContextDataRevert is a log parse operation binding the contract event 0xd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4.
//
// Solidity: event ContextDataRevert((address,address,uint256,bytes) revertContext)
func (_TestUniversalContract *TestUniversalContractFilterer) ParseContextDataRevert(log types.Log) (*TestUniversalContractContextDataRevert, error) {
	event := new(TestUniversalContractContextDataRevert)
	if err := _TestUniversalContract.contract.UnpackLog(event, "ContextDataRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
