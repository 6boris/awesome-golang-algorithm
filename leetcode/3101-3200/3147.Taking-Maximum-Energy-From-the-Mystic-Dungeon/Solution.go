package Solution

func Solution(energy []int, k int) int {
	dp := make([]int, len(energy))
	// k = 2
	// 1, 2, 3, 4,
	l := len(energy)
	ret := energy[l-1]
	for i := l - 1; i >= l-k; i-- {
		// 这些是无法跳的，所以默认就有自己的值
		dp[i] = energy[i]
		ret = max(ret, dp[i])
	}
	for i := l - k - 1; i >= 0; i-- {
		dp[i] = dp[i+k] + energy[i]
		ret = max(ret, dp[i])
	}
	return ret
}
