package Solution

import "sort"

func Solution(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	ans := 0
	l := make([]int, cols)
	for row := rows - 1; row >= 0; row-- {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0 {
				l[col] = 0
				continue
			}
			if row < rows-1 {
				matrix[row][col] += matrix[row+1][col]
			}
			l[col] = matrix[row][col]
		}
		sort.Ints(l)

		total := 0
		for i := cols - 1; i >= 0; i-- {
			if l[i] == 0 {
				break
			}
			total++
			if r := l[i] * total; r > ans {
				ans = r
			}
		}
	}
	return ans
}
