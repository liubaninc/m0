#!/bin/bash
function printHelp() {
  USAGE="$1"
  if [ "$USAGE" == "first" ]; then
    println "Usage: "
    println "  network.sh \033[0;32mfirst\033[0m [Flags]"
    println
    println "    Flags:"
    println "    -tag <image tag> -  m0镜像tag"
    println "    -chain-id <chain name> - 链名称"
    println "    -genesis-time <timestamp> - 创世块时间戳"
    println "    -num <num> - 验证人个数"
    println "    -validator-ip <ips> - 验证人节点列表 host:port,host:port"
    println "    -validator-key <keys> -  验证人私钥的助记词列表"
    println "    -id <node moniker> - 节点名字"
    println "    -key <node key> - 节点私钥助记词"
    println "    -seeds <seeds> - 种子节点列表 ID@host:port,ID@host:port"
    println "    -seed - 节点启动种子模式"
    println "    -unsync - 启动同步程序，开启浏览器、钱包服务"
    println "    -automatic - 启动自动发送交易"
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
    println "    -tag <image tag> -  m0镜像tag"
    println "    -chain-id <chain name> - 链名称"
    println "    -id <node moniker> - 节点名字"
    println "    -key <node key> - 节点私钥助记词"
    println "    -seeds <seeds> - 种子节点列表 ID@host:port,ID@host:port"
    println "    -seed - 节点启动种子模式"
    println "    -node <node rpc> - 可访问的节点RPC"
    println "    -unsync - 启动同步程序，开启浏览器、钱包服务"
    println "    -automatic - 启动自动发送交易"
    println "    -output <directory> - 结果输出目录(默认当前目录)"
    println
    println "    -h - Print this message"
    println " Examples:"
    println "   network.sh more --chain-id tttt --node-id n0"
  else
    println "Usage: "
    println "  network.sh <Mode> [Flags]"
    println "    Modes:"
    println "      \033[0;32mfirst\033[0m - 生成节点启动的docker-compose yaml, 适用于知晓链初始配置所需参数的启动节点"
    println "      \033[0;32mmore\033[0m - 生成节点启动的docker-compose yaml, 适用于非首个节点的启动节点"
    println
    println "    Flags:"
    println "    Used with \033[0;32mnetwork.sh first\033[0m, \033[0;32mnetwork.sh more\033[0m:"
    println "    -tag <image tag> -  m0镜像tag"
    println "    -chain-id <chain name> - 链名称"
    println "    -genesis-time <timestamp> - 创世块时间戳"
    println "    -num <num> - 验证人个数"
    println "    -validator-ip <ips> - 验证人节点列表 host:port,host:port"
    println "    -validator-key <keys> -  验证人私钥的助记词列表"
    println "    -id <node moniker> - 节点名字"
    println "    -key <node key> - 节点私钥助记词"
    println "    -seeds <seeds> - 种子节点列表 ID@host:port,ID@host:port"
    println "    -seed - 节点启动种子模式"
    println "    -unsync - 启动同步程序，开启浏览器、钱包服务"
    println "    -automatic - 启动自动发送交易"
    println "    -output <directory> - 结果输出目录(默认当前目录)"
    println
    println "    -h - Print this message"
    println
    println " Examples:"
    println "   network.sh first --chain-id tttt --node-id n0"
    println "   network.sh more --chain-id tttt --node-id n0"
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

export -f errorln
export -f successln
export -f infoln
export -f warnln