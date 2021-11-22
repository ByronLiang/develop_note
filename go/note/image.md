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

### 图片移除EXIF信息与图片编码解码引发丢失图片旋转信息

1. `jpeg.Encode` 将image 对象编码，写入文件或者生成字节流，会丢失图片旋转信息, 导致与原图的旋转角度不相同
2. 从图片字节数据里移除 EXIF 信息字节数据，即使不进行图片编码处理，将图片字节上传到云存储，下载图片因丢失图片旋转信息，导致与原图不一致

#### 解决方案

解析EXIF数据: `github.com/rwcarlsen/goexif/exif`

设置图片旋转: `github.com/disintegration/imagin` `imaging.Rotate180` 等方法

1. 从图片流里提取原图片旋转信息，当图片编码时，先设置图片旋转信息，再继续完成图片编码

2. 移除 EXIF 信息数据时，保留图片旋转的信息，并且不进行图片编码处理

### EXIF IFD 字节信息

#### 读取字节顺序

EXIF字节数据里, 前两个字节数据是字节的读取顺序

读取字节顺序由两种类型: `Little-endian` 与 `Big-endian`

- Big-endian：高位字节存入低地址，低位字节存入高地址

- Little-endian：低位字节存入低地址，高位字节存入高地址

#### IFD 标签信息

从第八个字节开始，将正式读取标签信息：

- 标签总数目(2bit)

- 每8bit包含标签基本信息: 标签ID, 标签类型, 标签数据长度与指定偏移位置读取内容

