package Solution

func Solution(grid [][]int, health int) bool {
	m, n := len(grid), len(grid[0])
	dir := [4][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}
	seen := make([][]bool, m)
	for i := range m {
		seen[i] = make([]bool, n)
	}

	cache := make(map[[3]int]bool)
	var dfs func(int, int, int) bool
	dfs = func(x, y, h int) bool {
		if x == m-1 && y == n-1 {
			return h >= 1
		}
		if h == 0 {
			return false
		}
		key := [3]int{x, y, h}
		if v, ok := cache[key]; ok {
			return v
		}
		seen[x][y] = true
		var ok bool
		for _, d := range dir {
			nx, ny := x+d[0], y+d[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || seen[nx][ny] {
				continue
			}
			ok = dfs(nx, ny, h-grid[nx][ny])
			if ok {
				break
			}
		}
		cache[key] = ok
		seen[x][y] = false
		return ok
	}

	return dfs(0, 0, health-grid[0][0])
}
