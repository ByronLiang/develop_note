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
