package Solution

import "sort"

func Solution(nums []int, queries []int) []int {
	ans := make([]int, len(queries))
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	for idx, q := range queries {
		sum := 0
		for ; sum < len(nums); sum++ {
			if nums[sum] > q {
				break
			}
		}
		ans[idx] = sum
	}
	return ans
}
