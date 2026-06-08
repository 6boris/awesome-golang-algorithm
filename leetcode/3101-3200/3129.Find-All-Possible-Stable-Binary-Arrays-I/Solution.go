package Solution

func Solution(zero, one, limit int) int {
	dp := make([][][2]int, zero+1)
	mod := int(1e9 + 7)
	for i := 0; i <= zero; i++ {
		dp[i] = make([][2]int, one+1)
	}
	for i := 0; i <= min(zero, limit); i++ {
		dp[i][0][0] = 1
	}
	for j := 0; j <= min(one, limit); j++ {
		dp[0][j][1] = 1
	}
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			if i > limit {
				dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1] - dp[i-limit-1][j][1]
			} else {
				dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1]
			}
			dp[i][j][0] = (dp[i][j][0]%mod + mod) % mod
			if j > limit {
				dp[i][j][1] = dp[i][j-1][1] + dp[i][j-1][0] - dp[i][j-limit-1][0]
			} else {
				dp[i][j][1] = dp[i][j-1][1] + dp[i][j-1][0]
			}
			dp[i][j][1] = (dp[i][j][1]%mod + mod) % mod
		}
	}
	return (dp[zero][one][0] + dp[zero][one][1]) % mod
}
