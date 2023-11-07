package Solution

func Solution(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int, [][]bool) int
	dfs = func(x, y int, visited [][]bool) int {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
			return 0
		}
		if visited[x][y] {
			return 0
		}
		visited[x][y] = true
		return grid[x][y] + dfs(x+1, y, visited) + dfs(x-1, y, visited) + dfs(x, y+1, visited) + dfs(x, y-1, visited)
	}
	ans := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 0 {
				continue
			}
			v := make([][]bool, m)
			for i := 0; i < n; i++ {
				v[i] = make([]bool, n)
			}
			if r := dfs(r, c, v); r > ans {
				ans = r
			}
		}
	}
	return ans
}
