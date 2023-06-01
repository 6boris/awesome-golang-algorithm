package Solution

type point1091 struct {
	x, y, dis int
}

func Solution(grid [][]int) int {
	tx, ty := len(grid)-1, len(grid[0])-1
	if grid[tx][ty] == 1 || grid[0][0] == 1 {
		return -1
	}
	dirs := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	queue := []point1091{{0, 0, 0}}
	for len(queue) > 0 {
		next := make([]point1091, 0)
		for _, item := range queue {
			if item.x == tx && item.y == ty {
				return item.dis + 1
			}
			for _, dir := range dirs {
				nx, ny := item.x+dir[0], item.y+dir[1]
				if nx < 0 || nx > tx || ny < 0 || ny > ty || grid[nx][ny] == 1 {
					continue
				}
				grid[nx][ny] = 1
				next = append(next, point1091{nx, ny, item.dis + 1})
			}
			grid[item.x][item.y] = 1
		}
		queue = next
	}
	return -1
}
