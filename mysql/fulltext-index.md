# MySQL 全文索引使用记录

全文索引 (FullText Index) 在 5.6 版本前，只能使用 MyISAM 驱动，无法使用在 InnoDB 驱动上。

在 5.6 版本后，支持全文索引，并且使用自然语言插件(ngram)，满足中文分词，使查询更加高效。

## 使用

1. 使用 where like "%%" 查询，无法使用全文索引。只能使用 match (字段名,字段名) against ("查询内容" 添加匹配模式)

备注：

1.1 `match (字段名)` 需要使用全文索引的字段，否则，异常: `Can't find FULLTEXT index matching the column list` 

1.2 若是多个字段组成的联合全文索引，则需要使用`match (多个字段名[使用逗号分隔])`

2. 匹配模式默认是 `IN NATURAL LANGUAGE MODE`, `IN BOOLEAN MODE`; 扩展：`WITH QUERY EXPANSION` (基于匹配内容，进行扩展查询数据)

3. 全文索引创建自然语言(汉字)，`CREATE FULLTEXT INDEX 索引名 ON 表名 (字段名) WITH PARSER ngram;` 对于后期添加全文索引，需要执行`optimize table`, 对当前表数据内容进行分词处理。

## 优化与排查

### optimize 命令

`optimize table 名` 优化全文索引数据表，重新整理倒排索引，清除已删除内容的倒排索引数据

### INFORMATION_SCHEMA 表

可以查询数据表，获取相关数据库运行情况，如 分词缓存表(`INNODB_FT_INDEX_TABLE`), 全文索引删除文档ID(`INNODB_FT_DELETE`)

#### 查询步骤

1. `SET GLOBAL innodb_ft_aux_table = 'database_name/table_name';` 配置全文索引分词指定的数据表名 

2. 先配置 `innodb_ft_aux_table`，后触发对其数据表进行写入数据, 再进行查询分词缓存数据: `SELECT * FROM INFORMATION_SCHEMA.INNODB_FT_INDEX_CACHE;` 
