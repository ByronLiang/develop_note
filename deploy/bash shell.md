# shell的相关笔记

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

