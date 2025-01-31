package Solution

func Solution(grid [][]int) int {
	group := 1
	groupMapCount := map[int]int{}
	rows, cols := len(grid), len(grid[0])
	groupGrid := make([][]int, rows)
	for i := range rows {
		groupGrid[i] = make([]int, cols)
	}

	var (
		bfs  func(int, int) int
		dirs = [][2]int{
			{0, 1}, {1, 0}, {0, -1}, {-1, 0},
		}
	)

	bfs = func(x, y int) int {
		q := [][2]int{{x, y}}
		cnt := 0
		groupGrid[x][y] = group
		for len(q) > 0 {
			cnt += len(q)
			nq := make([][2]int, 0)
			for _, cur := range q {
				for _, d := range dirs {
					nx, ny := cur[0]+d[0], cur[1]+d[1]
					if nx < 0 || nx >= rows || ny < 0 || ny >= cols {
						continue
					}
					if grid[nx][ny] == 1 && groupGrid[nx][ny] == 0 {
						groupGrid[nx][ny] = group
						nq = append(nq, [2]int{nx, ny})
					}
				}
			}
			q = nq
		}
		return cnt
	}

	ans := 0
	for i := range rows {
		for j := range cols {
			if grid[i][j] == 1 && groupGrid[i][j] == 0 {
				groupMapCount[group] = bfs(i, j)
				ans = max(ans, groupMapCount[group])
				group++
			}
		}
	}

	for i := range rows {
		for j := range cols {
			if grid[i][j] == 0 {
				tmp := 1
				used := map[int]struct{}{}
				for _, d := range dirs {
					nx, ny := i+d[0], j+d[1]
					if nx < 0 || nx >= rows || ny < 0 || ny >= cols {
						continue
					}
					if groupGrid[nx][ny] != 0 {
						if _, ok := used[groupGrid[nx][ny]]; !ok {
							tmp += groupMapCount[groupGrid[nx][ny]]
							used[groupGrid[nx][ny]] = struct{}{}
						}
					}
				}
				ans = max(ans, tmp)
			}
		}
	}

	return ans
}
