# 定时器应用

## 避免使用 time.After

`time.After(duration)` 定时函数, 每完成一个定时周期，都会重新初始化对象，会不断申请与回收时间内存

## time.Ticker 与 time.Reset 定时周期

ticker: 当满足定时，立即开启下一个定时周期，并执行定时任务。执行任务所耗的时间已经被计算进下一个定时周期内。
因此，当周期执行任务所耗时间大于定时长度, 每当定时任务结束后，将立即进入下一周期的任务

### time.Reset 定时模型

每次执行完定时任务, 重置计时，保证执行任务所耗时间不占用定时时长

适合高精细度定时(500ms)、执行定时任务耗时随机性较大，能确保每个定时任务严格按照相隔指定时间长度

```go
timer := time.NewTimer(duration)
// check timer is stoped
if !timer.Stop() {
    <-timer.C
}
for {
    // reset time, start to next round
    timer.Reset(duration)
    select {
		case <-timer.C:
            // process
	}
}
```
