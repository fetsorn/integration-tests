# integration-tests

This integration tests framework is designed for the purpose of writing and testing integrations of new Gravity target chains, where the data is supplied to by the Gravity network. It simplifies making modifications to the logic of USER-SC or the rules in NEBULA-SC without the need to deploy and update a full custom Gravity network. The framework helps to solve the task of initial configuration and functionality of Gravity nodes not directly related to target chains.


The standard development flow for integrating a new target chain consists of four steps: 

1. Implement SYSTEM-SC and NEBULA-SC contracts on a specific blockchain;
2. Create an identical set of integration tests for testing the contracts on the blockchain;
3. Run the integration tests to verify that the contracts are implemented properly;
4. Add an integration of the target chain into a Gravity Core interface `IBlockchainAdaptor`. 

The integration test framework can greatly facilitate the 2d and the 3d steps of target chain integration.

## Key concepts
The integration framework is focused on testing different aspects of the functionality of [SuSy](https://arxiv.org/ftp/arxiv/papers/2008/2008.13515.pdf), a blockchain-agnostic cross-chain asset transfer gateway protocol based on the Gravity protocol. As a reliable foundation for gateways, [Gravity](https://gravity.tech) allows them to remain trustless and decentralized. A crucial functionality required for cross-chain communication is the most basic procedure: cross-chain asset swaps. Let us start by introducing key terminology that we will use to explain Gravity, SuSy and their interaction:

ORIGIN-CHAIN: a blockchain network from which a transfer originates. That is, in this network, tokens are locked and unlocked.

DESTINATION-CHAIN: a blockchain to which transfers are made from the ORIGIN-CHAIN. Issuance and burning of wrapped tokens take place on this network.

IB-PORT is a smart contract in DESTINATION-CHAIN ​​that implements the functionality of issuance and burning of the wrapped token.

LU-PORT is a smart contract in ORIGIN-CHAIN that locks and unlocks the original token.

NEBULA-SC is one of the main architectural units of the Gravity protocol, a smart contract that accepts and verifies data from Gravity oracles. It implements checks of data relevance (blockchain height), availability of appropriate cryptographic signatures and threshold signature rules for transmitted data.

USER-SC is one of the main architectural units of the Gravity protocol. It is a smart contract that accepts data verified in NEBULA-SC and produces an action that is part of a custom application. In the case of SuSy, LU-PORT and IB-PORT are examples of USER-SC.

PULSE-TX is a transaction that will transfer hash from data to NEBULA-SC with  necessary signatures for verification and registration.

SEND-DATA-TX is a transaction that transfers data verified and registered in NEBULA-SC to USER-SC.

## Running tests

Host requirements: docker, docker-compose

Supported networks:

- Ethereum testnet ROPSTEN
- Ethereum private net (launched in a Docker container)
- TRON testnet SHASTA
- WAVES testnet "stagenet"
- WAVES private net (launched in a Docker container)

The source code of SYSTEM-SC and NEBULA-SC should be placed in the corresponding folders:

- solidity/0.7/Nebula|Token|IBPort|LUPort - for Ethereum testnet and privatenet
- Nebula|Token|IBPort|LUPort - for TRON
- waves/scripts/gravity|nebula (TODO: at this moment, it requires an already compiled base64, compilation from source code is a work in progress).

When using testnet, you may need to pre-fund the below addresses with test coins:

- for WAVES testnet STAGENET: 3MUBYCv4iHmmCA16z8GzKTSc3Fj1umn8Jb
- for Ethereum testnet ROPSTEN: 0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1
- for TRON testnet SHASTA: THxQonMCidbEgQSFxWcEgkyw1YfdpiqGz6

Run a `test-*.sh` bash-script with an appropriate name, for instance, test-eth-private.sh to test contracts in Ethereum private network.

The package already contains a compiled code of smart contracts. When modifying the source code of the contracts, launch `recompile.sh`.
The compilation script will produce a build in a container for the Solidity compiler versions 5 and 7, as well as abigen utilities: note that the contracts have a pre-generated ABI (application binary interface). These can be found, for instance, in `rh_tests/api/nebula` for NEBULA-SC, or `rh_tests/api/ibport` for the IB-Port (an example of a USER-SC). The compilation may take a long time during the first run.
