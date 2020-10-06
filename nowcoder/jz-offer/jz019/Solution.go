package Solution

func printMatrix(matrix [][]int) []int {
	ans := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ans
	}

	low, high, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for low <= high && left <= right {
		//	向右
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[low][i])
		}
		//	向下
		for i := low + 1; i <= high; i++ {
			ans = append(ans, matrix[i][right])
		}
		//	向左
		if low > 1 {
			for i := right - 1; i >= left; i-- {
				ans = append(ans, matrix[high][i])
			}
		}
		//	向上
		if high > 1 {
			for i := high - 1; i >= low+1; i-- {
				ans = append(ans, matrix[i][left])
			}
		}
		low, high, left, right = low+1, high-1, left+1, right-1

	}
	return ans
}
