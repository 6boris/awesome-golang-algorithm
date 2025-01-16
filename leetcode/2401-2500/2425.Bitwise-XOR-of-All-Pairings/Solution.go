package Solution

func Solution(nums1 []int, nums2 []int) int {
	l1 := len(nums1) & 1
	l2 := len(nums2) & 1
	ans := 0
	if l2 == 1 {
		for _, n := range nums1 {
			ans ^= n
		}
	}
	if l1 == 1 {
		for _, n := range nums2 {
			ans ^= n
		}
	}
	return ans
}
