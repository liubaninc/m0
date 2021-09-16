# 搭建测试网络

  节点加入网络需获取加入区块链的创世块文件genesis.json。创世块文件定义了该链初始数据，不同区块链的创世块文件是不一样的。

  搭建测试网络，首先定义创世块文件，如链启动时默认数据。比如链名称、链启动时间、验证节点列表、创世账户、创世资产等。然后启动节点(验证节点、数据节点)组建区块链网络。

## 区块链链定义

  **M0**提供了命令`m0d testnet`可以定义区块链网络

  ```shell
  m0d testnet --help
  ```

  ```
  testnet will create "n" number of directories and populate each with
necessary files (private validator, genesis, config, etc.).
  Note, strict routability for addresses is turned off in the config file.
  Example:
	m0d testnet --v 4 --output-dir ./output --starting-ip-address 192.168.10.2
  
  Usage:
    m0d testnet [flags]
  
  Flags:
        --algo string                        Key signing algorithm to generate keys for (default "secp256k1")
        --chain-id string                    genesis file chain-id, if left blank will be randomly created
        --dev-mode                           develop mode, skip some module's genesis state， as peer、permission
        --genesis-time int                   override genesis UNIX time instead of using a random UNIX time
    -h, --help                               help for testnet
        --keyring-backend string             Select keyring's backend (os|file|test) (default "test")
        --minimum-gas-prices string          Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01photino,0.001stake)
        --node-daemon-home string            Home directory of the node's daemon configuration (default ".m0")
        --node-dir-prefix string             Prefix the directory name for each node with (node results in node0, node1, ...) (default "node")
        --node-ip string                     Comma-delimited node ip address, if left blank will be randomly created 127.0.0.1:26656
        --node-validator-key string          Comma-delimited node validator key,  if left blank will be randomly created
    -n, --num int                            Number of validators to initialize the testnet with (default 4)
    -o, --output-dir string                  Directory to store initialization data for the testnet (default "./mytestnet")
        --reserved-account-mnemonic string   Reserved account mnemonic (default "key erupt service six thing spy noise heart giggle year oil fuel rival drop goat deal moral require knee pact bind brain word nuclear")
      --reserved-coin string               Reserved coin (default "100000000000m0token")
        --starting-ip-address string         Starting IP address (192.168.0.1 results in persistent peers list ID0@192.168.0.1:46656, ID1@192.168.0.2:46656, ...) (default "192.168.0.1")
  ```

  常用参数解释：
  - chain-id 区块链名称，默认随机生成。
  - genesis-time 区块链启动时间， 默认当前时间的时间戳
  - num 验证人个数，默认4个节点
  - node-validator-key 验证人节点私钥组记词， 默认随机生成
  - reserved-account-mnemonic 预留账户私钥组记词，默认账户地址 mc19dzfuxxv8vjeajjq475ahgrl0meudwexdmrnye
  - reserved-coin 预留资产， 默认100000000000m0token
  - dev-mode 开发者模式， 默认关闭。 开发者模式会关闭用户权限、多方合约部署、节点准入等机制

