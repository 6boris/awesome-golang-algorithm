package Solution

func Solution(grid [][]int, x, y, k int) [][]int {
	for s, e := x, x+k-1; s < e; s, e = s+1, e-1 {
		for j := y; j < y+k; j++ {
			grid[s][j], grid[e][j] = grid[e][j], grid[s][j]
		}
	}
	return grid
}
