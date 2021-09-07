## M0是什么?

**M0**是北京磁云唐泉金服开发的可扩展、高性能的区块链框架, 是构建超级联盟网络的底层方案。您可以在面向各类业务场景的区块链需求中使用M0，作为区块链基础设施构建合规的区块链网络。

核心特点

* **高性能**
    * 实现了智能合约的并发执行和验证。
    * POA+BFT算法确保快速共识且绝对一致性。
    * 支持WASM虚拟机

* **更安全**
    * 多私钥保护的账户体系。
    * 支持SM2国密账户

* **多语言开发智能合约**
    * 支持多种语言的虚拟机
    
* **高灵活性**
    * 模块化设计使得用户可以方便选择适合自己业务场景的解决方案。
    * 提供了全面的、高性能的各业务场景的模块化解决方案。

## 快速使用

如何获取M0的代码并部署一个基础的可用环境

### 环境配置
* 操作系统：支持Linux以及Mac OS
* 开发语言：Go 1.16.x及以上
* 编译器：GCC 4.8.x及以上
* 版本控制工具：Git

### 构建源码
**注意**: `master` 分支是日常开发分支，会包含最新的功能，但是 **不稳定**。生产环境请使用最新的已发布分支。

```
# 克隆仓库
git clone https://github.com/liubaninc/m0.git
# 编译
cd m0
make
```

在build目录生成可执行文件
- m0d 节点可执行文件，节点启动及命令行程序。
   ![m0d](./docs/m0d.jpg)
   
- synced 数据服务可执行文件，同步节点数据且向业务提供数据服务程序。
   ![synced](./docs/synced.jpg)
   
- xdev wasm合约编译器可执行文件，编译c++源码为wasm文件。
   ![xdev](./docs/xdev.jpg)
   
- wasm2c wasm工具，部署wasm合约时会用到。

### 构建源码并安装

**注意**: `master` 分支是日常开发分支，会包含最新的功能，但是 **不稳定**。生产环境请使用最新的已发布分支。

```
# 克隆仓库
git clone https://github.com/liubaninc/m0.git
# 编译
cd m0
make install
```

在$GOPATH/bin目录生成可执行文件

### 构建镜像

**注意**: `master` 分支是日常开发分支，会包含最新的功能，但是 **不稳定**。生产环境请使用最新的已发布分支。

```
# 克隆仓库
git clone https://github.com/liubaninc/m0.git
# 编译
cd m0
make m0-image
```

### 部署测试网络

#### 启动测试网络
```shell
# 节点配置文件的生成，尤其是创世块文件的生成
m0d testnet -n 1 --output-dir ~/mytestnet

# 启动节点
m0d start --home ~/mytestnet/node0/.m0/
```

单验证节点测试网络搭建完成，开启您的区块链之旅！


#### 部署节点数据同步服务
```shell script
# 默认连接127.0.0.1:26657的节点rpc服务，并监听8080端口提供数据服务
synced start
```

#### 部署节点浏览器界面
```shell script
# 默认连接127.0.0.1:8080的数据同步服务，并监听8088端口提供界面服务
cd m0/vue/browser
npm run serve
```
提供用户浏览与查询区块链所有信息的工具
#### 部署节点钱包界面
```shell script
# 默认连接127.0.0.1:8080的数据同步服务，并监听8086端口提供界面服务
cd m0/vue/wallet
npm run serve
```
提供用户操作与查询区块链所有信息的工具

## 详细文档

关于M0更详细、更深入的了解，请查看[M0文档](./docs/readme.md)

## 如何参与开发
1. 阅读源代码，了解我们当前的开发方向
2. 找到自己感兴趣的功能或模块
3. 进行开发，开发完成后自测功能是否正确，并运行make & make test
4. 发起 [issue](.gitee/ISSUE_TEMPLATE.zh-CN.md)
5. 提交 [pull request](.gitee/PULL_REQUEST_TEMPLATE.zh-CN.md)

## 许可证
**M0**使用的许可证是Apache 2.0

## 联系我们
商务合作，请联系。