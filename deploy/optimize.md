# 服务优化与排查指南

一般以容器化部署应用，每个应用实例的资源已被隔离，但所有资源(CPU、内存)都来自宿主机。

对于公共服务资源抢占比较常见，比如 Redis 与 MySQL 等

## 接口响应慢/高尾时延 现象

### 硬件机器层面

结合日志系统 与 监控平台 查看机器资源: 是否有告警日志，内存OOM，CPU 占用 与 内存占用情况

#### CPU 与 内存 

- `top` `free -m` 查看机器CPU 进程与负载值(load level) 内存使用；程序性能排查: GC 垃圾回收问题

#### 网络

- 网络 I/O 状况: QPS 瞬时过高, TCP 半连接队列已满, TCP Socket 缓冲读写值; 针对 Linux 的 TCP 优化参数

`netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'` 查看网络连接的各个连接状态数量

分析各TCP 状态的连接数目，(SYN 阶段)半连接数目较大，调整 内核半连接队列数目；TIME_WAIT 阶段数目较大，调整 TIME_WAIT 时长

`netstat -na|grep ESTABLISHED|wc -l`: `ESTABLISHED` 状态的数量就是当前并发连接数

`netstat -nat|grep -i "80"|wc -l` 查看 端口80 连接数

- TCP 参数配置: `TCP_NODELAY` 默认是开启，能提高数据传输，减小延时; 若关闭, 适合带宽小，减小数据传输次数，能容忍一定时延效果

- DNS 解析慢引发 (比较隐蔽, 所有问题排查完，针对这因素进行排查)

- 网络攻击: SYN 攻击，占用连接数 (不常见)

- 磁盘空间，因磁盘写满，导致写数据阻塞，句柄报错等信息

#### 高并发服务网络常见问题

- 进程可用 `FD` 不足，句柄报错异常

- IP 端口号不够 与 TIME-WAIT 状态过多，无法及时释放占用IP

- 针对容器网络 `netfilter/conntrack` 相关内核参数配置与 NAT 服务，conntrackTable 不足, 引发丢包问题, 呈现请求超时

### 应用层/代码层面

查看请求的中间件与应用层逻辑，可采用全链路追踪，定位时长较长的链路位置

- 涉及 Redis 缓存，排查是否慢查询，或其他服务存在慢查询，阻塞请求

- 涉及 MySQl 存储, 分析 SQL 语句，索引优化，是否有慢查询日志，查询数据量是否过大 I/O 问题

代码写法：

1. 锁阻塞；资源没及时释放

2. 线程池/连接池, 连接释放，失活等

3. 中间件相关配置

4. 日志log 过多，log 写阻塞

## 故障排查手段

### 基本步骤

#### 查看各个指标监控

- sentry 项目检测
- k8s 重启次数
- Grafana 性能指标
- 数据库 CPU&内存指标
- 慢 SQL 检测

#### 代码工程与项目状态排查，

- 针对代码或者项目进行分析，确认最近的修改

- 拉通相关人员进行一起分析

### 常见排查工具分析

[Linux 命令行监控程序](https://www.tecmint.com/command-line-tools-to-monitor-linux-performance/)

`sysclt` 命令：查看机器相关内核配置，如TCP 内核配置等; 同时可使用此命令进行更改内核配置: `sysctl -w`

```
net.ipv4.ip_local_port_range = 32768    60999
net.ipv4.tcp_timestamps = 1
net.ipv4.tcp_tw_reuse = 0
```

ip_local_port_range: 本地端口地址范围，会影响连接数，若端口使用完，会因无法分配连接地址而报错
tcp_tw_reuse: 若开启后，会减小 TIME_WAIT 等待时长，提高地址复用，有效处理瞬时海量请求连接地址问题

`strace` 对系统调用进行追踪，涉及内存、网络与文件处理。

`ss` 获取socket统计信息，优势在于它能够显示更多更详细的有关TCP和连接状态的信息，而且比netstat更快速更高效

`dmesg` 查看内核缓冲区日志信息: `dmesg -T` 针对机器开机日志与网络丢包等异常，都可以进行查看
