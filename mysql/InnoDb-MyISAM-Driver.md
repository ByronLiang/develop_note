# MyISAM与InnoDB区别与选择笔记

## 不同引擎的性能与索引等特点

### InnoDB

Innodb支持外键，事务和行级锁；

事务的ACID属性：atomicity, consistent, isolation, durable。

并发事务带来的几个问题：更新丢失，脏读，不可重复读，幻读。

事务隔离级别：未提交读(Read uncommitted)，已提交读(Read committed)，可重复读(Repeatable read)，可序列化(Serializable)

没有设定主键或者非空唯一索引，就会自动生成一个6字节的主键(用户不可见)

自动增长列必须是索引

### MyISAM

1. 只有表锁; 支持FullText全文索引;
2. 保存表的总行数，`select count(*) from table;` 会直接取出出该值
3. 允许没有任何索引和主键的表存在; 索引是保存行地址
4. 自动增长列必须是索引

## 业务场景影响引擎的选择

考虑因素：外键/事务/行级锁；数据操作方式/数据量；全文索引;

1. 外键具有保持数据完整性和一致性的机制，对业务处理有着很好的校验作用;

2. 行级锁，允许多线程访问不同行数据；表锁, 不会引发死锁；

3. 轻量级应用建议使用MyISAM引擎;

