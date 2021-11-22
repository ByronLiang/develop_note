# gRPC Health Check 应用

## 原理

通过官方自带的健康检测协议与服务，实现对服务状态的探测。及时获取当前服务的健康状态, 提高服务可用性

一般, 结合k8s里的探针API，服务启动后，探针与服务包都处于同一个pod里, 可以方便调用 Health Check 接口

### k8s probe

- `startupProbe`: 用于确定容器内的应用程序/功能是否已启动并准备好进行进一步处理。其他探测器将被禁用，直到startupProbe成功。如果失败，容器将根据restartPolicy定义重新启动。

- `readinessProbe`: 用于确定容器是否用于确定它是否准备好为任何请求提供服务

- `livenessProbe`: 用于判断容器是否正在运行。如果失败，kubelet将重新启动受 约束的容器restartPolicy

每个probe 都有三个action: `ExecAction`, `TCPSocketAction`, `HttpGetAction` 与 三个状态: `success` `failure` `unknown`

### 生态工具

(监测健康探针二进制包)[https://github.com/grpc-ecosystem/grpc-health-probe]

#### 应用

1. 通过将服务的二进制包与探针二进制包 生成镜像文件，部署在k8s 的pod里;

2. 配置 k8s 的 `readinessProbe` 与 `livenessProbe`: 调用探针二进制包, 使其调用gRPC 的 health check 接口

### 简易检测服务可用性

应用层的接口路由新增ping接口, 当pod启动时, 调用ping接口地址，实现检测服务实例可用性; ping接口只允许内部请求

```yaml
spec:
    containers:
    - name: srv
      startupProbe:
        httpGet:
            path: /ping
            port: 8080
```
