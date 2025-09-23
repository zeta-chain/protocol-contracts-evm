// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zeta

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

// ZetaEthMetaData contains all meta data concerning the ZetaEth contract.
var ZetaEthMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610deb380380610deb83398101604081905261002e91610225565b604051806040016040528060048152602001635a65746160e01b815250604051806040016040528060048152602001635a45544160e01b815250816003908161007791906102f4565b50600461008482826102f4565b5050506100b6826100996100bd60201b60201c565b6100a79060ff16600a6104a7565b6100b190846104b9565b6100c2565b50506104e3565b601290565b6001600160a01b0382166100f05760405163ec442f0560e01b81525f60048201526024015b60405180910390fd5b6100fb5f83836100ff565b5050565b6001600160a01b038316610129578060025f82825461011e91906104d0565b909155506101999050565b6001600160a01b0383165f908152602081905260409020548181101561017b5760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016100e7565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b0382166101b5576002805482900390556101d3565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161021891815260200190565b60405180910390a3505050565b5f8060408385031215610236575f80fd5b82516001600160a01b038116811461024c575f80fd5b6020939093015192949293505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061028457607f821691505b6020821081036102a257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156102ef57805f5260205f20601f840160051c810160208510156102cd5750805b601f840160051c820191505b818110156102ec575f81556001016102d9565b50505b505050565b81516001600160401b0381111561030d5761030d61025c565b6103218161031b8454610270565b846102a8565b6020601f821160018114610353575f831561033c5750848201515b5f19600385901b1c1916600184901b1784556102ec565b5f84815260208120601f198516915b828110156103825787850151825560209485019460019092019101610362565b508482101561039f57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b6001815b60018411156103fd578085048111156103e1576103e16103ae565b60018416156103ef57908102905b60019390931c9280026103c6565b935093915050565b5f82610413575060016104a1565b8161041f57505f6104a1565b8160018114610435576002811461043f5761045b565b60019150506104a1565b60ff841115610450576104506103ae565b50506001821b6104a1565b5060208310610133831016604e8410600b841016171561047e575081810a6104a1565b61048a5f1984846103c2565b805f190482111561049d5761049d6103ae565b0290505b92915050565b5f6104b28383610405565b9392505050565b80820281158282048414176104a1576104a16103ae565b808201808211156104a1576104a16103ae565b6108fb806104f05f395ff3fe608060405234801561000f575f80fd5b506004361061009f575f3560e01c8063313ce5671161007257806395d89b411161005857806395d89b411461014d578063a9059cbb14610155578063dd62ed3e14610168575f80fd5b8063313ce5671461010957806370a0823114610118575f80fd5b806306fdde03146100a3578063095ea7b3146100c157806318160ddd146100e457806323b872dd146100f6575b5f80fd5b6100ab6101ad565b6040516100b8919061070e565b60405180910390f35b6100d46100cf366004610789565b61023d565b60405190151581526020016100b8565b6002545b6040519081526020016100b8565b6100d46101043660046107b1565b610256565b604051601281526020016100b8565b6100e86101263660046107eb565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b6100ab610279565b6100d4610163366004610789565b610288565b6100e861017636600461080b565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b6060600380546101bc9061083c565b80601f01602080910402602001604051908101604052809291908181526020018280546101e89061083c565b80156102335780601f1061020a57610100808354040283529160200191610233565b820191905f5260205f20905b81548152906001019060200180831161021657829003601f168201915b5050505050905090565b5f3361024a818585610295565b60019150505b92915050565b5f336102638582856102a7565b61026e858585610379565b506001949350505050565b6060600480546101bc9061083c565b5f3361024a818585610379565b6102a28383836001610422565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103735781811015610365576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b61037384848484035f610422565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166103c8576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f600482015260240161035c565b73ffffffffffffffffffffffffffffffffffffffff8216610417576040517fec442f050000000000000000000000000000000000000000000000000000000081525f600482015260240161035c565b6102a2838383610567565b73ffffffffffffffffffffffffffffffffffffffff8416610471576040517fe602df050000000000000000000000000000000000000000000000000000000081525f600482015260240161035c565b73ffffffffffffffffffffffffffffffffffffffff83166104c0576040517f94280d620000000000000000000000000000000000000000000000000000000081525f600482015260240161035c565b73ffffffffffffffffffffffffffffffffffffffff8085165f9081526001602090815260408083209387168352929052208290558015610373578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161055991815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661059e578060025f828254610593919061088d565b9091555061064e9050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020819052604090205481811015610623576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602481018290526044810183905260640161035c565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610677576002805482900390556106a2565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161070191815260200190565b60405180910390a3505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610784575f80fd5b919050565b5f806040838503121561079a575f80fd5b6107a383610761565b946020939093013593505050565b5f805f606084860312156107c3575f80fd5b6107cc84610761565b92506107da60208501610761565b929592945050506040919091013590565b5f602082840312156107fb575f80fd5b61080482610761565b9392505050565b5f806040838503121561081c575f80fd5b61082583610761565b915061083360208401610761565b90509250929050565b600181811c9082168061085057607f821691505b602082108103610887577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b80820180821115610250577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea2646970667358221220df00c92d68103bdee848c88559acde5541cde3acd4d54d1b3a989c4cc790816164736f6c634300081a0033",
}

// ZetaEthABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaEthMetaData.ABI instead.
var ZetaEthABI = ZetaEthMetaData.ABI

// ZetaEthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaEthMetaData.Bin instead.
var ZetaEthBin = ZetaEthMetaData.Bin

// DeployZetaEth deploys a new Ethereum contract, binding an instance of ZetaEth to it.
func DeployZetaEth(auth *bind.TransactOpts, backend bind.ContractBackend, creator common.Address, initialSupply *big.Int) (common.Address, *types.Transaction, *ZetaEth, error) {
	parsed, err := ZetaEthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaEthBin), backend, creator, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaEth{ZetaEthCaller: ZetaEthCaller{contract: contract}, ZetaEthTransactor: ZetaEthTransactor{contract: contract}, ZetaEthFilterer: ZetaEthFilterer{contract: contract}}, nil
}

// ZetaEth is an auto generated Go binding around an Ethereum contract.
type ZetaEth struct {
	ZetaEthCaller     // Read-only binding to the contract
	ZetaEthTransactor // Write-only binding to the contract
	ZetaEthFilterer   // Log filterer for contract events
}

// ZetaEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaEthSession struct {
	Contract     *ZetaEth          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZetaEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaEthCallerSession struct {
	Contract *ZetaEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ZetaEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaEthTransactorSession struct {
	Contract     *ZetaEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZetaEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaEthRaw struct {
	Contract *ZetaEth // Generic contract binding to access the raw methods on
}

// ZetaEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaEthCallerRaw struct {
	Contract *ZetaEthCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaEthTransactorRaw struct {
	Contract *ZetaEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaEth creates a new instance of ZetaEth, bound to a specific deployed contract.
func NewZetaEth(address common.Address, backend bind.ContractBackend) (*ZetaEth, error) {
	contract, err := bindZetaEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaEth{ZetaEthCaller: ZetaEthCaller{contract: contract}, ZetaEthTransactor: ZetaEthTransactor{contract: contract}, ZetaEthFilterer: ZetaEthFilterer{contract: contract}}, nil
}

// NewZetaEthCaller creates a new read-only instance of ZetaEth, bound to a specific deployed contract.
func NewZetaEthCaller(address common.Address, caller bind.ContractCaller) (*ZetaEthCaller, error) {
	contract, err := bindZetaEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaEthCaller{contract: contract}, nil
}

// NewZetaEthTransactor creates a new write-only instance of ZetaEth, bound to a specific deployed contract.
func NewZetaEthTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaEthTransactor, error) {
	contract, err := bindZetaEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaEthTransactor{contract: contract}, nil
}

// NewZetaEthFilterer creates a new log filterer instance of ZetaEth, bound to a specific deployed contract.
func NewZetaEthFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaEthFilterer, error) {
	contract, err := bindZetaEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaEthFilterer{contract: contract}, nil
}

// bindZetaEth binds a generic wrapper to an already deployed contract.
func bindZetaEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaEthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaEth *ZetaEthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaEth.Contract.ZetaEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaEth *ZetaEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaEth.Contract.ZetaEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaEth *ZetaEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaEth.Contract.ZetaEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaEth *ZetaEthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaEth *ZetaEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaEth *ZetaEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaEth.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZetaEth *ZetaEthCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZetaEth.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZetaEth *ZetaEthSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZetaEth.Contract.Allowance(&_ZetaEth.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZetaEth *ZetaEthCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZetaEth.Contract.Allowance(&_ZetaEth.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZetaEth *ZetaEthCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZetaEth.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZetaEth *ZetaEthSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZetaEth.Contract.BalanceOf(&_ZetaEth.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZetaEth *ZetaEthCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZetaEth.Contract.BalanceOf(&_ZetaEth.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZetaEth *ZetaEthCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZetaEth.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZetaEth *ZetaEthSession) Decimals() (uint8, error) {
	return _ZetaEth.Contract.Decimals(&_ZetaEth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZetaEth *ZetaEthCallerSession) Decimals() (uint8, error) {
	return _ZetaEth.Contract.Decimals(&_ZetaEth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZetaEth *ZetaEthCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZetaEth.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZetaEth *ZetaEthSession) Name() (string, error) {
	return _ZetaEth.Contract.Name(&_ZetaEth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZetaEth *ZetaEthCallerSession) Name() (string, error) {
	return _ZetaEth.Contract.Name(&_ZetaEth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZetaEth *ZetaEthCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZetaEth.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZetaEth *ZetaEthSession) Symbol() (string, error) {
	return _ZetaEth.Contract.Symbol(&_ZetaEth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZetaEth *ZetaEthCallerSession) Symbol() (string, error) {
	return _ZetaEth.Contract.Symbol(&_ZetaEth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZetaEth *ZetaEthCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZetaEth.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZetaEth *ZetaEthSession) TotalSupply() (*big.Int, error) {
	return _ZetaEth.Contract.TotalSupply(&_ZetaEth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZetaEth *ZetaEthCallerSession) TotalSupply() (*big.Int, error) {
	return _ZetaEth.Contract.TotalSupply(&_ZetaEth.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ZetaEth *ZetaEthTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaEth.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ZetaEth *ZetaEthSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.Approve(&_ZetaEth.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ZetaEth *ZetaEthTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.Approve(&_ZetaEth.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ZetaEth *ZetaEthTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaEth.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ZetaEth *ZetaEthSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.Transfer(&_ZetaEth.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ZetaEth *ZetaEthTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.Transfer(&_ZetaEth.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ZetaEth *ZetaEthTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaEth.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ZetaEth *ZetaEthSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.TransferFrom(&_ZetaEth.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ZetaEth *ZetaEthTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaEth.Contract.TransferFrom(&_ZetaEth.TransactOpts, from, to, value)
}

// ZetaEthApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZetaEth contract.
type ZetaEthApprovalIterator struct {
	Event *ZetaEthApproval // Event containing the contract specifics and raw log

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
func (it *ZetaEthApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaEthApproval)
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
		it.Event = new(ZetaEthApproval)
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
func (it *ZetaEthApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaEthApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaEthApproval represents a Approval event raised by the ZetaEth contract.
type ZetaEthApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZetaEth *ZetaEthFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ZetaEthApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZetaEth.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZetaEthApprovalIterator{contract: _ZetaEth.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZetaEth *ZetaEthFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZetaEthApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZetaEth.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaEthApproval)
				if err := _ZetaEth.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ZetaEth *ZetaEthFilterer) ParseApproval(log types.Log) (*ZetaEthApproval, error) {
	event := new(ZetaEthApproval)
	if err := _ZetaEth.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaEthTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZetaEth contract.
type ZetaEthTransferIterator struct {
	Event *ZetaEthTransfer // Event containing the contract specifics and raw log

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
func (it *ZetaEthTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaEthTransfer)
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
		it.Event = new(ZetaEthTransfer)
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
func (it *ZetaEthTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaEthTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaEthTransfer represents a Transfer event raised by the ZetaEth contract.
type ZetaEthTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZetaEth *ZetaEthFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZetaEthTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaEth.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaEthTransferIterator{contract: _ZetaEth.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZetaEth *ZetaEthFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZetaEthTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaEth.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaEthTransfer)
				if err := _ZetaEth.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ZetaEth *ZetaEthFilterer) ParseTransfer(log types.Log) (*ZetaEthTransfer, error) {
	event := new(ZetaEthTransfer)
	if err := _ZetaEth.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
