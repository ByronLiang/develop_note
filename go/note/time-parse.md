# 时间解析开发注意点

## 时间字符串解析

`time.Parse()` 是以`UTC`时区解析时间; `time.Now()` 是以服务器设置的timezone 返回时间对象

由于时区不一致, 当进行时间比较(`Before()/After()`)，时间差(`Sub()`)值计算, 时区会引发异常

当对日期时间字符串解析成时间对象, 建议使用`time.ParseInLocation`, 使用`time.Location`服务器设置的timezone

## 时间戳解析

若涉及多时间地区, 每次解析时间对象，都需设置指定时区, 确保能正确显示用户所在地区的时间

`time.Unix(sec, nsec int64) Time`是按照服务器设置的timezone 来解析时间对象

`(t Time) In(loc *Location) Time` 设置时间对象的时区, 并返回所设置时区的时间对象

```go
l, err := time.LoadLocation("Asia/Tokyo")
if err != nil {
    return err
}
t := time.Unix(1637723229, 0)
// 设置时区
newT := t.In(l)
// 转换时间戳为东京时区的日期时间
fmt.Println(newT.Format("2006-01-02 15:04:05"))
```

## 开发建议

1. 按照国家维度，分表存储数据，确保按照指定国家地区解析日期时间

2. 日期时间字段存放时间戳，或者日期时间字符串带上时区: `2006-01-02 15:04:05 +0800 CST` 避免解析错误

3. 前端调接口时, 通过header 传递当前用户所在地区，服务端解析所属时区，使其能正确解析日期时间字段数据
