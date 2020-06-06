# Redis Cluster 相关总结

## TCP端口号

在redis.conf配置文件里的`port=6379`是指连接Redis服务的端口号；而在Redis Cluster的通信(节点之间的数据通信频道、宕机侦测，配置变更，故障转移认证)，所使用的端口号是在原redis端口上加10000；需要对开通相应的端口号，否则，影响集群的搭建

## 节点故障与slot无法使用

### cluster-require-full-coverage 参数

当一个节点的主从都发生故障，导致此节点的slot无法继续使用，并引发请求错误；`cluster-require-full-coverage` 默认为yes 当缺失slot, 整个集群将无法使用; 
若设置为`cluster-require-full-coverage no` 其余slot仍能继续使用, 只是故障节点无法进行读写操作

### cluster-node-timeout 参数

参数单位为毫秒；用作判断节点失效超时长度；

- 若设置长度过短，主节点处于慢查询时，容易被误认为主节点处于失效状态;

- 当处于失效超时时长内，若出现多个主节点故障，会影响failover的投票机制, 导致无法选取从节点作为主节点，来实现故障转移
