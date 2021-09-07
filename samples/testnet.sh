#!/bin/bash
a=`uname  -a`
b="Darwin"
c="centos"
d="ubuntu"

cd "$(dirname "$0")"

function printHelp() {
  USAGE="$1"
  if [ "$USAGE" == "first" ]; then
    println "Usage: "
    println "  network.sh \033[0;32mfirst\033[0m [Flags]"
    println
    println "    Flags:"
	  println "    -up  true直接在本机启动网络,默认不启动"
    println "    -tag <image tag> -  m0镜像tag"
    println "    -chain-id <chain name> - 链名称(默认testnet)"
    println "    -genesis-time <timestamp> - 创世块时间戳"
    println "    -num <num> - 验证人个数"
    println "    -validator-ip <ips> - 验证人节点列表 host:port,host:port"
    println "    -validator-key <keys> -  验证人私钥的助记词列表"
    println "    -id <node moniker> - 节点名字"
    println "    -key <node key> - 节点私钥助记词"
    println "    -seeds <seeds> - 种子节点列表 ID@host:port,ID@host:port"
    println "    -peers <peers> - 节点列表 ID@host:port,ID@host:port"
    println "    -seed - 节点启动种子模式"
    println "    -sync - 启动同步程序，开启浏览器、钱包服务,默认开启"
    println "    -automatic - 启动自动发送交易,默认开启"
    println "    -output <directory> - 结果输出目录(默认当前目录)"
    println
    println "    -h - Print this message"
    println
    println " Examples:"
    println "   network.sh first --chain-id tttt --node-id n0"
  elif [ "$USAGE" == "more" ]; then
    println "Usage: "
    println "  network.sh \033[0;32mmore\033[0m [Flags]"
    println
    println "    Flags:"
    println "    -up 直接在本机启动网络,默认不启动"
    println "    -tag <image tag> -  m0镜像tag"
    println "    -chain-id <chain name> - 链名称"
    println "    -id <node moniker> - 节点名字 node1"
    println "    -key <node key> - 节点私钥助记词"
    println "    -seeds <seeds> - 种子节点列表 ID@host:26656,ID@host:26656"
    println "    -peers <peers> - 节点列表 ID@host:port,ID@host:port"
    println "    -seed - 节点启动种子模式"
    println "    -node <node rpc> - 可访问的节点RPC host:26657"
    println "    -sync - 启动同步程序，开启浏览器、钱包服务(默认开启)"
    println "    -automatic - 启动自动发送交易(默认开启)"
    println "    -output <directory> - 结果输出目录(默认当前目录)"
    println
    println "    -h - Print this message"
    println " Examples:"
    println "   network.sh more --chain-id tttt --node-id n0"
  elif [ "$USAGE" == "mysql" ]; then
    println "Usage: "
    println "  network.sh \033[0;32mmysql\033[0m [Flags]"
    println
    println "    Flags:"
    println "    -up 直接在本机启动网络,默认不启动"
    println "    -chain-id <chain name> - 链名称"
    println "    -id <node moniker> - 节点名字 node1"
    println "    -output <directory> - 结果输出目录(默认当前目录)"
    println
    println "    -h - Print this message"
    println " Examples:"
    println "   network.sh mysql --chain-id tttt --node-id n0"
  elif [ "$USAGE" == "ca" ]; then
    println "Usage: "
    println "  network.sh \033[0;32mca\033[0m [Flags]"
    println
    println "    Flags:"
    println "    -up 直接在本机启动网络,默认不启动"
    println "    -chain-id <chain name> - 链名称"
    println "    -id <node moniker> - 节点名字 node1"
    println "    -output <directory> - 结果输出目录(默认当前目录)"
    println
    println "    -h - Print this message"
    println " Examples:"
    println "   network.sh ca --chain-id tttt --node-id n0"
  else
    println "Usage: "
    println "  network.sh <Mode> [Flags]"
    println "    Modes:"
    println "      \033[0;32mfirst\033[0m - 生成节点启动的docker-compose yaml, 适用于知晓链初始配置所需参数的启动节点"
    println "      \033[0;32mmore\033[0m - 生成节点启动的docker-compose yaml, 适用于非首个节点的启动节点"
    println "      \033[0;3232mmysql\033[0m - mysql启动的docker-compose yaml"
    println "      \033[0;3232mca\033[0m - ca启动的docker-compose yaml"
    println
    println "    Flags:"
    println "    Used with \033[0;32mnetwork.sh first\033[0m, \033[0;32mnetwork.sh more\033[0m:"
    println "    -up 直接在本机启动网络,默认不启动"
    println "    -tag <image tag> -  m0镜像tag"
    println "    -chain-id <chain name> - 链名称(默认testnet)"
    println "    -genesis-time <timestamp> - 创世块时间戳"
    println "    -num <num> - 验证人个数"
    println "    -validator-ip <ips> - 验证人节点列表 host:port,host:port"
    println "    -validator-key <keys> -  验证人私钥的助记词列表"
    println "    -id <node moniker> - 节点名字"
    println "    -key <node key> - 节点私钥助记词"
    println "    -seeds <seeds> - 种子节点列表 ID@host:port,ID@host:port"
    println "    -peers <peers> - 节点列表 ID@host:port,ID@host:port"
    println "    -seed - 节点启动种子模式"
    println "    -sync - 启动同步程序，开启浏览器、钱包服务"
    println "    -automatic - 启动自动发送交易"
    println "    -output <directory> - 结果输出目录(默认当前目录)"
    println
    println "    -h - Print this message"
    println
    println " Examples:"
    println "   network.sh first --chain-id tttt --node-id n0"
    println "   network.sh more --chain-id tttt --node-id n0"
    println "   network.sh mysql --chain-id tttt --node-id n0"
    println "   network.sh ca --chain-id tttt --node-id n0"
  fi
}

