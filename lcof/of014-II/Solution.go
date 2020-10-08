package Solution

func cuttingRope(n int) int {
	dp := make(map[int]int)
	dp[1] = 1 // 首项
	for i := 2; i < n+1; i++ {
		j, k := 1, i-1
		ans := 0
		for j <= k {
			ans = max(ans, max(j, dp[j])*max(k, dp[k])) // 递推公式
			j++
			k--
		}
		dp[i] = ans
	}
	return dp[n]
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
