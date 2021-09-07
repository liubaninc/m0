#!/bin/bash

chain_id=${1:-testnetA}

files=$(ls $(dirname $0)/$chain_id/*.yaml)
for file in $files
do
docker-compose -f $file stop
done