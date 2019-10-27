# elasticsearch 查询笔记总结

## multi_match

- 支持对要搜索的字段的名称使用通配符

```sh
{
    "query": {
        "multi_match": {
            "query": "php编程",
            "fields": ["title", "*_name"]
        }
    }
}
```

- 设置指定搜索字段的权重
```sh
{
    "query": {
        "multi_match": {
            "query": "php编程",
            "fields": ["title^3", "description"]
        }
    }
}
```