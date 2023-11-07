package Solution

func bfs1267(x, y, m, n int, grid [][]int) int {
	if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
		return 0
	}
	ans := 0
	grid[x][y] = 0
	queue := [][2]int{{x, y}}
	for len(queue) > 0 {
		ans += len(queue)
		nq := make([][2]int, 0)
		for _, item := range queue {
			// 判断同一行
			for i := 0; i < n; i++ {
				if grid[item[0]][i] != 0 {
					nq = append(nq, [2]int{item[0], i})
					grid[item[0]][i] = 0
				}
			}
			// 判断同一列
			for i := 0; i < m; i++ {
				if grid[i][item[1]] != 0 {
					nq = append(nq, [2]int{i, item[1]})
					grid[i][item[1]] = 0
				}
			}
		}
		queue = nq
	}
	return ans
}
func Solution(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] != 0 {
				if r := bfs1267(r, c, m, n, grid); r > 1 {
					ans += r
				}
			}
		}
	}
	return ans
}
