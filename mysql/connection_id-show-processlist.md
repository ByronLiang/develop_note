## connection_id() and show processlist 的函数使用

### `connection_id()`
- 返回连接的连接标识（线程标识）。 每个连接都有一组在当前连接的客户端中唯一的ID
- 返回的值与INFORMATION_SCHEMA.PROCESSLIST表的ID列，SHOW PROCESSLIST输出的Id列和Performance Schema线程表的PROCESSLIST_ID列中显示的值类型相同。

#### How to use
原生SQL语句: <br>
```
select connection_id() as connection_id;
```

Laravel: <br>
```
\DB::select('select CONNECTION_ID() as connection_id');
```
<br>

### `show processlist`
显示连接MySQL服务的线程相关信息；包括connection_id与状态<br>
- 查询内容显示如下:<br>

"id" "User" "host" "DB" "Command" "Time" "State" <br>
"2"	"root"	"localhost:51538"	"dreamer"	"Query"	"0" "starting"	"SHOW PROCESSLIST"<br>

"3"	"root"	"localhost:51561"	"dreamer"	"Sleep"	"891"	""

<br>

- MySQL文档注释:<br>
The SHOW PROCESSLIST statement is very useful if you get the “too many connections” error message
and want to find out what is going on. MySQL reserves one extra connection to be used by accounts that
have the SUPER privilege, to ensure that administrators should always be able to connect and check the
system (assuming that you are not giving this privilege to all your users).

Threads can be killed with the KILL statement.



#### Using sample in develop
- 可以有效避免高并发下对数据进行抢占处理；主要可用于队列数据的处理
- 避免重复处理数据与线程出现阻塞情况
<br>
- 标注线程标识<br>
当消费者处理队列数据时，将当期线程标识(connection_id)存储在未被消费数据的标识处，能有效避免消费者重复处理数据；<br>
1.先使用connection_id进行数据update.<br>
2.利用当期connection_id进行select数据，再将数据进行相关业务处理
<br>

- show processlist得出有效线程<br>


当部分线程因意外被中断，无法消费数据，需将对应的线程ID的消费数据进行复位处理
<br>

#### Get Idea from sample
- 减少使用 `select for update` 无需查询哪些记录未被处理，只需将被更新了connection_id 的数据进行处理即可;<br>
- 尽可能完成需要做的事情。尽量使用update 代替 先select for update 再 update 的写法。<br>事务提交的速度越快，持有的锁时间越短，可以大大减少竞争和加速串行执行效率。