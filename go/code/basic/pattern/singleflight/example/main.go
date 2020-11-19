package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "test/pattern/singleflight"
    "time"
)

var (
    g singleflight.Group
    callTime int32
)

func main()  {
    var wg sync.WaitGroup
    loadReq(&wg, 3, "abc")
    time.Sleep(1 * time.Second)
    loadReq(&wg, 3, "mm")
    wg.Wait()
    fmt.Println(atomic.LoadInt32(&callTime), g.FindMissKeys())
}

func loadReq(wg *sync.WaitGroup, n int, key string) {
    for i := 0; i < n; i++ {
        wg.Add(1)
        go func() {
            v, err := g.Do(key, HandleDo)
            if err == nil {
                switch res := v.(type) {
                case nil:
                    fmt.Println("none match")
                case string:
                    fmt.Println(res)
                }
            }
            wg.Done()
        }()
    }
}

func HandleDo(key string) (interface{}, error, bool) {
    atomic.AddInt32(&callTime, 1)
    if key == "abc" {
        return "cc", nil, true
    }
    return nil, nil, false
}
