// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20custody

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

// ERC20CustodyMetaData contains all meta data concerning the ERC20Custody contract.
var ERC20CustodyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STRANDED_FUNDS_REFUNDER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WHITELISTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"gateway_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refundStrandedFunds\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSupportsLegacy\",\"inputs\":[{\"name\":\"_supportsLegacy\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsLegacy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unwhitelist\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"whitelist\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"whitelisted\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"messageContext\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndRevert\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrandedFundsRefunded\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unwhitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedCustodyTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Whitelisted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LegacyMethodsNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedStrandedFundsRefunder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a0604052306080523480156012575f80fd5b50608051612a706100395f395f8181611a4701528181611a700152611c470152612a705ff3fe6080604052600436106101d0575f3560e01c806391d14854116100fd578063ad3cb1cc11610092578063d9caed1211610062578063d9caed121461062c578063e609055e1461064b578063e63ab1e91461066a578063eab103df1461069d575f80fd5b8063ad3cb1cc1461056b578063c0c53b8b146105c0578063d547741f146105df578063d936547e146105fe575f80fd5b80639b19251a116100cd5780639b19251a146104f35780639dbc9b1514610512578063a217fddf14610539578063ad0818521461054c575f80fd5b806391d1485414610433578063950837aa1461049657806399a3c356146104b55780639a590427146104d4575f80fd5b80633f4ba83a116101735780635b112591116101435780635b112591146103975780635c975abb146103b65780638456cb59146103ec57806385f438c114610400575f80fd5b80633f4ba83a146103295780634f1ef2861461033d57806352d1902d14610350578063570618e114610364575f80fd5b8063252f07bf116101ae578063252f07bf146102995780632f2ff15d146102ca57806336568abe146102eb5780633e1b96441461030a575f80fd5b806301ffc9a7146101d4578063116191b614610208578063248a9ca31461023e575b5f80fd5b3480156101df575f80fd5b506101f36101ee3660046122dc565b6106bc565b60405190151581526020015b60405180910390f35b348015610213575f80fd5b505f54610226906001600160a01b031681565b6040516001600160a01b0390911681526020016101ff565b348015610249575f80fd5b5061028b61025836600461231b565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016101ff565b3480156102a4575f80fd5b506002546101f39074010000000000000000000000000000000000000000900460ff1681565b3480156102d5575f80fd5b506102e96102e4366004612346565b610754565b005b3480156102f6575f80fd5b506102e9610305366004612346565b61079d565b348015610315575f80fd5b506102e9610324366004612374565b6107ee565b348015610334575f80fd5b506102e961095e565b6102e961034b3660046123df565b610993565b34801561035b575f80fd5b5061028b6109b2565b34801561036f575f80fd5b5061028b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a81565b3480156103a2575f80fd5b50600254610226906001600160a01b031681565b3480156103c1575f80fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166101f3565b3480156103f7575f80fd5b506102e96109e0565b34801561040b575f80fd5b5061028b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b34801561043e575f80fd5b506101f361044d366004612346565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156104a1575f80fd5b506102e96104b03660046124e2565b610a12565b3480156104c0575f80fd5b506102e96104cf366004612542565b610ba4565b3480156104df575f80fd5b506102e96104ee3660046124e2565b610d4a565b3480156104fe575f80fd5b506102e961050d3660046124e2565b610dfd565b34801561051d575f80fd5b50610226738c1b2e11f2b217caa6f95a31b8d9ec6ad93c880381565b348015610544575f80fd5b5061028b5f81565b348015610557575f80fd5b506102e96105663660046125de565b610eb6565b348015610576575f80fd5b506105b36040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101ff9190612663565b3480156105cb575f80fd5b506102e96105da3660046126b6565b611020565b3480156105ea575f80fd5b506102e96105f9366004612346565b611342565b348015610609575f80fd5b506101f36106183660046124e2565b60016020525f908152604090205460ff1681565b348015610637575f80fd5b506102e9610646366004612374565b611385565b348015610656575f80fd5b506102e96106653660046126fe565b61149b565b348015610675575f80fd5b5061028b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b3480156106a8575f80fd5b506102e96106b7366004612778565b6116e0565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061074e57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461078d81611735565b610797838361173f565b50505050565b6001600160a01b03811633146107df576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107e9828261180b565b505050565b6107f66118af565b33738c1b2e11f2b217caa6f95a31b8d9ec6ad93c880314610843576040517f04086d3700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038316610883576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382165f9081526001602052604090205460ff166108d4576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108e86001600160a01b0383168483611930565b816001600160a01b0316836001600160a01b03167f7b73729b322c3c294385b7454f9895d7975135d24a54310b0e4542962bd089ab8360405161092d91815260200190565b60405180910390a36107e960017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61098881611735565b6109906119ca565b50565b61099b611a3c565b6109a482611b0e565b6109ae8282611b18565b5050565b5f6109bb611c3c565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610a0a81611735565b610990611c9e565b5f610a1c81611735565b6001600160a01b038216610a5c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254610a93907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4906001600160a01b031661180b565b50600254610acb907f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a906001600160a01b031661180b565b50610af67f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48361173f565b50610b217f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a8361173f565b50600254604080516001600160a01b03928316815291841660208301527f4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6910160405180910390a150600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b610bac6118af565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610bd681611735565b610bde611cf9565b6001600160a01b0386165f9081526001602052604090205460ff16610c2f576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54610c48906001600160a01b03888116911687611930565b5f546040517faa0c0fc10000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063aa0c0fc190610c9a9089908b908a908a908a908a9060040161288b565b5f604051808303815f87803b158015610cb1575f80fd5b505af1158015610cc3573d5f803e3d5ffd5b50505050856001600160a01b0316876001600160a01b03167f7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb972187878787604051610d1094939291906128e1565b60405180910390a350610d4260017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050505050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a610d7481611735565b6001600160a01b038216610db4576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382165f81815260016020526040808220805460ff19169055517f51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da467919190a25050565b7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a610e2781611735565b6001600160a01b038216610e67576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382165f818152600160208190526040808320805460ff1916909217909155517faab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a549190a25050565b610ebe6118af565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610ee881611735565b610ef0611cf9565b6001600160a01b0385165f9081526001602052604090205460ff16610f41576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54610f5a906001600160a01b03878116911686611930565b5f546040517f7bbe9afa0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690637bbe9afa90610fac908a9089908b908a908a908a90600401612917565b5f604051808303815f87803b158015610fc3575f80fd5b505af1158015610fd5573d5f803e3d5ffd5b50505050846001600160a01b0316866001600160a01b03167f6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5868686604051610d1093929190612975565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f8115801561106a5750825b90505f8267ffffffffffffffff1660011480156110865750303b155b905081158015611094575080155b156110cb576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000166001178555831561112c5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038816158061114957506001600160a01b038716155b8061115b57506001600160a01b038616155b15611192576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61119a611d55565b6111a2611d5d565b6111aa611d55565b6111b2611d6d565b5f80546001600160a01b03808b167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617835560028054918b16919092161790556111ff908761173f565b5061122a7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8761173f565b506112557f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8861173f565b506112807f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e48861173f565b506112ab7f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a8761173f565b506112d67f8619cecd8b9e095ab43867f5b69d492180450fe862e6b50bfbfb24b75dd84c8a8861173f565b5083156113385784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461137b81611735565b610797838361180b565b61138d6118af565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46113b781611735565b6113bf611cf9565b6001600160a01b0383165f9081526001602052604090205460ff16611410576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114246001600160a01b0384168584611930565b826001600160a01b0316846001600160a01b03167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb8460405161146991815260200190565b60405180910390a3506107e960017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6114a36118af565b6114ab611cf9565b60025474010000000000000000000000000000000000000000900460ff166114ff576040517f73cba66300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0384165f9081526001602052604090205460ff16611550576040517f584a793800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f906001600160a01b038616906370a0823190602401602060405180830381865afa1580156115ad573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115d1919061298e565b90506115e86001600160a01b038616333087611d7d565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038616907f1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae9089908990859085906370a0823190602401602060405180830381865afa15801561166d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611691919061298e565b61169b91906129a5565b87876040516116ae9594939291906129dd565b60405180910390a250610d4260017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f6116ea81611735565b506002805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b6109908133611db6565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16611802575f848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556117b83390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061074e565b5f91505061074e565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615611802575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061074e565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0161192a576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b6040516001600160a01b038381166024830152604482018390526107e991859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611e42565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6119d2611ebc565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480611ad557507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316611ac97f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15611b0c576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f6109ae81611735565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611b90575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611b8d9181019061298e565b60015b611bd6576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611c32576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401611bcd565b6107e98383611f17565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611b0c576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611ca6611cf9565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833611a1e565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615611b0c576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611b0c611f6c565b611d65611f6c565b611b0c611fd3565b611d75611f6c565b611b0c611fdb565b6040516001600160a01b0384811660248301528381166044830152606482018390526107979186918216906323b872dd9060840161195d565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff166109ae576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401611bcd565b5f611e566001600160a01b0384168361200e565b905080515f14158015611e7a575080806020019051810190611e789190612a09565b155b156107e9576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611bcd565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16611b0c576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611f2082612022565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115611f64576107e982826120c9565b6109ae61213b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16611b0c576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119a4611f6c565b611fe3611f6c565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff19169055565b606061201b83835f612173565b9392505050565b806001600160a01b03163b5f03612070576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401611bcd565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f80846001600160a01b0316846040516120e59190612a24565b5f60405180830381855af49150503d805f811461211d576040519150601f19603f3d011682016040523d82523d5f602084013e612122565b606091505b5091509150612132858383612225565b95945050505050565b3415611b0c576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060814710156121b1576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401611bcd565b5f80856001600160a01b031684866040516121cc9190612a24565b5f6040518083038185875af1925050503d805f8114612206576040519150601f19603f3d011682016040523d82523d5f602084013e61220b565b606091505b509150915061221b868383612225565b9695505050505050565b60608261223a576122358261229a565b61201b565b815115801561225157506001600160a01b0384163b155b15612293576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611bcd565b508061201b565b8051156122aa5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f602082840312156122ec575f80fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461201b575f80fd5b5f6020828403121561232b575f80fd5b5035919050565b6001600160a01b0381168114610990575f80fd5b5f8060408385031215612357575f80fd5b82359150602083013561236981612332565b809150509250929050565b5f805f60608486031215612386575f80fd5b833561239181612332565b925060208401356123a181612332565b929592945050506040919091013590565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f80604083850312156123f0575f80fd5b82356123fb81612332565b9150602083013567ffffffffffffffff811115612416575f80fd5b8301601f81018513612426575f80fd5b803567ffffffffffffffff811115612440576124406123b2565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156124ac576124ac6123b2565b6040528181528282016020018710156124c3575f80fd5b816020840160208301375f602083830101528093505050509250929050565b5f602082840312156124f2575f80fd5b813561201b81612332565b5f8083601f84011261250d575f80fd5b50813567ffffffffffffffff811115612524575f80fd5b60208301915083602082850101111561253b575f80fd5b9250929050565b5f805f805f8060a08789031215612557575f80fd5b863561256281612332565b9550602087013561257281612332565b945060408701359350606087013567ffffffffffffffff811115612594575f80fd5b6125a089828a016124fd565b909450925050608087013567ffffffffffffffff8111156125bf575f80fd5b87016080818a0312156125d0575f80fd5b809150509295509295509295565b5f805f805f8086880360a08112156125f4575f80fd5b6020811215612601575f80fd5b50869550602087013561261381612332565b9450604087013561262381612332565b935060608701359250608087013567ffffffffffffffff811115612645575f80fd5b61265189828a016124fd565b979a9699509497509295939492505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f805f606084860312156126c8575f80fd5b83356126d381612332565b925060208401356126e381612332565b915060408401356126f381612332565b809150509250925092565b5f805f805f8060808789031215612713575f80fd5b863567ffffffffffffffff811115612729575f80fd5b61273589828a016124fd565b909750955050602087013561274981612332565b935060408701359250606087013567ffffffffffffffff811115612645575f80fd5b8015158114610990575f80fd5b5f60208284031215612788575f80fd5b813561201b8161276b565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b5f81356127e681612332565b6001600160a01b0316835260208201356127ff81612332565b6001600160a01b03166020840152604082810135908401526060820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe101811261284c575f80fd5b820160208101903567ffffffffffffffff811115612868575f80fd5b803603821315612876575f80fd5b60806060860152612132608086018284612793565b6001600160a01b03871681526001600160a01b038616602082015284604082015260a060608201525f6128c260a083018587612793565b82810360808401526128d481856127da565b9998505050505050505050565b848152606060208201525f6128fa606083018587612793565b828103604084015261290c81856127da565b979650505050505050565b5f873561292381612332565b6001600160a01b0381168352506001600160a01b03871660208301526001600160a01b038616604083015284606083015260a0608083015261296960a083018486612793565b98975050505050505050565b838152604060208201525f612132604083018486612793565b5f6020828403121561299e575f80fd5b5051919050565b8181038181111561074e577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b606081525f6129f0606083018789612793565b8560208401528281036040840152612969818587612793565b5f60208284031215612a19575f80fd5b815161201b8161276b565b5f82518060208501845e5f92019182525091905056fea2646970667358221220c0ce30ad69ab5eefe81c48c0aad8edbd8f6eb7e432ae6130c4b3acce21529d1664736f6c634300081a0033",
}

