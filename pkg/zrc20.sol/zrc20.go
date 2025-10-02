// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zrc20

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

// ZRC20MetaData contains all meta data concerning the ZRC20 contract.
var ZRC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"chainid_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"coinType_\",\"type\":\"uint8\",\"internalType\":\"enumCoinType\"},{\"name\":\"gasLimit_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"systemContractAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gatewayAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"COIN_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumCoinType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROTOCOL_FLAT_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SYSTEM_CONTRACT_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gatewayAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"newName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSymbol\",\"inputs\":[{\"name\":\"newSymbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateGasLimit\",\"inputs\":[{\"name\":\"gasLimit_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateGatewayAddress\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateProtocolFlatFee\",\"inputs\":[{\"name\":\"protocolFlatFee_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSystemContractAddress\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawGasFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawGasFeeWithGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"from\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGateway\",\"inputs\":[{\"name\":\"gateway\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedProtocolFlatFee\",\"inputs\":[{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedSystemContract\",\"inputs\":[{\"name\":\"systemContract\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotFungibleModule\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LowAllowance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LowBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroGasCoin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroGasPrice\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161207238038061207283398101604081905261002f916101f0565b6001600160a01b038216158061004c57506001600160a01b038116155b1561006a5760405163d92e233d60e01b815260040160405180910390fd5b60066100768982610342565b5060076100838882610342565b506008805460ff191660ff881617905560808590528360028111156100aa576100aa610400565b60a08160028111156100be576100be610400565b905250600192909255600080546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506104169350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261013357600080fd5b81516001600160401b0381111561014c5761014c61010c565b604051601f8201601f19908116603f011681016001600160401b038111828210171561017a5761017a61010c565b60405281815283820160200185101561019257600080fd5b60005b828110156101b157602081860181015183830182015201610195565b506000918101602001919091529392505050565b8051600381106101d457600080fd5b919050565b80516001600160a01b03811681146101d457600080fd5b600080600080600080600080610100898b03121561020d57600080fd5b88516001600160401b0381111561022357600080fd5b61022f8b828c01610122565b60208b015190995090506001600160401b0381111561024d57600080fd5b6102598b828c01610122565b975050604089015160ff8116811461027057600080fd5b60608a0151909650945061028660808a016101c5565b60a08a0151909450925061029c60c08a016101d9565b91506102aa60e08a016101d9565b90509295985092959890939650565b600181811c908216806102cd57607f821691505b6020821081036102ed57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561033d57806000526020600020601f840160051c8101602085101561031a5750805b601f840160051c820191505b8181101561033a5760008155600101610326565b50505b505050565b81516001600160401b0381111561035b5761035b61010c565b61036f8161036984546102b9565b846102f3565b6020601f8211600181146103a3576000831561038b5750848201515b600019600385901b1c1916600184901b17845561033a565b600084815260208120601f198516915b828110156103d357878501518255602094850194600190920191016103b3565b50848210156103f15786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a051611c1b61045760003960006103440152600081816102f001528181610bdc01528181610ce201528181610efe01526110040152611c1b6000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c806395d89b41116100f9578063ccc7759911610097578063eddeb12311610071578063eddeb12314610461578063f2441b3214610474578063f687d12a14610494578063fc5fecd5146104a757600080fd5b8063ccc77599146103d4578063d9eeebed146103e7578063dd62ed3e1461041b57600080fd5b8063b84c8246116100d3578063b84c824614610386578063c47f00271461039b578063c7012626146103ae578063c835d7cc146103c157600080fd5b806395d89b4114610337578063a3413d031461033f578063a9059cbb1461037357600080fd5b80633ce4a5bc116101665780634d8943bb116101405780634d8943bb146102ac57806370a08231146102b557806385e1f4d0146102eb5780638b851b951461031257600080fd5b80633ce4a5bc1461024657806342966c681461028657806347e7ef241461029957600080fd5b806318160ddd1161019757806318160ddd1461021657806323b872dd1461021e578063313ce5671461023157600080fd5b806306fdde03146101be578063091d2788146101dc578063095ea7b3146101f3575b600080fd5b6101c66104ba565b6040516101d39190611648565b60405180910390f35b6101e560015481565b6040519081526020016101d3565b610206610201366004611687565b61054c565b60405190151581526020016101d3565b6005546101e5565b61020661022c3660046116b3565b610563565b60085460405160ff90911681526020016101d3565b61026173735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d3565b6102066102943660046116f4565b6105fa565b6102066102a7366004611687565b61060e565b6101e560025481565b6101e56102c336600461170d565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6101e57f000000000000000000000000000000000000000000000000000000000000000081565b60085461026190610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101c6610767565b6103667f000000000000000000000000000000000000000000000000000000000000000081565b6040516101d3919061172a565b610206610381366004611687565b610776565b610399610394366004611832565b610783565b005b6103996103a9366004611832565b6107e0565b6102066103bc366004611883565b610839565b6103996103cf36600461170d565b610988565b6103996103e236600461170d565b610a9c565b6103ef610bb0565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101d3565b6101e56104293660046118dc565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260046020908152604080832093909416825291909152205490565b61039961046f3660046116f4565b610dce565b6000546102619073ffffffffffffffffffffffffffffffffffffffff1681565b6103996104a23660046116f4565b610e50565b6103ef6104b53660046116f4565b610ed2565b6060600680546104c990611915565b80601f01602080910402602001604051908101604052809291908181526020018280546104f590611915565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b5050505050905090565b60006105593384846110ee565b5060015b92915050565b60006105708484846111f7565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152902054828110156105db576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ef85336105ea8685611997565b6110ee565b506001949350505050565b600061060633836113b2565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab1480159061064c575060005473ffffffffffffffffffffffffffffffffffffffff163314155b80156106755750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b156106ac576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106b683836114f4565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526107569186906119aa565b60405180910390a250600192915050565b6060600780546104c990611915565b60006105593384846111f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107d0576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076107dc8282611a1b565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461082d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60066107dc8282611a1b565b6000806000610846610bb0565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af11580156108d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fc9190611b34565b610932576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61093c33856113b2565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161097591899189918791611b56565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109d5576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a22576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610ae9576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b36576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610c43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c679190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610cb6576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610d45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d699190611ba2565b905080600003610da5576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060025460015483610db89190611bbb565b610dc29190611bd2565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e1b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610a91565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e9d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610f65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f899190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610fd8576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015611067573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108b9190611ba2565b9050806000036110c7576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546000906110d78784611bbb565b6110e19190611bd2565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff831661113b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611188576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611244576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611291576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040902054818110156112f1576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112fb8282611997565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260036020526040808220939093559085168152908120805484929061133e908490611bd2565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113a491815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166113ff576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260409020548181101561145f576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114698282611997565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260036020526040812091909155600580548492906114a4908490611997565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016111ea565b73ffffffffffffffffffffffffffffffffffffffff8216611541576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008282546115539190611bd2565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260408120805483929061158d908490611bd2565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561160a576020818501810151868301820152016115ee565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061165b60208301846115e4565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461168457600080fd5b50565b6000806040838503121561169a57600080fd5b82356116a581611662565b946020939093013593505050565b6000806000606084860312156116c857600080fd5b83356116d381611662565b925060208401356116e381611662565b929592945050506040919091013590565b60006020828403121561170657600080fd5b5035919050565b60006020828403121561171f57600080fd5b813561165b81611662565b6020810160038310611765577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff8411156117b5576117b561176b565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156118025761180261176b565b60405283815290508082840185101561181a57600080fd5b83836020830137600060208583010152509392505050565b60006020828403121561184457600080fd5b813567ffffffffffffffff81111561185b57600080fd5b8201601f8101841361186c57600080fd5b61187b8482356020840161179a565b949350505050565b6000806040838503121561189657600080fd5b823567ffffffffffffffff8111156118ad57600080fd5b8301601f810185136118be57600080fd5b6118cd8582356020840161179a565b95602094909401359450505050565b600080604083850312156118ef57600080fd5b82356118fa81611662565b9150602083013561190a81611662565b809150509250929050565b600181811c9082168061192957607f821691505b602082108103611962577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561055d5761055d611968565b6040815260006119bd60408301856115e4565b90508260208301529392505050565b601f821115611a1657806000526020600020601f840160051c810160208510156119f35750805b601f840160051c820191505b81811015611a1357600081556001016119ff565b50505b505050565b815167ffffffffffffffff811115611a3557611a3561176b565b611a4981611a438454611915565b846119cc565b6020601f821160018114611a9b5760008315611a655750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611a13565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611ae95787850151825560209485019460019092019101611ac9565b5084821015611b2557868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b600060208284031215611b4657600080fd5b8151801515811461165b57600080fd5b608081526000611b6960808301876115e4565b6020830195909552506040810192909252606090910152919050565b600060208284031215611b9757600080fd5b815161165b81611662565b600060208284031215611bb457600080fd5b5051919050565b808202811582820484141761055d5761055d611968565b8082018082111561055d5761055d61196856fea2646970667358221220e7be4f6cbd9165fe86191f383968618f320a1c54b4d694652e1d738c851a10e064736f6c634300081a0033",
}

