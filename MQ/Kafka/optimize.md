# 弱网环境下 Kafka 相关优化

在不理想的物理网络链路(弱网), 高时延网络环境(海外机房与国内机房网络传输)，Kafka 消息服务会比较频繁出现 读写超时，容易进行消息重试与客户端会高频次进行服务探测(请求 Kafka 集群 metadata)，从而影响消息的吞吐量 (写入消息QPS)

## 常见现象

- Kafka 消息服务会比较频繁出现读写超时 (read/write i/o timeout), 属于传输层相关异常

- 因读写超时，会进行消息退避与重试。因需要优先完成重试消息，正常消息的发送会被搁置，会影响消息写入的吞吐量。

- 每次从 Kafka 服务获取 metaData，但都因网络因素，没法正常响应，客户端无法从 Kafka 集群获取可连接的 broke; 错误内容: `client/metadata got error from broker %d while fetching metadata`

## 相关解决方案

1. 弱网下，减少消息写入频率，降低网络请求频率: 增大 Flush 的时间间隔与最大消息数目，客户端在进程进行消息缓存。

2. 网络质量好，则可采取高频网络请求，以小体积消息发送 broker, 以达到高吞吐量，同时进程内存占用小。

3. 连接参数配置：将读超时与连接超时调大，增加连接等待时长。同时重试次数与退避时长也可以调大。

4. 对消息使用压缩率高算法，减小网络通讯传输体积。

## 讨论问题集合

[Async producer can overflow itself at high rate](https://github.com/Shopify/sarama/issues/805)

[AsyncProducer: i/o timeout error when message is big](https://github.com/Shopify/sarama/issues/834)

[Sarama Producer hangs when broker dies no errors are thrown](https://github.com/Shopify/sarama/issues/765)
