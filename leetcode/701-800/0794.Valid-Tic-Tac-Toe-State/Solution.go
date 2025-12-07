package Solution

func win794(board []string) (bool, bool) {
	x, o := false, false
	for i := range 3 {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			if board[i][0] == 'X' {
				x = true
			}
			if board[i][0] == 'O' {
				o = true
			}
		}

		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			if board[0][i] == 'X' {
				x = true
			}
			if board[0][i] == 'O' {
				o = true
			}
		}
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		if board[0][0] == 'X' {
			x = true
		}
		if board[0][0] == 'O' {
			o = true
		}
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		if board[0][2] == 'X' {
			x = true
		}
		if board[0][2] == 'O' {
			o = true
		}
	}
	return x, o
}

func Solution(board []string) bool {
	xWin, oWin := win794(board)
	if xWin && oWin {
		return false
	}
	x, o := 0, 0
	for i := range 3 {
		for j := range 3 {
			if board[i][j] == 'X' {
				x++
			}
			if board[i][j] == 'O' {
				o++
			}
		}
	}
	if !xWin && !oWin {
		return x == o || x == o+1
	}
	if xWin {
		return x == o+1
	}
	return x == o
}
