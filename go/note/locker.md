# Golang 锁应用

## 自旋锁与互斥锁背景

### 应用背景

互斥锁用于临界区持锁时间比较长的操作: 业务复杂；单核处理器

自旋锁就主要用在临界区持锁时间非常短且CPU资源不紧张的情况

### 总结

1. 自旋锁与互斥锁都是为了实现保护资源共享的机制。
2. 无论是自旋锁还是互斥锁，在任意时刻，都最多只能有一个保持者。
3. 获取互斥锁的线程，如果锁已经被占用，则该线程将进入睡眠状态；获取自旋锁的线程则不会睡眠，而是一直循环等待锁释放。

## 互斥锁

1.互斥锁有两种操作，获取锁和释放锁
2.当有一个goroutine获取了互斥锁后，任何goroutine都不可以获取互斥锁，只能等待这个goroutine将互斥锁释放
3.互斥锁适用于读写操作数量差不多的情况

## 读写锁

`sync.RWMutex`

1.读写锁有四种操作 读上锁 读解锁 写上锁 写解锁
2.写锁最多有一个，读锁可以有多个; 写锁的优先级高于读锁
3.一个goroutine获得写锁时，其他goroutine不可以获得读锁/写锁，需要等待写锁释放
5.一个goroutine获得读锁时，其他goroutine可以获得读锁，不能获得写锁; (写锁的优先级高于读锁, 有效避免长时间阻塞情况)
6.读写锁适用于读多写少的情景

## 使用总结

1.在单纯的只是获取锁和释放锁时，互斥锁的用时要少一些，这主要是因为多个线程同时获取读写锁的情况比较少出现
2.golang底层实现上，互斥锁确实要比读写锁的性能要好一些，这主要是因为读写锁的底层实现其实是互斥锁加上计数器
3.在增强协程互相冲突的效果后，读写锁的性能要明显高于互斥锁。
