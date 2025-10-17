// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "./utils/ReceiverEVM.sol";

import "./mocks/NonReturnApprovalToken.sol";
import "./mocks/RevertOnZeroApprovalToken.sol";
import "./utils/TestERC20.sol";
import "./utils/upgrades/GatewayEVMUpgradeTest.sol";

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
import "./utils/Zeta.non-eth.sol";

contract GatewayEVMTest is Test, IGatewayEVMErrors, IGatewayEVMEvents, IReceiverEVMEvents, IERC20CustodyEvents {
    using SafeERC20 for IERC20;

    GatewayEVM gateway;
    ReceiverEVM receiver;
    ERC20Custody custody;
    ZetaConnectorNonNative zetaConnector;
    TestERC20 token;
    ZetaNonEth zeta;
    address owner;
    address destination;
    address tssAddress;
    address foo;
    RevertOptions revertOptions;
    RevertContext revertContext;
    MessageContext arbitraryCallMessageContext = MessageContext({ sender: address(0) });

    error EnforcedPause();
    error AccessControlUnauthorizedAccount(address account, bytes32 neededRole);

    event ExecutedV2(address indexed destination, uint256 value, bytes data);

    bytes32 public constant TSS_ROLE = keccak256("TSS_ROLE");
    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
    bytes32 public constant ASSET_HANDLER_ROLE = keccak256("ASSET_HANDLER_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant DEFAULT_ADMIN_ROLE = 0x00;

    function setUp() public {
        owner = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);
        foo = address(0x9876);

        token = new TestERC20("test", "TTK");

        zeta = new ZetaNonEth(tssAddress, tssAddress);
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

        vm.prank(tssAddress);
        zeta.updateTssAndConnectorAddresses(tssAddress, address(zetaConnector));
        receiver = new ReceiverEVM();

        vm.deal(tssAddress, 1 ether);

        vm.startPrank(owner);
        gateway.setCustody(address(custody));
        gateway.setConnector(address(zetaConnector));
        vm.stopPrank();

        token.mint(owner, 1_000_000);
        token.transfer(address(custody), 500_000);

        vm.deal(tssAddress, 1 ether);

        revertContext = RevertContext({ sender: owner, asset: address(token), amount: 1, revertMessage: "" });
    }

    function testTSSUpgrade() public {
        address newTSSAddress = address(0x4321);

        bool newTSSAddressHasTSSRole = gateway.hasRole(TSS_ROLE, newTSSAddress);
        assertFalse(newTSSAddressHasTSSRole);
        bool oldTSSAddressHasTSSRole = gateway.hasRole(TSS_ROLE, tssAddress);
        assertTrue(oldTSSAddressHasTSSRole);

        vm.startPrank(owner);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit UpdatedGatewayTSSAddress(tssAddress, newTSSAddress);
        gateway.updateTSSAddress(newTSSAddress);
        assertEq(newTSSAddress, gateway.tssAddress());

        newTSSAddressHasTSSRole = gateway.hasRole(TSS_ROLE, newTSSAddress);
        assertTrue(newTSSAddressHasTSSRole);
        oldTSSAddressHasTSSRole = gateway.hasRole(TSS_ROLE, tssAddress);
        assertFalse(oldTSSAddressHasTSSRole);
    }

    function testTSSUpgradeFailsIfSenderIsNotTSSUpdater() public {
        vm.startPrank(tssAddress);
        vm.expectRevert(
            abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, tssAddress, DEFAULT_ADMIN_ROLE)
        );
        gateway.updateTSSAddress(owner);
    }

    function testTSSUpgradeFailsIfZeroAddress() public {
        vm.startPrank(owner);
        vm.expectRevert(ZeroAddress.selector);
        gateway.updateTSSAddress(address(0));
    }

    function testSetCustodyFailsIfSenderIsNotAdmin() public {
        vm.startPrank(tssAddress);
        vm.expectRevert(
            abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, tssAddress, DEFAULT_ADMIN_ROLE)
        );
        gateway.setCustody(address(custody));
    }

    function testSetCustodyFailsIfSet() public {
        vm.startPrank(owner);
        vm.expectRevert(CustodyInitialized.selector);
        gateway.setCustody(address(custody));
    }

    function testSetCustodyFailsIfZero() public {
        vm.startPrank(owner);
        vm.expectRevert(ZeroAddress.selector);
        gateway.setCustody(address(0));
    }

    function testSetConnectorFailsIfSenderIsNotAdmin() public {
        vm.startPrank(tssAddress);
        vm.expectRevert(
            abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, tssAddress, DEFAULT_ADMIN_ROLE)
        );
        gateway.setConnector(address(zetaConnector));
    }

    function testSetConnectorFailsIfSet() public {
        vm.startPrank(owner);
        vm.expectRevert(ConnectorInitialized.selector);
        gateway.setConnector(address(zetaConnector));
    }

    function testSetConnectorFailsIfZero() public {
        vm.startPrank(owner);
        vm.expectRevert(ZeroAddress.selector);
        gateway.setConnector(address(0));
    }

    function testForwardCallToReceiveNonPayable() public {
        string[] memory str = new string[](1);
        str[0] = "Hello, Foundry!";
        uint256[] memory num = new uint256[](1);
        num[0] = 42;
        bool flag = true;

        bytes memory data = abi.encodeWithSignature("receiveNonPayable(string[],uint256[],bool)", str, num, flag);

        vm.expectCall(address(receiver), 0, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNonPayable(address(gateway), str, num, flag);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), 0, data);
        vm.prank(tssAddress);
        gateway.execute(arbitraryCallMessageContext, address(receiver), data);
    }

    function testForwardCallToReceiveOnCallUsingAuthCall() public {
        vm.expectEmit(true, true, true, true, address(receiver));
        address sender = address(0x123);
        emit ReceivedOnCall(sender, bytes("1"));
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), 0, bytes("1"));
        vm.prank(tssAddress);
        gateway.execute(MessageContext({ sender: sender }), address(receiver), bytes("1"));
    }

    function testForwardCallToReceiveNonPayableFailsIfSenderIsNotTSS() public {
        string[] memory str = new string[](1);
        str[0] = "Hello, Foundry!";
        uint256[] memory num = new uint256[](1);
        num[0] = 42;
        bool flag = true;
        bytes memory data = abi.encodeWithSignature("receiveNonPayable(string[],uint256[],bool)", str, num, flag);

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, TSS_ROLE));
        gateway.execute(arbitraryCallMessageContext, address(receiver), data);
    }

    function testForwardCallToReceiveNonPayableWithMsgContextFailsIfSenderIsNotTSS() public {
        string[] memory str = new string[](1);
        str[0] = "Hello, Foundry!";
        uint256[] memory num = new uint256[](1);
        num[0] = 42;
        bool flag = true;
        bytes memory data = abi.encodeWithSignature("receiveNonPayable(string[],uint256[],bool)", str, num, flag);

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, TSS_ROLE));
        gateway.execute(MessageContext({ sender: address(0x123) }), address(receiver), data);
    }

    function testForwardCallToReceivePayable() public {
        string memory str = "Hello, Foundry!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;
        assertEq(0, address(receiver).balance);

        bytes memory data = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);

        vm.expectCall(address(receiver), 1 ether, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedPayable(address(gateway), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), 1 ether, data);
        vm.prank(tssAddress);
        gateway.execute{ value: value }(arbitraryCallMessageContext, address(receiver), data);

        assertEq(value, address(receiver).balance);
    }

    function testForwardCallToReceiveNoParams() public {
        bytes memory data = abi.encodeWithSignature("receiveNoParams()");

        vm.expectCall(address(receiver), 0, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNoParams(address(gateway));
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), 0, data);
        vm.prank(tssAddress);
        gateway.execute(arbitraryCallMessageContext, address(receiver), data);
    }

    function testForwardCallToReceiveOnCallFails() public {
        bytes memory data = abi.encodeWithSignature("onCall((address),bytes)", address(123), bytes(""));

        vm.prank(tssAddress);
        vm.expectRevert(NotAllowedToCallOnCall.selector);
        gateway.execute(arbitraryCallMessageContext, address(receiver), data);
    }

    function testForwardCallToReceiveOnRevertFails() public {
        bytes memory data = abi.encodeWithSignature("onRevert((address,address,uint256,bytes))");

        vm.prank(tssAddress);
        vm.expectRevert(NotAllowedToCallOnRevert.selector);
        gateway.execute(arbitraryCallMessageContext, address(receiver), data);
    }

    function testExecuteFailsIfDestinationIsZeroAddress() public {
        bytes memory data = abi.encodeWithSignature("receiveNoParams()");

        vm.prank(tssAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.execute(arbitraryCallMessageContext, address(0), data);
    }

    function testExecuteWithMsgContextFailsIfDestinationIsZeroAddress() public {
        bytes memory data = abi.encodeWithSignature("receiveNoParams()");

        vm.prank(tssAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.execute(MessageContext({ sender: address(0x123) }), address(0), data);
    }

    function testForwardCallToReceiveNoParamsTogglePause() public {
        vm.prank(foo);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, foo, PAUSER_ROLE));
        gateway.pause();

        vm.prank(foo);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, foo, PAUSER_ROLE));
        gateway.unpause();

        vm.prank(tssAddress);
        gateway.pause();

        bytes memory data = abi.encodeWithSignature("receiveNoParams()");

        vm.expectRevert(EnforcedPause.selector);
        vm.prank(tssAddress);
        gateway.execute(arbitraryCallMessageContext, address(receiver), data);

        vm.prank(owner);
        gateway.unpause();

        vm.expectCall(address(receiver), 0, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNoParams(address(gateway));
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Executed(address(receiver), 0, data);
        vm.prank(tssAddress);
        gateway.execute(arbitraryCallMessageContext, address(receiver), data);
    }

    function testExecuteWithTokenRevertingOnZeroApproval() public {
        RevertOnZeroApprovalToken revertingToken = new RevertOnZeroApprovalToken("BNB", "BNB");
        revertingToken.mint(owner, 1_000_000);

        vm.startPrank(owner);
        custody.whitelist(address(revertingToken));
        vm.stopPrank();

        uint256 amount = 100_000;
        revertingToken.transfer(address(custody), amount);

        bytes memory data = abi.encodeWithSignature("receiveNoParams()");

        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedNoParams(address(gateway));

        vm.expectEmit(true, true, true, true, address(gateway));
        emit ExecutedWithERC20(address(revertingToken), address(receiver), amount, data);

        vm.prank(address(custody));
        gateway.executeWithERC20(arbitraryCallMessageContext, address(revertingToken), address(receiver), amount, data);
    }

    function testExecuteWithNonReturnApprovalToken() public {
        NonReturnApprovalToken nonReturnToken = new NonReturnApprovalToken("USDT", "USDT");
        nonReturnToken.mint(owner, 1_000_000);

        vm.startPrank(owner);
        custody.whitelist(address(nonReturnToken));
        vm.stopPrank();

        uint256 amount = 100_000;
        nonReturnToken.transfer(address(custody), amount);

        bytes memory data = abi.encodeWithSignature("receiveNoParams()");

        vm.prank(address(custody));
        gateway.executeWithERC20(arbitraryCallMessageContext, address(nonReturnToken), address(receiver), amount, data);
    }

    function testExecuteWithERC20FailsIfNotCustodyOrConnector() public {
        uint256 amount = 100_000;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, ASSET_HANDLER_ROLE));
        gateway.executeWithERC20(arbitraryCallMessageContext, address(token), destination, amount, data);
    }

    function testRevertWithERC20FailsIfNotCustodyOrConnector() public {
        uint256 amount = 100_000;
        bytes memory data =
            abi.encodeWithSignature("receiveERC20(uint256,address,address)", amount, address(token), destination);

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, ASSET_HANDLER_ROLE));
        gateway.revertWithERC20(address(token), destination, amount, data, revertContext);
    }

    function testExecuteRevert() public {
        uint256 value = 1 ether;
        bytes memory data = abi.encodePacked("hello");
        uint256 balanceBefore = address(receiver).balance;
        assertEq(balanceBefore, 0);

        // Verify that onRevert callback was called
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedRevert(address(gateway), revertContext);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit Reverted(address(receiver), address(0), 1 ether, data, revertContext);
        vm.prank(tssAddress);
        gateway.executeRevert{ value: value }(address(receiver), data, revertContext);

        // Verify that the tokens were transferred to the receiver address
        uint256 balanceAfter = address(receiver).balance;
        assertEq(balanceAfter, 1 ether);
    }

    function testExecuteRevertFailsIfSenderIsNotTSS() public {
        uint256 value = 1 ether;
        bytes memory data = abi.encodePacked("hello");

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, owner, TSS_ROLE));
        gateway.executeRevert{ value: value }(address(receiver), data, revertContext);
    }

    function testExecuteRevertFailsIfReceiverIsZeroAddress() public {
        uint256 value = 1 ether;
        bytes memory data = abi.encodePacked("hello");

        vm.prank(tssAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.executeRevert{ value: value }(address(0), data, revertContext);
    }

    function testUpgradeAndForwardCallToReceivePayable() public {
        // upgrade
        Upgrades.upgradeProxy(address(gateway), "GatewayEVMUpgradeTest.sol", "", owner);
        GatewayEVMUpgradeTest gatewayUpgradeTest = GatewayEVMUpgradeTest(address(gateway));
        // call
        address custodyBeforeUpgrade = gateway.custody();
        address tssBeforeUpgrade = gateway.tssAddress();

        string memory str = "Hello, Foundry!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        bytes memory data = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);
        vm.expectCall(address(receiver), value, data);
        vm.expectEmit(true, true, true, true, address(receiver));
        emit ReceivedPayable(address(gateway), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gateway));
        emit ExecutedV2(address(receiver), value, data);
        vm.prank(tssAddress);
        gatewayUpgradeTest.execute{ value: value }(arbitraryCallMessageContext, address(receiver), data);

        assertEq(custodyBeforeUpgrade, gateway.custody());
        assertEq(tssBeforeUpgrade, gateway.tssAddress());
    }
}

