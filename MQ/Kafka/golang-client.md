# Go SDK 客户端选型思考

## 消费组

### 通过连接 zk 自建消费组

适合自建 zookeeper + Kafka 搭建集群，并且 zookeeper 节点地址能对外暴露。

通过连接 zookeeper, 监听 Kafka 集群事件, 从而控制消费组成员再均衡

### 连接 Kafka broker 集群

对于服务商的消息队列服务，因无法对外提供 zookeeper 地址，需要连接 Kafka broker 节点集群，从而实现获取消费组信息及再均衡业务

使用 `github.com/bsm/sarama-cluster` 能兼容低版本客户端

### 底层原理

每个分区都有一个协程监听获取消息，默认模式下，多个协程的消息向一个管道传送消息
