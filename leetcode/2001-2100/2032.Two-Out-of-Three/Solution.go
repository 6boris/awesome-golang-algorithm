package Solution

func Solution(nums1 []int, nums2 []int, nums3 []int) []int {
	ans := make([]int, 0)
	bucket := make([]int, 101)
	for _, n := range nums1 {
		bucket[n] |= 1
	}
	for _, n := range nums2 {
		bucket[n] |= 2
	}
	for _, n := range nums3 {
		bucket[n] |= 4
	}

	// 000, 001, 010, 011, 100, 101, 110,111
	for idx, n := range bucket {
		if n == 3 || n == 5 || n == 6 || n == 7 {
			ans = append(ans, idx)
		}
	}
	return ans
}
