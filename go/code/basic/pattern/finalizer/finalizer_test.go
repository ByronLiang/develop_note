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

func Test_RoutineFinalizer(t *testing.T) {
	wrapper := NewFinalizeObjWrapper()
	for i := 0; i < 3; i++ {
		wrapper.finalizeObj.rec <- fmt.Sprintf("data: abc-%v", i)
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println("routine num: ", runtime.NumGoroutine())
	// 不能确保每个变量都能触发回收
	runtime.GC()
	time.Sleep(1 * time.Second)
	// 观察被回收的变量里的routine是否被关闭
	fmt.Println("routine num: ", runtime.NumGoroutine())
	// 变量对象仍有使用，不对变量进行垃圾回收
	//wrapper.finalizeObj.rec <- "aa"
}

func TestFileRead(t *testing.T) {
	FileRead()
}
