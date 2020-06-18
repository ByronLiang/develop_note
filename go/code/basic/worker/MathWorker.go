package worker

import "fmt"

func Add(x float32, y float32) float32 {
	return x + y
}

func PushData(x string) (string, string) {
	return "ss" + x, "hhh"
}

func RoundNum(x int)  {
	fmt.Println(x)
	x--
	if x > 0 {
		RoundNum(x)
	}
}

func channelRoundNum(start int) chan int {
	bottom := make(chan int)
	go func() {
		for ; start > 0; start -- {
			bottom <- start
		}
		close(bottom)
	}()
	return bottom
}

func PrintRoundNum(x int)  {
	bottom := channelRoundNum(x)
	//遍历通道(chan) 无法获取索引
	for data := range bottom {
		fmt.Println(data)
	}
}
