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
    // 数字后缀长度
    NumSize     int
    // 字母前缀
    PrefixSize  int
}

func (cg *CodeGen) ReflectCode() string {
    n := cg.InitCode / cg.NumSize
    prefix := genPrefix(n, cg.PrefixSize)

    return fmt.Sprintf("%s%s", prefix, strconv.Itoa(cg.InitCode % cg.NumSize))
}

func genPrefix(code, length int) string {
    prefixByte := make([]rune, length)
    index := length - 1
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
    data := make([]string, 0, total)
    lastCode := cg.InitCode + total
    initPrefix := cg.InitCode / cg.NumSize
    initNum := cg.InitCode % cg.NumSize
    prefix := genPrefix(initPrefix, cg.PrefixSize)
    cg.InitCode = lastCode
    for i := 0; i < total; i++ {
        //num, _ := fmt.Printf("%0*d", 3, initNum)
        code := fmt.Sprintf("%s%d", prefix, initNum)
        data = append(data, code)
        if initNum == 999 {
            initPrefix += 1
            initNum = 0
            prefix = genPrefix(initPrefix, cg.PrefixSize)
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
