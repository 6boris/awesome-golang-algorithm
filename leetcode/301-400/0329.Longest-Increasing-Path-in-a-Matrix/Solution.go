package Solution

func Solution(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	path := make([][]int, rows)
	for i := 0; i < rows; i++ {
		path[i] = make([]int, cols)
	}
	var (
		dfs  func(int, int) int
		dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	)
	dfs = func(x, y int) int {
		if path[x][y] != 0 {
			return path[x][y]
		}
		selfMaxL := 1
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols || matrix[nx][ny] <= matrix[x][y] {
				continue
			}
			if r := dfs(nx, ny); r+1 > selfMaxL {
				selfMaxL = r + 1
			}
		}
		path[x][y] = selfMaxL
		return selfMaxL
	}
	ans := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if r := dfs(r, c); r > ans {
				ans = r
			}
		}
	}
	return ans
}
