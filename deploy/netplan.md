# netplan 网络设置

Netplan 是 Ubuntu 17.10 中引入的一种新的命令行网络配置实用程序，用于在 Ubuntu 系统中轻松管理和配置网络设置。 它允许您使用 YAML 格式的描述文件来抽像化定义网络接口的相关信息

在 Ubuntu 18.04 中如果再通过原来的 ifupdown 工具包继续在 /etc/network/interfaces 文件里配置管理网络接口是无效的。

## 文件配置与操作

对于 Ubuntu 系统，配置文件在 `/etc/netplan` 路径下; 默认文件名: 50-cloud-init.yaml

### 配置生效与校验操作

生效操作: `netplan apply`

校验操作: `netplan apply --debug`

## 常见配置

1. enp0s5 指定需配置网络接口的名称。
2. dhcp4  是否打开 IPv4 的 dhcp。
3. dhcp6  是否打开 IPv6 的 dhcp。
4. addresses 定义网络接口的静态 IP 地址。
5. gateway4  指定默认网关的 IPv4 地址。
6. nameservers  指定域名服务器的 IP 地址。

### 设置静态IP地址

- 配置dhcp4 为false

- addresses: 配置静态地址，如下图，则配置 `192.168.56.5` 为机器静态 IP 地址

- gateway4 可以不进行配置

- nameservers 若不进行配置, 会引发机器端口无法 ping 通外网地址

```yaml
network:
    ethernets:
        enp0s8:
            dhcp4: false
            addresses: [192.168.56.5/24]
            gateway4: 192.168.56.1
            nameservers:
                addresses: [8.8.8.8, 8.8.4.4]
    version: 2
```

当配置生效后, ifconfig 命令打印端口信息: `enp0s8 inet 192.168.56.5  netmask 255.255.255.0  broadcast 192.168.56.255`

#### DHCP 服务

针对 VirtualBox 虚拟机, 每次重启机器，机器 IP 地址都会变更。因开启了DHCP 服务，自动设定网络配置, 为机器提供动态 IP 地址

### 常用命令

`route` 查看机器端口路由表

`ip address` / `ip a` 机器端口 IP 地址信息

### 备注

静态地址 以 `IP/数字` 进行配置的原因: 斜杠后面的数字是子网掩码的信息, 也称作网络位; 网络位越大，可分配的子网IP数越小

一般配置 IP 地址，需要 IP 地址 与 子网掩码 两个配置，若以 `IP/数字` 则同时完成 IP 与 子网掩码 的配置

`192.168.56.5/24` 解读为 IP 地址 是 `192.168.56.5` 子网掩码是 `255.255.255.0` 子网IP数有254个, 主机网IP有8位

对于前24位都为1, 二进制下，高位前三个字节，都为1，则十进制为255

`192.168.56.5/28` 解读为 IP 地址 是 `192.168.56.5` 子网掩码是 `255.255.255.240` 子网IP有14个, 主机IP有4位
