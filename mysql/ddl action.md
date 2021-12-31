# DDL问题 与 MySQL、MongoDB 选型思考

## 选型思考

MongoDB 是NOSQL类型数据库 文档型数据库，无需设置表结构, 表数据能动态增减字段

MySQL 对于大表, 线上对表结构进行相关操作，会引发一段时间的表无法写入操作 

### DDL 问题

DDL行为: 对表新增/删除字段与索引

一般使用`online ddl`解决: pt-online-schema-chang
