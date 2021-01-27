package finalizer

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func buildFinalizerSample() *int {
	rand.Seed(time.Now().UnixNano())
	var i = rand.Intn(10)
	runtime.SetFinalizer(&i, finalizersHandle)

	return &i
}

func finalizersHandle(num *int)  {
	fmt.Println("gc for: ", *num)
}
