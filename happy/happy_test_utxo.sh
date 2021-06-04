#!/usr/bin/env bash

logfile=$(dirname $0)/happy_test_utxo.log
echo "" > $logfile

function issue() {
  from=$1
  to=$2
  amount=$3
  coin=$4

  echo "issue $from $to $amount $coin"


  before_supply=$(m0d query bank total --output json)
  before_token=$(m0d query utxo show-token $coin --output json)
  before_balance_token_from=$(m0d query bank balances $(m0d keys show $from -a) --denom $coin --output json)
  before_input_token_from=$(m0d query utxo list-input $(m0d keys show $from -a) $coin --count-total --output json)
  before_balance_token_to=$(m0d query bank balances $(m0d keys show $to -a) --denom $coin --output json)
  before_input_token_to=$(m0d query utxo list-input $(m0d keys show $to -a) $coin --count-total --output json)

  hash=$(m0d tx utxo issue $(m0d keys show $to -a) $amount$coin  --from $from --chain-id=happy -y --broadcast-mode block | jq ".txhash")
  tx=$(m0d query tx $(echo $hash | sed 's/\"//g'))

  after_supply=$(m0d query bank total --output json)
  after_token=$(m0d query utxo show-token $coin --output json)
  after_balance_token_from=$(m0d query bank balances $(m0d keys show $from -a) --denom $coin --output json)
  after_input_token_from=$(m0d query utxo list-input $(m0d keys show $from -a) $coin --count-total --output json)
  after_balance_token_to=$(m0d query bank balances $(m0d keys show $to -a) --denom $coin --output json)
  after_input_token_to=$(m0d query utxo list-input $(m0d keys show $to -a) $coin --count-total --output json)


  echo "issue $hash" >> $logfile
  echo "$tx" >> $logfile
  echo "total supply: issue $from $to $amount $coin]" >> $logfile
  echo "=== $before_supply" >> $logfile
  echo "*** $after_supply" >> $logfile

  echo "show token: issue $from $to $amount $coin]" >> $logfile
  echo "=== $before_token" >> $logfile
  echo "*** $after_token" >> $logfile

  echo "balance : issue $from $to $amount $coin]" >> $logfile
  echo "=== from $before_balance_token_from" >> $logfile
  echo "*** from $after_balance_token_from" >> $logfile
  echo "=== to $before_balance_token_to" >> $logfile
  echo "*** to $after_balance_token_to" >> $logfile

  echo "utxo : issue $from $to $amount $coin]" >> $logfile
  echo "=== from $before_input_token_from" >> $logfile
  echo "*** from $after_input_token_from" >> $logfile
  echo "=== to $before_input_token_to" >> $logfile
  echo "*** to $after_input_token_to" >> $logfile
  echo ""

  return 0
}

function send() {
  from=$1
  to=$2
  amount=$3
  coin=$4

  echo "send $from $to $amount $coin"

  before_supply=$(m0d query bank total --output json)
  before_token=$(m0d query utxo show-token $coin --output json)
  before_balance_token_from=$(m0d query bank balances $(m0d keys show $from -a) --denom $coin --output json)
  before_input_token_from=$(m0d query utxo list-input $(m0d keys show $from -a) $coin --count-total --output json)
  before_balance_token_to=$(m0d query bank balances $(m0d keys show $to -a) --denom $coin --output json)
  before_input_token_to=$(m0d query utxo list-input $(m0d keys show $to -a) $coin --count-total --output json)

  hash=$(m0d tx utxo send $(m0d keys show $to -a) $amount$coin  --from $from --chain-id=happy -y --broadcast-mode block | jq ".txhash")
  tx=$(m0d query tx $(echo $hash | sed 's/\"//g'))

  after_supply=$(m0d query bank total --output json)
  after_token=$(m0d query utxo show-token $coin --output json)
  after_balance_token_from=$(m0d query bank balances $(m0d keys show $from -a) --denom $coin --output json)
  after_input_token_from=$(m0d query utxo list-input $(m0d keys show $from -a) $coin --count-total --output json)
  after_balance_token_to=$(m0d query bank balances $(m0d keys show $to -a) --denom $coin --output json)
  after_input_token_to=$(m0d query utxo list-input $(m0d keys show $to -a) $coin --count-total --output json)


  echo "send $hash" >> $logfile
  echo "$tx" >> $logfile
  echo "total supply: send $from $to $amount $coin]" >> $logfile
  echo "=== $before_supply" >> $logfile
  echo "*** $after_supply" >> $logfile

  echo "show token: send $from $to $amount $coin]" >> $logfile
  echo "=== $before_token" >> $logfile
  echo "*** $after_token" >> $logfile

  echo "balance : send $from $to $amount $coin]" >> $logfile
  echo "=== from $before_balance_token_from" >> $logfile
  echo "*** from $after_balance_token_from" >> $logfile
  echo "=== to $before_balance_token_to" >> $logfile
  echo "*** to $after_balance_token_to" >> $logfile

  echo "utxo : send $from $to $amount $coin]" >> $logfile
  echo "=== from $before_input_token_from" >> $logfile
  echo "*** from $after_input_token_from" >> $logfile
  echo "=== to $before_input_token_to" >> $logfile
  echo "*** to $after_input_token_to" >> $logfile
  echo ""

  return 0
}

function destroy() {
  from=$1
  amount=$2
  coin=$3

  echo "destroy $from $amount $coin"

  before_supply=$(m0d query bank total --output json)
  before_token=$(m0d query utxo show-token $coin --output json)
  before_balance_token_from=$(m0d query bank balances $(m0d keys show $from -a) --denom $coin --output json)
  before_input_token_from=$(m0d query utxo list-input $(m0d keys show $from -a) $coin --count-total --output json)

  hash=$(m0d tx utxo destroy $amount$coin --from $from --chain-id=happy -y --broadcast-mode block | jq ".txhash")
  tx=$(m0d query tx $(echo $hash | sed 's/\"//g'))

  after_supply=$(m0d query bank total --output json)
  after_token=$(m0d query utxo show-token $coin --output json)
  after_balance_token_from=$(m0d query bank balances $(m0d keys show $from -a) --denom $coin --output json)
  after_input_token_from=$(m0d query utxo list-input $(m0d keys show $from -a) $coin --count-total --output json)


  echo "destroy $hash" >> $logfile
  echo "$tx" >> $logfile
  echo "total supply: destroy $from $to $amount $coin]" >> $logfile
  echo "=== $before_supply" >> $logfile
  echo "*** $after_supply" >> $logfile

  echo "show token: destroy $from $to $amount $coin]" >> $logfile
  echo "=== $before_token" >> $logfile
  echo "*** $after_token" >> $logfile

  echo "balance : destroy $from $to $amount $coin]" >> $logfile
  echo "=== from $before_balance_token_from" >> $logfile
  echo "*** from $after_balance_token_from" >> $logfile

  echo "utxo : destroy $from $to $amount $coin]" >> $logfile
  echo "=== from $before_input_token_from" >> $logfile
  echo "*** from $after_input_token_from" >> $logfile
  echo ""

  return 0
}

set -e

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