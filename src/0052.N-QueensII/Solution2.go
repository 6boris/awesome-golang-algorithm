package Solution

import "math"

func totalNQueens2(n int) (result int) {
	columns := make([]int, n)
	totalNQueensHelper(n, 0, columns, &result)
	return result
}

func totalNQueensHelper(n, row int, columns []int, result *int) {
	if row == n {
		*result++
		return
	}
	for col := 0; col < n; col++ {
		columns[row] = col
		if valid(columns, row) {
			totalNQueensHelper(n, row+1, columns, result)
		}
	}
}

func valid(columns []int, row int) bool {
	for i := 0; i < row; i++ {
		diff := int(math.Abs(float64(columns[i] - columns[row])))
		// diff == 0            : Queens are in the same column
		// diff == lastRow - i  : Queens are placed diagonally
		//                        i.e, |col2 - col1| == |row2 - row1|
		if diff == 0 || diff == row-i {
			return false
		}
	}
	return true
}