## 节点启动

  区块链链定义后，需要启动节点加入其网络，尤其是验证节点启动。验证节点是参与区块链共识的节点，负责打包区块。如果验证节点个数不足，共识无法完成，区块高度会停滞，交易不能进去区块。

  1. 生成节点配置文件 `m0d init` 

     ```shell
     m0d init --help
     ```

     ```
     Initialize validators's and node's configuration files.
     
     Usage:
       m0d init [moniker] [flags]
     
     Flags:
           --chain-id string   genesis file chain-id, if left blank will be randomly created
       -h, --help              help for init
       -o, --overwrite         overwrite the genesis.json file
           --recover           provide seed phrase to recover existing key instead of creating
     ```

     常用参数
     
     - recover 恢复节点私钥组记词。验证节点必须是恢复的，验证节点公钥已上链成为链验证人。
     
  2. 替换genesis.json

  3. 启动节点 `m0d start`

     ```shell
     m0d start --help
     ```

     ```
     Run the full node application with Tendermint in or out of process. By
     default, the application will run with Tendermint in process.
     
     Pruning options can be provided via the '--pruning' flag or alternatively with '--pruning-keep-recent',
     'pruning-keep-every', and 'pruning-interval' together.
     
     For '--pruning' the options are as follows:
     
     default: the last 100 states are kept in addition to every 500th state; pruning at 10 block intervals
     nothing: all historic states will be saved, nothing will be deleted (i.e. archiving node)
     everything: all saved states will be deleted, storing only the current state; pruning at 10 block intervals
     custom: allow pruning options to be manually specified through 'pruning-keep-recent', 'pruning-keep-every', and 'pruning-interval'
     
     Node halting configurations exist in the form of two flags: '--halt-height' and '--halt-time'. During
     the ABCI Commit phase, the node will check if the current block height is greater than or equal to
     the halt-height or if the current block time is greater than or equal to the halt-time. If so, the
     node will attempt to gracefully shutdown and the block will not be committed. In addition, the node
     will not be able to commit subsequent blocks.
     
     For profiling and benchmarking purposes, CPU profiling can be enabled via the '--cpu-profile' flag
     which accepts a path for the resulting pprof file.
     
     Usage:
       m0d start [flags]
     
     Flags:
           --abci string                                     specify abci transport (socket | grpc) (default "socket")
           --address string                                  Listen address (default "tcp://0.0.0.0:26658")
           --consensus.create_empty_blocks                   set this to false to only produce blocks when there are txs or when the AppHash changes (default true)
           --consensus.create_empty_blocks_interval string   the possible interval between empty blocks (default "0s")
           --consensus.double_sign_check_height int          how many blocks to look back to check existence of the node's consensus votes before joining consensus
           --cpu-profile string                              Enable CPU profiling and write to the provided file
           --db_backend string                               database backend: goleveldb | cleveldb | boltdb | rocksdb | badgerdb (default "goleveldb")
           --db_dir string                                   database directory (default "data")
           --fast_sync                                       fast blockchain syncing (default true)
           --genesis_hash bytesHex                           optional SHA-256 hash of the genesis file
           --grpc.address string                             the gRPC server address to listen on (default "0.0.0.0:9090")
           --grpc.enable                                     Define if the gRPC server should be enabled (default true)
           --halt-height uint                                Block height at which to gracefully halt the chain and shutdown the node
           --halt-time uint                                  Minimum block time (in Unix seconds) at which to gracefully halt the chain and shutdown the node
       -h, --help                                            help for start
           --inter-block-cache                               Enable inter-block caching (default true)
           --inv-check-period uint                           Assert registered invariants every N blocks
           --min-retain-blocks uint                          Minimum block height offset during ABCI commit to prune Tendermint blocks
           --minimum-gas-prices string                       Minimum gas prices to accept for transactions; Any fee in a tx must meet this minimum (e.g. 0.01photino;0.0001stake)
           --moniker string                                  node name (default "AppledeMacBook-Pro-3.local")
           --p2p.laddr string                                node listen address. (0.0.0.0:0 means any interface, any port) (default "tcp://0.0.0.0:26656")
           --p2p.persistent_peers string                     comma-delimited ID@host:port persistent peers
           --p2p.pex                                         enable/disable Peer-Exchange (default true)
           --p2p.private_peer_ids string                     comma-delimited private peer IDs
           --p2p.seed_mode                                   enable/disable seed mode
           --p2p.seeds string                                comma-delimited ID@host:port seed nodes
           --p2p.unconditional_peer_ids string               comma-delimited IDs of unconditional peers
           --p2p.upnp                                        enable/disable UPNP port forwarding
           --priv_validator_laddr string                     socket address to listen on for connections from external priv_validator process
           --proxy_app string                                proxy app address, or one of: 'kvstore', 'persistent_kvstore', 'counter', 'counter_serial' or 'noop' for local testing. (default "tcp://127.0.0.1:26658")
           --pruning string                                  Pruning strategy (default|nothing|everything|custom) (default "default")
           --pruning-interval uint                           Height interval at which pruned heights are removed from disk (ignored if pruning is not 'custom')
           --pruning-keep-every uint                         Offset heights to keep on disk after 'keep-every' (ignored if pruning is not 'custom')
           --pruning-keep-recent uint                        Number of recent heights to keep on disk (ignored if pruning is not 'custom')
           --rpc.grpc_laddr string                           GRPC listen address (BroadcastTx only). Port required
           --rpc.laddr string                                RPC listen address. Port required (default "tcp://127.0.0.1:26657")
           --rpc.pprof_laddr string                          pprof listen address (https://golang.org/pkg/net/http/pprof)
           --rpc.unsafe                                      enabled unsafe rpc methods
           --state-sync.snapshot-interval uint               State sync snapshot interval
           --state-sync.snapshot-keep-recent uint32          State sync snapshot to keep (default 2)
           --trace-store string                              Enable KVStore tracing to an output file
           --transport string                                Transport protocol: socket, grpc (default "socket")
           --unsafe-skip-upgrades ints                       Skip a set of upgrade heights to continue the old binary
           --with-tendermint                                 Run abci app embedded in-process with tendermint (default true)
           --x-crisis-skip-assert-invariants                 Skip x/crisis invariants check on startup
     ```

 

 ## 案例

