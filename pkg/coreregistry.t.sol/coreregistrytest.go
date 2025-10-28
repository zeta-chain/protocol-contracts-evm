// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coreregistry

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

// StdInvariantFuzzArtifactSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzArtifactSelector struct {
	Artifact  string
	Selectors [][4]byte
}

// StdInvariantFuzzInterface is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzInterface struct {
	Addr      common.Address
	Artifacts []string
}

// StdInvariantFuzzSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzSelector struct {
	Addr      common.Address
	Selectors [][4]byte
}

// CoreRegistryTestMetaData contains all meta data concerning the CoreRegistryTest contract.
var CoreRegistryTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testActivateAlreadyActiveChain\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testActivateChain\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testActivateChainUnauthorized\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testChangeAdmin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testChangeAdminUnauthorized\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testChangeAdminWithZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testChangeRegistryManager\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testChangeRegistryManagerUnauthorized\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testChangeRegistryManagerWithZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDeactivateChain\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testDeactivateNonActiveChain\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testGetAllChains\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testGetAllContracts\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testGetAllZRC20Tokens\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testInitialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testInitializeWithZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testPause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testPauseUnauthorized\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRegisterContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRegisterContractForNonActiveChain\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRegisterContractTwice\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRegisterContractWithEmptyAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRegisterContractWithEmptyType\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRegisterZRC20Token\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRegisterZRC20TokenTwice\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRegisterZRC20TokenWithDuplicateSymbol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRegisterZRC20TokenWithEmptyAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRegisterZRC20TokenWithEmptyOriginAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testRegisterZRC20TokenWithEmptySymbol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetContractActive\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetStatusForNonExistenContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetZRC20TokenActive\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testSetZRC20TokenActiveWithZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpdateChainMetadata\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpdateConfigForNonExistentContract\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpdateContractConfig\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpdateMetadataForNonActiveChain\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpdateMetadataUnauthorized\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpdateNonExistentZRC20Token\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWhenPaused\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"oldAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainMetadataUpdated\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainStatusChanged\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newStatus\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractConfigurationUpdated\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractStatusChanged\",\"inputs\":[{\"name\":\"addressBytes\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistryManagerChanged\",\"inputs\":[{\"name\":\"oldRegistryManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRegistryManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZRC20TokenRegistered\",\"inputs\":[{\"name\":\"originAddress\",\"type\":\"bytes\",\"indexed\":true,\"internalType\":\"bytes\"},{\"name\":\"address_\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"originChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ZRC20TokenUpdated\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"active\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ChainActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainNonActive\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ContractAlreadyRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"addressBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ContractNotFound\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"contractType\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidContractType\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20AlreadyRegistered\",\"inputs\":[{\"name\":\"address_\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZRC20SymbolAlreadyInUse\",\"inputs\":[{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602b575f80fd5b50620140a5806200003b5f395ff3fe608060405234801562000010575f80fd5b50600436106200038b575f3560e01c80638c18b40f11620001e3578063b5508aa91162000113578063d97aa89911620000ab578063f7146feb1162000083578063f7146feb14620005e0578063f81b656b14620005ea578063fa7626d414620005f4578063faab466a1462000602575f80fd5b8063d97aa89914620005c2578063e20c9f7114620005cc578063f48974af14620005d6575f80fd5b8063cff8f11111620000eb578063cff8f111146200059a578063d312ad7d14620005a4578063d52fe06614620005ae578063d668385614620005b8575f80fd5b8063b5508aa9146200056b578063ba414fa61462000575578063bc056f7f1462000590575f80fd5b80639dc744021162000187578063aa5ec565116200015f578063aa5ec5651462000543578063b0464fdc146200054d578063b17255801462000557578063b39377d81462000561575f80fd5b80639dc7440214620005255780639fe48db2146200052f578063a5ae07e21462000539575f80fd5b8063933e897011620001bb578063933e89701462000507578063993831b614620005115780639bf35597146200051b575f80fd5b80638c18b40f14620004da5780638c500b5c14620004e4578063916a17c614620004ee575f80fd5b80633f7286f411620002bf57806366d9a9a0116200026357806371748c09116200023b57806371748c0914620004a35780637d706cef14620004ad57806385226c8114620004b75780638932825814620004d0575f80fd5b806366d9a9a01462000476578063674270a0146200048f57806370cd41381462000499575f80fd5b8063599a9da31162000297578063599a9da31462000458578063599d117314620004625780635d737a7e146200046c575f80fd5b80633f7286f4146200043a57806340a3b50b14620004445780634a3b5640146200044e575f80fd5b80632013038b11620003335780632ade3880116200030b5780632ade3880146200040357806336328450146200041c57806339c1f2a414620004265780633e5e3c231462000430575f80fd5b80632013038b14620003e557806324196f3214620003ef57806324ffc31714620003f9575f80fd5b8063164c5b021162000367578063164c5b0214620003af578063173d959014620003b95780631ed7831c14620003c3575f80fd5b80624e57bd146200038f5780630a9254e4146200039b578063151e9be414620003a5575b5f80fd5b620003996200060c565b005b620003996200084c565b620003996200106f565b62000399620012e0565b620003996200146a565b620003cd620015e6565b604051620003dc91906200ec86565b60405180910390f35b6200039962001648565b620003996200186c565b6200039962001f26565b6200040d620022ca565b604051620003dc91906200ed01565b6200039962002412565b6200039962002697565b620003cd62002a0d565b620003cd62002a6d565b6200039962002acd565b6200039962002c87565b6200039962002f9a565b620003996200321f565b6200039962003383565b62000480620036f6565b604051620003dc91906200ee6a565b620003996200387b565b6200039962003d78565b62000399620040ec565b620003996200447b565b620004c1620046a4565b604051620003dc91906200ef0c565b6200039962004779565b6200039962005001565b62000399620053af565b620004f8620057b9565b604051620003dc91906200ef85565b62000399620058b3565b6200039962005a09565b6200039962005cce565b6200039962005f9b565b6200039962006614565b620003996200680b565b620003996200695c565b620004f862006b53565b6200039962006c4d565b62000399620072c7565b620004c162007696565b6200057f6200776b565b6040519015158152602001620003dc565b620003996200783f565b6200039962007be4565b6200039962007ee7565b62000399620080fe565b620003996200878e565b6200039962008951565b620003cd62008c43565b6200039962008ca3565b6200039962008f29565b62000399620099e0565b601f546200057f9060ff1681565b6200039962009cf9565b604080518082018252600981527f636f6e6e6563746f72000000000000000000000000000000000000000000000060208083019190915291516daa550000000000000000000000009281019290925260019161aa5591905f9060340160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620006cd575f80fd5b505af1158015620006e0573d5f803e3d5ffd5b505060408051602480820189905282518083039091018152604490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8e6feba50000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506200079991906004016200f01e565b5f604051808303815f87803b158015620007b1575f80fd5b505af1158015620007c4573d5f803e3d5ffd5b50506020546040517f631d62e40000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063631d62e4915062000817908790869086906004016200f032565b5f604051808303815f87803b1580156200082f575f80fd5b505af115801562000842573d5f803e3d5ffd5b5050505050505050565b602480547fffffffffffffffffffffffff000000000000000000000000000000000000000090811661abcd1790915560258054821661123417905560268054909116615678179055604051620008a2906200ebab565b604051809103905ff080158015620008bc573d5f803e3d5ffd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252601081527f436f726552656769737472792e736f6c000000000000000000000000000000006020820152602480546025549351908616918101919091529190931660448201526064810191909152620009b2919060840160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b000000000000000000000000000000000000000000000000000000001790526200a467565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0393841681029190911791829055602080549190920483167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116178155602154604080517f2722feee00000000000000000000000000000000000000000000000000000000815290519190931692632722feee9260048083019391928290030181865afa15801562000a78573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062000a9e91906200f087565b602780547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039290921691821790556040517f06447d560000000000000000000000000000000000000000000000000000000081526004810191909152737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b15801562000b3a575f80fd5b505af115801562000b4d573d5f803e3d5ffd5b505050505f805f60405162000b62906200ebb9565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103905ff08015801562000b9c573d5f803e3d5ffd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392831690811790915560215460405160129360019384935f939192169062000bf3906200ebc7565b62000c04969594939291906200f0a5565b604051809103905ff08015801562000c1e573d5f803e3d5ffd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081179091556023546040517fee2815ba0000000000000000000000000000000000000000000000000000000081526001600482015260248101929092529091169063ee2815ba906044015f604051808303815f87803b15801562000cb3575f80fd5b505af115801562000cc6573d5f803e3d5ffd5b50506023546040517fa7cb050700000000000000000000000000000000000000000000000000000000815260016004820181905260248201526001600160a01b03909116925063a7cb050791506044015f604051808303815f87803b15801562000d2e575f80fd5b505af115801562000d41573d5f803e3d5ffd5b50506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152633b9aca006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d91506044015f604051808303815f87803b15801562000dbf575f80fd5b505af115801562000dd2573d5f803e3d5ffd5b50506022546025546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e1006024820152911692506347e7ef2491506044016020604051808303815f875af115801562000e45573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062000e6b91906200f1a9565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562000ec7575f80fd5b505af115801562000eda573d5f803e3d5ffd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d5691506024015f604051808303815f87803b15801562000f4e575f80fd5b505af115801562000f61573d5f803e3d5ffd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e10060248201529116925063095ea7b391506044016020604051808303815f875af115801562000fd4573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062000ffa91906200f1a9565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562001056575f80fd5b505af115801562001069573d5f803e3d5ffd5b50505050565b60265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620010c6575f80fd5b505af1158015620010d9573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562001138575f80fd5b505af11580156200114b573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156200119c575f80fd5b505af1158015620011af573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b1580156200120a575f80fd5b505af11580156200121d573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156200127c575f80fd5b505af11580156200128f573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562001056575f80fd5b60255460405163ca669fa760e01b81526001600160a01b0390911660048201525f90737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562001339575f80fd5b505af11580156200134c573d5f803e3d5ffd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015b5f604051808303815f87803b158015620013d4575f80fd5b505af1158015620013e7573d5f803e3d5ffd5b50506020546040517f9060bda90000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301525f60248301529091169250639060bda991506044015b5f604051808303815f87803b15801562001450575f80fd5b505af115801562001463573d5f803e3d5ffd5b5050505050565b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f190181529082905260265463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620014ee575f80fd5b505af115801562001501573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562001560575f80fd5b505af115801562001573573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb969350620015b3928792169086906001906004016200f1c5565b5f604051808303815f87803b158015620015cb575f80fd5b505af1158015620015de573d5f803e3d5ffd5b505050505050565b606060168054806020026020016040519081016040528092919081815260200182805480156200163e57602002820191905f5260205f20905b81546001600160a01b031681526001909101906020018083116200161f575b5050505050905090565b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f190181528282018252600983527f636f6e6e6563746f720000000000000000000000000000000000000000000000602080850191909152825190810183525f8152602554925163ca669fa760e01b81526001600160a01b03909316600484015290935090737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200170b575f80fd5b505af11580156200171e573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb9693506200175e928992169088906001906004016200f1c5565b5f604051808303815f87803b15801562001776575f80fd5b505af115801562001789573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015620017e4575f80fd5b505af1158015620017f7573d5f803e3d5ffd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e0915060240162000799565b604080518082018252600381527f544b4e000000000000000000000000000000000000000000000000000000000060208083019190915291516d13570000000000000000000000009281019290925261dddd916001905f9060340160408051808303601f190181528282018252600583527f45524332300000000000000000000000000000000000000000000000000000006020840152602554915163ca669fa760e01b81526001600160a01b0390921660048301529250601290737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200195e575f80fd5b505af115801562001971573d5f803e3d5ffd5b50506020546040517f6e9e2d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250636e9e2d3f9150620019ca908990899089908990899089906004016200f201565b5f604051808303815f87803b158015620019e2575f80fd5b505af1158015620019f5573d5f803e3d5ffd5b50506020546040517fe9d6c5ba0000000000000000000000000000000000000000000000000000000081526001600160a01b038a811660048301525f9450909116915063e9d6c5ba906024015f60405180830381865afa15801562001a5c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405262001a8591908101906200f3ce565b5050505050905062001a97816200a489565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562001aee575f80fd5b505af115801562001b01573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562001b65575f80fd5b505af115801562001b78573d5f803e3d5ffd5b5050604080516001600160a01b038b1681525f60208201527f9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8935001905060405180910390a16020546040517f9060bda90000000000000000000000000000000000000000000000000000000081526001600160a01b0389811660048301525f602483015290911690639060bda9906044015f604051808303815f87803b15801562001c22575f80fd5b505af115801562001c35573d5f803e3d5ffd5b50506020546040517fe9d6c5ba0000000000000000000000000000000000000000000000000000000081526001600160a01b038b81166004830152909116925063e9d6c5ba91506024015f60405180830381865afa15801562001c9a573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405262001cc391908101906200f3ce565b5093945062001cd993508492506200a501915050565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562001d30575f80fd5b505af115801562001d43573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562001da7575f80fd5b505af115801562001dba573d5f803e3d5ffd5b5050604080516001600160a01b038b168152600160208201527f9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8935001905060405180910390a16020546040517f9060bda90000000000000000000000000000000000000000000000000000000081526001600160a01b0389811660048301526001602483015290911690639060bda9906044015f604051808303815f87803b15801562001e66575f80fd5b505af115801562001e79573d5f803e3d5ffd5b50506020546040517fe9d6c5ba0000000000000000000000000000000000000000000000000000000081526001600160a01b038b81166004830152909116925063e9d6c5ba91506024015f60405180830381865afa15801562001ede573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405262001f0791908101906200f3ce565b5093945062001f1d93508492506200a489915050565b50505050505050565b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562001faa575f80fd5b505af115801562001fbd573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb96935062001ffd928792169086906001906004016200f1c5565b5f604051808303815f87803b15801562002015575f80fd5b505af115801562002028573d5f803e3d5ffd5b505050505f60205f9054906101000a90046001600160a01b03166001600160a01b03166394cc86836040518163ffffffff1660e01b81526004015f60405180830381865afa1580156200207d573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052620020a691908101906200f4bb565b9050620020b6815160026200a554565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200210d575f80fd5b505af115801562002120573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562002184575f80fd5b505af115801562002197573d5f803e3d5ffd5b50506040515f81528592507fc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e915060200160405180910390a260205460225460405163547965cb60e11b81526001600160a01b039283169263a8f2cb96926200220d9288929091169087905f906004016200f1c5565b5f604051808303815f87803b15801562002225575f80fd5b505af115801562002238573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b03166394cc86836040518163ffffffff1660e01b81526004015f60405180830381865afa1580156200228c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052620022b591908101906200f4bb565b9050620022c5815160016200a554565b505050565b6060601e805480602002602001604051908101604052809291908181526020015f905b8282101562002409575f84815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b82821015620023f1578382905f5260205f200180546200235f906200f551565b80601f01602080910402602001604051908101604052809291908181526020018280546200238d906200f551565b8015620023dc5780601f10620023b257610100808354040283529160200191620023dc565b820191905f5260205f20905b815481529060010190602001808311620023be57829003601f168201915b5050505050815260200190600101906200233f565b505050508152505081526020019060010190620022ed565b50505050905090565b60255460405163ca669fa760e01b81526001600160a01b03909116600482015261777790737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200246d575f80fd5b505af115801562002480573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b158015620024df575f80fd5b505af1158015620024f2573d5f803e3d5ffd5b50506020546040517f3500c24b0000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301529091169250633500c24b91506024015f604051808303815f87803b15801562002554575f80fd5b505af115801562002567573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015620025c2575f80fd5b505af1158015620025d5573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562002634575f80fd5b505af115801562002647573d5f803e3d5ffd5b50506020546040517f3500c24b0000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301529091169250633500c24b915060240162001438565b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f190181528282018252600b83527f6e6f6e6578697374656e7400000000000000000000000000000000000000000060208085019190915282518084018452600881527f6761734c696d6974000000000000000000000000000000000000000000000000818301528351620493e09281019290925291945090915f910160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620027a0575f80fd5b505af1158015620027b3573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb969350620027f3928a92169089906001906004016200f1c5565b5f604051808303815f87803b1580156200280b575f80fd5b505af11580156200281e573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801562002879575f80fd5b505af11580156200288c573d5f803e3d5ffd5b5050604051737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507f2b4f9c860000000000000000000000000000000000000000000000000000000090620028e390899088906024016200f59e565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b909216825262002957916004016200f01e565b5f604051808303815f87803b1580156200296f575f80fd5b505af115801562002982573d5f803e3d5ffd5b50506020546040517ff354b31f0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063f354b31f9150620029d79088908790879087906004016200f5b8565b5f604051808303815f87803b158015620029ef575f80fd5b505af115801562002a02573d5f803e3d5ffd5b505050505050505050565b606060188054806020026020016040519081016040528092919081815260200182805480156200163e57602002820191905f5260205f209081546001600160a01b031681526001909101906020018083116200161f575050505050905090565b606060178054806020026020016040519081016040528092919081815260200182805480156200163e57602002820191905f5260205f209081546001600160a01b031681526001909101906020018083116200161f575050505050905090565b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562002b51575f80fd5b505af115801562002b64573d5f803e3d5ffd5b505060408051602480820187905282518083039091018152604490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8e6feba50000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925062002c1d91906004016200f01e565b5f604051808303815f87803b15801562002c35575f80fd5b505af115801562002c48573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb969350620015b3928792169086905f906004016200f1c5565b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f190181528282018252600b83527f6e6f6e6578697374656e740000000000000000000000000000000000000000006020840152602554915163ca669fa760e01b81526001600160a01b0390921660048301529250737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562002d3a575f80fd5b505af115801562002d4d573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb96935062002d8d928892169087906001906004016200f1c5565b5f604051808303815f87803b15801562002da5575f80fd5b505af115801562002db8573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801562002e13575f80fd5b505af115801562002e26573d5f803e3d5ffd5b5050604051737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507f2b4f9c86000000000000000000000000000000000000000000000000000000009062002e7d90879086906024016200f59e565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b909216825262002ef1916004016200f01e565b5f604051808303815f87803b15801562002f09575f80fd5b505af115801562002f1c573d5f803e3d5ffd5b50506020546040517f10d29b9e0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506310d29b9e915062002f6f90869085905f906004016200f5fc565b5f604051808303815f87803b15801562002f87575f80fd5b505af115801562001f1d573d5f803e3d5ffd5b60255460405163ca669fa760e01b81526001600160a01b03909116600482015261999990737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562002ff5575f80fd5b505af115801562003008573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562003067575f80fd5b505af11580156200307a573d5f803e3d5ffd5b50506020546040517f8f2839700000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301529091169250638f28397091506024015f604051808303815f87803b158015620030dc575f80fd5b505af1158015620030ef573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b1580156200314a575f80fd5b505af11580156200315d573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b158015620031bc575f80fd5b505af1158015620031cf573d5f803e3d5ffd5b50506020546040517f8f2839700000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301529091169250638f283970915060240162001438565b6024805460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791015f604051808303815f87803b15801562003275575f80fd5b505af115801562003288573d5f803e3d5ffd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b1580156200330f575f80fd5b505af115801562003322573d5f803e3d5ffd5b50506020546040517f3500c24b0000000000000000000000000000000000000000000000000000000081525f60048201526001600160a01b039091169250633500c24b91506024015b5f604051808303815f87803b15801562001056575f80fd5b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f19018152908290526024805463ca669fa760e01b84526001600160a01b03166004840152909250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791015f604051808303815f87803b15801562003407575f80fd5b505af11580156200341a573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156200346b575f80fd5b505af11580156200347e573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015620034d9575f80fd5b505af1158015620034ec573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156200354b575f80fd5b505af11580156200355e573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb9693506200359e928792169086906001906004016200f1c5565b5f604051808303815f87803b158015620035b6575f80fd5b505af1158015620035c9573d5f803e3d5ffd5b50506024805460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063ca669fa79250015f604051808303815f87803b15801562003623575f80fd5b505af115801562003636573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562003687575f80fd5b505af11580156200369a573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015b5f604051808303815f87803b15801562001560575f80fd5b6060601b805480602002602001604051908101604052809291908181526020015f905b8282101562002409578382905f5260205f2090600202016040518060400160405290815f820180546200374c906200f551565b80601f01602080910402602001604051908101604052809291908181526020018280546200377a906200f551565b8015620037c95780601f106200379f57610100808354040283529160200191620037c9565b820191905f5260205f20905b815481529060010190602001808311620037ab57829003601f168201915b50505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156200386257602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116200380e5790505b5050505050815250508152602001906001019062003719565b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f190181528282018252600983527f636f6e6e6563746f72000000000000000000000000000000000000000000000060208085019190915291516daa5500000000000000000000000092810192909252925061aa5591905f9060340160408051808303601f190181528282018252600883527f6761734c696d69740000000000000000000000000000000000000000000000006020808501919091528251620493e0918101919091529093505f910160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620039b6575f80fd5b505af1158015620039c9573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb96935062003a09928c9216908b906001906004016200f1c5565b5f604051808303815f87803b15801562003a21575f80fd5b505af115801562003a34573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801562003a8f575f80fd5b505af115801562003aa2573d5f803e3d5ffd5b50506020546040517f631d62e40000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063631d62e4915062003af5908a90889088906004016200f032565b5f604051808303815f87803b15801562003b0d575f80fd5b505af115801562003b20573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801562003b7b575f80fd5b505af115801562003b8e573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562003bf2575f80fd5b505af115801562003c05573d5f803e3d5ffd5b50505050867faea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c585848460405162003c3f939291906200f628565b60405180910390a26020546040517ff354b31f0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063f354b31f9062003c98908a908890879087906004016200f5b8565b5f604051808303815f87803b15801562003cb0575f80fd5b505af115801562003cc3573d5f803e3d5ffd5b50506020546040517fd3523ea20000000000000000000000000000000000000000000000000000000081525f93506001600160a01b03909116915063d3523ea29062003d18908b90899088906004016200f032565b5f60405180830381865afa15801562003d33573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405262003d5c91908101906200f666565b905062000842818051906020012083805190602001206200a5d2565b604080518082018252600381527f544b4e000000000000000000000000000000000000000000000000000000000060208083019190915291516d13570000000000000000000000009281019290925261dddd916001905f9060340160408051808303601f190181528282018252600583527f45524332300000000000000000000000000000000000000000000000000000006020840152602554915163ca669fa760e01b81526001600160a01b0390921660048301529250601290737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562003e6a575f80fd5b505af115801562003e7d573d5f803e3d5ffd5b50506020546040517f6e9e2d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250636e9e2d3f915062003ed6908990899089908990899089906004016200f201565b5f604051808303815f87803b15801562003eee575f80fd5b505af115801562003f01573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801562003f5c575f80fd5b505af115801562003f6f573d5f803e3d5ffd5b5050604080516001600160a01b038a1660248083019190915282518083039091018152604490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f63f4ee1f0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506200403391906004016200f01e565b5f604051808303815f87803b1580156200404b575f80fd5b505af11580156200405e573d5f803e3d5ffd5b50506020546040517f6e9e2d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250636e9e2d3f9150620040b590899088908890889088906004016200f69c565b5f604051808303815f87803b158015620040cd575f80fd5b505af1158015620040e0573d5f803e3d5ffd5b50505050505050505050565b5f6200413c6040518060400160405280601081526020017f436f726552656769737472792e736f6c0000000000000000000000000000000081525060405180602001604052805f8152506200a467565b6040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d0000000000000000000000000000000000000000000000000000000060048201529091508190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015f604051808303815f87803b158015620041c4575f80fd5b505af1158015620041d7573d5f803e3d5ffd5b50506025546021546040517fc0c53b8b0000000000000000000000000000000000000000000000000000000081525f60048201526001600160a01b0392831660248201529082166044820152908416925063c0c53b8b91506064015f604051808303815f87803b1580156200424a575f80fd5b505af11580156200425d573d5f803e3d5ffd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015620042e4575f80fd5b505af1158015620042f7573d5f803e3d5ffd5b5050602480546021546040517fc0c53b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b0392831660048201525f93810193909352811660448301528416925063c0c53b8b91506064015f604051808303815f87803b1580156200436b575f80fd5b505af11580156200437e573d5f803e3d5ffd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b15801562004405575f80fd5b505af115801562004418573d5f803e3d5ffd5b5050602480546025546040517fc0c53b8b0000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820152908216928101929092525f60448301528416925063c0c53b8b9150606401620015b3565b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620044ff575f80fd5b505af115801562004512573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb96935062004552928792169086906001906004016200f1c5565b5f604051808303815f87803b1580156200456a575f80fd5b505af11580156200457d573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015620045d8575f80fd5b505af1158015620045eb573d5f803e3d5ffd5b505060408051602480820187905282518083039091018152604490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa1452cb00000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250620036de91906004016200f01e565b6060601a805480602002602001604051908101604052809291908181526020015f905b8282101562002409578382905f5260205f20018054620046e7906200f551565b80601f016020809104026020016040519081016040528092919081815260200182805462004715906200f551565b8015620047645780601f106200473a5761010080835404028352916020019162004764565b820191905f5260205f20905b8154815290600101906020018083116200474657829003601f168201915b505050505081526020019060010190620046c7565b60208054604080517fa217fddf0000000000000000000000000000000000000000000000000000000081529051619999936200489d936001600160a01b0316926391d1485492849263a217fddf92600480820193918290030181865afa158015620047e6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200480c91906200f725565b6024805460405160e085901b7fffffffff0000000000000000000000000000000000000000000000000000000016815260048101939093526001600160a01b0316908201526044015b602060405180830381865afa15801562004871573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200489791906200f1a9565b6200a489565b60208054604080517fe63ab1e900000000000000000000000000000000000000000000000000000000815290516200490a936001600160a01b03909316926391d1485492849263e63ab1e9926004808401939192918290030181865afa158015620047e6573d5f803e3d5ffd5b60208054604080517fa217fddf000000000000000000000000000000000000000000000000000000008152905162004a2c936001600160a01b03909316926391d1485492849263a217fddf926004808401939192918290030181865afa15801562004977573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200499d91906200f725565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815260048101919091526001600160a01b03851660248201526044015b602060405180830381865afa15801562004a00573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004a2691906200f1a9565b6200a501565b60208054604080517fe63ab1e9000000000000000000000000000000000000000000000000000000008152905162004a99936001600160a01b03909316926391d1485492849263e63ab1e9926004808401939192918290030181865afa15801562004977573d5f803e3d5ffd5b60208054604080517ff851a440000000000000000000000000000000000000000000000000000000008152905162004b34936001600160a01b039093169263f851a44092600480820193918290030181865afa15801562004afc573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004b2291906200f087565b6024546001600160a01b03166200a62b565b6024805460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791015f604051808303815f87803b15801562004b8a575f80fd5b505af115801562004b9d573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562004c01575f80fd5b505af115801562004c14573d5f803e3d5ffd5b5050602454604080516001600160a01b03928316815291851660208301527f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f935001905060405180910390a16020546040517f8f2839700000000000000000000000000000000000000000000000000000000081526001600160a01b03838116600483015290911690638f283970906024015f604051808303815f87803b15801562004cbe575f80fd5b505af115801562004cd1573d5f803e3d5ffd5b505060208054604080517fa217fddf000000000000000000000000000000000000000000000000000000008152905162004db595506001600160a01b0390921693506391d1485492849263a217fddf9260048082019392918290030181865afa15801562004d41573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004d6791906200f725565b6024805460405160e085901b7fffffffff0000000000000000000000000000000000000000000000000000000016815260048101939093526001600160a01b031690820152604401620049e4565b60208054604080517fe63ab1e9000000000000000000000000000000000000000000000000000000008152905162004e22936001600160a01b03909316926391d1485492849263e63ab1e9926004808401939192918290030181865afa15801562004d41573d5f803e3d5ffd5b60208054604080517fa217fddf000000000000000000000000000000000000000000000000000000008152905162004f01936001600160a01b03909316926391d1485492849263a217fddf926004808401939192918290030181865afa15801562004e8f573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004eb591906200f725565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815260048101919091526001600160a01b038516602482015260440162004855565b60208054604080517fe63ab1e9000000000000000000000000000000000000000000000000000000008152905162004f6e936001600160a01b03909316926391d1485492849263e63ab1e9926004808401939192918290030181865afa15801562004e8f573d5f803e3d5ffd5b60208054604080517ff851a440000000000000000000000000000000000000000000000000000000008152905162004ffe936001600160a01b039093169263f851a44092600480820193918290030181865afa15801562004fd1573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062004ff791906200f087565b826200a62b565b50565b604080518082018252600381527f544b4e000000000000000000000000000000000000000000000000000000000060208083019190915291516d13570000000000000000000000009281019290925261dddd9161cccc91906001905f9060340160408051601f19818403018152908290526daaaa000000000000000000000000602083015291505f9060340160408051808303601f190181528282018252600583527f45524332300000000000000000000000000000000000000000000000000000006020840152602554915163ca669fa760e01b81526001600160a01b0390921660048301529250601290737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562005124575f80fd5b505af115801562005137573d5f803e3d5ffd5b50506020546040517f6e9e2d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250636e9e2d3f915062005190908b908a908a908a90899089906004016200f201565b5f604051808303815f87803b158015620051a8575f80fd5b505af1158015620051bb573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801562005216575f80fd5b505af115801562005229573d5f803e3d5ffd5b5050604051737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507fb295cac900000000000000000000000000000000000000000000000000000000906200527e908a906024016200f01e565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b9092168252620052f2916004016200f01e565b5f604051808303815f87803b1580156200530a575f80fd5b505af11580156200531d573d5f803e3d5ffd5b50506020546040517f6e9e2d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250636e9e2d3f915062005376908a908a908a908990899089906004016200f201565b5f604051808303815f87803b1580156200538e575f80fd5b505af1158015620053a1573d5f803e3d5ffd5b505050505050505050505050565b604080518082018252600381527f544b4e000000000000000000000000000000000000000000000000000000000060208083019190915291516d13570000000000000000000000009281019290925261dddd916001905f9060340160408051808303601f190181528282018252600583527f45524332300000000000000000000000000000000000000000000000000000006020840152602554915163ca669fa760e01b81526001600160a01b0390921660048301529250601290737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620054a1575f80fd5b505af1158015620054b4573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562005518575f80fd5b505af11580156200552b573d5f803e3d5ffd5b50505050856001600160a01b0316836040516200554991906200f754565b60405180910390207fa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e36783878960405162005586939291906200f761565b60405180910390a36020546040517f6e9e2d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636e9e2d3f90620055e3908990899089908990899089906004016200f201565b5f604051808303815f87803b158015620055fb575f80fd5b505af11580156200560e573d5f803e3d5ffd5b50506020546040517fe9d6c5ba0000000000000000000000000000000000000000000000000000000081526001600160a01b038a811660048301525f94508493508392839283928392169063e9d6c5ba906024015f60405180830381865afa1580156200567d573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052620056a691908101906200f3ce565b955095509550955095509550620056bd866200a489565b620056c9858c6200a68d565b620056d5848b6200a554565b620056ef83805190602001208a805190602001206200a5d2565b620056fb82896200a68d565b6200570d8160ff168860ff166200a554565b6020546040517faa808c060000000000000000000000000000000000000000000000000000000081525f916001600160a01b03169063aa808c06906200575a908e908e906004016200f59e565b602060405180830381865afa15801562005776573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200579c91906200f087565b9050620057aa818e6200a62b565b50505050505050505050505050565b6060601d805480602002602001604051908101604052809291908181526020015f905b8282101562002409575f8481526020908190206040805180820182526002860290920180546001600160a01b031683526001810180548351818702810187019094528084529394919385830193928301828280156200589a57602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411620058465790505b50505050508152505081526020019060010190620057dc565b60255460405163ca669fa760e01b81526001600160a01b03909116600482015261dead90737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200590e575f80fd5b505af115801562005921573d5f803e3d5ffd5b505060408051602060248201819052601460448301527f5a52433230206e6f7420726567697374657265640000000000000000000000006064808401919091528351808403909101815260849092018352810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ec10cfd0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb39250620013bc91906004016200f01e565b60208054604080517fa217fddf000000000000000000000000000000000000000000000000000000008152905162005a76936001600160a01b03909316926391d1485492849263a217fddf926004808401939192918290030181865afa158015620047e6573d5f803e3d5ffd5b60208054604080517fbd8fde1c000000000000000000000000000000000000000000000000000000008152905162005b57936001600160a01b03909316926391d1485492849263bd8fde1c926004808401939192918290030181865afa15801562005ae3573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005b0991906200f725565b60255460405160e084901b7fffffffff0000000000000000000000000000000000000000000000000000000016815260048101929092526001600160a01b0316602482015260440162004855565b60208054604080517fe63ab1e9000000000000000000000000000000000000000000000000000000008152905162005bc4936001600160a01b03909316926391d1485492849263e63ab1e9926004808401939192918290030181865afa15801562005ae3573d5f803e3d5ffd5b60208054604080517fe63ab1e9000000000000000000000000000000000000000000000000000000008152905162005c31936001600160a01b03909316926391d1485492849263e63ab1e9926004808401939192918290030181865afa158015620047e6573d5f803e3d5ffd5b60208054604080517fcc5ad8b6000000000000000000000000000000000000000000000000000000008152905162005ccc936001600160a01b039093169263cc5ad8b692600480820193918290030181865afa15801562005c94573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062005cba91906200f087565b6021546001600160a01b03166200a62b565b565b60208054604080517f5c975abb000000000000000000000000000000000000000000000000000000008152905162005d31936001600160a01b0390931692635c975abb92600480820193918290030181865afa15801562004a00573d5f803e3d5ffd5b6024805460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791015f604051808303815f87803b15801562005d87575f80fd5b505af115801562005d9a573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b0316638456cb596040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562005deb575f80fd5b505af115801562005dfe573d5f803e3d5ffd5b505060208054604080517f5c975abb000000000000000000000000000000000000000000000000000000008152905162005e6695506001600160a01b039092169350635c975abb9260048083019391928290030181865afa15801562004871573d5f803e3d5ffd5b6024805460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791015f604051808303815f87803b15801562005ebc575f80fd5b505af115801562005ecf573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b0316633f4ba83a6040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562005f20575f80fd5b505af115801562005f33573d5f803e3d5ffd5b505060208054604080517f5c975abb000000000000000000000000000000000000000000000000000000008152905162005ccc95506001600160a01b039092169350635c975abb9260048083019391928290030181865afa15801562004a00573d5f803e3d5ffd5b6040516d987600000000000000000000000060208201526001906002905f9060340160408051601f19818403018152908290526d8765000000000000000000000000602083015291505f9060340160408051808303601f19018152908290526025547f06447d560000000000000000000000000000000000000000000000000000000083526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b15801562006067575f80fd5b505af11580156200607a573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb969350620060ba928992169087906001906004016200f1c5565b5f604051808303815f87803b158015620060d2575f80fd5b505af1158015620060e5573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb96935062006125928892169086906001906004016200f1c5565b5f604051808303815f87803b1580156200613d575f80fd5b505af115801562006150573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015620061af575f80fd5b505af1158015620061c2573d5f803e3d5ffd5b505050505f60205f9054906101000a90046001600160a01b03166001600160a01b0316639ca220dd6040518163ffffffff1660e01b81526004015f60405180830381865afa15801562006217573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200624091908101906200f784565b905062006250815160036200a554565b6200627c815f815181106200626957620062696200f8b0565b602002602001015160200151466200a554565b620062a6815f815181106200629557620062956200f8b0565b60200260200101515f01516200a489565b620062d381600181518110620062c057620062c06200f8b0565b602002602001015160200151866200a554565b620062ed816001815181106200629557620062956200f8b0565b62006328816001815181106200630757620063076200f8b0565b6020908102919091010151604001516022546001600160a01b03166200a62b565b62006363816001815181106200634257620063426200f8b0565b6020026020010151606001518051906020012084805190602001206200a5d2565b62006390816002815181106200637d576200637d6200f8b0565b602002602001015160200151856200a554565b620063aa816002815181106200629557620062956200f8b0565b620063c4816002815181106200630757620063076200f8b0565b620063ff81600281518110620063de57620063de6200f8b0565b6020026020010151606001518051906020012083805190602001206200a5d2565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562006456575f80fd5b505af115801562006469573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb969350620064a8928a92169088905f906004016200f1c5565b5f604051808303815f87803b158015620064c0575f80fd5b505af1158015620064d3573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b0316639ca220dd6040518163ffffffff1660e01b81526004015f60405180830381865afa15801562006527573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200655091908101906200f784565b905062006560815160036200a554565b5f805b8251811015620065d257868382815181106200658357620065836200f8b0565b60200260200101516020015103620065c957620065bf838281518110620065ae57620065ae6200f8b0565b60200260200101515f01516200a501565b60019150620065d2565b60010162006563565b50620015de816040518060400160405280601d81526020017f4661696c656420746f2066696e6420696e61637469766520636861696e0000008152506200a6e1565b604080518082018252600981527f626c6f636b54696d65000000000000000000000000000000000000000000000060208083019190915282516005918101919091526001925f910160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620066c1575f80fd5b505af1158015620066d4573d5f803e3d5ffd5b505060408051602480820188905282518083039091018152604490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8e6feba50000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb392506200678d91906004016200f01e565b5f604051808303815f87803b158015620067a5575f80fd5b505af1158015620067b8573d5f803e3d5ffd5b50506020546040517f2259e9e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632259e9e5915062002f6f908690869086906004016200f032565b6024805460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791015f604051808303815f87803b15801562006861575f80fd5b505af115801562006874573d5f803e3d5ffd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015f604051808303815f87803b158015620068fb575f80fd5b505af11580156200690e573d5f803e3d5ffd5b50506020546040517f8f2839700000000000000000000000000000000000000000000000000000000081525f60048201526001600160a01b039091169250638f28397091506024016200336b565b604080518082018252600381527f544b4e000000000000000000000000000000000000000000000000000000000060208083019190915291516d1357000000000000000000000000928101929092525f91600190839060340160408051808303601f190181528282018252600583527f45524332300000000000000000000000000000000000000000000000000000006020840152602554915163ca669fa760e01b81526001600160a01b0390921660048301529250601290737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562006a4c575f80fd5b505af115801562006a5f573d5f803e3d5ffd5b50506040517fc31eb0e00000000000000000000000000000000000000000000000000000000081527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c31eb0e091506024015b5f604051808303815f87803b15801562006ae7575f80fd5b505af115801562006afa573d5f803e3d5ffd5b50506020546040517f6e9e2d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250636e9e2d3f9150620040b5908990899089908990899089906004016200f201565b6060601c805480602002602001604051908101604052809291908181526020015f905b8282101562002409575f8481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801562006c3457602002820191905f5260205f20905f905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001906004019060208260030104928301926001038202915080841162006be05790505b5050505050815250508152602001906001019062006b76565b60208054604080517fbd8fde1c00000000000000000000000000000000000000000000000000000000815290516177779362006cba936001600160a01b0316926391d1485492849263bd8fde1c92600480820193918290030181865afa15801562005ae3573d5f803e3d5ffd5b60208054604080517fe63ab1e9000000000000000000000000000000000000000000000000000000008152905162006d27936001600160a01b03909316926391d1485492849263e63ab1e9926004808401939192918290030181865afa15801562005ae3573d5f803e3d5ffd5b60208054604080517fbd8fde1c000000000000000000000000000000000000000000000000000000008152905162006d94936001600160a01b03909316926391d1485492849263bd8fde1c926004808401939192918290030181865afa15801562004977573d5f803e3d5ffd5b60208054604080517fe63ab1e9000000000000000000000000000000000000000000000000000000008152905162006e01936001600160a01b03909316926391d1485492849263e63ab1e9926004808401939192918290030181865afa15801562004977573d5f803e3d5ffd5b60208054604080517f0c63109e000000000000000000000000000000000000000000000000000000008152905162006e9c936001600160a01b0390931692630c63109e92600480820193918290030181865afa15801562006e64573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062006e8a91906200f087565b6025546001600160a01b03166200a62b565b6024805460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9163ca669fa791015f604051808303815f87803b15801562006ef2575f80fd5b505af115801562006f05573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562006f69575f80fd5b505af115801562006f7c573d5f803e3d5ffd5b5050602554604080516001600160a01b03928316815291851660208301527f6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279935001905060405180910390a16020546040517f3500c24b0000000000000000000000000000000000000000000000000000000081526001600160a01b03838116600483015290911690633500c24b906024015f604051808303815f87803b15801562007026575f80fd5b505af115801562007039573d5f803e3d5ffd5b505060208054604080517fbd8fde1c00000000000000000000000000000000000000000000000000000000815290516200711d95506001600160a01b0390921693506391d1485492849263bd8fde1c9260048082019392918290030181865afa158015620070a9573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620070cf91906200f725565b60255460405160e084901b7fffffffff0000000000000000000000000000000000000000000000000000000016815260048101929092526001600160a01b03166024820152604401620049e4565b60208054604080517fe63ab1e900000000000000000000000000000000000000000000000000000000815290516200718a936001600160a01b03909316926391d1485492849263e63ab1e9926004808401939192918290030181865afa158015620070a9573d5f803e3d5ffd5b60208054604080517fbd8fde1c0000000000000000000000000000000000000000000000000000000081529051620071f7936001600160a01b03909316926391d1485492849263bd8fde1c926004808401939192918290030181865afa15801562004e8f573d5f803e3d5ffd5b60208054604080517fe63ab1e9000000000000000000000000000000000000000000000000000000008152905162007264936001600160a01b03909316926391d1485492849263e63ab1e9926004808401939192918290030181865afa15801562004e8f573d5f803e3d5ffd5b60208054604080517f0c63109e000000000000000000000000000000000000000000000000000000008152905162004ffe936001600160a01b0390931692630c63109e92600480820193918290030181865afa15801562004fd1573d5f803e3d5ffd5b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f190181528282018252600983527f636f6e6e6563746f72000000000000000000000000000000000000000000000060208085019190915291516daa5500000000000000000000000092810192909252925061aa5591905f9060340160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620073af575f80fd5b505af1158015620073c2573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb96935062007402928a92169089906001906004016200f1c5565b5f604051808303815f87803b1580156200741a575f80fd5b505af11580156200742d573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801562007488575f80fd5b505af11580156200749b573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b158015620074ff575f80fd5b505af115801562007512573d5f803e3d5ffd5b50505050816040516200752691906200f754565b6040518091039020857f20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581836040516200756091906200f01e565b60405180910390a36020546040517f631d62e40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063631d62e490620075b7908890869086906004016200f032565b5f604051808303815f87803b158015620075cf575f80fd5b505af1158015620075e2573d5f803e3d5ffd5b50506020546040517f5cf92c9f0000000000000000000000000000000000000000000000000000000081525f93508392506001600160a01b0390911690635cf92c9f9062007637908a9088906004016200f59e565b5f60405180830381865afa15801562007652573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200767b91908101906200f8dd565b915091506200768a826200a489565b62001f1d81846200a735565b60606019805480602002602001604051908101604052809291908181526020015f905b8282101562002409578382905f5260205f20018054620076d9906200f551565b80601f016020809104026020016040519081016040528092919081815260200182805462007707906200f551565b8015620077565780601f106200772c5761010080835404028352916020019162007756565b820191905f5260205f20905b8154815290600101906020018083116200773857829003601f168201915b505050505081526020019060010190620076b9565b6008545f9060ff161562007783575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c6564000000000000000000000000000000000000000000000000000060248301525f9163667f9d7090604401602060405180830381865afa15801562007812573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200783891906200f725565b1415905090565b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f190181528282018252600983527f626c6f636b54696d65000000000000000000000000000000000000000000000060208085019190915282516005918101919091529093505f910160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562007914575f80fd5b505af115801562007927573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb96935062007967928992169088906001906004016200f1c5565b5f604051808303815f87803b1580156200797f575f80fd5b505af115801562007992573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b158015620079ed575f80fd5b505af115801562007a00573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562007a64575f80fd5b505af115801562007a77573d5f803e3d5ffd5b50505050837f40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee634838360405162007aaf9291906200f92e565b60405180910390a26020546040517f2259e9e50000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690632259e9e59062007b06908790869086906004016200f032565b5f604051808303815f87803b15801562007b1e575f80fd5b505af115801562007b31573d5f803e3d5ffd5b50506020546040517f7066b18d0000000000000000000000000000000000000000000000000000000081525f93506001600160a01b039091169150637066b18d9062007b8490889087906004016200f59e565b5f60405180830381865afa15801562007b9f573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405262007bc891908101906200f666565b905062001463818051906020012083805190602001206200a5d2565b6040516d987600000000000000000000000060208201526001905f9060340160408051601f198184030181528282526020547f94cc868300000000000000000000000000000000000000000000000000000000845291519093505f926001600160a01b03909216916394cc868391600480830192869291908290030181865afa15801562007c74573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405262007c9d91908101906200f4bb565b905062007cad815160016200a554565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562007d04575f80fd5b505af115801562007d17573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562007d7b575f80fd5b505af115801562007d8e573d5f803e3d5ffd5b5050604051600181528592507fc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e915060200160405180910390a260205460225460405163547965cb60e11b81526001600160a01b039283169263a8f2cb969262007e069288929091169087906001906004016200f1c5565b5f604051808303815f87803b15801562007e1e575f80fd5b505af115801562007e31573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b03166394cc86836040518163ffffffff1660e01b81526004015f60405180830381865afa15801562007e85573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405262007eae91908101906200f4bb565b905062007ebe815160026200a554565b620022c58160018151811062007ed85762007ed86200f8b0565b6020026020010151846200a554565b604080518082018252600381527f544b4e0000000000000000000000000000000000000000000000000000000000602080830191909152825180820184525f815283518085018552600581527f455243323000000000000000000000000000000000000000000000000000000092810192909252602554935163ca669fa760e01b81526001600160a01b03909416600485015261dddd9360019290601290737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562007fbc575f80fd5b505af115801562007fcf573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b15801562008033575f80fd5b505af115801562008046573d5f803e3d5ffd5b50505050856001600160a01b0316836040516200806491906200f754565b60405180910390207fa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367838789604051620080a1939291906200f761565b60405180910390a36020546040517f6e9e2d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690636e9e2d3f90620040b5908990899089908990899089906004016200f201565b60408051808201825260048082527f544b4e3100000000000000000000000000000000000000000000000000000000602080840191909152835180850185529182527f544b4e32000000000000000000000000000000000000000000000000000000008282015292516d11110000000000000000000000009381019390935261dddd9261eeee9291906001905f9060340160408051601f19818403018152908290526d2222000000000000000000000000602083015291505f9060340160408051808303601f190181528282018252600583527f4552433230000000000000000000000000000000000000000000000000000000602084015260255491517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b0390921660048301529250601290737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d56906024015f604051808303815f87803b1580156200826b575f80fd5b505af11580156200827e573d5f803e3d5ffd5b50506020546040517f6e9e2d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250636e9e2d3f9150620082d7908c908b908a908a90899089906004016200f201565b5f604051808303815f87803b158015620082ef575f80fd5b505af115801562008302573d5f803e3d5ffd5b50506020546040517f6e9e2d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250636e9e2d3f91506200835b908b908a908a908990899089906004016200f201565b5f604051808303815f87803b15801562008373575f80fd5b505af115801562008386573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015620083e5575f80fd5b505af1158015620083f8573d5f803e3d5ffd5b505050505f60205f9054906101000a90046001600160a01b03166001600160a01b031663c1bd469f6040518163ffffffff1660e01b81526004015f60405180830381865afa1580156200844d573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200847691908101906200f956565b905062008486815160026200a554565b5f805b825181101562008525578b6001600160a01b0316838281518110620084b257620084b26200f8b0565b6020026020010151602001516001600160a01b0316036200851c57620084e68382815181106200629557620062956200f8b0565b62008512838281518110620084ff57620084ff6200f8b0565b6020026020010151608001518b6200a68d565b6001915062008525565b60010162008489565b5062008567816040518060400160405280601681526020017f4661696c656420746f2066696e6420746f6b656e2031000000000000000000008152506200a6e1565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b158015620085be575f80fd5b505af1158015620085d1573d5f803e3d5ffd5b50506020546040517f9060bda90000000000000000000000000000000000000000000000000000000081526001600160a01b038e811660048301525f60248301529091169250639060bda991506044015f604051808303815f87803b15801562008639575f80fd5b505af11580156200864c573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b031663c1bd469f6040518163ffffffff1660e01b81526004015f60405180830381865afa158015620086a0573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052620086c991908101906200f956565b9150620086d9825160026200a554565b5f805b83518110156200874c578b6001600160a01b03168482815181106200870557620087056200f8b0565b6020026020010151602001516001600160a01b031603620087435762008739848281518110620065ae57620065ae6200f8b0565b600191506200874c565b600101620086dc565b50620053a1816040518060400160405280601d81526020017f4661696c656420746f2066696e6420696e61637469766520746f6b656e0000008152506200a6e1565b60408051602080820183525f80835292516d13570000000000000000000000009181019190915261dddd9260019160340160408051808303601f190181528282018252600583527f45524332300000000000000000000000000000000000000000000000000000006020840152602554915163ca669fa760e01b81526001600160a01b0390921660048301529250601290737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562008856575f80fd5b505af115801562008869573d5f803e3d5ffd5b505060408051602060248201819052601660448301527f53796d626f6c2063616e6e6f7420626520656d707479000000000000000000006064808401919091528351808403909101815260849092018352810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ec10cfd0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d935063f28dceb3925062006acf91906004016200f01e565b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f19018152602080840183525f80855292516daa550000000000000000000000009181019190915290935061aa5592919060340160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562008a11575f80fd5b505af115801562008a24573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb96935062008a64928a92169089906001906004016200f1c5565b5f604051808303815f87803b15801562008a7c575f80fd5b505af115801562008a8f573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801562008aea575f80fd5b505af115801562008afd573d5f803e3d5ffd5b5050604051737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507ec10cfd000000000000000000000000000000000000000000000000000000009062008b519086906024016200f01e565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e084901b909216825262008bc5916004016200f01e565b5f604051808303815f87803b15801562008bdd575f80fd5b505af115801562008bf0573d5f803e3d5ffd5b50506020546040517f631d62e40000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063631d62e49150620029d7908890869086906004016200f032565b606060158054806020026020016040519081016040528092919081815260200182805480156200163e57602002820191905f5260205f209081546001600160a01b031681526001909101906020018083116200161f575050505050905090565b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f190181528282018252600983527f626c6f636b54696d65000000000000000000000000000000000000000000000060208085019190915282516005918101919091529093505f910160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562008d78575f80fd5b505af115801562008d8b573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb96935062008dcb928992169088906001906004016200f1c5565b5f604051808303815f87803b15801562008de3575f80fd5b505af115801562008df6573d5f803e3d5ffd5b505060265460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801562008e51575f80fd5b505af115801562008e64573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b031663f48448146040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562008ec3575f80fd5b505af115801562008ed6573d5f803e3d5ffd5b50506020546040517f2259e9e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250632259e9e5915062000817908790869086906004016200f032565b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562008fad575f80fd5b505af115801562008fc0573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb96935062009000928792169086906001906004016200f1c5565b5f604051808303815f87803b15801562009018575f80fd5b505af11580156200902b573d5f803e3d5ffd5b5050604080518082018252600981527f636f6e6e6563746f72000000000000000000000000000000000000000000000060208083019190915282518084018452600781527f67617465776179000000000000000000000000000000000000000000000000008183015283518085018552600381527f74737300000000000000000000000000000000000000000000000000000000008184015284516daa5500000000000000000000000093810193909352845160148185030181526034840186526dbb660000000000000000000000006054850152855160488186030181526068850187526dcc7700000000000000000000000060888601528651607c818703018152609c86018089526025547f06447d56000000000000000000000000000000000000000000000000000000009091526001600160a01b031660a087015296519598509296509094909391929091737109709ecfa91a80626ff3989d68f67f5b1dd12d916306447d569160c0808301925f92919082900301818387803b158015620091b5575f80fd5b505af1158015620091c8573d5f803e3d5ffd5b50506020546040517f631d62e40000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063631d62e491506200921b908b908a9088906004016200f032565b5f604051808303815f87803b15801562009233575f80fd5b505af115801562009246573d5f803e3d5ffd5b50506020546040517f631d62e40000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063631d62e4915062009299908b90899087906004016200f032565b5f604051808303815f87803b158015620092b1575f80fd5b505af1158015620092c4573d5f803e3d5ffd5b50506020546040517f631d62e40000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063631d62e4915062009317908b90889086906004016200f032565b5f604051808303815f87803b1580156200932f575f80fd5b505af115801562009342573d5f803e3d5ffd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d5f1c6001600160a01b03166390c5013b6040518163ffffffff1660e01b81526004015f604051808303815f87803b158015620093a1575f80fd5b505af1158015620093b4573d5f803e3d5ffd5b505050505f60205f9054906101000a90046001600160a01b03166001600160a01b03166318d3ce966040518163ffffffff1660e01b81526004015f60405180830381865afa15801562009409573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200943291908101906200faea565b905062009442815160036200a554565b5f805f805b84518110156200972e57620094b26040518060400160405280600e81526020017f436f6e747261637420747970653a000000000000000000000000000000000000815250868381518110620094a057620094a06200f8b0565b6020026020010151604001516200a789565b620095136040518060400160405280600981526020017f436861696e2049443a00000000000000000000000000000000000000000000008152508683815181106200950157620095016200f8b0565b6020026020010151606001516200a804565b8a805190602001208582815181106200953057620095306200f8b0565b6020026020010151604001518051906020012003620095d257620095758582815181106200956257620095626200f8b0565b6020026020010151606001518e6200a554565b6200958e8582815181106200629557620062956200f8b0565b620095c8858281518110620095a757620095a76200f8b0565b6020026020010151602001518051906020012089805190602001206200a5d2565b6001935062009725565b8980519060200120858281518110620095ef57620095ef6200f8b0565b60200260200101516040015180519060200120036200967e57620096218582815181106200956257620095626200f8b0565b6200963a8582815181106200629557620062956200f8b0565b620096748582815181106200965357620096536200f8b0565b6020026020010151602001518051906020012088805190602001206200a5d2565b6001925062009725565b88805190602001208582815181106200969b576200969b6200f8b0565b60200260200101516040015180519060200120036200972557620096cd8582815181106200956257620095626200f8b0565b620096e68582815181106200629557620062956200f8b0565b62009720858281518110620096ff57620096ff6200f8b0565b6020026020010151602001518051906020012087805190602001206200a5d2565b600191505b60010162009447565b5062009754836040518060600160405280602181526020016201404f602191396200a6e1565b62009795826040518060400160405280601f81526020017f4661696c656420746f2066696e64206761746577617920636f6e7472616374008152506200a6e1565b620097d6816040518060400160405280601b81526020017f4661696c656420746f2066696e642074737320636f6e747261637400000000008152506200a6e1565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200982d575f80fd5b505af115801562009840573d5f803e3d5ffd5b50506020546040517f10d29b9e0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911692506310d29b9e915062009893908f908d905f906004016200f5fc565b5f604051808303815f87803b158015620098ab575f80fd5b505af1158015620098be573d5f803e3d5ffd5b5050505060205f9054906101000a90046001600160a01b03166001600160a01b03166318d3ce966040518163ffffffff1660e01b81526004015f60405180830381865afa15801562009912573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200993b91908101906200faea565b93506200994b845160036200a554565b5f805b8551811015620099ba578a805190602001208682815181106200997557620099756200f8b0565b6020026020010151604001518051906020012003620099b157620099a7868281518110620065ae57620065ae6200f8b0565b60019150620099ba565b6001016200994e565b50620057aa8160405180606001604052806028815260200162014027602891396200a6e1565b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f190181528282018252600983527f636f6e6e6563746f72000000000000000000000000000000000000000000000060208085019190915291516daa5500000000000000000000000092810192909252925061aa5591905f9060340160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562009ac8575f80fd5b505af115801562009adb573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb96935062009b1b928a92169089906001906004016200f1c5565b5f604051808303815f87803b15801562009b33575f80fd5b505af115801562009b46573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801562009ba1575f80fd5b505af115801562009bb4573d5f803e3d5ffd5b50506020546040517f631d62e40000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063631d62e4915062009c07908890869086906004016200f032565b5f604051808303815f87803b15801562009c1f575f80fd5b505af115801562009c32573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801562009c8d575f80fd5b505af115801562009ca0573d5f803e3d5ffd5b5050604051737109709ecfa91a80626ff3989d68f67f5b1dd12d925063f28dceb391507f2b192eab000000000000000000000000000000000000000000000000000000009062008b51908990879087906024016200f032565b6040516d987600000000000000000000000060208201526001905f9060340160408051808303601f190181528282018252600983527f636f6e6e6563746f72000000000000000000000000000000000000000000000060208085019190915291516daa5500000000000000000000000092810192909252925061aa5591905f9060340160408051808303601f190181529082905260255463ca669fa760e01b83526001600160a01b031660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b15801562009de1575f80fd5b505af115801562009df4573d5f803e3d5ffd5b505060205460225460405163547965cb60e11b81526001600160a01b03928316945063a8f2cb96935062009e34928a92169089906001906004016200f1c5565b5f604051808303815f87803b15801562009e4c575f80fd5b505af115801562009e5f573d5f803e3d5ffd5b505060255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa791506024015f604051808303815f87803b15801562009eba575f80fd5b505af115801562009ecd573d5f803e3d5ffd5b50506020546040517f631d62e40000000000000000000000000000000000000000000000000000000081526001600160a01b03909116925063631d62e4915062009f20908890869086906004016200f032565b5f604051808303815f87803b15801562009f38575f80fd5b505af115801562009f4b573d5f803e3d5ffd5b50506020546040517f5cf92c9f0000000000000000000000000000000000000000000000000000000081525f93506001600160a01b039091169150635cf92c9f9062009f9e90899087906004016200f59e565b5f60405180830381865afa15801562009fb9573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405262009fe291908101906200f8dd565b50905062009ff0816200a489565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200a047575f80fd5b505af11580156200a05a573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b1580156200a0be575f80fd5b505af11580156200a0d1573d5f803e3d5ffd5b505050507f6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c826040516200a10691906200f01e565b60405180910390a16020546040517f10d29b9e0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906310d29b9e906200a15d90899087905f906004016200f5fc565b5f604051808303815f87803b1580156200a175575f80fd5b505af11580156200a188573d5f803e3d5ffd5b50506020546040517f5cf92c9f0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250635cf92c9f91506200a1d990899087906004016200f59e565b5f60405180830381865afa1580156200a1f4573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200a21d91908101906200f8dd565b5090506200a22b816200a501565b60255460405163ca669fa760e01b81526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa7906024015f604051808303815f87803b1580156200a282575f80fd5b505af11580156200a295573d5f803e3d5ffd5b505060405163248e63e160e11b815260016004820181905260248201819052604482018190526064820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063491cc7c291506084015f604051808303815f87803b1580156200a2f9575f80fd5b505af11580156200a30c573d5f803e3d5ffd5b505050507f6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c826040516200a34191906200f01e565b60405180910390a16020546040517f10d29b9e0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116906310d29b9e906200a39990899087906001906004016200f5fc565b5f604051808303815f87803b1580156200a3b1575f80fd5b505af11580156200a3c4573d5f803e3d5ffd5b50506020546040517f5cf92c9f0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169250635cf92c9f91506200a41590899087906004016200f59e565b5f60405180830381865afa1580156200a430573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200a45991908101906200f8dd565b509050620015de816200a489565b5f6200a4726200ebd5565b6200a47f8484836200a87b565b9150505b92915050565b6040517f0c9fd5810000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90630c9fd581906024015b5f6040518083038186803b1580156200a4ee575f80fd5b505afa15801562001463573d5f803e3d5ffd5b6040517fa59828850000000000000000000000000000000000000000000000000000000081528115156004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063a5982885906024016200a4d7565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c54906044015b5f6040518083038186803b1580156200a5bf575f80fd5b505afa158015620015de573d5f803e3d5ffd5b6040517f7c84c69b0000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d90637c84c69b906044016200a5a8565b6040517f515361f60000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015282166024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063515361f6906044016200a5a8565b6040517ff320d963000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063f320d963906200a5a890859085906004016200f92e565b6040517fa34edc03000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063a34edc03906200a5a890859085906004016200fc28565b6040517f97624631000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d906397624631906200a5a890859085906004016200f92e565b6200a80082826040516024016200a7a29291906200f92e565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f4b5c4277000000000000000000000000000000000000000000000000000000001790526200a8fb565b5050565b6200a80082826040516024016200a81d9291906200fc44565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f9710a9d0000000000000000000000000000000000000000000000000000000001790526200a8fb565b5f806200a88985846200a906565b90506200a8f06040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f787900000081525082866040516020016200a8da9291906200fc67565b604051602081830303815290604052856200a913565b9150505b9392505050565b62004ffe816200a946565b5f6200a8f483836200a966565b60c0810151515f90156200a93a576200a93284848460c001516200a984565b90506200a8f4565b6200a93284846200ab38565b80516a636f6e736f6c652e6c6f67602083015f808483855afa5050505050565b5f6200a97383836200ac29565b6200a8f4838360200151846200a913565b5f806200a9906200ac36565b90505f6200a99f86836200ad0a565b90505f6200a9b782606001518360200151856200b1c2565b90505f6200a9c8838389896200b3e7565b90505f6200a9d6826200c329565b602081015181519192509060030b156200aa4e578982604001516040516020016200aa039291906200fc8a565b60408051601f19818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526200aa45916004016200f01e565b60405180910390fd5b5f6200aa926040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a2000000000000000000000008152508360016200c4ff565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d906200aae79084906004016200f01e565b602060405180830381865afa1580156200ab03573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200ab2991906200f087565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081525f908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc925906200ab8e9087906004016200f01e565b5f60405180830381865afa1580156200aba9573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200abd291908101906200f666565b90505f6200ac0382856040516020016200abee9291906200fcef565b6040516020818303038152906040526200c704565b90506001600160a01b0381166200a47f5784846040516020016200aa039291906200fd07565b6200a80082825f6200c715565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c906200acbf9084906004016200fd9b565b5f60405180830381865afa1580156200acda573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200ad0391908101906200f666565b9250505090565b6200ad3d6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b5f737109709ecfa91a80626ff3989d68f67f5b1dd12d90506200ad886040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6200ad93856200c822565b60208201525f6200ada4866200cc1b565b90505f836001600160a01b031663d930a0e66040518163ffffffff1660e01b81526004015f60405180830381865afa1580156200ade3573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200ae0c91908101906200f666565b868385602001516040516020016200ae2894939291906200fde3565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291505f906001600160a01b038616906360f9bb11906200ae819085906004016200f01e565b5f60405180830381865afa1580156200ae9c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200aec591908101906200f666565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f6906200af0f9084906004016200febb565b602060405180830381865afa1580156200af2b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200af5191906200f1a9565b6200af6957816040516020016200aa0391906200ff0e565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906200afb09084906004016200ff94565b5f60405180830381865afa1580156200afcb573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200aff491908101906200f666565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f6906200b03d9084906004016200ffe7565b602060405180830381865afa1580156200b059573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200b07f91906200f1a9565b156200b116576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac8906200b0cc9084906004016200ffe7565b5f60405180830381865afa1580156200b0e7573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200b11091908101906200f666565b60408501525b846001600160a01b03166349c4fac882865f01516040516020016200b13c91906201003a565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016200b16a9291906200f92e565b5f60405180830381865afa1580156200b185573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200b1ae91908101906200f666565b606085015250608083015250949350505050565b60408051600480825260a082019092526060915f9190816020015b60608152602001906001900390816200b1dd5790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250815f815181106200b240576200b2406200f8b0565b60200260200101819052506040518060400160405280600381526020017f2d726c0000000000000000000000000000000000000000000000000000000000815250816001815181106200b297576200b2976200f8b0565b6020026020010181905250846040516020016200b2b591906201009a565b604051602081830303815290604052816002815181106200b2da576200b2da6200f8b0565b6020026020010181905250826040516020016200b2f89190620100fa565b604051602081830303815290604052816003815181106200b31d576200b31d6200f8b0565b60200260200101819052505f6200b334826200c329565b602080820151604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000818501908152825180840184525f80825290860152825180840190935290518252928101929092529192506200b3c5906040805180820182525f80825260209182015281518083019092528451825280850190820152906200ceb0565b6200b3dd57856040516020016200aa03919062010134565b9695505050505050565b60a08101516040805180820182525f80825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d90156200b437565b511590565b6200b5b0578260200151156200b4f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a4016200aa45565b8260c00151156200b5b0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a4016200aa45565b6040805160ff80825261200082019092525f91816020015b60608152602001906001900390816200b5c85790505090505f6040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806200b62590620101e0565b935060ff16815181106200b63d576200b63d6200f8b0565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e37000000000000000000000000000000000000008152506040516020016200b690919062010201565b6040516020818303038152906040528282806200b6ad90620101e0565b935060ff16815181106200b6c5576200b6c56200f8b0565b60200260200101819052506040518060400160405280600681526020017f6465706c6f7900000000000000000000000000000000000000000000000000008152508282806200b71490620101e0565b935060ff16815181106200b72c576200b72c6200f8b0565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d650000000000000000000000000000000000008152508282806200b77b90620101e0565b935060ff16815181106200b793576200b7936200f8b0565b602002602001018190525087602001518282806200b7b190620101e0565b935060ff16815181106200b7c9576200b7c96200f8b0565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e7472616374506174680000000000000000000000000000000000008152508282806200b81890620101e0565b935060ff16815181106200b830576200b8306200f8b0565b6020908102919091010152875182826200b84a81620101e0565b935060ff16815181106200b862576200b8626200f8b0565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e496400000000000000000000000000000000000000000000008152508282806200b8b190620101e0565b935060ff16815181106200b8c9576200b8c96200f8b0565b60200260200101819052506200b8df466200cf16565b82826200b8ec81620101e0565b935060ff16815181106200b904576200b9046200f8b0565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c6500000000000000000000000000000000008152508282806200b95390620101e0565b935060ff16815181106200b96b576200b96b6200f8b0565b6020026020010181905250868282806200b98590620101e0565b935060ff16815181106200b99d576200b99d6200f8b0565b60209081029190910101528551156200bad05760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826200b9f181620101e0565b935060ff16815181106200ba09576200ba096200f8b0565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d906200ba5b9089906004016200f01e565b5f60405180830381865afa1580156200ba76573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200ba9f91908101906200f666565b82826200baac81620101e0565b935060ff16815181106200bac4576200bac46200f8b0565b60200260200101819052505b8460200151156200bbac5760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826200bb1c81620101e0565b935060ff16815181106200bb34576200bb346200f8b0565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806200bb8390620101e0565b935060ff16815181106200bb9b576200bb9b6200f8b0565b60200260200101819052506200bd8e565b6200bbe56200b4328660a001516040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6200bc825760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826200bc2b81620101e0565b935060ff16815181106200bc43576200bc436200f8b0565b60200260200101819052508460a001516040516020016200bc6591906201009a565b6040516020818303038152906040528282806200bb8390620101e0565b8460c001511580156200bcc65750604080890151815180830183525f808252602091820152825180840190935281518352908101908201526200bcc490511590565b155b156200bd8e5760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826200bd0d81620101e0565b935060ff16815181106200bd25576200bd256200f8b0565b60200260200101819052506200bd3b886200cfba565b6040516020016200bd4d91906201009a565b6040516020818303038152906040528282806200bd6a90620101e0565b935060ff16815181106200bd82576200bd826200f8b0565b60200260200101819052505b604080860151815180830183525f808252602091820152825180840190935281518352908101908201526200bdc290511590565b6200be625760408051808201909152600b81527f2d2d72656c617965724964000000000000000000000000000000000000000000602082015282826200be0881620101e0565b935060ff16815181106200be20576200be206200f8b0565b602002602001018190525084604001518282806200be3e90620101e0565b935060ff16815181106200be56576200be566200f8b0565b60200260200101819052505b6060850151156200bf8d5760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826200beae81620101e0565b935060ff16815181106200bec6576200bec66200f8b0565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e8906024015f60405180830381865afa1580156200bf33573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200bf5c91908101906200f666565b82826200bf6981620101e0565b935060ff16815181106200bf81576200bf816200f8b0565b60200260200101819052505b60e085015151156200c0405760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826200bfda81620101e0565b935060ff16815181106200bff2576200bff26200f8b0565b60200260200101819052506200c00f8560e001515f01516200cf16565b82826200c01c81620101e0565b935060ff16815181106200c034576200c0346200f8b0565b60200260200101819052505b60e085015160200151156200c0f75760408051808201909152600a81527f2d2d676173507269636500000000000000000000000000000000000000000000602082015282826200c09081620101e0565b935060ff16815181106200c0a8576200c0a86200f8b0565b60200260200101819052506200c0c68560e00151602001516200cf16565b82826200c0d381620101e0565b935060ff16815181106200c0eb576200c0eb6200f8b0565b60200260200101819052505b60e085015160400151156200c1ae5760408051808201909152600e81527f2d2d6d6178466565506572476173000000000000000000000000000000000000602082015282826200c14781620101e0565b935060ff16815181106200c15f576200c15f6200f8b0565b60200260200101819052506200c17d8560e00151604001516200cf16565b82826200c18a81620101e0565b935060ff16815181106200c1a2576200c1a26200f8b0565b60200260200101819052505b60e085015160600151156200c2655760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826200c1fe81620101e0565b935060ff16815181106200c216576200c2166200f8b0565b60200260200101819052506200c2348560e00151606001516200cf16565b82826200c24181620101e0565b935060ff16815181106200c259576200c2596200f8b0565b60200260200101819052505b5f8160ff1667ffffffffffffffff8111156200c285576200c2856200f26b565b6040519080825280602002602001820160405280156200c2ba57816020015b60608152602001906001900390816200c2a45790505b5090505f5b8260ff168160ff1610156200c31a57838160ff16815181106200c2e6576200c2e66200f8b0565b6020026020010151828260ff16815181106200c306576200c3066200f8b0565b60209081029190910101526001016200c2bf565b5093505050505b949350505050565b6200c35060405180606001604052805f60030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d925f91849163d145736c916200c3d7918691016201025a565b5f60405180830381865afa1580156200c3f2573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200c41b91908101906200f666565b90505f6200c42a86836200dad1565b90505f846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b81526004016200c45b91906200ef0c565b5f604051808303815f875af11580156200c477573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526200c4a09190810190620102a2565b805190915060030b158015906200c4ba5750602081015151155b80156200c4ca5750604081015151155b156200b3dd57815f815181106200c4e5576200c4e56200f8b0565b60200260200101516040516020016200aa0391906201035d565b60605f6200c533856040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925286518252808701908201529091506200c56b9082905b906200dc39565b156200c6d2575f6200c5ec826200c5e5846200c5de6200c5b18a6040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f808252602091820152815180830190925282518252918201519181019190915290565b906200dc61565b906200dcc5565b604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200c6519082906200dc39565b156200c6bd57604080518082018252600181527f0a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200c6ba905b82906200dd51565b90505b6200c6c8816200dd78565b925050506200a8f4565b82156200c6ee5784846040516020016200aa039291906201053c565b505060408051602081019091525f81526200a8f4565b5f808251602084015ff09392505050565b8160a00151156200c72557505050565b5f6200c7338484846200dde3565b90505f6200c741826200c329565b602081015181519192509060030b1580156200c7de5750604080518082018252600781527f53554343455353000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200c7de906040805180820182525f808252602091820152815180830190925284518252808501908201526200c564565b156200c7ec57505050505050565b604082015151156200c80f5781604001516040516020016200aa039190620105c7565b806040516020016200aa03919062010620565b60605f6200c856836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200c8bc905b82906200ceb0565b156200c92f57604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200a8f4906200c9299083906200e3d7565b6200dd78565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200c992905b82906200e469565b6001036200ca6357604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200c9fa906200c6b2565b50604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200a8f4906200c929905b83906200dd51565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200cac3906200c8b4565b156200cc0257604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201819052845180860190955292518452830152906200cb2d9083906200e50f565b90505f81600183516200cb41919062010679565b815181106200cb54576200cb546200f8b0565b602002602001015190506200cbf96200c9296200cbcc6040518060400160405280600581526020017f2e6a736f6e0000000000000000000000000000000000000000000000000000008152506040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b6040805180820182525f80825260209182015281518083019092528551825280860190820152906200e3d7565b95945050505050565b826040516020016200aa0391906201068f565b50919050565b60605f6200cc4f836040805180820182525f8082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201529091506200ccb2906200c8b4565b156200ccc3576200a8f4816200dd78565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200cd23906200c98a565b6001036200cd9057604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200a8f4906200c929906200ca5b565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200cdf0906200c8b4565b156200cc0257604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f808252908201819052845180860190955292518452830152906200ce5a9083906200e50f565b90506001815111156200ce9c5780600282516200ce78919062010679565b815181106200ce8b576200ce8b6200f8b0565b602002602001015192505050919050565b50826040516020016200aa0391906201068f565b805182515f9111156200cec557505f6200a483565b8151835160208501515f92916200cedc9162010761565b6200cee8919062010679565b9050826020015181036200cf015760019150506200a483565b82516020840151819020912014905092915050565b60605f6200cf24836200e5cb565b60010190505f8167ffffffffffffffff8111156200cf46576200cf466200f26b565b6040519080825280601f01601f1916602001820160405280156200cf71576020820181803683370190505b5090508181016020015b5f19017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846200cf7b57509392505050565b604081810151815180830183525f808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e53454400000000000000000000000000000000000000000000818401908152855180870187528381528401929092528451808601909552518452908301526060916200d047905b82906200e6b3565b156200d08857505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d0e8906200d03f565b156200d12957505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d495400000000000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d189906200d03f565b156200d1ca57505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d22a906200d03f565b806200d2915750604080518082018252601081527f47504c2d322e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d291906200d03f565b156200d2d257505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c7900000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d332906200d03f565b806200d3995750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d399906200d03f565b156200d3da57505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d43a906200d03f565b806200d4a15750604080518082018252601181527f4c47504c2d322e312d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d4a1906200d03f565b156200d4e257505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d542906200d03f565b806200d5a95750604080518082018252601181527f4c47504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d5a9906200d03f565b156200d5ea57505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d64a906200d03f565b156200d68b57505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d6eb906200d03f565b156200d72c57505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d78c906200d03f565b156200d7cd57505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d82d906200d03f565b156200d86e57505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d8ce906200d03f565b156200d90f57505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c79000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d96f906200d03f565b806200d9d65750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200d9d6906200d03f565b156200da1757505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185525f80825290820152835180850190945291518352908201526200da77906200d03f565b156200dab857505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b604080840151845191516200aa03929060200162010777565b6060805f5b84518110156200db6757818582815181106200daf6576200daf66200f8b0565b60200260200101516040516020016200db119291906200fcef565b6040516020818303038152906040529150600185516200db32919062010679565b81146200db5e57816040516020016200db4c9190620108c9565b60405160208183030381529060405291505b6001016200dad6565b50604080516003808252608082019092525f91816020015b60608152602001906001900390816200db7f57905050905083815f815181106200dbad576200dbad6200f8b0565b60200260200101819052506040518060400160405280600281526020017f2d63000000000000000000000000000000000000000000000000000000000000815250816001815181106200dc04576200dc046200f8b0565b602002602001018190525081816002815181106200dc26576200dc266200f8b0565b6020908102919091010152949350505050565b60208083015183518351928401515f936200dc5892918491906200e6c8565b14159392505050565b604080518082019091525f80825260208201525f6200dc91845f01518560200151855f015186602001516200e7fb565b90508360200151816200dca5919062010679565b845185906200dcb690839062010679565b90525060208401525090919050565b604080518082019091525f80825260208201528151835110156200dceb5750816200a483565b60208083015190840151600191146200dd135750815160208481015190840151829020919020145b80156200dd49578251845185906200dd2d90839062010679565b90525082516020850180516200dd4590839062010761565b9052505b509192915050565b604080518082019091525f80825260208201526200dd718383836200e93b565b5092915050565b60605f825f015167ffffffffffffffff8111156200dd9a576200dd9a6200f26b565b6040519080825280601f01601f1916602001820160405280156200ddc5576020820181803683370190505b5090505f6020820190506200dd71818560200151865f01516200e9ef565b60605f6200ddf06200ac36565b6040805160ff80825261200082019092529192505f9190816020015b60608152602001906001900390816200de0c5790505090505f6040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806200de6990620101e0565b935060ff16815181106200de81576200de816200f8b0565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e33000000000000000000000000000000000000000000000000008152506040516020016200ded4919062010903565b6040516020818303038152906040528282806200def190620101e0565b935060ff16815181106200df09576200df096200f8b0565b60200260200101819052506040518060400160405280600881526020017f76616c69646174650000000000000000000000000000000000000000000000008152508282806200df5890620101e0565b935060ff16815181106200df70576200df706200f8b0565b6020026020010181905250826040516020016200df8e9190620100fa565b6040516020818303038152906040528282806200dfab90620101e0565b935060ff16815181106200dfc3576200dfc36200f8b0565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e7472616374000000000000000000000000000000000000000000008152508282806200e01290620101e0565b935060ff16815181106200e02a576200e02a6200f8b0565b60200260200101819052506200e04187846200ea77565b82826200e04e81620101e0565b935060ff16815181106200e066576200e0666200f8b0565b6020908102919091010152855151156200e11e5760408051808201909152600b81527f2d2d7265666572656e6365000000000000000000000000000000000000000000602082015282826200e0bb81620101e0565b935060ff16815181106200e0d3576200e0d36200f8b0565b60200260200101819052506200e0ed865f0151846200ea77565b82826200e0fa81620101e0565b935060ff16815181106200e112576200e1126200f8b0565b60200260200101819052505b8560800151156200e1935760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b0000000000000000602082015282826200e16a81620101e0565b935060ff16815181106200e182576200e1826200f8b0565b60200260200101819052506200e1ff565b84156200e1ff5760408051808201909152601281527f2d2d726571756972655265666572656e63650000000000000000000000000000602082015282826200e1db81620101e0565b935060ff16815181106200e1f3576200e1f36200f8b0565b60200260200101819052505b604086015151156200e2a65760408051808201909152600d81527f2d2d756e73616665416c6c6f7700000000000000000000000000000000000000602082015282826200e24c81620101e0565b935060ff16815181106200e264576200e2646200f8b0565b602002602001018190525085604001518282806200e28290620101e0565b935060ff16815181106200e29a576200e29a6200f8b0565b60200260200101819052505b8560600151156200e3165760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d6573000000000000000000000000602082015282826200e2f281620101e0565b935060ff16815181106200e30a576200e30a6200f8b0565b60200260200101819052505b5f8160ff1667ffffffffffffffff8111156200e336576200e3366200f26b565b6040519080825280602002602001820160405280156200e36b57816020015b60608152602001906001900390816200e3555790505b5090505f5b8260ff168160ff1610156200e3cb57838160ff16815181106200e397576200e3976200f8b0565b6020026020010151828260ff16815181106200e3b7576200e3b76200f8b0565b60209081029190910101526001016200e370565b50979650505050505050565b604080518082019091525f80825260208201528151835110156200e3fd5750816200a483565b8151835160208501515f92916200e4149162010761565b6200e420919062010679565b602084015190915060019082146200e442575082516020840151819020908220145b80156200e460578351855186906200e45c90839062010679565b9052505b50929392505050565b5f80825f01516200e48b855f01518660200151865f015187602001516200e7fb565b6200e497919062010761565b90505b835160208501516200e4ad919062010761565b81116200dd7157816200e4c08162010936565b925050825f01516200e4fb8560200151836200e4dd919062010679565b86516200e4eb919062010679565b83865f015187602001516200e7fb565b6200e507919062010761565b90506200e49a565b60605f6200e51e84846200e469565b6200e52b90600162010761565b67ffffffffffffffff8111156200e546576200e5466200f26b565b6040519080825280602002602001820160405280156200e57b57816020015b60608152602001906001900390816200e5655790505b5090505f5b81518110156200e5c3576200e59a6200c92986866200dd51565b8282815181106200e5af576200e5af6200f8b0565b60209081029190910101526001016200e580565b509392505050565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106200e614577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106200e641576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106200e66057662386f26fc10000830492506010015b6305f5e10083106200e679576305f5e100830492506008015b61271083106200e68e57612710830492506004015b606483106200e6a1576064830492506002015b600a83106200a4835760010192915050565b5f6200e6c083836200eaba565b159392505050565b5f808584116200e7f157602084116200e791575f84156200e71c5760016200e6f286602062010679565b6200e6ff90600862010951565b6200e70c90600262010a67565b6200e718919062010679565b1990505b83518116856200e72d898962010761565b6200e739919062010679565b805190935082165b8181146200e779578784116200e75e57879450505050506200c321565b836200e76a8162010a74565b9450508284511690506200e741565b6200e785878562010761565b9450505050506200c321565b8383206200e7a0858862010679565b6200e7ac908762010761565b91505b8582106200e7ef578482208082036200e7d9576200e7ce868462010761565b93505050506200c321565b6200e7e660018462010679565b9250506200e7af565b505b5092949350505050565b5f83818685116200e92457602085116200e8ca575f85156200e8505760016200e82687602062010679565b6200e83390600862010951565b6200e84090600262010a67565b6200e84c919062010679565b1990505b845181165f876200e8628b8b62010761565b6200e86e919062010679565b855190915083165b8281146200e8bb578186106200e8a0576200e8928b8b62010761565b96505050505050506200c321565b856200e8ac8162010936565b9650508386511690506200e876565b8596505050505050506200c321565b508383205f905b6200e8dd868962010679565b82116200e922578583208082036200e8fc57839450505050506200c321565b6200e90960018562010761565b93505081806200e9199062010936565b9250506200e8d1565b505b6200e930878762010761565b979650505050505050565b604080518082019091525f80825260208201525f6200e96b855f01518660200151865f015187602001516200e7fb565b6020808701805191860191909152519091506200e989908262010679565b8352845160208601516200e99e919062010761565b81036200e9ae575f85526200e9e6565b835183516200e9be919062010761565b855186906200e9cf90839062010679565b90525083516200e9e0908262010761565b60208601525b50909392505050565b602081106200ea2f57815183526200ea0960208462010761565b92506200ea1860208362010761565b91506200ea2760208262010679565b90506200e9ef565b5f1981156200ea645760016200ea4783602062010679565b6200ea559061010062010a67565b6200ea61919062010679565b90505b9151835183169219169190911790915250565b60605f6200ea8684846200ad0a565b80516020808301516040519394506200eaa29390910162010a8c565b60405160208183030381529060405291505092915050565b815181515f91908111156200eacd575081515b602080850151908401515f5b838110156200eb9a57825182518082146200eb63575f1960208710156200eb40576001846200eb0a89602062010679565b6200eb16919062010761565b6200eb2390600862010951565b6200eb3090600262010a67565b6200eb3c919062010679565b1990505b81811683821681810391146200eb605797506200a4839650505050505050565b50505b6200eb7060208662010761565b94506200eb7f60208562010761565b935050506020816200eb92919062010761565b90506200ead9565b50845186516200b3dd919062010acb565b61054b8062010aee83390190565b61102e806201103983390190565b611fc0806201206783390190565b6040518060e001604052806060815260200160608152602001606081526020015f151581526020015f151581526020015f151581526020016200ec176200ec1c565b905290565b6040518061010001604052805f151581526020015f15158152602001606081526020015f801916815260200160608152602001606081526020015f151581526020016200ec1760405180608001604052805f81526020015f81526020015f81526020015f81525090565b602080825282518282018190525f918401906040840190835b818110156200ecc85783516001600160a01b03168352602093840193909201916001016200ec9f565b509095945050505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200ee00577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b8801810191908801905f5b818110156200ede5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a85030183526200edce8486516200ecd3565b60209586019590945092909201916001016200ed91565b5091975050506020948501949290920191506001016200ed27565b50929695505050505050565b5f8151808452602084019350602083015f5b828110156200ee605781517fffffffff00000000000000000000000000000000000000000000000000000000168652602095860195909101906001016200ee1e565b5093949350505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200ee00577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051604087526200eed760408801826200ecd3565b90506020820151915086810360208801526200eef481836200ee0c565b9650505060209384019391909101906001016200ee90565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200ee00577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184526200ef6f8583516200ecd3565b945060209384019391909101906001016200ef32565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b828110156200ee00577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b03815116865260208101519050604060208701526200f00760408701826200ee0c565b95505060209384019391909101906001016200efab565b602081525f6200a8f460208301846200ecd3565b838152606060208201525f6200f04c60608301856200ecd3565b82810360408401526200b3dd81856200ecd3565b6001600160a01b038116811462004ffe575f80fd5b80516200f082816200f060565b919050565b5f602082840312156200f098575f80fd5b81516200a8f4816200f060565b610100815260056101008201527f544f4b454e000000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f544b4e00000000000000000000000000000000000000000000000000000000006101608201525f6101808201905060ff88166040830152866060830152600386106200f15d577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b8560808301528460a08301526200f17f60c08301856001600160a01b03169052565b6001600160a01b03831660e0830152979650505050505050565b805180151581146200f082575f80fd5b5f602082840312156200f1ba575f80fd5b6200a8f4826200f199565b8481526001600160a01b0384166020820152608060408201525f6200f1ee60808301856200ecd3565b9050821515606083015295945050505050565b6001600160a01b038716815260c060208201525f6200f22460c08301886200ecd3565b86604084015282810360608401526200f23e81876200ecd3565b905082810360808401526200f25481866200ecd3565b91505060ff831660a0830152979650505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040516080810167ffffffffffffffff811182821017156200f2be576200f2be6200f26b565b60405290565b60405160e0810167ffffffffffffffff811182821017156200f2be576200f2be6200f26b565b6040516060810167ffffffffffffffff811182821017156200f2be576200f2be6200f26b565b604051601f8201601f1916810167ffffffffffffffff811182821017156200f33c576200f33c6200f26b565b604052919050565b5f82601f8301126200f354575f80fd5b8151602083015f8067ffffffffffffffff8411156200f377576200f3776200f26b565b50601f8301601f19166020016200f38e816200f310565b9150508281528583830111156200f3a3575f80fd5b8282602083015e5f92810160200192909252509392505050565b805160ff811681146200f082575f80fd5b5f805f805f8060c087890312156200f3e4575f80fd5b6200f3ef876200f199565b9550602087015167ffffffffffffffff8111156200f40b575f80fd5b6200f41989828a016200f344565b604089015160608a01519197509550905067ffffffffffffffff8111156200f43f575f80fd5b6200f44d89828a016200f344565b935050608087015167ffffffffffffffff8111156200f46a575f80fd5b6200f47889828a016200f344565b9250506200f48960a088016200f3bd565b90509295509295509295565b5f67ffffffffffffffff8211156200f4b1576200f4b16200f26b565b5060051b60200190565b5f602082840312156200f4cc575f80fd5b815167ffffffffffffffff8111156200f4e3575f80fd5b8201601f810184136200f4f4575f80fd5b80516200f50b6200f505826200f495565b6200f310565b8082825260208201915060208360051b8501019250868311156200f52d575f80fd5b6020840193505b828410156200b3dd5783518252602093840193909101906200f534565b600181811c908216806200f56657607f821691505b6020821081036200cc15577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b828152604060208201525f6200c32160408301846200ecd3565b848152608060208201525f6200f5d260808301866200ecd3565b82810360408401526200f5e681866200ecd3565b905082810360608401526200e93081856200ecd3565b838152606060208201525f6200f61660608301856200ecd3565b90508215156040830152949350505050565b606081525f6200f63c60608301866200ecd3565b82810360208401526200f65081866200ecd3565b905082810360408401526200b3dd81856200ecd3565b5f602082840312156200f677575f80fd5b815167ffffffffffffffff8111156200f68e575f80fd5b6200a47f848285016200f344565b6001600160a01b038616815260c06020820152600460c08201527f544b4e320000000000000000000000000000000000000000000000000000000060e082015284604082015261010060608201525f6200f6fb6101008301866200ecd3565b82810360808401526200f70f81866200ecd3565b91505060ff831660a08301529695505050505050565b5f602082840312156200f736575f80fd5b5051919050565b5f81518060208401855e5f93019283525090919050565b5f6200a8f482846200f73d565b60ff84168152826020820152606060408201525f6200cbf960608301846200ecd3565b5f602082840312156200f795575f80fd5b815167ffffffffffffffff8111156200f7ac575f80fd5b8201601f810184136200f7bd575f80fd5b80516200f7ce6200f505826200f495565b8082825260208201915060208360051b8501019250868311156200f7f0575f80fd5b602084015b838110156200f8a557805167ffffffffffffffff8111156200f815575f80fd5b85016080818a03601f190112156200f82b575f80fd5b6200f8356200f298565b6200f843602083016200f199565b81526040820151602082015260608201516200f85f816200f060565b6040820152608082015167ffffffffffffffff8111156200f87e575f80fd5b6200f88f8b6020838601016200f344565b606083015250845250602092830192016200f7f5565b509695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f80604083850312156200f8ef575f80fd5b6200f8fa836200f199565b9150602083015167ffffffffffffffff8111156200f916575f80fd5b6200f924858286016200f344565b9150509250929050565b604081525f6200f94260408301856200ecd3565b82810360208401526200a8f081856200ecd3565b5f602082840312156200f967575f80fd5b815167ffffffffffffffff8111156200f97e575f80fd5b8201601f810184136200f98f575f80fd5b80516200f9a06200f505826200f495565b8082825260208201915060208360051b8501019250868311156200f9c2575f80fd5b602084015b838110156200f8a557805167ffffffffffffffff8111156200f9e7575f80fd5b850160e0818a03601f190112156200f9fd575f80fd5b6200fa076200f2c4565b6200fa15602083016200f199565b81526200fa25604083016200f075565b6020820152606082015167ffffffffffffffff8111156200fa44575f80fd5b6200fa558b6020838601016200f344565b6040830152506080820151606082015260a082015167ffffffffffffffff8111156200fa7f575f80fd5b6200fa908b6020838601016200f344565b60808301525060c082015167ffffffffffffffff8111156200fab0575f80fd5b6200fac18b6020838601016200f344565b60a0830152506200fad560e083016200f3bd565b60c0820152845250602092830192016200f9c7565b5f602082840312156200fafb575f80fd5b815167ffffffffffffffff8111156200fb12575f80fd5b8201601f810184136200fb23575f80fd5b80516200fb346200f505826200f495565b8082825260208201915060208360051b8501019250868311156200fb56575f80fd5b602084015b838110156200f8a557805167ffffffffffffffff8111156200fb7b575f80fd5b85016080818a03601f190112156200fb91575f80fd5b6200fb9b6200f298565b6200fba9602083016200f199565b8152604082015167ffffffffffffffff8111156200fbc5575f80fd5b6200fbd68b6020838601016200f344565b602083015250606082015167ffffffffffffffff8111156200fbf6575f80fd5b6200fc078b6020838601016200f344565b6040830152506080919091015160608201528352602092830192016200fb5b565b8215158152604060208201525f6200c32160408301846200ecd3565b604081525f6200fc5860408301856200ecd3565b90508260208301529392505050565b6001600160a01b0383168152604060208201525f6200c32160408301846200ecd3565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f6200fcbd601a8301856200f73d565b7f3a2000000000000000000000000000000000000000000000000000000000000081526200a8f060028201856200f73d565b5f6200c3216200fd0083866200f73d565b846200f73d565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081525f6200fd3a601a8301856200f73d565b7f207573696e6720636f6e7374727563746f72206461746120220000000000000081526200fd6c60198201856200f73d565b7f2200000000000000000000000000000000000000000000000000000000000000815260010195945050505050565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201525f6200a8f460808301846200ecd3565b5f6200fdf082876200f73d565b7f2f0000000000000000000000000000000000000000000000000000000000000081526200fe2260018201876200f73d565b90507f2f0000000000000000000000000000000000000000000000000000000000000081526200fe5660018201866200f73d565b90507f2f0000000000000000000000000000000000000000000000000000000000000081526200fe8a60018201856200f73d565b7f2e6a736f6e0000000000000000000000000000000000000000000000000000008152600501979650505050505050565b604081525f6200fecf60408301846200ecd3565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081525f6200ff41601f8301846200f73d565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f81527f6d6c00000000000000000000000000000000000000000000000000000000000060208201526022019392505050565b604081525f6200ffa860408301846200ecd3565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b604081525f6200fffb60408301846200ecd3565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081525f6201006d60148301846200f73d565b7f275d2e6b656363616b32353600000000000000000000000000000000000000008152600c019392505050565b7f220000000000000000000000000000000000000000000000000000000000000081525f620100cd60018301846200f73d565b7f220000000000000000000000000000000000000000000000000000000000000081526001019392505050565b5f6201010782846200f73d565b7f2f6275696c642d696e666f0000000000000000000000000000000000000000008152600b019392505050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201525f6200a8f4604b8301846200f73d565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f60ff821660ff8103620101f857620101f8620101b3565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201525f6200a8f460298301846200f73d565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201525f6200a8f460808301846200ecd3565b5f60208284031215620102b3575f80fd5b815167ffffffffffffffff811115620102ca575f80fd5b820160608185031215620102dc575f80fd5b620102e66200f2ea565b81518060030b8114620102f7575f80fd5b8152602082015167ffffffffffffffff81111562010313575f80fd5b62010321868285016200f344565b602083015250604082015167ffffffffffffffff81111562010341575f80fd5b6201034f868285016200f344565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201525f620103b660218301846200f73d565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657481527f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960208201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560408201527f616c69666965642070617468206f66207468652062617368206578656375746160608201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960808201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60a08201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960c08201527f6f75722070726f6a65637420287573696e6720666f727761726420736c61736860e08201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101008201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061012082015261013b019392505050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201525f6201059560218301856200f73d565b7f2720696e206f75747075743a200000000000000000000000000000000000000081526200a8f0600d8201856200f73d565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201525f6200a8f460298301846200f73d565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201525f6200a8f460228301846200f73d565b818103818111156200a483576200a483620101b3565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081525f620106c2600e8301846200f73d565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e7472616381527f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e7360208201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e747260408201527f6163742e6a736f6e00000000000000000000000000000000000000000000000060608201526068019392505050565b808201808211156200a483576200a483620101b3565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081525f620107aa60188301856200f73d565b7f20696e20000000000000000000000000000000000000000000000000000000008152620107dc60048201856200f73d565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f727465642081527f6c6963656e736520666f7220626c6f636b206578706c6f72657220766572696660208201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f60408201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c60608201527f206f7220736574207468652060736b69704c6963656e73655479706560206f7060808201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060a082015260b70195945050505050565b5f620108d682846200f73d565b7f200000000000000000000000000000000000000000000000000000000000000081526001019392505050565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081525f6200a8f4601c8301846200f73d565b5f5f1982036201094a576201094a620101b3565b5060010190565b80820281158282048414176200a483576200a483620101b3565b6001815b6001841115620109ac578085048111156201098e576201098e620101b3565b60018416156201099d57908102905b60019390931c9280026201096f565b935093915050565b5f82620109c4575060016200a483565b81620109d257505f6200a483565b8160018114620109eb5760028114620109f65762010a16565b60019150506200a483565b60ff84111562010a0a5762010a0a620101b3565b50506001821b6200a483565b5060208310610133831016604e8410600b841016171562010a3b575081810a6200a483565b62010a495f1984846201096b565b805f190482111562010a5f5762010a5f620101b3565b029392505050565b5f6200a8f48383620109b4565b5f8162010a855762010a85620101b3565b505f190190565b5f62010a9982856200f73d565b7f3a0000000000000000000000000000000000000000000000000000000000000081526200a8f060018201856200f73d565b8181035f8312801583831316838312821617156200dd71576200dd71620101b356fe6080604052348015600e575f80fd5b5061052f8061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c806306cb8983146100385780632722feee1461004d575b5f80fd5b61004b61004636600461019a565b610091565b005b61006873735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b7ffa84a91b0ed9555afae4459021634264ec770c42989afa595d13944058e229e58686868686866040516100ca96959493929190610440565b60405180910390a1505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461012a575f80fd5b919050565b5f8083601f84011261013f575f80fd5b50813567ffffffffffffffff811115610156575f80fd5b60208301915083602082850101111561016d575f80fd5b9250929050565b5f60408284031215610184575f80fd5b50919050565b5f60a08284031215610184575f80fd5b5f805f805f8060c087890312156101af575f80fd5b863567ffffffffffffffff8111156101c5575f80fd5b8701601f810189136101d5575f80fd5b803567ffffffffffffffff8111156101ef576101ef6100da565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561025b5761025b6100da565b6040528181528282016020018b1015610272575f80fd5b816020840160208301375f6020838301015280985050505061029660208801610107565b9450604087013567ffffffffffffffff8111156102b1575f80fd5b6102bd89828a0161012f565b90955093506102d190508860608901610174565b915060a087013567ffffffffffffffff8111156102ec575f80fd5b6102f889828a0161018a565b9150509295509295509295565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8035801515811461012a575f80fd5b73ffffffffffffffffffffffffffffffffffffffff61037982610107565b1682526103886020820161034c565b1515602083015273ffffffffffffffffffffffffffffffffffffffff6103b060408301610107565b1660408301525f60608201357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18336030181126103eb575f80fd5b820160208101903567ffffffffffffffff811115610407575f80fd5b803603821315610415575f80fd5b60a0606086015261042a60a086018284610305565b6080948501359590940194909452509092915050565b60c081525f87518060c08401528060208a0160e085015e5f60e082850101527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8201168301905073ffffffffffffffffffffffffffffffffffffffff8816602084015260e08382030160408401526104be60e082018789610305565b8535606085015290506104d36020860161034c565b1515608084015282810360a08401526104ec818561035b565b999850505050505050505056fea2646970667358221220b4e3c494e4629c49be1aa581aadf6c83fb63d55d5741640647967e8906ba333d64736f6c634300081a003360c060405234801561000f575f80fd5b5060405161102e38038061102e83398101604081905261002e916100d8565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006257604051632b2add3d60e01b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0385811691909117909155828116608052811660a0526040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac5905f90a1505050610118565b80516001600160a01b03811681146100d3575f80fd5b919050565b5f805f606084860312156100ea575f80fd5b6100f3846100bd565b9250610101602085016100bd565b915061010f604085016100bd565b90509250925092565b60805160a051610eee6101405f395f6101dd01525f81816102b001526104510152610eee5ff3fe608060405234801561000f575f80fd5b50600436106100f0575f3560e01c806397770dff11610093578063c63585cc11610063578063c63585cc1461026b578063d7fd7afb1461027e578063d936a012146102ab578063ee2815ba146102d2575f80fd5b806397770dff14610212578063a7cb050714610225578063c39aca3714610238578063c62178ac1461024b575f80fd5b8063513a9c05116100ce578063513a9c0514610183578063569541b9146101b8578063842da36d146101d857806391dd645f146101ff575f80fd5b80630be15547146100f45780631f0e251b146101535780633ce4a5bc14610168575b5f80fd5b610129610102366004610bb9565b60016020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b610166610161366004610bf8565b6102e5565b005b61012973735b14bb79463307aacbed86daf3322b1e6226ab81565b610129610191366004610bb9565b60026020525f908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101299073ffffffffffffffffffffffffffffffffffffffff1681565b6101297f000000000000000000000000000000000000000000000000000000000000000081565b61016661020d366004610c18565b6103f9565b610166610220366004610bf8565b61051b565b610166610233366004610c42565b610628565b610166610246366004610c62565b6106c2565b6004546101299073ffffffffffffffffffffffffffffffffffffffff1681565b610129610279366004610d28565b6108b9565b61029d61028c366004610bb9565b5f6020819052908152604090205481565b60405190815260200161014a565b6101297f000000000000000000000000000000000000000000000000000000000000000081565b6101666102e0366004610c18565b6109ec565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610332576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811661037f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610446576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003545f9061048d907f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff16846108b9565b5f8481526002602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251878152918201529192507f0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e910160405180910390a1505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610568576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166105b5576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e906020016103ee565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610675576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461070f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831673735b14bb79463307aacbed86daf3322b1e6226ab148061075c575073ffffffffffffffffffffffffffffffffffffffff831630145b15610793576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690528616906347e7ef24906044016020604051808303815f875af1158015610805573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108299190610d68565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063de43156e906108849089908990899088908890600401610dce565b5f604051808303815f87803b15801561089b575f80fd5b505af11580156108ad573d5f803e3d5ffd5b50505050505050505050565b5f805f6108c68585610abc565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b16603482015291935091508690604801604051602081830303815290604052805190602001206040516020016109ac9291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a39576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d91016106b6565b5f808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610b23576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610b5d578284610b60565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610bb2576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b5f60208284031215610bc9575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610bf3575f80fd5b919050565b5f60208284031215610c08575f80fd5b610c1182610bd0565b9392505050565b5f8060408385031215610c29575f80fd5b82359150610c3960208401610bd0565b90509250929050565b5f8060408385031215610c53575f80fd5b50508035926020909101359150565b5f805f805f8060a08789031215610c77575f80fd5b863567ffffffffffffffff811115610c8d575f80fd5b87016060818a031215610c9e575f80fd5b9550610cac60208801610bd0565b945060408701359350610cc160608801610bd0565b9250608087013567ffffffffffffffff811115610cdc575f80fd5b8701601f81018913610cec575f80fd5b803567ffffffffffffffff811115610d02575f80fd5b896020828401011115610d13575f80fd5b60208201935080925050509295509295509295565b5f805f60608486031215610d3a575f80fd5b610d4384610bd0565b9250610d5160208501610bd0565b9150610d5f60408501610bd0565b90509250925092565b5f60208284031215610d78575f80fd5b81518015158114610c11575f80fd5b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b608081525f86357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018112610e04575f80fd5b870160208101903567ffffffffffffffff811115610e20575f80fd5b803603821315610e2e575f80fd5b60606080850152610e4360e085018284610d87565b91505073ffffffffffffffffffffffffffffffffffffffff610e6760208a01610bd0565b1660a0840152604088013560c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401528281036060840152610eac818587610d87565b9897505050505050505056fea2646970667358221220390f960888a4ebf4b1e6cbfae1a45754fc9f10c2947021b59d737acf45cd622864736f6c634300081a003360c060405234801561000f575f80fd5b50604051611fc0380380611fc083398101604081905261002e916101d0565b6001600160a01b038216158061004b57506001600160a01b038116155b156100695760405163d92e233d60e01b815260040160405180910390fd5b60066100758982610315565b5060076100828882610315565b506008805460ff191660ff881617905560808590528360028111156100a9576100a96103cf565b60a08160028111156100bd576100bd6103cf565b9052506001929092555f80546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506103e39350505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f83011261012d575f80fd5b81516001600160401b038111156101465761014661010a565b604051601f8201601f19908116603f011681016001600160401b03811182821017156101745761017461010a565b60405281815283820160200185101561018b575f80fd5b8160208501602083015e5f918101602001919091529392505050565b8051600381106101b5575f80fd5b919050565b80516001600160a01b03811681146101b5575f80fd5b5f805f805f805f80610100898b0312156101e8575f80fd5b88516001600160401b038111156101fd575f80fd5b6102098b828c0161011e565b60208b015190995090506001600160401b03811115610226575f80fd5b6102328b828c0161011e565b975050604089015160ff81168114610248575f80fd5b60608a0151909650945061025e60808a016101a7565b60a08a0151909450925061027460c08a016101ba565b915061028260e08a016101ba565b90509295985092959890939650565b600181811c908216806102a557607f821691505b6020821081036102c357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561031057805f5260205f20601f840160051c810160208510156102ee5750805b601f840160051c820191505b8181101561030d575f81556001016102fa565b50505b505050565b81516001600160401b0381111561032e5761032e61010a565b6103428161033c8454610291565b846102c9565b6020601f821160018114610374575f831561035d5750848201515b5f19600385901b1c1916600184901b17845561030d565b5f84815260208120601f198516915b828110156103a35787850151825560209485019460019092019101610383565b50848210156103c057868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52602160045260245ffd5b60805160a051611ba06104205f395f61033901525f81816102e501528181610bbf01528181610cc201528181610ed90152610fdc0152611ba05ff3fe608060405234801561000f575f80fd5b50600436106101b0575f3560e01c806395d89b41116100f3578063ccc7759911610093578063eddeb1231161006e578063eddeb12314610455578063f2441b3214610468578063f687d12a14610487578063fc5fecd51461049a575f80fd5b8063ccc77599146103c9578063d9eeebed146103dc578063dd62ed3e14610410575f80fd5b8063b84c8246116100ce578063b84c82461461037b578063c47f002714610390578063c7012626146103a3578063c835d7cc146103b6575f80fd5b806395d89b411461032c578063a3413d0314610334578063a9059cbb14610368575f80fd5b80633ce4a5bc1161015e5780634d8943bb116101395780634d8943bb146102a257806370a08231146102ab57806385e1f4d0146102e05780638b851b9514610307575f80fd5b80633ce4a5bc1461023c57806342966c681461027c57806347e7ef241461028f575f80fd5b806318160ddd1161018e57806318160ddd1461020c57806323b872dd14610214578063313ce56714610227575f80fd5b806306fdde03146101b4578063091d2788146101d2578063095ea7b3146101e9575b5f80fd5b6101bc6104ad565b6040516101c991906115fb565b60405180910390f35b6101db60015481565b6040519081526020016101c9565b6101fc6101f7366004611638565b61053d565b60405190151581526020016101c9565b6005546101db565b6101fc610222366004611662565b610553565b60085460405160ff90911681526020016101c9565b61025773735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c9565b6101fc61028a3660046116a0565b6105e8565b6101fc61029d366004611638565b6105fb565b6101db60025481565b6101db6102b93660046116b7565b73ffffffffffffffffffffffffffffffffffffffff165f9081526003602052604090205490565b6101db7f000000000000000000000000000000000000000000000000000000000000000081565b60085461025790610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101bc610752565b61035b7f000000000000000000000000000000000000000000000000000000000000000081565b6040516101c991906116d2565b6101fc610376366004611638565b610761565b61038e6103893660046117d3565b61076d565b005b61038e61039e3660046117d3565b6107ca565b6101fc6103b1366004611820565b610823565b61038e6103c43660046116b7565b61096d565b61038e6103d73660046116b7565b610a80565b6103e4610b94565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101c9565b6101db61041e366004611875565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260046020908152604080832093909416825291909152205490565b61038e6104633660046116a0565b610daa565b5f546102579073ffffffffffffffffffffffffffffffffffffffff1681565b61038e6104953660046116a0565b610e2c565b6103e46104a83660046116a0565b610eae565b6060600680546104bc906118ac565b80601f01602080910402602001604051908101604052809291908181526020018280546104e8906118ac565b80156105335780601f1061050a57610100808354040283529160200191610533565b820191905f5260205f20905b81548152906001019060200180831161051657829003601f168201915b5050505050905090565b5f6105493384846110c2565b5060015b92915050565b5f61055f8484846111ca565b73ffffffffffffffffffffffffffffffffffffffff84165f908152600460209081526040808320338452909152902054828110156105c9576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105dd85336105d8868561192a565b6110c2565b506001949350505050565b5f6105f33383611383565b506001919050565b5f3373735b14bb79463307aacbed86daf3322b1e6226ab1480159061063757505f5473ffffffffffffffffffffffffffffffffffffffff163314155b80156106605750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b15610697576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106a183836114c2565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261074191869061193d565b60405180910390a250600192915050565b6060600780546104bc906118ac565b5f6105493384846111ca565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107ba576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076107c682826119aa565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610817576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60066107c682826119aa565b5f805f61082e610b94565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303815f875af11580156108bd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108e19190611ac1565b610917576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109213385611383565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161095a91899189918791611ae0565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109ba576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a07576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610acd576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b1a576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610a75565b5f80546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610c24573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c489190611b0e565b905073ffffffffffffffffffffffffffffffffffffffff8116610c97576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610d23573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d479190611b29565b9050805f03610d82576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60025460015483610d949190611b40565b610d9e9190611b57565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610df7576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610a75565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e79576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610a75565b5f80546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610f3e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f629190611b0e565b905073ffffffffffffffffffffffffffffffffffffffff8116610fb1576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa15801561103d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110619190611b29565b9050805f0361109c576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002545f906110ab8784611b40565b6110b59190611b57565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff831661110f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821661115c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8381165f8181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611217576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611264576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f90815260036020526040902054818110156112c3576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112cd828261192a565b73ffffffffffffffffffffffffffffffffffffffff8086165f90815260036020526040808220939093559085168152908120805484929061130f908490611b57565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161137591815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166113d0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f908152600360205260409020548181101561142f576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611439828261192a565b73ffffffffffffffffffffffffffffffffffffffff84165f908152600360205260408120919091556005805484929061147390849061192a565b90915550506040518281525f9073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016111bd565b73ffffffffffffffffffffffffffffffffffffffff821661150f576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060055f8282546115209190611b57565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f9081526003602052604081208054839290611559908490611b57565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316905f907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61160d60208301846115af565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff81168114611635575f80fd5b50565b5f8060408385031215611649575f80fd5b823561165481611614565b946020939093013593505050565b5f805f60608486031215611674575f80fd5b833561167f81611614565b9250602084013561168f81611614565b929592945050506040919091013590565b5f602082840312156116b0575f80fd5b5035919050565b5f602082840312156116c7575f80fd5b813561160d81611614565b602081016003831061170b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b91905290565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8067ffffffffffffffff84111561175857611758611711565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156117a5576117a5611711565b6040528381529050808284018510156117bc575f80fd5b838360208301375f60208583010152509392505050565b5f602082840312156117e3575f80fd5b813567ffffffffffffffff8111156117f9575f80fd5b8201601f81018413611809575f80fd5b6118188482356020840161173e565b949350505050565b5f8060408385031215611831575f80fd5b823567ffffffffffffffff811115611847575f80fd5b8301601f81018513611857575f80fd5b6118668582356020840161173e565b95602094909401359450505050565b5f8060408385031215611886575f80fd5b823561189181611614565b915060208301356118a181611614565b809150509250929050565b600181811c908216806118c057607f821691505b6020821081036118f7577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561054d5761054d6118fd565b604081525f61194f60408301856115af565b90508260208301529392505050565b601f8211156119a557805f5260205f20601f840160051c810160208510156119835750805b601f840160051c820191505b818110156119a2575f815560010161198f565b50505b505050565b815167ffffffffffffffff8111156119c4576119c4611711565b6119d8816119d284546118ac565b8461195e565b6020601f821160018114611a29575f83156119f35750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556119a2565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611a765787850151825560209485019460019092019101611a56565b5084821015611ab257868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b5f60208284031215611ad1575f80fd5b8151801515811461160d575f80fd5b608081525f611af260808301876115af565b6020830195909552506040810192909252606090910152919050565b5f60208284031215611b1e575f80fd5b815161160d81611614565b5f60208284031215611b39575f80fd5b5051919050565b808202811582820484141761054d5761054d6118fd565b8082018082111561054d5761054d6118fd56fea26469706673582212206587bf9078f059498405d04f85296ef5da57d060c91cbc6dd678aa79dbaee6e064736f6c634300081a00334661696c656420746f2066696e6420696e616374697665206761746577617920636f6e74726163744661696c656420746f2066696e6420636f6e6e6563746f7220636f6e7472616374a264697066735822122057db8f1f7beadd9cb26ec4f47dc460008e466bfa259d04b4f2b6ad15bfaae1b664736f6c634300081a0033",
}

