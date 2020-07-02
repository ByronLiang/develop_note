# Nginx Cors Config

Nginx 通过 `add header` 添加跨域配置; 可指定 `location` / `server`

## 基本配置

```sh
add_header Access-Control-Allow-Origin "$http_origin" always;

add_header Access-Control-Allow-Methods 'HEAD, GET, POST, DELETE, PUT, OPTIONS' always;

add_header Access-Control-Allow-Headers 'DNT,Authorization,User-Agent,X-XSRF-TOKEN,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,X-APP-KEY,X-APP-SIGN' always;

add_header Access-Control-Allow-Credentials 'true' always;

add_header Access-Control-Max-Age 1728000 always;

if ($request_method = 'OPTIONS') {
    return 204;
}
```

- `$http_origin` 请求域名路径（唯一请求源进行跨域处理）协议 + 域名 + 端口

- `add_header ...(header配置) always` 不添加`always` 若发送50x内部错误, 无法加载其头部信息

### 预检请求 preflight request

浏览器必须首先使用 OPTIONS 方法发起一个预检请求（preflight request），从而获知服务端是否允许该跨域请求。服务器确认允许之后，才发起实际的 HTTP 请求

```sh
if ($request_method = 'OPTIONS') {
    return 204;
}
```

识别`options`请求 直接返回`204状态码` (没有实体内容, 即便返回数据，响应主体也会进行忽略, 减少不必要的数据传输), 因此, 预检请求不会对服务器资源产生性能影响

