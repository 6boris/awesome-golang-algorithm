package Solution

func Solution(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, 2)
	dp[0], dp[1] = make([]int, cols), make([]int, cols)
	visited := make([]bool, cols)
	loop := 1
	for col := 0; col < cols; col++ {
		dp[0][col] = matrix[0][col]
	}

	for row := 1; row < rows; row++ {
		for col := 0; col < cols; col++ {
			down := dp[1-loop][col] + matrix[row][col]
			if !visited[col] || dp[loop][col] > down {
				dp[loop][col] = down
				visited[col] = true
			}
			if col < cols-1 {
				downRight := dp[1-loop][col] + matrix[row][col+1]
				if !visited[col+1] || dp[loop][col+1] > downRight {
					dp[loop][col+1] = downRight
					visited[col+1] = true
				}
			}
			if col > 0 {
				downLeft := dp[1-loop][col] + matrix[row][col-1]
				if !visited[col-1] || dp[loop][col-1] > downLeft {
					dp[loop][col-1] = downLeft
					visited[col-1] = true
				}
			}
		}

		for col := 0; col < cols; col++ {
			visited[col] = false
		}
		loop = 1 - loop
	}
	ans := dp[1-loop][0]
	for i := 1; i < cols; i++ {
		if dp[1-loop][i] < ans {
			ans = dp[1-loop][i]
		}
	}
	return ans
}
