package alog

import (
    "bytes"
    "fmt"
)

/**
给定 S 和 T 两个字符串 # 代表退格字符。
https://leetcode-cn.com/problems/backspace-string-compare/
 */
func BackspaceCompare(s, t string) bool {
    var sList, tList []byte
    if len(s) == len(t) {
        for i := 0; i < len(s); i++ {
            compared(&sList, s[i])
            compared(&tList, t[i])
        }
    } else {
        for i := 0; i < len(s); i++ {
            compared(&sList, s[i])
        }
        for i := 0; i < len(t); i++ {
            compared(&tList, t[i])
        }
    }
    defer func(s, t []byte) {
        fmt.Println(string(s), string(t))
    }(sList, tList)
    // 判断切片数据
    //return reflect.DeepEqual(sList, tList)
    return bytes.Equal(sList, tList)
}

/**
切片值传递 影响原数据源
 */
func compared(list *[]byte, s byte)  {
    if s != '#' {
        *list = append(*list, s)
    } else {
        length := len(*list)
        if length > 0 {
            // 出栈
            *list = (*list)[:(length - 1)]
        }
    }
}
