package Solution

func Solution(nums1 []int, nums2 []int, k int) int64 {
	var ret int64
	cnt := map[int]int64{}
	for i := range nums2 {
		cnt[nums2[i]*k]++
	}
	for i := range nums1 {
		num := nums1[i]
		for d := 1; d*d <= num; d++ {
			if num%d == 0 {
				if c, ok := cnt[d]; ok {
					ret += c
				}

				other := num / d
				if other != d {
					if c, ok := cnt[other]; ok {
						ret += c
					}
				}
			}
		}
	}
	return ret
}
