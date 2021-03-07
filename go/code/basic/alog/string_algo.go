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

/**
https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/submissions/
 */
func ReverseLeftWords(s string, n int) string {
    bytes := []byte(s)
    length := len(bytes)
    newBytes := make([]byte, 0, length)
    for i := n; i < length + n; i ++ {
        newBytes = append(newBytes, bytes[i%length])
    }
    return string(newBytes)
}

/**
https://leetcode-cn.com/problems/longest-common-prefix/
 */
func LongestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    var loopLen, j, targetLen, compareTarLen int
    initTarget := strs[0]
    for i:= 1; i < len(strs); i++ {
        targetLen = len(initTarget)
        compareTar := strs[i]
        compareTarLen = len(compareTar)
        if targetLen > compareTarLen {
            loopLen = compareTarLen
        } else {
            loopLen = targetLen
        }

        for j = 0; j < loopLen; j++ {
            if initTarget[j] != compareTar[j] {
                break
            }
        }
        initTarget = initTarget[:j]
    }
    return initTarget
}

/**
https://leetcode-cn.com/problems/monotonic-array/
 */
func IsMonotonic(A []int) bool {
    inc, dec := true, true
    for i := 0; i < len(A)-1; i++ {
        if A[i] > A[i+1] {
            inc = false
        }
        if A[i] < A[i+1] {
            dec = false
        }
    }
    return inc || dec
}
