// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockerc721

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

// MockERC721MetaData contains all meta data concerning the MockERC721 contract.
var MockERC721MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f80fd5b506114458061001c5f395ff3fe6080604052600436106100d9575f3560e01c80636352211e1161007c578063a22cb46511610057578063a22cb46514610252578063b88d4fde14610271578063c87b56dd14610284578063e985e9c5146102a4575f80fd5b80636352211e146101f257806370a082311461021157806395d89b411461023e575f80fd5b8063095ea7b3116100b7578063095ea7b31461019857806323b872dd146101ad57806342842e0e146101c05780634cd88b76146101d3575f80fd5b806301ffc9a7146100dd57806306fdde0314610111578063081812fc14610132575b5f80fd5b3480156100e8575f80fd5b506100fc6100f7366004610df4565b6102f8565b60405190151581526020015b60405180910390f35b34801561011c575f80fd5b506101256103dc565b6040516101089190610e62565b34801561013d575f80fd5b5061017361014c366004610e74565b5f9081526004602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610108565b6101ab6101a6366004610eae565b61046b565b005b6101ab6101bb366004610ed6565b6105b7565b6101ab6101ce366004610ed6565b6108a6565b3480156101de575f80fd5b506101ab6101ed366004610ff0565b6109f6565b3480156101fd575f80fd5b5061017361020c366004610e74565b610aab565b34801561021c575f80fd5b5061023061022b366004611055565b610b3b565b604051908152602001610108565b348015610249575f80fd5b50610125610be1565b34801561025d575f80fd5b506101ab61026c36600461106e565b610bf0565b6101ab61027f3660046110a7565b610c86565b34801561028f575f80fd5b5061012561029e366004610e74565b50606090565b3480156102af575f80fd5b506100fc6102be36600461111e565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260056020908152604080832093909416825291909152205460ff1690565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316148061038a57507f80ac58cd000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b806103d657507f5b5e139f000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b60605f80546103ea9061114f565b80601f01602080910402602001604051908101604052809291908181526020018280546104169061114f565b80156104615780601f1061043857610100808354040283529160200191610461565b820191905f5260205f20905b81548152906001019060200180831161044457829003601f168201915b5050505050905090565b5f8181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff16338114806104cc575073ffffffffffffffffffffffffffffffffffffffff81165f90815260056020908152604080832033845290915290205460ff165b610537576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f4e4f545f415554484f52495a454400000000000000000000000000000000000060448201526064015b60405180910390fd5b5f8281526004602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff87811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b5f8181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff848116911614610646576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f57524f4e475f46524f4d00000000000000000000000000000000000000000000604482015260640161052e565b73ffffffffffffffffffffffffffffffffffffffff82166106c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f494e56414c49445f524543495049454e54000000000000000000000000000000604482015260640161052e565b3373ffffffffffffffffffffffffffffffffffffffff84161480610716575073ffffffffffffffffffffffffffffffffffffffff83165f90815260056020908152604080832033845290915290205460ff165b8061074357505f8181526004602052604090205473ffffffffffffffffffffffffffffffffffffffff1633145b6107a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f4e4f545f415554484f52495a4544000000000000000000000000000000000000604482015260640161052e565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526003602052604081208054916107d9836111cd565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f90815260036020526040812080549161080e83611201565b90915550505f818152600260209081526040808320805473ffffffffffffffffffffffffffffffffffffffff8088167fffffffffffffffffffffffff000000000000000000000000000000000000000092831681179093556004909452828520805490911690559051849391928716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6108b18383836105b7565b813b158061098b57506040517f150b7a020000000000000000000000000000000000000000000000000000000080825233600483015273ffffffffffffffffffffffffffffffffffffffff858116602484015260448301849052608060648401525f608484015290919084169063150b7a029060a4016020604051808303815f875af1158015610943573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109679190611238565b7fffffffff0000000000000000000000000000000000000000000000000000000016145b6109f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f554e534146455f524543495049454e5400000000000000000000000000000000604482015260640161052e565b505050565b60065460ff1615610a63576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f414c52454144595f494e495449414c495a454400000000000000000000000000604482015260640161052e565b5f610a6e838261129e565b506001610a7b828261129e565b5050600680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905550565b5f8181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff1680610b36576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e4f545f4d494e54454400000000000000000000000000000000000000000000604482015260640161052e565b919050565b5f73ffffffffffffffffffffffffffffffffffffffff8216610bb9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f5a45524f5f414444524553530000000000000000000000000000000000000000604482015260640161052e565b5073ffffffffffffffffffffffffffffffffffffffff165f9081526003602052604090205490565b6060600180546103ea9061114f565b335f81815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168085529083529281902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b610c918484846105b7565b823b1580610d5857506040517f150b7a02000000000000000000000000000000000000000000000000000000008082529073ffffffffffffffffffffffffffffffffffffffff85169063150b7a0290610cf49033908990889088906004016113b5565b6020604051808303815f875af1158015610d10573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d349190611238565b7fffffffff0000000000000000000000000000000000000000000000000000000016145b610dbe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f554e534146455f524543495049454e5400000000000000000000000000000000604482015260640161052e565b50505050565b7fffffffff0000000000000000000000000000000000000000000000000000000081168114610df1575f80fd5b50565b5f60208284031215610e04575f80fd5b8135610e0f81610dc4565b9392505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f610e0f6020830184610e16565b5f60208284031215610e84575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610b36575f80fd5b5f8060408385031215610ebf575f80fd5b610ec883610e8b565b946020939093013593505050565b5f805f60608486031215610ee8575f80fd5b610ef184610e8b565b9250610eff60208501610e8b565b929592945050506040919091013590565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8067ffffffffffffffff841115610f5757610f57610f10565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715610fa457610fa4610f10565b604052838152905080828401851015610fbb575f80fd5b838360208301375f60208583010152509392505050565b5f82601f830112610fe1575f80fd5b610e0f83833560208501610f3d565b5f8060408385031215611001575f80fd5b823567ffffffffffffffff811115611017575f80fd5b61102385828601610fd2565b925050602083013567ffffffffffffffff81111561103f575f80fd5b61104b85828601610fd2565b9150509250929050565b5f60208284031215611065575f80fd5b610e0f82610e8b565b5f806040838503121561107f575f80fd5b61108883610e8b565b91506020830135801515811461109c575f80fd5b809150509250929050565b5f805f80608085870312156110ba575f80fd5b6110c385610e8b565b93506110d160208601610e8b565b925060408501359150606085013567ffffffffffffffff8111156110f3575f80fd5b8501601f81018713611103575f80fd5b61111287823560208401610f3d565b91505092959194509250565b5f806040838503121561112f575f80fd5b61113883610e8b565b915061114660208401610e8b565b90509250929050565b600181811c9082168061116357607f821691505b60208210810361119a577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f816111db576111db6111a0565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611231576112316111a0565b5060010190565b5f60208284031215611248575f80fd5b8151610e0f81610dc4565b601f8211156109f157805f5260205f20601f840160051c810160208510156112785750805b601f840160051c820191505b81811015611297575f8155600101611284565b5050505050565b815167ffffffffffffffff8111156112b8576112b8610f10565b6112cc816112c6845461114f565b84611253565b6020601f82116001811461131d575f83156112e75750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611297565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b8281101561136a578785015182556020948501946001909201910161134a565b50848210156113a657868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b73ffffffffffffffffffffffffffffffffffffffff8516815273ffffffffffffffffffffffffffffffffffffffff84166020820152826040820152608060608201525f6114056080830184610e16565b969550505050505056fea2646970667358221220cf4b1f0b4883567666b196111b514c953badc5227770a8bcbbd0907dff948b3864736f6c634300081a0033",
}

// MockERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use MockERC721MetaData.ABI instead.
var MockERC721ABI = MockERC721MetaData.ABI

// MockERC721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockERC721MetaData.Bin instead.
var MockERC721Bin = MockERC721MetaData.Bin

// DeployMockERC721 deploys a new Ethereum contract, binding an instance of MockERC721 to it.
func DeployMockERC721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockERC721, error) {
	parsed, err := MockERC721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockERC721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockERC721{MockERC721Caller: MockERC721Caller{contract: contract}, MockERC721Transactor: MockERC721Transactor{contract: contract}, MockERC721Filterer: MockERC721Filterer{contract: contract}}, nil
}

// MockERC721 is an auto generated Go binding around an Ethereum contract.
type MockERC721 struct {
	MockERC721Caller     // Read-only binding to the contract
	MockERC721Transactor // Write-only binding to the contract
	MockERC721Filterer   // Log filterer for contract events
}

// MockERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type MockERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MockERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockERC721Session struct {
	Contract     *MockERC721       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockERC721CallerSession struct {
	Contract *MockERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MockERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockERC721TransactorSession struct {
	Contract     *MockERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MockERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type MockERC721Raw struct {
	Contract *MockERC721 // Generic contract binding to access the raw methods on
}

// MockERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockERC721CallerRaw struct {
	Contract *MockERC721Caller // Generic read-only contract binding to access the raw methods on
}

// MockERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockERC721TransactorRaw struct {
	Contract *MockERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMockERC721 creates a new instance of MockERC721, bound to a specific deployed contract.
func NewMockERC721(address common.Address, backend bind.ContractBackend) (*MockERC721, error) {
	contract, err := bindMockERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockERC721{MockERC721Caller: MockERC721Caller{contract: contract}, MockERC721Transactor: MockERC721Transactor{contract: contract}, MockERC721Filterer: MockERC721Filterer{contract: contract}}, nil
}

// NewMockERC721Caller creates a new read-only instance of MockERC721, bound to a specific deployed contract.
func NewMockERC721Caller(address common.Address, caller bind.ContractCaller) (*MockERC721Caller, error) {
	contract, err := bindMockERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockERC721Caller{contract: contract}, nil
}

// NewMockERC721Transactor creates a new write-only instance of MockERC721, bound to a specific deployed contract.
func NewMockERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*MockERC721Transactor, error) {
	contract, err := bindMockERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockERC721Transactor{contract: contract}, nil
}

