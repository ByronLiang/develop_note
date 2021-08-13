package app

import (
	"fmt"
	"test/tools"
	"test/tools/roundtimer"
	"time"
)

func RoundTimerPoolSrv() {
	roundtimer.NewRoundTimerPool()
	rt := roundtimer.Pool.Get()
	rt.SetInterval(2 * time.Second).
		SetId(1).
		SetHandle(roundTimerCallback, roundTimerResetAfter)
	err := rt.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	tools.SignCatch(nil, func() {
		roundtimer.Pool.Put(rt)
	})
}

func roundTimerCallback(rt *roundtimer.RoundTimer)  {
	fmt.Println("timer count")
}

func roundTimerResetAfter(rt *roundtimer.RoundTimer)  {
	fmt.Println("after reset timer count")
}
