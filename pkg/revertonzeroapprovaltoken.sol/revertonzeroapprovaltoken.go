// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package revertonzeroapprovaltoken

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

// RevertOnZeroApprovalTokenMetaData contains all meta data concerning the RevertOnZeroApprovalToken contract.
var RevertOnZeroApprovalTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610cb7380380610cb783398101604081905261002f91610111565b81818181600361003f8382610203565b50600461004c8282610203565b505050505050506102c1565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007f57600080fd5b81516001600160401b0381111561009857610098610058565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c6576100c6610058565b6040528181528382016020018510156100de57600080fd5b60005b828110156100fd576020818601810151838301820152016100e1565b506000918101602001919091529392505050565b6000806040838503121561012457600080fd5b82516001600160401b0381111561013a57600080fd5b6101468582860161006e565b602085015190935090506001600160401b0381111561016457600080fd5b6101708582860161006e565b9150509250929050565b600181811c9082168061018e57607f821691505b6020821081036101ae57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fe57806000526020600020601f840160051c810160208510156101db5750805b601f840160051c820191505b818110156101fb57600081556001016101e7565b50505b505050565b81516001600160401b0381111561021c5761021c610058565b6102308161022a845461017a565b846101b4565b6020601f821160018114610264576000831561024c5750848201515b600019600385901b1c1916600184901b1784556101fb565b600084815260208120601f198516915b828110156102945787850151825560209485019460019092019101610274565b50848210156102b25786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109e7806102d06000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107da565b60405180910390f35b6100f46100ef36600461086f565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f4610124366004610899565b610297565b604051601281526020016100d8565b61014b61014636600461086f565b6102bb565b005b61010861015b3660046108d6565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c9565b6100f461019936600461086f565b6102d8565b6101086101ac3660046108f1565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610924565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610924565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b600080821161028457600080fd5b61028e83836102f0565b90505b92915050565b6000336102a58582856102fe565b6102b08585856103d2565b506001949350505050565b6102c58282610482565b5050565b6060600480546101f390610924565b6000336102e68185856103d2565b5060019392505050565b6000336102e68185856104de565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103cc57818110156103bd576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103cc848484840360006104e7565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610422576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103b4565b73ffffffffffffffffffffffffffffffffffffffff8216610472576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103b4565b61047d83838361062f565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166104d2576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103b4565b6102c56000838361062f565b61047d83838360015b73ffffffffffffffffffffffffffffffffffffffff8416610537576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103b4565b73ffffffffffffffffffffffffffffffffffffffff8316610587576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103b4565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103cc578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161062191815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661066757806002600082825461065c9190610977565b909155506107199050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106ed576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103b4565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff82166107425760028054829003905561076e565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107cd91815260200190565b60405180910390a3505050565b602081526000825180602084015260005b8181101561080857602081860181015160408684010152016107eb565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461086a57600080fd5b919050565b6000806040838503121561088257600080fd5b61088b83610846565b946020939093013593505050565b6000806000606084860312156108ae57600080fd5b6108b784610846565b92506108c560208501610846565b929592945050506040919091013590565b6000602082840312156108e857600080fd5b61028e82610846565b6000806040838503121561090457600080fd5b61090d83610846565b915061091b60208401610846565b90509250929050565b600181811c9082168061093857607f821691505b602082108103610971577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b80820180821115610291577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea264697066735822122073db0e30c05200fc444905446d879bf6bfcd0ec4eb5c23d25c82318bbdc3416b64736f6c634300081a0033",
}

// RevertOnZeroApprovalTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use RevertOnZeroApprovalTokenMetaData.ABI instead.
var RevertOnZeroApprovalTokenABI = RevertOnZeroApprovalTokenMetaData.ABI

// RevertOnZeroApprovalTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RevertOnZeroApprovalTokenMetaData.Bin instead.
var RevertOnZeroApprovalTokenBin = RevertOnZeroApprovalTokenMetaData.Bin

// DeployRevertOnZeroApprovalToken deploys a new Ethereum contract, binding an instance of RevertOnZeroApprovalToken to it.
func DeployRevertOnZeroApprovalToken(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *RevertOnZeroApprovalToken, error) {
	parsed, err := RevertOnZeroApprovalTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RevertOnZeroApprovalTokenBin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RevertOnZeroApprovalToken{RevertOnZeroApprovalTokenCaller: RevertOnZeroApprovalTokenCaller{contract: contract}, RevertOnZeroApprovalTokenTransactor: RevertOnZeroApprovalTokenTransactor{contract: contract}, RevertOnZeroApprovalTokenFilterer: RevertOnZeroApprovalTokenFilterer{contract: contract}}, nil
}

