package Solution

import "sort"

func Solution(nums []int) []int {
	ans := make([]int, 0)
	sort.Ints(nums)
	need := len(nums) / 3
	count := 0
	// 1, 1, 2
	for idx := 0; idx < len(nums); idx++ {
		if count == 0 || nums[idx] == nums[idx-1] {
			count++
			continue
		}
		if count > need {
			ans = append(ans, nums[idx-1])
		}
		count = 1
	}
	if count > need {
		ans = append(ans, nums[len(nums)-1])
	}
	return ans
}
