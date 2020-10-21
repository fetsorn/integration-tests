#!/bin/bash

# TRON part needs to be copied to core
# and then referenced to.


# ETH
ln -s "$(pwd)/core/contracts/ethereum/contracts" 'solidity/0.7/source'

# WAVES
ln -s "$(pwd)/core/contracts/waves/script" 'waves/scripts'

# TRON
# ln -s "$(pwd)/core/contracts/ethereum/contracts" 'solidity/tron/source'

