# Redis相关优化

## 应用

### 哈希数据结构实现超时删除

一般情况，只能对hash的key设置超时，无法对hash里的每个成员设置超时，通过每个成员标注超时时间，利用定时任务，使用hscan命令来对成员进行超时移除处理

`hscan 哈希key 0 match 成员key.*/成员key指定数值 count 100` 主要使用hscan方法，设置浮标 match 通配符 设置返回条目数

## 内存优化

### 关注数据结构的编码

同一种数据结构，多种编码方式，为实现效率与空间平衡；

Redis写入数据时自动选择编码方式与编码类型转换；转换过程不可逆，遵循小内存编码向大内存编码转换

利用命令`object encoding key名字` 查看当前存储key的数据结构编码

配置制定数据大小，来动态使用不同的编码格式 `config set list-max-ziplist-entries 4` 当此key的list数据结构，超出4个元素，则会使用`linkedlist`编码类型

#### string

1. embstr 优化内存分配字符串编码 默认

2. raw 动态字符串编码

3. int 整数编码

#### hash

1. hashtable 散列表编码

2. ziplist 压缩列表编码

### 优化方向

1. 控制key的数量；对于存储相同的数据内容利用Redis的数据结构降低外层键的数量，也可以节省大量内存

2. ziplist编码类型可以大幅降低内存占用；长度不要超过1000，每个元素大小控制在512字节以内；ziplist适合存储的小对象，对于大对象不但内存优化效果不明显还会增加命令操作耗时