// CoreRegistryTestABI is the input ABI used to generate the binding from.
// Deprecated: Use CoreRegistryTestMetaData.ABI instead.
var CoreRegistryTestABI = CoreRegistryTestMetaData.ABI

// CoreRegistryTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CoreRegistryTestMetaData.Bin instead.
var CoreRegistryTestBin = CoreRegistryTestMetaData.Bin

// DeployCoreRegistryTest deploys a new Ethereum contract, binding an instance of CoreRegistryTest to it.
func DeployCoreRegistryTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CoreRegistryTest, error) {
	parsed, err := CoreRegistryTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CoreRegistryTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CoreRegistryTest{CoreRegistryTestCaller: CoreRegistryTestCaller{contract: contract}, CoreRegistryTestTransactor: CoreRegistryTestTransactor{contract: contract}, CoreRegistryTestFilterer: CoreRegistryTestFilterer{contract: contract}}, nil
}

// CoreRegistryTest is an auto generated Go binding around an Ethereum contract.
type CoreRegistryTest struct {
	CoreRegistryTestCaller     // Read-only binding to the contract
	CoreRegistryTestTransactor // Write-only binding to the contract
	CoreRegistryTestFilterer   // Log filterer for contract events
}

// CoreRegistryTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoreRegistryTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreRegistryTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoreRegistryTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreRegistryTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoreRegistryTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoreRegistryTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoreRegistryTestSession struct {
	Contract     *CoreRegistryTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoreRegistryTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoreRegistryTestCallerSession struct {
	Contract *CoreRegistryTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CoreRegistryTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoreRegistryTestTransactorSession struct {
	Contract     *CoreRegistryTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CoreRegistryTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoreRegistryTestRaw struct {
	Contract *CoreRegistryTest // Generic contract binding to access the raw methods on
}

// CoreRegistryTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoreRegistryTestCallerRaw struct {
	Contract *CoreRegistryTestCaller // Generic read-only contract binding to access the raw methods on
}

// CoreRegistryTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoreRegistryTestTransactorRaw struct {
	Contract *CoreRegistryTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoreRegistryTest creates a new instance of CoreRegistryTest, bound to a specific deployed contract.
func NewCoreRegistryTest(address common.Address, backend bind.ContractBackend) (*CoreRegistryTest, error) {
	contract, err := bindCoreRegistryTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTest{CoreRegistryTestCaller: CoreRegistryTestCaller{contract: contract}, CoreRegistryTestTransactor: CoreRegistryTestTransactor{contract: contract}, CoreRegistryTestFilterer: CoreRegistryTestFilterer{contract: contract}}, nil
}

// NewCoreRegistryTestCaller creates a new read-only instance of CoreRegistryTest, bound to a specific deployed contract.
func NewCoreRegistryTestCaller(address common.Address, caller bind.ContractCaller) (*CoreRegistryTestCaller, error) {
	contract, err := bindCoreRegistryTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestCaller{contract: contract}, nil
}

// NewCoreRegistryTestTransactor creates a new write-only instance of CoreRegistryTest, bound to a specific deployed contract.
func NewCoreRegistryTestTransactor(address common.Address, transactor bind.ContractTransactor) (*CoreRegistryTestTransactor, error) {
	contract, err := bindCoreRegistryTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestTransactor{contract: contract}, nil
}

// NewCoreRegistryTestFilterer creates a new log filterer instance of CoreRegistryTest, bound to a specific deployed contract.
func NewCoreRegistryTestFilterer(address common.Address, filterer bind.ContractFilterer) (*CoreRegistryTestFilterer, error) {
	contract, err := bindCoreRegistryTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestFilterer{contract: contract}, nil
}

// bindCoreRegistryTest binds a generic wrapper to an already deployed contract.
func bindCoreRegistryTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CoreRegistryTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoreRegistryTest *CoreRegistryTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoreRegistryTest.Contract.CoreRegistryTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoreRegistryTest *CoreRegistryTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.CoreRegistryTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoreRegistryTest *CoreRegistryTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.CoreRegistryTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoreRegistryTest *CoreRegistryTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CoreRegistryTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoreRegistryTest *CoreRegistryTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoreRegistryTest *CoreRegistryTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.contract.Transact(opts, method, params...)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_CoreRegistryTest *CoreRegistryTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CoreRegistryTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_CoreRegistryTest *CoreRegistryTestSession) ISTEST() (bool, error) {
	return _CoreRegistryTest.Contract.ISTEST(&_CoreRegistryTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_CoreRegistryTest *CoreRegistryTestCallerSession) ISTEST() (bool, error) {
	return _CoreRegistryTest.Contract.ISTEST(&_CoreRegistryTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_CoreRegistryTest *CoreRegistryTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _CoreRegistryTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_CoreRegistryTest *CoreRegistryTestSession) ExcludeArtifacts() ([]string, error) {
	return _CoreRegistryTest.Contract.ExcludeArtifacts(&_CoreRegistryTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_CoreRegistryTest *CoreRegistryTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _CoreRegistryTest.Contract.ExcludeArtifacts(&_CoreRegistryTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_CoreRegistryTest *CoreRegistryTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CoreRegistryTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_CoreRegistryTest *CoreRegistryTestSession) ExcludeContracts() ([]common.Address, error) {
	return _CoreRegistryTest.Contract.ExcludeContracts(&_CoreRegistryTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_CoreRegistryTest *CoreRegistryTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _CoreRegistryTest.Contract.ExcludeContracts(&_CoreRegistryTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_CoreRegistryTest *CoreRegistryTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _CoreRegistryTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_CoreRegistryTest *CoreRegistryTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _CoreRegistryTest.Contract.ExcludeSelectors(&_CoreRegistryTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_CoreRegistryTest *CoreRegistryTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _CoreRegistryTest.Contract.ExcludeSelectors(&_CoreRegistryTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_CoreRegistryTest *CoreRegistryTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CoreRegistryTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_CoreRegistryTest *CoreRegistryTestSession) ExcludeSenders() ([]common.Address, error) {
	return _CoreRegistryTest.Contract.ExcludeSenders(&_CoreRegistryTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_CoreRegistryTest *CoreRegistryTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _CoreRegistryTest.Contract.ExcludeSenders(&_CoreRegistryTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_CoreRegistryTest *CoreRegistryTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CoreRegistryTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_CoreRegistryTest *CoreRegistryTestSession) Failed() (bool, error) {
	return _CoreRegistryTest.Contract.Failed(&_CoreRegistryTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_CoreRegistryTest *CoreRegistryTestCallerSession) Failed() (bool, error) {
	return _CoreRegistryTest.Contract.Failed(&_CoreRegistryTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_CoreRegistryTest *CoreRegistryTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _CoreRegistryTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_CoreRegistryTest *CoreRegistryTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _CoreRegistryTest.Contract.TargetArtifactSelectors(&_CoreRegistryTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_CoreRegistryTest *CoreRegistryTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _CoreRegistryTest.Contract.TargetArtifactSelectors(&_CoreRegistryTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_CoreRegistryTest *CoreRegistryTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _CoreRegistryTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_CoreRegistryTest *CoreRegistryTestSession) TargetArtifacts() ([]string, error) {
	return _CoreRegistryTest.Contract.TargetArtifacts(&_CoreRegistryTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_CoreRegistryTest *CoreRegistryTestCallerSession) TargetArtifacts() ([]string, error) {
	return _CoreRegistryTest.Contract.TargetArtifacts(&_CoreRegistryTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_CoreRegistryTest *CoreRegistryTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CoreRegistryTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_CoreRegistryTest *CoreRegistryTestSession) TargetContracts() ([]common.Address, error) {
	return _CoreRegistryTest.Contract.TargetContracts(&_CoreRegistryTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_CoreRegistryTest *CoreRegistryTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _CoreRegistryTest.Contract.TargetContracts(&_CoreRegistryTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_CoreRegistryTest *CoreRegistryTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _CoreRegistryTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_CoreRegistryTest *CoreRegistryTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _CoreRegistryTest.Contract.TargetInterfaces(&_CoreRegistryTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_CoreRegistryTest *CoreRegistryTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _CoreRegistryTest.Contract.TargetInterfaces(&_CoreRegistryTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_CoreRegistryTest *CoreRegistryTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _CoreRegistryTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_CoreRegistryTest *CoreRegistryTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _CoreRegistryTest.Contract.TargetSelectors(&_CoreRegistryTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_CoreRegistryTest *CoreRegistryTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _CoreRegistryTest.Contract.TargetSelectors(&_CoreRegistryTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_CoreRegistryTest *CoreRegistryTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CoreRegistryTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_CoreRegistryTest *CoreRegistryTestSession) TargetSenders() ([]common.Address, error) {
	return _CoreRegistryTest.Contract.TargetSenders(&_CoreRegistryTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_CoreRegistryTest *CoreRegistryTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _CoreRegistryTest.Contract.TargetSenders(&_CoreRegistryTest.CallOpts)
}

// TestInitialize is a free data retrieval call binding the contract method 0x993831b6.
//
// Solidity: function testInitialize() view returns()
func (_CoreRegistryTest *CoreRegistryTestCaller) TestInitialize(opts *bind.CallOpts) error {
	var out []interface{}
	err := _CoreRegistryTest.contract.Call(opts, &out, "testInitialize")

	if err != nil {
		return err
	}

	return err

}

// TestInitialize is a free data retrieval call binding the contract method 0x993831b6.
//
// Solidity: function testInitialize() view returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestInitialize() error {
	return _CoreRegistryTest.Contract.TestInitialize(&_CoreRegistryTest.CallOpts)
}

// TestInitialize is a free data retrieval call binding the contract method 0x993831b6.
//
// Solidity: function testInitialize() view returns()
func (_CoreRegistryTest *CoreRegistryTestCallerSession) TestInitialize() error {
	return _CoreRegistryTest.Contract.TestInitialize(&_CoreRegistryTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) SetUp() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.SetUp(&_CoreRegistryTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.SetUp(&_CoreRegistryTest.TransactOpts)
}

// TestActivateAlreadyActiveChain is a paid mutator transaction binding the contract method 0x7d706cef.
//
// Solidity: function testActivateAlreadyActiveChain() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestActivateAlreadyActiveChain(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testActivateAlreadyActiveChain")
}

// TestActivateAlreadyActiveChain is a paid mutator transaction binding the contract method 0x7d706cef.
//
// Solidity: function testActivateAlreadyActiveChain() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestActivateAlreadyActiveChain() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestActivateAlreadyActiveChain(&_CoreRegistryTest.TransactOpts)
}

// TestActivateAlreadyActiveChain is a paid mutator transaction binding the contract method 0x7d706cef.
//
// Solidity: function testActivateAlreadyActiveChain() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestActivateAlreadyActiveChain() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestActivateAlreadyActiveChain(&_CoreRegistryTest.TransactOpts)
}

// TestActivateChain is a paid mutator transaction binding the contract method 0xcff8f111.
//
// Solidity: function testActivateChain() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestActivateChain(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testActivateChain")
}

// TestActivateChain is a paid mutator transaction binding the contract method 0xcff8f111.
//
// Solidity: function testActivateChain() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestActivateChain() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestActivateChain(&_CoreRegistryTest.TransactOpts)
}

// TestActivateChain is a paid mutator transaction binding the contract method 0xcff8f111.
//
// Solidity: function testActivateChain() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestActivateChain() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestActivateChain(&_CoreRegistryTest.TransactOpts)
}

// TestActivateChainUnauthorized is a paid mutator transaction binding the contract method 0x173d9590.
//
// Solidity: function testActivateChainUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestActivateChainUnauthorized(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testActivateChainUnauthorized")
}

// TestActivateChainUnauthorized is a paid mutator transaction binding the contract method 0x173d9590.
//
// Solidity: function testActivateChainUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestActivateChainUnauthorized() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestActivateChainUnauthorized(&_CoreRegistryTest.TransactOpts)
}

// TestActivateChainUnauthorized is a paid mutator transaction binding the contract method 0x173d9590.
//
// Solidity: function testActivateChainUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestActivateChainUnauthorized() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestActivateChainUnauthorized(&_CoreRegistryTest.TransactOpts)
}

// TestChangeAdmin is a paid mutator transaction binding the contract method 0x89328258.
//
// Solidity: function testChangeAdmin() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestChangeAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testChangeAdmin")
}

// TestChangeAdmin is a paid mutator transaction binding the contract method 0x89328258.
//
// Solidity: function testChangeAdmin() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestChangeAdmin() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestChangeAdmin(&_CoreRegistryTest.TransactOpts)
}

// TestChangeAdmin is a paid mutator transaction binding the contract method 0x89328258.
//
// Solidity: function testChangeAdmin() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestChangeAdmin() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestChangeAdmin(&_CoreRegistryTest.TransactOpts)
}

// TestChangeAdminUnauthorized is a paid mutator transaction binding the contract method 0x599a9da3.
//
// Solidity: function testChangeAdminUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestChangeAdminUnauthorized(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testChangeAdminUnauthorized")
}

// TestChangeAdminUnauthorized is a paid mutator transaction binding the contract method 0x599a9da3.
//
// Solidity: function testChangeAdminUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestChangeAdminUnauthorized() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestChangeAdminUnauthorized(&_CoreRegistryTest.TransactOpts)
}

// TestChangeAdminUnauthorized is a paid mutator transaction binding the contract method 0x599a9da3.
//
// Solidity: function testChangeAdminUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestChangeAdminUnauthorized() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestChangeAdminUnauthorized(&_CoreRegistryTest.TransactOpts)
}

// TestChangeAdminWithZeroAddress is a paid mutator transaction binding the contract method 0xa5ae07e2.
//
// Solidity: function testChangeAdminWithZeroAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestChangeAdminWithZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testChangeAdminWithZeroAddress")
}

// TestChangeAdminWithZeroAddress is a paid mutator transaction binding the contract method 0xa5ae07e2.
//
// Solidity: function testChangeAdminWithZeroAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestChangeAdminWithZeroAddress() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestChangeAdminWithZeroAddress(&_CoreRegistryTest.TransactOpts)
}

// TestChangeAdminWithZeroAddress is a paid mutator transaction binding the contract method 0xa5ae07e2.
//
// Solidity: function testChangeAdminWithZeroAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestChangeAdminWithZeroAddress() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestChangeAdminWithZeroAddress(&_CoreRegistryTest.TransactOpts)
}

// TestChangeRegistryManager is a paid mutator transaction binding the contract method 0xb1725580.
//
// Solidity: function testChangeRegistryManager() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestChangeRegistryManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testChangeRegistryManager")
}

// TestChangeRegistryManager is a paid mutator transaction binding the contract method 0xb1725580.
//
// Solidity: function testChangeRegistryManager() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestChangeRegistryManager() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestChangeRegistryManager(&_CoreRegistryTest.TransactOpts)
}

// TestChangeRegistryManager is a paid mutator transaction binding the contract method 0xb1725580.
//
// Solidity: function testChangeRegistryManager() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestChangeRegistryManager() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestChangeRegistryManager(&_CoreRegistryTest.TransactOpts)
}

// TestChangeRegistryManagerUnauthorized is a paid mutator transaction binding the contract method 0x36328450.
//
// Solidity: function testChangeRegistryManagerUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestChangeRegistryManagerUnauthorized(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testChangeRegistryManagerUnauthorized")
}

// TestChangeRegistryManagerUnauthorized is a paid mutator transaction binding the contract method 0x36328450.
//
// Solidity: function testChangeRegistryManagerUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestChangeRegistryManagerUnauthorized() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestChangeRegistryManagerUnauthorized(&_CoreRegistryTest.TransactOpts)
}

