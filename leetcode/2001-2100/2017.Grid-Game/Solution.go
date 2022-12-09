package Solution

func Solution(grid [][]int) int64 {
	column := len(grid[0])
	for col := 1; col < column; col++ {
		grid[0][col] += grid[0][col-1]
	}

	for col := column - 2; col >= 0; col-- {
		grid[1][col] += grid[1][col+1]
	}

	secondMax := int64(-1)
	for firstDown := 0; firstDown < column; firstDown++ {
		// 第一个机器人在哪个节点向下转动
		top := int64(grid[0][column-1] - grid[0][firstDown])
		down := int64(grid[1][0] - grid[1][firstDown])
		if top < down {
			top = down
		}
		if secondMax == -1 || top < secondMax {
			secondMax = top
		}
	}

	return secondMax
}
