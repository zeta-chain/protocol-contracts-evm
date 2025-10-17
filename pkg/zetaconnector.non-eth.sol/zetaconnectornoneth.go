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

// ZetaConnectorNonEthMetaData contains all meta data concerning the ZetaConnectorNonEth contract.
var ZetaConnectorNonEthMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"zetaTokenAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tssAddressUpdater_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pauserAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLockedAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReceive\",\"inputs\":[{\"name\":\"zetaTxSenderAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zetaValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRevert\",\"inputs\":[{\"name\":\"zetaTxSenderAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"remainingZetaValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceTssAddressUpdater\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"send\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structZetaInterfaces.SendInput\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destinationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"zetaValueAndGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"zetaParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxSupply\",\"inputs\":[{\"name\":\"maxSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssAddressUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePauserAddress\",\"inputs\":[{\"name\":\"pauserAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTssAddress\",\"inputs\":[{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MaxSupplyUpdated\",\"inputs\":[{\"name\":\"callerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newMaxSupply\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserAddressUpdated\",\"inputs\":[{\"name\":\"callerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTssAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TSSAddressUpdated\",\"inputs\":[{\"name\":\"callerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTssAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TSSAddressUpdaterUpdated\",\"inputs\":[{\"name\":\"callerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTssUpdaterAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZetaReceived\",\"inputs\":[{\"name\":\"zetaTxSenderAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zetaValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZetaReverted\",\"inputs\":[{\"name\":\"zetaTxSenderAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"remainingZetaValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZetaSent\",\"inputs\":[{\"name\":\"sourceTxOriginAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"zetaTxSenderAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zetaValueAndGas\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationGasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zetaParams\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotPauser\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTss\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTssOrUpdater\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTssUpdater\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsMaxSupply\",\"inputs\":[{\"name\":\"maxSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZetaTransferError\",\"inputs\":[]}]",
	Bin: "0x60a060405260001960035534801561001657600080fd5b506040516119323803806119328339810160408190526100359161011a565b6000805460ff19169055838383836001600160a01b038416158061006057506001600160a01b038316155b8061007257506001600160a01b038216155b8061008457506001600160a01b038116155b156100a25760405163e6c4247b60e01b815260040160405180910390fd5b6001600160a01b03938416608052600180549385166001600160a01b0319948516179055600280549285169290931691909117909155600080549190921661010002610100600160a81b03199091161790555061016e92505050565b80516001600160a01b038116811461011557600080fd5b919050565b6000806000806080858703121561013057600080fd5b610139856100fe565b9350610147602086016100fe565b9250610155604086016100fe565b9150610163606086016100fe565b905092959194509250565b60805161177f6101b36000396000818161010a015281816102ad0152818161038f015281816104b701528181610bf201528181610d1a0152610f44015261177f6000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80636f8b44b011610097578063942a5e1611610066578063942a5e1614610228578063d5abeb011461023b578063ec02690114610244578063f7fb869b1461025757600080fd5b80636f8b44b0146101f2578063779e3b63146102055780638456cb591461020d5780639122c3441461021557600080fd5b80633f4ba83a116100d35780633f4ba83a146101a15780635b112591146101a95780635c975abb146101c95780636128480f146101df57600080fd5b806321e093b114610105578063252bc8861461015657806329dd214d1461016c578063328a01d014610181575b600080fd5b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61015e61027c565b60405190815260200161014d565b61017f61017a3660046111eb565b610332565b005b60025461012c9073ffffffffffffffffffffffffffffffffffffffff1681565b61017f6106b0565b60015461012c9073ffffffffffffffffffffffffffffffffffffffff1681565b60005460ff16604051901515815260200161014d565b61017f6101ed36600461128b565b610712565b61017f6102003660046112ad565b61083f565b61017f6108cd565b61017f6109f2565b61017f61022336600461128b565b610a52565b61017f6102363660046112c6565b610b92565b61015e60035481565b61017f610252366004611370565b610f04565b60005461012c90610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610309573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061032d91906113ab565b905090565b60015473ffffffffffffffffffffffffffffffffffffffff16331461038a576040517fff70ace20000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6003547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061041c91906113ab565b61042690866113c4565b1115610464576003546040517f3d3dbc8300000000000000000000000000000000000000000000000000000000815260040161038191815260200190565b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201869052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b1580156104fb57600080fd5b505af115801561050f573d6000803e3d6000fd5b50508315915061064e9050578473ffffffffffffffffffffffffffffffffffffffff16633749c51a6040518060a001604052808b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060208082018b905273ffffffffffffffffffffffffffffffffffffffff8a16604080840191909152606083018a90528051601f890183900483028101830190915287815260809092019190889088908190840183828082843760009201919091525050509152506040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815261061b9190600401611468565b600060405180830381600087803b15801561063557600080fd5b505af1158015610649573d6000803e3d6000fd5b505050505b808573ffffffffffffffffffffffffffffffffffffffff16877ff1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d6988b8b89898960405161069e959493929190611545565b60405180910390a45050505050505050565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610708576040517f4677a0d3000000000000000000000000000000000000000000000000000000008152336004820152602401610381565b61071061102d565b565b600054610100900473ffffffffffffffffffffffffffffffffffffffff16331461076a576040517f4677a0d3000000000000000000000000000000000000000000000000000000008152336004820152602401610381565b73ffffffffffffffffffffffffffffffffffffffff81166107b7576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040805133815260208101929092527fd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d039791015b60405180910390a150565b60015473ffffffffffffffffffffffffffffffffffffffff163314610892576040517fff70ace2000000000000000000000000000000000000000000000000000000008152336004820152602401610381565b600381905560408051338152602081018390527f26843c619c87f9021bcc4ec5143177198076d9da3c13ce1bb2e941644c69d42e9101610834565b60025473ffffffffffffffffffffffffffffffffffffffff163314610920576040517fe700765e000000000000000000000000000000000000000000000000000000008152336004820152602401610381565b60015473ffffffffffffffffffffffffffffffffffffffff1661096f576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600154600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd091015b60405180910390a1565b600054610100900473ffffffffffffffffffffffffffffffffffffffff163314610a4a576040517f4677a0d3000000000000000000000000000000000000000000000000000000008152336004820152602401610381565b6107106110a5565b60015473ffffffffffffffffffffffffffffffffffffffff163314801590610a92575060025473ffffffffffffffffffffffffffffffffffffffff163314155b15610acb576040517fcdfcef97000000000000000000000000000000000000000000000000000000008152336004820152602401610381565b73ffffffffffffffffffffffffffffffffffffffff8116610b18576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff9101610834565b610b9a611100565b60015473ffffffffffffffffffffffffffffffffffffffff163314610bed576040517fff70ace2000000000000000000000000000000000000000000000000000000008152336004820152602401610381565b6003547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c7f91906113ab565b610c8990866113c4565b1115610cc7576003546040517f3d3dbc8300000000000000000000000000000000000000000000000000000000815260040161038191815260200190565b6040517f1e458bee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a8116600483015260248201869052604482018390527f00000000000000000000000000000000000000000000000000000000000000001690631e458bee90606401600060405180830381600087803b158015610d5e57600080fd5b505af1158015610d72573d6000803e3d6000fd5b505083159150610eb49050578873ffffffffffffffffffffffffffffffffffffffff16633ff0693c6040518060c001604052808c73ffffffffffffffffffffffffffffffffffffffff1681526020018b81526020018a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060208082018a905260408083018a90528051601f890183900483028101830190915287815260609092019190889088908190840183828082843760009201919091525050509152506040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b168152610e81919060040161157e565b600060405180830381600087803b158015610e9b57600080fd5b505af1158015610eaf573d6000803e3d6000fd5b505050505b80857f521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c888b8b8b8b8a8a8a604051610ef19796959493929190611613565b60405180910390a3505050505050505050565b610f0c611100565b6040517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152608082013560248201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906379cc679090604401600060405180830381600087803b158015610f9d57600080fd5b505af1158015610fb1573d6000803e3d6000fd5b5050823591503390507f7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e432610fe96020860186611670565b6080870135604088013561100060608a018a611670565b61100d60a08c018c611670565b604051611022999897969594939291906116d5565b60405180910390a350565b61103561113d565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016109e8565b6110ad611100565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586110803390565b60005460ff1615610710576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005460ff16610710576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008083601f84011261118b57600080fd5b50813567ffffffffffffffff8111156111a357600080fd5b6020830191508360208285010111156111bb57600080fd5b9250929050565b803573ffffffffffffffffffffffffffffffffffffffff811681146111e657600080fd5b919050565b60008060008060008060008060c0898b03121561120757600080fd5b883567ffffffffffffffff81111561121e57600080fd5b61122a8b828c01611179565b9099509750506020890135955061124360408a016111c2565b945060608901359350608089013567ffffffffffffffff81111561126657600080fd5b6112728b828c01611179565b999c989b50969995989497949560a00135949350505050565b60006020828403121561129d57600080fd5b6112a6826111c2565b9392505050565b6000602082840312156112bf57600080fd5b5035919050565b600080600080600080600080600060e08a8c0312156112e457600080fd5b6112ed8a6111c2565b985060208a0135975060408a013567ffffffffffffffff81111561131057600080fd5b61131c8c828d01611179565b90985096505060608a0135945060808a0135935060a08a013567ffffffffffffffff81111561134a57600080fd5b6113568c828d01611179565b9a9d999c50979a9699959894979660c00135949350505050565b60006020828403121561138257600080fd5b813567ffffffffffffffff81111561139957600080fd5b820160c081850312156112a657600080fd5b6000602082840312156113bd57600080fd5b5051919050565b808201808211156113fe577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b6000815180845260005b8181101561142a5760208185018101518683018201520161140e565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000825160a0602084015261148460c0840182611404565b90506020840151604084015273ffffffffffffffffffffffffffffffffffffffff60408501511660608401526060840151608084015260808401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160a08501526114f38282611404565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6060815260006115596060830187896114fc565b85602084015282810360408401526115728185876114fc565b98975050505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff8251166020820152602082015160408201526000604083015160c060608401526115c460e0840182611404565b905060608401516080840152608084015160a084015260a08401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160c08501526114f38282611404565b73ffffffffffffffffffffffffffffffffffffffff8816815286602082015260a06040820152600061164960a0830187896114fc565b85606084015282810360808401526116628185876114fc565b9a9950505050505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126116a557600080fd5b83018035915067ffffffffffffffff8211156116c057600080fd5b6020019150368190038213156111bb57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8a16815260c06020820152600061170560c083018a8c6114fc565b88604084015287606084015282810360808401526117248187896114fc565b905082810360a08401526117398185876114fc565b9c9b50505050505050505050505056fea2646970667358221220fda195e75ca152d19cdb4dd5afe156628718c9356b9a62990e267bfe8613817e64736f6c634300081a0033",
}

// ZetaConnectorNonEthABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNonEthMetaData.ABI instead.
var ZetaConnectorNonEthABI = ZetaConnectorNonEthMetaData.ABI

// ZetaConnectorNonEthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNonEthMetaData.Bin instead.
var ZetaConnectorNonEthBin = ZetaConnectorNonEthMetaData.Bin

// DeployZetaConnectorNonEth deploys a new Ethereum contract, binding an instance of ZetaConnectorNonEth to it.
func DeployZetaConnectorNonEth(auth *bind.TransactOpts, backend bind.ContractBackend, zetaTokenAddress_ common.Address, tssAddress_ common.Address, tssAddressUpdater_ common.Address, pauserAddress_ common.Address) (common.Address, *types.Transaction, *ZetaConnectorNonEth, error) {
	parsed, err := ZetaConnectorNonEthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNonEthBin), backend, zetaTokenAddress_, tssAddress_, tssAddressUpdater_, pauserAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorNonEth{ZetaConnectorNonEthCaller: ZetaConnectorNonEthCaller{contract: contract}, ZetaConnectorNonEthTransactor: ZetaConnectorNonEthTransactor{contract: contract}, ZetaConnectorNonEthFilterer: ZetaConnectorNonEthFilterer{contract: contract}}, nil
}

