// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package baseforktest

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

// StdInvariantFuzzArtifactSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzArtifactSelector struct {
	Artifact  string
	Selectors [][4]byte
}

// StdInvariantFuzzInterface is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzInterface struct {
	Addr      common.Address
	Artifacts []string
}

// StdInvariantFuzzSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzSelector struct {
	Addr      common.Address
	Selectors [][4]byte
}

// BaseForkTestMetaData contains all meta data concerning the BaseForkTest contract.
var BaseForkTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ARBITRUM_RPC_URL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"AVALANCHE_RPC_URL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BASE_RPC_URL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BSC_RPC_URL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ETHEREUM_RPC_URL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POLYGON_RPC_URL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ZETACHAIN_RPC_URL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chains\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"forkId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rpcUrl\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testCanSwitchForks\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testForkIdDiffer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testUpgradeGatewayOnAllChains\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
}

// BaseForkTestABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseForkTestMetaData.ABI instead.
var BaseForkTestABI = BaseForkTestMetaData.ABI

// BaseForkTest is an auto generated Go binding around an Ethereum contract.
type BaseForkTest struct {
	BaseForkTestCaller     // Read-only binding to the contract
	BaseForkTestTransactor // Write-only binding to the contract
	BaseForkTestFilterer   // Log filterer for contract events
}

// BaseForkTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseForkTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseForkTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseForkTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseForkTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseForkTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseForkTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseForkTestSession struct {
	Contract     *BaseForkTest     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseForkTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseForkTestCallerSession struct {
	Contract *BaseForkTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BaseForkTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseForkTestTransactorSession struct {
	Contract     *BaseForkTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BaseForkTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseForkTestRaw struct {
	Contract *BaseForkTest // Generic contract binding to access the raw methods on
}

// BaseForkTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseForkTestCallerRaw struct {
	Contract *BaseForkTestCaller // Generic read-only contract binding to access the raw methods on
}

// BaseForkTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseForkTestTransactorRaw struct {
	Contract *BaseForkTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseForkTest creates a new instance of BaseForkTest, bound to a specific deployed contract.
func NewBaseForkTest(address common.Address, backend bind.ContractBackend) (*BaseForkTest, error) {
	contract, err := bindBaseForkTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseForkTest{BaseForkTestCaller: BaseForkTestCaller{contract: contract}, BaseForkTestTransactor: BaseForkTestTransactor{contract: contract}, BaseForkTestFilterer: BaseForkTestFilterer{contract: contract}}, nil
}

// NewBaseForkTestCaller creates a new read-only instance of BaseForkTest, bound to a specific deployed contract.
func NewBaseForkTestCaller(address common.Address, caller bind.ContractCaller) (*BaseForkTestCaller, error) {
	contract, err := bindBaseForkTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseForkTestCaller{contract: contract}, nil
}

// NewBaseForkTestTransactor creates a new write-only instance of BaseForkTest, bound to a specific deployed contract.
func NewBaseForkTestTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseForkTestTransactor, error) {
	contract, err := bindBaseForkTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseForkTestTransactor{contract: contract}, nil
}

// NewBaseForkTestFilterer creates a new log filterer instance of BaseForkTest, bound to a specific deployed contract.
func NewBaseForkTestFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseForkTestFilterer, error) {
	contract, err := bindBaseForkTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseForkTestFilterer{contract: contract}, nil
}

// bindBaseForkTest binds a generic wrapper to an already deployed contract.
func bindBaseForkTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseForkTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseForkTest *BaseForkTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseForkTest.Contract.BaseForkTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseForkTest *BaseForkTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseForkTest.Contract.BaseForkTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseForkTest *BaseForkTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseForkTest.Contract.BaseForkTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseForkTest *BaseForkTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseForkTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseForkTest *BaseForkTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseForkTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseForkTest *BaseForkTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseForkTest.Contract.contract.Transact(opts, method, params...)
}

// ARBITRUMRPCURL is a free data retrieval call binding the contract method 0x699370b0.
//
// Solidity: function ARBITRUM_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestCaller) ARBITRUMRPCURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "ARBITRUM_RPC_URL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ARBITRUMRPCURL is a free data retrieval call binding the contract method 0x699370b0.
//
// Solidity: function ARBITRUM_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestSession) ARBITRUMRPCURL() (string, error) {
	return _BaseForkTest.Contract.ARBITRUMRPCURL(&_BaseForkTest.CallOpts)
}

