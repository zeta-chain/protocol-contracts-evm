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
	Bin: "0x6080604052348015600f57600080fd5b506114b88061001f6000396000f3fe6080604052600436106100dd5760003560e01c80636352211e1161007f578063a22cb46511610059578063a22cb4651461025f578063b88d4fde1461027f578063c87b56dd14610292578063e985e9c5146102b357600080fd5b80636352211e146101fc57806370a082311461021c57806395d89b411461024a57600080fd5b8063095ea7b3116100bb578063095ea7b3146101a157806323b872dd146101b657806342842e0e146101c95780634cd88b76146101dc57600080fd5b806301ffc9a7146100e257806306fdde0314610117578063081812fc14610139575b600080fd5b3480156100ee57600080fd5b506101026100fd366004610e1f565b610309565b60405190151581526020015b60405180910390f35b34801561012357600080fd5b5061012c6103ee565b60405161010e9190610ea7565b34801561014557600080fd5b5061017c610154366004610eba565b60009081526004602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010e565b6101b46101af366004610ef7565b610480565b005b6101b46101c4366004610f21565b6105cf565b6101b46101d7366004610f21565b6108c4565b3480156101e857600080fd5b506101b46101f7366004611045565b610a18565b34801561020857600080fd5b5061017c610217366004610eba565b610ace565b34801561022857600080fd5b5061023c6102373660046110ae565b610b5f565b60405190815260200161010e565b34801561025657600080fd5b5061012c610c07565b34801561026b57600080fd5b506101b461027a3660046110c9565b610c16565b6101b461028d366004611105565b610cad565b34801561029e57600080fd5b5061012c6102ad366004610eba565b50606090565b3480156102bf57600080fd5b506101026102ce366004611181565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260056020908152604080832093909416825291909152205460ff1690565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316148061039c57507f80ac58cd000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b806103e857507f5b5e139f000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6060600080546103fd906111b4565b80601f0160208091040260200160405190810160405280929190818152602001828054610429906111b4565b80156104765780601f1061044b57610100808354040283529160200191610476565b820191906000526020600020905b81548152906001019060200180831161045957829003601f168201915b5050505050905090565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff16338114806104e3575073ffffffffffffffffffffffffffffffffffffffff8116600090815260056020908152604080832033845290915290205460ff165b61054e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f4e4f545f415554484f52495a454400000000000000000000000000000000000060448201526064015b60405180910390fd5b60008281526004602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff87811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff84811691161461065f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f57524f4e475f46524f4d000000000000000000000000000000000000000000006044820152606401610545565b73ffffffffffffffffffffffffffffffffffffffff82166106dc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f494e56414c49445f524543495049454e540000000000000000000000000000006044820152606401610545565b3373ffffffffffffffffffffffffffffffffffffffff84161480610730575073ffffffffffffffffffffffffffffffffffffffff8316600090815260056020908152604080832033845290915290205460ff165b8061075e575060008181526004602052604090205473ffffffffffffffffffffffffffffffffffffffff1633145b6107c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f4e4f545f415554484f52495a45440000000000000000000000000000000000006044820152606401610545565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604081208054916107f583611236565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260036020526040812080549161082b8361126b565b90915550506000818152600260209081526040808320805473ffffffffffffffffffffffffffffffffffffffff8088167fffffffffffffffffffffffff000000000000000000000000000000000000000092831681179093556004909452828520805490911690559051849391928716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6108cf8383836105cf565b813b15806109ad57506040517f150b7a020000000000000000000000000000000000000000000000000000000080825233600483015273ffffffffffffffffffffffffffffffffffffffff858116602484015260448301849052608060648401526000608484015290919084169063150b7a029060a4016020604051808303816000875af1158015610965573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061098991906112a3565b7fffffffff0000000000000000000000000000000000000000000000000000000016145b610a13576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f554e534146455f524543495049454e54000000000000000000000000000000006044820152606401610545565b505050565b60065460ff1615610a85576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f414c52454144595f494e495449414c495a4544000000000000000000000000006044820152606401610545565b6000610a91838261130e565b506001610a9e828261130e565b5050600680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905550565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff1680610b5a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e4f545f4d494e544544000000000000000000000000000000000000000000006044820152606401610545565b919050565b600073ffffffffffffffffffffffffffffffffffffffff8216610bde576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f5a45524f5f4144445245535300000000000000000000000000000000000000006044820152606401610545565b5073ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6060600180546103fd906111b4565b33600081815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168085529083529281902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b610cb88484846105cf565b823b1580610d8257506040517f150b7a02000000000000000000000000000000000000000000000000000000008082529073ffffffffffffffffffffffffffffffffffffffff85169063150b7a0290610d1b903390899088908890600401611427565b6020604051808303816000875af1158015610d3a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d5e91906112a3565b7fffffffff0000000000000000000000000000000000000000000000000000000016145b610de8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f554e534146455f524543495049454e54000000000000000000000000000000006044820152606401610545565b50505050565b7fffffffff0000000000000000000000000000000000000000000000000000000081168114610e1c57600080fd5b50565b600060208284031215610e3157600080fd5b8135610e3c81610dee565b9392505050565b6000815180845260005b81811015610e6957602081850181015186830182015201610e4d565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610e3c6020830184610e43565b600060208284031215610ecc57600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610b5a57600080fd5b60008060408385031215610f0a57600080fd5b610f1383610ed3565b946020939093013593505050565b600080600060608486031215610f3657600080fd5b610f3f84610ed3565b9250610f4d60208501610ed3565b929592945050506040919091013590565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff841115610fa857610fa8610f5e565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715610ff557610ff5610f5e565b60405283815290508082840185101561100d57600080fd5b83836020830137600060208583010152509392505050565b600082601f83011261103657600080fd5b610e3c83833560208501610f8d565b6000806040838503121561105857600080fd5b823567ffffffffffffffff81111561106f57600080fd5b61107b85828601611025565b925050602083013567ffffffffffffffff81111561109857600080fd5b6110a485828601611025565b9150509250929050565b6000602082840312156110c057600080fd5b610e3c82610ed3565b600080604083850312156110dc57600080fd5b6110e583610ed3565b9150602083013580151581146110fa57600080fd5b809150509250929050565b6000806000806080858703121561111b57600080fd5b61112485610ed3565b935061113260208601610ed3565b925060408501359150606085013567ffffffffffffffff81111561115557600080fd5b8501601f8101871361116657600080fd5b61117587823560208401610f8d565b91505092959194509250565b6000806040838503121561119457600080fd5b61119d83610ed3565b91506111ab60208401610ed3565b90509250929050565b600181811c908216806111c857607f821691505b602082108103611201577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008161124557611245611207565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361129c5761129c611207565b5060010190565b6000602082840312156112b557600080fd5b8151610e3c81610dee565b601f821115610a1357806000526020600020601f840160051c810160208510156112e75750805b601f840160051c820191505b8181101561130757600081556001016112f3565b5050505050565b815167ffffffffffffffff81111561132857611328610f5e565b61133c8161133684546111b4565b846112c0565b6020601f82116001811461138e57600083156113585750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611307565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156113dc57878501518255602094850194600190920191016113bc565b508482101561141857868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b73ffffffffffffffffffffffffffffffffffffffff8516815273ffffffffffffffffffffffffffffffffffffffff841660208201528260408201526080606082015260006114786080830184610e43565b969550505050505056fea264697066735822122056107d06a5207b08640ffc20c4b6d71dd4c40dff8c35fe204a9a777b7c70e69b64736f6c634300081a0033",
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
