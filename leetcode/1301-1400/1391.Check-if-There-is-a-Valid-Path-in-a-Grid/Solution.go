package Solution

var dirFn = [6]func(int, int, [][]int) [][2]int{
	func(x, y int, grid [][]int) [][2]int {
		m, n := len(grid), len(grid[0])
		res := make([][2]int, 0)

		nx, ny := x, y-1
		if !(nx < 0 || nx >= m || ny < 0 || ny >= n) && (grid[nx][ny] == 1 || grid[nx][ny] == 4 || grid[nx][ny] == 6) {
			res = append(res, [2]int{nx, ny})
		}
		nx, ny = x, y+1
		if !(nx < 0 || nx >= m || ny < 0 || ny >= n) && (grid[nx][ny] == 1 || grid[nx][ny] == 3 || grid[nx][ny] == 5) {
			res = append(res, [2]int{nx, ny})
		}
		return res
	},
	func(x, y int, grid [][]int) [][2]int {
		m, n := len(grid), len(grid[0])
		res := make([][2]int, 0)

		nx, ny := x-1, y
		if !(nx < 0 || nx >= m || ny < 0 || ny >= n) && (grid[nx][ny] == 2 || grid[nx][ny] == 3 || grid[nx][ny] == 4) {
			res = append(res, [2]int{nx, ny})
		}
		nx, ny = x+1, y
		if !(nx < 0 || nx >= m || ny < 0 || ny >= n) && (grid[nx][ny] == 2 || grid[nx][ny] == 5 || grid[nx][ny] == 6) {
			res = append(res, [2]int{nx, ny})
		}
		return res
	},
	func(x, y int, grid [][]int) [][2]int {
		m, n := len(grid), len(grid[0])
		res := make([][2]int, 0)

		nx, ny := x, y-1
		if !(nx < 0 || nx >= m || ny < 0 || ny >= n) && (grid[nx][ny] == 1 || grid[nx][ny] == 4 || grid[nx][ny] == 6) {
			res = append(res, [2]int{nx, ny})
		}
		nx, ny = x+1, y
		if !(nx < 0 || nx >= m || ny < 0 || ny >= n) && (grid[nx][ny] == 2 || grid[nx][ny] == 5 || grid[nx][ny] == 6) {
			res = append(res, [2]int{nx, ny})
		}
		return res
	},
	func(x, y int, grid [][]int) [][2]int {
		m, n := len(grid), len(grid[0])
		res := make([][2]int, 0)

		nx, ny := x, y+1
		if !(nx < 0 || nx >= m || ny < 0 || ny >= n) && (grid[nx][ny] == 1 || grid[nx][ny] == 3 || grid[nx][ny] == 5) {
			res = append(res, [2]int{nx, ny})
		}
		nx, ny = x+1, y
		if !(nx < 0 || nx >= m || ny < 0 || ny >= n) && (grid[nx][ny] == 2 || grid[nx][ny] == 5 || grid[nx][ny] == 6) {
			res = append(res, [2]int{nx, ny})
		}
		return res
	},
	func(x, y int, grid [][]int) [][2]int {
		m, n := len(grid), len(grid[0])
		res := make([][2]int, 0)

		nx, ny := x, y-1
		if !(nx < 0 || nx >= m || ny < 0 || ny >= n) && (grid[nx][ny] == 1 || grid[nx][ny] == 4 || grid[nx][ny] == 6) {
			res = append(res, [2]int{nx, ny})
		}
		nx, ny = x-1, y
		if !(nx < 0 || nx >= m || ny < 0 || ny >= n) && (grid[nx][ny] == 2 || grid[nx][ny] == 3 || grid[nx][ny] == 4) {
			res = append(res, [2]int{nx, ny})
		}
		return res
	},
	func(x, y int, grid [][]int) [][2]int {
		m, n := len(grid), len(grid[0])
		res := make([][2]int, 0)

		nx, ny := x, y+1
		if !(nx < 0 || nx >= m || ny < 0 || ny >= n) && (grid[nx][ny] == 1 || grid[nx][ny] == 3 || grid[nx][ny] == 5) {
			res = append(res, [2]int{nx, ny})
		}
		nx, ny = x-1, y
		if !(nx < 0 || nx >= m || ny < 0 || ny >= n) && (grid[nx][ny] == 2 || grid[nx][ny] == 3 || grid[nx][ny] == 4) {
			res = append(res, [2]int{nx, ny})
		}
		return res
	},
}

func Solution(grid [][]int) bool {
	m, n := len(grid)-1, len(grid[0])-1
	q := [][2]int{{0, 0}}
	v := make(map[[2]int]struct{})
	v[[2]int{0, 0}] = struct{}{}
	for len(q) > 0 {
		nq := make([][2]int, 0)
		for _, pos := range q {
			x, y := pos[0], pos[1]
			if x == m && y == n {
				return true
			}
			for _, next := range dirFn[grid[x][y]-1](x, y, grid) {
				if _, ok := v[next]; ok {
					continue
				}
				nq = append(nq, next)
				v[next] = struct{}{}
			}
		}
		q = nq
	}
	return false
}
