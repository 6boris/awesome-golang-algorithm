package Solution

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	ans := nums[0] + nums[1] + nums[len(nums)-1]
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		start, end := i+1, len(nums)-1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum > target {
				end--
			} else {
				start++
			}

			if Abs(sum-target) < Abs(ans-target) {
				ans = sum
			}

		}
	}
	return ans
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