# println echos string
function println() {
  echo -e "$1"
}

# errorln echos i red color
function errorln() {
  println "${C_RED}${1}${C_RESET}"
}

# successln echos in green color
function successln() {
  println "${C_GREEN}${1}${C_RESET}"
}

# infoln echos in blue color
function infoln() {
  println "${C_BLUE}${1}${C_RESET}"
}

# warnln echos in yellow color
function warnln() {
  println "${C_YELLOW}${1}${C_RESET}"
}

# fatalln echos in red color and exits with fail status
function fatalln() {
  errorln "$1"
  exit 1
}

mnemonic=$(m0d keys mnemonic 2>&1)

OUTPUT=.
TAG=latest
NUM=1
CHAIN_ID=testnet
NODE_ID=node0
GENESIS_TIME=1615183366
VALIDATORS_IP=
VALIDATORS_KEY="$mnemonic"
NODE_KEY="$mnemonic"
DEV_MODE=false
SEEDS=
NODES=
SEED_MODE=false
SYNCED=false
AUTOMATIC=false
NODE_RPC=
UP=false

function template() {
  if [[ $a =~ $b ]];then
    sed -i '' "s/TAG/${TAG}/g" "$1"
    sed -i '' "s/VALIDATORNUM/${NUM}/g" "$1"
    sed -i '' "s/CHAINID/${CHAIN_ID}/g" "$1"
    sed -i '' "s/GENESISTIME/${GENESIS_TIME}/g" "$1"
    sed -i '' "s/VALIDATORIP/${VALIDATORS_IP}/g" "$1"
    sed -i '' "s/VALIDATORKEY/${VALIDATORS_KEY}/g" "$1"
    sed -i '' "s/NODEID/${NODE_ID}/g" "$1"
    sed -i '' "s/NODEKEY/${NODE_KEY}/g" "$1"
    sed -i '' "s/DEVMODE/${DEV_MODE}/g" "$1"
    sed -i '' "s/SEEDNODES/${SEEDS}/g" "$1"
    sed -i '' "s/PEERNODES/${PEERS}/g" "$1"
    sed -i '' "s/SEEDMODE/${SEED_MODE}/g" "$1"
    sed -i '' "s/NODERPC/${NODE_RPC}/g" "$1"
    sed -i '' "s/SYNCED/${SYNCED}/g" "$1"
    sed -i '' "s/AUTOMATIC/${AUTOMATIC}/g" "$1"
  else
    sed -i "s/TAG/${TAG}/g" "$1"
    sed -i "s/VALIDATORNUM/${NUM}/g" "$1"
    sed -i "s/CHAINID/${CHAIN_ID}/g" "$1"
    sed -i "s/GENESISTIME/${GENESIS_TIME}/g" "$1"
    sed -i "s/VALIDATORIP/${VALIDATORS_IP}/g" "$1"
    sed -i "s/VALIDATORKEY/${VALIDATORS_KEY}/g" "$1"
    sed -i "s/NODEID/${NODE_ID}/g" "$1"
    sed -i "s/NODEKEY/${NODE_KEY}/g" "$1"
    sed -i "s/DEVMODE/${DEV_MODE}/g" "$1"
    sed -i "s/SEEDNODES/${SEEDS}/g" "$1"
    sed -i "s/PEERNODES/${PEERS}/g" "$1"
    sed -i "s/SEEDMODE/${SEED_MODE}/g" "$1"
    sed -i "s/NODERPC/${NODE_RPC}/g" "$1"
    sed -i "s/SYNCED/${SYNCED}/g" "$1"
    sed -i "s/AUTOMATIC/${AUTOMATIC}/g" "$1"
  fi
}

