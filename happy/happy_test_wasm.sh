#!/usr/bin/env bash

set -e

source $(dirname $0)/happy_utxo.sh
source $(dirname $0)/happy_wasm.sh

m0d keys add wuser
m0d keys add iuser

issue alice $(m0d keys show wuser -a) 10000 wuser
issue alice $(m0d keys show iuser -a) 10000 wuser

codefile=$(dirname $0)/counter.wasm

deploy wuser happyc "$codefile" "{\"creator\": \"someone\"}"

upgrade wuser happyc "$codefile"

invoke iuser happyc increase "{\"key\": \"someone\"}"
query happyc get "{\"key\": \"someone\"}"

invoke iuser happyc increase "{\"key\": \"someone\"}"
query happyc get "{\"key\": \"someone\"}"

invoke iuser happyc increase "{\"key\": \"someone\"}"
query happyc get "{\"key\": \"someone\"}"
