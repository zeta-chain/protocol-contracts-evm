// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetaconnectornonnativeupgradetest

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

// ZetaConnectorNonNativeUpgradeTestMetaData contains all meta data concerning the ZetaConnectorNonNativeUpgradeTest contract.
var ZetaConnectorNonNativeUpgradeTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"gateway_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zetaToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveTokens\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxSupply\",\"inputs\":[{\"name\":\"maxSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"messageContext\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndRevert\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zetaToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxSupplyUpdated\",\"inputs\":[{\"name\":\"maxSupply\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedZetaConnectorTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsMaxSupply\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a060405230608052348015601357600080fd5b5060805161252a61003d60003960008181611428015281816114510152611627015261252a6000f3fe6080604052600436106101ac5760003560e01c80636f8b44b0116100ec578063a217fddf1161008a578063d547741f11610064578063d547741f1461057e578063d5abeb011461059e578063e63ab1e9146105b4578063f8c8765e146105e857600080fd5b8063a217fddf146104df578063a783c789146104f4578063ad3cb1cc1461052857600080fd5b80638456cb59116100c65780638456cb591461041157806385f438c11461042657806391d148541461045a578063950837aa146104bf57600080fd5b80636f8b44b0146103b15780636fb9a7af146103d1578063743e0c9b146103f157600080fd5b806336568abe1161015957806352d1902d1161013357806352d1902d146103255780635b1125911461033a5780635c975abb1461035a5780636f8728ad1461039157600080fd5b806336568abe146102dd5780633f4ba83a146102fd5780634f1ef2861461031257600080fd5b806321e093b11161018a57806321e093b114610240578063248a9ca3146102605780632f2ff15d146102bd57600080fd5b806301ffc9a7146101b1578063106e6290146101e6578063116191b614610208575b600080fd5b3480156101bd57600080fd5b506101d16101cc366004611e38565b610608565b60405190151581526020015b60405180910390f35b3480156101f257600080fd5b50610206610201366004611e96565b6106a1565b005b34801561021457600080fd5b50600054610228906001600160a01b031681565b6040516001600160a01b0390911681526020016101dd565b34801561024c57600080fd5b50600154610228906001600160a01b031681565b34801561026c57600080fd5b506102af61027b366004611ec9565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016101dd565b3480156102c957600080fd5b506102066102d8366004611ee2565b610758565b3480156102e957600080fd5b506102066102f8366004611ee2565b6107a2565b34801561030957600080fd5b506102066107ee565b610206610320366004611f3d565b610823565b34801561033157600080fd5b506102af610842565b34801561034657600080fd5b50600254610228906001600160a01b031681565b34801561036657600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166101d1565b34801561039d57600080fd5b506102066103ac36600461208d565b610871565b3480156103bd57600080fd5b506102066103cc366004611ec9565b6109c4565b3480156103dd57600080fd5b506102066103ec366004612125565b610a32565b3480156103fd57600080fd5b5061020661040c366004611ec9565b610b4a565b34801561041d57600080fd5b50610206610bd2565b34801561043257600080fd5b506102af7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b34801561046657600080fd5b506101d1610475366004611ee2565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156104cb57600080fd5b506102066104da3660046121a4565b610c04565b3480156104eb57600080fd5b506102af600081565b34801561050057600080fd5b506102af7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb81565b34801561053457600080fd5b506105716040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101dd91906121e3565b34801561058a57600080fd5b50610206610599366004611ee2565b610d97565b3480156105aa57600080fd5b506102af60035481565b3480156105c057600080fd5b506102af7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b3480156105f457600080fd5b50610206610603366004612234565b610ddb565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061069b57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6106a9610f86565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46106d381611007565b6106db611011565b6106e684848461106f565b836001600160a01b03167f3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee98460405161072191815260200190565b60405180910390a25061075360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461079281611007565b61079c83836111dc565b50505050565b6001600160a01b03811633146107e4576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61075382826112c9565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61081881611007565b61082061138d565b50565b61082b61141d565b610834826114ed565b61083e82826114f8565b5050565b600061084c61161c565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b610879610f86565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46108a381611007565b6108ab611011565b6000546108c2906001600160a01b0316878561106f565b6000546001546040517faa0c0fc10000000000000000000000000000000000000000000000000000000081526001600160a01b039283169263aa0c0fc192610919929116908b908b908b908b908a90600401612382565b600060405180830381600087803b15801561093357600080fd5b505af1158015610947573d6000803e3d6000fd5b50505050866001600160a01b03167f5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff08787878660405161098a94939291906123d9565b60405180910390a2506109bc60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050505050565b7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb6109ee81611007565b6109f6611011565b60038290556040518281527f7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c9060200160405180910390a15050565b610a3a610f86565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610a6481611007565b610a6c611011565b600054610a83906001600160a01b0316868461106f565b6000546001546040517f7bbe9afa0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831692637bbe9afa92610adb928c92909116908b908b908b908b90600401612410565b600060405180830381600087803b158015610af557600080fd5b505af1158015610b09573d6000803e3d6000fd5b50505050856001600160a01b03167f23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d86868660405161098a9392919061246b565b610b52611011565b6001546040517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018390526001600160a01b03909116906379cc679090604401600060405180830381600087803b158015610bb757600080fd5b505af1158015610bcb573d6000803e3d6000fd5b5050505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610bfc81611007565b61082061167e565b6000610c0f81611007565b6001600160a01b038216610c4f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254610c86907f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4906001600160a01b03166112c9565b50600254610cbe907f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb906001600160a01b03166112c9565b50610ce97f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4836111dc565b50610d147f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb836111dc565b50600254604080516001600160a01b03928316815291841660208301527f33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e910160405180910390a150600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610dd181611007565b61079c83836112c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff16600081158015610e265750825b905060008267ffffffffffffffff166001148015610e435750303b155b905081158015610e51575080155b15610e88576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001660011785558315610ee95784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b610ef5898989896116f7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6003558315610f7b5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2906020015b60405180910390a15b505050505050505050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611001576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b6108208133611a02565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff161561106d576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b600354600160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110e99190612485565b6110f3908461249e565b111561112b576040517fc30436e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040517f1e458bee0000000000000000000000000000000000000000000000000000000081526001600160a01b038581166004830152602482018590526044820184905290911690631e458bee90606401600060405180830381600087803b15801561119957600080fd5b505af11580156111ad573d6000803e3d6000fd5b50505050505050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff166112bf576000848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556112753390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061069b565b600091505061069b565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16156112bf576000848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061069b565b611395611a8f565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806114b657507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166114aa7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b1561106d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061083e81611007565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611570575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261156d91810190612485565b60015b6115b6576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114611612576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016115ad565b6107538383611aea565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461106d576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611686611011565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336113ff565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156117425750825b905060008267ffffffffffffffff16600114801561175f5750303b155b90508115801561176d575080155b156117a4576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156118055784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038916158061182257506001600160a01b038816155b8061183457506001600160a01b038716155b8061184657506001600160a01b038616155b1561187d576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611885611b40565b61188d611b48565b611895611b40565b61189d611b58565b600080546001600160a01b03808c167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316178355600180548c831690841617905560028054918b16919092161790556118f890876111dc565b506119237f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4886111dc565b5061194e7f0da06bffcb63442de88b7f8385468eaf51e47079d4fa96875938e2c27c451deb886111dc565b506119797f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a876111dc565b506119a47f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a886111dc565b508315610f7b5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602001610f72565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff1661083e576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044016115ad565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1661106d576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611af382611b68565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115611b38576107538282611c10565b61083e611c86565b61106d611cbe565b611b50611cbe565b61106d611d25565b611b60611cbe565b61106d611d2d565b806001600160a01b03163b600003611bb7576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016115ad565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051611c2d91906124d8565b600060405180830381855af49150503d8060008114611c68576040519150601f19603f3d011682016040523d82523d6000602084013e611c6d565b606091505b5091509150611c7d858383611d7e565b95945050505050565b341561106d576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff1661106d576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111b6611cbe565b611d35611cbe565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b606082611d9357611d8e82611df6565b611def565b8151158015611daa57506001600160a01b0384163b155b15611dec576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016115ad565b50805b9392505050565b805115611e065780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215611e4a57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611def57600080fd5b80356001600160a01b0381168114611e9157600080fd5b919050565b600080600060608486031215611eab57600080fd5b611eb484611e7a565b95602085013595506040909401359392505050565b600060208284031215611edb57600080fd5b5035919050565b60008060408385031215611ef557600080fd5b82359150611f0560208401611e7a565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060408385031215611f5057600080fd5b611f5983611e7a565b9150602083013567ffffffffffffffff811115611f7557600080fd5b8301601f81018513611f8657600080fd5b803567ffffffffffffffff811115611fa057611fa0611f0e565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561200c5761200c611f0e565b60405281815282820160200187101561202457600080fd5b816020840160208301376000602083830101528093505050509250929050565b60008083601f84011261205657600080fd5b50813567ffffffffffffffff81111561206e57600080fd5b60208301915083602082850101111561208657600080fd5b9250929050565b60008060008060008060a087890312156120a657600080fd5b6120af87611e7a565b955060208701359450604087013567ffffffffffffffff8111156120d257600080fd5b6120de89828a01612044565b90955093505060608701359150608087013567ffffffffffffffff81111561210557600080fd5b87016080818a03121561211757600080fd5b809150509295509295509295565b60008060008060008086880360a081121561213f57600080fd5b602081121561214d57600080fd5b5086955061215d60208801611e7a565b945060408701359350606087013567ffffffffffffffff81111561218057600080fd5b61218c89828a01612044565b979a9699509497949695608090950135949350505050565b6000602082840312156121b657600080fd5b611def82611e7a565b60005b838110156121da5781810151838201526020016121c2565b50506000910152565b60208152600082518060208401526122028160408501602087016121bf565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b6000806000806080858703121561224a57600080fd5b61225385611e7a565b935061226160208601611e7a565b925061226f60408601611e7a565b915061227d60608601611e7a565b905092959194509250565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6001600160a01b036122e282611e7a565b1682526001600160a01b036122f960208301611e7a565b1660208301526040818101359083015260006060820135368390037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe101811261234157600080fd5b820160208101903567ffffffffffffffff81111561235e57600080fd5b80360382131561236d57600080fd5b60806060860152611c7d608086018284612288565b6001600160a01b03871681526001600160a01b038616602082015284604082015260a0606082015260006123ba60a083018587612288565b82810360808401526123cc81856122d1565b9998505050505050505050565b8481526060602082015260006123f3606083018587612288565b828103604084015261240581856122d1565b979650505050505050565b6001600160a01b0361242188611e7a565b1681526001600160a01b03861660208201526001600160a01b038516604082015283606082015260a06080820152600061245f60a083018486612288565b98975050505050505050565b838152604060208201526000611c7d604083018486612288565b60006020828403121561249757600080fd5b5051919050565b8082018082111561069b577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082516124ea8184602087016121bf565b919091019291505056fea26469706673582212205b7d69ca110d73c61d90205bc8b310a76dae3263ff702f3c47a32da97121690964736f6c634300081a0033",
}

