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
		if index == length-1 && bt != ' ' {
			r = index
			revers(l, r, bytes)
		}
	}
	return string(bytes)
}

func revers(l, r int, bytes []byte) {
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
	for i := n; i < length+n; i++ {
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
	for i := 1; i < len(strs); i++ {
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

/**
https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/
*/
func RemoveDuplicates(S string) string {
	if len(S) < 2 {
		return S
	}
	tmp := make([]byte, 0, len(S))
	for i := 0; i < len(S); i++ {
		if len(tmp) == 0 {
			tmp = append(tmp, S[i])
		} else {
			if tmp[len(tmp)-1] != S[i] {
				tmp = append(tmp, S[i])
			} else {
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	return string(tmp)
}

// https://leetcode-cn.com/problems/roman-to-integer/
func RomanToInt(s string) int {
	set := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	i := 0
	for i < len(s) {
		c := s[i]
		if i == len(s)-1 {
			total += set[c]
			break
		}
		n := s[i+1]
		if (c == 'I') && (n == 'V' || n == 'X') {
			total += set[n] - 1
			i += 2
			continue
		}
		if (c == 'X') && (n == 'L' || n == 'C') {
			total += set[n] - 10
			i += 2
			continue
		}
		if (c == 'C') && (n == 'D' || n == 'M') {
			total += set[n] - 100
			i += 2
			continue
		}
		total += set[c]
		i++
	}
	return total
}
