package main

import (
    "fmt"
    "log"
    "test/tools"
)

func main()  {
    //tools.CasualTimeCount(10)
    //tools.LockSample()
    //tools.TimeNotify()
    //hostSample()
    //DecoratorSample()
    genCodeExample()
}

func DecoratorSample()  {
    //type MyFoo func(int, int, int) int
    origin := foo
    // 装饰器模式 对传参按照修饰方法进行处理
    _ = tools.Decorator(&origin, douFoo)
    ews := origin(1, 2, 3)
    log.Print(ews)
}

func douFoo(a, b, c int) int {
    fmt.Printf("%d, %d, %d \n", a, b, c)
    return a*2 + b*1 + c*3
}

func foo(a, b, c int) int {
    fmt.Printf("%d, %d, %d \n", a, b, c)
    return a + b + c
}

func hostSample() {
    //fmt.Println(tools.GetLocalHost())
    fmt.Println(tools.IpWithNames())
    fmt.Println(tools.FindIpByName("eth1"))
}

func genCodeExample()  {
    gen := tools.CodeGen{
       InitCode:   27995,
       NumSize:    1000,
       PrefixSize: 2,
       NumLength: 3,
    }
    fmt.Println(gen.GenTotalCode(10))
    fmt.Println(gen.GenTotalCode(10))
    fmt.Println(gen.ReflectCode())
    fmt.Println(gen.GenCodeNum("ZZZ"))
}