// TestChangeRegistryManagerUnauthorized is a paid mutator transaction binding the contract method 0x36328450.
//
// Solidity: function testChangeRegistryManagerUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestChangeRegistryManagerUnauthorized() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestChangeRegistryManagerUnauthorized(&_CoreRegistryTest.TransactOpts)
}

// TestChangeRegistryManagerWithZeroAddress is a paid mutator transaction binding the contract method 0x599d1173.
//
// Solidity: function testChangeRegistryManagerWithZeroAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestChangeRegistryManagerWithZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testChangeRegistryManagerWithZeroAddress")
}

// TestChangeRegistryManagerWithZeroAddress is a paid mutator transaction binding the contract method 0x599d1173.
//
// Solidity: function testChangeRegistryManagerWithZeroAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestChangeRegistryManagerWithZeroAddress() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestChangeRegistryManagerWithZeroAddress(&_CoreRegistryTest.TransactOpts)
}

// TestChangeRegistryManagerWithZeroAddress is a paid mutator transaction binding the contract method 0x599d1173.
//
// Solidity: function testChangeRegistryManagerWithZeroAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestChangeRegistryManagerWithZeroAddress() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestChangeRegistryManagerWithZeroAddress(&_CoreRegistryTest.TransactOpts)
}

// TestDeactivateChain is a paid mutator transaction binding the contract method 0x24ffc317.
//
// Solidity: function testDeactivateChain() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestDeactivateChain(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testDeactivateChain")
}

