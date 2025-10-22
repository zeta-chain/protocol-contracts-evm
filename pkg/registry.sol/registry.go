// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// ChainInfoDTO is an auto generated low-level Go binding around an user-defined struct.
type ChainInfoDTO struct {
	Active   bool
	ChainId  *big.Int
	GasZRC20 common.Address
	Registry []byte
}

// ContractInfoDTO is an auto generated low-level Go binding around an user-defined struct.
type ContractInfoDTO struct {
	Active       bool
	AddressBytes []byte
	ContractType string
	ChainId      *big.Int
}

// IRegistryChainMetadataEntry is an auto generated low-level Go binding around an user-defined struct.
type IRegistryChainMetadataEntry struct {
	ChainId *big.Int
	Key     string
	Value   []byte
}

// IRegistryContractConfigEntry is an auto generated low-level Go binding around an user-defined struct.
type IRegistryContractConfigEntry struct {
	ChainId      *big.Int
	ContractType string
	Key          string
	Value        []byte
}

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Sender common.Address
}

// ZRC20Info is an auto generated low-level Go binding around an user-defined struct.
type ZRC20Info struct {
	Active        bool
	Address       common.Address
	OriginAddress []byte
	OriginChainId *big.Int
	Symbol        string
	CoinType      string
	Decimals      uint8
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GATEWAY_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTRY_MANAGER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bootstrapChains\",\"inputs\":[{\"name\":\"chains\",\"type\":\"tuple[]\",\"internalType\":\"structChainInfoDTO[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasZRC20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registry\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"metadataEntries\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistry.ChainMetadataEntry[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bootstrapContracts\",\"inputs\":[{\"name\":\"contracts\",\"type\":\"tuple[]\",\"internalType\":\"structContractInfoDTO[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"configEntries\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistry.ContractConfigEntry[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bootstrapZRC20Tokens\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"tuple[]\",\"internalType\":\"structZRC20Info[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeChainStatus\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasZRC20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registry\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"activation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeRegistryManager\",\"inputs\":[{\"name\":\"newRegistryManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"coreRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gatewayEVM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllChains\",\"inputs\":[],\"outputs\":[{\"name\":\"chainsInfo\",\"type\":\"tuple[]\",\"internalType\":\"structChainInfoDTO[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasZRC20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registry\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"contractsInfo\",\"type\":\"tuple[]\",\"internalType\":\"structContractInfoDTO[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllZRC20Tokens\",\"inputs\":[],\"outputs\":[{\"name\":\"tokensInfo\",\"type\":\"tuple[]\",\"internalType\":\"structZRC20Info[]\",\"components\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainInfo\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"gasZRC20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registry\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainMetadata\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractConfiguration\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractInfo\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getZRC20AddressByForeignAsset\",\"inputs\":[{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getZRC20TokenInfo\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"registryManager_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gatewayEVM_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"coreRegistry_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerContract\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerZRC20Token\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"coinType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setContractActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setZRC20TokenActive\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateChainMetadata\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateContractConfiguration\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"oldAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainMetadataUpdated\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainStatusChanged\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newStatus\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractConfigurationUpdated\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractStatusChanged\",\"inputs\":[{\"name\":\"addressBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistryManagerChanged\",\"inputs\":[{\"name\":\"oldRegistryManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRegistryManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZRC20TokenRegistered\",\"inputs\":[{\"name\":\"originAddress\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"address_\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZRC20TokenUpdated\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ChainActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainNonActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ContractAlreadyRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ContractNotFound\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidContractType\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZRC20AlreadyRegistered\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZRC20SymbolAlreadyInUse\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614e136100fd60003960008181612bde01528181612c070152612dba0152614e136000f3fe6080604052600436106102e75760003560e01c80638f28397011610184578063bd8fde1c116100d6578063dc07f8661161008a578063f354b31f11610064578063f354b31f14610943578063f851a44014610963578063f8c8765e1461098357600080fd5b8063dc07f866146108bd578063e63ab1e9146108dd578063e9d6c5ba1461091157600080fd5b8063c50d6dd9116100bb578063c50d6dd914610849578063d3523ea21461087d578063d547741f1461089d57600080fd5b8063bd8fde1c146107f3578063c1bd469f1461082757600080fd5b80639ca220dd11610138578063a8f2cb9611610112578063a8f2cb961461076a578063aa808c061461078a578063ad3cb1cc146107aa57600080fd5b80639ca220dd14610713578063a217fddf14610735578063a4b7e5f91461074a57600080fd5b8063909b6a0311610169578063909b6a031461066c57806391d148541461068c57806394cc8683146106f157600080fd5b80638f2839701461062c5780639060bda91461064c57600080fd5b80634f1ef2861161023d578063676cc054116101f15780637066b18d116101cb5780637066b18d146105c9578063804ea334146105e95780638456cb591461061757600080fd5b8063676cc0541461055c5780636bf3d05a146105895780636e9e2d3f146105a957600080fd5b80635c975abb116102225780635c975abb146104d75780635cf92c9f1461050e578063631d62e41461053c57600080fd5b80634f1ef286146104af57806352d1902d146104c257600080fd5b8063248a9ca31161029f5780633500c24b116102795780633500c24b1461045a57806336568abe1461047a5780633f4ba83a1461049a57600080fd5b8063248a9ca3146103bd5780632c78f74c1461041a5780632f2ff15d1461043a57600080fd5b806310d29b9e116102d057806310d29b9e1461035957806318d3ce961461037b5780632259e9e51461039d57600080fd5b806301ffc9a7146102ec5780630c63109e14610321575b600080fd5b3480156102f857600080fd5b5061030c610307366004613ce7565b6109a3565b60405190151581526020015b60405180910390f35b34801561032d57600080fd5b50600154610341906001600160a01b031681565b6040516001600160a01b039091168152602001610318565b34801561036557600080fd5b50610379610374366004613d87565b610a3c565b005b34801561038757600080fd5b50610390610af1565b6040516103189190613e35565b3480156103a957600080fd5b506103796103b8366004613ef6565b610d93565b3480156103c957600080fd5b5061040c6103d8366004613f75565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b604051908152602001610318565b34801561042657600080fd5b50600b54610341906001600160a01b031681565b34801561044657600080fd5b50610379610455366004613fa5565b610e26565b34801561046657600080fd5b50610379610475366004613fd1565b610e70565b34801561048657600080fd5b50610379610495366004613fa5565b611003565b3480156104a657600080fd5b50610379611054565b6103796104bd36600461401b565b61106a565b3480156104ce57600080fd5b5061040c611089565b3480156104e357600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1661030c565b34801561051a57600080fd5b5061052e6105293660046140e6565b6110b8565b604051610318929190614132565b34801561054857600080fd5b50610379610557366004613ef6565b6111b3565b34801561056857600080fd5b5061057c610577366004614155565b61125a565b6040516103189190614197565b34801561059557600080fd5b50600c54610341906001600160a01b031681565b3480156105b557600080fd5b506103796105c43660046141bb565b611362565b3480156105d557600080fd5b5061057c6105e43660046140e6565b61141e565b3480156105f557600080fd5b50610609610604366004613f75565b6114e7565b60405161031892919061428b565b34801561062357600080fd5b5061037961159f565b34801561063857600080fd5b50610379610647366004613fd1565b6115d1565b34801561065857600080fd5b506103796106673660046142ad565b611725565b34801561067857600080fd5b5061037961068736600461431c565b6117b7565b34801561069857600080fd5b5061030c6106a7366004613fa5565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156106fd57600080fd5b50610706611878565b604051610318919061435e565b34801561071f57600080fd5b506107286118d0565b60405161031891906143a1565b34801561074157600080fd5b5061040c600081565b34801561075657600080fd5b5061037961076536600461444e565b611a8f565b34801561077657600080fd5b506103796107853660046144bf565b611b98565b34801561079657600080fd5b506103416107a53660046140e6565b611c1a565b3480156107b657600080fd5b5061057c6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b3480156107ff57600080fd5b5061040c7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa381565b34801561083357600080fd5b5061083c611c5a565b604051610318919061452e565b34801561085557600080fd5b5061040c7fb90e9995c6170fff8ea03e9ad6919878e483770c237f1a6f330ceaa7112b344a81565b34801561088957600080fd5b5061057c610898366004613ef6565b611f61565b3480156108a957600080fd5b506103796108b8366004613fa5565b612049565b3480156108c957600080fd5b506103796108d836600461444e565b61208d565b3480156108e957600080fd5b5061040c7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b34801561091d57600080fd5b5061093161092c366004613fd1565b612191565b60405161031896959493929190614625565b34801561094f57600080fd5b5061037961095e366004614683565b6123e3565b34801561096f57600080fd5b50600054610341906001600160a01b031681565b34801561098f57600080fd5b5061037961099e366004614733565b61247e565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610a3657507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b333014610a75576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a81848484846127a2565b7f6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c600760008681526020019081526020016000208484604051610ac592919061477c565b9081526020016040518091039020600101604051610ae391906147df565b60405180910390a150505050565b6004546060908067ffffffffffffffff811115610b1057610b10613fec565b604051908082528060200260200182016040528015610b6e57816020015b610b5b60405180608001604052806000151581526020016060815260200160608152602001600081525090565b815260200190600190039081610b2e5790505b50915060005b81811015610d8e57600060048281548110610b9157610b9161486b565b906000526020600020906002020160405180604001604052908160008201548152602001600182018054610bc49061478c565b80601f0160208091040260200160405190810160405280929190818152602001828054610bf09061478c565b8015610c3d5780601f10610c1257610100808354040283529160200191610c3d565b820191906000526020600020905b815481529060010190602001808311610c2057829003601f168201915b505050505081525050905060008160000151905060008260200151905060405180608001604052806007600085815260200190815260200160002083604051610c86919061489a565b90815260408051602092819003830190205460ff161515835260008681526007835281902090519290910191610cbd90859061489a565b90815260200160405180910390206001018054610cd99061478c565b80601f0160208091040260200160405190810160405280929190818152602001828054610d059061478c565b8015610d525780601f10610d2757610100808354040283529160200191610d52565b820191906000526020600020905b815481529060010190602001808311610d3557829003601f168201915b5050505050815260200182815260200183815250868581518110610d7857610d7861486b565b6020908102919091010152505050600101610b74565b505090565b333014610dcc576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610dd46128f6565b610de18585858585612954565b847f40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee63485858585604051610e1794939291906148e1565b60405180910390a25050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610e60816129e2565b610e6a83836129ec565b50505050565b6000610e7b816129e2565b6001600160a01b038216610ebb576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ee57ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3836129ec565b50610f107f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a836129ec565b50600154610f48907ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3906001600160a01b0316612abb565b50600154610f80907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a906001600160a01b0316612abb565b50600154604080516001600160a01b03928316815291841660208301527f6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279910160405180910390a150600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6001600160a01b0381163314611045576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61104f8282612abb565b505050565b600061105f816129e2565b611067612b61565b50565b611072612bd3565b61107b82612ca3565b6110858282612cae565b5050565b6000611093612daf565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b6000838152600760205260408082209051606091906110da908690869061477c565b908152604080519182900360209081018320546000898152600790925291902060ff90911693509061110f908690869061477c565b9081526020016040518091039020600101805461112b9061478c565b80601f01602080910402602001604051908101604052809291908181526020018280546111579061478c565b80156111a45780601f10611179576101008083540402835291602001916111a4565b820191906000526020600020905b81548152906001019060200180831161118757829003601f168201915b50505050509050935093915050565b3330146111ec576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111f46128f6565b6112018585858585612e11565b838360405161121192919061477c565b6040518091039020857f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581848460405161124b929190614913565b60405180910390a35050505050565b60607fb90e9995c6170fff8ea03e9ad6919878e483770c237f1a6f330ceaa7112b344a611286816129e2565b61128e6128f6565b600c546001600160a01b03166112a76020870187613fd1565b6001600160a01b0316146112e7576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080306001600160a01b0316868660405161130492919061477c565b6000604051808303816000865af19150503d8060008114611341576040519150601f19603f3d011682016040523d82523d6000602084013e611346565b606091505b50915091508161135857805160208201fd5b9695505050505050565b33301461139b576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113a36128f6565b6113b48989898989898989896130ac565b886001600160a01b031685856040516113ce92919061477c565b60405180910390207fa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e36783898c8c60405161140b9493929190614927565b60405180910390a3505050505050505050565b606060066000858152602001908152602001600020600401838360405161144692919061477c565b9081526020016040518091039020805461145f9061478c565b80601f016020809104026020016040519081016040528092919081815260200182805461148b9061478c565b80156114d85780601f106114ad576101008083540402835291602001916114d8565b820191906000526020600020905b8154815290600101906020018083116114bb57829003601f168201915b505050505090505b9392505050565b60008181526006602052604090206002810154600390910180546001600160a01b03909216916060919061151a9061478c565b80601f01602080910402602001604051908101604052809291908181526020018280546115469061478c565b80156115935780601f1061156857610100808354040283529160200191611593565b820191906000526020600020905b81548152906001019060200180831161157657829003601f168201915b50505050509050915091565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6115c9816129e2565b6110676133e2565b60006115dc816129e2565b6001600160a01b03821661161c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6116276000836129ec565b506116527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a836129ec565b506000805461166a91906001600160a01b0316612abb565b506000546116a2907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a906001600160a01b0316612abb565b50600054604080516001600160a01b03928316815291841660208301527f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f910160405180910390a150600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b33301461175e576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6117666128f6565b611770828261343d565b604080516001600160a01b038416815282151560208201527f9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8910160405180910390a15050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa36117e1816129e2565b6117e96128f6565b60005b82811015610e6a57368484838181106118075761180761486b565b9050602002810190611819919061494a565b905061186f61182e6040830160208401613fd1565b61183b608084018461497e565b606085013561184d604087018761497e565b61185a60a089018961497e565b61186a60e08b0160c08c016149e3565b6130ac565b506001016117ec565b606060028054806020026020016040519081016040528092919081815260200182805480156118c657602002820191906000526020600020905b8154815260200190600101908083116118b2575b5050505050905090565b6003546060908067ffffffffffffffff8111156118ef576118ef613fec565b60405190808252806020026020018201604052801561195e57816020015b60408051608081018252600080825260208083018290529282015260608082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90920191018161190d5790505b50915060005b81811015610d8e576000600382815481106119815761198161486b565b6000918252602080832090910154604080516080810182528285526006808552828620805460ff161515835282860185905260028101546001600160a01b03169383019390935294839052939092526003909101805491935060608301916119e89061478c565b80601f0160208091040260200160405190810160405280929190818152602001828054611a149061478c565b8015611a615780601f10611a3657610100808354040283529160200191611a61565b820191906000526020600020905b815481529060010190602001808311611a4457829003601f168201915b5050505050815250848381518110611a7b57611a7b61486b565b602090810291909101015250600101611964565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3611ab9816129e2565b611ac16128f6565b60005b84811015611b235736868683818110611adf57611adf61486b565b9050602002810190611af191906149fe565b9050611b1a6060820135611b08604084018461497e565b611b15602086018661497e565b612e11565b50600101611ac4565b5060005b82811015611b905736848483818110611b4257611b4261486b565b9050602002810190611b5491906149fe565b9050611b878135611b68602084018461497e565b611b75604086018661497e565b611b82606088018861497e565b61352d565b50600101611b27565b505050505050565b333014611bd1576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611bd96128f6565b611be6858585858561368f565b847fc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e82604051610e17911515815260200190565b6000838152600a60205260408082209051611c38908590859061477c565b908152604051908190036020019020546001600160a01b031690509392505050565b6005546060908067ffffffffffffffff811115611c7957611c79613fec565b604051908082528060200260200182016040528015611cf857816020015b611ce56040518060e0016040528060001515815260200160006001600160a01b0316815260200160608152602001600081526020016060815260200160608152602001600060ff1681525090565b815260200190600190039081611c975790505b50915060005b81811015610d8e57600060058281548110611d1b57611d1b61486b565b60009182526020808320909101546001600160a01b0390811680845260088352604093849020845160e081018652815460ff811615158252610100900490931693830193909352600183018054919550919384019190611d7a9061478c565b80601f0160208091040260200160405190810160405280929190818152602001828054611da69061478c565b8015611df35780601f10611dc857610100808354040283529160200191611df3565b820191906000526020600020905b815481529060010190602001808311611dd657829003601f168201915b5050505050815260200160028201548152602001600382018054611e169061478c565b80601f0160208091040260200160405190810160405280929190818152602001828054611e429061478c565b8015611e8f5780601f10611e6457610100808354040283529160200191611e8f565b820191906000526020600020905b815481529060010190602001808311611e7257829003601f168201915b50505050508152602001600482018054611ea89061478c565b80601f0160208091040260200160405190810160405280929190818152602001828054611ed49061478c565b8015611f215780601f10611ef657610100808354040283529160200191611f21565b820191906000526020600020905b815481529060010190602001808311611f0457829003601f168201915b50505091835250506005919091015460ff166020909101528451859084908110611f4d57611f4d61486b565b602090810291909101015250600101611cfe565b6060600760008781526020019081526020016000208585604051611f8692919061477c565b90815260200160405180910390206003018383604051611fa792919061477c565b90815260200160405180910390208054611fc09061478c565b80601f0160208091040260200160405190810160405280929190818152602001828054611fec9061478c565b80156120395780601f1061200e57610100808354040283529160200191612039565b820191906000526020600020905b81548152906001019060200180831161201c57829003601f168201915b5050505050905095945050505050565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154612083816129e2565b610e6a8383612abb565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa36120b7816129e2565b6120bf6128f6565b60005b8481101561213157368686838181106120dd576120dd61486b565b90506020028101906120ef91906149fe565b905061212860208201356121096060840160408501613fd1565b612116606085018561497e565b6121236020870187614a32565b61368f565b506001016120c2565b5060005b82811015611b9057368484838181106121505761215061486b565b90506020028101906121629190614a4d565b90506121888135612176602084018461497e565b612183604086018661497e565b612954565b50600101612135565b6001600160a01b038082166000908152600860209081526040808320815160e081018352815460ff811615158252610100900490951692850192909252600182018054939460609486948694859487948594909392840191906121f39061478c565b80601f016020809104026020016040519081016040528092919081815260200182805461221f9061478c565b801561226c5780601f106122415761010080835404028352916020019161226c565b820191906000526020600020905b81548152906001019060200180831161224f57829003601f168201915b505050505081526020016002820154815260200160038201805461228f9061478c565b80601f01602080910402602001604051908101604052809291908181526020018280546122bb9061478c565b80156123085780601f106122dd57610100808354040283529160200191612308565b820191906000526020600020905b8154815290600101906020018083116122eb57829003601f168201915b505050505081526020016004820180546123219061478c565b80601f016020809104026020016040519081016040528092919081815260200182805461234d9061478c565b801561239a5780601f1061236f5761010080835404028352916020019161239a565b820191906000526020600020905b81548152906001019060200180831161237d57829003601f168201915b50505091835250506005919091015460ff16602090910152805160808201516060830151604084015160a085015160c090950151939d929c50909a509850919650945092505050565b33301461241c576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6124246128f6565b6124338787878787878761352d565b867faea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c587878787878760405161246d96959493929190614a81565b60405180910390a250505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff166000811580156124c95750825b905060008267ffffffffffffffff1660011480156124e65750303b155b9050811580156124f4575080155b1561252b576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000166001178555831561258c5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b03891615806125a957506001600160a01b038716155b806125bb57506001600160a01b038616155b806125cd57506001600160a01b038816155b15612604576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61260c61384c565b61261461384c565b61261c613854565b61262760008a6129ec565b506126527ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3896129ec565b5061267d7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a896129ec565b506126a87f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8a6129ec565b506126d37fb90e9995c6170fff8ea03e9ad6919878e483770c237f1a6f330ceaa7112b344a886129ec565b50600080546001600160a01b03808c167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255600180548b8416908316179055600b80548a8416908316179055600c80549289169290911691909117905583156127975784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b60008481526006602052604090205460ff166127f2576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018590526024015b60405180910390fd5b60008290036128305782826040517ec10cfd0000000000000000000000000000000000000000000000000000000081526004016127e9929190614913565b60008481526007602052604090819020905161284f908590859061477c565b9081526020016040518091039020600101805461286b9061478c565b90506000036128ac578383836040517f2b4f9c860000000000000000000000000000000000000000000000000000000081526004016127e993929190614aca565b806007600086815260200190815260200160002084846040516128d092919061477c565b908152604051908190036020019020805491151560ff1990921691909117905550505050565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615612952576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60008581526006602052604090205460ff1661299f576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018690526024016127e9565b81816006600088815260200190815260200160002060040186866040516129c792919061477c565b90815260200160405180910390209182611b90929190614b2b565b6110678133613887565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612ab1576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055612a673390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610a36565b6000915050610a36565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612ab1576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610a36565b612b69613914565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480612c6c57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612c607f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15612952576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611085816129e2565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612d08575060408051601f3d908101601f19168201909252612d0591810190614c27565b60015b612d49576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016127e9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612da5576040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600481018290526024016127e9565b61104f838361396f565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612952576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008581526006602052604090205460ff16612e5c576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018690526024016127e9565b6000839003612e9a5783836040517ec10cfd0000000000000000000000000000000000000000000000000000000081526004016127e9929190614913565b6000819003612ed5576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000858152600760205260408082209051612ef3908790879061477c565b90815260200160405180910390206001018054612f0f9061478c565b90501115612f535784848484846040517f2b192eab0000000000000000000000000000000000000000000000000000000081526004016127e9959493929190614c40565b6001600760008781526020019081526020016000208585604051612f7892919061477c565b9081526040805160209281900383018120805460ff1916941515949094179093556000888152600790925290208391839190612fb7908890889061477c565b90815260200160405180910390206001019182612fd5929190614b2b565b508383600760008881526020019081526020016000208686604051612ffb92919061477c565b90815260200160405180910390206002019182613019929190614b2b565b506004604051806040016040528087815260200186868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509390945250508354600181810186559482526020918290208451600290920201908155908301519293909290830191506130999082614c79565b505050838360405161121192919061477c565b6001600160a01b0389166130ec576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000879003613156576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f53796d626f6c2063616e6e6f7420626520656d7074790000000000000000000060448201526064016127e9565b6001600160a01b03898116600090815260086020526040902054610100900416156131b8576040517f63f4ee1f0000000000000000000000000000000000000000000000000000000081526001600160a01b038a1660048201526024016127e9565b60006001600160a01b0316600989896040516131d592919061477c565b908152604051908190036020019020546001600160a01b0316146132295787876040517fb295cac90000000000000000000000000000000000000000000000000000000081526004016127e9929190614913565b6001600160a01b0389166000818152600860205260409020805460017fffffffffffffffffffffff000000000000000000000000000000000000000000909116610100909302929092178217815501613283858783614b2b565b506001600160a01b0389166000908152600860205260409020600281018790556003016132b1888a83614b2b565b506001600160a01b038916600090815260086020526040902060058101805460ff191660ff84161790556004016132e9838583614b2b565b5088600a6000888152602001908152602001600020868660405161330e92919061477c565b908152602001604051809103902060006101000a8154816001600160a01b0302191690836001600160a01b03160217905550886009898960405161335392919061477c565b90815260405190819003602001902080546001600160a01b039283167fffffffffffffffffffffffff000000000000000000000000000000000000000091821617909155600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180549b9092169a16999099179098555050505050505050565b6133ea6128f6565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612bb5565b6001600160a01b03821661347d576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260086020526040902054610100900416613502576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5a52433230206e6f74207265676973746572656400000000000000000000000060448201526064016127e9565b6001600160a01b03919091166000908152600860205260409020805460ff1916911515919091179055565b60008781526006602052604090205460ff16613578576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018890526024016127e9565b60008590036135b65785856040517ec10cfd0000000000000000000000000000000000000000000000000000000081526004016127e9929190614913565b6000878152600760205260409081902090516135d5908890889061477c565b9081526040519081900360200190205460ff16613624578686866040517f2b4f9c860000000000000000000000000000000000000000000000000000000081526004016127e993929190614aca565b8181600760008a8152602001908152602001600020888860405161364992919061477c565b9081526020016040518091039020600301868660405161366a92919061477c565b90815260200160405180910390209182613685929190614b2b565b5050505050505050565b60008581526006602052604090205460ff1680156136aa5750805b156136e4576040517fa1452cb0000000000000000000000000000000000000000000000000000000008152600481018690526024016127e9565b60008581526006602052604090205460ff16158015613701575080155b1561373b576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018690526024016127e9565b6000858152600660205260409020600201546001600160a01b03161580156137635750468514155b1561379e57600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018590555b6000858152600660205260409020805460ff19168215151781556002810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0387161790556003016137fb838583614b2b565b50801561383c57600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01859055613845565b613845856139c5565b5050505050565b612952613a73565b61385c613a73565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff19169055565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16611085576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602481018390526044016127e9565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16612952576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61397882613ada565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156139bd5761104f8282613b82565b611085613bf8565b60005b6002548110156110855781600282815481106139e6576139e661486b565b906000526020600020015403613a6b5760028054613a0690600190614d74565b81548110613a1657613a1661486b565b906000526020600020015460028281548110613a3457613a3461486b565b6000918252602090912001556002805480613a5157613a51614dae565b600190038181906000526020600020016000905590555050565b6001016139c8565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16612952576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806001600160a01b03163b600003613b29576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016127e9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051613b9f919061489a565b600060405180830381855af49150503d8060008114613bda576040519150601f19603f3d011682016040523d82523d6000602084013e613bdf565b606091505b5091509150613bef858383613c30565b95945050505050565b3415612952576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606082613c4557613c4082613ca5565b6114e0565b8151158015613c5c57506001600160a01b0384163b155b15613c9e576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016127e9565b50806114e0565b805115613cb55780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060208284031215613cf957600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146114e057600080fd5b60008083601f840112613d3b57600080fd5b50813567ffffffffffffffff811115613d5357600080fd5b602083019150836020828501011115613d6b57600080fd5b9250929050565b80358015158114613d8257600080fd5b919050565b60008060008060608587031215613d9d57600080fd5b84359350602085013567ffffffffffffffff811115613dbb57600080fd5b613dc787828801613d29565b9094509250613dda905060408601613d72565b905092959194509250565b60005b83811015613e00578181015183820152602001613de8565b50506000910152565b60008151808452613e21816020860160208601613de5565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613eea577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805115158652602081015160806020880152613ead6080880182613e09565b905060408201518782036040890152613ec68282613e09565b60609384015198909301979097525094506020938401939190910190600101613e5d565b50929695505050505050565b600080600080600060608688031215613f0e57600080fd5b85359450602086013567ffffffffffffffff811115613f2c57600080fd5b613f3888828901613d29565b909550935050604086013567ffffffffffffffff811115613f5857600080fd5b613f6488828901613d29565b969995985093965092949392505050565b600060208284031215613f8757600080fd5b5035919050565b80356001600160a01b0381168114613d8257600080fd5b60008060408385031215613fb857600080fd5b82359150613fc860208401613f8e565b90509250929050565b600060208284031215613fe357600080fd5b6114e082613f8e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806040838503121561402e57600080fd5b61403783613f8e565b9150602083013567ffffffffffffffff81111561405357600080fd5b8301601f8101851361406457600080fd5b803567ffffffffffffffff81111561407e5761407e613fec565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff821117156140ae576140ae613fec565b6040528181528282016020018710156140c657600080fd5b816020840160208301376000602083830101528093505050509250929050565b6000806000604084860312156140fb57600080fd5b83359250602084013567ffffffffffffffff81111561411957600080fd5b61412586828701613d29565b9497909650939450505050565b821515815260406020820152600061414d6040830184613e09565b949350505050565b6000806000838503604081121561416b57600080fd5b602081121561417957600080fd5b50839250602084013567ffffffffffffffff81111561411957600080fd5b6020815260006114e06020830184613e09565b803560ff81168114613d8257600080fd5b600080600080600080600080600060c08a8c0312156141d957600080fd5b6141e28a613f8e565b985060208a013567ffffffffffffffff8111156141fe57600080fd5b61420a8c828d01613d29565b90995097505060408a0135955060608a013567ffffffffffffffff81111561423157600080fd5b61423d8c828d01613d29565b90965094505060808a013567ffffffffffffffff81111561425d57600080fd5b6142698c828d01613d29565b909450925061427c905060a08b016141aa565b90509295985092959850929598565b6001600160a01b038316815260406020820152600061414d6040830184613e09565b600080604083850312156142c057600080fd5b6142c983613f8e565b9150613fc860208401613d72565b60008083601f8401126142e957600080fd5b50813567ffffffffffffffff81111561430157600080fd5b6020830191508360208260051b8501011115613d6b57600080fd5b6000806020838503121561432f57600080fd5b823567ffffffffffffffff81111561434657600080fd5b614352858286016142d7565b90969095509350505050565b602080825282518282018190526000918401906040840190835b81811015614396578351835260209384019390920191600101614378565b509095945050505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613eea577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805115158652602081015160208701526001600160a01b03604082015116604087015260608101519050608060608701526144386080870182613e09565b95505060209384019391909101906001016143c9565b6000806000806040858703121561446457600080fd5b843567ffffffffffffffff81111561447b57600080fd5b614487878288016142d7565b909550935050602085013567ffffffffffffffff8111156144a757600080fd5b6144b3878288016142d7565b95989497509550505050565b6000806000806000608086880312156144d757600080fd5b853594506144e760208701613f8e565b9350604086013567ffffffffffffffff81111561450357600080fd5b61450f88828901613d29565b9094509250614522905060608701613d72565b90509295509295909350565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613eea577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051151586526001600160a01b036020820151166020870152604081015160e060408801526145b960e0880182613e09565b905060608201516060880152608082015187820360808901526145dc8282613e09565b91505060a082015187820360a08901526145f68282613e09565b91505060c0820151915061460f60c088018360ff169052565b9550506020938401939190910190600101614556565b861515815260c06020820152600061464060c0830188613e09565b86604084015282810360608401526146588187613e09565b9050828103608084015261466c8186613e09565b91505060ff831660a0830152979650505050505050565b60008060008060008060006080888a03121561469e57600080fd5b87359650602088013567ffffffffffffffff8111156146bc57600080fd5b6146c88a828b01613d29565b909750955050604088013567ffffffffffffffff8111156146e857600080fd5b6146f48a828b01613d29565b909550935050606088013567ffffffffffffffff81111561471457600080fd5b6147208a828b01613d29565b989b979a50959850939692959293505050565b6000806000806080858703121561474957600080fd5b61475285613f8e565b935061476060208601613f8e565b925061476e60408601613f8e565b9150613dda60608601613f8e565b8183823760009101908152919050565b600181811c908216806147a057607f821691505b6020821081036147d9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b6020815260008083546147f18161478c565b8060208601526001821660008114614810576001811461482c57614860565b60ff1983166040870152604082151560051b8701019350614860565b86600052602060002060005b8381101561485757815488820160400152600190910190602001614838565b87016040019450505b509195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082516148ac818460208701613de5565b9190910192915050565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b6040815260006148f56040830186886148b6565b82810360208401526149088185876148b6565b979650505050505050565b60208152600061414d6020830184866148b6565b60ff851681528360208201526060604082015260006113586060830184866148b6565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff218336030181126148ac57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126149b357600080fd5b83018035915067ffffffffffffffff8211156149ce57600080fd5b602001915036819003821315613d6b57600080fd5b6000602082840312156149f557600080fd5b6114e0826141aa565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818336030181126148ac57600080fd5b600060208284031215614a4457600080fd5b6114e082613d72565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa18336030181126148ac57600080fd5b606081526000614a9560608301888a6148b6565b8281036020840152614aa88187896148b6565b90508281036040840152614abd8185876148b6565b9998505050505050505050565b838152604060208201526000613bef6040830184866148b6565b601f82111561104f57806000526020600020601f840160051c81016020851015614b0b5750805b601f840160051c820191505b818110156138455760008155600101614b17565b67ffffffffffffffff831115614b4357614b43613fec565b614b5783614b51835461478c565b83614ae4565b6000601f841160018114614ba95760008515614b735750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355613845565b600083815260209020601f19861690835b82811015614bda5786850135825560209485019460019092019101614bba565b5086821015614c15577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b600060208284031215614c3957600080fd5b5051919050565b858152606060208201526000614c5a6060830186886148b6565b8281036040840152614c6d8185876148b6565b98975050505050505050565b815167ffffffffffffffff811115614c9357614c93613fec565b614ca781614ca1845461478c565b84614ae4565b6020601f821160018114614cf95760008315614cc35750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455613845565b600084815260208120601f198516915b82811015614d295787850151825560209485019460019092019101614d09565b5084821015614d6557868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b81810381811115610a36577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea26469706673582212204c02e3015e563cbe3078d5d6d927b9232259cff40b51d903457d773f6cf268ab64736f6c634300081a0033",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// RegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegistryMetaData.Bin instead.
var RegistryBin = RegistryMetaData.Bin

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registry *RegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registry *RegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Registry.Contract.DEFAULTADMINROLE(&_Registry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registry *RegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Registry.Contract.DEFAULTADMINROLE(&_Registry.CallOpts)
}

// GATEWAYROLE is a free data retrieval call binding the contract method 0xc50d6dd9.
//
// Solidity: function GATEWAY_ROLE() view returns(bytes32)
func (_Registry *RegistryCaller) GATEWAYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "GATEWAY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GATEWAYROLE is a free data retrieval call binding the contract method 0xc50d6dd9.
//
// Solidity: function GATEWAY_ROLE() view returns(bytes32)
func (_Registry *RegistrySession) GATEWAYROLE() ([32]byte, error) {
	return _Registry.Contract.GATEWAYROLE(&_Registry.CallOpts)
}

// GATEWAYROLE is a free data retrieval call binding the contract method 0xc50d6dd9.
//
// Solidity: function GATEWAY_ROLE() view returns(bytes32)
func (_Registry *RegistryCallerSession) GATEWAYROLE() ([32]byte, error) {
	return _Registry.Contract.GATEWAYROLE(&_Registry.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Registry *RegistryCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Registry *RegistrySession) PAUSERROLE() ([32]byte, error) {
	return _Registry.Contract.PAUSERROLE(&_Registry.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Registry *RegistryCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Registry.Contract.PAUSERROLE(&_Registry.CallOpts)
}

// REGISTRYMANAGERROLE is a free data retrieval call binding the contract method 0xbd8fde1c.
//
// Solidity: function REGISTRY_MANAGER_ROLE() view returns(bytes32)
func (_Registry *RegistryCaller) REGISTRYMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "REGISTRY_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTRYMANAGERROLE is a free data retrieval call binding the contract method 0xbd8fde1c.
//
// Solidity: function REGISTRY_MANAGER_ROLE() view returns(bytes32)
func (_Registry *RegistrySession) REGISTRYMANAGERROLE() ([32]byte, error) {
	return _Registry.Contract.REGISTRYMANAGERROLE(&_Registry.CallOpts)
}

// REGISTRYMANAGERROLE is a free data retrieval call binding the contract method 0xbd8fde1c.
//
// Solidity: function REGISTRY_MANAGER_ROLE() view returns(bytes32)
func (_Registry *RegistryCallerSession) REGISTRYMANAGERROLE() ([32]byte, error) {
	return _Registry.Contract.REGISTRYMANAGERROLE(&_Registry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Registry *RegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Registry *RegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Registry.Contract.UPGRADEINTERFACEVERSION(&_Registry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Registry *RegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Registry.Contract.UPGRADEINTERFACEVERSION(&_Registry.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Registry *RegistryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Registry *RegistrySession) Admin() (common.Address, error) {
	return _Registry.Contract.Admin(&_Registry.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Registry *RegistryCallerSession) Admin() (common.Address, error) {
	return _Registry.Contract.Admin(&_Registry.CallOpts)
}

// CoreRegistry is a free data retrieval call binding the contract method 0x6bf3d05a.
//
// Solidity: function coreRegistry() view returns(address)
func (_Registry *RegistryCaller) CoreRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "coreRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CoreRegistry is a free data retrieval call binding the contract method 0x6bf3d05a.
//
// Solidity: function coreRegistry() view returns(address)
func (_Registry *RegistrySession) CoreRegistry() (common.Address, error) {
	return _Registry.Contract.CoreRegistry(&_Registry.CallOpts)
}

// CoreRegistry is a free data retrieval call binding the contract method 0x6bf3d05a.
//
// Solidity: function coreRegistry() view returns(address)
func (_Registry *RegistryCallerSession) CoreRegistry() (common.Address, error) {
	return _Registry.Contract.CoreRegistry(&_Registry.CallOpts)
}

// GatewayEVM is a free data retrieval call binding the contract method 0x2c78f74c.
//
// Solidity: function gatewayEVM() view returns(address)
func (_Registry *RegistryCaller) GatewayEVM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "gatewayEVM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatewayEVM is a free data retrieval call binding the contract method 0x2c78f74c.
//
// Solidity: function gatewayEVM() view returns(address)
func (_Registry *RegistrySession) GatewayEVM() (common.Address, error) {
	return _Registry.Contract.GatewayEVM(&_Registry.CallOpts)
}

// GatewayEVM is a free data retrieval call binding the contract method 0x2c78f74c.
//
// Solidity: function gatewayEVM() view returns(address)
func (_Registry *RegistryCallerSession) GatewayEVM() (common.Address, error) {
	return _Registry.Contract.GatewayEVM(&_Registry.CallOpts)
}

// GetActiveChains is a free data retrieval call binding the contract method 0x94cc8683.
//
// Solidity: function getActiveChains() view returns(uint256[])
func (_Registry *RegistryCaller) GetActiveChains(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getActiveChains")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActiveChains is a free data retrieval call binding the contract method 0x94cc8683.
//
// Solidity: function getActiveChains() view returns(uint256[])
func (_Registry *RegistrySession) GetActiveChains() ([]*big.Int, error) {
	return _Registry.Contract.GetActiveChains(&_Registry.CallOpts)
}

// GetActiveChains is a free data retrieval call binding the contract method 0x94cc8683.
//
// Solidity: function getActiveChains() view returns(uint256[])
func (_Registry *RegistryCallerSession) GetActiveChains() ([]*big.Int, error) {
	return _Registry.Contract.GetActiveChains(&_Registry.CallOpts)
}

// GetAllChains is a free data retrieval call binding the contract method 0x9ca220dd.
//
// Solidity: function getAllChains() view returns((bool,uint256,address,bytes)[] chainsInfo)
func (_Registry *RegistryCaller) GetAllChains(opts *bind.CallOpts) ([]ChainInfoDTO, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getAllChains")

	if err != nil {
		return *new([]ChainInfoDTO), err
	}

	out0 := *abi.ConvertType(out[0], new([]ChainInfoDTO)).(*[]ChainInfoDTO)

	return out0, err

}

// GetAllChains is a free data retrieval call binding the contract method 0x9ca220dd.
//
// Solidity: function getAllChains() view returns((bool,uint256,address,bytes)[] chainsInfo)
func (_Registry *RegistrySession) GetAllChains() ([]ChainInfoDTO, error) {
	return _Registry.Contract.GetAllChains(&_Registry.CallOpts)
}

// GetAllChains is a free data retrieval call binding the contract method 0x9ca220dd.
//
// Solidity: function getAllChains() view returns((bool,uint256,address,bytes)[] chainsInfo)
func (_Registry *RegistryCallerSession) GetAllChains() ([]ChainInfoDTO, error) {
	return _Registry.Contract.GetAllChains(&_Registry.CallOpts)
}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns((bool,bytes,string,uint256)[] contractsInfo)
func (_Registry *RegistryCaller) GetAllContracts(opts *bind.CallOpts) ([]ContractInfoDTO, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getAllContracts")

	if err != nil {
		return *new([]ContractInfoDTO), err
	}

	out0 := *abi.ConvertType(out[0], new([]ContractInfoDTO)).(*[]ContractInfoDTO)

	return out0, err

}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns((bool,bytes,string,uint256)[] contractsInfo)
func (_Registry *RegistrySession) GetAllContracts() ([]ContractInfoDTO, error) {
	return _Registry.Contract.GetAllContracts(&_Registry.CallOpts)
}

// GetAllContracts is a free data retrieval call binding the contract method 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns((bool,bytes,string,uint256)[] contractsInfo)
func (_Registry *RegistryCallerSession) GetAllContracts() ([]ContractInfoDTO, error) {
	return _Registry.Contract.GetAllContracts(&_Registry.CallOpts)
}

// GetAllZRC20Tokens is a free data retrieval call binding the contract method 0xc1bd469f.
//
// Solidity: function getAllZRC20Tokens() view returns((bool,address,bytes,uint256,string,string,uint8)[] tokensInfo)
func (_Registry *RegistryCaller) GetAllZRC20Tokens(opts *bind.CallOpts) ([]ZRC20Info, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getAllZRC20Tokens")

	if err != nil {
		return *new([]ZRC20Info), err
	}

	out0 := *abi.ConvertType(out[0], new([]ZRC20Info)).(*[]ZRC20Info)

	return out0, err

}

// GetAllZRC20Tokens is a free data retrieval call binding the contract method 0xc1bd469f.
//
// Solidity: function getAllZRC20Tokens() view returns((bool,address,bytes,uint256,string,string,uint8)[] tokensInfo)
func (_Registry *RegistrySession) GetAllZRC20Tokens() ([]ZRC20Info, error) {
	return _Registry.Contract.GetAllZRC20Tokens(&_Registry.CallOpts)
}

// GetAllZRC20Tokens is a free data retrieval call binding the contract method 0xc1bd469f.
//
// Solidity: function getAllZRC20Tokens() view returns((bool,address,bytes,uint256,string,string,uint8)[] tokensInfo)
func (_Registry *RegistryCallerSession) GetAllZRC20Tokens() ([]ZRC20Info, error) {
	return _Registry.Contract.GetAllZRC20Tokens(&_Registry.CallOpts)
}

// GetChainInfo is a free data retrieval call binding the contract method 0x804ea334.
//
// Solidity: function getChainInfo(uint256 chainId) view returns(address gasZRC20, bytes registry)
func (_Registry *RegistryCaller) GetChainInfo(opts *bind.CallOpts, chainId *big.Int) (struct {
	GasZRC20 common.Address
	Registry []byte
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getChainInfo", chainId)

	outstruct := new(struct {
		GasZRC20 common.Address
		Registry []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GasZRC20 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Registry = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetChainInfo is a free data retrieval call binding the contract method 0x804ea334.
//
// Solidity: function getChainInfo(uint256 chainId) view returns(address gasZRC20, bytes registry)
func (_Registry *RegistrySession) GetChainInfo(chainId *big.Int) (struct {
	GasZRC20 common.Address
	Registry []byte
}, error) {
	return _Registry.Contract.GetChainInfo(&_Registry.CallOpts, chainId)
}

// GetChainInfo is a free data retrieval call binding the contract method 0x804ea334.
//
// Solidity: function getChainInfo(uint256 chainId) view returns(address gasZRC20, bytes registry)
func (_Registry *RegistryCallerSession) GetChainInfo(chainId *big.Int) (struct {
	GasZRC20 common.Address
	Registry []byte
}, error) {
	return _Registry.Contract.GetChainInfo(&_Registry.CallOpts, chainId)
}

// GetChainMetadata is a free data retrieval call binding the contract method 0x7066b18d.
//
// Solidity: function getChainMetadata(uint256 chainId, string key) view returns(bytes)
func (_Registry *RegistryCaller) GetChainMetadata(opts *bind.CallOpts, chainId *big.Int, key string) ([]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getChainMetadata", chainId, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetChainMetadata is a free data retrieval call binding the contract method 0x7066b18d.
//
// Solidity: function getChainMetadata(uint256 chainId, string key) view returns(bytes)
func (_Registry *RegistrySession) GetChainMetadata(chainId *big.Int, key string) ([]byte, error) {
	return _Registry.Contract.GetChainMetadata(&_Registry.CallOpts, chainId, key)
}

// GetChainMetadata is a free data retrieval call binding the contract method 0x7066b18d.
//
// Solidity: function getChainMetadata(uint256 chainId, string key) view returns(bytes)
func (_Registry *RegistryCallerSession) GetChainMetadata(chainId *big.Int, key string) ([]byte, error) {
	return _Registry.Contract.GetChainMetadata(&_Registry.CallOpts, chainId, key)
}

// GetContractConfiguration is a free data retrieval call binding the contract method 0xd3523ea2.
//
// Solidity: function getContractConfiguration(uint256 chainId, string contractType, string key) view returns(bytes)
func (_Registry *RegistryCaller) GetContractConfiguration(opts *bind.CallOpts, chainId *big.Int, contractType string, key string) ([]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getContractConfiguration", chainId, contractType, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetContractConfiguration is a free data retrieval call binding the contract method 0xd3523ea2.
//
// Solidity: function getContractConfiguration(uint256 chainId, string contractType, string key) view returns(bytes)
func (_Registry *RegistrySession) GetContractConfiguration(chainId *big.Int, contractType string, key string) ([]byte, error) {
	return _Registry.Contract.GetContractConfiguration(&_Registry.CallOpts, chainId, contractType, key)
}

// GetContractConfiguration is a free data retrieval call binding the contract method 0xd3523ea2.
//
// Solidity: function getContractConfiguration(uint256 chainId, string contractType, string key) view returns(bytes)
func (_Registry *RegistryCallerSession) GetContractConfiguration(chainId *big.Int, contractType string, key string) ([]byte, error) {
	return _Registry.Contract.GetContractConfiguration(&_Registry.CallOpts, chainId, contractType, key)
}

// GetContractInfo is a free data retrieval call binding the contract method 0x5cf92c9f.
//
// Solidity: function getContractInfo(uint256 chainId, string contractType) view returns(bool active, bytes addressBytes)
func (_Registry *RegistryCaller) GetContractInfo(opts *bind.CallOpts, chainId *big.Int, contractType string) (struct {
	Active       bool
	AddressBytes []byte
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getContractInfo", chainId, contractType)

	outstruct := new(struct {
		Active       bool
		AddressBytes []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.AddressBytes = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetContractInfo is a free data retrieval call binding the contract method 0x5cf92c9f.
//
// Solidity: function getContractInfo(uint256 chainId, string contractType) view returns(bool active, bytes addressBytes)
func (_Registry *RegistrySession) GetContractInfo(chainId *big.Int, contractType string) (struct {
	Active       bool
	AddressBytes []byte
}, error) {
	return _Registry.Contract.GetContractInfo(&_Registry.CallOpts, chainId, contractType)
}

// GetContractInfo is a free data retrieval call binding the contract method 0x5cf92c9f.
//
// Solidity: function getContractInfo(uint256 chainId, string contractType) view returns(bool active, bytes addressBytes)
func (_Registry *RegistryCallerSession) GetContractInfo(chainId *big.Int, contractType string) (struct {
	Active       bool
	AddressBytes []byte
}, error) {
	return _Registry.Contract.GetContractInfo(&_Registry.CallOpts, chainId, contractType)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registry *RegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registry *RegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Registry.Contract.GetRoleAdmin(&_Registry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registry *RegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Registry.Contract.GetRoleAdmin(&_Registry.CallOpts, role)
}

// GetZRC20AddressByForeignAsset is a free data retrieval call binding the contract method 0xaa808c06.
//
// Solidity: function getZRC20AddressByForeignAsset(uint256 originChainId, bytes originAddress) view returns(address)
func (_Registry *RegistryCaller) GetZRC20AddressByForeignAsset(opts *bind.CallOpts, originChainId *big.Int, originAddress []byte) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getZRC20AddressByForeignAsset", originChainId, originAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetZRC20AddressByForeignAsset is a free data retrieval call binding the contract method 0xaa808c06.
//
// Solidity: function getZRC20AddressByForeignAsset(uint256 originChainId, bytes originAddress) view returns(address)
func (_Registry *RegistrySession) GetZRC20AddressByForeignAsset(originChainId *big.Int, originAddress []byte) (common.Address, error) {
	return _Registry.Contract.GetZRC20AddressByForeignAsset(&_Registry.CallOpts, originChainId, originAddress)
}

// GetZRC20AddressByForeignAsset is a free data retrieval call binding the contract method 0xaa808c06.
//
// Solidity: function getZRC20AddressByForeignAsset(uint256 originChainId, bytes originAddress) view returns(address)
func (_Registry *RegistryCallerSession) GetZRC20AddressByForeignAsset(originChainId *big.Int, originAddress []byte) (common.Address, error) {
	return _Registry.Contract.GetZRC20AddressByForeignAsset(&_Registry.CallOpts, originChainId, originAddress)
}

// GetZRC20TokenInfo is a free data retrieval call binding the contract method 0xe9d6c5ba.
//
// Solidity: function getZRC20TokenInfo(address address_) view returns(bool active, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals)
func (_Registry *RegistryCaller) GetZRC20TokenInfo(opts *bind.CallOpts, address_ common.Address) (struct {
	Active        bool
	Symbol        string
	OriginChainId *big.Int
	OriginAddress []byte
	CoinType      string
	Decimals      uint8
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getZRC20TokenInfo", address_)

	outstruct := new(struct {
		Active        bool
		Symbol        string
		OriginChainId *big.Int
		OriginAddress []byte
		CoinType      string
		Decimals      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.OriginChainId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OriginAddress = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.CoinType = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Decimals = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetZRC20TokenInfo is a free data retrieval call binding the contract method 0xe9d6c5ba.
//
// Solidity: function getZRC20TokenInfo(address address_) view returns(bool active, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals)
func (_Registry *RegistrySession) GetZRC20TokenInfo(address_ common.Address) (struct {
	Active        bool
	Symbol        string
	OriginChainId *big.Int
	OriginAddress []byte
	CoinType      string
	Decimals      uint8
}, error) {
	return _Registry.Contract.GetZRC20TokenInfo(&_Registry.CallOpts, address_)
}

// GetZRC20TokenInfo is a free data retrieval call binding the contract method 0xe9d6c5ba.
//
// Solidity: function getZRC20TokenInfo(address address_) view returns(bool active, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals)
func (_Registry *RegistryCallerSession) GetZRC20TokenInfo(address_ common.Address) (struct {
	Active        bool
	Symbol        string
	OriginChainId *big.Int
	OriginAddress []byte
	CoinType      string
	Decimals      uint8
}, error) {
	return _Registry.Contract.GetZRC20TokenInfo(&_Registry.CallOpts, address_)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Registry.Contract.HasRole(&_Registry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registry *RegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Registry.Contract.HasRole(&_Registry.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registry *RegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registry *RegistrySession) Paused() (bool, error) {
	return _Registry.Contract.Paused(&_Registry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registry *RegistryCallerSession) Paused() (bool, error) {
	return _Registry.Contract.Paused(&_Registry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Registry *RegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Registry *RegistrySession) ProxiableUUID() ([32]byte, error) {
	return _Registry.Contract.ProxiableUUID(&_Registry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Registry *RegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Registry.Contract.ProxiableUUID(&_Registry.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_Registry *RegistryCaller) RegistryManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "registryManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_Registry *RegistrySession) RegistryManager() (common.Address, error) {
	return _Registry.Contract.RegistryManager(&_Registry.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_Registry *RegistryCallerSession) RegistryManager() (common.Address, error) {
	return _Registry.Contract.RegistryManager(&_Registry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Registry.Contract.SupportsInterface(&_Registry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registry *RegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Registry.Contract.SupportsInterface(&_Registry.CallOpts, interfaceId)
}

// BootstrapChains is a paid mutator transaction binding the contract method 0xdc07f866.
//
// Solidity: function bootstrapChains((bool,uint256,address,bytes)[] chains, (uint256,string,bytes)[] metadataEntries) returns()
func (_Registry *RegistryTransactor) BootstrapChains(opts *bind.TransactOpts, chains []ChainInfoDTO, metadataEntries []IRegistryChainMetadataEntry) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "bootstrapChains", chains, metadataEntries)
}

// BootstrapChains is a paid mutator transaction binding the contract method 0xdc07f866.
//
// Solidity: function bootstrapChains((bool,uint256,address,bytes)[] chains, (uint256,string,bytes)[] metadataEntries) returns()
func (_Registry *RegistrySession) BootstrapChains(chains []ChainInfoDTO, metadataEntries []IRegistryChainMetadataEntry) (*types.Transaction, error) {
	return _Registry.Contract.BootstrapChains(&_Registry.TransactOpts, chains, metadataEntries)
}

// BootstrapChains is a paid mutator transaction binding the contract method 0xdc07f866.
//
// Solidity: function bootstrapChains((bool,uint256,address,bytes)[] chains, (uint256,string,bytes)[] metadataEntries) returns()
func (_Registry *RegistryTransactorSession) BootstrapChains(chains []ChainInfoDTO, metadataEntries []IRegistryChainMetadataEntry) (*types.Transaction, error) {
	return _Registry.Contract.BootstrapChains(&_Registry.TransactOpts, chains, metadataEntries)
}

// BootstrapContracts is a paid mutator transaction binding the contract method 0xa4b7e5f9.
//
// Solidity: function bootstrapContracts((bool,bytes,string,uint256)[] contracts, (uint256,string,string,bytes)[] configEntries) returns()
func (_Registry *RegistryTransactor) BootstrapContracts(opts *bind.TransactOpts, contracts []ContractInfoDTO, configEntries []IRegistryContractConfigEntry) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "bootstrapContracts", contracts, configEntries)
}

// BootstrapContracts is a paid mutator transaction binding the contract method 0xa4b7e5f9.
//
// Solidity: function bootstrapContracts((bool,bytes,string,uint256)[] contracts, (uint256,string,string,bytes)[] configEntries) returns()
func (_Registry *RegistrySession) BootstrapContracts(contracts []ContractInfoDTO, configEntries []IRegistryContractConfigEntry) (*types.Transaction, error) {
	return _Registry.Contract.BootstrapContracts(&_Registry.TransactOpts, contracts, configEntries)
}

// BootstrapContracts is a paid mutator transaction binding the contract method 0xa4b7e5f9.
//
// Solidity: function bootstrapContracts((bool,bytes,string,uint256)[] contracts, (uint256,string,string,bytes)[] configEntries) returns()
func (_Registry *RegistryTransactorSession) BootstrapContracts(contracts []ContractInfoDTO, configEntries []IRegistryContractConfigEntry) (*types.Transaction, error) {
	return _Registry.Contract.BootstrapContracts(&_Registry.TransactOpts, contracts, configEntries)
}

// BootstrapZRC20Tokens is a paid mutator transaction binding the contract method 0x909b6a03.
//
// Solidity: function bootstrapZRC20Tokens((bool,address,bytes,uint256,string,string,uint8)[] tokens) returns()
func (_Registry *RegistryTransactor) BootstrapZRC20Tokens(opts *bind.TransactOpts, tokens []ZRC20Info) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "bootstrapZRC20Tokens", tokens)
}

// BootstrapZRC20Tokens is a paid mutator transaction binding the contract method 0x909b6a03.
//
// Solidity: function bootstrapZRC20Tokens((bool,address,bytes,uint256,string,string,uint8)[] tokens) returns()
func (_Registry *RegistrySession) BootstrapZRC20Tokens(tokens []ZRC20Info) (*types.Transaction, error) {
	return _Registry.Contract.BootstrapZRC20Tokens(&_Registry.TransactOpts, tokens)
}

// BootstrapZRC20Tokens is a paid mutator transaction binding the contract method 0x909b6a03.
//
// Solidity: function bootstrapZRC20Tokens((bool,address,bytes,uint256,string,string,uint8)[] tokens) returns()
func (_Registry *RegistryTransactorSession) BootstrapZRC20Tokens(tokens []ZRC20Info) (*types.Transaction, error) {
	return _Registry.Contract.BootstrapZRC20Tokens(&_Registry.TransactOpts, tokens)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Registry *RegistryTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Registry *RegistrySession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ChangeAdmin(&_Registry.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Registry *RegistryTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ChangeAdmin(&_Registry.TransactOpts, newAdmin)
}

// ChangeChainStatus is a paid mutator transaction binding the contract method 0xa8f2cb96.
//
// Solidity: function changeChainStatus(uint256 chainId, address gasZRC20, bytes registry, bool activation) returns()
func (_Registry *RegistryTransactor) ChangeChainStatus(opts *bind.TransactOpts, chainId *big.Int, gasZRC20 common.Address, registry []byte, activation bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "changeChainStatus", chainId, gasZRC20, registry, activation)
}

// ChangeChainStatus is a paid mutator transaction binding the contract method 0xa8f2cb96.
//
// Solidity: function changeChainStatus(uint256 chainId, address gasZRC20, bytes registry, bool activation) returns()
func (_Registry *RegistrySession) ChangeChainStatus(chainId *big.Int, gasZRC20 common.Address, registry []byte, activation bool) (*types.Transaction, error) {
	return _Registry.Contract.ChangeChainStatus(&_Registry.TransactOpts, chainId, gasZRC20, registry, activation)
}

// ChangeChainStatus is a paid mutator transaction binding the contract method 0xa8f2cb96.
//
// Solidity: function changeChainStatus(uint256 chainId, address gasZRC20, bytes registry, bool activation) returns()
func (_Registry *RegistryTransactorSession) ChangeChainStatus(chainId *big.Int, gasZRC20 common.Address, registry []byte, activation bool) (*types.Transaction, error) {
	return _Registry.Contract.ChangeChainStatus(&_Registry.TransactOpts, chainId, gasZRC20, registry, activation)
}

// ChangeRegistryManager is a paid mutator transaction binding the contract method 0x3500c24b.
//
// Solidity: function changeRegistryManager(address newRegistryManager) returns()
func (_Registry *RegistryTransactor) ChangeRegistryManager(opts *bind.TransactOpts, newRegistryManager common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "changeRegistryManager", newRegistryManager)
}

// ChangeRegistryManager is a paid mutator transaction binding the contract method 0x3500c24b.
//
// Solidity: function changeRegistryManager(address newRegistryManager) returns()
func (_Registry *RegistrySession) ChangeRegistryManager(newRegistryManager common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ChangeRegistryManager(&_Registry.TransactOpts, newRegistryManager)
}

// ChangeRegistryManager is a paid mutator transaction binding the contract method 0x3500c24b.
//
// Solidity: function changeRegistryManager(address newRegistryManager) returns()
func (_Registry *RegistryTransactorSession) ChangeRegistryManager(newRegistryManager common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ChangeRegistryManager(&_Registry.TransactOpts, newRegistryManager)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.GrantRole(&_Registry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.GrantRole(&_Registry.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address admin_, address registryManager_, address gatewayEVM_, address coreRegistry_) returns()
func (_Registry *RegistryTransactor) Initialize(opts *bind.TransactOpts, admin_ common.Address, registryManager_ common.Address, gatewayEVM_ common.Address, coreRegistry_ common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "initialize", admin_, registryManager_, gatewayEVM_, coreRegistry_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address admin_, address registryManager_, address gatewayEVM_, address coreRegistry_) returns()
func (_Registry *RegistrySession) Initialize(admin_ common.Address, registryManager_ common.Address, gatewayEVM_ common.Address, coreRegistry_ common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts, admin_, registryManager_, gatewayEVM_, coreRegistry_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address admin_, address registryManager_, address gatewayEVM_, address coreRegistry_) returns()
func (_Registry *RegistryTransactorSession) Initialize(admin_ common.Address, registryManager_ common.Address, gatewayEVM_ common.Address, coreRegistry_ common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts, admin_, registryManager_, gatewayEVM_, coreRegistry_)
}

// OnCall is a paid mutator transaction binding the contract method 0x676cc054.
//
// Solidity: function onCall((address) context, bytes data) returns(bytes)
func (_Registry *RegistryTransactor) OnCall(opts *bind.TransactOpts, context MessageContext, data []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "onCall", context, data)
}

// OnCall is a paid mutator transaction binding the contract method 0x676cc054.
//
// Solidity: function onCall((address) context, bytes data) returns(bytes)
func (_Registry *RegistrySession) OnCall(context MessageContext, data []byte) (*types.Transaction, error) {
	return _Registry.Contract.OnCall(&_Registry.TransactOpts, context, data)
}

// OnCall is a paid mutator transaction binding the contract method 0x676cc054.
//
// Solidity: function onCall((address) context, bytes data) returns(bytes)
func (_Registry *RegistryTransactorSession) OnCall(context MessageContext, data []byte) (*types.Transaction, error) {
	return _Registry.Contract.OnCall(&_Registry.TransactOpts, context, data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registry *RegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registry *RegistrySession) Pause() (*types.Transaction, error) {
	return _Registry.Contract.Pause(&_Registry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registry *RegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _Registry.Contract.Pause(&_Registry.TransactOpts)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x631d62e4.
//
// Solidity: function registerContract(uint256 chainId, string contractType, bytes addressBytes) returns()
func (_Registry *RegistryTransactor) RegisterContract(opts *bind.TransactOpts, chainId *big.Int, contractType string, addressBytes []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerContract", chainId, contractType, addressBytes)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x631d62e4.
//
// Solidity: function registerContract(uint256 chainId, string contractType, bytes addressBytes) returns()
func (_Registry *RegistrySession) RegisterContract(chainId *big.Int, contractType string, addressBytes []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterContract(&_Registry.TransactOpts, chainId, contractType, addressBytes)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x631d62e4.
//
// Solidity: function registerContract(uint256 chainId, string contractType, bytes addressBytes) returns()
func (_Registry *RegistryTransactorSession) RegisterContract(chainId *big.Int, contractType string, addressBytes []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterContract(&_Registry.TransactOpts, chainId, contractType, addressBytes)
}

// RegisterZRC20Token is a paid mutator transaction binding the contract method 0x6e9e2d3f.
//
// Solidity: function registerZRC20Token(address address_, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals) returns()
func (_Registry *RegistryTransactor) RegisterZRC20Token(opts *bind.TransactOpts, address_ common.Address, symbol string, originChainId *big.Int, originAddress []byte, coinType string, decimals uint8) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerZRC20Token", address_, symbol, originChainId, originAddress, coinType, decimals)
}

// RegisterZRC20Token is a paid mutator transaction binding the contract method 0x6e9e2d3f.
//
// Solidity: function registerZRC20Token(address address_, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals) returns()
func (_Registry *RegistrySession) RegisterZRC20Token(address_ common.Address, symbol string, originChainId *big.Int, originAddress []byte, coinType string, decimals uint8) (*types.Transaction, error) {
	return _Registry.Contract.RegisterZRC20Token(&_Registry.TransactOpts, address_, symbol, originChainId, originAddress, coinType, decimals)
}

// RegisterZRC20Token is a paid mutator transaction binding the contract method 0x6e9e2d3f.
//
// Solidity: function registerZRC20Token(address address_, string symbol, uint256 originChainId, bytes originAddress, string coinType, uint8 decimals) returns()
func (_Registry *RegistryTransactorSession) RegisterZRC20Token(address_ common.Address, symbol string, originChainId *big.Int, originAddress []byte, coinType string, decimals uint8) (*types.Transaction, error) {
	return _Registry.Contract.RegisterZRC20Token(&_Registry.TransactOpts, address_, symbol, originChainId, originAddress, coinType, decimals)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Registry *RegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Registry *RegistrySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RenounceRole(&_Registry.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Registry *RegistryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RenounceRole(&_Registry.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RevokeRole(&_Registry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registry *RegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RevokeRole(&_Registry.TransactOpts, role, account)
}

// SetContractActive is a paid mutator transaction binding the contract method 0x10d29b9e.
//
// Solidity: function setContractActive(uint256 chainId, string contractType, bool active) returns()
func (_Registry *RegistryTransactor) SetContractActive(opts *bind.TransactOpts, chainId *big.Int, contractType string, active bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setContractActive", chainId, contractType, active)
}

// SetContractActive is a paid mutator transaction binding the contract method 0x10d29b9e.
//
// Solidity: function setContractActive(uint256 chainId, string contractType, bool active) returns()
func (_Registry *RegistrySession) SetContractActive(chainId *big.Int, contractType string, active bool) (*types.Transaction, error) {
	return _Registry.Contract.SetContractActive(&_Registry.TransactOpts, chainId, contractType, active)
}

// SetContractActive is a paid mutator transaction binding the contract method 0x10d29b9e.
//
// Solidity: function setContractActive(uint256 chainId, string contractType, bool active) returns()
func (_Registry *RegistryTransactorSession) SetContractActive(chainId *big.Int, contractType string, active bool) (*types.Transaction, error) {
	return _Registry.Contract.SetContractActive(&_Registry.TransactOpts, chainId, contractType, active)
}

// SetZRC20TokenActive is a paid mutator transaction binding the contract method 0x9060bda9.
//
// Solidity: function setZRC20TokenActive(address address_, bool active) returns()
func (_Registry *RegistryTransactor) SetZRC20TokenActive(opts *bind.TransactOpts, address_ common.Address, active bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setZRC20TokenActive", address_, active)
}

// SetZRC20TokenActive is a paid mutator transaction binding the contract method 0x9060bda9.
//
// Solidity: function setZRC20TokenActive(address address_, bool active) returns()
func (_Registry *RegistrySession) SetZRC20TokenActive(address_ common.Address, active bool) (*types.Transaction, error) {
	return _Registry.Contract.SetZRC20TokenActive(&_Registry.TransactOpts, address_, active)
}

// SetZRC20TokenActive is a paid mutator transaction binding the contract method 0x9060bda9.
//
// Solidity: function setZRC20TokenActive(address address_, bool active) returns()
func (_Registry *RegistryTransactorSession) SetZRC20TokenActive(address_ common.Address, active bool) (*types.Transaction, error) {
	return _Registry.Contract.SetZRC20TokenActive(&_Registry.TransactOpts, address_, active)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registry *RegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registry *RegistrySession) Unpause() (*types.Transaction, error) {
	return _Registry.Contract.Unpause(&_Registry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registry *RegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _Registry.Contract.Unpause(&_Registry.TransactOpts)
}

// UpdateChainMetadata is a paid mutator transaction binding the contract method 0x2259e9e5.
//
// Solidity: function updateChainMetadata(uint256 chainId, string key, bytes value) returns()
func (_Registry *RegistryTransactor) UpdateChainMetadata(opts *bind.TransactOpts, chainId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateChainMetadata", chainId, key, value)
}

// UpdateChainMetadata is a paid mutator transaction binding the contract method 0x2259e9e5.
//
// Solidity: function updateChainMetadata(uint256 chainId, string key, bytes value) returns()
func (_Registry *RegistrySession) UpdateChainMetadata(chainId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateChainMetadata(&_Registry.TransactOpts, chainId, key, value)
}

// UpdateChainMetadata is a paid mutator transaction binding the contract method 0x2259e9e5.
//
// Solidity: function updateChainMetadata(uint256 chainId, string key, bytes value) returns()
func (_Registry *RegistryTransactorSession) UpdateChainMetadata(chainId *big.Int, key string, value []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateChainMetadata(&_Registry.TransactOpts, chainId, key, value)
}

// UpdateContractConfiguration is a paid mutator transaction binding the contract method 0xf354b31f.
//
// Solidity: function updateContractConfiguration(uint256 chainId, string contractType, string key, bytes value) returns()
func (_Registry *RegistryTransactor) UpdateContractConfiguration(opts *bind.TransactOpts, chainId *big.Int, contractType string, key string, value []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateContractConfiguration", chainId, contractType, key, value)
}

// UpdateContractConfiguration is a paid mutator transaction binding the contract method 0xf354b31f.
//
// Solidity: function updateContractConfiguration(uint256 chainId, string contractType, string key, bytes value) returns()
func (_Registry *RegistrySession) UpdateContractConfiguration(chainId *big.Int, contractType string, key string, value []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateContractConfiguration(&_Registry.TransactOpts, chainId, contractType, key, value)
}

// UpdateContractConfiguration is a paid mutator transaction binding the contract method 0xf354b31f.
//
// Solidity: function updateContractConfiguration(uint256 chainId, string contractType, string key, bytes value) returns()
func (_Registry *RegistryTransactorSession) UpdateContractConfiguration(chainId *big.Int, contractType string, key string, value []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpdateContractConfiguration(&_Registry.TransactOpts, chainId, contractType, key, value)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Registry *RegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Registry *RegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpgradeToAndCall(&_Registry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Registry *RegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Registry.Contract.UpgradeToAndCall(&_Registry.TransactOpts, newImplementation, data)
}

// RegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Registry contract.
type RegistryAdminChangedIterator struct {
	Event *RegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *RegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryAdminChanged)
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
		it.Event = new(RegistryAdminChanged)
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
func (it *RegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryAdminChanged represents a AdminChanged event raised by the Registry contract.
type RegistryAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address oldAdmin, address newAdmin)
func (_Registry *RegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*RegistryAdminChangedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &RegistryAdminChangedIterator{contract: _Registry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address oldAdmin, address newAdmin)
func (_Registry *RegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *RegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryAdminChanged)
				if err := _Registry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address oldAdmin, address newAdmin)
func (_Registry *RegistryFilterer) ParseAdminChanged(log types.Log) (*RegistryAdminChanged, error) {
	event := new(RegistryAdminChanged)
	if err := _Registry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryChainMetadataUpdatedIterator is returned from FilterChainMetadataUpdated and is used to iterate over the raw logs and unpacked data for ChainMetadataUpdated events raised by the Registry contract.
type RegistryChainMetadataUpdatedIterator struct {
	Event *RegistryChainMetadataUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryChainMetadataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryChainMetadataUpdated)
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
		it.Event = new(RegistryChainMetadataUpdated)
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
func (it *RegistryChainMetadataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryChainMetadataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryChainMetadataUpdated represents a ChainMetadataUpdated event raised by the Registry contract.
type RegistryChainMetadataUpdated struct {
	ChainId *big.Int
	Key     string
	Value   []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainMetadataUpdated is a free log retrieval operation binding the contract event 0x40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee634.
//
// Solidity: event ChainMetadataUpdated(uint256 indexed chainId, string key, bytes value)
func (_Registry *RegistryFilterer) FilterChainMetadataUpdated(opts *bind.FilterOpts, chainId []*big.Int) (*RegistryChainMetadataUpdatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ChainMetadataUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryChainMetadataUpdatedIterator{contract: _Registry.contract, event: "ChainMetadataUpdated", logs: logs, sub: sub}, nil
}

// WatchChainMetadataUpdated is a free log subscription operation binding the contract event 0x40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee634.
//
// Solidity: event ChainMetadataUpdated(uint256 indexed chainId, string key, bytes value)
func (_Registry *RegistryFilterer) WatchChainMetadataUpdated(opts *bind.WatchOpts, sink chan<- *RegistryChainMetadataUpdated, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ChainMetadataUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryChainMetadataUpdated)
				if err := _Registry.contract.UnpackLog(event, "ChainMetadataUpdated", log); err != nil {
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

// ParseChainMetadataUpdated is a log parse operation binding the contract event 0x40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee634.
//
// Solidity: event ChainMetadataUpdated(uint256 indexed chainId, string key, bytes value)
func (_Registry *RegistryFilterer) ParseChainMetadataUpdated(log types.Log) (*RegistryChainMetadataUpdated, error) {
	event := new(RegistryChainMetadataUpdated)
	if err := _Registry.contract.UnpackLog(event, "ChainMetadataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryChainStatusChangedIterator is returned from FilterChainStatusChanged and is used to iterate over the raw logs and unpacked data for ChainStatusChanged events raised by the Registry contract.
type RegistryChainStatusChangedIterator struct {
	Event *RegistryChainStatusChanged // Event containing the contract specifics and raw log

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
func (it *RegistryChainStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryChainStatusChanged)
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
		it.Event = new(RegistryChainStatusChanged)
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
func (it *RegistryChainStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryChainStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryChainStatusChanged represents a ChainStatusChanged event raised by the Registry contract.
type RegistryChainStatusChanged struct {
	ChainId   *big.Int
	NewStatus bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChainStatusChanged is a free log retrieval operation binding the contract event 0xc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e.
//
// Solidity: event ChainStatusChanged(uint256 indexed chainId, bool newStatus)
func (_Registry *RegistryFilterer) FilterChainStatusChanged(opts *bind.FilterOpts, chainId []*big.Int) (*RegistryChainStatusChangedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ChainStatusChanged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryChainStatusChangedIterator{contract: _Registry.contract, event: "ChainStatusChanged", logs: logs, sub: sub}, nil
}

// WatchChainStatusChanged is a free log subscription operation binding the contract event 0xc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e.
//
// Solidity: event ChainStatusChanged(uint256 indexed chainId, bool newStatus)
func (_Registry *RegistryFilterer) WatchChainStatusChanged(opts *bind.WatchOpts, sink chan<- *RegistryChainStatusChanged, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ChainStatusChanged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryChainStatusChanged)
				if err := _Registry.contract.UnpackLog(event, "ChainStatusChanged", log); err != nil {
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

// ParseChainStatusChanged is a log parse operation binding the contract event 0xc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e.
//
// Solidity: event ChainStatusChanged(uint256 indexed chainId, bool newStatus)
func (_Registry *RegistryFilterer) ParseChainStatusChanged(log types.Log) (*RegistryChainStatusChanged, error) {
	event := new(RegistryChainStatusChanged)
	if err := _Registry.contract.UnpackLog(event, "ChainStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryContractConfigurationUpdatedIterator is returned from FilterContractConfigurationUpdated and is used to iterate over the raw logs and unpacked data for ContractConfigurationUpdated events raised by the Registry contract.
type RegistryContractConfigurationUpdatedIterator struct {
	Event *RegistryContractConfigurationUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryContractConfigurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractConfigurationUpdated)
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
		it.Event = new(RegistryContractConfigurationUpdated)
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
func (it *RegistryContractConfigurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractConfigurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractConfigurationUpdated represents a ContractConfigurationUpdated event raised by the Registry contract.
type RegistryContractConfigurationUpdated struct {
	ChainId      *big.Int
	ContractType string
	Key          string
	Value        []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractConfigurationUpdated is a free log retrieval operation binding the contract event 0xaea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c5.
//
// Solidity: event ContractConfigurationUpdated(uint256 indexed chainId, string contractType, string key, bytes value)
func (_Registry *RegistryFilterer) FilterContractConfigurationUpdated(opts *bind.FilterOpts, chainId []*big.Int) (*RegistryContractConfigurationUpdatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ContractConfigurationUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistryContractConfigurationUpdatedIterator{contract: _Registry.contract, event: "ContractConfigurationUpdated", logs: logs, sub: sub}, nil
}

// WatchContractConfigurationUpdated is a free log subscription operation binding the contract event 0xaea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c5.
//
// Solidity: event ContractConfigurationUpdated(uint256 indexed chainId, string contractType, string key, bytes value)
func (_Registry *RegistryFilterer) WatchContractConfigurationUpdated(opts *bind.WatchOpts, sink chan<- *RegistryContractConfigurationUpdated, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ContractConfigurationUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractConfigurationUpdated)
				if err := _Registry.contract.UnpackLog(event, "ContractConfigurationUpdated", log); err != nil {
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

// ParseContractConfigurationUpdated is a log parse operation binding the contract event 0xaea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c5.
//
// Solidity: event ContractConfigurationUpdated(uint256 indexed chainId, string contractType, string key, bytes value)
func (_Registry *RegistryFilterer) ParseContractConfigurationUpdated(log types.Log) (*RegistryContractConfigurationUpdated, error) {
	event := new(RegistryContractConfigurationUpdated)
	if err := _Registry.contract.UnpackLog(event, "ContractConfigurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryContractRegisteredIterator is returned from FilterContractRegistered and is used to iterate over the raw logs and unpacked data for ContractRegistered events raised by the Registry contract.
type RegistryContractRegisteredIterator struct {
	Event *RegistryContractRegistered // Event containing the contract specifics and raw log

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
func (it *RegistryContractRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractRegistered)
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
		it.Event = new(RegistryContractRegistered)
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
func (it *RegistryContractRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractRegistered represents a ContractRegistered event raised by the Registry contract.
type RegistryContractRegistered struct {
	ChainId      *big.Int
	ContractType common.Hash
	AddressBytes []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractRegistered is a free log retrieval operation binding the contract event 0x20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581.
//
// Solidity: event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes)
func (_Registry *RegistryFilterer) FilterContractRegistered(opts *bind.FilterOpts, chainId []*big.Int, contractType []string) (*RegistryContractRegisteredIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ContractRegistered", chainIdRule, contractTypeRule)
	if err != nil {
		return nil, err
	}
	return &RegistryContractRegisteredIterator{contract: _Registry.contract, event: "ContractRegistered", logs: logs, sub: sub}, nil
}

// WatchContractRegistered is a free log subscription operation binding the contract event 0x20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581.
//
// Solidity: event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes)
func (_Registry *RegistryFilterer) WatchContractRegistered(opts *bind.WatchOpts, sink chan<- *RegistryContractRegistered, chainId []*big.Int, contractType []string) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ContractRegistered", chainIdRule, contractTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractRegistered)
				if err := _Registry.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
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

// ParseContractRegistered is a log parse operation binding the contract event 0x20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581.
//
// Solidity: event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes)
func (_Registry *RegistryFilterer) ParseContractRegistered(log types.Log) (*RegistryContractRegistered, error) {
	event := new(RegistryContractRegistered)
	if err := _Registry.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryContractStatusChangedIterator is returned from FilterContractStatusChanged and is used to iterate over the raw logs and unpacked data for ContractStatusChanged events raised by the Registry contract.
type RegistryContractStatusChangedIterator struct {
	Event *RegistryContractStatusChanged // Event containing the contract specifics and raw log

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
func (it *RegistryContractStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryContractStatusChanged)
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
		it.Event = new(RegistryContractStatusChanged)
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
func (it *RegistryContractStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryContractStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryContractStatusChanged represents a ContractStatusChanged event raised by the Registry contract.
type RegistryContractStatusChanged struct {
	AddressBytes []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractStatusChanged is a free log retrieval operation binding the contract event 0x6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c.
//
// Solidity: event ContractStatusChanged(bytes addressBytes)
func (_Registry *RegistryFilterer) FilterContractStatusChanged(opts *bind.FilterOpts) (*RegistryContractStatusChangedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ContractStatusChanged")
	if err != nil {
		return nil, err
	}
	return &RegistryContractStatusChangedIterator{contract: _Registry.contract, event: "ContractStatusChanged", logs: logs, sub: sub}, nil
}

// WatchContractStatusChanged is a free log subscription operation binding the contract event 0x6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c.
//
// Solidity: event ContractStatusChanged(bytes addressBytes)
func (_Registry *RegistryFilterer) WatchContractStatusChanged(opts *bind.WatchOpts, sink chan<- *RegistryContractStatusChanged) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ContractStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryContractStatusChanged)
				if err := _Registry.contract.UnpackLog(event, "ContractStatusChanged", log); err != nil {
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

// ParseContractStatusChanged is a log parse operation binding the contract event 0x6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c.
//
// Solidity: event ContractStatusChanged(bytes addressBytes)
func (_Registry *RegistryFilterer) ParseContractStatusChanged(log types.Log) (*RegistryContractStatusChanged, error) {
	event := new(RegistryContractStatusChanged)
	if err := _Registry.contract.UnpackLog(event, "ContractStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Registry contract.
type RegistryInitializedIterator struct {
	Event *RegistryInitialized // Event containing the contract specifics and raw log

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
func (it *RegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryInitialized)
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
		it.Event = new(RegistryInitialized)
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
func (it *RegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryInitialized represents a Initialized event raised by the Registry contract.
type RegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Registry *RegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*RegistryInitializedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RegistryInitializedIterator{contract: _Registry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Registry *RegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryInitialized)
				if err := _Registry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseInitialized(log types.Log) (*RegistryInitialized, error) {
	event := new(RegistryInitialized)
	if err := _Registry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Registry contract.
type RegistryPausedIterator struct {
	Event *RegistryPaused // Event containing the contract specifics and raw log

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
func (it *RegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryPaused)
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
		it.Event = new(RegistryPaused)
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
func (it *RegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryPaused represents a Paused event raised by the Registry contract.
type RegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Registry *RegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*RegistryPausedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RegistryPausedIterator{contract: _Registry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Registry *RegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RegistryPaused) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryPaused)
				if err := _Registry.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Registry *RegistryFilterer) ParsePaused(log types.Log) (*RegistryPaused, error) {
	event := new(RegistryPaused)
	if err := _Registry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRegistryManagerChangedIterator is returned from FilterRegistryManagerChanged and is used to iterate over the raw logs and unpacked data for RegistryManagerChanged events raised by the Registry contract.
type RegistryRegistryManagerChangedIterator struct {
	Event *RegistryRegistryManagerChanged // Event containing the contract specifics and raw log

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
func (it *RegistryRegistryManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRegistryManagerChanged)
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
		it.Event = new(RegistryRegistryManagerChanged)
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
func (it *RegistryRegistryManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRegistryManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRegistryManagerChanged represents a RegistryManagerChanged event raised by the Registry contract.
type RegistryRegistryManagerChanged struct {
	OldRegistryManager common.Address
	NewRegistryManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRegistryManagerChanged is a free log retrieval operation binding the contract event 0x6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279.
//
// Solidity: event RegistryManagerChanged(address oldRegistryManager, address newRegistryManager)
func (_Registry *RegistryFilterer) FilterRegistryManagerChanged(opts *bind.FilterOpts) (*RegistryRegistryManagerChangedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RegistryManagerChanged")
	if err != nil {
		return nil, err
	}
	return &RegistryRegistryManagerChangedIterator{contract: _Registry.contract, event: "RegistryManagerChanged", logs: logs, sub: sub}, nil
}

// WatchRegistryManagerChanged is a free log subscription operation binding the contract event 0x6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279.
//
// Solidity: event RegistryManagerChanged(address oldRegistryManager, address newRegistryManager)
func (_Registry *RegistryFilterer) WatchRegistryManagerChanged(opts *bind.WatchOpts, sink chan<- *RegistryRegistryManagerChanged) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RegistryManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRegistryManagerChanged)
				if err := _Registry.contract.UnpackLog(event, "RegistryManagerChanged", log); err != nil {
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

// ParseRegistryManagerChanged is a log parse operation binding the contract event 0x6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279.
//
// Solidity: event RegistryManagerChanged(address oldRegistryManager, address newRegistryManager)
func (_Registry *RegistryFilterer) ParseRegistryManagerChanged(log types.Log) (*RegistryRegistryManagerChanged, error) {
	event := new(RegistryRegistryManagerChanged)
	if err := _Registry.contract.UnpackLog(event, "RegistryManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Registry contract.
type RegistryRoleAdminChangedIterator struct {
	Event *RegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRoleAdminChanged)
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
		it.Event = new(RegistryRoleAdminChanged)
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
func (it *RegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRoleAdminChanged represents a RoleAdminChanged event raised by the Registry contract.
type RegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Registry *RegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRoleAdminChangedIterator{contract: _Registry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Registry *RegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRoleAdminChanged)
				if err := _Registry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseRoleAdminChanged(log types.Log) (*RegistryRoleAdminChanged, error) {
	event := new(RegistryRoleAdminChanged)
	if err := _Registry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Registry contract.
type RegistryRoleGrantedIterator struct {
	Event *RegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *RegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRoleGranted)
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
		it.Event = new(RegistryRoleGranted)
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
func (it *RegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRoleGranted represents a RoleGranted event raised by the Registry contract.
type RegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRoleGrantedIterator{contract: _Registry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRoleGranted)
				if err := _Registry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseRoleGranted(log types.Log) (*RegistryRoleGranted, error) {
	event := new(RegistryRoleGranted)
	if err := _Registry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Registry contract.
type RegistryRoleRevokedIterator struct {
	Event *RegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRoleRevoked)
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
		it.Event = new(RegistryRoleRevoked)
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
func (it *RegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRoleRevoked represents a RoleRevoked event raised by the Registry contract.
type RegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRoleRevokedIterator{contract: _Registry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registry *RegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRoleRevoked)
				if err := _Registry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseRoleRevoked(log types.Log) (*RegistryRoleRevoked, error) {
	event := new(RegistryRoleRevoked)
	if err := _Registry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Registry contract.
type RegistryUnpausedIterator struct {
	Event *RegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *RegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryUnpaused)
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
		it.Event = new(RegistryUnpaused)
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
func (it *RegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryUnpaused represents a Unpaused event raised by the Registry contract.
type RegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Registry *RegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RegistryUnpausedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RegistryUnpausedIterator{contract: _Registry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Registry *RegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryUnpaused)
				if err := _Registry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseUnpaused(log types.Log) (*RegistryUnpaused, error) {
	event := new(RegistryUnpaused)
	if err := _Registry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Registry contract.
type RegistryUpgradedIterator struct {
	Event *RegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *RegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryUpgraded)
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
		it.Event = new(RegistryUpgraded)
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
func (it *RegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryUpgraded represents a Upgraded event raised by the Registry contract.
type RegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Registry *RegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RegistryUpgradedIterator{contract: _Registry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Registry *RegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryUpgraded)
				if err := _Registry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseUpgraded(log types.Log) (*RegistryUpgraded, error) {
	event := new(RegistryUpgraded)
	if err := _Registry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryZRC20TokenRegisteredIterator is returned from FilterZRC20TokenRegistered and is used to iterate over the raw logs and unpacked data for ZRC20TokenRegistered events raised by the Registry contract.
type RegistryZRC20TokenRegisteredIterator struct {
	Event *RegistryZRC20TokenRegistered // Event containing the contract specifics and raw log

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
func (it *RegistryZRC20TokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryZRC20TokenRegistered)
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
		it.Event = new(RegistryZRC20TokenRegistered)
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
func (it *RegistryZRC20TokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryZRC20TokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryZRC20TokenRegistered represents a ZRC20TokenRegistered event raised by the Registry contract.
type RegistryZRC20TokenRegistered struct {
	OriginAddress common.Hash
	Address       common.Address
	Decimals      uint8
	OriginChainId *big.Int
	Symbol        string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterZRC20TokenRegistered is a free log retrieval operation binding the contract event 0xa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367.
//
// Solidity: event ZRC20TokenRegistered(bytes indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol)
func (_Registry *RegistryFilterer) FilterZRC20TokenRegistered(opts *bind.FilterOpts, originAddress [][]byte, address_ []common.Address) (*RegistryZRC20TokenRegisteredIterator, error) {

	var originAddressRule []interface{}
	for _, originAddressItem := range originAddress {
		originAddressRule = append(originAddressRule, originAddressItem)
	}
	var address_Rule []interface{}
	for _, address_Item := range address_ {
		address_Rule = append(address_Rule, address_Item)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ZRC20TokenRegistered", originAddressRule, address_Rule)
	if err != nil {
		return nil, err
	}
	return &RegistryZRC20TokenRegisteredIterator{contract: _Registry.contract, event: "ZRC20TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchZRC20TokenRegistered is a free log subscription operation binding the contract event 0xa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367.
//
// Solidity: event ZRC20TokenRegistered(bytes indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol)
func (_Registry *RegistryFilterer) WatchZRC20TokenRegistered(opts *bind.WatchOpts, sink chan<- *RegistryZRC20TokenRegistered, originAddress [][]byte, address_ []common.Address) (event.Subscription, error) {

	var originAddressRule []interface{}
	for _, originAddressItem := range originAddress {
		originAddressRule = append(originAddressRule, originAddressItem)
	}
	var address_Rule []interface{}
	for _, address_Item := range address_ {
		address_Rule = append(address_Rule, address_Item)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ZRC20TokenRegistered", originAddressRule, address_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryZRC20TokenRegistered)
				if err := _Registry.contract.UnpackLog(event, "ZRC20TokenRegistered", log); err != nil {
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

// ParseZRC20TokenRegistered is a log parse operation binding the contract event 0xa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367.
//
// Solidity: event ZRC20TokenRegistered(bytes indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol)
func (_Registry *RegistryFilterer) ParseZRC20TokenRegistered(log types.Log) (*RegistryZRC20TokenRegistered, error) {
	event := new(RegistryZRC20TokenRegistered)
	if err := _Registry.contract.UnpackLog(event, "ZRC20TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryZRC20TokenUpdatedIterator is returned from FilterZRC20TokenUpdated and is used to iterate over the raw logs and unpacked data for ZRC20TokenUpdated events raised by the Registry contract.
type RegistryZRC20TokenUpdatedIterator struct {
	Event *RegistryZRC20TokenUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryZRC20TokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryZRC20TokenUpdated)
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
		it.Event = new(RegistryZRC20TokenUpdated)
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
func (it *RegistryZRC20TokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryZRC20TokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryZRC20TokenUpdated represents a ZRC20TokenUpdated event raised by the Registry contract.
type RegistryZRC20TokenUpdated struct {
	Address common.Address
	Active  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterZRC20TokenUpdated is a free log retrieval operation binding the contract event 0x9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8.
//
// Solidity: event ZRC20TokenUpdated(address address_, bool active)
func (_Registry *RegistryFilterer) FilterZRC20TokenUpdated(opts *bind.FilterOpts) (*RegistryZRC20TokenUpdatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ZRC20TokenUpdated")
	if err != nil {
		return nil, err
	}
	return &RegistryZRC20TokenUpdatedIterator{contract: _Registry.contract, event: "ZRC20TokenUpdated", logs: logs, sub: sub}, nil
}

// WatchZRC20TokenUpdated is a free log subscription operation binding the contract event 0x9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8.
//
// Solidity: event ZRC20TokenUpdated(address address_, bool active)
func (_Registry *RegistryFilterer) WatchZRC20TokenUpdated(opts *bind.WatchOpts, sink chan<- *RegistryZRC20TokenUpdated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ZRC20TokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryZRC20TokenUpdated)
				if err := _Registry.contract.UnpackLog(event, "ZRC20TokenUpdated", log); err != nil {
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

// ParseZRC20TokenUpdated is a log parse operation binding the contract event 0x9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8.
//
// Solidity: event ZRC20TokenUpdated(address address_, bool active)
func (_Registry *RegistryFilterer) ParseZRC20TokenUpdated(log types.Log) (*RegistryZRC20TokenUpdated, error) {
	event := new(RegistryZRC20TokenUpdated)
	if err := _Registry.contract.UnpackLog(event, "ZRC20TokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
