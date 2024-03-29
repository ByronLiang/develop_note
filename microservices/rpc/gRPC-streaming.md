# gRPC Streaming 流请求 与 Unary 一元请求 思考

## 两种请求模式: 流请求 与 一元请求

Unary: 直接使用请求与响应结构体对象在 HTTP2 的请求与响应进行映射

Streaming: 多个请求和响应通过长期存活 `(long lived tcp connection)` 的 HTTP2 流进行交换，该流可以是单向或双向的

### 连接比较

只需要一个 TCP 连接，每次发消息无需额外连接开销

Unary: 一元请求需要为每个请求建立一个新的 HTTP2 流，包括通过网络发送的附加标头帧

Streaming: 无需新建连接流与标头帧, 只需在流式中，添加数据帧就可以发送数据。

## 优势与应用场景

### Unary请求优势

- 能有效应对高并发请求场景, 通过多个协程处理连接请求与响应。并且具有负载均衡策略

- 复杂度低，对网络短暂异常，只影响部分请求

### Streaming 流请求优势

- 只需维持一个 TCP 连接, 高效传输数据, 减小相关网络帧

- 无法实现负载均衡，只适合对单一服务或具有连接状态的数据实现传输, 如订阅/推送模式

- 针对高并发请求流量，吞吐量不够好，可配置多个流客户端提升数据吞吐量

### 应用场景

微服务场景下，服务进程间数据请求与响应适合使用一元请求

对于即时通讯平台, websocket 连接具有状态性, 当连接层将消息数据请求到指定 IM节点 时，适合使用 stream 流传输
