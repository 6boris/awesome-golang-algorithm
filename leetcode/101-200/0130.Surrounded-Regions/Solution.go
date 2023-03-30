package Solution

func Solution(board [][]byte) {
	rows, cols := len(board), len(board[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if row == 0 || row == rows-1 || col == 0 || col == cols-1 {
				fillEdges(board, row, col)
			}
		}
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == '1' {
				board[row][col] = 'O'
				continue
			}
			if board[row][col] == 'O' {
				board[row][col] = 'X'
			}
		}
	}
}

// 边上填充‘1’
func fillEdges(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != 'O' {
		return
	}

	board[x][y] = '1'
	fillEdges(board, x+1, y)
	fillEdges(board, x-1, y)
	fillEdges(board, x, y-1)
	fillEdges(board, x, y+1)
}
