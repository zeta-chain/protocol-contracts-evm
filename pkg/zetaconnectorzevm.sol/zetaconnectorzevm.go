// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectorzevm

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

// ZetaConnectorZEVMMetaData contains all meta data concerning the ZetaConnectorZEVM contract.
var ZetaConnectorZEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"wzeta_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReceive\",\"inputs\":[{\"name\":\"zetaTxSenderAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zetaValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"onRevert\",\"inputs\":[{\"name\":\"zetaTxSenderAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"remainingZetaValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"send\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structZetaInterfaces.SendInput\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destinationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zetaValueAndGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zetaParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWzetaAddress\",\"inputs\":[{\"name\":\"wzeta_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"wzeta\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"SetWZETA\",\"inputs\":[{\"name\":\"wzeta_\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZetaReceived\",\"inputs\":[{\"name\":\"zetaTxSenderAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zetaValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZetaReverted\",\"inputs\":[{\"name\":\"zetaTxSenderAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"remainingZetaValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZetaSent\",\"inputs\":[{\"name\":\"sourceTxOriginAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"zetaTxSenderAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zetaValueAndGas\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationGasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zetaParams\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyFungibleModule\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrFungible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WZETATransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongValue\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600f57600080fd5b5060405161122f38038061122f833981016040819052602c916050565b600080546001600160a01b0319166001600160a01b0392909216919091179055607e565b600060208284031215606157600080fd5b81516001600160a01b0381168114607757600080fd5b9392505050565b6111a28061008d6000396000f3fe6080604052600436106100685760003560e01c8063942a5e1611610043578063942a5e1614610178578063eb3bacbd1461018b578063ec026901146101ab57600080fd5b8062173d46146100e757806329dd214d1461013d5780633ce4a5bc1461015057600080fd5b366100e25760005473ffffffffffffffffffffffffffffffffffffffff1633148015906100a957503373735b14bb79463307aacbed86daf3322b1e6226ab14155b156100e0576040517f229930b200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b3480156100f357600080fd5b506000546101149073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100e061014b366004610c5e565b6101cb565b34801561015c57600080fd5b5061011473735b14bb79463307aacbed86daf3322b1e6226ab81565b6100e0610186366004610cfe565b610547565b34801561019757600080fd5b506100e06101a6366004610da8565b6108b4565b3480156101b757600080fd5b506100e06101c6366004610dca565b61097a565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610218576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b833414610251576040517f98d4901c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b1580156102b957600080fd5b505af11580156102cd573d6000803e3d6000fd5b50506000546040517f23b872dd00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8a81166024830152604482018a905290911693506323b872dd925060640190506020604051808303816000875af1158015610352573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103769190610e05565b6103ac576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81156104e5578473ffffffffffffffffffffffffffffffffffffffff16633749c51a6040518060a001604052808b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060208082018b905273ffffffffffffffffffffffffffffffffffffffff8a16604080840191909152606083018a90528051601f890183900483028101830190915287815260809092019190889088908190840183828082843760009201919091525050509152506040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526104b29190600401610e8b565b600060405180830381600087803b1580156104cc57600080fd5b505af11580156104e0573d6000803e3d6000fd5b505050505b808573ffffffffffffffffffffffffffffffffffffffff16877ff1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d6988b8b898989604051610535959493929190610f68565b60405180910390a45050505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610594576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8334146105cd576040517f98d4901c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b15801561063557600080fd5b505af1158015610649573d6000803e3d6000fd5b50506000546040517f23b872dd00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8e81166024830152604482018a905290911693506323b872dd925060640190506020604051808303816000875af11580156106ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f29190610e05565b610728576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8115610864578873ffffffffffffffffffffffffffffffffffffffff16633ff0693c6040518060c001604052808c73ffffffffffffffffffffffffffffffffffffffff1681526020018b81526020018a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060208082018a905260408083018a90528051601f890183900483028101830190915287815260609092019190889088908190840183828082843760009201919091525050509152506040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526108319190600401610fa1565b600060405180830381600087803b15801561084b57600080fd5b505af115801561085f573d6000803e3d6000fd5b505050505b80857f521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c888b8b8b8b8a8a8a6040516108a19796959493929190611036565b60405180910390a3505050505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610901576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f7325870b05f8f3412c318a35fc6a74feca51ea15811ec7a257676ca4db9d41769060200160405180910390a150565b6000546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526080830135604482015273ffffffffffffffffffffffffffffffffffffffff909116906323b872dd906064016020604051808303816000875af11580156109fa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1e9190610e05565b610a54576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000546040517f2e1a7d4d0000000000000000000000000000000000000000000000000000000081526080830135600482015273ffffffffffffffffffffffffffffffffffffffff90911690632e1a7d4d90602401600060405180830381600087803b158015610ac357600080fd5b505af1158015610ad7573d6000803e3d6000fd5b50506040516000925073735b14bb79463307aacbed86daf3322b1e6226ab91506080840135908381818185875af1925050503d8060008114610b35576040519150601f19603f3d011682016040523d82523d6000602084013e610b3a565b606091505b5050905080610b75576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8135337f7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e432610ba76020870187611093565b60808801356040890135610bbe60608b018b611093565b610bcb60a08d018d611093565b604051610be0999897969594939291906110f8565b60405180910390a35050565b60008083601f840112610bfe57600080fd5b50813567ffffffffffffffff811115610c1657600080fd5b602083019150836020828501011115610c2e57600080fd5b9250929050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c5957600080fd5b919050565b60008060008060008060008060c0898b031215610c7a57600080fd5b883567ffffffffffffffff811115610c9157600080fd5b610c9d8b828c01610bec565b90995097505060208901359550610cb660408a01610c35565b945060608901359350608089013567ffffffffffffffff811115610cd957600080fd5b610ce58b828c01610bec565b999c989b50969995989497949560a00135949350505050565b600080600080600080600080600060e08a8c031215610d1c57600080fd5b610d258a610c35565b985060208a0135975060408a013567ffffffffffffffff811115610d4857600080fd5b610d548c828d01610bec565b90985096505060608a0135945060808a0135935060a08a013567ffffffffffffffff811115610d8257600080fd5b610d8e8c828d01610bec565b9a9d999c50979a9699959894979660c00135949350505050565b600060208284031215610dba57600080fd5b610dc382610c35565b9392505050565b600060208284031215610ddc57600080fd5b813567ffffffffffffffff811115610df357600080fd5b820160c08185031215610dc357600080fd5b600060208284031215610e1757600080fd5b81518015158114610dc357600080fd5b6000815180845260005b81811015610e4d57602081850181015186830182015201610e31565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000825160a06020840152610ea760c0840182610e27565b90506020840151604084015273ffffffffffffffffffffffffffffffffffffffff60408501511660608401526060840151608084015260808401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160a0850152610f168282610e27565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b606081526000610f7c606083018789610f1f565b8560208401528281036040840152610f95818587610f1f565b98975050505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff8251166020820152602082015160408201526000604083015160c06060840152610fe760e0840182610e27565b905060608401516080840152608084015160a084015260a08401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160c0850152610f168282610e27565b73ffffffffffffffffffffffffffffffffffffffff8816815286602082015260a06040820152600061106c60a083018789610f1f565b8560608401528281036080840152611085818587610f1f565b9a9950505050505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126110c857600080fd5b83018035915067ffffffffffffffff8211156110e357600080fd5b602001915036819003821315610c2e57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8a16815260c06020820152600061112860c083018a8c610f1f565b8860408401528760608401528281036080840152611147818789610f1f565b905082810360a084015261115c818587610f1f565b9c9b50505050505050505050505056fea26469706673582212208e7e07539f932936c5220343de6c6c7ba79341a2a292f69d42e058537070becf64736f6c634300081a0033",
}

// ZetaConnectorZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorZEVMMetaData.ABI instead.
var ZetaConnectorZEVMABI = ZetaConnectorZEVMMetaData.ABI

// ZetaConnectorZEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorZEVMMetaData.Bin instead.
var ZetaConnectorZEVMBin = ZetaConnectorZEVMMetaData.Bin

// DeployZetaConnectorZEVM deploys a new Ethereum contract, binding an instance of ZetaConnectorZEVM to it.
func DeployZetaConnectorZEVM(auth *bind.TransactOpts, backend bind.ContractBackend, wzeta_ common.Address) (common.Address, *types.Transaction, *ZetaConnectorZEVM, error) {
	parsed, err := ZetaConnectorZEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorZEVMBin), backend, wzeta_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorZEVM{ZetaConnectorZEVMCaller: ZetaConnectorZEVMCaller{contract: contract}, ZetaConnectorZEVMTransactor: ZetaConnectorZEVMTransactor{contract: contract}, ZetaConnectorZEVMFilterer: ZetaConnectorZEVMFilterer{contract: contract}}, nil
}

// ZetaConnectorZEVM is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorZEVM struct {
	ZetaConnectorZEVMCaller     // Read-only binding to the contract
	ZetaConnectorZEVMTransactor // Write-only binding to the contract
	ZetaConnectorZEVMFilterer   // Log filterer for contract events
}

// ZetaConnectorZEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorZEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorZEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorZEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorZEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorZEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorZEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorZEVMSession struct {
	Contract     *ZetaConnectorZEVM // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZetaConnectorZEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorZEVMCallerSession struct {
	Contract *ZetaConnectorZEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ZetaConnectorZEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorZEVMTransactorSession struct {
	Contract     *ZetaConnectorZEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ZetaConnectorZEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorZEVMRaw struct {
	Contract *ZetaConnectorZEVM // Generic contract binding to access the raw methods on
}

// ZetaConnectorZEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorZEVMCallerRaw struct {
	Contract *ZetaConnectorZEVMCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorZEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorZEVMTransactorRaw struct {
	Contract *ZetaConnectorZEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorZEVM creates a new instance of ZetaConnectorZEVM, bound to a specific deployed contract.
func NewZetaConnectorZEVM(address common.Address, backend bind.ContractBackend) (*ZetaConnectorZEVM, error) {
	contract, err := bindZetaConnectorZEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVM{ZetaConnectorZEVMCaller: ZetaConnectorZEVMCaller{contract: contract}, ZetaConnectorZEVMTransactor: ZetaConnectorZEVMTransactor{contract: contract}, ZetaConnectorZEVMFilterer: ZetaConnectorZEVMFilterer{contract: contract}}, nil
}

// NewZetaConnectorZEVMCaller creates a new read-only instance of ZetaConnectorZEVM, bound to a specific deployed contract.
func NewZetaConnectorZEVMCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorZEVMCaller, error) {
	contract, err := bindZetaConnectorZEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMCaller{contract: contract}, nil
}

// NewZetaConnectorZEVMTransactor creates a new write-only instance of ZetaConnectorZEVM, bound to a specific deployed contract.
func NewZetaConnectorZEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorZEVMTransactor, error) {
	contract, err := bindZetaConnectorZEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMTransactor{contract: contract}, nil
}

// NewZetaConnectorZEVMFilterer creates a new log filterer instance of ZetaConnectorZEVM, bound to a specific deployed contract.
func NewZetaConnectorZEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorZEVMFilterer, error) {
	contract, err := bindZetaConnectorZEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMFilterer{contract: contract}, nil
}

// bindZetaConnectorZEVM binds a generic wrapper to an already deployed contract.
func bindZetaConnectorZEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorZEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorZEVM.Contract.ZetaConnectorZEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.ZetaConnectorZEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.ZetaConnectorZEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorZEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.contract.Transact(opts, method, params...)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCaller) FUNGIBLEMODULEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorZEVM.contract.Call(opts, &out, "FUNGIBLE_MODULE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.FUNGIBLEMODULEADDRESS(&_ZetaConnectorZEVM.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCallerSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.FUNGIBLEMODULEADDRESS(&_ZetaConnectorZEVM.CallOpts)
}

// Wzeta is a free data retrieval call binding the contract method 0x00173d46.
//
// Solidity: function wzeta() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCaller) Wzeta(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorZEVM.contract.Call(opts, &out, "wzeta")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wzeta is a free data retrieval call binding the contract method 0x00173d46.
//
// Solidity: function wzeta() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) Wzeta() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.Wzeta(&_ZetaConnectorZEVM.CallOpts)
}

