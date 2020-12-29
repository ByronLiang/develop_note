package alog

import "fmt"

func SearchInsert(data []int, target int) int {
    var (
        l = 0
        r = len(data) - 1
    )
    for l <= r {
        mid := l + (r - l) >> 1
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
        l = 0
        r = len(data) - 1
    )
    for l <= r {
        mid := l + (r - l) >> 1
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
            res = append(res, match + 1 + rRes)
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
        mid := l + (r - l) / 2
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
        mid := l + (r - l) / 2
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
    l, r := 0, len(data) - 1
    for l < r {
        mid := l + (r - l) >> 1
        if data[mid] > data[r] {
            l = mid + 1
        } else {
            r = r - 1
        }
        fmt.Println(l, r)
    }
    return data[l]
}
