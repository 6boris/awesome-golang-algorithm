package Solution

import "strconv"

func Solution(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		digits := len(strconv.Itoa(nums[i]))
		if digits&1 == 0 {
			count++
		}
	}
	return count
}
