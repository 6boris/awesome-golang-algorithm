package Solution

func Solution(grid [][]int) int {
	ans := 0
	rows, cols := len(grid), len(grid[0])
	for r := 0; r <= rows-3; r++ {
		for c := 0; c <= cols-3; c++ {
			sum1 := grid[r][c] + grid[r][c+1] + grid[r][c+2]
			sum2 := grid[r+2][c] + grid[r+2][c+1] + grid[r+2][c+2]
			if sum3 := sum1 + sum2 + grid[r+1][c+1]; sum3 > ans {
				ans = sum3
			}
		}
	}
	return ans
}
