package main

import (
    "fmt"
    "test/alog"
)

func main()  {
    //byteExp()
    //kmpTest()
    //search()
    isUnique()
}

func kmpTest()  {
    //res := alog.KmpSearch("hello", "ll")
    //res := alog.PreKMP("ll")
    //demo := alog.PreKMP("abcdabca")
    demo := alog.PreKMP("aabaabaaa")
    fmt.Println(demo)
    //res := alog.Kmp("hello", "abcdabca")
    res := alog.KmpSearch("abxabcabcaby", "abcaby")
    fmt.Println(res)
}

func search()  {
    //fmt.Println(alog.SearchInsert([]int{1,3,5,6}, 7))
    fmt.Println(alog.SearchRange([]int{2,3,5,7,7,9,9,10}, 7))
}

func byteExp() {
    fmt.Println(alog.ByteExpand("abc1[a2[k]x2[mm]]zz"))
}

func isUnique()  {
    if  ! alog.CheckUnique("letepo") {
        fmt.Println("none unique")
    } else {
        fmt.Println("unique")
    }
}