function nodeMysql() {
  infoln "生成MYSQL服务$NODE_ID... yaml:$OUTPUT/${CHAIN_ID}/$NODE_ID.yaml"
  mkdir -p  $OUTPUT/$CHAIN_ID
  cp docker-compose-mysql.yaml $OUTPUT/${CHAIN_ID}/${NODE_ID}.yaml
  if [[ $a =~ $b ]];then
    sed -i '' "s/CHAINID/${CHAIN_ID}/g" $OUTPUT/${CHAIN_ID}/${NODE_ID}.yaml
    sed -i '' "s/NODEID/${NODE_ID}/g" $OUTPUT/${CHAIN_ID}/${NODE_ID}.yaml
  else
    sed -i "s/CHAINID/${CHAIN_ID}/g" $OUTPUT/${CHAIN_ID}/${NODE_ID}.yaml
    sed -i "s/NODEID/${NODE_ID}/g" $OUTPUT/${CHAIN_ID}/${NODE_ID}.yaml
  fi
}

function nodeCA() {
  infoln "生成CA服务$NODE_ID... yaml:$OUTPUT/${CHAIN_ID}/$NODE_ID.yaml"
  mkdir -p  $OUTPUT/$CHAIN_ID
  cp docker-compose-ca.yaml $OUTPUT/${CHAIN_ID}/${NODE_ID}.yaml
  if [[ $a =~ $b ]];then
    sed -i '' "s/CHAINID/${CHAIN_ID}/g" $OUTPUT/${CHAIN_ID}/${NODE_ID}.yaml
    sed -i '' "s/NODEID/${NODE_ID}/g" $OUTPUT/${CHAIN_ID}/${NODE_ID}.yaml
  else
    sed -i "s/CHAINID/${CHAIN_ID}/g" $OUTPUT/${CHAIN_ID}/${NODE_ID}.yaml
    sed -i "s/NODEID/${NODE_ID}/g" $OUTPUT/${CHAIN_ID}/${NODE_ID}.yaml
  fi
}

function nodeFirst() {
  nid=$(echo "$NODE_KEY"	| m0d init ss --overwrite --recover 2>&1 | jq .node_id | sed 's/"//g')@$CHAIN_ID-$NODE_ID:26656
  infoln "生成节点$NODE_ID($nid)... yaml:$OUTPUT/${CHAIN_ID}/$NODE_ID.yaml"
  mkdir -p  $OUTPUT/$CHAIN_ID
  cp docker-compose-first.yaml $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  template $OUTPUT/$CHAIN_ID/$NODE_ID.yaml

  if [ $UP == true ];then
    infoln "启动节点$NODE_ID($nid)..."
    docker-compose -f $OUTPUT/$CHAIN_ID/$NODE_ID.yaml up -d 
  fi
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

  nid=$(echo "$NODE_KEY"	| m0d init ss --overwrite --recover 2>&1 | jq .node_id | sed 's/"//g')@$CHAIN_ID-$NODE_ID:26656
  infoln "生成节点$NODE_ID($nid)... yaml:$OUTPUT/${CHAIN_ID}/$NODE_ID.yaml"
  mkdir -p  $OUTPUT/$CHAIN_ID
  cp docker-compose-more.yaml $OUTPUT/${CHAIN_ID}/$NODE_ID.yaml
  template $OUTPUT/$CHAIN_ID/$NODE_ID.yaml
  
  if [ $UP == true ];then
    infoln "启动节点$NODE_ID($nid)..."
    docker-compose -f $OUTPUT/$CHAIN_ID/$NODE_ID.yaml up -d 
  fi
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
  -peers )
    PEERS="$2"
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
  -dev )
    DEV_MODE=true
    ;;
  -seed )
    SEED_MODE=true
    ;;
  -sync )
    SYNCED=true
    ;;
  -automatic )
    AUTOMATIC=true
    ;;
  -up )
    UP=true
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
elif [ "${MODE}" == "mysql" ]; then
  nodeMysql
elif [ "${MODE}" == "ca" ]; then
  nodeCA
else
  printHelp
  exit 1
fi