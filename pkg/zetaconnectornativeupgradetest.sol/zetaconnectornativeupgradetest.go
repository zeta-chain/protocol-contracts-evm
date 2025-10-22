// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectornativeupgradetest

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

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Sender common.Address
}

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// ZetaConnectorNativeUpgradeTestMetaData contains all meta data concerning the ZetaConnectorNativeUpgradeTest contract.
var ZetaConnectorNativeUpgradeTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"gateway_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveTokens\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"messageContext\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndRevert\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedZetaConnectorTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a060405230608052348015601357600080fd5b5060805161248e61003d60003960008181611227015281816112500152611426015261248e6000f3fe6080604052600436106101965760003560e01c80636f8728ad116100e1578063950837aa1161008a578063ad3cb1cc11610064578063ad3cb1cc146104f2578063d547741f14610548578063e63ab1e914610568578063f8c8765e1461059c57600080fd5b8063950837aa14610489578063a217fddf146104a9578063a783c789146104be57600080fd5b80638456cb59116100bb5780638456cb59146103db57806385f438c1146103f057806391d148541461042457600080fd5b80636f8728ad1461037b5780636fb9a7af1461039b578063743e0c9b146103bb57600080fd5b806336568abe1161014357806352d1902d1161011d57806352d1902d1461030f5780635b112591146103245780635c975abb1461034457600080fd5b806336568abe146102c75780633f4ba83a146102e75780634f1ef286146102fc57600080fd5b806321e093b11161017457806321e093b11461022a578063248a9ca31461024a5780632f2ff15d146102a757600080fd5b806301ffc9a71461019b578063106e6290146101d0578063116191b6146101f2575b600080fd5b3480156101a757600080fd5b506101bb6101b6366004611db4565b6105bc565b60405190151581526020015b60405180910390f35b3480156101dc57600080fd5b506101f06101eb366004611e12565b610655565b005b3480156101fe57600080fd5b50600054610212906001600160a01b031681565b6040516001600160a01b0390911681526020016101c7565b34801561023657600080fd5b50600154610212906001600160a01b031681565b34801561025657600080fd5b50610299610265366004611e45565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016101c7565b3480156102b357600080fd5b506101f06102c2366004611e5e565b610718565b3480156102d357600080fd5b506101f06102e2366004611e5e565b610762565b3480156102f357600080fd5b506101f06107ae565b6101f061030a366004611eb9565b6107e3565b34801561031b57600080fd5b50610299610802565b34801561033057600080fd5b50600254610212906001600160a01b031681565b34801561035057600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166101bb565b34801561038757600080fd5b506101f0610396366004612009565b610831565b3480156103a757600080fd5b506101f06103b63660046120a1565b61098a565b3480156103c757600080fd5b506101f06103d6366004611e45565b610aa8565b3480156103e757600080fd5b506101f0610ac8565b3480156103fc57600080fd5b506102997f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b34801561043057600080fd5b506101bb61043f366004611e5e565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561049557600080fd5b506101f06104a4366004612120565b610afa565b3480156104b557600080fd5b50610299600081565b3480156104ca57600080fd5b506102997f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b3480156104fe57600080fd5b5061053b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101c7919061215f565b34801561055457600080fd5b506101f0610563366004611e5e565b610c8d565b34801561057457600080fd5b506102997f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b3480156105a857600080fd5b506101f06105b73660046121b0565b610cd1565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061064f57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b61065d610e58565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461068781610ed9565b61068f610ee3565b6001546106a6906001600160a01b03168585610f41565b836001600160a01b03167f3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9846040516106e191815260200190565b60405180910390a25061071360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461075281610ed9565b61075c8383610fdb565b50505050565b6001600160a01b03811633146107a4576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61071382826110c8565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107d881610ed9565b6107e061118c565b50565b6107eb61121c565b6107f4826112ec565b6107fe82826112f7565b5050565b600061080c61141b565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b610839610e58565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461086381610ed9565b61086b610ee3565b600054600154610888916001600160a01b03918216911688610f41565b6000546001546040517faa0c0fc10000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263aa0c0fc1926108df929116908b908b908b908b908a906004016122fe565b600060405180830381600087803b1580156108f957600080fd5b505af115801561090d573d6000803e3d6000fd5b50505050866001600160a01b03167f5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0878787866040516109509493929190612355565b60405180910390a25061098260017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050505050565b610992610e58565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46109bc81610ed9565b6109c4610ee3565b6000546001546109e1916001600160a01b03918216911687610f41565b6000546001546040517f7bbe9afa0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692637bbe9afa92610a39928c92909116908b908b908b908b9060040161238c565b600060405180830381600087803b158015610a5357600080fd5b505af1158015610a67573d6000803e3d6000fd5b50505050856001600160a01b03167f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d868686604051610950939291906123e7565b610ab0610ee3565b6001546107e0906001600160a01b031633308461147d565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610af281610ed9565b6107e06114b6565b6000610b0581610ed9565b6001600160a01b038216610b45576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254610b7c907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4906001600160a01b03166110c8565b50600254610bb4907f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb906001600160a01b03166110c8565b50610bdf7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e483610fdb565b50610c0a7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb83610fdb565b50600254604080516001600160a01b03928316815291841660208301527f33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e910160405180910390a150600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610cc781610ed9565b61075c83836110c8565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610d1c5750825b905060008267ffffffffffffffff166001148015610d395750303b155b905081158015610d47575080155b15610d7e576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610ddf5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610deb8989898961152f565b8315610e4d5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2906020015b60405180910390a15b505050505050505050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01610ed3576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b6107e0813361183a565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615610f3f576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6040516001600160a01b0383811660248301526044820183905261071391859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506118c7565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff166110be576000848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556110743390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061064f565b600091505061064f565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16156110be576000848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061064f565b611194611943565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806112b557507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166112a97f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15610f3f576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006107fe81610ed9565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561136f575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261136c91810190612401565b60015b6113b5576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611411576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016113ac565b610713838361199e565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610f3f576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040516001600160a01b03848116602483015283811660448301526064820183905261075c9186918216906323b872dd90608401610f6e565b6114be610ee3565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336111fe565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff1660008115801561157a5750825b905060008267ffffffffffffffff1660011480156115975750303b155b9050811580156115a5575080155b156115dc576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000166001178555831561163d5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038916158061165a57506001600160a01b038816155b8061166c57506001600160a01b038716155b8061167e57506001600160a01b038616155b156116b5576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6116bd6119f4565b6116c56119fc565b6116cd6119f4565b6116d5611a0c565b600080546001600160a01b03808c167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316178355600180548c831690841617905560028054918b16919092161790556117309087610fdb565b5061175b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e488610fdb565b506117867f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb88610fdb565b506117b17f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a87610fdb565b506117dc7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a88610fdb565b508315610e4d5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602001610e44565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff166107fe576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044016113ac565b60006118dc6001600160a01b03841683611a1c565b905080516000141580156119015750808060200190518101906118ff919061241a565b155b15610713576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201526024016113ac565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610f3f576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119a782611a31565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156119ec576107138282611ad9565b6107fe611b4f565b610f3f611b87565b611a04611b87565b610f3f611bee565b611a14611b87565b610f3f611bf6565b6060611a2a83836000611c47565b9392505050565b806001600160a01b03163b600003611a80576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016113ac565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051611af6919061243c565b600060405180830381855af49150503d8060008114611b31576040519150601f19603f3d011682016040523d82523d6000602084013e611b36565b606091505b5091509150611b46858383611cfd565b95945050505050565b3415610f3f576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610f3f576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fb5611b87565b611bfe611b87565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b606081471015611c85576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024016113ac565b600080856001600160a01b03168486604051611ca1919061243c565b60006040518083038185875af1925050503d8060008114611cde576040519150601f19603f3d011682016040523d82523d6000602084013e611ce3565b606091505b5091509150611cf3868383611cfd565b9695505050505050565b606082611d1257611d0d82611d72565b611a2a565b8151158015611d2957506001600160a01b0384163b155b15611d6b576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016113ac565b5080611a2a565b805115611d825780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215611dc657600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611a2a57600080fd5b80356001600160a01b0381168114611e0d57600080fd5b919050565b600080600060608486031215611e2757600080fd5b611e3084611df6565b95602085013595506040909401359392505050565b600060208284031215611e5757600080fd5b5035919050565b60008060408385031215611e7157600080fd5b82359150611e8160208401611df6565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060408385031215611ecc57600080fd5b611ed583611df6565b9150602083013567ffffffffffffffff811115611ef157600080fd5b8301601f81018513611f0257600080fd5b803567ffffffffffffffff811115611f1c57611f1c611e8a565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715611f8857611f88611e8a565b604052818152828201602001871015611fa057600080fd5b816020840160208301376000602083830101528093505050509250929050565b60008083601f840112611fd257600080fd5b50813567ffffffffffffffff811115611fea57600080fd5b60208301915083602082850101111561200257600080fd5b9250929050565b60008060008060008060a0878903121561202257600080fd5b61202b87611df6565b955060208701359450604087013567ffffffffffffffff81111561204e57600080fd5b61205a89828a01611fc0565b90955093505060608701359150608087013567ffffffffffffffff81111561208157600080fd5b87016080818a03121561209357600080fd5b809150509295509295509295565b60008060008060008086880360a08112156120bb57600080fd5b60208112156120c957600080fd5b508695506120d960208801611df6565b945060408701359350606087013567ffffffffffffffff8111156120fc57600080fd5b61210889828a01611fc0565b979a9699509497949695608090950135949350505050565b60006020828403121561213257600080fd5b611a2a82611df6565b60005b8381101561215657818101518382015260200161213e565b50506000910152565b602081526000825180602084015261217e81604085016020870161213b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b600080600080608085870312156121c657600080fd5b6121cf85611df6565b93506121dd60208601611df6565b92506121eb60408601611df6565b91506121f960608601611df6565b905092959194509250565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6001600160a01b0361225e82611df6565b1682526001600160a01b0361227560208301611df6565b1660208301526040818101359083015260006060820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe10181126122bd57600080fd5b820160208101903567ffffffffffffffff8111156122da57600080fd5b8036038213156122e957600080fd5b60806060860152611b46608086018284612204565b6001600160a01b03871681526001600160a01b038616602082015284604082015260a06060820152600061233660a083018587612204565b8281036080840152612348818561224d565b9998505050505050505050565b84815260606020820152600061236f606083018587612204565b8281036040840152612381818561224d565b979650505050505050565b6001600160a01b0361239d88611df6565b1681526001600160a01b03861660208201526001600160a01b038516604082015283606082015260a0608082015260006123db60a083018486612204565b98975050505050505050565b838152604060208201526000611b46604083018486612204565b60006020828403121561241357600080fd5b5051919050565b60006020828403121561242c57600080fd5b81518015158114611a2a57600080fd5b6000825161244e81846020870161213b565b919091019291505056fea2646970667358221220dae538362f28ba76655cc1bd2bd86c0030f581de43980c289f67432c151216c264736f6c634300081a0033",
}

// ZetaConnectorNativeUpgradeTestABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNativeUpgradeTestMetaData.ABI instead.
var ZetaConnectorNativeUpgradeTestABI = ZetaConnectorNativeUpgradeTestMetaData.ABI

// ZetaConnectorNativeUpgradeTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNativeUpgradeTestMetaData.Bin instead.
var ZetaConnectorNativeUpgradeTestBin = ZetaConnectorNativeUpgradeTestMetaData.Bin

// DeployZetaConnectorNativeUpgradeTest deploys a new Ethereum contract, binding an instance of ZetaConnectorNativeUpgradeTest to it.
func DeployZetaConnectorNativeUpgradeTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZetaConnectorNativeUpgradeTest, error) {
	parsed, err := ZetaConnectorNativeUpgradeTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNativeUpgradeTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorNativeUpgradeTest{ZetaConnectorNativeUpgradeTestCaller: ZetaConnectorNativeUpgradeTestCaller{contract: contract}, ZetaConnectorNativeUpgradeTestTransactor: ZetaConnectorNativeUpgradeTestTransactor{contract: contract}, ZetaConnectorNativeUpgradeTestFilterer: ZetaConnectorNativeUpgradeTestFilterer{contract: contract}}, nil
}

// ZetaConnectorNativeUpgradeTest is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNativeUpgradeTest struct {
	ZetaConnectorNativeUpgradeTestCaller     // Read-only binding to the contract
	ZetaConnectorNativeUpgradeTestTransactor // Write-only binding to the contract
	ZetaConnectorNativeUpgradeTestFilterer   // Log filterer for contract events
}

// ZetaConnectorNativeUpgradeTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNativeUpgradeTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeUpgradeTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNativeUpgradeTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeUpgradeTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNativeUpgradeTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNativeUpgradeTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNativeUpgradeTestSession struct {
	Contract     *ZetaConnectorNativeUpgradeTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ZetaConnectorNativeUpgradeTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNativeUpgradeTestCallerSession struct {
	Contract *ZetaConnectorNativeUpgradeTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// ZetaConnectorNativeUpgradeTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNativeUpgradeTestTransactorSession struct {
	Contract     *ZetaConnectorNativeUpgradeTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ZetaConnectorNativeUpgradeTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNativeUpgradeTestRaw struct {
	Contract *ZetaConnectorNativeUpgradeTest // Generic contract binding to access the raw methods on
}

// ZetaConnectorNativeUpgradeTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNativeUpgradeTestCallerRaw struct {
	Contract *ZetaConnectorNativeUpgradeTestCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNativeUpgradeTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNativeUpgradeTestTransactorRaw struct {
	Contract *ZetaConnectorNativeUpgradeTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNativeUpgradeTest creates a new instance of ZetaConnectorNativeUpgradeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeUpgradeTest(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNativeUpgradeTest, error) {
	contract, err := bindZetaConnectorNativeUpgradeTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTest{ZetaConnectorNativeUpgradeTestCaller: ZetaConnectorNativeUpgradeTestCaller{contract: contract}, ZetaConnectorNativeUpgradeTestTransactor: ZetaConnectorNativeUpgradeTestTransactor{contract: contract}, ZetaConnectorNativeUpgradeTestFilterer: ZetaConnectorNativeUpgradeTestFilterer{contract: contract}}, nil
}

// NewZetaConnectorNativeUpgradeTestCaller creates a new read-only instance of ZetaConnectorNativeUpgradeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeUpgradeTestCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNativeUpgradeTestCaller, error) {
	contract, err := bindZetaConnectorNativeUpgradeTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestCaller{contract: contract}, nil
}

// NewZetaConnectorNativeUpgradeTestTransactor creates a new write-only instance of ZetaConnectorNativeUpgradeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeUpgradeTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNativeUpgradeTestTransactor, error) {
	contract, err := bindZetaConnectorNativeUpgradeTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestTransactor{contract: contract}, nil
}

