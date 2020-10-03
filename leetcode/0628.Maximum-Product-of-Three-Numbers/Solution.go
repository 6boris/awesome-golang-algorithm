package Solution

import (
	"sort"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 暴力排序解法
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return max(nums[n-1]*nums[n-2]*nums[n-3], nums[n-1]*nums[0]*nums[1])
}

// 桶排序排序 O(N)
func maximumProduct2(nums []int) int {
	max1, max2, max3, min1, min2 := -1001, -1001, -1001, 1001, 1001

	for _, v := range nums {
		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}

		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 {
			min2 = v
		}

	}
	return max(max1*min1*min2, max1*max2*max3)
}
