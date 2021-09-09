# 公钥基础设施管理

## 简介

## 执行命令 `m0d tx pki`

```shell
m0d tx pki --help
```

```
pki transactions subcommands

Usage:
  m0d tx pki [flags]
  m0d tx pki [command]

Available Commands:
  add-cert         add an intermediate or leaf certificate signed by a chain of certificates which must be already present
  add-root-cert    add a new self-signed root certificate
  freeze-cert      freeze an intermediate or leaf certificate
  revoke-cert      revoke the given intermediate or leaf certificate.All the certificates in the subtree signed by the revoked certificate will be revoked as well.
  revoke-root-cert revoke the given root certificate.All the certificates in the subtree signed by the revoked certificate will be revoked as well.
  unfreeze-cert    unfreeze an intermediate or leaf certificate
```

- 上传根证书

```shell
m0d tx pki add-root-cert --help
```

```
add a new self-signed root certificate

Usage:
  m0d tx pki add-root-cert [certificate] [flags]
```

- 吊销根证书

```shell
m0d tx pki revoke-root-cert --help
```

```
revoke the given root certificate.All the certificates in the subtree signed by the revoked certificate will be revoked as well.

Usage:
  m0d tx pki revoke-root-cert [subject] [subjectKeyID] [flags]
```

- 上传证书

```shell
m0d tx pki add-cert --help
```

```
add an intermediate or leaf certificate signed by a chain of certificates which must be already present

Usage:
  m0d tx pki add-cert [certificate] [flags]
```

- 吊销证书

```shell
m0d tx pki revoke-cert --help
```

```
revoke the given intermediate or leaf certificate.All the certificates in the subtree signed by the revoked certificate will be revoked as well.

Usage:
  m0d tx pki revoke-cert [subject] [subjectKeyID] [flags]
```

- 冻结证书

```shell
m0d tx pki freeze-cert --help
```

```
freeze an intermediate or leaf certificate

Usage:
  m0d tx pki freeze-cert [subject] [subjectKeyID] [flags]
```

- 解冻证书

```shell
m0d tx pki unfreeze-cert --help
```

```
unfreeze an intermediate or leaf certificate

Usage:
  m0d tx pki unfreeze-cert [subject] [subjectKeyID] [flags]
```

##  查询命令 `m0d query pki`

```shell
m0d query pki --help
```

```
Querying commands for the pki module

Usage:
  m0d query pki [flags]
  m0d query pki [command]

Available Commands:
  list-certificate  list all certificate
  list-certificates list all certificates
  show-certificate  shows a certificate
  show-certificates shows a certificates
```

