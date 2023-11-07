package Solution

func setZeroes(matrix [][]int) {
	zRow := make([]bool, len(matrix))
	zCol := make([]bool, len(matrix[0]))

	for i := range matrix {
		for j, v := range matrix[i] {
			if v == 0 {
				zRow[i] = true
				zCol[j] = true
			}
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if zRow[i] || zCol[j] {
				matrix[i][j] = 0
			}
		}
	}
}
