# 容器网络

## Nat 分类

### SNAT(源地址转换)

SNAT操作的是iptables中NAT表的`Postrouting`链 (封包传出阶段)，将ip包的源地址替换为网关的公网地址

### DNAT (目的地地址转换)

DNAT操作的是iptables中NAT表的Prerouting链，将ip包的目的地址替换为内网要映射端口的主机地址。

### k8s 网络说明

k8s 使用 ipvs、kube-proxy 组件 (NodePort)，属于 DNAT 类型 会有一定网络性能损耗

#### 流量请求流程

请求流量进入负载均衡组件，并转发到某一个节点的NodePort；KubeProxy将来自NodePort的流量进行NAT转发，目的地址是随机的一个Pod。请求进入容器网络，并根据Pod地址转发到对应节点；请求来到Pod所属节点，转发到Pod

#### 直连模式

减少进行 NAT 转发，负载均衡后，流量直接转发到 Pod 里 (Pod 直连负载网关)
