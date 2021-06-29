#!/usr/bin/env sh

set -e
mnemonic="key erupt service six thing spy noise heart giggle year oil fuel rival drop goat deal moral require knee pact bind brain word nuclear"
rm -rf ~/.m0
m0d init happys --chain-id=happy
echo "$mnemonic" | m0d keys add alice --recover
m0d add-genesis-account $(m0d keys show alice -a) 100000000stake
m0d gentx alice 100000000stake --chain-id=happy
m0d collect-gentxs
sed -i 's/enabled = false/enabled = true/g' ~/.m0/config/app.toml
sed -i 's/prometheus-retention-time = 0/prometheus-retention-time = 10/g' ~/.m0/config/app.toml
sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "60s"/g' ~/.m0/config/config.toml
sed -i 's/prometheus = false/prometheus = true/g' ~/.m0/config/config.toml
m0d start --pruning nothing --grpc.address 0.0.0.0:9090