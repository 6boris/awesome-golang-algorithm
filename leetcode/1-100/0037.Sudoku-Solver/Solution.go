package Solution

func solveSudoku(board [][]byte) {
	solve(board, 0)
}

/* k 是把 board 转换成一维数组后，元素的索引值 */
func solve(board [][]byte, k int) bool {
	if k == 81 {
		return true
	}

	r, c := k/9, k%9
	if board[r][c] != '.' {
		return solve(board, k+1)
	}

	/* bi, bj 是 rc 所在块的左上角元素的索引值 */
	bi, bj := r/3*3, c/3*3

	// 按照数独的规则，检查 b 能否放在 board[r][c]
	isValid := func(b byte) bool {
		for n := 0; n < 9; n++ {
			if board[r][n] == b ||
				board[n][c] == b ||
				board[bi+n/3][bj+n%3] == b {
				return false
			}
		}
		return true
	}

	for b := byte('1'); b <= '9'; b++ {
		if isValid(b) {
			board[r][c] = b
			if solve(board, k+1) {
				return true
			}
		}
	}

	board[r][c] = '.'

	return false
}
