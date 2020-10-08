package Solution

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
