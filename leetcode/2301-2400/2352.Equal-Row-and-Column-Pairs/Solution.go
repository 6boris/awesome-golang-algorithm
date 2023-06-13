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
			ssb.WriteString(fmt.Sprintf("%d-", grid[row][col]))
			cols[col].WriteString(fmt.Sprintf("%d-", grid[row][col]))
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
