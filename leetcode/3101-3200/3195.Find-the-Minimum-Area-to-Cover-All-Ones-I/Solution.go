package Solution

func Solution(grid [][]int) int {
	leftTop, rightBottom := [2]int{1001, 1001}, [2]int{-1, -1}
	rows, cols := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				leftTop[0] = min(leftTop[0], i)
				leftTop[1] = min(leftTop[1], j)

				rightBottom[0] = max(rightBottom[0], i)
				rightBottom[1] = max(rightBottom[1], j)
			}
		}
	}
	x := rightBottom[0] - leftTop[0]
	y := rightBottom[1] - leftTop[1]
	return (x + 1) * (y + 1)
}
