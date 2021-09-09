# wasm合约管理

## 简介

## 执行命令 `m0d tx wasm`

```shell
m0d tx wasm --help
```

```
wasm transactions subcommands

Usage:
  m0d tx wasm [flags]
  m0d tx wasm [command]

Available Commands:
  approve-deploy-contract approve a deploy of contract
  deploy                  deploy an wasm contract
  freeze                  freeze an wasm contract
  invoke                  invoke an wasm contract's method
  propose-deploy-contract propose a deploy of contract
  undeploy                kill an wasm contract
  unfreeze                unfreeze an wasm contract
  upgrade                 upgrade an wasm contract
```

- 部署合约多方确认提议

```shell
m0d tx wasm propose-deploy-contract --help
```

```
propose a deploy of contract

Usage:
  m0d tx wasm propose-deploy-contract [name] [code-file] [init-args] [approval] [[approval]] [flags]
```

- 部署合约多方确认审批

```shell
m0d tx wasm approve-deploy-contract --help
```

```
approve a deploy of contract

Usage:
  m0d tx wasm approve-deploy-contract [name] [flags]
```

- 部署合约

```shell
m0d tx wasm deploy --help
```

```
deploy an wasm contract

Usage:
  m0d tx wasm deploy [name] [code-file] [init-args] [flags]
```

- 调用合约

```shell
m0d tx wasm invoke --help
```

```
invoke an wasm contract's method

Usage:
  m0d tx wasm invoke [name] [method] [args] [flags]
```

- 升级合约

```shell
m0d tx wasm upgrade --help
```

```
upgrade an wasm contract

Usage:
  m0d tx wasm upgrade [name] [code-file] [flags]
```

- 冻结合约

```shell
m0d tx wasm freeze --help
```

```
freeze an wasm contract

Usage:
  m0d tx wasm freeze [name] [flags]
```

- 解冻合约

```shell
m0d tx wasm unfreeze --help
```

```
unfreeze an wasm contract

Usage:
  m0d tx wasm unfreeze [name] [flags]
```

- 销毁合约

```shell
m0d tx wasm undeploy --help
```

```
kill an wasm contract

Usage:
  m0d tx wasm undeploy [name] [flags]
```

## 查询命令 `m0d query wasm`

```shell
m0d query wasm --help
```

```
Querying commands for the wasm module

Usage:
  m0d query wasm [flags]
  m0d query wasm [command]

Available Commands:
  list-approve-deploy list all approveDeploy
  list-contract       list all contract
  list-propose-deploy list all proposeDeploy
  query               query contract method
  show-approve-deploy shows a approveDeploy
  show-contract       shows a contract
  show-propose-deploy shows a proposeDeploy
```

- 查询合约部署多方确认提议列表

```shell
m0d query wasm list-propose-deploy --help
```

```
list all proposeDeploy

Usage:
  m0d query wasm list-propose-deploy [flags]
```

- 查询合约部署多方确认提议详情

```shell
m0d query wasm show-propose-deploy --help
```

```
shows a proposeDeploy

Usage:
  m0d query wasm show-propose-deploy [name] [flags]
```

- 查询合约部署多方确认审批列表

```shell
m0d query wasm list-approve-deploy --help
```

```
list all approveDeploy

Usage:
  m0d query wasm list-approve-deploy [flags]
```

- 查询合约部署多方确认审批详情

```shell
m0d query wasm show-approve-deploy --help
```

```
shows a approveDeploy

Usage:
  m0d query wasm show-approve-deploy [name] [flags]
```

- 查询合约列表

```shell
m0d query wasm list-contract --help
```

```
list all contract

Usage:
  m0d query wasm list-contract [[address]] [flags]
```

- 查询合约详情

```shell
m0d query wasm show-contract --help
```

```
shows a contract

Usage:
  m0d query wasm show-contract [name] [flags]
```

- 合约查询

```shell
m0d query wasm query --help
```

```
query contract method

Usage:
  m0d query wasm query [contract] [method] [args] [flags]
```