# Swoole框架类与WebSocket协议的部署笔记

## Swoole的启动

一般，会使用`supervisor`守护swoole进程, Http/Websocket服务会在指定端口进行开放

## Nginx配置

### 配置upstream

- 通过配置反向代理，将请求转向Swoole的服务端口；nginx应用里，反向代理的模块名称`swoole_dev`不能重复
- 反向代理都是以Http请求其服务地址

```sh
upstream swoole_dev {
    # 通过 IP:Port 连接
    server 127.0.0.1:5200 weight=5 max_fails=3 fail_timeout=30s;
    keepalive 16;
}
```

对/api/的路径都使用路径别名为`proxyApi`

```sh
location ~ ^/api {
    try_files $uri @proxyApi;
}
```

### 不基于upstream配置转发

```sh
location ~ ^/api {
    # 转发请求地址
    proxy_pass http://127.0.0.1:5200;
    include proxy_params;
    # 使用HTTP 1.1
    proxy_http_version 1.1;
    # 开启 websocket
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "Upgrade";
}
```

配置路径别名`proxyApi`

```sh
location @proxyApi {
    # http版本
    proxy_http_version 1.1;

    proxy_set_header Connection "";
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Real-PORT $remote_port;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Host $http_host;
    proxy_set_header Scheme $scheme;
    proxy_set_header Server-Protocol $server_protocol;
    proxy_set_header Server-Name $server_name;
    proxy_set_header Server-Addr $server_addr;
    proxy_set_header Server-Port $server_port;
    
    # 反向代理 寻找upstream模块里 匹配的名称
    proxy_pass http://swoole_dev;
}
```

## WebSocket配置

### map指令作用

- 为实现对`http_upgrade` `connection_upgrade`参数进行值设置，设置的规则在`{}`进行配置
- 可以将变量组合成为新的变量，会根据客户端传来的连接中是否带有Upgrade头来决定是否给源站传递Connection头，这样做的方法比直接全部传递upgrade更加优雅。

```sh
map $http_upgrade $connection_upgrade {
    default upgrade;
    ''      close;
}
```

### 路径配置

- 指定地址路径处理websocket请求
- http与websocket请求差别：

1. WebSocket协议里的Upgrade`[proxy_set_header Upgrade $http_upgrade]` 需使用websocket
2. WebSocket协议里Connection`[proxy_set_header Connection]` 会使用`Upgrade`

```sh
location =/ws {
    # proxy_connect_timeout 60s;
    # proxy_send_timeout 60s;
    # proxy_read_timeout：如果60秒内被代理的服务器没有响应数据给Nginx，那么Nginx会关闭当前连接；同时，Swoole的心跳设置也会影响连接的关闭
    # proxy_read_timeout 60s;

    # 与普通http请求一致
    proxy_http_version 1.1;

    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Real-PORT $remote_port;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Host $http_host;
    proxy_set_header Scheme $scheme;
    proxy_set_header Server-Protocol $server_protocol;
    proxy_set_header Server-Name $server_name;
    proxy_set_header Server-Addr $server_addr;
    proxy_set_header Server-Port $server_port;

    # 与普通http请求的差别
    proxy_set_header Upgrade $http_upgrade;
    # 与普通http请求的差别
    proxy_set_header Connection $connection_upgrade;
    proxy_pass http://swoole_dev;
}
```

### proxy_http_version 参数

若不设置`proxy_http_version` 默认是使用`HTTP 1.0`

在HTTP 1.0 需要设置请求head `Connection: keep-alive`; 

HTTP 1.1 默认使用keepalive 设置请求head使用 `Connection: close` 能实现短连接

基于Swoole的WebSocket服务, 其底层连接依赖HTTP1.1的keepalive特性；若使用HTTP1.0 引发连接断开情况

#### Nginx keepalive_timeout

`keepalive_timeout` 若设置为0 则全部请求都是短连接; 

若设定大于0的数值(单位:s[秒]), 则决定此长连接的存活时长
