#!/usr/bin/env sh

set -e
rm -rf ~/.m0
m0d init happys --chain-id=happy
m0d keys add alice
m0d add-genesis-account $(m0d keys show alice -a) 100000000stake
m0d gentx alice 100000000stake --chain-id=happy
m0d collect-gentxs
sed -i .bak 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "60s"/g' ~/.m0/config/config.toml
m0d start --pruning nothing --grpc.address 0.0.0.0:9090