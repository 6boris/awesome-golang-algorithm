package Solution

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	rLen := len(matrix)
	cLen := len(matrix[0])

	//	从右上角的元素找起来
	for r, c := 0, cLen-1; r < rLen && c >= 0; {
		if matrix[r][c] == target {
			return true
		}
		if matrix[r][c] > target {
			c--
		} else {
			r++
		}
	}
	return false
}
