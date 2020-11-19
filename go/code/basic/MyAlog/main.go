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
    //numbersThanCurrent()
    uniqueOccurrences()
    //sumNumbers()
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
    //node := alog.InitListNode([]int{2,2,1})
    node := alog.InitListNode([]int{2,3,4})
    alog.ShowListNode(node)
    //alog.IsPalindrome(node)
    newNode := alog.SwapPairs(node)
    alog.ShowListNode(newNode)
}

// 计数排序算法: 得出当前排名数值
func numbersThanCurrent() {
    fmt.Println(
        alog.SmallerNumbersThanCurrent([]int{4, 4, 2, 4}))
}

func uniqueOccurrences() {
    //alog.UniqueOccurrences([]int{1,2})
    if alog.ValidMountainArray([]int{0,1,2,3,4,5,6,7,8,9}) {
        fmt.Println("match")
    } else {
        fmt.Println("none")
    }
}
/**
https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/submissions/
 */
func sumNumbers() int {
    root := alog.InitTreeData()
    amount := alog.SumNumbers(root)
    fmt.Println(amount)
    return amount
}
