package Solution

func Solution(nums1 []int, nums2 []int) []int {
	in := [1001]int{}
	for _, n := range nums1 {
		in[n]++
	}
	ans := make([]int, 0)
	for _, n := range nums2 {
		if in[n] > 0 {
			ans = append(ans, n)
			in[n]--
		}
	}
	return ans
}
