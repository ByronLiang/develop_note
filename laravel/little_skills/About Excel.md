# 涉及导出&导入数据相关总结

## 避免导出的数值超出11位长度被强制转换成科学计数法显示

- 处理：先判断导出数据是否属于数字类型，数字类型则加上制表符'\t'，使数字变为字符串

```php
if (is_numeric($val)) {
    $val = $val.'\t';
}
```

- Laravel Excel包的处理方案

对指定的列进行格式设定，判断其属于数字数据结构，则使用内置的`DataType`进行转换处理

参考：[Custom Formatting Values](https://docs.laravel-excel.com/3.1/exports/custom-formatting-values.html)
