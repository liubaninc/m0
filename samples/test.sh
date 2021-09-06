#!/bin/bash
# 构建区块链网络拓扑案例, 生成各进程的docker-compose yaml文件
c_num=${1:-1} # CA服务个数 默认1个
s_num=${2:-1} # 种子节点个数 默认1个
v_num=${3:-4} # 验证节点个数 默认4个
n_num=${4:-2} # 数据节点个数 默认2个
chain_id=${5:-testnetA} # 链名称 默认为testnetA
genesis_time=${6:-1615183366} # 链启动时间 默认为1615183366
nids=""
# 生成初始验证节点助记词列表
validators=""
array=""
for((i=1;i<="$v_num";i++));
do
  key=$(m0d keys mnemonic 2>&1)  # 节点助记词
  array[$i-1]=$key
  if [ $i == 1 ];then
    validators="$key"
  else
    validators="$validators,$key"
  fi
done
# 生成ca服务yaml文件
for((i=1;i<="$c_num";i++));
do
  $(dirname $0)/testnet.sh ca -chain-id "$chain_id" -id "C$i"
done
# 生成种子节点yaml文件
seeds="" # 种子节点NodeID列表
for((i=1;i<="$s_num";i++));
do
  key=$(m0d keys mnemonic 2>&1)  # 节点助记词
  $(dirname $0)/testnet.sh first -chain-id "$chain_id" -genesis-time "$genesis_time" -num "$v_num" -validator-key "$validators" -id "S$i" -key "$key" -seeds "$seeds" -seed -dev
  nid=$(echo "$key"	| m0d init ss --overwrite --recover 2>&1 | jq .node_id | sed 's/"//g') #节点NodeID
  nids="$nids $nid"
  if [ $i == 1 ];then
    seeds="$nid@${chain_id}-S$i:26656"
  else
    seeds="$seeds,$nid@${chain_id}-S$i:26656"
  fi
done
# 生成验证节点yaml文件
peers="" # 节点NodeID列表
i=1
for key in "${array[@]}"
do
  $(dirname $0)/testnet.sh more -chain-id "$chain_id" -id "V$i" -key "$key" -seeds "$seeds" -peers "$peers" -node "${chain_id}"-S1:26657
  nid=$(echo "$key"	| m0d init ss --overwrite --recover 2>&1 | jq .node_id | sed 's/"//g') #节点NodeID
  if [ $i == 1 ];then
    peers="$nid@${chain_id}-V$i:26656"
  else
    peers="$peers,$nid@${chain_id}-V$i:26656"
  fi
  let i++
done
# 生成数据节点yaml文件
for((i=1;i<="$n_num";i++));
do
  key=$(m0d keys mnemonic 2>&1)  # 节点助记词
  nid=$(echo "$key"	| m0d init ss --overwrite --recover 2>&1 | jq .node_id | sed 's/"//g')
  nids="$nids $nid"
  $(dirname $0)/testnet.sh more -chain-id "$chain_id" -id "N$i" -key "$key" -seeds "$seeds" -node "${chain_id}"-S1:26657 -sync
done
echo ""
echo "*************备份信息********************"
echo "链信息: "
echo "名称: $chain_id"
echo "启动时间: $genesis_time"
echo "验证人助记词: $validators"
echo "种子列表: $seeds"
echo "节点列表: $peers"
echo "数据节点": $nids
echo "***************************************"
echo "$seeds" > ./${chain_id}/seeds
echo "$peers" > ./${chain_id}/peers
echo "$validators" > ./${chain_id}/validators
