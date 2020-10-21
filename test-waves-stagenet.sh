#!/bin/bash

bash symlink.sh

docker-compose run golang sh -c 'cd /go/src/rh_tests/waves; cp config-stagenet.json config.json; go test'
