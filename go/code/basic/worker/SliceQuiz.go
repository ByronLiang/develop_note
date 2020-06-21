package worker

import (
    "fmt"
)

func BasicInit()  {

    /**
    初始化空切片
    初始化指定下标1 与 2 的值为 2 与 9 的切片
     */
    var lpp []int
    lpp = []int{1:2, 2:9}

    // 初始化内存地址 默认不返回值
    tt := new([10]string)[0:1]
    // 引发异常 tt[1] = "qq"
    tt[0] = "pq"
    tt = append(tt, "w", "qq")
    //初始化10个元素 容量是20
    kk := make([]int, 10, 20)
    /**
    不能访问超出初始化长度的地址[容量未合并到切片的长度里, 不能进行访问]
    kk[11] = 11 会引发异常
    只能通过后期操作, 合并才能进行访问 kk = append(kk, 11)
     */
    kk[9] = 11
    fmt.Println(len(kk), cap(kk), tt, tt[0], kk, lpp)
    kk1 := []int{2,5,6,88,22,11,23,45,67}
    //重新分片操作
    // 向后以三位 但是末尾没有移动。切片只能向后移动
    kk2 := kk1[3:6:6]
    /**
    第三个参数限制容量 否则剩余容量都被分配到kk2里
    kk2 := kk1[3:6:7]
     */
    //当切片进行扩容时 开辟新的内存地址 之后的值变更不影响原切片
    kk2 = append(kk2, 101)
    //基于内存地址变更值 同时改变原分片值

    kk2[0] = 10
    fmt.Println(cap(kk1), kk2, cap(kk2), kk2[:cap(kk2)], kk1)

    //length := len(kk1)
    //for i := 1; i <= length; i++ {
    //    kk1 = kk1[1:]
    //    fmt.Println(kk1, len(kk1), cap(kk1), i)
    //}
}

func SliceTest() {
    //var blo []int

    var identifier = []int{1,3,2,0,0}
    l1 := len(identifier)
    aa := identifier[:len(identifier) - 2]
    identifier = identifier[:cap(identifier)]
    fmt.Println(l1)
    fmt.Println(aa)
    fmt.Println(identifier)
    fmt.Println(len(identifier))
}

func SliceQuiz()  {
    var no []int
    no = append(no, 1,2,3,4)
    blo := make([]int, 5)
    no1 := no[1:3]
    fmt.Println(no1, "", cap(no1))
    no1 = no1[0:3]
    fmt.Println(blo, cap(no1), no, no1)
}

func TestQuiz(items []int)  {
    for index, item := range items {
        fmt.Println(&item, &items[index])
        items[index] *= 2
    }
    fmt.Println(items)
}

/**
当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
 */
func DeferSample()  {
    for i := 0; i < 3; i++ {
        defer fmt.Printf("%d ", i)
    }
}

func PointSlice() {
    //指针初始化切片数据结构 长度为0 一个空切片
    var emptySlice = new([]int)
    //使用make初始化切片 初始长度为5 默认值为0
    var initSlice = make([]int, 5)
    //指针初始化数组数据结构 长度为5 默认值为0 不能初始化具体值
    var emptyArray = new([5]int)

    fmt.Println(len(*emptySlice), cap(*emptySlice), len(initSlice), cap(initSlice))

    //空指针切片初始化值 指定下标的值
    *emptySlice = []int{2:9,3:10}
    initSlice = []int{2:9,3:10}
    //指针追加操作(扩容)
    *emptySlice = append(*emptySlice, 1)
    initSlice = append(initSlice, 1)
    fmt.Println(*emptySlice, &emptySlice, len(*emptySlice), cap(*emptySlice))
    newPoint := emptyArray
    (*newPoint)[1] = 3
    fmt.Println(*emptyArray, *newPoint)
}

