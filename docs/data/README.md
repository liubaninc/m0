# 消息结构

- 花费utxo

```protobuf
message Input {
  string ref_tx = 1;
  int32  ref_msg = 2;
  int32  ref_offset = 3;
  string from_addr = 4;
  cosmos.base.v1beta1.Coin amount = 5 [(gogoproto.nullable) = false];
  int64  frozen_height = 7;
}
```

- 未花费utxo

```protobuf
message Output {
  string to_addr = 1;
  cosmos.base.v1beta1.Coin amount = 2 [(gogoproto.nullable) = false];
  int64  frozen_height = 4;
  bool change = 5;
}
```



> 交易可知信息：
>
> - 交易手续费金额
>
> - 交易gas消耗
> - 交易描述
> - msg 列表

## [资产管理](utxo.md)

## [合约管理](wasm.md)

## [验证人管理](validator.md)

## [权限角色列表](permission.md)

## [节点管理](peer.md)

## [证书管理](pki.md)

