// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnector

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

// ZetaInterfacesSendInput is an auto generated low-level Go binding around an user-defined struct.
type ZetaInterfacesSendInput struct {
	DestinationChainId  *big.Int
	DestinationAddress  []byte
	DestinationGasLimit *big.Int
	Message             []byte
	ZetaValueAndGas     *big.Int
	ZetaParams          []byte
}

// ZetaConnectorBaseMetaData contains all meta data concerning the ZetaConnectorBase contract.
var ZetaConnectorBaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tssAddressUpdater_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pauserAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onReceive\",\"inputs\":[{\"name\":\"zetaTxSenderAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zetaValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRevert\",\"inputs\":[{\"name\":\"zetaTxSenderAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"remainingZetaValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceTssAddressUpdater\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"send\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structZetaInterfaces.SendInput\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destinationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zetaValueAndGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zetaParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssAddressUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePauserAddress\",\"inputs\":[{\"name\":\"pauserAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTssAddress\",\"inputs\":[{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserAddressUpdated\",\"inputs\":[{\"name\":\"callerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTssAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TSSAddressUpdated\",\"inputs\":[{\"name\":\"callerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTssAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TSSAddressUpdaterUpdated\",\"inputs\":[{\"name\":\"callerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTssUpdaterAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZetaReceived\",\"inputs\":[{\"name\":\"zetaTxSenderAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zetaValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZetaReverted\",\"inputs\":[{\"name\":\"zetaTxSenderAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"remainingZetaValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZetaSent\",\"inputs\":[{\"name\":\"sourceTxOriginAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"zetaTxSenderAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zetaValueAndGas\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationGasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zetaParams\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotPauser\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTss\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTssOrUpdater\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTssUpdater\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsMaxSupply\",\"inputs\":[{\"name\":\"maxSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZetaTransferError\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610b9f380380610b9f83398101604081905261002f9161010b565b6000805460ff191690556001600160a01b038416158061005657506001600160a01b038316155b8061006857506001600160a01b038216155b8061007a57506001600160a01b038116155b156100985760405163e6c4247b60e01b815260040160405180910390fd5b6001600160a01b03938416608052600180549385166001600160a01b0319948516179055600280549285169290931691909117909155600080549190921661010002610100600160a81b031990911617905561015f565b80516001600160a01b038116811461010657600080fd5b919050565b6000806000806080858703121561012157600080fd5b61012a856100ef565b9350610138602086016100ef565b9250610146604086016100ef565b9150610154606086016100ef565b905092959194509250565b608051610a26610179600039600060e90152610a266000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c80636128480f1161008c5780639122c344116100665780639122c344146101d0578063942a5e16146101e3578063ec026901146101fc578063f7fb869b1461020d57600080fd5b80636128480f146101ad578063779e3b63146101c05780638456cb59146101c857600080fd5b80633f4ba83a116100bd5780633f4ba83a1461016f5780635b112591146101775780635c975abb1461019757600080fd5b806321e093b1146100e457806329dd214d14610135578063328a01d01461014f575b600080fd5b61010b7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61014d610143366004610849565b5050505050505050565b005b60025461010b9073ffffffffffffffffffffffffffffffffffffffff1681565b61014d610232565b60015461010b9073ffffffffffffffffffffffffffffffffffffffff1681565b60005460ff16604051901515815260200161012c565b61014d6101bb3660046108e9565b610299565b61014d6103c6565b61014d6104eb565b61014d6101de3660046108e9565b61054b565b61014d6101f136600461090b565b505050505050505050565b61014d61020a3660046109b5565b50565b60005461010b90610100900473ffffffffffffffffffffffffffffffffffffffff1681565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16331461028f576040517f4677a0d30000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61029761068b565b565b600054610100900473ffffffffffffffffffffffffffffffffffffffff1633146102f1576040517f4677a0d3000000000000000000000000000000000000000000000000000000008152336004820152602401610286565b73ffffffffffffffffffffffffffffffffffffffff811661033e576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040805133815260208101929092527fd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d039791015b60405180910390a150565b60025473ffffffffffffffffffffffffffffffffffffffff163314610419576040517fe700765e000000000000000000000000000000000000000000000000000000008152336004820152602401610286565b60015473ffffffffffffffffffffffffffffffffffffffff16610468576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600154600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd091015b60405180910390a1565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610543576040517f4677a0d3000000000000000000000000000000000000000000000000000000008152336004820152602401610286565b610297610703565b60015473ffffffffffffffffffffffffffffffffffffffff16331480159061058b575060025473ffffffffffffffffffffffffffffffffffffffff163314155b156105c4576040517fcdfcef97000000000000000000000000000000000000000000000000000000008152336004820152602401610286565b73ffffffffffffffffffffffffffffffffffffffff8116610611576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff91016103bb565b61069361075e565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016104e1565b61070b61079a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586106de3390565b60005460ff16610297576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005460ff1615610297576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008083601f8401126107e957600080fd5b50813567ffffffffffffffff81111561080157600080fd5b60208301915083602082850101111561081957600080fd5b9250929050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084457600080fd5b919050565b60008060008060008060008060c0898b03121561086557600080fd5b883567ffffffffffffffff81111561087c57600080fd5b6108888b828c016107d7565b909950975050602089013595506108a160408a01610820565b945060608901359350608089013567ffffffffffffffff8111156108c457600080fd5b6108d08b828c016107d7565b999c989b50969995989497949560a00135949350505050565b6000602082840312156108fb57600080fd5b61090482610820565b9392505050565b600080600080600080600080600060e08a8c03121561092957600080fd5b6109328a610820565b985060208a0135975060408a013567ffffffffffffffff81111561095557600080fd5b6109618c828d016107d7565b90985096505060608a0135945060808a0135935060a08a013567ffffffffffffffff81111561098f57600080fd5b61099b8c828d016107d7565b9a9d999c50979a9699959894979660c00135949350505050565b6000602082840312156109c757600080fd5b813567ffffffffffffffff8111156109de57600080fd5b820160c0818503121561090457600080fdfea26469706673582212205ed12068a6ae2902ad3ba063d1fd76bca564e296433ff272b8dcf9f49faf22d864736f6c634300081a0033",
}

// ZetaConnectorBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorBaseMetaData.ABI instead.
var ZetaConnectorBaseABI = ZetaConnectorBaseMetaData.ABI

// ZetaConnectorBaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorBaseMetaData.Bin instead.
var ZetaConnectorBaseBin = ZetaConnectorBaseMetaData.Bin

// DeployZetaConnectorBase deploys a new Ethereum contract, binding an instance of ZetaConnectorBase to it.
func DeployZetaConnectorBase(auth *bind.TransactOpts, backend bind.ContractBackend, zetaToken_ common.Address, tssAddress_ common.Address, tssAddressUpdater_ common.Address, pauserAddress_ common.Address) (common.Address, *types.Transaction, *ZetaConnectorBase, error) {
	parsed, err := ZetaConnectorBaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorBaseBin), backend, zetaToken_, tssAddress_, tssAddressUpdater_, pauserAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorBase{ZetaConnectorBaseCaller: ZetaConnectorBaseCaller{contract: contract}, ZetaConnectorBaseTransactor: ZetaConnectorBaseTransactor{contract: contract}, ZetaConnectorBaseFilterer: ZetaConnectorBaseFilterer{contract: contract}}, nil
}

// ZetaConnectorBase is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorBase struct {
	ZetaConnectorBaseCaller     // Read-only binding to the contract
	ZetaConnectorBaseTransactor // Write-only binding to the contract
	ZetaConnectorBaseFilterer   // Log filterer for contract events
}

// ZetaConnectorBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorBaseSession struct {
	Contract     *ZetaConnectorBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZetaConnectorBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorBaseCallerSession struct {
	Contract *ZetaConnectorBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ZetaConnectorBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorBaseTransactorSession struct {
	Contract     *ZetaConnectorBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ZetaConnectorBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorBaseRaw struct {
	Contract *ZetaConnectorBase // Generic contract binding to access the raw methods on
}

// ZetaConnectorBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorBaseCallerRaw struct {
	Contract *ZetaConnectorBaseCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorBaseTransactorRaw struct {
	Contract *ZetaConnectorBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorBase creates a new instance of ZetaConnectorBase, bound to a specific deployed contract.
func NewZetaConnectorBase(address common.Address, backend bind.ContractBackend) (*ZetaConnectorBase, error) {
	contract, err := bindZetaConnectorBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBase{ZetaConnectorBaseCaller: ZetaConnectorBaseCaller{contract: contract}, ZetaConnectorBaseTransactor: ZetaConnectorBaseTransactor{contract: contract}, ZetaConnectorBaseFilterer: ZetaConnectorBaseFilterer{contract: contract}}, nil
}

// NewZetaConnectorBaseCaller creates a new read-only instance of ZetaConnectorBase, bound to a specific deployed contract.
func NewZetaConnectorBaseCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorBaseCaller, error) {
	contract, err := bindZetaConnectorBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseCaller{contract: contract}, nil
}

// NewZetaConnectorBaseTransactor creates a new write-only instance of ZetaConnectorBase, bound to a specific deployed contract.
func NewZetaConnectorBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorBaseTransactor, error) {
	contract, err := bindZetaConnectorBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseTransactor{contract: contract}, nil
}

// NewZetaConnectorBaseFilterer creates a new log filterer instance of ZetaConnectorBase, bound to a specific deployed contract.
func NewZetaConnectorBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorBaseFilterer, error) {
	contract, err := bindZetaConnectorBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseFilterer{contract: contract}, nil
}

