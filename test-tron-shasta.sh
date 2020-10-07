#!/bin/bash

docker-compose run golang sh -c 'cd /go/src/rh_tests/solidity/tron; go run rh_tests && go test'

