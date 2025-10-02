// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nonreturnapprovaltoken

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

// NonReturnApprovalTokenMetaData contains all meta data concerning the NonReturnApprovalToken contract.
var NonReturnApprovalTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60806040526002805460ff1916601217905534801561001d57600080fd5b50604051610bf8380380610bf883398101604081905261003c91610116565b60006100488382610208565b5060016100558282610208565b5050506102c6565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261008457600080fd5b81516001600160401b0381111561009d5761009d61005d565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100cb576100cb61005d565b6040528181528382016020018510156100e357600080fd5b60005b82811015610102576020818601810151838301820152016100e6565b506000918101602001919091529392505050565b6000806040838503121561012957600080fd5b82516001600160401b0381111561013f57600080fd5b61014b85828601610073565b602085015190935090506001600160401b0381111561016957600080fd5b61017585828601610073565b9150509250929050565b600181811c9082168061019357607f821691505b6020821081036101b357634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561020357806000526020600020601f840160051c810160208510156101e05750805b601f840160051c820191505b8181101561020057600081556001016101ec565b50505b505050565b81516001600160401b038111156102215761022161005d565b6102358161022f845461017f565b846101b9565b6020601f82116001811461026957600083156102515750848201515b600019600385901b1c1916600184901b178455610200565b600084815260208120601f198516915b828110156102995787850151825560209485019460019092019101610279565b50848210156102b75786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610923806102d56000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610182578063a9059cbb1461018a578063dd62ed3e1461019d57600080fd5b806340c10f191461014f57806370a082311461016257600080fd5b806318160ddd116100a757806318160ddd146100f657806323b872dd1461010d578063313ce5671461013057600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101c8565b6040516100d891906106f4565b60405180910390f35b6100f46100ef366004610789565b610256565b005b6100ff60035481565b6040519081526020016100d8565b61012061011b3660046107b3565b6102c3565b60405190151581526020016100d8565b60025461013d9060ff1681565b60405160ff90911681526020016100d8565b6100f461015d366004610789565b61051d565b6100ff6101703660046107f0565b60046020526000908152604090205481565b6100cb6105b9565b610120610198366004610789565b6105c6565b6100ff6101ab366004610812565b600560209081526000928352604080842090915290825290205481565b600080546101d590610845565b80601f016020809104026020016040519081016040528092919081815260200182805461020190610845565b801561024e5780601f106102235761010080835404028352916020019161024e565b820191906000526020600020905b81548152906001019060200180831161023157829003601f168201915b505050505081565b33600081815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff871680855290835292819020859055518481529192917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260046020526040812054821115610357576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e636500000000000000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff841660009081526005602090815260408083203384529091529020548211156103f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e73756666696369656e7420616c6c6f77616e636500000000000000000000604482015260640161034e565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260046020526040812080548492906104269084906108c7565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600090815260046020526040812080548492906104609084906108da565b909155505073ffffffffffffffffffffffffffffffffffffffff84166000908152600560209081526040808320338452909152812080548492906104a59084906108c7565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161050b91815260200190565b60405180910390a35060019392505050565b806003600082825461052f91906108da565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260046020526040812080548392906105699084906108da565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016102b7565b600180546101d590610845565b3360009081526004602052604081205482111561063f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e6365000000000000000000000000604482015260640161034e565b336000908152600460205260408120805484929061065e9084906108c7565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600090815260046020526040812080548492906106989084906108da565b909155505060405182815273ffffffffffffffffffffffffffffffffffffffff84169033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35060015b92915050565b602081526000825180602084015260005b818110156107225760208186018101516040868401015201610705565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461078457600080fd5b919050565b6000806040838503121561079c57600080fd5b6107a583610760565b946020939093013593505050565b6000806000606084860312156107c857600080fd5b6107d184610760565b92506107df60208501610760565b929592945050506040919091013590565b60006020828403121561080257600080fd5b61080b82610760565b9392505050565b6000806040838503121561082557600080fd5b61082e83610760565b915061083c60208401610760565b90509250929050565b600181811c9082168061085957607f821691505b602082108103610892577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156106ee576106ee610898565b808201808211156106ee576106ee61089856fea264697066735822122041013599089452b2c9cdcab1a4eb2194e2bcab74511cb3cee684a048cea2529f64736f6c634300081a0033",
}

// NonReturnApprovalTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use NonReturnApprovalTokenMetaData.ABI instead.
var NonReturnApprovalTokenABI = NonReturnApprovalTokenMetaData.ABI

// NonReturnApprovalTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NonReturnApprovalTokenMetaData.Bin instead.
var NonReturnApprovalTokenBin = NonReturnApprovalTokenMetaData.Bin

// DeployNonReturnApprovalToken deploys a new Ethereum contract, binding an instance of NonReturnApprovalToken to it.
func DeployNonReturnApprovalToken(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *NonReturnApprovalToken, error) {
	parsed, err := NonReturnApprovalTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NonReturnApprovalTokenBin), backend, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NonReturnApprovalToken{NonReturnApprovalTokenCaller: NonReturnApprovalTokenCaller{contract: contract}, NonReturnApprovalTokenTransactor: NonReturnApprovalTokenTransactor{contract: contract}, NonReturnApprovalTokenFilterer: NonReturnApprovalTokenFilterer{contract: contract}}, nil
}

// NonReturnApprovalToken is an auto generated Go binding around an Ethereum contract.
type NonReturnApprovalToken struct {
	NonReturnApprovalTokenCaller     // Read-only binding to the contract
	NonReturnApprovalTokenTransactor // Write-only binding to the contract
	NonReturnApprovalTokenFilterer   // Log filterer for contract events
}

// NonReturnApprovalTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type NonReturnApprovalTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonReturnApprovalTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NonReturnApprovalTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonReturnApprovalTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NonReturnApprovalTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonReturnApprovalTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NonReturnApprovalTokenSession struct {
	Contract     *NonReturnApprovalToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NonReturnApprovalTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NonReturnApprovalTokenCallerSession struct {
	Contract *NonReturnApprovalTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// NonReturnApprovalTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NonReturnApprovalTokenTransactorSession struct {
	Contract     *NonReturnApprovalTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// NonReturnApprovalTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type NonReturnApprovalTokenRaw struct {
	Contract *NonReturnApprovalToken // Generic contract binding to access the raw methods on
}

// NonReturnApprovalTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NonReturnApprovalTokenCallerRaw struct {
	Contract *NonReturnApprovalTokenCaller // Generic read-only contract binding to access the raw methods on
}

// NonReturnApprovalTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NonReturnApprovalTokenTransactorRaw struct {
	Contract *NonReturnApprovalTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNonReturnApprovalToken creates a new instance of NonReturnApprovalToken, bound to a specific deployed contract.
func NewNonReturnApprovalToken(address common.Address, backend bind.ContractBackend) (*NonReturnApprovalToken, error) {
	contract, err := bindNonReturnApprovalToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NonReturnApprovalToken{NonReturnApprovalTokenCaller: NonReturnApprovalTokenCaller{contract: contract}, NonReturnApprovalTokenTransactor: NonReturnApprovalTokenTransactor{contract: contract}, NonReturnApprovalTokenFilterer: NonReturnApprovalTokenFilterer{contract: contract}}, nil
}

// NewNonReturnApprovalTokenCaller creates a new read-only instance of NonReturnApprovalToken, bound to a specific deployed contract.
func NewNonReturnApprovalTokenCaller(address common.Address, caller bind.ContractCaller) (*NonReturnApprovalTokenCaller, error) {
	contract, err := bindNonReturnApprovalToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NonReturnApprovalTokenCaller{contract: contract}, nil
}

// NewNonReturnApprovalTokenTransactor creates a new write-only instance of NonReturnApprovalToken, bound to a specific deployed contract.
func NewNonReturnApprovalTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*NonReturnApprovalTokenTransactor, error) {
	contract, err := bindNonReturnApprovalToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NonReturnApprovalTokenTransactor{contract: contract}, nil
}

