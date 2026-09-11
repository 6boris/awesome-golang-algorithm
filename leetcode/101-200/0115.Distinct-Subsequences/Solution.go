package Solution

func Solution(s string, t string) int {
	ls, lt := len(s), len(t)
	dp := make([]int, lt+1)
	dp[0] = 1
	var diag, tmp int
	for i := range ls {
		diag = dp[0]
		for j := 1; j <= lt; j++ {
			tmp, diag = diag, dp[j]
			diag = dp[j]
			if s[i] == t[j-1] {
				dp[j] += tmp
			}
		}
	}
	return dp[lt]
}