func NullSlice()  {
    var ww, cc []string
    //空切片
    ww = nil
    cc = []string{"1", "2", "5"}
    //截取为空切片
    cc = cc[:0]
    fmt.Println(ww, cc, cc[0:2], cap(cc))
}

/**
函数里的形参使用切片 data 使用引用传递
 */
func RemoveIndexData(index int, data []string) []string {
    if index > 0 && index < (len(data) - 1) {
        // append 多于一个值 无需添加... 当追加一个切片 需要追加...
        head := append([]string{}, data[:index]...)
        return append(head, data[(index + 1):]...)
        //return append(data[:index], data[(index + 1):]...)
    } else if index == 0 {
        return data[1:]
    } else if index == (len(data) - 1) {
        return data[:(len(data) - 1)]
    } else {
        panic("index error")
    }
}

/**
copy复制会比等号复制慢。但是copy复制为值复制，改变原切片的值不会影响新切片。
等号复制为指针复制，改变原切片或新切片都会对另一个产生影响。
 */
func CopySlice()  {
    cp := make([]string, 3, 4)
    cc := []string{"1", "2", "5"}
    //切片拷贝 地址完全一致
    mm := cc
    //值拷贝 不拷贝cc的内存地址
    copy(cp, cc)
    //影响cc变量
    mm[0] = "99"
    //与cc内存地址不相同 不影响cc值
    cp[2] = "100"
    fmt.Printf("%p, %p, %p \n", cc, mm, cp)
    // cc 变量地址发生变更
    // cc = append(cc, "909")
    //拷贝进行扩容 与原cc变量地址不相同
    mm = append(mm, "44")
    mm[1] = "77"
    //未超出cap 无需扩容 地址不变
    cp = append(cp, "777")
    // 超出cap 会产生扩容 地址变更
    //cp = append(cp, "777", "009")
    fmt.Println(cc, mm, cp)
    fmt.Printf("%p, %p, %p \n", cc, mm, cp)
}

func ModifiedSlice() {
    //确保append扩容不会溢出
    //a := make([]int, 3, 4)
    //a[0] = 7
    //a[1] = 8
    //a[2] = 9
    a := []int{7, 8, 9}
    txt := "abc"
    fmt.Printf("len: %d cap: %d data: %+v add: %p \n", len(a), cap(a), a, a)
    fmt.Printf("txt add %p \n", &txt)
    // 值传递 地址操作仍受影响 append 操作不影响
    upgrade(a, txt)
    // 引用传递 影响append
    change(&a)
    fmt.Printf("len: %d cap: %d data: %+v add: %p \n", len(a), cap(a), a, a)
    fmt.Println(a, txt)
}

func upgrade(a []int, txt string) {
    // 函数内开辟新地址 存储形参的值(值传递) 没有传递内存地址
    txt = "qq"
    fmt.Printf("upgrading add: %p, txt add %p \n", a, &txt)
    // 形参类型未切片 函数里进行地址传递
    // 对内存地址的操作 会影响外部的数据
    a[0] = 1
    // 不影响函数外的值 即便内存地址相同
    // 在函数内，append操作超过了原始切片的容量，将会有一个新建底层数组的过程
    // 那么此时再修改函数返回切片，应该不会再影响原始切片
    a = append(a, 10)
    fmt.Printf("finished upgrade add: %p \n", a)
    fmt.Println(a)
}

/**
值传递
 */
func change(a *[]int) {
    fmt.Printf("upgrading add: %p \n", a)
    // 形参传递切片 函数里进行引用传递
    // 对内存地址的操作 会影响外部的数据
    (*a)[0] = 1
    //不影响函数外的值 即便内存地址相同
    // 在函数内，append操作超过了原始切片的容量，将会有一个新建底层数组的过程
    // 那么此时再修改函数返回切片，应该不会再影响原始切片
    *a = append(*a, 10)
    fmt.Printf("finished upgrade add: %p \n", a)
    fmt.Println(a)
}
