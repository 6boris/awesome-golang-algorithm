package Solution

func Solution(n int) bool {
	dp := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if !dp[i-j*j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

func Solution1(n int) bool {
	cache := map[int]bool{}
	var dfs func(int) bool
	dfs = func(c int) bool {
		if c == 0 {
			return false
		}

		if v, ok := cache[c]; ok {
			return v
		}

		for x := 1; x*x <= c; x++ {
			left := c - x*x
			if !dfs(left) {
				cache[c] = true
				break
			}
		}
		return cache[c]
	}

	return dfs(n)
}
