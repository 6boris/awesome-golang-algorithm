package Solution

func Solution(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	ans := rows * cols
	need := 0
	for i := 0; i < rows; i++ {
		for s, e := 0, cols-1; s < e; s, e = s+1, e-1 {
			if grid[i][s] != grid[i][e] {
				need++
			}
		}
	}
	ans = min(ans, need)
	need = 0
	for i := 0; i < cols; i++ {
		for s, e := 0, rows-1; s < e; s, e = s+1, e-1 {
			if grid[s][i] != grid[e][i] {
				need++
			}
		}
	}
	return min(ans, need)
}