// bindZetaConnectorBase binds a generic wrapper to an already deployed contract.
func bindZetaConnectorBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorBase *ZetaConnectorBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorBase.Contract.ZetaConnectorBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorBase *ZetaConnectorBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.ZetaConnectorBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorBase *ZetaConnectorBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.ZetaConnectorBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorBase *ZetaConnectorBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorBase *ZetaConnectorBaseCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorBase.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorBase *ZetaConnectorBaseSession) Paused() (bool, error) {
	return _ZetaConnectorBase.Contract.Paused(&_ZetaConnectorBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorBase *ZetaConnectorBaseCallerSession) Paused() (bool, error) {
	return _ZetaConnectorBase.Contract.Paused(&_ZetaConnectorBase.CallOpts)
}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseCaller) PauserAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorBase.contract.Call(opts, &out, "pauserAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseSession) PauserAddress() (common.Address, error) {
	return _ZetaConnectorBase.Contract.PauserAddress(&_ZetaConnectorBase.CallOpts)
}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseCallerSession) PauserAddress() (common.Address, error) {
	return _ZetaConnectorBase.Contract.PauserAddress(&_ZetaConnectorBase.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorBase.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorBase.Contract.TssAddress(&_ZetaConnectorBase.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseCallerSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorBase.Contract.TssAddress(&_ZetaConnectorBase.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseCaller) TssAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorBase.contract.Call(opts, &out, "tssAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseSession) TssAddressUpdater() (common.Address, error) {
	return _ZetaConnectorBase.Contract.TssAddressUpdater(&_ZetaConnectorBase.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseCallerSession) TssAddressUpdater() (common.Address, error) {
	return _ZetaConnectorBase.Contract.TssAddressUpdater(&_ZetaConnectorBase.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorBase.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorBase.Contract.ZetaToken(&_ZetaConnectorBase.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorBase *ZetaConnectorBaseCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorBase.Contract.ZetaToken(&_ZetaConnectorBase.CallOpts)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactor) OnReceive(opts *bind.TransactOpts, zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.contract.Transact(opts, "onReceive", zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseSession) OnReceive(zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.OnReceive(&_ZetaConnectorBase.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorSession) OnReceive(zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.OnReceive(&_ZetaConnectorBase.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactor) OnRevert(opts *bind.TransactOpts, zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.contract.Transact(opts, "onRevert", zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseSession) OnRevert(zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.OnRevert(&_ZetaConnectorBase.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorSession) OnRevert(zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.OnRevert(&_ZetaConnectorBase.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorBase.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorBase *ZetaConnectorBaseSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.Pause(&_ZetaConnectorBase.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.Pause(&_ZetaConnectorBase.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactor) RenounceTssAddressUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorBase.contract.Transact(opts, "renounceTssAddressUpdater")
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaConnectorBase *ZetaConnectorBaseSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.RenounceTssAddressUpdater(&_ZetaConnectorBase.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.RenounceTssAddressUpdater(&_ZetaConnectorBase.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactor) Send(opts *bind.TransactOpts, input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorBase.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseSession) Send(input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.Send(&_ZetaConnectorBase.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorSession) Send(input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.Send(&_ZetaConnectorBase.TransactOpts, input)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorBase.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorBase *ZetaConnectorBaseSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.Unpause(&_ZetaConnectorBase.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.Unpause(&_ZetaConnectorBase.TransactOpts)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactor) UpdatePauserAddress(opts *bind.TransactOpts, pauserAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorBase.contract.Transact(opts, "updatePauserAddress", pauserAddress_)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseSession) UpdatePauserAddress(pauserAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.UpdatePauserAddress(&_ZetaConnectorBase.TransactOpts, pauserAddress_)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorSession) UpdatePauserAddress(pauserAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.UpdatePauserAddress(&_ZetaConnectorBase.TransactOpts, pauserAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactor) UpdateTssAddress(opts *bind.TransactOpts, tssAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorBase.contract.Transact(opts, "updateTssAddress", tssAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseSession) UpdateTssAddress(tssAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.UpdateTssAddress(&_ZetaConnectorBase.TransactOpts, tssAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_ZetaConnectorBase *ZetaConnectorBaseTransactorSession) UpdateTssAddress(tssAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorBase.Contract.UpdateTssAddress(&_ZetaConnectorBase.TransactOpts, tssAddress_)
}

// ZetaConnectorBasePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ZetaConnectorBase contract.
type ZetaConnectorBasePausedIterator struct {
	Event *ZetaConnectorBasePaused // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorBasePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorBasePaused)
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
		it.Event = new(ZetaConnectorBasePaused)
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
func (it *ZetaConnectorBasePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorBasePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorBasePaused represents a Paused event raised by the ZetaConnectorBase contract.
type ZetaConnectorBasePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) FilterPaused(opts *bind.FilterOpts) (*ZetaConnectorBasePausedIterator, error) {

	logs, sub, err := _ZetaConnectorBase.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBasePausedIterator{contract: _ZetaConnectorBase.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorBasePaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorBase.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorBasePaused)
				if err := _ZetaConnectorBase.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) ParsePaused(log types.Log) (*ZetaConnectorBasePaused, error) {
	event := new(ZetaConnectorBasePaused)
	if err := _ZetaConnectorBase.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorBasePauserAddressUpdatedIterator is returned from FilterPauserAddressUpdated and is used to iterate over the raw logs and unpacked data for PauserAddressUpdated events raised by the ZetaConnectorBase contract.
type ZetaConnectorBasePauserAddressUpdatedIterator struct {
	Event *ZetaConnectorBasePauserAddressUpdated // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorBasePauserAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorBasePauserAddressUpdated)
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
		it.Event = new(ZetaConnectorBasePauserAddressUpdated)
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
func (it *ZetaConnectorBasePauserAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorBasePauserAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorBasePauserAddressUpdated represents a PauserAddressUpdated event raised by the ZetaConnectorBase contract.
type ZetaConnectorBasePauserAddressUpdated struct {
	CallerAddress common.Address
	NewTssAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPauserAddressUpdated is a free log retrieval operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) FilterPauserAddressUpdated(opts *bind.FilterOpts) (*ZetaConnectorBasePauserAddressUpdatedIterator, error) {

	logs, sub, err := _ZetaConnectorBase.contract.FilterLogs(opts, "PauserAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBasePauserAddressUpdatedIterator{contract: _ZetaConnectorBase.contract, event: "PauserAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchPauserAddressUpdated is a free log subscription operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) WatchPauserAddressUpdated(opts *bind.WatchOpts, sink chan<- *ZetaConnectorBasePauserAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorBase.contract.WatchLogs(opts, "PauserAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorBasePauserAddressUpdated)
				if err := _ZetaConnectorBase.contract.UnpackLog(event, "PauserAddressUpdated", log); err != nil {
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

// ParsePauserAddressUpdated is a log parse operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) ParsePauserAddressUpdated(log types.Log) (*ZetaConnectorBasePauserAddressUpdated, error) {
	event := new(ZetaConnectorBasePauserAddressUpdated)
	if err := _ZetaConnectorBase.contract.UnpackLog(event, "PauserAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorBaseTSSAddressUpdatedIterator is returned from FilterTSSAddressUpdated and is used to iterate over the raw logs and unpacked data for TSSAddressUpdated events raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseTSSAddressUpdatedIterator struct {
	Event *ZetaConnectorBaseTSSAddressUpdated // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorBaseTSSAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorBaseTSSAddressUpdated)
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
		it.Event = new(ZetaConnectorBaseTSSAddressUpdated)
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
func (it *ZetaConnectorBaseTSSAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorBaseTSSAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorBaseTSSAddressUpdated represents a TSSAddressUpdated event raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseTSSAddressUpdated struct {
	CallerAddress common.Address
	NewTssAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTSSAddressUpdated is a free log retrieval operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) FilterTSSAddressUpdated(opts *bind.FilterOpts) (*ZetaConnectorBaseTSSAddressUpdatedIterator, error) {

	logs, sub, err := _ZetaConnectorBase.contract.FilterLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseTSSAddressUpdatedIterator{contract: _ZetaConnectorBase.contract, event: "TSSAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchTSSAddressUpdated is a free log subscription operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) WatchTSSAddressUpdated(opts *bind.WatchOpts, sink chan<- *ZetaConnectorBaseTSSAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorBase.contract.WatchLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorBaseTSSAddressUpdated)
				if err := _ZetaConnectorBase.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
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

// ParseTSSAddressUpdated is a log parse operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) ParseTSSAddressUpdated(log types.Log) (*ZetaConnectorBaseTSSAddressUpdated, error) {
	event := new(ZetaConnectorBaseTSSAddressUpdated)
	if err := _ZetaConnectorBase.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorBaseTSSAddressUpdaterUpdatedIterator is returned from FilterTSSAddressUpdaterUpdated and is used to iterate over the raw logs and unpacked data for TSSAddressUpdaterUpdated events raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseTSSAddressUpdaterUpdatedIterator struct {
	Event *ZetaConnectorBaseTSSAddressUpdaterUpdated // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorBaseTSSAddressUpdaterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorBaseTSSAddressUpdaterUpdated)
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
		it.Event = new(ZetaConnectorBaseTSSAddressUpdaterUpdated)
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
func (it *ZetaConnectorBaseTSSAddressUpdaterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorBaseTSSAddressUpdaterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorBaseTSSAddressUpdaterUpdated represents a TSSAddressUpdaterUpdated event raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseTSSAddressUpdaterUpdated struct {
	CallerAddress        common.Address
	NewTssUpdaterAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTSSAddressUpdaterUpdated is a free log retrieval operation binding the contract event 0x5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0.
//
// Solidity: event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) FilterTSSAddressUpdaterUpdated(opts *bind.FilterOpts) (*ZetaConnectorBaseTSSAddressUpdaterUpdatedIterator, error) {

	logs, sub, err := _ZetaConnectorBase.contract.FilterLogs(opts, "TSSAddressUpdaterUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseTSSAddressUpdaterUpdatedIterator{contract: _ZetaConnectorBase.contract, event: "TSSAddressUpdaterUpdated", logs: logs, sub: sub}, nil
}

// WatchTSSAddressUpdaterUpdated is a free log subscription operation binding the contract event 0x5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0.
//
// Solidity: event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) WatchTSSAddressUpdaterUpdated(opts *bind.WatchOpts, sink chan<- *ZetaConnectorBaseTSSAddressUpdaterUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorBase.contract.WatchLogs(opts, "TSSAddressUpdaterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorBaseTSSAddressUpdaterUpdated)
				if err := _ZetaConnectorBase.contract.UnpackLog(event, "TSSAddressUpdaterUpdated", log); err != nil {
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

// ParseTSSAddressUpdaterUpdated is a log parse operation binding the contract event 0x5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0.
//
// Solidity: event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) ParseTSSAddressUpdaterUpdated(log types.Log) (*ZetaConnectorBaseTSSAddressUpdaterUpdated, error) {
	event := new(ZetaConnectorBaseTSSAddressUpdaterUpdated)
	if err := _ZetaConnectorBase.contract.UnpackLog(event, "TSSAddressUpdaterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorBaseUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseUnpausedIterator struct {
	Event *ZetaConnectorBaseUnpaused // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorBaseUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorBaseUnpaused)
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
		it.Event = new(ZetaConnectorBaseUnpaused)
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
func (it *ZetaConnectorBaseUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorBaseUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorBaseUnpaused represents a Unpaused event raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ZetaConnectorBaseUnpausedIterator, error) {

	logs, sub, err := _ZetaConnectorBase.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseUnpausedIterator{contract: _ZetaConnectorBase.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorBaseUnpaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorBase.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorBaseUnpaused)
				if err := _ZetaConnectorBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) ParseUnpaused(log types.Log) (*ZetaConnectorBaseUnpaused, error) {
	event := new(ZetaConnectorBaseUnpaused)
	if err := _ZetaConnectorBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorBaseZetaReceivedIterator is returned from FilterZetaReceived and is used to iterate over the raw logs and unpacked data for ZetaReceived events raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseZetaReceivedIterator struct {
	Event *ZetaConnectorBaseZetaReceived // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorBaseZetaReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorBaseZetaReceived)
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
		it.Event = new(ZetaConnectorBaseZetaReceived)
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
func (it *ZetaConnectorBaseZetaReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorBaseZetaReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorBaseZetaReceived represents a ZetaReceived event raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseZetaReceived struct {
	ZetaTxSenderAddress []byte
	SourceChainId       *big.Int
	DestinationAddress  common.Address
	ZetaValue           *big.Int
	Message             []byte
	InternalSendHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterZetaReceived is a free log retrieval operation binding the contract event 0xf1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d698.
//
// Solidity: event ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) FilterZetaReceived(opts *bind.FilterOpts, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (*ZetaConnectorBaseZetaReceivedIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorBase.contract.FilterLogs(opts, "ZetaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseZetaReceivedIterator{contract: _ZetaConnectorBase.contract, event: "ZetaReceived", logs: logs, sub: sub}, nil
}

// WatchZetaReceived is a free log subscription operation binding the contract event 0xf1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d698.
//
// Solidity: event ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) WatchZetaReceived(opts *bind.WatchOpts, sink chan<- *ZetaConnectorBaseZetaReceived, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorBase.contract.WatchLogs(opts, "ZetaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorBaseZetaReceived)
				if err := _ZetaConnectorBase.contract.UnpackLog(event, "ZetaReceived", log); err != nil {
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

// ParseZetaReceived is a log parse operation binding the contract event 0xf1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d698.
//
// Solidity: event ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) ParseZetaReceived(log types.Log) (*ZetaConnectorBaseZetaReceived, error) {
	event := new(ZetaConnectorBaseZetaReceived)
	if err := _ZetaConnectorBase.contract.UnpackLog(event, "ZetaReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorBaseZetaRevertedIterator is returned from FilterZetaReverted and is used to iterate over the raw logs and unpacked data for ZetaReverted events raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseZetaRevertedIterator struct {
	Event *ZetaConnectorBaseZetaReverted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorBaseZetaRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorBaseZetaReverted)
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
		it.Event = new(ZetaConnectorBaseZetaReverted)
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
func (it *ZetaConnectorBaseZetaRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorBaseZetaRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorBaseZetaReverted represents a ZetaReverted event raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseZetaReverted struct {
	ZetaTxSenderAddress common.Address
	SourceChainId       *big.Int
	DestinationChainId  *big.Int
	DestinationAddress  []byte
	RemainingZetaValue  *big.Int
	Message             []byte
	InternalSendHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterZetaReverted is a free log retrieval operation binding the contract event 0x521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c88.
//
// Solidity: event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) FilterZetaReverted(opts *bind.FilterOpts, destinationChainId []*big.Int, internalSendHash [][32]byte) (*ZetaConnectorBaseZetaRevertedIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorBase.contract.FilterLogs(opts, "ZetaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseZetaRevertedIterator{contract: _ZetaConnectorBase.contract, event: "ZetaReverted", logs: logs, sub: sub}, nil
}

// WatchZetaReverted is a free log subscription operation binding the contract event 0x521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c88.
//
// Solidity: event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) WatchZetaReverted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorBaseZetaReverted, destinationChainId []*big.Int, internalSendHash [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorBase.contract.WatchLogs(opts, "ZetaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorBaseZetaReverted)
				if err := _ZetaConnectorBase.contract.UnpackLog(event, "ZetaReverted", log); err != nil {
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

// ParseZetaReverted is a log parse operation binding the contract event 0x521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c88.
//
// Solidity: event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) ParseZetaReverted(log types.Log) (*ZetaConnectorBaseZetaReverted, error) {
	event := new(ZetaConnectorBaseZetaReverted)
	if err := _ZetaConnectorBase.contract.UnpackLog(event, "ZetaReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorBaseZetaSentIterator is returned from FilterZetaSent and is used to iterate over the raw logs and unpacked data for ZetaSent events raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseZetaSentIterator struct {
	Event *ZetaConnectorBaseZetaSent // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorBaseZetaSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorBaseZetaSent)
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
		it.Event = new(ZetaConnectorBaseZetaSent)
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
func (it *ZetaConnectorBaseZetaSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorBaseZetaSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorBaseZetaSent represents a ZetaSent event raised by the ZetaConnectorBase contract.
type ZetaConnectorBaseZetaSent struct {
	SourceTxOriginAddress common.Address
	ZetaTxSenderAddress   common.Address
	DestinationChainId    *big.Int
	DestinationAddress    []byte
	ZetaValueAndGas       *big.Int
	DestinationGasLimit   *big.Int
	Message               []byte
	ZetaParams            []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterZetaSent is a free log retrieval operation binding the contract event 0x7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e4.
//
// Solidity: event ZetaSent(address sourceTxOriginAddress, address indexed zetaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) FilterZetaSent(opts *bind.FilterOpts, zetaTxSenderAddress []common.Address, destinationChainId []*big.Int) (*ZetaConnectorBaseZetaSentIterator, error) {

	var zetaTxSenderAddressRule []interface{}
	for _, zetaTxSenderAddressItem := range zetaTxSenderAddress {
		zetaTxSenderAddressRule = append(zetaTxSenderAddressRule, zetaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _ZetaConnectorBase.contract.FilterLogs(opts, "ZetaSent", zetaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorBaseZetaSentIterator{contract: _ZetaConnectorBase.contract, event: "ZetaSent", logs: logs, sub: sub}, nil
}

// WatchZetaSent is a free log subscription operation binding the contract event 0x7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e4.
//
// Solidity: event ZetaSent(address sourceTxOriginAddress, address indexed zetaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) WatchZetaSent(opts *bind.WatchOpts, sink chan<- *ZetaConnectorBaseZetaSent, zetaTxSenderAddress []common.Address, destinationChainId []*big.Int) (event.Subscription, error) {

	var zetaTxSenderAddressRule []interface{}
	for _, zetaTxSenderAddressItem := range zetaTxSenderAddress {
		zetaTxSenderAddressRule = append(zetaTxSenderAddressRule, zetaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _ZetaConnectorBase.contract.WatchLogs(opts, "ZetaSent", zetaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorBaseZetaSent)
				if err := _ZetaConnectorBase.contract.UnpackLog(event, "ZetaSent", log); err != nil {
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

// ParseZetaSent is a log parse operation binding the contract event 0x7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e4.
//
// Solidity: event ZetaSent(address sourceTxOriginAddress, address indexed zetaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
func (_ZetaConnectorBase *ZetaConnectorBaseFilterer) ParseZetaSent(log types.Log) (*ZetaConnectorBaseZetaSent, error) {
	event := new(ZetaConnectorBaseZetaSent)
	if err := _ZetaConnectorBase.contract.UnpackLog(event, "ZetaSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
