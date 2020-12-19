package alog

import "fmt"

func StringToInt(s string) {
    amount := 0
    bytes := []byte(s)
    for _, b := range bytes {
        num := int(b - '0')
        if num >= 0 && num <= 9 {
            amount = amount * 10 + num
        }
    }
    fmt.Println(amount)
}

/**
https://leetcode-cn.com/problems/find-the-difference/
 */
func CompareString(s string, t string) byte {
    var res, res2 int
    byteForS, byteForT := []byte(s), []byte(t)
    for _, by := range byteForS {
        res += int(by)
    }
    for _, by := range byteForT {
        res2 += int(by)
    }
    return byte(res2 - res)
}
