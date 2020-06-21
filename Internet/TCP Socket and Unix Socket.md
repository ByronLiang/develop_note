# TCP Socket 与 Unix Socket 差异比较与应用

## Nginx 与 php-fpm 通信方式

进程间通信方式: 采用socket通信；类型: TCP Socket 与 Unix Socket

### TCP Socket

TCP Socket是IP加端口, nginx默认的通信方式,可以跨服务器,非常适合做负载均衡；适合Nginx 与 php-fpm 不在同一机器上

#### 通信流程

```sh
TCP Socket(本地回环):
Nginx <=> socket <=> TCP/IP <=> socket <=> PHP-FPM

TCP Socket(Nginx和PHP-FPM位于不同服务器):

Nginx <=> socket(应用层) <=> TCP/IP(传输层) <=> (数据链路)物理层 <=> 路由器(实现转发) <=> 物理层 <=> TCP/IP <=> socket <=> PHP-FPM
```

每完成一次通信请求(Nginx与php-fpm)，都需经过两次TCP四层协议; 完成数据请求与响应, 对于高请求, 会耗费内存与时延较大

### Unix Domain Socket

1. 由于不经过TCP层(传输层)，通信过程无需进行TCP四层协议; 是属于一种内核传播通信 

2. 采用Unix内核系统的方法实现通信: buffer之间缓存区;`socket, bind, send, recv`函数实现初始化, 监听绑定 发送接收；直接以文件形式,以stream socket通讯

3. 只适用于Nginx 与 PHP-FPM 都在同一机器上

#### 通信流程

```sh
Nginx <=> socket <=> PHP-FPM
```

## 补充

### Redis

Redis服务可配置使用Unix Socket; 默认使用TCP连接；若Redis不对外开启访问, 可使用unix socket 连接方式 具有优化作用

```sh
unixsocket /var/run/redis/redis.sock
unixsocketperm 700
```

采用 Unix Socket 会对远程连接 产生一定影响；若只以单机运行，不对外开放连接, 可采用Unix Socket 的连接方式进行优化

### MySQL

连接方式种类：

- `shared_memory` `pipe`（前两种方式适用于Windows系统）
- `TCP` `Unix Socket` Unix系统下

连接方案的选择：

- 若`host=localhost` 或不配置 默认都采用 Unix Socket
- 若`host=127.0.0.1` (填写IP地址) 采用TCP的连接方式

MySQL客户端连接MySQL服务 默认采用Unix Socket 连接

```sh
// 使用TCP协议连接MySQL服务
mysql -uroot -p --protocol=tcp
```
