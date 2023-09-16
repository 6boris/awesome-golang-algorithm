package Solution

func Solution(heights [][]int) int {
	// binary search + bfs
	var (
		left  = 0
		right = 1000000
		bfs   func(int, int, int) bool
		dirs  = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	)

	rows, cols := len(heights), len(heights[0])
	bfs = func(x, y, limit int) bool {
		visited := [100][100]bool{}
		queue := [][]int{{x, y}}
		visited[x][y] = true
		for len(queue) > 0 {
			next := make([][]int, 0)
			for _, item := range queue {
				if item[0] == rows-1 && item[1] == cols-1 {
					return true
				}
				for _, dir := range dirs {
					nx, ny := item[0]+dir[0], item[1]+dir[1]
					if nx < 0 || nx >= rows || ny < 0 || ny >= cols || visited[nx][ny] {
						continue
					}
					diff := heights[nx][ny] - heights[item[0]][item[1]]
					if diff < 0 {
						diff = ^diff + 1
					}
					if diff > limit {
						continue
					}
					next = append(next, []int{nx, ny})
					visited[nx][ny] = true
				}
			}
			queue = next
		}
		return false
	}
	for left < right {
		m := left + (right-left)/2
		if bfs(0, 0, m) {
			right = m
			continue
		}
		left = m + 1
	}

	return left
}
