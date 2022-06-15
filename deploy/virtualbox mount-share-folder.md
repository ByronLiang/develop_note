# VirtualBox平台对虚拟机挂载宿主机文件夹

常见错误: 无法识别nfs; 或者 缺乏辅助程序引发 `missing codepage or helper program`

## 安装辅助软件

`sudo apt install nfs-common`

`sudo apt-get install virtualbox-guest-utils`

## 设置共享文件夹与新建对应命名路径

在 VirtualBox 设置页面进行设置，选择 `自动挂载`选型，并设置宿主机文件路径及共享文件名

在虚拟机系统的 `/mnt/` 路径下，创建对应的共享文件名, 权限可以是 `root`

## 执行挂载命令

`sudo mount -t vboxsf 共享文件名 共享文件名指定路径` 如 `sudo mount -t vboxsf www /mnt/www`
