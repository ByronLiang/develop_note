# Minikube 常用命令

## 配置

### 虚拟机安装目录

在宿主机环境变量，新建key：`MINIKUBE_HOME` 配置minikube 节点虚拟安装位置 (Virtual Box)

### 节点挂载路径

节点(node)默认挂载宿主机路径：`/c/Users/`

配置文件路径需要属于节点挂载路径: `kubectl create configmap 名称 -n 命名空间 --from-file 配置文件名 -o yaml`

## 开发

#### 打包本地镜像

`minikube image build -t 镜像标签 -f 指向Dockerfile路径`

- 若部署多个节点, 需要为每个节点都需重新打包镜像

- 容器配置: `imagePullPolicy: Never` 优先使用本地镜像文件

#### 更新与部署

`kubectl edit pods pod名称 -n 命名空间` 变更 pod 配置并自动更新 pod

`kubectl apply -f 配置.yaml` 部署配置文件

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