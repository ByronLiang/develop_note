# TCP Keepalive 长连接与心跳包

## Linux内核配置

### 配置文件

Linux 配置参数路径: `/proc/sys/net/ipv4` 

1. `tcp_keepalive_intvl` 在tcp_keepalive_time之后，没有接收到对方确认，继续发送保活探测包的发送频率，默认值为75s。

2. `tcp_keepalive_probes` 在tcp_keepalive_time之后，没有接收到对方确认，继续发送保活探测包次数，默认值为9（次）

3. `tcp_keepalive_time` 在TCP保活打开的情况下，最后一次数据交换到TCP发送第一个保活探测包的间隔，即允许的持续空闲时长，或者说每次正常发送心跳的周期，默认值为7200s（2h）。

4. 开始探测到放弃探测确定连接断开的时间 = `tcp_keepalive_probes` * `tcp_keepalive_intvl`

5. TCP keep-alive是通过在空闲时发送TCP Keep-Alive数据包，然后对方回应TCP Keep-Alive ACK来实现的。

### 应用层的使用配置

1. 默认下，TCP Keepalive是关闭的；只有服务端是基于socket协议的TCP服务, 并且socket配置使用Keepalive, 才会使用TCP Keepalive内核配置；而WebSocket 是基于Socket的具体实现；

2. WebSocket协议是独立的基于TCP协议；应用上, 主要在业务层进行维持心跳`(Ping/Pong数据帧)` 较少依赖内核来维持心跳、检测连接；

## Http 与 TCP 的 Keepalive

1. HTTP协议(Http1.1, 2.0默认开启keep-alive)的Keep-Alive主要应用连接复用，同一个连接上串行方式传递请求-响应数据
2. TCP的keepalive机制意图在于保活、心跳，检测连接错误。(IM服务，WebSocket)

## WebSocket 与 TCP Keepalive

1. TCP Keepalive 默认关闭的; 开启时，需要服务端与客户端都需开启；

2. 不能应用于代理层, 只能在点对点里起作用

3. `(Ping/Pong数据帧)`是能动态调整, 基于应用层配置