// TestDeactivateChain is a paid mutator transaction binding the contract method 0x24ffc317.
//
// Solidity: function testDeactivateChain() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestDeactivateChain() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestDeactivateChain(&_CoreRegistryTest.TransactOpts)
}

// TestDeactivateChain is a paid mutator transaction binding the contract method 0x24ffc317.
//
// Solidity: function testDeactivateChain() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestDeactivateChain() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestDeactivateChain(&_CoreRegistryTest.TransactOpts)
}

// TestDeactivateNonActiveChain is a paid mutator transaction binding the contract method 0x40a3b50b.
//
// Solidity: function testDeactivateNonActiveChain() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestDeactivateNonActiveChain(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testDeactivateNonActiveChain")
}

// TestDeactivateNonActiveChain is a paid mutator transaction binding the contract method 0x40a3b50b.
//
// Solidity: function testDeactivateNonActiveChain() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestDeactivateNonActiveChain() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestDeactivateNonActiveChain(&_CoreRegistryTest.TransactOpts)
}

// TestDeactivateNonActiveChain is a paid mutator transaction binding the contract method 0x40a3b50b.
//
// Solidity: function testDeactivateNonActiveChain() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestDeactivateNonActiveChain() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestDeactivateNonActiveChain(&_CoreRegistryTest.TransactOpts)
}

