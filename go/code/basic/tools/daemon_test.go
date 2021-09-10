package tools

import (
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

// 需要将程序进行包编译，执行./包名 -d 程序将在后台运行
func TestDaemon(t *testing.T) {
	Daemon(handle)
}

func handle(c chan struct{}) {
	file, err := os.OpenFile("dae_test.txt", os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	for {
		select {
		case <-c:
			return
		default:
			file.Write([]byte(strconv.Itoa((int)(time.Now().Unix())) + "\n"))
			time.Sleep(time.Second * 1)
		}
	}
}
