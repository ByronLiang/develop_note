# 手机号和短信验证码登录需求及代码实现


阅读此篇文章的相关总结

[https://insights.thoughtworks.cn/sms-authentication-login-api/](https://insights.thoughtworks.cn/sms-authentication-login-api/ "Phone Message API Develop")

### 短信验证码有效期为2分钟

- 需求理解

避免用户因手机信号不好，而无法收到短信.

- Laravel & Redis代码实现

`\Cache::put('156xxx', '123456', 2);`

`\Redis::setex('156xxx', 120, '123456');`

### 验证码为6位纯数字

- 需求理解

验证码长度只有4位而且还是纯数字，黑客来个多线程并发请求，或者拿一个集群来暴力登录，有可能会赶在有效期内破解出合法的验证码；

### 验证码，至多可被使用3次

- 需求理解

保存于服务器端的验证码，至多可被使用3次（无论和请求中的验证码是否匹配），随后立即作废，以防止暴力攻击

- Laravel and Redis代码实现


每次请求验证码接口，记录使用次数


`\Redis::incr('156xxx_count');`  作为统计的key可以复杂点；

当返回数值大于3且验证码错误/验证码正确: 使验证码失效且清除使用次数的统计

`\Redis::del('156xxx');` 

`\Redis::del('156xxx_count');`

### 集成第三方API做登录保护 & 设置图形验证码

- 第三方 `Google firebase` 或者 APP可以直接走移动供应商手机验证