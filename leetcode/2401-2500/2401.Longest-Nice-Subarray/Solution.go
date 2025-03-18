package Solution

func Solution(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	cur := 0
	start, end := 0, 0
	ans := 0
	for ; end < len(nums); end++ {
		if cur&nums[end] == 0 {
			cur |= nums[end]
			continue
		}
		ans = max(ans, end-start)

		for ; start <= end; start++ {
			cur ^= nums[start]
			if cur&nums[end] == 0 {
				start++
				cur |= nums[end]
				break
			}
		}
	}
	ans = max(ans, end-start)
	return ans
}
