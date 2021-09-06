# 快速入门

本章节将指导您获取M0的代码并部署一个基础的可用环境，还会展示一些基本操作

## 环境部署

### 准备环境
M0主要由Golang开发，需要首先准备编译运行的环境
- 安装go语言编译环境，版本为1.16或更高。[下载](https://golang.org/dl/)
- 安装git。[下载](https://git-scm.com/download)

### 编译

- 使用git下载源码到本地
```shell script
git clone https://github.com/liubaninc/m0.git
```

- 源码编译
```shell script
cd src/github.com/liubaninc/m0
make
```
build目录生成可执行文件

- 源码安装
```shell script
cd src/github.com/liubaninc/m0
make
```
$GOPATH/bin目录生成可执行文件

- 编译镜像
```shell script
cd src/github.com/liubaninc/m0
make m0-image
```

## 搭建网络

使用docker-compose工具部署测试网络
> 本地容器化运行的示例，实际场景中可以用卷的方式挂载并覆盖配置。

## 基本操作

### 账户管理

### 发送交易

### 数据查询