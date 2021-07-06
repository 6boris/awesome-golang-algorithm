package Solution

import "strings"

func Solution(x int) [][]string {
	board := make([][]string, x)
	for i := range board {
		board[i] = make([]string, x)
		for j := range board[i] {
			board[i][j] = ".";
		}
	}
	res := make([][]string, 0)
	helper(board, &res, 0);
	return res
}

func helper(board [][]string, res *[][]string, row int) {
	if len(board) == row {
		temp := make([]string, 0);
		for i := range board {
			temp = append(temp, strings.Join(board[i], ""));
		}
		*res = append(*res, temp);
		return;
	}
	for col := 0; col < len(board); col++ {
		if isValid(board, row, col) {
			board[row][col] = "Q";
			helper(board, res, row + 1);
			board[row][col] = ".";
		}
	}
}

func isValid(board [][]string, r , c int) bool {
	for i := 0; i < r; i++ {
		if board[i][c] == "Q" {
			return false;
		}
	}
	for i, j := r, c; i >= 0 && j >= 0; i, j = i - 1, j - 1 {
		if board[i][j] == "Q" {
			return false;
		}
	}
	for i, j := r, c; i >= 0 && j < len(board); i, j = i - 1, j + 1 {
		if board[i][j] == "Q" {
			return false;
		}
	}
	return true;
}
