package Solution

func Solution(grid [][]int) []int {
	m := len(grid)
	base, sum := m*m, 0
	n := make([]int, base)
	target := base * (base + 1) / 2
	repeate := 1
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			n[grid[i][j]-1]++
			if n[grid[i][j]-1] == 2 {
				repeate = grid[i][j]
			}
			sum += grid[i][j]
		}
	}
	diff := sum - target
	ans := []int{repeate, repeate - diff}
	return ans
}
