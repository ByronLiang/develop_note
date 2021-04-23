package consumer

import (
	"fmt"
	"sync"
)

const (
	QueueLen = 10
	WorkLen  = 5
)

type DispatcherHandler interface {
	Do(data interface{})
}

type Dispatcher struct {
	DispatcherHandler
	queue   chan interface{} //数据队列
	workLen int              //协程数
	exit    chan bool        //退出协程
	close   bool
	wg      *sync.WaitGroup //用于保证所有的任务都已完成才关闭调度器
}

//新建一个调度器
func NewDispatcher(queueLen int, workLen int, handler DispatcherHandler) *Dispatcher {
	d := new(Dispatcher)
	if queueLen == 0 {
		queueLen = QueueLen
	}
	if workLen == 0 {
		workLen = WorkLen
	}
	d.queue = make(chan interface{}, queueLen)
	d.workLen = workLen
	d.DispatcherHandler = handler
	d.exit = make(chan bool)
	d.wg = new(sync.WaitGroup)

	for i := 0; i < d.workLen; i++ {
		go d.start()
	}
	return d
}

//启动协程
func (d *Dispatcher) start() {
	defer func() {
		// 避免雪崩
		if err := recover(); err != nil {
			return
		}
		d.wg.Wait()
	}()
	for {
		select {
		case j, ok := <-d.queue:
			if !ok {
				return
			}
			d.DispatcherHandler.Do(j)
			d.wg.Done()
		case <-d.exit:
			fmt.Println("dispatch exist")
			return
		}
	}
}

//把任务添加到队列
func (d *Dispatcher) Add(data interface{}) {
	if d.close {
		return
	}
	d.wg.Add(1)
	d.queue <- data
}

//优雅退出调度器
func (d *Dispatcher) Exit() {
	d.wg.Wait()
	for i := 0; i < d.workLen; i++ {
		d.exit <- true
	}
	d.close = true
}
