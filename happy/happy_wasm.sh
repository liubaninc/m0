#!/usr/bin/env bash

logfile=$(dirname $0)/happy_test_wasm.log
echo "" > "$logfile"

function deploy() {
  from=$1
  contract_name=$2
  contract_code_file=$3
  contract_init_args=$4

  echo "deploy $contract_name $contract_code_file $contract_init_args"

  before_contract=$(m0d query wasm show-contract "$contract_name")

  hash=$(m0d tx wasm deploy "$contract_name" "$contract_code_file" "$contract_init_args" --from "$from" --chain-id=happy --gas auto --gas-adjustment 1.5 -y --broadcast-mode block | jq ".txhash")
  tx=$(m0d query tx $(echo $hash | sed 's/\"//g') --output json)

  after_contract=$(m0d query wasm show-contract "$contract_name")

  echo "deploy $hash" >> "$logfile"
  echo $(echo $tx | jq .code) >> "$logfile"
  echo $(echo $tx | jq .raw_log) >> "$logfile"
  echo "==== $before_contract" >> "$logfile"
  echo "**** $after_contract" >> "$logfile"

  return 0
}

function upgrade() {
  from=$1
  contract_name=$2
  contract_code_file=$3

  echo "upgrade $contract_name $contract_code_file"

  before_contract=$(m0d query wasm show-contract "$contract_name")

  hash=$(m0d tx wasm upgrade "$contract_name" "$contract_code_file" --from "$from" --chain-id=happy --gas auto --gas-adjustment 1.5 -y --broadcast-mode block | jq ".txhash")
  tx=$(m0d query tx $(echo $hash | sed 's/\"//g') --output json)

  after_contract=$(m0d query wasm show-contract "$contract_name")

  echo "upgrade $hash" >> "$logfile"
  echo $(echo $tx | jq .code) >> "$logfile"
  echo $(echo $tx | jq .raw_log) >> "$logfile"
  echo "==== $before_contract" >> "$logfile"
  echo "**** $after_contract" >> "$logfile"

  return 0
}

function invoke() {
  from=$1
  contract_name=$2
  contract_method=$3
  contract_method_args=$4

  echo "invoke $contract_name $contract_method $contract_method_args"

  hash=$(m0d tx wasm invoke "$contract_name" "$contract_method" "$contract_method_args" --from "$from" --chain-id=happy --gas auto --gas-adjustment 1.5 -y --broadcast-mode block | jq ".txhash")
  tx=$(m0d query tx $(echo $hash | sed 's/\"//g') --output json)

  echo "invoke $hash" >> "$logfile"
  echo $(echo $tx | jq .code) >> "$logfile"
  echo $(echo $tx | jq .raw_log) >> "$logfile"

  return 0
}

function query() {
  contract_name=$1
  contract_method=$2
  contract_method_args=$3

  echo "query $contract_name $contract_method $contract_method_args"

  val=$(m0d query wasm query "$contract_name" "$contract_method" "$contract_method_args")

  echo "query $contract_name $contract_method $contract_method_args" >> "$logfile"
  echo "$val" >> "$logfile"
  return 0
}
