package Solution

func Solution(n int, delay int, forget int) int {
	MOD := 1_000_000_007
	dp := make([]int, n+forget+1)
	inc := make([]int, n+forget+1)

	// init
	dp[1] = 1
	dp[1+forget] = -1

	for j := 1 + delay; j < 1+forget; j++ {
		inc[j] = 1
	}

	for i := 2; i <= n; i++ {
		dp[i] = (dp[i] + dp[i-1] + inc[i]) % MOD
		dp[i+forget] = (dp[i+forget] - inc[i]) % MOD
		for j := i + delay; j < i+forget; j++ {
			inc[j] = (inc[j] + inc[i]) % MOD
		}
	}
	result := dp[n] % MOD
	if result < 0 {
		result += MOD
	}
	return result
}
