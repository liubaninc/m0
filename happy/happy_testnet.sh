#!/usr/bin/env bash

set -e

mnemonic="spray gallery hazard hidden math unusual butter ranch chief enter party huge spread priority pony indoor hover essay melody crane area bomb voice dress"

home=~/happy/.m0
# first
echo "$mnemonic" | m0d init node0 --recover --home $home 
nodeKey=$(md5sum $home/config/node_key.json | awk '{print $4}')
validatorKey=$(md5sum $home/config/node_key.json | awk '{print $4}')
rm -rf $home

# more
echo "$mnemonic" | m0d init node0 --recover --home $home 
nodeKey2=$(md5sum $home/config/node_key.json | awk '{print $4}')
validatorKey2=$(md5sum $home/config/node_key.json | awk '{print $4}')
rm -rf $home

if [ "$nodeKey" != "$nodeKey2" ];then
  echo "error node key file diff"
  exit 1
fi

if [ "$validatorKey" != "$validatorKey2" ];then
  echo "error validator key file diff"
  exit 1
fi


genesisTime=1623119877
chain=222
num=1
outputDir=happynet
m0d testnet --num $num --genesis-time $genesisTime --chain-id $chain --node-validator-key "$mnemonic" --output-dir $outputDir 
genesis=$(md5sum $outputDir/node0/.m0/config/genesis.json | awk '{print $4}')
rm -rf $outputDir

m0d testnet --num $num --genesis-time $genesisTime --chain-id $chain --node-validator-key "$mnemonic" --output-dir $outputDir 
genesis2=$(md5sum $outputDir/node0/.m0/config/genesis.json | awk '{print $4}')
rm -rf $outputDir
if [ "$genesis" != "$genesis2" ];then
  echo "error genesis file diff"
  exit 1
fi