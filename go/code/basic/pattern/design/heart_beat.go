package design

import (
    "fmt"
    "time"
)

func heartBeatSys(d time.Duration, done chan interface{}) chan time.Time {
    beat := make(chan time.Time)
    // 一次性定时器
    // time.NewTimer(d)
    // 打点器 间隔d秒触发
    t := time.NewTicker(d)
    go func() {
        for {
            select {
            case <-done:
                t.Stop()
                fmt.Println("beat end")
                return
            case res := <- t.C:
                beat <- res
           }
        }
    }()
    return beat
}

func HeartBeat()  {
    var done = make(chan interface{})
    res := heartBeatSys(2 * time.Second, done)
    time.AfterFunc(5 * time.Second, func() {
        close(done)
    })
    for {
        select {
        case r := <-res:
            fmt.Println("beat coming: ", r.Second())
        case <-time.After(7 * time.Second):
            // 超时处理
            fmt.Println("system timeout")
            return
       }
    }
}
