package Solution

import "sort"

func Solution(nums []int) int64 {
	var ret, tmp int64
	// 1, 2, ,3 4, 5
	got := make([]int64, len(nums))
	for i := range nums {
		tmp = int64(nums[i])
		got[i] = tmp * tmp
	}
	sort.Slice(got, func(i, j int) bool {
		return got[i] < got[j]
	})
	n := len(got)
	for i := n - 1; i >= n/2; i-- {
		ret += got[i]
	}
	for i := n/2 - 1; i >= 0; i-- {
		ret -= got[i]
	}
	return ret
}
