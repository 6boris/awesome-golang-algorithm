package Solution

import "slices"

const INF2304 = 0x7fffffff

func Solution(grid [][]int, moveCost [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dp := [2][]int{}
	for i := range 2 {
		dp[i] = make([]int, cols)
		for j := range cols {
			dp[i][j] = INF2304
		}
	}
	for i := range cols {
		dp[0][i] = INF2304
	}
	for i := range cols {
		dp[1][i] = grid[rows-1][i]
	}

	index := 0
	for i := rows - 2; i >= 0; i-- {
		for c := 0; c < cols; c++ {
			for next := 0; next < cols; next++ {
				dp[index][c] = min(dp[index][c], grid[i][c]+dp[1-index][next]+moveCost[grid[i][c]][next])
			}
		}
		index = 1 - index
		for j := range cols {
			dp[index][j] = INF2304
		}
	}
	return slices.Min(dp[1-index])
}
