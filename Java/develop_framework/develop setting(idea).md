# 建立Java基本开发环境

1. [安装JDK](http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html) 
2. 安装IDEA [破解 version](https://www.jiweichengzhu.com/article/c6ae011e3fce496fb11b6ba60c1a0e41)

3. 配置IDEA

- 变更文件编码

`Editor -> File Encodings -> transparent native-to-ascii conversion`

- 代码编写设置

`Editor ->Code Style ->Java ->Code Generration ->Field Prefix m & Static Field prefix s`

- lombok插件注释启用

启用注释:

`Build, Execution, Deployment -> Compiler -> Annotation Process`

`勾选 Enable annotation processing`

- 安装相关插件：

Rainbow Brackets
Lombok Plugin
Relative Line Numbers
RestfulToolkit
GsonFormat

## 基本使用操作

### 导入新项目

使用`open`

### 引入新包/变更关联库

在`Gradle`里进行刷新, 重新拉取包的配置

### 控制引入包为*显示配置

`Code Style -> Java -> Class count to use import with *` 配置相关数量

### Windows平台 CMD命令行启动SpringBoot（启动内置Tomcat)

`java -Duser.timezone=Asia/Shanghai -Dfile.encoding=UTF-8 -Xms2048m -Xmx2048m -jar xxx.jar`

1. `xxx.jar` 是项目的打包文件；`./gradlew --no-daemon bootJar`(打包命令)

2. `-Dfile.encoding=UTF-8` 配置jvm(tomcat)内置文件编码 (Window 默认使用GBK) 

解决遇到的错误: `The Java Virtual Machine has not been configured to use the desired default character encoding`

3. `-Xms2048m -Xmx2048m` 配置运行最小最大内存

4. `-Duser.timezone=Asia/Shanghai` 设置jvm默认时区
