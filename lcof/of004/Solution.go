package Solution

func findNumberIn2DArray(board [][]int, target int) bool {
	rLen := len(board)
	cLen := len(board[0])

	//	从右上角的元素找起来
	for r, c := 0, cLen-1; r < rLen && c >= 0; {
		if board[r][c] == target {
			return true
		}
		if board[r][c] > target {
			c--
		} else {
			r++
		}
	}
	return false
}
