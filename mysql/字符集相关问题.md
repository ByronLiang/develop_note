# MySQL字符集与大小写敏感问题

背景：注册账号（比如邮箱) 是不因大小写而影响存储数据与查询数据的；主要基于字符集的设定来实现的

## 字符集相关知识

- utf8与utf8mb4（utf8 most bytes 4）

1. MySQL 5.5.3之后增加了utfmb4字符编码。
2. 支持BMP（Basic Multilingual Plane，基本多文种平面）和补充字符 最多使用四个字节存储字符

- utf8mb4是utf8的超集并完全兼容utf8，能够用四个字节存储更多的字符。

1. 标准的UTF-8字符集编码是可以使用1-4个字节去编码21位字符，这几乎包含了世界上所有能看见的语言
2. MySQL里面实现的utf8最长使用3个字符，包含了大多数字符但并不是所有。例如emoji和一些不常用的汉字，如“墅”，这些需要四个字节才能编码的就不支持。

- utf8mb4对应的排序字符集有utf8mb4_unicode_ci、utf8mb4_general_ci

### utf8mb4_unicode_ci和utf8mb4_general_ci的对比准确性

- utf8mb4_unicode_ci是基于标准的Unicode来排序和比较，能够在各种语言之间精确排序
- utf8mb4_general_ci没有实现Unicode排序规则，在遇到某些特殊语言或者字符集，排序结果可能不一致。但是，在绝大多数情况下，这些特殊字符的顺序并不需要那么精确。

## 相关字符集配置

- 表或行的 collation 约定的命名方法如下

1. *_bin: 表示的是 binary case sensitive collation，也就是说是区分大小写的
2. *_cs: case sensitive collation，区分大小写
3. *_ci: case insensitive collation，不区分大小写

- chartset & collation

1. MySQL 查看列表 `show collation` 指令可以查看到mysql所支持的所有COLLATE;

2. Laravel migration 文件配置

```php
$table->string('key')->unique()->charset('utf8')->collation('utf8_bin');
```
