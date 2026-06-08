package Solution

func Solution(board [][]int) {
	rows, cols := len(board), len(board[0])

	ones := func(x, y int) int {
		count := 0
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}

				a, b := x+i, y+j

				if a >= 0 && a < rows && b >= 0 && b < cols && (board[a][b] == 1 || board[a][b] == 2) {
					count++
				}
			}
		}
		return count
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			count := ones(r, c)
			if count < 2 || count > 3 {
				if board[r][c] == 1 {
					// 之前是活着的，但是被迫挂了
					board[r][c] = 2
				}
			}

			// 被迫复活
			if count == 3 && board[r][c] == 0 {
				board[r][c] = 3
			}
		}
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 2 {
				board[r][c] = 0
			}
			if board[r][c] == 3 {
				board[r][c] = 1
			}
		}
	}
}
