# 命令行操作

**M0** 模块化设计，按照业务划分不同模块。每个模块有自己的消息结构及相应的执行命令、数据状态查询命令。

`m0d tx 模块名 消息类型 消息参数列表` 执行命令，提交消息到节点

```shell
m0d tx --help
```

```
Transactions subcommands

Usage:
  m0d tx [flags]
  m0d tx [command]

Available Commands:
              
              
  broadcast           Broadcast transactions generated offline
  decode              Decode an binary encoded transaction string.
  encode              Encode transactions generated offline
  ibc                 IBC transaction subcommands
  mibc                mibc transactions subcommands
  multisign           Generate multisig signatures for transactions generated offline
  peer                peer transactions subcommands
  permission          permission transactions subcommands
  pki                 pki transactions subcommands
  sign                Sign a transaction generated offline
  sign-batch          Sign transaction batch files
  storage             storage transactions subcommands
  utxo                utxo transactions subcommands
  validate-signatures Validate transactions signatures
  validator           validator transactions subcommands
  wasm                wasm transactions subcommands
```

`m0d query 模块名 消息类型 消息参数列表` 状态查询，查询节点中数据状态

```shell
m0d query --help
```

```
Querying subcommands

Usage:
  m0d query [flags]
  m0d query [command]

Aliases:
  query, q

Available Commands:
  auth                     Querying commands for the auth module
  bank                     Querying commands for the bank module
  block                    Get verified data for a the block at given height
  ibc                      Querying commands for the IBC module
  mibc                     Querying commands for the mibc module
  params                   Querying commands for the params module
  peer                     Querying commands for the peer module
  permission               Querying commands for the permission module
  pki                      Querying commands for the pki module
  storage                  Querying commands for the storage module
  tendermint-validator-set Get the full tendermint validator set at given height
  tx                       Query for a transaction by hash in a committed block
  txs                      Query for paginated transactions that match a set of events
  utxo                     Querying commands for the utxo module
  validator                Querying commands for the validator module
  wasm                     Querying commands for the wasm module
```

支持的模块

- [utxo](../spec/utxo/README.md) 资产管理, 发行/增发、转移、销毁资产
- [wasm](../spec/wasm/README.md) 合约管理，部署、升级、调用wasm合约 
- [validator](../spec/validator/README.md) 验证人管理，新增、删除验证人
- [permission](../spec/permission/README.md) 角色权限管理，设置账户权限角色列表
- [peer](../spec/peer/README.md) 节点管理， 新增、删除准入节点
- [pki](../spec/pki/README.md) 公钥基础设施管理， 上传、吊销、冻结证书

