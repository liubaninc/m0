## 容器方式部署实例

> 本地容器化运行的示例，实际场景中可以用卷的方式挂载并覆盖配置。

docker-compose服务分类：
  - M0节点
    
    区块链节点不限于验证节点、数据节点等类型。
  - M0浏览器及钱包
    
    区块链数据同步、[浏览器界面](http://127.0.0.1:8088)、[钱包界面](http://127.0.0.1:8086)
    
  - M0自动发送交易
    

#### 准备工作

构建容器镜像
```shell script
# 构建m0节点镜像
make m0-image
```

#### 单个节点网络
 
##### 部署链
 ```shell script
 # 启动
 docker-compose -f samples/docker-compose-one.yaml --env-file samples/env up
```

#### 测试

[钱包]服务(http://localhost:8081)

[浏览器]服务(http://localhost:8082)

##### 删除链
```shell script
 # 删除
 docker-compose -f samples/docker-compose-one.yaml --env-file samples/env down -v
```
 
#### 4个节点网络

##### 部署链
 ```shell script
 # 启动
 docker-compose -f samples/docker-compose-four.yaml --env-file samples/env up
```

#### 测试

[钱包]服务(http://localhost:8081)

[浏览器]服务(http://localhost:8082)

##### 删除链
```shell script
 # 删除
 docker-compose -f samples/docker-compose-four.yaml --env-file samples/env down -v
```