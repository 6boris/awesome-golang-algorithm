package Solution

func Solution(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]struct{})
	for _, n := range nums1 {
		m1[n] = struct{}{}
	}
	m2 := make(map[int]struct{})
	for _, n := range nums2 {
		m2[n] = struct{}{}
	}
	ans := []int{0, 0}
	for _, n := range nums1 {
		if _, ok := m2[n]; ok {
			ans[0]++
		}
	}
	for _, n := range nums2 {
		if _, ok := m1[n]; ok {
			ans[1]++
		}
	}
	return ans
}