// ARBITRUMRPCURL is a free data retrieval call binding the contract method 0x699370b0.
//
// Solidity: function ARBITRUM_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestCallerSession) ARBITRUMRPCURL() (string, error) {
	return _BaseForkTest.Contract.ARBITRUMRPCURL(&_BaseForkTest.CallOpts)
}

// AVALANCHERPCURL is a free data retrieval call binding the contract method 0xca747493.
//
// Solidity: function AVALANCHE_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestCaller) AVALANCHERPCURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "AVALANCHE_RPC_URL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AVALANCHERPCURL is a free data retrieval call binding the contract method 0xca747493.
//
// Solidity: function AVALANCHE_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestSession) AVALANCHERPCURL() (string, error) {
	return _BaseForkTest.Contract.AVALANCHERPCURL(&_BaseForkTest.CallOpts)
}

// AVALANCHERPCURL is a free data retrieval call binding the contract method 0xca747493.
//
// Solidity: function AVALANCHE_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestCallerSession) AVALANCHERPCURL() (string, error) {
	return _BaseForkTest.Contract.AVALANCHERPCURL(&_BaseForkTest.CallOpts)
}

// BASERPCURL is a free data retrieval call binding the contract method 0x2e8ade6e.
//
// Solidity: function BASE_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestCaller) BASERPCURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "BASE_RPC_URL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BASERPCURL is a free data retrieval call binding the contract method 0x2e8ade6e.
//
// Solidity: function BASE_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestSession) BASERPCURL() (string, error) {
	return _BaseForkTest.Contract.BASERPCURL(&_BaseForkTest.CallOpts)
}

// BASERPCURL is a free data retrieval call binding the contract method 0x2e8ade6e.
//
// Solidity: function BASE_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestCallerSession) BASERPCURL() (string, error) {
	return _BaseForkTest.Contract.BASERPCURL(&_BaseForkTest.CallOpts)
}

// BSCRPCURL is a free data retrieval call binding the contract method 0x88de7b87.
//
// Solidity: function BSC_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestCaller) BSCRPCURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "BSC_RPC_URL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BSCRPCURL is a free data retrieval call binding the contract method 0x88de7b87.
//
// Solidity: function BSC_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestSession) BSCRPCURL() (string, error) {
	return _BaseForkTest.Contract.BSCRPCURL(&_BaseForkTest.CallOpts)
}

// BSCRPCURL is a free data retrieval call binding the contract method 0x88de7b87.
//
// Solidity: function BSC_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestCallerSession) BSCRPCURL() (string, error) {
	return _BaseForkTest.Contract.BSCRPCURL(&_BaseForkTest.CallOpts)
}

// ETHEREUMRPCURL is a free data retrieval call binding the contract method 0x9d3c5161.
//
// Solidity: function ETHEREUM_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestCaller) ETHEREUMRPCURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "ETHEREUM_RPC_URL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ETHEREUMRPCURL is a free data retrieval call binding the contract method 0x9d3c5161.
//
// Solidity: function ETHEREUM_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestSession) ETHEREUMRPCURL() (string, error) {
	return _BaseForkTest.Contract.ETHEREUMRPCURL(&_BaseForkTest.CallOpts)
}

// ETHEREUMRPCURL is a free data retrieval call binding the contract method 0x9d3c5161.
//
// Solidity: function ETHEREUM_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestCallerSession) ETHEREUMRPCURL() (string, error) {
	return _BaseForkTest.Contract.ETHEREUMRPCURL(&_BaseForkTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_BaseForkTest *BaseForkTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_BaseForkTest *BaseForkTestSession) ISTEST() (bool, error) {
	return _BaseForkTest.Contract.ISTEST(&_BaseForkTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_BaseForkTest *BaseForkTestCallerSession) ISTEST() (bool, error) {
	return _BaseForkTest.Contract.ISTEST(&_BaseForkTest.CallOpts)
}

// POLYGONRPCURL is a free data retrieval call binding the contract method 0x94897fb1.
//
// Solidity: function POLYGON_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestCaller) POLYGONRPCURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "POLYGON_RPC_URL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// POLYGONRPCURL is a free data retrieval call binding the contract method 0x94897fb1.
//
// Solidity: function POLYGON_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestSession) POLYGONRPCURL() (string, error) {
	return _BaseForkTest.Contract.POLYGONRPCURL(&_BaseForkTest.CallOpts)
}

