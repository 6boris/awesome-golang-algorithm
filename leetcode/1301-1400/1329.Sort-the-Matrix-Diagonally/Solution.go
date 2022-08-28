package Solution

func Solution(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])
	x, y := 0, 0
	for ; y < cols-1; y++ {
		sort1329(mat, x, y, rows, cols)
	}
	y = 0
	for ; x < rows-1; x++ {
		sort1329(mat, x, y, rows, cols)
	}
	return mat
}

func sort1329(mat [][]int, x, y, rows, cols int) {
	for cmpTimes := 0; cmpTimes < rows-x-1; cmpTimes++ {
		for x1, y1 := x, y; x1 < rows-1 && y1 < cols-1; x1, y1 = x1+1, y1+1 {
			nx, ny := x1+1, y1+1
			if mat[x1][y1] > mat[nx][ny] {
				mat[nx][ny], mat[x1][y1] = mat[x1][y1], mat[nx][ny]
			}
		}
	}
}
