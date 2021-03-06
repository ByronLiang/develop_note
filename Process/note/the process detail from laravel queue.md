# Laravel框架下的消息队列涉及的进程信号笔记

## 流程

- 当启动消息队列时`php artisan queue:work` 会建立信号监听处理机制；
- 消息队列消费是以单进程进行处理，能保持消息的顺序

1. `pcntl_async_signals` 开启信号接收并回调处理；
2. `pcntl_signal(SIGTERM, handle())` 建立监听进程终止的事件回调
3. `pcntl_signal(SIGUSR2, handle())` `pcntl_signal(SIGCONT, handle())` 监听进程暂停与重新启动的信号
4. 未避免消费队列消息的时候出现超时情况: 为进程设置一个alarm闹钟信号`pcntl_alarm(int $seconds)`
5. 监听定时信号`pcntl_signal(SIGALRM, handle())`

当手动关闭消息队列 进程会发出`SIGTERM`信号；当暂无消息可以消费的时候，进程将暂停, 处于休眠状态；

### 产生消息流程

1. 消息组成：消息内容，消息唯一性ID, 消息已重试次数`attempts`
2. 对于Redis驱动的消息队列，以`rpush`命令放置消息命令
3. 生产消息成功，会返回消息唯一性ID, 保证已成功投递消息到消息队列里

### 消息消费流程

1. 对于以Redis驱动的消息队列，会先取延时队列`delay[zsort]`与待重试`reserved[zsort]`并将已到期的消息转移到待消费队列[list]
2. 以`lpop`命令取出消息，若消息消费过程出现错误，会被放置待重试的队列中
3. 为确保原子性，消息消费与消息队列迁移都需要使用`lua`脚本

### 多进程处理消费队列

1. 初始一个进程，建立定时器， 取出延时消息，并将其转移到待消费队列里
2. 初始化指定数量的进程，从待消费队列里消费消息
3. 当监听到相关中断信号量, 父进程会对子进程进行回收处理, 避免出现僵尸进程
