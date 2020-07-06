package main

import (
    "fmt"
    "test/alog"
)

func main()  {
    //worker.ByteBuff()
    //worker.EncodeJsonSample()
    res := alog.BuildStatusTable("ABCDABD")
    fmt.Println(res)
}
