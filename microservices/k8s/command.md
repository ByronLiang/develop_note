# 常用指令

## Pods 相关

### 配置类

获取Pod 名称与服务网络配置信息:`kube get pods -o wide -n 指定命名空间` / `kube get pods -o wide -n 指定命名空间 | grep pod名称`

### 开发

进入Pod内部: `kube exec -it pod名称 -n 指定命名空间 sh`

