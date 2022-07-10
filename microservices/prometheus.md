# Prometheus 记录

## 概念

1. 使用 TSDB 存储以时序维度的监控数据, 并通过查询语句，显示监控内容

2. 以 `node_` 开头的监控指标, 其数据来源于节点采集器`node_exporter`

3. 由 指标采集程序(metric) 和 指标时序服务(prometheus server) 组成。使用`拉数据(pull)`方式, 调取指标采集程序的配置路径，获取指标数据，从而存储在其 TSDB, 以供后期数据查询

4. job: 任务；instance: 实例、targets 一个job 能对应多个实例(targets)

## 开发问题

### 虚拟机无法采集电池服务属性

针对使用 `node_exporter` 于虚拟机采集机器指标, 会有异常日志：`could not get power_supply class info: error obtaining power_supply class info: failed to read file \"/sys/class/power_supply/BAT0/power_now\"`

会因无法识别 `device` 而无法读取当前机器电池指标 

### 指标采集客户端注册器

[官方客户端 SDK](github.com/prometheus/client_golang)

默认的注册器会进行采集`processCollector` 会当前进程的 FD 目录`(/proc/pid/fd)`进行遍历，若进程拥有大量连接时，句柄文件较多，会引发采集超时，并影响其余指标上报数据，从而使数据查询时，出现数据缺失情况

```go
func init() {
	MustRegister(NewProcessCollector(ProcessCollectorOpts{}))
	MustRegister(NewGoCollector())
}
```

若不希望采集无关指标, 使用新的注册器，减少不必要的指标采集

### 指标采集客户端数据采集行为

若应用层面里的指标采集程序，指标标签、指标数据过长, 会引发 TSDB 写入阻塞，大量占用内存等问题

```sh
level=info ts=2019-07-21T05:10:05.029Z caller=main.go:668 msg="TSDB started"
level=info ts=2019-07-21T05:10:05.040Z caller=main.go:738 msg="Loading configuration file" filename=/etc/prometheus/prometheus.yml
level=info ts=2019-07-21T05:10:05.060Z caller=main.go:766 msg="Completed loading of configuration file" filename=/etc/prometheus/prometheus.yml
level=info ts=2019-07-21T05:10:05.060Z caller=main.go:621 msg="Server is ready to receive web requests."
level=info ts=2019-07-21T05:10:22.567Z caller=compact.go:495 component=tsdb msg="write block" mint=1563674400000 maxt=1563681600000 ulid=01DG9EZZ2R2GNAB73MF3TF85ED duration=9.230158995s
level=info ts=2019-07-21T05:10:24.785Z caller=head.go:586 component=tsdb msg="head GC completed" duration=742.465478ms
level=info ts=2019-07-21T05:10:40.581Z caller=head.go:656 component=tsdb msg="WAL checkpoint complete" first=42 last=45 duration=15.795350131s
level=warn ts=2019-07-21T06:33:54.405Z caller=scrape.go:952 component="scrape manager" scrape_pool=game target=http://172.17.125.122:3301/metrics msg="appending scrape report failed" err="out of bounds"
level=info ts=2019-07-21T07:30:34.769Z caller=compact.go:495 component=tsdb msg="write block" mint=1563681600000 maxt=1563688800000 ulid=01DG9P39F74KY4H1E8CB711G5G duration=16m12.466554245s
```
### 指标时序服务相关优化

1. 增加从指标采集程序拉取指标数据的时间间隔，减少 TSDB 数据写入频次

2. 控制指标数据的标签等数据长度, 使其具有意义
