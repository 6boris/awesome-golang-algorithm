package Solution

func Solution(grid [][]int) int {
	var ret int
	m, n := len(grid), len(grid[0])
	index := n - 1
	for row := 0; row < m; row++ {
		for ; index >= 0 && grid[row][index] < 0; index-- {
		}
		ret += n - index - 1
	}
	return ret
}
