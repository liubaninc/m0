#!/bin/bash
set -e

echo "before send-ibc-utxo... earth ==> mars"
echo "earth balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.earth/) --output json)"
echo "mars balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.mars/) --output json  --node tcp://127.0.0.1:26659)"
echo "hash $(m0d tx mibc send-ibc-utxo mibc channel-0 $(m0d keys show alice -a --home ~/.mars) 13etoken --from alice --home ~/.earth --chain-id earth -y | jq .txhash)"
sleep 10
echo "after send-ibc-utxo... earth ==> mars"
echo "earth balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.earth/) --output json)"
echo "mars balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.mars/) --output json --node tcp://127.0.0.1:26659)"