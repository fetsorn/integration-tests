#!/bin/bash

docker-compose up -d geth-bsc
docker-compose run golang sh -c 'cd /go/src/rh_tests/solidity/0.7; cp config.geth.json config.json; go run rh_tests && go test -v'


