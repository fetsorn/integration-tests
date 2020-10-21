#!/bin/bash

# TRON part needs to be copied to core
# and then referenced to.

eth_contracts_dest='solidity/0.7/source'
waves_contracts_dest='waves/scripts'

# ETH
if [ -d "$eth_contracts_dest" ]
then
  echo "Symlink to $eth_contracts_dest already exists"
else
  ln -s "$(pwd)/core/contracts/ethereum/contracts" "$eth_contracts_dest"
  echo "Created symlink from core to $eth_contracts_dest"
fi

# WAVES
if [ -d "$waves_contracts_dest" ]
then
  echo "Symlink to $waves_contracts_dest already exists"
else
  ln -s "$(pwd)/core/contracts/waves/script" "$waves_contracts_dest"
  echo "Created symlink from core to $waves_contracts_dest"
fi

# TRON
# ln -s "$(pwd)/core/contracts/ethereum/contracts" 'solidity/tron/source'

