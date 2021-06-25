#!/bin/bash

cd $(dirname "$0")/..

starport chain serve -c ./relayer/earth.yml > earth.log 2>&1

starport chain serve -c ./relayer/mars.yml  > mars.log 2>&1

rm -rf ~/.starport/

starport relayer configure -a \
--source-rpc "http://0.0.0.0:26657" \
--source-faucet "http://0.0.0.0:4500" \
--source-port "mibc" \
--source-version "mibc-1" \
--source-gasprice "0.0000025stake" \
--source-prefix "mc" \
--target-rpc "http://0.0.0.0:26659" \
--target-faucet "http://0.0.0.0:4501" \
--target-port "mibc" \
--target-version "mibc-1" \
--target-gasprice "0.0000025stake" \
--target-prefix "mc"

starport relayer connect