package Solution

func spiralOrder(matrix [][]int) []int {
	var (
		rows, columns            = len(matrix), len(matrix[0])
		order                    = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return order
	}
	for left <= right && top <= bottom {
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}
