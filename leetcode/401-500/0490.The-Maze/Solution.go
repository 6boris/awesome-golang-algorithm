package Solution

func hasPath(maze [][]int, start []int, destination []int) bool {

	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}

	// 右， 左，下，上
	var dfs func(int, int) bool
	dfs = func(x, y int) bool {
		if visited[x][y] {
			return false
		}
		if x == destination[0] && y == destination[1] {
			return true
		}

		visited[x][y] = true
		up, down, left, right := x-1, x+1, y-1, y+1

		for right < len(maze[0]) && maze[x][right] == 0 {
			right++
		}
		if dfs(x, right-1) {
			return true
		}

		for left >= 0 && maze[x][left] == 0 {
			left--
		}
		if dfs(x, left+1) {
			return true
		}

		for up >= 0 && maze[up][y] == 0 {
			up--
		}
		if dfs(up+1, y) {
			return true
		}

		for down < len(maze) && maze[down][y] == 0 {
			down++
		}
		if dfs(down-1, y) {
			return true
		}

		return false
	}

	return dfs(start[0], start[1])
}