// NewNonReturnApprovalTokenFilterer creates a new log filterer instance of NonReturnApprovalToken, bound to a specific deployed contract.
func NewNonReturnApprovalTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*NonReturnApprovalTokenFilterer, error) {
	contract, err := bindNonReturnApprovalToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NonReturnApprovalTokenFilterer{contract: contract}, nil
}

// bindNonReturnApprovalToken binds a generic wrapper to an already deployed contract.
func bindNonReturnApprovalToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NonReturnApprovalTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonReturnApprovalToken *NonReturnApprovalTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonReturnApprovalToken.Contract.NonReturnApprovalTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonReturnApprovalToken *NonReturnApprovalTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonReturnApprovalToken.Contract.NonReturnApprovalTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonReturnApprovalToken *NonReturnApprovalTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonReturnApprovalToken.Contract.NonReturnApprovalTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonReturnApprovalToken *NonReturnApprovalTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonReturnApprovalToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonReturnApprovalToken *NonReturnApprovalTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonReturnApprovalToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonReturnApprovalToken *NonReturnApprovalTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonReturnApprovalToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_NonReturnApprovalToken *NonReturnApprovalTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NonReturnApprovalToken.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_NonReturnApprovalToken *NonReturnApprovalTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _NonReturnApprovalToken.Contract.Allowance(&_NonReturnApprovalToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_NonReturnApprovalToken *NonReturnApprovalTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _NonReturnApprovalToken.Contract.Allowance(&_NonReturnApprovalToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_NonReturnApprovalToken *NonReturnApprovalTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NonReturnApprovalToken.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_NonReturnApprovalToken *NonReturnApprovalTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _NonReturnApprovalToken.Contract.BalanceOf(&_NonReturnApprovalToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_NonReturnApprovalToken *NonReturnApprovalTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _NonReturnApprovalToken.Contract.BalanceOf(&_NonReturnApprovalToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NonReturnApprovalToken *NonReturnApprovalTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NonReturnApprovalToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NonReturnApprovalToken *NonReturnApprovalTokenSession) Decimals() (uint8, error) {
	return _NonReturnApprovalToken.Contract.Decimals(&_NonReturnApprovalToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NonReturnApprovalToken *NonReturnApprovalTokenCallerSession) Decimals() (uint8, error) {
	return _NonReturnApprovalToken.Contract.Decimals(&_NonReturnApprovalToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NonReturnApprovalToken *NonReturnApprovalTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NonReturnApprovalToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NonReturnApprovalToken *NonReturnApprovalTokenSession) Name() (string, error) {
	return _NonReturnApprovalToken.Contract.Name(&_NonReturnApprovalToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NonReturnApprovalToken *NonReturnApprovalTokenCallerSession) Name() (string, error) {
	return _NonReturnApprovalToken.Contract.Name(&_NonReturnApprovalToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NonReturnApprovalToken *NonReturnApprovalTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NonReturnApprovalToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NonReturnApprovalToken *NonReturnApprovalTokenSession) Symbol() (string, error) {
	return _NonReturnApprovalToken.Contract.Symbol(&_NonReturnApprovalToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NonReturnApprovalToken *NonReturnApprovalTokenCallerSession) Symbol() (string, error) {
	return _NonReturnApprovalToken.Contract.Symbol(&_NonReturnApprovalToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NonReturnApprovalToken *NonReturnApprovalTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NonReturnApprovalToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NonReturnApprovalToken *NonReturnApprovalTokenSession) TotalSupply() (*big.Int, error) {
	return _NonReturnApprovalToken.Contract.TotalSupply(&_NonReturnApprovalToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NonReturnApprovalToken *NonReturnApprovalTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _NonReturnApprovalToken.Contract.TotalSupply(&_NonReturnApprovalToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns()
func (_NonReturnApprovalToken *NonReturnApprovalTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NonReturnApprovalToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns()
func (_NonReturnApprovalToken *NonReturnApprovalTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NonReturnApprovalToken.Contract.Approve(&_NonReturnApprovalToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns()
func (_NonReturnApprovalToken *NonReturnApprovalTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NonReturnApprovalToken.Contract.Approve(&_NonReturnApprovalToken.TransactOpts, spender, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_NonReturnApprovalToken *NonReturnApprovalTokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NonReturnApprovalToken.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_NonReturnApprovalToken *NonReturnApprovalTokenSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NonReturnApprovalToken.Contract.Mint(&_NonReturnApprovalToken.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_NonReturnApprovalToken *NonReturnApprovalTokenTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NonReturnApprovalToken.Contract.Mint(&_NonReturnApprovalToken.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NonReturnApprovalToken *NonReturnApprovalTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NonReturnApprovalToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NonReturnApprovalToken *NonReturnApprovalTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NonReturnApprovalToken.Contract.Transfer(&_NonReturnApprovalToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NonReturnApprovalToken *NonReturnApprovalTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NonReturnApprovalToken.Contract.Transfer(&_NonReturnApprovalToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NonReturnApprovalToken *NonReturnApprovalTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NonReturnApprovalToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NonReturnApprovalToken *NonReturnApprovalTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NonReturnApprovalToken.Contract.TransferFrom(&_NonReturnApprovalToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NonReturnApprovalToken *NonReturnApprovalTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NonReturnApprovalToken.Contract.TransferFrom(&_NonReturnApprovalToken.TransactOpts, from, to, value)
}

// NonReturnApprovalTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NonReturnApprovalToken contract.
type NonReturnApprovalTokenApprovalIterator struct {
	Event *NonReturnApprovalTokenApproval // Event containing the contract specifics and raw log

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
func (it *NonReturnApprovalTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NonReturnApprovalTokenApproval)
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
		it.Event = new(NonReturnApprovalTokenApproval)
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
func (it *NonReturnApprovalTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NonReturnApprovalTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NonReturnApprovalTokenApproval represents a Approval event raised by the NonReturnApprovalToken contract.
type NonReturnApprovalTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NonReturnApprovalToken *NonReturnApprovalTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NonReturnApprovalTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NonReturnApprovalToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NonReturnApprovalTokenApprovalIterator{contract: _NonReturnApprovalToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NonReturnApprovalToken *NonReturnApprovalTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NonReturnApprovalTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NonReturnApprovalToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NonReturnApprovalTokenApproval)
				if err := _NonReturnApprovalToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_NonReturnApprovalToken *NonReturnApprovalTokenFilterer) ParseApproval(log types.Log) (*NonReturnApprovalTokenApproval, error) {
	event := new(NonReturnApprovalTokenApproval)
	if err := _NonReturnApprovalToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NonReturnApprovalTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NonReturnApprovalToken contract.
type NonReturnApprovalTokenTransferIterator struct {
	Event *NonReturnApprovalTokenTransfer // Event containing the contract specifics and raw log

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
func (it *NonReturnApprovalTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NonReturnApprovalTokenTransfer)
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
		it.Event = new(NonReturnApprovalTokenTransfer)
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
func (it *NonReturnApprovalTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NonReturnApprovalTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NonReturnApprovalTokenTransfer represents a Transfer event raised by the NonReturnApprovalToken contract.
type NonReturnApprovalTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NonReturnApprovalToken *NonReturnApprovalTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NonReturnApprovalTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NonReturnApprovalToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NonReturnApprovalTokenTransferIterator{contract: _NonReturnApprovalToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NonReturnApprovalToken *NonReturnApprovalTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NonReturnApprovalTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NonReturnApprovalToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NonReturnApprovalTokenTransfer)
				if err := _NonReturnApprovalToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_NonReturnApprovalToken *NonReturnApprovalTokenFilterer) ParseTransfer(log types.Log) (*NonReturnApprovalTokenTransfer, error) {
	event := new(NonReturnApprovalTokenTransfer)
	if err := _NonReturnApprovalToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