// ZetaConnectorNonNativeUpgradeTestABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaConnectorNonNativeUpgradeTestMetaData.ABI instead.
var ZetaConnectorNonNativeUpgradeTestABI = ZetaConnectorNonNativeUpgradeTestMetaData.ABI

// ZetaConnectorNonNativeUpgradeTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaConnectorNonNativeUpgradeTestMetaData.Bin instead.
var ZetaConnectorNonNativeUpgradeTestBin = ZetaConnectorNonNativeUpgradeTestMetaData.Bin

// DeployZetaConnectorNonNativeUpgradeTest deploys a new Ethereum contract, binding an instance of ZetaConnectorNonNativeUpgradeTest to it.
func DeployZetaConnectorNonNativeUpgradeTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZetaConnectorNonNativeUpgradeTest, error) {
	parsed, err := ZetaConnectorNonNativeUpgradeTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaConnectorNonNativeUpgradeTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaConnectorNonNativeUpgradeTest{ZetaConnectorNonNativeUpgradeTestCaller: ZetaConnectorNonNativeUpgradeTestCaller{contract: contract}, ZetaConnectorNonNativeUpgradeTestTransactor: ZetaConnectorNonNativeUpgradeTestTransactor{contract: contract}, ZetaConnectorNonNativeUpgradeTestFilterer: ZetaConnectorNonNativeUpgradeTestFilterer{contract: contract}}, nil
}

