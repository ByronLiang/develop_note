package tools

import (
    "fmt"
    "strconv"
)

/**
https://leetcode-cn.com/problems/excel-sheet-column-title/
 */

type CodeGen struct {
    InitCode    int
    // 数字后缀长度转换
    NumSize     int
    // 字母前缀
    PrefixSize  int
    // 数字后缀字符长度
    NumLength   int
}

func (cg *CodeGen) ReflectCode() string {
    n := cg.InitCode / cg.NumSize
    prefix := cg.genPrefix(n)

    return fmt.Sprintf("%s%s", prefix, strconv.Itoa(cg.InitCode % cg.NumSize))
}

func (cg *CodeGen) genPrefix(code int) string {
    prefixByte := make([]rune, cg.PrefixSize)
    index := cg.PrefixSize - 1
    for code > 0 {
        m := code % 26
        if m == 0 {
            m = 26
            code -= 1
        }
        code = code / 26
        prefixByte[index] = rune(m+64)
        index --
    }
    return string(prefixByte)
}

func (cg *CodeGen) GenTotalCode(total int) []string {
    var codeIndex = cg.InitCode
    cg.InitCode += total
    data := make([]string, 0, total)
    initPrefix := codeIndex / cg.NumSize
    initNum := codeIndex % cg.NumSize
    prefix := cg.genPrefix(initPrefix)
    for i := 0; i < total; i++ {
        code := fmt.Sprintf("%s%0*d", prefix, cg.NumLength, initNum)
        data = append(data, code)
        if initNum == 999 {
            initPrefix += 1
            initNum = 0
            prefix = cg.genPrefix(initPrefix)
        } else {
            initNum ++
        }
    }
    return data
}

func (cg *CodeGen) GenCodeNum(code string) int {
    ret := 0
    runes := []rune(code)
    for _, c := range runes {
        ret = 26 * ret + (int(c-'A') + 1)
    }
    return ret
}
