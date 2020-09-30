# integration-tests

This integration tests framework is designed for the purpose of writing and testing integrations of new Gravity target chains, where the data is supplied to by the Gravity network. It simplifies making modifications to the logic of USER-SC or the rules in NEBULA-SC without the need to deploy and update a full custom Gravity network. The framework helps to solve the task of initial configuration and functionality of Gravity nodes not directly related to target chains.

The standard development flow for integrating a new target chain consists of four steps: 

1. Implement SYSTEM-SC and NEBULA-SC contracts on a specific blockchain;
2. Create an identical set of integration tests for testing the contracts on the blockchain;
3. Run the integration tests to verify that the contracts are implemented properly;
4. Add an integration of the target chain into a Gravity Core interface `IBlockchainAdaptor`. 

The integration test framework can greatly facilitate the 2d and the 3d steps of target chain integration.

--- 

Host requirements: docker, docker-compose

Supported networks:

- Ethereum testnet ROPSTEN
- Ethereum private net (is launched in a Docker container)
- TRON testnet SHASTA
- WAVES testnet "stagenet"
- WAVES private net (is launched in a Docker container)

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
The compilation script will produce a build in a container for the Solidity compiler versions 5 and 7, as well as abigen utilities. The compilation may take a long time during the first run.
