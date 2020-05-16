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

- 方法名称是publics，作用是以数组形式返回这个对象里面的所有公开属性及值，而私有属性不返回

```php
trait Publics
{
    // 获取一个类的所有公共属性
    public function publics()
    {
        $public_properties = [];
        $varArray = get_object_vars($this);
        $ref = new \ReflectionClass($this);
        $data = $ref->getProperties(\ReflectionProperty::IS_PUBLIC);
        foreach ($data as $value) {
            if (isset($varArray[$value->getName()])) {
                $public_properties = array_merge(
                    $public_properties, 
                    [$value->getName() => $varArray[$value->getName()]]
                );
            }
        }

        return $public_properties;
    }
}
```
