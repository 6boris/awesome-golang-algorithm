package Solution

func Solution(n int) int {
	if n <= 3 {
		return n
	}
	base := 1
	target := make([]int, n+1)
	for ; base*base <= n; base++ {
		target[base*base] = base * base
	}
	for idx := 1; idx <= n; idx++ {
		if target[idx] != 0 {
			base = target[idx]
			continue
		}
		target[idx] = base
	}

	dp := make([]int, n+1)
	dp[1], dp[2], dp[3] = 1, 2, 3
	for idx := 4; idx <= n; idx++ {
		if target[idx] == idx {
			dp[idx] = 1
			continue
		}

		start := target[idx]
		for ; start >= idx-start; start-- {
			diff := idx - start
			if dp[idx] == 0 || dp[idx] > dp[diff]+dp[start] {
				dp[idx] = dp[diff] + dp[start]
			}
		}
	}

	return dp[n]
}
