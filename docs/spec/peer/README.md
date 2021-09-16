# 节点管理

## 简介

## 执行命令 `m0d tx peer`

```shell
m0d tx peer --help
```

```
peer transactions subcommands

Usage:
  m0d tx peer [flags]
  m0d tx peer [command]

Available Commands:
  create-peer-id Create a new peerID
  delete-peer-id Delete a peerID
  update-peer-id Update a peerID
```

- 准入节点

```shell
m0d tx peer create-peer-id --help
```

```
Create a new peerID

Usage:
  m0d tx peer create-peer-id [nodeID] [certIssuer] [certSerialNum] [flags]
```

- 准出节点

```shell
m0d tx peer delete-peer-id --help
```

```
Delete a peerID

Usage:
  m0d tx peer delete-peer-id [nodeID] [flags]
```

- 更新节点证书

```shell
m0d tx peer update-peer-id --help
```

```
Update a peerID

Usage:
  m0d tx peer update-peer-id [nodeID] [certIssuer] [certSerialNum] [flags]
```

##  查询命令 `m0d query peer`

```shell
m0d query peer --help
```

```
Querying commands for the peer module

Usage:
  m0d query peer [flags]
  m0d query peer [command]

Available Commands:
  list-peer-id list all peerID
  show-peer-id shows a peerID
```

- 查询节点列表

```shell
m0d query peer list-peer-id --help
```

```
list all peerID

Usage:
  m0d query peer list-peer-id [flags]
```

- 查询节点详情

```shell
m0d query peer show-peer-id --help
```

```
shows a peerID

Usage:
  m0d query peer show-peer-id [index] [flags]
```