// ZRC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ZRC20MetaData.ABI instead.
var ZRC20ABI = ZRC20MetaData.ABI

// ZRC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZRC20MetaData.Bin instead.
var ZRC20Bin = ZRC20MetaData.Bin

// DeployZRC20 deploys a new Ethereum contract, binding an instance of ZRC20 to it.
func DeployZRC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8, chainid_ *big.Int, coinType_ uint8, gasLimit_ *big.Int, systemContractAddress_ common.Address, gatewayAddress_ common.Address) (common.Address, *types.Transaction, *ZRC20, error) {
	parsed, err := ZRC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZRC20Bin), backend, name_, symbol_, decimals_, chainid_, coinType_, gasLimit_, systemContractAddress_, gatewayAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZRC20{ZRC20Caller: ZRC20Caller{contract: contract}, ZRC20Transactor: ZRC20Transactor{contract: contract}, ZRC20Filterer: ZRC20Filterer{contract: contract}}, nil
}

// ZRC20 is an auto generated Go binding around an Ethereum contract.
type ZRC20 struct {
	ZRC20Caller     // Read-only binding to the contract
	ZRC20Transactor // Write-only binding to the contract
	ZRC20Filterer   // Log filterer for contract events
}

// ZRC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ZRC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZRC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ZRC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZRC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZRC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZRC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZRC20Session struct {
	Contract     *ZRC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZRC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZRC20CallerSession struct {
	Contract *ZRC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZRC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZRC20TransactorSession struct {
	Contract     *ZRC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZRC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ZRC20Raw struct {
	Contract *ZRC20 // Generic contract binding to access the raw methods on
}

// ZRC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZRC20CallerRaw struct {
	Contract *ZRC20Caller // Generic read-only contract binding to access the raw methods on
}

// ZRC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZRC20TransactorRaw struct {
	Contract *ZRC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewZRC20 creates a new instance of ZRC20, bound to a specific deployed contract.
func NewZRC20(address common.Address, backend bind.ContractBackend) (*ZRC20, error) {
	contract, err := bindZRC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZRC20{ZRC20Caller: ZRC20Caller{contract: contract}, ZRC20Transactor: ZRC20Transactor{contract: contract}, ZRC20Filterer: ZRC20Filterer{contract: contract}}, nil
}

// NewZRC20Caller creates a new read-only instance of ZRC20, bound to a specific deployed contract.
func NewZRC20Caller(address common.Address, caller bind.ContractCaller) (*ZRC20Caller, error) {
	contract, err := bindZRC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZRC20Caller{contract: contract}, nil
}

// NewZRC20Transactor creates a new write-only instance of ZRC20, bound to a specific deployed contract.
func NewZRC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ZRC20Transactor, error) {
	contract, err := bindZRC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZRC20Transactor{contract: contract}, nil
}