// TestGetAllChains is a paid mutator transaction binding the contract method 0x9dc74402.
//
// Solidity: function testGetAllChains() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestGetAllChains(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testGetAllChains")
}

// TestGetAllChains is a paid mutator transaction binding the contract method 0x9dc74402.
//
// Solidity: function testGetAllChains() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestGetAllChains() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestGetAllChains(&_CoreRegistryTest.TransactOpts)
}

// TestGetAllChains is a paid mutator transaction binding the contract method 0x9dc74402.
//
// Solidity: function testGetAllChains() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestGetAllChains() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestGetAllChains(&_CoreRegistryTest.TransactOpts)
}

// TestGetAllContracts is a paid mutator transaction binding the contract method 0xf7146feb.
//
// Solidity: function testGetAllContracts() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestGetAllContracts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testGetAllContracts")
}

// TestGetAllContracts is a paid mutator transaction binding the contract method 0xf7146feb.
//
// Solidity: function testGetAllContracts() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestGetAllContracts() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestGetAllContracts(&_CoreRegistryTest.TransactOpts)
}

// TestGetAllContracts is a paid mutator transaction binding the contract method 0xf7146feb.
//
// Solidity: function testGetAllContracts() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestGetAllContracts() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestGetAllContracts(&_CoreRegistryTest.TransactOpts)
}

