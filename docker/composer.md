# Docker基本笔记

## Docker Compose 基本命令记录

`docker-compose up -d 服务名称` 对服务镜像建立容器，并进行启动

`docker-compose up --build -d 服务名称` 对服务重新进行编译与生成镜像, 并启动服务

## <none>:<none> 空镜像产生原理

1. 专业名词: dangling images; 产生条件: 没有标签; 不再被容器使用

2. 从docker的文件层组织，镜像的生成是依赖不同层封装组成; 属于中间层的镜像


### 针对<none>:<none> 空镜像处理: 

此过程可被理解为手动进行垃圾回收

```sh
# 移除已停止弃用的容器
docker container prune
# 移除空置镜像文件
docker image prune
```
