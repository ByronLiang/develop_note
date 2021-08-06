# 性能分析与优化

## 采样方式

1. 线上分析：http pprof 提供访问pprof的HTTP接口，获得性能数据；底层也是调用的 runtime/pprof提供的函数，封装成接口对外提供网络访问。
   当进行采集时，通过访问不同地址，能得出相应的性能指标，性能指标数据全面；

2. 运行分析：runtime/pprof; 调用pprof包提供的函数，手动开启性能数据采集，适合对单一性能指标进行采集；可配置采集
   对单独执行的函数与一段程序进行性能分析；可参考封装的库：`pkg/profile`

## pprof工具查看与分析

1. 可使用go tool pprof 命令 指向提供访问pprof的HTTP 或 pprof文件；
如`go tool pprof http://localhost:6060/debug/pprof/profile` 或者 `go tool pprof mem.pprof`

2. 加上 `-sample_index` 参数后，可以切换内存分析的类型：`go tool pprof -sample_index=alloc_space http://localhost:6060/debug/pprof/heap`

### 常用操作

进入到pprof交互界面, 主要操作：

1. `top`查看指标，资源使用占比, 所属函数方法；
   
2. `list 指定函数方法` 查看函数方法资源使用情况

3. `web`对资源申请的可视化

### 内存指标

1. 初始化一个占用较大内存块的对象或多个对象（且他们占用的总内存交大），且初始化后不释放, 程序运行占用较大的内存，可以通过 inuse_space 来体现

2. 不断在初始化多个对象，且占用较多的内存，可以通过 inuse_objects 和 alloc_objects 来体现
   
3.存在非常频繁的 GC 活动，通常意味着 alloc_space 非常高，而程序运行过程中并没有消耗太多的内存（体现为 inuse_space 并不高） 
当然也可能出现 GC 来不及回收，因此出现 inuse_space 也变高的情况 内存泄漏，通常 alloc_space 较高，且 inuse_space 也较高

`go tool pprof -alloc_space mem.pprof` 指定一个指标查看; 默认是查看`inuse_space` 已分配但未释放的内存指标
