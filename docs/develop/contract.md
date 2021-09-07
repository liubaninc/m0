# 智能合约接口使用说明

|API|功能|
|---|---|
|rpc PutObject(PutRequest) returns (PutResponse)|	产生一个读加一个写|
|rpc GetObject(GetRequest) returns (GetResponse)|	生成一个读请求|
|rpc DeleteObject(DeleteRequest) returns (DeleteResponse)|	产生一个读加一个特殊的写|
|rpc NewIterator(IteratorRequest) returns (IteratorResponse)|	对迭代的key产生读|
|rpc QueryTx(QueryTxRequest) returns (QueryTxResponse)|	查询交易|
|rpc QueryBlock(QueryBlockRequest) returns (QueryBlockResponse)|	查询区块|
|rpc ContractCall(ContractCallRequest) returns (ContractCallResponse)|	合约调用|
|rpc Ping(PingRequest) returns (PingResponse)|	探测是否存活|
|rpc GetCallArgs(GetCallArgsRequest) returns (CallArgs)|	得到合约调用参数|

## C++接口API
get_object
```cgo
bool ContextImpl::get_object(const std::string& key, std::string* value)
```
输入
| 参数 | 说明 |
| :---- | :---- |
| key | 查询的key值 | 
| value | 根据key查到的value值 | 

输出
| 参数	| 说明 | 
| :---- | :---- |
| true	| key值查询成功，返回value值 | 
| false	| key值不存在 | 

put_object
```cgo
bool ContextImpl::put_object(const std::string& key, const std::string& value)
```

输入
| 参数	| 说明| 
| :---- | :---- |
| key	| 存入的key值| 
| value	| 存入key值对应的value值| 

输出
| 参数	| 说明| 
| :---- | :---- |
| true	| 存入db成功| 
| false	| 存入db失败| 

delete_object
```cgo
bool ContextImpl::delete_object(const std::string& key)
```

输入
| 参数	| 说明 | 
| :---- | :---- |
| key	| 将要删除的key值 | 

输出
| 参数	| 说明| 
| :---- | :---- |
| true	| 删除成功| 
| false	| 删除失败| 

query_tx
```cgo
bool ContextImpl::query_tx(const std::string &txid, Transaction* tx)
```
输入
| 参数	| 说明| 
| :---- | :---- |
| txid	| 待查询的txid| 
| tx	| 得到此txid的transaction| 

输出
| 参数	| 说明| 
| :---- | :---- |
| true	| 查询交易成功| 
| false	| 查询交易失败| 

query_block
```cgo
bool ContextImpl::query_block(const std::string &blockid, Block* block)
```

输入
| 参数	| 说明 | 
| :---- | :---- |
| blockid	| 待查询的blockid| 
| block	| 得到此blockid的block|

输出
| 参数	| 说明| 
| :---- | :---- |
| true	| 查询block成功| 
| false	| 查询block失败|
