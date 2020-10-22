#!/bin/bash

docker-compose up -d ganache
docker-compose run golang sh -c 'cd /go/src/rh_tests/solidity/0.7; cp config.ganache.json config.json; go get -u github.com/Gravity-Tech/gravity-core@master; go run rh_tests && go test'


