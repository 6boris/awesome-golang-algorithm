package Solution

var dirs1905 = [][]int{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
}

func Solution(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	var dfs func(int, int, *bool)
	dfs = func(x, y int, ok *bool) {
		if x < 0 || x >= m || y < 0 || y >= n || grid2[x][y] == -1 || grid2[x][y] == 0 {
			return
		}
		grid2[x][y] = -1
		if grid1[x][y] != 1 {
			*ok = false
		}
		for _, dir := range dirs1905 {
			nx, ny := dir[0]+x, dir[1]+y
			dfs(nx, ny, ok)
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				ok := true
				dfs(i, j, &ok)
				if ok {
					ans++
				}
			}
		}
	}
	return ans
}
