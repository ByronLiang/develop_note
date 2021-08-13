package app

import (
	"fmt"
	"sync/atomic"
	"test/tools"
	"test/tools/roundtimer"
	"time"
)

type CountPara struct {
	Round int64
	Total int64
}

func RoundTimerParaPoolSrv() {
	roundtimer.NewRoundTimerPool()
	rt := roundtimer.Pool.Get()
	rt.SetInterval(1 * time.Second).
		SetId(1).
		SetParameters(&CountPara{
			Round: 0,
			Total: 10,
		}).
		SetHandle(paraCallback, nil)
	err := rt.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
	tools.SignCatch(nil, func() {
		roundtimer.Pool.Put(rt)
	})
}

func paraCallback(rt *roundtimer.RoundTimer) {
	if para, ok := rt.GetPara().(*CountPara); ok {
		if para.Total - para.Round == 0 {
			rt.SetInterval(2 * time.Second)
			rt.SetParameters(&CountPara{
				Round: 0,
				Total: 5,
			})
		} else {
			if para.Round == 3 {
				fmt.Println("round 3 counter")
			} else {
				fmt.Printf("timer count: %d \n", para.Total - para.Round)
			}
			atomic.AddInt64(&para.Round, 1)
			rt.SetParameters(para)
		}
		if para.Total == 5 && para.Round == 5 {
			rt.StopWithHandle(func(_ *roundtimer.RoundTimer) {
				fmt.Println("end the timer")
			})
		}
	}
}