// POLYGONRPCURL is a free data retrieval call binding the contract method 0x94897fb1.
//
// Solidity: function POLYGON_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestCallerSession) POLYGONRPCURL() (string, error) {
	return _BaseForkTest.Contract.POLYGONRPCURL(&_BaseForkTest.CallOpts)
}

// ZETACHAINRPCURL is a free data retrieval call binding the contract method 0x3f8f61dd.
//
// Solidity: function ZETACHAIN_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestCaller) ZETACHAINRPCURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "ZETACHAIN_RPC_URL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ZETACHAINRPCURL is a free data retrieval call binding the contract method 0x3f8f61dd.
//
// Solidity: function ZETACHAIN_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestSession) ZETACHAINRPCURL() (string, error) {
	return _BaseForkTest.Contract.ZETACHAINRPCURL(&_BaseForkTest.CallOpts)
}

// ZETACHAINRPCURL is a free data retrieval call binding the contract method 0x3f8f61dd.
//
// Solidity: function ZETACHAIN_RPC_URL() view returns(string)
func (_BaseForkTest *BaseForkTestCallerSession) ZETACHAINRPCURL() (string, error) {
	return _BaseForkTest.Contract.ZETACHAINRPCURL(&_BaseForkTest.CallOpts)
}

// Chains is a free data retrieval call binding the contract method 0x550325b5.
//
// Solidity: function chains(uint256 ) view returns(uint256 forkId, address contractAddress, address admin, string rpcUrl, string name)
func (_BaseForkTest *BaseForkTestCaller) Chains(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ForkId          *big.Int
	ContractAddress common.Address
	Admin           common.Address
	RpcUrl          string
	Name            string
}, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "chains", arg0)

	outstruct := new(struct {
		ForkId          *big.Int
		ContractAddress common.Address
		Admin           common.Address
		RpcUrl          string
		Name            string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ForkId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ContractAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Admin = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.RpcUrl = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// Chains is a free data retrieval call binding the contract method 0x550325b5.
//
// Solidity: function chains(uint256 ) view returns(uint256 forkId, address contractAddress, address admin, string rpcUrl, string name)
func (_BaseForkTest *BaseForkTestSession) Chains(arg0 *big.Int) (struct {
	ForkId          *big.Int
	ContractAddress common.Address
	Admin           common.Address
	RpcUrl          string
	Name            string
}, error) {
	return _BaseForkTest.Contract.Chains(&_BaseForkTest.CallOpts, arg0)
}

// Chains is a free data retrieval call binding the contract method 0x550325b5.
//
// Solidity: function chains(uint256 ) view returns(uint256 forkId, address contractAddress, address admin, string rpcUrl, string name)
func (_BaseForkTest *BaseForkTestCallerSession) Chains(arg0 *big.Int) (struct {
	ForkId          *big.Int
	ContractAddress common.Address
	Admin           common.Address
	RpcUrl          string
	Name            string
}, error) {
	return _BaseForkTest.Contract.Chains(&_BaseForkTest.CallOpts, arg0)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_BaseForkTest *BaseForkTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_BaseForkTest *BaseForkTestSession) ExcludeArtifacts() ([]string, error) {
	return _BaseForkTest.Contract.ExcludeArtifacts(&_BaseForkTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_BaseForkTest *BaseForkTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _BaseForkTest.Contract.ExcludeArtifacts(&_BaseForkTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_BaseForkTest *BaseForkTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_BaseForkTest *BaseForkTestSession) ExcludeContracts() ([]common.Address, error) {
	return _BaseForkTest.Contract.ExcludeContracts(&_BaseForkTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_BaseForkTest *BaseForkTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _BaseForkTest.Contract.ExcludeContracts(&_BaseForkTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_BaseForkTest *BaseForkTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_BaseForkTest *BaseForkTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _BaseForkTest.Contract.ExcludeSelectors(&_BaseForkTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_BaseForkTest *BaseForkTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _BaseForkTest.Contract.ExcludeSelectors(&_BaseForkTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_BaseForkTest *BaseForkTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_BaseForkTest *BaseForkTestSession) ExcludeSenders() ([]common.Address, error) {
	return _BaseForkTest.Contract.ExcludeSenders(&_BaseForkTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_BaseForkTest *BaseForkTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _BaseForkTest.Contract.ExcludeSenders(&_BaseForkTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_BaseForkTest *BaseForkTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_BaseForkTest *BaseForkTestSession) Failed() (bool, error) {
	return _BaseForkTest.Contract.Failed(&_BaseForkTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_BaseForkTest *BaseForkTestCallerSession) Failed() (bool, error) {
	return _BaseForkTest.Contract.Failed(&_BaseForkTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_BaseForkTest *BaseForkTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_BaseForkTest *BaseForkTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _BaseForkTest.Contract.TargetArtifactSelectors(&_BaseForkTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_BaseForkTest *BaseForkTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _BaseForkTest.Contract.TargetArtifactSelectors(&_BaseForkTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_BaseForkTest *BaseForkTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_BaseForkTest *BaseForkTestSession) TargetArtifacts() ([]string, error) {
	return _BaseForkTest.Contract.TargetArtifacts(&_BaseForkTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_BaseForkTest *BaseForkTestCallerSession) TargetArtifacts() ([]string, error) {
	return _BaseForkTest.Contract.TargetArtifacts(&_BaseForkTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_BaseForkTest *BaseForkTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_BaseForkTest *BaseForkTestSession) TargetContracts() ([]common.Address, error) {
	return _BaseForkTest.Contract.TargetContracts(&_BaseForkTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_BaseForkTest *BaseForkTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _BaseForkTest.Contract.TargetContracts(&_BaseForkTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_BaseForkTest *BaseForkTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_BaseForkTest *BaseForkTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _BaseForkTest.Contract.TargetInterfaces(&_BaseForkTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_BaseForkTest *BaseForkTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _BaseForkTest.Contract.TargetInterfaces(&_BaseForkTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_BaseForkTest *BaseForkTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_BaseForkTest *BaseForkTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _BaseForkTest.Contract.TargetSelectors(&_BaseForkTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_BaseForkTest *BaseForkTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _BaseForkTest.Contract.TargetSelectors(&_BaseForkTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_BaseForkTest *BaseForkTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_BaseForkTest *BaseForkTestSession) TargetSenders() ([]common.Address, error) {
	return _BaseForkTest.Contract.TargetSenders(&_BaseForkTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_BaseForkTest *BaseForkTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _BaseForkTest.Contract.TargetSenders(&_BaseForkTest.CallOpts)
}

// TestForkIdDiffer is a free data retrieval call binding the contract method 0x6cf73eea.
//
// Solidity: function testForkIdDiffer() view returns()
func (_BaseForkTest *BaseForkTestCaller) TestForkIdDiffer(opts *bind.CallOpts) error {
	var out []interface{}
	err := _BaseForkTest.contract.Call(opts, &out, "testForkIdDiffer")

	if err != nil {
		return err
	}

	return err

}

// TestForkIdDiffer is a free data retrieval call binding the contract method 0x6cf73eea.
//
// Solidity: function testForkIdDiffer() view returns()
func (_BaseForkTest *BaseForkTestSession) TestForkIdDiffer() error {
	return _BaseForkTest.Contract.TestForkIdDiffer(&_BaseForkTest.CallOpts)
}

// TestForkIdDiffer is a free data retrieval call binding the contract method 0x6cf73eea.
//
// Solidity: function testForkIdDiffer() view returns()
func (_BaseForkTest *BaseForkTestCallerSession) TestForkIdDiffer() error {
	return _BaseForkTest.Contract.TestForkIdDiffer(&_BaseForkTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_BaseForkTest *BaseForkTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseForkTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_BaseForkTest *BaseForkTestSession) SetUp() (*types.Transaction, error) {
	return _BaseForkTest.Contract.SetUp(&_BaseForkTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_BaseForkTest *BaseForkTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _BaseForkTest.Contract.SetUp(&_BaseForkTest.TransactOpts)
}

// TestCanSwitchForks is a paid mutator transaction binding the contract method 0x8cf03dbc.
//
// Solidity: function testCanSwitchForks() returns()
func (_BaseForkTest *BaseForkTestTransactor) TestCanSwitchForks(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseForkTest.contract.Transact(opts, "testCanSwitchForks")
}

// TestCanSwitchForks is a paid mutator transaction binding the contract method 0x8cf03dbc.
//
// Solidity: function testCanSwitchForks() returns()
func (_BaseForkTest *BaseForkTestSession) TestCanSwitchForks() (*types.Transaction, error) {
	return _BaseForkTest.Contract.TestCanSwitchForks(&_BaseForkTest.TransactOpts)
}

// TestCanSwitchForks is a paid mutator transaction binding the contract method 0x8cf03dbc.
//
// Solidity: function testCanSwitchForks() returns()
func (_BaseForkTest *BaseForkTestTransactorSession) TestCanSwitchForks() (*types.Transaction, error) {
	return _BaseForkTest.Contract.TestCanSwitchForks(&_BaseForkTest.TransactOpts)
}

// TestUpgradeGatewayOnAllChains is a paid mutator transaction binding the contract method 0x03b965ec.
//
// Solidity: function testUpgradeGatewayOnAllChains() returns()
func (_BaseForkTest *BaseForkTestTransactor) TestUpgradeGatewayOnAllChains(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseForkTest.contract.Transact(opts, "testUpgradeGatewayOnAllChains")
}

// TestUpgradeGatewayOnAllChains is a paid mutator transaction binding the contract method 0x03b965ec.
//
// Solidity: function testUpgradeGatewayOnAllChains() returns()
func (_BaseForkTest *BaseForkTestSession) TestUpgradeGatewayOnAllChains() (*types.Transaction, error) {
	return _BaseForkTest.Contract.TestUpgradeGatewayOnAllChains(&_BaseForkTest.TransactOpts)
}

// TestUpgradeGatewayOnAllChains is a paid mutator transaction binding the contract method 0x03b965ec.
//
// Solidity: function testUpgradeGatewayOnAllChains() returns()
func (_BaseForkTest *BaseForkTestTransactorSession) TestUpgradeGatewayOnAllChains() (*types.Transaction, error) {
	return _BaseForkTest.Contract.TestUpgradeGatewayOnAllChains(&_BaseForkTest.TransactOpts)
}

// BaseForkTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the BaseForkTest contract.
type BaseForkTestLogIterator struct {
	Event *BaseForkTestLog // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLog)
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
		it.Event = new(BaseForkTestLog)
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
func (it *BaseForkTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLog represents a Log event raised by the BaseForkTest contract.
type BaseForkTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_BaseForkTest *BaseForkTestFilterer) FilterLog(opts *bind.FilterOpts) (*BaseForkTestLogIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogIterator{contract: _BaseForkTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_BaseForkTest *BaseForkTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *BaseForkTestLog) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLog)
				if err := _BaseForkTest.contract.UnpackLog(event, "log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_BaseForkTest *BaseForkTestFilterer) ParseLog(log types.Log) (*BaseForkTestLog, error) {
	event := new(BaseForkTestLog)
	if err := _BaseForkTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the BaseForkTest contract.
type BaseForkTestLogAddressIterator struct {
	Event *BaseForkTestLogAddress // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogAddress)
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
		it.Event = new(BaseForkTestLogAddress)
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
func (it *BaseForkTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogAddress represents a LogAddress event raised by the BaseForkTest contract.
type BaseForkTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*BaseForkTestLogAddressIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogAddressIterator{contract: _BaseForkTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogAddress)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_address", log); err != nil {
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

// ParseLogAddress is a log parse operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogAddress(log types.Log) (*BaseForkTestLogAddress, error) {
	event := new(BaseForkTestLogAddress)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the BaseForkTest contract.
type BaseForkTestLogArrayIterator struct {
	Event *BaseForkTestLogArray // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogArray)
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
		it.Event = new(BaseForkTestLogArray)
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
func (it *BaseForkTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogArray represents a LogArray event raised by the BaseForkTest contract.
type BaseForkTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*BaseForkTestLogArrayIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogArrayIterator{contract: _BaseForkTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogArray) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogArray)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_array", log); err != nil {
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

// ParseLogArray is a log parse operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogArray(log types.Log) (*BaseForkTestLogArray, error) {
	event := new(BaseForkTestLogArray)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the BaseForkTest contract.
type BaseForkTestLogArray0Iterator struct {
	Event *BaseForkTestLogArray0 // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogArray0)
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
		it.Event = new(BaseForkTestLogArray0)
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
func (it *BaseForkTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogArray0 represents a LogArray0 event raised by the BaseForkTest contract.
type BaseForkTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*BaseForkTestLogArray0Iterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogArray0Iterator{contract: _BaseForkTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogArray0)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_array0", log); err != nil {
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

// ParseLogArray0 is a log parse operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogArray0(log types.Log) (*BaseForkTestLogArray0, error) {
	event := new(BaseForkTestLogArray0)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the BaseForkTest contract.
type BaseForkTestLogArray1Iterator struct {
	Event *BaseForkTestLogArray1 // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogArray1)
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
		it.Event = new(BaseForkTestLogArray1)
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
func (it *BaseForkTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogArray1 represents a LogArray1 event raised by the BaseForkTest contract.
type BaseForkTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*BaseForkTestLogArray1Iterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogArray1Iterator{contract: _BaseForkTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogArray1)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_array1", log); err != nil {
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

// ParseLogArray1 is a log parse operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogArray1(log types.Log) (*BaseForkTestLogArray1, error) {
	event := new(BaseForkTestLogArray1)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the BaseForkTest contract.
type BaseForkTestLogBytesIterator struct {
	Event *BaseForkTestLogBytes // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogBytes)
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
		it.Event = new(BaseForkTestLogBytes)
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
func (it *BaseForkTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogBytes represents a LogBytes event raised by the BaseForkTest contract.
type BaseForkTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*BaseForkTestLogBytesIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogBytesIterator{contract: _BaseForkTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogBytes)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
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

// ParseLogBytes is a log parse operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogBytes(log types.Log) (*BaseForkTestLogBytes, error) {
	event := new(BaseForkTestLogBytes)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the BaseForkTest contract.
type BaseForkTestLogBytes32Iterator struct {
	Event *BaseForkTestLogBytes32 // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogBytes32)
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
		it.Event = new(BaseForkTestLogBytes32)
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
func (it *BaseForkTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogBytes32 represents a LogBytes32 event raised by the BaseForkTest contract.
type BaseForkTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*BaseForkTestLogBytes32Iterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogBytes32Iterator{contract: _BaseForkTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogBytes32)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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

// ParseLogBytes32 is a log parse operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogBytes32(log types.Log) (*BaseForkTestLogBytes32, error) {
	event := new(BaseForkTestLogBytes32)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the BaseForkTest contract.
type BaseForkTestLogIntIterator struct {
	Event *BaseForkTestLogInt // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogInt)
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
		it.Event = new(BaseForkTestLogInt)
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
func (it *BaseForkTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogInt represents a LogInt event raised by the BaseForkTest contract.
type BaseForkTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*BaseForkTestLogIntIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogIntIterator{contract: _BaseForkTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogInt) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogInt)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_int", log); err != nil {
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

// ParseLogInt is a log parse operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogInt(log types.Log) (*BaseForkTestLogInt, error) {
	event := new(BaseForkTestLogInt)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the BaseForkTest contract.
type BaseForkTestLogNamedAddressIterator struct {
	Event *BaseForkTestLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogNamedAddress)
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
		it.Event = new(BaseForkTestLogNamedAddress)
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
func (it *BaseForkTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogNamedAddress represents a LogNamedAddress event raised by the BaseForkTest contract.
type BaseForkTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*BaseForkTestLogNamedAddressIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogNamedAddressIterator{contract: _BaseForkTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogNamedAddress)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
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

// ParseLogNamedAddress is a log parse operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogNamedAddress(log types.Log) (*BaseForkTestLogNamedAddress, error) {
	event := new(BaseForkTestLogNamedAddress)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the BaseForkTest contract.
type BaseForkTestLogNamedArrayIterator struct {
	Event *BaseForkTestLogNamedArray // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogNamedArray)
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
		it.Event = new(BaseForkTestLogNamedArray)
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
func (it *BaseForkTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogNamedArray represents a LogNamedArray event raised by the BaseForkTest contract.
type BaseForkTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*BaseForkTestLogNamedArrayIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogNamedArrayIterator{contract: _BaseForkTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogNamedArray)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
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

// ParseLogNamedArray is a log parse operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogNamedArray(log types.Log) (*BaseForkTestLogNamedArray, error) {
	event := new(BaseForkTestLogNamedArray)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the BaseForkTest contract.
type BaseForkTestLogNamedArray0Iterator struct {
	Event *BaseForkTestLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogNamedArray0)
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
		it.Event = new(BaseForkTestLogNamedArray0)
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
func (it *BaseForkTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogNamedArray0 represents a LogNamedArray0 event raised by the BaseForkTest contract.
type BaseForkTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*BaseForkTestLogNamedArray0Iterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogNamedArray0Iterator{contract: _BaseForkTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogNamedArray0)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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

// ParseLogNamedArray0 is a log parse operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogNamedArray0(log types.Log) (*BaseForkTestLogNamedArray0, error) {
	event := new(BaseForkTestLogNamedArray0)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the BaseForkTest contract.
type BaseForkTestLogNamedArray1Iterator struct {
	Event *BaseForkTestLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogNamedArray1)
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
		it.Event = new(BaseForkTestLogNamedArray1)
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
func (it *BaseForkTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogNamedArray1 represents a LogNamedArray1 event raised by the BaseForkTest contract.
type BaseForkTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*BaseForkTestLogNamedArray1Iterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogNamedArray1Iterator{contract: _BaseForkTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogNamedArray1)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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

// ParseLogNamedArray1 is a log parse operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogNamedArray1(log types.Log) (*BaseForkTestLogNamedArray1, error) {
	event := new(BaseForkTestLogNamedArray1)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the BaseForkTest contract.
type BaseForkTestLogNamedBytesIterator struct {
	Event *BaseForkTestLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogNamedBytes)
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
		it.Event = new(BaseForkTestLogNamedBytes)
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
func (it *BaseForkTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogNamedBytes represents a LogNamedBytes event raised by the BaseForkTest contract.
type BaseForkTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*BaseForkTestLogNamedBytesIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogNamedBytesIterator{contract: _BaseForkTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogNamedBytes)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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

// ParseLogNamedBytes is a log parse operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogNamedBytes(log types.Log) (*BaseForkTestLogNamedBytes, error) {
	event := new(BaseForkTestLogNamedBytes)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the BaseForkTest contract.
type BaseForkTestLogNamedBytes32Iterator struct {
	Event *BaseForkTestLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogNamedBytes32)
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
		it.Event = new(BaseForkTestLogNamedBytes32)
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
func (it *BaseForkTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the BaseForkTest contract.
type BaseForkTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*BaseForkTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogNamedBytes32Iterator{contract: _BaseForkTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogNamedBytes32)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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

// ParseLogNamedBytes32 is a log parse operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogNamedBytes32(log types.Log) (*BaseForkTestLogNamedBytes32, error) {
	event := new(BaseForkTestLogNamedBytes32)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the BaseForkTest contract.
type BaseForkTestLogNamedDecimalIntIterator struct {
	Event *BaseForkTestLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogNamedDecimalInt)
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
		it.Event = new(BaseForkTestLogNamedDecimalInt)
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
func (it *BaseForkTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the BaseForkTest contract.
type BaseForkTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*BaseForkTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogNamedDecimalIntIterator{contract: _BaseForkTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogNamedDecimalInt)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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

// ParseLogNamedDecimalInt is a log parse operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*BaseForkTestLogNamedDecimalInt, error) {
	event := new(BaseForkTestLogNamedDecimalInt)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the BaseForkTest contract.
type BaseForkTestLogNamedDecimalUintIterator struct {
	Event *BaseForkTestLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogNamedDecimalUint)
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
		it.Event = new(BaseForkTestLogNamedDecimalUint)
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
func (it *BaseForkTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the BaseForkTest contract.
type BaseForkTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*BaseForkTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogNamedDecimalUintIterator{contract: _BaseForkTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogNamedDecimalUint)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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

// ParseLogNamedDecimalUint is a log parse operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*BaseForkTestLogNamedDecimalUint, error) {
	event := new(BaseForkTestLogNamedDecimalUint)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the BaseForkTest contract.
type BaseForkTestLogNamedIntIterator struct {
	Event *BaseForkTestLogNamedInt // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogNamedInt)
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
		it.Event = new(BaseForkTestLogNamedInt)
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
func (it *BaseForkTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogNamedInt represents a LogNamedInt event raised by the BaseForkTest contract.
type BaseForkTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*BaseForkTestLogNamedIntIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogNamedIntIterator{contract: _BaseForkTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogNamedInt)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
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

// ParseLogNamedInt is a log parse operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogNamedInt(log types.Log) (*BaseForkTestLogNamedInt, error) {
	event := new(BaseForkTestLogNamedInt)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the BaseForkTest contract.
type BaseForkTestLogNamedStringIterator struct {
	Event *BaseForkTestLogNamedString // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogNamedString)
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
		it.Event = new(BaseForkTestLogNamedString)
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
func (it *BaseForkTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogNamedString represents a LogNamedString event raised by the BaseForkTest contract.
type BaseForkTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*BaseForkTestLogNamedStringIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogNamedStringIterator{contract: _BaseForkTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogNamedString)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
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

// ParseLogNamedString is a log parse operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogNamedString(log types.Log) (*BaseForkTestLogNamedString, error) {
	event := new(BaseForkTestLogNamedString)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the BaseForkTest contract.
type BaseForkTestLogNamedUintIterator struct {
	Event *BaseForkTestLogNamedUint // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogNamedUint)
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
		it.Event = new(BaseForkTestLogNamedUint)
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
func (it *BaseForkTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogNamedUint represents a LogNamedUint event raised by the BaseForkTest contract.
type BaseForkTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*BaseForkTestLogNamedUintIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogNamedUintIterator{contract: _BaseForkTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogNamedUint)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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

// ParseLogNamedUint is a log parse operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogNamedUint(log types.Log) (*BaseForkTestLogNamedUint, error) {
	event := new(BaseForkTestLogNamedUint)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the BaseForkTest contract.
type BaseForkTestLogStringIterator struct {
	Event *BaseForkTestLogString // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogString)
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
		it.Event = new(BaseForkTestLogString)
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
func (it *BaseForkTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogString represents a LogString event raised by the BaseForkTest contract.
type BaseForkTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogString(opts *bind.FilterOpts) (*BaseForkTestLogStringIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogStringIterator{contract: _BaseForkTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogString) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogString)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_string", log); err != nil {
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

// ParseLogString is a log parse operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogString(log types.Log) (*BaseForkTestLogString, error) {
	event := new(BaseForkTestLogString)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the BaseForkTest contract.
type BaseForkTestLogUintIterator struct {
	Event *BaseForkTestLogUint // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogUint)
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
		it.Event = new(BaseForkTestLogUint)
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
func (it *BaseForkTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogUint represents a LogUint event raised by the BaseForkTest contract.
type BaseForkTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*BaseForkTestLogUintIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogUintIterator{contract: _BaseForkTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogUint) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogUint)
				if err := _BaseForkTest.contract.UnpackLog(event, "log_uint", log); err != nil {
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

// ParseLogUint is a log parse operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogUint(log types.Log) (*BaseForkTestLogUint, error) {
	event := new(BaseForkTestLogUint)
	if err := _BaseForkTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseForkTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the BaseForkTest contract.
type BaseForkTestLogsIterator struct {
	Event *BaseForkTestLogs // Event containing the contract specifics and raw log

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
func (it *BaseForkTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseForkTestLogs)
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
		it.Event = new(BaseForkTestLogs)
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
func (it *BaseForkTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseForkTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseForkTestLogs represents a Logs event raised by the BaseForkTest contract.
type BaseForkTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_BaseForkTest *BaseForkTestFilterer) FilterLogs(opts *bind.FilterOpts) (*BaseForkTestLogsIterator, error) {

	logs, sub, err := _BaseForkTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &BaseForkTestLogsIterator{contract: _BaseForkTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_BaseForkTest *BaseForkTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *BaseForkTestLogs) (event.Subscription, error) {

	logs, sub, err := _BaseForkTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseForkTestLogs)
				if err := _BaseForkTest.contract.UnpackLog(event, "logs", log); err != nil {
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

// ParseLogs is a log parse operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_BaseForkTest *BaseForkTestFilterer) ParseLogs(log types.Log) (*BaseForkTestLogs, error) {
	event := new(BaseForkTestLogs)
	if err := _BaseForkTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
