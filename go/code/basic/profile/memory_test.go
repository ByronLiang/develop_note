package profile

import (
	"fmt"
	"os"
	"runtime/pprof"
	"testing"
)

func TestLastNumsByCopy(t *testing.T) {
	ans := make([][]int, 0)
	for k := 0; k < 100; k++ {
		origin := GenerateWithCap(128 * 1024) // 1M
		res := LastNumsByCopy(origin)
		ans = append(ans, res)
	}
	file, err := os.Create("./copy_mem.pprof")
	defer file.Close()
	if err != nil {
		fmt.Printf("create mem pprof failed, err:%v\n", err)
		return
	}
	pprof.WriteHeapProfile(file)
}

func TestLastNumsBySlice(t *testing.T) {
	ans := make([][]int, 0)
	for k := 0; k < 100; k++ {
		origin := GenerateWithCap(128 * 1024) // 1M
		res := LastNumsBySlice(origin)
		ans = append(ans, res)
	}
	file, err := os.Create("./slice_mem.pprof")
	defer file.Close()
	if err != nil {
		fmt.Printf("create mem pprof failed, err:%v\n", err)
		return
	}
	pprof.WriteHeapProfile(file)
}
