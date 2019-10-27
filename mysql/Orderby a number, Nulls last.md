# 实现指定排序，对未设定的排序进行最后排序处理

实现置顶排序显示内容，则需要对某字段进行赋值处理，而其余数据进行NUll处理；但无论以升序、降序排列数据，Null数据都会排于首位

## sample

```sql
SELECT * FROM tablename WHERE visible=1 ORDER BY position ASC, id DESC
```

The problem with this is that NULL entries for 'position' are treated as 0. Therefore all entries with position as NULL appear before those with 1,2,3,4. eg:

```php
NULL, NULL, NULL, 1, 2, 3, 4
```

the following ordering:

```php
1, 2, 3, 4, NULL, NULL, NULL
```

## solutions

- Place a minus sign (-) before the column name and switch the ASC to DESC

It is essentially the inverse of position DESC placing the NULL values last but otherwise the same as position ASC.

```sql
SELECT * FROM tablename WHERE visible=1 ORDER BY -position DESC, id DESC
```

- make sure data be ordered is `true` use order by asc; If `false` also using field to order asc

```sql
SELECT * FROM table ORDER BY ISNULL(field), field ASC;
```

other sample: one_field == 0 is `true` order by desc. If not  another_field order by desc too

```sql
SELECT * FROM table ORDER BY `one_field` != 0 IS NOT TRUE, `another_field` desc
```

## note

1. 当可销库存为0时，按照销量排序；当可销库存大于0时，按照库存进行排序
2. 内容置顶处理
