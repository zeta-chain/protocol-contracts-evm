# ZetaChain Protocol Contracts

Contracts of official protocol contracts deployed by the core ZetaChain team.

## Learn more about ZetaChain

* Check our [website](https://www.zetachain.com/).
* Read our [docs](https://docs.zetachain.com/).

## Codebase

The protocol contracts codebase is separated into two sections:

- `zevm`: contracts deployed on ZetaChain
- `evm`: contracts deployed on EVM connected chains (Ethereum, Base, etc..)

**ZEVM Contracts**

- `GatewayZEVM`: entrypoints for interaction. The users call the gateway contract to initiate cctx. The gateway contract is also the contract calling the target when a smart contract call is initiated on a connected chain
- `ZRC20`: is a ERC20 compliant contract that represents fungible assets from connected chains
- `WZETA`: wrapped ZETA (fork of WETH)

**EVM Contracts**

- `GatewayEVM`: similar to GatewayZEVM for connected chains. Entrypoint for users.
- `ERC20Custody`: hold ERC20 assets being sent to ZetaChain
- `ZetaConnector`: manage ZETA for connected chains. There are two version of the contract:
    - Native: when the ZETA tokens are native to the chain (ZETA where initially deployed as a ERC20 on Ethereum, not emitted fully on ZetaChain). Use lock/unlock model.
    - Non-Native: when ZETA tokens where never native to the chains but withdrawn from ZetaChain. Use mint/burn model.
- TSS (EOA): Threshold-signature-scheme wallet, this address holds the gas token sent to ZetaChain (like Ethers)

[Learn more about the Gateway contracts](https://www.zetachain.com/docs/developers/evm/gateway/)

## Usage

**Install dependencies**

```shell
$ yarn
$ forge soldeer update
```

**Build**

```shell
$ forge build
```

**Test**

```shell
$ forge test
```

**Format**

```shell
$ forge fmt
```

**Deploy**

```shell
$ forge script script/<DeployScript>.s.sol:<DeployScript> --rpc-url <your_rpc_url> --private-key <your_private_key>
```

To view detailed instructions on how to deploy the contracts, please refer to the [Deployment Guide](./scripts/deploy/readme.md).

This guide covers all steps required to deploy the contracts, including environment setup, commands, and best practices.

## Community

[X](https://x.com/zetachain) | [Discord](https://discord.com/invite/zetachain) | [Telegram](https://t.me/zetachainofficial) | [Website](https://zetachain.com)
