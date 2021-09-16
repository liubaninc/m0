# 资产管理

支持一对一，一对多，不支持多对多

>  **string desc = 2; **  将来会删除

- 资产发行

```protobuf
message MsgIssue {
  string creator = 1;
  string desc = 2;

  repeated Input inputs = 10;
  repeated Output outputs = 11;
}
```

- 资产转账

```protobuf
message MsgSend {
  string creator = 1;
  string desc = 2;

  repeated Input inputs = 10;
  repeated Output outputs = 11;
}
```

- 资产销毁

```protobuf
message MsgDestroy {
  string creator = 1;
  string desc = 2;

  repeated Input inputs = 10;
  repeated Output outputs = 11;
}
```



>  可知信息：
>
> - 发送人&发送金额
>
> - 接收人&接收金额
>
> - 地址及其地址金额变化
>
> - 输入总金额与输出总金额差（资产发行、销毁时为***非零***）

