package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"test/worker"
	"time"
)

func main() {
	/*
	上述情况表明
	 */
	var kk3, lop *int
	var kk1, kk2 int
	kk1 = 1
	kk2 = 2
	kk3 = &kk1
	kk1 = 9
	const lol string = "400"
	ll := len(lol)
	// ASCII码 -> 转换字符显示
	ll1 := string(lol[0])

	// 指针赋值操作
	//lop = &kk1
	lop = new(int)
	*lop = 100
	const(
		Male = 1
		Female = 2
		Middle = iota + 1
		Kol
	)
	//lol = 400
	fmt.Printf("xx %p, val: %d\n", kk3, *kk3)
	fmt.Println("Hello, 世界", kk1, kk2, kk3, lol, Kol, ll, ll1, *lop)
	fmt.Println("ss")
	x := time.Now().Unix()
	rand.Seed(time.Now().Unix())
	fmt.Println("time is", x)
	fmt.Println("My favorite number is ", rand.Intn(100))
	y := worker.Add(5.0, 9.0) * 5
	fmt.Println(y)
	//var name int = 15
	//var kk float64 = 2.1
	kk := 2.1
	var str = "ss"
	//var tea string
	tea := strconv.FormatFloat(kk, 'E', -1, 32)
	k2, error := strconv.Atoi("20")
	if error != nil {
		println(error)
	}
	sen, extra := worker.PushData("abc")
	fmt.Println(tea, "ss", "xx")
	fmt.Printf("ass is %s, %d, %8.10f", str+sen+extra, k2, kk)
	if "aa" == "aa" {
		fmt.Printf("xx")
	} else {
		fmt.Printf("qq")
	}
}