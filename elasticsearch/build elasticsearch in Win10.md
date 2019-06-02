# Win10系统安装elasticsearch相关注意事项

主要记录使用Win10系统下安装elasticsearch 与 Kibana窗口交互及elasticsearch-analysis-ik 中文分词插件的相关流程

## 解决elasticsearch的ik analyzer access denied "java.io.FilePermission"

- 报错信息

`
access denied ("java.io.FilePermission" "D:\Program%20Files\elasticsearch-6.2.4\plugins\analysis-ik\config\IKAnalyzer.cfg.xml" "read");
`

从异常提示信息上看是权限问题，就是说 IKAnalyzer.cfg.xml 这个文件，java 程序只有读取权限，而对于 IKAnalyzer.cfg.xml 文件，elasticsearch 程序代码可能还需求其他一些权限，所以不能只有读权限.

- 解决

调整文件夹权限，开放访问权限

## Kibana初始化报错

- 报错信息

index_not_found_exception

`"message":"closed: [index_closed_exception] closed`

- 解决

.kibana索引被关闭；可以执行命令：`curl -X POST "localhost:9200/.kibana/_open"`

