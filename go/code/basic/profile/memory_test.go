package profile

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"runtime"
	"runtime/pprof"
	"testing"
	"time"
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

// 图片处理性能测试
func TestProfileCompressImage(t *testing.T) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	t.Logf("before: %d MB\n", m.Alloc/1024/1024)
	for i := 1; i < 4; i++ {
		filename := fmt.Sprintf("view%d.%s", i, "jpg")
		fileBt, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Error(err)
			continue
		}
		buf := bytes.NewReader(fileBt)
		_, format, err := image.Decode(buf)
		// 比 Decode 节省4倍内存
		//conf, format, err := image.DecodeConfig(buf)
		if err != nil {
			t.Error(err)
			continue
		}
		t.Logf("name: %s, format: %s", filename, format)
		//t.Logf("w: %d, h: %d", conf.Width, conf.Height)
		time.Sleep(500 * time.Millisecond)
		// 强制回收资源
		runtime.GC()
	}
	runtime.ReadMemStats(&m)
	t.Logf("after: %d MB\n", m.Alloc/1024/1024)
}
