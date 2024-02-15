package Solution

import "sort"

func Solution(nums []int) int64 {

	sort.Ints(nums)
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}

	ans := int64(-1)
	for end := len(nums) - 1; end >= 2; end-- {
		idx := sort.Search(end, func(i int) bool {
			return sum[i] > nums[end]
		})
		if idx == end {
			continue
		}
		if r := int64(nums[end] + sum[end-1]); r > ans {
			ans = r
		}
	}
	return ans
}
