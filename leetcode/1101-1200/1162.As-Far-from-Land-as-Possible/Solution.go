package Solution

type loc struct {
	x, y int
}

func Solution(grid [][]int) int {
	waterCell := make([]loc, 0)
	landCell := make([]loc, 0)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 1 {
				landCell = append(landCell, loc{row, col})
				continue
			}
			waterCell = append(waterCell, loc{row, col})
		}
	}
	ans := -1
	for _, water := range waterCell {
		x, y := water.x, water.y
		minDis := -1
		for _, land := range landCell {
			x1, y1 := land.x, land.y
			diff1 := x - x1
			if diff1 < 0 {
				diff1 = -diff1
			}
			diff2 := y - y1
			if diff2 < 0 {
				diff2 = -diff2
			}
			if minDis == -1 || minDis > diff1+diff2 {
				minDis = diff1 + diff2
			}
		}
		if minDis > ans {
			ans = minDis
		}
	}
	return ans
}
