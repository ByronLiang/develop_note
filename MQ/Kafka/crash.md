# 常见故障分析与官方分析工具

## 消费请求卡死

### 分析流程

#### 官方工具脚本

`kafka-topic.sh --zookeeper 127.0.0.1:9092 --list`

`kafka-topic.sh --zookeeper 127.0.0.1:9092 --describe` 查看当前 broker 节点的topic 和 消费分区详情 `(_consumer_offsets)`

#### 日志

默认是在 kafka 目录的 logs，以 controller 前缀为 broker 控制器日志

消费线程无法对分区进行消费数据

#### 现象

1. 从 `_consumer_offsets` 详情里，某些 partition 的 leader 那列数据为 -1，表明暂无法选出主控制器。

#### 结果

因网络异常后，引发集群脑裂: 由于新老 controller 之间感知的可用节点不同，导致新 controller 对某个分区的 Leader 在内存中的信息与 zk 记录元数据的信息不一致，导致 controller 选举流程出现错误，选不出 Leader

