# Questions

- 写一下这条SQL的对应Mongo Query：

``` SQL
SELECT * FROM users WHERE name!="kingmax" and name!="soul"
```

- Answer :

```php
$db->users->find(['name' => ['$nin' => ['kingmax', 'soul’]]]);
```

- 把以下符合美元价格格式的字符串匹配并可从一段字符串中提取出来："$200.49"、"$1,999.00"、"$99"、"50.00美元"

- Answer：

`(\d+\.?\d+)` 提取小数

`|` 条件分隔符；可以在美元符号前与后进行相关匹配

```php
preg_match('/[\$|美元](\d+\.?\d+)|(\d+\.?\d+)[\$|美元]/', '美元200', $money);

echo last($money);
```
