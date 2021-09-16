# 快速入门

本节将指导您获取节点程序并部署一个可用的测试网络环境，还会展示一些基本操作。

## 获取节点程序

- 程序列表

  - m0d 节点程序，提供节点启动、命令行操作等功能。

    ```shell
    m0d --help
    ```

    ```
    m0 Daemon (server)
    
    Usage:
      m0d [command]
    
    Available Commands:
                  
                  
      automatic        automatic send tx
      ca-server        Manage your application's certificate
      debug            Tool for helping with debugging your application
      export           Export state to JSON
      help             Help about any command
      init             Initialize private validator, p2p, genesis, and application configuration files
      keys             Manage your application's keys
      query            Querying subcommands
      start            Run the full node
      status           Query remote node for status
      tendermint       Tendermint subcommands
      testnet          Initialize files for a m0 testnet
      tx               Transactions subcommands
      unsafe-reset-all Resets the blockchain database, removes address book files, and resets data/priv_validator_state.json to the genesis state
      validate-genesis validates the genesis file at the default location or at the location passed as an arg
      version          Print the application binary version information
    ```

    

  - synced 数据同步服务程序，提供同步节点数据、中心化私钥管理、查询数据服务。

    ```shell
    synced --help
    ```
  
    ```
    m0 synced Daemon (server)
    
    Usage:
      synced [command]
    
    Available Commands:
      help        Help about any command
      start       start synced service
      version     Print tendermint libraries' version
    ```
  
    
  
  - xdev wasm合约编译器，编译c++源码为wasm文件。
  
    ```shell
    xdev --help
    ```
  
    ```
    Usage:
      xdev [command]
    
    Available Commands:
      build       build command builds a project
      help        Help about any command
      init        init initializes a new project
      test        test perform unit test
      version     Print tendermint libraries' version
    ```
  
  - wasm2c wasm工具，部署wasm合约时节点会用到。

---

