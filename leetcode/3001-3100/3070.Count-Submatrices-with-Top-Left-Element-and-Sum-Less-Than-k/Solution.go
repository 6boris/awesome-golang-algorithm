package Solution

func Solution(grid [][]int, k int) int {
	if grid[0][0] > k {
		return 0
	}
	ans := 1
	for i := 1; i < len(grid[0]); i++ {
		grid[0][i] = grid[0][i-1] + grid[0][i]
		if grid[0][i] <= k {
			ans++
		}

	}
	for r := 1; r < len(grid); r++ {
		cur := 0
		for c := 0; c < len(grid[0]); c++ {
			cur += grid[r][c]
			tmp := cur + grid[r-1][c]
			if tmp <= k {
				ans++
			}
			grid[r][c] = tmp
		}
	}
	return ans
}
