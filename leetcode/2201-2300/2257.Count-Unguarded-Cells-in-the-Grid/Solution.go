package Solution

func Solution(m int, n int, guards [][]int, walls [][]int) int {
	count := m * n
	count -= len(guards)
	count -= len(walls)
	wm := make(map[int]map[int]struct{})
	for _, w := range walls {
		if _, ok := wm[w[0]]; !ok {
			wm[w[0]] = make(map[int]struct{})
		}
		wm[w[0]][w[1]] = struct{}{}
	}
	gm := make(map[int]map[int]struct{})
	for _, g := range guards {
		if _, ok := gm[g[0]]; !ok {
			gm[g[0]] = make(map[int]struct{})
		}
		gm[g[0]][g[1]] = struct{}{}
	}

	used := make(map[[2]int]struct{})
	for _, g := range guards {
		x, y := g[0], g[1]
		for i := x - 1; i >= 0; i-- {
			if v, ok := gm[i]; ok {
				if _, ok := v[y]; ok {
					break
				}
			}
			if v, ok := wm[i]; ok {
				if _, ok := v[y]; ok {
					break
				}
			}
			used[[2]int{i, y}] = struct{}{}
		}

		for i := x + 1; i < m; i++ {
			if v, ok := gm[i]; ok {
				if _, ok := v[y]; ok {
					break
				}
			}
			if v, ok := wm[i]; ok {
				if _, ok := v[y]; ok {
					break
				}
			}
			used[[2]int{i, y}] = struct{}{}
		}

		rows := gm[x]
		for i := y + 1; i < n; i++ {
			if _, ok := rows[i]; ok {
				break
			}
			if v, ok := wm[x]; ok {
				if _, ok := v[i]; ok {
					break
				}
			}
			used[[2]int{x, i}] = struct{}{}
		}
		for i := y - 1; i >= 0; i-- {
			if _, ok := rows[i]; ok {
				break
			}
			if v, ok := wm[x]; ok {
				if _, ok := v[i]; ok {
					break
				}
			}
			used[[2]int{x, i}] = struct{}{}
		}
	}
	return count - len(used)
}
