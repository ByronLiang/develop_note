# 记录实用shell命令

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

