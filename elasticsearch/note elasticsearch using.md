# elasticsearch 基本配置要点与使用笔记

## type类型的移除问题

- elasticsearch 由 `index` `type` `document`组成；但逐渐会移除type, 以一个索引一个type; 并且以每个数据表为一个index作为索引数据

- 去掉type能够使数据存储在独立的index中，这样即使有相同的字段名称也不会出现冲突;去掉type就是为了提高ES处理数据的效率。

## 查询基本要义与表达式

1. QUERY_NAME 查询表达式 (query, filters)
2. ARGUMENT 搜索类型 (match, multi_match_query)
3. Value 查询值

```sh
{
    QUERY_NAME: {
        ARGUMENT: VALUE,
        ARGUMENT: VALUE,...
    }
}
```

- 针对`FIELD_NAME`字段作出的搜索查询
```sh
{
    QUERY_NAME: {
        FIELD_NAME: {
            ARGUMENT: VALUE,
            ARGUMENT: VALUE,...
        }
    }
}
```


