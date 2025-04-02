package Solution

func Solution(nums []int) int64 {
	var ans int64
	cut := make([]int, len(nums))
	_max := nums[0]
	for i := 1; i < len(nums); i++ {
		cut[i] = _max - nums[i]
		_max = max(_max, nums[i])
	}
	_max = cut[1]
	for i := 2; i < len(nums); i++ {
		ans = max(ans, int64(nums[i])*int64(_max))
		_max = max(_max, cut[i])
	}
	return ans
}
