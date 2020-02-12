# HashSet And TreeSet 处理List去重

## HashSet

### 基本原理

- 存储唯一元素与可存储空值
- 初始化内部为HashMap
- 不保持插入顺序

### 去重原理

1. 对象放入HashSet时，会使用对象里的hashcode值确定元素是否已经存在该集合中
2. 当存在相同的对象(相同的哈希值)，则使用equals()方法进行比较，再决定是否存入集合里

## TreeSet

### 基本原理

- 存储具有唯一性，可对元素进行排序
- 底层结构是二叉树；节点进行存储与取出

## 例子

- 无序去重使用HashSet，有序去重TreeSet

```java
public static List removeDuplicationByTreeSet(List<Integer> list) {
    TreeSet set = new TreeSet(list);
    //把List集合所有元素清空
    list.clear();
    //把TreeSet对象添加至List集合
    list.addAll(set);

    return list;
}
```

- 原生方法处理去重

```java
public static List removeDuplicationByStream(List<Integer> list) {
    List newList = list.stream().distinct().collect(Collectors.toList());
    return newList;
}
```
