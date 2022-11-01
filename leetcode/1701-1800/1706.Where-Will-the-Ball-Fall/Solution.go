package Solution

func Solution(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	ans := make([]int, n)

	for col := 0; col < n; col++ {
		cur := col
		for row := 0; row < m; row++ {
			nextCol := cur + grid[row][cur]

			/* /\  \/ */
			if !(nextCol >= 0 && nextCol < n && grid[row][cur] == grid[row][nextCol]) {
				ans[col] = -1
				break
			}
			ans[col] = nextCol
			cur = nextCol
		}
	}
	return ans
}
