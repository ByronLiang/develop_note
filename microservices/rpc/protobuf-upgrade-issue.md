# Protobuf 版本升级的兼容问题

`github.com/golang/protobuf@v1.3.3` 是版本升级的分水岭。同一级别协议文件，大多数情况可以兼容处理, 无需重新生成协议文件

升级主要带来：消息对象结构体差异，序列化方法；接口反射方式。

若 `proto-gen-go` 代码生成工具版本与 prtobuf 版本不一致，也会引发异常问题。

若进行跨级升级 protobuf 库，需要重新编译协议文件

## mismatching message name

因进行跨级升级 `protobuf` 库，引发不兼容原协议文件。

异常代码：`proto/registry.go -> RegisterType(m Message, s messageName)`

## 序列化与反序列化

因结构体方法可调用序列化方法，而新版本取消暴露的序列化方法，应需统一使用 `proto.Unmarshal`

## 扩展

当只生成 proto 协议，只需 `protoc --go_out=生成文件路径 *.proto`

若生成含有 gRPC 服务方法, 需要引用插件: 
