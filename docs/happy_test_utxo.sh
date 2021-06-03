#!/usr/bin/env sh

function issue() {
  from=$1
  to=$2
  amount=$3
  coin=$4

  echo "issue before: $from ---> $to $amount $coin]"
  m0d query bank total $coin
  m0d query utxo show-token $coin
  m0d query bank balances $(m0d keys show $from -a)
  m0d query utxo list-input $(m0d keys show $from -a) $coin
  m0d query bank balances $(m0d keys show $to -a)
  m0d query utxo list-input $(m0d keys show $to -a) $coin

  m0d tx utxo issue $(m0d keys show $to -a) $amount  --from $from | jq ".txhash" |  xargs $(sleep 6) m0d query tx

  echo "issue after: $from ---> $to $amount $coin]"
  m0d query bank total $coin
  m0d query utxo show-token $coin
  m0d query bank balances $(m0d keys show $from -a)
  m0d query utxo list-input $(m0d keys show $from -a) $coin
  m0d query bank balances $(m0d keys show $to -a)
  m0d query utxo list-input $(m0d keys show $to -a) $coin
  return 0
}

m0d keys add user1
m0d keys add user2

issue alice user1 1000 user1token
issue alice user2 1000 user2token
