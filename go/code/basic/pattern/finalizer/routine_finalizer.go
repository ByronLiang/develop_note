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
	wrapper := &finalizeObjWrapper{obj}
	// 只要协程一直在运行，Gc就无法选中对象执行垃圾回收。
	//实际上，给ttlmap设置了SetFinalizer后，如果ttlmap.clear()协程没有return，而是阻塞了，Gc就可以将ttlmap回收掉。
	runtime.SetFinalizer(wrapper, wrapperFinalizerHandle)
	return wrapper
}

func (obj *finalizeObj) handle()  {
	for {
		select {
		case <-obj.close:
			fmt.Println("close finalize obj")
			return
		case data := <- obj.rec:
			fmt.Println(data)
		}
	}
}

func wrapperFinalizerHandle(wrapper *finalizeObjWrapper) {
	wrapper.finalizeObj.close <- struct{}{}
}
