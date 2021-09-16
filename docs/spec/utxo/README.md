# 资产管理

## 简介

## 执行命令 `m0d tx utxo`

```shell
m0d tx utxo --help
```

```
utxo transactions subcommands

Usage:
  m0d tx utxo [flags]
  m0d tx utxo [command]

Available Commands:
  destroy     destroy an utxo token (denom)
  issue       issue an utxo token (denom)
  send        send an utxo token (denom)
```



- 资产发行/资产增发

```shell
m0d tx utxo issue --help
```

```
issue an utxo token (denom)

Usage:
  m0d tx utxo issue [to_address] [amount] [[to_address] [amount]] [flags]
```

- 资产转移

```shell
m0d tx utxo send --help
```

```
send an utxo token (denom)

Usage:
  m0d tx utxo send [to_address] [amount] [[to_address] [amount]] [flags]
```

- 资产销毁

```shell
m0d tx utxo destroy --help
```

```
destroy an utxo token (denom)

Usage:
  m0d tx utxo destroy [amount][,[amount]] [flags]
```

## 查询命令 `m0d query utxo`

```shell
m0d query utxo --help
```

```
Querying commands for the utxo module

Usage:
  m0d query utxo [flags]
  m0d query utxo [command]

Available Commands:
  list-input  query all available unspent outputs list of a specific address and denom.
  list-token  list all token
  show-input  query for the available unspent outputs of a specific address and specific amount of coins
  show-token  shows a token
```

- 查询资产列表

```shell
m0d query utxo list-token --help
```

```
list all token

Usage:
  m0d query utxo list-token [flags]
```

- 查询资产详情

```shell
m0d query utxo show-token --help
```

```
shows a token

Usage:
  m0d query utxo show-token [name] [flags]
```

- 查询指定地址、指定资产所有utxo列表

```shell
m0d query utxo list-input --help
```

```
query all available unspent outputs list of a specific address and denom.

Usage:
  m0d query utxo list-input [address] [denom] [flags]
```

- 查询指定地址、指定资产、指定金额utxo列表

```shell
m0d query utxo show-input --help
```

```
query for the available unspent outputs of a specific address and specific amount of coins

Usage:
  m0d query utxo show-input [address] [amount][,[amount]] [flags]
```