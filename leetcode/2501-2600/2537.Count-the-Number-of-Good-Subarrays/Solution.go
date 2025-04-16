package Solution

func Solution(nums []int, k int) int64 {
	l := len(nums)
	count := make(map[int]int64)
	start, end := 0, 0

	var (
		ans, pairs int64
	)

	i64k := int64(k)
	for ; end < l; end++ {
		pairs += count[nums[end]]
		count[nums[end]]++

		for ; start < end && pairs >= i64k; start++ {
			ans += int64(l - end)
			count[nums[start]]--
			pairs -= count[nums[start]]
		}
	}
	return ans
}
