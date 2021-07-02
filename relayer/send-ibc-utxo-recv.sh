#!/bin/bash
set -e

echo "before send-ibc-utxo... mars ==> earth"
echo "earth balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.earth/) --output json)"
echo "mars balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.mars/) --output json  --node tcp://127.0.0.1:26659)"
echo "hash $(m0d tx mibc send-ibc-utxo mibc channel-0 $(m0d keys show alice -a --home ~/.earth/) 13mibc/165B0BF0A59FDB92C95AD814618B96DF63B7E5AAF74F321E009FCE94F72A1895 --from alice --home ~/.mars --chain-id mars -y --node tcp://127.0.0.1:26659 | jq .txhash)"
sleep 10
echo "after send-ibc-utxo... mars ==> earth"
echo "earth balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.earth/) --output json)"
echo "mars balances $(m0d q bank balances $(m0d keys show alice -a --home ~/.mars/) --output json --node tcp://127.0.0.1:26659)"