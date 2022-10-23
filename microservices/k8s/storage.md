# 存储

## ConfigMap

配置文件以软链(symbolic links)方式在容器配置的挂载路径上, 而软链的目标地址再经过软链，都在同一挂载目录下

```sh
drwxr-xr-x    2 root     root          4096 Mar 31 09:39 ..2022_03_31_09_39_45.609364522
lrwxrwxrwx    1 root     root            31 Mar 31 09:39 ..data -> ..2022_03_31_09_39_45.609364522
lrwxrwxrwx    1 root     root            27 Mar 10  2022 dev.xx.config.yaml -> ..data/dev.xx.config.yaml
```

因此，容器里进程需要监听文件变更(`Linux 的 inotify`) 需要监听整个挂载目录

### 共享卷(volume) 选型

#### Empty Dir

临时挂载路径。当容器销毁，挂载路径的文件数据同时被销毁

#### HostPath (主机路径挂载)

适合于写数据: 涉及容器实例写数据到挂载路径文件, 而宿主机的路径能获取其写入文件数据。同时当容器销毁时，挂载路径文件仍得到保留

#### 分布式存储管理驱动

常见网络文件系统: NFS、Glusterfs、cephfs 等

适合多个 pods 共享一份数据, 解决 hostpath 挂载宿主机, 无法跨主机节点共享数据

容器实例从挂载路径读取数据，但pods的副本实例分布多个节点，需要使用分布式文件管理系统，确保每个节点的挂载路径文件，文件数据一致，从而 pods 读取数据能实现一致性。

#### Glusterfs 驱动

k8s 共享卷作为 NFS 的 client, 通过连接 NFS 的server, 为共享卷提供数据。其驱动主要调用 `github.com/heketi/heketi` api 库
