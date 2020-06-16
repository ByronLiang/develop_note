package worker

import (
    "fmt"
    "time"
)

var lop chan string

func BasicChan()  {
    cc := make(chan string)
    //go sendChan(cc, "hope")
    //go sendChan(cc, "hope123")
    go getChan(cc)

    //word := <- cc
    fmt.Println("start")
    cc <- "xx"
    //go longWait()
    fmt.Println("end")
    //time.Sleep(4 * 1e9)
}

func longWait() {
    fmt.Println("Beginning longWait()")
    time.Sleep(1 * 1e9) // sleep for 5 seconds
    fmt.Println("End of longWait()")
}

func shortWait() {
    fmt.Println("Beginning shortWait()")
    time.Sleep(2 * 1e9) // sleep for 2 seconds
    fmt.Println("End of shortWait()")
}

func sendChan(channel chan string, data string)  {
    channel <- data
    fmt.Println("finished send")
}

func getChan(channel chan string) {
    time.Sleep(3 * time.Second)
    word := <-channel
    fmt.Println(word)
}

func BlockSample()  {
    c := make(chan int)
    go func() {
        //阻塞消费管道消息 引发维持三秒 main线程等待协程完成
        time.Sleep(2 * 1e9)
        x := <-c
        fmt.Println("received", x)
        //协程线程结束
    }()
    fmt.Println("sending", 10)
    //利用协程线程阻塞对main主线程实现阻塞
    //若使用缓冲 无需被协程线程阻塞 直接结束主线程 导致协程线程无法输出数据(被结束)
    c <- 10
    //协程线程结束 管道消息消费结束 再生成管道消息，无线程进行消费 形成死锁
    //c <- 100
    //管道通信使用协程线程 不造成main主线程阻塞 主线程结束 协程线程无法输出数据(被结束了)
    //go func() {
    //    c <- 10
    //}()
    fmt.Println("sent", 10)
}

func f1(in chan int) {
    // 初始化协程 sleep两秒
    time.Sleep(2 * 1e9)
    //维持协程线程不结束 无限循环中读取通道
    for  {
        fmt.Println("blocked")
        //阻塞 等待读取消息
        fmt.Println(<-in)
        fmt.Println("finished blocked")
        //完成消息消费 重新进行阻塞等待读取消息
    }
}

func Blocking() {
    //初始化 非缓冲区通道
    out := make(chan int)
    //初始化协程
    //协程线程随着main线程结束而结束
    go f1(out)
    //对管道提供数值
    out <- 2
    //协程里读取消息 解除阻塞 进入sleep 2s
    time.Sleep(2 * 1e9)
    //对管道提供数值
    out <- 19
    time.Sleep(2 * 1e9)
    //若设置缓冲区作用于主线程投递消息 并在之后建立协程处理消息 才不构成阻塞主线程 并且确保主线程不过早结束
    //不设置缓冲区 产生死锁
    //go f1(out)
    fmt.Println("end")
    //main线程结束 协程线程同样结束 即使使用for循环维持线程
}

func channelConsumer(in chan int)  {
    fmt.Println(<-in)
}

func NoneBlocking()  {
    //设置容量5的缓冲区通道
    out := make(chan int, 1)
    //投递生产
    //out <- 2
    //异步消费 非阻塞
    //go channelConsumer(out)
    go func() {
        time.Sleep(4 * 1e9)
        out <- 100
    }()
    go func() {
        time.Sleep(2 * 1e9)
        out <- 222
    }()
    fmt.Println("pending")
    //即便设置缓冲区 由于投递使用协程线程 主线程里取出消息 仍被阻塞
    //通道先进先出 与协程顺序无关
    kk1 := <-out
    fmt.Println(kk1, time.Now().Unix())
    kk2 := <-out
    fmt.Println(kk2, time.Now().Unix())
    fmt.Println("end")
    //预留时间让协程线程处理数据 避免主线程过早结束
    //time.Sleep(2 * 1e9)
}

func RoutineBreak()  {
    cc := make(chan int)
    go func() {
        time.Sleep(2 * 1e9)
        cc <- 100
        cc <- 9
        cc <- 20
        //不关闭通道 会引发主线程阻塞 导致死锁
        close(cc)
    }()
    fmt.Println("starting")
    //for {
    //   num, res := <- cc
    //   if !res {
    //       fmt.Println("chan is closed")
    //       break
    //   }
    //   fmt.Println(num, time.Now().Unix())
    //}
    for data := range cc {
        fmt.Println(data)
    }
    fmt.Println("end")
}

func selectWorker(target chan int, len int)  {
    for i := 0; i < len; i++ {
        select {
        case target <- 1:
        case target <- 0:
        }
    }
    close(target)
}

func CreateBitNum()  {
    var list []int
    contain := make(chan int)
    go selectWorker(contain, 5)
    for data := range contain {
        list = append(list, data)
    }
    fmt.Println(list)
}

func TimeTicker()  {
    counter := 0
    //kk := time.NewTicker(1 * 1e9)
    //kk.Stop()
    // 时间类型的输出管道
    //times := kk.C
    ticker := time.Tick(2 * time.Second)
    //一次性计时器
    time.After(5e9)
    //监听定时器通道 异步
    for {
       select {
       case <-ticker:
           counter ++
           fmt.Println("time reach", counter)
           if counter == 3 {
               fmt.Println("reached count")
               return
           }
       }
    }
}