# 图像处理总结

## 常见异常

[图片编码异常](https://stackoverflow.com/questions/62846156/image-decode-unknown-format)

[图片编码解码](https://github.com/golang/go/issues/10532)

### 内存占用原理

一般，4k分辨率图片，约耗30MB-100MB内存, 并且对GC 也会产生一定压力。服务端应减小处理高分辨图片的业务

1. `image.NewYCbCr` 函数: 将字节流转码成`YCbCr` 采样成图片对象, 从而满足SDK里调用图片处理的方法。而图片的分辨率(width, height) 将影响采样的数量，从而对占用存放空间

2. `image/jpeg.(*decoder).processSOS` 当识别标志位，会触发执行, 同样消耗一定的内存资源

### 引入必要的图片类型包

`image.Decode()` 将字节码转换图片对象, 需要引入相应的图片类型包，否则，将引发`unknown format`

### 获取图片字节码的基本信息

`image.DecodeConfig()` 能将字节码转换基本图片配置对象。
若只需获取图片的宽高，图片类型，可以使用此函数，无需转码成整个图片对象`(image.Decode())`，减少使用内存, 至少节省4倍内存

一般情况下，`image.DecodeConfig()` 与 `image.Decode()` 很少两者同时使用。后者编码后能获取图片的基本信息，无需重复获取