// RevertOnZeroApprovalToken is an auto generated Go binding around an Ethereum contract.
type RevertOnZeroApprovalToken struct {
	RevertOnZeroApprovalTokenCaller     // Read-only binding to the contract
	RevertOnZeroApprovalTokenTransactor // Write-only binding to the contract
	RevertOnZeroApprovalTokenFilterer   // Log filterer for contract events
}

// RevertOnZeroApprovalTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RevertOnZeroApprovalTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertOnZeroApprovalTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RevertOnZeroApprovalTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertOnZeroApprovalTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RevertOnZeroApprovalTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertOnZeroApprovalTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RevertOnZeroApprovalTokenSession struct {
	Contract     *RevertOnZeroApprovalToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RevertOnZeroApprovalTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RevertOnZeroApprovalTokenCallerSession struct {
	Contract *RevertOnZeroApprovalTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// RevertOnZeroApprovalTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RevertOnZeroApprovalTokenTransactorSession struct {
	Contract     *RevertOnZeroApprovalTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// RevertOnZeroApprovalTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type RevertOnZeroApprovalTokenRaw struct {
	Contract *RevertOnZeroApprovalToken // Generic contract binding to access the raw methods on
}

// RevertOnZeroApprovalTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RevertOnZeroApprovalTokenCallerRaw struct {
	Contract *RevertOnZeroApprovalTokenCaller // Generic read-only contract binding to access the raw methods on
}

// RevertOnZeroApprovalTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RevertOnZeroApprovalTokenTransactorRaw struct {
	Contract *RevertOnZeroApprovalTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRevertOnZeroApprovalToken creates a new instance of RevertOnZeroApprovalToken, bound to a specific deployed contract.
func NewRevertOnZeroApprovalToken(address common.Address, backend bind.ContractBackend) (*RevertOnZeroApprovalToken, error) {
	contract, err := bindRevertOnZeroApprovalToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RevertOnZeroApprovalToken{RevertOnZeroApprovalTokenCaller: RevertOnZeroApprovalTokenCaller{contract: contract}, RevertOnZeroApprovalTokenTransactor: RevertOnZeroApprovalTokenTransactor{contract: contract}, RevertOnZeroApprovalTokenFilterer: RevertOnZeroApprovalTokenFilterer{contract: contract}}, nil
}

// NewRevertOnZeroApprovalTokenCaller creates a new read-only instance of RevertOnZeroApprovalToken, bound to a specific deployed contract.
func NewRevertOnZeroApprovalTokenCaller(address common.Address, caller bind.ContractCaller) (*RevertOnZeroApprovalTokenCaller, error) {
	contract, err := bindRevertOnZeroApprovalToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RevertOnZeroApprovalTokenCaller{contract: contract}, nil
}

// NewRevertOnZeroApprovalTokenTransactor creates a new write-only instance of RevertOnZeroApprovalToken, bound to a specific deployed contract.
func NewRevertOnZeroApprovalTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*RevertOnZeroApprovalTokenTransactor, error) {
	contract, err := bindRevertOnZeroApprovalToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RevertOnZeroApprovalTokenTransactor{contract: contract}, nil
}

// NewRevertOnZeroApprovalTokenFilterer creates a new log filterer instance of RevertOnZeroApprovalToken, bound to a specific deployed contract.
func NewRevertOnZeroApprovalTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*RevertOnZeroApprovalTokenFilterer, error) {
	contract, err := bindRevertOnZeroApprovalToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RevertOnZeroApprovalTokenFilterer{contract: contract}, nil
}

