package Solution

func distance(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a + b
}

func Solution(ghosts [][]int, target []int) bool {
	a, b := target[0], target[1]
	cmp := distance(a, b)
	for _, g := range ghosts {
		a, b = target[0]-g[0], target[1]-g[1]
		tmp := distance(a, b)
		if tmp <= cmp {
			return false
		}
	}
	return true
}
