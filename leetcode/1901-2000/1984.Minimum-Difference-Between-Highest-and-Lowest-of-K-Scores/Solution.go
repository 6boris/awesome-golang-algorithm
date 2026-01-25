package Solution

import "sort"

func Solution(nums []int, k int) int {
	if k == 1 {
		return 0
	}

	ret, end := 100001, len(nums)-k
	sort.Ints(nums)
	for i := 0; i <= end; i++ {
		ret = min(ret, nums[i+k-1]-nums[i])
	}
	return ret
}
