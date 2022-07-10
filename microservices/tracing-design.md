# 分布式追踪系统搭建


为显示了执行时间的上下文，相关服务间的层次关系，进程或者任务的串行或并行调用关系。

这样的视图有助于发现系统调用的关键路径。通过关注关键路径的执行过程，项目团队可能专注于优化路径中的关键位置，最大幅度的提升系统性能

## 核心概念

### Traces

具备潜在分布式的执行轨迹的系统；是作为多个span的有向无环图

### Spans

一个span可以和一个或者多个span间存在因果关系。OpenTracing定义了两种关系：ChildOf 和 FollowsFrom。这两种引用类型代表了子节点和父节点间的直接因果关系。

#### Spans 与 Baggage

Baggage在全局范围内，（伴随业务系统的调用）跨进程传输数据。Span的tag不会进行传输，因为他们不会被子级的span继承。

span的tag可以用来记录业务相关的数据，并存储于追踪系统中。实现OpenTracing时，可以选择是否存储Baggage中的非业务数据