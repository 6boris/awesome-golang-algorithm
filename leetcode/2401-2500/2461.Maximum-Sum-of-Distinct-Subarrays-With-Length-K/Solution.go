package Solution

func Solution(nums []int, k int) int64 {
	l := len(nums)
	count := make(map[int]int)
	cur := int64(0)
	for i := 0; i < k; i++ {
		count[nums[i]]++
		cur += int64(nums[i])
	}

	ans := int64(0)
	if len(count) == k {
		ans = cur
	}

	// [1, 2], 3, 4,
	s, e := 0, k
	for ; e < l; s, e = s+1, e+1 {
		count[nums[s]]--
		cur -= int64(nums[s])
		if count[nums[s]] == 0 {
			delete(count, nums[s])
		}
		count[nums[e]]++
		cur += int64(nums[e])
		if len(count) == k {
			ans = max(ans, cur)
		}
	}
	return ans
}
