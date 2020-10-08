package Solution

//	每次只检查网格的2边
func islandPerimeter(grid [][]int) int {
	ans := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			//	当前网格是陆地
			if grid[r][c] == 1 {
				ans += 4
				//	每次比较上边/和左边，如果这个方想有方块，将会减少2条边
				//	当前网格不是第一行且上一个方格是陆地
				if r > 0 && grid[r-1][c] == 1 {
					ans -= 2
				}
				//	当前网格不是第一列且右边的方格是陆地
				if c > 0 && grid[r][c-1] == 1 {
					ans -= 2
				}
			}
		}
	}
	return ans
}

//	依次检查每个网格的上下左右
func islandPerimeter2(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				//	上边：处于第一行或者上面一个网格为空
				if i == 0 || grid[i-1][j] == 0 {
					ans++
				}
				//	下边：处于最后一行或者下边一个网格为空
				if i == len(grid)-1 || grid[i+1][j] == 0 {
					ans++
				}
				//	左边：处于最左一行或者左边一个网格为空
				if j == 0 || grid[i][j-1] == 0 {
					ans++
				}
				//	右边：处于最右一行或者右边一个网格为空
				if j == len(grid[0])-1 || grid[i][j+1] == 0 {
					ans++
				}
			}
		}
	}
	return ans
}
