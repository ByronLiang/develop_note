package finalizer

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test_FinalizerSample(t *testing.T) {
	for i := 0; i < 3; i++ {
		res := buildFinalizerSample()
		fmt.Println("num: ", *res)
		if i % 2 != 0 {
			// 奇数生成的对象取消触发回收回调
			runtime.SetFinalizer(res, nil)
		}
	}
	time.Sleep(500 * time.Millisecond)
	// 不能确保每个变量都能触发回收
	runtime.GC()
	time.Sleep(1 * time.Second)
}
