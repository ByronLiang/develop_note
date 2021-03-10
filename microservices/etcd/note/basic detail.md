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

## 网络分区问题

