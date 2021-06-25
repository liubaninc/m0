
m0path=$(dirname $0)../build

# 发行资产m0token
m0d keys add ibcuser1 --home $HOME/.earth

hash=$($m0path/m0d tx utxo issue $(m0d keys show "$ibcuser1" -a) 100000etoken  --from alice --chain-id=earth -y --broadcast-mode block  --home $HOME/.earth | jq ".txhash")

m0d tx mibc send-ibc-utxo mibc channel-0 mc1hfaex4s0gw6gjn4eg556jl9w9yvqaxy438d28l 10etoken --from ibcuser1 --chain-id earth --home ~/.earth