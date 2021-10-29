# gRPC记录

## 基本理论

1. gRPC是HTTP协议的(HTTP/2); 不能在`HTTP/1.x`进行传递; header标识: `content-type: application/grpc`

2. gRPC把元数据(metadata)放到HTTP/2 Headers里，请求参数序列化之后放到 DATA frame里; 

3. gRPC的请求是Http的POST请求; 请求body是使用`protocol buffers`编码

### Nginx反向代理gRPC服务

1. gRPC协议: `grpc://`, `grpcs://`(ssl加密) Nginx反向代理配置: `grpc_pass grpc://localhost:50051;`

## 基本工程环境搭建

1. 安装`protoc` 根据系统类型，[protobuf](https://github.com/protocolbuffers/protobuf/releases) 将安装路径添加到环境变量里

2. 安装`protoc-gen-go` 插件: `go get -u github.com/golang/protobuf/protoc-gen-go` 可指定版本: `go get -u github.com/golang/protobuf/protoc-gen-go@1.3.2`[此版本较稳定]

二进制文件会存放在GOPATH的bin文件里，此文件路径也需要添加进环境变量里

### 关联生态

1. protoc-gen-gogofaster: [gogo/protobuf](https://github.com/gogo/protobuf) 基于protoc-gen-go 做出相关优化

2. protoc-go-inject-tag: 对生成的Message成员进行自定义注入

#### 查看Http2 数据帧传递

`GODEBUG=http2debug=2 go run main.go` 监控gRPC 应用层数据帧传递日志, 但无法查看传输层TCP，IP、端口等行为

## 开发细节

### gRPC 超时处理流程

超时时长的设置是由Client进行配置, 调用时传入了带 timeout 的 ctx: `context.WithTimeout`

1. gRPC的超时计算是可以跨进程累计, 超时传递；若请求链路如: `SrvA -> SrvB -> SrvC` 配置的超时时长，需要考虑每个请求的耗时时长，从而，设置合理的总超时时长；timeout = 5s => srvA(3s)->srvb(2s)->srvc(报超时异常)

2. 超时时长剩余时长都存储在 Http2 的metadata 里 (Headers 帧)；key 为 `grpc-timeout`; 当请求进入服务里，会解析剩余超时时长, 若已超出时长, 则不再完成剩余的请求, 响应超时的异常码

### gRPC 优雅关闭连接

利用Http2 GoAway 帧信号关闭连接 服务端对每一个连接都发送关闭连接信号；客户端会从发送流里识别 GoAway 帧信号，从而客户端关闭活动的stream


