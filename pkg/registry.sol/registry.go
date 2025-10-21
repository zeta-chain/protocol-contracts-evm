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
	Bin: "0x60a060405230608052348015610013575f80fd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051614c9c6100f95f395f8181612b3701528181612b600152612d120152614c9c5ff3fe6080604052600436106102d7575f3560e01c80638f2839701161017b578063bd8fde1c116100d1578063dc07f86611610087578063f354b31f11610062578063f354b31f14610907578063f851a44014610926578063f8c8765e14610944575f80fd5b8063dc07f86614610884578063e63ab1e9146108a3578063e9d6c5ba146108d6575f80fd5b8063c50d6dd9116100b7578063c50d6dd914610813578063d3523ea214610846578063d547741f14610865575f80fd5b8063bd8fde1c146107bf578063c1bd469f146107f2575f80fd5b80639ca220dd11610131578063a8f2cb961161010c578063a8f2cb9614610739578063aa808c0614610758578063ad3cb1cc14610777575f80fd5b80639ca220dd146106e6578063a217fddf14610707578063a4b7e5f91461071a575f80fd5b8063909b6a0311610161578063909b6a031461064357806391d148541461066257806394cc8683146106c5575f80fd5b80638f283970146106055780639060bda914610624575f80fd5b80634f1ef28611610230578063676cc054116101e65780637066b18d116101c15780637066b18d146105a5578063804ea334146105c45780638456cb59146105f1575f80fd5b8063676cc0541461053b5780636bf3d05a146105675780636e9e2d3f14610586575f80fd5b80635c975abb116102165780635c975abb146104b95780635cf92c9f146104ef578063631d62e41461051c575f80fd5b80634f1ef2861461049257806352d1902d146104a5575f80fd5b8063248a9ca3116102905780633500c24b1161026b5780633500c24b1461044057806336568abe1461045f5780633f4ba83a1461047e575f80fd5b8063248a9ca3146103a75780632c78f74c146104025780632f2ff15d14610421575f80fd5b806310d29b9e116102c057806310d29b9e1461034657806318d3ce96146103675780632259e9e514610388575f80fd5b806301ffc9a7146102db5780630c63109e1461030f575b5f80fd5b3480156102e6575f80fd5b506102fa6102f5366004613c0e565b610963565b60405190151581526020015b60405180910390f35b34801561031a575f80fd5b5060015461032e906001600160a01b031681565b6040516001600160a01b039091168152602001610306565b348015610351575f80fd5b50610365610360366004613ca6565b6109fb565b005b348015610372575f80fd5b5061037b610aae565b6040516103069190613d2e565b348015610393575f80fd5b506103656103a2366004613ded565b610d3f565b3480156103b2575f80fd5b506103f46103c1366004613e66565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b604051908152602001610306565b34801561040d575f80fd5b50600b5461032e906001600160a01b031681565b34801561042c575f80fd5b5061036561043b366004613e93565b610dd2565b34801561044b575f80fd5b5061036561045a366004613ebd565b610e1b565b34801561046a575f80fd5b50610365610479366004613e93565b610fad565b348015610489575f80fd5b50610365610ffe565b6103656104a0366004613f03565b611013565b3480156104b0575f80fd5b506103f4611032565b3480156104c4575f80fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166102fa565b3480156104fa575f80fd5b5061050e610509366004613fc8565b611060565b604051610306929190614010565b348015610527575f80fd5b50610365610536366004613ded565b611157565b348015610546575f80fd5b5061055a610555366004614032565b6111fe565b604051610306919061406f565b348015610572575f80fd5b50600c5461032e906001600160a01b031681565b348015610591575f80fd5b506103656105a0366004614091565b611301565b3480156105b0575f80fd5b5061055a6105bf366004613fc8565b6113bd565b3480156105cf575f80fd5b506105e36105de366004613e66565b611482565b604051610306929190614158565b3480156105fc575f80fd5b50610365611537565b348015610610575f80fd5b5061036561061f366004613ebd565b611569565b34801561062f575f80fd5b5061036561063e366004614179565b6116b7565b34801561064e575f80fd5b5061036561065d3660046141e2565b611749565b34801561066d575f80fd5b506102fa61067c366004613e93565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156106d0575f80fd5b506106d9611809565b6040516103069190614221565b3480156106f1575f80fd5b506106fa61185f565b6040516103069190614263565b348015610712575f80fd5b506103f45f81565b348015610725575f80fd5b5061036561073436600461430e565b611a18565b348015610744575f80fd5b5061036561075336600461437a565b611b1f565b348015610763575f80fd5b5061032e610772366004613fc8565b611ba1565b348015610782575f80fd5b5061055a6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b3480156107ca575f80fd5b506103f47ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa381565b3480156107fd575f80fd5b50610806611be0565b60405161030691906143e4565b34801561081e575f80fd5b506103f47fb90e9995c6170fff8ea03e9ad6919878e483770c237f1a6f330ceaa7112b344a81565b348015610851575f80fd5b5061055a610860366004613ded565b611eda565b348015610870575f80fd5b5061036561087f366004613e93565b611fbe565b34801561088f575f80fd5b5061036561089e36600461430e565b612001565b3480156108ae575f80fd5b506103f47f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b3480156108e1575f80fd5b506108f56108f0366004613ebd565b612103565b604051610306969594939291906144d9565b348015610912575f80fd5b50610365610921366004614536565b61234e565b348015610931575f80fd5b505f5461032e906001600160a01b031681565b34801561094f575f80fd5b5061036561095e3660046145de565b6123e9565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806109f557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b333014610a34576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a4084848484612709565b7f6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c60075f8681526020019081526020015f208484604051610a82929190614624565b9081526020016040518091039020600101604051610aa09190614684565b60405180910390a150505050565b6004546060908067ffffffffffffffff811115610acd57610acd613ed6565b604051908082528060200260200182016040528015610b2957816020015b610b1660405180608001604052805f1515815260200160608152602001606081526020015f81525090565b815260200190600190039081610aeb5790505b5091505f5b81811015610d3a575f60048281548110610b4a57610b4a61470b565b905f5260205f2090600202016040518060400160405290815f8201548152602001600182018054610b7a90614633565b80601f0160208091040260200160405190810160405280929190818152602001828054610ba690614633565b8015610bf15780601f10610bc857610100808354040283529160200191610bf1565b820191905f5260205f20905b815481529060010190602001808311610bd457829003601f168201915b50505050508152505090505f815f015190505f82602001519050604051806080016040528060075f8581526020019081526020015f2083604051610c35919061474f565b90815260408051602092819003830190205460ff16151583525f8681526007835281902090519290910191610c6b90859061474f565b90815260200160405180910390206001018054610c8790614633565b80601f0160208091040260200160405190810160405280929190818152602001828054610cb390614633565b8015610cfe5780601f10610cd557610100808354040283529160200191610cfe565b820191905f5260205f20905b815481529060010190602001808311610ce157829003601f168201915b5050505050815260200182815260200183815250868581518110610d2457610d2461470b565b6020908102919091010152505050600101610b2e565b505090565b333014610d78576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d80612857565b610d8d85858585856128b5565b847f40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee63485858585604051610dc39493929190614783565b60405180910390a25050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610e0b81612940565b610e15838361294a565b50505050565b5f610e2581612940565b6001600160a01b038216610e65576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e8f7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa38361294a565b50610eba7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8361294a565b50600154610ef2907ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3906001600160a01b0316612a16565b50600154610f2a907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a906001600160a01b0316612a16565b50600154604080516001600160a01b03928316815291841660208301527f6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279910160405180910390a150600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6001600160a01b0381163314610fef576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ff98282612a16565b505050565b5f61100881612940565b611010612aba565b50565b61101b612b2c565b61102482612bfc565b61102e8282612c06565b5050565b5f61103b612d07565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f838152600760205260408082209051606091906110819086908690614624565b908152604080519182900360209081018320545f898152600790925291902060ff9091169350906110b59086908690614624565b908152602001604051809103902060010180546110d190614633565b80601f01602080910402602001604051908101604052809291908181526020018280546110fd90614633565b80156111485780601f1061111f57610100808354040283529160200191611148565b820191905f5260205f20905b81548152906001019060200180831161112b57829003601f168201915b50505050509050935093915050565b333014611190576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611198612857565b6111a58585858585612d69565b83836040516111b5929190614624565b6040518091039020857f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce258184846040516111ef9291906147b4565b60405180910390a35050505050565b60607fb90e9995c6170fff8ea03e9ad6919878e483770c237f1a6f330ceaa7112b344a61122a81612940565b611232612857565b600c546001600160a01b031661124b6020870187613ebd565b6001600160a01b03161461128b576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80306001600160a01b031686866040516112a7929190614624565b5f604051808303815f865af19150503d805f81146112e0576040519150601f19603f3d011682016040523d82523d5f602084013e6112e5565b606091505b5091509150816112f757805160208201fd5b9695505050505050565b33301461133a576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611342612857565b611353898989898989898989612ffa565b886001600160a01b0316858560405161136d929190614624565b60405180910390207fa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e36783898c8c6040516113aa94939291906147c7565b60405180910390a3505050505050505050565b606060065f8581526020019081526020015f2060040183836040516113e3929190614624565b908152602001604051809103902080546113fc90614633565b80601f016020809104026020016040519081016040528092919081815260200182805461142890614633565b80156114735780601f1061144a57610100808354040283529160200191611473565b820191905f5260205f20905b81548152906001019060200180831161145657829003601f168201915b505050505090505b9392505050565b5f8181526006602052604090206002810154600390910180546001600160a01b0390921691606091906114b490614633565b80601f01602080910402602001604051908101604052809291908181526020018280546114e090614633565b801561152b5780601f106115025761010080835404028352916020019161152b565b820191905f5260205f20905b81548152906001019060200180831161150e57829003601f168201915b50505050509050915091565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61156181612940565b611010613326565b5f61157381612940565b6001600160a01b0382166115b3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115bd5f8361294a565b506115e87f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8361294a565b505f80546115ff91906001600160a01b0316612a16565b505f54611636907f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a906001600160a01b0316612a16565b505f54604080516001600160a01b03928316815291841660208301527f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f910160405180910390a1505f80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b3330146116f0576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6116f8612857565b6117028282613381565b604080516001600160a01b038416815282151560208201527f9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8910160405180910390a15050565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361177381612940565b61177b612857565b5f5b82811015610e1557368484838181106117985761179861470b565b90506020028101906117aa91906147e9565b90506118006117bf6040830160208401613ebd565b6117cc6080840184614825565b60608501356117de6040870187614825565b6117eb60a0890189614825565b6117fb60e08b0160c08c01614886565b612ffa565b5060010161177d565b6060600280548060200260200160405190810160405280929190818152602001828054801561185557602002820191905f5260205f20905b815481526020019060010190808311611841575b5050505050905090565b6003546060908067ffffffffffffffff81111561187e5761187e613ed6565b6040519080825280602002602001820160405280156118ec57816020015b604080516080810182525f80825260208083018290529282015260608082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90920191018161189c5790505b5091505f5b81811015610d3a575f6003828154811061190d5761190d61470b565b5f918252602080832090910154604080516080810182528285526006808552828620805460ff161515835282860185905260028101546001600160a01b031693830193909352948390529390925260039091018054919350606083019161197390614633565b80601f016020809104026020016040519081016040528092919081815260200182805461199f90614633565b80156119ea5780601f106119c1576101008083540402835291602001916119ea565b820191905f5260205f20905b8154815290600101906020018083116119cd57829003601f168201915b5050505050815250848381518110611a0457611a0461470b565b6020908102919091010152506001016118f1565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa3611a4281612940565b611a4a612857565b5f5b84811015611aab5736868683818110611a6757611a6761470b565b9050602002810190611a79919061489f565b9050611aa26060820135611a906040840184614825565b611a9d6020860186614825565b612d69565b50600101611a4c565b505f5b82811015611b175736848483818110611ac957611ac961470b565b9050602002810190611adb919061489f565b9050611b0e8135611aef6020840184614825565b611afc6040860186614825565b611b096060880188614825565b61346f565b50600101611aae565b505050505050565b333014611b58576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611b60612857565b611b6d85858585856135cc565b847fc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e82604051610dc3911515815260200190565b5f838152600a60205260408082209051611bbe9085908590614624565b908152604051908190036020019020546001600160a01b031690509392505050565b6005546060908067ffffffffffffffff811115611bff57611bff613ed6565b604051908082528060200260200182016040528015611c7a57816020015b611c676040518060e001604052805f151581526020015f6001600160a01b03168152602001606081526020015f815260200160608152602001606081526020015f60ff1681525090565b815260200190600190039081611c1d5790505b5091505f5b81811015610d3a575f60058281548110611c9b57611c9b61470b565b5f9182526020808320909101546001600160a01b0390811680845260088352604093849020845160e081018652815460ff811615158252610100900490931693830193909352600183018054919550919384019190611cf990614633565b80601f0160208091040260200160405190810160405280929190818152602001828054611d2590614633565b8015611d705780601f10611d4757610100808354040283529160200191611d70565b820191905f5260205f20905b815481529060010190602001808311611d5357829003601f168201915b5050505050815260200160028201548152602001600382018054611d9390614633565b80601f0160208091040260200160405190810160405280929190818152602001828054611dbf90614633565b8015611e0a5780601f10611de157610100808354040283529160200191611e0a565b820191905f5260205f20905b815481529060010190602001808311611ded57829003601f168201915b50505050508152602001600482018054611e2390614633565b80601f0160208091040260200160405190810160405280929190818152602001828054611e4f90614633565b8015611e9a5780601f10611e7157610100808354040283529160200191611e9a565b820191905f5260205f20905b815481529060010190602001808311611e7d57829003601f168201915b50505091835250506005919091015460ff166020909101528451859084908110611ec657611ec661470b565b602090810291909101015250600101611c7f565b606060075f8781526020019081526020015f208585604051611efd929190614624565b90815260200160405180910390206003018383604051611f1e929190614624565b90815260200160405180910390208054611f3790614633565b80601f0160208091040260200160405190810160405280929190818152602001828054611f6390614633565b8015611fae5780601f10611f8557610100808354040283529160200191611fae565b820191905f5260205f20905b815481529060010190602001808311611f9157829003601f168201915b5050505050905095945050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611ff781612940565b610e158383612a16565b7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa361202b81612940565b612033612857565b5f5b848110156120a457368686838181106120505761205061470b565b9050602002810190612062919061489f565b905061209b602082013561207c6060840160408501613ebd565b6120896060850185614825565b61209660208701876148d1565b6135cc565b50600101612035565b505f5b82811015611b1757368484838181106120c2576120c261470b565b90506020028101906120d491906148ea565b90506120fa81356120e86020840184614825565b6120f56040860186614825565b6128b5565b506001016120a7565b6001600160a01b038082165f908152600860209081526040808320815160e081018352815460ff8116151582526101009004909516928501929092526001820180549394606094869486948594879485949093928401919061216490614633565b80601f016020809104026020016040519081016040528092919081815260200182805461219090614633565b80156121db5780601f106121b2576101008083540402835291602001916121db565b820191905f5260205f20905b8154815290600101906020018083116121be57829003601f168201915b50505050508152602001600282015481526020016003820180546121fe90614633565b80601f016020809104026020016040519081016040528092919081815260200182805461222a90614633565b80156122755780601f1061224c57610100808354040283529160200191612275565b820191905f5260205f20905b81548152906001019060200180831161225857829003601f168201915b5050505050815260200160048201805461228e90614633565b80601f01602080910402602001604051908101604052809291908181526020018280546122ba90614633565b80156123055780601f106122dc57610100808354040283529160200191612305565b820191905f5260205f20905b8154815290600101906020018083116122e857829003601f168201915b50505091835250506005919091015460ff16602090910152805160808201516060830151604084015160a085015160c090950151939d929c50909a509850919650945092505050565b333014612387576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61238f612857565b61239e8787878787878761346f565b867faea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c58787878787876040516123d89695949392919061491c565b60405180910390a250505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156124335750825b90505f8267ffffffffffffffff16600114801561244f5750303b155b90508115801561245d575080155b15612494576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156124f55784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b038916158061251257506001600160a01b038716155b8061252457506001600160a01b038616155b8061253657506001600160a01b038816155b1561256d576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612575613783565b61257d613783565b61258561378b565b61258f5f8a61294a565b506125ba7ff7a450ef335e1892cb42c8ca72e7242359d7711924b75db5717410da3f614aa38961294a565b506125e57f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8961294a565b506126107f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8a61294a565b5061263b7fb90e9995c6170fff8ea03e9ad6919878e483770c237f1a6f330ceaa7112b344a8861294a565b505f80546001600160a01b03808c167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255600180548b8416908316179055600b80548a8416908316179055600c80549289169290911691909117905583156126fe5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b5f8481526006602052604090205460ff16612758576040517f8e6feba5000000000000000000000000000000000000000000000000000000008152600481018590526024015b60405180910390fd5b5f8290036127955782826040517ec10cfd00000000000000000000000000000000000000000000000000000000815260040161274f9291906147b4565b5f848152600760205260409081902090516127b39085908590614624565b908152602001604051809103902060010180546127cf90614633565b90505f0361280f578383836040517f2b4f9c8600000000000000000000000000000000000000000000000000000000815260040161274f93929190614964565b8060075f8681526020019081526020015f208484604051612831929190614624565b908152604051908190036020019020805491151560ff1990921691909117905550505050565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16156128b3576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f8581526006602052604090205460ff166128ff576040517f8e6feba50000000000000000000000000000000000000000000000000000000081526004810186905260240161274f565b818160065f8881526020019081526020015f206004018686604051612925929190614624565b90815260200160405180910390209182611b179291906149c1565b61101081336137be565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612a0d575f848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556129c33390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506109f5565b5f9150506109f5565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612a0d575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506109f5565b612ac261384a565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480612bc557507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612bb97f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b156128b3576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61102e81612940565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612c60575060408051601f3d908101601f19168201909252612c5d91810190614ab9565b60015b612ca1576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161274f565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612cfd576040517faa1d49a40000000000000000000000000000000000000000000000000000000081526004810182905260240161274f565b610ff983836138a5565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146128b3576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8581526006602052604090205460ff16612db3576040517f8e6feba50000000000000000000000000000000000000000000000000000000081526004810186905260240161274f565b5f839003612df05783836040517ec10cfd00000000000000000000000000000000000000000000000000000000815260040161274f9291906147b4565b5f819003612e2a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f858152600760205260408082209051612e479087908790614624565b90815260200160405180910390206001018054612e6390614633565b90501115612ea75784848484846040517f2b192eab00000000000000000000000000000000000000000000000000000000815260040161274f959493929190614ad0565b600160075f8781526020019081526020015f208585604051612eca929190614624565b9081526040805160209281900383018120805460ff1916941515949094179093555f888152600790925290208391839190612f089088908890614624565b90815260200160405180910390206001019182612f269291906149c1565b50838360075f8881526020019081526020015f208686604051612f4a929190614624565b90815260200160405180910390206002019182612f689291906149c1565b506004604051806040016040528087815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920182905250939094525050835460018181018655948252602091829020845160029092020190815590830151929390929083019150612fe79082614b08565b50505083836040516111b5929190614624565b6001600160a01b03891661303a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8790036130a3576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f53796d626f6c2063616e6e6f7420626520656d70747900000000000000000000604482015260640161274f565b6001600160a01b038981165f9081526008602052604090205461010090041615613104576040517f63f4ee1f0000000000000000000000000000000000000000000000000000000081526001600160a01b038a16600482015260240161274f565b5f6001600160a01b031660098989604051613120929190614624565b908152604051908190036020019020546001600160a01b0316146131745787876040517fb295cac900000000000000000000000000000000000000000000000000000000815260040161274f9291906147b4565b6001600160a01b0389165f818152600860205260409020805460017fffffffffffffffffffffff0000000000000000000000000000000000000000009091166101009093029290921782178155016131cd8587836149c1565b506001600160a01b0389165f908152600860205260409020600281018790556003016131fa888a836149c1565b506001600160a01b0389165f90815260086020526040902060058101805460ff191660ff84161790556004016132318385836149c1565b5088600a5f8881526020019081526020015f208686604051613254929190614624565b90815260200160405180910390205f6101000a8154816001600160a01b0302191690836001600160a01b031602179055508860098989604051613298929190614624565b90815260405190819003602001902080546001600160a01b039283167fffffffffffffffffffffffff000000000000000000000000000000000000000091821617909155600580546001810182555f919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180549b9092169a16999099179098555050505050505050565b61332e612857565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612b0e565b6001600160a01b0382166133c1576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038281165f90815260086020526040902054610100900416613445576040517ec10cfd00000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5a52433230206e6f742072656769737465726564000000000000000000000000604482015260640161274f565b6001600160a01b03919091165f908152600860205260409020805460ff1916911515919091179055565b5f8781526006602052604090205460ff166134b9576040517f8e6feba50000000000000000000000000000000000000000000000000000000081526004810188905260240161274f565b5f8590036134f65785856040517ec10cfd00000000000000000000000000000000000000000000000000000000815260040161274f9291906147b4565b5f878152600760205260409081902090516135149088908890614624565b9081526040519081900360200190205460ff16613563578686866040517f2b4f9c8600000000000000000000000000000000000000000000000000000000815260040161274f93929190614964565b818160075f8a81526020019081526020015f208888604051613586929190614624565b908152602001604051809103902060030186866040516135a7929190614624565b908152602001604051809103902091826135c29291906149c1565b5050505050505050565b5f8581526006602052604090205460ff1680156135e65750805b15613620576040517fa1452cb00000000000000000000000000000000000000000000000000000000081526004810186905260240161274f565b5f8581526006602052604090205460ff1615801561363c575080155b15613676576040517f8e6feba50000000000000000000000000000000000000000000000000000000081526004810186905260240161274f565b5f858152600660205260409020600201546001600160a01b031615801561369d5750468514155b156136d757600380546001810182555f919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b018590555b5f858152600660205260409020805460ff19168215151781556002810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0387161790556003016137338385836149c1565b50801561377357600280546001810182555f919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0185905561377c565b61377c856138fa565b5050505050565b6128b361399f565b61379361399f565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff19169055565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff1661102e576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024810183905260440161274f565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166128b3576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6138ae82613a06565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156138f257610ff98282613aad565b61102e613b1f565b5f5b60025481101561102e57816002828154811061391a5761391a61470b565b905f5260205f20015403613997576002805461393890600190614c01565b815481106139485761394861470b565b905f5260205f200154600282815481106139645761396461470b565b5f91825260209091200155600280548061398057613980614c39565b600190038181905f5260205f20015f905590555050565b6001016138fc565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff166128b3576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806001600160a01b03163b5f03613a54576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161274f565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60605f80846001600160a01b031684604051613ac9919061474f565b5f60405180830381855af49150503d805f8114613b01576040519150601f19603f3d011682016040523d82523d5f602084013e613b06565b606091505b5091509150613b16858383613b57565b95945050505050565b34156128b3576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606082613b6c57613b6782613bcc565b61147b565b8151158015613b8357506001600160a01b0384163b155b15613bc5576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161274f565b508061147b565b805115613bdc5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60208284031215613c1e575f80fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461147b575f80fd5b5f8083601f840112613c5d575f80fd5b50813567ffffffffffffffff811115613c74575f80fd5b602083019150836020828501011115613c8b575f80fd5b9250929050565b80358015158114613ca1575f80fd5b919050565b5f805f8060608587031215613cb9575f80fd5b84359350602085013567ffffffffffffffff811115613cd6575f80fd5b613ce287828801613c4d565b9094509250613cf5905060408601613c92565b905092959194509250565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015613de1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805115158652602081015160806020880152613da46080880182613d00565b905060408201518782036040890152613dbd8282613d00565b60609384015198909301979097525094506020938401939190910190600101613d54565b50929695505050505050565b5f805f805f60608688031215613e01575f80fd5b85359450602086013567ffffffffffffffff811115613e1e575f80fd5b613e2a88828901613c4d565b909550935050604086013567ffffffffffffffff811115613e49575f80fd5b613e5588828901613c4d565b969995985093965092949392505050565b5f60208284031215613e76575f80fd5b5035919050565b80356001600160a01b0381168114613ca1575f80fd5b5f8060408385031215613ea4575f80fd5b82359150613eb460208401613e7d565b90509250929050565b5f60208284031215613ecd575f80fd5b61147b82613e7d565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8060408385031215613f14575f80fd5b613f1d83613e7d565b9150602083013567ffffffffffffffff811115613f38575f80fd5b8301601f81018513613f48575f80fd5b803567ffffffffffffffff811115613f6257613f62613ed6565b604051601f19603f601f19601f8501160116810181811067ffffffffffffffff82111715613f9257613f92613ed6565b604052818152828201602001871015613fa9575f80fd5b816020840160208301375f602083830101528093505050509250929050565b5f805f60408486031215613fda575f80fd5b83359250602084013567ffffffffffffffff811115613ff7575f80fd5b61400386828701613c4d565b9497909650939450505050565b8215158152604060208201525f61402a6040830184613d00565b949350505050565b5f805f8385036040811215614045575f80fd5b6020811215614052575f80fd5b50839250602084013567ffffffffffffffff811115613ff7575f80fd5b602081525f61147b6020830184613d00565b803560ff81168114613ca1575f80fd5b5f805f805f805f805f60c08a8c0312156140a9575f80fd5b6140b28a613e7d565b985060208a013567ffffffffffffffff8111156140cd575f80fd5b6140d98c828d01613c4d565b90995097505060408a0135955060608a013567ffffffffffffffff8111156140ff575f80fd5b61410b8c828d01613c4d565b90965094505060808a013567ffffffffffffffff81111561412a575f80fd5b6141368c828d01613c4d565b9094509250614149905060a08b01614081565b90509295985092959850929598565b6001600160a01b0383168152604060208201525f61402a6040830184613d00565b5f806040838503121561418a575f80fd5b61419383613e7d565b9150613eb460208401613c92565b5f8083601f8401126141b1575f80fd5b50813567ffffffffffffffff8111156141c8575f80fd5b6020830191508360208260051b8501011115613c8b575f80fd5b5f80602083850312156141f3575f80fd5b823567ffffffffffffffff811115614209575f80fd5b614215858286016141a1565b90969095509350505050565b602080825282518282018190525f918401906040840190835b8181101561425857835183526020938401939092019160010161423a565b509095945050505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015613de1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805115158652602081015160208701526001600160a01b03604082015116604087015260608101519050608060608701526142f86080870182613d00565b9550506020938401939190910190600101614289565b5f805f8060408587031215614321575f80fd5b843567ffffffffffffffff811115614337575f80fd5b614343878288016141a1565b909550935050602085013567ffffffffffffffff811115614362575f80fd5b61436e878288016141a1565b95989497509550505050565b5f805f805f6080868803121561438e575f80fd5b8535945061439e60208701613e7d565b9350604086013567ffffffffffffffff8111156143b9575f80fd5b6143c588828901613c4d565b90945092506143d8905060608701613c92565b90509295509295909350565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b82811015613de1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051151586526001600160a01b036020820151166020870152604081015160e0604088015261446d60e0880182613d00565b905060608201516060880152608082015187820360808901526144908282613d00565b91505060a082015187820360a08901526144aa8282613d00565b91505060c082015191506144c360c088018360ff169052565b955050602093840193919091019060010161440a565b861515815260c060208201525f6144f360c0830188613d00565b866040840152828103606084015261450b8187613d00565b9050828103608084015261451f8186613d00565b91505060ff831660a0830152979650505050505050565b5f805f805f805f6080888a03121561454c575f80fd5b87359650602088013567ffffffffffffffff811115614569575f80fd5b6145758a828b01613c4d565b909750955050604088013567ffffffffffffffff811115614594575f80fd5b6145a08a828b01613c4d565b909550935050606088013567ffffffffffffffff8111156145bf575f80fd5b6145cb8a828b01613c4d565b989b979a50959850939692959293505050565b5f805f80608085870312156145f1575f80fd5b6145fa85613e7d565b935061460860208601613e7d565b925061461660408601613e7d565b9150613cf560608601613e7d565b818382375f9101908152919050565b600181811c9082168061464757607f821691505b60208210810361467e577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b602081525f80835461469581614633565b806020860152600182165f81146146b357600181146146cf57614700565b60ff1983166040870152604082151560051b8701019350614700565b865f5260205f205f5b838110156146f7578154888201604001526001909101906020016146d8565b87016040019450505b509195945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81518060208401855e5f93019283525090919050565b5f61147b8284614738565b81835281816020850137505f602082840101525f6020601f19601f840116840101905092915050565b604081525f61479660408301868861475a565b82810360208401526147a981858761475a565b979650505050505050565b602081525f61402a60208301848661475a565b60ff85168152836020820152606060408201525f6112f760608301848661475a565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2183360301811261481b575f80fd5b9190910192915050565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112614858575f80fd5b83018035915067ffffffffffffffff821115614872575f80fd5b602001915036819003821315613c8b575f80fd5b5f60208284031215614896575f80fd5b61147b82614081565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8183360301811261481b575f80fd5b5f602082840312156148e1575f80fd5b61147b82613c92565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa183360301811261481b575f80fd5b606081525f61492f60608301888a61475a565b828103602084015261494281878961475a565b9050828103604084015261495781858761475a565b9998505050505050505050565b838152604060208201525f613b1660408301848661475a565b601f821115610ff957805f5260205f20601f840160051c810160208510156149a25750805b601f840160051c820191505b8181101561377c575f81556001016149ae565b67ffffffffffffffff8311156149d9576149d9613ed6565b6149ed836149e78354614633565b8361497d565b5f601f841160018114614a3d575f8515614a075750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b17835561377c565b5f83815260208120601f198716915b82811015614a6c5786850135825560209485019460019092019101614a4c565b5086821015614aa7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b5f60208284031215614ac9575f80fd5b5051919050565b858152606060208201525f614ae960608301868861475a565b8281036040840152614afc81858761475a565b98975050505050505050565b815167ffffffffffffffff811115614b2257614b22613ed6565b614b3681614b308454614633565b8461497d565b6020601f821160018114614b87575f8315614b515750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b17845561377c565b5f84815260208120601f198516915b82811015614bb65787850151825560209485019460019092019101614b96565b5084821015614bf257868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b818103818111156109f5577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea264697066735822122007d64bc5d097e5ce9b840029ae9959fd969a3c5c11a6c27ac58114ed9441fd3e64736f6c634300081a0033",
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
