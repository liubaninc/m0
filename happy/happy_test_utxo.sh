#!/usr/bin/env bash

set -e

source $(dirname $0)/happy_utxo.sh

m0d keys add user1
m0d keys add user2
m0d keys add user3

issue alice user1 1000 user1token
issue alice user2 1600 user2token

send user1 user3 200 user1token
send user2 user3 300 user2token

destroy user3 200 user1token
destroy user3 300 user2token

destroy user1 200 user1token
destroy user2 300 user2token

issue alice user1 400 user1token
issue alice user2 400 user2token