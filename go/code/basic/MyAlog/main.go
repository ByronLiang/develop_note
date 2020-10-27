package main

import (
    "fmt"
    "test/alog"
)

func main()  {
    //byteExp()
    //kmpTest()
    //search()
    //isUnique()
    //backspace()
    //linkList()
    numbersThanCurrent()
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

func backspace() {
    if alog.BackspaceCompare("#csl#", "#csl#") {
        fmt.Println("match")
    } else {
        fmt.Println("failed")
    }
}

func linkList()  {
    node := alog.InitListNode()
    alog.ShowListNode(node)
    alog.IsPalindrome(node)
}

// 计数排序算法: 得出当前排名数值
func numbersThanCurrent() {
    fmt.Println(
        alog.SmallerNumbersThanCurrent([]int{4, 4, 2, 4}))
}
