package alog

import (
	"testing"
)

func TestMinArray(t *testing.T) {
	res := MinArray([]int{4, 5, 5, 6, 1, 2, 3, 4})
	if res != 1 {
		t.Error("Not min data")
	}
}

func TestBitSearch(t *testing.T) {
	res := BitSearch(-10, 3)
	t.Log(res)
}
