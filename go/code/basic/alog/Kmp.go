package alog

import (
    "fmt"
)

func KmpSearch(text string, target string) int {
    var (
        res = 0
        i = 0
        j = 0
    )
    textLen := len(text)
    targetLen := len(target)
    table := PreMP(target)
    textByte := []byte(text)
    targetByte := []byte(target)
    //当可搜索长度小于目标长度 停止搜索
    for j < textLen {
        // 回退上一次匹配的下标
        for i > -1 && targetByte[i] != textByte[j] {
            // 根据部分匹配表进行下标重定位
            i = table[i]
        }
        // 针对匹配字符进行步进扫描
        i ++
        // 继续向前扫描
        j ++
        // 完成目标字符扫描 得出匹配起始下标值
        if i >= targetLen {
            fmt.Println("match", j, i)
            res = j - i
            break
        }
    }
    return res
}

func PreMP(x string) [10]int {
    var i, j int
    length := len(x) - 1
    var mpNext [10]int
    i = 0
    j = -1
    // 标志起点 以便区分数值0
    mpNext[0] = -1

    for i < length {
        // 回退到指定边界:
        // 1. 回退到上一次匹配的下标位置
        // 2. 从未匹配过下标位置，回到起点下标
        for j > -1 && x[i] != x[j] {
            fmt.Println("before", j, i, mpNext[j])
            j = mpNext[j]
        }
        i++
        j++
        fmt.Println("count plus", i, j)
        mpNext[i] = j
    }
    return mpNext
}

func PreKMP(x string) [10]int {
    var i, j int
    length := len(x) - 1
    var mpNext [10]int
    i = 0
    j = 0
    // 标志起点 以便区分数值0
    mpNext[0] = 0

    for i < length {
        // 回退到指定边界:
        // 1. 回退到上一次匹配的下标位置
        // 2. 从未匹配过下标位置，回到起点下标
        i++
        for j > 0 && x[i] != x[j] {
            fmt.Println("before", j, i, mpNext[j])
            // 根据以前的标志 决定回退的下标
            j = mpNext[(j-1)]
        }
        if x[i] == x[j] {
            j ++
            fmt.Println("count plus", i, j)
            mpNext[i] = j
        }
    }
    return mpNext
}

func Kmp(text string, target string) int {
    var (
        res = 0
        i = 0
        j = 0
    )
    textLen := len(text)
    targetLen := len(target)
    table := PreKMP(target)
    textByte := []byte(text)
    targetByte := []byte(target)
    //当可搜索长度小于目标长度 停止搜索
    for j < textLen {
        // 回退上一次匹配的下标
        for i > 0 && targetByte[i] != textByte[j] {
            fmt.Println("cal", j, i, table[i])
            // 根据部分匹配表进行下标重定位
            i = table[(i-1)]
        }
        if targetByte[i] == textByte[j] {
            // 针对匹配字符进行步进扫描
            i ++
        }
        // 完成目标字符扫描 得出匹配起始下标值
        if i >= targetLen {
            fmt.Println("match", j, i)
            res = j - (i - 1)
            break
        }
        // 继续向前扫描
        j ++
    }
    fmt.Println(j, textLen)
    return res
}
