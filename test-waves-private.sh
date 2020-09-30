#!/bin/bash

docker-compose up -d waves
docker-compose run golang sh -c 'cd /go/src/rh_tests/waves; cp config-private.json config.json; go test'


