# 合约开发

## WASM合约开发接口

C++接口详细可参见： [core/contractsdk/cpp/mchain/mchain.h](https://github.com/liubaninc/m0/blob/master/x/wasm/xmodel/contractsdk/cpp/src/mchain/account.h)

| API                                                          | 功能                         |
| ------------------------------------------------------------ | ---------------------------- |
| map<string, string>& args()                                  | 获取传入合约的参数表         |
| string& arg(string& name)                                    | 获取传入合约的指定参数值     |
| string& initiator()                                          | 获取发起此合约调用的账号     |
| int auth_require_size()                                      | 获取授权此合约调用的账号数   |
| string& auth_require(int idx)                                | 获取授权此合约调用的指定账号 |
| bool get_object(string& key, string* value)                  | 进行一次读操作               |
| bool put_object(string& key, string& value)                  | 进行一次写操作               |
| bool delete_object(string& key)                              | 进行一次删除操作             |
| bool query_tx(string &txid, Transaction* tx)                 | 查询指定id的交易内容         |
| bool query_block(string &blockid, Block* block)              | 查询指定id的区块内容         |
| void ok(string& body)                                        | 构造状态码为成功的返回       |
| void error(string& body)                                     | 构造状态码为失败的返回       |
| string& transfer_amount()                                    | 获取合约调用操作中的转账数额 |
| unique_ptr<Iterator> new_iterator(string& start, string& limit) | 获得遍历合约存储的迭代器     |
| call(module, contract, method, args ... )                    | 调用其他合约                 |

-  get_object

bool ContextImpl::get_object(const std::string& key, std::string* value)

输入

| 参数  | 说明                 |
| ----- | -------------------- |
| key   | 查询的key值          |
| value | 根据key查到的value值 |

输出

| 参数  | 说明                       |
| ----- | -------------------------- |
| true  | key值查询成功，返回value值 |
| false | key值不存在                |

- put_object

bool ContextImpl::put_object(const std::string& key, const std::string& value)

输入

| 参数  | 说明                   |
| ----- | ---------------------- |
| key   | 存入的key值            |
| value | 存入key值对应的value值 |

输出

| 参数  | 说明       |
| ----- | ---------- |
| true  | 存入db成功 |
| false | 存入db失败 |

- delete_object

bool ContextImpl::delete_object(const std::string& key)

输入

| 参数 | 说明            |
| ---- | --------------- |
| key  | 将要删除的key值 |

输出

| 参数  | 说明     |
| ----- | -------- |
| true  | 删除成功 |
| false | 删除失败 |

- query_tx

bool ContextImpl::query_tx(const std::string &txid, Transaction* tx)

输入

| 参数 | 说明                    |
| ---- | ----------------------- |
| txid | 待查询的txid            |
| tx   | 得到此txid的transaction |

输出

| 参数  | 说明         |
| ----- | ------------ |
| true  | 查询交易成功 |
| false | 查询交易失败 |

- query_block

bool ContextImpl::query_block(const std::string &blockid, Block* block)

输入

| 参数    | 说明                 |
| ------- | -------------------- |
| blockid | 待查询的blockid      |
| block   | 得到此blockid的block |

输出

| 参数  | 说明          |
| ----- | ------------- |
| true  | 查询block成功 |
| false | 查询block失败 |

## WASM合约开发案例

1. 编写wasm合约

参考源码样例 xmodel/contractsdk/cpp/example/counter.cc 主要实现一个struct的任意方法，来实现自己的逻辑

2. 编译wasm合约

```shell
export XDEV_ROOT=$GOPATH/src/github.com/liubaninc/m0/x/wasm/xmodel/contractsdk/cpp 
xdev counter.cc -o counter.wasm
```
3. 部署wasm合约

```shell
# 部署合约
m0d tx utxo deploy testcounter counter.wasm --chain-id=testnet {\"creator\":\"aaa\"} --gas auto --from alice
```
4. 调用合约

```shell
# 预估gas
m0d tx utxo invoke testcounter increase --chain-id=testnet {\"key\":\"aaa\"} --gas auto --from alice
```

5. 查询合约

```shell
m0d query utxo query testcounter get {\"key\":\"aaa\"}
```

6. 升级合约

```shell
# c++合约
m0d tx utxo upgrade testcounter counter.wasm --chain-id=m0-testnet --gas auto --from alice
```

