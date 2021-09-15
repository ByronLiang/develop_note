# Git Bash 基本配置

## CRLF警告 针对Unix与Win 字符处理

- CRLF: windows-style `Carriage-Return Line-Feed`
- LF: unix-style

### 显示告警原因

常见告警: LF will be replaced by CRLF

配置项: core.autocrlf; 查看命令`git config --get core.autocrlf`

当 `core.autocrlf` = `true` 会显示其告警

#### 解决

core.autocrlf 可设置为 true, false, input

`git config --global core.autocrlf false`

[详情参考](https://stackoverflow.com/questions/1967370/git-replacing-lf-with-crlf)

## 凭证存储

针对Http Clone 的项目，每次提交都需要输入用户名与密码，通过启用`credential.helper` 为 cache; 可指定缓存文件路径与缓存过期时间

### 类型

1. cache: 凭证放在内存中，不会存储在磁盘中; 命令: `git config --global credential.helper 'cache --timeout=3000'` timeout单位为秒

2. store: 凭证是明文存储在磁盘中, 不会过期；配置指定位置；命令: `git config --global credential.helper 'store --file ~/.credential'` file 指定凭证存储的路径

## 子模块submodule

### 基本使用

1. `git submodule add 仓库地址` 会生成或基于原.gitmodules文件追加子模块配置

path: 子模块目录名
url: 子模块仓库地址

.gitmodules 文件内容

```sh
[submodule "tools"]
        path = tools
        url = git@github.com:xx/xxx.git
```

2. 在父仓库更新/拉取子模块: `git submodule update --init`

3. 子模块的变更发布：需要进入子模块目录路径，才能完成子模块的仓库发布提交；在父模块无法提交子模块的变更的内容


