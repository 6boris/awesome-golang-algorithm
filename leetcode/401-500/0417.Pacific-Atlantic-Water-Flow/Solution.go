package Solution

func Solution(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	state := make([][]int, m)
	pacific, atlantic := 1, 2
	for i := 0; i < m; i++ {
		state[i] = make([]int, n)
	}
	for i := 0; i < n-1; i++ {
		state[0][i] = pacific
		state[m-1][i+1] = atlantic
	}
	for i := 0; i < m-1; i++ {
		state[i][0] = pacific
		state[i+1][n-1] = atlantic
	}

	state[0][n-1] = 3
	state[m-1][0] = 3
	var (
		leftTop     func(int, int, int)
		rightBottom func(int, int, int)
	)
	leftTop = func(x, y, pre int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if pre != -1 && state[x][y]&pacific == pacific {
			return
		}
		if heights[x][y] < pre {
			return
		}
		state[x][y] |= pacific
		leftTop(x-1, y, heights[x][y])
		leftTop(x+1, y, heights[x][y])
		leftTop(x, y+1, heights[x][y])
		leftTop(x, y-1, heights[x][y])
	}
	rightBottom = func(x, y, pre int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if heights[x][y] < pre {
			return
		}
		if pre != -1 && state[x][y]&atlantic == atlantic {
			return
		}
		state[x][y] |= atlantic
		rightBottom(x-1, y, heights[x][y])
		rightBottom(x+1, y, heights[x][y])
		rightBottom(x, y-1, heights[x][y])
		rightBottom(x, y+1, heights[x][y])
	}

	for i := 0; i < n; i++ {
		leftTop(0, i, -1)
		rightBottom(m-1, i, -1)
	}
	for i := 0; i < m; i++ {
		leftTop(i, 0, -1)
		rightBottom(i, n-1, -1)
	}
	res := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if state[i][j] == 3 {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
