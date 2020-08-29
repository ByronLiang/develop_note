# WebSocket Detail

## Basic

- websocket协议和http协议都属于应用层协议(TCP/IP之上)
- WebSocket协议里，由`握手`和`传输数据`组成；握手是建立在HTTP协议上；握手成功，则运行在TCP协议之上的；
- WebSocket是客户端-服务端全双工通信

### 握手协议特有字段

- Upgrade 升级到的协议；字段值为websocket

- Connection 升级协议；字段值为Upgrade

- Sec-WebSocket-Key 和 Sec-WebSocket-Accept 客户端与服务端握手使用

- Sec-WebSocket-Version websocket协议版本号

### Http协议与WebSocket协议的区别

1. TCP/IP 协议，是网络七层协议的第四层，本身没有长连接或短连接的区别；
2. HTTP 是基于 TCP 协议之上的「短连接」应用层协议，它的出现极大简化了网络应用的实现门槛，丰富了应用；单向连接(只能客户端发送请求，服务端被动接收；服务端没有主动发起对话的能力)
3. Socket 是操作系统内置的一套操作 TCP/IP 协议的网络（套接字）的方法；
4. Websocket 是跟 HTTP 对应的，基于 TCP 协议之上的「长连接」协议。

### Http协议 与 TCP协议区别

1. Http协议属于应用层；TCP协议是属于传输层
2. Http协议于TCP协议之上建立的，http在发起请求时通过tcp协议建立起连接服务器的通道，请求结束后，立即断开tcp连接
3. HTTP连接是`无状态短连接`; TCP连接属于有状态(三次握手/四次挥手)长连接(TCP keepalive机制)

#### TCP状态

1. 三次握手: LISTEN, SYN_SENT, SYN_RECV, ESTABLISHED
2. 四次挥手: FIN_WAIT1, CLOSE_WAIT, FIN_WAIT2, LAST_ACK, TIME_WAIT

### Socket 与 WebSocket的区别

1. socket为套接字；是应用层和传输层之间的抽象层；socket是一组接口，把复杂的TCP/IP协议族隐藏在socket接口后面；
2. WebSocket在网络七层协议上的层级等同于Http，而Socket位置处于七层协议中的第四层，Socket是操作系统对TCP、UDP的封装。WebSocket处在上层，Socket处在下层，WebSocket依赖于Socket，Socket为WebSocket服务