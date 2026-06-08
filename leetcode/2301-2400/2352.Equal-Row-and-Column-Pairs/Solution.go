package Solution

import (
	"fmt"
	"strings"
)

func Solution(grid [][]int) int {
	length := len(grid)
	ans := 0
	cols := make([]strings.Builder, length)
	rowsMap := make(map[string]int)

	var ssb strings.Builder
	for row := 0; row < length; row++ {
		for col := 0; col < length; col++ {
			_, _ = fmt.Fprintf(&ssb, "%d-", grid[row][col])
			_, _ = fmt.Fprintf(&cols[col], "%d-", grid[row][col])
		}
		rowsMap[ssb.String()]++
		ssb.Reset()
	}

	for _, sb := range cols {
		if count, ok := rowsMap[sb.String()]; ok {
			ans += count
		}
	}
	return ans
}
