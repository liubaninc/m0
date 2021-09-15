# 证书管理

- 上传根证书

```protobuf
message MsgAddRootCert {
  string creator = 1;
  string certificate = 2;
}
```

> 可知信息：
>
> - 发起人
> - 证书

- 吊销根证书

```protobuf
message MsgRevokeRootCert {
  string creator = 1;
  string subject = 2;
  string subjectKeyID = 3;
}
```

> 可知信息：
>
> - 发起人
> - 证书主题

- 上传证书

```protobuf
message MsgAddCert {
  string creator = 1;
  string certificate = 2;
}
```

> 可知信息：
>
> - 发起人
> - 证书

- 吊销证书

```protobuf
message MsgRevokeCert {
  string creator = 1;
  string subject = 2;
  string subjectKeyID = 3;
}
```

> 可知信息：
>
> - 发起人
> - 证书主题

- 冻结证书

```protobuf
message MsgFreezeCert {
  string creator = 1;
  string subject = 2;
  string subjectKeyID = 3;
}
```

> 可知信息：
>
> - 发起人
> - 证书主题

- 解冻证书

```protobuf
message MsgUnfreezeCert {
  string creator = 1;
  string subject = 2;
  string subjectKeyID = 3;
}
```

> 可知信息：
>
> - 发起人
> - 证书主题

