package Solution

import "math"

func Solution(nums []int) int {
	start, end, min, max := 0, -1, math.MaxInt32, math.MinInt32
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		if nums[i] >= max {
			max = nums[i]
		} else {
			end = i
		}
		if nums[j] <= min {
			min = nums[j]
		} else {
			start = j
		}
	}
	return end - start + 1
}
