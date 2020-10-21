#!/bin/bash

bash symlink.sh

docker-compose run golang sh -c 'cd /go/src/rh_tests/solidity/0.7; cp config.bsc-testnet.json config.json; go run rh_tests && go test -v'
