package main

import (
    "fmt"
)

type dog struct {
    name string
    age int
    hobby string
}

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

func MyStruct()  {
    none := struct {
        area string
    }{
        area: "abc",
    }
    cong := dog{
        name:  "cong",
        age:   10,
        hobby: "play",
    }
    //同时拷贝内存地址与拷贝数据
    boy := &cong
    //新开辟内存地址拷贝数据
    //boy2 := cong
    boy.name = "kk"
    fmt.Println(cong.name, cong, &cong.name)
    fmt.Println(boy.name, boy, &boy.name)
    fmt.Println(none, &none.area)
}
