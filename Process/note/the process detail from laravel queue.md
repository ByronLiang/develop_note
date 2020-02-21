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
