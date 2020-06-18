# Nginx 高可用细节记录

## 支撑高可用的细节

- Nginx 采用多进程+异步非阻塞方式（IO 多路复用 Epoll）。
- 请求的完整过程：建立连接→读取请求→解析请求→处理请求→响应请求。
- 请求的完整过程对应到底层就是：读写 Socket 事件。

## Nginx 处理连接

1. Nginx 启动时，Master 进程，加载配置文件。
2. Master 进程，初始化监听的 Socket。
3. Master 进程，Fork 出多个 Worker 进程。
4. Worker 进程，竞争新的连接`(accept_mutex) 理解为共享锁`，获胜方通过三次握手，建立 Socket 连接，并处理请求。
5. event模块里, 参数`multi_accept` 主进程决定对子进程处理请求(一个个的)

## 惊群问题

### 起因

master进程首先通过 socket() 来创建一个 sock 文件描述符用来监听，然后fork生成子进程（workers 进程），子进程将继承父进程的 sockfd（socket 文件描述符

之后子进程 accept() 后将创建已连接描述符（connected descriptor），然后通过已连接描述符来与客户端通信。

连接进来时，所有子进程都将收到通知并“争着”与它建立连接；当最终只有一个子进程能获取连接

### 优点

1. 当并发请求量大时，能保证一定的吞吐量, 请求效率高

### 缺点

1. 所有子进程都被唤醒，若进程数量较大，易引发资源消耗(上下文切换, 负载上升)

### 参数配置 accept_mutex

- `connection processing methods` 使用`epoll`方式, 通过`EPOLLEXCLUSIVE`标志, 解决惊群问题：不让多个进程在同一时间监听接受连接的socket，而是让每个进程轮流监听，这样当有连接过来的时候，就只有一个进程在监听;

- `accept_mutex` 参数默认是关闭；