contract GatewayEVMInboundTest is
    Test,
    IGatewayEVMErrors,
    IGatewayEVMEvents,
    IReceiverEVMEvents,
    INotSupportedMethods
{
    using SafeERC20 for IERC20;

    address proxy;
    GatewayEVM gateway;
    ERC20Custody custody;
    ZetaConnectorNonNative zetaConnector;
    TestERC20 token;
    ZetaNonEth zeta;
    address owner;
    address destination;
    address tssAddress;
    RevertOptions revertOptions;

    uint256 ownerAmount = 1_000_000;

    function setUp() public {
        owner = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);

        token = new TestERC20("test", "TTK");

        zeta = new ZetaNonEth(tssAddress, tssAddress);
        proxy = Upgrades.deployUUPSProxy(
            "GatewayEVM.sol", abi.encodeCall(GatewayEVM.initialize, (tssAddress, address(zeta), owner))
        );
        gateway = GatewayEVM(proxy);

        proxy = Upgrades.deployUUPSProxy(
            "ERC20Custody.sol", abi.encodeCall(ERC20Custody.initialize, (address(gateway), tssAddress, owner))
        );
        custody = ERC20Custody(proxy);
        proxy = Upgrades.deployUUPSProxy(
            "ZetaConnectorNonNative.sol",
            abi.encodeCall(ZetaConnectorBase.initialize, (address(gateway), address(zeta), tssAddress, owner))
        );
        zetaConnector = ZetaConnectorNonNative(proxy);

        vm.prank(tssAddress);
        zeta.updateTssAndConnectorAddresses(tssAddress, address(zetaConnector));
        vm.deal(tssAddress, 1 ether);

        vm.startPrank(owner);
        gateway.setCustody(address(custody));
        gateway.setConnector(address(zetaConnector));

        custody.whitelist(address(token));
        vm.stopPrank();

        token.mint(owner, ownerAmount);
        vm.prank(tssAddress);
        zetaConnector.withdraw(owner, ownerAmount, "");

        revertOptions = RevertOptions({
            revertAddress: address(0x321),
            callOnRevert: false,
            abortAddress: address(0x321),
            revertMessage: "",
            onRevertGasLimit: 0
        });
    }

    function testDepositERC20ToCustody() public {
        uint256 amount = 100_000;
        uint256 custodyBalanceBefore = token.balanceOf(address(custody));
        assertEq(0, custodyBalanceBefore);

        token.approve(address(gateway), amount);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Deposited(owner, destination, amount, address(token), "", revertOptions);
        gateway.deposit(destination, amount, address(token), revertOptions);

        uint256 custodyBalanceAfter = token.balanceOf(address(custody));
        assertEq(amount, custodyBalanceAfter);

        uint256 ownerAmountAfter = token.balanceOf(owner);
        assertEq(ownerAmount - amount, ownerAmountAfter);
    }

    function testDepositERC20ToCustodyFailsIfTokenIsNotWhitelisted() public {
        uint256 amount = 100_000;
        token.approve(address(gateway), amount);

        vm.prank(owner);
        custody.unwhitelist(address(token));

        vm.expectRevert(NotWhitelistedInCustody.selector);
        gateway.deposit(destination, amount, address(token), revertOptions);
    }

    function testDepositERC20ToCustodyFailsIfPayloadSizeExceeded() public {
        uint256 amount = 100_000;
        token.approve(address(gateway), amount);
        revertOptions.revertMessage = new bytes(gateway.MAX_PAYLOAD_SIZE() + 1);

        vm.expectRevert(PayloadSizeExceeded.selector);
        gateway.deposit(destination, amount, address(token), revertOptions);
    }

    function testDepositZetaToConnector() public {
        uint256 amount = 100_000;
        zeta.approve(address(gateway), amount);

        // TODO: replace error to check once ZETA supported back
        // https://github.com/zeta-chain/protocol-contracts-evm/issues/394
        // vm.expectEmit(true, true, true, true, address(gateway));
        vm.expectRevert(ZETANotSupported.selector);

        emit Deposited(owner, destination, amount, address(zeta), "", revertOptions);
        gateway.deposit(destination, amount, address(zeta), revertOptions);

        // uint256 ownerAmountAfter = zeta.balanceOf(owner);
        // assertEq(ownerAmount - amount, ownerAmountAfter);
    }

    function testRevertDepositERC20ToCustodyIfAmountIs0() public {
        uint256 amount = 0;

        token.approve(address(gateway), amount);

        vm.expectRevert(InsufficientERC20Amount.selector);
        gateway.deposit(destination, amount, address(token), revertOptions);
    }

    function testRevertDepositERC20ToCustodyIfReceiverIsZeroAddress() public {
        uint256 amount = 1;
        vm.expectRevert(ZeroAddress.selector);
        gateway.deposit(address(0), amount, address(token), revertOptions);
    }

    function testDepositEthToTss() public {
        uint256 amount = 100_000;
        uint256 tssBalanceBefore = tssAddress.balance;

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Deposited(owner, destination, amount, address(0), "", revertOptions);
        gateway.deposit{ value: amount }(destination, revertOptions);

        uint256 tssBalanceAfter = tssAddress.balance;
        assertEq(tssBalanceBefore + amount, tssBalanceAfter);
    }

    function testRevertDepositEthToTssIfAmountIs0() public {
        uint256 amount = 0;

        vm.expectRevert(InsufficientETHAmount.selector);
        gateway.deposit{ value: amount }(destination, revertOptions);
    }

    function testRevertDepositEthToTssIfPayloadSizeExceeded() public {
        revertOptions.revertMessage = new bytes(gateway.MAX_PAYLOAD_SIZE() + 1);
        vm.expectRevert(PayloadSizeExceeded.selector);
        gateway.deposit{ value: 1 }(destination, revertOptions);
    }

    function testRevertDepositEthToTssIfReceiverIsZeroAddress() public {
        uint256 amount = 1;

        vm.expectRevert(ZeroAddress.selector);
        gateway.deposit{ value: amount }(address(0), revertOptions);
    }

    function testDepositERC20ToCustodyWithPayloadFailsIfTokenIsNotWhitelisted() public {
        uint256 amount = 100_000;
        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        token.approve(address(gateway), amount);

        vm.prank(owner);
        custody.unwhitelist(address(token));

        vm.expectRevert(NotWhitelistedInCustody.selector);
        gateway.depositAndCall(destination, amount, address(token), payload, revertOptions);
    }

    function testDepositERC20ToCustodyWithPayloadFailsIfPayloadSizeExceeded() public {
        uint256 amount = 100_000;
        bytes memory payload = new bytes(gateway.MAX_PAYLOAD_SIZE() / 2);
        revertOptions.revertMessage = new bytes(gateway.MAX_PAYLOAD_SIZE() / 2 + 1);

        token.approve(address(gateway), amount);

        vm.expectRevert(PayloadSizeExceeded.selector);
        gateway.depositAndCall(destination, amount, address(token), payload, revertOptions);
    }

    function testDepositERC20ToCustodyWithPayload() public {
        uint256 amount = 100_000;
        uint256 custodyBalanceBefore = token.balanceOf(address(custody));
        assertEq(0, custodyBalanceBefore);

        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        token.approve(address(gateway), amount);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit DepositedAndCalled(owner, destination, amount, address(token), payload, revertOptions);
        gateway.depositAndCall(destination, amount, address(token), payload, revertOptions);

        uint256 custodyBalanceAfter = token.balanceOf(address(custody));
        assertEq(amount, custodyBalanceAfter);

        uint256 ownerAmountAfter = token.balanceOf(owner);
        assertEq(ownerAmount - amount, ownerAmountAfter);
    }

    function testRevertDepositERC20ToCustodyWithPayloadIfAmountIs0() public {
        uint256 amount = 0;

        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectRevert(InsufficientERC20Amount.selector);
        gateway.depositAndCall(destination, amount, address(token), payload, revertOptions);
    }

    function testRevertDepositERC20ToCustodyWithPayloadIfReceiverIsZeroAddress() public {
        uint256 amount = 1;

        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(address(0), amount, address(token), payload, revertOptions);
    }

    function testDepositEthToTssWithPayload() public {
        uint256 amount = 100_000;
        uint256 tssBalanceBefore = tssAddress.balance;
        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit DepositedAndCalled(owner, destination, amount, address(0), payload, revertOptions);
        gateway.depositAndCall{ value: amount }(destination, payload, revertOptions);

        uint256 tssBalanceAfter = tssAddress.balance;
        assertEq(tssBalanceBefore + amount, tssBalanceAfter);
    }

    function testDepositEthToTssWithPayloadFailsIfPayloadSizeExceeded() public {
        uint256 amount = 100_000;
        bytes memory payload = new bytes(gateway.MAX_PAYLOAD_SIZE() / 2);
        revertOptions.revertMessage = new bytes(gateway.MAX_PAYLOAD_SIZE() / 2 + 1);

        vm.expectRevert(PayloadSizeExceeded.selector);
        gateway.depositAndCall{ value: amount }(destination, payload, revertOptions);
    }

    function testRevertDepositEthToTssWithPayloadIfAmountIs0() public {
        uint256 amount = 0;
        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectRevert(InsufficientETHAmount.selector);
        gateway.depositAndCall{ value: amount }(destination, payload, revertOptions);
    }

    function testRevertDepositEthToTssWithPayloadIfReceiverIsZeroAddress() public {
        uint256 amount = 1;
        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall{ value: amount }(address(0), payload, revertOptions);
    }

    function testCallWithPayload() public {
        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Called(owner, destination, payload, revertOptions);
        gateway.call(destination, payload, revertOptions);
    }

    function testCallWithPayloadFailsIfPayloadSizeExceeded() public {
        bytes memory payload = new bytes(gateway.MAX_PAYLOAD_SIZE() / 2);
        revertOptions.revertMessage = new bytes(gateway.MAX_PAYLOAD_SIZE() / 2 + 1);

        vm.expectRevert(PayloadSizeExceeded.selector);
        gateway.call(destination, payload, revertOptions);
    }

    function testCallWithPayloadFailsIfCallOnRevertIsTrue() public {
        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);
        revertOptions.callOnRevert = true;

        vm.expectRevert(CallOnRevertNotSupported.selector);
        gateway.call(destination, payload, revertOptions);
    }

    function testCallWithPayloadFailsIfDestinationIsZeroAddress() public {
        bytes memory payload = abi.encodeWithSignature("hello(address)", destination);

        vm.expectRevert(ZeroAddress.selector);
        gateway.call(address(0), payload, revertOptions);
    }
}
