package Solution

func Solution(strs []string, m, n int) int {
	store := make([][2]int, len(strs))
	dp := make([][][]int, len(strs)+1)
	for x := 0; x < len(strs)+1; x++ {
		dp[x] = make([][]int, m+1)
		for y := 0; y < m+1; y++ {
			dp[x][y] = make([]int, n+1)
		}
		if x < len(strs) {
			for _, b := range []byte(strs[x]) {
				store[x][b-'0']++
			}
		}
	}

	max := 0
	for idx := 1; idx <= len(strs); idx++ {
		for zero := 0; zero <= m; zero++ {
			for one := 0; one <= n; one++ {
				dp[idx][zero][one] = dp[idx-1][zero][one]
				zeros, ones := store[idx-1][0], store[idx-1][1]

				if zero >= zeros && one >= ones && dp[idx-1][zero-zeros][one-ones]+1 > dp[idx][zero][one] {
					dp[idx][zero][one] = dp[idx-1][zero-zeros][one-ones] + 1
				}
				if dp[idx][zero][one] > max {
					max = dp[idx][zero][one]
				}
			}
		}
	}
	return max
}
