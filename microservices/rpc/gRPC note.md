# gRPC记录

## 基本理论

1. gRPC是HTTP协议的(HTTP/2); 不能在`HTTP/1.x`进行传递; header标识: `content-type: application/grpc`

2. gRPC把元数据(metadata)放到HTTP/2 Headers里，请求参数序列化之后放到 DATA frame里; 

3. gRPC的请求是Http的POST请求; 请求body是使用`protocol buffers`编码

### Nginx反向代理gRPC服务

1. gRPC协议: `grpc://`, `grpcs://`(ssl加密) Nginx反向代理配置: `grpc_pass grpc://localhost:50051;`
