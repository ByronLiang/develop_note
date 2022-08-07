# stream 消息队列

[官方文档](https://redis.io/docs/manual/data-types/streams)

- 轻量消息中间件，具有生产者与消费组，分区概念。相比分布式消息中间件 Kafka, stream 消息队列可应用范围较低

- 具备一定可用性，具有消息消费响应(ACK)

## 设计

自增消息ID: 纳秒级别时间戳拼接序列数，当同一纳秒数，对序列数进行递增

