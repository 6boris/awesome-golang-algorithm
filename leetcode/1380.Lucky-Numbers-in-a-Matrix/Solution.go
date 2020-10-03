package Solution

func Solution(matrix [][]int) []int {
	ans, m, n := []int{}, len(matrix), len(matrix[0])
	for j := 0; j < n; j++ {
		max, rowIndex := 1, 0
		for i := 0; i < m; i++ {
			if matrix[i][j] > max {
				max = matrix[i][j]
				rowIndex = i
			}
		}
		if Min(matrix[rowIndex]) == max {
			ans = append(ans, max)
		}
	}
	return ans
}

func Min(arr []int) int {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
