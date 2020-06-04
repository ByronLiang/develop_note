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

### 总结

1. 关闭过程里, 发送`FIN` 可以理解为发起连接的`SYN`
2. 发回ACK标识, 是确保已收到FIN标识；并且ACK的数据是基于FIN序号+1
3. 关闭连接进行四次挥手, 为实现主动关闭与被动关闭方都发起一次FIN标识 与 回应一次ACK 的验证流程 此流程为确保双方都不再进行传输数据
4. 两端谁发送完数据都需要自己告诉对方一次，并且对方确认一次。
5. TCP是全双工的，client与server都占用各自的资源发送segment（同一通道，同时双向传输seq和ack），所以，双方都需要关闭自己的资源（向对方发送FIN）并确认对方资源已关闭（回复对方Ack）

