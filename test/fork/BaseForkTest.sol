// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Test.sol";
import "forge-std/Vm.sol";

abstract contract BaseForkTest is Test {
    // RPCs of all supported chains.
    string public ETHEREUM_RPC_URL = vm.envString("ETHEREUM_RPC_URL");
    string public BSC_RPC_URL = vm.envString("BSC_RPC_URL");
    string public POLYGON_RPC_URL = vm.envString("POLYGON_RPC_URL");
    string public BASE_RPC_URL = vm.envString("BASE_RPC_URL");
    string public ARBITRUM_RPC_URL = vm.envString("ARBITRUM_RPC_URL");
    string public AVALANCHE_RPC_URL = vm.envString("AVALANCHE_RPC_URL");
    string public ZETACHAIN_RPC_URL = vm.envString("ZETACHAIN_RPC_URL");

    // The fork identifiers.
    uint256 ethereumForkId;
    uint256 bscForkId;
    uint256 polygonForkId;
    uint256 baseForkId;
    uint256 arbitrumForkId;
    uint256 avalancheForkId;
    uint256 zetachainForkId;

    // Mainnet admins.
    address constant ETHEREUM_ADMIN = 0xaeB6dDB7708467814D557e340283248be8E43124;
    address constant BSC_ADMIN = 0xaf28a257D292e7f0E531073f70a175b57E0261a8;
    address constant POLYGON_ADMIN = 0x7828F92E7d79E141189f24C98aceF71Bc07bad3f;
    address constant BASE_ADMIN = 0xF43CF8b3F3D22d4cC33f964c59076eaB2F8A108E;
    address constant ARBITRUM_ADMIN = 0xdE3fb63723f0EEed8967ff9124e1c3bA89871b03;
    address constant AVALANCHE_ADMIN = 0xdE3fb63723f0EEed8967ff9124e1c3bA89871b03;
    address constant ZETACHAIN_ADMIN = 0x01d8207AfF7a7d029114Ee35afc72d0e133B7a0A;

    struct ChainConfig {
        uint256 forkId;
        address contractAddress;
        address admin;
        string rpcUrl;
        string name;
    }

    ChainConfig[] public chains;

    function setUp() public virtual {
        ethereumForkId = vm.createFork(ETHEREUM_RPC_URL);
        bscForkId = vm.createFork(BSC_RPC_URL);
        polygonForkId = vm.createFork(POLYGON_RPC_URL);
        baseForkId = vm.createFork(BASE_RPC_URL);
        arbitrumForkId = vm.createFork(ARBITRUM_RPC_URL);
        avalancheForkId = vm.createFork(AVALANCHE_RPC_URL);
        zetachainForkId = vm.createFork(ZETACHAIN_RPC_URL);

        _setupChains();
    }

    function _setupChains() internal virtual;

    function testForkIdDiffer() public view {
        assert(ethereumForkId != bscForkId);
        assert(bscForkId != polygonForkId);
        assert(polygonForkId != baseForkId);
        assert(baseForkId != arbitrumForkId);
        assert(arbitrumForkId != avalancheForkId);
        assert(avalancheForkId != zetachainForkId);
    }

    function testCanSwitchForks() public {
        for (uint256 i = 0; i < chains.length; i++) {
            vm.selectFork(chains[i].forkId);
            assertEq(vm.activeFork(), chains[i].forkId);
        }
    }

    function testUpgradeGatewayOnAllChains() public {
        for (uint256 i = 0; i < chains.length; i++) {
            console.log("\n=== Testing upgrade on", chains[i].name, "===");
            _testUpgradeContract(chains[i]);
        }
    }

    function _testUpgradeContract(ChainConfig memory config) internal virtual;

    function _getImplementation(address proxy) internal view returns (address) {
        bytes32 slot = 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;
        return address(uint160(uint256(vm.load(proxy, slot))));
    }
}
