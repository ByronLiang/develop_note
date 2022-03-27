# K8s 运维生态

## 网络服务

K8s Ingress 网络服务，原生K8s服务

Istio 网络服务: 在K8s 基础上开发 有sidecar proxy 与 envoy 可视化平台; 不再使用 kube-proxy 组件做流量转发，而是依托在每个 pod 中注入的 sidecar proxy，所有的 proxy 组成了 Istio 的数据平面；

### CoreDNS

k8s 原生集群DNS解析: kube-dns

k8s集群下的DNS解析插件: CoreDNS

主要实现为集群内gRPC提供服务发现(服务名发现与解析)

### Minikube 常用命令

#### 打包本地镜像

`minikube image build -t 镜像标签 -f 指向Dockerfile路径`

- 若部署多个节点, 需要为每个节点都需重新打包镜像

- 容器配置: `imagePullPolicy: Never` 优先使用本地镜像文件

#### 查看服务对外地址

`minikube service -n 命名空间 list` 查看此命名空间的服务列表

#### Pod 无法通过 Service IP 连接到它本身

现象: pod 里能正确DNS解析, 但对 service 的 IP 地址无法请求

- 检测能否正确解析域名的 IP 地址

当service name 与 pod 在同一个 namespace 无需标明namespace

`nslookup service` / `nslookup service.namespace` 

- hairpin-mode 标志必须被设置为 hairpin-veth 或者 promiscuous-bridge。

```sh
# 登录节点
minikube ssh
# 配置 docker0 网卡 为promisc 模式
sudo ip link set docker0 promisc on
```

### 基本排查

以下操作为登录集群节点, 可进行的排查手段

#### 查看组件进程

登录集群节点, `ps auxw | grep kube-proxy`

#### 使用 journalctl 查看组件日志

`journalctl -u docker` docker 内核日志

`journalctl -u kubelet -f` k8s 输出日志

#### k8s 组件与日志

`kubectl get pods -n kube-system` 查看k8s 组件的pod

`kubectl logs pod的名称 -n kube-system --tail 最新数目` 输出指定数目日志条数

`kubectl logs pod的名称 -n kube-system -f` 输出即时日志流
