package Solution

func Solution(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, 3)
			for k := range memo[i][j] {
				memo[i][j][k] = -1 << 31
			}
		}
	}

	var dfs func(i, j, k int) int
	dfs = func(i, j, k int) int {
		if i >= m || j >= n {
			return -1 << 31
		}

		x := coins[i][j]
		// arrive at the destination
		if i == m-1 && j == n-1 {
			if k > 0 {
				return max(0, x)
			}
			return x
		}

		if memo[i][j][k] != -1<<31 {
			return memo[i][j][k]
		}

		// not neutralize
		res := max(dfs(i+1, j, k), dfs(i, j+1, k)) + x
		if k > 0 && x < 0 {
			// neutralize
			res = max(res, max(dfs(i+1, j, k-1), dfs(i, j+1, k-1)))
		}

		memo[i][j][k] = res
		return res
	}

	return dfs(0, 0, 2)
}
