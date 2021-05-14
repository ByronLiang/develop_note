# HTTP断点续传原理与实现

## 断点续传原理

### 服务端

- 服务端响应的header有`accept-ranges: bytes` 表明支持设置返回指定起始的byte字节数据

- 识别客户端请求的header里有Range，发起分片返回数据，响应header里`Content-Range: bytes=起始值-结束值/数据长度` 返回当前返回分片数据的起始标志值

- 识别客户端发起的正确的分片请求，响应状态码：`206 Partial Content`; 客户端header参数有误, 响应状态码: `416 Range Not Satisfiable`

### 客户端

- 当识别服务端支持分片续传数据后，请求header里配置`Range: bytes=from数值-to数值`

## 实现

- 参考Go SDK 里 即时文件数据处理分块`http.ServeContent()` 与 针对服务路径文件处理分块`http.ServeFile()`

### goroutine实现多线程发起分块文件下载请求

- 识别服务端是否支持分片下载及获取下载文件的content-length

- 按照配置的长度进行分片，每个分片数据长度越小, 响应越快速；需要确保每个分片的顺序一致；

#### 可用性与一致性

1. 对比合并文件后的总长度`length与content-length`;

2. `文件SHA-256校验`: sha256加密算法处理下载的文件内容，对比原文件的sha256, 确保下载的文件没有募改

