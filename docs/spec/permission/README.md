# 角色权限管理

## 简介

## 执行命令 `m0d tx permission`

```shell
m0d tx permission --help
```

```
permission transactions subcommands

Usage:
  m0d tx permission [flags]
  m0d tx permission [command]

Available Commands:
  set-permission modify the permissions of account
```

- 设置角色权限列表

```shell
m0d tx permission set-permission --help
```

```
modify the permissions of account

Usage:
  m0d tx permission set-permission [address] [perm] [[perm]] [flags]
```

##  查询命令 `m0d query permission`

```shell
m0d query permission --help
```

```
Querying commands for the permission module

Usage:
  m0d query permission [flags]
  m0d query permission [command]

Available Commands:
  list-account list all account
  show-account shows a account
```

- 查询所有账户角色权限列表

```shell
m0d query permission list-account --help
```

```
list all account

Usage:
  m0d query permission list-account [flags]
```

- 查询指定账户角色权限列表

```shell
m0d query permission show-account --help
```

```
shows a account

Usage:
  m0d query permission show-account [address] [flags]
```