# phpunit 使用注意事项

## 单元测试配置文件

- phpunit.xml文件的配置

核心配置文件

```[php]
<php>
    <env name="APP_ENV" value="testing"/>
    <env name="CACHE_DRIVER" value="array"/>
    <env name="SESSION_DRIVER" value="array"/>
    <env name="QUEUE_DRIVER" value="sync"/>
</php>
```

可以对单元测试文件的缓存驱动、Seesion驱动与队列驱动进行设置；

当进行测试时，尽量将单元测试环境配置成生产环境的对应配置

- 问题关注点

当使用缓存驱动为array时，同一环境下，不同的单元测试文件会出现无法获取缓存内容的问题；

因此，尽量使用默认的缓存驱动来进行相关缓存测试