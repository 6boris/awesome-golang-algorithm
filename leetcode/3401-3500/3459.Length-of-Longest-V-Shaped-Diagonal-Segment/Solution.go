package Solution

func Solution(grid [][]int) int {
	dirs := [4][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
	m, n := len(grid), len(grid[0])
	memo := make([][][4][2]int, m)
	for i := range memo {
		memo[i] = make([][4][2]int, n)
		for j := range memo[i] {
			for k := range memo[i][j] {
				for l := range memo[i][j][k] {
					memo[i][j][k][l] = -1
				}
			}
		}
	}

	var dfs func(cx, cy, direction int, turn bool, target int) int
	dfs = func(cx, cy, direction int, turn bool, target int) int {
		nx, ny := cx+dirs[direction][0], cy+dirs[direction][1]
		/* If it goes beyond the boundary or the next node's value is not the target value, then return */
		if nx < 0 || ny < 0 || nx >= m || ny >= n || grid[nx][ny] != target {
			return 0
		}

		turnInt := 0
		if turn {
			turnInt = 1
		}
		if memo[nx][ny][direction][turnInt] != -1 {
			return memo[nx][ny][direction][turnInt]
		}

		/* Continue walking in the original direction. */
		maxStep := dfs(nx, ny, direction, turn, 2-target)
		if turn {
			/* Clockwise rotate 90 degrees turn */
			maxStep = max(maxStep, dfs(nx, ny, (direction+1)%4, false, 2-target))
		}
		memo[nx][ny][direction][turnInt] = maxStep + 1
		return maxStep + 1
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				for direction := 0; direction < 4; direction++ {
					res = max(res, dfs(i, j, direction, true, 2)+1)
				}
			}
		}
	}
	return res
}
