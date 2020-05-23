package main

import (
    "fmt"
    "test/worker"
)

//var mam [2]int

func main()  {
    //var mam = [3]int{1,3,9}
    //mamPoint := &[]int{1,3,9}
    //CopyArray(mam, &mam)
    worker.SliceQuiz()
    //mam = AddNum(20, 21, 10, 10)
    //fmt.Println(mam)
    //res := worker.WindowMax([]int{1,13,3,2,8,4}, 3)
    //fmt.Println(res)
}

func AddNum(newNum int, nums ...int) []int {
    res := append(nums, newNum)
    return res
}

func CopyArray(target [3]int, pointTarget *[3]int)  {
    var(
        obj [3]int
    )
    //内存地址拷贝 同时影响原值
    boo := &target
    boo[1] = 999

    //非内存地址拷贝 不影响原值
    //obj = *pointTarget
    obj = target
    obj[2] = 39

    fmt.Println(boo)
    fmt.Println(obj)
    fmt.Println(target)
    fmt.Println(*pointTarget)
}
