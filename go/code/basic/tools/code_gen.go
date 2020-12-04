package tools

import (
    "fmt"
    "strconv"
)

/**
https://leetcode-cn.com/problems/excel-sheet-column-title/
 */

type CodeGen struct {
    InitCode    string
    // 数字后缀长度
    NumSize     int
    // 字母前缀
    PrefixSize  int
}

func (cg *CodeGen) ReflectCode() string {
    prefixByte := make([]rune, cg.PrefixSize)
    byt := []rune(cg.InitCode)
    WordCode := byt[:len(byt) - cg.NumSize]
    n, _ := strconv.Atoi(string(WordCode))
    index := cg.PrefixSize - 1
    for n > 0 {
        m := n % 26
        if m == 0 {
            m = 26
            n -= 1
        }
        n = n / 26
        prefixByte[index] = rune(m+64)
        index --
    }

    return fmt.Sprintf("%s%s", string(prefixByte), string(byt[(len(byt) - cg.NumSize):]))
}

func (cg *CodeGen) GenCodeNum(code string) int {
    ret := 0
    runes := []rune(code)
    for _, c := range runes {
        ret = 26 * ret + (int(c-'A') + 1)
    }
    return ret
}
