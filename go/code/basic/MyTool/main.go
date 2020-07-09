package main

import (
    "fmt"
    "test/alog"
)

func main()  {
    //tools.CasualTimeCount(10)
    res := alog.KmpSearch("hello", "ll")
    fmt.Println(res)
}
