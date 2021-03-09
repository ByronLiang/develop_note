package tools

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestPointExchange(t *testing.T) {
	v2 := 13
	p := *(*uint)(unsafe.Pointer(&v2))
	t.Logf("type: %v; val: %v", reflect.TypeOf(p), p)
}

func TestStringToBytes(t *testing.T) {
	bt := StringToBytes("abc")
	t.Logf("conv data: %v", bt)
	str := BytesToString(bt)
	t.Logf("conv data string: %v", str)
}
