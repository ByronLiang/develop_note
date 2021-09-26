package worker

import "fmt"

var (
	n       = 10
	pp, pp1 *string
	//数组指针声明
	pp2, pp3 *[3]int
)

func Quiz1() {
	//指针初始化方式 可采用 & new
	pp3 = &[3]int{7, 5, 3}
	pp2 = new([3]int)
	//指针数组赋值
	(*pp2)[0] = 2
	fmt.Println((*pp3)[0:2])
	//指针赋值初始化
	pp = new(string)
	*pp = "ss"

	var lo = &n
	trans(5, lo)
	println(*lo, *pp)
}

func trans(val int, point *int) {
	*point = *point * val
}
