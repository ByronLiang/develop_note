# Mattermost 工程代码阅读笔记

## Web 中间件上下文与鉴权

`web`目录下包含对请求的生命周期的相关处理; 

`handlers.go` 文件通过自定义Handle, 封装请求中间件逻辑, 包括对请求鉴权，执行HandleFun

每一个请求都对应申请一个`Context` 对象

## WS 连接句柄存放

初始化一定数量的 Hub 存放连接, 每个 Hub 使用select + channel 处理连接与断开连接事件, 保证线程安全；

多个 Hub 能保证一定负载性能，提高连接效率。通过对 `userId` 哈希处理，定位到对应的 Hub 对象里

### 同一用户多端在线设计

以用户ID为分组, 存放不同`ConnectionId`的WS句柄，从而实现一个用户ID 对应多个连接

多端消息同步，对多个连接进行写入数据操作

### 定时检测连接活性

每个 Hub 配置定时器, 当连接活跃时刻超出指定时长, 会被进行关闭

客户端每次探测(ping) 或 连接数据写入 都会刷新连接的活跃时刻

## WS 连接写数据

WS 连接句柄写操作需要线程安全，不能并发写，否则会有异常: `panic concurrent write to websocket connection`

### 使用一个协程单独消费写缓冲队列channel (Mattermost 方案)

每个连接，对于写操作，会初始化一个带有缓冲长度的channel, 并开启一个协程，从channel 消费数据进行写操作

适合高频写操作业务情景

#### 优点

具备聚合性, 所有涉及写操作事件被写入channel里, 并被唯一的协程进行消费，确保单线程写操作，保证写操作线程安全

#### 缺点

随着连接数增加，协程数量也随之增加, 内存消耗较大，对调度压力比较大

由于每个协程调度频率因素，会诱发消息堆积，无法及时消费channel 消息，引发消息发送时延

### 多个连接共享一个channel 单协程消费 (一个生产者对于一个消费协程)

确保指定连接写操作消息放进同一个channel 里, 适合低频写操作业务

#### 优点

协程数量与channel 数量 一对一关系, 确保线程安全, 无需额外锁

#### 缺点

单线程消费，由于channel 涉及多个连接的写操作数据，比较容易积压数据，引发消息发送时延

### 多协程消费同一个 channel (一生产者对多个消费协程)

写操作时，需要对连接进行加锁，避免并发写操作发生，减少写操作数据积压

#### 优点

可以动态增加与减少协程数量，有效避免数据积压

#### 缺点

每次写操作，都需要对连接进行加锁操作，确保同一连接，不会并发写数据



