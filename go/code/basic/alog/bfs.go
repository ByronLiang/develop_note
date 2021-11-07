package alog

/**
https://leetcode-cn.com/problems/jump-game-iii/
*/
func CanReach(arr []int, start int) bool {
	queue := []int{start}
	for len(queue) > 0 {
		// get ele from queue
		ele := queue[0]
		queue = queue[1:]
		if arr[ele] == 0 {
			return true
		}
		// move front or back. if match, push num into queue
		numbers := []int{ele + arr[ele], ele - arr[ele]}
		for _, num := range numbers {
			if 0 <= num && num < len(arr) && arr[num] >= 0 {
				queue = append(queue, num)
			}
		}
		arr[ele] = -1
	}
	return false
}
