# K8s 运维生态

## 网络服务

K8s Ingress 网络服务，原生K8s服务

Istio 网络服务: 在K8s 基础上开发 有sidecar proxy 与 envoy 可视化平台; 不再使用 kube-proxy 组件做流量转发，而是依托在每个 pod 中注入的 sidecar proxy，所有的 proxy 组成了 Istio 的数据平面；

### CoreDNS

k8s 原生集群DNS解析: kube-dns

k8s集群下的DNS解析插件: CoreDNS

主要实现为集群内gRPC提供服务发现(服务名发现与解析)