// Wzeta is a free data retrieval call binding the contract method 0x00173d46.
//
// Solidity: function wzeta() view returns(address)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMCallerSession) Wzeta() (common.Address, error) {
	return _ZetaConnectorZEVM.Contract.Wzeta(&_ZetaConnectorZEVM.CallOpts)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) OnReceive(opts *bind.TransactOpts, zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "onReceive", zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) OnReceive(zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.OnReceive(&_ZetaConnectorZEVM.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) OnReceive(zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.OnReceive(&_ZetaConnectorZEVM.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) OnRevert(opts *bind.TransactOpts, zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "onRevert", zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) OnRevert(zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.OnRevert(&_ZetaConnectorZEVM.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) OnRevert(zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.OnRevert(&_ZetaConnectorZEVM.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) Send(opts *bind.TransactOpts, input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) Send(input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Send(&_ZetaConnectorZEVM.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) Send(input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Send(&_ZetaConnectorZEVM.TransactOpts, input)
}

// SetWzetaAddress is a paid mutator transaction binding the contract method 0xeb3bacbd.
//
// Solidity: function setWzetaAddress(address wzeta_) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) SetWzetaAddress(opts *bind.TransactOpts, wzeta_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.Transact(opts, "setWzetaAddress", wzeta_)
}

// SetWzetaAddress is a paid mutator transaction binding the contract method 0xeb3bacbd.
//
// Solidity: function setWzetaAddress(address wzeta_) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) SetWzetaAddress(wzeta_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.SetWzetaAddress(&_ZetaConnectorZEVM.TransactOpts, wzeta_)
}

// SetWzetaAddress is a paid mutator transaction binding the contract method 0xeb3bacbd.
//
// Solidity: function setWzetaAddress(address wzeta_) returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) SetWzetaAddress(wzeta_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.SetWzetaAddress(&_ZetaConnectorZEVM.TransactOpts, wzeta_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorZEVM.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMSession) Receive() (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Receive(&_ZetaConnectorZEVM.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaConnectorZEVM *ZetaConnectorZEVMTransactorSession) Receive() (*types.Transaction, error) {
	return _ZetaConnectorZEVM.Contract.Receive(&_ZetaConnectorZEVM.TransactOpts)
}

// ZetaConnectorZEVMSetWZETAIterator is returned from FilterSetWZETA and is used to iterate over the raw logs and unpacked data for SetWZETA events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMSetWZETAIterator struct {
	Event *ZetaConnectorZEVMSetWZETA // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorZEVMSetWZETAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMSetWZETA)
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
		it.Event = new(ZetaConnectorZEVMSetWZETA)
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
func (it *ZetaConnectorZEVMSetWZETAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMSetWZETAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMSetWZETA represents a SetWZETA event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMSetWZETA struct {
	Wzeta common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetWZETA is a free log retrieval operation binding the contract event 0x7325870b05f8f3412c318a35fc6a74feca51ea15811ec7a257676ca4db9d4176.
//
// Solidity: event SetWZETA(address wzeta_)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterSetWZETA(opts *bind.FilterOpts) (*ZetaConnectorZEVMSetWZETAIterator, error) {

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "SetWZETA")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMSetWZETAIterator{contract: _ZetaConnectorZEVM.contract, event: "SetWZETA", logs: logs, sub: sub}, nil
}

// WatchSetWZETA is a free log subscription operation binding the contract event 0x7325870b05f8f3412c318a35fc6a74feca51ea15811ec7a257676ca4db9d4176.
//
// Solidity: event SetWZETA(address wzeta_)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchSetWZETA(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMSetWZETA) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "SetWZETA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMSetWZETA)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "SetWZETA", log); err != nil {
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

// ParseSetWZETA is a log parse operation binding the contract event 0x7325870b05f8f3412c318a35fc6a74feca51ea15811ec7a257676ca4db9d4176.
//
// Solidity: event SetWZETA(address wzeta_)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParseSetWZETA(log types.Log) (*ZetaConnectorZEVMSetWZETA, error) {
	event := new(ZetaConnectorZEVMSetWZETA)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "SetWZETA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorZEVMZetaReceivedIterator is returned from FilterZetaReceived and is used to iterate over the raw logs and unpacked data for ZetaReceived events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaReceivedIterator struct {
	Event *ZetaConnectorZEVMZetaReceived // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorZEVMZetaReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMZetaReceived)
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
		it.Event = new(ZetaConnectorZEVMZetaReceived)
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
func (it *ZetaConnectorZEVMZetaReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMZetaReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMZetaReceived represents a ZetaReceived event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaReceived struct {
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
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterZetaReceived(opts *bind.FilterOpts, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (*ZetaConnectorZEVMZetaReceivedIterator, error) {

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

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "ZetaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMZetaReceivedIterator{contract: _ZetaConnectorZEVM.contract, event: "ZetaReceived", logs: logs, sub: sub}, nil
}

// WatchZetaReceived is a free log subscription operation binding the contract event 0xf1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d698.
//
// Solidity: event ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchZetaReceived(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMZetaReceived, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "ZetaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMZetaReceived)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaReceived", log); err != nil {
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
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParseZetaReceived(log types.Log) (*ZetaConnectorZEVMZetaReceived, error) {
	event := new(ZetaConnectorZEVMZetaReceived)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorZEVMZetaRevertedIterator is returned from FilterZetaReverted and is used to iterate over the raw logs and unpacked data for ZetaReverted events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaRevertedIterator struct {
	Event *ZetaConnectorZEVMZetaReverted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorZEVMZetaRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMZetaReverted)
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
		it.Event = new(ZetaConnectorZEVMZetaReverted)
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
func (it *ZetaConnectorZEVMZetaRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMZetaRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMZetaReverted represents a ZetaReverted event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaReverted struct {
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
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterZetaReverted(opts *bind.FilterOpts, destinationChainId []*big.Int, internalSendHash [][32]byte) (*ZetaConnectorZEVMZetaRevertedIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "ZetaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMZetaRevertedIterator{contract: _ZetaConnectorZEVM.contract, event: "ZetaReverted", logs: logs, sub: sub}, nil
}

// WatchZetaReverted is a free log subscription operation binding the contract event 0x521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c88.
//
// Solidity: event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchZetaReverted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMZetaReverted, destinationChainId []*big.Int, internalSendHash [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "ZetaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMZetaReverted)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaReverted", log); err != nil {
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
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParseZetaReverted(log types.Log) (*ZetaConnectorZEVMZetaReverted, error) {
	event := new(ZetaConnectorZEVMZetaReverted)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorZEVMZetaSentIterator is returned from FilterZetaSent and is used to iterate over the raw logs and unpacked data for ZetaSent events raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaSentIterator struct {
	Event *ZetaConnectorZEVMZetaSent // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorZEVMZetaSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorZEVMZetaSent)
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
		it.Event = new(ZetaConnectorZEVMZetaSent)
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
func (it *ZetaConnectorZEVMZetaSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorZEVMZetaSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorZEVMZetaSent represents a ZetaSent event raised by the ZetaConnectorZEVM contract.
type ZetaConnectorZEVMZetaSent struct {
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
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) FilterZetaSent(opts *bind.FilterOpts, zetaTxSenderAddress []common.Address, destinationChainId []*big.Int) (*ZetaConnectorZEVMZetaSentIterator, error) {

	var zetaTxSenderAddressRule []interface{}
	for _, zetaTxSenderAddressItem := range zetaTxSenderAddress {
		zetaTxSenderAddressRule = append(zetaTxSenderAddressRule, zetaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.FilterLogs(opts, "ZetaSent", zetaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorZEVMZetaSentIterator{contract: _ZetaConnectorZEVM.contract, event: "ZetaSent", logs: logs, sub: sub}, nil
}

// WatchZetaSent is a free log subscription operation binding the contract event 0x7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e4.
//
// Solidity: event ZetaSent(address sourceTxOriginAddress, address indexed zetaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) WatchZetaSent(opts *bind.WatchOpts, sink chan<- *ZetaConnectorZEVMZetaSent, zetaTxSenderAddress []common.Address, destinationChainId []*big.Int) (event.Subscription, error) {

	var zetaTxSenderAddressRule []interface{}
	for _, zetaTxSenderAddressItem := range zetaTxSenderAddress {
		zetaTxSenderAddressRule = append(zetaTxSenderAddressRule, zetaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _ZetaConnectorZEVM.contract.WatchLogs(opts, "ZetaSent", zetaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorZEVMZetaSent)
				if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaSent", log); err != nil {
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
func (_ZetaConnectorZEVM *ZetaConnectorZEVMFilterer) ParseZetaSent(log types.Log) (*ZetaConnectorZEVMZetaSent, error) {
	event := new(ZetaConnectorZEVMZetaSent)
	if err := _ZetaConnectorZEVM.contract.UnpackLog(event, "ZetaSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
