package Solution

import "math"

func Solution(nums []int) int {
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	for i := 0; i < n; i++ {
		if nums[i] >= max1 {
			if nums[i] != max1 {
				max3, max2, max1 = max2, max1, nums[i]
			}
		} else if nums[i] >= max2 {
			if nums[i] != max2 {
				max3, max2 = max2, nums[i]
			}
		} else if nums[i] > max3 {
			max3 = nums[i]
		}
	}
	if max3 != math.MinInt64 {
		return max3
	} else {
		return max1
	}
}
