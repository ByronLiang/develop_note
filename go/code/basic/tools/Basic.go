package tools

import (
	"fmt"
	"math/rand"
	"time"
)

/**
随机睡眠定时器
*/
func CasualTimeCount(size int) {
	rand.Seed(time.Now().UnixNano())
	count := rand.Intn(size)
	fmt.Println(count)
	time.Sleep(time.Duration(count) * time.Second)
}
