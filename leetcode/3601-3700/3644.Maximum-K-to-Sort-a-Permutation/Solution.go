package Solution

import "math"

func Solution(nums []int) int {
	n := len(nums)
	res := math.MaxInt32
	for i := 0; i < n; i++ {
		if nums[i] == i {
			continue
		}
		res &= nums[i]
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}
