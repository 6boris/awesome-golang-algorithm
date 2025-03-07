package Solution

func Solution(nums1 [][]int, nums2 [][]int) [][]int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	res := [][]int{}
	if nums1[0][0] == nums2[0][0] {
		res = append(res, []int{nums1[0][0], nums1[0][1] + nums2[0][1]})
		nums1 = nums1[1:]
		nums2 = nums2[1:]
	} else {
		if nums1[0][0] < nums2[0][0] {
			res = append(res, nums1[0])
			nums1 = nums1[1:]
		} else {
			res = append(res, nums2[0])
			nums2 = nums2[1:]
		}
	}
	res = append(res, Solution(nums1, nums2)...)
	return res
}
