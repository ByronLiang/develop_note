package alog

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	res := Permute([]int{1, 2, 3})
	fmt.Println(res)
}

func TestPartitionTxt(t *testing.T) {
	res := PartitionTxt("aab")
	fmt.Println(res)
}

func TestSubSort(t *testing.T) {
	res := SubSort([]int{1, 2, 3})
	fmt.Println(res)
}

func TestLetterCombinations(t *testing.T) {
	res := LetterCombinations("23")
	t.Log(res)
}
