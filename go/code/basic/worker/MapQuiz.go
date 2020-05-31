package worker

import "fmt"

func MyMapQuiz()  {
    var (
        //空映射 nil映射
        kk map[string]string
        bi = make(map[string]string)
    )

    //assignment to entry in nil map 空映射不能直接赋值
    //kk["gender"] = "male"
    kk = map[string]string{"name":"byron"}

    bi["name"] = "biu"
    cp := bi
    cp["gender"] = "male"
    cp["name"] = "dude"
    fmt.Println(kk["name"], bi, cp, len(bi))
}

func MapPoint()  {
    mf := map[int]func() int{
        1: func() int { return 10 },
        2: func() int { return 20 },
        5: func() int { return 50 },
    }
    kk := mf[1]()
    //一对多 [初始化多个整形切片]
    mp := make(map[string][]int)
    mp["apple"] = []int{3,1}
    mp["apple"] = append(mp["apple"], 9,20)
    mp["banana"] = make([]int, 5)
    fmt.Println(kk, mp["apple"], mp["banana"])
}

func InitMap()  {
    type oo []int
    type pp map[string]string
    //map 不应使用new进行初始化
    //new(pp)
    cc := new(oo)
    *cc = append(*cc, 2)
    //map 使用make 初始化
    ll := make(pp, 1)
    ll["ss"] = "qq"
    ll["qq"] = "xx"
    fmt.Println(*cc, ll)
}
