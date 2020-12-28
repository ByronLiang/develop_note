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

## 切片append操作线程安全性

现象: slice在并发执行中`append操作`不会报错，但是数据会丢失

原因: 每次append操作不会强行指向新的内存地址；当多个go协程对共享切片进行append，部分数据会因竞争失败而产生丢失

## map并发读写异常

异常: `concurrent map read and map write`

解决: 不要做map的并发，如果用并发要加锁，保证map的操作要么读，要么写
