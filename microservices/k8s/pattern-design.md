# K8S 常用设计模式

## SideCar 模式

一个Pod至少存在两个容器(container), 它们共享同一个pod的资源(网络、物理资源)，容器间能相互通信，

### 优势

因减小应用间网络通信消耗，而将多个容器部署至同一个pod

指标数据采集任务，如Prometheus 探测器(exporter)，减少对探测目标资源损耗

#### 常见配置

同一个pod下多个container暴露相应端口号, 并对应相应 service 的配置端口

```yaml
apiVersion: v1
kind: Service
metadata:
  name: xx
  labels:
    app: x
spec:
  ports:
    - name: xx1
      port: 27017
      targetPort: 27017
    - name: xx2
      port: 9104
      targetPort: 9104
  selector:
    app: xx
---
## Deployment
apiVersion: apps/v1
kind: Deployment
metadata:
  name: xx
  labels:
    app: xx
spec:
  replicas: 1
  selector:
    matchLabels:
      app: xx
  template:
    metadata:
      labels:
        app: xx
    spec:
      containers:
        - name: xx
          image: xx:0.0.1
          ports:
            - containerPort: 27017
        - name: xx2
          image: xx2:latest
          ports:
            - containerPort: 9104
```
