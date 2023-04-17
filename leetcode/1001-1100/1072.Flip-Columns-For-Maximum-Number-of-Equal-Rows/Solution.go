package Solution

func Solution(matrix [][]int) int {
	rowSum := make([]int, len(matrix))
	ans := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			rowSum[row] += matrix[row][col]
		}
		if rowSum[row] == len(matrix[row]) || rowSum[row] == 0 {
			ans++
		}
	}

	for row := 0; row < len(matrix); row++ {
		for target := 0; target < 2; target++ {
			n := 1
			for inRow := 0; inRow < len(matrix); inRow++ {
				if inRow == row {
					continue
				}
				sum := rowSum[inRow]
				for col := 0; col < len(matrix[row]); col++ {
					if matrix[row][col] == target {
						continue
					}
					if matrix[inRow][col] == 1 {
						sum -= 1
					} else {
						sum += 1
					}
				}
				if sum == 0 || sum == len(matrix[inRow]) {
					n++
				}
			}
			if n > ans {
				ans = n
			}
		}
	}
	return ans
}
