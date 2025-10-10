// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "./utils/ReceiverEVM.sol";

import "./utils/TestERC20.sol";
import "./utils/upgrades/ERC20CustodyUpgradeTest.sol";

import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

import "../contracts/evm/ERC20Custody.sol";
import "../contracts/evm/GatewayEVM.sol";
import "../contracts/evm/ZetaConnectorNonNative.sol";
import "../contracts/evm/interfaces/IERC20Custody.sol";
import "../contracts/evm/interfaces/IGatewayEVM.sol";
import "./utils/IReceiverEVM.sol";

contract ERC20CustodyTest is Test, IGatewayEVMErrors, IGatewayEVMEvents, IReceiverEVMEvents, IERC20CustodyEvents {
    using SafeERC20 for IERC20;

    GatewayEVM gateway;
    ReceiverEVM receiver;
    ERC20Custody custody;
    ZetaConnectorNonNative zetaConnector;
    TestERC20 token;
    TestERC20 zeta;
    address owner;
    address destination;
    address tssAddress;
    address foo;
    RevertContext revertContext;
    MessageContextV2 arbitraryCallMessageContext =
        MessageContextV2({ sender: address(0), asset: address(0), amount: 0 });

    error EnforcedPause();
    error NotWhitelisted();
    error AccessControlUnauthorizedAccount(address account, bytes32 neededRole);
    error LegacyMethodsNotSupported();

    event WithdrawnV2(address indexed to, address indexed token, uint256 amount);

    bytes32 public constant TSS_ROLE = keccak256("TSS_ROLE");
    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
    bytes32 public constant ASSET_HANDLER_ROLE = keccak256("ASSET_HANDLER_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant WHITELISTER_ROLE = keccak256("WHITELISTER_ROLE");
    bytes32 public constant DEFAULT_ADMIN_ROLE = 0x00;

    function setUp() public {
        owner = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);
        foo = address(0x9876);

        token = new TestERC20("test", "TTK");
        zeta = new TestERC20("zeta", "ZETA");

        address proxy = Upgrades.deployUUPSProxy(
            "GatewayEVM.sol", abi.encodeCall(GatewayEVM.initialize, (tssAddress, address(zeta), owner))
        );
        gateway = GatewayEVM(proxy);
        proxy = Upgrades.deployUUPSProxy(
            "ERC20Custody.sol", abi.encodeCall(ERC20Custody.initialize, (address(gateway), tssAddress, owner))
        );
        custody = ERC20Custody(proxy);
        proxy = Upgrades.deployUUPSProxy(
            "ZetaConnectorNonNative.sol",
            abi.encodeCall(ZetaConnectorNonNative.initialize, (address(gateway), address(zeta), tssAddress, owner))
        );
        zetaConnector = ZetaConnectorNonNative(proxy);
        receiver = new ReceiverEVM();

        vm.deal(tssAddress, 1 ether);

        vm.startPrank(owner);
        gateway.setCustody(address(custody));
        gateway.setConnector(address(zetaConnector));

        custody.whitelist(address(token));
        vm.stopPrank();

        token.mint(owner, 1_000_000);
        zeta.mint(owner, 1_000_000);
        token.transfer(address(custody), 500_000);

        vm.deal(tssAddress, 1 ether);

        revertContext = RevertContext({ sender: owner, asset: address(token), amount: 1, revertMessage: "" });
    }

    function testTSSUpgrade() public {
        address newTSSAddress = address(0x4321);

        bool newTSSAddressWithdrawerRole = custody.hasRole(WITHDRAWER_ROLE, newTSSAddress);
        assertFalse(newTSSAddressWithdrawerRole);
        bool newTSSAddressWhitelisterRole = custody.hasRole(WHITELISTER_ROLE, newTSSAddress);
        assertFalse(newTSSAddressWhitelisterRole);

        bool oldTSSAddressHasWithdrawerRole = custody.hasRole(WITHDRAWER_ROLE, tssAddress);
        assertTrue(oldTSSAddressHasWithdrawerRole);
        bool oldTSSAddressHasWhitelisterRole = custody.hasRole(WHITELISTER_ROLE, tssAddress);
        assertTrue(oldTSSAddressHasWhitelisterRole);

        vm.startPrank(owner);
        vm.expectEmit(true, true, true, true, address(custody));
        emit UpdatedCustodyTSSAddress(tssAddress, newTSSAddress);
        custody.updateTSSAddress(newTSSAddress);
        assertEq(newTSSAddress, custody.tssAddress());

        newTSSAddressWithdrawerRole = custody.hasRole(WITHDRAWER_ROLE, newTSSAddress);
        assertTrue(newTSSAddressWithdrawerRole);
        newTSSAddressWhitelisterRole = custody.hasRole(WHITELISTER_ROLE, newTSSAddress);
        assertTrue(newTSSAddressWhitelisterRole);

        oldTSSAddressHasWithdrawerRole = custody.hasRole(WITHDRAWER_ROLE, tssAddress);
        assertFalse(oldTSSAddressHasWithdrawerRole);
        oldTSSAddressHasWhitelisterRole = custody.hasRole(WHITELISTER_ROLE, tssAddress);
        assertFalse(oldTSSAddressHasWhitelisterRole);
    }

    function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() public {
        vm.startPrank(tssAddress);
        vm.expectRevert(
            abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, tssAddress, DEFAULT_ADMIN_ROLE)
        );
        custody.updateTSSAddress(owner);
    }

    function testTSSUpgradeFailsIfZeroAddress() public {
        vm.startPrank(owner);
        vm.expectRevert(ZeroAddress.selector);
        custody.updateTSSAddress(address(0));
    }

    function testWhitelistFailsIfZeroAddress() public {
        vm.expectRevert(ZeroAddress.selector);
        custody.whitelist(address(0));
    }

    function testUnwhitelistFailsIfZeroAddress() public {
        vm.expectRevert(ZeroAddress.selector);
        custody.unwhitelist(address(0));
    }

    function testWhitelistFailsIfSenderIsNotWhitelister() public {
        vm.prank(address(0x123));
        vm.expectRevert(
            abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, address(0x123), WHITELISTER_ROLE)
        );
        custody.whitelist(address(zeta));
    }

    function testUnwhitelistFailsIfSenderIsNotWhitelister() public {
        vm.prank(address(0x123));
        vm.expectRevert(
            abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, address(0x123), WHITELISTER_ROLE)
        );
        custody.unwhitelist(address(zeta));
    }

    function testWhitelist() public {
        bool whitelisted = custody.whitelisted(address(zeta));
        assertEq(false, whitelisted);

        vm.expectEmit(true, true, true, true, address(custody));
        emit Whitelisted(address(zeta));
        vm.prank(owner);
        custody.whitelist(address(zeta));

        whitelisted = custody.whitelisted(address(zeta));
        assertEq(true, whitelisted);
    }

    function testUnwhitelist() public {
        bool whitelisted = custody.whitelisted(address(token));
        assertEq(true, whitelisted);

        vm.expectEmit(true, true, true, true, address(custody));
        emit Unwhitelisted(address(token));
        vm.prank(owner);
        custody.unwhitelist(address(token));

        whitelisted = custody.whitelisted(address(token));
        assertEq(false, whitelisted);
    }

    function testForwardCallToReceiveERC20ThroughCustody() public {
        uint256 amount = 100_000;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount, address(token), destination);
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawnAndCalled(address(receiver), address(token), amount, data);
        vm.prank(tssAddress);
        custody.withdrawAndCall(arbitraryCallMessageContext, address(receiver), address(token), amount, data);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, amount);

        // Verify that the remaining tokens were refunded to the Custody contract
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testForwardCallToReceiveERC20ThroughCustodyTogglePause() public {
        vm.prank(foo);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, foo, PAUSER_ROLE));
        custody.pause();

        vm.prank(foo);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, foo, PAUSER_ROLE));
        gateway.unpause();

        vm.prank(tssAddress);
        custody.pause();

        uint256 amount = 100_000;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);

        vm.expectRevert(EnforcedPause.selector);
        vm.prank(tssAddress);
        custody.withdrawAndCall(arbitraryCallMessageContext, address(receiver), address(token), amount, data);

        vm.prank(owner);
        custody.unpause();

        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount, address(token), destination);
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawnAndCalled(address(receiver), address(token), amount, data);
        vm.prank(tssAddress);
        custody.withdrawAndCall(arbitraryCallMessageContext, address(receiver), address(token), amount, data);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, amount);

        // Verify that the remaining tokens were refunded to the Custody contract
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testForwardCallToReceiveERC20ThroughCustodyFailsIfSenderIsNotWithdrawer() public {
        uint256 amount = 100_000;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, WITHDRAWER_ROLE));
        custody.withdrawAndCall(arbitraryCallMessageContext, address(receiver), address(token), amount, data);
    }

    function testForwardCallToReceiveERC20ThroughCustodyFailsIfAmountIs0() public {
        uint256 amount = 0;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);

        vm.prank(tssAddress);
        vm.expectRevert(InsufficientERC20Amount.selector);
        custody.withdrawAndCall(arbitraryCallMessageContext, address(receiver), address(token), amount, data);
    }

    function testForwardCallToReceiveERC20ThroughCustodyFailsIfReceiverIsZeroAddress() public {
        uint256 amount = 1;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);

        vm.prank(tssAddress);
        vm.expectRevert(ZeroAddress.selector);
        custody.withdrawAndCall(arbitraryCallMessageContext, address(0), address(token), amount, data);
    }

    function testForwardCallToReceiveERC20PartialThroughCustody() public {
        uint256 amount = 100_000;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20Partial(uint256,address,address)", amount, address(token), destination);
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedERC20(address(gateway), amount / 2, address(token), destination);
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawnAndCalled(address(receiver), address(token), amount, data);
        vm.prank(tssAddress);
        custody.withdrawAndCall(arbitraryCallMessageContext, address(receiver), address(token), amount, data);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, amount / 2);

        // Verify that the remaining tokens were refunded to the Custody contract
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount / 2);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testForwardCallToReceiveOnCallThroughCustody() public {
        uint256 amount = 100_000;
        address sender = address(0x123);
        bytes memory message = bytes("1");
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedOnCallV2(sender, address(token), amount, message);
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawnAndCalled(address(receiver), address(token), amount, message);
        vm.prank(tssAddress);
        custody.withdrawAndCall(
            MessageContextV2({ sender: sender, asset: address(token), amount: amount }),
            address(receiver),
            address(token),
            amount,
            message
        );

        // Verify that the tokens were not transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, 0);

        // Verify that the remaining tokens were refunded to the Custody contract
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testForwardCallToReceiveOnCallThroughCustodyNotAllowedWithArbitraryCall() public {
        uint256 amount = 100_000;
        address sender = address(0x123);
        bytes memory message = bytes("1");
        bytes memory data = abi.encodeWithSignature("onCall((address),bytes)", sender, message);

        vm.expectRevert(NotAllowedToCallOnCall.selector);
        vm.prank(tssAddress);
        custody.withdrawAndCall(arbitraryCallMessageContext, address(receiver), address(token), amount, data);
    }

    function testForwardCallOnRevertThroughCustodyNotAllowedWithArbitraryCall() public {
        uint256 amount = 100_000;
        bytes memory data = abi.encodeWithSignature("onRevert((address,address,uint256,bytes))");

        vm.expectRevert(NotAllowedToCallOnRevert.selector);
        vm.prank(tssAddress);
        custody.withdrawAndCall(arbitraryCallMessageContext, address(receiver), address(token), amount, data);
    }

    function testForwardCallToReceiveERC20PartialThroughCustodyFailsIfSenderIsNotWithdrawer() public {
        uint256 amount = 100_000;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20Partial(uint256,address,address)", amount, address(token), destination);

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, WITHDRAWER_ROLE));
        custody.withdrawAndCall(arbitraryCallMessageContext, address(receiver), address(token), amount, data);
    }

    function testForwardCallToReceiveERC20PartialThroughCustodyFailsIfAmountIs0() public {
        uint256 amount = 0;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20Partial(uint256,address,address)", amount, address(token), destination);

        vm.prank(tssAddress);
        vm.expectRevert(InsufficientERC20Amount.selector);
        custody.withdrawAndCall(arbitraryCallMessageContext, address(receiver), address(token), amount, data);
    }

    function testForwardCallToReceiveNoParamsThroughCustody() public {
        uint256 amount = 100_000;
        bytes memory data = abi.encodeWithSignature("receiveNoParams()");
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNoParams(address(gateway));
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawnAndCalled(address(receiver), address(token), amount, data);
        vm.prank(tssAddress);
        custody.withdrawAndCall(arbitraryCallMessageContext, address(receiver), address(token), amount, data);

        // Verify that the tokens were not transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, 0);

        // Verify that the remaining tokens were refunded to the Custody contract
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawFailsIfTokenIsNotWhitelisted() public {
        vm.startPrank(tssAddress);
        custody.unwhitelist(address(token));
        vm.expectRevert(NotWhitelisted.selector);
        custody.withdraw(destination, address(token), 1);
        vm.stopPrank();
    }

    function testWithdrawAndCallFailsIfTokenIsNotWhitelisted() public {
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", 1, address(token), destination);
        vm.startPrank(tssAddress);
        custody.unwhitelist(address(token));
        vm.expectRevert(NotWhitelisted.selector);
        custody.withdrawAndCall(arbitraryCallMessageContext, address(receiver), address(token), 1, data);
        vm.stopPrank();
    }

    function testWithdrawAndRevertFailsIfTokenIsNotWhitelisted() public {
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", 1, address(token), destination);
        vm.startPrank(tssAddress);
        custody.unwhitelist(address(token));
        vm.expectRevert(NotWhitelisted.selector);
        custody.withdrawAndRevert(address(receiver), address(token), 1, data, revertContext);
        vm.stopPrank();
    }

    function testWithdrawThroughCustody() public {
        uint256 amount = 100_000;
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(destination), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(custody));
        emit Withdrawn(destination, address(token), amount);
        vm.prank(tssAddress);
        custody.withdraw(destination, address(token), amount);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, amount);

        // Verify that the tokens were substracted from custody
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawThroughCustodyFailsIfSenderIsNotWithdrawer() public {
        uint256 amount = 100_000;

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, WITHDRAWER_ROLE));
        custody.withdraw(destination, address(token), amount);
    }

    function testWithdrawAndRevertThroughCustody() public {
        uint256 amount = 100_000;
        bytes memory data = abi.encodePacked("hello");
        uint256 balanceBefore = token.balanceOf(address(receiver));
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custody));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(gateway), amount);
        vm.expectCall(address(token), 0, transferData);
        // Verify that onRevert callback was called
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedRevert(address(gateway), revertContext);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Reverted(address(receiver), address(token), amount, data, revertContext);
        vm.expectEmit(true, true, true, true, address(custody));
        emit WithdrawnAndReverted(address(receiver), address(token), amount, data, revertContext);
        vm.prank(tssAddress);
        custody.withdrawAndRevert(address(receiver), address(token), amount, data, revertContext);

        // Verify that the tokens were transferred to the receiver address
        uint256 balanceAfter = token.balanceOf(address(receiver));
        assertEq(balanceAfter, amount);

        // Verify that zeta connector doesn't get more tokens
        uint256 balanceAfterCustody = token.balanceOf(address(custody));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount);

        // Verify that the approval was reset
        uint256 allowance = token.allowance(address(gateway), address(receiver));
        assertEq(allowance, 0);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }

    function testWithdrawAndRevertThroughCustodyFailsIfSenderIsNotWithdrawer() public {
        uint256 amount = 100_000;
        bytes memory data = abi.encodePacked("hello");

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, WITHDRAWER_ROLE));
        custody.withdrawAndRevert(address(receiver), address(token), amount, data, revertContext);
    }

    function testWithdrawAndRevertThroughCustodyFailsIfAmountIs0() public {
        uint256 amount = 0;
        bytes memory data = abi.encodePacked("hello");

        vm.prank(tssAddress);
        vm.expectRevert(InsufficientERC20Amount.selector);
        custody.withdrawAndRevert(address(receiver), address(token), amount, data, revertContext);
    }

    function testWithdrawAndRevertThroughCustodyFailsIfReceiverIsZeroAddress() public {
        uint256 amount = 1;
        bytes memory data = abi.encodePacked("hello");

        vm.prank(tssAddress);
        vm.expectRevert(ZeroAddress.selector);
        custody.withdrawAndRevert(address(0), address(token), amount, data, revertContext);
    }

    function testDepositLegacy() public {
        uint256 amount = 1000;
        custody.setSupportsLegacy(true);
        token.approve(address(custody), 1_000_000);
        zeta.approve(address(custody), 1_000_000);

        uint256 custodyTokenBalanceBefore = token.balanceOf(address(custody));

        bytes memory message = abi.encodePacked("hello");

        vm.expectEmit(true, true, true, true, address(custody));
        emit Deposited(abi.encodePacked(destination), token, amount, message);
        custody.deposit(abi.encodePacked(destination), token, amount, message);

        assertEq(custodyTokenBalanceBefore + amount, token.balanceOf(address(custody)));
    }

    function testDepositLegacyFailsIfNotSupported() public {
        custody.setSupportsLegacy(false);
        token.approve(address(custody), 1_000_000);
        zeta.approve(address(custody), 1_000_000);

        bytes memory message = abi.encodePacked("hello");

        vm.expectRevert(LegacyMethodsNotSupported.selector);
        custody.deposit(abi.encodePacked(destination), token, 1000, message);
    }

    function testDepositLegacyFailsIfTokenNotWhitelisted() public {
        custody.setSupportsLegacy(true);
        custody.unwhitelist(address(token));
        token.approve(address(custody), 1_000_000);
        zeta.approve(address(custody), 1_000_000);

        bytes memory message = abi.encodePacked("hello");

        vm.expectRevert(NotWhitelisted.selector);
        custody.deposit(abi.encodePacked(destination), token, 1000, message);
    }

    function testUpgradeAndWithdraw() public {
        // upgrade
        Upgrades.upgradeProxy(address(custody), "ERC20CustodyUpgradeTest.sol", "", owner);
        ERC20CustodyUpgradeTest custodyV2 = ERC20CustodyUpgradeTest(address(custody));
        // withdraw
        uint256 amount = 100_000;
        uint256 balanceBefore = token.balanceOf(destination);
        assertEq(balanceBefore, 0);
        uint256 balanceBeforeCustody = token.balanceOf(address(custodyV2));

        bytes memory transferData = abi.encodeWithSignature("transfer(address,uint256)", address(destination), amount);
        vm.expectCall(address(token), 0, transferData);
        vm.expectEmit(true, true, true, true, address(custodyV2));
        emit WithdrawnV2(destination, address(token), amount);
        vm.prank(tssAddress);
        custodyV2.withdraw(destination, address(token), amount);

        // Verify that the tokens were transferred to the destination address
        uint256 balanceAfter = token.balanceOf(destination);
        assertEq(balanceAfter, amount);

        // Verify that the tokens were substracted from custody
        uint256 balanceAfterCustody = token.balanceOf(address(custodyV2));
        assertEq(balanceAfterCustody, balanceBeforeCustody - amount);

        // Verify that gateway doesn't hold any tokens
        uint256 balanceGateway = token.balanceOf(address(gateway));
        assertEq(balanceGateway, 0);
    }
}
