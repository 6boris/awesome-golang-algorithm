package Solution

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}

	n := len(matrix[0])
	if n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		dp[0][j] = int(matrix[0][j] - '0')

		for i := 1; i < m; i++ {
			if matrix[i][j] == '1' {
				dp[i][j] = dp[i-1][j] + 1
			}
		}
	}

	max := 0
	for i := 0; i < m; i++ {
		area := largestRectangleArea(dp[i])
		if area > max {
			max = area
		}
	}

	return max
}

func largestRectangleArea(heights []int) int {
	heights = append(heights, -1)
	hLen := len(heights)

	left, right := 0, 0
	area, maxArea := 0, 0

	stack := []int{}

	for right < hLen {
		if len(stack) == 0 || heights[stack[len(stack)-1]] <= heights[right] {
			stack = append(stack, right)
			right++
			continue
		}

		height := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]

		if len(stack) == 0 {
			left = -1
		} else {
			left = stack[len(stack)-1]
		}

		area = height * (right - left - 1)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
