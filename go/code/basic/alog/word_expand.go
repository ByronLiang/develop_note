package alog

import (
    "fmt"
    "strconv"
    "strings"
)

/**
给定一个经过编码的字符串，返回它解码后的字符串
https://leetcode-cn.com/problems/decode-string/
 */
func ByteExpand(txt string) string {
    var (
        tar = []byte(txt)
        newStr = ""
        buf []byte
        num int
        steak []*pak
        build strings.Builder
    )
    for _, bt := range tar {
        // 识别数值
        if bt >= 49 && bt <= 57 {
            num, _ = strconv.Atoi(string(bt))
            fmt.Println(num)
            continue
        }
        if string(bt) == "[" {
            // 入栈对象 倍数字符串 与 待拼接字符串
            i := &pak{
                Storage: string(buf),
                Size:    num,
            }
            steak = append(steak, i)
            buf = nil
            continue
        }
        if string(bt) == "]" {
            build.Reset()
            current := steak[len(steak) - 1]
            steak = steak[:len(steak) - 1]
            fmt.Println(string(buf), "process target")
            for i := 0; i < current.Size; i ++ {
                build.WriteString(string(buf))
            }
            // 取出倍数处理的字符串
            processedText := build.String()
            // 清空 并取出栈里内容与倍数处理的字符串 拼接
            build.Reset()
            build.WriteString(current.Storage)
            build.WriteString(processedText)
            // 拼接完成 重新放回缓冲区
            newStr = build.String()
            buf = []byte(newStr)
            fmt.Println(newStr)
            continue
        }
        buf = append(buf, bt)
    }
    return string(buf)
}

type pak struct {
    Storage     string
    Size        int
}

func CheckUnique(data string) bool {
    bt := []byte(data)
    for i := 0; i < len(bt); i++ {
        res := strings.LastIndex(data, string(bt[i]))
        //fmt.Println(res, string(bt[i]))
        if res != i {
            return false
        }
    }
    return true
}
