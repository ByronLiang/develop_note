# K8S 相关异常记录

## OOM引发Pod 服务重启

### 背景

- 因超出容器设置的limit memory, 引发熔断，并向客户端响应503; 并对pod重启

- 针对大分辨率尺寸图片，5000x3000 (4K图片) 会瞬时向机器申请30M 左右内存, 同时可能引发`image/jpeg.(*decoder).processSOS`申请更多内存。

- 当对内存资源进行回收时，消耗CPU, 多次触发强制GC (force GC) 情况。高频率地图片进行编码与解码操作，会极度消耗机器资源，当未能及时GC, 会引发机器OOM 内存过载

### 图片编码解码

查看程序瞬时内存分配

```go
var m runtime.MemStats
runtime.ReadMemStats(&m)
mer := m.Alloc / 1024
fmt.Printf("%d Kb\n", m.Alloc / 1024)
```

[图片编码异常](https://stackoverflow.com/questions/62846156/image-decode-unknown-format)

[图片编码解码](https://github.com/golang/go/issues/10532)

#### 内存占用原理

1. `image.NewYCbCr` 函数: 将字节流转码成`YCbCr` 采样成图片对象, 从而满足SDK里调用图片处理的方法。而图片的分辨率(width, height) 将影响采样的数量，从而对占用存放空间

2. `image/jpeg.(*decoder).processSOS` 当识别标志位，会触发执行, 同样消耗一定的内存资源
