package Solution

func Solution(grid [][]int) int {
	n := len(grid)
	pos := make([]int, n)
	for i := range pos {
		pos[i] = -1
	}

	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				pos[i] = j
				break
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		k := -1
		for j := i; j < n; j++ {
			if pos[j] <= i {
				ans += j - i
				k = j
				break
			}
		}

		if k != -1 {
			for j := k; j > i; j-- {
				pos[j], pos[j-1] = pos[j-1], pos[j]
			}
		} else {
			return -1
		}
	}
	return ans
}
