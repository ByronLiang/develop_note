package main

import (
	"fmt"
	"test/tools/roundtimer"
	"test/tools/roundtimer/example/app"
)

func main()  {
	//app.RoundTimerPoolSrv()
	app.RoundTimerParaPoolSrv()
}



func roundTimerCallback(rt *roundtimer.RoundTimer)  {
	fmt.Println("timer count")
}

func roundTimerResetAfter(rt *roundtimer.RoundTimer)  {
	fmt.Println("after reset timer count")
}
