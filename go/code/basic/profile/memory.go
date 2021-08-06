package profile

import (
	"math/rand"
	"time"
)

func GenerateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func LastNumsBySlice(origin []int) []int {
	// 无法释放原地址内存
	return origin[len(origin)-2:]
}

func LastNumsByCopy(origin []int) []int {
	result := make([]int, 2)
	// 独立拷贝数据, 避免引用数据，能对剩余数组的数据进行回收
	copy(result, origin[len(origin)-2:])
	return result
}
