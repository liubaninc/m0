## m0是什么?

**m0**是北京磁云唐泉公司开发的多业务模块的区块链项目, 同时支持WASM智能合约，是构建超级联盟网络的底层方案。

核心特点

* **高性能**
    * 实现了智能合约的并发执行和验证。

* **更安全**
    * 多私钥保护的账户体系。

* **多语言开发智能合约**
    * 支持多语言开发智能合约。

* **高灵活性**
    * 插件化的设计使得用户可以方便选择适合自己业务场景的解决方案。

### 环境配置

* 操作系统：支持Linux以及Mac OS
* 开发语言：Go 1.12.x及以上
* 编译器：GCC 4.8.x及以上
* 版本控制工具：Git

### 程序构建

**注意**: `master` 分支是日常开发分支，会包含最新的功能，但是 **不稳定**。生产环境请使用最新的已发布分支。

```
# 下载代码
git clone --recursive https://gitee.com/tq_bc/m0.git
# 编译
make
```

build目录下程序介绍
- m0d
- synced
- xdev
- testnet
- wasm2c


### 快速搭建区块链网络
#### [进程方式运行](samples/binary.md)
#### [容器方式运行](samples/docker.md)

### 快速入门

## 详细文档

关于M0更详细、更深入的使用方法链接请查看[wiki](https://gitee.com/tq_bc/m0/wikis)

## 如何参与开发
1. 阅读源代码，了解我们当前的开发方向
2. 找到自己感兴趣的功能或模块
3. 进行开发，开发完成后自测功能是否正确，并运行make & make test
4. 发起 [issue](.gitee/ISSUE_TEMPLATE.zh-CN.md)
5. 提交 [pull request](.gitee/PULL_REQUEST_TEMPLATE.zh-CN.md)

## 许可证
**m0**使用的许可证是Apache 2.0

## 联系我们
商务合作，请联系。