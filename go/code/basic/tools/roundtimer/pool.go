package roundtimer

import "sync"

// 池化技术: 池化频繁创建的对象
// 复用原有对象, 减少创建和销毁的动作

var DefaultRoundTimerPool = NewRoundTimerPool()

type roundTimerPool struct {
	pool *sync.Pool
}

// 对象池 避免频繁创建与回收对象
func NewRoundTimerPool() *roundTimerPool {
	return &roundTimerPool{
		pool: &sync.Pool{
			New: func() interface{} {
				// 创建原始对象
				return NewRoundTimer()
			},
		},
	}
}

func (r *roundTimerPool) Get() *RoundTimer {
	return r.pool.Get().(*RoundTimer)
}

func (r *roundTimerPool) Put(rt *RoundTimer) {
	rt.Reset()
	r.pool.Put(rt)
}
