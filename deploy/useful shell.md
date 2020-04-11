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
