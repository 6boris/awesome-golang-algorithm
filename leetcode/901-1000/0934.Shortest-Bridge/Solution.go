package Solution

type pair934 struct {
	x, y int
}

var dirs934 = [][]int{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
}

func Solution(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	x, y := -1, -1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				x, y = i, j
				break
			}
		}
		if x != -1 || y != -1 {
			break
		}
	}
	pairs := make([]pair934, 0)
	var dfs func(int, int)
	dfs = func(x, y int) {
		if x < 0 || x >= rows || y < 0 || y >= cols {
			return
		}
		if grid[x][y] == 0 || grid[x][y] == 2 {
			return
		}
		// 第一组陆地
		grid[x][y] = 2
		pairs = append(pairs, pair934{x, y})
		for _, dir := range dirs934 {
			nx, ny := x+dir[0], y+dir[1]
			dfs(nx, ny)
		}
	}
	dfs(x, y)
	var bfs func(int, int) int
	visited := make([][]bool, rows)
	for r := 0; r < rows; r++ {
		visited[r] = make([]bool, cols)
	}
	bfs = func(x, y int) int {
		path := 0
		q := []pair934{{x, y}}
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				visited[r][c] = false
			}
		}
		visited[x][y] = true
		for len(q) > 0 {
			next := make([]pair934, 0)
			for _, item := range q {
				if grid[item.x][item.y] == 1 {
					return path - 1
				}
				for _, dir := range dirs934 {
					nx, ny := item.x+dir[0], item.y+dir[1]
					if nx < 0 || nx >= rows || ny < 0 || ny >= cols || grid[nx][ny] == 2 || visited[nx][ny] {
						continue
					}
					visited[nx][ny] = true
					next = append(next, pair934{nx, ny})
				}
			}
			path++
			q = next
		}
		return -1
	}
	ans := -1
	for _, p := range pairs {
		path := bfs(p.x, p.y)
		if ans == -1 || (path != -1 && path < ans) {
			ans = path
		}
	}
	return ans
}
