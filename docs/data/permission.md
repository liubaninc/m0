# 权限角色管理

- 设置权限角色列表

```
message MsgSetPermission {
  string creator = 1;
  string address = 2;
  repeated string perms = 3;
}
```

> 可知信息：
>
> - 发起人
> - 账户地址
> - 权限列表