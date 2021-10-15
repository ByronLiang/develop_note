package worker

import (
	"fmt"
	"sync"
	"time"
)

var lop chan string

func BasicChan() {
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
	time.Sleep(time.Second) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * time.Second) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}

func sendChan(channel chan string, data string) {
	channel <- data
	fmt.Println("finished send")
}

func getChan(channel chan string) {
	time.Sleep(3 * time.Second)
	word := <-channel
	fmt.Println(word)
}

func BlockSample() {
	c := make(chan int)
	go func() {
		//阻塞消费管道消息 引发维持三秒 main线程等待协程完成
		time.Sleep(2 * time.Second)
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
	time.Sleep(2 * time.Second)
	//维持协程线程不结束 无限循环中读取通道
	for {
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
	time.Sleep(2 * time.Second)
	//对管道提供数值
	out <- 19
	time.Sleep(2 * time.Second)
	//若设置缓冲区作用于主线程投递消息 并在之后建立协程处理消息 才不构成阻塞主线程 并且确保主线程不过早结束
	//不设置缓冲区 产生死锁
	//go f1(out)
	fmt.Println("end")
	//main线程结束 协程线程同样结束 即使使用for循环维持线程
}

func channelConsumer(in chan int) {
	fmt.Println(<-in)
}

func NoneBlocking() {
	//设置容量5的缓冲区通道
	out := make(chan int, 1)
	//投递生产
	//out <- 2
	//异步消费 非阻塞
	//go channelConsumer(out)
	go func() {
		time.Sleep(4 * time.Second)
		out <- 100
	}()
	go func() {
		time.Sleep(2 * time.Second)
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

func RoutineBreak() {
	cc := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
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

func selectWorker(target chan int, len int) {
	for i := 0; i < len; i++ {
		select {
		case target <- 1:
		case target <- 0:
		}
	}
	close(target)
}

func CreateBitNum() {
	var list []int
	contain := make(chan int)
	go selectWorker(contain, 5)
	for data := range contain {
		list = append(list, data)
	}
	fmt.Println(list)
}

func TimeTicker() {
	counter := 0
	//kk := time.NewTicker(1 * 1e9)
	//kk.Stop()
	// 时间类型的输出管道
	//times := kk.C
	ticker := time.Tick(2 * time.Second)
	overTimer := time.NewTimer(5 * time.Second)
	defer overTimer.Stop()
	// 一次性计时器
	// time.After(5 * time.Second)
	//监听定时器通道 异步
tickerLoop:
	for {
		select {
		case <-ticker:
			counter++
			fmt.Println("time reach", counter)
		case <-overTimer.C:
			fmt.Println("time over")
			// 针对for 跳出循环, 可执行for循环以外的业务
			break tickerLoop
			// 直接跳出函数
			return
			//case <-time.After(5 * time.Second): // 禁止: 每次select时，都会重新初始化一个全新的计时器, 无法被回收 引发内存泄露
			//    return
		}
	}
	fmt.Println("all end at last")
}

func CheckChanFullDemo() {
	input := make(chan interface{}, 5)
	input <- "one"
	if checkChanIsFull(input, "two") {
		fmt.Println("chan is full")
	} else {
		fmt.Println("chan is not full")
		close(input)
		for data := range input {
			fmt.Println(data.(string))
		}
	}
}

/**
利用select检测 缓冲通道是否已满
*/
func checkChanIsFull(input chan<- interface{}, data interface{}) bool {
	select {
	case input <- data:
		return false
	default:
		return true
	}
}

func CloseSign() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err.(string))
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	data, done := make(chan int), make(chan struct{})
	go watchSign(data, done, &wg)
	data <- 10
	data <- 1
	close(data)
	done <- struct{}{}
	// 对chan仍会发出信号
	close(done)
	//panic("abc")
	wg.Wait()
	fmt.Println("end")
}

func watchSign(data <-chan int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	for {
		select {
		case res, ok := <-data:
			if ok {
				fmt.Println(res)
			}
		case _, ok := <-done:
			if ok {
				fmt.Println("receive")
			} else {
				fmt.Println("done handle")
				return
			}
		}
	}
}

func VerifyChanClose() {
	demo := make(chan struct{})
	fmt.Println(verifyChan(demo))
	close(demo)
	// 关闭chan, 每次在select里都会进入输出语句里
	fmt.Println(verifyChan(demo))
	// 将chan设置为nil 读写都阻塞; 避免进入select的输出语句里
	demo = nil
	fmt.Println(verifyChan(demo))
}

func verifyChan(c chan struct{}) (isClose bool) {
	select {
	case _, open := <-c:
		isClose = !open
	default:
		isClose = false
	}
	return
}
