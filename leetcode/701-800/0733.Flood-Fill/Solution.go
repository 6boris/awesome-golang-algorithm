package Solution

var localDirs = [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func floodFillDFS(image, visited [][]int, sr, sc, rows, cols int, sourceColor, newColor int) {
	if sr < 0 || sr >= rows || sc < 0 || sc >= cols {
		return
	}

	if visited[sr][sc] > 0 || image[sr][sc] != sourceColor {
		return
	}

	image[sr][sc] = newColor
	visited[sr][sc] = 1
	for _, dir := range localDirs {
		floodFillDFS(image, visited, sr+dir[0], sc+dir[1], rows, cols, sourceColor, newColor)
	}
}

func Solution(image [][]int, sr, sc, newColor int) [][]int {
	rows, cols := len(image), len(image[0])
	visited := make([][]int, len(image))
	for idx := 0; idx < rows; idx++ {
		visited[idx] = make([]int, cols)
	}

	floodFillDFS(image, visited, sr, sc, rows, cols, image[sr][sc], newColor)
	return image
}

func floodFill_dfs(image [][]int, sr int, sc int, newColor int) [][]int {
	var dfs func([][]int, int, int, int, int)
	dfs = func(image [][]int, r, c, oldColor, newColor int) {
		if c < 0 || r < 0 || r >= len(image) || c >= len(image[0]) || image[r][c] != oldColor {
			return
		}
		image[r][c] = newColor
		for _, dir := range localDirs {
			dfs(image, r+dir[0], c+dir[1], oldColor, newColor)
		}
	}
	if image[sr][sc] != newColor {
		dfs(image, sr, sc, image[sr][sc], newColor)
	}
	return image
}

func floodFill_bfs(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	n, m := len(image), len(image[0])
	queue := [][]int{}
	queue = append(queue, []int{sr, sc})
	image[sr][sc] = newColor
	for len(queue) > 0 {
		tmp := queue[0]
		queue = queue[1:]
		for _, dir := range localDirs {
			r, c := tmp[0]+dir[0], tmp[1]+dir[1]
			if r >= 0 && r < n && c >= 0 && c < m && image[r][c] == oldColor {
				queue = append(queue, []int{r, c})
				image[r][c] = newColor
			}
		}
	}
	return image
}
