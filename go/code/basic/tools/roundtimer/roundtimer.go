package roundtimer

import (
	"fmt"
	"sync"
	"time"
)

type RoundTimer struct {
	id int64
	parameters interface{}
	interval time.Duration
	running  bool
	timerHandle func(rt *RoundTimer)
	timerAfterHandle func(rt *RoundTimer)
	timer    *time.Timer

	sync.Mutex
}

// sample对象
type Parameter struct {
	id int64
	round int
}

func NewRoundTimer() *RoundTimer {
	return &RoundTimer{}
}

// 初次重置计时器回调
func (rt *RoundTimer) initTimer(afterInitHandle func()) {
	rt.timer = time.AfterFunc(rt.interval, rt.resetHandle)
	// 初次配置计时器的回调设置
	afterInitHandle()
}

// 重置回调与续期计时器
func (rt *RoundTimer) resetHandle() {
	if !rt.running {
		return
	}
	// 执行定时操作
	// TODO: 新增参数进行配置同步/异步执行
	rt.timerHandle(rt)
	if !rt.running {
		return
	}
	// 进行续期
	rt.timer = time.AfterFunc(rt.interval, rt.resetHandle)
	// 完成续期的操作
	rt.timerAfterHandle(rt)
}

// 首次启动定时器
func (rt *RoundTimer) Start(startHandle func()) error {
	rt.Lock()
	defer rt.Unlock()
	if rt.running {
		return fmt.Errorf("id: %d timer already running", rt.id)
	}
	rt.initTimer(startHandle)
	return nil
}

// 重置定时间隔
func (rt *RoundTimer) ResetInterval(interval time.Duration) error {
	rt.Mutex.Lock()
	defer rt.Mutex.Unlock()

	if !rt.running {
		return fmt.Errorf("id: %d timer already stop", rt.id)
	}

	rt.interval = interval
	rt.timer.Reset(interval)
	return nil
}

func (rt *RoundTimer) Stop() {
	rt.Mutex.Lock()
	defer rt.Mutex.Unlock()

	rt.running = false
	rt.timer.Stop()
}

// 重置对象参数
func (rt *RoundTimer) Reset() {
	rt.Mutex.Lock()
	defer rt.Mutex.Unlock()
	rt.id = 0
	rt.running = false
	rt.interval = 0 * time.Second
	rt.timerHandle = nil
	rt.timerAfterHandle = nil
	rt.parameters = nil
}