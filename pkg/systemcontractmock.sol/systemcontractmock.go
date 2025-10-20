// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package systemcontractmock

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

// SystemContractMockMetaData contains all meta data concerning the SystemContractMock contract.
var SystemContractMockMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"wzeta_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uniswapv2Factory_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uniswapv2Router02_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gasCoinZRC20ByChainId\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasPriceByChainId\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasZetaPoolByChainId\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onCrossChainCall\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasCoinZRC20\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zrc20\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasPrice\",\"inputs\":[{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWZETAContractAddress\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"uniswapv2FactoryAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"uniswapv2PairFor\",\"inputs\":[{\"name\":\"factory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenA\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenB\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"pair\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"uniswapv2Router02Address\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"wZetaContractAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"SetGasCoin\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetGasPrice\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetGasZetaPool\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetWZeta\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SystemContractDeployed\",\"inputs\":[],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotFungibleModule\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CantBeIdenticalAddresses\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CantBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610b3f380380610b3f83398101604081905261002f916100b9565b600380546001600160a01b038086166001600160a01b0319928316179092556004805485841690831617905560058054928416929091169190911790556040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac590600090a15050506100fc565b80516001600160a01b03811681146100b457600080fd5b919050565b6000806000606084860312156100ce57600080fd5b6100d78461009d565b92506100e56020850161009d565b91506100f36040850161009d565b90509250925092565b610a348061010b6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806397770dff11610081578063d7fd7afb1161005b578063d7fd7afb146101f2578063d936a01214610220578063ee2815ba1461024057600080fd5b806397770dff146101b9578063a7cb0507146101cc578063c63585cc146101df57600080fd5b8063513a9c05116100b2578063513a9c0514610143578063569541b914610179578063842da36d1461019957600080fd5b80630be15547146100ce5780633c669d551461012e575b600080fd5b6101046100dc36600461071e565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61014161013c366004610760565b610253565b005b61010461015136600461071e565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101049073ffffffffffffffffffffffffffffffffffffffff1681565b6005546101049073ffffffffffffffffffffffffffffffffffffffff1681565b6101416101c73660046107fd565b6103a0565b6101416101da36600461081f565b610419565b6101046101ed366004610841565b610467565b61021261020036600461071e565b60006020819052908152604090205481565b604051908152602001610125565b6004546101049073ffffffffffffffffffffffffffffffffffffffff1681565b61014161024e366004610884565b61059c565b604080516080810182526000606082019081528152336020820152468183015290517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820186905286169063a9059cbb906044016020604051808303816000875af11580156102e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030b91906108b0565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87169063de43156e90610366908490899089908990899060040161091b565b600060405180830381600087803b15801561038057600080fd5b505af1158015610394573d6000803e3d6000fd5b50505050505050505050565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e9060200160405180910390a150565b6000828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b60008060006104768585610620565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b166034820152919350915086906048016040516020818303038152906040528051906020012060405160200161055c9291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b60008281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d910161045b565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610688576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16106106c25782846106c5565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610717576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b60006020828403121561073057600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461075b57600080fd5b919050565b60008060008060006080868803121561077857600080fd5b61078186610737565b945061078f60208701610737565b935060408601359250606086013567ffffffffffffffff8111156107b257600080fd5b8601601f810188136107c357600080fd5b803567ffffffffffffffff8111156107da57600080fd5b8860208284010111156107ec57600080fd5b959894975092955050506020019190565b60006020828403121561080f57600080fd5b61081882610737565b9392505050565b6000806040838503121561083257600080fd5b50508035926020909101359150565b60008060006060848603121561085657600080fd5b61085f84610737565b925061086d60208501610737565b915061087b60408501610737565b90509250925092565b6000806040838503121561089757600080fd5b823591506108a760208401610737565b90509250929050565b6000602082840312156108c257600080fd5b8151801515811461081857600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60808152600086516060608084015280518060e085015260005b81811015610953576020818401810151610100878401015201610935565b5060008482016101000152602089015173ffffffffffffffffffffffffffffffffffffffff811660a0860152601f9091017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168401915050604088015160c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401526101008382030160608401526109f2610100820185876108d2565b9897505050505050505056fea2646970667358221220e2543a307670188761049782bdaa8a44f7c7370f5674a3896ec2c52018aa509b64736f6c634300081a0033",
}

// SystemContractMockABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemContractMockMetaData.ABI instead.
var SystemContractMockABI = SystemContractMockMetaData.ABI

// SystemContractMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemContractMockMetaData.Bin instead.
var SystemContractMockBin = SystemContractMockMetaData.Bin

// DeploySystemContractMock deploys a new Ethereum contract, binding an instance of SystemContractMock to it.
func DeploySystemContractMock(auth *bind.TransactOpts, backend bind.ContractBackend, wzeta_ common.Address, uniswapv2Factory_ common.Address, uniswapv2Router02_ common.Address) (common.Address, *types.Transaction, *SystemContractMock, error) {
	parsed, err := SystemContractMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemContractMockBin), backend, wzeta_, uniswapv2Factory_, uniswapv2Router02_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemContractMock{SystemContractMockCaller: SystemContractMockCaller{contract: contract}, SystemContractMockTransactor: SystemContractMockTransactor{contract: contract}, SystemContractMockFilterer: SystemContractMockFilterer{contract: contract}}, nil
}

// SystemContractMock is an auto generated Go binding around an Ethereum contract.
type SystemContractMock struct {
	SystemContractMockCaller     // Read-only binding to the contract
	SystemContractMockTransactor // Write-only binding to the contract
	SystemContractMockFilterer   // Log filterer for contract events
}

// SystemContractMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemContractMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemContractMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemContractMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemContractMockSession struct {
	Contract     *SystemContractMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SystemContractMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemContractMockCallerSession struct {
	Contract *SystemContractMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SystemContractMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemContractMockTransactorSession struct {
	Contract     *SystemContractMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SystemContractMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemContractMockRaw struct {
	Contract *SystemContractMock // Generic contract binding to access the raw methods on
}

// SystemContractMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemContractMockCallerRaw struct {
	Contract *SystemContractMockCaller // Generic read-only contract binding to access the raw methods on
}

// SystemContractMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemContractMockTransactorRaw struct {
	Contract *SystemContractMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemContractMock creates a new instance of SystemContractMock, bound to a specific deployed contract.
func NewSystemContractMock(address common.Address, backend bind.ContractBackend) (*SystemContractMock, error) {
	contract, err := bindSystemContractMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemContractMock{SystemContractMockCaller: SystemContractMockCaller{contract: contract}, SystemContractMockTransactor: SystemContractMockTransactor{contract: contract}, SystemContractMockFilterer: SystemContractMockFilterer{contract: contract}}, nil
}

// NewSystemContractMockCaller creates a new read-only instance of SystemContractMock, bound to a specific deployed contract.
func NewSystemContractMockCaller(address common.Address, caller bind.ContractCaller) (*SystemContractMockCaller, error) {
	contract, err := bindSystemContractMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemContractMockCaller{contract: contract}, nil
}

// NewSystemContractMockTransactor creates a new write-only instance of SystemContractMock, bound to a specific deployed contract.
func NewSystemContractMockTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemContractMockTransactor, error) {
	contract, err := bindSystemContractMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemContractMockTransactor{contract: contract}, nil
}

// NewSystemContractMockFilterer creates a new log filterer instance of SystemContractMock, bound to a specific deployed contract.
func NewSystemContractMockFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemContractMockFilterer, error) {
	contract, err := bindSystemContractMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemContractMockFilterer{contract: contract}, nil
}

// bindSystemContractMock binds a generic wrapper to an already deployed contract.
func bindSystemContractMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SystemContractMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemContractMock *SystemContractMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemContractMock.Contract.SystemContractMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemContractMock *SystemContractMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SystemContractMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemContractMock *SystemContractMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SystemContractMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemContractMock *SystemContractMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemContractMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemContractMock *SystemContractMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContractMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemContractMock *SystemContractMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemContractMock.Contract.contract.Transact(opts, method, params...)
}

