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

## Docker 容器健康检测

在 dockerfile 与 docker-composer 文件都有 HEALTHCHECK 标签, 可在启动时对容器进行探测, 可配置间隔时间、超时时间与重试次数

一般探测可使用调用服务的ping 接口，容器内部需要安装 `curl` 命令包 或者 调用程序自带的应用

MySQL 可使用自带的 `mysqladmin` 程序进行程序可用性探测: `mysqladmin ping -h localhost`

```yaml
healthcheck:
      test: ["CMD", "mysqladmin" ,"ping", "-h", "localhost"]
      interval: 5s
      timeout: 10s
      retries: 3
```

### 服务状态 Condition

对于依赖其他服务，可对所依赖的服务进行服务探测, 确保依赖服务正常, 然后再启动当前服务

若对依赖服务进行 condition 选项，被依赖的服务需要加上 `healthcheck`, 从而得出探测结果

```yaml
depends_on:
      leader:
        condition: service_healthy
```
