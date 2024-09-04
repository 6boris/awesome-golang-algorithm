package Solution

import "sort"

func nextDir(cur byte, op int) byte {
	if cur == 'n' {
		if op == -1 {
			return 'e'
		}
		return 'w'
	}
	if cur == 'w' {
		if op == -1 {
			return 'n'
		}
		return 's'
	}
	if cur == 's' {
		if op == -1 {
			return 'w'
		}
		return 'e'
	}
	if op == -1 {
		return 's'
	}
	return 'n'
}

func nextSteps(cur byte) [2]int {
	if cur == 'n' {
		return [2]int{0, 1}
	}
	if cur == 'w' {
		return [2]int{-1, 0}
	}
	if cur == 's' {
		return [2]int{0, -1}
	}
	return [2]int{1, 0}
}

func Solution(commands []int, obstacles [][]int) int {
	x, y := 0, 0
	ans := 0
	snKeys := make([]int, 0)
	ewKeys := make([]int, 0)
	snIn := make(map[int][]int)
	ewIn := make(map[int][]int)

	for _, o := range obstacles {
		if _, ok := snIn[o[0]]; !ok {
			snIn[o[0]] = []int{}
			snKeys = append(snKeys, o[0])
		}
		snIn[o[0]] = append(snIn[o[0]], o[1])

		if _, ok := ewIn[o[1]]; !ok {
			ewIn[o[1]] = []int{}
			ewKeys = append(ewKeys, o[1])
		}
		ewIn[o[1]] = append(ewIn[o[1]], o[0])
	}
	sort.Ints(snKeys)
	sort.Ints(ewKeys)

	for k := range snIn {
		sort.Ints(snIn[k])
	}
	for k := range ewIn {
		sort.Ints(ewIn[k])
	}

	var check func(x, y, tx, ty int, cur byte) (int, int)
	check = func(x, y, tx, ty int, cur byte) (int, int) {
		if cur == 'n' || cur == 's' {
			v, ok := snIn[x]
			if !ok {
				return tx, ty
			}
			idx := sort.Search(len(v), func(i int) bool {
				return v[i] > y
			})
			if cur == 'n' {
				if idx == len(v) || v[idx] > ty || v[idx] == y {
					return tx, ty
				}
				return x, v[idx] - 1
			}
			if idx == 0 || v[idx-1] < ty {
				return tx, ty
			}
			return x, v[idx-1] + 1
		}
		if cur == 'w' || cur == 'e' {
			v, ok := ewIn[y]
			if !ok {
				return tx, ty
			}
			idx := sort.Search(len(v), func(i int) bool {
				return v[i] > x
			})
			if cur == 'w' {
				if idx == 0 || v[idx-1] < tx || v[idx-1] == x {
					return tx, ty
				}
				return v[idx-1] + 1, y
			}
			if idx == len(v) || v[idx] > tx {
				return tx, ty
			}
			return v[idx] - 1, y
		}
		return tx, ty
	}
	dir := byte('n')
	cur := nextSteps(dir)
	for _, cmd := range commands {
		if cmd > 0 {
			nx, ny := x+cur[0]*cmd, y+cur[1]*cmd
			nx, ny = check(x, y, nx, ny, dir)
			ans = max(ans, nx*nx+ny*ny)
			x, y = nx, ny
			continue
		}
		dir = nextDir(dir, cmd)
		cur = nextSteps(dir)
	}
	return ans
}