// bindRevertOnZeroApprovalToken binds a generic wrapper to an already deployed contract.
func bindRevertOnZeroApprovalToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RevertOnZeroApprovalTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevertOnZeroApprovalToken.Contract.RevertOnZeroApprovalTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.Contract.RevertOnZeroApprovalTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.Contract.RevertOnZeroApprovalTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevertOnZeroApprovalToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RevertOnZeroApprovalToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _RevertOnZeroApprovalToken.Contract.Allowance(&_RevertOnZeroApprovalToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _RevertOnZeroApprovalToken.Contract.Allowance(&_RevertOnZeroApprovalToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RevertOnZeroApprovalToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RevertOnZeroApprovalToken.Contract.BalanceOf(&_RevertOnZeroApprovalToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _RevertOnZeroApprovalToken.Contract.BalanceOf(&_RevertOnZeroApprovalToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RevertOnZeroApprovalToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenSession) Decimals() (uint8, error) {
	return _RevertOnZeroApprovalToken.Contract.Decimals(&_RevertOnZeroApprovalToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenCallerSession) Decimals() (uint8, error) {
	return _RevertOnZeroApprovalToken.Contract.Decimals(&_RevertOnZeroApprovalToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RevertOnZeroApprovalToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenSession) Name() (string, error) {
	return _RevertOnZeroApprovalToken.Contract.Name(&_RevertOnZeroApprovalToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenCallerSession) Name() (string, error) {
	return _RevertOnZeroApprovalToken.Contract.Name(&_RevertOnZeroApprovalToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RevertOnZeroApprovalToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenSession) Symbol() (string, error) {
	return _RevertOnZeroApprovalToken.Contract.Symbol(&_RevertOnZeroApprovalToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenCallerSession) Symbol() (string, error) {
	return _RevertOnZeroApprovalToken.Contract.Symbol(&_RevertOnZeroApprovalToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RevertOnZeroApprovalToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenSession) TotalSupply() (*big.Int, error) {
	return _RevertOnZeroApprovalToken.Contract.TotalSupply(&_RevertOnZeroApprovalToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _RevertOnZeroApprovalToken.Contract.TotalSupply(&_RevertOnZeroApprovalToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.Contract.Approve(&_RevertOnZeroApprovalToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.Contract.Approve(&_RevertOnZeroApprovalToken.TransactOpts, spender, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.Contract.Mint(&_RevertOnZeroApprovalToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.Contract.Mint(&_RevertOnZeroApprovalToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.Contract.Transfer(&_RevertOnZeroApprovalToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.Contract.Transfer(&_RevertOnZeroApprovalToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.Contract.TransferFrom(&_RevertOnZeroApprovalToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RevertOnZeroApprovalToken.Contract.TransferFrom(&_RevertOnZeroApprovalToken.TransactOpts, from, to, value)
}

// RevertOnZeroApprovalTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RevertOnZeroApprovalToken contract.
type RevertOnZeroApprovalTokenApprovalIterator struct {
	Event *RevertOnZeroApprovalTokenApproval // Event containing the contract specifics and raw log

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
func (it *RevertOnZeroApprovalTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RevertOnZeroApprovalTokenApproval)
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
		it.Event = new(RevertOnZeroApprovalTokenApproval)
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
func (it *RevertOnZeroApprovalTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RevertOnZeroApprovalTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RevertOnZeroApprovalTokenApproval represents a Approval event raised by the RevertOnZeroApprovalToken contract.
type RevertOnZeroApprovalTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RevertOnZeroApprovalTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RevertOnZeroApprovalToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RevertOnZeroApprovalTokenApprovalIterator{contract: _RevertOnZeroApprovalToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RevertOnZeroApprovalTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RevertOnZeroApprovalToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RevertOnZeroApprovalTokenApproval)
				if err := _RevertOnZeroApprovalToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenFilterer) ParseApproval(log types.Log) (*RevertOnZeroApprovalTokenApproval, error) {
	event := new(RevertOnZeroApprovalTokenApproval)
	if err := _RevertOnZeroApprovalToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RevertOnZeroApprovalTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RevertOnZeroApprovalToken contract.
type RevertOnZeroApprovalTokenTransferIterator struct {
	Event *RevertOnZeroApprovalTokenTransfer // Event containing the contract specifics and raw log

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
func (it *RevertOnZeroApprovalTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RevertOnZeroApprovalTokenTransfer)
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
		it.Event = new(RevertOnZeroApprovalTokenTransfer)
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
func (it *RevertOnZeroApprovalTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RevertOnZeroApprovalTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RevertOnZeroApprovalTokenTransfer represents a Transfer event raised by the RevertOnZeroApprovalToken contract.
type RevertOnZeroApprovalTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RevertOnZeroApprovalTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RevertOnZeroApprovalToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RevertOnZeroApprovalTokenTransferIterator{contract: _RevertOnZeroApprovalToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RevertOnZeroApprovalTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RevertOnZeroApprovalToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RevertOnZeroApprovalTokenTransfer)
				if err := _RevertOnZeroApprovalToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RevertOnZeroApprovalToken *RevertOnZeroApprovalTokenFilterer) ParseTransfer(log types.Log) (*RevertOnZeroApprovalTokenTransfer, error) {
	event := new(RevertOnZeroApprovalTokenTransfer)
	if err := _RevertOnZeroApprovalToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
