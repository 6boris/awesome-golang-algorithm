package Solution

func Solution(board [][]byte) int {
	rows, cols := len(board), len(board[0])
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'X' {
				if i > 0 && board[i-1][j] == 'X' {
					continue
				}
				if j > 0 && board[i][j-1] == 'X' {
					continue
				}
				ans++
			}
		}
	}
	return ans
}
