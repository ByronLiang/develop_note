# 基本记录

## Docker 容器信号

### 容器关闭信号

容器只向1号进程发送信号; 需确保容器里程序在1号进程执行, 并且程序做好相关信号捕捉, 从而实现优雅关闭业务程序

#### 信号流程

- stop 命令会首先发送`SIGTERM`信号，并等待应用优雅的结束。如果发现应用没有结束(用户可以指定等待的时间)，就再发送一个 SIGKILL 信号强行结束程序。`docker stop ----time=30 容器名`

- kill 命令默认发送`SIGKILL`信号；可以通过 -s 选项指定信号: `docker kill ----signal=SIGINT 容器名`

- `docker rm -f` 强制删除容器, 依然是先发出`SIGKILL`信号

- 信号`SIGINT` 是`Ctrl+C`执行发出; 针对容器不是以守护进程启动的情景

#### 容器中PID为1(进程号为1)捕捉到信号

配置程序以1号进程启动，在DockerFile命令里，主要是`EntryPoint`, `CMD`, `RUN` 启动容器程序
