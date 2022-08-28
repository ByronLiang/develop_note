# 日常开发总结

## 索引

### 部分索引 (Partial Indexes)

只为集合中满足指定筛选条件的文档创建索引

比如订单列表: 针对未支付、待发货与售后等状态订单建立索引。减少其余订单状态建立索引，起到优化磁盘空间等作用

#### 注意

当部分索引具有唯一约束，只有满足其部分索引的条件`(partialFilterExpression)`, 约束条件才其阻止作用

针对订单状态不小于2的枚举值，才能对订单号起到唯一性约束

```
db.orders.createIndex(
    { no: 1 },
    { unique: true, partialFilterExpression: { status: { $gte: 2 } } }
)
```
