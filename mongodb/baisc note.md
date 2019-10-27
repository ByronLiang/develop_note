# MongoDB 基本笔记

1. MongoDB 是一种典型的非关系数据库（NOSQL）与传统的关系数据库（RDBMS）有数据关系差别
2. MongoDB以数据库、集合与文档结构组成，灵活变更文档字段及索引
3. MongoDB的数据文档存储使用BSON，有效描述非结构化数据和结构化数据

## BSON相关概念

- 具有轻量、 可遍历与高效

1. 一种类JSON的二进制的存储格式；支持内嵌文档对象与数组对象；
2. 具备JSON没有的数据类型，Date、 BinData
3. 遍历原理更高效：会将JSON的每一个元素的长度存在元素的头部，只需要读取到元素长度就能直接seek到指定的点上进行读取了；

## PHP驱动配置问题

- Q: 在PHP7中, MongoDB属于长连接。一个请求完，对应得 php-fpm 进程没有被 kill 掉化，这个连接不会断开，会一直保持。所以要减少连接个数，需要合理设置 php-fpm 空闲进程数

- A: 配置php-fpm 限制最大连接数与空闲数的进程占用：`pm.min_spare_servers` & `pm.max_spare_servers` [php-fpm.d/www.conf文件]
