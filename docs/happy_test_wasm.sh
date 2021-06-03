#!/usr/bin/env sh

m0d keys add user1
m0d keys add user2

m0d utxo issue $(m0d keys show user1 -a) 1000user1token $(m0d keys show user1 -a) 2000user2token --from alice --gas=auto

m0d utxo issue $(m0d keys show user1 -a) 1000user1token $(m0d keys show user1 -a) 2000user2token --from alice --gas=auto