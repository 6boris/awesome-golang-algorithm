package Solution

var dirs1034 = [][2]int{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
}

func Solution(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	initColor := grid[row][col]

	queue := [][2]int{{row, col}}
	v := map[[2]int]struct{}{
		[2]int{row, col}: struct{}{},
	}

	pos := make([][2]int, 0)
	for len(queue) > 0 {
		nq := make([][2]int, 0)
		for _, i := range queue {
			for _, d := range dirs1034 {
				nx, ny := i[0]+d[0], i[1]+d[1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] != initColor {
					pos = append(pos, i)
					continue
				}
				key := [2]int{nx, ny}
				if _, ok := v[key]; !ok {
					v[key] = struct{}{}
					nq = append(nq, key)
				}
			}
		}
		queue = nq
	}
	for _, p := range pos {
		grid[p[0]][p[1]] = color
	}
	return grid
}
