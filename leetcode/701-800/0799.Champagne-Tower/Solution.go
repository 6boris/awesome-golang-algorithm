package Solution

func Solution(poured int, query_row int, query_glass int) float64 {
	tower := make([][]float64, 101)
	for i := 0; i < 101; i++ {
		tower[i] = make([]float64, 101)
	}
	tower[0][0] = float64(poured)
	for row := 1; row <= query_row; row++ {
		for col := 0; col < row; col++ {
			if tower[row-1][col] > 1.0 {
				left := (tower[row-1][col] - 1.0) / 2.0
				tower[row][col] += left
				tower[row][col+1] += left
			}
		}
	}
	if tower[query_row][query_glass] > 1.0 {
		return 1.0
	}
	return tower[query_row][query_glass]
}
