# integration-tests

The package is designed for integration testing of Gravity smart contracts: Nebula, IB-Port, LU-Port.

Host requirements: docker, docker-compose

Supported networks:

- Ethereum testnet ROPSTEN
- Ethereum private net (is launched in a Docker container)
- TRON testnet SHASTA
- WAVES testnet "stagenet"
- WAVES private net (is launched in a Docker container)

The source codes of contracts should be placed in the corresponding folders:

- solidity/0.7/Nebula|Token|IBPort|LUPort - for Ethereum testnet and privatenet
- Nebula|Token|IBPort|LUPort - for TRON
- waves/scripts/gravity|nebula (TODO: at this moment, it requires already an compiled base64, compilation from source code is work in progress).

When using testnet, you may need to pre-fund the addresses with test coins:

- for WAVES testnet STAGENET: 3MUBYCv4iHmmCA16z8GzKTSc3Fj1umn8Jb
- for Ethereum testnet ROPSTEN: 0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1
- for TRON testnet SHASTA: THxQonMCidbEgQSFxWcEgkyw1YfdpiqGz6

Run a test-*.sh bash-script with an appropriate name, for instance, test-eth-private.sh to test contracts in Ethereum private network.

The package already contains the compiled code of smart contracts. When modifying the source code, launch recompile.sh.
The compilation script will produce a build in a container for the Solidity compiler versions 5 and 7, as well as abigen utilities. The compilation may take a long time during the first run.
