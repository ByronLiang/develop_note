package consumer

import (
	"sync"
)

type Item = int

type Queue struct {
	items     []Item
	*sync.Cond
}

func NewQueue() *Queue {
	q := new(Queue)
	q.Cond = sync.NewCond(&sync.Mutex{})
	return q
}

func (q *Queue) Put(item Item) {
	q.L.Lock()
	defer q.L.Unlock()
	q.items = append(q.items, item)
	// 完成生产, 对任一处于阻塞的goroutine进行唤醒
	q.Signal()
	// 唤醒全部条件阻塞的goroutine
	//q.Broadcast()
}

func (q *Queue) GetDefineSize(n int) []Item {
	q.L.Lock()
	defer q.L.Unlock()
	// 当前队列长度不满足指定消费长度, 处于阻塞
	for len(q.items) < n {
		q.Wait()
	}
	items := q.items[:n]
	q.items = q.items[n:]
	return items
}
