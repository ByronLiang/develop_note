package alog

import (
    "fmt"
    "log"
)

func BuildStatusTable(s string) []int {
    next := make([]int, len(s))
    next[0] = -1
    fmt.Println(next)
    i, j := 0, -1
    for ; i< len(s)-1;  {
        log.Print(i, j)
        if j == -1 || s[i] == s[j] {
            i++
            j++
            log.Print("update", i, j)
            next[i] = j
        } else {
            j = next[j]
        }
    }
    return next
}
