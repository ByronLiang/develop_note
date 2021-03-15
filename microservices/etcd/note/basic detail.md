# ETCD底层相关知识

## 协议与连接处理

1. 多数情况下, 集群里一个leader节点，其他节点是Follower状态；

2. 当客户端发出请求，leader处理请求；当请求发向follower节点，会将请求重定向给集群里的leader节点进行处理；

3. 选举节点是由follower节点转换而来：长时间无法从leader收到心跳消息，会令选举定时器过期，从而变成选举节点，进行leader节点的选举流程

## 分布式一致性处理

以日志记录，维护leader 与 各个follower 节点数据一致；涉及日志索引(复制位置)

### 流程

1. leader节点接收客户端的请求；将更新操作以消息形式发送到各个follower节点；

2. follower节点收到记录，存储本地日志里，并响应给leader节点

3. leader节点收到半数follower响应，会对客户端进行应答；

4. leader节点对请求操作进行commit(执行处理); 通知follower节点，follower节点也进行操作到本地状态机中

### 新节点加入集群

1. 新节点需要与leader节点同步log数据；当log数据一致，才能成为follower节点

### 持久层存储

1. 数据的查询因不涉及范围的筛选，只基于单一key进行数据查询；存储层数据结构使用BoltDB

2. 请求数据存储，利用MVCC与CAS机制, 处理并发请求处理数据

## 网络分区问题

对外提供Http接口请求etcd服务；集群节点使用gRPC进行通讯

### 网络分区与Quorum机制

1. 分区里节点数与Quorum机制有关系；若分区节点数不满足Quorum机制，此分区因无法选出leader节点导致无法可用

2. 出现不同分区，被隔离的从节点因无法与原leader节点通信，因此会触发选举，建立leader节点

3. 当网络分区完成修复，集群出现两个leader, 使用最新成为leader节点(term值更大)为总集群leader；其余leader节点的日志将提交到总leader的节点上


