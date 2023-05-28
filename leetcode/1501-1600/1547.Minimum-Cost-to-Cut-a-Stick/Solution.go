package Solution

func Solution(n int, cuts []int) int {
	dp := make(map[int]map[int]int)
	/*
		       n  很大，这么开内存受不了
				dp := make([][]int, n+1)
				for i := 0; i <= n; i++ {
					dp[i] = make([]int, n+1)
				}
	*/
	dp[0] = make(map[int]int)
	dp[n] = make(map[int]int)
	for _, c := range cuts {
		dp[c] = make(map[int]int)
	}

	var dfs func(int, int) int
	dfs = func(start, end int) int {
		if dp[start][end] != 0 {
			return dp[start][end]
		}
		need := end - start
		if need <= 1 {
			return 0
		}

		m := 0
		found := false
		for _, pos := range cuts {
			if pos > start && pos < end {
				found = true
				left := dfs(start, pos)
				right := dfs(pos, end)
				if m == 0 || left+right < m {
					m = left + right
				}
			}
		}
		if !found {
			return 0
		}
		dp[start][end] = need + m
		return dp[start][end]
	}
	return dfs(0, n)
}
