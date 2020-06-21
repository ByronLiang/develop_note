package main

import (
    "fmt"
    "math/rand"
    "test/worker"
)

//var mam [2]int

func main()  {
    //var mam = [3]int{1,3,9}
    //mamPoint := &[]int{1,3,9}
    //defer CopyArray(mam, &mam)
    //defer worker.DeferSample()
    //items := []int{10, 20, 30, 40, 50}
    //worker.TestQuiz(items)
    //worker.PointSlice()
    //worker.BasicInit()
    //worker.SliceQuiz()
    //testArrayPara()
    //fmt.Println(mam)
    //res := worker.WindowMax([]int{1,13,3,2,8,4}, 3)
    //fmt.Println(res)
    //worker.NullSlice()
    testRemoveIndexData()
    //worker.CopySlice()
    //worker.ModifiedSlice()
    //worker.RoundNum(5)
    //worker.PrintRoundNum(5)
}

func AddNum(newNum int, nums ...int) []int {
    res := append(nums, newNum)
    return res
}

func CopyArray(target [3]int, pointTarget *[3]int)  {
    var(
        obj [3]int
        point *[3]int
    )
    //内存地址拷贝 同时影响原值
    boo := &target
    boo[1] = 999

    //非内存地址拷贝 不影响原值
    //obj = *pointTarget
    obj = target
    obj[2] = 39

    //指针地址拷贝 对原值有影响
    point = pointTarget
    point[1] = 100

    fmt.Println(boo)
    fmt.Println(obj)
    fmt.Println(target)
    fmt.Println(*pointTarget, point)
}

func testRemoveIndexData()  {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("something error: ", err)
        }
    }()
    data := []string{"apple", "banana", "cat", "dog", "bird"}
    fmt.Println(data, cap(data))
    res := worker.RemoveIndexData(13, data)
    //若异常 无法进行 直接从recover() 完成函数执行
    fmt.Println(res, cap(res), "qq")
    fmt.Println(data, cap(data))
}

/**
函数动态形参长度处理实例
 */
func testArrayPara()  {
    data1 := AddNum(rand.Intn(10), 21, 10, 10)
    sample := []int{20, 10, 33}
    data2 := AddNum(rand.Intn(10), sample...)
    fmt.Println(data1, data2)
}
