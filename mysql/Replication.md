# MySQL 主从复制 与 binlog基本配置

## binlog理论

MySQL的二进制日志可以说是MySQL最重要的日志了，它记录了所有的DDL和DML(除了数据查询语句)语句，以事件形式记录，还包含语句所执行的消耗的时间，MySQL的二进制日志是事务安全型的

### 作用

- MySQL Replication在Master端开启binlog，Mster把它的二进制日志传递给slaves来达到master-slave数据一致的目的

- 通过使用mysqlbinlog工具来使恢复数据

### 文件特点

二进制日志包括两类文件:
    
- 二进制日志索引文件（mysql-bin.index）用于记录所有的二进制文件;
- 二进制日志文件（mysql-bin.00000*）记录数据库所有的DDL和DML(除了数据查询语句)语句事件。

## binlog配置

在`/etc/mysql/my.cnf` 的`[mysqld]` 区块进行相关配置: 

1. `log-bin=mysql-bin` [`mysql-bin` 是日志的基本名或前缀名)]; 
    - 若只配置日志的基本名称，则log-bin 的存放路径则为`datadir= /var/lib/mysql`处
    - `log-bin=/var/log/mysql_bin_log/mysql-bin` 配置绝对路径，则日志文件会存放指定路径的文件夹里；并且此文件夹的权限需要属于mysql的用户与用户组

2. `binlog_format=row` 配置日志格式；
3. `binlog-do-db` 与 `binlog-ignore-db` 指定对哪些数据表进行日志记录 和 不对指定的数据表进行日志记录

## binlog 查看方式

1. 使用mysqlbinlog自带查看命令法：`mysqlbinlog /var/lib/mysql/mysql-bin.000013`
2. 进入MySQL命令行, `show binlog events [IN 'log_name'] [FROM pos] [LIMIT [offset,] row_count];`

```sh
选项解析：
IN 'log_name'   指定要查询的binlog文件名(不指定就是第一个binlog文件)
FROM pos        指定从哪个pos起始点开始查起(不指定就是从整个文件首个pos点开始算)
LIMIT [offset,] 偏移量(不指定就是0)
row_count       查询总条数(不指定就是所有行)
```

## MySQl主从复制

### 配置

- 在`/etc/mysql/my.cnf` 的`[mysqld]` 区块进行相关配置: 

1. 为每个MySQL实例配置`server-id = ` id 数值需要唯一;

- 对需要作为从机实例进行配置:

1. `replicate_do_db` 配置对某数据库进行复制；`replicate_ignore_db` 指定数据库不进行复制
2. `replicate_wild_ignore_table = targetdb.job` 指定数据库里的表进行不复制 (若以数据库作为队列驱动，不对队列表进行主从复制)

- 需进入从机实例的命令行进行配置:

1. 配置从机实例的主机地址信息: `change master to` 命令

```sh
master_host='192.168.137.26'       远程主(Master)机IP      
master_user='phpslave',            登录远程主(Master)机 用户名
master_password='789',             登录远程主(Master)机 密码      
MASTER_LOG_FILE='mysql-bin.000009' 指定从哪个binlog文件复制(如果不加此参数默认是最早的binlog日志)
MASTER_LOG_POS=107                 指定从哪个binlog文件的哪个pos点开始复制(如果不加此参数默认是最早的pos点)
MASTER_PORT=3306,                  远程主(Master)机端口
MASTER_HEARTBEAT_PERIOD=60         多长时间探测一次主服务器是否在线 单位：秒
MASTER_CONNECT_RETRY=10;           无法连接主服务器的时候重试连接的间隔时间 单位：秒
```

2. 查看从机状况 `show slave status\G;`
3. 启动/关闭: `start slave` / `stop slave`

### 原理

复制类型: 同步复制和异步复制，实际复制架构中大部分为异步复制。

#### 主从同步是采用Push方式

主库上数据更新时，将消息推送给从库，然后从库再进行更新(Push方式)；

从库若通过轮询的方式（Pull方式）不断查询主库是否更新了数据，会耗费资源；

#### 复制的基本过程 

1. Slave上面的IO进程连接上Master，并请求从指定日志文件的指定位置（或者从最开始的日志）之后的日志内容；  

2. Master接收到来自Slave的IO进程的请求后，通过负责复制的IO进程根据请求信息读取指定日志指定位置之后的日志信息，返回给Slave的IO进程。返回信息中除了日志所包含的信息之外，还包括本次返回的信息已经到Master端的bin-log文件的名称以及bin-log的pos位置；  

3. Slave的IO进程接收到信息后，将接收到的日志内容依次添加到Slave端的中继日志(relay-log)文件的最末端，并将读取到的Master端的bin-log的文件名和位置记录到master.info文件中，以便在下一次读取的时候能够清楚的告诉Master“我需要从某个bin-log的哪个pos位置开始往后的日志内容，请发给我”；

4. Slave的sql进程检测到relay-log中新增加了内容后，会马上解析relay-log的内容成为在Master端真实执行时候的那些可执行的内容，并在自身执行。

### 应用

1. 对应用进行配置: 写请求向MySQL主机实例；读请求向MySQL的从机实例；
2. 利用MySQL中间件处理请求，自动分发请求到相关实例；常用开源架构: `MyCat` `MHA`

#### 主从不一致优化方案

引入缓存组件 标记查询对象是否最近进行数据操作

- 写操作设置一个key(`db:table:主键`) 记录缓存里(cache) 并设置有效时长, 如1秒左右 作为主从延迟的时间窗

- 读操作：在缓存进行查询key(`db:table:主键`), 若key存在 处于时间窗内, 为保证准确, 对主库进行读操作; 否则, 直接对从库进行读操作

