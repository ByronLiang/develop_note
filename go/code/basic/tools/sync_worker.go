package tools

import (
    "fmt"
    "sync"
    "time"
)

type wallet struct {
    sync.RWMutex
    total   int
}

func LockSample() {
    target := &wallet{
        total: 10,
    }
    go func(obj *wallet) {
        time.Sleep(1 * time.Second)
        obj.RLock()
        defer obj.RUnlock()
        //time.Sleep(1 * time.Second)
        fmt.Println("read: ", obj.total)
    }(target)
    go func(obj *wallet) {
        obj.Lock()
        defer obj.Unlock()
        time.Sleep(2 * time.Second)
        fmt.Println("write: ", obj.total)
        obj.total += 10
    }(target)
    time.Sleep(3 * time.Second)
}

var timerPool sync.Pool

func TimeNotify()  {
    timer := AcquireTimer(3 * time.Second)
    go timeProcess(timer, "cc")
    zz := AcquireTimer(2 * time.Second)
    go timeProcess(zz, "zz")
    time.Sleep(6 * time.Second)
}

func timeProcess(tar *time.Timer, name string)  {
    for {
        select {
        case <-tar.C:
            fmt.Println("time to take action", name)
            ReleaseTimer(tar)
            break
        default:

        }
    }
}

// AcquireTimer returns time from pool if possible.
func AcquireTimer(d time.Duration) *time.Timer {
    v := timerPool.Get()
    if v == nil {
        return time.NewTimer(d)
    }
    tm := v.(*time.Timer)
    if tm.Reset(d) {
        // active timer?
        return time.NewTimer(d)
    }
    return tm
}

// ReleaseTimer returns timer into pool.
func ReleaseTimer(tm *time.Timer) {
    if !tm.Stop() {
        return
    }
    timerPool.Put(tm)
}
