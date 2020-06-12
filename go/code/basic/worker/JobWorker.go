package worker

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

//使用空结构体，作为任务结束的通知
var notifyCh = make(chan struct{}, 50)

func worker(n int, jobs <-chan int, results chan<- int) {
    for x := range jobs {
        fmt.Printf("协程: %d处理开始\n", n)
        //模拟处理时间
        time.Sleep(time.Second)
        results <- x * 2
        //每完成一个job，计数一次
        notifyCh <- struct{}{}
        fmt.Printf("协程: %d处理结束\n", n)
    }
}

func InitJob()  {
    //预留缓冲
    jobs := make(chan int, 20)
    results := make(chan int, 20)
    //使用一个协程生成50个任务；
    go func(xx string) {
        for i := 1; i <= 10; i++ {
            jobs <- i
        }
        fmt.Println(xx + " none job created")
        close(jobs)
    }("kk")
    //使用3个goroutine处理任务；
    for j := 1; j <= 2; j++ {
        go worker(j, jobs, results)
    }
    //使用一个goroutine去notifyCh取值，通知通道有50次,说明50个任务处理结束
    go func() {
        for w := 1; w <= 10; w++ {
            <-notifyCh
        }
        //关闭result通道
        close(results)
    }()
    //输出
    for r := range results {
        fmt.Println(r)
    }
}

var (
    wg      sync.WaitGroup
    mutex   sync.Mutex
)

func initCounter(i int)  {
    defer func() {
        fmt.Println("the last point", i)
        wg.Done()
    }()
    for j := 1; j < 3; j ++ {
        val := j
        fmt.Println("thread prepared to exchange", i, j)
        //切换其他协程
        runtime.Gosched()
        //当再次被切换回来 再进行此流程
        val ++
        fmt.Println("thread had came back", i, val)
    }
}

func lockCounter(i int) {
    defer func() {
        fmt.Println("the last point", i)
        wg.Done()
    }()
    for j := 1; j < 3; j ++ {
        //加锁 确保即便切换其他协程 也确保流程不会被中断
        //当进入到其他协程 发现无法获取锁 重新换回原协程进行继续处理
        mutex.Lock()
        val := j
        fmt.Println("thread prepared to exchange", i, j)
        //切换其他协程
        runtime.Gosched()
        val ++
        fmt.Println("thread had came back", i, val)
        mutex.Unlock()
    }
}

func ThreadExchange()  {
    runtime.GOMAXPROCS(1)
    wg.Add(2)
    //go initCounter(1)
    //go initCounter(2)
    go lockCounter(1)
    go lockCounter(2)
    wg.Wait()
    //time.Sleep(time.Millisecond * 500)
    fmt.Println("end")
}
