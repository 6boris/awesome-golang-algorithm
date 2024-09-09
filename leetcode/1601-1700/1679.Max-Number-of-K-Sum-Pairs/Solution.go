package Solution

import "sort"

func Solution(nums []int, k int) int {
	sort.Ints(nums)
	s, e := 0, len(nums)-1
	ans := 0
	for s < e {
		sum := nums[s] + nums[e]
		if sum == k {
			ans++
			s, e = s+1, e-1
			continue
		}
		if sum < k {
			s++
		} else {
			e--
		}
	}
	return ans
}
