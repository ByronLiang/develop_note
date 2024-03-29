# 生成随机数据

从数据库以随机顺序获取指定数量的数据

## 依赖数据库函数生成随机顺序数据

### MySQL / ClickHouse

`order by rand() limit 数据量`: 从一次查询中随机返回指定数据量的数据

#### order by rand() 原理

在数据表里新增虚拟列`(临时表)`，存放生成的随机数，每次依据随机数排序`(文件排序)`，从而每次查询，都能获取随机顺序的数据

### MongoDB

利用聚合函数`(aggregate)`的条件: `$sample: {size: 随机数量}`

场景: 利用管道(pipeline)与聚合函数(aggregate): `db.Test.aggregate([{$match: {"Status": 1}},{$sample: { size: 2 }}]);`

## 利用偏移分块随机到指定偏移数据

获取记录总数，利用随机数据数量，对总数据进行分块处理，并随机从分块数据里取出分块数据。

如：100 条数据总量，获取5条随机顺序数据，会将数据分成20块，从20块随机选取一个数据块，返回数据
