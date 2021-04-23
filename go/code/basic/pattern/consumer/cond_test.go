package consumer

import (
	"fmt"
	"sync"
	"testing"
)

func TestCondConsumer(t *testing.T) {
	CondConsumer()
}

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	var wg sync.WaitGroup
	// 建立五个不同消费数量的消费者
	// 总共只能消费15个item
	for n := 5; n > 0; n-- {
		wg.Add(1)
		go func(n int) {
			items := q.GetDefineSize(n)
			fmt.Printf("%2d: %2d\n", n, items)
			wg.Done()
		}(n)
	}
	for i := 0; i < 20; i++ {
		q.Put(i)
	}
	wg.Wait()
}