// NewMockERC721Filterer creates a new log filterer instance of MockERC721, bound to a specific deployed contract.
func NewMockERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*MockERC721Filterer, error) {
	contract, err := bindMockERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockERC721Filterer{contract: contract}, nil
}

// bindMockERC721 binds a generic wrapper to an already deployed contract.
func bindMockERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockERC721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockERC721 *MockERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockERC721.Contract.MockERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockERC721 *MockERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockERC721.Contract.MockERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockERC721 *MockERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockERC721.Contract.MockERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockERC721 *MockERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockERC721 *MockERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockERC721 *MockERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MockERC721 *MockERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MockERC721 *MockERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MockERC721.Contract.BalanceOf(&_MockERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MockERC721 *MockERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MockERC721.Contract.BalanceOf(&_MockERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address)
func (_MockERC721 *MockERC721Caller) GetApproved(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "getApproved", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address)
func (_MockERC721 *MockERC721Session) GetApproved(id *big.Int) (common.Address, error) {
	return _MockERC721.Contract.GetApproved(&_MockERC721.CallOpts, id)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address)
func (_MockERC721 *MockERC721CallerSession) GetApproved(id *big.Int) (common.Address, error) {
	return _MockERC721.Contract.GetApproved(&_MockERC721.CallOpts, id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MockERC721 *MockERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MockERC721 *MockERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MockERC721.Contract.IsApprovedForAll(&_MockERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MockERC721 *MockERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MockERC721.Contract.IsApprovedForAll(&_MockERC721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockERC721 *MockERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockERC721 *MockERC721Session) Name() (string, error) {
	return _MockERC721.Contract.Name(&_MockERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockERC721 *MockERC721CallerSession) Name() (string, error) {
	return _MockERC721.Contract.Name(&_MockERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_MockERC721 *MockERC721Caller) OwnerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "ownerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_MockERC721 *MockERC721Session) OwnerOf(id *big.Int) (common.Address, error) {
	return _MockERC721.Contract.OwnerOf(&_MockERC721.CallOpts, id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_MockERC721 *MockERC721CallerSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _MockERC721.Contract.OwnerOf(&_MockERC721.CallOpts, id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockERC721 *MockERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockERC721 *MockERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockERC721.Contract.SupportsInterface(&_MockERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MockERC721 *MockERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockERC721.Contract.SupportsInterface(&_MockERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockERC721 *MockERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockERC721 *MockERC721Session) Symbol() (string, error) {
	return _MockERC721.Contract.Symbol(&_MockERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockERC721 *MockERC721CallerSession) Symbol() (string, error) {
	return _MockERC721.Contract.Symbol(&_MockERC721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_MockERC721 *MockERC721Caller) TokenURI(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _MockERC721.contract.Call(opts, &out, "tokenURI", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_MockERC721 *MockERC721Session) TokenURI(id *big.Int) (string, error) {
	return _MockERC721.Contract.TokenURI(&_MockERC721.CallOpts, id)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_MockERC721 *MockERC721CallerSession) TokenURI(id *big.Int) (string, error) {
	return _MockERC721.Contract.TokenURI(&_MockERC721.CallOpts, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) payable returns()
func (_MockERC721 *MockERC721Transactor) Approve(opts *bind.TransactOpts, spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.contract.Transact(opts, "approve", spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) payable returns()
func (_MockERC721 *MockERC721Session) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.Contract.Approve(&_MockERC721.TransactOpts, spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) payable returns()
func (_MockERC721 *MockERC721TransactorSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.Contract.Approve(&_MockERC721.TransactOpts, spender, id)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_MockERC721 *MockERC721Transactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _MockERC721.contract.Transact(opts, "initialize", name_, symbol_)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_MockERC721 *MockERC721Session) Initialize(name_ string, symbol_ string) (*types.Transaction, error) {
	return _MockERC721.Contract.Initialize(&_MockERC721.TransactOpts, name_, symbol_)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_MockERC721 *MockERC721TransactorSession) Initialize(name_ string, symbol_ string) (*types.Transaction, error) {
	return _MockERC721.Contract.Initialize(&_MockERC721.TransactOpts, name_, symbol_)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) payable returns()
func (_MockERC721 *MockERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.contract.Transact(opts, "safeTransferFrom", from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) payable returns()
func (_MockERC721 *MockERC721Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.Contract.SafeTransferFrom(&_MockERC721.TransactOpts, from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) payable returns()
func (_MockERC721 *MockERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.Contract.SafeTransferFrom(&_MockERC721.TransactOpts, from, to, id)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) payable returns()
func (_MockERC721 *MockERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _MockERC721.contract.Transact(opts, "safeTransferFrom0", from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) payable returns()
func (_MockERC721 *MockERC721Session) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _MockERC721.Contract.SafeTransferFrom0(&_MockERC721.TransactOpts, from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) payable returns()
func (_MockERC721 *MockERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _MockERC721.Contract.SafeTransferFrom0(&_MockERC721.TransactOpts, from, to, id, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockERC721 *MockERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockERC721 *MockERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockERC721.Contract.SetApprovalForAll(&_MockERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MockERC721 *MockERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MockERC721.Contract.SetApprovalForAll(&_MockERC721.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) payable returns()
func (_MockERC721 *MockERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.contract.Transact(opts, "transferFrom", from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) payable returns()
func (_MockERC721 *MockERC721Session) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.Contract.TransferFrom(&_MockERC721.TransactOpts, from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) payable returns()
func (_MockERC721 *MockERC721TransactorSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _MockERC721.Contract.TransferFrom(&_MockERC721.TransactOpts, from, to, id)
}

// MockERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MockERC721 contract.
type MockERC721ApprovalIterator struct {
	Event *MockERC721Approval // Event containing the contract specifics and raw log

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
func (it *MockERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockERC721Approval)
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
		it.Event = new(MockERC721Approval)
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
func (it *MockERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockERC721Approval represents a Approval event raised by the MockERC721 contract.
type MockERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_MockERC721 *MockERC721Filterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*MockERC721ApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _MockERC721.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MockERC721ApprovalIterator{contract: _MockERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_MockERC721 *MockERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MockERC721Approval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _MockERC721.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockERC721Approval)
				if err := _MockERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_MockERC721 *MockERC721Filterer) ParseApproval(log types.Log) (*MockERC721Approval, error) {
	event := new(MockERC721Approval)
	if err := _MockERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MockERC721 contract.
type MockERC721ApprovalForAllIterator struct {
	Event *MockERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MockERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockERC721ApprovalForAll)
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
		it.Event = new(MockERC721ApprovalForAll)
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
func (it *MockERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockERC721ApprovalForAll represents a ApprovalForAll event raised by the MockERC721 contract.
type MockERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_MockERC721 *MockERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*MockERC721ApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _MockERC721.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &MockERC721ApprovalForAllIterator{contract: _MockERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_MockERC721 *MockERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MockERC721ApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _MockERC721.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockERC721ApprovalForAll)
				if err := _MockERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_MockERC721 *MockERC721Filterer) ParseApprovalForAll(log types.Log) (*MockERC721ApprovalForAll, error) {
	event := new(MockERC721ApprovalForAll)
	if err := _MockERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MockERC721 contract.
type MockERC721TransferIterator struct {
	Event *MockERC721Transfer // Event containing the contract specifics and raw log

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
func (it *MockERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockERC721Transfer)
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
		it.Event = new(MockERC721Transfer)
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
func (it *MockERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockERC721Transfer represents a Transfer event raised by the MockERC721 contract.
type MockERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_MockERC721 *MockERC721Filterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*MockERC721TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _MockERC721.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MockERC721TransferIterator{contract: _MockERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_MockERC721 *MockERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MockERC721Transfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _MockERC721.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockERC721Transfer)
				if err := _MockERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_MockERC721 *MockERC721Filterer) ParseTransfer(log types.Log) (*MockERC721Transfer, error) {
	event := new(MockERC721Transfer)
	if err := _MockERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
