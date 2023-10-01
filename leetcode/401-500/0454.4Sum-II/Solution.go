package Solution

func Solution(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	l := len(nums1)
	cache := make(map[int]int)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			cache[nums1[i]+nums2[j]]++
		}
	}
	ans := 0
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			target := 0 - nums3[i] - nums4[j]
			if v, ok := cache[target]; ok {
				ans += v
			}
		}
	}
	return ans
}
