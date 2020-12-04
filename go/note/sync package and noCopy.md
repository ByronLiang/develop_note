# Golang sync库函数开发注意 与 noCopy接口

## sync库函数初始化后不建议进行拷贝

```go
type MyMutex struct {
	sync.Mutex
}

var mu MyMutex
mu.Lock()
var mu2 = mu
mu.Unlock()
mu2.Lock()
```
mu处于加锁状态, 变量mu2被拷贝, 锁状态也会被拷贝; `mu2.Lock()` 会一直处于阻塞, 形成死锁;

## noCopy接口

针对拥有noCopy的结构体, 当进行变量拷贝时, 可被`Go vet`检测
