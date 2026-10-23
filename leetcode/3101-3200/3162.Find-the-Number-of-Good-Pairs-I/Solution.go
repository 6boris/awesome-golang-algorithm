package Solution

func Solution(nums1 []int, nums2 []int, k int) int {
	var ret int
	for i := range nums2 {
		nums2[i] *= k
	}
	for i := range nums1 {
		for j := range nums2 {
			if nums1[i]%nums2[j] == 0 {
				ret++
			}
		}
	}
	return ret
}
