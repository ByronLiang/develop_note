# Nginx相关知识总结

## 简介

- Nginx 是一款轻量的、高性能的 HTTP服务器，同时也可以用作反向代理、负载平衡器和 HTTP 缓存。
- Nginx 采用了模块化、事件驱动的架构设计，使用了异步非阻塞的事件处理机制处理请求，使得在高负载下也能提供更可靠的性能。

## 核心模块

### Http模块与Mail模块主要处理与Http协议与Mail相关协议有关事件
	
- Http模块提供Http与Https协议，处理客户端发起的请求(Ajax & WebSocket)
- Mail模块提供邮件代理服务, 提供POP3、IMAP与SMTP协议的服务器

### 六大配置区域

1. main(全局设置)

- user 来指定Nginx Worker进程运行用户以及用户组，默认由nobody账号运行。

- worker_processes来指定了Nginx要开启的子进程数。每个Nginx进程平均耗费10M~12M内存。根据经验，一般指定1个进程就足够了，如果是多核CPU，建议指定和CPU的数量一样的进程数即可。我这里写2，那么就会开启2个子进程，总共3个进程。
- error_log用来定义全局错误日志文件。日志输出级别有debug、info、notice、warn、error、crit可供选择，其中，debug输出日志最为最详细，而crit输出日志最少。

- pid用来指定进程id的存储文件位置。

- worker_rlimit_nofile用于指定一个nginx进程可以打开的最多文件描述符数目，这里是65535，需要使用命令“ulimit -n 65535”来设置

2. events(nginx工作模式)

- use用来指定Nginx的工作模式。
	1. Nginx支持的工作模式有select、poll、kqueue、epoll、rtsig和/dev/poll。
	2. 其中select和poll都是标准的工作模式，kqueue和epoll是高效的工作模式，不同的是epoll用在Linux平台上，而kqueue用在BSD系统中，因为Mac基于BSD,所以Mac也得用这个模式，
	3. 对于Linux系统，epoll工作模式是首选。
- worker_connections用于定义Nginx每个进程的最大连接数，即接收前端的最大请求数，默认是1024。
最大客户端连接数由worker_processes和worker_connections决定，即Max_clients=worker_processes*worker_connections，
在作为反向代理时，Max_clients变为：Max_clients = worker_processes * worker_connections/4。进程的最大连接数受Linux系统进程的最大打开文件数限制，在执行操作系统命令“ulimit -n 65536”后worker_connections的设置才能生效
3. http(http设置)
4. server(主机设置) 【虚拟服务器设置：由listen和server_name指令组合定义】
5. location(url匹配) 【where when how】
6. upstream(负载均衡设置)