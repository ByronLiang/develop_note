# 正向代理与反向代理

## Http Transport Proxy

在 Http 传输层配置代理地址，若代理地址`(ip:port)`所属机器处于关闭状态, 则会出现 `no route to host` 异常。表明在网络层面上，

### 如何理解 No Route to Host

主要表现在网络层面上无法连通: 因服务器防火墙拦截、无法 PING 服务器 (ICMP返回响应此 host 不可达)

### Transport Proxy 底层流程

支持 `socks5` `http` 与 `https` 代理类型

## 反向代理

[REVERSEPROXY](https://pandaychen.github.io/2021/07/01/GOLANG-REVERSEPROXY-LIB-ANALYSIS/)