// ERC20CustodyABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20CustodyMetaData.ABI instead.
var ERC20CustodyABI = ERC20CustodyMetaData.ABI

// ERC20CustodyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20CustodyMetaData.Bin instead.
var ERC20CustodyBin = ERC20CustodyMetaData.Bin

// DeployERC20Custody deploys a new Ethereum contract, binding an instance of ERC20Custody to it.
func DeployERC20Custody(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Custody, error) {
	parsed, err := ERC20CustodyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20CustodyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Custody{ERC20CustodyCaller: ERC20CustodyCaller{contract: contract}, ERC20CustodyTransactor: ERC20CustodyTransactor{contract: contract}, ERC20CustodyFilterer: ERC20CustodyFilterer{contract: contract}}, nil
}

// ERC20Custody is an auto generated Go binding around an Ethereum contract.
type ERC20Custody struct {
	ERC20CustodyCaller     // Read-only binding to the contract
	ERC20CustodyTransactor // Write-only binding to the contract
	ERC20CustodyFilterer   // Log filterer for contract events
}

// ERC20CustodyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20CustodyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20CustodyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20CustodyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20CustodySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20CustodySession struct {
	Contract     *ERC20Custody     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CustodyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CustodyCallerSession struct {
	Contract *ERC20CustodyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC20CustodyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20CustodyTransactorSession struct {
	Contract     *ERC20CustodyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC20CustodyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20CustodyRaw struct {
	Contract *ERC20Custody // Generic contract binding to access the raw methods on
}

// ERC20CustodyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CustodyCallerRaw struct {
	Contract *ERC20CustodyCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20CustodyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20CustodyTransactorRaw struct {
	Contract *ERC20CustodyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Custody creates a new instance of ERC20Custody, bound to a specific deployed contract.
func NewERC20Custody(address common.Address, backend bind.ContractBackend) (*ERC20Custody, error) {
	contract, err := bindERC20Custody(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Custody{ERC20CustodyCaller: ERC20CustodyCaller{contract: contract}, ERC20CustodyTransactor: ERC20CustodyTransactor{contract: contract}, ERC20CustodyFilterer: ERC20CustodyFilterer{contract: contract}}, nil
}

// NewERC20CustodyCaller creates a new read-only instance of ERC20Custody, bound to a specific deployed contract.
func NewERC20CustodyCaller(address common.Address, caller bind.ContractCaller) (*ERC20CustodyCaller, error) {
	contract, err := bindERC20Custody(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyCaller{contract: contract}, nil
}

// NewERC20CustodyTransactor creates a new write-only instance of ERC20Custody, bound to a specific deployed contract.
func NewERC20CustodyTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20CustodyTransactor, error) {
	contract, err := bindERC20Custody(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyTransactor{contract: contract}, nil
}

// NewERC20CustodyFilterer creates a new log filterer instance of ERC20Custody, bound to a specific deployed contract.
func NewERC20CustodyFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20CustodyFilterer, error) {
	contract, err := bindERC20Custody(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyFilterer{contract: contract}, nil
}

// bindERC20Custody binds a generic wrapper to an already deployed contract.
func bindERC20Custody(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20CustodyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Custody *ERC20CustodyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Custody.Contract.ERC20CustodyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Custody *ERC20CustodyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Custody.Contract.ERC20CustodyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Custody *ERC20CustodyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Custody.Contract.ERC20CustodyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Custody *ERC20CustodyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Custody.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Custody *ERC20CustodyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Custody.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Custody *ERC20CustodyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Custody.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20Custody *ERC20CustodyCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20Custody *ERC20CustodySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20Custody.Contract.DEFAULTADMINROLE(&_ERC20Custody.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20Custody *ERC20CustodyCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20Custody.Contract.DEFAULTADMINROLE(&_ERC20Custody.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20Custody *ERC20CustodyCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20Custody *ERC20CustodySession) PAUSERROLE() ([32]byte, error) {
	return _ERC20Custody.Contract.PAUSERROLE(&_ERC20Custody.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20Custody *ERC20CustodyCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ERC20Custody.Contract.PAUSERROLE(&_ERC20Custody.CallOpts)
}

// STRANDEDFUNDSREFUNDER is a free data retrieval call binding the contract method 0x9dbc9b15.
//
// Solidity: function STRANDED_FUNDS_REFUNDER() view returns(address)
func (_ERC20Custody *ERC20CustodyCaller) STRANDEDFUNDSREFUNDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "STRANDED_FUNDS_REFUNDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STRANDEDFUNDSREFUNDER is a free data retrieval call binding the contract method 0x9dbc9b15.
//
// Solidity: function STRANDED_FUNDS_REFUNDER() view returns(address)
func (_ERC20Custody *ERC20CustodySession) STRANDEDFUNDSREFUNDER() (common.Address, error) {
	return _ERC20Custody.Contract.STRANDEDFUNDSREFUNDER(&_ERC20Custody.CallOpts)
}

// STRANDEDFUNDSREFUNDER is a free data retrieval call binding the contract method 0x9dbc9b15.
//
// Solidity: function STRANDED_FUNDS_REFUNDER() view returns(address)
func (_ERC20Custody *ERC20CustodyCallerSession) STRANDEDFUNDSREFUNDER() (common.Address, error) {
	return _ERC20Custody.Contract.STRANDEDFUNDSREFUNDER(&_ERC20Custody.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ERC20Custody *ERC20CustodyCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ERC20Custody *ERC20CustodySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ERC20Custody.Contract.UPGRADEINTERFACEVERSION(&_ERC20Custody.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ERC20Custody *ERC20CustodyCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ERC20Custody.Contract.UPGRADEINTERFACEVERSION(&_ERC20Custody.CallOpts)
}

// WHITELISTERROLE is a free data retrieval call binding the contract method 0x570618e1.
//
// Solidity: function WHITELISTER_ROLE() view returns(bytes32)
func (_ERC20Custody *ERC20CustodyCaller) WHITELISTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "WHITELISTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WHITELISTERROLE is a free data retrieval call binding the contract method 0x570618e1.
//
// Solidity: function WHITELISTER_ROLE() view returns(bytes32)
func (_ERC20Custody *ERC20CustodySession) WHITELISTERROLE() ([32]byte, error) {
	return _ERC20Custody.Contract.WHITELISTERROLE(&_ERC20Custody.CallOpts)
}

// WHITELISTERROLE is a free data retrieval call binding the contract method 0x570618e1.
//
// Solidity: function WHITELISTER_ROLE() view returns(bytes32)
func (_ERC20Custody *ERC20CustodyCallerSession) WHITELISTERROLE() ([32]byte, error) {
	return _ERC20Custody.Contract.WHITELISTERROLE(&_ERC20Custody.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ERC20Custody *ERC20CustodyCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ERC20Custody *ERC20CustodySession) WITHDRAWERROLE() ([32]byte, error) {
	return _ERC20Custody.Contract.WITHDRAWERROLE(&_ERC20Custody.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ERC20Custody *ERC20CustodyCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ERC20Custody.Contract.WITHDRAWERROLE(&_ERC20Custody.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ERC20Custody *ERC20CustodyCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ERC20Custody *ERC20CustodySession) Gateway() (common.Address, error) {
	return _ERC20Custody.Contract.Gateway(&_ERC20Custody.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ERC20Custody *ERC20CustodyCallerSession) Gateway() (common.Address, error) {
	return _ERC20Custody.Contract.Gateway(&_ERC20Custody.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20Custody *ERC20CustodyCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20Custody *ERC20CustodySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC20Custody.Contract.GetRoleAdmin(&_ERC20Custody.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ERC20Custody *ERC20CustodyCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ERC20Custody.Contract.GetRoleAdmin(&_ERC20Custody.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20Custody *ERC20CustodyCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20Custody *ERC20CustodySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC20Custody.Contract.HasRole(&_ERC20Custody.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20Custody *ERC20CustodyCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC20Custody.Contract.HasRole(&_ERC20Custody.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Custody *ERC20CustodyCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Custody *ERC20CustodySession) Paused() (bool, error) {
	return _ERC20Custody.Contract.Paused(&_ERC20Custody.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC20Custody *ERC20CustodyCallerSession) Paused() (bool, error) {
	return _ERC20Custody.Contract.Paused(&_ERC20Custody.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ERC20Custody *ERC20CustodyCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ERC20Custody *ERC20CustodySession) ProxiableUUID() ([32]byte, error) {
	return _ERC20Custody.Contract.ProxiableUUID(&_ERC20Custody.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ERC20Custody *ERC20CustodyCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ERC20Custody.Contract.ProxiableUUID(&_ERC20Custody.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC20Custody *ERC20CustodyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC20Custody *ERC20CustodySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC20Custody.Contract.SupportsInterface(&_ERC20Custody.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC20Custody *ERC20CustodyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC20Custody.Contract.SupportsInterface(&_ERC20Custody.CallOpts, interfaceId)
}

// SupportsLegacy is a free data retrieval call binding the contract method 0x252f07bf.
//
// Solidity: function supportsLegacy() view returns(bool)
func (_ERC20Custody *ERC20CustodyCaller) SupportsLegacy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "supportsLegacy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsLegacy is a free data retrieval call binding the contract method 0x252f07bf.
//
// Solidity: function supportsLegacy() view returns(bool)
func (_ERC20Custody *ERC20CustodySession) SupportsLegacy() (bool, error) {
	return _ERC20Custody.Contract.SupportsLegacy(&_ERC20Custody.CallOpts)
}

// SupportsLegacy is a free data retrieval call binding the contract method 0x252f07bf.
//
// Solidity: function supportsLegacy() view returns(bool)
func (_ERC20Custody *ERC20CustodyCallerSession) SupportsLegacy() (bool, error) {
	return _ERC20Custody.Contract.SupportsLegacy(&_ERC20Custody.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ERC20Custody *ERC20CustodyCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ERC20Custody *ERC20CustodySession) TssAddress() (common.Address, error) {
	return _ERC20Custody.Contract.TssAddress(&_ERC20Custody.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ERC20Custody *ERC20CustodyCallerSession) TssAddress() (common.Address, error) {
	return _ERC20Custody.Contract.TssAddress(&_ERC20Custody.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_ERC20Custody *ERC20CustodyCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Custody.contract.Call(opts, &out, "whitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_ERC20Custody *ERC20CustodySession) Whitelisted(arg0 common.Address) (bool, error) {
	return _ERC20Custody.Contract.Whitelisted(&_ERC20Custody.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_ERC20Custody *ERC20CustodyCallerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _ERC20Custody.Contract.Whitelisted(&_ERC20Custody.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xe609055e.
//
// Solidity: function deposit(bytes recipient, address asset, uint256 amount, bytes message) returns()
func (_ERC20Custody *ERC20CustodyTransactor) Deposit(opts *bind.TransactOpts, recipient []byte, asset common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "deposit", recipient, asset, amount, message)
}

// Deposit is a paid mutator transaction binding the contract method 0xe609055e.
//
// Solidity: function deposit(bytes recipient, address asset, uint256 amount, bytes message) returns()
func (_ERC20Custody *ERC20CustodySession) Deposit(recipient []byte, asset common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Deposit(&_ERC20Custody.TransactOpts, recipient, asset, amount, message)
}

// Deposit is a paid mutator transaction binding the contract method 0xe609055e.
//
// Solidity: function deposit(bytes recipient, address asset, uint256 amount, bytes message) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) Deposit(recipient []byte, asset common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Deposit(&_ERC20Custody.TransactOpts, recipient, asset, amount, message)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20Custody *ERC20CustodyTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20Custody *ERC20CustodySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.GrantRole(&_ERC20Custody.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.GrantRole(&_ERC20Custody.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address gateway_, address tssAddress_, address admin_) returns()
func (_ERC20Custody *ERC20CustodyTransactor) Initialize(opts *bind.TransactOpts, gateway_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "initialize", gateway_, tssAddress_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address gateway_, address tssAddress_, address admin_) returns()
func (_ERC20Custody *ERC20CustodySession) Initialize(gateway_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Initialize(&_ERC20Custody.TransactOpts, gateway_, tssAddress_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address gateway_, address tssAddress_, address admin_) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) Initialize(gateway_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Initialize(&_ERC20Custody.TransactOpts, gateway_, tssAddress_, admin_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Custody *ERC20CustodyTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Custody *ERC20CustodySession) Pause() (*types.Transaction, error) {
	return _ERC20Custody.Contract.Pause(&_ERC20Custody.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) Pause() (*types.Transaction, error) {
	return _ERC20Custody.Contract.Pause(&_ERC20Custody.TransactOpts)
}

// RefundStrandedFunds is a paid mutator transaction binding the contract method 0x3e1b9644.
//
// Solidity: function refundStrandedFunds(address to, address token, uint256 amount) returns()
func (_ERC20Custody *ERC20CustodyTransactor) RefundStrandedFunds(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "refundStrandedFunds", to, token, amount)
}

// RefundStrandedFunds is a paid mutator transaction binding the contract method 0x3e1b9644.
//
// Solidity: function refundStrandedFunds(address to, address token, uint256 amount) returns()
func (_ERC20Custody *ERC20CustodySession) RefundStrandedFunds(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Custody.Contract.RefundStrandedFunds(&_ERC20Custody.TransactOpts, to, token, amount)
}

// RefundStrandedFunds is a paid mutator transaction binding the contract method 0x3e1b9644.
//
// Solidity: function refundStrandedFunds(address to, address token, uint256 amount) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) RefundStrandedFunds(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Custody.Contract.RefundStrandedFunds(&_ERC20Custody.TransactOpts, to, token, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ERC20Custody *ERC20CustodyTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ERC20Custody *ERC20CustodySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.RenounceRole(&_ERC20Custody.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.RenounceRole(&_ERC20Custody.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20Custody *ERC20CustodyTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20Custody *ERC20CustodySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.RevokeRole(&_ERC20Custody.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.RevokeRole(&_ERC20Custody.TransactOpts, role, account)
}

// SetSupportsLegacy is a paid mutator transaction binding the contract method 0xeab103df.
//
// Solidity: function setSupportsLegacy(bool _supportsLegacy) returns()
func (_ERC20Custody *ERC20CustodyTransactor) SetSupportsLegacy(opts *bind.TransactOpts, _supportsLegacy bool) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "setSupportsLegacy", _supportsLegacy)
}

// SetSupportsLegacy is a paid mutator transaction binding the contract method 0xeab103df.
//
// Solidity: function setSupportsLegacy(bool _supportsLegacy) returns()
func (_ERC20Custody *ERC20CustodySession) SetSupportsLegacy(_supportsLegacy bool) (*types.Transaction, error) {
	return _ERC20Custody.Contract.SetSupportsLegacy(&_ERC20Custody.TransactOpts, _supportsLegacy)
}

// SetSupportsLegacy is a paid mutator transaction binding the contract method 0xeab103df.
//
// Solidity: function setSupportsLegacy(bool _supportsLegacy) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) SetSupportsLegacy(_supportsLegacy bool) (*types.Transaction, error) {
	return _ERC20Custody.Contract.SetSupportsLegacy(&_ERC20Custody.TransactOpts, _supportsLegacy)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20Custody *ERC20CustodyTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20Custody *ERC20CustodySession) Unpause() (*types.Transaction, error) {
	return _ERC20Custody.Contract.Unpause(&_ERC20Custody.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) Unpause() (*types.Transaction, error) {
	return _ERC20Custody.Contract.Unpause(&_ERC20Custody.TransactOpts)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address token) returns()
func (_ERC20Custody *ERC20CustodyTransactor) Unwhitelist(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "unwhitelist", token)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address token) returns()
func (_ERC20Custody *ERC20CustodySession) Unwhitelist(token common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Unwhitelist(&_ERC20Custody.TransactOpts, token)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address token) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) Unwhitelist(token common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Unwhitelist(&_ERC20Custody.TransactOpts, token)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ERC20Custody *ERC20CustodyTransactor) UpdateTSSAddress(opts *bind.TransactOpts, newTSSAddress common.Address) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "updateTSSAddress", newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ERC20Custody *ERC20CustodySession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.UpdateTSSAddress(&_ERC20Custody.TransactOpts, newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.UpdateTSSAddress(&_ERC20Custody.TransactOpts, newTSSAddress)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ERC20Custody *ERC20CustodyTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ERC20Custody *ERC20CustodySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Custody.Contract.UpgradeToAndCall(&_ERC20Custody.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ERC20Custody.Contract.UpgradeToAndCall(&_ERC20Custody.TransactOpts, newImplementation, data)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address token) returns()
func (_ERC20Custody *ERC20CustodyTransactor) Whitelist(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "whitelist", token)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address token) returns()
func (_ERC20Custody *ERC20CustodySession) Whitelist(token common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Whitelist(&_ERC20Custody.TransactOpts, token)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address token) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) Whitelist(token common.Address) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Whitelist(&_ERC20Custody.TransactOpts, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address to, address token, uint256 amount) returns()
func (_ERC20Custody *ERC20CustodyTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "withdraw", to, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address to, address token, uint256 amount) returns()
func (_ERC20Custody *ERC20CustodySession) Withdraw(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Withdraw(&_ERC20Custody.TransactOpts, to, token, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address to, address token, uint256 amount) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) Withdraw(to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Custody.Contract.Withdraw(&_ERC20Custody.TransactOpts, to, token, amount)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0xad081852.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, address token, uint256 amount, bytes data) returns()
func (_ERC20Custody *ERC20CustodyTransactor) WithdrawAndCall(opts *bind.TransactOpts, messageContext MessageContext, to common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "withdrawAndCall", messageContext, to, token, amount, data)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0xad081852.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, address token, uint256 amount, bytes data) returns()
func (_ERC20Custody *ERC20CustodySession) WithdrawAndCall(messageContext MessageContext, to common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Custody.Contract.WithdrawAndCall(&_ERC20Custody.TransactOpts, messageContext, to, token, amount, data)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0xad081852.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, address token, uint256 amount, bytes data) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) WithdrawAndCall(messageContext MessageContext, to common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC20Custody.Contract.WithdrawAndCall(&_ERC20Custody.TransactOpts, messageContext, to, token, amount, data)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x99a3c356.
//
// Solidity: function withdrawAndRevert(address to, address token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_ERC20Custody *ERC20CustodyTransactor) WithdrawAndRevert(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ERC20Custody.contract.Transact(opts, "withdrawAndRevert", to, token, amount, data, revertContext)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x99a3c356.
//
// Solidity: function withdrawAndRevert(address to, address token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_ERC20Custody *ERC20CustodySession) WithdrawAndRevert(to common.Address, token common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ERC20Custody.Contract.WithdrawAndRevert(&_ERC20Custody.TransactOpts, to, token, amount, data, revertContext)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x99a3c356.
//
// Solidity: function withdrawAndRevert(address to, address token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext) returns()
func (_ERC20Custody *ERC20CustodyTransactorSession) WithdrawAndRevert(to common.Address, token common.Address, amount *big.Int, data []byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ERC20Custody.Contract.WithdrawAndRevert(&_ERC20Custody.TransactOpts, to, token, amount, data, revertContext)
}

// ERC20CustodyDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ERC20Custody contract.
type ERC20CustodyDepositedIterator struct {
	Event *ERC20CustodyDeposited // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyDeposited)
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
		it.Event = new(ERC20CustodyDeposited)
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
func (it *ERC20CustodyDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyDeposited represents a Deposited event raised by the ERC20Custody contract.
type ERC20CustodyDeposited struct {
	Recipient []byte
	Asset     common.Address
	Amount    *big.Int
	Message   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_ERC20Custody *ERC20CustodyFilterer) FilterDeposited(opts *bind.FilterOpts, asset []common.Address) (*ERC20CustodyDepositedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyDepositedIterator{contract: _ERC20Custody.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1dafa057cc5c3bccb5ad974129a2bccd3c74002d9dfd7062404ba9523b18d6ae.
//
// Solidity: event Deposited(bytes recipient, address indexed asset, uint256 amount, bytes message)
func (_ERC20Custody *ERC20CustodyFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ERC20CustodyDeposited, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "Deposited", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyDeposited)
				if err := _ERC20Custody.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_ERC20Custody *ERC20CustodyFilterer) ParseDeposited(log types.Log) (*ERC20CustodyDeposited, error) {
	event := new(ERC20CustodyDeposited)
	if err := _ERC20Custody.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20Custody contract.
type ERC20CustodyInitializedIterator struct {
	Event *ERC20CustodyInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyInitialized)
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
		it.Event = new(ERC20CustodyInitialized)
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
func (it *ERC20CustodyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyInitialized represents a Initialized event raised by the ERC20Custody contract.
type ERC20CustodyInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20Custody *ERC20CustodyFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20CustodyInitializedIterator, error) {

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyInitializedIterator{contract: _ERC20Custody.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20Custody *ERC20CustodyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20CustodyInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyInitialized)
				if err := _ERC20Custody.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ERC20Custody *ERC20CustodyFilterer) ParseInitialized(log types.Log) (*ERC20CustodyInitialized, error) {
	event := new(ERC20CustodyInitialized)
	if err := _ERC20Custody.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC20Custody contract.
type ERC20CustodyPausedIterator struct {
	Event *ERC20CustodyPaused // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyPaused)
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
		it.Event = new(ERC20CustodyPaused)
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
func (it *ERC20CustodyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyPaused represents a Paused event raised by the ERC20Custody contract.
type ERC20CustodyPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20Custody *ERC20CustodyFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC20CustodyPausedIterator, error) {

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyPausedIterator{contract: _ERC20Custody.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC20Custody *ERC20CustodyFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC20CustodyPaused) (event.Subscription, error) {

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyPaused)
				if err := _ERC20Custody.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ERC20Custody *ERC20CustodyFilterer) ParsePaused(log types.Log) (*ERC20CustodyPaused, error) {
	event := new(ERC20CustodyPaused)
	if err := _ERC20Custody.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ERC20Custody contract.
type ERC20CustodyRoleAdminChangedIterator struct {
	Event *ERC20CustodyRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyRoleAdminChanged)
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
		it.Event = new(ERC20CustodyRoleAdminChanged)
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
func (it *ERC20CustodyRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyRoleAdminChanged represents a RoleAdminChanged event raised by the ERC20Custody contract.
type ERC20CustodyRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC20Custody *ERC20CustodyFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ERC20CustodyRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyRoleAdminChangedIterator{contract: _ERC20Custody.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ERC20Custody *ERC20CustodyFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC20CustodyRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyRoleAdminChanged)
				if err := _ERC20Custody.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ERC20Custody *ERC20CustodyFilterer) ParseRoleAdminChanged(log types.Log) (*ERC20CustodyRoleAdminChanged, error) {
	event := new(ERC20CustodyRoleAdminChanged)
	if err := _ERC20Custody.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ERC20Custody contract.
type ERC20CustodyRoleGrantedIterator struct {
	Event *ERC20CustodyRoleGranted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyRoleGranted)
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
		it.Event = new(ERC20CustodyRoleGranted)
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
func (it *ERC20CustodyRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyRoleGranted represents a RoleGranted event raised by the ERC20Custody contract.
type ERC20CustodyRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20Custody *ERC20CustodyFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC20CustodyRoleGrantedIterator, error) {

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

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyRoleGrantedIterator{contract: _ERC20Custody.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20Custody *ERC20CustodyFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyRoleGranted)
				if err := _ERC20Custody.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ERC20Custody *ERC20CustodyFilterer) ParseRoleGranted(log types.Log) (*ERC20CustodyRoleGranted, error) {
	event := new(ERC20CustodyRoleGranted)
	if err := _ERC20Custody.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ERC20Custody contract.
type ERC20CustodyRoleRevokedIterator struct {
	Event *ERC20CustodyRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyRoleRevoked)
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
		it.Event = new(ERC20CustodyRoleRevoked)
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
func (it *ERC20CustodyRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyRoleRevoked represents a RoleRevoked event raised by the ERC20Custody contract.
type ERC20CustodyRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20Custody *ERC20CustodyFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ERC20CustodyRoleRevokedIterator, error) {

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

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyRoleRevokedIterator{contract: _ERC20Custody.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ERC20Custody *ERC20CustodyFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ERC20CustodyRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyRoleRevoked)
				if err := _ERC20Custody.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ERC20Custody *ERC20CustodyFilterer) ParseRoleRevoked(log types.Log) (*ERC20CustodyRoleRevoked, error) {
	event := new(ERC20CustodyRoleRevoked)
	if err := _ERC20Custody.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyStrandedFundsRefundedIterator is returned from FilterStrandedFundsRefunded and is used to iterate over the raw logs and unpacked data for StrandedFundsRefunded events raised by the ERC20Custody contract.
type ERC20CustodyStrandedFundsRefundedIterator struct {
	Event *ERC20CustodyStrandedFundsRefunded // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyStrandedFundsRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyStrandedFundsRefunded)
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
		it.Event = new(ERC20CustodyStrandedFundsRefunded)
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
func (it *ERC20CustodyStrandedFundsRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyStrandedFundsRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyStrandedFundsRefunded represents a StrandedFundsRefunded event raised by the ERC20Custody contract.
type ERC20CustodyStrandedFundsRefunded struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStrandedFundsRefunded is a free log retrieval operation binding the contract event 0x7b73729b322c3c294385b7454f9895d7975135d24a54310b0e4542962bd089ab.
//
// Solidity: event StrandedFundsRefunded(address indexed to, address indexed token, uint256 amount)
func (_ERC20Custody *ERC20CustodyFilterer) FilterStrandedFundsRefunded(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ERC20CustodyStrandedFundsRefundedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "StrandedFundsRefunded", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyStrandedFundsRefundedIterator{contract: _ERC20Custody.contract, event: "StrandedFundsRefunded", logs: logs, sub: sub}, nil
}

// WatchStrandedFundsRefunded is a free log subscription operation binding the contract event 0x7b73729b322c3c294385b7454f9895d7975135d24a54310b0e4542962bd089ab.
//
// Solidity: event StrandedFundsRefunded(address indexed to, address indexed token, uint256 amount)
func (_ERC20Custody *ERC20CustodyFilterer) WatchStrandedFundsRefunded(opts *bind.WatchOpts, sink chan<- *ERC20CustodyStrandedFundsRefunded, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "StrandedFundsRefunded", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyStrandedFundsRefunded)
				if err := _ERC20Custody.contract.UnpackLog(event, "StrandedFundsRefunded", log); err != nil {
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

// ParseStrandedFundsRefunded is a log parse operation binding the contract event 0x7b73729b322c3c294385b7454f9895d7975135d24a54310b0e4542962bd089ab.
//
// Solidity: event StrandedFundsRefunded(address indexed to, address indexed token, uint256 amount)
func (_ERC20Custody *ERC20CustodyFilterer) ParseStrandedFundsRefunded(log types.Log) (*ERC20CustodyStrandedFundsRefunded, error) {
	event := new(ERC20CustodyStrandedFundsRefunded)
	if err := _ERC20Custody.contract.UnpackLog(event, "StrandedFundsRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC20Custody contract.
type ERC20CustodyUnpausedIterator struct {
	Event *ERC20CustodyUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUnpaused)
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
		it.Event = new(ERC20CustodyUnpaused)
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
func (it *ERC20CustodyUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUnpaused represents a Unpaused event raised by the ERC20Custody contract.
type ERC20CustodyUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20Custody *ERC20CustodyFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC20CustodyUnpausedIterator, error) {

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUnpausedIterator{contract: _ERC20Custody.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC20Custody *ERC20CustodyFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUnpaused)
				if err := _ERC20Custody.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ERC20Custody *ERC20CustodyFilterer) ParseUnpaused(log types.Log) (*ERC20CustodyUnpaused, error) {
	event := new(ERC20CustodyUnpaused)
	if err := _ERC20Custody.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUnwhitelistedIterator is returned from FilterUnwhitelisted and is used to iterate over the raw logs and unpacked data for Unwhitelisted events raised by the ERC20Custody contract.
type ERC20CustodyUnwhitelistedIterator struct {
	Event *ERC20CustodyUnwhitelisted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUnwhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUnwhitelisted)
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
		it.Event = new(ERC20CustodyUnwhitelisted)
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
func (it *ERC20CustodyUnwhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUnwhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUnwhitelisted represents a Unwhitelisted event raised by the ERC20Custody contract.
type ERC20CustodyUnwhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnwhitelisted is a free log retrieval operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_ERC20Custody *ERC20CustodyFilterer) FilterUnwhitelisted(opts *bind.FilterOpts, token []common.Address) (*ERC20CustodyUnwhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUnwhitelistedIterator{contract: _ERC20Custody.contract, event: "Unwhitelisted", logs: logs, sub: sub}, nil
}

// WatchUnwhitelisted is a free log subscription operation binding the contract event 0x51085ddf9ebdded84b76e829eb58c4078e4b5bdf97d9a94723f336039da46791.
//
// Solidity: event Unwhitelisted(address indexed token)
func (_ERC20Custody *ERC20CustodyFilterer) WatchUnwhitelisted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUnwhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "Unwhitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUnwhitelisted)
				if err := _ERC20Custody.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
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
func (_ERC20Custody *ERC20CustodyFilterer) ParseUnwhitelisted(log types.Log) (*ERC20CustodyUnwhitelisted, error) {
	event := new(ERC20CustodyUnwhitelisted)
	if err := _ERC20Custody.contract.UnpackLog(event, "Unwhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpdatedCustodyTSSAddressIterator is returned from FilterUpdatedCustodyTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedCustodyTSSAddress events raised by the ERC20Custody contract.
type ERC20CustodyUpdatedCustodyTSSAddressIterator struct {
	Event *ERC20CustodyUpdatedCustodyTSSAddress // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpdatedCustodyTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpdatedCustodyTSSAddress)
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
		it.Event = new(ERC20CustodyUpdatedCustodyTSSAddress)
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
func (it *ERC20CustodyUpdatedCustodyTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpdatedCustodyTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpdatedCustodyTSSAddress represents a UpdatedCustodyTSSAddress event raised by the ERC20Custody contract.
type ERC20CustodyUpdatedCustodyTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCustodyTSSAddress is a free log retrieval operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ERC20Custody *ERC20CustodyFilterer) FilterUpdatedCustodyTSSAddress(opts *bind.FilterOpts) (*ERC20CustodyUpdatedCustodyTSSAddressIterator, error) {

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpdatedCustodyTSSAddressIterator{contract: _ERC20Custody.contract, event: "UpdatedCustodyTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedCustodyTSSAddress is a free log subscription operation binding the contract event 0x4d3470c839d3c4dd664eec934b920c12fe0966e3185103dd40149496815df2b6.
//
// Solidity: event UpdatedCustodyTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ERC20Custody *ERC20CustodyFilterer) WatchUpdatedCustodyTSSAddress(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpdatedCustodyTSSAddress) (event.Subscription, error) {

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "UpdatedCustodyTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpdatedCustodyTSSAddress)
				if err := _ERC20Custody.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
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
func (_ERC20Custody *ERC20CustodyFilterer) ParseUpdatedCustodyTSSAddress(log types.Log) (*ERC20CustodyUpdatedCustodyTSSAddress, error) {
	event := new(ERC20CustodyUpdatedCustodyTSSAddress)
	if err := _ERC20Custody.contract.UnpackLog(event, "UpdatedCustodyTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC20Custody contract.
type ERC20CustodyUpgradedIterator struct {
	Event *ERC20CustodyUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyUpgraded)
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
		it.Event = new(ERC20CustodyUpgraded)
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
func (it *ERC20CustodyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyUpgraded represents a Upgraded event raised by the ERC20Custody contract.
type ERC20CustodyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC20Custody *ERC20CustodyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ERC20CustodyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyUpgradedIterator{contract: _ERC20Custody.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC20Custody *ERC20CustodyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC20CustodyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyUpgraded)
				if err := _ERC20Custody.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ERC20Custody *ERC20CustodyFilterer) ParseUpgraded(log types.Log) (*ERC20CustodyUpgraded, error) {
	event := new(ERC20CustodyUpgraded)
	if err := _ERC20Custody.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the ERC20Custody contract.
type ERC20CustodyWhitelistedIterator struct {
	Event *ERC20CustodyWhitelisted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyWhitelisted)
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
		it.Event = new(ERC20CustodyWhitelisted)
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
func (it *ERC20CustodyWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyWhitelisted represents a Whitelisted event raised by the ERC20Custody contract.
type ERC20CustodyWhitelisted struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_ERC20Custody *ERC20CustodyFilterer) FilterWhitelisted(opts *bind.FilterOpts, token []common.Address) (*ERC20CustodyWhitelistedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyWhitelistedIterator{contract: _ERC20Custody.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xaab7954e9d246b167ef88aeddad35209ca2489d95a8aeb59e288d9b19fae5a54.
//
// Solidity: event Whitelisted(address indexed token)
func (_ERC20Custody *ERC20CustodyFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyWhitelisted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "Whitelisted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyWhitelisted)
				if err := _ERC20Custody.contract.UnpackLog(event, "Whitelisted", log); err != nil {
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
func (_ERC20Custody *ERC20CustodyFilterer) ParseWhitelisted(log types.Log) (*ERC20CustodyWhitelisted, error) {
	event := new(ERC20CustodyWhitelisted)
	if err := _ERC20Custody.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ERC20Custody contract.
type ERC20CustodyWithdrawnIterator struct {
	Event *ERC20CustodyWithdrawn // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyWithdrawn)
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
		it.Event = new(ERC20CustodyWithdrawn)
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
func (it *ERC20CustodyWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyWithdrawn represents a Withdrawn event raised by the ERC20Custody contract.
type ERC20CustodyWithdrawn struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_ERC20Custody *ERC20CustodyFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ERC20CustodyWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyWithdrawnIterator{contract: _ERC20Custody.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed to, address indexed token, uint256 amount)
func (_ERC20Custody *ERC20CustodyFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ERC20CustodyWithdrawn, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "Withdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyWithdrawn)
				if err := _ERC20Custody.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_ERC20Custody *ERC20CustodyFilterer) ParseWithdrawn(log types.Log) (*ERC20CustodyWithdrawn, error) {
	event := new(ERC20CustodyWithdrawn)
	if err := _ERC20Custody.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the ERC20Custody contract.
type ERC20CustodyWithdrawnAndCalledIterator struct {
	Event *ERC20CustodyWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyWithdrawnAndCalled)
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
		it.Event = new(ERC20CustodyWithdrawnAndCalled)
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
func (it *ERC20CustodyWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the ERC20Custody contract.
type ERC20CustodyWithdrawnAndCalled struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_ERC20Custody *ERC20CustodyFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ERC20CustodyWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyWithdrawnAndCalledIterator{contract: _ERC20Custody.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x6478cbb6e28c0823c691dfd74c01c985634faddd4c401b990fe4ec26277ea8d5.
//
// Solidity: event WithdrawnAndCalled(address indexed to, address indexed token, uint256 amount, bytes data)
func (_ERC20Custody *ERC20CustodyFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *ERC20CustodyWithdrawnAndCalled, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyWithdrawnAndCalled)
				if err := _ERC20Custody.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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
func (_ERC20Custody *ERC20CustodyFilterer) ParseWithdrawnAndCalled(log types.Log) (*ERC20CustodyWithdrawnAndCalled, error) {
	event := new(ERC20CustodyWithdrawnAndCalled)
	if err := _ERC20Custody.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20CustodyWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the ERC20Custody contract.
type ERC20CustodyWithdrawnAndRevertedIterator struct {
	Event *ERC20CustodyWithdrawnAndReverted // Event containing the contract specifics and raw log

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
func (it *ERC20CustodyWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20CustodyWithdrawnAndReverted)
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
		it.Event = new(ERC20CustodyWithdrawnAndReverted)
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
func (it *ERC20CustodyWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20CustodyWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20CustodyWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the ERC20Custody contract.
type ERC20CustodyWithdrawnAndReverted struct {
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
func (_ERC20Custody *ERC20CustodyFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*ERC20CustodyWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20Custody.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ERC20CustodyWithdrawnAndRevertedIterator{contract: _ERC20Custody.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x7b53ec10a80164e60591c43d9c222e9354886981b880a3fba19c9ceb77fb9721.
//
// Solidity: event WithdrawnAndReverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ERC20Custody *ERC20CustodyFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *ERC20CustodyWithdrawnAndReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ERC20Custody.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20CustodyWithdrawnAndReverted)
				if err := _ERC20Custody.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
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
func (_ERC20Custody *ERC20CustodyFilterer) ParseWithdrawnAndReverted(log types.Log) (*ERC20CustodyWithdrawnAndReverted, error) {
	event := new(ERC20CustodyWithdrawnAndReverted)
	if err := _ERC20Custody.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
