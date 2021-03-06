# 分布式事务

1. 在微服务生态下，本地事务无法满足数据一致性, 事务需要跨越多个数据源`(跨库)`

2. 只针对业务价值非常大的领域采用分布式事务(强一致性/时效性高);
对于一般数据处理, 不建议过于依赖事务进行处理; 确保实现BASE理论-最终一致性

3. 部分业务场景很难实现达到`可撤销/可补偿`的数据效果

4. 多采用补偿方式, 兜底方式来确保数据的一致性, 减少依赖事务处理异常

## 原理与特点

### 2PC

主要由参与者 与 协调者 组成; 流程: 准备阶段-提交阶段

缺陷: 协调者和参与者都挂了的情况，有可能导致数据不一致

### 3PC

三阶段提交就有`CanCommit`、`PreCommit`、`DoCommit`三个阶段。

1. 只是询问所有参与者是否可可以执行事务操作，并不在本阶段执行事务操作。

2. 当协调者收到所有的参与者都返回YES时，在才执行事务操作

3. 执行commit或者rollback。

### 事件型事务通知

利用MQ消息发送与消费: 主服务完成后将结果通过事件传递给从服务，从服务在接受到消息后进行消费，完成业务，从而达到主服务与从服务间的消息一致性。

1. 确保事件投递MQ成功，并对投递响应结果做成正确处理

2. 建立一定的重试机制, 保证事件系统（MQ消息队列）可用

### saga补偿模式

1. 业务补偿为主导, 对执行的行为进行log记录, 方便进行补偿时, 回滚已执行的操作

2. 考虑多次执行补偿, 幂等性问题

### TCC模式

如果Try在所有服务中都成功，那么执行Confirm操作，Confirm操作不做任何的业务检查（因为try中已经做过），只是用Try阶段预留的业务资源进行业务处理；否则进行Cancel操作，Cancel操作释放Try阶段预留的业务资源。

1. 每个服务都需要实现Confirm和Cancel两个接口

2. 上下游的服务依赖问题
