package Solution

func Solution(moves [][]int) string {
	board := make([][]string, 3)
	toggle := true
	for i := range board {
		board[i] = make([]string, 3)
	}
	for _, move := range moves {
		i, j := move[0], move[1]
		if toggle {
			board[i][j] = "X"
			toggle = false
		} else {
			board[i][j] = "O"
			toggle = true
		}
	}
	return winner(board)
}

func winner(board [][]string) string {
	player := checkRow(board)
	switch player {
	case "X":
		return "A"
	case "O":
		return "B"
	}
	player = checkColumn(board)
	switch player {
	case "X":
		return "A"
	case "O":
		return "B"
	}
	player = checkDiagonals(board)
	switch player {
	case "X":
		return "A"
	case "O":
		return "B"
	}
	count := 0
	for _, v := range board {
		for _, val := range v {
			if val == "X" || val == "O" {
				count++
			}
		}
	}
	if count == 9 {
		return "Draw"
	} else {
		return "Pending"
	}
}

func checkRow(moves [][]string) string {
	for i := 0; i < 3; i++ {
		x, o := 0, 0
		for j := 0; j < 3; j++ {
			switch moves[i][j] {
			case "X":
				x++
			case "O":
				o++
			}
		}
		if x == 3 {
			return "X"
		} else if o == 3 {
			return "O"
		}
	}
	return "D"
}

func checkColumn(moves [][]string) string {
	for i := 0; i < 3; i++ {
		x, o := 0, 0
		for j := 0; j < 3; j++ {
			switch moves[j][i] {
			case "X":
				x++
			case "O":
				o++
			}
		}
		if x == 3 {
			return "X"
		} else if o == 3 {
			return "O"
		}
	}
	return "D"
}

func checkDiagonals(moves [][]string) string {
	x, o := 0, 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				switch moves[i][j] {
				case "X":
					x++
				case "O":
					o++
				}
			}
		}
		if x == 3 {
			return "X"
		} else if o == 3 {
			return "O"
		}
	}
	x, o = 0, 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i+j == 2 {
				switch moves[i][j] {
				case "X":
					x++
				case "O":
					o++
				}
			}
		}
		if x == 3 {
			return "X"
		} else if o == 3 {
			return "O"
		}
	}
	return "D"
}
