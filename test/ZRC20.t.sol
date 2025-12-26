// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "../contracts/zevm/SystemContract.sol";
import "./utils/WZETA.sol";

import "../contracts/zevm/GatewayZEVM.sol";
import "../contracts/zevm/ZRC20.sol";
import "../contracts/zevm/interfaces/IGatewayZEVM.sol";
import "../contracts/zevm/interfaces/IZRC20.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract ZRC20Test is Test, ZRC20Errors {
    ZRC20 zrc20;
    SystemContract systemContract;
    GatewayZEVM gateway;
    address payable proxy;
    WETH9 zetaToken;

    address owner;
    address addr1;
    address protocolAddress;

    function setUp() public {
        owner = address(this);
        addr1 = address(0x1234);

        zetaToken = new WETH9();

        proxy = payable(Upgrades.deployUUPSProxy(
                "GatewayZEVM.sol", abi.encodeCall(GatewayZEVM.initialize, (address(zetaToken), owner))
            ));
        gateway = GatewayZEVM(proxy);
        protocolAddress = gateway.PROTOCOL_ADDRESS();

        vm.startPrank(protocolAddress);
        systemContract = new SystemContract(address(0), address(0), address(0));
        zrc20 = new ZRC20("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);
        vm.deal(protocolAddress, 1_000_000_000);
        vm.deal(proxy, 1_000_000_000);
        zrc20.deposit(owner, 100_000);
        vm.stopPrank();
    }

    function testZRC20BasicInfo() public {
        // silence view function warning
        zrc20 = zrc20;

        string memory name = zrc20.name();
        assertEq("TOKEN", name);

        string memory symbol = zrc20.symbol();
        assertEq("TKN", symbol);

        uint8 decimals = zrc20.decimals();
        assertEq(18, decimals);

        uint256 totalSupply = zrc20.totalSupply();
        assertEq(100_000, totalSupply);
    }

    function testUpdateNameAndSymbol() public {
        assertEq("TOKEN", zrc20.name());
        assertEq("TKN", zrc20.symbol());

        vm.prank(protocolAddress);
        zrc20.setName("TOKEN2");
        vm.prank(protocolAddress);
        zrc20.setSymbol("TKN2");

        assertEq("TOKEN2", zrc20.name());
        assertEq("TKN2", zrc20.symbol());
    }

    function testUpdateNameAndSymbolFailsIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        zrc20.setName("TOKEN2");

        vm.expectRevert(CallerIsNotFungibleModule.selector);
        zrc20.setSymbol("TKN2");
    }

    function testTransfer() public {
        uint256 balanceStart = zrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 amount = 50_000;
        zrc20.transfer(addr1, amount);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(amount, balanceAfter);
    }

    function testTransferFailsIfNoBalance() public {
        uint256 balanceStart = zrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 amount = 500_000;
        zrc20.approve(owner, amount);
        vm.expectRevert(LowBalance.selector);
        zrc20.transfer(addr1, amount);
    }

    function testTransferFailsIfRecipientIsZeroAddress() public {
        uint256 amount = 500_000;
        zrc20.approve(owner, amount);
        vm.expectRevert(ZeroAddress.selector);
        zrc20.transfer(address(0), amount);
    }

    function testTransferFrom() public {
        uint256 balanceStart = zrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 amount = 50_000;
        zrc20.approve(owner, amount);
        zrc20.transferFrom(owner, addr1, amount);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(amount, balanceAfter);
    }

    function testTransferFromFailsIfNoAllowance() public {
        uint256 balanceStart = zrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 allowance = zrc20.allowance(owner, owner);
        assertEq(0, allowance);

        uint256 amount = 50_000;
        vm.expectRevert(LowAllowance.selector);
        zrc20.transferFrom(owner, addr1, amount);
    }

    function testTransferFromFailsIfNoBalance() public {
        uint256 balanceStart = zrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 amount = 500_000;
        zrc20.approve(owner, amount);
        vm.expectRevert(LowBalance.selector);
        zrc20.transferFrom(owner, addr1, amount);
    }

    function testTransferFromFailsIfSenderIsZeroAddress() public {
        uint256 balanceStart = zrc20.balanceOf(addr1);
        assertEq(0, balanceStart);

        uint256 amount = 500_000;
        zrc20.approve(owner, amount);
        vm.expectRevert(ZeroAddress.selector);
        zrc20.transferFrom(address(0), addr1, amount);
    }

    function testTransferFromFailsIfRecipientIsZeroAddress() public {
        uint256 amount = 500_000;
        zrc20.approve(owner, amount);
        vm.expectRevert(ZeroAddress.selector);
        zrc20.transferFrom(owner, address(0), amount);
    }

    function testBurn() public {
        uint256 balanceStart = zrc20.balanceOf(owner);
        assertEq(100_000, balanceStart);
        uint256 totalSupplyStart = zrc20.totalSupply();
        assertEq(100_000, totalSupplyStart);

        zrc20.burn(50_000);
        uint256 balanceAfter = zrc20.balanceOf(owner);
        assertEq(50_000, balanceAfter);
        uint256 totalSupplyAfter = zrc20.totalSupply();
        assertEq(50_000, totalSupplyAfter);
    }

    function testApproveFailsIfRecipientIsZeroAddress() public {
        vm.expectRevert(ZeroAddress.selector);
        zrc20.approve(address(0), 10);
    }

    function testBurnFailsIfNoBalance() public {
        vm.expectRevert(LowBalance.selector);
        zrc20.burn(150_000);
    }

    function testDeposit() public {
        uint256 totalSupplyStart = zrc20.totalSupply();
        assertEq(100_000, totalSupplyStart);

        vm.prank(proxy);
        zrc20.deposit(owner, 100_000);

        uint256 totalSupplyEnd = zrc20.totalSupply();
        assertEq(200_000, totalSupplyEnd);
    }

    function testDepositFailsIfRecipientIsZeroAddress() public {
        vm.prank(proxy);
        vm.expectRevert(ZeroAddress.selector);
        zrc20.deposit(address(0), 100_000);
    }

    function testWithdrawGasFee() public {
        vm.prank(protocolAddress);
        uint256 gasLimit = 10;
        uint256 protocolFlatFee = 10;

        zrc20.updateGasLimit(10);

        vm.prank(protocolAddress);
        zrc20.updateProtocolFlatFee(10);

        (address gasZRC20, uint256 gasFee) = zrc20.withdrawGasFee();
        assertEq(address(zrc20), gasZRC20);
        assertEq(gasLimit + protocolFlatFee, gasFee);
    }

    function testWithdrawGasFeeFailsIfGasCoinNotSetForChainId() public {
        vm.prank(protocolAddress);
        systemContract.setGasCoinZRC20(1, address(0));

        vm.expectRevert(ZeroGasCoin.selector);
        zrc20.withdrawGasFee();
    }

    function testWithdrawGasFeeFailsIfGasPriceNotSetForChainId() public {
        vm.prank(protocolAddress);
        systemContract.setGasPrice(1, 0);

        vm.expectRevert(ZeroGasPrice.selector);
        zrc20.withdrawGasFee();
    }

    function testWithdraw() public {
        uint256 gasLimit = 10;
        uint256 protocolFlatFee = 10;

        vm.prank(protocolAddress);
        zrc20.updateGasLimit(gasLimit);

        vm.prank(protocolAddress);
        zrc20.updateProtocolFlatFee(protocolFlatFee);

        uint256 balanceStart = zrc20.balanceOf(owner);
        assertEq(100_000, balanceStart);
        uint256 totalSupplyStart = zrc20.totalSupply();
        assertEq(100_000, totalSupplyStart);

        uint256 protocolAddressBalanceStart = zrc20.balanceOf(protocolAddress);

        zrc20.approve(address(zrc20), 50_000);
        zrc20.withdraw(abi.encodePacked(addr1), 50_000);

        uint256 protocolAddressBalanceAfter = zrc20.balanceOf(protocolAddress);
        assertEq(protocolAddressBalanceStart + gasLimit + protocolFlatFee, protocolAddressBalanceAfter);

        uint256 balanceAfter = zrc20.balanceOf(owner);
        assertEq(50_000 - gasLimit - protocolFlatFee, balanceAfter);
        uint256 totalSupplyAfter = zrc20.totalSupply();
        assertEq(50_000, totalSupplyAfter);
    }

    function testWithdrawFailsIfNoBalance() public {
        uint256 gasLimit = 10;
        uint256 protocolFlatFee = 100_000_000;

        vm.prank(protocolAddress);
        zrc20.updateGasLimit(gasLimit);

        vm.prank(protocolAddress);
        zrc20.updateProtocolFlatFee(protocolFlatFee);

        zrc20.approve(address(zrc20), 200_000_000);

        vm.expectRevert(LowBalance.selector);
        zrc20.withdraw(abi.encodePacked(addr1), 100);
    }

    function testWithdrawFailsIfNoAllowance() public {
        uint256 gasLimit = 10;
        uint256 protocolFlatFee = 10;

        vm.prank(protocolAddress);
        zrc20.updateGasLimit(gasLimit);

        vm.prank(protocolAddress);
        zrc20.updateProtocolFlatFee(protocolFlatFee);

        zrc20.approve(address(zrc20), 0);

        vm.expectRevert();
        zrc20.withdraw(abi.encodePacked(addr1), 1);
    }

    function testDepositFailsIfSenderIsNotGateway() public {
        vm.expectRevert(InvalidSender.selector);
        zrc20.deposit(owner, 100_000);
    }

    function testUpdateSystemContractAddress() public {
        vm.prank(protocolAddress);
        zrc20.updateSystemContractAddress(address(0x3211));
        assertEq(zrc20.SYSTEM_CONTRACT_ADDRESS(), address(0x3211));
    }

    function testUpdateSystemContractAddressFailsIfZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        zrc20.updateSystemContractAddress(address(0));
    }

    function testUpdateSystemContractAddressFailsIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        zrc20.updateSystemContractAddress(address(0x3211));
    }

    function testUpdateGatewayAddress() public {
        vm.prank(protocolAddress);
        zrc20.updateGatewayAddress(address(0x3211));
        assertEq(zrc20.gatewayAddress(), address(0x3211));
    }

    function testUpdateGatewayAddressFailsIfZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        zrc20.updateGatewayAddress(address(0));
    }

    function testUpdateGatewayAddressFailsIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        zrc20.updateGatewayAddress(address(0x3211));
    }

    function testUpdateGasLimit() public {
        vm.prank(protocolAddress);
        zrc20.updateGasLimit(10);
        assertEq(10, zrc20.GAS_LIMIT());
    }

    function testUpdateGasLimitFailsIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        zrc20.updateGasLimit(10);
    }

    function testUpdateProtocolFlatFee() public {
        vm.prank(protocolAddress);
        zrc20.updateProtocolFlatFee(10);
        assertEq(10, zrc20.PROTOCOL_FLAT_FEE());
    }

    function testUpdateProtocolFlatFeeFailsIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotFungibleModule.selector);
        zrc20.updateProtocolFlatFee(10);
    }
}
