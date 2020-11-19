package alog

import "fmt"

func SmallerNumbersThanCurrent(nums []int) []int {
    var (
        max = 0
        // 初始化二维数组
        data = make([][10]int, 2)
        )
    for _, num := range nums {
        if max < num {
            max = num
        }
        // 存储出现次数
        data[0][num]++
    }

    for i := 1; i <= max; i++ {
        // 计算比当前数值小的出现次数
        data[1][i] = data[1][i-1] + data[0][i-1]
    }

    for index, num := range nums {
        nums[index] = data[1][num]
    }
    return nums
}

/**
https://leetcode-cn.com/problems/unique-number-of-occurrences/
 */
func UniqueOccurrences(arr []int) bool {
    var temp, storage = make(map[int]int), make(map[int]struct{})
    for _, num := range arr {
       temp[num]++
    }
    for _, val := range temp {
       if _, ok := storage[val]; ok{
           fmt.Println("none match")
           return false
       } else {
           storage[val] = struct {}{}
       }
    }
    fmt.Println("success")
    return true
}

func ValidMountainArray(A []int) bool {
    length := len(A)
    if length < 3 {
        return false
    }
    i, j := 0, length- 1
    l, r := 0, 0
    for i < j {
        if A[i] < A[i+1] {
            i++
        } else if A[i] == A[i+1] {
            return false
        } else {
            l = i
            if r != 0 {
                return l == r
            }
        }
        if A[j] < A[j-1] {
            j--
        } else if A[j] == A[j-1] {
            return false
        } else {
            r = j
            if l != 0 {
                return r == l
            }
        }
    }
    return i > 0 && j < length -1 && i == j
}