// ZetaConnectorNonEth is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNonEth struct {
	ZetaConnectorNonEthCaller     // Read-only binding to the contract
	ZetaConnectorNonEthTransactor // Write-only binding to the contract
	ZetaConnectorNonEthFilterer   // Log filterer for contract events
}

// ZetaConnectorNonEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNonEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNonEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNonEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNonEthSession struct {
	Contract     *ZetaConnectorNonEth // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZetaConnectorNonEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNonEthCallerSession struct {
	Contract *ZetaConnectorNonEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ZetaConnectorNonEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNonEthTransactorSession struct {
	Contract     *ZetaConnectorNonEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ZetaConnectorNonEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNonEthRaw struct {
	Contract *ZetaConnectorNonEth // Generic contract binding to access the raw methods on
}

// ZetaConnectorNonEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNonEthCallerRaw struct {
	Contract *ZetaConnectorNonEthCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNonEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNonEthTransactorRaw struct {
	Contract *ZetaConnectorNonEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNonEth creates a new instance of ZetaConnectorNonEth, bound to a specific deployed contract.
func NewZetaConnectorNonEth(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNonEth, error) {
	contract, err := bindZetaConnectorNonEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonEth{ZetaConnectorNonEthCaller: ZetaConnectorNonEthCaller{contract: contract}, ZetaConnectorNonEthTransactor: ZetaConnectorNonEthTransactor{contract: contract}, ZetaConnectorNonEthFilterer: ZetaConnectorNonEthFilterer{contract: contract}}, nil
}

// NewZetaConnectorNonEthCaller creates a new read-only instance of ZetaConnectorNonEth, bound to a specific deployed contract.
func NewZetaConnectorNonEthCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNonEthCaller, error) {
	contract, err := bindZetaConnectorNonEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonEthCaller{contract: contract}, nil
}

// NewZetaConnectorNonEthTransactor creates a new write-only instance of ZetaConnectorNonEth, bound to a specific deployed contract.
func NewZetaConnectorNonEthTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNonEthTransactor, error) {
	contract, err := bindZetaConnectorNonEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonEthTransactor{contract: contract}, nil
}

