# Golang的Http Client相关总结

## 默认初始的Client

初始化无指定任何参数，无超时设置，使用默认的transport；

## DNS域名解析

针对域名请求，需要将域名解析成IP地址(address:port/IPv4/IPv6), IP地址为实现传输层的连接

### DNS域名解析流程

1. 默认的解析`DefaultResolver`；在`net/dial.go`包里`DialContext`, `resolveAddrList` 方法进行域名解析流程

2. 具体解析可参考: `net.LookupHost`对域名进行解析方法； `net\lookup.go`里的`lookupIPAddr`

3. `lookupIPAddr`流程: 读取本地`/etc/hosts`文件;
   `goLookupIPCNAMEOrder`解析是否对应域名的IP地址; 然后再到读取`/etc/resolv.conf`, 通过访问DNS服务器进行域名解析

### 优化域名解析

1. 复用传输层的连接(transport); 针对Transport的配置，指定dial的超时设置与keepalive; 

2. 避免对域名进行请求，转换成IP地址进行请求

3. 多次创建TCP请求时，每次还是需要依赖DNS的解析；可以自定义解析流程，增加DNS缓存
