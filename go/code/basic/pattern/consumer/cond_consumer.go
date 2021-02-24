package consumer

import (
	"fmt"
	"sync"
	"time"
)

var (
	m sync.Mutex
	queue = make([]struct{}, 0, 10)
	c = sync.NewCond(&m)
)

func CondConsumer()  {
	go consumerHandle()
	for i := 0; i < 5; i++ {
		m.Lock()
		if len(queue) == 2 {
			fmt.Println("waiting")
			// 阻塞中 c.Signal() 解除阻塞
			c.Wait()
		}
		queue = append(queue, struct {}{})
		fmt.Println("add data into queue")
		m.Unlock()
	}
}

func consumerHandle()  {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("pending lock")
		m.Lock()
		fmt.Println("get lock")
		if len(queue) > 0 {
			queue = queue[1:]
			fmt.Println("consume data in queue")
			// 解除c.Wait()阻塞
			c.Signal()
		}
		m.Unlock()
		fmt.Println("finished consume")
	}
}
