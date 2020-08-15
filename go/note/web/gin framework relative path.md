# Gin 框架 相对路径定义

## Go的文件相对路径是相对于执行命令时的目录

1. 一般规范，相对路径是指相对于`main.go`(被编译文件)所处的目录路径; 若路径配置为`./stroage/static/`; 则`./`指代与`main.go`同级的文件目录路径

2. 若执行命令不在与`main.go`同级路径目录; 如以`go run /src/main.go`(main.go文件的上一级目录);  `./`指代`/src/`此级文件路径, 会引发无法匹配`stroage/static`文件路径引发错误
