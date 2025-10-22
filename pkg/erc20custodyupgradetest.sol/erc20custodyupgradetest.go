// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20custodyupgradetest

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

// ERC20CustodyUpgradeTestMetaData contains all meta data concerning the ERC20CustodyUpgradeTest contract.
var ERC20CustodyUpgradeTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WHITELISTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"gateway_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSupportsLegacy\",\"inputs\":[{\"name\":\"_supportsLegacy\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsLegacy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unwhitelist\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"whitelist\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"whitelisted\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"messageContext\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndRevert\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unwhitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedCustodyTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Whitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LegacyMethodsNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a060405230608052348015601357600080fd5b5060805161296f61003d600039600081816117b1015281816117da01526119b2015261296f6000f3fe6080604052600436106101c25760003560e01c806391d14854116100f7578063ad3cb1cc11610095578063d9caed1211610064578063d9caed12146105f6578063e609055e14610616578063e63ab1e914610636578063eab103df1461066a57600080fd5b8063ad3cb1cc14610530578063c0c53b8b14610586578063d547741f146105a6578063d936547e146105c657600080fd5b80639a590427116100d15780639a590427146104bb5780639b19251a146104db578063a217fddf146104fb578063ad0818521461051057600080fd5b806391d1485414610416578063950837aa1461047b57806399a3c3561461049b57600080fd5b80634f1ef286116101645780635b1125911161013e5780635b112591146103765780635c975abb146103965780638456cb59146103cd57806385f438c1146103e257600080fd5b80634f1ef2861461031a57806352d1902d1461032d578063570618e11461034257600080fd5b8063252f07bf116101a0578063252f07bf146102915780632f2ff15d146102c357806336568abe146102e55780633f4ba83a1461030557600080fd5b806301ffc9a7146101c7578063116191b6146101fc578063248a9ca314610234575b600080fd5b3480156101d357600080fd5b506101e76101e2366004612170565b61068a565b60405190151581526020015b60405180910390f35b34801561020857600080fd5b5060005461021c906001600160a01b031681565b6040516001600160a01b0390911681526020016101f3565b34801561024057600080fd5b5061028361024f3660046121b2565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016101f3565b34801561029d57600080fd5b506002546101e79074010000000000000000000000000000000000000000900460ff1681565b3480156102cf57600080fd5b506102e36102de3660046121e0565b610723565b005b3480156102f157600080fd5b506102e36103003660046121e0565b61076d565b34801561031157600080fd5b506102e36107be565b6102e361032836600461223f565b6107f3565b34801561033957600080fd5b50610283610812565b34801561034e57600080fd5b506102837f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b34801561038257600080fd5b5060025461021c906001600160a01b031681565b3480156103a257600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166101e7565b3480156103d957600080fd5b506102e3610841565b3480156103ee57600080fd5b506102837f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b34801561042257600080fd5b506101e76104313660046121e0565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561048757600080fd5b506102e3610496366004612348565b610873565b3480156104a757600080fd5b506102e36104b63660046123ae565b610a06565b3480156104c757600080fd5b506102e36104d6366004612348565b610bb4565b3480156104e757600080fd5b506102e36104f6366004612348565b610c68565b34801561050757600080fd5b50610283600081565b34801561051c57600080fd5b506102e361052b366004612451565b610d22565b34801561053c57600080fd5b506105796040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101f39190612500565b34801561059257600080fd5b506102e36105a1366004612551565b610e94565b3480156105b257600080fd5b506102e36105c13660046121e0565b6111b9565b3480156105d257600080fd5b506101e76105e1366004612348565b60016020526000908152604090205460ff1681565b34801561060257600080fd5b506102e361061136600461259c565b6111fd565b34801561062257600080fd5b506102e36106313660046125dd565b611314565b34801561064257600080fd5b506102837f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b34801561067657600080fd5b506102e361068536600461265e565b61155f565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061071d57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461075d816115b5565b61076783836115bf565b50505050565b6001600160a01b03811633146107af576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107b9828261168e565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6107e8816115b5565b6107f0611734565b50565b6107fb6117a6565b61080482611878565b61080e8282611883565b5050565b600061081c6119a7565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61086b816115b5565b6107f0611a09565b600061087e816115b5565b6001600160a01b0382166108be576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546108f5907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4906001600160a01b031661168e565b5060025461092d907f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a906001600160a01b031661168e565b506109587f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4836115bf565b506109837f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a836115bf565b50600254604080516001600160a01b03928316815291841660208301527f4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6910160405180910390a150600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b610a0e611a64565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610a38816115b5565b610a40611ae5565b6001600160a01b03861660009081526001602052604090205460ff16610a92576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610aac906001600160a01b03888116911687611b41565b6000546040517faa0c0fc10000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063aa0c0fc190610aff9089908b908a908a908a908a90600401612779565b600060405180830381600087803b158015610b1957600080fd5b505af1158015610b2d573d6000803e3d6000fd5b50505050856001600160a01b0316876001600160a01b03167f7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb972187878787604051610b7a94939291906127d0565b60405180910390a350610bac60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050505050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a610bde816115b5565b6001600160a01b038216610c1e576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216600081815260016020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a610c92816115b5565b6001600160a01b038216610cd2576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382166000818152600160208190526040808320805460ff1916909217909155517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b610d2a611a64565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610d54816115b5565b610d5c611ae5565b6001600160a01b03851660009081526001602052604090205460ff16610dae576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610dc8906001600160a01b03878116911686611b41565b6000546040517f7bbe9afa0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690637bbe9afa90610e1b908a9089908b908a908a908a90600401612807565b600060405180830381600087803b158015610e3557600080fd5b505af1158015610e49573d6000803e3d6000fd5b50505050846001600160a01b0316866001600160a01b03167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5868686604051610b7a93929190612866565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610edf5750825b905060008267ffffffffffffffff166001148015610efc5750303b155b905081158015610f0a575080155b15610f41576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610fa25784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b0388161580610fbf57506001600160a01b038716155b80610fd157506001600160a01b038616155b15611008576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611010611bdb565b611018611be3565b611020611bdb565b611028611bf3565b600080546001600160a01b03808b167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617835560028054918b169190921617905561107690876115bf565b506110a17f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a876115bf565b506110cc7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a886115bf565b506110f77f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4886115bf565b506111227f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a876115bf565b5061114d7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a886115bf565b5083156111af5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546111f3816115b5565b610767838361168e565b611205611a64565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461122f816115b5565b611237611ae5565b6001600160a01b03831660009081526001602052604090205460ff16611289576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61129d6001600160a01b0384168584611b41565b826001600160a01b0316846001600160a01b03167fd4dabfe72081670cc78f2ebda8e2eddaf3feebde6288dcb8fe673b3dc201b5a4846040516112e291815260200190565b60405180910390a3506107b960017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b61131c611a64565b611324611ae5565b60025474010000000000000000000000000000000000000000900460ff16611378576040517f73cba66300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03841660009081526001602052604090205460ff166113ca576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000906001600160a01b038616906370a0823190602401602060405180830381865afa15801561142a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061144e9190612880565b90506114656001600160a01b038616333087611c03565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038616907f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae9089908990859085906370a0823190602401602060405180830381865afa1580156114ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115109190612880565b61151a9190612899565b878760405161152d9594939291906128d3565b60405180910390a250610bac60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b600061156a816115b5565b506002805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b6107f08133611c3c565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16611684576000848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561163a3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061071d565b600091505061071d565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615611684576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061071d565b61173c611cc9565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061183f57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166118337f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15611876576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b600061080e816115b5565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156118fb575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526118f891810190612880565b60015b611941576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc811461199d576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401611938565b6107b98383611d24565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611876576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611a11611ae5565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833611788565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611adf576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615611876576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040516001600160a01b038381166024830152604482018390526107b991859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611d7a565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b611876611df6565b611beb611df6565b611876611e5d565b611bfb611df6565b611876611e65565b6040516001600160a01b0384811660248301528381166044830152606482018390526107679186918216906323b872dd90608401611b6e565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff1661080e576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401611938565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16611876576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611d2d82611e98565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115611d72576107b98282611f40565b61080e611fb6565b6000611d8f6001600160a01b03841683611fee565b90508051600014158015611db4575080806020019051810190611db29190612900565b155b156107b9576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611938565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16611876576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611bb5611df6565b611e6d611df6565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff19169055565b806001600160a01b03163b600003611ee7576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401611938565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051611f5d919061291d565b600060405180830381855af49150503d8060008114611f98576040519150601f19603f3d011682016040523d82523d6000602084013e611f9d565b606091505b5091509150611fad858383612003565b95945050505050565b3415611876576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060611ffc83836000612078565b9392505050565b606082612018576120138261212e565b611ffc565b815115801561202f57506001600160a01b0384163b155b15612071576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611938565b5092915050565b6060814710156120b6576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611938565b600080856001600160a01b031684866040516120d2919061291d565b60006040518083038185875af1925050503d806000811461210f576040519150601f19603f3d011682016040523d82523d6000602084013e612114565b606091505b5091509150612124868383612003565b9695505050505050565b80511561213e5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561218257600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611ffc57600080fd5b6000602082840312156121c457600080fd5b5035919050565b6001600160a01b03811681146107f057600080fd5b600080604083850312156121f357600080fd5b823591506020830135612205816121cb565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806040838503121561225257600080fd5b823561225d816121cb565b9150602083013567ffffffffffffffff81111561227957600080fd5b8301601f8101851361228a57600080fd5b803567ffffffffffffffff8111156122a4576122a4612210565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561231057612310612210565b60405281815282820160200187101561232857600080fd5b816020840160208301376000602083830101528093505050509250929050565b60006020828403121561235a57600080fd5b8135611ffc816121cb565b60008083601f84011261237757600080fd5b50813567ffffffffffffffff81111561238f57600080fd5b6020830191508360208285010111156123a757600080fd5b9250929050565b60008060008060008060a087890312156123c757600080fd5b86356123d2816121cb565b955060208701356123e2816121cb565b945060408701359350606087013567ffffffffffffffff81111561240557600080fd5b61241189828a01612365565b909450925050608087013567ffffffffffffffff81111561243157600080fd5b87016080818a03121561244357600080fd5b809150509295509295509295565b60008060008060008086880360a081121561246b57600080fd5b602081121561247957600080fd5b50869550602087013561248b816121cb565b9450604087013561249b816121cb565b935060608701359250608087013567ffffffffffffffff8111156124be57600080fd5b6124ca89828a01612365565b979a9699509497509295939492505050565b60005b838110156124f75781810151838201526020016124df565b50506000910152565b602081526000825180602084015261251f8160408501602087016124dc565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60008060006060848603121561256657600080fd5b8335612571816121cb565b92506020840135612581816121cb565b91506040840135612591816121cb565b809150509250925092565b6000806000606084860312156125b157600080fd5b83356125bc816121cb565b925060208401356125cc816121cb565b929592945050506040919091013590565b600080600080600080608087890312156125f657600080fd5b863567ffffffffffffffff81111561260d57600080fd5b61261989828a01612365565b909750955050602087013561262d816121cb565b935060408701359250606087013567ffffffffffffffff8111156124be57600080fd5b80151581146107f057600080fd5b60006020828403121561267057600080fd5b8135611ffc81612650565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b600081356126d1816121cb565b6001600160a01b0316835260208201356126ea816121cb565b6001600160a01b03166020840152604082810135908401526060820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe101811261273857600080fd5b820160208101903567ffffffffffffffff81111561275557600080fd5b80360382131561276457600080fd5b60806060860152611fad60808601828461267b565b6001600160a01b03871681526001600160a01b038616602082015284604082015260a0606082015260006127b160a08301858761267b565b82810360808401526127c381856126c4565b9998505050505050505050565b8481526060602082015260006127ea60608301858761267b565b82810360408401526127fc81856126c4565b979650505050505050565b60008735612814816121cb565b6001600160a01b0381168352506001600160a01b03871660208301526001600160a01b038616604083015284606083015260a0608083015261285a60a08301848661267b565b98975050505050505050565b838152604060208201526000611fad60408301848661267b565b60006020828403121561289257600080fd5b5051919050565b8181038181111561071d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6060815260006128e760608301878961267b565b856020840152828103604084015261285a81858761267b565b60006020828403121561291257600080fd5b8151611ffc81612650565b6000825161292f8184602087016124dc565b919091019291505056fea2646970667358221220b4f260470f6b091a7ad3812ff64f5f940e6590b37d90215988fdb61122fc4ea464736f6c634300081a0033",
}

