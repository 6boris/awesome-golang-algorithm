package Solution

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	l, r := 0, n*m-1

	for l < r {
		mid := (l + r) / 2
		if matrix[mid/m][mid%m] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return matrix[r/m][r%m] == target
}
