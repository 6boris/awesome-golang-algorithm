package Solution

func Solution(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	// 啥都不做
	if board[click[0]][click[1]] != 'E' {
		return board
	}
	board[click[0]][click[1]] = 'B'
	queue := [][2]int{{click[0], click[1]}}
	rows, cols := len(board), len(board[0])
	for len(queue) > 0 {
		// E的时候传播, 其他都静止
		nq := make([][2]int, 0)
		for _, cur := range queue {
			x, y := cur[0], cur[1]
			count := byte('0')
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}
					nx, ny := x+i, y+j
					if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
						if board[nx][ny] == 'M' {
							count++
						}
					}
				}
			}
			if count != '0' {
				board[x][y] = count
			} else {
				board[x][y] = 'B'
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						if i == 0 && j == 0 {
							continue
						}
						nx, ny := x+i, y+j
						if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
							if board[nx][ny] == 'E' {
								board[nx][ny] = 'B'
								nq = append(nq, [2]int{nx, ny})
							}
						}
					}
				}
			}
		}
		queue = nq
	}
	return board
}
