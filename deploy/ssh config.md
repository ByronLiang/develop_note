# SSH服务端与客户端的相关配置笔记

## SSH服务端

- 进入`/etc/ssh/` 查看是否拥有`sshd_config`文件；
- `service ssh status` 查看ssh服务状态
- `apt-get install openssh-server` 安装SSH服务端

### ssh密匙生成

- 利用ssh-keygen 生成密匙: `ssh-keygen -t rsa -P ""`
- 生成密匙的被访问对象: 通过以不同用户身份来生成访问密匙(root: `/root/.ssh/`) 或 (xxxuser:`/home/xxxuser/.ssh`)
- 在文件夹里`/root/.ssh/`里会有`id_rsa`(私匙)与`id_rsa.pub`(公匙) 
- 同时，需手动新建`authorized_keys`: `(chmod 600 authorized_keys)`
- `cat id_rsa.pub >> authorized_keys` 复制公匙到`authorized_keys` [上述重定向时使用>>进行追加, 不要用>, 那会清空原有内容.]

### 密匙认证配置

- 配置`/etc/ssh/sshd_config`文件，对`~/.ssh/authorized_keys`进行密匙认证

## SSH客户端

### 查看软件安装

- 查看ssh 客户端配置: `/etc/ssh/ssh_config`
- `apt-get install openssh-client` 安装ssh 客户端

### 密匙授权请求配置

- 将`~/.ssh/id_rsa`私匙 发放到指定ssh客户端里；可对`ssh_config`进行配置连接主机的相关配置

```sh
Host wsl
    User wsl
    HostName 192.168.92.1
    PORT 8022
    IdentityFile ~/.ssh/wsl/wsl_rsa.key
```
上述配置里，定义主机地址`wsl`的请求配置信息，`IdentityFile`为私匙配置存放文件；

发起连接: `ssh wsl`
