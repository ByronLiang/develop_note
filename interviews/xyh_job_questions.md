# Questions

- 写一下这条SQL的对应Mongo Query：

``` SQL

SELECT * FROM users WHERE name!="kingmax" and name!="soul"

```

- Answer :

```php

$db->users->find(['name' => ['$nin' => ['kingmax', 'soul’]]]);
```