// ZetaConnectorNonNativeUpgradeTest is an auto generated Go binding around an Ethereum contract.
type ZetaConnectorNonNativeUpgradeTest struct {
	ZetaConnectorNonNativeUpgradeTestCaller     // Read-only binding to the contract
	ZetaConnectorNonNativeUpgradeTestTransactor // Write-only binding to the contract
	ZetaConnectorNonNativeUpgradeTestFilterer   // Log filterer for contract events
}

// ZetaConnectorNonNativeUpgradeTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeUpgradeTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonNativeUpgradeTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeUpgradeTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonNativeUpgradeTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaConnectorNonNativeUpgradeTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaConnectorNonNativeUpgradeTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaConnectorNonNativeUpgradeTestSession struct {
	Contract     *ZetaConnectorNonNativeUpgradeTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                      // Call options to use throughout this session
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ZetaConnectorNonNativeUpgradeTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaConnectorNonNativeUpgradeTestCallerSession struct {
	Contract *ZetaConnectorNonNativeUpgradeTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                            // Call options to use throughout this session
}

// ZetaConnectorNonNativeUpgradeTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaConnectorNonNativeUpgradeTestTransactorSession struct {
	Contract     *ZetaConnectorNonNativeUpgradeTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                            // Transaction auth options to use throughout this session
}

// ZetaConnectorNonNativeUpgradeTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaConnectorNonNativeUpgradeTestRaw struct {
	Contract *ZetaConnectorNonNativeUpgradeTest // Generic contract binding to access the raw methods on
}

// ZetaConnectorNonNativeUpgradeTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeUpgradeTestCallerRaw struct {
	Contract *ZetaConnectorNonNativeUpgradeTestCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaConnectorNonNativeUpgradeTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaConnectorNonNativeUpgradeTestTransactorRaw struct {
	Contract *ZetaConnectorNonNativeUpgradeTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaConnectorNonNativeUpgradeTest creates a new instance of ZetaConnectorNonNativeUpgradeTest, bound to a specific deployed contract.
func NewZetaConnectorNonNativeUpgradeTest(address common.Address, backend bind.ContractBackend) (*ZetaConnectorNonNativeUpgradeTest, error) {
	contract, err := bindZetaConnectorNonNativeUpgradeTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTest{ZetaConnectorNonNativeUpgradeTestCaller: ZetaConnectorNonNativeUpgradeTestCaller{contract: contract}, ZetaConnectorNonNativeUpgradeTestTransactor: ZetaConnectorNonNativeUpgradeTestTransactor{contract: contract}, ZetaConnectorNonNativeUpgradeTestFilterer: ZetaConnectorNonNativeUpgradeTestFilterer{contract: contract}}, nil
}

// NewZetaConnectorNonNativeUpgradeTestCaller creates a new read-only instance of ZetaConnectorNonNativeUpgradeTest, bound to a specific deployed contract.
func NewZetaConnectorNonNativeUpgradeTestCaller(address common.Address, caller bind.ContractCaller) (*ZetaConnectorNonNativeUpgradeTestCaller, error) {
	contract, err := bindZetaConnectorNonNativeUpgradeTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestCaller{contract: contract}, nil
}

// NewZetaConnectorNonNativeUpgradeTestTransactor creates a new write-only instance of ZetaConnectorNonNativeUpgradeTest, bound to a specific deployed contract.
func NewZetaConnectorNonNativeUpgradeTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaConnectorNonNativeUpgradeTestTransactor, error) {
	contract, err := bindZetaConnectorNonNativeUpgradeTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestTransactor{contract: contract}, nil
}

// NewZetaConnectorNonNativeUpgradeTestFilterer creates a new log filterer instance of ZetaConnectorNonNativeUpgradeTest, bound to a specific deployed contract.
func NewZetaConnectorNonNativeUpgradeTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaConnectorNonNativeUpgradeTestFilterer, error) {
	contract, err := bindZetaConnectorNonNativeUpgradeTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestFilterer{contract: contract}, nil
}

