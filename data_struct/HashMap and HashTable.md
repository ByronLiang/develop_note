# HashTable 与 HashMap 及 PHP数组底层实现

## 区别

以Java语言解释

1. Hashtable是线程安全，而HashMap则非线程安全

- Hashtable的实现方法里面都添加了synchronized关键字来确保线程同步，因此相对而言HashMap性能会高一些，我们平时使用时若无特殊需求建议使用HashMap，在多线程环境下若使用HashMap需要使用Collections.synchronizedMap()方法来获取一个线程安全的集合。

- 效率低: 当一个线程访问HashTable的同步方法时，其他线程访问HashTable的同步方法时，可能会进入阻塞或轮询状态

2. HashMap可以使用null作为key，而Hashtable则不允许null作为key

3. 4.HashMap扩容时是当前容量翻倍即:capacity*2，Hashtable扩容时是容量翻倍+1即:capacity*2+1

4. Hashtable计算hash是直接使用key的hashcode对table数组的长度直接进行取模; HashMap计算hash对key的hashcode进行了二次hash，以获得更好的散列值，然后对table数组长度取摸

## PHP数组底层实现原理

哈希表+双向链表: 由HashTable组成; 解决哈希冲突是使用`链地址法` 将同一个slot中的bucket通过链表链接起来；

## Redis哈希类型底层原理

### 装载因子

Hash表的装载因子 = 填入表中的元素个数 / Hash表的长度

#### 指标作用

- 当元素越多，空闲位置越少，散列冲突的机率越大，性能就会下降；

- 当超过值，通过动态扩容，申请更大的Hash表, 并将原旧数据拷贝到新Hash表里

- 当hash表的元素不断被移除，空闲位置越多，若对内存空间敏感，则会触发动态缩容，回收部分内存

### ReHash渐进式扩容

1. 当发送扩容，迁移原旧数据时，并非一次性完成数据迁移；而是以`分治思想`渐进地完成数据迁移

2. 新增的数据会优先存放在新扩容的表里；旧哈希表只会因不断迁移数据而数据量不断减小

3. Redis的哈希类型里，底层有两个哈希表, 其中一个存储数据, 另一个用作扩容的哈希表;

## Hash 对 key 查询复杂度为 O(1)

Hash 底层存储数据是数组, 若没有出现哈希碰撞, key 进行哈希化得出hashCode, 然后以数组长度的值对hashCode 取模, 定位数组的下标位置存储数据。

一般初始数组长度为8, 若哈希扩容, 会以倍数进行扩容。同时需要将旧Bucket数据迁移至新扩容的数组