// NewZetaConnectorNonEthFilterer creates a new log filterer instance of ZetaConnectorNonEth, bound to a specific deployed contract.
func NewZetaConnectorNonEthFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNonEthFilterer, error) {
	contract, err := bindZetaConnectorNonEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonEthFilterer{contract: contract}, nil
}

// bindZetaConnectorNonEth binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNonEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNonEthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNonEth *ZetaConnectorNonEthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNonEth.Contract.ZetaConnectorNonEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNonEth *ZetaConnectorNonEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.ZetaConnectorNonEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNonEth *ZetaConnectorNonEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.ZetaConnectorNonEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNonEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.contract.Transact(opts, method, params...)
}

// GetLockedAmount is a free data retrieval call binding the contract method 0x252bc886.
//
// Solidity: function getLockedAmount() view returns(uint256)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCaller) GetLockedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZetaConnectorNonEth.contract.Call(opts, &out, "getLockedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedAmount is a free data retrieval call binding the contract method 0x252bc886.
//
// Solidity: function getLockedAmount() view returns(uint256)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) GetLockedAmount() (*big.Int, error) {
	return _ZetaConnectorNonEth.Contract.GetLockedAmount(&_ZetaConnectorNonEth.CallOpts)
}

// GetLockedAmount is a free data retrieval call binding the contract method 0x252bc886.
//
// Solidity: function getLockedAmount() view returns(uint256)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCallerSession) GetLockedAmount() (*big.Int, error) {
	return _ZetaConnectorNonEth.Contract.GetLockedAmount(&_ZetaConnectorNonEth.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZetaConnectorNonEth.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) MaxSupply() (*big.Int, error) {
	return _ZetaConnectorNonEth.Contract.MaxSupply(&_ZetaConnectorNonEth.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCallerSession) MaxSupply() (*big.Int, error) {
	return _ZetaConnectorNonEth.Contract.MaxSupply(&_ZetaConnectorNonEth.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNonEth.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) Paused() (bool, error) {
	return _ZetaConnectorNonEth.Contract.Paused(&_ZetaConnectorNonEth.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCallerSession) Paused() (bool, error) {
	return _ZetaConnectorNonEth.Contract.Paused(&_ZetaConnectorNonEth.CallOpts)
}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCaller) PauserAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonEth.contract.Call(opts, &out, "pauserAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) PauserAddress() (common.Address, error) {
	return _ZetaConnectorNonEth.Contract.PauserAddress(&_ZetaConnectorNonEth.CallOpts)
}

// PauserAddress is a free data retrieval call binding the contract method 0xf7fb869b.
//
// Solidity: function pauserAddress() view returns(address)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCallerSession) PauserAddress() (common.Address, error) {
	return _ZetaConnectorNonEth.Contract.PauserAddress(&_ZetaConnectorNonEth.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonEth.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNonEth.Contract.TssAddress(&_ZetaConnectorNonEth.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCallerSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNonEth.Contract.TssAddress(&_ZetaConnectorNonEth.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCaller) TssAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonEth.contract.Call(opts, &out, "tssAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) TssAddressUpdater() (common.Address, error) {
	return _ZetaConnectorNonEth.Contract.TssAddressUpdater(&_ZetaConnectorNonEth.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCallerSession) TssAddressUpdater() (common.Address, error) {
	return _ZetaConnectorNonEth.Contract.TssAddressUpdater(&_ZetaConnectorNonEth.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonEth.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNonEth.Contract.ZetaToken(&_ZetaConnectorNonEth.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNonEth.Contract.ZetaToken(&_ZetaConnectorNonEth.CallOpts)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactor) OnReceive(opts *bind.TransactOpts, zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.contract.Transact(opts, "onReceive", zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) OnReceive(zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.OnReceive(&_ZetaConnectorNonEth.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes zetaTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 zetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactorSession) OnReceive(zetaTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, zetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.OnReceive(&_ZetaConnectorNonEth.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, zetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactor) OnRevert(opts *bind.TransactOpts, zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.contract.Transact(opts, "onRevert", zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) OnRevert(zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.OnRevert(&_ZetaConnectorNonEth.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address zetaTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingZetaValue, bytes message, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactorSession) OnRevert(zetaTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingZetaValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.OnRevert(&_ZetaConnectorNonEth.TransactOpts, zetaTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingZetaValue, message, internalSendHash)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.Pause(&_ZetaConnectorNonEth.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactorSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.Pause(&_ZetaConnectorNonEth.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactor) RenounceTssAddressUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.contract.Transact(opts, "renounceTssAddressUpdater")
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.RenounceTssAddressUpdater(&_ZetaConnectorNonEth.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactorSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.RenounceTssAddressUpdater(&_ZetaConnectorNonEth.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactor) Send(opts *bind.TransactOpts, input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) Send(input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.Send(&_ZetaConnectorNonEth.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactorSession) Send(input ZetaInterfacesSendInput) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.Send(&_ZetaConnectorNonEth.TransactOpts, input)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactor) SetMaxSupply(opts *bind.TransactOpts, maxSupply_ *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.contract.Transact(opts, "setMaxSupply", maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.SetMaxSupply(&_ZetaConnectorNonEth.TransactOpts, maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactorSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.SetMaxSupply(&_ZetaConnectorNonEth.TransactOpts, maxSupply_)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.Unpause(&_ZetaConnectorNonEth.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactorSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.Unpause(&_ZetaConnectorNonEth.TransactOpts)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactor) UpdatePauserAddress(opts *bind.TransactOpts, pauserAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.contract.Transact(opts, "updatePauserAddress", pauserAddress_)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) UpdatePauserAddress(pauserAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.UpdatePauserAddress(&_ZetaConnectorNonEth.TransactOpts, pauserAddress_)
}

// UpdatePauserAddress is a paid mutator transaction binding the contract method 0x6128480f.
//
// Solidity: function updatePauserAddress(address pauserAddress_) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactorSession) UpdatePauserAddress(pauserAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.UpdatePauserAddress(&_ZetaConnectorNonEth.TransactOpts, pauserAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactor) UpdateTssAddress(opts *bind.TransactOpts, tssAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.contract.Transact(opts, "updateTssAddress", tssAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthSession) UpdateTssAddress(tssAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.UpdateTssAddress(&_ZetaConnectorNonEth.TransactOpts, tssAddress_)
}

// UpdateTssAddress is a paid mutator transaction binding the contract method 0x9122c344.
//
// Solidity: function updateTssAddress(address tssAddress_) returns()
func (_ZetaConnectorNonEth *ZetaConnectorNonEthTransactorSession) UpdateTssAddress(tssAddress_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonEth.Contract.UpdateTssAddress(&_ZetaConnectorNonEth.TransactOpts, tssAddress_)
}

// ZetaConnectorNonEthMaxSupplyUpdatedIterator is returned from FilterMaxSupplyUpdated and is used to iterate over the raw logs and unpacked data for MaxSupplyUpdated events raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthMaxSupplyUpdatedIterator struct {
	Event *ZetaConnectorNonEthMaxSupplyUpdated // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonEthMaxSupplyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonEthMaxSupplyUpdated)
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
		it.Event = new(ZetaConnectorNonEthMaxSupplyUpdated)
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
func (it *ZetaConnectorNonEthMaxSupplyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonEthMaxSupplyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonEthMaxSupplyUpdated represents a MaxSupplyUpdated event raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthMaxSupplyUpdated struct {
	CallerAddress common.Address
	NewMaxSupply  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyUpdated is a free log retrieval operation binding the contract event 0x26843c619c87f9021bcc4ec5143177198076d9da3c13ce1bb2e941644c69d42e.
//
// Solidity: event MaxSupplyUpdated(address callerAddress, uint256 newMaxSupply)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) FilterMaxSupplyUpdated(opts *bind.FilterOpts) (*ZetaConnectorNonEthMaxSupplyUpdatedIterator, error) {

	logs, sub, err := _ZetaConnectorNonEth.contract.FilterLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonEthMaxSupplyUpdatedIterator{contract: _ZetaConnectorNonEth.contract, event: "MaxSupplyUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyUpdated is a free log subscription operation binding the contract event 0x26843c619c87f9021bcc4ec5143177198076d9da3c13ce1bb2e941644c69d42e.
//
// Solidity: event MaxSupplyUpdated(address callerAddress, uint256 newMaxSupply)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) WatchMaxSupplyUpdated(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonEthMaxSupplyUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonEth.contract.WatchLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonEthMaxSupplyUpdated)
				if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
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

// ParseMaxSupplyUpdated is a log parse operation binding the contract event 0x26843c619c87f9021bcc4ec5143177198076d9da3c13ce1bb2e941644c69d42e.
//
// Solidity: event MaxSupplyUpdated(address callerAddress, uint256 newMaxSupply)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) ParseMaxSupplyUpdated(log types.Log) (*ZetaConnectorNonEthMaxSupplyUpdated, error) {
	event := new(ZetaConnectorNonEthMaxSupplyUpdated)
	if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonEthPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthPausedIterator struct {
	Event *ZetaConnectorNonEthPaused // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonEthPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonEthPaused)
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
		it.Event = new(ZetaConnectorNonEthPaused)
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
func (it *ZetaConnectorNonEthPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonEthPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonEthPaused represents a Paused event raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) FilterPaused(opts *bind.FilterOpts) (*ZetaConnectorNonEthPausedIterator, error) {

	logs, sub, err := _ZetaConnectorNonEth.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonEthPausedIterator{contract: _ZetaConnectorNonEth.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonEthPaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonEth.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonEthPaused)
				if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) ParsePaused(log types.Log) (*ZetaConnectorNonEthPaused, error) {
	event := new(ZetaConnectorNonEthPaused)
	if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonEthPauserAddressUpdatedIterator is returned from FilterPauserAddressUpdated and is used to iterate over the raw logs and unpacked data for PauserAddressUpdated events raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthPauserAddressUpdatedIterator struct {
	Event *ZetaConnectorNonEthPauserAddressUpdated // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonEthPauserAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonEthPauserAddressUpdated)
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
		it.Event = new(ZetaConnectorNonEthPauserAddressUpdated)
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
func (it *ZetaConnectorNonEthPauserAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonEthPauserAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonEthPauserAddressUpdated represents a PauserAddressUpdated event raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthPauserAddressUpdated struct {
	CallerAddress common.Address
	NewTssAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPauserAddressUpdated is a free log retrieval operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) FilterPauserAddressUpdated(opts *bind.FilterOpts) (*ZetaConnectorNonEthPauserAddressUpdatedIterator, error) {

	logs, sub, err := _ZetaConnectorNonEth.contract.FilterLogs(opts, "PauserAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonEthPauserAddressUpdatedIterator{contract: _ZetaConnectorNonEth.contract, event: "PauserAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchPauserAddressUpdated is a free log subscription operation binding the contract event 0xd41d83655d484bdf299598751c371b2d92088667266fe3774b25a97bdd5d0397.
//
// Solidity: event PauserAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) WatchPauserAddressUpdated(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonEthPauserAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonEth.contract.WatchLogs(opts, "PauserAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonEthPauserAddressUpdated)
				if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "PauserAddressUpdated", log); err != nil {
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
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) ParsePauserAddressUpdated(log types.Log) (*ZetaConnectorNonEthPauserAddressUpdated, error) {
	event := new(ZetaConnectorNonEthPauserAddressUpdated)
	if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "PauserAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonEthTSSAddressUpdatedIterator is returned from FilterTSSAddressUpdated and is used to iterate over the raw logs and unpacked data for TSSAddressUpdated events raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthTSSAddressUpdatedIterator struct {
	Event *ZetaConnectorNonEthTSSAddressUpdated // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonEthTSSAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonEthTSSAddressUpdated)
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
		it.Event = new(ZetaConnectorNonEthTSSAddressUpdated)
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
func (it *ZetaConnectorNonEthTSSAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonEthTSSAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonEthTSSAddressUpdated represents a TSSAddressUpdated event raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthTSSAddressUpdated struct {
	CallerAddress common.Address
	NewTssAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTSSAddressUpdated is a free log retrieval operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) FilterTSSAddressUpdated(opts *bind.FilterOpts) (*ZetaConnectorNonEthTSSAddressUpdatedIterator, error) {

	logs, sub, err := _ZetaConnectorNonEth.contract.FilterLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonEthTSSAddressUpdatedIterator{contract: _ZetaConnectorNonEth.contract, event: "TSSAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchTSSAddressUpdated is a free log subscription operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address callerAddress, address newTssAddress)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) WatchTSSAddressUpdated(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonEthTSSAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonEth.contract.WatchLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonEthTSSAddressUpdated)
				if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
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
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) ParseTSSAddressUpdated(log types.Log) (*ZetaConnectorNonEthTSSAddressUpdated, error) {
	event := new(ZetaConnectorNonEthTSSAddressUpdated)
	if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonEthTSSAddressUpdaterUpdatedIterator is returned from FilterTSSAddressUpdaterUpdated and is used to iterate over the raw logs and unpacked data for TSSAddressUpdaterUpdated events raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthTSSAddressUpdaterUpdatedIterator struct {
	Event *ZetaConnectorNonEthTSSAddressUpdaterUpdated // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonEthTSSAddressUpdaterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonEthTSSAddressUpdaterUpdated)
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
		it.Event = new(ZetaConnectorNonEthTSSAddressUpdaterUpdated)
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
func (it *ZetaConnectorNonEthTSSAddressUpdaterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonEthTSSAddressUpdaterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonEthTSSAddressUpdaterUpdated represents a TSSAddressUpdaterUpdated event raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthTSSAddressUpdaterUpdated struct {
	CallerAddress        common.Address
	NewTssUpdaterAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTSSAddressUpdaterUpdated is a free log retrieval operation binding the contract event 0x5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0.
//
// Solidity: event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) FilterTSSAddressUpdaterUpdated(opts *bind.FilterOpts) (*ZetaConnectorNonEthTSSAddressUpdaterUpdatedIterator, error) {

	logs, sub, err := _ZetaConnectorNonEth.contract.FilterLogs(opts, "TSSAddressUpdaterUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonEthTSSAddressUpdaterUpdatedIterator{contract: _ZetaConnectorNonEth.contract, event: "TSSAddressUpdaterUpdated", logs: logs, sub: sub}, nil
}

// WatchTSSAddressUpdaterUpdated is a free log subscription operation binding the contract event 0x5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0.
//
// Solidity: event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) WatchTSSAddressUpdaterUpdated(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonEthTSSAddressUpdaterUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonEth.contract.WatchLogs(opts, "TSSAddressUpdaterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonEthTSSAddressUpdaterUpdated)
				if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "TSSAddressUpdaterUpdated", log); err != nil {
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
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) ParseTSSAddressUpdaterUpdated(log types.Log) (*ZetaConnectorNonEthTSSAddressUpdaterUpdated, error) {
	event := new(ZetaConnectorNonEthTSSAddressUpdaterUpdated)
	if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "TSSAddressUpdaterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonEthUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthUnpausedIterator struct {
	Event *ZetaConnectorNonEthUnpaused // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonEthUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonEthUnpaused)
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
		it.Event = new(ZetaConnectorNonEthUnpaused)
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
func (it *ZetaConnectorNonEthUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonEthUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonEthUnpaused represents a Unpaused event raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ZetaConnectorNonEthUnpausedIterator, error) {

	logs, sub, err := _ZetaConnectorNonEth.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonEthUnpausedIterator{contract: _ZetaConnectorNonEth.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonEthUnpaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonEth.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonEthUnpaused)
				if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) ParseUnpaused(log types.Log) (*ZetaConnectorNonEthUnpaused, error) {
	event := new(ZetaConnectorNonEthUnpaused)
	if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonEthZetaReceivedIterator is returned from FilterZetaReceived and is used to iterate over the raw logs and unpacked data for ZetaReceived events raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthZetaReceivedIterator struct {
	Event *ZetaConnectorNonEthZetaReceived // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonEthZetaReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonEthZetaReceived)
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
		it.Event = new(ZetaConnectorNonEthZetaReceived)
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
func (it *ZetaConnectorNonEthZetaReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonEthZetaReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonEthZetaReceived represents a ZetaReceived event raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthZetaReceived struct {
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
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) FilterZetaReceived(opts *bind.FilterOpts, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (*ZetaConnectorNonEthZetaReceivedIterator, error) {

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

	logs, sub, err := _ZetaConnectorNonEth.contract.FilterLogs(opts, "ZetaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonEthZetaReceivedIterator{contract: _ZetaConnectorNonEth.contract, event: "ZetaReceived", logs: logs, sub: sub}, nil
}

// WatchZetaReceived is a free log subscription operation binding the contract event 0xf1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d698.
//
// Solidity: event ZetaReceived(bytes zetaTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 zetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) WatchZetaReceived(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonEthZetaReceived, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ZetaConnectorNonEth.contract.WatchLogs(opts, "ZetaReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonEthZetaReceived)
				if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "ZetaReceived", log); err != nil {
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
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) ParseZetaReceived(log types.Log) (*ZetaConnectorNonEthZetaReceived, error) {
	event := new(ZetaConnectorNonEthZetaReceived)
	if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "ZetaReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonEthZetaRevertedIterator is returned from FilterZetaReverted and is used to iterate over the raw logs and unpacked data for ZetaReverted events raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthZetaRevertedIterator struct {
	Event *ZetaConnectorNonEthZetaReverted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonEthZetaRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonEthZetaReverted)
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
		it.Event = new(ZetaConnectorNonEthZetaReverted)
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
func (it *ZetaConnectorNonEthZetaRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonEthZetaRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonEthZetaReverted represents a ZetaReverted event raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthZetaReverted struct {
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
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) FilterZetaReverted(opts *bind.FilterOpts, destinationChainId []*big.Int, internalSendHash [][32]byte) (*ZetaConnectorNonEthZetaRevertedIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorNonEth.contract.FilterLogs(opts, "ZetaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonEthZetaRevertedIterator{contract: _ZetaConnectorNonEth.contract, event: "ZetaReverted", logs: logs, sub: sub}, nil
}

// WatchZetaReverted is a free log subscription operation binding the contract event 0x521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c88.
//
// Solidity: event ZetaReverted(address zetaTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingZetaValue, bytes message, bytes32 indexed internalSendHash)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) WatchZetaReverted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonEthZetaReverted, destinationChainId []*big.Int, internalSendHash [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _ZetaConnectorNonEth.contract.WatchLogs(opts, "ZetaReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonEthZetaReverted)
				if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "ZetaReverted", log); err != nil {
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
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) ParseZetaReverted(log types.Log) (*ZetaConnectorNonEthZetaReverted, error) {
	event := new(ZetaConnectorNonEthZetaReverted)
	if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "ZetaReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonEthZetaSentIterator is returned from FilterZetaSent and is used to iterate over the raw logs and unpacked data for ZetaSent events raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthZetaSentIterator struct {
	Event *ZetaConnectorNonEthZetaSent // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonEthZetaSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonEthZetaSent)
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
		it.Event = new(ZetaConnectorNonEthZetaSent)
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
func (it *ZetaConnectorNonEthZetaSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonEthZetaSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonEthZetaSent represents a ZetaSent event raised by the ZetaConnectorNonEth contract.
type ZetaConnectorNonEthZetaSent struct {
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
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) FilterZetaSent(opts *bind.FilterOpts, zetaTxSenderAddress []common.Address, destinationChainId []*big.Int) (*ZetaConnectorNonEthZetaSentIterator, error) {

	var zetaTxSenderAddressRule []interface{}
	for _, zetaTxSenderAddressItem := range zetaTxSenderAddress {
		zetaTxSenderAddressRule = append(zetaTxSenderAddressRule, zetaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _ZetaConnectorNonEth.contract.FilterLogs(opts, "ZetaSent", zetaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonEthZetaSentIterator{contract: _ZetaConnectorNonEth.contract, event: "ZetaSent", logs: logs, sub: sub}, nil
}

// WatchZetaSent is a free log subscription operation binding the contract event 0x7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e4.
//
// Solidity: event ZetaSent(address sourceTxOriginAddress, address indexed zetaTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 zetaValueAndGas, uint256 destinationGasLimit, bytes message, bytes zetaParams)
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) WatchZetaSent(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonEthZetaSent, zetaTxSenderAddress []common.Address, destinationChainId []*big.Int) (event.Subscription, error) {

	var zetaTxSenderAddressRule []interface{}
	for _, zetaTxSenderAddressItem := range zetaTxSenderAddress {
		zetaTxSenderAddressRule = append(zetaTxSenderAddressRule, zetaTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _ZetaConnectorNonEth.contract.WatchLogs(opts, "ZetaSent", zetaTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonEthZetaSent)
				if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "ZetaSent", log); err != nil {
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
func (_ZetaConnectorNonEth *ZetaConnectorNonEthFilterer) ParseZetaSent(log types.Log) (*ZetaConnectorNonEthZetaSent, error) {
	event := new(ZetaConnectorNonEthZetaSent)
	if err := _ZetaConnectorNonEth.contract.UnpackLog(event, "ZetaSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
