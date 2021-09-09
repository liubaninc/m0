# 技术文档

**M0** 模块化设计，按照业务划分不同模块。每个模块有自己的消息结构及相应的执行命令、数据状态查询命令。

支持的模块

- [utxo](spec/utxo/README.md) 资产管理, 发行/增发、转移、销毁资产
- [wasm](spec/wasm/README.md) 合约管理，部署、升级、调用wasm合约 
- [validator](spec/validator/README.md) 验证人管理，新增、删除验证人
- [permssion](spec/permssion/README.md) 角色权限管理，设置账户权限角色列表
- [peer](spec/peer/README.md) 节点管理， 新增、删除准入节点
- [pki](spec/pki/README.md) 公钥基础设施管理， 上传、吊销、冻结证书