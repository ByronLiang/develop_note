# IM网关系统平台设计与总结

## 平台功能

### 开关流量功能，通过配置开启拒绝应用接入

##### 限流开启功能

1.配置开启灰度开关
2.可以设置灰度百分比
3.有一定概率拒绝应用内的用户接入
4.令牌桶算法

#### 广播区分appid

1. 应用内可以支持区分不同渠道包进行推送
2. 梳理开发协议 优化协议; 如平台内部的RPC
3. 支持函数式闭包条件

#### 推送特点

1. 支持开发使用复杂的条件进行推送
2. 支持开发对使用函数作为条件判断推送
3. 支持针对客户端不同机型不同版本进行推送
4. 支持不同国家版本进行不同语言文案推送
5. 性能测试 benchmark测试

#### 目标

1. 支持大规模应用集群, 可扩展添加部署服务器, 支持大量在线用户
2. 不同项目与项目之间不造成影响，又可以互通
3. 相关数据采集；已推送消息数量等
