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

## Redis Sentinel哨兵

### 基本架构

1. sentinel集群是监视Redis实例，选举Redis主实例；sentinel集群节点数量是`quorum`法定人数，通过raft等算法，确保实例高可用

2. Redis主从架构: 由至少一个主节点与两个从节点组成；主节点被sentinel集群监控; 主节点复制数据到从节点

### 实例故障

1. 主观下线: 主节点无法响应；因执行耗时长的lua脚本(忙状态)，无法响应

2. 客观下线: 被sentinel集群，选举而判断下线

### Gossip分布式协议

Redis Cluster集群，节点间通信是使用Gossip分布式协议; 它是一种冗余的容错算法和最终一致性算法(遵循BASE)， 不是强一致性协议

#### 原理

各个节点的信息行为(新上线/下线/异常节点) 会被传递出去；是周期性地随机选择一些节点进行传递给这些节点。收到信息的节点再会传递给其他一些随机选择的节点。一般而言，信息会周期性的传递给N个目标节点

- 优势

1. 节点扩展，不影响消息传递

2. 节点可以重复传递消息，遇到网络故障, 可补偿重新传递消息

- 劣势

1. 由于依靠节点间层层传递信息，信息行为具有一定的延迟；但最终，整个集群里都会收到信息

2. 消息冗余: 同一节点多次接收同一消息，增加消息处理的压力。一次通信会对网路带宽、CUP资源造成很大的负载