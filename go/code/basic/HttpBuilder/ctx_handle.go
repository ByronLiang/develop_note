package HttpBuilder

import (
    "context"
    "fmt"
    "time"
)

type Rsp struct {
    Data    interface{}
    Code    int
}
const MAX_WORKING_DURATION = 2 * time.Second
//达到实际工作时长后，manager可以提前叫停
const ACTUAL_WORKING_DURATION = 10 * time.Second

const SUCCESS int = 1
const ERROR int = 2

func CheckDemo() {
    ctxWithCancel, cancel := context.WithTimeout(context.Background(), MAX_WORKING_DURATION)

    go workerHandle(ctxWithCancel, "[1]", 1 * time.Second)
    go workerHandle(ctxWithCancel, "[2]", 3 * time.Second)
    
    go manager(cancel)

    <-ctxWithCancel.Done()
    //暂停1秒便于协程的打印输出
    time.Sleep(1 * time.Second)
    fmt.Println("company closed")
}

func manager(cancel func()) {
    time.Sleep(ACTUAL_WORKING_DURATION)
    fmt.Println("manager called cancel()")
    cancel()
}

func worker(ctxWithCancel context.Context, name string, t time.Duration) {
    for {
        select {
        case <-ctxWithCancel.Done():
            fmt.Println(name, "return for ctxWithCancel.Done()")
            return
        default:
            fmt.Println(name, "work done")
        }
        time.Sleep(t)
    }
}

func workerHandle(ctxWithCancel context.Context, name string, t time.Duration) {
    for {
        select {
        case <-ctxWithCancel.Done():
            fmt.Println(name, "return for ctxWithCancel.Done()")
            return
        case <-time.After(t):
            fmt.Println(name, "work done")
            return
        }
    }
}

func FetchSys()  {
    var data []interface{}
    var RevBuf = make(chan Rsp)
    ctx := context.WithValue(context.Background(), "name", "qq")
    ctx = context.WithValue(ctx, "pp", "ll")
    ctxCancel, cancel := context.WithCancel(ctx)
    go FetchHandle(ctxCancel, "order", RevBuf, 3 * time.Second, Fetch("zz", true))
    go FetchHandle(ctxCancel, "product", RevBuf, 1 * time.Second, Fetch("mm", false))
    for rev := range RevBuf {
        if rev.Code == ERROR {
            cancel()
            close(RevBuf)
        } else {
            fmt.Println(rev.Code)
            data = append(data, rev.Data)
        }
        if len(data) == 2 {
            close(RevBuf)
        }
    }
    fmt.Println(data)
}

func FetchHandle(ctxCancel context.Context, name string, cal chan Rsp,  t time.Duration, f Rsp) {
    for {
        select {
        case <-ctxCancel.Done():
            fmt.Println(name, "return for ctxWithCancel.Done()")
            return
        case <-time.After(t):
            cal <- f
            fmt.Println(name, "work done", ctxCancel.Value("name"))
            return
        }
    }
}

func Fetch(name string, isSuccess bool) Rsp {
    var rsp Rsp
    if isSuccess {
        rsp.Code = SUCCESS
    } else {
        rsp.Code = ERROR
    }
    rsp.Data = name
    return rsp
}