// TestGetAllZRC20Tokens is a paid mutator transaction binding the contract method 0xd52fe066.
//
// Solidity: function testGetAllZRC20Tokens() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestGetAllZRC20Tokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testGetAllZRC20Tokens")
}

// TestGetAllZRC20Tokens is a paid mutator transaction binding the contract method 0xd52fe066.
//
// Solidity: function testGetAllZRC20Tokens() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestGetAllZRC20Tokens() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestGetAllZRC20Tokens(&_CoreRegistryTest.TransactOpts)
}

// TestGetAllZRC20Tokens is a paid mutator transaction binding the contract method 0xd52fe066.
//
// Solidity: function testGetAllZRC20Tokens() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestGetAllZRC20Tokens() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestGetAllZRC20Tokens(&_CoreRegistryTest.TransactOpts)
}

// TestInitializeWithZeroAddress is a paid mutator transaction binding the contract method 0x71748c09.
//
// Solidity: function testInitializeWithZeroAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestInitializeWithZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testInitializeWithZeroAddress")
}

// TestInitializeWithZeroAddress is a paid mutator transaction binding the contract method 0x71748c09.
//
// Solidity: function testInitializeWithZeroAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestInitializeWithZeroAddress() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestInitializeWithZeroAddress(&_CoreRegistryTest.TransactOpts)
}

// TestInitializeWithZeroAddress is a paid mutator transaction binding the contract method 0x71748c09.
//
// Solidity: function testInitializeWithZeroAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestInitializeWithZeroAddress() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestInitializeWithZeroAddress(&_CoreRegistryTest.TransactOpts)
}

// TestPause is a paid mutator transaction binding the contract method 0x9bf35597.
//
// Solidity: function testPause() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testPause")
}

// TestPause is a paid mutator transaction binding the contract method 0x9bf35597.
//
// Solidity: function testPause() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestPause() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestPause(&_CoreRegistryTest.TransactOpts)
}

// TestPause is a paid mutator transaction binding the contract method 0x9bf35597.
//
// Solidity: function testPause() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestPause() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestPause(&_CoreRegistryTest.TransactOpts)
}

// TestPauseUnauthorized is a paid mutator transaction binding the contract method 0x151e9be4.
//
// Solidity: function testPauseUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestPauseUnauthorized(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testPauseUnauthorized")
}

// TestPauseUnauthorized is a paid mutator transaction binding the contract method 0x151e9be4.
//
// Solidity: function testPauseUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestPauseUnauthorized() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestPauseUnauthorized(&_CoreRegistryTest.TransactOpts)
}

// TestPauseUnauthorized is a paid mutator transaction binding the contract method 0x151e9be4.
//
// Solidity: function testPauseUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestPauseUnauthorized() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestPauseUnauthorized(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterContract is a paid mutator transaction binding the contract method 0xb39377d8.
//
// Solidity: function testRegisterContract() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestRegisterContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testRegisterContract")
}

// TestRegisterContract is a paid mutator transaction binding the contract method 0xb39377d8.
//
// Solidity: function testRegisterContract() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestRegisterContract() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterContract(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterContract is a paid mutator transaction binding the contract method 0xb39377d8.
//
// Solidity: function testRegisterContract() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestRegisterContract() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterContract(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterContractForNonActiveChain is a paid mutator transaction binding the contract method 0x004e57bd.
//
// Solidity: function testRegisterContractForNonActiveChain() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestRegisterContractForNonActiveChain(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testRegisterContractForNonActiveChain")
}

// TestRegisterContractForNonActiveChain is a paid mutator transaction binding the contract method 0x004e57bd.
//
// Solidity: function testRegisterContractForNonActiveChain() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestRegisterContractForNonActiveChain() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterContractForNonActiveChain(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterContractForNonActiveChain is a paid mutator transaction binding the contract method 0x004e57bd.
//
// Solidity: function testRegisterContractForNonActiveChain() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestRegisterContractForNonActiveChain() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterContractForNonActiveChain(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterContractTwice is a paid mutator transaction binding the contract method 0xf81b656b.
//
// Solidity: function testRegisterContractTwice() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestRegisterContractTwice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testRegisterContractTwice")
}