// bindZetaConnectorNonNativeUpgradeTest binds a generic wrapper to an already deployed contract.
func bindZetaConnectorNonNativeUpgradeTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaConnectorNonNativeUpgradeTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.ZetaConnectorNonNativeUpgradeTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.ZetaConnectorNonNativeUpgradeTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.ZetaConnectorNonNativeUpgradeTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeUpgradeTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.DEFAULTADMINROLE(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.DEFAULTADMINROLE(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeUpgradeTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.PAUSERROLE(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.PAUSERROLE(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeUpgradeTest.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.TSSROLE(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerSession) TSSROLE() ([32]byte, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.TSSROLE(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeUpgradeTest.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeUpgradeTest.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.WITHDRAWERROLE(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.WITHDRAWERROLE(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeUpgradeTest.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.Gateway(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerSession) Gateway() (common.Address, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.Gateway(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeUpgradeTest.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.GetRoleAdmin(&_ZetaConnectorNonNativeUpgradeTest.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.GetRoleAdmin(&_ZetaConnectorNonNativeUpgradeTest.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeUpgradeTest.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.HasRole(&_ZetaConnectorNonNativeUpgradeTest.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.HasRole(&_ZetaConnectorNonNativeUpgradeTest.CallOpts, role, account)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeUpgradeTest.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) MaxSupply() (*big.Int, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.MaxSupply(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerSession) MaxSupply() (*big.Int, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.MaxSupply(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeUpgradeTest.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) Paused() (bool, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.Paused(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerSession) Paused() (bool, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.Paused(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeUpgradeTest.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) ProxiableUUID() ([32]byte, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.ProxiableUUID(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.ProxiableUUID(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeUpgradeTest.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.SupportsInterface(&_ZetaConnectorNonNativeUpgradeTest.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.SupportsInterface(&_ZetaConnectorNonNativeUpgradeTest.CallOpts, interfaceId)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeUpgradeTest.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.TssAddress(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerSession) TssAddress() (common.Address, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.TssAddress(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCaller) ZetaToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaConnectorNonNativeUpgradeTest.contract.Call(opts, &out, "zetaToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.ZetaToken(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// ZetaToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function zetaToken() view returns(address)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestCallerSession) ZetaToken() (common.Address, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.ZetaToken(&_ZetaConnectorNonNativeUpgradeTest.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.GrantRole(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.GrantRole(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address gateway_, address zetaToken_, address tssAddress_, address admin_) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactor) Initialize(opts *bind.TransactOpts, gateway_ common.Address, zetaToken_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.contract.Transact(opts, "initialize", gateway_, zetaToken_, tssAddress_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address gateway_, address zetaToken_, address tssAddress_, address admin_) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) Initialize(gateway_ common.Address, zetaToken_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.Initialize(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, gateway_, zetaToken_, tssAddress_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address gateway_, address zetaToken_, address tssAddress_, address admin_) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorSession) Initialize(gateway_ common.Address, zetaToken_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.Initialize(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, gateway_, zetaToken_, tssAddress_, admin_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.Pause(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorSession) Pause() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.Pause(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactor) ReceiveTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.contract.Transact(opts, "receiveTokens", amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.ReceiveTokens(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.ReceiveTokens(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.RenounceRole(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.RenounceRole(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.RevokeRole(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.RevokeRole(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, role, account)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactor) SetMaxSupply(opts *bind.TransactOpts, maxSupply_ *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.contract.Transact(opts, "setMaxSupply", maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.SetMaxSupply(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.SetMaxSupply(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, maxSupply_)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.Unpause(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorSession) Unpause() (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.Unpause(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactor) UpdateTSSAddress(opts *bind.TransactOpts, newTSSAddress common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.contract.Transact(opts, "updateTSSAddress", newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.UpdateTSSAddress(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.UpdateTSSAddress(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, newTSSAddress)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.UpgradeToAndCall(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.UpgradeToAndCall(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.contract.Transact(opts, "withdraw", to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.Withdraw(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.Withdraw(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, to, amount, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x6fb9a7af.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactor) WithdrawAndCall(opts *bind.TransactOpts, messageContext MessageContext, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.contract.Transact(opts, "withdrawAndCall", messageContext, to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x6fb9a7af.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) WithdrawAndCall(messageContext MessageContext, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.WithdrawAndCall(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, messageContext, to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x6fb9a7af.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorSession) WithdrawAndCall(messageContext MessageContext, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.WithdrawAndCall(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, messageContext, to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x6f8728ad.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash, (address,address,uint256,bytes) revertContext) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactor) WithdrawAndRevert(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.contract.Transact(opts, "withdrawAndRevert", to, amount, data, internalSendHash, revertContext)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x6f8728ad.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash, (address,address,uint256,bytes) revertContext) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.WithdrawAndRevert(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, to, amount, data, internalSendHash, revertContext)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x6f8728ad.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash, (address,address,uint256,bytes) revertContext) returns()
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestTransactorSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte, revertContext RevertContext) (*types.Transaction, error) {
	return _ZetaConnectorNonNativeUpgradeTest.Contract.WithdrawAndRevert(&_ZetaConnectorNonNativeUpgradeTest.TransactOpts, to, amount, data, internalSendHash, revertContext)
}

// ZetaConnectorNonNativeUpgradeTestInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestInitializedIterator struct {
	Event *ZetaConnectorNonNativeUpgradeTestInitialized // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeUpgradeTestInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeUpgradeTestInitialized)
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
		it.Event = new(ZetaConnectorNonNativeUpgradeTestInitialized)
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
func (it *ZetaConnectorNonNativeUpgradeTestInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeUpgradeTestInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeUpgradeTestInitialized represents a Initialized event raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZetaConnectorNonNativeUpgradeTestInitializedIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestInitializedIterator{contract: _ZetaConnectorNonNativeUpgradeTest.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeUpgradeTestInitialized) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeUpgradeTestInitialized)
				if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) ParseInitialized(log types.Log) (*ZetaConnectorNonNativeUpgradeTestInitialized, error) {
	event := new(ZetaConnectorNonNativeUpgradeTestInitialized)
	if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdatedIterator is returned from FilterMaxSupplyUpdated and is used to iterate over the raw logs and unpacked data for MaxSupplyUpdated events raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdatedIterator struct {
	Event *ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdated // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdated)
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
		it.Event = new(ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdated)
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
func (it *ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdated represents a MaxSupplyUpdated event raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdated struct {
	MaxSupply *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyUpdated is a free log retrieval operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 maxSupply)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) FilterMaxSupplyUpdated(opts *bind.FilterOpts) (*ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdatedIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.FilterLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdatedIterator{contract: _ZetaConnectorNonNativeUpgradeTest.contract, event: "MaxSupplyUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyUpdated is a free log subscription operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 maxSupply)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) WatchMaxSupplyUpdated(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdated) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.WatchLogs(opts, "MaxSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdated)
				if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
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

// ParseMaxSupplyUpdated is a log parse operation binding the contract event 0x7810bd47de260c3e9ee10061cf438099dd12256c79485f12f94dbccc981e806c.
//
// Solidity: event MaxSupplyUpdated(uint256 maxSupply)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) ParseMaxSupplyUpdated(log types.Log) (*ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdated, error) {
	event := new(ZetaConnectorNonNativeUpgradeTestMaxSupplyUpdated)
	if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "MaxSupplyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeUpgradeTestPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestPausedIterator struct {
	Event *ZetaConnectorNonNativeUpgradeTestPaused // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeUpgradeTestPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeUpgradeTestPaused)
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
		it.Event = new(ZetaConnectorNonNativeUpgradeTestPaused)
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
func (it *ZetaConnectorNonNativeUpgradeTestPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeUpgradeTestPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeUpgradeTestPaused represents a Paused event raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) FilterPaused(opts *bind.FilterOpts) (*ZetaConnectorNonNativeUpgradeTestPausedIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestPausedIterator{contract: _ZetaConnectorNonNativeUpgradeTest.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeUpgradeTestPaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeUpgradeTestPaused)
				if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) ParsePaused(log types.Log) (*ZetaConnectorNonNativeUpgradeTestPaused, error) {
	event := new(ZetaConnectorNonNativeUpgradeTestPaused)
	if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeUpgradeTestRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestRoleAdminChangedIterator struct {
	Event *ZetaConnectorNonNativeUpgradeTestRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeUpgradeTestRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeUpgradeTestRoleAdminChanged)
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
		it.Event = new(ZetaConnectorNonNativeUpgradeTestRoleAdminChanged)
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
func (it *ZetaConnectorNonNativeUpgradeTestRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeUpgradeTestRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeUpgradeTestRoleAdminChanged represents a RoleAdminChanged event raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZetaConnectorNonNativeUpgradeTestRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestRoleAdminChangedIterator{contract: _ZetaConnectorNonNativeUpgradeTest.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeUpgradeTestRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeUpgradeTestRoleAdminChanged)
				if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) ParseRoleAdminChanged(log types.Log) (*ZetaConnectorNonNativeUpgradeTestRoleAdminChanged, error) {
	event := new(ZetaConnectorNonNativeUpgradeTestRoleAdminChanged)
	if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeUpgradeTestRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestRoleGrantedIterator struct {
	Event *ZetaConnectorNonNativeUpgradeTestRoleGranted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeUpgradeTestRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeUpgradeTestRoleGranted)
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
		it.Event = new(ZetaConnectorNonNativeUpgradeTestRoleGranted)
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
func (it *ZetaConnectorNonNativeUpgradeTestRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeUpgradeTestRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeUpgradeTestRoleGranted represents a RoleGranted event raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZetaConnectorNonNativeUpgradeTestRoleGrantedIterator, error) {

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

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestRoleGrantedIterator{contract: _ZetaConnectorNonNativeUpgradeTest.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeUpgradeTestRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeUpgradeTestRoleGranted)
				if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) ParseRoleGranted(log types.Log) (*ZetaConnectorNonNativeUpgradeTestRoleGranted, error) {
	event := new(ZetaConnectorNonNativeUpgradeTestRoleGranted)
	if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeUpgradeTestRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestRoleRevokedIterator struct {
	Event *ZetaConnectorNonNativeUpgradeTestRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeUpgradeTestRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeUpgradeTestRoleRevoked)
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
		it.Event = new(ZetaConnectorNonNativeUpgradeTestRoleRevoked)
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
func (it *ZetaConnectorNonNativeUpgradeTestRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeUpgradeTestRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeUpgradeTestRoleRevoked represents a RoleRevoked event raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZetaConnectorNonNativeUpgradeTestRoleRevokedIterator, error) {

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

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestRoleRevokedIterator{contract: _ZetaConnectorNonNativeUpgradeTest.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeUpgradeTestRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeUpgradeTestRoleRevoked)
				if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) ParseRoleRevoked(log types.Log) (*ZetaConnectorNonNativeUpgradeTestRoleRevoked, error) {
	event := new(ZetaConnectorNonNativeUpgradeTestRoleRevoked)
	if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeUpgradeTestUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestUnpausedIterator struct {
	Event *ZetaConnectorNonNativeUpgradeTestUnpaused // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeUpgradeTestUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeUpgradeTestUnpaused)
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
		it.Event = new(ZetaConnectorNonNativeUpgradeTestUnpaused)
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
func (it *ZetaConnectorNonNativeUpgradeTestUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeUpgradeTestUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeUpgradeTestUnpaused represents a Unpaused event raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ZetaConnectorNonNativeUpgradeTestUnpausedIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestUnpausedIterator{contract: _ZetaConnectorNonNativeUpgradeTest.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeUpgradeTestUnpaused) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeUpgradeTestUnpaused)
				if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) ParseUnpaused(log types.Log) (*ZetaConnectorNonNativeUpgradeTestUnpaused, error) {
	event := new(ZetaConnectorNonNativeUpgradeTestUnpaused)
	if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator is returned from FilterUpdatedZetaConnectorTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedZetaConnectorTSSAddress events raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator struct {
	Event *ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddress // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddress)
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
		it.Event = new(ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddress)
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
func (it *ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddress represents a UpdatedZetaConnectorTSSAddress event raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedZetaConnectorTSSAddress is a free log retrieval operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) FilterUpdatedZetaConnectorTSSAddress(opts *bind.FilterOpts) (*ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator, error) {

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.FilterLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddressIterator{contract: _ZetaConnectorNonNativeUpgradeTest.contract, event: "UpdatedZetaConnectorTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedZetaConnectorTSSAddress is a free log subscription operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedZetaConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) WatchUpdatedZetaConnectorTSSAddress(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddress) (event.Subscription, error) {

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.WatchLogs(opts, "UpdatedZetaConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddress)
				if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "UpdatedZetaConnectorTSSAddress", log); err != nil {
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
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) ParseUpdatedZetaConnectorTSSAddress(log types.Log) (*ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddress, error) {
	event := new(ZetaConnectorNonNativeUpgradeTestUpdatedZetaConnectorTSSAddress)
	if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "UpdatedZetaConnectorTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeUpgradeTestUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestUpgradedIterator struct {
	Event *ZetaConnectorNonNativeUpgradeTestUpgraded // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeUpgradeTestUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeUpgradeTestUpgraded)
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
		it.Event = new(ZetaConnectorNonNativeUpgradeTestUpgraded)
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
func (it *ZetaConnectorNonNativeUpgradeTestUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeUpgradeTestUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeUpgradeTestUpgraded represents a Upgraded event raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ZetaConnectorNonNativeUpgradeTestUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestUpgradedIterator{contract: _ZetaConnectorNonNativeUpgradeTest.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeUpgradeTestUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeUpgradeTestUpgraded)
				if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) ParseUpgraded(log types.Log) (*ZetaConnectorNonNativeUpgradeTestUpgraded, error) {
	event := new(ZetaConnectorNonNativeUpgradeTestUpgraded)
	if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeUpgradeTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestWithdrawnIterator struct {
	Event *ZetaConnectorNonNativeUpgradeTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeUpgradeTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeUpgradeTestWithdrawn)
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
		it.Event = new(ZetaConnectorNonNativeUpgradeTestWithdrawn)
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
func (it *ZetaConnectorNonNativeUpgradeTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeUpgradeTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeUpgradeTestWithdrawn represents a Withdrawn event raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNonNativeUpgradeTestWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.FilterLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestWithdrawnIterator{contract: _ZetaConnectorNonNativeUpgradeTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeUpgradeTestWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.WatchLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeUpgradeTestWithdrawn)
				if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) ParseWithdrawn(log types.Log) (*ZetaConnectorNonNativeUpgradeTestWithdrawn, error) {
	event := new(ZetaConnectorNonNativeUpgradeTestWithdrawn)
	if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalledIterator struct {
	Event *ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalled)
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
		it.Event = new(ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalled)
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
func (it *ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalled struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalledIterator{contract: _ZetaConnectorNonNativeUpgradeTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalled, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalled)
				if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalled, error) {
	event := new(ZetaConnectorNonNativeUpgradeTestWithdrawnAndCalled)
	if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeUpgradeTestWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestWithdrawnAndRevertedIterator struct {
	Event *ZetaConnectorNonNativeUpgradeTestWithdrawnAndReverted // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeUpgradeTestWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeUpgradeTestWithdrawnAndReverted)
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
		it.Event = new(ZetaConnectorNonNativeUpgradeTestWithdrawnAndReverted)
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
func (it *ZetaConnectorNonNativeUpgradeTestWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeUpgradeTestWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeUpgradeTestWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestWithdrawnAndReverted struct {
	To            common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNonNativeUpgradeTestWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestWithdrawnAndRevertedIterator{contract: _ZetaConnectorNonNativeUpgradeTest.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeUpgradeTestWithdrawnAndReverted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeUpgradeTestWithdrawnAndReverted)
				if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
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
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) ParseWithdrawnAndReverted(log types.Log) (*ZetaConnectorNonNativeUpgradeTestWithdrawnAndReverted, error) {
	event := new(ZetaConnectorNonNativeUpgradeTestWithdrawnAndReverted)
	if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaConnectorNonNativeUpgradeTestWithdrawnV2Iterator is returned from FilterWithdrawnV2 and is used to iterate over the raw logs and unpacked data for WithdrawnV2 events raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestWithdrawnV2Iterator struct {
	Event *ZetaConnectorNonNativeUpgradeTestWithdrawnV2 // Event containing the contract specifics and raw log

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
func (it *ZetaConnectorNonNativeUpgradeTestWithdrawnV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaConnectorNonNativeUpgradeTestWithdrawnV2)
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
		it.Event = new(ZetaConnectorNonNativeUpgradeTestWithdrawnV2)
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
func (it *ZetaConnectorNonNativeUpgradeTestWithdrawnV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaConnectorNonNativeUpgradeTestWithdrawnV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaConnectorNonNativeUpgradeTestWithdrawnV2 represents a WithdrawnV2 event raised by the ZetaConnectorNonNativeUpgradeTest contract.
type ZetaConnectorNonNativeUpgradeTestWithdrawnV2 struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnV2 is a free log retrieval operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) FilterWithdrawnV2(opts *bind.FilterOpts, to []common.Address) (*ZetaConnectorNonNativeUpgradeTestWithdrawnV2Iterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.FilterLogs(opts, "WithdrawnV2", toRule)
	if err != nil {
		return nil, err
	}
	return &ZetaConnectorNonNativeUpgradeTestWithdrawnV2Iterator{contract: _ZetaConnectorNonNativeUpgradeTest.contract, event: "WithdrawnV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnV2 is a free log subscription operation binding the contract event 0x3e35ef4326151011878c9e8e968a0f3913fe98ca68f72a1e0c2e9be13ffb3ee9.
//
// Solidity: event WithdrawnV2(address indexed to, uint256 amount)
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) WatchWithdrawnV2(opts *bind.WatchOpts, sink chan<- *ZetaConnectorNonNativeUpgradeTestWithdrawnV2, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZetaConnectorNonNativeUpgradeTest.contract.WatchLogs(opts, "WithdrawnV2", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaConnectorNonNativeUpgradeTestWithdrawnV2)
				if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
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
func (_ZetaConnectorNonNativeUpgradeTest *ZetaConnectorNonNativeUpgradeTestFilterer) ParseWithdrawnV2(log types.Log) (*ZetaConnectorNonNativeUpgradeTestWithdrawnV2, error) {
	event := new(ZetaConnectorNonNativeUpgradeTestWithdrawnV2)
	if err := _ZetaConnectorNonNativeUpgradeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
