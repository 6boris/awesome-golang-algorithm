package Solution

import "sort"

func Solution(nums1 []int, nums2 []int) []int {
	indies := make([]int, len(nums2))
	res := make([]int, len(nums2))
	for i := range nums2 {
		indies[i] = i
		res[i] = -1
	}
	sort.Slice(indies, func(i, j int) bool {
		a, b := indies[i], indies[j]
		return nums2[a] < nums2[b]
	})
	sort.Ints(nums1)
	i1, i2 := len(nums1)-1, len(nums2)-1
	for i1 >= 0 && i2 >= 0 {
		if nums1[i1] > nums2[indies[i2]] {
			res[indies[i2]] = nums1[i1]
			nums1[i1] = -1
			i1, i2 = i1-1, i2-1
			continue
		}
		i2--
	}
	index := 0
	for i := range nums1 {
		if nums1[i] == -1 {
			continue
		}
		for ; index < len(nums2) && res[index] != -1; index++ {
		}
		res[index] = nums1[i]
	}

	return res
}
