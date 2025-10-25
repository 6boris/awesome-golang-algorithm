package Solution

func nBits(n int) int {
	if n == 0 {
		return 1
	}
	bits := 0
	if n < 0 {
		bits++
		n = -n
	}
	for n > 0 {
		n /= 10
		bits++
	}
	return bits
}

func Solution(grid [][]int) []int {
	rows, cols := len(grid), len(grid[0])
	ret := make([]int, cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			ret[c] = max(ret[c], nBits(grid[r][c]))
		}
	}
	return ret
}
