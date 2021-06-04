#!/usr/bin/env sh

logfile=$(dirname $0)/happy_test_wasm.log
echo "" > $logfile

function deploy() {
  from=$1
  contract_name=$2
  contract_code_file=$3
  contract_init_args=$4

  echo "deploy $contract_name $contract_code_file $contract_init_args"

  hash=$(m0d tx wasm deploy "$contract_name" "$contract_code_file" "$contract_init_args" --from "$from" --chain-id=happy --gas auto -y --broadcast-mode block | jq ".txhash")
  tx=$(m0d query tx $(echo $hash | sed 's/\"//g'))

  echo "deploy $hash" >> $logfile
  echo "$tx" >> $logfile

  return 0
}

function upgrade() {
  from=$1
  contract_name=$2
  contract_code_file=$3

  echo "upgrade $contract_name $contract_code_file"

  hash=$(m0d tx wasm upgrade "$contract_name" "$contract_code_file" --from "$from" --chain-id=happy --gas auto -y --broadcast-mode block | jq ".txhash")
  tx=$(m0d query tx $(echo $hash | sed 's/\"//g'))

  echo "upgrade $hash" >> $logfile
  echo "$tx" >> $logfile

  return 0
}

function invoke() {
  from=$1
  contract_name=$2
  contract_method=$3
  contract_method_args=$4

  echo "invoke $contract_name $contract_method $contract_method_args"

  hash=$(m0d tx wasm upgrade "$contract_name" "$contract_method" "$contract_method_args" --from "$from" --chain-id=happy --gas auto -y --broadcast-mode block | jq ".txhash")
  tx=$(m0d query tx $(echo $hash | sed 's/\"//g'))

  echo "invoke $hash" >> $logfile
  echo "$tx" >> $logfile

  return 0
}

set -e

m0d keys add wuser

codefile=$(dirname $0)/counter.wasm

deploy wuser happyc "$codefile" "{\"creator\": \"someone\"}"

upgrade wuser happyc "$codefile"

m0d keys add iuser

invoke iuser happyc increase "{\"key\": \"someone\"}"

upgrade iuser happyc "$codefile"
