package app

import (
	"fmt"
	"test/tools"
	"test/tools/roundtimer"
	"time"
)

func RoundTimerPoolSrv() {
	pool := roundtimer.DefaultRoundTimerPool
	rt := pool.Get()
	rt.SetInterval(2 * time.Second).
		SetId(1).
		SetHandle(roundTimerCallback, roundTimerResetAfter)
	err := rt.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	tools.SignCatch(nil, func() {
		pool.Put(rt)
	})
}

func roundTimerCallback(rt *roundtimer.RoundTimer)  {
	fmt.Println("timer count")
}

func roundTimerResetAfter(rt *roundtimer.RoundTimer)  {
	fmt.Println("after reset timer count")
}
