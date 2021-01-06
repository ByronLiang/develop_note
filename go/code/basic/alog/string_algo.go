package alog

/**
https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/
 */
func ReverseWords(s string) string {
    var l, r = 0, 0
    bytes := []byte(s)
    length := len(bytes)
    for index, bt := range bytes {
        if bt == ' ' {
            r = index - 1
            revers(l, r, bytes)
            l = index + 1
        }
        if index == length - 1 && bt != ' ' {
            r = index
            revers(l, r, bytes)
        }
    }
    return string(bytes)
}

func revers(l, r int, bytes []byte)  {
    for l < r {
        bytes[l], bytes[r] = bytes[r], bytes[l]
        l++
        r--
    }
}