// GasCoinZRC20ByChainId is a free data retrieval call binding the contract method 0x0be15547.
//
// Solidity: function gasCoinZRC20ByChainId(uint256 ) view returns(address)
func (_SystemContractMock *SystemContractMockCaller) GasCoinZRC20ByChainId(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SystemContractMock.contract.Call(opts, &out, "gasCoinZRC20ByChainId", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasCoinZRC20ByChainId is a free data retrieval call binding the contract method 0x0be15547.
//
// Solidity: function gasCoinZRC20ByChainId(uint256 ) view returns(address)
func (_SystemContractMock *SystemContractMockSession) GasCoinZRC20ByChainId(arg0 *big.Int) (common.Address, error) {
	return _SystemContractMock.Contract.GasCoinZRC20ByChainId(&_SystemContractMock.CallOpts, arg0)
}

// GasCoinZRC20ByChainId is a free data retrieval call binding the contract method 0x0be15547.
//
// Solidity: function gasCoinZRC20ByChainId(uint256 ) view returns(address)
func (_SystemContractMock *SystemContractMockCallerSession) GasCoinZRC20ByChainId(arg0 *big.Int) (common.Address, error) {
	return _SystemContractMock.Contract.GasCoinZRC20ByChainId(&_SystemContractMock.CallOpts, arg0)
}

// GasPriceByChainId is a free data retrieval call binding the contract method 0xd7fd7afb.
//
// Solidity: function gasPriceByChainId(uint256 ) view returns(uint256)
func (_SystemContractMock *SystemContractMockCaller) GasPriceByChainId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SystemContractMock.contract.Call(opts, &out, "gasPriceByChainId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasPriceByChainId is a free data retrieval call binding the contract method 0xd7fd7afb.
//
// Solidity: function gasPriceByChainId(uint256 ) view returns(uint256)
func (_SystemContractMock *SystemContractMockSession) GasPriceByChainId(arg0 *big.Int) (*big.Int, error) {
	return _SystemContractMock.Contract.GasPriceByChainId(&_SystemContractMock.CallOpts, arg0)
}

// GasPriceByChainId is a free data retrieval call binding the contract method 0xd7fd7afb.
//
// Solidity: function gasPriceByChainId(uint256 ) view returns(uint256)
func (_SystemContractMock *SystemContractMockCallerSession) GasPriceByChainId(arg0 *big.Int) (*big.Int, error) {
	return _SystemContractMock.Contract.GasPriceByChainId(&_SystemContractMock.CallOpts, arg0)
}

// GasZetaPoolByChainId is a free data retrieval call binding the contract method 0x513a9c05.
//
// Solidity: function gasZetaPoolByChainId(uint256 ) view returns(address)
func (_SystemContractMock *SystemContractMockCaller) GasZetaPoolByChainId(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SystemContractMock.contract.Call(opts, &out, "gasZetaPoolByChainId", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasZetaPoolByChainId is a free data retrieval call binding the contract method 0x513a9c05.
//
// Solidity: function gasZetaPoolByChainId(uint256 ) view returns(address)
func (_SystemContractMock *SystemContractMockSession) GasZetaPoolByChainId(arg0 *big.Int) (common.Address, error) {
	return _SystemContractMock.Contract.GasZetaPoolByChainId(&_SystemContractMock.CallOpts, arg0)
}

// GasZetaPoolByChainId is a free data retrieval call binding the contract method 0x513a9c05.
//
// Solidity: function gasZetaPoolByChainId(uint256 ) view returns(address)
func (_SystemContractMock *SystemContractMockCallerSession) GasZetaPoolByChainId(arg0 *big.Int) (common.Address, error) {
	return _SystemContractMock.Contract.GasZetaPoolByChainId(&_SystemContractMock.CallOpts, arg0)
}

// Uniswapv2FactoryAddress is a free data retrieval call binding the contract method 0xd936a012.
//
// Solidity: function uniswapv2FactoryAddress() view returns(address)
func (_SystemContractMock *SystemContractMockCaller) Uniswapv2FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContractMock.contract.Call(opts, &out, "uniswapv2FactoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uniswapv2FactoryAddress is a free data retrieval call binding the contract method 0xd936a012.
//
// Solidity: function uniswapv2FactoryAddress() view returns(address)
func (_SystemContractMock *SystemContractMockSession) Uniswapv2FactoryAddress() (common.Address, error) {
	return _SystemContractMock.Contract.Uniswapv2FactoryAddress(&_SystemContractMock.CallOpts)
}

// Uniswapv2FactoryAddress is a free data retrieval call binding the contract method 0xd936a012.
//
// Solidity: function uniswapv2FactoryAddress() view returns(address)
func (_SystemContractMock *SystemContractMockCallerSession) Uniswapv2FactoryAddress() (common.Address, error) {
	return _SystemContractMock.Contract.Uniswapv2FactoryAddress(&_SystemContractMock.CallOpts)
}

// Uniswapv2PairFor is a free data retrieval call binding the contract method 0xc63585cc.
//
// Solidity: function uniswapv2PairFor(address factory, address tokenA, address tokenB) pure returns(address pair)
func (_SystemContractMock *SystemContractMockCaller) Uniswapv2PairFor(opts *bind.CallOpts, factory common.Address, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _SystemContractMock.contract.Call(opts, &out, "uniswapv2PairFor", factory, tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uniswapv2PairFor is a free data retrieval call binding the contract method 0xc63585cc.
//
// Solidity: function uniswapv2PairFor(address factory, address tokenA, address tokenB) pure returns(address pair)
func (_SystemContractMock *SystemContractMockSession) Uniswapv2PairFor(factory common.Address, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _SystemContractMock.Contract.Uniswapv2PairFor(&_SystemContractMock.CallOpts, factory, tokenA, tokenB)
}

// Uniswapv2PairFor is a free data retrieval call binding the contract method 0xc63585cc.
//
// Solidity: function uniswapv2PairFor(address factory, address tokenA, address tokenB) pure returns(address pair)
func (_SystemContractMock *SystemContractMockCallerSession) Uniswapv2PairFor(factory common.Address, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _SystemContractMock.Contract.Uniswapv2PairFor(&_SystemContractMock.CallOpts, factory, tokenA, tokenB)
}

// Uniswapv2Router02Address is a free data retrieval call binding the contract method 0x842da36d.
//
// Solidity: function uniswapv2Router02Address() view returns(address)
func (_SystemContractMock *SystemContractMockCaller) Uniswapv2Router02Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContractMock.contract.Call(opts, &out, "uniswapv2Router02Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uniswapv2Router02Address is a free data retrieval call binding the contract method 0x842da36d.
//
// Solidity: function uniswapv2Router02Address() view returns(address)
func (_SystemContractMock *SystemContractMockSession) Uniswapv2Router02Address() (common.Address, error) {
	return _SystemContractMock.Contract.Uniswapv2Router02Address(&_SystemContractMock.CallOpts)
}

// Uniswapv2Router02Address is a free data retrieval call binding the contract method 0x842da36d.
//
// Solidity: function uniswapv2Router02Address() view returns(address)
func (_SystemContractMock *SystemContractMockCallerSession) Uniswapv2Router02Address() (common.Address, error) {
	return _SystemContractMock.Contract.Uniswapv2Router02Address(&_SystemContractMock.CallOpts)
}

// WZetaContractAddress is a free data retrieval call binding the contract method 0x569541b9.
//
// Solidity: function wZetaContractAddress() view returns(address)
func (_SystemContractMock *SystemContractMockCaller) WZetaContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContractMock.contract.Call(opts, &out, "wZetaContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WZetaContractAddress is a free data retrieval call binding the contract method 0x569541b9.
//
// Solidity: function wZetaContractAddress() view returns(address)
func (_SystemContractMock *SystemContractMockSession) WZetaContractAddress() (common.Address, error) {
	return _SystemContractMock.Contract.WZetaContractAddress(&_SystemContractMock.CallOpts)
}

// WZetaContractAddress is a free data retrieval call binding the contract method 0x569541b9.
//
// Solidity: function wZetaContractAddress() view returns(address)
func (_SystemContractMock *SystemContractMockCallerSession) WZetaContractAddress() (common.Address, error) {
	return _SystemContractMock.Contract.WZetaContractAddress(&_SystemContractMock.CallOpts)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0x3c669d55.
//
// Solidity: function onCrossChainCall(address target, address zrc20, uint256 amount, bytes message) returns()
func (_SystemContractMock *SystemContractMockTransactor) OnCrossChainCall(opts *bind.TransactOpts, target common.Address, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _SystemContractMock.contract.Transact(opts, "onCrossChainCall", target, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0x3c669d55.
//
// Solidity: function onCrossChainCall(address target, address zrc20, uint256 amount, bytes message) returns()
func (_SystemContractMock *SystemContractMockSession) OnCrossChainCall(target common.Address, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _SystemContractMock.Contract.OnCrossChainCall(&_SystemContractMock.TransactOpts, target, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0x3c669d55.
//
// Solidity: function onCrossChainCall(address target, address zrc20, uint256 amount, bytes message) returns()
func (_SystemContractMock *SystemContractMockTransactorSession) OnCrossChainCall(target common.Address, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _SystemContractMock.Contract.OnCrossChainCall(&_SystemContractMock.TransactOpts, target, zrc20, amount, message)
}

// SetGasCoinZRC20 is a paid mutator transaction binding the contract method 0xee2815ba.
//
// Solidity: function setGasCoinZRC20(uint256 chainID, address zrc20) returns()
func (_SystemContractMock *SystemContractMockTransactor) SetGasCoinZRC20(opts *bind.TransactOpts, chainID *big.Int, zrc20 common.Address) (*types.Transaction, error) {
	return _SystemContractMock.contract.Transact(opts, "setGasCoinZRC20", chainID, zrc20)
}

// SetGasCoinZRC20 is a paid mutator transaction binding the contract method 0xee2815ba.
//
// Solidity: function setGasCoinZRC20(uint256 chainID, address zrc20) returns()
func (_SystemContractMock *SystemContractMockSession) SetGasCoinZRC20(chainID *big.Int, zrc20 common.Address) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SetGasCoinZRC20(&_SystemContractMock.TransactOpts, chainID, zrc20)
}

// SetGasCoinZRC20 is a paid mutator transaction binding the contract method 0xee2815ba.
//
// Solidity: function setGasCoinZRC20(uint256 chainID, address zrc20) returns()
func (_SystemContractMock *SystemContractMockTransactorSession) SetGasCoinZRC20(chainID *big.Int, zrc20 common.Address) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SetGasCoinZRC20(&_SystemContractMock.TransactOpts, chainID, zrc20)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xa7cb0507.
//
// Solidity: function setGasPrice(uint256 chainID, uint256 price) returns()
func (_SystemContractMock *SystemContractMockTransactor) SetGasPrice(opts *bind.TransactOpts, chainID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _SystemContractMock.contract.Transact(opts, "setGasPrice", chainID, price)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xa7cb0507.
//
// Solidity: function setGasPrice(uint256 chainID, uint256 price) returns()
func (_SystemContractMock *SystemContractMockSession) SetGasPrice(chainID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SetGasPrice(&_SystemContractMock.TransactOpts, chainID, price)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xa7cb0507.
//
// Solidity: function setGasPrice(uint256 chainID, uint256 price) returns()
func (_SystemContractMock *SystemContractMockTransactorSession) SetGasPrice(chainID *big.Int, price *big.Int) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SetGasPrice(&_SystemContractMock.TransactOpts, chainID, price)
}

// SetWZETAContractAddress is a paid mutator transaction binding the contract method 0x97770dff.
//
// Solidity: function setWZETAContractAddress(address addr) returns()
func (_SystemContractMock *SystemContractMockTransactor) SetWZETAContractAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _SystemContractMock.contract.Transact(opts, "setWZETAContractAddress", addr)
}

// SetWZETAContractAddress is a paid mutator transaction binding the contract method 0x97770dff.
//
// Solidity: function setWZETAContractAddress(address addr) returns()
func (_SystemContractMock *SystemContractMockSession) SetWZETAContractAddress(addr common.Address) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SetWZETAContractAddress(&_SystemContractMock.TransactOpts, addr)
}

// SetWZETAContractAddress is a paid mutator transaction binding the contract method 0x97770dff.
//
// Solidity: function setWZETAContractAddress(address addr) returns()
func (_SystemContractMock *SystemContractMockTransactorSession) SetWZETAContractAddress(addr common.Address) (*types.Transaction, error) {
	return _SystemContractMock.Contract.SetWZETAContractAddress(&_SystemContractMock.TransactOpts, addr)
}

// SystemContractMockSetGasCoinIterator is returned from FilterSetGasCoin and is used to iterate over the raw logs and unpacked data for SetGasCoin events raised by the SystemContractMock contract.
type SystemContractMockSetGasCoinIterator struct {
	Event *SystemContractMockSetGasCoin // Event containing the contract specifics and raw log

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
func (it *SystemContractMockSetGasCoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractMockSetGasCoin)
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
		it.Event = new(SystemContractMockSetGasCoin)
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
func (it *SystemContractMockSetGasCoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractMockSetGasCoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractMockSetGasCoin represents a SetGasCoin event raised by the SystemContractMock contract.
type SystemContractMockSetGasCoin struct {
	Arg0 *big.Int
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetGasCoin is a free log retrieval operation binding the contract event 0xd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d.
//
// Solidity: event SetGasCoin(uint256 arg0, address arg1)
func (_SystemContractMock *SystemContractMockFilterer) FilterSetGasCoin(opts *bind.FilterOpts) (*SystemContractMockSetGasCoinIterator, error) {

	logs, sub, err := _SystemContractMock.contract.FilterLogs(opts, "SetGasCoin")
	if err != nil {
		return nil, err
	}
	return &SystemContractMockSetGasCoinIterator{contract: _SystemContractMock.contract, event: "SetGasCoin", logs: logs, sub: sub}, nil
}

// WatchSetGasCoin is a free log subscription operation binding the contract event 0xd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d.
//
// Solidity: event SetGasCoin(uint256 arg0, address arg1)
func (_SystemContractMock *SystemContractMockFilterer) WatchSetGasCoin(opts *bind.WatchOpts, sink chan<- *SystemContractMockSetGasCoin) (event.Subscription, error) {

	logs, sub, err := _SystemContractMock.contract.WatchLogs(opts, "SetGasCoin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractMockSetGasCoin)
				if err := _SystemContractMock.contract.UnpackLog(event, "SetGasCoin", log); err != nil {
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

// ParseSetGasCoin is a log parse operation binding the contract event 0xd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d.
//
// Solidity: event SetGasCoin(uint256 arg0, address arg1)
func (_SystemContractMock *SystemContractMockFilterer) ParseSetGasCoin(log types.Log) (*SystemContractMockSetGasCoin, error) {
	event := new(SystemContractMockSetGasCoin)
	if err := _SystemContractMock.contract.UnpackLog(event, "SetGasCoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractMockSetGasPriceIterator is returned from FilterSetGasPrice and is used to iterate over the raw logs and unpacked data for SetGasPrice events raised by the SystemContractMock contract.
type SystemContractMockSetGasPriceIterator struct {
	Event *SystemContractMockSetGasPrice // Event containing the contract specifics and raw log

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
func (it *SystemContractMockSetGasPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractMockSetGasPrice)
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
		it.Event = new(SystemContractMockSetGasPrice)
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
func (it *SystemContractMockSetGasPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractMockSetGasPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractMockSetGasPrice represents a SetGasPrice event raised by the SystemContractMock contract.
type SystemContractMockSetGasPrice struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetGasPrice is a free log retrieval operation binding the contract event 0x49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d.
//
// Solidity: event SetGasPrice(uint256 arg0, uint256 arg1)
func (_SystemContractMock *SystemContractMockFilterer) FilterSetGasPrice(opts *bind.FilterOpts) (*SystemContractMockSetGasPriceIterator, error) {

	logs, sub, err := _SystemContractMock.contract.FilterLogs(opts, "SetGasPrice")
	if err != nil {
		return nil, err
	}
	return &SystemContractMockSetGasPriceIterator{contract: _SystemContractMock.contract, event: "SetGasPrice", logs: logs, sub: sub}, nil
}

// WatchSetGasPrice is a free log subscription operation binding the contract event 0x49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d.
//
// Solidity: event SetGasPrice(uint256 arg0, uint256 arg1)
func (_SystemContractMock *SystemContractMockFilterer) WatchSetGasPrice(opts *bind.WatchOpts, sink chan<- *SystemContractMockSetGasPrice) (event.Subscription, error) {

	logs, sub, err := _SystemContractMock.contract.WatchLogs(opts, "SetGasPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractMockSetGasPrice)
				if err := _SystemContractMock.contract.UnpackLog(event, "SetGasPrice", log); err != nil {
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

// ParseSetGasPrice is a log parse operation binding the contract event 0x49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d.
//
// Solidity: event SetGasPrice(uint256 arg0, uint256 arg1)
func (_SystemContractMock *SystemContractMockFilterer) ParseSetGasPrice(log types.Log) (*SystemContractMockSetGasPrice, error) {
	event := new(SystemContractMockSetGasPrice)
	if err := _SystemContractMock.contract.UnpackLog(event, "SetGasPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractMockSetGasZetaPoolIterator is returned from FilterSetGasZetaPool and is used to iterate over the raw logs and unpacked data for SetGasZetaPool events raised by the SystemContractMock contract.
type SystemContractMockSetGasZetaPoolIterator struct {
	Event *SystemContractMockSetGasZetaPool // Event containing the contract specifics and raw log

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
func (it *SystemContractMockSetGasZetaPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractMockSetGasZetaPool)
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
		it.Event = new(SystemContractMockSetGasZetaPool)
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
func (it *SystemContractMockSetGasZetaPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractMockSetGasZetaPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractMockSetGasZetaPool represents a SetGasZetaPool event raised by the SystemContractMock contract.
type SystemContractMockSetGasZetaPool struct {
	Arg0 *big.Int
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetGasZetaPool is a free log retrieval operation binding the contract event 0x0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e.
//
// Solidity: event SetGasZetaPool(uint256 arg0, address arg1)
func (_SystemContractMock *SystemContractMockFilterer) FilterSetGasZetaPool(opts *bind.FilterOpts) (*SystemContractMockSetGasZetaPoolIterator, error) {

	logs, sub, err := _SystemContractMock.contract.FilterLogs(opts, "SetGasZetaPool")
	if err != nil {
		return nil, err
	}
	return &SystemContractMockSetGasZetaPoolIterator{contract: _SystemContractMock.contract, event: "SetGasZetaPool", logs: logs, sub: sub}, nil
}

// WatchSetGasZetaPool is a free log subscription operation binding the contract event 0x0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e.
//
// Solidity: event SetGasZetaPool(uint256 arg0, address arg1)
func (_SystemContractMock *SystemContractMockFilterer) WatchSetGasZetaPool(opts *bind.WatchOpts, sink chan<- *SystemContractMockSetGasZetaPool) (event.Subscription, error) {

	logs, sub, err := _SystemContractMock.contract.WatchLogs(opts, "SetGasZetaPool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractMockSetGasZetaPool)
				if err := _SystemContractMock.contract.UnpackLog(event, "SetGasZetaPool", log); err != nil {
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

// ParseSetGasZetaPool is a log parse operation binding the contract event 0x0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e.
//
// Solidity: event SetGasZetaPool(uint256 arg0, address arg1)
func (_SystemContractMock *SystemContractMockFilterer) ParseSetGasZetaPool(log types.Log) (*SystemContractMockSetGasZetaPool, error) {
	event := new(SystemContractMockSetGasZetaPool)
	if err := _SystemContractMock.contract.UnpackLog(event, "SetGasZetaPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractMockSetWZetaIterator is returned from FilterSetWZeta and is used to iterate over the raw logs and unpacked data for SetWZeta events raised by the SystemContractMock contract.
type SystemContractMockSetWZetaIterator struct {
	Event *SystemContractMockSetWZeta // Event containing the contract specifics and raw log

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
func (it *SystemContractMockSetWZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractMockSetWZeta)
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
		it.Event = new(SystemContractMockSetWZeta)
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
func (it *SystemContractMockSetWZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractMockSetWZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractMockSetWZeta represents a SetWZeta event raised by the SystemContractMock contract.
type SystemContractMockSetWZeta struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetWZeta is a free log retrieval operation binding the contract event 0xdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e.
//
// Solidity: event SetWZeta(address arg0)
func (_SystemContractMock *SystemContractMockFilterer) FilterSetWZeta(opts *bind.FilterOpts) (*SystemContractMockSetWZetaIterator, error) {

	logs, sub, err := _SystemContractMock.contract.FilterLogs(opts, "SetWZeta")
	if err != nil {
		return nil, err
	}
	return &SystemContractMockSetWZetaIterator{contract: _SystemContractMock.contract, event: "SetWZeta", logs: logs, sub: sub}, nil
}

// WatchSetWZeta is a free log subscription operation binding the contract event 0xdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e.
//
// Solidity: event SetWZeta(address arg0)
func (_SystemContractMock *SystemContractMockFilterer) WatchSetWZeta(opts *bind.WatchOpts, sink chan<- *SystemContractMockSetWZeta) (event.Subscription, error) {

	logs, sub, err := _SystemContractMock.contract.WatchLogs(opts, "SetWZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractMockSetWZeta)
				if err := _SystemContractMock.contract.UnpackLog(event, "SetWZeta", log); err != nil {
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

// ParseSetWZeta is a log parse operation binding the contract event 0xdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e.
//
// Solidity: event SetWZeta(address arg0)
func (_SystemContractMock *SystemContractMockFilterer) ParseSetWZeta(log types.Log) (*SystemContractMockSetWZeta, error) {
	event := new(SystemContractMockSetWZeta)
	if err := _SystemContractMock.contract.UnpackLog(event, "SetWZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractMockSystemContractDeployedIterator is returned from FilterSystemContractDeployed and is used to iterate over the raw logs and unpacked data for SystemContractDeployed events raised by the SystemContractMock contract.
type SystemContractMockSystemContractDeployedIterator struct {
	Event *SystemContractMockSystemContractDeployed // Event containing the contract specifics and raw log

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
func (it *SystemContractMockSystemContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractMockSystemContractDeployed)
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
		it.Event = new(SystemContractMockSystemContractDeployed)
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
func (it *SystemContractMockSystemContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractMockSystemContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractMockSystemContractDeployed represents a SystemContractDeployed event raised by the SystemContractMock contract.
type SystemContractMockSystemContractDeployed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSystemContractDeployed is a free log retrieval operation binding the contract event 0x80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac5.
//
// Solidity: event SystemContractDeployed()
func (_SystemContractMock *SystemContractMockFilterer) FilterSystemContractDeployed(opts *bind.FilterOpts) (*SystemContractMockSystemContractDeployedIterator, error) {

	logs, sub, err := _SystemContractMock.contract.FilterLogs(opts, "SystemContractDeployed")
	if err != nil {
		return nil, err
	}
	return &SystemContractMockSystemContractDeployedIterator{contract: _SystemContractMock.contract, event: "SystemContractDeployed", logs: logs, sub: sub}, nil
}

// WatchSystemContractDeployed is a free log subscription operation binding the contract event 0x80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac5.
//
// Solidity: event SystemContractDeployed()
func (_SystemContractMock *SystemContractMockFilterer) WatchSystemContractDeployed(opts *bind.WatchOpts, sink chan<- *SystemContractMockSystemContractDeployed) (event.Subscription, error) {

	logs, sub, err := _SystemContractMock.contract.WatchLogs(opts, "SystemContractDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractMockSystemContractDeployed)
				if err := _SystemContractMock.contract.UnpackLog(event, "SystemContractDeployed", log); err != nil {
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

// ParseSystemContractDeployed is a log parse operation binding the contract event 0x80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac5.
//
// Solidity: event SystemContractDeployed()
func (_SystemContractMock *SystemContractMockFilterer) ParseSystemContractDeployed(log types.Log) (*SystemContractMockSystemContractDeployed, error) {
	event := new(SystemContractMockSystemContractDeployed)
	if err := _SystemContractMock.contract.UnpackLog(event, "SystemContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
