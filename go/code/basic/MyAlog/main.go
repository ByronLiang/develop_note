package main

import (
    "fmt"
    "test/alog"
)

func main()  {
    //res := alog.KmpSearch("hello", "ll")
    //res := alog.PreKMP("ll")
    //demo := alog.PreKMP("abcdabca")
    demo := alog.PreKMP("aabaabaaa")
    fmt.Println(demo)
    //res := alog.Kmp("hello", "abcdabca")
    res := alog.KmpSearch("abxabcabcaby", "abcaby")
    fmt.Println(res)

}
