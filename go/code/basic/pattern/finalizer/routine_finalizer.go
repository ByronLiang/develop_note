package finalizer

import (
	"fmt"
	"runtime"
)

type finalizeObj struct {
	rec		chan string
	close 	chan struct{}
}

type finalizeObjWrapper struct {
	*finalizeObj
}

func NewFinalizeObjWrapper() *finalizeObjWrapper {
	obj := &finalizeObj{
		rec:   make(chan string),
		close: make(chan struct{}),
	}
	// 开启协程异步处理数据
	go obj.handle()
	wrapper := &finalizeObjWrapper{obj}
	// 只要协程一直在运行，Gc就无法选中对象执行垃圾回收。
	// 实际上，给finalizeObj设置了SetFinalizer后，如果finalizeObj.handle()协程没有return，而是阻塞了，Gc就可以将finalizeObj回收掉。
	runtime.SetFinalizer(wrapper, wrapperFinalizerHandle)
	return wrapper
}

func (obj *finalizeObj) handle() {
	for {
		select {
		case <-obj.close:
			fmt.Println("close finalize obj")
			return
		case data := <- obj.rec:
			fmt.Println("rec data: ", data)
		}
	}
}

func wrapperFinalizerHandle(wrapper *finalizeObjWrapper) {
	fmt.Println("finalizer working")
	// 变量进行垃圾回收触发关闭routine 避免变量被回收，里面的协程未关闭出现内存泄露
	wrapper.finalizeObj.close <- struct{}{}
}
