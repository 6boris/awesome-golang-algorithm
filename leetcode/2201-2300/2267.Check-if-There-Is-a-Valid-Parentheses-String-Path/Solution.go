package Solution

func Solution(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	if grid[0][0] == ')' || grid[m-1][n-1] == '(' {
		return false
	}
	if (m+n)%2 == 0 {
		return false
	}
	cache := map[[3]int]bool{
		[3]int{0, 0, 1}: true,
	}
	var dfs func(int, int, int) bool
	dfs = func(x, y, c int) bool {
		if x < 0 || y < 0 || c < 0 || c > x+y+1 {
			return false
		}

		key := [3]int{x, y, c}
		if v, ok := cache[key]; ok {
			return v
		}
		ret := false
		if grid[x][y] == ')' {
			// 前面的需要c+1
			ret = dfs(x-1, y, c+1) || dfs(x, y-1, c+1)
		} else {
			// 如果是(
			ret = dfs(x-1, y, c-1) || dfs(x, y-1, c-1)
		}

		cache[key] = ret
		return ret
	}

	return dfs(m-1, n-1, 0)
}
