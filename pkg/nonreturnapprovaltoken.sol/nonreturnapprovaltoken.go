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
	Bin: "0x60806040526002805460ff1916601217905534801561001c575f80fd5b50604051610b8d380380610b8d83398101604081905261003b916100f8565b5f61004683826101e1565b50600161005382826101e1565b50505061029b565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261007e575f80fd5b81516001600160401b038111156100975761009761005b565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c5576100c561005b565b6040528181528382016020018510156100dc575f80fd5b8160208501602083015e5f918101602001919091529392505050565b5f8060408385031215610109575f80fd5b82516001600160401b0381111561011e575f80fd5b61012a8582860161006f565b602085015190935090506001600160401b03811115610147575f80fd5b6101538582860161006f565b9150509250929050565b600181811c9082168061017157607f821691505b60208210810361018f57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101dc57805f5260205f20601f840160051c810160208510156101ba5750805b601f840160051c820191505b818110156101d9575f81556001016101c6565b50505b505050565b81516001600160401b038111156101fa576101fa61005b565b61020e81610208845461015d565b84610195565b6020601f821160018114610240575f83156102295750848201515b5f19600385901b1c1916600184901b1784556101d9565b5f84815260208120601f198516915b8281101561026f578785015182556020948501946001909201910161024f565b508482101561028c57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6108e5806102a85f395ff3fe608060405234801561000f575f80fd5b50600436106100b9575f3560e01c806340c10f191161007257806395d89b411161005857806395d89b411461017b578063a9059cbb14610183578063dd62ed3e14610196575f80fd5b806340c10f191461014957806370a082311461015c575f80fd5b806318160ddd116100a257806318160ddd146100f057806323b872dd14610107578063313ce5671461012a575f80fd5b806306fdde03146100bd578063095ea7b3146100db575b5f80fd5b6100c56101c0565b6040516100d291906106dd565b60405180910390f35b6100ee6100e9366004610758565b61024b565b005b6100f960035481565b6040519081526020016100d2565b61011a610115366004610780565b6102b7565b60405190151581526020016100d2565b6002546101379060ff1681565b60405160ff90911681526020016100d2565b6100ee610157366004610758565b61050c565b6100f961016a3660046107ba565b60046020525f908152604090205481565b6100c56105a5565b61011a610191366004610758565b6105b2565b6100f96101a43660046107da565b600560209081525f928352604080842090915290825290205481565b5f80546101cc9061080b565b80601f01602080910402602001604051908101604052809291908181526020018280546101f89061080b565b80156102435780601f1061021a57610100808354040283529160200191610243565b820191905f5260205f20905b81548152906001019060200180831161022657829003601f168201915b505050505081565b335f81815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff871680855290835292819020859055518481529192917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526004602052604081205482111561034a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e636500000000000000000000000060448201526064015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff84165f9081526005602090815260408083203384529091529020548211156103e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e73756666696369656e7420616c6c6f77616e6365000000000000000000006044820152606401610341565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526004602052604081208054849290610417908490610889565b909155505073ffffffffffffffffffffffffffffffffffffffff83165f908152600460205260408120805484929061045090849061089c565b909155505073ffffffffffffffffffffffffffffffffffffffff84165f90815260056020908152604080832033845290915281208054849290610494908490610889565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516104fa91815260200190565b60405180910390a35060019392505050565b8060035f82825461051d919061089c565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f908152600460205260408120805483929061055690849061089c565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316905f907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016102ab565b600180546101cc9061080b565b335f9081526004602052604081205482111561062a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f496e73756666696369656e742062616c616e63650000000000000000000000006044820152606401610341565b335f9081526004602052604081208054849290610648908490610889565b909155505073ffffffffffffffffffffffffffffffffffffffff83165f908152600460205260408120805484929061068190849061089c565b909155505060405182815273ffffffffffffffffffffffffffffffffffffffff84169033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35060015b92915050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610753575f80fd5b919050565b5f8060408385031215610769575f80fd5b61077283610730565b946020939093013593505050565b5f805f60608486031215610792575f80fd5b61079b84610730565b92506107a960208501610730565b929592945050506040919091013590565b5f602082840312156107ca575f80fd5b6107d382610730565b9392505050565b5f80604083850312156107eb575f80fd5b6107f483610730565b915061080260208401610730565b90509250929050565b600181811c9082168061081f57607f821691505b602082108103610856577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b818103818111156106d7576106d761085c565b808201808211156106d7576106d761085c56fea26469706673582212200cfd2a3b3ac61410f384f5158efa9e45b9e8b44cb7226c0b479f845f3c50d43664736f6c634300081a0033",
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
