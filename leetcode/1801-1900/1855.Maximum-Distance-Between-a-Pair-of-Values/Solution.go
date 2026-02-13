package Solution

func Solution(nums1 []int, nums2 []int) int {
	var ret int
	l := len(nums2)
	index := 0
	for idx := range nums1 {
		for ; index < l && nums2[index] >= nums1[idx]; index++ {
		}
		ret = max(ret, index-1-idx)
	}
	if index == 0 {
		return 0
	}
	return ret
}
