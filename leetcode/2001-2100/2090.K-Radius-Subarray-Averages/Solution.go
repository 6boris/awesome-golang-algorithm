package Solution

func Solution(nums []int, k int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = -1
	}
	start, end := 0, 2*k
	if end >= len(nums) {
		return ans
	}

	windowSize := 2*k + 1
	sum := 0
	for i := 0; i <= end; i++ {
		sum += nums[i]
	}
	ans[end-k] = sum / windowSize
	end++
	for ; end < len(nums); start, end = start+1, end+1 {
		sum -= nums[start]
		sum += nums[end]
		ans[end-k] = sum / windowSize
	}
	return ans
}
