package alog

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
