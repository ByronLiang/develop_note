# 记录实用shell命令

## Windows下创建shell文件的异常

1. Windows创建的shell文件是dos格式的； 而执行shell文件需要使用unix文件格式;

2. 若shell的文件格式为dos, 则执行shell时则发生异常: `/bin/bash^M: bad interpreter`

### 查看文件格式

```sh

vi 文件名

# 进入文件内容页面, 命令行输入, 查看当前文件的fileformat
# 一般会显示: fileformat=dos [windows创建] 一般Linux/Mac os系统创建的，则为 fileformat=unix
: set ff

```

### 转换文件格式

```sh

vi 文件名
# 进入文件内容页面, 命令行输入转换, 将当前文件格式转换为unix
: set ff=unix
# 按照转换后的格式保存文件
:wq
```

## 进程与文件描述符

1. 一个进程启动以后，除了会分配堆、栈空间以外，还会默认分配三个文件描述符句柄：`键盘输入-stdin[标志号：0]`、`屏幕输出-stdout[标志号：1]`，`屏幕输出-stderr[标志号：2]`

2. `/dev/null` 是一个特殊的设备文件，所有接收到的数据都会被丢弃;从 `/dev/null` 读数据会立即返回 EOF

3. 文件描述的配置，可以对进程的日志输出交互进行配置；如程序日志是否显示在终端屏幕上

### 2>&1 shell脚本含义

`2>`表示`重定向 stderr` ，`&1` 表示 `stdout`，连起来的含义就是将标准错误输出 stderr 改写为标准输出 stdout 相同的输出方式。常用于日志写入

## 查看目录大小

`du -h 文件目录` / `du -sh 文件目录`

`ls -lh` 无法深度遍历每个文件目录的文件大小，只能计算文件块的大小。

### ls 与 du 命令

若 `ls -lh` 命令统计非文件夹对象，那么统计的大小与 `du -sh` 命令是一致的

针对统计具有文件夹的对象，`ls` 显示的文件大小会比 `du` 命令显示的体积要小

## scp

scp就是secure copy的简写; 用于在Linux下进行远程拷贝文件的命令，和它类似的命令有cp，不过cp只是在本机进行拷贝不能跨服务器

命令格式: `scp [参数] [原路径] [目标路径]`

### 将本机文件复制到远程服务器上

`scp -P 8022 /home/test/xxx.sh root@192.168.11.2:/home/test/sample`

- 指定远程服务器端口: `-P 8022`; 
- 本机文件路径: `/home/test/xxx.sh` 
- 远程服务器地址: `root@192.168.11.2`
- 远程服务器存放文件路径: `/home/test/sample`

### 从远程服务器下载文件到本地

`scp -P 8022 root@192.168.11.2:/opt/soft/xxx.sh /home/test/sample`

- 远程服务器地址与文件路径 `用户名@服务器地址`:`文件路径`
- 本机下载文件到指定路径: `/home/test/sample`

## 本地转发与远程转发

### 本地转发

将服务器的本地地址端口转发到本地机器上: `ssh -L 12708:127.0.0.1:12707 root@10.1.0.2`
密匙认证转发: `ssh -i 密匙文件 本地端口:远程本地:端口号 远程服务名@远程服务IP地址`

命令执行于本地机器上(接收隧道转发)：ssh客户端，又是隧道的应用客户端; 远程机器: ssh服务端，隧道应用服务端

### 额外参数 

1. `-N` 创建ssh隧道时，并不会打开远程shell连接到目标主机
2. `-f` 后台运行ssh隧道

### 远程转发

`ssh -f -N -R 10.1.0.1:9906:10.1.0.2:3306 root@10.1.0.1` 

远程本地机器: 10.1.0.2:3306 转发到隧道远程机器: 10.1.0.1:9906

命令执行于需要被转发的远程机器上: 远程本地主机: ssh客户端，隧道应用服务端; 远程机器: ssh服务端, 隧道应用客户端

## kill

格式: `kill[参数][进程号]` kill命令默认的信号就是15; (15: 终止; 9: 强制终止)

### kill -15信号

#### 特点：

- 进行"安全、干净的退出"; 
- 可以被阻塞和忽略的

##### 流程: 

