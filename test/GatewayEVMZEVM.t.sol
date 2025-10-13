// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "./utils/ReceiverEVM.sol";

import "../contracts/evm/ERC20Custody.sol";
import { GatewayEVM } from "../contracts/evm/GatewayEVM.sol";
import "../contracts/evm/ZetaConnectorNonNative.sol";
import "./utils/TestERC20.sol";

import "./utils/SenderZEVM.sol";

import { SystemContractMock } from "./utils/SystemContractMock.sol";

import { GatewayZEVM } from "../contracts/zevm/GatewayZEVM.sol";
import { IGatewayZEVM } from "../contracts/zevm/GatewayZEVM.sol";
import { CallOptions, IGatewayZEVMErrors, IGatewayZEVMEvents } from "../contracts/zevm/interfaces/IGatewayZEVM.sol";

import { IGatewayEVMErrors } from "../contracts/evm/interfaces/IGatewayEVM.sol";
import { IGatewayEVMEvents } from "../contracts/evm/interfaces/IGatewayEVM.sol";

import "../contracts/zevm/ZRC20.sol";

import "./utils/IReceiverEVM.sol";

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract GatewayEVMZEVMTest is
    Test,
    IGatewayEVMErrors,
    IGatewayEVMEvents,
    IGatewayZEVMEvents,
    IGatewayZEVMErrors,
    IReceiverEVMEvents
{
    // evm
    using SafeERC20 for IERC20;

    GatewayEVM gatewayEVM;
    ERC20Custody custody;
    ZetaConnectorNonNative zetaConnector;
    TestERC20 token;
    TestERC20 zeta;
    ReceiverEVM receiverEVM;
    address ownerEVM;
    address destination;
    address tssAddress;
    MessageContext arbitraryCallMessageContext = MessageContext({ sender: address(0), asset: address(0), amount: 0 });

    // zevm
    address payable proxyZEVM;
    GatewayZEVM gatewayZEVM;
    SenderZEVM senderZEVM;
    SystemContractMock systemContract;
    ZRC20 zrc20;
    address ownerZEVM;

    CallOptions callOptions;
    RevertOptions revertOptions;

    function setUp() public {
        // evm
        ownerEVM = address(this);
        destination = address(0x1234);
        tssAddress = address(0x5678);
        ownerZEVM = address(0x4321);

        token = new TestERC20("test", "TTK");
        zeta = new TestERC20("zeta", "ZETA");

        address proxy = Upgrades.deployUUPSProxy(
            "GatewayEVM.sol", abi.encodeCall(GatewayEVM.initialize, (tssAddress, address(zeta), ownerEVM))
        );
        gatewayEVM = GatewayEVM(proxy);
        proxy = Upgrades.deployUUPSProxy(
            "ERC20Custody.sol", abi.encodeCall(ERC20Custody.initialize, (address(gatewayEVM), tssAddress, ownerEVM))
        );
        custody = ERC20Custody(proxy);
        proxy = Upgrades.deployUUPSProxy(
            "ZetaConnectorNonNative.sol",
            abi.encodeCall(
                ZetaConnectorNonNative.initialize, (address(gatewayEVM), address(zeta), tssAddress, ownerEVM)
            )
        );
        zetaConnector = ZetaConnectorNonNative(proxy);

        vm.deal(tssAddress, 1 ether);

        vm.startPrank(ownerEVM);
        gatewayEVM.setCustody(address(custody));
        gatewayEVM.setConnector(address(zetaConnector));
        vm.stopPrank();

        token.mint(ownerEVM, 1_000_000);
        token.transfer(address(custody), 500_000);

        receiverEVM = new ReceiverEVM();

        // zevm
        proxyZEVM = payable(
            Upgrades.deployUUPSProxy(
                "GatewayZEVM.sol", abi.encodeCall(GatewayZEVM.initialize, (address(zeta), ownerZEVM))
            )
        );
        gatewayZEVM = GatewayZEVM(proxyZEVM);

        senderZEVM = new SenderZEVM(address(gatewayZEVM));
        address protocolAddress = address(0x735b14BB79463307AAcBED86DAf3322B1e6226aB);
        vm.startPrank(protocolAddress);
        systemContract = new SystemContractMock(address(0), address(0), address(0));
        zrc20 = new ZRC20("TOKEN", "TKN", 18, 1, CoinType.Zeta, 0, address(systemContract), address(gatewayZEVM));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);
        zrc20.deposit(ownerZEVM, 1_000_000);
        zrc20.deposit(address(senderZEVM), 1_000_000);
        vm.stopPrank();

        vm.prank(ownerZEVM);
        zrc20.approve(address(gatewayZEVM), 1_000_000);

        vm.deal(tssAddress, 1 ether);

        callOptions = CallOptions({ gasLimit: 100_000, isArbitraryCall: false });

        revertOptions = RevertOptions({
            revertAddress: address(0x321),
            callOnRevert: true,
            abortAddress: address(0x321),
            revertMessage: "",
            onRevertGasLimit: 0
        });
    }

    function testCallReceiverEVMFromZEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on zevm
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        vm.prank(ownerZEVM);
        vm.expectEmit(true, true, true, true, address(gatewayZEVM));
        emit Called(
            address(ownerZEVM),
            address(zrc20),
            abi.encodePacked(receiverEVM),
            message,
            CallOptions({ gasLimit: 100_000, isArbitraryCall: true }),
            revertOptions
        );
        gatewayZEVM.call(
            abi.encodePacked(receiverEVM),
            address(zrc20),
            message,
            CallOptions({ gasLimit: 100_000, isArbitraryCall: true }),
            revertOptions
        );

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        vm.prank(tssAddress);
        gatewayEVM.execute{ value: value }(arbitraryCallMessageContext, address(receiverEVM), message);
    }

    function testCallReceiverEVMFromSenderZEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on zevm
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        bytes memory data = abi.encodeWithSignature(
            "call(bytes,address,bytes,(uint256,bool),(address,bool,address,bytes,uint256))",
            abi.encodePacked(receiverEVM),
            address(zrc20),
            message,
            callOptions,
            revertOptions
        );
        vm.expectCall(address(gatewayZEVM), 0, data);
        vm.prank(ownerZEVM);
        senderZEVM.callReceiver(abi.encodePacked(receiverEVM), address(zrc20), str, num, flag);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(receiverEVM));
        emit ReceivedPayable(address(gatewayEVM), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        vm.prank(tssAddress);
        gatewayEVM.execute{ value: value }(arbitraryCallMessageContext, address(receiverEVM), message);
    }

    function testWithdrawAndCallReceiverEVMFromZEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on zevm
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        uint256 expectedGasFee = 100_000;
        vm.expectEmit(true, true, true, true, address(gatewayZEVM));
        emit WithdrawnAndCalled(
            ownerZEVM,
            0,
            abi.encodePacked(receiverEVM),
            address(zrc20),
            500_000,
            expectedGasFee,
            zrc20.PROTOCOL_FLAT_FEE(),
            message,
            CallOptions({ gasLimit: 100_000, isArbitraryCall: true }),
            revertOptions
        );
        vm.prank(ownerZEVM);
        gatewayZEVM.withdrawAndCall(
            abi.encodePacked(receiverEVM),
            500_000,
            address(zrc20),
            message,
            CallOptions({ gasLimit: 100_000, isArbitraryCall: true }),
            revertOptions
        );

        // Check the balance after withdrawal
        uint256 balanceOfAfterWithdrawal = zrc20.balanceOf(ownerZEVM);
        assertEq(balanceOfAfterWithdrawal, 500_000 - expectedGasFee);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(receiverEVM));
        emit ReceivedPayable(address(gatewayEVM), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        vm.prank(tssAddress);
        gatewayEVM.execute{ value: value }(arbitraryCallMessageContext, address(receiverEVM), message);
    }

    function testWithdrawAndCallReceiverEVMFromSenderZEVM() public {
        string memory str = "Hello!";
        uint256 num = 42;
        bool flag = true;
        uint256 value = 1 ether;

        // Encode the function call data and call on zevm
        uint256 senderBalanceBeforeWithdrawal = IZRC20(zrc20).balanceOf(address(senderZEVM));
        bytes memory message = abi.encodeWithSelector(receiverEVM.receivePayable.selector, str, num, flag);
        bytes memory data = abi.encodeWithSignature(
            "withdrawAndCall(bytes,uint256,address,bytes,(uint256,bool),(address,bool,address,bytes,uint256))",
            abi.encodePacked(receiverEVM),
            500_000,
            address(zrc20),
            message,
            callOptions,
            revertOptions
        );
        vm.expectCall(address(gatewayZEVM), 0, data);
        vm.prank(ownerZEVM);
        senderZEVM.withdrawAndCallReceiver(abi.encodePacked(receiverEVM), 500_000, address(zrc20), str, num, flag);

        // Call execute on evm
        vm.deal(address(gatewayEVM), value);
        vm.expectEmit(true, true, true, true, address(receiverEVM));
        emit ReceivedPayable(address(gatewayEVM), value, str, num, flag);
        vm.expectEmit(true, true, true, true, address(gatewayEVM));
        emit Executed(address(receiverEVM), value, message);
        vm.prank(tssAddress);
        gatewayEVM.execute{ value: value }(arbitraryCallMessageContext, address(receiverEVM), message);

        // Check the balance after withdrawal
        uint256 senderBalanceAfterWithdrawal = IZRC20(zrc20).balanceOf(address(senderZEVM));
        // Expected gas fee 1
        assertEq(senderBalanceAfterWithdrawal, senderBalanceBeforeWithdrawal - 500_000 - 100_000);
    }
}
