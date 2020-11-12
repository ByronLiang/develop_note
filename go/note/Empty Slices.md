# 空切片声明

```go
// 初始长度为0 nil切片
var t []string

// 初始长度为0 非nil切片
// 类比: var tt []string = make([]string, 0)
tt := []string{}
```

- 当声明空数组时，推荐使用第一种方法; 当在JSON编码时，推荐的是后两种方式

- 一个nil空数组会被编码为null，但非nil空数组会被编码为JSON array []
