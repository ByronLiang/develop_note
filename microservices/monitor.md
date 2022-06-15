# 系统监控

## ELK

### elasticalert 告警

[elastalert 插件](https://github.com/Yelp/elastalert)

[基于elastalert 钉钉通知插件](https://github.com/xuyaoqiang/elastalert-dingtalk-plugin)

使用 Python 开发的 ES 插件。独立部署于宿主机，通过配置 ES 主机地址，当日志内容`(output)` 命中配置的规则`(rule)`, 执行 alter 配置方法

#### 自定义 alter 方法

当触发 alter, 向通知平台发送告警信息，通知平台接收到解析告警信息，可基于通知策略，执行相关通知行为。(如钉钉通知: 在某群众，@相关开发人员, 处理相关告警内容)
