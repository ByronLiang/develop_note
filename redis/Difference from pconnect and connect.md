# PhpRedis pconnect 与 connect 区别

## 连接类型

### connect 连接

1. 每次完成脚本请求, 自动释放连接`(即使不执行close方法，也会释放Redis连接)`; 不会复用同一连接

### pconnect 长连接

1. 对于首次创建, 则创建一个Redis客户端；或者重用一个通过pconnect已连接就绪的客户端
2. 连接结束后, 不会关闭连接, `进程(php-fpm)`保留一个Redis客户端，以作连接重复使用
3. 只有PHP进程结束, 才会关闭Redis连接

#### 使用注意

1. 当php-fpm进程增加，Redis的长连接同样是增加，无法控制数量；

2. 一旦网络异常导致长连接失效，没有办法自动关闭重新连接，以至于后续请求全部失败，通过重启服务解决

## Laravel框架里Redis客户端连接配置

Redis客户端选用phpredis为驱动

配置参数: `persistent = true` 若配置使用持久连接, 则使用pconnect创建Redis客户端 

```php
protected function establishConnection($client, array $config)
{
    $client->{($config['persistent'] ?? false) === true ? 'pconnect' : 'connect'}(
        $config['host'], $config['port'], Arr::get($config, 'timeout', 0)
    );
}
```
