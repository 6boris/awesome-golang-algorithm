package Solution

func Find(board [][]int, target int) bool {
	rlen := len(board)
	clen := len(board[0])

	//	从右上角的元素找起来
	for r, c := 0, clen-1; r < rlen && c >= 0; {
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