// NewZRC20Filterer creates a new log filterer instance of ZRC20, bound to a specific deployed contract.
func NewZRC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ZRC20Filterer, error) {
	contract, err := bindZRC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZRC20Filterer{contract: contract}, nil
}

// bindZRC20 binds a generic wrapper to an already deployed contract.
func bindZRC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZRC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZRC20 *ZRC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZRC20.Contract.ZRC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZRC20 *ZRC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZRC20.Contract.ZRC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZRC20 *ZRC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZRC20.Contract.ZRC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZRC20 *ZRC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZRC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZRC20 *ZRC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZRC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZRC20 *ZRC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZRC20.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_ZRC20 *ZRC20Caller) CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_ZRC20 *ZRC20Session) CHAINID() (*big.Int, error) {
	return _ZRC20.Contract.CHAINID(&_ZRC20.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint256)
func (_ZRC20 *ZRC20CallerSession) CHAINID() (*big.Int, error) {
	return _ZRC20.Contract.CHAINID(&_ZRC20.CallOpts)
}

// COINTYPE is a free data retrieval call binding the contract method 0xa3413d03.
//
// Solidity: function COIN_TYPE() view returns(uint8)
func (_ZRC20 *ZRC20Caller) COINTYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "COIN_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// COINTYPE is a free data retrieval call binding the contract method 0xa3413d03.
//
// Solidity: function COIN_TYPE() view returns(uint8)
func (_ZRC20 *ZRC20Session) COINTYPE() (uint8, error) {
	return _ZRC20.Contract.COINTYPE(&_ZRC20.CallOpts)
}

// COINTYPE is a free data retrieval call binding the contract method 0xa3413d03.
//
// Solidity: function COIN_TYPE() view returns(uint8)
func (_ZRC20 *ZRC20CallerSession) COINTYPE() (uint8, error) {
	return _ZRC20.Contract.COINTYPE(&_ZRC20.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20Caller) FUNGIBLEMODULEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "FUNGIBLE_MODULE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20Session) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _ZRC20.Contract.FUNGIBLEMODULEADDRESS(&_ZRC20.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20CallerSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _ZRC20.Contract.FUNGIBLEMODULEADDRESS(&_ZRC20.CallOpts)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_ZRC20 *ZRC20Caller) GASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_ZRC20 *ZRC20Session) GASLIMIT() (*big.Int, error) {
	return _ZRC20.Contract.GASLIMIT(&_ZRC20.CallOpts)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_ZRC20 *ZRC20CallerSession) GASLIMIT() (*big.Int, error) {
	return _ZRC20.Contract.GASLIMIT(&_ZRC20.CallOpts)
}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_ZRC20 *ZRC20Caller) PROTOCOLFLATFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "PROTOCOL_FLAT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_ZRC20 *ZRC20Session) PROTOCOLFLATFEE() (*big.Int, error) {
	return _ZRC20.Contract.PROTOCOLFLATFEE(&_ZRC20.CallOpts)
}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_ZRC20 *ZRC20CallerSession) PROTOCOLFLATFEE() (*big.Int, error) {
	return _ZRC20.Contract.PROTOCOLFLATFEE(&_ZRC20.CallOpts)
}

