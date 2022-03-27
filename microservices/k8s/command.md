# 常用指令

## Pods 相关

### 配置类

获取Pod 名称与服务网络配置信息:`kube get pods -o wide -n 指定命名空间` / `kube get pods -o wide -n 指定命名空间 | grep pod名称`

### 开发

进入Pod内部: `kube exec -it pod名称 -n 指定命名空间 sh`

更新 configmap 配置: `kubectl create configmap configmap名称 -n 命名空间 --from-file 配置文件路径 -o yaml --dry-run | kubectl apply -f -`

#### Service 端口配置

```yaml
ports:
    - name: http
      protocol: TCP
      nodePort: 8090 # 集群外暴露外部端口访问, 若不配置, 则随机端口号. 配置时，不能配置已启用的端口号
      port: 8090 # 集群端口号, 主要用于集群内部调用
      targetPort: 8090 # 配置容器本身暴露端口号: 需要与DockerFile中的EXPOSE; 最终流量经过kube-proxy进入到容器端口
```

#### 滚动升级

单独升级镜像: `kubectl set image deployment {deployment名称} {容器名称}={镜像名称} -n {命名空间} --record`

基于配置更新: `kubectl replace -f {配置文件}.yaml -n {命名空间} --record`
