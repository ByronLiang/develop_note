# Go 开发环境记录

## Go SDK 下载

- wget下载指定版本的go版本 `linux-amd64.tar.gz`; 指定路径解压 `tar -xzf`

- 在`/etc/bash.bashrc` 配置环境变量; 设定Go Path 地址（存放扩展包地址） 

```sh
export GOPATH=/usr/local/src/go/path
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

## 配置与模块初始化

### Go 配置设定

1. `go env` 查看配置信息; 

2. `go env -w GO111MODULE=on` 开启包管理

3. `go env -w GOPROXY=https://goproxy.io,direct` 设置下载包的代理镜像; `https://mirrors.aliyun.com/goproxy/` 阿里云镜像地址

### 模块项目初始化

1. `go mod init xxx` 对xxx模块初始化, 生成`go.mod` 包管理相关信息

2. `go mod tidy` 安装缺失依赖

3. `go mod clean` 移除无关依赖
