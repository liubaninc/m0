# 验证人管理

- 验证人信息

```protobuf
message Description {
  string moniker = 1;
  string identity = 2;
  string website = 3;
  string details = 4;
}
```

- 新增验证人

```protobuf
message MsgCreateValidator {
  string creator = 1;
  string pubKey = 3;
  Description description = 5;
}
```

> 可知信息：
>
> - 发起人
> - 验证人
> - 验证人信息

- 更新验证人

```protobuf
message MsgEditValidator {
  string creator = 1;
  string pubKey = 3;
  Description description = 5;
}
```

> 可知信息：
>
> - 发起人
> - 验证人
> - 验证人信息

- 删除验证人

```protobuf
message MsgLeaveValidator {
  string creator = 1;
  string pubKey = 3;
}
```

> 可知信息：
>
> - 发起人
> - 验证人