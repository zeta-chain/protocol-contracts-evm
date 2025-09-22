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
	Bin: "0x608060405234801561000f575f80fd5b50604051610c49380380610c4983398101604081905261002e916100f4565b81818181600361003e83826101dd565b50600461004b82826101dd565b50505050505050610297565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261007a575f80fd5b81516001600160401b0381111561009357610093610057565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c1576100c1610057565b6040528181528382016020018510156100d8575f80fd5b8160208501602083015e5f918101602001919091529392505050565b5f8060408385031215610105575f80fd5b82516001600160401b0381111561011a575f80fd5b6101268582860161006b565b602085015190935090506001600160401b03811115610143575f80fd5b61014f8582860161006b565b9150509250929050565b600181811c9082168061016d57607f821691505b60208210810361018b57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101d857805f5260205f20601f840160051c810160208510156101b65750805b601f840160051c820191505b818110156101d5575f81556001016101c2565b50505b505050565b81516001600160401b038111156101f6576101f6610057565b61020a816102048454610159565b84610191565b6020601f82116001811461023c575f83156102255750848201515b5f19600385901b1c1916600184901b1784556101d5565b5f84815260208120601f198516915b8281101561026b578785015182556020948501946001909201910161024b565b508482101561028857868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6109a5806102a45f395ff3fe608060405234801561000f575f80fd5b50600436106100b9575f3560e01c806340c10f191161007257806395d89b411161005857806395d89b411461017c578063a9059cbb14610184578063dd62ed3e14610197575f80fd5b806340c10f191461013257806370a0823114610147575f80fd5b806318160ddd116100a257806318160ddd146100fe57806323b872dd14610110578063313ce56714610123575f80fd5b806306fdde03146100bd578063095ea7b3146100db575b5f80fd5b6100c56101dc565b6040516100d291906107bf565b60405180910390f35b6100ee6100e936600461083a565b61026c565b60405190151581526020016100d2565b6002545b6040519081526020016100d2565b6100ee61011e366004610862565b61028b565b604051601281526020016100d2565b61014561014036600461083a565b6102ae565b005b61010261015536600461089c565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b6100c56102bc565b6100ee61019236600461083a565b6102cb565b6101026101a53660046108b5565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b6060600380546101eb906108e6565b80601f0160208091040260200160405190810160405280929190818152602001828054610217906108e6565b80156102625780601f1061023957610100808354040283529160200191610262565b820191905f5260205f20905b81548152906001019060200180831161024557829003601f168201915b5050505050905090565b5f808211610278575f80fd5b61028283836102e2565b90505b92915050565b5f336102988582856102ef565b6102a38585856103c2565b506001949350505050565b6102b88282610470565b5050565b6060600480546101eb906108e6565b5f336102d88185856103c2565b5060019392505050565b5f336102d88185856104ca565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156103bc57818110156103ae576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bc84848484035f6104d3565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610411576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f60048201526024016103a5565b73ffffffffffffffffffffffffffffffffffffffff8216610460576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016103a5565b61046b838383610618565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166104bf576040517fec442f050000000000000000000000000000000000000000000000000000000081525f60048201526024016103a5565b6102b85f8383610618565b61046b83838360015b73ffffffffffffffffffffffffffffffffffffffff8416610522576040517fe602df050000000000000000000000000000000000000000000000000000000081525f60048201526024016103a5565b73ffffffffffffffffffffffffffffffffffffffff8316610571576040517f94280d620000000000000000000000000000000000000000000000000000000081525f60048201526024016103a5565b73ffffffffffffffffffffffffffffffffffffffff8085165f90815260016020908152604080832093871683529290522082905580156103bc578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060a91815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064f578060025f8282546106449190610937565b909155506106ff9050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260208190526040902054818110156106d4576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a5565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072857600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610835575f80fd5b919050565b5f806040838503121561084b575f80fd5b61085483610812565b946020939093013593505050565b5f805f60608486031215610874575f80fd5b61087d84610812565b925061088b60208501610812565b929592945050506040919091013590565b5f602082840312156108ac575f80fd5b61028282610812565b5f80604083850312156108c6575f80fd5b6108cf83610812565b91506108dd60208401610812565b90509250929050565b600181811c908216806108fa57607f821691505b602082108103610931577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b80820180821115610285577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea26469706673582212208dd9c783e1e1c143434dee6e3109982215f9998235c611bae6c2378b81aeb5f664736f6c634300081a0033",
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
