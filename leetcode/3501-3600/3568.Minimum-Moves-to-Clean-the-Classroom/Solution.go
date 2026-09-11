package Solution

func Solution(classroom []string, energy int) int {
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	m := len(classroom)
	n := len(classroom[0])
	id := make([][]int, m)
	for i := 0; i < m; i++ {
		id[i] = make([]int, n)
	}
	var sx, sy int
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if classroom[i][j] == 'S' {
				sx = i
				sy = j
			} else if classroom[i][j] == 'L' {
				id[i][j] = 1 << cnt
				cnt++
			}
		}
	}

	full := 1 << cnt
	bestEnergy := make([][][]int, m)
	for i := 0; i < m; i++ {
		bestEnergy[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			bestEnergy[i][j] = make([]int, full)
			for k := 0; k < full; k++ {
				bestEnergy[i][j][k] = -1
			}
		}
	}
	bestEnergy[sx][sy][0] = energy

	type Info struct {
		x, y, mask, e, steps int
	}
	q := make([]Info, 0)
	q = append(q, Info{sx, sy, 0, energy, 0})
	head := 0
	for head < len(q) {
		t := q[head]
		head++
		if t.mask == full-1 {
			return t.steps
		}
		if t.e == 0 {
			continue
		}
		for d := 0; d < 4; d++ {
			nx := t.x + dx[d]
			ny := t.y + dy[d]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || classroom[nx][ny] == 'X' {
				continue
			}
			ne := t.e - 1
			if classroom[nx][ny] == 'R' {
				ne = energy
			}
			nmask := t.mask | id[nx][ny]
			if ne > bestEnergy[nx][ny][nmask] {
				bestEnergy[nx][ny][nmask] = ne
				q = append(q, Info{nx, ny, nmask, ne, t.steps + 1})
			}
		}
	}
	return -1
}
