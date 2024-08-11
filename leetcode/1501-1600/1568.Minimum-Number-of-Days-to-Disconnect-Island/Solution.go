package Solution

var dirs1568 = [][]int{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
}

func bfs1568(x, y int, grid [][]int) {
	q := [][2]int{{x, y}}
	grid[x][y] = 0
	for len(q) > 0 {
		nq := make([][2]int, 0)
		for _, cur := range q {
			for _, dir := range dirs1568 {
				nx, ny := cur[0]+dir[0], cur[1]+dir[1]
				if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) || grid[nx][ny] == 0 {
					continue
				}
				grid[nx][ny] = 0
				nq = append(nq, [2]int{nx, ny})
			}
		}
		q = nq
	}
}

// 判断联通, 如果开始就有多个区域，直击返回0
func connected1568(grid [][]int) bool {
	rows, cols := len(grid), len(grid[0])
	tmp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		tmp[i] = make([]int, cols)
		copy(tmp[i], grid[i])
	}
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if tmp[i][j] == 1 {
				if count == 1 {
					return false
				}
				bfs1568(i, j, tmp)
				count++
			}
		}
	}
	// 有且只有一个陆地的时候才是联通的
	return count == 1
}

// 我的思路是首先通过bfs搜索有几个联通
func Solution(grid [][]int) int {
	// 如果去掉一个1的时候无法将联通区域变成多个，那就至少需要两个操作
	// 两个操作，就类似于两步将一个角的1陆地裁剪出去。
	if !connected1568(grid) {
		// 一步都不用做
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				//尝试更新
				grid[i][j] = 0
				if !connected1568(grid) {
					return 1
				}
				grid[i][j] = 1
			}
		}
	}
	return 2
}
