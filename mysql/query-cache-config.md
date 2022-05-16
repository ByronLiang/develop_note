# Query Cache

1. Query Cache 作用于整个MySQL实例，主要用于缓存ResultSet。
2. 当满足Query Cache的要求，MySQL会直接根据预先设定好的HASH算法将收到的select语句以字符串方式进行hash；然后到Query Cache中查找是否已经缓存。若已经有结果在缓存里，直接取出结果，无需再进行SQL语句解析与存储引擎请求数据等流程，极大提高性能。


## 配置&关键参数

- query_cache_size 缓存ResultSet的内存大小；query_cache_type 何种场景使用缓存
- 获取Query Cache 的配置参数与检验合理性

```sql
show variables like '%query_cache%';
show status like 'Qcache%';
```
query_cache_size设置合理性判断：
- Qcache inserts 
- Qcache hits 
- Qcache lowmem prunes
- Qcache free blocks