// NewZetaConnectorNativeUpgradeTestFilterer creates a new log filterer instance of ZetaConnectorNativeUpgradeTest, bound to a specific deployed contract.
func NewZetaConnectorNativeUpgradeTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNativeUpgradeTestFilterer, error) {
	contract, err := bindZetaConnectorNativeUpgradeTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestFilterer{contract: contract}, nil
}

// bindZetaConnectorNativeUpgradeTest binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNativeUpgradeTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNativeUpgradeTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNativeUpgradeTest.Contract.ZetaConnectorNativeUpgradeTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.ZetaConnectorNativeUpgradeTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.ZetaConnectorNativeUpgradeTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNativeUpgradeTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.DEFAULTADMINROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.DEFAULTADMINROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.PAUSERROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.PAUSERROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.TSSROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.TSSROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.WITHDRAWERROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.WITHDRAWERROLE(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Gateway(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Gateway(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.GetRoleAdmin(&_ZetaConnectorNativeUpgradeTest.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.GetRoleAdmin(&_ZetaConnectorNativeUpgradeTest.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.HasRole(&_ZetaConnectorNativeUpgradeTest.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.HasRole(&_ZetaConnectorNativeUpgradeTest.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) Paused() (bool, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Paused(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) Paused() (bool, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Paused(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) ProxiableUUID() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.ProxiableUUID(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.ProxiableUUID(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.SupportsInterface(&_ZetaConnectorNativeUpgradeTest.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.SupportsInterface(&_ZetaConnectorNativeUpgradeTest.CallOpts, interfaceId)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.TssAddress(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.TssAddress(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNativeUpgradeTest.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.ZetaToken(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.ZetaToken(&_ZetaConnectorNativeUpgradeTest.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.GrantRole(&_ZetaConnectorNativeUpgradeTest.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.GrantRole(&_ZetaConnectorNativeUpgradeTest.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address gateway_, address zetaToken_, address tssAddress_, address admin_) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) Initialize(opts *bind.TransactOpts, gateway_ common.Address, zetaToken_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "initialize", gateway_, zetaToken_, tssAddress_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address gateway_, address zetaToken_, address tssAddress_, address admin_) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) Initialize(gateway_ common.Address, zetaToken_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Initialize(&_ZetaConnectorNativeUpgradeTest.TransactOpts, gateway_, zetaToken_, tssAddress_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address gateway_, address zetaToken_, address tssAddress_, address admin_) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) Initialize(gateway_ common.Address, zetaToken_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Initialize(&_ZetaConnectorNativeUpgradeTest.TransactOpts, gateway_, zetaToken_, tssAddress_, admin_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Pause(&_ZetaConnectorNativeUpgradeTest.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Pause(&_ZetaConnectorNativeUpgradeTest.TransactOpts)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) ReceiveTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "receiveTokens", amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.ReceiveTokens(&_ZetaConnectorNativeUpgradeTest.TransactOpts, amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.ReceiveTokens(&_ZetaConnectorNativeUpgradeTest.TransactOpts, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.RenounceRole(&_ZetaConnectorNativeUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.RenounceRole(&_ZetaConnectorNativeUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.RevokeRole(&_ZetaConnectorNativeUpgradeTest.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.RevokeRole(&_ZetaConnectorNativeUpgradeTest.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Unpause(&_ZetaConnectorNativeUpgradeTest.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Unpause(&_ZetaConnectorNativeUpgradeTest.TransactOpts)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) UpdateTSSAddress(opts *bind.TransactOpts, newTSSAddress common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "updateTSSAddress", newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.UpdateTSSAddress(&_ZetaConnectorNativeUpgradeTest.TransactOpts, newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.UpdateTSSAddress(&_ZetaConnectorNativeUpgradeTest.TransactOpts, newTSSAddress)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.UpgradeToAndCall(&_ZetaConnectorNativeUpgradeTest.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.UpgradeToAndCall(&_ZetaConnectorNativeUpgradeTest.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 ) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int, arg2 [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "withdraw", to, amount, arg2)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 ) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) Withdraw(to common.Address, amount *big.Int, arg2 [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Withdraw(&_ZetaConnectorNativeUpgradeTest.TransactOpts, to, amount, arg2)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 ) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) Withdraw(to common.Address, amount *big.Int, arg2 [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.Withdraw(&_ZetaConnectorNativeUpgradeTest.TransactOpts, to, amount, arg2)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x6fb9a7af.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, uint256 amount, bytes data, bytes32 ) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) WithdrawAndCall(opts *bind.TransactOpts, messageContext MessageContext, to common.Address, amount *big.Int, data []byte, arg4 [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "withdrawAndCall", messageContext, to, amount, data, arg4)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x6fb9a7af.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, uint256 amount, bytes data, bytes32 ) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) WithdrawAndCall(messageContext MessageContext, to common.Address, amount *big.Int, data []byte, arg4 [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.WithdrawAndCall(&_ZetaConnectorNativeUpgradeTest.TransactOpts, messageContext, to, amount, data, arg4)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x6fb9a7af.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, uint256 amount, bytes data, bytes32 ) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) WithdrawAndCall(messageContext MessageContext, to common.Address, amount *big.Int, data []byte, arg4 [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.WithdrawAndCall(&_ZetaConnectorNativeUpgradeTest.TransactOpts, messageContext, to, amount, data, arg4)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x6f8728ad.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 , (address,address,uint256,bytes) revertContext) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactor) WithdrawAndRevert(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, arg3 [32]byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.contract.Transact(opts, "withdrawAndRevert", to, amount, data, arg3, revertContext)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x6f8728ad.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 , (address,address,uint256,bytes) revertContext) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, arg3 [32]byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.WithdrawAndRevert(&_ZetaConnectorNativeUpgradeTest.TransactOpts, to, amount, data, arg3, revertContext)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x6f8728ad.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 , (address,address,uint256,bytes) revertContext) returns()
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestTransactorSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, arg3 [32]byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ZetaConnectorNativeUpgradeTest.Contract.WithdrawAndRevert(&_ZetaConnectorNativeUpgradeTest.TransactOpts, to, amount, data, arg3, revertContext)
}

// ZetaConnectorNativeUpgradeTestInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestInitializedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestInitialized // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestInitialized)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestInitialized)
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
func (it *ZetaConnectorNativeUpgradeTestInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestInitialized represents a Initialized event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZetaConnectorNativeUpgradeTestInitializedIterator, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestInitializedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestInitialized) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestInitialized)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseInitialized(log types.Log) (*ZetaConnectorNativeUpgradeTestInitialized, error) {
	event := new(ZetaConnectorNativeUpgradeTestInitialized)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestPausedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestPaused // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestPaused)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestPaused)
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
func (it *ZetaConnectorNativeUpgradeTestPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestPaused represents a Paused event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterPaused(opts *bind.FilterOpts) (*ZetaConnectorNativeUpgradeTestPausedIterator, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestPausedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestPaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestPaused)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParsePaused(log types.Log) (*ZetaConnectorNativeUpgradeTestPaused, error) {
	event := new(ZetaConnectorNativeUpgradeTestPaused)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestRoleAdminChangedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestRoleAdminChanged)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestRoleAdminChanged)
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
func (it *ZetaConnectorNativeUpgradeTestRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestRoleAdminChanged represents a RoleAdminChanged event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZetaConnectorNativeUpgradeTestRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestRoleAdminChangedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestRoleAdminChanged)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseRoleAdminChanged(log types.Log) (*ZetaConnectorNativeUpgradeTestRoleAdminChanged, error) {
	event := new(ZetaConnectorNativeUpgradeTestRoleAdminChanged)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestRoleGrantedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestRoleGranted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestRoleGranted)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestRoleGranted)
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
func (it *ZetaConnectorNativeUpgradeTestRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestRoleGranted represents a RoleGranted event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZetaConnectorNativeUpgradeTestRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestRoleGrantedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestRoleGranted)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseRoleGranted(log types.Log) (*ZetaConnectorNativeUpgradeTestRoleGranted, error) {
	event := new(ZetaConnectorNativeUpgradeTestRoleGranted)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestRoleRevokedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestRoleRevoked)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestRoleRevoked)
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
func (it *ZetaConnectorNativeUpgradeTestRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestRoleRevoked represents a RoleRevoked event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZetaConnectorNativeUpgradeTestRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestRoleRevokedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestRoleRevoked)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseRoleRevoked(log types.Log) (*ZetaConnectorNativeUpgradeTestRoleRevoked, error) {
	event := new(ZetaConnectorNativeUpgradeTestRoleRevoked)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestUnpausedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestUnpaused // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestUnpaused)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestUnpaused)
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
func (it *ZetaConnectorNativeUpgradeTestUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestUnpaused represents a Unpaused event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ZetaConnectorNativeUpgradeTestUnpausedIterator, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestUnpausedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestUnpaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestUnpaused)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseUnpaused(log types.Log) (*ZetaConnectorNativeUpgradeTestUnpaused, error) {
	event := new(ZetaConnectorNativeUpgradeTestUnpaused)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator is returned from FilterUpdatedZetaConnectorTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedZetaConnectorTSSAddress events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator struct {
	Event *ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress)
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
func (it *ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress represents a UpdatedZetaConnectorTSSAddress event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedZetaConnectorTSSAddress is a free log retrieval operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterUpdatedZetaConnectorTSSAddress(opts *bind.FilterOpts) (*ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "UpdatedZetaConnectorTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedZetaConnectorTSSAddress is a free log subscription operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchUpdatedZetaConnectorTSSAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "UpdatedZetaConnectorTSSAddress", log); err != nil {
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

// ParseUpdatedZetaConnectorTSSAddress is a log parse operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseUpdatedZetaConnectorTSSAddress(log types.Log) (*ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress, error) {
	event := new(ZetaConnectorNativeUpgradeTestUpdatedZetaConnectorTSSAddress)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "UpdatedZetaConnectorTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestUpgradedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestUpgraded // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestUpgraded)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestUpgraded)
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
func (it *ZetaConnectorNativeUpgradeTestUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestUpgraded represents a Upgraded event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ZetaConnectorNativeUpgradeTestUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestUpgradedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestUpgraded)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseUpgraded(log types.Log) (*ZetaConnectorNativeUpgradeTestUpgraded, error) {
	event := new(ZetaConnectorNativeUpgradeTestUpgraded)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawnIterator struct {
	Event *ZetaConnectorNativeUpgradeTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawn)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawn)
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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestWithdrawn represents a Withdrawn event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeUpgradeTestWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestWithdrawnIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestWithdrawn)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseWithdrawn(log types.Log) (*ZetaConnectorNativeUpgradeTestWithdrawn, error) {
	event := new(ZetaConnectorNativeUpgradeTestWithdrawn)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawnAndCalledIterator struct {
	Event *ZetaConnectorNativeUpgradeTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawnAndCalled)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawnAndCalled)
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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawnAndCalled struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeUpgradeTestWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestWithdrawnAndCalledIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestWithdrawnAndCalled, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestWithdrawnAndCalled)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*ZetaConnectorNativeUpgradeTestWithdrawnAndCalled, error) {
	event := new(ZetaConnectorNativeUpgradeTestWithdrawnAndCalled)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawnAndRevertedIterator struct {
	Event *ZetaConnectorNativeUpgradeTestWithdrawnAndReverted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawnAndReverted)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawnAndReverted)
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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawnAndReverted struct {
	To            common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeUpgradeTestWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestWithdrawnAndRevertedIterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestWithdrawnAndReverted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestWithdrawnAndReverted)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
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

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseWithdrawnAndReverted(log types.Log) (*ZetaConnectorNativeUpgradeTestWithdrawnAndReverted, error) {
	event := new(ZetaConnectorNativeUpgradeTestWithdrawnAndReverted)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNativeUpgradeTestWithdrawnV2Iterator is returned from FilterWithdrawnV2 and is used to iterate over the raw logs and unpacked data for WithdrawnV2 events raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawnV2Iterator struct {
	Event *ZetaConnectorNativeUpgradeTestWithdrawnV2 // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawnV2)
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
		it.Event = new(ZetaConnectorNativeUpgradeTestWithdrawnV2)
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
func (it *ZetaConnectorNativeUpgradeTestWithdrawnV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNativeUpgradeTestWithdrawnV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNativeUpgradeTestWithdrawnV2 represents a WithdrawnV2 event raised by the ZetaConnectorNativeUpgradeTest contract.
type ZetaConnectorNativeUpgradeTestWithdrawnV2 struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnV2 is a free log retrieval operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) FilterWithdrawnV2(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNativeUpgradeTestWithdrawnV2Iterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.FilterLogs(opts, "WithdrawnV2", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNativeUpgradeTestWithdrawnV2Iterator{contract: _ZetaConnectorNativeUpgradeTest.contract, event: "WithdrawnV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnV2 is a free log subscription operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) WatchWithdrawnV2(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNativeUpgradeTestWithdrawnV2, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNativeUpgradeTest.contract.WatchLogs(opts, "WithdrawnV2", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNativeUpgradeTestWithdrawnV2)
				if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
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

// ParseWithdrawnV2 is a log parse operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNativeUpgradeTest *ZetaConnectorNativeUpgradeTestFilterer) ParseWithdrawnV2(log types.Log) (*ZetaConnectorNativeUpgradeTestWithdrawnV2, error) {
	event := new(ZetaConnectorNativeUpgradeTestWithdrawnV2)
	if err := _ZetaConnectorNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
