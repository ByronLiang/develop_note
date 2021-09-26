package worker

import (
	"fmt"
)

/**
当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
*/
func DeferSample() {
	for i := 0; i < 3; i++ {
		defer fmt.Printf("%d ", i)
	}
}

// 非命名值返回
func DeferHandle(i *int) int {
	defer func() {
		*i++
	}()
	return *i
}

// 命名值返回
func DeferHandleWithName(i *int) (m int) {
	m = *i
	defer func() {
		// 影响最终返回的m值
		m++
	}()
	return m + 10
}
