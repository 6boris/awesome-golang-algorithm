package Solution

import "sort"

func Solution(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)

	ans := n
	right := 0

	for left := 0; left < n; left++ {
		for right < n && int64(nums[right]) <= int64(nums[left])*int64(k) {
			right++
		}
		current := n - (right - left)
		if current < ans {
			ans = current
		}
	}

	return ans
}
