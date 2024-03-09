package Solution

import "sort"

func Solution(nums1 []int, nums2 []int) int {
	l2 := len(nums2)
	for _, n := range nums1 {
		if idx := sort.Search(l2, func(i int) bool {
			return nums2[i] >= n
		}); idx != l2 && nums2[idx] == n {
			return n
		}
	}
	return -1
}
