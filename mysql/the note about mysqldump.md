# mysqldump 的日常用法实例

- 导出数据表的数据表及其数据
`mysqldump -q(快速导出配置) -h (链接数据库主机地址) -u(链接数据库用户名) -p(链接数据密码) 数据库名称(db_name) 指定导出的表名(table_name)> 导出SQL文件名`

- 只导出数据表的数据
`mysqldump -q(快速导出配置) -h (链接数据库主机地址) -u(链接数据库用户名) -p(链接数据密码) --compact -t 数据库名称(db_name) 指定导出的表名(table_name) > 导出SQL文件名`