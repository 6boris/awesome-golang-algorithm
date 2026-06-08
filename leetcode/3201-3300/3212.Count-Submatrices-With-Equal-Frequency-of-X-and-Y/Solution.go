package Solution

func Solution(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	prefX := make([][]int, rows+1)
	prefY := make([][]int, rows+1)
	for i := range prefX {
		prefX[i] = make([]int, cols+1)
		prefY[i] = make([]int, cols+1)
	}

	res := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			xVal, yVal := 0, 0
			switch grid[i][j] {
			case 'X':
				xVal = 1
			case 'Y':
				yVal = 1
			}

			prefX[i+1][j+1] = prefX[i][j+1] + prefX[i+1][j] - prefX[i][j] + xVal
			prefY[i+1][j+1] = prefY[i][j+1] + prefY[i+1][j] - prefY[i][j] + yVal

			if prefX[i+1][j+1] > 0 && prefX[i+1][j+1] == prefY[i+1][j+1] {
				res++
			}
		}
	}

	return res
}
