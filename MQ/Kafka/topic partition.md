# 分区数量配置思考

## Topic 与 Partition 关系

一个topic 对应多个 partition；partition数量会影响生产者的吞吐量与消费者数量

## 配置思考因素

1. partition 创建数量与 Linux 可使用 socket 数目`(FD文件描述)`有关系, 若超出，会报异常: `Too many open file`

2. 分区数量过多，主从副本需要复制。broke 节点数据同步与切换耗时

3. 扩容分区数，会影响基于 key 哈希进行指定分区写入逻辑
