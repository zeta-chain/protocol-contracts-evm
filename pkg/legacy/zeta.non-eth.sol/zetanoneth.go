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

// ZetaNonEthMetaData contains all meta data concerning the ZetaNonEth contract.
var ZetaNonEthMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tssAddressUpdater_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"connectorAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"mintee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceTssAddressUpdater\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssAddressUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateTssAndConnectorAddresses\",\"inputs\":[{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"connectorAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Burnt\",\"inputs\":[{\"name\":\"burnee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConnectorAddressUpdated\",\"inputs\":[{\"name\":\"callerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newConnectorAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Minted\",\"inputs\":[{\"name\":\"mintee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TSSAddressUpdated\",\"inputs\":[{\"name\":\"callerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTssAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TSSAddressUpdaterUpdated\",\"inputs\":[{\"name\":\"callerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTssUpdaterAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotConnector\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTss\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTssOrUpdater\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTssUpdater\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZetaTransferError\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161124c38038061124c83398101604081905261002e9161010e565b604051806040016040528060048152602001635a65746160e01b815250604051806040016040528060048152602001635a45544160e01b815250816003908161007791906101d7565b50600461008482826101d7565b5050506001600160a01b03821615806100a457506001600160a01b038116155b156100c25760405163e6c4247b60e01b815260040160405180910390fd5b600680546001600160a01b039384166001600160a01b03199182161790915560078054929093169116179055610291565b80516001600160a01b0381168114610109575f80fd5b919050565b5f806040838503121561011f575f80fd5b610128836100f3565b9150610136602084016100f3565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061016757607f821691505b60208210810361018557634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101d257805f5260205f20601f840160051c810160208510156101b05750805b601f840160051c820191505b818110156101cf575f81556001016101bc565b50505b505050565b81516001600160401b038111156101f0576101f061013f565b610204816101fe8454610153565b8461018b565b6020601f821160018114610236575f831561021f5750848201515b5f19600385901b1c1916600184901b1784556101cf565b5f84815260208120601f198516915b828110156102655787850151825560209485019460019092019101610245565b508482101561028257868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b610fae8061029e5f395ff3fe608060405234801561000f575f80fd5b5060043610610115575f3560e01c806342966c68116100ad57806379cc67901161007d578063a9059cbb11610063578063a9059cbb14610286578063bff9662a14610299578063dd62ed3e146102b9575f80fd5b806379cc67901461026b57806395d89b411461027e575f80fd5b806342966c68146101fb5780635b1125911461020e57806370a082311461022e578063779e3b6314610263575f80fd5b80631e458bee116100e85780631e458bee1461018157806323b872dd14610194578063313ce567146101a7578063328a01d0146101b6575f80fd5b806306fdde0314610119578063095ea7b31461013757806315d57fd41461015a57806318160ddd1461016f575b5f80fd5b6101216102fe565b60405161012e9190610d7a565b60405180910390f35b61014a610145366004610df5565b61038e565b604051901515815260200161012e565b61016d610168366004610e1d565b6103a7565b005b6002545b60405190815260200161012e565b61016d61018f366004610e4e565b610572565b61014a6101a2366004610e7e565b610625565b6040516012815260200161012e565b6007546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161012e565b61016d610209366004610eb8565b610648565b6006546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b61017361023c366004610ecf565b73ffffffffffffffffffffffffffffffffffffffff165f9081526020819052604090205490565b61016d610655565b61016d610279366004610df5565b610779565b61012161082a565b61014a610294366004610df5565b610839565b6005546101d69073ffffffffffffffffffffffffffffffffffffffff1681565b6101736102c7366004610e1d565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260016020908152604080832093909416825291909152205490565b60606003805461030d90610eef565b80601f016020809104026020016040519081016040528092919081815260200182805461033990610eef565b80156103845780601f1061035b57610100808354040283529160200191610384565b820191905f5260205f20905b81548152906001019060200180831161036757829003601f168201915b5050505050905090565b5f3361039b818585610846565b60019150505b92915050565b60075473ffffffffffffffffffffffffffffffffffffffff1633148015906103e7575060065473ffffffffffffffffffffffffffffffffffffffff163314155b15610425576040517fcdfcef970000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216158061045c575073ffffffffffffffffffffffffffffffffffffffff8116155b15610493576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316811790935560058054918516919092161790556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff910160405180910390a16040805133815273ffffffffffffffffffffffffffffffffffffffff831660208201527f1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c910160405180910390a15050565b60055473ffffffffffffffffffffffffffffffffffffffff1633146105c5576040517f3fe32fba00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b6105cf8383610858565b808373ffffffffffffffffffffffffffffffffffffffff167fc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb8460405161061891815260200190565b60405180910390a3505050565b5f336106328582856108b6565b61063d858585610983565b506001949350505050565b6106523382610a2c565b50565b60075473ffffffffffffffffffffffffffffffffffffffff1633146106a8576040517fe700765e00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b60065473ffffffffffffffffffffffffffffffffffffffff166106f7576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600654600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0910160405180910390a1565b60055473ffffffffffffffffffffffffffffffffffffffff1633146107cc576040517f3fe32fba00000000000000000000000000000000000000000000000000000000815233600482015260240161041c565b6107d68282610a86565b8173ffffffffffffffffffffffffffffffffffffffff167f919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b18260405161081e91815260200190565b60405180910390a25050565b60606004805461030d90610eef565b5f3361039b818585610983565b6108538383836001610a9b565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166108a7576040517fec442f050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b6108b25f8383610be0565b5050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461097d578181101561096f576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602481018290526044810183905260640161041c565b61097d84848484035f610a9b565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166109d2576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8216610a21576040517fec442f050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b610853838383610be0565b73ffffffffffffffffffffffffffffffffffffffff8216610a7b576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b6108b2825f83610be0565b610a918233836108b6565b6108b28282610a2c565b73ffffffffffffffffffffffffffffffffffffffff8416610aea576040517fe602df050000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8316610b39576040517f94280d620000000000000000000000000000000000000000000000000000000081525f600482015260240161041c565b73ffffffffffffffffffffffffffffffffffffffff8085165f908152600160209081526040808320938716835292905220829055801561097d578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bd291815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c17578060025f828254610c0c9190610f40565b90915550610cc79050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526020819052604090205481811015610c9c576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602481018290526044810183905260640161041c565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610cf057600280548290039055610d1b565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161061891815260200190565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610df0575f80fd5b919050565b5f8060408385031215610e06575f80fd5b610e0f83610dcd565b946020939093013593505050565b5f8060408385031215610e2e575f80fd5b610e3783610dcd565b9150610e4560208401610dcd565b90509250929050565b5f805f60608486031215610e60575f80fd5b610e6984610dcd565b95602085013595506040909401359392505050565b5f805f60608486031215610e90575f80fd5b610e9984610dcd565b9250610ea760208501610dcd565b929592945050506040919091013590565b5f60208284031215610ec8575f80fd5b5035919050565b5f60208284031215610edf575f80fd5b610ee882610dcd565b9392505050565b600181811c90821680610f0357607f821691505b602082108103610f3a577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b808201808211156103a1577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffdfea2646970667358221220707a28036036225d7550d60f6cd194bb9df5e81f70d69970e578c66d0fd7b8eb64736f6c634300081a0033",
}

// ZetaNonEthABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaNonEthMetaData.ABI instead.
var ZetaNonEthABI = ZetaNonEthMetaData.ABI

// ZetaNonEthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaNonEthMetaData.Bin instead.
var ZetaNonEthBin = ZetaNonEthMetaData.Bin

// DeployZetaNonEth deploys a new Ethereum contract, binding an instance of ZetaNonEth to it.
func DeployZetaNonEth(auth *bind.TransactOpts, backend bind.ContractBackend, tssAddress_ common.Address, tssAddressUpdater_ common.Address) (common.Address, *types.Transaction, *ZetaNonEth, error) {
	parsed, err := ZetaNonEthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaNonEthBin), backend, tssAddress_, tssAddressUpdater_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaNonEth{ZetaNonEthCaller: ZetaNonEthCaller{contract: contract}, ZetaNonEthTransactor: ZetaNonEthTransactor{contract: contract}, ZetaNonEthFilterer: ZetaNonEthFilterer{contract: contract}}, nil
}

// ZetaNonEth is an auto generated Go binding around an Ethereum contract.
type ZetaNonEth struct {
	ZetaNonEthCaller     // Read-only binding to the contract
	ZetaNonEthTransactor // Write-only binding to the contract
	ZetaNonEthFilterer   // Log filterer for contract events
}

// ZetaNonEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaNonEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaNonEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaNonEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaNonEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaNonEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaNonEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaNonEthSession struct {
	Contract     *ZetaNonEth       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZetaNonEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaNonEthCallerSession struct {
	Contract *ZetaNonEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZetaNonEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaNonEthTransactorSession struct {
	Contract     *ZetaNonEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZetaNonEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaNonEthRaw struct {
	Contract *ZetaNonEth // Generic contract binding to access the raw methods on
}

// ZetaNonEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaNonEthCallerRaw struct {
	Contract *ZetaNonEthCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaNonEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaNonEthTransactorRaw struct {
	Contract *ZetaNonEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaNonEth creates a new instance of ZetaNonEth, bound to a specific deployed contract.
func NewZetaNonEth(address common.Address, backend bind.ContractBackend) (*ZetaNonEth, error) {
	contract, err := bindZetaNonEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEth{ZetaNonEthCaller: ZetaNonEthCaller{contract: contract}, ZetaNonEthTransactor: ZetaNonEthTransactor{contract: contract}, ZetaNonEthFilterer: ZetaNonEthFilterer{contract: contract}}, nil
}

// NewZetaNonEthCaller creates a new read-only instance of ZetaNonEth, bound to a specific deployed contract.
func NewZetaNonEthCaller(address common.Address, caller bind.ContractCaller) (*ZetaNonEthCaller, error) {
	contract, err := bindZetaNonEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthCaller{contract: contract}, nil
}

// NewZetaNonEthTransactor creates a new write-only instance of ZetaNonEth, bound to a specific deployed contract.
func NewZetaNonEthTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaNonEthTransactor, error) {
	contract, err := bindZetaNonEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthTransactor{contract: contract}, nil
}

// NewZetaNonEthFilterer creates a new log filterer instance of ZetaNonEth, bound to a specific deployed contract.
func NewZetaNonEthFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaNonEthFilterer, error) {
	contract, err := bindZetaNonEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthFilterer{contract: contract}, nil
}

// bindZetaNonEth binds a generic wrapper to an already deployed contract.
func bindZetaNonEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaNonEthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaNonEth *ZetaNonEthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaNonEth.Contract.ZetaNonEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaNonEth *ZetaNonEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.ZetaNonEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaNonEth *ZetaNonEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.ZetaNonEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaNonEth *ZetaNonEthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaNonEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaNonEth *ZetaNonEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaNonEth *ZetaNonEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZetaNonEth *ZetaNonEthCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZetaNonEth *ZetaNonEthSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZetaNonEth.Contract.Allowance(&_ZetaNonEth.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZetaNonEth *ZetaNonEthCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZetaNonEth.Contract.Allowance(&_ZetaNonEth.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZetaNonEth *ZetaNonEthCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZetaNonEth *ZetaNonEthSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZetaNonEth.Contract.BalanceOf(&_ZetaNonEth.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZetaNonEth *ZetaNonEthCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZetaNonEth.Contract.BalanceOf(&_ZetaNonEth.CallOpts, account)
}

// ConnectorAddress is a free data retrieval call binding the contract method 0xbff9662a.
//
// Solidity: function connectorAddress() view returns(address)
func (_ZetaNonEth *ZetaNonEthCaller) ConnectorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "connectorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorAddress is a free data retrieval call binding the contract method 0xbff9662a.
//
// Solidity: function connectorAddress() view returns(address)
func (_ZetaNonEth *ZetaNonEthSession) ConnectorAddress() (common.Address, error) {
	return _ZetaNonEth.Contract.ConnectorAddress(&_ZetaNonEth.CallOpts)
}

// ConnectorAddress is a free data retrieval call binding the contract method 0xbff9662a.
//
// Solidity: function connectorAddress() view returns(address)
func (_ZetaNonEth *ZetaNonEthCallerSession) ConnectorAddress() (common.Address, error) {
	return _ZetaNonEth.Contract.ConnectorAddress(&_ZetaNonEth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZetaNonEth *ZetaNonEthCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZetaNonEth *ZetaNonEthSession) Decimals() (uint8, error) {
	return _ZetaNonEth.Contract.Decimals(&_ZetaNonEth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZetaNonEth *ZetaNonEthCallerSession) Decimals() (uint8, error) {
	return _ZetaNonEth.Contract.Decimals(&_ZetaNonEth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZetaNonEth *ZetaNonEthCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZetaNonEth *ZetaNonEthSession) Name() (string, error) {
	return _ZetaNonEth.Contract.Name(&_ZetaNonEth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZetaNonEth *ZetaNonEthCallerSession) Name() (string, error) {
	return _ZetaNonEth.Contract.Name(&_ZetaNonEth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZetaNonEth *ZetaNonEthCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZetaNonEth *ZetaNonEthSession) Symbol() (string, error) {
	return _ZetaNonEth.Contract.Symbol(&_ZetaNonEth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZetaNonEth *ZetaNonEthCallerSession) Symbol() (string, error) {
	return _ZetaNonEth.Contract.Symbol(&_ZetaNonEth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZetaNonEth *ZetaNonEthCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZetaNonEth *ZetaNonEthSession) TotalSupply() (*big.Int, error) {
	return _ZetaNonEth.Contract.TotalSupply(&_ZetaNonEth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZetaNonEth *ZetaNonEthCallerSession) TotalSupply() (*big.Int, error) {
	return _ZetaNonEth.Contract.TotalSupply(&_ZetaNonEth.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaNonEth *ZetaNonEthCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaNonEth *ZetaNonEthSession) TssAddress() (common.Address, error) {
	return _ZetaNonEth.Contract.TssAddress(&_ZetaNonEth.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaNonEth *ZetaNonEthCallerSession) TssAddress() (common.Address, error) {
	return _ZetaNonEth.Contract.TssAddress(&_ZetaNonEth.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaNonEth *ZetaNonEthCaller) TssAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaNonEth.contract.Call(opts, &out, "tssAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaNonEth *ZetaNonEthSession) TssAddressUpdater() (common.Address, error) {
	return _ZetaNonEth.Contract.TssAddressUpdater(&_ZetaNonEth.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaNonEth *ZetaNonEthCallerSession) TssAddressUpdater() (common.Address, error) {
	return _ZetaNonEth.Contract.TssAddressUpdater(&_ZetaNonEth.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ZetaNonEth *ZetaNonEthSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Approve(&_ZetaNonEth.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Approve(&_ZetaNonEth.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ZetaNonEth *ZetaNonEthTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ZetaNonEth *ZetaNonEthSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Burn(&_ZetaNonEth.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ZetaNonEth *ZetaNonEthTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Burn(&_ZetaNonEth.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ZetaNonEth *ZetaNonEthTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ZetaNonEth *ZetaNonEthSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.BurnFrom(&_ZetaNonEth.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ZetaNonEth *ZetaNonEthTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.BurnFrom(&_ZetaNonEth.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_ZetaNonEth *ZetaNonEthTransactor) Mint(opts *bind.TransactOpts, mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "mint", mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_ZetaNonEth *ZetaNonEthSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Mint(&_ZetaNonEth.TransactOpts, mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_ZetaNonEth *ZetaNonEthTransactorSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Mint(&_ZetaNonEth.TransactOpts, mintee, value, internalSendHash)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaNonEth *ZetaNonEthTransactor) RenounceTssAddressUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "renounceTssAddressUpdater")
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaNonEth *ZetaNonEthSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _ZetaNonEth.Contract.RenounceTssAddressUpdater(&_ZetaNonEth.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaNonEth *ZetaNonEthTransactorSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _ZetaNonEth.Contract.RenounceTssAddressUpdater(&_ZetaNonEth.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ZetaNonEth *ZetaNonEthSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Transfer(&_ZetaNonEth.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.Transfer(&_ZetaNonEth.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ZetaNonEth *ZetaNonEthSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.TransferFrom(&_ZetaNonEth.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ZetaNonEth *ZetaNonEthTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.TransferFrom(&_ZetaNonEth.TransactOpts, from, to, value)
}

// UpdateTssAndConnectorAddresses is a paid mutator transaction binding the contract method 0x15d57fd4.
//
// Solidity: function updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) returns()
func (_ZetaNonEth *ZetaNonEthTransactor) UpdateTssAndConnectorAddresses(opts *bind.TransactOpts, tssAddress_ common.Address, connectorAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaNonEth.contract.Transact(opts, "updateTssAndConnectorAddresses", tssAddress_, connectorAddress_)
}

// UpdateTssAndConnectorAddresses is a paid mutator transaction binding the contract method 0x15d57fd4.
//
// Solidity: function updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) returns()
func (_ZetaNonEth *ZetaNonEthSession) UpdateTssAndConnectorAddresses(tssAddress_ common.Address, connectorAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.UpdateTssAndConnectorAddresses(&_ZetaNonEth.TransactOpts, tssAddress_, connectorAddress_)
}

// UpdateTssAndConnectorAddresses is a paid mutator transaction binding the contract method 0x15d57fd4.
//
// Solidity: function updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) returns()
func (_ZetaNonEth *ZetaNonEthTransactorSession) UpdateTssAndConnectorAddresses(tssAddress_ common.Address, connectorAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaNonEth.Contract.UpdateTssAndConnectorAddresses(&_ZetaNonEth.TransactOpts, tssAddress_, connectorAddress_)
}

// ZetaNonEthApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZetaNonEth contract.
type ZetaNonEthApprovalIterator struct {
	Event *ZetaNonEthApproval // Event containing the contract specifics and raw log

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
func (it *ZetaNonEthApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaNonEthApproval)
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
		it.Event = new(ZetaNonEthApproval)
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
func (it *ZetaNonEthApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaNonEthApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaNonEthApproval represents a Approval event raised by the ZetaNonEth contract.
type ZetaNonEthApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZetaNonEth *ZetaNonEthFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ZetaNonEthApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZetaNonEth.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthApprovalIterator{contract: _ZetaNonEth.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZetaNonEth *ZetaNonEthFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZetaNonEthApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZetaNonEth.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaNonEthApproval)
				if err := _ZetaNonEth.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ZetaNonEth *ZetaNonEthFilterer) ParseApproval(log types.Log) (*ZetaNonEthApproval, error) {
	event := new(ZetaNonEthApproval)
	if err := _ZetaNonEth.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaNonEthBurntIterator is returned from FilterBurnt and is used to iterate over the raw logs and unpacked data for Burnt events raised by the ZetaNonEth contract.
type ZetaNonEthBurntIterator struct {
	Event *ZetaNonEthBurnt // Event containing the contract specifics and raw log

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
func (it *ZetaNonEthBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaNonEthBurnt)
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
		it.Event = new(ZetaNonEthBurnt)
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
func (it *ZetaNonEthBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaNonEthBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaNonEthBurnt represents a Burnt event raised by the ZetaNonEth contract.
type ZetaNonEthBurnt struct {
	Burnee common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnt is a free log retrieval operation binding the contract event 0x919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b1.
//
// Solidity: event Burnt(address indexed burnee, uint256 amount)
func (_ZetaNonEth *ZetaNonEthFilterer) FilterBurnt(opts *bind.FilterOpts, burnee []common.Address) (*ZetaNonEthBurntIterator, error) {

	var burneeRule []interface{}
	for _, burneeItem := range burnee {
		burneeRule = append(burneeRule, burneeItem)
	}

	logs, sub, err := _ZetaNonEth.contract.FilterLogs(opts, "Burnt", burneeRule)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthBurntIterator{contract: _ZetaNonEth.contract, event: "Burnt", logs: logs, sub: sub}, nil
}

// WatchBurnt is a free log subscription operation binding the contract event 0x919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b1.
//
// Solidity: event Burnt(address indexed burnee, uint256 amount)
func (_ZetaNonEth *ZetaNonEthFilterer) WatchBurnt(opts *bind.WatchOpts, sink chan<- *ZetaNonEthBurnt, burnee []common.Address) (event.Subscription, error) {

	var burneeRule []interface{}
	for _, burneeItem := range burnee {
		burneeRule = append(burneeRule, burneeItem)
	}

	logs, sub, err := _ZetaNonEth.contract.WatchLogs(opts, "Burnt", burneeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaNonEthBurnt)
				if err := _ZetaNonEth.contract.UnpackLog(event, "Burnt", log); err != nil {
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

// ParseBurnt is a log parse operation binding the contract event 0x919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b1.
//
// Solidity: event Burnt(address indexed burnee, uint256 amount)
func (_ZetaNonEth *ZetaNonEthFilterer) ParseBurnt(log types.Log) (*ZetaNonEthBurnt, error) {
	event := new(ZetaNonEthBurnt)
	if err := _ZetaNonEth.contract.UnpackLog(event, "Burnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaNonEthConnectorAddressUpdatedIterator is returned from FilterConnectorAddressUpdated and is used to iterate over the raw logs and unpacked data for ConnectorAddressUpdated events raised by the ZetaNonEth contract.
type ZetaNonEthConnectorAddressUpdatedIterator struct {
	Event *ZetaNonEthConnectorAddressUpdated // Event containing the contract specifics and raw log

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
func (it *ZetaNonEthConnectorAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaNonEthConnectorAddressUpdated)
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
		it.Event = new(ZetaNonEthConnectorAddressUpdated)
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
func (it *ZetaNonEthConnectorAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaNonEthConnectorAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaNonEthConnectorAddressUpdated represents a ConnectorAddressUpdated event raised by the ZetaNonEth contract.
type ZetaNonEthConnectorAddressUpdated struct {
	CallerAddress       common.Address
	NewConnectorAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterConnectorAddressUpdated is a free log retrieval operation binding the contract event 0x1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c.
//
// Solidity: event ConnectorAddressUpdated(address callerAddress, address newConnectorAddress)
func (_ZetaNonEth *ZetaNonEthFilterer) FilterConnectorAddressUpdated(opts *bind.FilterOpts) (*ZetaNonEthConnectorAddressUpdatedIterator, error) {

	logs, sub, err := _ZetaNonEth.contract.FilterLogs(opts, "ConnectorAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthConnectorAddressUpdatedIterator{contract: _ZetaNonEth.contract, event: "ConnectorAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchConnectorAddressUpdated is a free log subscription operation binding the contract event 0x1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c.
//
// Solidity: event ConnectorAddressUpdated(address callerAddress, address newConnectorAddress)
func (_ZetaNonEth *ZetaNonEthFilterer) WatchConnectorAddressUpdated(opts *bind.WatchOpts, sink chan<- *ZetaNonEthConnectorAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaNonEth.contract.WatchLogs(opts, "ConnectorAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaNonEthConnectorAddressUpdated)
				if err := _ZetaNonEth.contract.UnpackLog(event, "ConnectorAddressUpdated", log); err != nil {
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

// ParseConnectorAddressUpdated is a log parse operation binding the contract event 0x1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c.
//
// Solidity: event ConnectorAddressUpdated(address callerAddress, address newConnectorAddress)
func (_ZetaNonEth *ZetaNonEthFilterer) ParseConnectorAddressUpdated(log types.Log) (*ZetaNonEthConnectorAddressUpdated, error) {
	event := new(ZetaNonEthConnectorAddressUpdated)
	if err := _ZetaNonEth.contract.UnpackLog(event, "ConnectorAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaNonEthMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the ZetaNonEth contract.
type ZetaNonEthMintedIterator struct {
	Event *ZetaNonEthMinted // Event containing the contract specifics and raw log

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
func (it *ZetaNonEthMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaNonEthMinted)
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
		it.Event = new(ZetaNonEthMinted)
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
func (it *ZetaNonEthMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaNonEthMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaNonEthMinted represents a Minted event raised by the ZetaNonEth contract.
type ZetaNonEthMinted struct {
	Mintee           common.Address
	Amount           *big.Int
	InternalSendHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0xc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb.
//
// Solidity: event Minted(address indexed mintee, uint256 amount, bytes32 indexed internalSendHash)
func (_ZetaNonEth *ZetaNonEthFilterer) FilterMinted(opts *bind.FilterOpts, mintee []common.Address, internalSendHash [][32]byte) (*ZetaNonEthMintedIterator, error) {

	var minteeRule []interface{}
	for _, minteeItem := range mintee {
		minteeRule = append(minteeRule, minteeItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaNonEth.contract.FilterLogs(opts, "Minted", minteeRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthMintedIterator{contract: _ZetaNonEth.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0xc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb.
//
// Solidity: event Minted(address indexed mintee, uint256 amount, bytes32 indexed internalSendHash)
func (_ZetaNonEth *ZetaNonEthFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *ZetaNonEthMinted, mintee []common.Address, internalSendHash [][32]byte) (event.Subscription, error) {

	var minteeRule []interface{}
	for _, minteeItem := range mintee {
		minteeRule = append(minteeRule, minteeItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaNonEth.contract.WatchLogs(opts, "Minted", minteeRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaNonEthMinted)
				if err := _ZetaNonEth.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0xc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb.
//
// Solidity: event Minted(address indexed mintee, uint256 amount, bytes32 indexed internalSendHash)
func (_ZetaNonEth *ZetaNonEthFilterer) ParseMinted(log types.Log) (*ZetaNonEthMinted, error) {
	event := new(ZetaNonEthMinted)
	if err := _ZetaNonEth.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaNonEthTSSAddressUpdatedIterator is returned from FilterTSSAddressUpdated and is used to iterate over the raw logs and unpacked data for TSSAddressUpdated events raised by the ZetaNonEth contract.
type ZetaNonEthTSSAddressUpdatedIterator struct {
	Event *ZetaNonEthTSSAddressUpdated // Event containing the contract specifics and raw log

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
func (it *ZetaNonEthTSSAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaNonEthTSSAddressUpdated)
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
		it.Event = new(ZetaNonEthTSSAddressUpdated)
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
func (it *ZetaNonEthTSSAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaNonEthTSSAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaNonEthTSSAddressUpdated represents a TSSAddressUpdated event raised by the ZetaNonEth contract.
type ZetaNonEthTSSAddressUpdated struct {
	CallerAddress common.Address
	NewTssAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTSSAddressUpdated is a free log retrieval operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaNonEth *ZetaNonEthFilterer) FilterTSSAddressUpdated(opts *bind.FilterOpts) (*ZetaNonEthTSSAddressUpdatedIterator, error) {

	logs, sub, err := _ZetaNonEth.contract.FilterLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthTSSAddressUpdatedIterator{contract: _ZetaNonEth.contract, event: "TSSAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchTSSAddressUpdated is a free log subscription operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaNonEth *ZetaNonEthFilterer) WatchTSSAddressUpdated(opts *bind.WatchOpts, sink chan<- *ZetaNonEthTSSAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaNonEth.contract.WatchLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaNonEthTSSAddressUpdated)
				if err := _ZetaNonEth.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
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
func (_ZetaNonEth *ZetaNonEthFilterer) ParseTSSAddressUpdated(log types.Log) (*ZetaNonEthTSSAddressUpdated, error) {
	event := new(ZetaNonEthTSSAddressUpdated)
	if err := _ZetaNonEth.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaNonEthTSSAddressUpdaterUpdatedIterator is returned from FilterTSSAddressUpdaterUpdated and is used to iterate over the raw logs and unpacked data for TSSAddressUpdaterUpdated events raised by the ZetaNonEth contract.
type ZetaNonEthTSSAddressUpdaterUpdatedIterator struct {
	Event *ZetaNonEthTSSAddressUpdaterUpdated // Event containing the contract specifics and raw log

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
func (it *ZetaNonEthTSSAddressUpdaterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaNonEthTSSAddressUpdaterUpdated)
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
		it.Event = new(ZetaNonEthTSSAddressUpdaterUpdated)
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
func (it *ZetaNonEthTSSAddressUpdaterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaNonEthTSSAddressUpdaterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaNonEthTSSAddressUpdaterUpdated represents a TSSAddressUpdaterUpdated event raised by the ZetaNonEth contract.
type ZetaNonEthTSSAddressUpdaterUpdated struct {
	CallerAddress        common.Address
	NewTssUpdaterAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTSSAddressUpdaterUpdated is a free log retrieval operation binding the contract event 0x5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0.
//
// Solidity: event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress)
func (_ZetaNonEth *ZetaNonEthFilterer) FilterTSSAddressUpdaterUpdated(opts *bind.FilterOpts) (*ZetaNonEthTSSAddressUpdaterUpdatedIterator, error) {

	logs, sub, err := _ZetaNonEth.contract.FilterLogs(opts, "TSSAddressUpdaterUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthTSSAddressUpdaterUpdatedIterator{contract: _ZetaNonEth.contract, event: "TSSAddressUpdaterUpdated", logs: logs, sub: sub}, nil
}

// WatchTSSAddressUpdaterUpdated is a free log subscription operation binding the contract event 0x5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0.
//
// Solidity: event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress)
func (_ZetaNonEth *ZetaNonEthFilterer) WatchTSSAddressUpdaterUpdated(opts *bind.WatchOpts, sink chan<- *ZetaNonEthTSSAddressUpdaterUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaNonEth.contract.WatchLogs(opts, "TSSAddressUpdaterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaNonEthTSSAddressUpdaterUpdated)
				if err := _ZetaNonEth.contract.UnpackLog(event, "TSSAddressUpdaterUpdated", log); err != nil {
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
func (_ZetaNonEth *ZetaNonEthFilterer) ParseTSSAddressUpdaterUpdated(log types.Log) (*ZetaNonEthTSSAddressUpdaterUpdated, error) {
	event := new(ZetaNonEthTSSAddressUpdaterUpdated)
	if err := _ZetaNonEth.contract.UnpackLog(event, "TSSAddressUpdaterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaNonEthTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZetaNonEth contract.
type ZetaNonEthTransferIterator struct {
	Event *ZetaNonEthTransfer // Event containing the contract specifics and raw log

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
func (it *ZetaNonEthTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaNonEthTransfer)
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
		it.Event = new(ZetaNonEthTransfer)
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
func (it *ZetaNonEthTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaNonEthTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaNonEthTransfer represents a Transfer event raised by the ZetaNonEth contract.
type ZetaNonEthTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZetaNonEth *ZetaNonEthFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZetaNonEthTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaNonEth.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaNonEthTransferIterator{contract: _ZetaNonEth.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZetaNonEth *ZetaNonEthFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZetaNonEthTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaNonEth.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaNonEthTransfer)
				if err := _ZetaNonEth.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ZetaNonEth *ZetaNonEthFilterer) ParseTransfer(log types.Log) (*ZetaNonEthTransfer, error) {
	event := new(ZetaNonEthTransfer)
	if err := _ZetaNonEth.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
