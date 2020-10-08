package Solution

func findMaxAverage(nums []int, k int) float64 {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	ans := float64(sum[k-1]) / float64(k)
	for i := k; i < len(nums); i++ {
		ans = max(ans, float64(sum[i]-sum[i-k])/float64(k))
	}
	return ans
}

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}