- 来自于源码

  - 前期准备
    主要由Golang开发，需要首先准备源码编译运行的环境
    - 安装go语言编译环境，版本为1.16或更高。[下载地址](https://golang.org/dl/)
    - 安装git。[下载地址](https://git-scm.com/download)
  - 下载源码
  **注意**: `master` 分支是日常开发分支，会包含最新的功能，但是 **不稳定**。生产环境请使用最新的已发布分支。
  
```shell
  # git下载源码
  git clone https://github.com/liubaninc/m0.git
```
  - 编译程序
  
    输出在./build文件夹里
  ```shell
  # 切换源码目录
  cd src/github.com/liubaninc/m0
  # 编译
  make build
  ```
  - 安装程序
  
    输出在$GOPATH/bin文件夹里
  ```shell
  # 切换源码目录
  cd src/github.com/liubaninc/m0
  # 安装
  make install
  ```
  - 编译镜像
  
    构建ciyuntangquan/m0:latest镜像
  ```shell
  # 切换源码目录
  cd src/github.com/liubaninc/m0
  # 构建镜像
  make m0-image
  ```

## 搭建可用的测试网络

- 本地搭建
  1. 创世块文件

     ```shell 
     # 生成验证节点配置文件，尤其是创世块文件
     m0d testnet -n 1 --chain-id testnet --dev-mode --output-dir ~/mytestnet
     ```

  2. 启动网络节点

     ```shell
     # 启动验证节点
     m0d start --pruning nothing --home ~/mytestnet/node0/.m0/
     ```

  3. 数据同步服务

     ```shell
     # 默认连接127.0.0.1:26657的节点rpc服务，并监听8080端口提供数据服务
     synced start
     ```
     访问http://127.0.0.1:8080/swagger/index.html，提供api文档

  4. 钱包

     ```shell
     cd m0/vue/browser
     # 默认连接127.0.0.1:8080的数据同步服务，并监听8086端口提供界面服务
     npm run serve
     ```
     访问http://127.0.0.1:8086，提供用户操作与查询区块链数据。

  5. 浏览器

     ```shell
     cd m0/vue/wallet
     # 默认连接127.0.0.1:8080的数据同步服务，并监听8088端口提供界面服务
     npm run serve
     ```

     访问http://127.0.0.1:8088，提供浏览与查询区块链数据。

  测试网络搭建完成，开启您的区块链之旅！

- docker-compose 容器搭建

  1. 编写 node0.yaml
  ```yaml
  # 启动区块链网络节点，适用于知晓链初始配置所需参数的启动节点，如区块链首个节点
  # 前置条件：
  # 1. 提供验证人列表信息，如助记词、IP、端口等
  # 2. 提供链相关信息， 如链ID、创世块时间等
  # 3. 提供节点相关信息, 如助记词、IP:端口
  # 4. 提供节点列表 ID@host:port,ID@host:port,ID@host:port
  # 步骤:
  # 1、生成链配置信息（根据前置条件中1、2）
  # 2、生成节点配置（根据前置条件中3）
  # 3、使用1结果
  # 4、启动节点
  # 结果:
  # 启动节点
  
  version: '2.4'
  
  volumes:
    data-testnet: # 链存储数据定义
  
  networks:
    default:
      name: network-testnet #链网络定义
  
  services:
    testnet-node0:
      container_name: testnet-node0
      image: ciyuntangquan/m0:latest
      # restart: always
      labels:
        service: testnet
      environment:
        # 链参数
        - M0D_NUM=1
        - M0D_CHAIN_ID=testnet
        - M0D_GENESIS_TIME=1615183366
        - M0D_NODE_IP=
        - M0D_NODE_VALIDATOR_KEY=witness relief differ hobby shrimp million math left plate hood submit cover evoke stem street still gas voice nothing reward pull window meadow ability
        - M0D_DEV_MODE=true
        # 节点参数
        - M0D_MONIKER=node0
        - NODE_MNEMONIC=witness relief differ hobby shrimp million math left plate hood submit cover evoke stem street still gas voice nothing reward pull window meadow ability
        - M0D_CONSENSUS_CREATE_EMPTY_BLOCKS=true
        - M0D_HOME=/var/tq/production/testnet-node0/.m0
        - M0D_LOG_LEVEL=info
        - M0D_INSTRUMENTATION_PROMETHEUS=true
        - M0D_P2P_PEX=true
        - M0D_P2P_SEED_MODE=false
        - M0D_P2P_SEEDS=
        - M0D_P2P_PERSISTENT_PEERS=
        - WAIT_HOSTS_TIMEOUT=300
        - WAIT_SLEEP_INTERVAL=30
        - WAIT_HOST_CONNECT_TIMEOUT=30
      command:
        - "sh"
        - "-c"
        - |
          set -ex
          if [ ! -e /var/tq/production/testnet-node0/.m0/config/genesis.json ]; then
            echo "witness relief differ hobby shrimp million math left plate hood submit cover evoke stem street still gas voice nothing reward pull window meadow ability" | m0d init node0 --recover --home /var/tq/production/testnet-node0/.m0
            m0d testnet
            cp -r mytestnet/node0/.m0/config/genesis.json /var/tq/production/testnet-node0/.m0/config
            cp -r mytestnet/ca /var/tq/production/ca
          fi
  
          if [ ! -d /var/tq/production/testnet-node0/.m0/ca ]; then
            ln -s /var/tq/production/ca /var/tq/production/testnet-node0/.m0/ca
          fi
          sed -i "s/addr_book_strict = true/addr_book_strict = false/g" /var/tq/production/testnet-node0/.m0/config/config.toml
          nohup m0d start --pruning nothing --p2p.laddr tcp://testnet-node0:26656 > m0.log 2>&1 &
          
          export WAIT_HOSTS=127.0.0.1:26657
          /wait
          if [ true == true ];then
            nginx -g "daemon off;" &
            nohup synced start  > synced.log 2>&1 &
          fi
          if [ false == true ];then
            m0d automatic > automatic.log 2>&1 &
          fi
          
          tail -f m0.log
      volumes:
        - data-testnet:/var/tq/production
        - /var/run/docker.sock:/var/run/docker.sock
      ports:
        #- 1317  # API(HTTP)
        #- 9090  # GRPC (TCP)
        - 26656 # P2P (TCP)
        - 26657 # RPC(HTTP)
        - 26660 # prometheus (HTTP)
        #- 6060 # pprof 
        - 8080 # synced
        - 8086 # wallet
        - 8088 # browser
      networks:
        - default
  ```

  2. 启动node0节点
  ```shell
  docker-compose -f node0.yaml up -d
  ```

   使用命令`docker ps`查看8080，8086，8088端口随机映射在本地端口，可访问数据同步服务、节点浏览器、钱包等程序
  
   
  
  测试网络搭建完成，开启您的区块链之旅！

## 基本操作(命令行方式)

- 节点查询

```shell
# 节点状态
m0d status
# 节点NODEID
m0d tendermint show-node-id
```

- 区块查询

```shell
# 根据高度查询区块，不指定高度为最新区块
m0d query block 5
```

- 交易查询

```shell
# 根据哈希查询交易
m0d query tx 54086047504F5A7536EEC241ADDDE09377CD1882ECE70EA6C6D21CA326F2D85A
# 根据事件查询交易列表
m0d query txs --events message.module=utxo
```

- 账户私钥管理

```shell
# 账户列表
m0d keys list
# 新增账户: 导入助记词。账户名 alice
echo "key erupt service six thing spy noise heart giggle year oil fuel rival drop goat deal moral require knee pact bind brain word nuclear" | m0d keys add alice --recover
# 新增账户: 随机组记词。账户名 bob
m0d keys add bob
```

- 资产管理
  - 资产查询

  ```shell
  # 资产列表
  m0d query utxo list-token
  # 资产详情: m0token
  m0d query utxo show-token m0token
  ```

  - 查询账户余额

  ```shell
  # alice地址的余额
  m0d query bank balances $(m0d keys show alice --address)
  # alice地址的资产m0token utxo列表
m0d query utxo list-input $(m0d keys show alice --address) m0token
  # alice地址的资产10m0token utxo列表
m0d query utxo show-input $(m0d keys show alice --address) m0token 10m0token
  ```
  
  - 发行资产/增发资产
  
```shell
  # alice 发行资产 1000m1token 给 alice
m0d tx utxo issue $(m0d keys show alice --address) 1000m1token --from alice --chain-id=testnet --broadcast-mode block -y
  ```
  
  - 转账资产
  
```shell
  # alice 转账资产 10m0token 给 bob
m0d tx utxo send $(m0d keys show bob -a) 10m0token --from alice --chain-id=testnet --broadcast-mode block -y
  ```
  
  - 销毁资产
  
  ```shell
  # alice 销毁资产 10m1token
  m0d tx utxo destroy 10m1token --from alice --chain-id=testnet --broadcast-mode block -y
  ```

- wasm合约

  - 部署合约

  ```shell
  # 部署合约 counter00
  m0d tx wasm deploy counter00 wasm/counter.wasm "{\"creator\": \"alice\"}" --from alice --broadcast-mode block --chain-id testnet --gas auto --gas-adjustment 1.5 -y
  ```

  - 升级合约

  ```shell
  # 升级合约 counter00
  m0d tx wasm upgrade counter00 wasm/counter.wasm --from alice --broadcast-mode block --chain-id testnet --gas auto --gas-adjustment 1.5 -y
  ```

  - 调用合约

  ```shell
  # 调用合约 counter00
  m0d tx wasm invoke counter00 increase "{\"key\": \"a\"}" --from alice --broadcast-mode block --chain-id testnet --gas auto --gas-adjustment 1.5 -y
  ```

  - 查询合约

  ```shell
  # 查询合约 counter00
  m0d query wasm query counter00 get "{\"key\": \"a\"}"
  ```

  