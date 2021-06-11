package tools

import (
	"testing"
)

func TestNewRingBuffer(t *testing.T) {
	ringBuf, err := NewRingBuffer(5)
	if err != nil {
		t.Fatal(err)
	}
	ringBuf.Write(10)
	ringBuf.Write(20)
	val, err := ringBuf.Peak()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(printData(val, t))
	ringBuf.Read()
	ringBuf.Write(30)
	ringBuf.Write(40)
	//ringBuf.Write(50)
	//ringBuf.Write(60)
	allData, err := ringBuf.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	list := formatSlice(allData, t)
	t.Log(list)
	ringBuf.Write(50)
	ringBuf.Write(60)
	res, err := ringBuf.Read()
	t.Log(printData(res, t))
}

func printData(i interface{}, t *testing.T) int {
	if data, ok := i.(int); ok {
		return data
	} else {
		t.Error("type error")
	}
	return 0
}

func formatSlice(arr []interface{}, t *testing.T) []int {
	data := make([]int, len(arr))
	for index := range arr {
		data[index] = printData(arr[index], t)
	}
	return data
}
