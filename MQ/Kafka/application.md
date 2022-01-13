# 应用

## 消息顺序

多个Partition时，不能保证Topic级别的数据有序性; 若要实现消息顺序, 需要牺牲并发吞吐性

### 实现方式

1. 全局使用一个生产者

2. 全局使用一个消费者，只能是一个消费线程

3. 全局使用一个分区 (不同的表可以使用不同的分区或者topic实现隔离与扩展)

### 潜在问题

消息处理异常，消息重试引发顺序问题。需要将全流程进行回滚，重新执行

## 广播消费-共享订阅

### 组间广播

不同的消费组都对同一个topic进行订阅，能实现 topic 广播: 这两个消费组里的其中一个消费线程都监听相同的分区时，则它们都能获取消息, 两个过程是独立的

### 组内单播

同一消费组里的各个消费线程无法实现广播，它们都监听订阅topic的各个分区

## 消费选型: 消费组与单一消费者

### 消费组优势

1. 伸缩性, 负载均衡，新加入topic的分区, 能重新均衡化

### 劣势

1. rebalance 会耗时, 特别在多分区的topic下, 严重影响消费, 积压大量消息