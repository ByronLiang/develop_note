# 涉及swoole基本概念

- swoole启动后, 会产生Master进程、Manager进程；master控制各种线程；manager控制不同的worker进程
- `ps aux | grep xxx.php 看下一下进程PID`

- 当`worker_num` 大于1时，进程间无法共享数据; 需要借助Redis 或 Swoole Table 存储key-value数据
