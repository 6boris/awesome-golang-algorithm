package Solution

func Solution(board [][]byte) int {
	rook := [2]int{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				rook[0], rook[1] = i, j
			}
		}
	}
	pawns := 0
	for i := rook[0]; i < 8; i++ {
		if board[i][rook[1]] == 'B' {
			break
		} else if board[i][rook[1]] == 'p' {
			pawns++
			break
		}
	}
	for i := rook[0]; i >= 0; i-- {
		if board[i][rook[1]] == 'B' {
			break
		} else if board[i][rook[1]] == 'p' {
			pawns++
			break
		}
	}
	for i := rook[0]; i < 8; i++ {
		if board[rook[0]][i] == 'B' {
			break
		} else if board[rook[0]][i] == 'p' {
			pawns++
			break
		}
	}
	for i := rook[0]; i >= 0; i-- {
		if board[rook[0]][i] == 'B' {
			break
		} else if board[rook[0]][i] == 'p' {
			pawns++
			break
		}
	}
	return pawns
}
