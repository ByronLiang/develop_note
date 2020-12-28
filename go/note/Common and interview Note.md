# 日常开发问题/面试笔记

## 简单记录

1. 函数里不支持默认参数或可选参数

2. 表示枚举值(enums) `const 常量 ioat 常量自增计数器 组合`

## 针对切片数据结构比较是否相同

1. 若数据类型是`[]byte` 可以使用`bytes.Equal`;

2. `reflect.DeepEqual()` 低效地进行数据比较(适用数据量不大)

3. 遍历两者数据内容, 得出一致性结果

## 空 struct{} 的用途

`struct{}` 无需消耗内存

```go
set := make(map[string]struct{})
for _, value := range []string{"apple", "orange", "apple"} {
    // 节省内存
   set[value] = struct{}{}
}
fmt.Println(set)
// Output: map[orange:{} apple:{}]
```

## 函数返回局部变量/逃逸分析

Go编译器会自动决定把一个变量放在栈还是放在堆，编译器会做逃逸分析(escape analysis)；
当发现变量的作用域没有跑出函数范围，就可以在栈上，反之则必须分配在堆。

```go
func fun() *int {    // int类型指针函数
    var tmp := 1
    return &tmp      // 返回局部变量tmp的地址
}

func main() {
    var p *int
    p = fun()
    fmt.Printf("%d\n", *p) // 返回变量V的值1
}
```

### 背景

堆用于动态内存分配，与栈不同，程序需要使用指针在堆中查找数据

页堆: 存储动态数据，最大的内存块, GC(垃圾回收)扫描/标记/回收的地方
    - 驻留内存(resident set)被划分为每个大小为8KB的页，并由一个全局mheap对象管理

栈存储区: 每个Goroutine（G）有一个栈。在这里存储了静态数据，包括函数栈帧，静态结构，原生类型值和指向动态结构的指针
    - 许多对象直接在程序栈上分配, 避免处于垃圾回收的堆内存里

### 程序执行流程

```go
package main

import "fmt"

type Employee struct {
    name   string
    salary int
    sales  int
    bonus  int
}

const BONUS_PERCENTAGE = 10

func getBonusPercentage(salary int) int {
    percentage := (salary * BONUS_PERCENTAGE) / 100
    return percentage
}

func findEmployeeBonus(salary, noOfSales int) int {
    bonusPercentage := getBonusPercentage(salary)
    bonus := bonusPercentage * noOfSales
    return bonus
}

func main() {
    var john = Employee{"John", 5000, 5, 0}
    john.bonus = findEmployeeBonus(john.salary, john.sales)
    fmt.Println(john.bonus)
}
```

```
main函数被保存栈中的"main栈帧"中

每个函数调用都作为一个栈帧块被添加到栈中

包括参数和返回值在内的所有静态变量都保存在函数的栈帧块内

无论类型如何，所有静态值都直接存储在栈中。这也适用于全局范畴

所有动态类型都在堆上创建，并且被栈上的指针所引用。小于32Kb的对象由P的mcache分配。这同样适用于全局范畴

具有静态数据的结构体保留在栈上，直到在该位置将任何动态值添加到该结构中为止。该结构被移到堆上。

从当前函数调用的函数被推入堆顶部

当函数返回时，其栈帧将从栈中删除

一旦主过程(main)完成，堆上的对象将不再具有来自Stack的指针的引用，并成为孤立对象
```


### 代码优化gc

1. 减少对象分配；尽量做到，对象的重用

```go
// 没有注入形参 每次开辟新的切片返回数据
func(r *Reader)Read() ([]byte, error)
// 复用原形参地址内存 避免开辟/回收临时的地址内存数据
func(r *Reader)Read(tar []byte) (int, error)
```

2. 不使用+拼接string 由于采用+来进行string的连接会生成新的对象，降低gc的效率;

3. append操作：正确预估数组的长度的话，最初分配空间做好空间规划操作，有效降低gc的压力，提升代码的效率

### map随机读取底层原理

1. 在java语言里map是按照顺序读取的;golang的map读取是随机性的;

2. golang的map底层使用hash表实现，插入数据位置是随机的，所以遍历过程中新插入的数据不能保证遍历到;

3. 读取过程: 提前取一个随机数，把桶的遍历顺序随机化

### 浅拷贝与深拷贝

1. 若拷贝对象是struct, 成员有Slice/Map数据类型, 浅拷贝下, 对这类成员进行数据变更(非赋予新slice/Map), 会影响原对象

2. 若使用深拷贝, 拷贝对象与被拷贝对象不会互相影响; 拷贝的对象中没有引用类型，只需浅拷贝即可

### defer相关问题

```go
defer func() {
    // do A 
}()

defer func() {
    // panic("B") 
}()

defer func() {
    // do C 
}()

// 执行顺序 do C -> do A -> panic B

```

defer是后进先出顺序下；defer里发生panic, 优先将正常defer方法执行完, 最后再处理发生panic的defer方法

### 切片扩容

源码: `runtime/slice.go  growslice方法`

```go
newcap := old.cap
doublecap := newcap + newcap
if old.len < 1024 {
    newcap = doublecap
} else {
    // Check 0 < newcap to detect overflow
    // and prevent an infinite loop.
    for 0 < newcap && newcap < cap {
        newcap += newcap / 4
    }
    // Set newcap to the requested cap when
    // the newcap calculation overflowed.
    if newcap <= 0 {
        newcap = cap
    }
}
```

结论: slice的长度在超过一个阈值(1024)后便不再翻倍，而是每次以25%的幅度增长，直到满足所需的容量。

### 变量寻址与不可寻址

不可以寻址, 指的是不能通过&获得其地址。

#### 不可寻址特性

- 不可变的; 类型`const`常量不可寻址；基本类型值的字面量`[匿名变量值]`不可寻址`&(123)`

- 临时结果; `&(1+1)`

- 不安全的;

#### 常见例子

```go
_ = &([3]int{1, 2, 3}[0]) // 对数组字面量的索引结果值不可寻址。
_ = &([3]int{1, 2, 3}[0:2]) // 对数组字面量的切片结果值不可寻址。
_ = &([]int{1, 2, 3}[0]) // 对切片字面量的索引结果值却是可寻址的。
_ = &([]int{1, 2, 3}[0:2]) // 对切片字面量的切片结果值不可寻址。
_ = &(map[int]string{1: "a"}[0]) // 对字典字面量的索引结果值不可寻址
```


