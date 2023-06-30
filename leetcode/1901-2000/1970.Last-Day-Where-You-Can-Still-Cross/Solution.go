package Solution

func Solution(row int, col int, cells [][]int) int {

	var (
		bfs  func(int) bool
		dirs = [][]int{
			// 开始想只向下搜索，缺少-1,0这个方向，结果过了115个例子
			{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		}
	)
	bfs = func(day int) bool {
		matrix := make([][]int, row)
		for i := 0; i < row; i++ {
			matrix[i] = make([]int, col)
		}
		for _, cell := range cells[:day] {
			x, y := cell[0]-1, cell[1]-1
			matrix[x][y] = 1
		}

		queue := [][2]int{}
		for i := 0; i < col; i++ {
			if matrix[0][i] == 0 {
				queue = append(queue, [2]int{0, i})
				matrix[0][i] = -1
			}
		}

		for len(queue) > 0 {
			next := make([][2]int, 0)
			for _, item := range queue {
				if item[0] == row-1 {
					return true
				}
				for _, dir := range dirs {
					nx, ny := item[0]+dir[0], item[1]+dir[1]
					if nx >= 0 && nx < row && ny >= 0 && ny < col && matrix[nx][ny] == 0 {
						next = append(next, [2]int{nx, ny})
						matrix[nx][ny] = -1
					}
				}
			}
			queue = next
		}
		return false

	}
	//
	left, right := 1, len(cells)
	for left < right {
		//mid := (right-left)/2 + left
		mid := right - (right-left)/2
		if bfs(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
