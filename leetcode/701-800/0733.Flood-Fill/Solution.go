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