// ERC20CustodyUpgradeTestABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20CustodyUpgradeTestMetaData.ABI instead.
var ERC20CustodyUpgradeTestABI = ERC20CustodyUpgradeTestMetaData.ABI

// ERC20CustodyUpgradeTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20CustodyUpgradeTestMetaData.Bin instead.
var ERC20CustodyUpgradeTestBin = ERC20CustodyUpgradeTestMetaData.Bin

// DeployERC20CustodyUpgradeTest deploys a new Ethereum contract, binding an instance of ERC20CustodyUpgradeTest to it.
func DeployERC20CustodyUpgradeTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20CustodyUpgradeTest, error) {
	parsed, err := ERC20CustodyUpgradeTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20CustodyUpgradeTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20CustodyUpgradeTest{ERC20CustodyUpgradeTestCaller: ERC20CustodyUpgradeTestCaller{contract: contract}, ERC20CustodyUpgradeTestTransactor: ERC20CustodyUpgradeTestTransactor{contract: contract}, ERC20CustodyUpgradeTestFilterer: ERC20CustodyUpgradeTestFilterer{contract: contract}}, nil
}

// ERC20CustodyUpgradeTest is an auto generated Go binding around an Ethereum contract.
type ERC20CustodyUpgradeTest struct {
	ERC20CustodyUpgradeTestCaller     // Read-only binding to the contract
	ERC20CustodyUpgradeTestTransactor // Write-only binding to the contract
	ERC20CustodyUpgradeTestFilterer   // Log filterer for contract events
}

