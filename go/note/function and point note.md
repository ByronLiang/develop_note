# 函数与指针的相关原理

```go
package main

func main() {
    // Declare variable of type int with a value of 10.
    count := 10
    // Display the "value of" and "address of" count.
    println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
    // Pass the "address of" count.
    increment(&count)
    println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
}

func increment(inc *int) {
    *inc++
    println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}
```

## 帧边界

1. 函数在各自单独的内存空间中执行，允许函数在自己的上下文中操作。

2. 函数可以直接访问帧内的内存，但是帧外部的内存只能间接访问。

3. 要访问帧外部的内存，这块内存必须与函数共享。

## 函数与栈

每个单独函数的帧在`栈里申请内存(初始内存为2k)`，以作为物理内存空间

当函数被调用，在两个帧间会发生转换。代码从调动函数的空间中转换到被调用函数的空间中。数据必须从一个空间传送到另一个空间, 传输是`按值`传递的

函数调用意味着协程要在栈上框出一段内存; 每当有函数调用，都会事先清扫帧所在的栈内存。

### 共享值 与 函数返回

使用指针作为形参，就是与函数共享某个值，即使这个值不直接存在自己的帧内，也可以对其读写

若对于共享型的数据类型(slice, map等) 因append操作进行扩容, 而产生新地址, 建议函数返回值
