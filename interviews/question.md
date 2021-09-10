# 常见问题

## 范围广，涉及操作系统、网络

1. 浏览器输入网址, 涉及相关知识

### SQL语句的执行流程

1. 传统的C/S请求; 客服端请求MySQL服务器; 涉及网络请求(连接池)

2. MySQl对语句的解析(select语句/insert等), 解析树；查询是否命中缓存

3. 优化器对语句处理；缓存缓冲区查询/处理索引, 定位磁盘页，取出数据

## 并发问题

1. 针对并发情况，若不使用锁，则使用CAS，确保变更前的值没有发生变更
