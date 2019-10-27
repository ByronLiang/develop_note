# Nginx URL重写(rewrite)

- rewrite的组要功能是实现RUL地址的重定向。
- Nginx的rewrite功能需要PCRE软件的支持，即通过perl兼容正则表达式语句进行规则匹配的。
- 默认参数编译nginx就会支持rewrite的模块，但是也必须要PCRE的支持

## 语法格式与参考语法

```
rewrite    <regex>    <replacement>    [flag];
关键字      正则        替代内容         flag标记

关键字：其中关键字error_log不能改变
正则：perl兼容正则表达式语句进行规则匹配
替代内容：将正则匹配的内容替换成replacement
flag标记：rewrite支持的flag标记

rewrite参数的标签段位置：
server,location,if

flag标记说明：
last  #本条规则匹配完成后，继续向下匹配新的location URI规则
break  #本条规则匹配完成即终止，不再匹配后面的任何规则
redirect  #返回302临时重定向，浏览器地址会显示跳转后的URL地址
permanent  #返回301永久重定向，浏览器地址栏会显示跳转后的URL地址
```

## 常用涉及rewrite的内置Nginx变量

```
HTTP_USER_AGENT      用户使用的代理，例如浏览器；
HTTP_REFERER         告知服务器，从哪个页面来访问的；
HTTP_COOKIE          客户端缓存，主要用于存储用户名和密码等信息；
HTTP_HOST            匹配服务器ServerName域名；
HTTP_ACCEPT          客户端的浏览器支持的MIME类型；      
REMOTE_ADDR          客户端的IP地址
QUERY_STRING         URL中访问的字符串；
DOCUMENT_ROOT        服务器发布目录；
SERVER_PORT          服务器端口；
SERVER_PROTOCOL      服务器端协议；
TIME_YEAR            年；
TIME_MON             月；
TIME_DAY              日；
```

- 更多例子可参考：[nginx URL重写（rewrite）配置](https://www.jianshu.com/p/a8261a1a64f8).