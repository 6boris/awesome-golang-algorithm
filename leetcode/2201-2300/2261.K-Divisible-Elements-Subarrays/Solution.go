package Solution

func hash2261(nums []int, start, end int) uint64 {
	h := uint64(0)
	for i := start; i < end; i++ {
		h = h*211 + uint64(nums[i])
	}
	return h
}

func Solution(nums []int, k int, p int) int {
	ans := 0
	count := make([]int, len(nums))
	check := make(map[uint64]struct{})
	check[hash2261(nums, 0, 1)] = struct{}{}
	ans++
	if nums[0]%p == 0 {
		count[0] = 1
	}
	// 6,20,5,18
	for i := 1; i < len(nums); i++ {
		add := 0
		if nums[i]%p == 0 {
			add = 1
		}
		// 1, 2, 3, 4
		count[i] = count[i-1] + add
		for pre := i; pre >= 0; pre-- {
			cut := 0
			if pre > 0 {
				cut = count[pre-1]
			}

			diff := count[i] - cut
			if diff > k {
				break
			}
			key := hash2261(nums, pre, i+1)
			if _, ok := check[key]; ok {
				continue
			}
			check[key] = struct{}{}
			ans++
		}
	}
	return ans
}
