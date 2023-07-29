package Solution

func Solution(heights []int) []int {
	// 单调栈
	// 1, 2, 3
	l := len(heights)
	ans := make([]int, l)
	ans[l-1] = 0
	stack := make([]int, l)
	stack[l-1] = heights[l-1]
	index := l - 1
	// a, b, c, d, e
	for idx := l - 2; idx >= 0; idx-- {
		if heights[idx] < stack[index] {
			index--
			stack[index] = heights[idx]
			ans[idx] = 1
			continue
		}

		for ; index < l && heights[idx] > stack[index]; index++ {
			ans[idx]++
		}
		if index != l {
			ans[idx]++
		}
		index--
		stack[index] = heights[idx]
	}
	// abc c是b的nextGreater，a可以看到
	// 单调增和减
	// 先尝试n^2
	/*
		   m := make([]int, l)
		   // 10, 8, 6, 7
			for idx := 0; idx < l-2; idx++ {
				// +1 是可以看到的
				ans[idx] = 1
				max := heights[idx+1]
				for next := idx + 2; next < l; next++ {
					if heights[idx] < heights[next] && heights[idx] > max ||
						heights[idx] > heights[next] && heights[next] > max {
						ans[idx]++
					}
					if heights[next] > max {
						max = heights[next]
					}
				}
			}
	*/
	return ans
}