1、链定义
```
# 链定义, 输出在./mytestnet目录
# 创世块文件 node*/.m0/config/genesis.json
# 验证节点私钥助记词 node*/.m0/key_seed.json
m0d testnet -n 4 --chain-id testnetA
```
2、启动节点

不同的节点在不同物理机启动, 假设物理机IP分别为192.168.6.1、192.168.6.2、192.168.6.3、192.168.6.4

- 节点ID列表

```shell
peers=
# node0 
key=$(cat mytestnet/node0/.m0/key_seed.json | jq .secret | sed 's/\"//g')
nid=$(echo "$key"	| m0d init ss --overwrite --recover 2>&1 | jq .node_id | sed 's/"//g') 
peers=$nid@192.168.6.1:26656

# node1
key=$(cat mytestnet/node1/.m0/key_seed.json | jq .secret | sed 's/\"//g')
nid=$(echo "$key"	| m0d init ss --overwrite --recover 2>&1 | jq .node_id | sed 's/"//g') 
peers=$peers,$nid@192.168.6.2:26656
```

- 启动节点 node0在192.168.6.1 

```shell
key=$(cat mytestnet/node0/.m0/key_seed.json | jq .secret | sed 's/\"//g')
# 节点配置
echo $key | m0d init V1 --recover
# 创世块覆盖
cp mytestnet/node0/.m0/config/genesis.json  ~/.m0/config/genesis.json
# 节点启动
m0d start --pruning nothing --p2p.persistent_peers $peers
```

- 启动节点 node1 在192.168.6.2
```shell
key=$(cat mytestnet/node1/.m0/key_seed.json | jq .secret | sed 's/\"//g')
# 节点配置
echo $key | m0d init V1 --recover
# 创世块覆盖
cp mytestnet/node0/.m0/config/genesis.json  ~/.m0/config/genesis.json
# 节点启动
m0d start --pruning nothing --p2p.persistent_peers $peers
```

- 启动节点 node2 在192.168.6.3
```shell
key=$(cat mytestnet/node2/.m0/key_seed.json | jq .secret | sed 's/\"//g')
# 节点配置
echo $key | m0d init V1 --recover
# 创世块覆盖
cp mytestnet/node0/.m0/config/genesis.json  ~/.m0/config/genesis.json
# 节点启动
m0d start --pruning nothing --p2p.persistent_peers $peers
```

- 启动节点 node3 在192.168.6.4
```shell
key=$(cat mytestnet/node3/.m0/key_seed.json | jq .secret | sed 's/\"//g')
# 节点配置
echo $key | m0d init V1 --recover
# 创世块覆盖
cp mytestnet/node0/.m0/config/genesis.json  ~/.m0/config/genesis.json
# 节点启动
m0d start --pruning nothing --p2p.persistent_peers $peers
```