# 验证人管理

## 简介

## 执行命令 `m0d tx validator`

```shell
m0d tx validator --help
```

```
validator transactions subcommands

Usage:
  m0d tx validator [flags]
  m0d tx validator [command]

Available Commands:
  create-validator Adds a new validator node
  edit-validator   edit an existing validator node
  leave-validator  leave an existing validator node
```

- 新增验证人

```shell
m0d tx validator create-validator --help
```

```
Adds a new validator node

Usage:
  m0d tx validator create-validator [validator-pubkey] [flags]
```

- 修改验证人信息

```shell
m0d tx validator edit-validator --help
```

```
edit an existing validator node

Usage:
  m0d tx validator edit-validator [flags]
```

- 删除验证人

```shell
m0d tx validator leave-validator --help
```

```
leave an existing validator node

Usage:
  m0d tx validator leave-validator [flags]
```

## 查询命令 `m0d query validator`

```shell
m0d query validator --help
```

```
Querying commands for the validator module

Usage:
  m0d query validator [flags]
  m0d query validator [command]

Available Commands:
  list-validator list all validator
  show-validator shows a validator
```

- 查询验证人列表

```shell
m0d query validator list-validator --help
```

```
list all validator

Usage:
  m0d query validator list-validator [flags]
```

- 查询验证人详情

```shell
m0d query validator show-validator --help
```

```
shows a validator

Usage:
  m0d query validator show-validator [validator-address] [flags]
```

