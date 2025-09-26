// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

// TODO: https://github.com/zeta-chain/protocol-contracts/issues/384

// import "../../contracts/evm/ERC20Custody.sol";
// import "../../contracts/evm/GatewayEVM.sol";
// import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
// import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
// import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";
// import "test/utils/TestERC20.sol";

// contract ERC20CustodyEchidnaTest is ERC20Custody {
// using SafeERC20 for IERC20;

// TestERC20 public testERC20;
// address public echidnaCaller = msg.sender;

// address proxy = Upgrades.deployUUPSProxy(
// "GatewayEVM.sol", abi.encodeCall(GatewayEVM.initialize, (echidnaCaller, address(0x123), address(0x123)))
//);
// GatewayEVM testGateway = GatewayEVM(proxy);

// constructor() ERC20Custody(address(testGateway), echidnaCaller, echidnaCaller) {
// testERC20 = new TestERC20("test", "TEST");
// testGateway.setCustody(address(this));
//}

// function testWithdrawAndCall(address to, uint256 amount, bytes calldata data) public {
// // mint more than amount
// testERC20.mint(address(this), amount + 5);
// // transfer overhead to gateway
// testERC20.transfer(address(testGateway), 5);

// withdrawAndCall(address(testERC20), to, amount, data);

// // Assertion to ensure no ERC20 tokens are left in the gateway contract after execution
// assert(testERC20.balanceOf(address(gateway)) == 0);
//}
//}
