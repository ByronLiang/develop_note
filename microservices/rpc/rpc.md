# RPC服务基本理论

## 核心

- 服务寻址
- 数据流的序列化和反序列化
- 网络传输

### 服务寻址

- 每个调用的函数都有唯一的Call ID 是一种映射关系

- 当客户端需要进行远程调用时，查找对应表，找出相应的 Call ID，然后把它传给服务端，服务端也通过查表，来确定客户端需要调用的函数，然后执行相应函数的代码

### 数据流的序列化和反序列化

一般使用ProtoBuf, 也可使用文本传输JSON

### 网络传输

传输层: TCP/UDP; 应用层: HTTP

## RPC类型

数据传输方式: TCP/UDP(传输层)；Http(应用层)

### 基于TCP协议

通过长连接减少连接的建立所产生的花费

### 基于HTTP协议的RPC

使用Http2.0：多路复用特性; 增加请求并发性; 有效避免`request-response的阻塞模型`

## RPC框架使用ProtoBuf为消息序列化方式

### 数据结构差异

JSON序列化的数据结构是文本, 较好的可读性; ProtoBuf序列化的数据结构是二进制信息格式(字节流)，可读性差

### 编码解码性能

JSON针对数值、浮点数据结构, 编码性能较弱；ProtoBuf更小的数据体积，利于高效传输；高速的编码与解码性能