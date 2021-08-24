#!/bin/bash
set -e

echo "send-ibc-utxo... earth ==> mars"
echo "before"
echo "earth chain balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.earth/) --output json)"
echo "mars chain balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.mars/) --output json  --node tcp://127.0.0.1:26659)"
m0d tx mibc send-ibc-utxo mibc channel-0 $(m0d keys show alice -a --home ~/.mars) 13m0token --from alice --home ~/.earth --chain-id earth -y
sleep 60
echo "after"
echo "earth chain  balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.earth/) --output json)"
echo "mars chain  balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.mars/) --output json --node tcp://127.0.0.1:26659)"

t=$(m0d q bank balances $(m0d keys show alice -a --home ~/.mars/) --output json --node tcp://127.0.0.1:26659 | jq .balances[1].denom | sed 's/\"//g')

echo "send-ibc-utxo... mars==> earth"
echo "before"
echo "earth chain balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.earth/) --output json)"
echo "mars chain balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.mars/) --output json  --node tcp://127.0.0.1:26659)"
m0d tx mibc send-ibc-utxo mibc channel-0 $(m0d keys show alice -a --home ~/.earth) 13$t --from alice --home ~/.mars --chain-id mars -y --node tcp://127.0.0.1:26659
sleep 60
echo "after"
echo "earth chain  balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.earth/) --output json)"
echo "mars chain  balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.mars/) --output json --node tcp://127.0.0.1:26659)"