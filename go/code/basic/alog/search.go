package alog

import (
	"fmt"
	"math"
	"sort"
)

func SearchInsert(data []int, target int) int {
	var (
		l = 0
		r = len(data) - 1
	)
	for l <= r {
		mid := l + (r-l)>>1
		if data[mid] == target {
			return mid
		}
		if data[mid] < target {
			l = mid + 1
		}
		if data[mid] > target {
			r = mid - 1
		}
	}
	return l
}

func SearchRange(data []int, target int) (res []int) {
	var (
		match = -1
		l     = 0
		r     = len(data) - 1
	)
	for l <= r {
		mid := l + (r-l)>>1
		if data[mid] == target {
			match = mid
			break
		}
		if data[mid] < target {
			l = mid + 1
		}
		if data[mid] > target {
			r = mid - 1
		}
	}
	if match > -1 {
		lRes := leftSearch(data[0:match], target)
		rRes := rightSearch(data[(match+1):], target)
		if lRes > -1 {
			res = append(res, lRes)
		} else {
			res = append(res, match)
		}
		if rRes > -1 {
			res = append(res, match+1+rRes)
		} else {
			res = append(res, match)
		}
	} else {
		res = []int{-1, -1}
	}
	return
}

func leftSearch(data []int, target int) int {
	var (
		l = 0
		r = len(data) - 1
	)
	for l <= r {
		mid := l + (r-l)/2
		if data[mid] == target {
			r = mid - 1
		}
		if data[mid] < target {
			l = mid + 1
		}
		if data[mid] > target {
			r = mid - 1
		}
	}
	// 检查出界情况
	if l >= len(data) || data[l] != target {
		return -1
	}
	return l
}

func rightSearch(data []int, target int) int {
	var (
		l = 0
		r = len(data) - 1
	)
	for l <= r {
		mid := l + (r-l)/2
		if data[mid] == target {
			l = mid + 1
		}
		if data[mid] < target {
			l = mid + 1
		}
		if data[mid] > target {
			r = mid - 1
		}
	}
	// 检查出界情况
	if r < 0 || data[r] != target {
		return -1
	}
	return r
}

/**
https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
*/
func MinArray(data []int) int {
	l, r := 0, len(data)-1
	for l < r {
		mid := l + (r-l)>>1
		if data[mid] > data[r] {
			l = mid + 1
		} else {
			r = r - 1
		}
		fmt.Println(l, r)
	}
	return data[l]
}

// 二进制的二分法
// https://leetcode-cn.com/problems/divide-two-integers/
func BitSearch(a, b int32) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	// enhance: 识别正负值
	var sign bool
	if a > 0 && b > 0 || a < 0 && b < 0 {
		sign = true
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	res := 0
	// a是被除数; b是除数
	for a >= b {
		temp := b
		bit := 1
		// 检测下一次二分有没超出边界
		for a >= temp<<1 {
			temp <<= 1
			bit <<= 1
		}
		a = a - temp
		res += bit
	}
	if !sign {
		res = -res
	}
	return res
}

func backTrackBit(a, b int32, res int) int {
	// 终止递归条件
	if a < b {
		return res
	}
	temp := b
	bit := 1
	for a >= temp<<1 {
		temp <<= 1
		bit <<= 1
	}
	a = a - temp
	res += bit
	res = backTrackBit(a, b, res)
	return res
}

/**
https://leetcode-cn.com/problems/merge-intervals/
*/
func MergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	newIntervals := make([][]int, 0)
	newIntervals = append(newIntervals, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lastInterval := newIntervals[len(newIntervals)-1]
		// 最大值无法达到当前区间的最小值
		if lastInterval[1] < intervals[i][0] {
			newIntervals = append(newIntervals, intervals[i])
			continue
		}
		// 最大值到达当前区间值, 进行区间合并
		if lastInterval[1] < intervals[i][1] {
			newIntervals[len(newIntervals)-1][1] = intervals[i][1]
		}
	}
	return newIntervals
}
