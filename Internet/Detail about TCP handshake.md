# TCP 与 UDP 相关网络协议底层知识

## TCP handshake process

1. 客户端发送syn包(seq=x)到服务器，并进入SYN_SEND(发送)状态，等待服务器确认；

2. 服务器收到syn包，必须确认客户的SYN（ack=x+1），同时自己也发送一个SYN包（seq=y），即SYN+ACK包，此时服务器进入SYN_RECV（接收）状态；

3. 客户端收到服务器的SYN＋ACK包，向服务器发送确认包ACK(ack=y+1)，此包发送完毕，客户端和服务器进入ESTABLISHED（已建立）状态，完成三次握手。

### Why 3-way handshake

```sh
Client ---> Server         SYN
Client <--- Server     SYN ACK 
Client ---> Server     ACK
```

1. TCP是双向通讯协议；若只进行两次握手, 只能保证进行单向通讯(Client收到Server的ACK);确保能双向通讯，双方都需接收到对方发出的`ACK`

2.  SYN 标识含义: 请求证明此数据包通过; ACK 标识含义: 仅在SYN通过后才发送，以证明SYN通过了

3. 能有效避免失效连接出现，增强通讯信道的有效性

### 4-way handshake close connect

```sh
Client ------FIN-----> Server

Client <-----ACK------ Server

Client <-----FIN------ Server

Client ------ACK-----> Server
```

1. 主动发起关闭连接可在服务端或者客户端；具体可参考是否长连接, 有没提供具体长度的发送数据；主动发发起关闭，即发送一个FIN，用来关闭主动方到被动关闭方的数据传送

2. 被动关闭方收到主动关闭方的FIN包后，发送一个ACK给对方，确认序号为收到序号+1`(FIN + 1)`

3. 被动关闭方同样也向主动关闭方发一个FIN包

4. 主动关闭方收到FIN后，发送一个ACK给被动关闭方，确认序号为收到序号+1`(FIN + 1)`

#### 补充

- 当进程异常退出了，内核就会发送 RST 报文来关闭，它可以不走四次挥手流程，是一个暴力关闭连接的方式

- 安全关闭连接的方式必须通过四次挥手，它由进程调用close和shutdown函数发起 FIN 报文

### 总结

1. 关闭过程里, 发送`FIN` 可以理解为发起连接的`SYN`
2. 发回ACK标识, 是确保已收到FIN标识；并且ACK的数据是基于FIN序号+1
3. 关闭连接进行四次挥手, 为实现主动关闭与被动关闭方都发起一次FIN标识 与 回应一次ACK 的验证流程 此流程为确保双方都不再进行传输数据
4. 两端谁发送完数据都需要自己告诉对方一次，并且对方确认一次。
5. TCP是全双工的，client与server都占用各自的资源发送segment（同一通道，同时双向传输seq和ack），所以，双方都需要关闭自己的资源（向对方发送FIN）并确认对方资源已关闭（回复对方Ack）

## TCP TimeWait

### 流程

1. time_wait 是主动关闭TCP连接一方的状态; 当被动关闭方发起FIN至主动关闭方，并且主动方收到FIN, 响应发送回ACK, 会处于Time_Wait状态

2. 超时长度为2MSL: 接收FIN与发送Ack正好2个MSL

### 必要性 对TCP的可靠性传输的作用

1. 提供缓冲时间；避免ACK响应丢失，影响正确关闭连接；为提供重传预留时间`[被动方没有收到Ack，就会触发被动方重传FIN]`
2. 处理延迟到达的报文：由于路由器可能抖动，TCP 报文会延迟到达

### 优化

1. Time_Wait状态会比较耗费资源；客户端尽量成为主动关闭连接的一方

## TCP BackLog

backlog 是一个连接队列; 由半连接状态 与 全连接状态 两种队列大小

TCP三次握手的过程里，都于半连接状态队列里；完成三次握手，连接被放置于全连接的队列里

### 半连接状态

当服务端接收到客户端的SYN包，服务端会将此连接放进半连接队列`(SYN_RCVD状态)[syns queue]`

相关配置: SYN queue 队列长度由 `/proc/sys/net/ipv4/tcp_max_syn_backlog` 指定，默认为2048。

当队列已满时 客户端的请求将被抛弃，引发客户端无法连接到服务端

### 全连接状态

当服务器接收到客户端的ACK报文后(完成三次握手)，此连接将从半连接队列搬到全连接队列尾部，即 accept queue （服务器端口状态为：ESTABLISHED）

Accept queue 队列长度由 `/proc/sys/net/core/somaxconn` 和使用listen函数时传入的参数，二者取最小值。默认为128

若全连接队列已满, 将会丢弃客户端的ACK请求；容易引发客户端误认为已完成连接，出现调用超时；服务端需要进行重传(重传SYN+ACK给客户端)

### 总结

相关参数适用于Redis服务配置`[tcp-backlog]`：需要同时增加半连接状态`(tcp_max_syn_backlog)` 与 全连接状态的队列容量`(somaxconn)` ，实现高请求下，连接不被抛弃

## Enhance TCP Connect

### 优化发起连接

#### tcp_syncookies

1. 避免因为半连接队列`syns queue`已满，对请求连接进行丢弃

2. syncookies原理: 服务器根据当前状态计算出一个值，放在己方发出的 SYN+ACK 报文中发出，当客户端返回 ACK 报文时，取出该值验证，如果合法，就认为连接建立成功

#### tcp_abort_on_overflow

1. 当开启此参数设置, 全连接队列已满下, 服务端不会放弃与客户端的连接, 但会告知客户端, 丢弃本次握手过程连接(向客户端发送`RST`); 客户端收到异常为: `connection reset by peer`

2. 若服务端抛弃客户端连接，能有效应对突发高流量的访问特性


