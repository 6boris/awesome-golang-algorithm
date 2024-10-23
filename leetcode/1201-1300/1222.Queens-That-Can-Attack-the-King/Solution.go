package Solution

func Solution(queens [][]int, king []int) [][]int {
	ans := make([][]int, 0)
	board := [8][8]int{}
	for _, q := range queens {
		board[q[0]][q[1]] = 1
	}
	x, y := king[0], king[1]
	for r := y - 1; r >= 0; r-- {
		if board[x][r] == 1 {
			ans = append(ans, []int{x, r})
			break
		}
	}
	for r := y + 1; r < 8; r++ {
		if board[x][r] == 1 {
			ans = append(ans, []int{x, r})
			break
		}
	}

	for r := x - 1; r >= 0; r-- {
		if board[r][y] == 1 {
			ans = append(ans, []int{r, y})
			break
		}
	}

	for r := x + 1; r < 8; r++ {
		if board[r][y] == 1 {
			ans = append(ans, []int{r, y})
			break
		}
	}

	// 对
	for l, r := x-1, y-1; l >= 0 && r >= 0; l, r = l-1, r-1 {
		if board[l][r] == 1 {
			ans = append(ans, []int{l, r})
			break
		}
	}

	for l, r := x-1, y+1; l >= 0 && r < 8; l, r = l-1, r+1 {
		if board[l][r] == 1 {
			ans = append(ans, []int{l, r})
			break
		}
	}

	for l, r := x+1, y+1; l < 8 && r < 8; l, r = l+1, r+1 {
		if board[l][r] == 1 {
			ans = append(ans, []int{l, r})
			break
		}
	}

	for l, r := x+1, y-1; l < 8 && r >= 0; l, r = l+1, r-1 {
		if board[l][r] == 1 {
			ans = append(ans, []int{l, r})
			break
		}
	}

	return ans
}