// SYSTEMCONTRACTADDRESS is a free data retrieval call binding the contract method 0xf2441b32.
//
// Solidity: function SYSTEM_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20Caller) SYSTEMCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "SYSTEM_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMCONTRACTADDRESS is a free data retrieval call binding the contract method 0xf2441b32.
//
// Solidity: function SYSTEM_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20Session) SYSTEMCONTRACTADDRESS() (common.Address, error) {
	return _ZRC20.Contract.SYSTEMCONTRACTADDRESS(&_ZRC20.CallOpts)
}

// SYSTEMCONTRACTADDRESS is a free data retrieval call binding the contract method 0xf2441b32.
//
// Solidity: function SYSTEM_CONTRACT_ADDRESS() view returns(address)
func (_ZRC20 *ZRC20CallerSession) SYSTEMCONTRACTADDRESS() (common.Address, error) {
	return _ZRC20.Contract.SYSTEMCONTRACTADDRESS(&_ZRC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZRC20 *ZRC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZRC20 *ZRC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZRC20.Contract.Allowance(&_ZRC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ZRC20 *ZRC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ZRC20.Contract.Allowance(&_ZRC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZRC20 *ZRC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZRC20 *ZRC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZRC20.Contract.BalanceOf(&_ZRC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ZRC20 *ZRC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ZRC20.Contract.BalanceOf(&_ZRC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZRC20 *ZRC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZRC20 *ZRC20Session) Decimals() (uint8, error) {
	return _ZRC20.Contract.Decimals(&_ZRC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ZRC20 *ZRC20CallerSession) Decimals() (uint8, error) {
	return _ZRC20.Contract.Decimals(&_ZRC20.CallOpts)
}

// GatewayAddress is a free data retrieval call binding the contract method 0x8b851b95.
//
// Solidity: function gatewayAddress() view returns(address)
func (_ZRC20 *ZRC20Caller) GatewayAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "gatewayAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatewayAddress is a free data retrieval call binding the contract method 0x8b851b95.
//
// Solidity: function gatewayAddress() view returns(address)
func (_ZRC20 *ZRC20Session) GatewayAddress() (common.Address, error) {
	return _ZRC20.Contract.GatewayAddress(&_ZRC20.CallOpts)
}

// GatewayAddress is a free data retrieval call binding the contract method 0x8b851b95.
//
// Solidity: function gatewayAddress() view returns(address)
func (_ZRC20 *ZRC20CallerSession) GatewayAddress() (common.Address, error) {
	return _ZRC20.Contract.GatewayAddress(&_ZRC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZRC20 *ZRC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZRC20 *ZRC20Session) Name() (string, error) {
	return _ZRC20.Contract.Name(&_ZRC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZRC20 *ZRC20CallerSession) Name() (string, error) {
	return _ZRC20.Contract.Name(&_ZRC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZRC20 *ZRC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZRC20 *ZRC20Session) Symbol() (string, error) {
	return _ZRC20.Contract.Symbol(&_ZRC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZRC20 *ZRC20CallerSession) Symbol() (string, error) {
	return _ZRC20.Contract.Symbol(&_ZRC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZRC20 *ZRC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZRC20 *ZRC20Session) TotalSupply() (*big.Int, error) {
	return _ZRC20.Contract.TotalSupply(&_ZRC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZRC20 *ZRC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ZRC20.Contract.TotalSupply(&_ZRC20.CallOpts)
}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_ZRC20 *ZRC20Caller) WithdrawGasFee(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "withdrawGasFee")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_ZRC20 *ZRC20Session) WithdrawGasFee() (common.Address, *big.Int, error) {
	return _ZRC20.Contract.WithdrawGasFee(&_ZRC20.CallOpts)
}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_ZRC20 *ZRC20CallerSession) WithdrawGasFee() (common.Address, *big.Int, error) {
	return _ZRC20.Contract.WithdrawGasFee(&_ZRC20.CallOpts)
}

// WithdrawGasFeeWithGasLimit is a free data retrieval call binding the contract method 0xfc5fecd5.
//
// Solidity: function withdrawGasFeeWithGasLimit(uint256 gasLimit) view returns(address, uint256)
func (_ZRC20 *ZRC20Caller) WithdrawGasFeeWithGasLimit(opts *bind.CallOpts, gasLimit *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ZRC20.contract.Call(opts, &out, "withdrawGasFeeWithGasLimit", gasLimit)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// WithdrawGasFeeWithGasLimit is a free data retrieval call binding the contract method 0xfc5fecd5.
//
// Solidity: function withdrawGasFeeWithGasLimit(uint256 gasLimit) view returns(address, uint256)
func (_ZRC20 *ZRC20Session) WithdrawGasFeeWithGasLimit(gasLimit *big.Int) (common.Address, *big.Int, error) {
	return _ZRC20.Contract.WithdrawGasFeeWithGasLimit(&_ZRC20.CallOpts, gasLimit)
}

// WithdrawGasFeeWithGasLimit is a free data retrieval call binding the contract method 0xfc5fecd5.
//
// Solidity: function withdrawGasFeeWithGasLimit(uint256 gasLimit) view returns(address, uint256)
func (_ZRC20 *ZRC20CallerSession) WithdrawGasFeeWithGasLimit(gasLimit *big.Int) (common.Address, *big.Int, error) {
	return _ZRC20.Contract.WithdrawGasFeeWithGasLimit(&_ZRC20.CallOpts, gasLimit)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Approve(&_ZRC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Approve(&_ZRC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Burn(&_ZRC20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_ZRC20 *ZRC20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Burn(&_ZRC20.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Transactor) Deposit(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "deposit", to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Session) Deposit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Deposit(&_ZRC20.TransactOpts, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20TransactorSession) Deposit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Deposit(&_ZRC20.TransactOpts, to, amount)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_ZRC20 *ZRC20Transactor) SetName(opts *bind.TransactOpts, newName string) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "setName", newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_ZRC20 *ZRC20Session) SetName(newName string) (*types.Transaction, error) {
	return _ZRC20.Contract.SetName(&_ZRC20.TransactOpts, newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_ZRC20 *ZRC20TransactorSession) SetName(newName string) (*types.Transaction, error) {
	return _ZRC20.Contract.SetName(&_ZRC20.TransactOpts, newName)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_ZRC20 *ZRC20Transactor) SetSymbol(opts *bind.TransactOpts, newSymbol string) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "setSymbol", newSymbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_ZRC20 *ZRC20Session) SetSymbol(newSymbol string) (*types.Transaction, error) {
	return _ZRC20.Contract.SetSymbol(&_ZRC20.TransactOpts, newSymbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_ZRC20 *ZRC20TransactorSession) SetSymbol(newSymbol string) (*types.Transaction, error) {
	return _ZRC20.Contract.SetSymbol(&_ZRC20.TransactOpts, newSymbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Transfer(&_ZRC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Transfer(&_ZRC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.TransferFrom(&_ZRC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.TransferFrom(&_ZRC20.TransactOpts, sender, recipient, amount)
}

// UpdateGasLimit is a paid mutator transaction binding the contract method 0xf687d12a.
//
// Solidity: function updateGasLimit(uint256 gasLimit_) returns()
func (_ZRC20 *ZRC20Transactor) UpdateGasLimit(opts *bind.TransactOpts, gasLimit_ *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "updateGasLimit", gasLimit_)
}

// UpdateGasLimit is a paid mutator transaction binding the contract method 0xf687d12a.
//
// Solidity: function updateGasLimit(uint256 gasLimit_) returns()
func (_ZRC20 *ZRC20Session) UpdateGasLimit(gasLimit_ *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateGasLimit(&_ZRC20.TransactOpts, gasLimit_)
}

// UpdateGasLimit is a paid mutator transaction binding the contract method 0xf687d12a.
//
// Solidity: function updateGasLimit(uint256 gasLimit_) returns()
func (_ZRC20 *ZRC20TransactorSession) UpdateGasLimit(gasLimit_ *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateGasLimit(&_ZRC20.TransactOpts, gasLimit_)
}

// UpdateGatewayAddress is a paid mutator transaction binding the contract method 0xccc77599.
//
// Solidity: function updateGatewayAddress(address addr) returns()
func (_ZRC20 *ZRC20Transactor) UpdateGatewayAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "updateGatewayAddress", addr)
}

// UpdateGatewayAddress is a paid mutator transaction binding the contract method 0xccc77599.
//
// Solidity: function updateGatewayAddress(address addr) returns()
func (_ZRC20 *ZRC20Session) UpdateGatewayAddress(addr common.Address) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateGatewayAddress(&_ZRC20.TransactOpts, addr)
}

// UpdateGatewayAddress is a paid mutator transaction binding the contract method 0xccc77599.
//
// Solidity: function updateGatewayAddress(address addr) returns()
func (_ZRC20 *ZRC20TransactorSession) UpdateGatewayAddress(addr common.Address) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateGatewayAddress(&_ZRC20.TransactOpts, addr)
}

// UpdateProtocolFlatFee is a paid mutator transaction binding the contract method 0xeddeb123.
//
// Solidity: function updateProtocolFlatFee(uint256 protocolFlatFee_) returns()
func (_ZRC20 *ZRC20Transactor) UpdateProtocolFlatFee(opts *bind.TransactOpts, protocolFlatFee_ *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "updateProtocolFlatFee", protocolFlatFee_)
}

// UpdateProtocolFlatFee is a paid mutator transaction binding the contract method 0xeddeb123.
//
// Solidity: function updateProtocolFlatFee(uint256 protocolFlatFee_) returns()
func (_ZRC20 *ZRC20Session) UpdateProtocolFlatFee(protocolFlatFee_ *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateProtocolFlatFee(&_ZRC20.TransactOpts, protocolFlatFee_)
}

// UpdateProtocolFlatFee is a paid mutator transaction binding the contract method 0xeddeb123.
//
// Solidity: function updateProtocolFlatFee(uint256 protocolFlatFee_) returns()
func (_ZRC20 *ZRC20TransactorSession) UpdateProtocolFlatFee(protocolFlatFee_ *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateProtocolFlatFee(&_ZRC20.TransactOpts, protocolFlatFee_)
}

// UpdateSystemContractAddress is a paid mutator transaction binding the contract method 0xc835d7cc.
//
// Solidity: function updateSystemContractAddress(address addr) returns()
func (_ZRC20 *ZRC20Transactor) UpdateSystemContractAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "updateSystemContractAddress", addr)
}

// UpdateSystemContractAddress is a paid mutator transaction binding the contract method 0xc835d7cc.
//
// Solidity: function updateSystemContractAddress(address addr) returns()
func (_ZRC20 *ZRC20Session) UpdateSystemContractAddress(addr common.Address) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateSystemContractAddress(&_ZRC20.TransactOpts, addr)
}

// UpdateSystemContractAddress is a paid mutator transaction binding the contract method 0xc835d7cc.
//
// Solidity: function updateSystemContractAddress(address addr) returns()
func (_ZRC20 *ZRC20TransactorSession) UpdateSystemContractAddress(addr common.Address) (*types.Transaction, error) {
	return _ZRC20.Contract.UpdateSystemContractAddress(&_ZRC20.TransactOpts, addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Transactor) Withdraw(opts *bind.TransactOpts, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20Session) Withdraw(to []byte, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Withdraw(&_ZRC20.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_ZRC20 *ZRC20TransactorSession) Withdraw(to []byte, amount *big.Int) (*types.Transaction, error) {
	return _ZRC20.Contract.Withdraw(&_ZRC20.TransactOpts, to, amount)
}

// ZRC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZRC20 contract.
type ZRC20ApprovalIterator struct {
	Event *ZRC20Approval // Event containing the contract specifics and raw log

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
func (it *ZRC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20Approval)
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
		it.Event = new(ZRC20Approval)
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
func (it *ZRC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20Approval represents a Approval event raised by the ZRC20 contract.
type ZRC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZRC20 *ZRC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ZRC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZRC20ApprovalIterator{contract: _ZRC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ZRC20 *ZRC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZRC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20Approval)
				if err := _ZRC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ZRC20 *ZRC20Filterer) ParseApproval(log types.Log) (*ZRC20Approval, error) {
	event := new(ZRC20Approval)
	if err := _ZRC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ZRC20 contract.
type ZRC20DepositIterator struct {
	Event *ZRC20Deposit // Event containing the contract specifics and raw log

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
func (it *ZRC20DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20Deposit)
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
		it.Event = new(ZRC20Deposit)
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
func (it *ZRC20DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20Deposit represents a Deposit event raised by the ZRC20 contract.
type ZRC20Deposit struct {
	From  []byte
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab3.
//
// Solidity: event Deposit(bytes from, address indexed to, uint256 value)
func (_ZRC20 *ZRC20Filterer) FilterDeposit(opts *bind.FilterOpts, to []common.Address) (*ZRC20DepositIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "Deposit", toRule)
	if err != nil {
		return nil, err
	}
	return &ZRC20DepositIterator{contract: _ZRC20.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab3.
//
// Solidity: event Deposit(bytes from, address indexed to, uint256 value)
func (_ZRC20 *ZRC20Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ZRC20Deposit, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "Deposit", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20Deposit)
				if err := _ZRC20.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab3.
//
// Solidity: event Deposit(bytes from, address indexed to, uint256 value)
func (_ZRC20 *ZRC20Filterer) ParseDeposit(log types.Log) (*ZRC20Deposit, error) {
	event := new(ZRC20Deposit)
	if err := _ZRC20.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZRC20 contract.
type ZRC20TransferIterator struct {
	Event *ZRC20Transfer // Event containing the contract specifics and raw log

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
func (it *ZRC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20Transfer)
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
		it.Event = new(ZRC20Transfer)
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
func (it *ZRC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20Transfer represents a Transfer event raised by the ZRC20 contract.
type ZRC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZRC20 *ZRC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZRC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZRC20TransferIterator{contract: _ZRC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ZRC20 *ZRC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZRC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20Transfer)
				if err := _ZRC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ZRC20 *ZRC20Filterer) ParseTransfer(log types.Log) (*ZRC20Transfer, error) {
	event := new(ZRC20Transfer)
	if err := _ZRC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20UpdatedGasLimitIterator is returned from FilterUpdatedGasLimit and is used to iterate over the raw logs and unpacked data for UpdatedGasLimit events raised by the ZRC20 contract.
type ZRC20UpdatedGasLimitIterator struct {
	Event *ZRC20UpdatedGasLimit // Event containing the contract specifics and raw log

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
func (it *ZRC20UpdatedGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20UpdatedGasLimit)
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
		it.Event = new(ZRC20UpdatedGasLimit)
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
func (it *ZRC20UpdatedGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20UpdatedGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20UpdatedGasLimit represents a UpdatedGasLimit event raised by the ZRC20 contract.
type ZRC20UpdatedGasLimit struct {
	GasLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGasLimit is a free log retrieval operation binding the contract event 0xff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a.
//
// Solidity: event UpdatedGasLimit(uint256 gasLimit)
func (_ZRC20 *ZRC20Filterer) FilterUpdatedGasLimit(opts *bind.FilterOpts) (*ZRC20UpdatedGasLimitIterator, error) {

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "UpdatedGasLimit")
	if err != nil {
		return nil, err
	}
	return &ZRC20UpdatedGasLimitIterator{contract: _ZRC20.contract, event: "UpdatedGasLimit", logs: logs, sub: sub}, nil
}

// WatchUpdatedGasLimit is a free log subscription operation binding the contract event 0xff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a.
//
// Solidity: event UpdatedGasLimit(uint256 gasLimit)
func (_ZRC20 *ZRC20Filterer) WatchUpdatedGasLimit(opts *bind.WatchOpts, sink chan<- *ZRC20UpdatedGasLimit) (event.Subscription, error) {

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "UpdatedGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20UpdatedGasLimit)
				if err := _ZRC20.contract.UnpackLog(event, "UpdatedGasLimit", log); err != nil {
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

// ParseUpdatedGasLimit is a log parse operation binding the contract event 0xff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a.
//
// Solidity: event UpdatedGasLimit(uint256 gasLimit)
func (_ZRC20 *ZRC20Filterer) ParseUpdatedGasLimit(log types.Log) (*ZRC20UpdatedGasLimit, error) {
	event := new(ZRC20UpdatedGasLimit)
	if err := _ZRC20.contract.UnpackLog(event, "UpdatedGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20UpdatedGatewayIterator is returned from FilterUpdatedGateway and is used to iterate over the raw logs and unpacked data for UpdatedGateway events raised by the ZRC20 contract.
type ZRC20UpdatedGatewayIterator struct {
	Event *ZRC20UpdatedGateway // Event containing the contract specifics and raw log

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
func (it *ZRC20UpdatedGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20UpdatedGateway)
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
		it.Event = new(ZRC20UpdatedGateway)
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
func (it *ZRC20UpdatedGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20UpdatedGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20UpdatedGateway represents a UpdatedGateway event raised by the ZRC20 contract.
type ZRC20UpdatedGateway struct {
	Gateway common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGateway is a free log retrieval operation binding the contract event 0x88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a387.
//
// Solidity: event UpdatedGateway(address gateway)
func (_ZRC20 *ZRC20Filterer) FilterUpdatedGateway(opts *bind.FilterOpts) (*ZRC20UpdatedGatewayIterator, error) {

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "UpdatedGateway")
	if err != nil {
		return nil, err
	}
	return &ZRC20UpdatedGatewayIterator{contract: _ZRC20.contract, event: "UpdatedGateway", logs: logs, sub: sub}, nil
}

// WatchUpdatedGateway is a free log subscription operation binding the contract event 0x88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a387.
//
// Solidity: event UpdatedGateway(address gateway)
func (_ZRC20 *ZRC20Filterer) WatchUpdatedGateway(opts *bind.WatchOpts, sink chan<- *ZRC20UpdatedGateway) (event.Subscription, error) {

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "UpdatedGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20UpdatedGateway)
				if err := _ZRC20.contract.UnpackLog(event, "UpdatedGateway", log); err != nil {
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

// ParseUpdatedGateway is a log parse operation binding the contract event 0x88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a387.
//
// Solidity: event UpdatedGateway(address gateway)
func (_ZRC20 *ZRC20Filterer) ParseUpdatedGateway(log types.Log) (*ZRC20UpdatedGateway, error) {
	event := new(ZRC20UpdatedGateway)
	if err := _ZRC20.contract.UnpackLog(event, "UpdatedGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20UpdatedProtocolFlatFeeIterator is returned from FilterUpdatedProtocolFlatFee and is used to iterate over the raw logs and unpacked data for UpdatedProtocolFlatFee events raised by the ZRC20 contract.
type ZRC20UpdatedProtocolFlatFeeIterator struct {
	Event *ZRC20UpdatedProtocolFlatFee // Event containing the contract specifics and raw log

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
func (it *ZRC20UpdatedProtocolFlatFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20UpdatedProtocolFlatFee)
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
		it.Event = new(ZRC20UpdatedProtocolFlatFee)
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
func (it *ZRC20UpdatedProtocolFlatFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20UpdatedProtocolFlatFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20UpdatedProtocolFlatFee represents a UpdatedProtocolFlatFee event raised by the ZRC20 contract.
type ZRC20UpdatedProtocolFlatFee struct {
	ProtocolFlatFee *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdatedProtocolFlatFee is a free log retrieval operation binding the contract event 0xef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f.
//
// Solidity: event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) FilterUpdatedProtocolFlatFee(opts *bind.FilterOpts) (*ZRC20UpdatedProtocolFlatFeeIterator, error) {

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "UpdatedProtocolFlatFee")
	if err != nil {
		return nil, err
	}
	return &ZRC20UpdatedProtocolFlatFeeIterator{contract: _ZRC20.contract, event: "UpdatedProtocolFlatFee", logs: logs, sub: sub}, nil
}

// WatchUpdatedProtocolFlatFee is a free log subscription operation binding the contract event 0xef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f.
//
// Solidity: event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) WatchUpdatedProtocolFlatFee(opts *bind.WatchOpts, sink chan<- *ZRC20UpdatedProtocolFlatFee) (event.Subscription, error) {

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "UpdatedProtocolFlatFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20UpdatedProtocolFlatFee)
				if err := _ZRC20.contract.UnpackLog(event, "UpdatedProtocolFlatFee", log); err != nil {
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

// ParseUpdatedProtocolFlatFee is a log parse operation binding the contract event 0xef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f.
//
// Solidity: event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) ParseUpdatedProtocolFlatFee(log types.Log) (*ZRC20UpdatedProtocolFlatFee, error) {
	event := new(ZRC20UpdatedProtocolFlatFee)
	if err := _ZRC20.contract.UnpackLog(event, "UpdatedProtocolFlatFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20UpdatedSystemContractIterator is returned from FilterUpdatedSystemContract and is used to iterate over the raw logs and unpacked data for UpdatedSystemContract events raised by the ZRC20 contract.
type ZRC20UpdatedSystemContractIterator struct {
	Event *ZRC20UpdatedSystemContract // Event containing the contract specifics and raw log

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
func (it *ZRC20UpdatedSystemContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20UpdatedSystemContract)
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
		it.Event = new(ZRC20UpdatedSystemContract)
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
func (it *ZRC20UpdatedSystemContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20UpdatedSystemContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20UpdatedSystemContract represents a UpdatedSystemContract event raised by the ZRC20 contract.
type ZRC20UpdatedSystemContract struct {
	SystemContract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSystemContract is a free log retrieval operation binding the contract event 0xd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae.
//
// Solidity: event UpdatedSystemContract(address systemContract)
func (_ZRC20 *ZRC20Filterer) FilterUpdatedSystemContract(opts *bind.FilterOpts) (*ZRC20UpdatedSystemContractIterator, error) {

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "UpdatedSystemContract")
	if err != nil {
		return nil, err
	}
	return &ZRC20UpdatedSystemContractIterator{contract: _ZRC20.contract, event: "UpdatedSystemContract", logs: logs, sub: sub}, nil
}

// WatchUpdatedSystemContract is a free log subscription operation binding the contract event 0xd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae.
//
// Solidity: event UpdatedSystemContract(address systemContract)
func (_ZRC20 *ZRC20Filterer) WatchUpdatedSystemContract(opts *bind.WatchOpts, sink chan<- *ZRC20UpdatedSystemContract) (event.Subscription, error) {

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "UpdatedSystemContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20UpdatedSystemContract)
				if err := _ZRC20.contract.UnpackLog(event, "UpdatedSystemContract", log); err != nil {
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

// ParseUpdatedSystemContract is a log parse operation binding the contract event 0xd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae.
//
// Solidity: event UpdatedSystemContract(address systemContract)
func (_ZRC20 *ZRC20Filterer) ParseUpdatedSystemContract(log types.Log) (*ZRC20UpdatedSystemContract, error) {
	event := new(ZRC20UpdatedSystemContract)
	if err := _ZRC20.contract.UnpackLog(event, "UpdatedSystemContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZRC20WithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ZRC20 contract.
type ZRC20WithdrawalIterator struct {
	Event *ZRC20Withdrawal // Event containing the contract specifics and raw log

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
func (it *ZRC20WithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZRC20Withdrawal)
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
		it.Event = new(ZRC20Withdrawal)
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
func (it *ZRC20WithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZRC20WithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZRC20Withdrawal represents a Withdrawal event raised by the ZRC20 contract.
type ZRC20Withdrawal struct {
	From            common.Address
	To              []byte
	Value           *big.Int
	GasFee          *big.Int
	ProtocolFlatFee *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) FilterWithdrawal(opts *bind.FilterOpts, from []common.Address) (*ZRC20WithdrawalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ZRC20.contract.FilterLogs(opts, "Withdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return &ZRC20WithdrawalIterator{contract: _ZRC20.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ZRC20Withdrawal, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _ZRC20.contract.WatchLogs(opts, "Withdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZRC20Withdrawal)
				if err := _ZRC20.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee)
func (_ZRC20 *ZRC20Filterer) ParseWithdrawal(log types.Log) (*ZRC20Withdrawal, error) {
	event := new(ZRC20Withdrawal)
	if err := _ZRC20.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
