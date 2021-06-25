#!/bin/bash

ps -ef | grep 0.0.0.0:366 | grep -v grep |  awk -F ' ' '{print $2}' | xargs kill -9

outputDir=${1:-./chains}

m0dPath=$(dirname "$0")/../build
if [ ! -d "$outputDir"/chainone ]
then
   "$m0dPath"/m0d testnet --output-dir "$outputDir"/chainone --num 1 --node-ip "127.0.0.1:26656" --chain-id chainone
fi

if [ ! -d "$outputDir"/chaintwo ]
then
   "$m0dPath"/m0d testnet --output-dir "$outputDir"/chaintwo --num 1 --node-ip "127.0.0.1:26656" --chain-id chaintwo
fi

sed -i .bak 's/enable = true/enable = false/g' "$outputDir"/chainone/node0/.m0/config/app.toml
sed -i .bak 's/prometheus = true/prometheus = false/g' "$outputDir"/chainone/node0/.m0/config/config.toml
nohup "$m0dPath"/m0d start --pruning=nothing --consensus.create_empty_blocks false --p2p.laddr tcp://0.0.0.0:36656 --rpc.laddr tcp://0.0.0.0:36657 --rpc.pprof_laddr 127.0.0.1:36060 --home "$outputDir"/chainone/node0/.m0 > "$outputDir"/chainone.log 2>&1 &

sed -i .bak 's/enable = true/enable = false/g' "$outputDir"/chaintwo/node0/.m0/config/app.toml
sed -i .bak 's/prometheus = true/prometheus = false/g' "$outputDir"/chaintwo/node0/.m0/config/config.toml
nohup "$m0dPath"/m0d start --pruning=nothing --consensus.create_empty_blocks false --p2p.laddr tcp://0.0.0.0:36666 --rpc.laddr tcp://0.0.0.0:36667 --rpc.pprof_laddr 127.0.0.1:36061 --home "$outputDir"/chaintwo/node0/.m0 > "$outputDir"/chaintwo.log 2>&1 &

starport relayer configure -a \
--source-rpc "http://0.0.0.0:36657" \
--source-faucet "" \
--source-port "mibc" \
--source-version "mibc-1" \
--source-gasprice "0.00025stake" \
--source-prefix "mc" \
--target-rpc "http://0.0.0.0:36667" \
--target-faucet "" \
--target-port "mibc" \
--target-version "mibc-1" \
--target-gasprice "0.00025stake" \
--target-prefix "mc"
