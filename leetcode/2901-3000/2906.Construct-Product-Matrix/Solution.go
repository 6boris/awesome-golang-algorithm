package Solution

func Solution(grid [][]int) [][]int {
	/*
			1, 2, 3
		   4, 5, 6
		   7, 8, 9
	*/
	m, n := len(grid), len(grid[0])
	ret := make([][]int, m)
	total := m * n
	left, right := make([]int, total), make([]int, total)
	index, prod := 0, 1

	for i := range m {
		ret[i] = make([]int, n)
		for j := range n {
			prod = (prod * grid[i][j]) % 12345
			left[index] = prod
			index++
		}
	}

	index = total - 1
	prod = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			prod = (prod * grid[i][j]) % 12345
			right[index] = prod
			index--
		}
	}

	base := 1
	for i := range m {
		for j := range n {
			base = 1
			index = i*n + j
			if index >= 1 {
				base = (base * left[index-1]) % 12345
			}
			if index < total-1 {
				base = (base * right[index+1]) % 12345
			}
			ret[i][j] = base
		}
	}
	return ret
}
