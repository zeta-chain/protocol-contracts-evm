// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

// TODO: https://github.com/zeta-chain/protocol-contracts-evm/issues/384

// import "../../contracts/evm/ERC20Custody.sol";
// import "../../contracts/evm/GatewayEVM.sol";
// import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
// import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
// import "test/utils/TestERC20.sol";

// contract GatewayEVMEchidnaTest is GatewayEVM {
//     using SafeERC20 for IERC20;

//     TestERC20 public testERC20;
//     address public echidnaCaller = msg.sender;

//     constructor() {
//         tssAddress = echidnaCaller;
//         zetaConnector = address(0x123);
//         testERC20 = new TestERC20("test", "TEST");
//         custody = address(new ERC20Custody(address(this), tssAddress, address(this)));
//     }

//     function testExecuteWithERC20(address to, uint256 amount, bytes calldata data) public {
//         testERC20.mint(address(this), amount);

//         executeWithERC20(address(testERC20), to, amount, data);

//         // Assertion to ensure no ERC20 tokens are left in the contract after execution
//         assert(testERC20.balanceOf(address(this)) == 0);
//     }
// }