// TestRegisterContractTwice is a paid mutator transaction binding the contract method 0xf81b656b.
//
// Solidity: function testRegisterContractTwice() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestRegisterContractTwice() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterContractTwice(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterContractTwice is a paid mutator transaction binding the contract method 0xf81b656b.
//
// Solidity: function testRegisterContractTwice() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestRegisterContractTwice() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterContractTwice(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterContractWithEmptyAddress is a paid mutator transaction binding the contract method 0x2013038b.
//
// Solidity: function testRegisterContractWithEmptyAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestRegisterContractWithEmptyAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testRegisterContractWithEmptyAddress")
}

// TestRegisterContractWithEmptyAddress is a paid mutator transaction binding the contract method 0x2013038b.
//
// Solidity: function testRegisterContractWithEmptyAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestRegisterContractWithEmptyAddress() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterContractWithEmptyAddress(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterContractWithEmptyAddress is a paid mutator transaction binding the contract method 0x2013038b.
//
// Solidity: function testRegisterContractWithEmptyAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestRegisterContractWithEmptyAddress() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterContractWithEmptyAddress(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterContractWithEmptyType is a paid mutator transaction binding the contract method 0xd97aa899.
//
// Solidity: function testRegisterContractWithEmptyType() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestRegisterContractWithEmptyType(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testRegisterContractWithEmptyType")
}

// TestRegisterContractWithEmptyType is a paid mutator transaction binding the contract method 0xd97aa899.
//
// Solidity: function testRegisterContractWithEmptyType() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestRegisterContractWithEmptyType() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterContractWithEmptyType(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterContractWithEmptyType is a paid mutator transaction binding the contract method 0xd97aa899.
//
// Solidity: function testRegisterContractWithEmptyType() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestRegisterContractWithEmptyType() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterContractWithEmptyType(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterZRC20Token is a paid mutator transaction binding the contract method 0x8c500b5c.
//
// Solidity: function testRegisterZRC20Token() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestRegisterZRC20Token(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testRegisterZRC20Token")
}

// TestRegisterZRC20Token is a paid mutator transaction binding the contract method 0x8c500b5c.
//
// Solidity: function testRegisterZRC20Token() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestRegisterZRC20Token() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterZRC20Token(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterZRC20Token is a paid mutator transaction binding the contract method 0x8c500b5c.
//
// Solidity: function testRegisterZRC20Token() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestRegisterZRC20Token() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterZRC20Token(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterZRC20TokenTwice is a paid mutator transaction binding the contract method 0x70cd4138.
//
// Solidity: function testRegisterZRC20TokenTwice() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestRegisterZRC20TokenTwice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testRegisterZRC20TokenTwice")
}

// TestRegisterZRC20TokenTwice is a paid mutator transaction binding the contract method 0x70cd4138.
//
// Solidity: function testRegisterZRC20TokenTwice() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestRegisterZRC20TokenTwice() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterZRC20TokenTwice(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterZRC20TokenTwice is a paid mutator transaction binding the contract method 0x70cd4138.
//
// Solidity: function testRegisterZRC20TokenTwice() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestRegisterZRC20TokenTwice() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterZRC20TokenTwice(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterZRC20TokenWithDuplicateSymbol is a paid mutator transaction binding the contract method 0x8c18b40f.
//
// Solidity: function testRegisterZRC20TokenWithDuplicateSymbol() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestRegisterZRC20TokenWithDuplicateSymbol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testRegisterZRC20TokenWithDuplicateSymbol")
}

// TestRegisterZRC20TokenWithDuplicateSymbol is a paid mutator transaction binding the contract method 0x8c18b40f.
//
// Solidity: function testRegisterZRC20TokenWithDuplicateSymbol() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestRegisterZRC20TokenWithDuplicateSymbol() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterZRC20TokenWithDuplicateSymbol(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterZRC20TokenWithDuplicateSymbol is a paid mutator transaction binding the contract method 0x8c18b40f.
//
// Solidity: function testRegisterZRC20TokenWithDuplicateSymbol() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestRegisterZRC20TokenWithDuplicateSymbol() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterZRC20TokenWithDuplicateSymbol(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterZRC20TokenWithEmptyAddress is a paid mutator transaction binding the contract method 0xaa5ec565.
//
// Solidity: function testRegisterZRC20TokenWithEmptyAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestRegisterZRC20TokenWithEmptyAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testRegisterZRC20TokenWithEmptyAddress")
}

// TestRegisterZRC20TokenWithEmptyAddress is a paid mutator transaction binding the contract method 0xaa5ec565.
//
// Solidity: function testRegisterZRC20TokenWithEmptyAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestRegisterZRC20TokenWithEmptyAddress() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterZRC20TokenWithEmptyAddress(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterZRC20TokenWithEmptyAddress is a paid mutator transaction binding the contract method 0xaa5ec565.
//
// Solidity: function testRegisterZRC20TokenWithEmptyAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestRegisterZRC20TokenWithEmptyAddress() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterZRC20TokenWithEmptyAddress(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterZRC20TokenWithEmptyOriginAddress is a paid mutator transaction binding the contract method 0xd312ad7d.
//
// Solidity: function testRegisterZRC20TokenWithEmptyOriginAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestRegisterZRC20TokenWithEmptyOriginAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testRegisterZRC20TokenWithEmptyOriginAddress")
}

// TestRegisterZRC20TokenWithEmptyOriginAddress is a paid mutator transaction binding the contract method 0xd312ad7d.
//
// Solidity: function testRegisterZRC20TokenWithEmptyOriginAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestRegisterZRC20TokenWithEmptyOriginAddress() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterZRC20TokenWithEmptyOriginAddress(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterZRC20TokenWithEmptyOriginAddress is a paid mutator transaction binding the contract method 0xd312ad7d.
//
// Solidity: function testRegisterZRC20TokenWithEmptyOriginAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestRegisterZRC20TokenWithEmptyOriginAddress() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterZRC20TokenWithEmptyOriginAddress(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterZRC20TokenWithEmptySymbol is a paid mutator transaction binding the contract method 0xd6683856.
//
// Solidity: function testRegisterZRC20TokenWithEmptySymbol() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestRegisterZRC20TokenWithEmptySymbol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testRegisterZRC20TokenWithEmptySymbol")
}

// TestRegisterZRC20TokenWithEmptySymbol is a paid mutator transaction binding the contract method 0xd6683856.
//
// Solidity: function testRegisterZRC20TokenWithEmptySymbol() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestRegisterZRC20TokenWithEmptySymbol() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterZRC20TokenWithEmptySymbol(&_CoreRegistryTest.TransactOpts)
}

// TestRegisterZRC20TokenWithEmptySymbol is a paid mutator transaction binding the contract method 0xd6683856.
//
// Solidity: function testRegisterZRC20TokenWithEmptySymbol() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestRegisterZRC20TokenWithEmptySymbol() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestRegisterZRC20TokenWithEmptySymbol(&_CoreRegistryTest.TransactOpts)
}

// TestSetContractActive is a paid mutator transaction binding the contract method 0xfaab466a.
//
// Solidity: function testSetContractActive() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestSetContractActive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testSetContractActive")
}

// TestSetContractActive is a paid mutator transaction binding the contract method 0xfaab466a.
//
// Solidity: function testSetContractActive() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestSetContractActive() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestSetContractActive(&_CoreRegistryTest.TransactOpts)
}

// TestSetContractActive is a paid mutator transaction binding the contract method 0xfaab466a.
//
// Solidity: function testSetContractActive() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestSetContractActive() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestSetContractActive(&_CoreRegistryTest.TransactOpts)
}

// TestSetStatusForNonExistenContract is a paid mutator transaction binding the contract method 0x4a3b5640.
//
// Solidity: function testSetStatusForNonExistenContract() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestSetStatusForNonExistenContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testSetStatusForNonExistenContract")
}

// TestSetStatusForNonExistenContract is a paid mutator transaction binding the contract method 0x4a3b5640.
//
// Solidity: function testSetStatusForNonExistenContract() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestSetStatusForNonExistenContract() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestSetStatusForNonExistenContract(&_CoreRegistryTest.TransactOpts)
}

// TestSetStatusForNonExistenContract is a paid mutator transaction binding the contract method 0x4a3b5640.
//
// Solidity: function testSetStatusForNonExistenContract() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestSetStatusForNonExistenContract() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestSetStatusForNonExistenContract(&_CoreRegistryTest.TransactOpts)
}

// TestSetZRC20TokenActive is a paid mutator transaction binding the contract method 0x24196f32.
//
// Solidity: function testSetZRC20TokenActive() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestSetZRC20TokenActive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testSetZRC20TokenActive")
}

// TestSetZRC20TokenActive is a paid mutator transaction binding the contract method 0x24196f32.
//
// Solidity: function testSetZRC20TokenActive() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestSetZRC20TokenActive() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestSetZRC20TokenActive(&_CoreRegistryTest.TransactOpts)
}

// TestSetZRC20TokenActive is a paid mutator transaction binding the contract method 0x24196f32.
//
// Solidity: function testSetZRC20TokenActive() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestSetZRC20TokenActive() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestSetZRC20TokenActive(&_CoreRegistryTest.TransactOpts)
}

// TestSetZRC20TokenActiveWithZeroAddress is a paid mutator transaction binding the contract method 0x164c5b02.
//
// Solidity: function testSetZRC20TokenActiveWithZeroAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestSetZRC20TokenActiveWithZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testSetZRC20TokenActiveWithZeroAddress")
}

// TestSetZRC20TokenActiveWithZeroAddress is a paid mutator transaction binding the contract method 0x164c5b02.
//
// Solidity: function testSetZRC20TokenActiveWithZeroAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestSetZRC20TokenActiveWithZeroAddress() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestSetZRC20TokenActiveWithZeroAddress(&_CoreRegistryTest.TransactOpts)
}

// TestSetZRC20TokenActiveWithZeroAddress is a paid mutator transaction binding the contract method 0x164c5b02.
//
// Solidity: function testSetZRC20TokenActiveWithZeroAddress() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestSetZRC20TokenActiveWithZeroAddress() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestSetZRC20TokenActiveWithZeroAddress(&_CoreRegistryTest.TransactOpts)
}

// TestUpdateChainMetadata is a paid mutator transaction binding the contract method 0xbc056f7f.
//
// Solidity: function testUpdateChainMetadata() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestUpdateChainMetadata(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testUpdateChainMetadata")
}

// TestUpdateChainMetadata is a paid mutator transaction binding the contract method 0xbc056f7f.
//
// Solidity: function testUpdateChainMetadata() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestUpdateChainMetadata() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestUpdateChainMetadata(&_CoreRegistryTest.TransactOpts)
}

// TestUpdateChainMetadata is a paid mutator transaction binding the contract method 0xbc056f7f.
//
// Solidity: function testUpdateChainMetadata() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestUpdateChainMetadata() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestUpdateChainMetadata(&_CoreRegistryTest.TransactOpts)
}

// TestUpdateConfigForNonExistentContract is a paid mutator transaction binding the contract method 0x39c1f2a4.
//
// Solidity: function testUpdateConfigForNonExistentContract() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestUpdateConfigForNonExistentContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testUpdateConfigForNonExistentContract")
}

// TestUpdateConfigForNonExistentContract is a paid mutator transaction binding the contract method 0x39c1f2a4.
//
// Solidity: function testUpdateConfigForNonExistentContract() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestUpdateConfigForNonExistentContract() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestUpdateConfigForNonExistentContract(&_CoreRegistryTest.TransactOpts)
}

// TestUpdateConfigForNonExistentContract is a paid mutator transaction binding the contract method 0x39c1f2a4.
//
// Solidity: function testUpdateConfigForNonExistentContract() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestUpdateConfigForNonExistentContract() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestUpdateConfigForNonExistentContract(&_CoreRegistryTest.TransactOpts)
}

// TestUpdateContractConfig is a paid mutator transaction binding the contract method 0x674270a0.
//
// Solidity: function testUpdateContractConfig() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestUpdateContractConfig(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testUpdateContractConfig")
}

// TestUpdateContractConfig is a paid mutator transaction binding the contract method 0x674270a0.
//
// Solidity: function testUpdateContractConfig() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestUpdateContractConfig() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestUpdateContractConfig(&_CoreRegistryTest.TransactOpts)
}

// TestUpdateContractConfig is a paid mutator transaction binding the contract method 0x674270a0.
//
// Solidity: function testUpdateContractConfig() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestUpdateContractConfig() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestUpdateContractConfig(&_CoreRegistryTest.TransactOpts)
}

// TestUpdateMetadataForNonActiveChain is a paid mutator transaction binding the contract method 0x9fe48db2.
//
// Solidity: function testUpdateMetadataForNonActiveChain() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestUpdateMetadataForNonActiveChain(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testUpdateMetadataForNonActiveChain")
}

// TestUpdateMetadataForNonActiveChain is a paid mutator transaction binding the contract method 0x9fe48db2.
//
// Solidity: function testUpdateMetadataForNonActiveChain() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestUpdateMetadataForNonActiveChain() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestUpdateMetadataForNonActiveChain(&_CoreRegistryTest.TransactOpts)
}

// TestUpdateMetadataForNonActiveChain is a paid mutator transaction binding the contract method 0x9fe48db2.
//
// Solidity: function testUpdateMetadataForNonActiveChain() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestUpdateMetadataForNonActiveChain() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestUpdateMetadataForNonActiveChain(&_CoreRegistryTest.TransactOpts)
}

// TestUpdateMetadataUnauthorized is a paid mutator transaction binding the contract method 0xf48974af.
//
// Solidity: function testUpdateMetadataUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestUpdateMetadataUnauthorized(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testUpdateMetadataUnauthorized")
}

// TestUpdateMetadataUnauthorized is a paid mutator transaction binding the contract method 0xf48974af.
//
// Solidity: function testUpdateMetadataUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestUpdateMetadataUnauthorized() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestUpdateMetadataUnauthorized(&_CoreRegistryTest.TransactOpts)
}

// TestUpdateMetadataUnauthorized is a paid mutator transaction binding the contract method 0xf48974af.
//
// Solidity: function testUpdateMetadataUnauthorized() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestUpdateMetadataUnauthorized() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestUpdateMetadataUnauthorized(&_CoreRegistryTest.TransactOpts)
}

// TestUpdateNonExistentZRC20Token is a paid mutator transaction binding the contract method 0x933e8970.
//
// Solidity: function testUpdateNonExistentZRC20Token() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestUpdateNonExistentZRC20Token(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testUpdateNonExistentZRC20Token")
}

// TestUpdateNonExistentZRC20Token is a paid mutator transaction binding the contract method 0x933e8970.
//
// Solidity: function testUpdateNonExistentZRC20Token() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestUpdateNonExistentZRC20Token() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestUpdateNonExistentZRC20Token(&_CoreRegistryTest.TransactOpts)
}

// TestUpdateNonExistentZRC20Token is a paid mutator transaction binding the contract method 0x933e8970.
//
// Solidity: function testUpdateNonExistentZRC20Token() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestUpdateNonExistentZRC20Token() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestUpdateNonExistentZRC20Token(&_CoreRegistryTest.TransactOpts)
}

// TestWhenPaused is a paid mutator transaction binding the contract method 0x5d737a7e.
//
// Solidity: function testWhenPaused() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactor) TestWhenPaused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoreRegistryTest.contract.Transact(opts, "testWhenPaused")
}

// TestWhenPaused is a paid mutator transaction binding the contract method 0x5d737a7e.
//
// Solidity: function testWhenPaused() returns()
func (_CoreRegistryTest *CoreRegistryTestSession) TestWhenPaused() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestWhenPaused(&_CoreRegistryTest.TransactOpts)
}

// TestWhenPaused is a paid mutator transaction binding the contract method 0x5d737a7e.
//
// Solidity: function testWhenPaused() returns()
func (_CoreRegistryTest *CoreRegistryTestTransactorSession) TestWhenPaused() (*types.Transaction, error) {
	return _CoreRegistryTest.Contract.TestWhenPaused(&_CoreRegistryTest.TransactOpts)
}

// CoreRegistryTestAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the CoreRegistryTest contract.
type CoreRegistryTestAdminChangedIterator struct {
	Event *CoreRegistryTestAdminChanged // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestAdminChanged)
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
		it.Event = new(CoreRegistryTestAdminChanged)
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
func (it *CoreRegistryTestAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestAdminChanged represents a AdminChanged event raised by the CoreRegistryTest contract.
type CoreRegistryTestAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address oldAdmin, address newAdmin)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*CoreRegistryTestAdminChangedIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestAdminChangedIterator{contract: _CoreRegistryTest.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address oldAdmin, address newAdmin)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestAdminChanged) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestAdminChanged)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseAdminChanged(log types.Log) (*CoreRegistryTestAdminChanged, error) {
	event := new(CoreRegistryTestAdminChanged)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestChainMetadataUpdatedIterator is returned from FilterChainMetadataUpdated and is used to iterate over the raw logs and unpacked data for ChainMetadataUpdated events raised by the CoreRegistryTest contract.
type CoreRegistryTestChainMetadataUpdatedIterator struct {
	Event *CoreRegistryTestChainMetadataUpdated // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestChainMetadataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestChainMetadataUpdated)
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
		it.Event = new(CoreRegistryTestChainMetadataUpdated)
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
func (it *CoreRegistryTestChainMetadataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestChainMetadataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestChainMetadataUpdated represents a ChainMetadataUpdated event raised by the CoreRegistryTest contract.
type CoreRegistryTestChainMetadataUpdated struct {
	ChainId *big.Int
	Key     string
	Value   []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainMetadataUpdated is a free log retrieval operation binding the contract event 0x40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee634.
//
// Solidity: event ChainMetadataUpdated(uint256 indexed chainId, string key, bytes value)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterChainMetadataUpdated(opts *bind.FilterOpts, chainId []*big.Int) (*CoreRegistryTestChainMetadataUpdatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "ChainMetadataUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestChainMetadataUpdatedIterator{contract: _CoreRegistryTest.contract, event: "ChainMetadataUpdated", logs: logs, sub: sub}, nil
}

// WatchChainMetadataUpdated is a free log subscription operation binding the contract event 0x40c66d0452b5a398a7ebd687f5c3b020e21aa673375087ff6eb7ad214cfee634.
//
// Solidity: event ChainMetadataUpdated(uint256 indexed chainId, string key, bytes value)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchChainMetadataUpdated(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestChainMetadataUpdated, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "ChainMetadataUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestChainMetadataUpdated)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "ChainMetadataUpdated", log); err != nil {
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
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseChainMetadataUpdated(log types.Log) (*CoreRegistryTestChainMetadataUpdated, error) {
	event := new(CoreRegistryTestChainMetadataUpdated)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "ChainMetadataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestChainStatusChangedIterator is returned from FilterChainStatusChanged and is used to iterate over the raw logs and unpacked data for ChainStatusChanged events raised by the CoreRegistryTest contract.
type CoreRegistryTestChainStatusChangedIterator struct {
	Event *CoreRegistryTestChainStatusChanged // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestChainStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestChainStatusChanged)
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
		it.Event = new(CoreRegistryTestChainStatusChanged)
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
func (it *CoreRegistryTestChainStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestChainStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestChainStatusChanged represents a ChainStatusChanged event raised by the CoreRegistryTest contract.
type CoreRegistryTestChainStatusChanged struct {
	ChainId   *big.Int
	NewStatus bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChainStatusChanged is a free log retrieval operation binding the contract event 0xc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e.
//
// Solidity: event ChainStatusChanged(uint256 indexed chainId, bool newStatus)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterChainStatusChanged(opts *bind.FilterOpts, chainId []*big.Int) (*CoreRegistryTestChainStatusChangedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "ChainStatusChanged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestChainStatusChangedIterator{contract: _CoreRegistryTest.contract, event: "ChainStatusChanged", logs: logs, sub: sub}, nil
}

// WatchChainStatusChanged is a free log subscription operation binding the contract event 0xc98ceea113f96d4762d49d2885a10d7d0d16d07243bb17df97e53e9035e1415e.
//
// Solidity: event ChainStatusChanged(uint256 indexed chainId, bool newStatus)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchChainStatusChanged(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestChainStatusChanged, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "ChainStatusChanged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestChainStatusChanged)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "ChainStatusChanged", log); err != nil {
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
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseChainStatusChanged(log types.Log) (*CoreRegistryTestChainStatusChanged, error) {
	event := new(CoreRegistryTestChainStatusChanged)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "ChainStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestContractConfigurationUpdatedIterator is returned from FilterContractConfigurationUpdated and is used to iterate over the raw logs and unpacked data for ContractConfigurationUpdated events raised by the CoreRegistryTest contract.
type CoreRegistryTestContractConfigurationUpdatedIterator struct {
	Event *CoreRegistryTestContractConfigurationUpdated // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestContractConfigurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestContractConfigurationUpdated)
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
		it.Event = new(CoreRegistryTestContractConfigurationUpdated)
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
func (it *CoreRegistryTestContractConfigurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestContractConfigurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestContractConfigurationUpdated represents a ContractConfigurationUpdated event raised by the CoreRegistryTest contract.
type CoreRegistryTestContractConfigurationUpdated struct {
	ChainId      *big.Int
	ContractType string
	Key          string
	Value        []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractConfigurationUpdated is a free log retrieval operation binding the contract event 0xaea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c5.
//
// Solidity: event ContractConfigurationUpdated(uint256 indexed chainId, string contractType, string key, bytes value)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterContractConfigurationUpdated(opts *bind.FilterOpts, chainId []*big.Int) (*CoreRegistryTestContractConfigurationUpdatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "ContractConfigurationUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestContractConfigurationUpdatedIterator{contract: _CoreRegistryTest.contract, event: "ContractConfigurationUpdated", logs: logs, sub: sub}, nil
}

// WatchContractConfigurationUpdated is a free log subscription operation binding the contract event 0xaea6b6dd1ea232db3e0fc64d54fd642518ed2932043ff697ac7c8a83d651c7c5.
//
// Solidity: event ContractConfigurationUpdated(uint256 indexed chainId, string contractType, string key, bytes value)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchContractConfigurationUpdated(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestContractConfigurationUpdated, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "ContractConfigurationUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestContractConfigurationUpdated)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "ContractConfigurationUpdated", log); err != nil {
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
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseContractConfigurationUpdated(log types.Log) (*CoreRegistryTestContractConfigurationUpdated, error) {
	event := new(CoreRegistryTestContractConfigurationUpdated)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "ContractConfigurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestContractRegisteredIterator is returned from FilterContractRegistered and is used to iterate over the raw logs and unpacked data for ContractRegistered events raised by the CoreRegistryTest contract.
type CoreRegistryTestContractRegisteredIterator struct {
	Event *CoreRegistryTestContractRegistered // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestContractRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestContractRegistered)
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
		it.Event = new(CoreRegistryTestContractRegistered)
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
func (it *CoreRegistryTestContractRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestContractRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestContractRegistered represents a ContractRegistered event raised by the CoreRegistryTest contract.
type CoreRegistryTestContractRegistered struct {
	ChainId      *big.Int
	ContractType common.Hash
	AddressBytes []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractRegistered is a free log retrieval operation binding the contract event 0x20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581.
//
// Solidity: event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterContractRegistered(opts *bind.FilterOpts, chainId []*big.Int, contractType []string) (*CoreRegistryTestContractRegisteredIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "ContractRegistered", chainIdRule, contractTypeRule)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestContractRegisteredIterator{contract: _CoreRegistryTest.contract, event: "ContractRegistered", logs: logs, sub: sub}, nil
}

// WatchContractRegistered is a free log subscription operation binding the contract event 0x20319e67335097991b9d6add94a71632118372c1a0b5650654f069668dce2581.
//
// Solidity: event ContractRegistered(uint256 indexed chainId, string indexed contractType, bytes addressBytes)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchContractRegistered(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestContractRegistered, chainId []*big.Int, contractType []string) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "ContractRegistered", chainIdRule, contractTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestContractRegistered)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
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
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseContractRegistered(log types.Log) (*CoreRegistryTestContractRegistered, error) {
	event := new(CoreRegistryTestContractRegistered)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "ContractRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestContractStatusChangedIterator is returned from FilterContractStatusChanged and is used to iterate over the raw logs and unpacked data for ContractStatusChanged events raised by the CoreRegistryTest contract.
type CoreRegistryTestContractStatusChangedIterator struct {
	Event *CoreRegistryTestContractStatusChanged // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestContractStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestContractStatusChanged)
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
		it.Event = new(CoreRegistryTestContractStatusChanged)
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
func (it *CoreRegistryTestContractStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestContractStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestContractStatusChanged represents a ContractStatusChanged event raised by the CoreRegistryTest contract.
type CoreRegistryTestContractStatusChanged struct {
	AddressBytes []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractStatusChanged is a free log retrieval operation binding the contract event 0x6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c.
//
// Solidity: event ContractStatusChanged(bytes addressBytes)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterContractStatusChanged(opts *bind.FilterOpts) (*CoreRegistryTestContractStatusChangedIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "ContractStatusChanged")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestContractStatusChangedIterator{contract: _CoreRegistryTest.contract, event: "ContractStatusChanged", logs: logs, sub: sub}, nil
}

// WatchContractStatusChanged is a free log subscription operation binding the contract event 0x6db122b2555e642c944e09ae6d733a3f7600404765f612912f72b3c921c0b88c.
//
// Solidity: event ContractStatusChanged(bytes addressBytes)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchContractStatusChanged(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestContractStatusChanged) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "ContractStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestContractStatusChanged)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "ContractStatusChanged", log); err != nil {
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
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseContractStatusChanged(log types.Log) (*CoreRegistryTestContractStatusChanged, error) {
	event := new(CoreRegistryTestContractStatusChanged)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "ContractStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestRegistryManagerChangedIterator is returned from FilterRegistryManagerChanged and is used to iterate over the raw logs and unpacked data for RegistryManagerChanged events raised by the CoreRegistryTest contract.
type CoreRegistryTestRegistryManagerChangedIterator struct {
	Event *CoreRegistryTestRegistryManagerChanged // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestRegistryManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestRegistryManagerChanged)
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
		it.Event = new(CoreRegistryTestRegistryManagerChanged)
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
func (it *CoreRegistryTestRegistryManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestRegistryManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestRegistryManagerChanged represents a RegistryManagerChanged event raised by the CoreRegistryTest contract.
type CoreRegistryTestRegistryManagerChanged struct {
	OldRegistryManager common.Address
	NewRegistryManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRegistryManagerChanged is a free log retrieval operation binding the contract event 0x6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279.
//
// Solidity: event RegistryManagerChanged(address oldRegistryManager, address newRegistryManager)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterRegistryManagerChanged(opts *bind.FilterOpts) (*CoreRegistryTestRegistryManagerChangedIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "RegistryManagerChanged")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestRegistryManagerChangedIterator{contract: _CoreRegistryTest.contract, event: "RegistryManagerChanged", logs: logs, sub: sub}, nil
}

// WatchRegistryManagerChanged is a free log subscription operation binding the contract event 0x6e85328c26aff795a4964abbab261c488200d2708225db359ab39f4152645279.
//
// Solidity: event RegistryManagerChanged(address oldRegistryManager, address newRegistryManager)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchRegistryManagerChanged(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestRegistryManagerChanged) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "RegistryManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestRegistryManagerChanged)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "RegistryManagerChanged", log); err != nil {
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
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseRegistryManagerChanged(log types.Log) (*CoreRegistryTestRegistryManagerChanged, error) {
	event := new(CoreRegistryTestRegistryManagerChanged)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "RegistryManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestZRC20TokenRegisteredIterator is returned from FilterZRC20TokenRegistered and is used to iterate over the raw logs and unpacked data for ZRC20TokenRegistered events raised by the CoreRegistryTest contract.
type CoreRegistryTestZRC20TokenRegisteredIterator struct {
	Event *CoreRegistryTestZRC20TokenRegistered // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestZRC20TokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestZRC20TokenRegistered)
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
		it.Event = new(CoreRegistryTestZRC20TokenRegistered)
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
func (it *CoreRegistryTestZRC20TokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestZRC20TokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestZRC20TokenRegistered represents a ZRC20TokenRegistered event raised by the CoreRegistryTest contract.
type CoreRegistryTestZRC20TokenRegistered struct {
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
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterZRC20TokenRegistered(opts *bind.FilterOpts, originAddress [][]byte, address_ []common.Address) (*CoreRegistryTestZRC20TokenRegisteredIterator, error) {

	var originAddressRule []interface{}
	for _, originAddressItem := range originAddress {
		originAddressRule = append(originAddressRule, originAddressItem)
	}
	var address_Rule []interface{}
	for _, address_Item := range address_ {
		address_Rule = append(address_Rule, address_Item)
	}

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "ZRC20TokenRegistered", originAddressRule, address_Rule)
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestZRC20TokenRegisteredIterator{contract: _CoreRegistryTest.contract, event: "ZRC20TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchZRC20TokenRegistered is a free log subscription operation binding the contract event 0xa9edd2fd29fc8cab6015c2725afa5bc5f3b8d709a02d9e89990ef20fd781e367.
//
// Solidity: event ZRC20TokenRegistered(bytes indexed originAddress, address indexed address_, uint8 decimals, uint256 originChainId, string symbol)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchZRC20TokenRegistered(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestZRC20TokenRegistered, originAddress [][]byte, address_ []common.Address) (event.Subscription, error) {

	var originAddressRule []interface{}
	for _, originAddressItem := range originAddress {
		originAddressRule = append(originAddressRule, originAddressItem)
	}
	var address_Rule []interface{}
	for _, address_Item := range address_ {
		address_Rule = append(address_Rule, address_Item)
	}

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "ZRC20TokenRegistered", originAddressRule, address_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestZRC20TokenRegistered)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "ZRC20TokenRegistered", log); err != nil {
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
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseZRC20TokenRegistered(log types.Log) (*CoreRegistryTestZRC20TokenRegistered, error) {
	event := new(CoreRegistryTestZRC20TokenRegistered)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "ZRC20TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestZRC20TokenUpdatedIterator is returned from FilterZRC20TokenUpdated and is used to iterate over the raw logs and unpacked data for ZRC20TokenUpdated events raised by the CoreRegistryTest contract.
type CoreRegistryTestZRC20TokenUpdatedIterator struct {
	Event *CoreRegistryTestZRC20TokenUpdated // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestZRC20TokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestZRC20TokenUpdated)
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
		it.Event = new(CoreRegistryTestZRC20TokenUpdated)
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
func (it *CoreRegistryTestZRC20TokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestZRC20TokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestZRC20TokenUpdated represents a ZRC20TokenUpdated event raised by the CoreRegistryTest contract.
type CoreRegistryTestZRC20TokenUpdated struct {
	Address common.Address
	Active  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterZRC20TokenUpdated is a free log retrieval operation binding the contract event 0x9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8.
//
// Solidity: event ZRC20TokenUpdated(address address_, bool active)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterZRC20TokenUpdated(opts *bind.FilterOpts) (*CoreRegistryTestZRC20TokenUpdatedIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "ZRC20TokenUpdated")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestZRC20TokenUpdatedIterator{contract: _CoreRegistryTest.contract, event: "ZRC20TokenUpdated", logs: logs, sub: sub}, nil
}

// WatchZRC20TokenUpdated is a free log subscription operation binding the contract event 0x9542d02d4224477c9e9b53628bf5eae8b59520ea6bf2809cec7f24f76bba8ff8.
//
// Solidity: event ZRC20TokenUpdated(address address_, bool active)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchZRC20TokenUpdated(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestZRC20TokenUpdated) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "ZRC20TokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestZRC20TokenUpdated)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "ZRC20TokenUpdated", log); err != nil {
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
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseZRC20TokenUpdated(log types.Log) (*CoreRegistryTestZRC20TokenUpdated, error) {
	event := new(CoreRegistryTestZRC20TokenUpdated)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "ZRC20TokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogIterator struct {
	Event *CoreRegistryTestLog // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLog)
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
		it.Event = new(CoreRegistryTestLog)
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
func (it *CoreRegistryTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLog represents a Log event raised by the CoreRegistryTest contract.
type CoreRegistryTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLog(opts *bind.FilterOpts) (*CoreRegistryTestLogIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogIterator{contract: _CoreRegistryTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLog) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLog)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLog(log types.Log) (*CoreRegistryTestLog, error) {
	event := new(CoreRegistryTestLog)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogAddressIterator struct {
	Event *CoreRegistryTestLogAddress // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogAddress)
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
		it.Event = new(CoreRegistryTestLogAddress)
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
func (it *CoreRegistryTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogAddress represents a LogAddress event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*CoreRegistryTestLogAddressIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogAddressIterator{contract: _CoreRegistryTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogAddress)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_address", log); err != nil {
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

// ParseLogAddress is a log parse operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogAddress(log types.Log) (*CoreRegistryTestLogAddress, error) {
	event := new(CoreRegistryTestLogAddress)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogArrayIterator struct {
	Event *CoreRegistryTestLogArray // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogArray)
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
		it.Event = new(CoreRegistryTestLogArray)
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
func (it *CoreRegistryTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogArray represents a LogArray event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*CoreRegistryTestLogArrayIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogArrayIterator{contract: _CoreRegistryTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogArray) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogArray)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_array", log); err != nil {
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

// ParseLogArray is a log parse operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogArray(log types.Log) (*CoreRegistryTestLogArray, error) {
	event := new(CoreRegistryTestLogArray)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogArray0Iterator struct {
	Event *CoreRegistryTestLogArray0 // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogArray0)
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
		it.Event = new(CoreRegistryTestLogArray0)
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
func (it *CoreRegistryTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogArray0 represents a LogArray0 event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*CoreRegistryTestLogArray0Iterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogArray0Iterator{contract: _CoreRegistryTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogArray0)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_array0", log); err != nil {
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

// ParseLogArray0 is a log parse operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogArray0(log types.Log) (*CoreRegistryTestLogArray0, error) {
	event := new(CoreRegistryTestLogArray0)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogArray1Iterator struct {
	Event *CoreRegistryTestLogArray1 // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogArray1)
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
		it.Event = new(CoreRegistryTestLogArray1)
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
func (it *CoreRegistryTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogArray1 represents a LogArray1 event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*CoreRegistryTestLogArray1Iterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogArray1Iterator{contract: _CoreRegistryTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogArray1)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_array1", log); err != nil {
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

// ParseLogArray1 is a log parse operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogArray1(log types.Log) (*CoreRegistryTestLogArray1, error) {
	event := new(CoreRegistryTestLogArray1)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogBytesIterator struct {
	Event *CoreRegistryTestLogBytes // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogBytes)
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
		it.Event = new(CoreRegistryTestLogBytes)
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
func (it *CoreRegistryTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogBytes represents a LogBytes event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*CoreRegistryTestLogBytesIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogBytesIterator{contract: _CoreRegistryTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogBytes)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
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

// ParseLogBytes is a log parse operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogBytes(log types.Log) (*CoreRegistryTestLogBytes, error) {
	event := new(CoreRegistryTestLogBytes)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogBytes32Iterator struct {
	Event *CoreRegistryTestLogBytes32 // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogBytes32)
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
		it.Event = new(CoreRegistryTestLogBytes32)
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
func (it *CoreRegistryTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogBytes32 represents a LogBytes32 event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*CoreRegistryTestLogBytes32Iterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogBytes32Iterator{contract: _CoreRegistryTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogBytes32)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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

// ParseLogBytes32 is a log parse operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogBytes32(log types.Log) (*CoreRegistryTestLogBytes32, error) {
	event := new(CoreRegistryTestLogBytes32)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogIntIterator struct {
	Event *CoreRegistryTestLogInt // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogInt)
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
		it.Event = new(CoreRegistryTestLogInt)
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
func (it *CoreRegistryTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogInt represents a LogInt event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*CoreRegistryTestLogIntIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogIntIterator{contract: _CoreRegistryTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogInt) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogInt)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_int", log); err != nil {
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

// ParseLogInt is a log parse operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogInt(log types.Log) (*CoreRegistryTestLogInt, error) {
	event := new(CoreRegistryTestLogInt)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedAddressIterator struct {
	Event *CoreRegistryTestLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogNamedAddress)
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
		it.Event = new(CoreRegistryTestLogNamedAddress)
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
func (it *CoreRegistryTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogNamedAddress represents a LogNamedAddress event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*CoreRegistryTestLogNamedAddressIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogNamedAddressIterator{contract: _CoreRegistryTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogNamedAddress)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
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

// ParseLogNamedAddress is a log parse operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogNamedAddress(log types.Log) (*CoreRegistryTestLogNamedAddress, error) {
	event := new(CoreRegistryTestLogNamedAddress)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedArrayIterator struct {
	Event *CoreRegistryTestLogNamedArray // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogNamedArray)
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
		it.Event = new(CoreRegistryTestLogNamedArray)
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
func (it *CoreRegistryTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogNamedArray represents a LogNamedArray event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*CoreRegistryTestLogNamedArrayIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogNamedArrayIterator{contract: _CoreRegistryTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogNamedArray)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
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

// ParseLogNamedArray is a log parse operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogNamedArray(log types.Log) (*CoreRegistryTestLogNamedArray, error) {
	event := new(CoreRegistryTestLogNamedArray)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedArray0Iterator struct {
	Event *CoreRegistryTestLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogNamedArray0)
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
		it.Event = new(CoreRegistryTestLogNamedArray0)
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
func (it *CoreRegistryTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogNamedArray0 represents a LogNamedArray0 event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*CoreRegistryTestLogNamedArray0Iterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogNamedArray0Iterator{contract: _CoreRegistryTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogNamedArray0)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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

// ParseLogNamedArray0 is a log parse operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogNamedArray0(log types.Log) (*CoreRegistryTestLogNamedArray0, error) {
	event := new(CoreRegistryTestLogNamedArray0)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedArray1Iterator struct {
	Event *CoreRegistryTestLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogNamedArray1)
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
		it.Event = new(CoreRegistryTestLogNamedArray1)
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
func (it *CoreRegistryTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogNamedArray1 represents a LogNamedArray1 event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*CoreRegistryTestLogNamedArray1Iterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogNamedArray1Iterator{contract: _CoreRegistryTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogNamedArray1)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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

// ParseLogNamedArray1 is a log parse operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogNamedArray1(log types.Log) (*CoreRegistryTestLogNamedArray1, error) {
	event := new(CoreRegistryTestLogNamedArray1)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedBytesIterator struct {
	Event *CoreRegistryTestLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogNamedBytes)
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
		it.Event = new(CoreRegistryTestLogNamedBytes)
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
func (it *CoreRegistryTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogNamedBytes represents a LogNamedBytes event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*CoreRegistryTestLogNamedBytesIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogNamedBytesIterator{contract: _CoreRegistryTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogNamedBytes)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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

// ParseLogNamedBytes is a log parse operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogNamedBytes(log types.Log) (*CoreRegistryTestLogNamedBytes, error) {
	event := new(CoreRegistryTestLogNamedBytes)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedBytes32Iterator struct {
	Event *CoreRegistryTestLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogNamedBytes32)
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
		it.Event = new(CoreRegistryTestLogNamedBytes32)
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
func (it *CoreRegistryTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*CoreRegistryTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogNamedBytes32Iterator{contract: _CoreRegistryTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogNamedBytes32)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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

// ParseLogNamedBytes32 is a log parse operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogNamedBytes32(log types.Log) (*CoreRegistryTestLogNamedBytes32, error) {
	event := new(CoreRegistryTestLogNamedBytes32)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedDecimalIntIterator struct {
	Event *CoreRegistryTestLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogNamedDecimalInt)
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
		it.Event = new(CoreRegistryTestLogNamedDecimalInt)
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
func (it *CoreRegistryTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*CoreRegistryTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogNamedDecimalIntIterator{contract: _CoreRegistryTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogNamedDecimalInt)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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

// ParseLogNamedDecimalInt is a log parse operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*CoreRegistryTestLogNamedDecimalInt, error) {
	event := new(CoreRegistryTestLogNamedDecimalInt)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedDecimalUintIterator struct {
	Event *CoreRegistryTestLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogNamedDecimalUint)
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
		it.Event = new(CoreRegistryTestLogNamedDecimalUint)
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
func (it *CoreRegistryTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*CoreRegistryTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogNamedDecimalUintIterator{contract: _CoreRegistryTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogNamedDecimalUint)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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

// ParseLogNamedDecimalUint is a log parse operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*CoreRegistryTestLogNamedDecimalUint, error) {
	event := new(CoreRegistryTestLogNamedDecimalUint)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedIntIterator struct {
	Event *CoreRegistryTestLogNamedInt // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogNamedInt)
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
		it.Event = new(CoreRegistryTestLogNamedInt)
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
func (it *CoreRegistryTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogNamedInt represents a LogNamedInt event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*CoreRegistryTestLogNamedIntIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogNamedIntIterator{contract: _CoreRegistryTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogNamedInt)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
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

// ParseLogNamedInt is a log parse operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogNamedInt(log types.Log) (*CoreRegistryTestLogNamedInt, error) {
	event := new(CoreRegistryTestLogNamedInt)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedStringIterator struct {
	Event *CoreRegistryTestLogNamedString // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogNamedString)
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
		it.Event = new(CoreRegistryTestLogNamedString)
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
func (it *CoreRegistryTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogNamedString represents a LogNamedString event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*CoreRegistryTestLogNamedStringIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogNamedStringIterator{contract: _CoreRegistryTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogNamedString)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
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

// ParseLogNamedString is a log parse operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogNamedString(log types.Log) (*CoreRegistryTestLogNamedString, error) {
	event := new(CoreRegistryTestLogNamedString)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedUintIterator struct {
	Event *CoreRegistryTestLogNamedUint // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogNamedUint)
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
		it.Event = new(CoreRegistryTestLogNamedUint)
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
func (it *CoreRegistryTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogNamedUint represents a LogNamedUint event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*CoreRegistryTestLogNamedUintIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogNamedUintIterator{contract: _CoreRegistryTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogNamedUint)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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

// ParseLogNamedUint is a log parse operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogNamedUint(log types.Log) (*CoreRegistryTestLogNamedUint, error) {
	event := new(CoreRegistryTestLogNamedUint)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogStringIterator struct {
	Event *CoreRegistryTestLogString // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogString)
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
		it.Event = new(CoreRegistryTestLogString)
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
func (it *CoreRegistryTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogString represents a LogString event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogString(opts *bind.FilterOpts) (*CoreRegistryTestLogStringIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogStringIterator{contract: _CoreRegistryTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogString) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogString)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_string", log); err != nil {
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

// ParseLogString is a log parse operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogString(log types.Log) (*CoreRegistryTestLogString, error) {
	event := new(CoreRegistryTestLogString)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogUintIterator struct {
	Event *CoreRegistryTestLogUint // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogUint)
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
		it.Event = new(CoreRegistryTestLogUint)
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
func (it *CoreRegistryTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogUint represents a LogUint event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*CoreRegistryTestLogUintIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogUintIterator{contract: _CoreRegistryTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogUint) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogUint)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "log_uint", log); err != nil {
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

// ParseLogUint is a log parse operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogUint(log types.Log) (*CoreRegistryTestLogUint, error) {
	event := new(CoreRegistryTestLogUint)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CoreRegistryTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the CoreRegistryTest contract.
type CoreRegistryTestLogsIterator struct {
	Event *CoreRegistryTestLogs // Event containing the contract specifics and raw log

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
func (it *CoreRegistryTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoreRegistryTestLogs)
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
		it.Event = new(CoreRegistryTestLogs)
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
func (it *CoreRegistryTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoreRegistryTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoreRegistryTestLogs represents a Logs event raised by the CoreRegistryTest contract.
type CoreRegistryTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) FilterLogs(opts *bind.FilterOpts) (*CoreRegistryTestLogsIterator, error) {

	logs, sub, err := _CoreRegistryTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &CoreRegistryTestLogsIterator{contract: _CoreRegistryTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *CoreRegistryTestLogs) (event.Subscription, error) {

	logs, sub, err := _CoreRegistryTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoreRegistryTestLogs)
				if err := _CoreRegistryTest.contract.UnpackLog(event, "logs", log); err != nil {
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

// ParseLogs is a log parse operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_CoreRegistryTest *CoreRegistryTestFilterer) ParseLogs(log types.Log) (*CoreRegistryTestLogs, error) {
	event := new(CoreRegistryTestLogs)
	if err := _CoreRegistryTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
