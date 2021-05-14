package finalizer

import (
	"fmt"
	"runtime"
	"syscall"
	"time"
)

type File struct { d int }

func FileRead() {
	p := openFile("t.txt")
	// 只传递fd副本, p被回收, fd句柄关闭 触发error: bad file describe
	content := readFile(p.d)
	// 不引用变量下，延长变量生命周期, 避免垃圾回收
	//runtime.KeepAlive(p)
	// 有直接引用File结构体
	//content := readContent(p)
	println("Here is the content: "+content)
	time.Sleep(100 * time.Millisecond)
}

func openFile(path string) *File {
	d, err := syscall.Open(path, syscall.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}

	p := &File{d}
	runtime.SetFinalizer(p, func(p *File) {
		fmt.Println("File point gc")
		// 垃圾回收-触发关闭文件句柄(file describe)
		syscall.Close(p.d)
	})

	return p
}

func readFile(descriptor int) string {
	fullAllocation()
	runtime.GC()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("read: ", descriptor)
	var buf [1000]byte
	_, err := syscall.Read(descriptor, buf[:])
	if err != nil {
		panic(err)
	}
	return string(buf[:])
}

func fullAllocation() {
	var a *int
	// memory increase to force the GC
	for i:= 0; i < 10000000; i++ {
		i := 1
		a = &i
	}
	_ = a
}
