#!/usr/bin/env sh

set -e
mnemonic="key erupt service six thing spy noise heart giggle year oil fuel rival drop goat deal moral require knee pact bind brain word nuclear"
rm -rf ~/output
m0d testnet -n 1 -o ~/output --chain-id happy
echo "$mnemonic" | m0d keys add alice --recover --home ~/output/node0/.m0/

m0d start --grpc.address 0.0.0.0:9090 --home ~/output/node0/.m0/