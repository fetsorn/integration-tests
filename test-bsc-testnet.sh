#!/bin/bash

docker-compose run golang sh -c 'cd /go/src/rh_tests/solidity/0.7; cp config.bsc-testnet.json config.json; go get -u github.com/Gravity-Tech/gravity-core@master; go run rh_tests && go test -v'


