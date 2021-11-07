package alog

import "testing"

func TestReverseLeftWords(t *testing.T) {
	res := ReverseLeftWords("abcdefg", 2)
	t.Log(res)
}

func TestLongestCommonPrefix(t *testing.T) {
	res := LongestCommonPrefix([]string{"flower", "flow", "flight"})
	fail := LongestCommonPrefix([]string{"dog", "flower", "car"})
	t.Logf("first: %v; second: %v", res, fail)
}

func TestIsMonotonic(t *testing.T) {
	if IsMonotonic([]int{1, 1, 2}) {
		t.Log("true")
	} else {
		t.Log("false")
	}
}

func TestRemoveDuplicates(t *testing.T) {
	t.Log(RemoveDuplicates("abbaca"))
}

func TestRomanToInt(t *testing.T) {
	t.Log(RomanToInt("MCMXCIV"))
}
