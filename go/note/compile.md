# 编译笔记

## 跨平台交叉编译

`CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go`

针对Linux平台的编译命令, 编译后可执行文件名为`server`

### CGO_ENABLED参数对编译影响

- 如果仅仅是非net、os/user等的普通包，那么你的程序默认将是纯静态的，不依赖任何c lib等外部动态链接库；

- 如果使用了net这样的包含cgo代码的标准库包，那么CGO_ENABLED的值将影响你的程序编译后的属性：是静态的还是动态链接的；

- CGO_ENABLED=0的情况下，Go采用纯静态编译；若进行跨平台编译, 一般需要设置`CGO_ENABLED=0` 不依赖任何c, lib等外部动态链接库 

- 如果CGO_ENABLED=1，但依然要强制静态编译，需传递-linkmode=external给cmd/link
