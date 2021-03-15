package alog

import "fmt"

func Permute(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	var tmp []int
	var visited = make([]bool, len(nums))
	backtracking(nums, &res, tmp, visited)
	return res
}

/**
模板 choose -- explore -- unchoose：

用 for 循环枚举出当前的选择
作出一个选择，基于这个选择，继续递归
递归结束了，撤销这个选择，进入下一轮迭代
*/
func backtracking(nums []int, res *[][]int, tmp []int, visited []bool) {
	// 成功找到一组
	if len(tmp) == len(nums) {
		var c = make([]int, len(tmp))
		copy(c, tmp)
		*res = append(*res, c)
		return
	}
	// 回溯
	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			// 做选择 进入递归-展开其他选择
			visited[i] = true
			tmp = append(tmp, nums[i])
			backtracking(nums, res, tmp, visited)
			fmt.Println("loop", i, tmp, visited)
			// 递归终结-剪枝处理
			tmp = tmp[:len(tmp)-1]
			visited[i] = false
		}
	}
	fmt.Println("end", visited, tmp)
}
