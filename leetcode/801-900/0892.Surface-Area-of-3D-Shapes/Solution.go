package Solution

func Solution(grid [][]int) int {
	length := len(grid)
	surfaces := make([][]int, len(grid))
	for row := 0; row < length; row++ {
		surfaces[row] = make([]int, length)
		for col := 0; col < length; col++ {
			if grid[row][col] != 0 {
				surfaces[row][col] = 4*grid[row][col] + 2
			}
		}
	}

	dirs := [4][2]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}
	ans := 0
	for row := 0; row < length; row++ {
		for col := 0; col < length; col++ {
			for _, dir := range dirs {
				nextRow, nextCol := row+dir[0], col+dir[1]
				if nextRow >= 0 && nextRow < length && nextCol >= 0 && nextCol < length && grid[nextRow][nextCol] != 0 {
					diff := grid[row][col]
					if grid[nextRow][nextCol] < diff {
						diff = grid[nextRow][nextCol]
					}
					surfaces[row][col] -= diff
				}
			}
			ans += surfaces[row][col]
		}
	}
	return ans
}
