// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

import "../contracts/zevm/SystemContract.sol";

import "./utils/TestUniversalContract.sol";

import "./utils/WZETA.sol";
import "./utils/upgrades/GatewayZEVMUpgradeTest.sol";

import { CoreRegistry } from "../contracts/zevm/CoreRegistry.sol";
import "../contracts/zevm/GatewayZEVM.sol";
import "../contracts/zevm/ZRC20.sol";
import "../contracts/zevm/interfaces/IGatewayZEVM.sol";
import "../contracts/zevm/interfaces/IZRC20.sol";
import { Upgrades } from "openzeppelin-foundry-upgrades/Upgrades.sol";

contract GatewayZEVMInboundTest is Test, IGatewayZEVMEvents, IGatewayZEVMErrors {
    address payable proxy;
    GatewayZEVM gateway;
    ZRC20 zrc20;
    ZRC20 solanaZRC20;
    WETH9 zetaToken;
    SystemContract systemContract;
    TestUniversalContract testUniversalContract;
    CoreRegistry coreRegistry;
    address owner;
    address addr1;
    address protocolAddress;
    RevertOptions revertOptions;
    CallOptions callOptions;

    error ZeroAddress();
    error LowBalance();
    error ZETANotSupported();

    event WithdrawnV2(
        address indexed sender,
        uint256 indexed chainId,
        bytes receiver,
        address zrc20,
        uint256 value,
        uint256 gasfee,
        uint256 protocolFlatFee,
        bytes message,
        CallOptions callOptions,
        RevertOptions revertOptions
    );

    uint256 constant MIN_GAS_LIMIT = 100_000;

    function setUp() public {
        owner = address(this);
        addr1 = address(0x1234);

        zetaToken = new WETH9();

        proxy = payable(
            Upgrades.deployUUPSProxy(
                "GatewayZEVM.sol", abi.encodeCall(GatewayZEVM.initialize, (address(zetaToken), owner))
            )
        );
        gateway = GatewayZEVM(proxy);

        address expectedRegistryAddress = 0x7CCE3Eb018bf23e1FE2a32692f2C77592D110394;
        bytes memory creationCode = abi.encodePacked(type(CoreRegistry).creationCode);
        address deployedRegistry;
        assembly {
            deployedRegistry := create2(0, add(creationCode, 0x20), mload(creationCode), 0)
        }
        vm.etch(expectedRegistryAddress, deployedRegistry.code);
        CoreRegistry(expectedRegistryAddress).initialize(owner, owner, address(gateway));
        coreRegistry = CoreRegistry(expectedRegistryAddress);

        protocolAddress = gateway.PROTOCOL_ADDRESS();
        testUniversalContract = new TestUniversalContract();

        vm.startPrank(protocolAddress);
        systemContract = new SystemContract(address(0), address(0), address(0));

        // test token
        zrc20 = new ZRC20("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);

        // test Solana token
        solanaZRC20 = new ZRC20("SOLANA", "SOL", 9, 900, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinZRC20(900, address(solanaZRC20));
        systemContract.setGasPrice(900, 1);

        vm.deal(protocolAddress, 1_000_000_000);
        zetaToken.deposit{ value: 10 }();
        zetaToken.approve(address(gateway), 10);
        zrc20.deposit(owner, 100_000_000);
        solanaZRC20.deposit(owner, 100_000_000);
        vm.stopPrank();

        vm.startPrank(owner);
        zrc20.approve(address(gateway), 100_000_000);
        solanaZRC20.approve(address(gateway), 100_000_000);
        zetaToken.deposit{ value: 10 }();
        zetaToken.approve(address(gateway), 10);
        vm.stopPrank();

        revertOptions = RevertOptions({
            revertAddress: address(0x321),
            callOnRevert: true,
            abortAddress: address(0x321),
            revertMessage: "",
            onRevertGasLimit: 0
        });

        callOptions = CallOptions({ gasLimit: MIN_GAS_LIMIT, isArbitraryCall: true });
    }

    function testWithdrawZRC20() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawn(
            owner,
            0,
            abi.encodePacked(addr1),
            address(zrc20),
            amount,
            0,
            zrc20.PROTOCOL_FLAT_FEE(),
            "",
            CallOptions({ gasLimit: 0, isArbitraryCall: true }),
            revertOptions
        );
        gateway.withdraw(abi.encodePacked(addr1), amount, address(zrc20), revertOptions);

        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount, ownerBalanceAfter);
    }

    function testWithdrawZRC20Solana() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawn(
            owner,
            0,
            abi.encodePacked(addr1),
            address(solanaZRC20),
            amount,
            0,
            solanaZRC20.PROTOCOL_FLAT_FEE(),
            "",
            CallOptions({ gasLimit: 0, isArbitraryCall: true }),
            revertOptions
        );
        gateway.withdraw(abi.encodePacked(addr1), amount, address(solanaZRC20), revertOptions);

        uint256 ownerBalanceAfter = solanaZRC20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount, ownerBalanceAfter);
    }

    function testWithdrawZRC20FailsIfNoBalanceForGasFee() public {
        uint256 amount = 1;
        uint256 ownerBalance = zrc20.balanceOf(owner);
        zrc20.transfer(address(0x123), ownerBalance);

        vm.prank(protocolAddress);
        zrc20.updateGasLimit(10);

        vm.expectRevert(LowBalance.selector);
        gateway.withdraw(abi.encodePacked(addr1), amount, address(zrc20), revertOptions);
    }

    function testWithdrawZRC20FailsIfNoBalanceForTransfer() public {
        uint256 amount = 2;
        uint256 ownerBalance = zrc20.balanceOf(owner);
        zrc20.transfer(address(0x123), ownerBalance - 1);

        vm.expectRevert(LowBalance.selector);
        gateway.withdraw(abi.encodePacked(addr1), amount, address(zrc20), revertOptions);
    }

    function testWithdrawZRC20FailsIfMessageSizeExceeded() public {
        revertOptions.revertMessage = new bytes(gateway.MAX_MESSAGE_SIZE() + 1);
        vm.expectRevert(MessageSizeExceeded.selector);
        gateway.withdraw(abi.encodePacked(addr1), 2, address(zrc20), revertOptions);
    }

    function testWithdrawZRC20FailsIsAmountIs0() public {
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.withdraw(abi.encodePacked(addr1), 0, address(zrc20), revertOptions);
    }

    function testWithdrawZRC20FailsIfReceiverIsZeroAddress() public {
        vm.expectRevert(ZeroAddress.selector);
        gateway.withdraw(abi.encodePacked(""), 1, address(zrc20), revertOptions);
    }

    function testWithdrawZRC20FailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        // Remove allowance for gateway
        vm.prank(owner);
        zrc20.approve(address(gateway), 0);

        vm.expectRevert();
        gateway.withdraw(abi.encodePacked(addr1), amount, address(zrc20), revertOptions);

        // Check that balance didn't change
        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);
    }

    function testWithdrawZRC20WithCustomGasLimit() public {
        uint256 amount = 1;
        uint256 customGasLimit = 200_000;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);
        uint256 expectedGasFee = customGasLimit;

        vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawn(
            owner,
            0,
            abi.encodePacked(addr1),
            address(zrc20),
            amount,
            expectedGasFee,
            zrc20.PROTOCOL_FLAT_FEE(),
            "",
            CallOptions({ gasLimit: customGasLimit, isArbitraryCall: true }),
            revertOptions
        );

        gateway.withdraw(abi.encodePacked(addr1), amount, address(zrc20), customGasLimit, revertOptions);
        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount - expectedGasFee, ownerBalanceAfter);
    }

    function testWithdrawZRC20WithCustomGasLimitFailsIfGasLimitTooLow() public {
        uint256 amount = 1;
        uint256 lowGasLimit = MIN_GAS_LIMIT - 1;
        vm.expectRevert(InsufficientGasLimit.selector);
        gateway.withdraw(abi.encodePacked(addr1), amount, address(zrc20), lowGasLimit, revertOptions);
    }

    function testWithdrawZRC20WithCustomGasLimitFailsIfAmountIsZero() public {
        uint256 customGasLimit = 150_000;

        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.withdraw(abi.encodePacked(addr1), 0, address(zrc20), customGasLimit, revertOptions);
    }

    function testWithdrawZRC20WithCustomGasLimitFailsIfReceiverIsEmpty() public {
        uint256 amount = 1;
        uint256 customGasLimit = 150_000;

        vm.expectRevert(ZeroAddress.selector);
        gateway.withdraw(abi.encodePacked(""), amount, address(zrc20), customGasLimit, revertOptions);
    }

    function testWithdrawZRC20WithCustomGasLimitFailsIfMessageSizeExceeded() public {
        uint256 amount = 1;
        uint256 customGasLimit = 150_000;
        revertOptions.revertMessage = new bytes(gateway.MAX_MESSAGE_SIZE() + 1);

        vm.expectRevert(MessageSizeExceeded.selector);
        gateway.withdraw(abi.encodePacked(addr1), amount, address(zrc20), customGasLimit, revertOptions);
    }

    function testWithdrawAndCallZRC20FailsIfReceiverIsZeroAddress() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(ZeroAddress.selector);
        gateway.withdrawAndCall(
            abi.encodePacked(""),
            1,
            address(zrc20),
            message,
            CallOptions({ gasLimit: MIN_GAS_LIMIT, isArbitraryCall: false }),
            revertOptions
        );
    }

    function testWithdrawAndCallZRC20FailsIfMessageSizeExceeded() public {
        bytes memory message = new bytes(gateway.MAX_MESSAGE_SIZE() / 2);
        revertOptions.revertMessage = new bytes(gateway.MAX_MESSAGE_SIZE() / 2 + 1);

        vm.expectRevert(MessageSizeExceeded.selector);
        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            1,
            address(zrc20),
            message,
            CallOptions({ gasLimit: 100_000, isArbitraryCall: false }),
            revertOptions
        );
    }

    function testWithdrawAndCallZRC20FailsIfGasLimitIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(InsufficientGasLimit.selector);
        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            1,
            address(zrc20),
            message,
            CallOptions({ gasLimit: 0, isArbitraryCall: false }),
            revertOptions
        );
    }

    function testWithdrawAndCallZRC20FailsIfGasLimitIsBelowMin() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(InsufficientGasLimit.selector);
        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            1,
            address(zrc20),
            message,
            CallOptions({ gasLimit: MIN_GAS_LIMIT - 1, isArbitraryCall: false }),
            revertOptions
        );
    }

    function testWithdrawAndCallZRC20FailsIfAmountIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            0,
            address(zrc20),
            message,
            CallOptions({ gasLimit: MIN_GAS_LIMIT, isArbitraryCall: false }),
            revertOptions
        );
    }

    function testWithdrawAndCallZRC20WithCallOptsV2FailsIfReceiverIsZeroAddress() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        CallOptionsV2 memory callOptionsV2 =
            CallOptionsV2({ gasLimit: MIN_GAS_LIMIT, isArbitraryCall: true, isMsgContextV2: true });

        vm.expectRevert(ZeroAddress.selector);
        gateway.withdrawAndCall(abi.encodePacked(""), 1, address(zrc20), message, callOptionsV2, revertOptions);
    }

    function testWithdrawAndCallZRC20WithCallOptsV2FailsIfMessageSizeExceeded() public {
        uint256 maxSize = gateway.MAX_MESSAGE_SIZE();
        bytes memory message = new bytes(maxSize / 2);
        revertOptions.revertMessage = new bytes(maxSize / 2 + 1);
        CallOptionsV2 memory callOptionsV2 =
            CallOptionsV2({ gasLimit: MIN_GAS_LIMIT, isArbitraryCall: true, isMsgContextV2: true });

        uint256 messageSize = message.length + revertOptions.revertMessage.length;

        vm.expectRevert(abi.encodeWithSelector(MessageSizeExceeded.selector, messageSize, maxSize));
        gateway.withdrawAndCall(abi.encodePacked(addr1), 1, address(zrc20), message, callOptionsV2, revertOptions);
    }

    function testWithdrawAndCallZRC20WithCallOptsV2FailsIfGasLimitIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        CallOptionsV2 memory callOptionsV2 = CallOptionsV2({ gasLimit: 0, isArbitraryCall: true, isMsgContextV2: true });

        vm.expectRevert(InsufficientGasLimit.selector);
        gateway.withdrawAndCall(abi.encodePacked(addr1), 1, address(zrc20), message, callOptionsV2, revertOptions);
    }

    function testWithdrawAndCallZRC20WithCallOptsV2FailsIfGasLimitIsBelowMin() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        CallOptionsV2 memory callOptionsV2 =
            CallOptionsV2({ gasLimit: MIN_GAS_LIMIT - 1, isArbitraryCall: true, isMsgContextV2: true });

        vm.expectRevert(InsufficientGasLimit.selector);
        gateway.withdrawAndCall(abi.encodePacked(addr1), 1, address(zrc20), message, callOptionsV2, revertOptions);
    }

    function testWithdrawAndCallZRC20WithCallOptsV2FailsIfAmountIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        CallOptionsV2 memory callOptionsV2 =
            CallOptionsV2({ gasLimit: MIN_GAS_LIMIT, isArbitraryCall: true, isMsgContextV2: true });

        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.withdrawAndCall(abi.encodePacked(addr1), 0, address(zrc20), message, callOptionsV2, revertOptions);
    }

    function testWithdrawAndCallZRC20WithCallOptsV2WithMessageContextV2() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 expectedGasFee = MIN_GAS_LIMIT;
        CallOptionsV2 memory callOptionsV2 =
            CallOptionsV2({ gasLimit: MIN_GAS_LIMIT, isArbitraryCall: true, isMsgContextV2: true });

        vm.expectEmit(true, true, true, true, address(gateway));
        emit WithdrawnAndCalledV2(
            owner,
            0,
            abi.encodePacked(addr1),
            address(zrc20),
            amount,
            expectedGasFee,
            zrc20.PROTOCOL_FLAT_FEE(),
            message,
            callOptionsV2,
            revertOptions
        );

        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, address(zrc20), message, callOptionsV2, revertOptions);

        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount - expectedGasFee, ownerBalanceAfter);
    }

    function testWithdrawZRC20WithMessageFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        // Remove allowance for gateway
        vm.prank(owner);
        zrc20.approve(address(gateway), 0);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert();
        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            amount,
            address(zrc20),
            message,
            CallOptions({ gasLimit: MIN_GAS_LIMIT, isArbitraryCall: false }),
            revertOptions
        );

        // Check that balance didn't change
        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);
    }

    function testWithdrawAndCallZRC20SolanaFails() public {
        uint256 amount = 1;
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 gasLimit = 1;

        vm.expectRevert();
        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            amount,
            address(solanaZRC20),
            message,
            CallOptions({ gasLimit: gasLimit, isArbitraryCall: true }),
            revertOptions
        );
    }

    function testWithdrawZRC20WithMessage() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 expectedGasFee = MIN_GAS_LIMIT;
        uint256 gasLimit = MIN_GAS_LIMIT;
        vm.expectEmit(true, true, true, true, address(gateway));
        emit WithdrawnAndCalled(
            owner,
            0,
            abi.encodePacked(addr1),
            address(zrc20),
            amount,
            expectedGasFee,
            zrc20.PROTOCOL_FLAT_FEE(),
            message,
            CallOptions({ gasLimit: gasLimit, isArbitraryCall: true }),
            revertOptions
        );
        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            amount,
            address(zrc20),
            message,
            CallOptions({ gasLimit: gasLimit, isArbitraryCall: true }),
            revertOptions
        );

        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount - expectedGasFee, ownerBalanceAfter);
    }

    function testWithdrawAndCallZRC20WithCallOptsFailsIfReceiverIsZeroAddress() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(ZeroAddress.selector);
        gateway.withdrawAndCall(abi.encodePacked(""), 1, address(zrc20), message, callOptions, revertOptions);
    }

    function testWithdrawAndCallZRC20WithCallOptsFailsIfMessageSizeExceeded() public {
        bytes memory message = new bytes(gateway.MAX_MESSAGE_SIZE() / 2);
        revertOptions.revertMessage = new bytes(gateway.MAX_MESSAGE_SIZE() / 2 + 1);
        vm.expectRevert(MessageSizeExceeded.selector);
        gateway.withdrawAndCall(abi.encodePacked(addr1), 1, address(zrc20), message, callOptions, revertOptions);
    }

    function testWithdrawAndCallZRC20WithCallOptsFailsIfGasLimitIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        callOptions.gasLimit = 0;
        vm.expectRevert(InsufficientGasLimit.selector);
        gateway.withdrawAndCall(abi.encodePacked(addr1), 1, address(zrc20), message, callOptions, revertOptions);
    }

    function testWithdrawAndCallZRC20WithCallOptsFailsIfAmountIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.withdrawAndCall(abi.encodePacked(addr1), 0, address(zrc20), message, callOptions, revertOptions);
    }

    function testWithdrawZRC20WithMessageWithCallOptsFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        // Remove allowance for gateway
        vm.prank(owner);
        zrc20.approve(address(gateway), 0);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert();
        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, address(zrc20), message, callOptions, revertOptions);

        // Check that balance didn't change
        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);
    }

    function testWithdrawZRC20WithCallOptsWithMessage() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 expectedGasFee = MIN_GAS_LIMIT;
        uint256 gasLimit = MIN_GAS_LIMIT;
        vm.expectEmit(true, true, true, true, address(gateway));
        emit WithdrawnAndCalled(
            owner,
            0,
            abi.encodePacked(addr1),
            address(zrc20),
            amount,
            expectedGasFee,
            zrc20.PROTOCOL_FLAT_FEE(),
            message,
            CallOptions({ gasLimit: gasLimit, isArbitraryCall: true }),
            revertOptions
        );

        gateway.withdrawAndCall(
            abi.encodePacked(addr1),
            amount,
            address(zrc20),
            message,
            CallOptions({ gasLimit: gasLimit, isArbitraryCall: true }),
            revertOptions
        );

        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount - expectedGasFee, ownerBalanceAfter);
    }

    function testWithdrawZETAFailsIfAmountIsZero() public {
        // TODO: replace error to check once ZETA supported back
        // https://github.com/zeta-chain/protocol-contracts/issues/394
        // vm.expectRevert(InsufficientZetaAmount.selector);
        vm.expectRevert(ZETANotSupported.selector);

        gateway.withdraw(abi.encodePacked(addr1), 0, 1, revertOptions);
    }

    function testWithdrawZETAFailsIfMessageSizeExceeded() public {
        revertOptions.revertMessage = new bytes(gateway.MAX_MESSAGE_SIZE() + 1);

        // TODO: replace error to check once ZETA supported back
        // https://github.com/zeta-chain/protocol-contracts/issues/394
        // vm.expectRevert(MessageSizeExceeded.selector);
        vm.expectRevert(ZETANotSupported.selector);

        gateway.withdraw(abi.encodePacked(addr1), 1, 1, revertOptions);
    }

    function testWithdrawZETAFailsIfReceiverIsZeroAddress() public {
        // TODO: replace error to check once ZETA supported back
        // https://github.com/zeta-chain/protocol-contracts/issues/394
        // vm.expectRevert(ZeroAddress.selector);
        vm.expectRevert(ZETANotSupported.selector);

        gateway.withdraw(abi.encodePacked(""), 0, 1, revertOptions);
    }

    function testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);

        // TODO: replace error to check once ZETA supported back
        // https://github.com/zeta-chain/protocol-contracts/issues/394
        // vm.expectRevert(InsufficientZetaAmount.selector);
        vm.expectRevert(ZETANotSupported.selector);

        gateway.withdrawAndCall(abi.encodePacked(addr1), 0, 1, message, callOptions, revertOptions);
    }

    function testWithdrawAndCallZETAWithCallOptsFailsIfMessageSizeExceeded() public {
        bytes memory message = new bytes(gateway.MAX_MESSAGE_SIZE() / 2);
        revertOptions.revertMessage = new bytes(gateway.MAX_MESSAGE_SIZE() / 2 + 1);

        // TODO: replace error to check once ZETA supported back
        // https://github.com/zeta-chain/protocol-contracts/issues/394
        // vm.expectRevert(MessageSizeExceeded.selector);
        vm.expectRevert(ZETANotSupported.selector);

        gateway.withdrawAndCall(abi.encodePacked(addr1), 1, 1, message, callOptions, revertOptions);
    }

    function testWithdrawAndCallZETAWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);

        // TODO: replace error to check once ZETA supported back
        // https://github.com/zeta-chain/protocol-contracts/issues/394
        // vm.expectRevert(ZeroAddress.selector);
        vm.expectRevert(ZETANotSupported.selector);

        gateway.withdrawAndCall(abi.encodePacked(""), 1, 1, message, callOptions, revertOptions);
    }

    function testWithdrawZETA() public {
        uint256 amount = 1;
        // uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        // uint256 protocolBalanceBefore = protocolAddress.balance;
        uint256 chainId = 1;

        // vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawn(
            owner,
            chainId,
            abi.encodePacked(addr1),
            address(zetaToken),
            amount,
            0,
            0,
            "",
            CallOptions({ gasLimit: 0, isArbitraryCall: true }),
            revertOptions
        );

        // TODO: remove error once ZETA supported back
        // https://github.com/zeta-chain/protocol-contracts/issues/394
        vm.expectRevert(ZETANotSupported.selector);

        gateway.withdraw(abi.encodePacked(addr1), amount, chainId, revertOptions);

        // uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        // assertEq(ownerBalanceBefore - 1, ownerBalanceAfter);

        // uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        // assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        // // Verify amount is transfered to protocol address
        // assertEq(protocolBalanceBefore + 1, protocolAddress.balance);
    }

    function testWithdrawZETAFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 protocolBalanceBefore = protocolAddress.balance;
        uint256 chainId = 1;

        // Remove allowance for gateway
        vm.prank(owner);
        zetaToken.approve(address(gateway), 0);

        vm.expectRevert();
        gateway.withdraw(abi.encodePacked(addr1), amount, chainId, revertOptions);

        // Verify balances not changed
        uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        assertEq(protocolBalanceBefore, protocolAddress.balance);
    }

    function testWithdrawZETAFailsIfNoBalance() public {
        uint256 amount = 1;
        uint256 ownerBalance = zetaToken.balanceOf(owner);
        zetaToken.transfer(address(0x123), ownerBalance);
        uint256 chainId = 1;

        vm.expectRevert();
        gateway.withdraw(abi.encodePacked(addr1), amount, chainId, revertOptions);
    }

    function testWithdrawZETAWithCallOptsWithMessage() public {
        uint256 amount = 1;
        // uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        // uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        // uint256 protocolAddressBalanceBefore = protocolAddress.balance;
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 chainId = 1;

        // vm.expectEmit(true, true, true, true, address(gateway));
        emit Withdrawn(
            owner,
            chainId,
            abi.encodePacked(addr1),
            address(zetaToken),
            amount,
            0,
            0,
            message,
            callOptions,
            revertOptions
        );

        // TODO: remove error once ZETA supported back
        // vm.expectRevert(InsufficientGasLimit.selector);
        vm.expectRevert(ZETANotSupported.selector);

        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, chainId, message, callOptions, revertOptions);

        // uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        // assertEq(ownerBalanceBefore - 1, ownerBalanceAfter);

        // uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        // assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        // // Verify amount is transfered to fungible module
        // assertEq(protocolAddressBalanceBefore + 1, protocolAddress.balance);
    }

    function testWithdrawZETAWithCallOptsWithMessageFailsIfNoAllowance() public {
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zetaToken.balanceOf(owner);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 protocolAddressBalanceBefore = protocolAddress.balance;
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 chainId = 1;

        // Remove allowance for gateway
        vm.prank(owner);
        zetaToken.approve(address(gateway), 0);

        vm.expectRevert();
        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, chainId, message, callOptions, revertOptions);

        // Verify balances not changed
        uint256 ownerBalanceAfter = zetaToken.balanceOf(owner);
        assertEq(ownerBalanceBefore, ownerBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        assertEq(protocolAddressBalanceBefore, protocolAddress.balance);
    }

    function testWithdrawZETAWithCallOptsWithMessageFailsIfGasLimitIsZero() public {
        uint256 amount = 1;
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        uint256 chainId = 1;

        // Remove allowance for gateway
        vm.prank(owner);
        zetaToken.approve(address(gateway), 0);

        callOptions.gasLimit = 0;

        // TODO: replace error to check once ZETA supported back
        // https://github.com/zeta-chain/protocol-contracts/issues/394
        // vm.expectRevert(InsufficientGasLimit.selector);
        vm.expectRevert(ZETANotSupported.selector);

        gateway.withdrawAndCall(abi.encodePacked(addr1), amount, chainId, message, callOptions, revertOptions);
    }

    function testCallWithCallOptsFailsIfReceiverIsZeroAddress() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectRevert(ZeroAddress.selector);
        gateway.call(abi.encodePacked(""), address(zrc20), message, callOptions, revertOptions);
    }

    function testCallWithCallOptsFailsIfMessageSizeExceeded() public {
        bytes memory message = new bytes(gateway.MAX_MESSAGE_SIZE() / 2);
        revertOptions.revertMessage = new bytes(gateway.MAX_MESSAGE_SIZE() / 2 + 1);
        vm.expectRevert(MessageSizeExceeded.selector);
        gateway.call(abi.encodePacked(addr1), address(zrc20), message, callOptions, revertOptions);
    }

    function testCallWithCallOptsFailsIfGasLimitIsZero() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        callOptions.gasLimit = 0;
        vm.expectRevert(InsufficientGasLimit.selector);
        gateway.call(abi.encodePacked(addr1), address(zrc20), message, callOptions, revertOptions);
    }

    function testCallWithCallOptsFailsIfGasLimitIsBelowMin() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        callOptions.gasLimit = MIN_GAS_LIMIT - 1;
        vm.expectRevert(InsufficientGasLimit.selector);
        gateway.call(abi.encodePacked(addr1), address(zrc20), message, callOptions, revertOptions);
    }

    function testCallWithCallOpts() public {
        bytes memory message = abi.encodeWithSignature("hello(address)", addr1);
        vm.expectEmit(true, true, true, true, address(gateway));

        emit Called(owner, address(zrc20), abi.encodePacked(addr1), message, callOptions, revertOptions);
        gateway.call(abi.encodePacked(addr1), address(zrc20), message, callOptions, revertOptions);
    }

    function testUpgradeAndWithdrawZRC20() public {
        // upgrade
        Upgrades.upgradeProxy(proxy, "GatewayZEVMUpgradeTest.sol", "", owner);
        GatewayZEVMUpgradeTest gatewayUpgradeTest = GatewayZEVMUpgradeTest(proxy);
        // withdraw
        uint256 amount = 1;
        uint256 ownerBalanceBefore = zrc20.balanceOf(owner);

        vm.expectEmit(true, true, true, true, address(gatewayUpgradeTest));
        emit WithdrawnV2(
            owner,
            0,
            abi.encodePacked(addr1),
            address(zrc20),
            amount,
            0,
            zrc20.PROTOCOL_FLAT_FEE(),
            "",
            CallOptions({ gasLimit: 0, isArbitraryCall: true }),
            revertOptions
        );
        gatewayUpgradeTest.withdraw(abi.encodePacked(addr1), amount, address(zrc20), revertOptions);

        uint256 ownerBalanceAfter = zrc20.balanceOf(owner);
        assertEq(ownerBalanceBefore - amount, ownerBalanceAfter);
    }
}

