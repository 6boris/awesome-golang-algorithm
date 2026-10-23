package Solution

func Solution(board [][]byte, rMove int, cMove int, color byte) bool {
	m, n := len(board), len(board[0])
	helper := func(xDir, yDir int) bool {
		x, y := rMove+xDir, cMove+yDir
		oppositeColor := byte('W')
		if color == 'W' {
			oppositeColor = 'B'
		}
		cnt := 1 // endpoint
		for ; x >= 0 && x < m && y >= 0 && y < n; x, y = x+xDir, y+yDir {
			if board[x][y] == color {
				return cnt >= 2
			}

			if board[x][y] != oppositeColor {
				return false
			}
			cnt++
		}
		return false
	}

	for _, dir := range [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	} {
		if helper(dir[0], dir[1]) {
			return true
		}
	}
	return false
}