// ERC20CustodyUpgradeTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20CustodyUpgradeTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyUpgradeTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20CustodyUpgradeTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyUpgradeTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20CustodyUpgradeTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyUpgradeTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20CustodyUpgradeTestSession struct {
	Contract     *ERC20CustodyUpgradeTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20CustodyUpgradeTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CustodyUpgradeTestCallerSession struct {
	Contract *ERC20CustodyUpgradeTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ERC20CustodyUpgradeTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20CustodyUpgradeTestTransactorSession struct {
	Contract     *ERC20CustodyUpgradeTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ERC20CustodyUpgradeTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20CustodyUpgradeTestRaw struct {
	Contract *ERC20CustodyUpgradeTest // Generic contract binding to access the raw methods on
}

// ERC20CustodyUpgradeTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CustodyUpgradeTestCallerRaw struct {
	Contract *ERC20CustodyUpgradeTestCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20CustodyUpgradeTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20CustodyUpgradeTestTransactorRaw struct {
	Contract *ERC20CustodyUpgradeTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20CustodyUpgradeTest creates a new instance of ERC20CustodyUpgradeTest, bound to a specific deployed contract.
func NewERC20CustodyUpgradeTest(address common.Address, backend bind.ContractBackend) (*ERC20CustodyUpgradeTest, error) {
	contract, err := bindERC20CustodyUpgradeTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTest{ERC20CustodyUpgradeTestCaller: ERC20CustodyUpgradeTestCaller{contract: contract}, ERC20CustodyUpgradeTestTransactor: ERC20CustodyUpgradeTestTransactor{contract: contract}, ERC20CustodyUpgradeTestFilterer: ERC20CustodyUpgradeTestFilterer{contract: contract}}, nil
}

// NewERC20CustodyUpgradeTestCaller creates a new read-only instance of ERC20CustodyUpgradeTest, bound to a specific deployed contract.
func NewERC20CustodyUpgradeTestCaller(address common.Address, caller bind.ContractCaller) (*ERC20CustodyUpgradeTestCaller, error) {
	contract, err := bindERC20CustodyUpgradeTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestCaller{contract: contract}, nil
}

// NewERC20CustodyUpgradeTestTransactor creates a new write-only instance of ERC20CustodyUpgradeTest, bound to a specific deployed contract.
func NewERC20CustodyUpgradeTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20CustodyUpgradeTestTransactor, error) {
	contract, err := bindERC20CustodyUpgradeTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestTransactor{contract: contract}, nil
}

// NewERC20CustodyUpgradeTestFilterer creates a new log filterer instance of ERC20CustodyUpgradeTest, bound to a specific deployed contract.
func NewERC20CustodyUpgradeTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20CustodyUpgradeTestFilterer, error) {
	contract, err := bindERC20CustodyUpgradeTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestFilterer{contract: contract}, nil
}

// bindERC20CustodyUpgradeTest binds a generic wrapper to an already deployed contract.
func bindERC20CustodyUpgradeTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20CustodyUpgradeTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20CustodyUpgradeTest.Contract.ERC20CustodyUpgradeTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.ERC20CustodyUpgradeTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.ERC20CustodyUpgradeTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20CustodyUpgradeTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20CustodyUpgradeTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20CustodyUpgradeTest.Contract.DEFAULTADMINROLE(&_ERC20CustodyUpgradeTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20CustodyUpgradeTest.Contract.DEFAULTADMINROLE(&_ERC20CustodyUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20CustodyUpgradeTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) PAUSERROLE() ([32]byte, error) {
	return _ERC20CustodyUpgradeTest.Contract.PAUSERROLE(&_ERC20CustodyUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ERC20CustodyUpgradeTest.Contract.PAUSERROLE(&_ERC20CustodyUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20CustodyUpgradeTest.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ERC20CustodyUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_ERC20CustodyUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ERC20CustodyUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_ERC20CustodyUpgradeTest.CallOpts)
}

// WHITELISTERROLE is a free data retrieval call binding the contract method 0x570618e1.
//
// Solidity: function WHITELISTER_ROLE() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCaller) WHITELISTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20CustodyUpgradeTest.contract.Call(opts, &out, "WHITELISTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WHITELISTERROLE is a free data retrieval call binding the contract method 0x570618e1.
//
// Solidity: function WHITELISTER_ROLE() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) WHITELISTERROLE() ([32]byte, error) {
	return _ERC20CustodyUpgradeTest.Contract.WHITELISTERROLE(&_ERC20CustodyUpgradeTest.CallOpts)
}

// WHITELISTERROLE is a free data retrieval call binding the contract method 0x570618e1.
//
// Solidity: function WHITELISTER_ROLE() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerSession) WHITELISTERROLE() ([32]byte, error) {
	return _ERC20CustodyUpgradeTest.Contract.WHITELISTERROLE(&_ERC20CustodyUpgradeTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20CustodyUpgradeTest.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ERC20CustodyUpgradeTest.Contract.WITHDRAWERROLE(&_ERC20CustodyUpgradeTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ERC20CustodyUpgradeTest.Contract.WITHDRAWERROLE(&_ERC20CustodyUpgradeTest.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20CustodyUpgradeTest.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) Gateway() (common.Address, error) {
	return _ERC20CustodyUpgradeTest.Contract.Gateway(&_ERC20CustodyUpgradeTest.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerSession) Gateway() (common.Address, error) {
	return _ERC20CustodyUpgradeTest.Contract.Gateway(&_ERC20CustodyUpgradeTest.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC20CustodyUpgradeTest.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC20CustodyUpgradeTest.Contract.GetRoleAdmin(&_ERC20CustodyUpgradeTest.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC20CustodyUpgradeTest.Contract.GetRoleAdmin(&_ERC20CustodyUpgradeTest.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20CustodyUpgradeTest.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC20CustodyUpgradeTest.Contract.HasRole(&_ERC20CustodyUpgradeTest.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC20CustodyUpgradeTest.Contract.HasRole(&_ERC20CustodyUpgradeTest.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20CustodyUpgradeTest.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) Paused() (bool, error) {
	return _ERC20CustodyUpgradeTest.Contract.Paused(&_ERC20CustodyUpgradeTest.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerSession) Paused() (bool, error) {
	return _ERC20CustodyUpgradeTest.Contract.Paused(&_ERC20CustodyUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20CustodyUpgradeTest.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) ProxiableUUID() ([32]byte, error) {
	return _ERC20CustodyUpgradeTest.Contract.ProxiableUUID(&_ERC20CustodyUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ERC20CustodyUpgradeTest.Contract.ProxiableUUID(&_ERC20CustodyUpgradeTest.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC20CustodyUpgradeTest.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC20CustodyUpgradeTest.Contract.SupportsInterface(&_ERC20CustodyUpgradeTest.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC20CustodyUpgradeTest.Contract.SupportsInterface(&_ERC20CustodyUpgradeTest.CallOpts, interfaceId)
}

// SupportsLegacy is a free data retrieval call binding the contract method 0x252f07bf.
//
// Solidity: function supportsLegacy() view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCaller) SupportsLegacy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20CustodyUpgradeTest.contract.Call(opts, &out, "supportsLegacy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsLegacy is a free data retrieval call binding the contract method 0x252f07bf.
//
// Solidity: function supportsLegacy() view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) SupportsLegacy() (bool, error) {
	return _ERC20CustodyUpgradeTest.Contract.SupportsLegacy(&_ERC20CustodyUpgradeTest.CallOpts)
}

// SupportsLegacy is a free data retrieval call binding the contract method 0x252f07bf.
//
// Solidity: function supportsLegacy() view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerSession) SupportsLegacy() (bool, error) {
	return _ERC20CustodyUpgradeTest.Contract.SupportsLegacy(&_ERC20CustodyUpgradeTest.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20CustodyUpgradeTest.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) TssAddress() (common.Address, error) {
	return _ERC20CustodyUpgradeTest.Contract.TssAddress(&_ERC20CustodyUpgradeTest.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerSession) TssAddress() (common.Address, error) {
	return _ERC20CustodyUpgradeTest.Contract.TssAddress(&_ERC20CustodyUpgradeTest.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20CustodyUpgradeTest.contract.Call(opts, &out, "whitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _ERC20CustodyUpgradeTest.Contract.Whitelisted(&_ERC20CustodyUpgradeTest.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestCallerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _ERC20CustodyUpgradeTest.Contract.Whitelisted(&_ERC20CustodyUpgradeTest.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xe609055e.
//
// Solidity: function deposit(bytes recipient, address asset, uint256 amount, bytes message) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) Deposit(opts *bind.TransactOpts, recipient []byte, asset common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "deposit", recipient, asset, amount, message)
}

// Deposit is a paid mutator transaction binding the contract method 0xe609055e.
//
// Solidity: function deposit(bytes recipient, address asset, uint256 amount, bytes message) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) Deposit(recipient []byte, asset common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.Deposit(&_ERC20CustodyUpgradeTest.TransactOpts, recipient, asset, amount, message)
}

// Deposit is a paid mutator transaction binding the contract method 0xe609055e.
//
// Solidity: function deposit(bytes recipient, address asset, uint256 amount, bytes message) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) Deposit(recipient []byte, asset common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.Deposit(&_ERC20CustodyUpgradeTest.TransactOpts, recipient, asset, amount, message)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.GrantRole(&_ERC20CustodyUpgradeTest.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.GrantRole(&_ERC20CustodyUpgradeTest.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address gateway_, address tssAddress_, address admin_) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) Initialize(opts *bind.TransactOpts, gateway_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "initialize", gateway_, tssAddress_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address gateway_, address tssAddress_, address admin_) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) Initialize(gateway_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.Initialize(&_ERC20CustodyUpgradeTest.TransactOpts, gateway_, tssAddress_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address gateway_, address tssAddress_, address admin_) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) Initialize(gateway_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.Initialize(&_ERC20CustodyUpgradeTest.TransactOpts, gateway_, tssAddress_, admin_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) Pause() (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.Pause(&_ERC20CustodyUpgradeTest.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) Pause() (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.Pause(&_ERC20CustodyUpgradeTest.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.RenounceRole(&_ERC20CustodyUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.RenounceRole(&_ERC20CustodyUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.RevokeRole(&_ERC20CustodyUpgradeTest.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.RevokeRole(&_ERC20CustodyUpgradeTest.TransactOpts, role, account)
}

// SetSupportsLegacy is a paid mutator transaction binding the contract method 0xeab103df.
//
// Solidity: function setSupportsLegacy(bool _supportsLegacy) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) SetSupportsLegacy(opts *bind.TransactOpts, _supportsLegacy bool) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "setSupportsLegacy", _supportsLegacy)
}

// SetSupportsLegacy is a paid mutator transaction binding the contract method 0xeab103df.
//
// Solidity: function setSupportsLegacy(bool _supportsLegacy) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) SetSupportsLegacy(_supportsLegacy bool) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.SetSupportsLegacy(&_ERC20CustodyUpgradeTest.TransactOpts, _supportsLegacy)
}

// SetSupportsLegacy is a paid mutator transaction binding the contract method 0xeab103df.
//
// Solidity: function setSupportsLegacy(bool _supportsLegacy) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) SetSupportsLegacy(_supportsLegacy bool) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.SetSupportsLegacy(&_ERC20CustodyUpgradeTest.TransactOpts, _supportsLegacy)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) Unpause() (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.Unpause(&_ERC20CustodyUpgradeTest.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) Unpause() (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.Unpause(&_ERC20CustodyUpgradeTest.TransactOpts)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address token) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) Unwhitelist(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "unwhitelist", token)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address token) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) Unwhitelist(token common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.Unwhitelist(&_ERC20CustodyUpgradeTest.TransactOpts, token)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address token) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) Unwhitelist(token common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.Unwhitelist(&_ERC20CustodyUpgradeTest.TransactOpts, token)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) UpdateTSSAddress(opts *bind.TransactOpts, newTSSAddress common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "updateTSSAddress", newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.UpdateTSSAddress(&_ERC20CustodyUpgradeTest.TransactOpts, newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.UpdateTSSAddress(&_ERC20CustodyUpgradeTest.TransactOpts, newTSSAddress)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.UpgradeToAndCall(&_ERC20CustodyUpgradeTest.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.UpgradeToAndCall(&_ERC20CustodyUpgradeTest.TransactOpts, newImplementation, data)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address token) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) Whitelist(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "whitelist", token)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address token) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) Whitelist(token common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.Whitelist(&_ERC20CustodyUpgradeTest.TransactOpts, token)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address token) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) Whitelist(token common.Address) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.Whitelist(&_ERC20CustodyUpgradeTest.TransactOpts, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address to, address token, uint256 amount) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "withdraw", to, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address to, address token, uint256 amount) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) Withdraw(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.Withdraw(&_ERC20CustodyUpgradeTest.TransactOpts, to, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address to, address token, uint256 amount) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) Withdraw(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.Withdraw(&_ERC20CustodyUpgradeTest.TransactOpts, to, token, amount)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0xad081852.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, address token, uint256 amount, bytes data) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) WithdrawAndCall(opts *bind.TransactOpts, messageContext MessageContext, to common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "withdrawAndCall", messageContext, to, token, amount, data)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0xad081852.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, address token, uint256 amount, bytes data) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) WithdrawAndCall(messageContext MessageContext, to common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.WithdrawAndCall(&_ERC20CustodyUpgradeTest.TransactOpts, messageContext, to, token, amount, data)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0xad081852.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, address token, uint256 amount, bytes data) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) WithdrawAndCall(messageContext MessageContext, to common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.WithdrawAndCall(&_ERC20CustodyUpgradeTest.TransactOpts, messageContext, to, token, amount, data)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x99a3c356.
//
// Solidity: function withdrawAndRevert(address to, address token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactor) WithdrawAndRevert(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.contract.Transact(opts, "withdrawAndRevert", to, token, amount, data, revertContext)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x99a3c356.
//
// Solidity: function withdrawAndRevert(address to, address token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestSession) WithdrawAndRevert(to common.Address, token common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.WithdrawAndRevert(&_ERC20CustodyUpgradeTest.TransactOpts, to, token, amount, data, revertContext)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x99a3c356.
//
// Solidity: function withdrawAndRevert(address to, address token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestTransactorSession) WithdrawAndRevert(to common.Address, token common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ERC20CustodyUpgradeTest.Contract.WithdrawAndRevert(&_ERC20CustodyUpgradeTest.TransactOpts, to, token, amount, data, revertContext)
}

// ERC20CustodyUpgradeTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestDepositedIterator struct {
	Event *ERC20CustodyUpgradeTestDeposited // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestDeposited)
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
		it.Event = new(ERC20CustodyUpgradeTestDeposited)
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
func (it *ERC20CustodyUpgradeTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestDeposited represents a Deposited event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestDeposited struct {
	Recipient []byte
	Asset     common.Address
	Amount    *big.Int
	Message   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterDeposited(opts *bind.FilterOpts, asset []common.Address) (*ERC20CustodyUpgradeTestDepositedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestDepositedIterator{contract: _ERC20CustodyUpgradeTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestDeposited, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestDeposited)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParseDeposited(log types.Log) (*ERC20CustodyUpgradeTestDeposited, error) {
	event := new(ERC20CustodyUpgradeTestDeposited)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradeTestInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestInitializedIterator struct {
	Event *ERC20CustodyUpgradeTestInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestInitialized)
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
		it.Event = new(ERC20CustodyUpgradeTestInitialized)
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
func (it *ERC20CustodyUpgradeTestInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestInitialized represents a Initialized event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20CustodyUpgradeTestInitializedIterator, error) {

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestInitializedIterator{contract: _ERC20CustodyUpgradeTest.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestInitialized)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParseInitialized(log types.Log) (*ERC20CustodyUpgradeTestInitialized, error) {
	event := new(ERC20CustodyUpgradeTestInitialized)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradeTestPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestPausedIterator struct {
	Event *ERC20CustodyUpgradeTestPaused // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestPaused)
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
		it.Event = new(ERC20CustodyUpgradeTestPaused)
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
func (it *ERC20CustodyUpgradeTestPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestPaused represents a Paused event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC20CustodyUpgradeTestPausedIterator, error) {

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestPausedIterator{contract: _ERC20CustodyUpgradeTest.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestPaused) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestPaused)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParsePaused(log types.Log) (*ERC20CustodyUpgradeTestPaused, error) {
	event := new(ERC20CustodyUpgradeTestPaused)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradeTestRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestRoleAdminChangedIterator struct {
	Event *ERC20CustodyUpgradeTestRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestRoleAdminChanged)
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
		it.Event = new(ERC20CustodyUpgradeTestRoleAdminChanged)
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
func (it *ERC20CustodyUpgradeTestRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestRoleAdminChanged represents a RoleAdminChanged event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ERC20CustodyUpgradeTestRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestRoleAdminChangedIterator{contract: _ERC20CustodyUpgradeTest.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestRoleAdminChanged)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParseRoleAdminChanged(log types.Log) (*ERC20CustodyUpgradeTestRoleAdminChanged, error) {
	event := new(ERC20CustodyUpgradeTestRoleAdminChanged)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradeTestRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestRoleGrantedIterator struct {
	Event *ERC20CustodyUpgradeTestRoleGranted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestRoleGranted)
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
		it.Event = new(ERC20CustodyUpgradeTestRoleGranted)
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
func (it *ERC20CustodyUpgradeTestRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestRoleGranted represents a RoleGranted event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC20CustodyUpgradeTestRoleGrantedIterator, error) {

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

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestRoleGrantedIterator{contract: _ERC20CustodyUpgradeTest.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestRoleGranted)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParseRoleGranted(log types.Log) (*ERC20CustodyUpgradeTestRoleGranted, error) {
	event := new(ERC20CustodyUpgradeTestRoleGranted)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradeTestRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestRoleRevokedIterator struct {
	Event *ERC20CustodyUpgradeTestRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestRoleRevoked)
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
		it.Event = new(ERC20CustodyUpgradeTestRoleRevoked)
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
func (it *ERC20CustodyUpgradeTestRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestRoleRevoked represents a RoleRevoked event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC20CustodyUpgradeTestRoleRevokedIterator, error) {

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

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestRoleRevokedIterator{contract: _ERC20CustodyUpgradeTest.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestRoleRevoked)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParseRoleRevoked(log types.Log) (*ERC20CustodyUpgradeTestRoleRevoked, error) {
	event := new(ERC20CustodyUpgradeTestRoleRevoked)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradeTestUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestUnpausedIterator struct {
	Event *ERC20CustodyUpgradeTestUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestUnpaused)
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
		it.Event = new(ERC20CustodyUpgradeTestUnpaused)
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
func (it *ERC20CustodyUpgradeTestUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestUnpaused represents a Unpaused event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC20CustodyUpgradeTestUnpausedIterator, error) {

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestUnpausedIterator{contract: _ERC20CustodyUpgradeTest.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestUnpaused)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParseUnpaused(log types.Log) (*ERC20CustodyUpgradeTestUnpaused, error) {
	event := new(ERC20CustodyUpgradeTestUnpaused)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradeTestUnwhitelistedIterator is returned from FilterUnwhitelisted and is used to iterate over the raw logs and unpacked data for Unwhitelisted events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestUnwhitelistedIterator struct {
	Event *ERC20CustodyUpgradeTestUnwhitelisted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestUnwhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestUnwhitelisted)
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
		it.Event = new(ERC20CustodyUpgradeTestUnwhitelisted)
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
func (it *ERC20CustodyUpgradeTestUnwhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestUnwhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestUnwhitelisted represents a Unwhitelisted event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestUnwhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnwhitelisted is a free log retrieval operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterUnwhitelisted(opts *bind.FilterOpts, token []common.Address) (*ERC20CustodyUpgradeTestUnwhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestUnwhitelistedIterator{contract: _ERC20CustodyUpgradeTest.contract, event: "Unwhitelisted", logs: logs, sub: sub}, nil
}

// WatchUnwhitelisted is a free log subscription operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchUnwhitelisted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestUnwhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestUnwhitelisted)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
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

// ParseUnwhitelisted is a log parse operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParseUnwhitelisted(log types.Log) (*ERC20CustodyUpgradeTestUnwhitelisted, error) {
	event := new(ERC20CustodyUpgradeTestUnwhitelisted)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradeTestUpdatedCustodyTSSAddressIterator is returned from FilterUpdatedCustodyTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedCustodyTSSAddress events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestUpdatedCustodyTSSAddressIterator struct {
	Event *ERC20CustodyUpgradeTestUpdatedCustodyTSSAddress // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestUpdatedCustodyTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestUpdatedCustodyTSSAddress)
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
		it.Event = new(ERC20CustodyUpgradeTestUpdatedCustodyTSSAddress)
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
func (it *ERC20CustodyUpgradeTestUpdatedCustodyTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestUpdatedCustodyTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestUpdatedCustodyTSSAddress represents a UpdatedCustodyTSSAddress event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestUpdatedCustodyTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCustodyTSSAddress is a free log retrieval operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterUpdatedCustodyTSSAddress(opts *bind.FilterOpts) (*ERC20CustodyUpgradeTestUpdatedCustodyTSSAddressIterator, error) {

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestUpdatedCustodyTSSAddressIterator{contract: _ERC20CustodyUpgradeTest.contract, event: "UpdatedCustodyTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedCustodyTSSAddress is a free log subscription operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchUpdatedCustodyTSSAddress(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestUpdatedCustodyTSSAddress) (event.Subscription, error) {

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestUpdatedCustodyTSSAddress)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
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

// ParseUpdatedCustodyTSSAddress is a log parse operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParseUpdatedCustodyTSSAddress(log types.Log) (*ERC20CustodyUpgradeTestUpdatedCustodyTSSAddress, error) {
	event := new(ERC20CustodyUpgradeTestUpdatedCustodyTSSAddress)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradeTestUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestUpgradedIterator struct {
	Event *ERC20CustodyUpgradeTestUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestUpgraded)
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
		it.Event = new(ERC20CustodyUpgradeTestUpgraded)
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
func (it *ERC20CustodyUpgradeTestUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestUpgraded represents a Upgraded event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ERC20CustodyUpgradeTestUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestUpgradedIterator{contract: _ERC20CustodyUpgradeTest.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestUpgraded)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParseUpgraded(log types.Log) (*ERC20CustodyUpgradeTestUpgraded, error) {
	event := new(ERC20CustodyUpgradeTestUpgraded)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradeTestWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestWhitelistedIterator struct {
	Event *ERC20CustodyUpgradeTestWhitelisted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestWhitelisted)
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
		it.Event = new(ERC20CustodyUpgradeTestWhitelisted)
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
func (it *ERC20CustodyUpgradeTestWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestWhitelisted represents a Whitelisted event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestWhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterWhitelisted(opts *bind.FilterOpts, token []common.Address) (*ERC20CustodyUpgradeTestWhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestWhitelistedIterator{contract: _ERC20CustodyUpgradeTest.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestWhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestWhitelisted)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Whitelisted", log); err != nil {
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

// ParseWhitelisted is a log parse operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParseWhitelisted(log types.Log) (*ERC20CustodyUpgradeTestWhitelisted, error) {
	event := new(ERC20CustodyUpgradeTestWhitelisted)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradeTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestWithdrawnIterator struct {
	Event *ERC20CustodyUpgradeTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestWithdrawn)
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
		it.Event = new(ERC20CustodyUpgradeTestWithdrawn)
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
func (it *ERC20CustodyUpgradeTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestWithdrawn represents a Withdrawn event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestWithdrawn struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ERC20CustodyUpgradeTestWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestWithdrawnIterator{contract: _ERC20CustodyUpgradeTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestWithdrawn, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestWithdrawn)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParseWithdrawn(log types.Log) (*ERC20CustodyUpgradeTestWithdrawn, error) {
	event := new(ERC20CustodyUpgradeTestWithdrawn)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradeTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestWithdrawnAndCalledIterator struct {
	Event *ERC20CustodyUpgradeTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestWithdrawnAndCalled)
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
		it.Event = new(ERC20CustodyUpgradeTestWithdrawnAndCalled)
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
func (it *ERC20CustodyUpgradeTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestWithdrawnAndCalled struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ERC20CustodyUpgradeTestWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestWithdrawnAndCalledIterator{contract: _ERC20CustodyUpgradeTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestWithdrawnAndCalled, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestWithdrawnAndCalled)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*ERC20CustodyUpgradeTestWithdrawnAndCalled, error) {
	event := new(ERC20CustodyUpgradeTestWithdrawnAndCalled)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradeTestWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestWithdrawnAndRevertedIterator struct {
	Event *ERC20CustodyUpgradeTestWithdrawnAndReverted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestWithdrawnAndReverted)
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
		it.Event = new(ERC20CustodyUpgradeTestWithdrawnAndReverted)
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
func (it *ERC20CustodyUpgradeTestWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestWithdrawnAndReverted struct {
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ERC20CustodyUpgradeTestWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestWithdrawnAndRevertedIterator{contract: _ERC20CustodyUpgradeTest.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestWithdrawnAndReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestWithdrawnAndReverted)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
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

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParseWithdrawnAndReverted(log types.Log) (*ERC20CustodyUpgradeTestWithdrawnAndReverted, error) {
	event := new(ERC20CustodyUpgradeTestWithdrawnAndReverted)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradeTestWithdrawnV2Iterator is returned from FilterWithdrawnV2 and is used to iterate over the raw logs and unpacked data for WithdrawnV2 events raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestWithdrawnV2Iterator struct {
	Event *ERC20CustodyUpgradeTestWithdrawnV2 // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradeTestWithdrawnV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgradeTestWithdrawnV2)
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
		it.Event = new(ERC20CustodyUpgradeTestWithdrawnV2)
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
func (it *ERC20CustodyUpgradeTestWithdrawnV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradeTestWithdrawnV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgradeTestWithdrawnV2 represents a WithdrawnV2 event raised by the ERC20CustodyUpgradeTest contract.
type ERC20CustodyUpgradeTestWithdrawnV2 struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnV2 is a free log retrieval operation binding the contract event 0xd4dabfe72081670cc78f2ebda8e2eddaf3feebde6288dcb8fe673b3dc201b5a4.
//
// Solidity: event WithdrawnV2(address indexed to, address indexed token, uint256 amount)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) FilterWithdrawnV2(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ERC20CustodyUpgradeTestWithdrawnV2Iterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.FilterLogs(opts, "WithdrawnV2", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradeTestWithdrawnV2Iterator{contract: _ERC20CustodyUpgradeTest.contract, event: "WithdrawnV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnV2 is a free log subscription operation binding the contract event 0xd4dabfe72081670cc78f2ebda8e2eddaf3feebde6288dcb8fe673b3dc201b5a4.
//
// Solidity: event WithdrawnV2(address indexed to, address indexed token, uint256 amount)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) WatchWithdrawnV2(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgradeTestWithdrawnV2, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20CustodyUpgradeTest.contract.WatchLogs(opts, "WithdrawnV2", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgradeTestWithdrawnV2)
				if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
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

// ParseWithdrawnV2 is a log parse operation binding the contract event 0xd4dabfe72081670cc78f2ebda8e2eddaf3feebde6288dcb8fe673b3dc201b5a4.
//
// Solidity: event WithdrawnV2(address indexed to, address indexed token, uint256 amount)
func (_ERC20CustodyUpgradeTest *ERC20CustodyUpgradeTestFilterer) ParseWithdrawnV2(log types.Log) (*ERC20CustodyUpgradeTestWithdrawnV2, error) {
	event := new(ERC20CustodyUpgradeTestWithdrawnV2)
	if err := _ERC20CustodyUpgradeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
