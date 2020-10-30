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
