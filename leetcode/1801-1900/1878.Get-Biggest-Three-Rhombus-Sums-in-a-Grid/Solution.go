package Solution

import "sort"

func Solution(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	// 从左向右
	left := make([][]int, m+2)
	// 从右向左
	right := make([][]int, m+2)
	for i := range m + 2 {
		left[i] = make([]int, n+2)
		right[i] = make([]int, n+2)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			left[i][j] = left[i-1][j-1] + grid[i-1][j-1]
			right[i][j] = right[i-1][j+1] + grid[i-1][j-1]
		}
	}

	maxEdge := min(m, n)
	set := make(map[int]struct{})

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			set[grid[i][j]] = struct{}{}
			for edge := 1; edge <= maxEdge; edge++ {
				if i+2*edge >= m || j+edge >= n || j-edge < 0 {
					// 出界了
					break
				}
				// 左 -> 右
				rtEdge := left[i+edge+1][j+edge+1] - left[i][j]
				lbEdge := left[i+2*edge+1][j+1] - left[i+edge][j-edge]
				// 右 -> 左
				// right[i+1-1][j+1+1]
				ltEdge := right[i+edge+1][j-edge+1] - right[i][j+2]
				rbEdge := right[i+2*edge+1][j+1] - right[i+edge][j+edge+2]
				res = rtEdge + lbEdge + ltEdge + rbEdge - grid[i][j] - grid[i+edge][j-edge] - grid[i+2*edge][j] - grid[i+edge][j+edge]
				set[res] = struct{}{}
			}
		}
	}

	filter := make([]int, 0)
	for k := range set {
		filter = append(filter, k)
	}
	cut := min(3, len(filter))
	sort.Slice(filter, func(i, j int) bool {
		return filter[i] > filter[j]
	})
	return filter[:cut]
}