contract GatewayZEVMOutboundTest is Test, IGatewayZEVMEvents, IGatewayZEVMErrors {
    address payable proxy;
    GatewayZEVM gateway;
    ZRC20 zrc20;
    WETH9 zetaToken;
    SystemContract systemContract;
    TestUniversalContract testUniversalContract;
    CoreRegistry coreRegistry;
    address owner;
    address addr1;
    address protocolAddress;
    RevertOptions revertOptions;
    RevertContext revertContext;
    AbortContext abortContext;

    event ContextData(bytes origin, address sender, uint256 chainID, address msgSender, string message);
    event ContextDataRevert(RevertContext revertContext);
    event ContextDataAbort(AbortContext abortContext);

    error ZeroAddress();
    error EnforcedPause();
    error AccessControlUnauthorizedAccount(address account, bytes32 neededRole);

    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    uint256 constant MIN_GAS_LIMIT = 100_000;

    function setUp() public {
        owner = address(this);
        addr1 = address(0x1234);

        zetaToken = new WETH9();

        proxy = payable(
            Upgrades.deployUUPSProxy(
                "GatewayZEVM.sol", abi.encodeCall(GatewayZEVM.initialize, (address(zetaToken), owner))
            )
        );
        gateway = GatewayZEVM(proxy);

        address expectedRegistryAddress = 0x7CCE3Eb018bf23e1FE2a32692f2C77592D110394;
        bytes memory creationCode = abi.encodePacked(type(CoreRegistry).creationCode);
        address deployedRegistry;
        assembly {
            deployedRegistry := create2(0, add(creationCode, 0x20), mload(creationCode), 0)
        }
        vm.etch(expectedRegistryAddress, deployedRegistry.code);
        CoreRegistry(expectedRegistryAddress).initialize(owner, owner, address(gateway));
        coreRegistry = CoreRegistry(expectedRegistryAddress);

        protocolAddress = gateway.PROTOCOL_ADDRESS();
        testUniversalContract = new TestUniversalContract();

        vm.startPrank(protocolAddress);
        systemContract = new SystemContract(address(0), address(0), address(0));
        zrc20 = new ZRC20("TOKEN", "TKN", 18, 1, CoinType.Gas, 0, address(systemContract), address(gateway));
        systemContract.setGasCoinZRC20(1, address(zrc20));
        systemContract.setGasPrice(1, 1);
        vm.deal(protocolAddress, 1_000_000_000);
        zetaToken.deposit{ value: 10 }();
        zetaToken.approve(address(gateway), 10);
        zrc20.deposit(owner, MIN_GAS_LIMIT);
        vm.stopPrank();

        vm.startPrank(owner);
        zrc20.approve(address(gateway), MIN_GAS_LIMIT);
        zetaToken.deposit{ value: 10 }();
        zetaToken.approve(address(gateway), 10);
        vm.stopPrank();

        revertContext = RevertContext({ sender: owner, asset: address(0), amount: 1, revertMessage: "" });
        abortContext = AbortContext({
            sender: abi.encodePacked(owner),
            asset: address(0),
            amount: 1,
            outgoing: false,
            chainID: 1,
            revertMessage: ""
        });
    }

    function testDepositFailsIfZRC20IsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.deposit(address(0), 1, addr1);
    }

    function testDepositFailsIfTargetIsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.deposit(address(zrc20), 1, address(0));
    }

    function testDepositFailsIfAmountIs0() public {
        vm.prank(protocolAddress);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.deposit(address(zrc20), 0, addr1);
    }

    function testDeposit() public {
        uint256 amount = 1;
        uint256 balanceBefore = zrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(protocolAddress);
        gateway.deposit(address(zrc20), amount, addr1);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(amount, balanceAfter);
    }

    function testDepositTogglePause() public {
        vm.prank(protocolAddress);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, protocolAddress, PAUSER_ROLE));
        gateway.pause();

        vm.prank(protocolAddress);
        vm.expectRevert(abi.encodeWithSelector(AccessControlUnauthorizedAccount.selector, protocolAddress, PAUSER_ROLE));
        gateway.unpause();

        vm.prank(owner);
        gateway.pause();

        uint256 amount = 1;

        vm.expectRevert(EnforcedPause.selector);
        vm.prank(protocolAddress);
        gateway.deposit(address(zrc20), amount, addr1);

        vm.prank(owner);
        gateway.unpause();

        uint256 balanceBefore = zrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(protocolAddress);
        gateway.deposit(address(zrc20), amount, addr1);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(amount, balanceAfter);
    }

    function testDepositFailsIfSenderNotProtocol() public {
        uint256 amount = 1;
        uint256 balanceBefore = zrc20.balanceOf(addr1);
        assertEq(0, balanceBefore);

        vm.prank(owner);
        vm.expectRevert(CallerIsNotProtocol.selector);
        gateway.deposit(address(zrc20), amount, addr1);

        uint256 balanceAfter = zrc20.balanceOf(addr1);
        assertEq(0, balanceAfter);
    }

    function testDepositFailsIfTargetIsGateway() public {
        uint256 amount = 1;

        vm.prank(protocolAddress);
        vm.expectRevert(InvalidTarget.selector);
        gateway.deposit(address(zrc20), amount, address(gateway));
    }

    function testDepositFailsIfTargetIsProtocol() public {
        uint256 amount = 1;
        vm.prank(protocolAddress);
        vm.expectRevert(InvalidTarget.selector);
        gateway.deposit(address(zrc20), amount, protocolAddress);
    }

    function testExecuteFailsIfZRC20IsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.execute(context, address(0), 1, address(testUniversalContract), message);
    }

    function testExecuteFailsIfTargetIsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.execute(context, address(zrc20), 1, address(0), message);
    }

    function testExecuteUniversalContractFailsIfZeroAddress() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.execute(context, address(0), 1, address(testUniversalContract), message);
    }

    function testExecuteUniversalContract() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextData(abi.encodePacked(gateway), protocolAddress, 1, address(gateway), "hello");
        vm.prank(protocolAddress);
        gateway.execute(context, address(zrc20), 1, address(testUniversalContract), message);
    }

    function testExecuteUniversalContractFailsIfSenderIsNotProtocol() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.execute(context, address(zrc20), 1, address(testUniversalContract), message);
    }

    function testExecuteRevertUniversalContractFailsIfTargetIsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.executeRevert(address(0), revertContext);
    }

    function testExecuteRevertUniversalContract() public {
        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextDataRevert(revertContext);
        vm.prank(protocolAddress);
        gateway.executeRevert(address(testUniversalContract), revertContext);
    }

    function testExecuteRevertUniversalContractIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.executeRevert(address(testUniversalContract), revertContext);
    }

    function testDepositZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, address(0), 1, address(testUniversalContract), message);
    }

    function testDepositZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, address(zrc20), 1, address(0), message);
    }

    function testDepositZRC20AndCallUniversalContractFailsIfAmountIsZero() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.depositAndCall(context, address(zrc20), 0, address(testUniversalContract), message);
    }

    function testDepositZRC20AndCallUniversalContract() public {
        uint256 balanceBefore = zrc20.balanceOf(address(testUniversalContract));
        assertEq(0, balanceBefore);

        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextData(abi.encodePacked(gateway), protocolAddress, 1, address(gateway), "hello");
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, address(zrc20), 1, address(testUniversalContract), message);

        uint256 balanceAfter = zrc20.balanceOf(address(testUniversalContract));
        assertEq(1, balanceAfter);
    }

    function testDepositZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.depositAndCall(context, address(zrc20), 1, address(testUniversalContract), message);
    }

    function testDepositZRC20AndCallUniversalContractIfTargetIsProtocol() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, address(zrc20), 1, protocolAddress, message);
    }

    function testDepositZRC20AndCallUniversalContractIfTargetIsGateway() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, address(zrc20), 1, address(gateway), message);
    }

    function testDepositAndRevertZRC20AndCallUniversalContractFailsIfZRC20IsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndRevert(address(0), 1, address(testUniversalContract), revertContext);
    }

    function testDepositAndRevertZRC20AndCallUniversalContractFailsIfTargetIsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndRevert(address(zrc20), 1, address(0), revertContext);
    }

    function testDepositAndRevertZRC20AndCallUniversalContractFailsIfAmountIsZero() public {
        vm.prank(protocolAddress);
        vm.expectRevert(InsufficientZRC20Amount.selector);
        gateway.depositAndRevert(address(zrc20), 0, address(testUniversalContract), revertContext);
    }

    function testDepositAndRevertZRC20AndCallUniversalContract() public {
        uint256 balanceBefore = zrc20.balanceOf(address(testUniversalContract));
        assertEq(0, balanceBefore);

        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextDataRevert(revertContext);
        vm.prank(protocolAddress);
        gateway.depositAndRevert(address(zrc20), 1, address(testUniversalContract), revertContext);

        uint256 balanceAfter = zrc20.balanceOf(address(testUniversalContract));
        assertEq(1, balanceAfter);
    }

    function testDepositAndRevertZRC20AndCallUniversalContractFailsIfSenderIsNotProtocol() public {
        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.depositAndRevert(address(zrc20), 1, address(testUniversalContract), revertContext);
    }

    function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsProtocol() public {
        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndRevert(address(zrc20), 1, protocolAddress, revertContext);
    }

    function testDepositAndRevertZRC20AndCallUniversalContractFailsITargetIsGateway() public {
        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndRevert(address(zrc20), 1, address(gateway), revertContext);
    }

    function testDepositZETAAndCallUniversalContractFailsIfTargetIsZeroAddress() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, 1, address(0), message);
    }

    function testDepositZETAAndCallUniversalContractFailsIfTargetIsAmountIsZero() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(InsufficientZetaAmount.selector);
        gateway.depositAndCall(context, 0, address(zrc20), message);
    }

    function testDepositZETAAndCallUniversalContractFailsIfZeroAddress() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, 1, address(0), message);
    }

    function testDepositZETAAndCallUniversal() public {
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.depositAndCall(context, 1, address(0), message);
    }

    function testDepositZETAAndCallUniversalContract() public {
        uint256 amount = 1;
        uint256 protocolBalanceBefore = zetaToken.balanceOf(protocolAddress);
        uint256 gatewayBalanceBefore = zetaToken.balanceOf(address(gateway));
        uint256 destinationBalanceBefore = address(testUniversalContract).balance;
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextData(abi.encodePacked(gateway), protocolAddress, amount, address(gateway), "hello");
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, amount, address(testUniversalContract), message);

        uint256 protocolBalanceAfter = zetaToken.balanceOf(protocolAddress);
        assertEq(protocolBalanceBefore - amount, protocolBalanceAfter);

        uint256 gatewayBalanceAfter = zetaToken.balanceOf(address(gateway));
        assertEq(gatewayBalanceBefore, gatewayBalanceAfter);

        // Verify amount is transfered to destination
        assertEq(destinationBalanceBefore + amount, address(testUniversalContract).balance);
    }

    function testDepositZETAAndCallUniversalContractFailsIfSenderIsNotProtocol() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectRevert(CallerIsNotProtocol.selector);
        vm.prank(owner);
        gateway.depositAndCall(context, amount, address(testUniversalContract), message);
    }

    function testDepositZETAAndCallUniversalContractFailsIfTargetIsProtocol() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, amount, protocolAddress, message);
    }

    function testDepositZETAAndCallUniversalContractFailsIfTargetIsGateway() public {
        uint256 amount = 1;
        bytes memory message = abi.encode("hello");
        MessageContext memory context =
            MessageContext({ sender: abi.encodePacked(address(gateway)), senderEVM: protocolAddress, chainID: 1 });

        vm.expectRevert(InvalidTarget.selector);
        vm.prank(protocolAddress);
        gateway.depositAndCall(context, amount, address(gateway), message);
    }

    function testExecuteAbortUniversalContract() public {
        vm.expectEmit(true, true, true, true, address(testUniversalContract));
        emit ContextDataAbort(abortContext);
        vm.prank(protocolAddress);
        gateway.executeAbort(address(testUniversalContract), abortContext);
    }

    function testExecuteAbortUniversalContractFailsIfTargetIsZeroAddress() public {
        vm.prank(protocolAddress);
        vm.expectRevert(ZeroAddress.selector);
        gateway.executeAbort(address(0), abortContext);
    }
}
