package Solution

func findMaxAverage(nums []int, k int) float64 {
	tmp, ans := 0, -1
	for i := 0; i < len(nums); i++ {
		tmp += nums[i]
		if i >= k-1 {
			ans = max(ans, tmp)
			tmp -= nums[i-k+1]
		}
	}

	return float64(ans) / float64(k)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