退出前一般会进行一些"准备工作"，如资源释放、临时文件清理等等，如果准备工作做完了，再进行程序的终止。如果在"准备工作"进行过程中，遇到阻塞或者其他问题导致无法成功，那么应用程序可以选择忽略该终止信号。

### kill -9信号

#### 特点
- 具备强制性，不能被阻塞或者忽略
- 数据丢失或者终端无法恢复到正常状态等

#### 流程

系统会发出SIGKILL信号，他要求接收到该信号的程序应该立即结束运行。

### 关于popd、pushd、dirs的使用方式

文件目录栈的使用

1、dirs (区别于dir：dir的功能与ls类似)

显示栈顶到栈底的文件目录列表，通常显示的第一个文件目录为当前的工作目录；

2、pushd+文件目录(相对/绝对路径)

向目录栈中添加相应的文件目录，同时当前工作目录调整到相应的目录下；

3、pushd 单独使用

调换目录栈中最近的两个目录，同时刻当前的工作目录会发生变化；

4、popd的工作方式

弹出当前目录栈中保存的工作目录，工作目录变为栈中相应的目录；

### > /dev/null的使用方式

`> /dev/null`简单的理解，将命令的输出信息输入到 /dev/null中，标准输出几口将输出信息显示到屏幕上，而使用`> /dev/null`将消除命令回显信息的显示。

如：`pwd > /dev/null`将无任何信息显示

### 常用日志查看技巧

`tail -f 文件名 -n 200`

查看文件名 并从倒数n行开始查阅

### 常见统计

 `grep "查询字符" 目标查询文件 | wc -l`: 统计当前文件出现查询字符的数目

[awk 用法](https://www.junmajinlong.com/shell/awk/index/)

 `tail -n 行数 查询文件名 | awk -F "字符串查询" '{print $2}' | awk -F '\t' '{for(i=1;i<NF;i++) if(match($i,/bfl/)) print $i}'|sort |uniq -c|sort` 
 
 对查询文件内容进行字符串查询统计命令

`-F 查询字符内容` 分割相应内容 

$1：`被分割的左边字符内容`

$2: `被分割的右边内容`

`awk -F "分割字符串" '{print $2}' |  awk -F "分割字符串" '{print $1}'`: 打印基于第一个分割字符串与第二个分割字符串之间的内容

 `match($i,/bfl/)` 正则查询

### shell 获取日期与当前小时

使用 `date` 配合格式化，得出年月日等时间内容，常用于生成以时间日期等维度的文件名

`date +%Y%m%d` 显示年月日

`date --date "-2 hours" +%k | tr -d "[:space:]"` -d / --date 涉及时间操作，加减小时数

#### 输出24小时制

若使用 `date +%k` 显示小时数值，对于 0-9 的数值，前置会有空格。

因此，只能采取前置移除空格字符的措施，`tr -d "[:space:]"`

使用管道处理:`date +%k | tr -d "[:space:]"`, 可以确保能正常显示 0-9 的小时数值，并且数值前置不会有空格字符

## nohub 命令

将执行命令放置后台进行，效果与 `supervisor` 类似，但无法实现进程守护，当进程被 kill, 无法自动重启

### 用法

`nohup command &`

后台运行程序，将程序日志输出到指定文件

`nohup ./二进制文件 >> /var/log/nohup.log 2>&1 &`

`nohup /usr/local/bin/sh dev.sh >> /var/log/nohup.log 2>&1 &`

#### 查看 jobs

`job -l`: 查看运行的后台进程

#### sed 兼容

[sed-command](https://stackoverflow.com/questions/4247068/sed-command-with-i-option-failing-on-mac-but-works-on-linux)

`sed -i 's/old_link/new_link/g' *` 

`sed -i` 不兼容 mac 系统，异常：`command expects \ followed by text`

兼容 GNU and BSD 系统的 sed 写法：`sed -i'' -e 's/old_link/new_link/g' *`

#### 变量大小写处理

[bad-substitution](https://stackoverflow.com/questions/47815637/getting-bad-substitution-error-with-a-shell-script-on-a-mac)

`${name^}`/`${name,}`: 对变量的字符串进行首字母大小写处理。

`^`, `,` 和 `,,` (大写、小写) 字符处理，针对 bash 4.0 以上版本才兼容，若低版本执行，会抛出异常：`bad-substitution`
