#!/bin/bash

export PATH=${PWD}/../build:$PATH
. utils.sh

mnemonic="luxury hour naive special rib turtle glad discover accident forget estate virus super like strike country equip bar false three town game wet pigeon"
mnemonic=$(m0d keys mnemonic 2>&1)

OUTPUT=.
TAG=latest
NUM=1
CHAIN_ID=testnet
NODE_ID=node0
GENESIS_TIME=1615183366
VALIDATORS_IP=${CHAIN_ID}-${NODE_ID}:26657
VALIDATORS_KEY="$mnemonic"
NODE_KEY="$mnemonic"
SEEDS=
SEED_MODE=false
SYNCED=true
NODE_RPC=

function nodeFirst() {
  if [ ${SEED_MODE} ]; then
    echo "todo ..."
  fi
  mkdir -p  $OUTPUT/$CHAIN_ID
  cp ./template/docker-compose-first.yaml $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/TAG/${TAG}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/VALIDATORNUM/${NUM}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/CHAINID/${CHAIN_ID}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/GENESISTIME/${GENESIS_TIME}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/VALIDATORIP/${VALIDATORS_IP}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/VALIDATORKEY/${VALIDATORS_KEY}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/NODEID/${NODE_ID}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/NODEKEY/${NODE_KEY}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/SEEDNODES/${SEEDS}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/SEEDMODE/${SEED_MODE}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/NODERPC/${NODE_RPC}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/SYNCED/${SYNCED}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
}

function nodeMore() {
  if [ -z "${NODE_RPC}" ]; then
    echo "-node was empty, please set"
    printHelp more
    exit 0
  fi
  if [ -z "${SEEDS}" ]; then
    echo "-seeds was empty, please set"
    printHelp more
    exit 0
  fi
  mkdir -p  $OUTPUT/$CHAIN_ID
  cp ./template/docker-compose-more.yaml $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/TAG/${TAG}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/VALIDATORNUM/${NUM}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/CHAINID/${CHAIN_ID}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/GENESISTIME/${GENESIS_TIME}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/VALIDATORIP/${VALIDATORS_IP}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/VALIDATORKEY/${VALIDATORS_KEY}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/NODEID/${NODE_ID}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/NODEKEY/${NODE_KEY}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/SEEDNODES/${SEEDS}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/SEEDMODE/${SEED_MODE}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/NODERPC/${NODE_RPC}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  sed -i '' "s/SYNCED/${SYNCED}/g" $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
}

## Parse mode
if [[ $# -lt 1 ]] ; then
  printHelp
  exit 0
else
  MODE=$1
  shift
fi

while [[ $# -ge 1 ]] ; do
  key="$1"
  case $key in
  -h )
    printHelp $MODE
    exit 0
    ;;
  -tag )
    TAG="$2"
    shift
    ;;
  -chain-id )
    CHAIN_ID="$2"
    shift
    ;;
  -genesis-time )
    GENESIS_TIME="$2"
    shift
    ;;
  -num )
    NUM="$2"
    shift
    ;;
  -validator-ip )
    VALIDATORS_IP="$2"
    shift
    ;;
  -validator-key )
    VALIDATORS_KEY="$2"
    shift
    ;;
  -id )
    NODE_ID="$2"
    shift
    ;;
  -key )
    NODE_KEY="$2"
    shift
    ;;
  -seeds )
    SEEDS="$2"
    shift
    ;;
  -node )
    NODE_RPC="$2"
    shift
    ;;
  -output )
    OUTPUT="$2"
    shift
    ;;
  -seed )
    SEED_MODE=true
    ;;
  -unsync )
    SYNCED=false
    ;;
  * )
    errorln "Unknown flag: $key"
    printHelp
    exit 1
    ;;
  esac
  shift
done

if [ "${MODE}" == "first" ]; then
  nodeFirst
elif [ "${MODE}" == "more" ]; then
  nodeMore
else
  printHelp
  exit 1
fi