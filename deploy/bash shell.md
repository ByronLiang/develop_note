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
