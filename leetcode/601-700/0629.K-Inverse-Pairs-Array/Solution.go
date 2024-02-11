package Solution

const mod629 = 1000000007

func Solution(n int, k int) int {
	dp := [2][]int{}
	sum := [2][]int{}
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, k+1)
		sum[i] = make([]int, k+1)
	}
	dp[1][0] = 1
	for i := 0; i <= k; i++ {
		sum[1][i] = 1
	}
	loop := 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= k; j++ {
			dp[1-loop][j] = 0
			if i > j {
				dp[1-loop][j] = (dp[1-loop][j] + sum[loop][j]) % mod629
			} else {
				dp[1-loop][j] = (dp[1-loop][j] + sum[loop][j] - sum[loop][j-i]) % mod629
			}

			add := 0
			if j > 0 {
				add = sum[1-loop][j-1]
			}
			sum[1-loop][j] = dp[1-loop][j] + add
		}
		loop = 1 - loop
	}
	return dp[loop][k]
}
