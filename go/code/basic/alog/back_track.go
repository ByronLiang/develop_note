package alog

import (
	"fmt"
)

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
	// 确定结束条件
	if len(tmp) == len(nums) {
		var c = make([]int, len(tmp))
		copy(c, tmp)
		*res = append(*res, c)
		return
	}
	// 回溯
	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			fmt.Println("before loop", i, tmp, visited)
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

func PartitionTxt(s string) (res [][]string) {
	dfsTxt([]string{}, 0, &res, s) // 当前的部分解为空数组，从索引0开始，往后切回文串
	return
}

// dfs：基于当前的部分解temp，去切从start到末尾的子串
func dfsTxt(temp []string, start int, res *[][]string, s string) {
	if start == len(s) { // 当start指针越界了，一直切出回文才走到这，将当前的部分解temp加入解集res
		t := make([]string, len(temp)) // 新建一个和temp等长的切片
		copy(t, temp)                  // temp还要在递归中继续被修改，不能将它的引用推入res
		*res = append(*res, t)         // 将temp的拷贝 加入解集res
		return                         // 结束掉当前这个递归（分支）
	}

	for i := start; i < len(s); i++ { // 枚举出当前的所有选项，从索引start到末尾索引
		fmt.Println("before", start, i, temp)
		temp = append(temp, s[start:i+1]) // 切出来，加入到部分解temp
		fmt.Println("after", temp)
		dfsTxt(temp, i+1, res, s) // 基于这个选择，继续往下递归，继续切
		temp = temp[:len(temp)-1] // 上面递归结束了，撤销当前选择i，去下一轮迭代
		//if isPali(s, start, i) { // 当前选择i，如果 start到 i 是回文串，就切它
		//
		//}
	}
}

func trackingAll(i int, target []int, tmp []int, res *[][]int) {
	container := make([]int, len(tmp))
	copy(container, tmp)
	*res = append(*res, container)
	fmt.Println("before", tmp, i)
	for j := i; j < len(target); j++ {
		// 处理重复子集合
		if j > i && target[j] == target[j-1] {
			fmt.Println("check", j)
			continue
		}
		tmp = append(tmp, target[j])
		trackingAll(j+1, target, tmp, res)
		fmt.Println("end", tmp, j)
		tmp = tmp[:len(tmp)-1]
	}
}

/**
子集合
*/
func SubSort(t []int) [][]int {
	var res [][]int
	trackingAll(0, t, []int{}, &res)
	return res
}

var (
	letterMap = []string{
		" ",    //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	res = make([]string, 0)
)

// DFS
func LetterCombinations(digits string) []string {
	if digits == "" {
		return res
	}
	findCombination(&digits, 0, "")
	return res
}

func findCombination(digits *string, index int, s string) {
	if index == len(*digits) {
		res = append(res, s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num-'0']
	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, s+string(letter[i]))
		// 由于只传形参, 无需对s进行剪枝处理
		fmt.Println(letter[i], s)
	}
	return
}
