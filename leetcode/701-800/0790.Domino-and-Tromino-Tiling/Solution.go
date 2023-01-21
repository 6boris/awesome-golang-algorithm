package Solution

const mod790 = 1000000007

func Solution(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 3)
	}
	dp[1][0], dp[1][1], dp[1][2] = 1, 0, 0
	dp[2][0], dp[2][1], dp[2][2] = 2, 1, 1

	for i := 3; i <= n; i++ {
		dp[i][0] = (dp[i-1][0] + dp[i-2][0] + dp[i-1][1] + dp[i-1][2]) % mod790
		dp[i][1] = (dp[i-1][2] + dp[i-2][0]) % mod790
		dp[i][2] = (dp[i-1][1] + dp[i-2][0]) % mod790
	}
	return dp[n][0]
}
