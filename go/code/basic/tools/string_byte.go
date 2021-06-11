package tools

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

/**
直接对内存地址进行拷贝 转换数据类型
 */
func StringToBytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sliceHeader := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}

func BytesToString(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	stringHeader := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&stringHeader))
}

// 移除精度以外为0的数值
func FormatFloatWithPre(num float64, pre int) string {
	const (
		zero = "0"
		dot  = "."
	)
	str := fmt.Sprintf("%." + strconv.Itoa(pre) +"f", num)
	return strings.TrimRight(str, zero)
	//return strings.TrimRight(strings.TrimRight(str, zero), dot)
}
