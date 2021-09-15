# 合约管理

>  **string desc = 2; **  将来会删除

- 部署合约

```protobuf
message MsgDeploy {
  string creator = 1;
  string desc = 2;

  repeated utxo.Input inputs = 20;
  repeated utxo.Output outputs = 21;
  repeated InputExt inputs_ext = 22;
  repeated OutputExt outputs_ext = 23;
  string contract_name = 24;
  bytes contract_code = 25;
  xmodel.WasmCodeDesc contract_desc = 26;
  string args = 28;
  repeated xmodel.ResourceLimit resource_limits = 30;
}
```

>  可知信息：
>
> - 发送人&发送金额
>
> - 接收人&接收金额
>
> - 地址及其地址金额变化
> - 合约名称
> - 合约代码
> - 合约描述
> - 合约初始化参数

- 升级合约

```protobuf
message MsgUpgrade {
  string creator = 1;
  string desc = 2;

  repeated utxo.Input inputs = 20;
  repeated utxo.Output outputs = 21;
  repeated InputExt inputs_ext = 22;
  repeated OutputExt outputs_ext = 23;
  string contract_name = 24;
  bytes contract_code = 25;
  repeated xmodel.ResourceLimit resource_limits = 30;
}
```

>  可知信息：
>
> - 发送人&发送金额
>
> - 接收人&接收金额
>
> - 地址及其地址金额变化
> - 合约名称
> - 合约代码

- 调用合约

```protobuf
message InvokeRequest {
  string module_name = 1;
  string contract_name = 2;
  string method_name = 3;
  string args = 4;
  repeated xmodel.ResourceLimit resource_limits = 5;
  // amount is the amount transfer to the contract
  repeated cosmos.base.v1beta1.Coin amount = 6 [(gogoproto.nullable) = false, (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins"];
}
```



```protobuf
message MsgInvoke {
  string creator = 1;
  string desc = 2;

  repeated utxo.Input inputs = 20;
  repeated utxo.Output outputs = 21;
  repeated InputExt inputs_ext = 22;
  repeated OutputExt outputs_ext = 23;
  repeated InvokeRequest contract_requests = 24;
}
```

> 可知信息：
>
> - 发送人&发送金额
>
> - 接收人&接收金额
>
> - 地址及其地址金额变化
> - 合约调用列表
>   - 合约名
>   - 合约方法
>   - 方法参数

- 多方确认申请

**不支持资产，不支持支付手续费**

```protobuf
message MsgProposeDeployContract {
  string creator = 1;
  string contractName = 2;
  bytes contractCodeHash = 3;
  string initArgs = 4;
  repeated string approval = 5;
}
```

> 可知信息：
>
> - 发起人
> - 合约名称
> - 合约代码哈希
> - 合约初始化参数
> - 审批人列表

- 多方确认审批

**不支持资产，不支持支付手续费**

```protobuf
message MsgApproveDeployContract {
  string creator = 1;
  string index = 2;
}
```

> 可知信息：
>
> - 发起人
> - 合约名称

- 冻结

**不支持资产，不支持支付手续费**

```protobuf
message MsgFreeze {
  string creator = 1;
  string contract_name = 2;
}
```

> 可知信息：
>
> - 发起人
> - 合约名称

- 解冻

**不支持资产，不支持支付手续费**

```protobuf
message MsgUnfreeze {
  string creator = 1;
  string contract_name = 2;
}
```

> 可知信息：
>
> - 发起人
> - 合约名称

- 销毁

**不支持资产，不支持支付手续费**

```protobuf
message MsgUndeploy {
  string creator = 1;
  string contract_name = 2;
}
```

> 可知信息：
>
> - 发起人
> - 合约名称

