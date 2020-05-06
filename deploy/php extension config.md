# PHP 扩展安装与配置

## 问题

使用nginx搭配php-fpm下，无法加载Redis扩展服务；`use Redis; 引发报错`

## 解决

- 一般按照好扩展服务, 都需在`php.ini`文件配置扩展服务引用; 如`extension = redis.so`;而PHP的配置分为cli与fpm模式下, 它们配置分别位于`/etc/php/7.2/cli/php.ini`与 `/etc/php/7.2/fpm/php.ini`

- 若不在fpm配置下设置其扩展服务，nginx转发php文件的请求内容, 会引发无法识别扩展服务等错误
