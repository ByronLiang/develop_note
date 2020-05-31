package main

import (
    "fmt"
)

func main()  {
    var lol int
    lol = 1
    if val:= 100; val > 99 && lol != 1 {
        println( "ss", val)
    } else {
        fmt.Println("kk", val)
    }
    /**
    switch 使用
     */
    //MySwitch()

    /**
    无限循环
     */
    //ForTarget()
}

func MySwitch() {
    num := 2
    switch num {
    case 1, 2:
        //继续下面的case (除了default) 并输出逻辑
        println("ss"); fallthrough
    case 3:
        println("33")
    default:
        println("no match")
    }
}

func ForTarget() {
    for i := 0; ; i++ {
        var v int
        fmt.Println("Value of i is now:", i)
        fmt.Printf("%d, %d \n", v, i)
        v = 5
        if i == 5 {
            break
        }
    }
STT:
    for i := 0; i <= 3; i++ {
        for j := 0; j <= 3; j++ {
            if j == 2 {
                continue STT
            }
            fmt.Printf("i is: %d, and j is: %d\n", i, j)
        }
    }

    a := 1
    goto TARGET // compile error
TARGET:
    var b = 10
    b += a
    fmt.Printf("a is %d *** b is %d", a, b)
}
