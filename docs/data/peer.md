# 节点管理

- 新增节点

```
message MsgCreatePeerID {
  string creator = 1;
  string index = 2;
  string certIssuer = 3;
  string certSerialNum = 4;
}
```

> 可知信息：
>
> - 发起人
> - 节点ID
> - 证书发行者
> - 证书序列号

- 更新节点

```
message MsgUpdatePeerID {
  string creator = 1;
  string index = 2;
  string certIssuer = 3;
  string certSerialNum = 4;
}
```

> 可知信息：
>
> - 发起人
> - 节点ID
> - 证书发行者
> - 证书序列号

- 删除节点

```
message MsgDeletePeerID {
  string creator = 1;
  string index = 2;
}
```

> 可知信息：
>
> - 发起人
> - 节点ID