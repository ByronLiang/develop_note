# 协程 与 通道 笔记

## chan 底层

一个环形队列和两个链表展开

## chan关闭后相关操作现象

1. 写操作: 会引发异常
2. 读操作: 

- 如果 chan 关闭前，buffer 内有元素已经被读完，chan 内无值，接下来所有接收的值都会`非阻塞`直接成功，返回 channel 元素的零值，但是第二个 bool 值一直为 false。

- chan 关闭后，并被赋值为`nil`，则接收值操作会一直处于`阻塞`；(原理: 将当前chan 进行未初始化处理)

```go

// 未初始化chan
// 读写都会产生阻塞
var cc chan int

// 已初始化chan
c := make(chan int)

```